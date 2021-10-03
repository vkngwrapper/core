package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestBufferCreateNilIndices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)

	expectedBuffer := mocks.NewFakeBufferHandle()

	mockDriver.EXPECT().VkCreateBuffer(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, createInfo *core.VkBufferCreateInfo, allocator *core.VkAllocationCallbacks, buffer *core.VkBuffer) (core.VkResult, error) {
			v := reflect.ValueOf(*createInfo)
			require.Equal(t, v.FieldByName("sType").Uint(), uint64(12)) //VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, v.FieldByName("flags").Uint(), uint64(0))
			require.Equal(t, v.FieldByName("size").Uint(), uint64(5))
			require.Equal(t, v.FieldByName("usage").Uint(), uint64(0x81)) //VK_BUFFER_USAGE_VERTEX_BUFFER_BIT|VK_BUFFER_USAGE_TRANSFER_SRC_BIT
			require.Equal(t, v.FieldByName("queueFamilyIndexCount").Uint(), uint64(0))

			indicesVal := v.FieldByName("pQueueFamilyIndices")
			require.True(t, indicesVal.IsNil())

			*buffer = expectedBuffer

			return core.VKSuccess, nil
		})

	buffer, res, err := loader.CreateBuffer(device, &core.BufferOptions{
		BufferSize:         5,
		Usages:             common.UsageVertexBuffer | common.UsageTransferSrc,
		SharingMode:        common.SharingExclusive,
		QueueFamilyIndices: []int{},
	})

	require.Equal(t, res, core.VKSuccess)
	require.NoError(t, err)
	require.Equal(t, expectedBuffer, buffer.Handle())

}

func TestBasicBufferCreateWithReqs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)

	expectedBuffer := mocks.NewFakeBufferHandle()

	mockDriver.EXPECT().VkCreateBuffer(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, createInfo *core.VkBufferCreateInfo, allocator *core.VkAllocationCallbacks, buffer *core.VkBuffer) (core.VkResult, error) {
			v := reflect.ValueOf(*createInfo)
			require.Equal(t, v.FieldByName("sType").Uint(), uint64(12)) //VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, v.FieldByName("flags").Uint(), uint64(0))
			require.Equal(t, v.FieldByName("size").Uint(), uint64(5))
			require.Equal(t, v.FieldByName("usage").Uint(), uint64(0x81)) //VK_BUFFER_USAGE_VERTEX_BUFFER_BIT|VK_BUFFER_USAGE_TRANSFER_SRC_BIT
			require.Equal(t, v.FieldByName("queueFamilyIndexCount").Uint(), uint64(4))

			indicesVal := v.FieldByName("pQueueFamilyIndices")
			require.False(t, indicesVal.IsNil())

			indicesPtrUnsafe := unsafe.Pointer(indicesVal.Elem().UnsafeAddr())
			indicesSlice := ([]core.Uint32)(unsafe.Slice((*core.Uint32)(indicesPtrUnsafe), 4))
			require.Equal(t, []core.Uint32{1, 2, 3, 4}, indicesSlice)

			*buffer = expectedBuffer

			return core.VKSuccess, nil
		})

	buffer, res, err := loader.CreateBuffer(device, &core.BufferOptions{
		BufferSize:         5,
		Usages:             common.UsageVertexBuffer | common.UsageTransferSrc,
		SharingMode:        common.SharingExclusive,
		QueueFamilyIndices: []int{1, 2, 3, 4},
	})

	require.Equal(t, res, core.VKSuccess)
	require.NoError(t, err)
	require.Equal(t, expectedBuffer, buffer.Handle())

	mockDriver.EXPECT().VkGetBufferMemoryRequirements(device.Handle(), expectedBuffer, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, buffer core.VkBuffer, requirements *core.VkMemoryRequirements) error {
			v := reflect.ValueOf(requirements).Elem()
			*(*uint64)(unsafe.Pointer(v.FieldByName("size").UnsafeAddr())) = 5
			*(*uint64)(unsafe.Pointer(v.FieldByName("alignment").UnsafeAddr())) = 8
			*(*uint32)(unsafe.Pointer(v.FieldByName("memoryTypeBits").UnsafeAddr())) = 0xff

			return nil
		})

	reqs, err := buffer.MemoryRequirements()
	require.Equal(t, 5, reqs.Size)
	require.Equal(t, 8, reqs.Alignment)
	require.Equal(t, uint32(0xFF), reqs.MemoryType)
	require.NoError(t, err)

}
