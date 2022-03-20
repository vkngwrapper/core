package core1_0_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	internal_mocks "github.com/CannibalVox/VKng/core/internal/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanDeviceMemory_MapMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := internal_mocks.EasyDummyDevice(t, ctrl, loader)

	memory := internal_mocks.EasyDummyDeviceMemory(t, loader, device, 1)
	memoryPtr := unsafe.Pointer(t)

	mockDriver.EXPECT().VkMapMemory(device.Handle(), memory.Handle(), driver.VkDeviceSize(1), driver.VkDeviceSize(3), driver.VkMemoryMapFlags(0), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, memory driver.VkDeviceMemory, offset driver.VkDeviceSize, size driver.VkDeviceSize, flags driver.VkMemoryMapFlags, ppData *unsafe.Pointer) (common.VkResult, error) {
			*ppData = memoryPtr

			return core1_0.VKSuccess, nil
		})

	ptr, _, err := memory.MapMemory(1, 3, 0)
	require.Equal(t, memoryPtr, ptr)
	require.NoError(t, err)
}

func TestVulkanDeviceMemory_UnmapMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := internal_mocks.EasyDummyDevice(t, ctrl, loader)
	memory := internal_mocks.EasyDummyDeviceMemory(t, loader, device, 1)

	mockDriver.EXPECT().VkUnmapMemory(device.Handle(), memory.Handle())

	memory.UnmapMemory()
}

func TestVulkanDeviceMemory_Commitment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := internal_mocks.EasyDummyDevice(t, ctrl, loader)
	memory := internal_mocks.EasyDummyDeviceMemory(t, loader, device, 1)

	mockDriver.EXPECT().VkGetDeviceMemoryCommitment(device.Handle(), memory.Handle(), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, memory driver.VkDeviceMemory, pCommitment *driver.VkDeviceSize) {
			*pCommitment = 3
		})

	require.Equal(t, 3, memory.Commitment())
}

func TestVulkanDeviceMemory_Flush(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := internal_mocks.EasyDummyDevice(t, ctrl, loader)
	memory := internal_mocks.EasyDummyDeviceMemory(t, loader, device, 113)

	mockDriver.EXPECT().VkFlushMappedMemoryRanges(device.Handle(), driver.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, memoryRangeCount driver.Uint32, pMemoryRanges *driver.VkMappedMemoryRange) (common.VkResult, error) {
			val := reflect.ValueOf(pMemoryRanges).Elem()

			require.Equal(t, uint64(6), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, memory.Handle(), (driver.VkDeviceMemory)(unsafe.Pointer(val.FieldByName("memory").Elem().UnsafeAddr())))
			require.Equal(t, uint64(0), val.FieldByName("offset").Uint())
			require.Equal(t, uint64(113), val.FieldByName("size").Uint())

			return core1_0.VKSuccess, nil
		})

	_, err = memory.FlushAll()
	require.NoError(t, err)
}

func TestVulkanDeviceMemory_Invalidate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := internal_mocks.EasyDummyDevice(t, ctrl, loader)
	memory := internal_mocks.EasyDummyDeviceMemory(t, loader, device, 113)

	mockDriver.EXPECT().VkInvalidateMappedMemoryRanges(device.Handle(), driver.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, memoryRangeCount driver.Uint32, pMemoryRanges *driver.VkMappedMemoryRange) (common.VkResult, error) {
			val := reflect.ValueOf(pMemoryRanges).Elem()

			require.Equal(t, uint64(6), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, memory.Handle(), (driver.VkDeviceMemory)(unsafe.Pointer(val.FieldByName("memory").Elem().UnsafeAddr())))
			require.Equal(t, uint64(0), val.FieldByName("offset").Uint())
			require.Equal(t, uint64(113), val.FieldByName("size").Uint())

			return core1_0.VKSuccess, nil
		})

	_, err = memory.InvalidateAll()
	require.NoError(t, err)
}

func TestVulkanDeviceMemory_AllocateAndFreeMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := internal_mocks.EasyDummyDevice(t, ctrl, loader)
	memoryHandle := mocks.NewFakeDeviceMemoryHandle()

	mockDriver.EXPECT().VkAllocateMemory(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkMemoryAllocateInfo, pAllocator *driver.VkAllocationCallbacks, pMemory *driver.VkDeviceMemory) (common.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(5), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(7), val.FieldByName("allocationSize").Uint())
			require.Equal(t, uint64(3), val.FieldByName("memoryTypeIndex").Uint())

			*pMemory = memoryHandle
			return core1_0.VKSuccess, nil
		})
	mockDriver.EXPECT().VkFreeMemory(device.Handle(), memoryHandle, nil)

	memory, _, err := loader.AllocateMemory(device, nil, core1_0.DeviceMemoryOptions{
		AllocationSize:  7,
		MemoryTypeIndex: 3,
	})
	require.NoError(t, err)
	require.NotNil(t, memory)
	require.Equal(t, memoryHandle, memory.Handle())

	memory.Free(nil)
}
