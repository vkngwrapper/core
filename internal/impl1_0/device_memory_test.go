package impl1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanDeviceMemory_MapMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	device := mocks.EasyMockDevice(ctrl, mockDriver)

	builder := &impl1_0.DeviceObjectBuilderImpl{}
	memory := builder.CreateDeviceMemoryObject(mockDriver, device.Handle(), mocks.NewFakeDeviceMemoryHandle(), common.Vulkan1_0, 1)
	memoryPtr := unsafe.Pointer(t)

	mockDriver.EXPECT().VkMapMemory(device.Handle(), memory.Handle(), driver.VkDeviceSize(1), driver.VkDeviceSize(3), driver.VkMemoryMapFlags(0), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, memory driver.VkDeviceMemory, offset driver.VkDeviceSize, size driver.VkDeviceSize, flags driver.VkMemoryMapFlags, ppData *unsafe.Pointer) (common.VkResult, error) {
			*ppData = memoryPtr

			return core1_0.VKSuccess, nil
		})

	ptr, _, err := memory.Map(1, 3, 0)
	require.Equal(t, memoryPtr, ptr)
	require.NoError(t, err)
}

func TestVulkanDeviceMemory_UnmapMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	memory := builder.CreateDeviceMemoryObject(mockDriver, device.Handle(), mocks.NewFakeDeviceMemoryHandle(), common.Vulkan1_0, 1)

	mockDriver.EXPECT().VkUnmapMemory(device.Handle(), memory.Handle())

	memory.Unmap()
}

func TestVulkanDeviceMemory_Commitment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	memory := builder.CreateDeviceMemoryObject(mockDriver, device.Handle(), mocks.NewFakeDeviceMemoryHandle(), common.Vulkan1_0, 1)

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

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	memory := builder.CreateDeviceMemoryObject(mockDriver, device.Handle(), mocks.NewFakeDeviceMemoryHandle(), common.Vulkan1_0, 113)

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

	_, err := memory.FlushAll()
	require.NoError(t, err)
}

func TestVulkanDeviceMemory_Invalidate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	memory := builder.CreateDeviceMemoryObject(mockDriver, device.Handle(), mocks.NewFakeDeviceMemoryHandle(), common.Vulkan1_0, 113)

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

	_, err := memory.InvalidateAll()
	require.NoError(t, err)
}

func TestVulkanDeviceMemory_AllocateAndFreeMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
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

	memory, _, err := device.AllocateMemory(nil, core1_0.MemoryAllocateInfo{
		AllocationSize:  7,
		MemoryTypeIndex: 3,
	})
	require.NoError(t, err)
	require.NotNil(t, memory)
	require.Equal(t, memoryHandle, memory.Handle())

	memory.Free(nil)
}
