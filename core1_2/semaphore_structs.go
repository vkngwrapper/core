package core1_2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type SemaphoreType int32

var semaphoreTypeMapping = make(map[SemaphoreType]string)

func (e SemaphoreType) Register(str string) {
	semaphoreTypeMapping[e] = str
}

func (e SemaphoreType) String() string {
	return semaphoreTypeMapping[e]
}

////

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
	SemaphoreTypeBinary   SemaphoreType = C.VK_SEMAPHORE_TYPE_BINARY
	SemaphoreTypeTimeline SemaphoreType = C.VK_SEMAPHORE_TYPE_TIMELINE

	SemaphoreWaitAny SemaphoreWaitFlags = C.VK_SEMAPHORE_WAIT_ANY_BIT
)

func init() {
	SemaphoreTypeBinary.Register("Binary")
	SemaphoreTypeTimeline.Register("Timeline")

	SemaphoreWaitAny.Register("Any")
}

////

type SemaphoreSignalOptions struct {
	Semaphore core1_0.Semaphore
	Value     uint64

	common.NextOptions
}

func (o SemaphoreSignalOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSemaphoreSignalInfo{})))
	}

	if o.Semaphore == nil {
		return nil, errors.New("the 'Semaphore' field of SemaphoreSignalOptions must be non-nil")
	}

	info := (*C.VkSemaphoreSignalInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_SIGNAL_INFO
	info.pNext = next
	info.semaphore = C.VkSemaphore(unsafe.Pointer(o.Semaphore.Handle()))
	info.value = C.uint64_t(o.Value)

	return preallocatedPointer, nil
}

////

type SemaphoreWaitOptions struct {
	Flags      SemaphoreWaitFlags
	Semaphores []core1_0.Semaphore
	Values     []uint64

	common.NextOptions
}

func (o SemaphoreWaitOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSemaphoreWaitInfo{})))
	}

	if len(o.Semaphores) != len(o.Values) {
		return nil, errors.Newf("the SemaphoreWaitOptions 'Semaphores' list has %d elements, but the 'Values' list has %d elements- these lists must be the same size", len(o.Semaphores), len(o.Values))
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
				return nil, errors.Newf("the SemaphoreWaitOptions 'Semaphores' list has a nil semaphore at element %d- all elements must be non-nil", i)
			}

			semaphoreSlice[i] = C.VkSemaphore(unsafe.Pointer(o.Semaphores[i].Handle()))
			valueSlice[i] = C.uint64_t(o.Values[i])
		}
	}

	return preallocatedPointer, nil
}

////

type SemaphoreTypeCreateOptions struct {
	SemaphoreType SemaphoreType
	InitialValue  uint64

	common.NextOptions
}

func (o SemaphoreTypeCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

type TimelineSemaphoreSubmitOptions struct {
	WaitSemaphoreValues   []uint64
	SignalSemaphoreValues []uint64

	common.NextOptions
}

func (o TimelineSemaphoreSubmitOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
