package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

func (c *VulkanCommandBuffer) CmdCopyBuffer(srcBuffer core1_0.Buffer, dstBuffer core1_0.Buffer, copyRegions []core1_0.BufferCopy) error {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	copyRegionPtr, err := common.AllocSlice[C.VkBufferCopy, core1_0.BufferCopy](allocator, copyRegions)
	if err != nil {
		return err
	}

	c.DeviceDriver.VkCmdCopyBuffer(c.CommandBufferHandle, srcBuffer.Handle(), dstBuffer.Handle(), driver.Uint32(len(copyRegions)), (*driver.VkBufferCopy)(unsafe.Pointer(copyRegionPtr)))
	return nil
}

func (c *VulkanCommandBuffer) CmdCopyImage(srcImage core1_0.Image, srcImageLayout common.ImageLayout, dstImage core1_0.Image, dstImageLayout common.ImageLayout, regions []core1_0.ImageCopy) error {
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
