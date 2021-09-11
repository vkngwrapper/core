package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/cgoalloc"
	"strings"
)

type MemoryPropertyFlags int32

const (
	MemoryDeviceLocal       MemoryPropertyFlags = C.VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT
	MemoryHostVisible       MemoryPropertyFlags = C.VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT
	MemoryHostCoherent      MemoryPropertyFlags = C.VK_MEMORY_PROPERTY_HOST_COHERENT_BIT
	MemoryLazilyAllocated   MemoryPropertyFlags = C.VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT
	MemoryProtected         MemoryPropertyFlags = C.VK_MEMORY_PROPERTY_PROTECTED_BIT
	MemoryDeviceCoherentAMD MemoryPropertyFlags = C.VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD
	MemoryDeviceUncachedAMD MemoryPropertyFlags = C.VK_MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD
	MemoryRDMACapableNV     MemoryPropertyFlags = C.VK_MEMORY_PROPERTY_RDMA_CAPABLE_BIT_NV
)

var memoryPropertyFlagsToString = map[MemoryPropertyFlags]string{
	MemoryDeviceLocal:       "Device Local",
	MemoryHostVisible:       "Host Visible",
	MemoryHostCoherent:      "Host Coherent",
	MemoryLazilyAllocated:   "Lazily Allocated",
	MemoryProtected:         "Protected",
	MemoryDeviceCoherentAMD: "Device Coherent (AMD)",
	MemoryDeviceUncachedAMD: "Device Uncached (AMD)",
	MemoryRDMACapableNV:     "RDMA Capable (Nvidia)",
}

func (f MemoryPropertyFlags) String() string {
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := MemoryPropertyFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := memoryPropertyFlagsToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteRune('|')
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}

type MemoryType struct {
	Properties MemoryPropertyFlags
	HeapIndex  int
}

type MemoryHeapFlags int32

const (
	HeapDeviceLocal   MemoryHeapFlags = C.VK_MEMORY_HEAP_DEVICE_LOCAL_BIT
	HeapMultiInstance MemoryHeapFlags = C.VK_MEMORY_HEAP_MULTI_INSTANCE_BIT
)

var memoryHeapFlagsToString = map[MemoryHeapFlags]string{
	HeapDeviceLocal:   "Device Local",
	HeapMultiInstance: "Multi-Instance",
}

func (f MemoryHeapFlags) String() string {
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := MemoryHeapFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := memoryHeapFlagsToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteRune('|')
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}

type MemoryHeap struct {
	Size  uint64
	Flags MemoryHeapFlags
}

type PhysicalDeviceMemoryProperties struct {
	MemoryTypes []MemoryType
	MemoryHeaps []MemoryHeap
}

func (d *VulkanPhysicalDevice) MemoryProperties(allocator cgoalloc.Allocator) *PhysicalDeviceMemoryProperties {
	propsUnsafe := allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceMemoryProperties)
	defer allocator.Free(propsUnsafe)

	d.loader.VkGetPhysicalDeviceMemoryProperties(d.handle, (*loader.VkPhysicalDeviceMemoryProperties)(propsUnsafe))

	props := (*C.VkPhysicalDeviceMemoryProperties)(propsUnsafe)

	outProps := &PhysicalDeviceMemoryProperties{}
	typeCount := int(props.memoryTypeCount)
	heapCount := int(props.memoryHeapCount)

	for i := 0; i < typeCount; i++ {
		t := props.memoryTypes[i]
		outProps.MemoryTypes = append(outProps.MemoryTypes, MemoryType{
			Properties: MemoryPropertyFlags(t.propertyFlags),
			HeapIndex:  int(t.heapIndex),
		})
	}

	for i := 0; i < heapCount; i++ {
		heap := props.memoryHeaps[i]
		outProps.MemoryHeaps = append(outProps.MemoryHeaps, MemoryHeap{
			Size:  uint64(heap.size),
			Flags: MemoryHeapFlags(heap.flags),
		})
	}

	return outProps
}
