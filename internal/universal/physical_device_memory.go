package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
)

func (d *VulkanPhysicalDevice) MemoryProperties() *core1_0.PhysicalDeviceMemoryProperties {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	propsUnsafe := allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceMemoryProperties)

	d.driver.VkGetPhysicalDeviceMemoryProperties(d.handle, (*driver.VkPhysicalDeviceMemoryProperties)(propsUnsafe))

	props := (*C.VkPhysicalDeviceMemoryProperties)(propsUnsafe)

	outProps := &core1_0.PhysicalDeviceMemoryProperties{}
	typeCount := int(props.memoryTypeCount)
	heapCount := int(props.memoryHeapCount)

	for i := 0; i < typeCount; i++ {
		t := props.memoryTypes[i]
		outProps.MemoryTypes = append(outProps.MemoryTypes, core1_0.MemoryType{
			Properties: core1_0.MemoryPropertyFlags(t.propertyFlags),
			HeapIndex:  int(t.heapIndex),
		})
	}

	for i := 0; i < heapCount; i++ {
		heap := props.memoryHeaps[i]
		outProps.MemoryHeaps = append(outProps.MemoryHeaps, core1_0.MemoryHeap{
			Size:  uint64(heap.size),
			Flags: core1_0.MemoryHeapFlags(heap.flags),
		})
	}

	return outProps
}
