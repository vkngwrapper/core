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
	InternalAllocationTypeExecutable = C.VK_INTERNAL_ALLOCATION_TYPE_EXECUTABLE
)

var internalAllocationTypeToString = map[InternalAllocationType]string{
	InternalAllocationTypeExecutable: "Executable Allocation",
}

type AllocationFunction func(userData any, size int, alignment int, allocationScope SystemAllocationScope) unsafe.Pointer
type ReallocationFunction func(userData any, original unsafe.Pointer, size int, allignment int, allocationScope SystemAllocationScope) unsafe.Pointer
type FreeFunction func(userData any, memory unsafe.Pointer)
type InternalAllocationNotification func(userData any, size int, allocationType InternalAllocationType, allocationScope SystemAllocationScope)
type InternalFreeNotification func(userData any, size int, allocationType InternalAllocationType, allocationScope SystemAllocationScope)

type AllocationCallbackOptions struct {
	UserData           any
	Allocation         AllocationFunction
	Reallocation       ReallocationFunction
	Free               FreeFunction
	InternalAllocation InternalAllocationNotification
	InternalFree       InternalFreeNotification
}
