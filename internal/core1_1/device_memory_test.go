package internal1_1_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestMemoryDedicatedAllocateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil).AnyTimes()
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)

	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	deviceHandle := mocks.EasyMockDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkCreateDevice(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pCreateInfo *driver.VkDeviceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDevice *driver.VkDevice) (common.VkResult, error) {
			*pDevice = deviceHandle.Handle()

			return core1_0.VKSuccess, nil
		})

	device, _, err := loader.CreateDevice(physicalDevice, nil, core1_0.DeviceCreateOptions{
		QueueFamilies: []core1_0.DeviceQueueCreateOptions{
			{
				CreatedQueuePriorities: []float32{0},
			},
		},
	})
	require.NoError(t, err)

	buffer := mocks.EasyMockBuffer(ctrl)
	expectedMemory := mocks.EasyMockDeviceMemory(ctrl)

	coreDriver.EXPECT().VkAllocateMemory(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pAllocateInfo *driver.VkMemoryAllocateInfo, pAllocator *driver.VkAllocationCallbacks, pMemory *driver.VkDeviceMemory) (common.VkResult, error) {
			*pMemory = expectedMemory.Handle()

			options := reflect.ValueOf(pAllocateInfo).Elem()
			require.Equal(t, uint64(5), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
			require.Equal(t, uint64(1), options.FieldByName("allocationSize").Uint())
			require.Equal(t, uint64(3), options.FieldByName("memoryTypeIndex").Uint())

			dedicatedPtr := (*driver.VkMemoryDedicatedAllocateInfo)(options.FieldByName("pNext").UnsafePointer())
			dedicated := reflect.ValueOf(dedicatedPtr).Elem()

			require.Equal(t, uint64(1000127001), dedicated.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO
			require.True(t, dedicated.FieldByName("pNext").IsNil())
			require.Equal(t, buffer.Handle(), driver.VkBuffer(dedicated.FieldByName("buffer").UnsafePointer()))
			require.True(t, dedicated.FieldByName("image").IsNil())

			return core1_0.VKSuccess, nil
		})

	memory, _, err := loader.AllocateMemory(device, nil, core1_0.MemoryAllocateOptions{
		AllocationSize:  1,
		MemoryTypeIndex: 3,
		HaveNext: common.HaveNext{Next: core1_1.MemoryDedicatedAllocationOptions{
			Buffer: buffer,
		}},
	})
	require.NoError(t, err)
	require.Equal(t, expectedMemory.Handle(), memory.Handle())

}
