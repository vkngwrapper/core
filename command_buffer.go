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
	"unsafe"
)

type vulkanCommandBuffer struct {
	driver Driver
	device VkDevice
	pool   VkCommandPool
	handle VkCommandBuffer
}

func (c *vulkanCommandBuffer) Handle() VkCommandBuffer {
	return c.handle
}

func (c *vulkanCommandBuffer) Begin(o *BeginOptions) (VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return VKErrorUnknown, err
	}

	return c.driver.VkBeginCommandBuffer(c.handle, (*VkCommandBufferBeginInfo)(createInfo))
}

func (c *vulkanCommandBuffer) End() (VkResult, error) {
	return c.driver.VkEndCommandBuffer(c.handle)
}

func (c *vulkanCommandBuffer) CmdBeginRenderPass(contents SubpassContents, o *RenderPassBeginOptions) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	return c.driver.VkCmdBeginRenderPass(c.handle, (*VkRenderPassBeginInfo)(createInfo), VkSubpassContents(contents))
}

func (c *vulkanCommandBuffer) CmdEndRenderPass() error {
	return c.driver.VkCmdEndRenderPass(c.handle)
}

func (c *vulkanCommandBuffer) CmdBindPipeline(bindPoint common.PipelineBindPoint, pipeline Pipeline) error {
	return c.driver.VkCmdBindPipeline(c.handle, VkPipelineBindPoint(bindPoint), pipeline.Handle())
}

func (c *vulkanCommandBuffer) CmdDraw(vertexCount, instanceCount int, firstVertex, firstInstance uint32) error {
	return c.driver.VkCmdDraw(c.handle, Uint32(vertexCount), Uint32(instanceCount), Uint32(firstVertex), Uint32(firstInstance))
}

func (c *vulkanCommandBuffer) CmdDrawIndexed(indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32) error {
	return c.driver.VkCmdDrawIndexed(c.handle, Uint32(indexCount), Uint32(instanceCount), Uint32(firstIndex), Int32(vertexOffset), Uint32(firstInstance))
}

func (c *vulkanCommandBuffer) CmdBindVertexBuffers(firstBinding uint32, buffers []Buffer, bufferOffsets []int) error {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	bufferCount := len(buffers)

	bufferArrayUnsafe := allocator.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkBuffer{})))
	offsetArrayUnsafe := allocator.Malloc(bufferCount * int(unsafe.Sizeof(C.VkDeviceSize(0))))

	bufferArrayPtr := (*VkBuffer)(bufferArrayUnsafe)
	offsetArrayPtr := (*VkDeviceSize)(offsetArrayUnsafe)

	bufferArraySlice := ([]VkBuffer)(unsafe.Slice(bufferArrayPtr, bufferCount))
	offsetArraySlice := ([]VkDeviceSize)(unsafe.Slice(offsetArrayPtr, bufferCount))

	for i := 0; i < bufferCount; i++ {
		bufferArraySlice[i] = buffers[i].Handle()
		offsetArraySlice[i] = VkDeviceSize(bufferOffsets[i])
	}

	return c.driver.VkCmdBindVertexBuffers(c.handle, Uint32(firstBinding), Uint32(bufferCount), bufferArrayPtr, offsetArrayPtr)
}

func (c *vulkanCommandBuffer) CmdBindIndexBuffer(buffer Buffer, offset int, indexType common.IndexType) error {
	return c.driver.VkCmdBindIndexBuffer(c.handle, buffer.Handle(), VkDeviceSize(offset), VkIndexType(indexType))
}

func (c *vulkanCommandBuffer) CmdBindDescriptorSets(bindPoint common.PipelineBindPoint, layout PipelineLayout, firstSet int, sets []DescriptorSet, dynamicOffsets []int) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	setCount := len(sets)
	dynamicOffsetCount := len(dynamicOffsets)

	var setPtr unsafe.Pointer
	var dynamicOffsetPtr unsafe.Pointer

	if setCount > 0 {
		setPtr = arena.Malloc(setCount * int(unsafe.Sizeof([1]C.VkDescriptorSet{})))
		setSlice := ([]C.VkDescriptorSet)(unsafe.Slice((*C.VkDescriptorSet)(setPtr), setCount))
		for i := 0; i < setCount; i++ {
			setSlice[i] = (C.VkDescriptorSet)(unsafe.Pointer(sets[i].Handle()))
		}
	}

	if dynamicOffsetCount > 0 {
		dynamicOffsetPtr = arena.Malloc(dynamicOffsetCount * int(unsafe.Sizeof(C.uint32_t(0))))
		dynamicOffsetSlice := ([]C.uint32_t)(unsafe.Slice((*C.uint32_t)(dynamicOffsetPtr), dynamicOffsetCount))

		for i := 0; i < dynamicOffsetCount; i++ {
			dynamicOffsetSlice[i] = (C.uint32_t)(dynamicOffsets[i])
		}
	}

	return c.driver.VkCmdBindDescriptorSets(c.handle,
		VkPipelineBindPoint(bindPoint),
		layout.Handle(),
		Uint32(firstSet),
		Uint32(setCount),
		(*VkDescriptorSet)(setPtr),
		Uint32(dynamicOffsetCount),
		(*Uint32)(dynamicOffsetPtr))
}

func (c *vulkanCommandBuffer) CmdPipelineBarrier(srcStageMask, dstStageMask common.PipelineStages, dependencies common.DependencyFlags, memoryBarriers []*MemoryBarrierOptions, bufferMemoryBarriers []*BufferMemoryBarrierOptions, imageMemoryBarriers []*ImageMemoryBarrierOptions) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	barrierCount := len(memoryBarriers)
	bufferBarrierCount := len(bufferMemoryBarriers)
	imageBarrierCount := len(imageMemoryBarriers)

	var barrierPtr *C.VkMemoryBarrier
	var bufferBarrierPtr *C.VkBufferMemoryBarrier
	var imageBarrierPtr *C.VkImageMemoryBarrier

	if barrierCount > 0 {
		barrierPtr = (*C.VkMemoryBarrier)(arena.Malloc(barrierCount * C.sizeof_struct_VkMemoryBarrier))
		barrierSlice := ([]C.VkMemoryBarrier)(unsafe.Slice(barrierPtr, barrierCount))

		for i := 0; i < barrierCount; i++ {
			next, err := common.AllocNext(arena, memoryBarriers[i])
			if err != nil {
				return err
			}

			err = memoryBarriers[i].populate(&barrierSlice[i], next)
			if err != nil {
				return err
			}
		}
	}

	if bufferBarrierCount > 0 {
		bufferBarrierPtr = (*C.VkBufferMemoryBarrier)(arena.Malloc(bufferBarrierCount * C.sizeof_struct_VkBufferMemoryBarrier))
		bufferBarrierSlice := ([]C.VkBufferMemoryBarrier)(unsafe.Slice(bufferBarrierPtr, bufferBarrierCount))

		for i := 0; i < bufferBarrierCount; i++ {
			next, err := common.AllocNext(arena, bufferMemoryBarriers[i])
			if err != nil {
				return err
			}

			err = bufferMemoryBarriers[i].populate(&bufferBarrierSlice[i], next)
			if err != nil {
				return err
			}
		}
	}

	if imageBarrierCount > 0 {
		imageBarrierPtr = (*C.VkImageMemoryBarrier)(arena.Malloc(imageBarrierCount * C.sizeof_struct_VkImageMemoryBarrier))
		imageBarrierSlice := ([]C.VkImageMemoryBarrier)(unsafe.Slice(imageBarrierPtr, imageBarrierCount))

		for i := 0; i < imageBarrierCount; i++ {
			next, err := common.AllocNext(arena, imageMemoryBarriers[i])
			if err != nil {
				return err
			}

			err = imageMemoryBarriers[i].populate(&imageBarrierSlice[i], next)
			if err != nil {
				return err
			}
		}
	}

	return c.driver.VkCmdPipelineBarrier(c.handle, VkPipelineStageFlags(srcStageMask), VkPipelineStageFlags(dstStageMask), VkDependencyFlags(dependencies), Uint32(barrierCount), (*VkMemoryBarrier)(barrierPtr), Uint32(bufferBarrierCount), (*VkBufferMemoryBarrier)(bufferBarrierPtr), Uint32(imageBarrierCount), (*VkImageMemoryBarrier)(imageBarrierPtr))
}

func (c *vulkanCommandBuffer) CmdCopyBufferToImage(buffer Buffer, image Image, layout common.ImageLayout, regions []*BufferImageCopy) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	regionCount := len(regions)
	var regionPtr *C.VkBufferImageCopy

	if regionCount > 0 {
		regionPtr = (*C.VkBufferImageCopy)(arena.Malloc(regionCount * C.sizeof_struct_VkBufferImageCopy))
		regionSlice := ([]C.VkBufferImageCopy)(unsafe.Slice(regionPtr, regionCount))

		for i := 0; i < regionCount; i++ {
			err := regions[i].populate(&regionSlice[i])
			if err != nil {
				return err
			}
		}
	}

	return c.driver.VkCmdCopyBufferToImage(c.handle, buffer.Handle(), image.Handle(), VkImageLayout(layout), Uint32(regionCount), (*VkBufferImageCopy)(regionPtr))
}

func (c *vulkanCommandBuffer) CmdBlitImage(sourceImage Image, sourceImageLayout common.ImageLayout, destinationImage Image, destinationImageLayout common.ImageLayout, regions []*ImageBlit, filter common.Filter) error {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	regionCount := len(regions)
	regionPtr := (*C.VkImageBlit)(allocator.Malloc(regionCount * C.sizeof_struct_VkImageBlit))
	regionSlice := ([]C.VkImageBlit)(unsafe.Slice(regionPtr, regionCount))

	for i := range regionSlice {
		err := regions[i].Populate(&regionSlice[i])
		if err != nil {
			return err
		}
	}

	return c.driver.VkCmdBlitImage(
		c.handle,
		VkImage(sourceImage.Handle()),
		VkImageLayout(sourceImageLayout),
		VkImage(destinationImage.Handle()),
		VkImageLayout(destinationImageLayout),
		Uint32(regionCount),
		(*VkImageBlit)(regionPtr),
		VkFilter(filter))
}

type CommandBufferOptions struct {
	Level       common.CommandBufferLevel
	BufferCount int
	commandPool CommandPool

	common.HaveNext
}

func (o *CommandBufferOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Level == common.LevelUnset {
		return nil, errors.New("attempted to create command buffers without setting Level")
	}
	if o.BufferCount == 0 {
		return nil, errors.New("attempted to create 0 command buffers")
	}

	createInfo := (*C.VkCommandBufferAllocateInfo)(allocator.Malloc(int(unsafe.Sizeof([1]C.VkCommandBufferAllocateInfo{}))))
	createInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
	createInfo.pNext = next

	createInfo.level = C.VkCommandBufferLevel(o.Level)
	createInfo.commandBufferCount = C.uint32_t(o.BufferCount)
	createInfo.commandPool = C.VkCommandPool(unsafe.Pointer(o.commandPool.Handle()))

	return unsafe.Pointer(createInfo), nil
}

func (o *CommandBufferOptions) MustBeRootOptions() bool {
	return true
}
