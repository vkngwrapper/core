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
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *DeviceVulkanDriver) CmdCopyBuffer(commandBuffer types.CommandBuffer, srcBuffer types.Buffer, dstBuffer types.Buffer, copyRegions ...core1_0.BufferCopy) error {
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

	v.LoaderObj.VkCmdCopyBuffer(commandBuffer.Handle(), srcBuffer.Handle(), dstBuffer.Handle(), loader.Uint32(len(copyRegions)), (*loader.VkBufferCopy)(unsafe.Pointer(copyRegionPtr)))
	return nil
}

func (v *DeviceVulkanDriver) CmdCopyImage(commandBuffer types.CommandBuffer, srcImage types.Image, srcImageLayout core1_0.ImageLayout, dstImage types.Image, dstImageLayout core1_0.ImageLayout, regions ...core1_0.ImageCopy) error {
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

	v.LoaderObj.VkCmdCopyImage(commandBuffer.Handle(), srcImage.Handle(), loader.VkImageLayout(srcImageLayout), dstImage.Handle(), loader.VkImageLayout(dstImageLayout), loader.Uint32(copyRegionCount), (*loader.VkImageCopy)(unsafe.Pointer(copyRegionUnsafe)))
	return nil
}
