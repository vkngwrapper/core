package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type vulkanBufferView struct {
	driver Driver
	device VkDevice
	handle VkBufferView
}

func (v *vulkanBufferView) Handle() VkBufferView {
	return v.handle
}

func (v *vulkanBufferView) Destroy(callbacks *AllocationCallbacks) {
	v.driver.VkDestroyBufferView(v.device, v.handle, callbacks.Handle())
}

type BufferViewOptions struct {
	Buffer Buffer
	Format common.DataFormat
	Offset int
	Range  int

	common.HaveNext
}

func (o *BufferViewOptions) AllocForC(alloc *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkBufferViewCreateInfo)(alloc.Malloc(C.sizeof_struct_VkBufferViewCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO
	createInfo.pNext = next
	createInfo.flags = 0
	createInfo.buffer = C.VkBuffer(o.Buffer.Handle())
	createInfo.format = C.VkFormat(o.Format)
	createInfo.offset = C.VkDeviceSize(o.Offset)
	createInfo._range = C.VkDeviceSize(o.Range)

	return unsafe.Pointer(createInfo), nil
}
