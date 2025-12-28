package impl1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestBuffer_Create_NilIndices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	expectedBuffer := mocks.NewFakeBufferHandle()

	mockLoader.EXPECT().VkCreateBuffer(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, createInfo *loader.VkBufferCreateInfo, allocator *loader.VkAllocationCallbacks, buffer *loader.VkBuffer) (common.VkResult, error) {
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

	buffer, res, err := driver.CreateBuffer(nil, core1_0.BufferCreateInfo{
		Size:               5,
		Usage:              core1_0.BufferUsageVertexBuffer | core1_0.BufferUsageTransferSrc,
		SharingMode:        core1_0.SharingModeExclusive,
		QueueFamilyIndices: []int{},
	})

	require.Equal(t, res, core1_0.VKSuccess)
	require.NoError(t, err)
	require.Equal(t, expectedBuffer, buffer.Handle())
}

func TestBasicBuffer_Create_QueueFamilyIndices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	expectedBuffer := mocks.NewFakeBufferHandle()
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateBuffer(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, createInfo *loader.VkBufferCreateInfo, allocator *loader.VkAllocationCallbacks, buffer *loader.VkBuffer) (common.VkResult, error) {
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
			indicesSlice := ([]loader.Uint32)(unsafe.Slice((*loader.Uint32)(indicesPtrUnsafe), 4))
			require.Equal(t, []loader.Uint32{1, 2, 3, 4}, indicesSlice)

			*buffer = expectedBuffer

			return core1_0.VKSuccess, nil
		})

	buffer, res, err := driver.CreateBuffer(nil, core1_0.BufferCreateInfo{
		Size:               5,
		Usage:              core1_0.BufferUsageVertexBuffer | core1_0.BufferUsageTransferSrc,
		SharingMode:        core1_0.SharingModeExclusive,
		QueueFamilyIndices: []int{1, 2, 3, 4},
	})

	require.Equal(t, res, core1_0.VKSuccess)
	require.NoError(t, err)
	require.Equal(t, expectedBuffer, buffer.Handle())

}

func TestBuffer_MemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	buffer := mocks.NewDummyBuffer(device)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkGetBufferMemoryRequirements(device.Handle(), buffer.Handle(), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, buffer loader.VkBuffer, requirements *loader.VkMemoryRequirements) {
			v := reflect.ValueOf(requirements).Elem()
			*(*uint64)(unsafe.Pointer(v.FieldByName("size").UnsafeAddr())) = 5
			*(*uint64)(unsafe.Pointer(v.FieldByName("alignment").UnsafeAddr())) = 8
			*(*uint32)(unsafe.Pointer(v.FieldByName("memoryTypeBits").UnsafeAddr())) = 0xff
		})

	reqs := driver.GetBufferMemoryRequirements(buffer)
	require.Equal(t, 5, reqs.Size)
	require.Equal(t, 8, reqs.Alignment)
	require.Equal(t, uint32(0xFF), reqs.MemoryTypeBits)
}

func TestBuffer_BindBufferMemory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	buffer := mocks.NewDummyBuffer(device)
	memory := mocks.NewDummyDeviceMemory(device, 1)

	mockLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkBindBufferMemory(device.Handle(), buffer.Handle(), memory.Handle(), loader.VkDeviceSize(3)).Return(core1_0.VKSuccess, nil)
	_, err := driver.BindBufferMemory(buffer, memory, 3)
	require.NoError(t, err)
}

func TestBuffer_BindBufferMemory_FailNilMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	buffer := mocks.NewDummyBuffer(device)

	mockLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	_, err := driver.BindBufferMemory(buffer, core.DeviceMemory{}, 3)
	require.EqualError(t, err, "received uninitialized DeviceMemory")
}
