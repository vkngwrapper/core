package driver

/*
#include <stdlib.h>
#include "../common/vulkan.h"
#include "allocation_callbacks.h"

VKAPI_ATTR void* VKAPI_CALL allocationCallback(
	void *pUserData,
	size_t size,
	size_t alignment,
	VkSystemAllocationScope allocationScope) {

	return goAllocationCallback(pUserData, size, alignment, allocationScope);
}

VKAPI_ATTR void* VKAPI_CALL reallocationCallback(
	void *pUserData,
	void *pOriginal,
	size_t size,
	size_t alignment,
	VkSystemAllocationScope allocationScope) {

	return goReallocationCallback(pUserData, pOriginal, size, alignment, allocationScope);
}

VKAPI_ATTR void VKAPI_CALL freeCallback(
	void *pUserData,
	void *pMemory) {

	return goFreeCallback(pUserData, pMemory);
}

VKAPI_ATTR void VKAPI_CALL internalAllocationCallback(
	void *pUserData,
	size_t size,
	VkInternalAllocationType allocationType,
	VkSystemAllocationScope allocationScope) {

	return goInternalAllocationCallback(pUserData, size, allocationType, allocationScope);
}

VKAPI_ATTR void VKAPI_CALL internalFreeCallback(
	void *pUserData,
	size_t size,
	VkInternalAllocationType allocationType,
	VkSystemAllocationScope allocationScope) {

	return goInternalFreeCallback(pUserData, size, allocationType, allocationScope);
}
*/
import "C"
import (
	"runtime/cgo"
	"unsafe"

	"github.com/vkngwrapper/core/v3/common"
)

// CreateAllocationCallbacks accepts a (mutable) AllocationCallbacksOptions and produces an
// (immutable) AllocationCallbacks object representing the callbacks in the
// AllocationCallbacksOptions
func CreateAllocationCallbacks(o *common.AllocationCallbackOptions) *AllocationCallbacks {
	callbacks := &AllocationCallbacks{
		allocation:         o.Allocation,
		reallocation:       o.Reallocation,
		free:               o.Free,
		internalAllocation: o.InternalAllocation,
		internalFree:       o.InternalFree,
		userData:           o.UserData,
	}
	callbacks.initHandle()

	return callbacks
}

// AllocationCallbacks is a Vulkan structure that controls allocation and deallocation behavior
// for Vulkan objects. It works by passing it to Create/Free methods. This object is immutable and
// must be created with CreateAllocationCallbacks.
type AllocationCallbacks struct {
	userData           any
	allocation         common.AllocationFunction
	reallocation       common.ReallocationFunction
	free               common.FreeFunction
	internalAllocation common.InternalAllocationNotification
	internalFree       common.InternalFreeNotification

	handle *C.VkAllocationCallbacks
}

func (c *AllocationCallbacks) initHandle() {
	handle := (*C.VkAllocationCallbacks)(C.malloc(C.sizeof_struct_VkAllocationCallbacks))
	handle.pUserData = unsafe.Pointer(cgo.NewHandle(c))
	handle.pfnAllocation = nil
	handle.pfnReallocation = nil
	handle.pfnFree = nil
	handle.pfnInternalAllocation = nil
	handle.pfnInternalFree = nil

	if c.allocation != nil {
		handle.pfnAllocation = (C.PFN_vkAllocationFunction)(C.allocationCallback)
	}

	if c.reallocation != nil {
		handle.pfnReallocation = (C.PFN_vkReallocationFunction)(C.reallocationCallback)
	}

	if c.free != nil {
		handle.pfnFree = (C.PFN_vkFreeFunction)(C.freeCallback)
	}

	if c.internalAllocation != nil {
		handle.pfnInternalAllocation = (C.PFN_vkInternalAllocationNotification)(C.internalAllocationCallback)
	}

	if c.internalFree != nil {
		handle.pfnInternalFree = (C.PFN_vkInternalFreeNotification)(C.internalFreeCallback)
	}
	c.handle = handle
}

func (c *AllocationCallbacks) Handle() *VkAllocationCallbacks {
	if c == nil {
		return nil
	}

	return (*VkAllocationCallbacks)(unsafe.Pointer(c.handle))
}

func (c *AllocationCallbacks) Destroy() {
	cgo.Handle(c.handle.pUserData).Delete()
	C.free(unsafe.Pointer(c.handle))
	c.handle = nil
}
