package core1_2_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_2"
	"go.uber.org/mock/gomock"
)

func TestBufferOpaqueCaptureAddressCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	mockBuffer := mocks.NewDummyBuffer(device)

	coreLoader.EXPECT().VkCreateBuffer(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pCreateInfo *loader.VkBufferCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pBuffer *loader.VkBuffer) (common.VkResult, error) {

		*pBuffer = mockBuffer.Handle()
		val := reflect.ValueOf(pCreateInfo).Elem()

		require.Equal(t, uint64(12), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO

		next := (*loader.VkBufferOpaqueCaptureAddressCreateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000257002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_OPAQUE_CAPTURE_ADDRESS_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(13), val.FieldByName("opaqueCaptureAddress").Uint())

		return core1_0.VKSuccess, nil
	})

	buffer, _, err := driver.CreateBuffer(
		device,
		nil,
		core1_0.BufferCreateInfo{
			NextOptions: common.NextOptions{
				core1_2.BufferOpaqueCaptureAddressCreateInfo{
					OpaqueCaptureAddress: 13,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockBuffer.Handle(), buffer.Handle())
}

func TestMemoryOpaqueCaptureAddressAllocateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_2.InternalDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	mockMemory := mocks.NewDummyDeviceMemory(device, 1)

	coreLoader.EXPECT().VkAllocateMemory(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pAllocateInfo *loader.VkMemoryAllocateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pMemory *loader.VkDeviceMemory) (common.VkResult, error) {

		*pMemory = mockMemory.Handle()
		val := reflect.ValueOf(pAllocateInfo).Elem()

		require.Equal(t, uint64(5), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO

		next := (*loader.VkMemoryOpaqueCaptureAddressAllocateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000257003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_OPAQUE_CAPTURE_ADDRESS_ALLOCATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(17), val.FieldByName("opaqueCaptureAddress").Uint())

		return core1_0.VKSuccess, nil
	})

	memory, _, err := driver.AllocateMemory(
		device,
		nil,
		core1_0.MemoryAllocateInfo{
			NextOptions: common.NextOptions{
				core1_2.MemoryOpaqueCaptureAddressAllocateInfo{
					OpaqueCaptureAddress: 17,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockMemory.Handle(), memory.Handle())
}
