package common

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type Options interface {
	AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error)
	NextInChain() Options
}

type mustBeRootOptions interface {
	MustBeRootOptions() bool
}

type HaveNext struct {
	Next Options
}

func (n *HaveNext) NextInChain() Options {
	return n.Next
}

func AllocOptions(allocator *cgoparam.Allocator, o Options) (unsafe.Pointer, error) {
	nextPtr, err := AllocNext(allocator, o)
	if err != nil {
		return nil, err
	}

	return o.AllocForC(allocator, nextPtr)
}

func AllocNext(allocator *cgoparam.Allocator, o Options) (unsafe.Pointer, error) {
	next := o.NextInChain()
	if next == nil {
		return nil, nil
	}

	mustBeRoot, hasMustBeRoot := next.(mustBeRootOptions)
	if hasMustBeRoot && mustBeRoot.MustBeRootOptions() {
		return nil, errors.Newf("attempted to use %v as chained options, but it may only be used as the root in a chain", next)
	}

	return AllocOptions(allocator, next)
}
