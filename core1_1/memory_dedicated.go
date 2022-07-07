package core1_1

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
	"github.com/cockroachdb/errors"
	"unsafe"
)

type MemoryDedicatedAllocationInfo struct {
	Image  core1_0.Image
	Buffer core1_0.Buffer

	common.NextOptions
}

func (o MemoryDedicatedAllocationInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Image != nil && o.Buffer != nil {
		return nil, errors.New("both Image and Buffer fields are set in MemoryDedicatedAllocationInfo- only one must be set")
	} else if o.Image == nil && o.Buffer == nil {
		return nil, errors.New("neither Image nor Buffer fields are set in MemoryDedicatedAllocationInfo- one must be set")
	}

	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryDedicatedAllocateInfo{})))
	}

	createInfo := (*C.VkMemoryDedicatedAllocateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO
	createInfo.pNext = next
	createInfo.image = nil
	createInfo.buffer = nil

	if o.Image != nil {
		createInfo.image = C.VkImage(unsafe.Pointer(o.Image.Handle()))
	} else if o.Buffer != nil {
		createInfo.buffer = C.VkBuffer(unsafe.Pointer(o.Buffer.Handle()))
	}

	return preallocatedPointer, nil
}

////

type MemoryDedicatedRequirements struct {
	PrefersDedicatedAllocation  bool
	RequiresDedicatedAllocation bool

	common.NextOutData
}

func (o *MemoryDedicatedRequirements) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryDedicatedRequirements{})))
	}

	outData := (*C.VkMemoryDedicatedRequirements)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *MemoryDedicatedRequirements) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkMemoryDedicatedRequirements)(cDataPointer)
	o.RequiresDedicatedAllocation = driver.VkBool32(outData.requiresDedicatedAllocation) != driver.VkBool32(0)
	o.PrefersDedicatedAllocation = driver.VkBool32(outData.prefersDedicatedAllocation) != driver.VkBool32(0)

	return outData.pNext, nil
}
