package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

const (
	// DescriptorPoolCreateFreeDescriptorSet specifies that DescriptorSet objects can return their
	// individual allocations to the pool
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateFlagBits.html
	DescriptorPoolCreateFreeDescriptorSet DescriptorPoolCreateFlags = C.VK_DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT
)

func init() {
	DescriptorPoolCreateFreeDescriptorSet.Register("Free Descriptor Set")
}

// DescriptorPoolSize specifies DescriptorPool size
type DescriptorPoolSize struct {
	// Type is the type of descriptor
	Type DescriptorType
	// DescriptorCount is the number of descriptors of that type ot allocate
	DescriptorCount int
}

// DescriptorPoolCreateInfo specifies parameters of a newly-created DescriptorPool
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateInfo.html
type DescriptorPoolCreateInfo struct {
	// Flags specifies certain supported operations on the pool
	Flags DescriptorPoolCreateFlags

	// MaxSets is the maximum number of DescriptorSet objects that can be allocated from the pool
	MaxSets int
	// PoolSizes is a slice of DescriptorPoolSize structures, each containing a descriptor type
	// and number of descriptors of that type to be allocated in the pool
	PoolSizes []DescriptorPoolSize

	common.NextOptions
}

func (o DescriptorPoolCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDescriptorPoolCreateInfo)
	}

	createInfo := (*C.VkDescriptorPoolCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO
	createInfo.flags = C.VkDescriptorPoolCreateFlags(o.Flags)
	createInfo.pNext = next

	createInfo.maxSets = C.uint32_t(o.MaxSets)
	sizeCount := len(o.PoolSizes)
	createInfo.poolSizeCount = C.uint32_t(sizeCount)
	createInfo.pPoolSizes = nil

	if sizeCount > 0 {
		poolsPtr := (*C.VkDescriptorPoolSize)(allocator.Malloc(sizeCount * C.sizeof_struct_VkDescriptorPoolSize))
		poolsSlice := ([]C.VkDescriptorPoolSize)(unsafe.Slice(poolsPtr, sizeCount))

		for i := 0; i < sizeCount; i++ {
			poolsSlice[i]._type = C.VkDescriptorType(o.PoolSizes[i].Type)
			poolsSlice[i].descriptorCount = C.uint32_t(o.PoolSizes[i].DescriptorCount)
		}

		createInfo.pPoolSizes = poolsPtr
	}

	return unsafe.Pointer(createInfo), nil
}
