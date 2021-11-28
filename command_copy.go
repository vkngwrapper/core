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

type BufferCopy struct {
	SrcOffset int
	DstOffset int
	Size      int
}

func (c *vulkanCommandBuffer) CmdCopyBuffer(srcBuffer Buffer, dstBuffer Buffer, copyRegions []BufferCopy) error {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	copyRegionCount := len(copyRegions)
	copyRegionUnsafe := allocator.Malloc(copyRegionCount * C.sizeof_struct_VkBufferCopy)
	copyRegionPtr := (*C.VkBufferCopy)(copyRegionUnsafe)
	copyRegionSlice := ([]C.VkBufferCopy)(unsafe.Slice(copyRegionPtr, copyRegionCount))

	for i := 0; i < copyRegionCount; i++ {
		copyRegionSlice[i].srcOffset = C.VkDeviceSize(copyRegions[i].SrcOffset)
		copyRegionSlice[i].dstOffset = C.VkDeviceSize(copyRegions[i].DstOffset)
		copyRegionSlice[i].size = C.VkDeviceSize(copyRegions[i].Size)
	}

	c.driver.VkCmdCopyBuffer(c.handle, srcBuffer.Handle(), dstBuffer.Handle(), Uint32(copyRegionCount), (*VkBufferCopy)(copyRegionUnsafe))
	return nil
}

type ImageCopy struct {
	SrcSubresource common.ImageSubresourceLayers
	SrcOffset      common.Offset3D
	DstSubresource common.ImageSubresourceLayers
	DstOffset      common.Offset3D
	Extent         common.Extent3D
}

func (c *vulkanCommandBuffer) CmdCopyImage(srcImage Image, srcImageLayout common.ImageLayout, dstImage Image, dstImageLayout common.ImageLayout, regions []ImageCopy) error {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	copyRegionCount := len(regions)
	copyRegionUnsafe := allocator.Malloc(copyRegionCount * C.sizeof_struct_VkImageCopy)
	copyRegionPtr := (*C.VkImageCopy)(copyRegionUnsafe)
	copyRegionSlice := ([]C.VkImageCopy)(unsafe.Slice(copyRegionPtr, copyRegionCount))

	for i := 0; i < copyRegionCount; i++ {
		copyRegionSlice[i].srcSubresource.aspectMask = C.VkImageAspectFlags(regions[i].SrcSubresource.AspectMask)
		copyRegionSlice[i].srcSubresource.mipLevel = C.uint32_t(regions[i].SrcSubresource.MipLevel)
		copyRegionSlice[i].srcSubresource.baseArrayLayer = C.uint32_t(regions[i].SrcSubresource.BaseArrayLayer)
		copyRegionSlice[i].srcSubresource.layerCount = C.uint32_t(regions[i].SrcSubresource.LayerCount)

		copyRegionSlice[i].dstSubresource.aspectMask = C.VkImageAspectFlags(regions[i].DstSubresource.AspectMask)
		copyRegionSlice[i].dstSubresource.mipLevel = C.uint32_t(regions[i].DstSubresource.MipLevel)
		copyRegionSlice[i].dstSubresource.baseArrayLayer = C.uint32_t(regions[i].DstSubresource.BaseArrayLayer)
		copyRegionSlice[i].dstSubresource.layerCount = C.uint32_t(regions[i].DstSubresource.LayerCount)

		copyRegionSlice[i].srcOffset.x = C.int32_t(regions[i].SrcOffset.X)
		copyRegionSlice[i].srcOffset.y = C.int32_t(regions[i].SrcOffset.Y)
		copyRegionSlice[i].srcOffset.z = C.int32_t(regions[i].SrcOffset.Z)

		copyRegionSlice[i].dstOffset.x = C.int32_t(regions[i].DstOffset.X)
		copyRegionSlice[i].dstOffset.y = C.int32_t(regions[i].DstOffset.Y)
		copyRegionSlice[i].dstOffset.z = C.int32_t(regions[i].DstOffset.Z)

		copyRegionSlice[i].extent.width = C.uint32_t(regions[i].Extent.Width)
		copyRegionSlice[i].extent.height = C.uint32_t(regions[i].Extent.Height)
		copyRegionSlice[i].extent.depth = C.uint32_t(regions[i].Extent.Depth)
	}

	c.driver.VkCmdCopyImage(c.handle, srcImage.Handle(), VkImageLayout(srcImageLayout), dstImage.Handle(), VkImageLayout(dstImageLayout), Uint32(copyRegionCount), (*VkImageCopy)(copyRegionUnsafe))
	return nil
}
