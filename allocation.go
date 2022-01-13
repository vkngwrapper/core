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

type AllocationCallbackOptions struct {
	UserData           interface{}
	Allocation         AllocationFunction
	Reallocation       ReallocationFunction
	Free               FreeFunction
	InternalAllocation InternalAllocationNotification
	InternalFree       InternalFreeNotification
}

type AllocationCallbacks struct {
	userData           interface{}
	allocation         AllocationFunction
	reallocation       ReallocationFunction
	free               FreeFunction
	internalAllocation InternalAllocationNotification
	internalFree       InternalFreeNotification

	handle *VkAllocationCallbacks
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
	c.handle = (*VkAllocationCallbacks)(handle)
}

func (c *AllocationCallbacks) Handle() *VkAllocationCallbacks {
	if c == nil {
		return nil
	}

	return c.handle
}

func (c *AllocationCallbacks) Destroy() {
	cgo.Handle(c.handle.pUserData).Delete()
	C.free(unsafe.Pointer(c.handle))
	c.handle = nil
}
