package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "unsafe"

// ClearValue specifies a clear value
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearValue.html
type ClearValue interface {
	PopulateValueUnion(v unsafe.Pointer)
}

// ClearColorValue specifies a clear color value
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearColorValue.html
type ClearColorValue interface {
	PopulateColorUnion(c unsafe.Pointer)
}

// ClearValueInt32 is a ClearValue and ClearColorValue representing 4 signed 32-bit integer color channels
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

// ClearValueUint32 is a ClearValue and ClearColorValue representing 4 unsigned 32-bit integer
// color channels
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

// ClearValueFloat is a ClearValue and ClearColorValue representing 4 32-bit float color channels
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

// ClearValueDepthStencil is a ClearValue specifying a clear depth stencil value
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearDepthStencilValue.html
type ClearValueDepthStencil struct {
	// Depth is the clear value for the depth aspect of the depth/stencil attachment
	Depth float32
	// Stencil is the clear value of the stencil aspect of the depth/stencil attachment
	Stencil uint32
}

func (s ClearValueDepthStencil) PopulateValueUnion(c unsafe.Pointer) {
	depthStencil := (*C.VkClearDepthStencilValue)(c)
	depthStencil.depth = C.float(s.Depth)
	depthStencil.stencil = C.uint32_t(s.Stencil)
}
