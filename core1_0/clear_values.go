package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "unsafe"

type ClearValue interface {
	PopulateValueUnion(v unsafe.Pointer)
}

type ClearColorValue interface {
	PopulateColorUnion(c unsafe.Pointer)
}

type ClearValueInt32 [4]int32

func (v ClearValueInt32) PopulateValueUnion(c unsafe.Pointer) {
	colorInt32 := unsafe.Slice((*C.int32_t)(c), 4)
	for i := 0; i < 4; i++ {
		colorInt32[i] = C.int32_t(v[i])
	}
}

func (v ClearValueInt32) PopulateColorUnion(c unsafe.Pointer) {
	colorInt32 := unsafe.Slice((*C.int32_t)(c), 4)
	for i := 0; i < 4; i++ {
		colorInt32[i] = C.int32_t(v[i])
	}
}

type ClearValueUint32 [4]uint32

func (v ClearValueUint32) PopulateValueUnion(c unsafe.Pointer) {
	colorUint32 := unsafe.Slice((*C.uint32_t)(c), 4)
	for i := 0; i < 4; i++ {
		colorUint32[i] = C.uint32_t(v[i])
	}
}

func (v ClearValueUint32) PopulateColorUnion(c unsafe.Pointer) {
	colorUint32 := unsafe.Slice((*C.uint32_t)(c), 4)
	for i := 0; i < 4; i++ {
		colorUint32[i] = C.uint32_t(v[i])
	}
}

type ClearValueFloat [4]float32

func (v ClearValueFloat) PopulateValueUnion(c unsafe.Pointer) {
	colorFloat := unsafe.Slice((*C.float)(c), 4)
	for i := 0; i < 4; i++ {
		colorFloat[i] = C.float(v[i])
	}
}

func (v ClearValueFloat) PopulateColorUnion(c unsafe.Pointer) {
	colorFloat := unsafe.Slice((*C.float)(c), 4)
	for i := 0; i < 4; i++ {
		colorFloat[i] = C.float(v[i])
	}
}

type ClearValueDepthStencil struct {
	Depth   float32
	Stencil uint32
}

func (s ClearValueDepthStencil) PopulateValueUnion(c unsafe.Pointer) {
	depthStencil := (*C.VkClearDepthStencilValue)(c)
	depthStencil.depth = C.float(s.Depth)
	depthStencil.stencil = C.uint32_t(s.Stencil)
}
