package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
)

type BufferImageCopy struct {
	BufferOffset      int
	BufferRowLength   int
	BufferImageHeight int

	ImageSubresource common.ImageSubresourceLayers
	ImageOffset      common.Offset3D
	ImageExtent      common.Extent3D
}

func (c *BufferImageCopy) populate(createInfo *C.VkBufferImageCopy) error {
	if c.BufferImageHeight < 0 {
		return errors.New("provided BufferImageHeight of <0")
	}

	createInfo.bufferOffset = C.VkDeviceSize(c.BufferOffset)
	createInfo.bufferRowLength = C.uint32_t(c.BufferRowLength)
	createInfo.bufferImageHeight = C.uint32_t(c.BufferImageHeight)
	createInfo.imageSubresource.aspectMask = C.VkImageAspectFlags(c.ImageSubresource.AspectMask)
	createInfo.imageSubresource.mipLevel = C.uint32_t(c.ImageSubresource.MipLevel)
	createInfo.imageSubresource.baseArrayLayer = C.uint32_t(c.ImageSubresource.BaseArrayLayer)
	createInfo.imageSubresource.layerCount = C.uint32_t(c.ImageSubresource.LayerCount)
	createInfo.imageOffset.x = C.int32_t(c.ImageOffset.X)
	createInfo.imageOffset.y = C.int32_t(c.ImageOffset.Y)
	createInfo.imageOffset.z = C.int32_t(c.ImageOffset.Z)
	createInfo.imageExtent.width = C.uint32_t(c.ImageExtent.Width)
	createInfo.imageExtent.height = C.uint32_t(c.ImageExtent.Height)
	createInfo.imageExtent.depth = C.uint32_t(c.ImageExtent.Depth)

	return nil
}

func (c *BufferImageCopy) AllocForC(allocator *cgoparam.Allocator) (*C.VkBufferImageCopy, error) {
	createInfo := (*C.VkBufferImageCopy)(allocator.Malloc(C.sizeof_struct_VkBufferImageCopy))
	err := c.populate(createInfo)

	return createInfo, err
}
