package common

import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

// CAllocatable is implemented by vkngwrapper structures that are not chainable (do not have
// Next* fields) but need to be allocatable with AllocSlice or PopulateCPointer. If you are
// not contributing to vkngwrapper or writing a Vulkan extension wrapper, you do not need
// to understand this type.
type CAllocatable interface {
	PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error)
}

// Options is implemented by vkngwrapper structures that pass input into Vulkan and are
// chainable via NextOptions. If you are not contributing to vkngwrapper or writing a Vulkan
// extension wrapper, you do not need to understand this type.
type Options interface {
	PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error)

	NextOptionsInChain() Options
}

// OutData is implemented by vkngwrapper structures that receive output from Vulkan and
// are chainable via NextOutData. If you are not contributing to vkngwrapper or writing a
// Vulkan extension wrapper, you do not need to understand this type.
type OutData interface {
	PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error)
	PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error)

	NextOutDataInChain() OutData
}

// NextOptions is embedded by all vkngwrapper structures that pass input into Vulkan and
// are chainable. If you are wondering which structures can chain onto which other structures,
// all structures which embed NextOptions can be chained onto any other structures which embeds
// NextOptions, without exploding. However, you will need to refer to the Vulkan spec to see
// which chains actually do something worthwhile, and the validation layer may take issue
// with chains that do not match the spec.
type NextOptions struct {
	Next Options
}

func (n NextOptions) NextOptionsInChain() Options {
	return n.Next
}

// NextOutData is embedded by all vkngwrapper structures that receive output from Vulkan and
// are chainable. If you are wondering which structures can chain onto which other structures,
// all structures which embed NextOutData can be chained onto any other structures which embeds
// NextOutData, without exploding. However, you will need to refer to the Vulkan spec to see
// which chains actually do something worthwhile, and the validation layer may take issue
// with chains that do not match the spec.
type NextOutData struct {
	Next OutData
}

func (n NextOutData) NextOutDataInChain() OutData {
	return n.Next
}

// AllocOutDataHeader is used to receive a C-allocated pointer for a chain of OutData structs.
// If you are not contributing to vkngwrapper or writing a Vulkan extension wrapper, you do
// not need to understand this method.
func AllocOutDataHeader(allocator *cgoparam.Allocator, o OutData, preallocatedPointer ...unsafe.Pointer) (unsafe.Pointer, error) {
	nextPtr, err := allocNextOutData(allocator, o)
	if err != nil {
		return nil, err
	}

	var preallocatedPointerToPass unsafe.Pointer
	if len(preallocatedPointer) > 0 {
		preallocatedPointerToPass = preallocatedPointer[0]
	}

	return o.PopulateHeader(allocator, preallocatedPointerToPass, nextPtr)
}

// AllocOutDataHeaderSlice is used to receive a C-allocated array for several chains of OutData
// structs. If you are not contributing to vkngwrapper or writing a Vulkan extension wrapper,
// you do not need to understand this method.
func AllocOutDataHeaderSlice[T any, O OutData](allocator *cgoparam.Allocator, o []O) (*T, error) {
	optionCount := len(o)
	optionPtr := (*T)(allocator.Malloc(optionCount * int(unsafe.Sizeof([1]T{}))))
	optionSlice := unsafe.Slice(optionPtr, optionCount)
	for i := 0; i < optionCount; i++ {
		next, err := allocNextOutData(allocator, o[i])
		if err != nil {
			return nil, err
		}

		_, err = o[i].PopulateHeader(allocator, unsafe.Pointer(&optionSlice[i]), next)
		if err != nil {
			return nil, err
		}
	}

	return optionPtr, nil
}

func allocNextOutData(allocator *cgoparam.Allocator, o OutData) (unsafe.Pointer, error) {
	next := o.NextOutDataInChain()
	if next == nil {
		return nil, nil
	}

	return AllocOutDataHeader(allocator, next)
}

// AllocOptions is used to receive a C-allocated pointer for a chain of Options structs.
// If you are not contributing to vkngwrapper or writing a Vulkan extension wrapper, you do
// not need to understand this method.
func AllocOptions(allocator *cgoparam.Allocator, o Options, preallocatedPointer ...unsafe.Pointer) (unsafe.Pointer, error) {
	nextPtr, err := allocNextOptions(allocator, o)
	if err != nil {
		return nil, err
	}

	var preallocatedPointerToPass unsafe.Pointer
	if len(preallocatedPointer) > 0 {
		preallocatedPointerToPass = preallocatedPointer[0]
	}

	return o.PopulateCPointer(allocator, preallocatedPointerToPass, nextPtr)
}

// AllocOptionSlice is used to receive a C-allocated array for several chains of Options
// structs. If you are not contributing to vkngwrapper or writing a Vulkan extension wrapper,
// you do not need to understand this method.
func AllocOptionSlice[T any, O Options](allocator *cgoparam.Allocator, o []O) (*T, error) {
	optionCount := len(o)
	optionPtr := (*T)(allocator.Malloc(optionCount * int(unsafe.Sizeof([1]T{}))))
	optionSlice := unsafe.Slice(optionPtr, optionCount)
	for i := 0; i < optionCount; i++ {
		next, err := allocNextOptions(allocator, o[i])
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

func allocNextOptions(allocator *cgoparam.Allocator, o Options) (unsafe.Pointer, error) {
	next := o.NextOptionsInChain()
	if next == nil {
		return nil, nil
	}

	return AllocOptions(allocator, next)
}

// AllocSlice is used to receive a C-allocated array for several CAllocator objects.
// If you are not contributing to vkngwrapper or writing a Vulkan extension wrapper,
// you do not need to understand this method.
func AllocSlice[T any, A CAllocatable](allocator *cgoparam.Allocator, a []A) (*T, error) {
	optionCount := len(a)
	optionSize := unsafe.Sizeof([1]T{})
	optionPtr := allocator.Malloc(optionCount * int(optionSize))
	optionIterPtr := optionPtr

	for i := 0; i < optionCount; i++ {
		_, err := a[i].PopulateCPointer(allocator, optionIterPtr)
		if err != nil {
			return nil, err
		}

		optionIterPtr = unsafe.Add(optionIterPtr, optionSize)
	}

	return (*T)(optionPtr), nil
}

// PopulateOutData populates fields in a chain of OutData objects from the same chain in a C pointer.
// The C pointer should have been previously allocated with AllocOutDataHeader and populated
// with a call to a Vulkan command. If you are not contributing to vkngwrapper or writing a
// Vulkan extension wrapper, you do not need to understand this method.
func PopulateOutData(o OutData, cPointer unsafe.Pointer, helpers ...any) error {
	next, err := o.PopulateOutData(cPointer, helpers...)
	if err != nil {
		return err
	}

	nextOptions := o.NextOutDataInChain()
	if nextOptions != nil {
		return PopulateOutData(nextOptions, next, helpers...)
	}

	return nil
}

// PopulateOutDataSlice populates fields in a slice of OutData chains from an array of chains
// in a C pointer. The C pointer should have been previously allocated with AllocOutDataHeaderSlice
// and populated with a call to a Vulkan command. If you are not contributing to vkngwrapper
// or writing a Vulkan extension wrapper, you do not need to understand this method.
func PopulateOutDataSlice[T any, O OutData](o []O, cSlicePointer unsafe.Pointer, helpers ...any) error {
	cElementSize := unsafe.Sizeof([1]T{})

	for i := 0; i < len(o); i++ {
		err := PopulateOutData(o[i], cSlicePointer, helpers...)
		if err != nil {
			return err
		}

		cSlicePointer = unsafe.Add(cSlicePointer, cElementSize)
	}

	return nil
}

// OfType locates the first value in an untyped array which matches the parameter type and
// returns it, if any. Otherwise, it will return the default value of the parameter type.
// an `ok` return value is included, for convenience. If you are not contributing to
// vkngwrapper or writing a Vulkan extension wrapper, you do not need to understand this
// method.
func OfType[T any](values []any) (T, bool) {
	for _, val := range values {
		typed, ok := val.(T)
		if ok {
			return typed, true
		}
	}

	var zero T
	return zero, false
}
