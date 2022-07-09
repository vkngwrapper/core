package common

/*
#include <stdlib.h>
#include "vulkan.h"
*/
import "C"
import "unsafe"

const (
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSystemAllocationScope.html

	// SystemAllocationScopeCommand specifies that the allocation is scoped to the duration of the Vulkan command
	SystemAllocationScopeCommand SystemAllocationScope = C.VK_SYSTEM_ALLOCATION_SCOPE_COMMAND
	// SystemAllocationScopeObject specifies that the allocation is scoped to the lifetime of the Vulkan object that is being created or used.
	SystemAllocationScopeObject SystemAllocationScope = C.VK_SYSTEM_ALLOCATION_SCOPE_OBJECT
	// SystemAllocationScopeCache specifies that the allocation is scoped to the lifetime of a VkPipelineCache or VkValidationCacheEXT object.
	SystemAllocationScopeCache SystemAllocationScope = C.VK_SYSTEM_ALLOCATION_SCOPE_CACHE
	// SystemAllocationScopeDevice specifies that the allocation is scoped to the lifetime of the Vulkan device.
	SystemAllocationScopeDevice SystemAllocationScope = C.VK_SYSTEM_ALLOCATION_SCOPE_DEVICE
	// SystemAllocationScopeInstance specifies that the allocation is scoped to the lifetime of the Vulkan instance.
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
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkInternalAllocationType.html

	// InternalAllocationTypeExecutable specifies that the allocation is intended for execution by the host.
	InternalAllocationTypeExecutable InternalAllocationType = C.VK_INTERNAL_ALLOCATION_TYPE_EXECUTABLE
)

var internalAllocationTypeToString = map[InternalAllocationType]string{
	InternalAllocationTypeExecutable: "Executable Allocation",
}

// AllocationFunction is an application-defined memory allocation function. It can be registered with
// Vulkan via driver.CreateAllocationCallbacks in order to allow users to override the default Vulkan
// host memory allocation behavior. Beware: this should provide unsafe.Pointers to **C** memory, not go
// memory!
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/PFN_vkAllocationFunction.html
type AllocationFunction func(userData any, size int, alignment int, allocationScope SystemAllocationScope) unsafe.Pointer

// ReallocationFunction is an application-defined memory reallocation function. It can be registered with
// Vulkan via driver.CreateAllocationCallbacks in order to allow users to override the default Vulkan
// host memory reallocation behavior. `original` is a pointer to the memory being reallocated, and
// may be nil. When `original` is nil, ReallocationFunction must act equivalently to the AllocationFunction
// in the same AllocationCallbacks object.
// Beware: this should provide unsafe.Pointers to **C** memory, not Go memory!
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/PFN_vkReallocationFunction.html
type ReallocationFunction func(userData any, original unsafe.Pointer, size int, alignment int, allocationScope SystemAllocationScope) unsafe.Pointer

// FreeFunction is an application-defined memory free function. It can be registered with Vulkan via
// driver.CreateAllocationCallbacks in order to allow users to override the default Vulkan host memory
// free behavior.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/PFN_vkFreeFunction.html
type FreeFunction func(userData any, memory unsafe.Pointer)

// InternalAllocationNotification is an application-defined function to received allocation notifications.
// It is purely an informational callback. It can be registered with Vulkan via driver.CreateAllocationCallbacks
// in order to allow users to receive notifications when the Vulkan host makes memory allocations.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/PFN_vkInternalAllocationNotification.html
type InternalAllocationNotification func(userData any, size int, allocationType InternalAllocationType, allocationScope SystemAllocationScope)

// InternalFreeNotification is an application-defined function to receive free notifications. It is
// purely an informational callback. It can be registered with Vulkan via driver.CreateAllocationCallbacks
// in order to allow users to receive notifications when the Vulkan host frees memory.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/PFN_vkInternalFreeNotification.html
type InternalFreeNotification func(userData any, size int, allocationType InternalAllocationType, allocationScope SystemAllocationScope)

// AllocationCallbackOptions defines a set of callbacks and a piece of userdata that can be baked
// into a driver.AllocationCallbacks with driver.CreateAllocationCallbacks. driver.AllocationCallbacks
// objects are immutable, so it may be valuable to hold onto your AllocationCallbackOptions object
// in order to allow changes over time, if that is necessary.
//
// All fields in this struct are optional: any nil callback will not be executed by Vulkan, and
// nil UserData will be passed as-is to any executed callback.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html
type AllocationCallbackOptions struct {
	// An arbitrary go object that will be passed to all registered callbacks
	UserData any
	// Callback, executed to allocate host memory
	Allocation AllocationFunction
	// Callback, executed to reallocate host memory
	Reallocation ReallocationFunction
	// Callback, executed to free host memory
	Free FreeFunction
	// Callback, executed on all internal memory allocations
	InternalAllocation InternalAllocationNotification
	// Callback, executed on all internal memory frees
	InternalFree InternalFreeNotification
}
