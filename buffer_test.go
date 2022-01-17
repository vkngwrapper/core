package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestBuffer_Create_NilIndices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)

	expectedBuffer := mocks.NewFakeBufferHandle()

	mockDriver.EXPECT().VkCreateBuffer(mocks.Exactly(device.Handle()), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, createInfo *driver.VkBufferCreateInfo, allocator *driver.VkAllocationCallbacks, buffer *driver.VkBuffer) (common.VkResult, error) {
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

			return common.VKSuccess, nil
		})

	buffer, res, err := loader.CreateBuffer(device, nil, &core.BufferOptions{
		BufferSize:         5,
		Usage:              common.UsageVertexBuffer | common.UsageTransferSrc,
		SharingMode:        common.SharingExclusive,
		QueueFamilyIndices: []int{},
	})

	require.Equal(t, res, common.VKSuccess)
	require.NoError(t, err)
	require.Same(t, expectedBuffer, buffer.Handle())
}

func TestBasicBuffer_Create_QueueFamilyIndices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)

	expectedBuffer := mocks.NewFakeBufferHandle()

	mockDriver.EXPECT().VkCreateBuffer(mocks.Exactly(device.Handle()), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, createInfo *driver.VkBufferCreateInfo, allocator *driver.VkAllocationCallbacks, buffer *driver.VkBuffer) (common.VkResult, error) {
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
			indicesSlice := ([]driver.Uint32)(unsafe.Slice((*driver.Uint32)(indicesPtrUnsafe), 4))
			require.Equal(t, []driver.Uint32{1, 2, 3, 4}, indicesSlice)

			*buffer = expectedBuffer

			return common.VKSuccess, nil
		})

	buffer, res, err := loader.CreateBuffer(device, nil, &core.BufferOptions{
		BufferSize:         5,
		Usage:              common.UsageVertexBuffer | common.UsageTransferSrc,
		SharingMode:        common.SharingExclusive,
		QueueFamilyIndices: []int{1, 2, 3, 4},
	})

	require.Equal(t, res, common.VKSuccess)
	require.NoError(t, err)
	require.Same(t, expectedBuffer, buffer.Handle())

}

func TestBuffer_MemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	buffer := mocks.EasyDummyBuffer(t, loader, device)

	mockDriver.EXPECT().VkGetBufferMemoryRequirements(mocks.Exactly(device.Handle()), mocks.Exactly(buffer.Handle()), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, buffer driver.VkBuffer, requirements *driver.VkMemoryRequirements) {
			v := reflect.ValueOf(requirements).Elem()
			*(*uint64)(unsafe.Pointer(v.FieldByName("size").UnsafeAddr())) = 5
			*(*uint64)(unsafe.Pointer(v.FieldByName("alignment").UnsafeAddr())) = 8
			*(*uint32)(unsafe.Pointer(v.FieldByName("memoryTypeBits").UnsafeAddr())) = 0xff
		})

	reqs := buffer.MemoryRequirements()
	require.Equal(t, 5, reqs.Size)
	require.Equal(t, 8, reqs.Alignment)
	require.Equal(t, uint32(0xFF), reqs.MemoryType)
	require.NoError(t, err)
}

func TestBuffer_BindBufferMemory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	buffer := mocks.EasyDummyBuffer(t, loader, device)
	memory := mocks.EasyMockDeviceMemory(ctrl)

	mockDriver.EXPECT().VkBindBufferMemory(mocks.Exactly(device.Handle()), mocks.Exactly(buffer.Handle()), mocks.Exactly(memory.Handle()), driver.VkDeviceSize(3)).Return(common.VKSuccess, nil)
	_, err = buffer.BindBufferMemory(memory, 3)
	require.NoError(t, err)
}

func TestBuffer_BindBufferMemory_FailNilMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	buffer := mocks.EasyDummyBuffer(t, loader, device)

	_, err = buffer.BindBufferMemory(nil, 3)
	require.EqualError(t, err, "received nil DeviceMemory")
}
