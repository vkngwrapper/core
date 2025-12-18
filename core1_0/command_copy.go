package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
)

func (c *VulkanCommandBuffer) CmdCopyBuffer(srcBuffer Buffer, dstBuffer Buffer, copyRegions []BufferCopy) error {
	if srcBuffer == nil {
		panic("srcBuffer cannot be nil")
	}
	if dstBuffer == nil {
		panic("dstBuffer cannot be nil")
	}

	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	copyRegionPtr, err := common.AllocSlice[C.VkBufferCopy, BufferCopy](allocator, copyRegions)
	if err != nil {
		return err
	}

	c.deviceDriver.VkCmdCopyBuffer(c.commandBufferHandle, srcBuffer.Handle(), dstBuffer.Handle(), driver.Uint32(len(copyRegions)), (*driver.VkBufferCopy)(unsafe.Pointer(copyRegionPtr)))
	return nil
}

func (c *VulkanCommandBuffer) CmdCopyImage(srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regions []ImageCopy) error {
	if srcImage == nil {
		panic("srcImage cannot be nil")
	}
	if dstImage == nil {
		panic("dstImage cannot be nil")
	}
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	copyRegionCount := len(regions)
	copyRegionUnsafe, err := common.AllocSlice[C.VkImageCopy, ImageCopy](allocator, regions)
	if err != nil {
		return err
	}

	c.deviceDriver.VkCmdCopyImage(c.commandBufferHandle, srcImage.Handle(), driver.VkImageLayout(srcImageLayout), dstImage.Handle(), driver.VkImageLayout(dstImageLayout), driver.Uint32(copyRegionCount), (*driver.VkImageCopy)(unsafe.Pointer(copyRegionUnsafe)))
	return nil
}
