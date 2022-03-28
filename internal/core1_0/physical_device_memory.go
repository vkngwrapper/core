package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
)

func (d *VulkanPhysicalDevice) MemoryProperties() *core1_0.PhysicalDeviceMemoryProperties {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	propsUnsafe := allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceMemoryProperties)

	d.InstanceDriver.VkGetPhysicalDeviceMemoryProperties(d.PhysicalDeviceHandle, (*driver.VkPhysicalDeviceMemoryProperties)(propsUnsafe))

	props := (*C.VkPhysicalDeviceMemoryProperties)(propsUnsafe)

	outProps := &core1_0.PhysicalDeviceMemoryProperties{}
	typeCount := int(props.memoryTypeCount)
	heapCount := int(props.memoryHeapCount)

	for i := 0; i < typeCount; i++ {
		t := props.memoryTypes[i]
		outProps.MemoryTypes = append(outProps.MemoryTypes, common.MemoryType{
			Properties: common.MemoryProperties(t.propertyFlags),
			HeapIndex:  int(t.heapIndex),
		})
	}

	for i := 0; i < heapCount; i++ {
		heap := props.memoryHeaps[i]
		outProps.MemoryHeaps = append(outProps.MemoryHeaps, common.MemoryHeap{
			Size:  int(heap.size),
			Flags: common.MemoryHeapFlags(heap.flags),
		})
	}

	return outProps
}
