package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
)

func (d *VulkanPhysicalDevice) MemoryProperties() *PhysicalDeviceMemoryProperties {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	propsUnsafe := allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceMemoryProperties)

	d.instanceDriver.VkGetPhysicalDeviceMemoryProperties(d.physicalDeviceHandle, (*driver.VkPhysicalDeviceMemoryProperties)(propsUnsafe))

	props := (*C.VkPhysicalDeviceMemoryProperties)(propsUnsafe)

	outProps := &PhysicalDeviceMemoryProperties{}
	typeCount := int(props.memoryTypeCount)
	heapCount := int(props.memoryHeapCount)

	for i := 0; i < typeCount; i++ {
		t := props.memoryTypes[i]
		outProps.MemoryTypes = append(outProps.MemoryTypes, MemoryType{
			Properties: MemoryProperties(t.propertyFlags),
			HeapIndex:  int(t.heapIndex),
		})
	}

	for i := 0; i < heapCount; i++ {
		heap := props.memoryHeaps[i]
		outProps.MemoryHeaps = append(outProps.MemoryHeaps, MemoryHeap{
			Size:  int(heap.size),
			Flags: MemoryHeapFlags(heap.flags),
		})
	}

	return outProps
}
