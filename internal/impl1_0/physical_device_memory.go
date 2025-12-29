package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *InstanceVulkanDriver) GetPhysicalDeviceMemoryProperties(physicalDevice core.PhysicalDevice) *core1_0.PhysicalDeviceMemoryProperties {
	if !physicalDevice.Initialized() {
		panic("physicalDevice was uninitialized")
	}

	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	propsUnsafe := allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceMemoryProperties)

	v.LoaderObj.VkGetPhysicalDeviceMemoryProperties(physicalDevice.Handle(), (*loader.VkPhysicalDeviceMemoryProperties)(propsUnsafe))

	props := (*C.VkPhysicalDeviceMemoryProperties)(propsUnsafe)

	outProps := &core1_0.PhysicalDeviceMemoryProperties{}
	typeCount := int(props.memoryTypeCount)
	heapCount := int(props.memoryHeapCount)

	for i := 0; i < typeCount; i++ {
		t := props.memoryTypes[i]
		outProps.MemoryTypes = append(outProps.MemoryTypes, core1_0.MemoryType{
			PropertyFlags: core1_0.MemoryPropertyFlags(t.propertyFlags),
			HeapIndex:     int(t.heapIndex),
		})
	}

	for i := 0; i < heapCount; i++ {
		heap := props.memoryHeaps[i]
		outProps.MemoryHeaps = append(outProps.MemoryHeaps, core1_0.MemoryHeap{
			Size:  int(heap.size),
			Flags: core1_0.MemoryHeapFlags(heap.flags),
		})
	}

	return outProps
}
