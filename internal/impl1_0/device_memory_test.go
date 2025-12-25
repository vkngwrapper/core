package impl1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestVulkanDeviceMemory_MapMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	memory := mocks.NewDummyDeviceMemory(device, 1)
	memoryPtr := unsafe.Pointer(t)

	mockLoader.EXPECT().VkMapMemory(device.Handle(), memory.Handle(), loader.VkDeviceSize(1), loader.VkDeviceSize(3), loader.VkMemoryMapFlags(0), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, memory loader.VkDeviceMemory, offset loader.VkDeviceSize, size loader.VkDeviceSize, flags loader.VkMemoryMapFlags, ppData *unsafe.Pointer) (common.VkResult, error) {
			*ppData = memoryPtr

			return core1_0.VKSuccess, nil
		})

	ptr, _, err := driver.MapMemory(memory, 1, 3, 0)
	require.Equal(t, memoryPtr, ptr)
	require.NoError(t, err)
}

func TestVulkanDeviceMemory_UnmapMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	memory := mocks.NewDummyDeviceMemory(device, 1)

	mockLoader.EXPECT().VkUnmapMemory(device.Handle(), memory.Handle())

	driver.UnmapMemory(memory)
}

func TestVulkanDeviceMemory_Commitment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	memory := mocks.NewDummyDeviceMemory(device, 1)

	mockLoader.EXPECT().VkGetDeviceMemoryCommitment(device.Handle(), memory.Handle(), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, memory loader.VkDeviceMemory, pCommitment *loader.VkDeviceSize) {
			*pCommitment = 3
		})

	require.Equal(t, 3, driver.GetDeviceMemoryCommitment(memory))
}

func TestVulkanDeviceMemory_Flush(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	mockLoader.EXPECT().DeviceHandle().Return(device.Handle()).AnyTimes()
	driver := mocks1_0.InternalDeviceDriver(mockLoader)
	memory := mocks.NewDummyDeviceMemory(device, 113)

	mockLoader.EXPECT().VkFlushMappedMemoryRanges(device.Handle(), loader.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, memoryRangeCount loader.Uint32, pMemoryRanges *loader.VkMappedMemoryRange) (common.VkResult, error) {
			val := reflect.ValueOf(pMemoryRanges).Elem()

			require.Equal(t, uint64(6), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, memory.Handle(), (loader.VkDeviceMemory)(unsafe.Pointer(val.FieldByName("memory").Elem().UnsafeAddr())))
			require.Equal(t, uint64(0), val.FieldByName("offset").Uint())
			require.Equal(t, uint64(113), val.FieldByName("size").Uint())

			return core1_0.VKSuccess, nil
		})

	_, err := driver.FlushMappedMemoryRanges(core1_0.MappedMemoryRange{
		Memory: memory,
		Offset: 0,
		Size:   memory.Size(),
	})
	require.NoError(t, err)
}

func TestVulkanDeviceMemory_Invalidate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	mockLoader.EXPECT().DeviceHandle().Return(device.Handle()).AnyTimes()
	driver := mocks1_0.InternalDeviceDriver(mockLoader)
	memory := mocks.NewDummyDeviceMemory(device, 113)

	mockLoader.EXPECT().VkInvalidateMappedMemoryRanges(device.Handle(), loader.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, memoryRangeCount loader.Uint32, pMemoryRanges *loader.VkMappedMemoryRange) (common.VkResult, error) {
			val := reflect.ValueOf(pMemoryRanges).Elem()

			require.Equal(t, uint64(6), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, memory.Handle(), (loader.VkDeviceMemory)(unsafe.Pointer(val.FieldByName("memory").Elem().UnsafeAddr())))
			require.Equal(t, uint64(0), val.FieldByName("offset").Uint())
			require.Equal(t, uint64(113), val.FieldByName("size").Uint())

			return core1_0.VKSuccess, nil
		})

	_, err := driver.InvalidateMappedMemoryRanges(core1_0.MappedMemoryRange{
		Memory: memory,
		Offset: 0,
		Size:   memory.Size(),
	})
	require.NoError(t, err)
}

func TestVulkanDeviceMemory_AllocateAndFreeMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	expectedMemory := mocks.NewDummyDeviceMemory(device, 1)

	mockLoader.EXPECT().VkAllocateMemory(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkMemoryAllocateInfo, pAllocator *loader.VkAllocationCallbacks, pMemory *loader.VkDeviceMemory) (common.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(5), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(7), val.FieldByName("allocationSize").Uint())
			require.Equal(t, uint64(3), val.FieldByName("memoryTypeIndex").Uint())

			*pMemory = expectedMemory.Handle()
			return core1_0.VKSuccess, nil
		})
	mockLoader.EXPECT().VkFreeMemory(device.Handle(), expectedMemory.Handle(), nil)

	memory, _, err := driver.AllocateMemory(device, nil, core1_0.MemoryAllocateInfo{
		AllocationSize:  7,
		MemoryTypeIndex: 3,
	})
	require.NoError(t, err)
	require.NotNil(t, memory)
	require.Equal(t, expectedMemory.Handle(), memory.Handle())

	driver.FreeMemory(memory, nil)
}
