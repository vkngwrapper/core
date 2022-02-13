package core1_0

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

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
	return common.FlagsToString(f, memoryPropertyFlagsToString)
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
	return common.FlagsToString(f, memoryHeapFlagsToString)
}

type MemoryHeap struct {
	Size  uint64
	Flags MemoryHeapFlags
}

type PhysicalDeviceMemoryProperties struct {
	MemoryTypes []MemoryType
	MemoryHeaps []MemoryHeap
}
