package common

import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type CAllocatable interface {
	PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error)
}

type Options interface {
	PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error)
	NextInChain() Options
}

type HaveNext struct {
	Next Options
}

func (n HaveNext) NextInChain() Options {
	return n.Next
}

func AllocOptions(allocator *cgoparam.Allocator, o Options, preallocatedPointer ...unsafe.Pointer) (unsafe.Pointer, error) {
	nextPtr, err := allocNext(allocator, o)
	if err != nil {
		return nil, err
	}

	var preallocatedPointerToPass unsafe.Pointer
	if len(preallocatedPointer) > 0 {
		preallocatedPointerToPass = preallocatedPointer[0]
	}

	return o.PopulateCPointer(allocator, preallocatedPointerToPass, nextPtr)
}

func AllocOptionSlice[T any, O Options](allocator *cgoparam.Allocator, o []O) (*T, error) {
	optionCount := len(o)
	optionPtr := (*T)(allocator.Malloc(optionCount * int(unsafe.Sizeof([1]T{}))))
	optionSlice := unsafe.Slice(optionPtr, optionCount)
	for i := 0; i < optionCount; i++ {
		next, err := allocNext(allocator, o[i])
		if err != nil {
			return nil, err
		}

		_, err = o[i].PopulateCPointer(allocator, unsafe.Pointer(&optionSlice[i]), next)
		if err != nil {
			return nil, err
		}
	}

	return optionPtr, nil
}

func allocNext(allocator *cgoparam.Allocator, o Options) (unsafe.Pointer, error) {
	next := o.NextInChain()
	if next == nil {
		return nil, nil
	}

	return AllocOptions(allocator, next)
}

func AllocSlice[T any, A CAllocatable](allocator *cgoparam.Allocator, a []A) (*T, error) {
	optionCount := len(a)
	optionPtr := (*T)(allocator.Malloc(optionCount * int(unsafe.Sizeof([1]T{}))))
	optionSlice := unsafe.Slice(optionPtr, optionCount)
	for i := 0; i < optionCount; i++ {
		_, err := a[i].PopulateCPointer(allocator, unsafe.Pointer(&optionSlice[i]))
		if err != nil {
			return nil, err
		}
	}

	return optionPtr, nil
}
