package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

type ImageViewUsageCreateInfo struct {
	Usage core1_0.ImageUsageFlags

	common.NextOptions
}

func (o ImageViewUsageCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageViewUsageCreateInfo{})))
	}

	createInfo := (*C.VkImageViewUsageCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO
	createInfo.pNext = next
	createInfo.usage = C.VkImageUsageFlags(o.Usage)

	return preallocatedPointer, nil
}
