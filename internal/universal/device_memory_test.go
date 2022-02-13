package universal_test

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

func TestVulkanDeviceMemory_MapMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)

	memory := mocks.EasyDummyDeviceMemory(t, device, 1)
	memoryPtr := unsafe.Pointer(t)

	mockDriver.EXPECT().VkMapMemory(mocks.Exactly(device.Handle()), mocks.Exactly(memory.Handle()), driver.VkDeviceSize(1), driver.VkDeviceSize(3), driver.VkMemoryMapFlags(0), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, memory driver.VkDeviceMemory, offset driver.VkDeviceSize, size driver.VkDeviceSize, flags driver.VkMemoryMapFlags, ppData *unsafe.Pointer) (common.VkResult, error) {
			*ppData = memoryPtr

			return common.VKSuccess, nil
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
	memory := mocks.EasyDummyDeviceMemory(t, device, 1)

	mockDriver.EXPECT().VkUnmapMemory(mocks.Exactly(device.Handle()), mocks.Exactly(memory.Handle()))

	memory.UnmapMemory()
}

func TestVulkanDeviceMemory_Commitment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	memory := mocks.EasyDummyDeviceMemory(t, device, 1)

	mockDriver.EXPECT().VkGetDeviceMemoryCommitment(mocks.Exactly(device.Handle()), mocks.Exactly(memory.Handle()), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, memory driver.VkDeviceMemory, pCommitment *driver.VkDeviceSize) {
			*pCommitment = 3
		})

	require.Equal(t, 3, memory.Commitment())
}

func TestVulkanDeviceMemory_Flush(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	memory := mocks.EasyDummyDeviceMemory(t, device, 113)

	mockDriver.EXPECT().VkFlushMappedMemoryRanges(mocks.Exactly(device.Handle()), driver.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, memoryRangeCount driver.Uint32, pMemoryRanges *driver.VkMappedMemoryRange) (common.VkResult, error) {
			val := reflect.ValueOf(pMemoryRanges).Elem()

			require.Equal(t, uint64(6), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Same(t, memory.Handle(), (driver.VkDeviceMemory)(unsafe.Pointer(val.FieldByName("memory").Elem().UnsafeAddr())))
			require.Equal(t, uint64(0), val.FieldByName("offset").Uint())
			require.Equal(t, uint64(113), val.FieldByName("size").Uint())

			return common.VKSuccess, nil
		})

	_, err = memory.Flush()
	require.NoError(t, err)
}

func TestVulkanDeviceMemory_Invalidate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	memory := mocks.EasyDummyDeviceMemory(t, device, 113)

	mockDriver.EXPECT().VkInvalidateMappedMemoryRanges(mocks.Exactly(device.Handle()), driver.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, memoryRangeCount driver.Uint32, pMemoryRanges *driver.VkMappedMemoryRange) (common.VkResult, error) {
			val := reflect.ValueOf(pMemoryRanges).Elem()

			require.Equal(t, uint64(6), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Same(t, memory.Handle(), (driver.VkDeviceMemory)(unsafe.Pointer(val.FieldByName("memory").Elem().UnsafeAddr())))
			require.Equal(t, uint64(0), val.FieldByName("offset").Uint())
			require.Equal(t, uint64(113), val.FieldByName("size").Uint())

			return common.VKSuccess, nil
		})

	_, err = memory.Invalidate()
	require.NoError(t, err)
}

func TestVulkanDeviceMemory_AllocateAndFreeMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	memoryHandle := mocks.NewFakeDeviceMemoryHandle()

	mockDriver.EXPECT().VkAllocateMemory(mocks.Exactly(device.Handle()), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkMemoryAllocateInfo, pAllocator *driver.VkAllocationCallbacks, pMemory *driver.VkDeviceMemory) (common.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(5), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(7), val.FieldByName("allocationSize").Uint())
			require.Equal(t, uint64(3), val.FieldByName("memoryTypeIndex").Uint())

			*pMemory = memoryHandle
			return common.VKSuccess, nil
		})
	mockDriver.EXPECT().VkFreeMemory(mocks.Exactly(device.Handle()), mocks.Exactly(memoryHandle), nil)

	memory, _, err := device.AllocateMemory(nil, &core.DeviceMemoryOptions{
		AllocationSize:  7,
		MemoryTypeIndex: 3,
	})
	require.NoError(t, err)
	require.NotNil(t, memory)
	require.Same(t, memoryHandle, memory.Handle())

	memory.Free(nil)
}
