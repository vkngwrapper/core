package core1_2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

// SemaphoreType specifies the type of a Semaphore object
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html
type SemaphoreType int32

var semaphoreTypeMapping = make(map[SemaphoreType]string)

func (e SemaphoreType) Register(str string) {
	semaphoreTypeMapping[e] = str
}

func (e SemaphoreType) String() string {
	return semaphoreTypeMapping[e]
}

////

// SemaphoreWaitFlags specifies additional parameters of a Semaphore wait operation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitFlagBits.html
type SemaphoreWaitFlags int32

var semaphoreWaitFlagsMapping = common.NewFlagStringMapping[SemaphoreWaitFlags]()

func (f SemaphoreWaitFlags) Register(str string) {
	semaphoreWaitFlagsMapping.Register(f, str)
}
func (f SemaphoreWaitFlags) String() string {
	return semaphoreWaitFlagsMapping.FlagsToString(f)
}

////

const (
	// SemaphoreTypeBinary specifies a binary Semaphore type that has a boolean payload
	// indicating whether the Semaphore is currently signaled or unsignaled
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html
	SemaphoreTypeBinary SemaphoreType = C.VK_SEMAPHORE_TYPE_BINARY
	// SemaphoreTypeTimeline specifies a timeline Semaphore type that has a strictly
	// increasing 64-bit unsigned integer payload indicating whether the Semaphore is signaled
	// with respect to a particular reference value
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html
	SemaphoreTypeTimeline SemaphoreType = C.VK_SEMAPHORE_TYPE_TIMELINE

	// SemaphoreWaitAny specifies that the Semaphore wait condition is that at least one of
	// the Semaphore objects in SemaphoreWaitInfo.Semaphores has reached the value specified
	// by the corresponding element of SemaphoreWaitInfo.Values
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitFlagBits.html
	SemaphoreWaitAny SemaphoreWaitFlags = C.VK_SEMAPHORE_WAIT_ANY_BIT
)

func init() {
	SemaphoreTypeBinary.Register("Binary")
	SemaphoreTypeTimeline.Register("Timeline")

	SemaphoreWaitAny.Register("Any")
}

////

// SemaphoreSignalInfo contains information about a Semaphore signal operation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreSignalInfo.html
type SemaphoreSignalInfo struct {
	// Semaphore is the Semaphore object to signal
	Semaphore core1_0.Semaphore
	// Value is the value to signal
	Value uint64

	common.NextOptions
}

func (o SemaphoreSignalInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSemaphoreSignalInfo{})))
	}

	if o.Semaphore == nil {
		return nil, errors.New("the 'Semaphore' field of SemaphoreSignalInfo must be non-nil")
	}

	info := (*C.VkSemaphoreSignalInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_SIGNAL_INFO
	info.pNext = next
	info.semaphore = C.VkSemaphore(unsafe.Pointer(o.Semaphore.Handle()))
	info.value = C.uint64_t(o.Value)

	return preallocatedPointer, nil
}

////

// SemaphoreWaitInfo contains information about the Semaphore wait condition
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitInfo.html
type SemaphoreWaitInfo struct {
	// Flags specifies additional parameters for the Semaphore wait operation
	Flags SemaphoreWaitFlags
	// Semaphores is a slice of Semaphore objects to wait on
	Semaphores []core1_0.Semaphore
	// Values is a slice of timeline Semaphore values
	Values []uint64

	common.NextOptions
}

func (o SemaphoreWaitInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSemaphoreWaitInfo{})))
	}

	if len(o.Semaphores) != len(o.Values) {
		return nil, errors.Newf("the SemaphoreWaitInfo 'Semaphores' list has %d elements, but the 'Values' list has %d elements- these lists must be the same size", len(o.Semaphores), len(o.Values))
	}

	info := (*C.VkSemaphoreWaitInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_WAIT_INFO
	info.pNext = next
	info.flags = C.VkSemaphoreWaitFlags(o.Flags)

	count := len(o.Semaphores)
	info.semaphoreCount = C.uint32_t(count)
	info.pSemaphores = nil
	info.pValues = nil

	if count > 0 {
		info.pSemaphores = (*C.VkSemaphore)(allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkSemaphore{}))))
		info.pValues = (*C.uint64_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint64_t(0)))))

		semaphoreSlice := unsafe.Slice(info.pSemaphores, count)
		valueSlice := unsafe.Slice(info.pValues, count)

		for i := 0; i < count; i++ {
			if o.Semaphores[i] == nil {
				return nil, errors.Newf("the SemaphoreWaitInfo 'Semaphores' list has a nil semaphore at element %d- all elements must be non-nil", i)
			}

			semaphoreSlice[i] = C.VkSemaphore(unsafe.Pointer(o.Semaphores[i].Handle()))
			valueSlice[i] = C.uint64_t(o.Values[i])
		}
	}

	return preallocatedPointer, nil
}

////

// SemaphoreTypeCreateInfo specifies the type of a newly-created Semaphore
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeCreateInfo.html
type SemaphoreTypeCreateInfo struct {
	// SemaphoreType specifies the type of the Semaphore
	SemaphoreType SemaphoreType
	// InitialValue is the initial payload value if SemaphoreType is SemaphoreTypeTimeline
	InitialValue uint64

	common.NextOptions
}

func (o SemaphoreTypeCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSemaphoreTypeCreateInfo{})))
	}

	info := (*C.VkSemaphoreTypeCreateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_TYPE_CREATE_INFO
	info.pNext = next
	info.semaphoreType = C.VkSemaphoreType(o.SemaphoreType)
	info.initialValue = C.uint64_t(o.InitialValue)

	return preallocatedPointer, nil
}

////

// TimelineSemaphoreSubmitInfo specifies signal and wait values for timeline Semaphore objects
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimelineSemaphoreSubmitInfoKHR.html
type TimelineSemaphoreSubmitInfo struct {
	// WaitSemaphoreValues is a slice of values for the corresponding Semaphore objects in
	// SubmitInfo.WaitSemaphores to wait for
	WaitSemaphoreValues []uint64
	// SignalSemaphoreValues is a slice of values for the corresponding Semaphore objects in
	// SubmitInfo.SignalSemaphores to set when signaled
	SignalSemaphoreValues []uint64

	common.NextOptions
}

func (o TimelineSemaphoreSubmitInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkTimelineSemaphoreSubmitInfo{})))
	}

	info := (*C.VkTimelineSemaphoreSubmitInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_TIMELINE_SEMAPHORE_SUBMIT_INFO
	info.pNext = next

	count := len(o.WaitSemaphoreValues)
	info.waitSemaphoreValueCount = C.uint32_t(count)
	info.pWaitSemaphoreValues = nil

	if count > 0 {
		info.pWaitSemaphoreValues = (*C.uint64_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint64_t(0)))))
		waitSlice := unsafe.Slice(info.pWaitSemaphoreValues, count)

		for i := 0; i < count; i++ {
			waitSlice[i] = C.uint64_t(o.WaitSemaphoreValues[i])
		}
	}

	count = len(o.SignalSemaphoreValues)
	info.signalSemaphoreValueCount = C.uint32_t(count)
	info.pSignalSemaphoreValues = nil

	if count > 0 {
		info.pSignalSemaphoreValues = (*C.uint64_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint64_t(0)))))
		signalSlice := unsafe.Slice(info.pSignalSemaphoreValues, count)

		for i := 0; i < count; i++ {
			signalSlice[i] = C.uint64_t(o.SignalSemaphoreValues[i])
		}
	}

	return preallocatedPointer, nil
}
