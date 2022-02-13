package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/iface"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

func (c *VulkanCommandBuffer) CmdCopyBuffer(srcBuffer iface.Buffer, dstBuffer iface.Buffer, copyRegions []core1_0.BufferCopy) error {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	copyRegionPtr, err := core.AllocSlice[C.VkCommandBuffer, core1_0.BufferCopy](allocator, copyRegions)
	if err != nil {
		return err
	}

	c.driver.VkCmdCopyBuffer(c.handle, srcBuffer.Handle(), dstBuffer.Handle(), driver.Uint32(len(copyRegions)), (*driver.VkBufferCopy)(unsafe.Pointer(copyRegionPtr)))
	return nil
}

func (c *VulkanCommandBuffer) CmdCopyImage(srcImage iface.Image, srcImageLayout common.ImageLayout, dstImage iface.Image, dstImageLayout common.ImageLayout, regions []core1_0.ImageCopy) error {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	copyRegionCount := len(regions)
	copyRegionUnsafe, err := core.AllocSlice[C.VkImageCopy, core1_0.ImageCopy](allocator, regions)
	if err != nil {
		return err
	}

	c.driver.VkCmdCopyImage(c.handle, srcImage.Handle(), driver.VkImageLayout(srcImageLayout), dstImage.Handle(), driver.VkImageLayout(dstImageLayout), driver.Uint32(copyRegionCount), (*driver.VkImageCopy)(unsafe.Pointer(copyRegionUnsafe)))
	return nil
}
