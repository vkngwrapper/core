package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

type BufferViewCreateFlags int32

var bufferViewCreateFlagsMapping = common.NewFlagStringMapping[BufferViewCreateFlags]()

func (f BufferViewCreateFlags) Register(str string) {
	bufferViewCreateFlagsMapping.Register(f, str)
}

func (f BufferViewCreateFlags) String() string {
	return bufferViewCreateFlagsMapping.FlagsToString(f)
}

////

type BufferViewCreateInfo struct {
	Buffer Buffer
	Flags  BufferViewCreateFlags
	Format Format
	Offset int
	Range  int

	common.NextOptions
}

func (o BufferViewCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkBufferViewCreateInfo)
	}
	createInfo := (*C.VkBufferViewCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO
	createInfo.pNext = next
	createInfo.flags = C.VkBufferViewCreateFlags(o.Flags)
	createInfo.buffer = C.VkBuffer(unsafe.Pointer(o.Buffer.Handle()))
	createInfo.format = C.VkFormat(o.Format)
	createInfo.offset = C.VkDeviceSize(o.Offset)
	createInfo._range = C.VkDeviceSize(o.Range)

	return unsafe.Pointer(createInfo), nil
}
