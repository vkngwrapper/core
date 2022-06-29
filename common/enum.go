package common

import "github.com/cockroachdb/errors"

////

type InternalAllocationType int32

var internalAllocationTypeMapping = make(map[InternalAllocationType]string)

func (e InternalAllocationType) Register(str string) {
	internalAllocationTypeMapping[e] = str
}

func (e InternalAllocationType) String() string {
	return internalAllocationTypeMapping[e]
}

////

type SystemAllocationScope int32

var systemAllocationScopeMapping = make(map[SystemAllocationScope]string)

func (e SystemAllocationScope) Register(str string) {
	systemAllocationScopeMapping[e] = str
}

func (e SystemAllocationScope) String() string {
	return systemAllocationScopeMapping[e]
}

////

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
