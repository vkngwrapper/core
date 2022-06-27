package core1_0_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	internal_mocks "github.com/CannibalVox/VKng/core/internal/dummies"
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

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := internal_mocks.EasyDummyDevice(mockDriver)

	expectedBuffer := mocks.NewFakeBufferHandle()

	mockDriver.EXPECT().VkCreateBuffer(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
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

			return core1_0.VKSuccess, nil
		})

	buffer, res, err := device.CreateBuffer(nil, core1_0.BufferCreateOptions{
		BufferSize:         5,
		Usage:              core1_0.BufferUsageVertexBuffer | core1_0.BufferUsageTransferSrc,
		SharingMode:        core1_0.SharingExclusive,
		QueueFamilyIndices: []int{},
	})

	require.Equal(t, res, core1_0.VKSuccess)
	require.NoError(t, err)
	require.Equal(t, expectedBuffer, buffer.Handle())
}

func TestBasicBuffer_Create_QueueFamilyIndices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := internal_mocks.EasyDummyDevice(mockDriver)

	expectedBuffer := mocks.NewFakeBufferHandle()

	mockDriver.EXPECT().VkCreateBuffer(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
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

			return core1_0.VKSuccess, nil
		})

	buffer, res, err := device.CreateBuffer(nil, core1_0.BufferCreateOptions{
		BufferSize:         5,
		Usage:              core1_0.BufferUsageVertexBuffer | core1_0.BufferUsageTransferSrc,
		SharingMode:        core1_0.SharingExclusive,
		QueueFamilyIndices: []int{1, 2, 3, 4},
	})

	require.Equal(t, res, core1_0.VKSuccess)
	require.NoError(t, err)
	require.Equal(t, expectedBuffer, buffer.Handle())

}

func TestBuffer_MemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	buffer := internal_mocks.EasyDummyBuffer(mockDriver, device)

	mockDriver.EXPECT().VkGetBufferMemoryRequirements(device.Handle(), buffer.Handle(), gomock.Not(nil)).DoAndReturn(
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
}

func TestBuffer_BindBufferMemory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	buffer := internal_mocks.EasyDummyBuffer(mockDriver, device)
	memory := mocks.EasyMockDeviceMemory(ctrl)

	mockDriver.EXPECT().VkBindBufferMemory(device.Handle(), buffer.Handle(), memory.Handle(), driver.VkDeviceSize(3)).Return(core1_0.VKSuccess, nil)
	_, err := buffer.BindBufferMemory(memory, 3)
	require.NoError(t, err)
}

func TestBuffer_BindBufferMemory_FailNilMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	buffer := internal_mocks.EasyDummyBuffer(mockDriver, device)

	_, err := buffer.BindBufferMemory(nil, 3)
	require.EqualError(t, err, "received nil DeviceMemory")
}
