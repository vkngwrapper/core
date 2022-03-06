package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type DescriptorPoolResetFlags int32
type DescriptorPoolFlags int32

const (
	DescriptorPoolFreeDescriptorSet DescriptorPoolFlags = C.VK_DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT
	DescriptorPoolUpdateAfterBind   DescriptorPoolFlags = C.VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT
	DescriptorPoolHostOnlyValve     DescriptorPoolFlags = C.VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_VALVE
)

var descriptorPoolFlagsToString = map[DescriptorPoolFlags]string{
	DescriptorPoolFreeDescriptorSet: "Free Descriptor Set",
	DescriptorPoolUpdateAfterBind:   "Update After Bind",
	DescriptorPoolHostOnlyValve:     "Host-Only (Valve)",
}

func (f DescriptorPoolFlags) String() string {
	return common.FlagsToString(f, descriptorPoolFlagsToString)
}

type PoolSize struct {
	Type            common.DescriptorType
	DescriptorCount int
}

type DescriptorPoolOptions struct {
	Flags DescriptorPoolFlags

	MaxSets   int
	PoolSizes []PoolSize

	core.HaveNext
}

func (o DescriptorPoolOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
