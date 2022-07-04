package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

const (
	DescriptorPoolCreateFreeDescriptorSet DescriptorPoolCreateFlags = C.VK_DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT
)

func init() {
	DescriptorPoolCreateFreeDescriptorSet.Register("Free Descriptor Set")
}

type PoolSize struct {
	Type            DescriptorType
	DescriptorCount int
}

type DescriptorPoolCreateOptions struct {
	Flags DescriptorPoolCreateFlags

	MaxSets   int
	PoolSizes []PoolSize

	common.NextOptions
}

func (o DescriptorPoolCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
