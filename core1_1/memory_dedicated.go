package core1_1

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

// MemoryDedicatedAllocateInfo specifies a dedicated memory allocation resource
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html
type MemoryDedicatedAllocateInfo struct {
	// Image is nil or the Image object which this memory will be bound to
	Image core1_0.Image
	// Buffer is nil or the Buffer object this memory will be bound to
	Buffer core1_0.Buffer

	common.NextOptions
}

func (o MemoryDedicatedAllocateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Image.Initialized() && o.Buffer.Initialized() {
		return nil, errors.New("both Image and Buffer fields are set in MemoryDedicatedAllocateInfo- only one must be set")
	} else if !o.Image.Initialized() && !o.Buffer.Initialized() {
		return nil, errors.New("neither Image nor Buffer fields are set in MemoryDedicatedAllocateInfo- one must be set")
	}

	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryDedicatedAllocateInfo{})))
	}

	createInfo := (*C.VkMemoryDedicatedAllocateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO
	createInfo.pNext = next
	createInfo.image = nil
	createInfo.buffer = nil

	if o.Image.Initialized() {
		createInfo.image = C.VkImage(unsafe.Pointer(o.Image.Handle()))
	} else if o.Buffer.Initialized() {
		createInfo.buffer = C.VkBuffer(unsafe.Pointer(o.Buffer.Handle()))
	}

	return preallocatedPointer, nil
}

////

// MemoryDedicatedRequirements describes dedicated allocation requirements of Buffer and Image
// resources
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedRequirements.html
type MemoryDedicatedRequirements struct {
	// PrefersDedicatedAllocation specifies that the implementation would prefer a dedicated
	// allocation for this resource. The application is still free to suballocate the resource
	// but it may get better performance if a dedicated allocation is used
	PrefersDedicatedAllocation bool
	// RequiresDedicatedAllocation specifies that a dedicated allocation is required for this resource
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
	o.RequiresDedicatedAllocation = loader.VkBool32(outData.requiresDedicatedAllocation) != loader.VkBool32(0)
	o.PrefersDedicatedAllocation = loader.VkBool32(outData.prefersDedicatedAllocation) != loader.VkBool32(0)

	return outData.pNext, nil
}
