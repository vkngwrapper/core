package render_pass

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/resource"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type FramebufferHandle C.VkFramebuffer
type Framebuffer struct {
	device C.VkDevice
	handle C.VkFramebuffer
}

func CreateFrameBuffer(allocator cgoalloc.Allocator, device *resource.Device, o *FramebufferOptions) (*Framebuffer, core.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	deviceHandle := C.VkDevice(unsafe.Pointer(device.Handle()))
	var framebuffer C.VkFramebuffer

	res := core.Result(C.vkCreateFramebuffer(deviceHandle, (*C.VkFramebufferCreateInfo)(createInfo), nil, &framebuffer))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &Framebuffer{device: deviceHandle, handle: framebuffer}, res, nil
}

func (b *Framebuffer) Handle() FramebufferHandle {
	return FramebufferHandle(b.handle)
}

func (b *Framebuffer) Destroy() {
	C.vkDestroyFramebuffer(b.device, b.handle, nil)
}
