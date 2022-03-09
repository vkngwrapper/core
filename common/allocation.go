package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "unsafe"

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

const (
	InternalAllocationExecutable = C.VK_INTERNAL_ALLOCATION_TYPE_EXECUTABLE
)

var internalAllocationTypeToString = map[InternalAllocationType]string{
	InternalAllocationExecutable: "Executable Allocation",
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
