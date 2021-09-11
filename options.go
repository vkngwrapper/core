package core

import (
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type Options interface {
	AllocForC(allocator *cgoparam.Allocator) (unsafe.Pointer, error)
}
