package core

import (
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type Options interface {
	AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error)
	NextInChain() Options
}

type HaveNext struct {
	Next Options
}

func (n *HaveNext) NextInChain() Options {
	return n.Next
}

func AllocOptions(allocator *cgoparam.Allocator, o Options) (unsafe.Pointer, error) {
	next := o.NextInChain()
	var nextPtr unsafe.Pointer
	var err error
	if next != nil {
		nextPtr, err = AllocOptions(allocator, next)
		if err != nil {
			return nil, err
		}
	}

	return o.AllocForC(allocator, nextPtr)
}

func AllocNext(allocator *cgoparam.Allocator, o Options) (unsafe.Pointer, error) {
	next := o.NextInChain()
	if next == nil {
		return nil, nil
	}

	return AllocOptions(allocator, next)
}
