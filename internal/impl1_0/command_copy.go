package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

func (c *VulkanCommandBuffer) CmdCopyBuffer(srcBuffer core1_0.Buffer, dstBuffer core1_0.Buffer, copyRegions []core1_0.BufferCopy) error {
	if srcBuffer == nil {
		panic("srcBuffer cannot be nil")
	}
	if dstBuffer == nil {
		panic("dstBuffer cannot be nil")
	}

	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	copyRegionPtr, err := common.AllocSlice[C.VkBufferCopy, core1_0.BufferCopy](allocator, copyRegions)
	if err != nil {
		return err
	}

	c.DeviceDriver.VkCmdCopyBuffer(c.CommandBufferHandle, srcBuffer.Handle(), dstBuffer.Handle(), driver.Uint32(len(copyRegions)), (*driver.VkBufferCopy)(unsafe.Pointer(copyRegionPtr)))
	return nil
}

func (c *VulkanCommandBuffer) CmdCopyImage(srcImage core1_0.Image, srcImageLayout core1_0.ImageLayout, dstImage core1_0.Image, dstImageLayout core1_0.ImageLayout, regions []core1_0.ImageCopy) error {
	if srcImage == nil {
		panic("srcImage cannot be nil")
	}
	if dstImage == nil {
		panic("dstImage cannot be nil")
	}
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	copyRegionCount := len(regions)
	copyRegionUnsafe, err := common.AllocSlice[C.VkImageCopy, core1_0.ImageCopy](allocator, regions)
	if err != nil {
		return err
	}

	c.DeviceDriver.VkCmdCopyImage(c.CommandBufferHandle, srcImage.Handle(), driver.VkImageLayout(srcImageLayout), dstImage.Handle(), driver.VkImageLayout(dstImageLayout), driver.Uint32(copyRegionCount), (*driver.VkImageCopy)(unsafe.Pointer(copyRegionUnsafe)))
	return nil
}
