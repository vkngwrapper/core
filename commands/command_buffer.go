package commands

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/VKng/core/pipeline"
	"github.com/CannibalVox/VKng/core/resources"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type vulkanCommandBuffer struct {
	loader loader.Loader
	device loader.VkDevice
	pool   loader.VkCommandPool
	handle loader.VkCommandBuffer
}

func CreateCommandBuffers(device resources.Device, o *CommandBufferOptions) ([]CommandBuffer, loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	commandBufferPtr := (*loader.VkCommandBuffer)(arena.Malloc(o.BufferCount * int(unsafe.Sizeof([1]loader.VkCommandBuffer{}))))

	res, err := device.Loader().VkAllocateCommandBuffers(device.Handle(), (*loader.VkCommandBufferAllocateInfo)(createInfo), commandBufferPtr)
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	commandBufferArray := ([]loader.VkCommandBuffer)(unsafe.Slice(commandBufferPtr, o.BufferCount))
	var result []CommandBuffer
	for i := 0; i < o.BufferCount; i++ {
		result = append(result, &vulkanCommandBuffer{loader: device.Loader(), pool: o.CommandPool.Handle(), device: device.Handle(), handle: commandBufferArray[i]})
	}

	return result, res, nil
}

func (c *vulkanCommandBuffer) Handle() loader.VkCommandBuffer {
	return c.handle
}

func (c *vulkanCommandBuffer) Destroy() error {
	// cgocheckpointer considers &(c.handle) to be a go pointer containing a go pointer, probably
	// because loader is a go pointer?  Weird but passing a pointer just to the handle works
	handle := c.handle
	return c.loader.VkFreeCommandBuffers(c.device, c.pool, 1, &handle)
}

func (c *vulkanCommandBuffer) Begin(o *BeginOptions) (loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return loader.VKErrorUnknown, err
	}

	return c.loader.VkBeginCommandBuffer(c.handle, (*loader.VkCommandBufferBeginInfo)(createInfo))
}

func (c *vulkanCommandBuffer) End() (loader.VkResult, error) {
	return c.loader.VkEndCommandBuffer(c.handle)
}

func (c *vulkanCommandBuffer) CmdBeginRenderPass(contents SubpassContents, o *RenderPassBeginOptions) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return err
	}

	return c.loader.VkCmdBeginRenderPass(c.handle, (*loader.VkRenderPassBeginInfo)(createInfo), loader.VkSubpassContents(contents))
}

func (c *vulkanCommandBuffer) CmdEndRenderPass() error {
	return c.loader.VkCmdEndRenderPass(c.handle)
}

func (c *vulkanCommandBuffer) CmdBindPipeline(bindPoint core.PipelineBindPoint, pipeline pipeline.Pipeline) error {
	return c.loader.VkCmdBindPipeline(c.handle, loader.VkPipelineBindPoint(bindPoint), pipeline.Handle())
}

func (c *vulkanCommandBuffer) CmdDraw(vertexCount, instanceCount int, firstVertex, firstInstance uint32) error {
	return c.loader.VkCmdDraw(c.handle, loader.Uint32(vertexCount), loader.Uint32(instanceCount), loader.Uint32(firstVertex), loader.Uint32(firstInstance))
}

func (c *vulkanCommandBuffer) CmdDrawIndexed(indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32) error {
	return c.loader.VkCmdDrawIndexed(c.handle, loader.Uint32(indexCount), loader.Uint32(instanceCount), loader.Uint32(firstIndex), loader.Int32(vertexOffset), loader.Uint32(firstInstance))
}

func (c *vulkanCommandBuffer) CmdBindVertexBuffers(firstBinding uint32, buffers []resources.Buffer, bufferOffsets []int) error {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	bufferCount := len(buffers)

	bufferArrayUnsafe := allocator.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkBuffer{})))
	offsetArrayUnsafe := allocator.Malloc(bufferCount * int(unsafe.Sizeof(C.VkDeviceSize(0))))

	bufferArrayPtr := (*loader.VkBuffer)(bufferArrayUnsafe)
	offsetArrayPtr := (*loader.VkDeviceSize)(offsetArrayUnsafe)

	bufferArraySlice := ([]loader.VkBuffer)(unsafe.Slice(bufferArrayPtr, bufferCount))
	offsetArraySlice := ([]loader.VkDeviceSize)(unsafe.Slice(offsetArrayPtr, bufferCount))

	for i := 0; i < bufferCount; i++ {
		bufferArraySlice[i] = buffers[i].Handle()
		offsetArraySlice[i] = loader.VkDeviceSize(bufferOffsets[i])
	}

	return c.loader.VkCmdBindVertexBuffers(c.handle, loader.Uint32(firstBinding), loader.Uint32(bufferCount), bufferArrayPtr, offsetArrayPtr)
}

func (c *vulkanCommandBuffer) CmdBindIndexBuffer(buffer resources.Buffer, offset int, indexType core.IndexType) error {
	return c.loader.VkCmdBindIndexBuffer(c.handle, buffer.Handle(), loader.VkDeviceSize(offset), loader.VkIndexType(indexType))
}
