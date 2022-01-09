package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
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
	"github.com/CannibalVox/cgoparam"
	"runtime/cgo"
	"unsafe"
)

type SystemAllocationScope int32

const (
	SystemAllocationScopeCommand  SystemAllocationScope = C.VK_SYSTEM_ALLOCATION_SCOPE_COMMAND
	SystemAllocationScopeObject   SystemAllocationScope = C.VK_SYSTEM_ALLOCATION_SCOPE_OBJECT
	SystemAllocationScopeCache    SystemAllocationScope = C.VK_SYSTEM_ALLOCATION_SCOPE_CACHE
	SystemAllocationScopeDevice   SystemAllocationScope = C.VK_SYSTEM_ALLOCATION_SCOPE_DEVICE
	SystemAllocationScopeInstance SystemAllocationScope = C.VK_SYSTEM_ALLOCATION_SCOPE_INSTANCE
)

var systemAllocationScopeToString = map[SystemAllocationScope]string{
	SystemAllocationScopeCommand:  "Command Scope",
	SystemAllocationScopeObject:   "Object Scope",
	SystemAllocationScopeCache:    "Cache Scope",
	SystemAllocationScopeDevice:   "Device Scope",
	SystemAllocationScopeInstance: "Instance Scope",
}

func (s SystemAllocationScope) String() string {
	return systemAllocationScopeToString[s]
}

type InternalAllocationType int32

const (
	InternalAllocationExecutable = C.VK_INTERNAL_ALLOCATION_TYPE_EXECUTABLE
)

var internalAllocationTypeToString = map[InternalAllocationType]string{
	InternalAllocationExecutable: "Executable Allocation",
}

func (t InternalAllocationType) String() string {
	return internalAllocationTypeToString[t]
}

type AllocationFunction func(userData interface{}, size int, alignment int, allocationScope SystemAllocationScope) unsafe.Pointer
type ReallocationFunction func(userData interface{}, original unsafe.Pointer, size int, allignment int, allocationScope SystemAllocationScope) unsafe.Pointer
type FreeFunction func(userData interface{}, memory unsafe.Pointer)
type InternalAllocationNotification func(userData interface{}, size int, allocationType InternalAllocationType, allocationScope SystemAllocationScope)
type InternalFreeNotification func(userData interface{}, size int, allocationType InternalAllocationType, allocationScope SystemAllocationScope)

type AllocationCallbacks struct {
	userData           interface{}
	allocation         AllocationFunction
	reallocation       ReallocationFunction
	free               FreeFunction
	internalAllocation InternalAllocationNotification
	internalFree       InternalFreeNotification

	thisHandle     cgo.Handle
	wantsToDestroy bool
	destroyed      bool
	internalRefs   int
}

func (c *AllocationCallbacks) Destroy() {
	c.wantsToDestroy = true
	if c.internalRefs == 0 {
		c.destroyed = true
		c.thisHandle.Delete()
	}
}

func (c *AllocationCallbacks) BuildHandle(allocator *cgoparam.Allocator) *VkAllocationCallbacks {
	if c == nil {
		return nil
	}
	if c.destroyed {
		return nil
	}

	outHandle := (*C.VkAllocationCallbacks)(allocator.Malloc(C.sizeof_struct_VkAllocationCallbacks))
	outHandle.pUserData = unsafe.Pointer(c.thisHandle)
	outHandle.pfnAllocation = nil
	outHandle.pfnReallocation = nil
	outHandle.pfnFree = nil
	outHandle.pfnInternalAllocation = nil
	outHandle.pfnInternalFree = nil

	if c.allocation != nil {
		outHandle.pfnAllocation = (C.PFN_vkAllocationFunction)(C.allocationCallback)
	}

	if c.reallocation != nil {
		outHandle.pfnReallocation = (C.PFN_vkReallocationFunction)(C.reallocationCallback)
	}

	if c.free != nil {
		outHandle.pfnFree = (C.PFN_vkFreeFunction)(C.freeCallback)
	}

	if c.internalAllocation != nil {
		outHandle.pfnInternalAllocation = (C.PFN_vkInternalAllocationNotification)(C.internalAllocationCallback)
	}

	if c.internalFree != nil {
		outHandle.pfnInternalFree = (C.PFN_vkInternalFreeNotification)(C.internalFreeCallback)
	}

	return (*VkAllocationCallbacks)(outHandle)
}

func (c *AllocationCallbacks) BeginReference() {
	if c == nil {
		return
	}
	c.internalRefs += 1
}

func (c *AllocationCallbacks) EndReference() {
	if c == nil {
		return
	}
	c.internalRefs -= 1

	if c.internalRefs == 0 && c.wantsToDestroy && !c.destroyed {
		c.destroyed = true
		c.thisHandle.Delete()
	}
}
