package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"unsafe"
)

func TestVulkanDeviceMemory_MapMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)

	memory := mocks.EasyDummyDeviceMemory(t, device)
	memoryPtr := unsafe.Pointer(t)

	mockDriver.EXPECT().VkMapMemory(device.Handle(), memory.Handle(), core.VkDeviceSize(1), core.VkDeviceSize(3), core.VkMemoryMapFlags(0), gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, memory core.VkDeviceMemory, offset core.VkDeviceSize, size core.VkDeviceSize, flags core.VkMemoryMapFlags, ppData *unsafe.Pointer) (core.VkResult, error) {
			*ppData = memoryPtr

			return core.VKSuccess, nil
		})

	ptr, _, err := memory.MapMemory(1, 3, 0)
	require.Equal(t, memoryPtr, ptr)
	require.NoError(t, err)
}

func TestVulkanDeviceMemory_UnmapMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	memory := mocks.EasyDummyDeviceMemory(t, device)

	mockDriver.EXPECT().VkUnmapMemory(device.Handle(), memory.Handle()).Return(nil)

	err = memory.UnmapMemory()
	require.NoError(t, err)
}

func TestVulkanDeviceMemory_WriteData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	memory := mocks.EasyDummyDeviceMemory(t, device)

	type anonType struct {
		A int32
		B int32
		C int32
	}
	inputData := &anonType{
		A: 1,
		B: 2,
		C: 3,
	}
	outputData := &anonType{}
	dataPtr := unsafe.Pointer(outputData)

	mockDriver.EXPECT().VkMapMemory(device.Handle(), memory.Handle(), core.VkDeviceSize(0), core.VkDeviceSize(12), core.VkMemoryMapFlags(0), gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, memory core.VkDeviceMemory, offset core.VkDeviceSize, size core.VkDeviceSize, flags core.VkMemoryMapFlags, ppData *unsafe.Pointer) (core.VkResult, error) {
			*ppData = dataPtr

			return core.VKSuccess, nil
		})
	mockDriver.EXPECT().VkUnmapMemory(device.Handle(), memory.Handle()).Return(nil)

	_, err = memory.WriteData(0, inputData)
	require.NoError(t, err)
	require.Equal(t, *inputData, *outputData)
}
