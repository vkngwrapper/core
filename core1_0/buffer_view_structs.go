package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
)

// BufferViewCreateFlags is a set of flags reserved for future use
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferViewCreateFlags.html
type BufferViewCreateFlags int32

var bufferViewCreateFlagsMapping = common.NewFlagStringMapping[BufferViewCreateFlags]()

func (f BufferViewCreateFlags) Register(str string) {
	bufferViewCreateFlagsMapping.Register(f, str)
}

func (f BufferViewCreateFlags) String() string {
	return bufferViewCreateFlagsMapping.FlagsToString(f)
}

////

// BufferViewCreateInfo specifies the parameters of a newly-created BufferView object
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferViewCreateInfo.html
type BufferViewCreateInfo struct {
	// Buffer is the Buffer on which the view will be created
	Buffer core.Buffer
	// Flags is reserved for future use
	Flags BufferViewCreateFlags
	// Format describes the format of the data element in the Buffer
	Format Format
	// Offset is the offset in bytes from the base address of the Buffer
	Offset int
	// Range is the size in bytes of the BufferView
	Range int

	common.NextOptions
}

func (o BufferViewCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if !o.Buffer.Initialized() {
		return nil, errors.New("core1_0.BufferViewCreateInfo.Buffer cannot be left unset")
	}

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
