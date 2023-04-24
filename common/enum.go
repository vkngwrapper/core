package common

import "github.com/pkg/errors"

////

// InternalAllocationType classifies in an internal allocation in an AllocationCallbackOptions
// notification callback.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkInternalAllocationType.html
type InternalAllocationType int32

var internalAllocationTypeMapping = make(map[InternalAllocationType]string)

func (e InternalAllocationType) Register(str string) {
	internalAllocationTypeMapping[e] = str
}

func (e InternalAllocationType) String() string {
	return internalAllocationTypeMapping[e]
}

////

// SystemAllocationScope indicates how long an allocation is intended to last in an
// AllocationCallbacksOptions allocation callback
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSystemAllocationScope.html
type SystemAllocationScope int32

var systemAllocationScopeMapping = make(map[SystemAllocationScope]string)

func (e SystemAllocationScope) Register(str string) {
	systemAllocationScopeMapping[e] = str
}

func (e SystemAllocationScope) String() string {
	return systemAllocationScopeMapping[e]
}

////

// VkResult is a return code used by many Vulkan commands to indicate failure cases or other
// status information.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkResult.html
type VkResult int

var vkResultMapping = make(map[VkResult]string)

func (e VkResult) Register(str string) {
	vkResultMapping[e] = str
}

func (e VkResult) String() string {
	return vkResultMapping[e]
}

func (e VkResult) ToError() error {
	if e >= 0 {
		// All VKError* are <0
		return nil
	}

	return errors.WithStack(&VkResultError{e})
}
