package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) CmdCopyBuffer(commandBuffer types.CommandBuffer, srcBuffer types.Buffer, dstBuffer types.Buffer, copyRegions ...core1_0.BufferCopy) error {
	if commandBuffer.Handle() == 0 {
		return errors.New("commandBuffer cannot be uninitialized")
	}
	if srcBuffer.Handle() == 0 {
		return errors.New("srcBuffer cannot be uninitialized")
	}
	if dstBuffer.Handle() == 0 {
		return errors.New("dstBuffer cannot be uninitialized")
	}

	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	copyRegionPtr, err := common.AllocSlice[C.VkBufferCopy, core1_0.BufferCopy](allocator, copyRegions)
	if err != nil {
		return err
	}

	v.Driver.VkCmdCopyBuffer(commandBuffer.Handle(), srcBuffer.Handle(), dstBuffer.Handle(), driver.Uint32(len(copyRegions)), (*driver.VkBufferCopy)(unsafe.Pointer(copyRegionPtr)))
	return nil
}

func (v *Vulkan) CmdCopyImage(commandBuffer types.CommandBuffer, srcImage types.Image, srcImageLayout core1_0.ImageLayout, dstImage types.Image, dstImageLayout core1_0.ImageLayout, regions ...core1_0.ImageCopy) error {
	if commandBuffer.Handle() == 0 {
		return errors.New("commandBuffer cannot be uninitialized")
	}
	if srcImage.Handle() == 0 {
		return errors.New("srcImage cannot be uninitialized")
	}
	if dstImage.Handle() == 0 {
		return errors.New("dstImage cannot be uninitialized")
	}
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	copyRegionCount := len(regions)
	copyRegionUnsafe, err := common.AllocSlice[C.VkImageCopy, core1_0.ImageCopy](allocator, regions)
	if err != nil {
		return err
	}

	v.Driver.VkCmdCopyImage(commandBuffer.Handle(), srcImage.Handle(), driver.VkImageLayout(srcImageLayout), dstImage.Handle(), driver.VkImageLayout(dstImageLayout), driver.Uint32(copyRegionCount), (*driver.VkImageCopy)(unsafe.Pointer(copyRegionUnsafe)))
	return nil
}
