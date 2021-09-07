package commands

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/pipeline"
	"github.com/CannibalVox/VKng/core/resource"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type CommandBufferHandle C.VkCommandBuffer
type CommandBuffer struct {
	device C.VkDevice
	pool   C.VkCommandPool
	handle C.VkCommandBuffer
}

func CreateCommandBuffers(allocator cgoalloc.Allocator, device *resource.Device, o *CommandBufferOptions) ([]*CommandBuffer, core.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	deviceHandle := (C.VkDevice)(unsafe.Pointer(device.Handle()))

	commandBufferPtr := (*C.VkCommandBuffer)(arena.Malloc(o.BufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{}))))

	res := core.Result(C.vkAllocateCommandBuffers(deviceHandle, (*C.VkCommandBufferAllocateInfo)(createInfo), commandBufferPtr))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	commandBufferArray := ([]C.VkCommandBuffer)(unsafe.Slice(commandBufferPtr, o.BufferCount))
	var result []*CommandBuffer
	for i := 0; i < o.BufferCount; i++ {
		result = append(result, &CommandBuffer{pool: o.CommandPool.handle, device: deviceHandle, handle: commandBufferArray[i]})
	}

	return result, res, nil
}

func (c *CommandBuffer) Handle() CommandBufferHandle {
	return CommandBufferHandle(c.handle)
}

func (c *CommandBuffer) Destroy() {
	C.vkFreeCommandBuffers(c.device, c.pool, 1, &c.handle)
}

func DestroyBuffers(allocator cgoalloc.Allocator, pool *CommandPool, buffers []*CommandBuffer) {
	bufferCount := len(buffers)
	if bufferCount == 0 {
		return
	}

	destroyPtr := allocator.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{})))
	defer allocator.Free(destroyPtr)

	destroySlice := ([]C.VkCommandBuffer)(unsafe.Slice((*C.VkCommandBuffer)(destroyPtr), bufferCount))
	for i := 0; i < bufferCount; i++ {
		destroySlice[i] = buffers[i].handle
	}

	C.vkFreeCommandBuffers(pool.device, pool.handle, C.uint32_t(bufferCount), (*C.VkCommandBuffer)(destroyPtr))
}

func (c *CommandBuffer) Begin(allocator cgoalloc.Allocator, o *BeginOptions) (core.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return core.VKErrorUnknown, err
	}

	res := core.Result(C.vkBeginCommandBuffer(c.handle, (*C.VkCommandBufferBeginInfo)(createInfo)))
	return res, res.ToError()
}

func (c *CommandBuffer) End() (core.Result, error) {
	res := core.Result(C.vkEndCommandBuffer(c.handle))
	return res, res.ToError()
}

func (c *CommandBuffer) CmdBeginRenderPass(allocator cgoalloc.Allocator, contents SubpassContents, o *RenderPassBeginOptions) error {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return err
	}

	C.vkCmdBeginRenderPass(c.handle, (*C.VkRenderPassBeginInfo)(createInfo), C.VkSubpassContents(contents))

	return nil
}

func (c *CommandBuffer) CmdEndRenderPass() {
	C.vkCmdEndRenderPass(c.handle)
}

func (c *CommandBuffer) CmdBindPipeline(bindPoint core.PipelineBindPoint, pipeline *pipeline.Pipeline) {
	pipelineHandle := (C.VkPipeline)(unsafe.Pointer(pipeline.Handle()))
	C.vkCmdBindPipeline(c.handle, C.VkPipelineBindPoint(bindPoint), pipelineHandle)
}

func (c *CommandBuffer) CmdDraw(vertexCount, instanceCount int, firstVertex, firstInstance uint32) {
	C.vkCmdDraw(c.handle, C.uint32_t(vertexCount), C.uint32_t(instanceCount), C.uint32_t(firstVertex), C.uint32_t(firstInstance))
}

func (c *CommandBuffer) CmdDrawIndexed(indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32) {
	C.vkCmdDrawIndexed(c.handle, C.uint32_t(indexCount), C.uint32_t(instanceCount), C.uint32_t(firstIndex), C.int(vertexOffset), C.uint32_t(firstInstance))
}

func (c *CommandBuffer) CmdBindVertexBuffers(allocator cgoalloc.Allocator, firstBinding uint32, buffers []*resource.Buffer, bufferOffsets []int) {
	bufferCount := len(buffers)

	bufferArrayUnsafe := allocator.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkBuffer{})))
	defer allocator.Free(bufferArrayUnsafe)

	offsetArrayUnsafe := allocator.Malloc(bufferCount * int(unsafe.Sizeof(C.VkDeviceSize(0))))
	defer allocator.Free(offsetArrayUnsafe)

	bufferArrayPtr := (*C.VkBuffer)(bufferArrayUnsafe)
	offsetArrayPtr := (*C.VkDeviceSize)(offsetArrayUnsafe)

	bufferArraySlice := ([]C.VkBuffer)(unsafe.Slice(bufferArrayPtr, bufferCount))
	offsetArraySlice := ([]C.VkDeviceSize)(unsafe.Slice(offsetArrayPtr, bufferCount))

	for i := 0; i < bufferCount; i++ {
		bufferArraySlice[i] = (C.VkBuffer)(unsafe.Pointer(buffers[i].Handle()))
		offsetArraySlice[i] = C.VkDeviceSize(bufferOffsets[i])
	}

	C.vkCmdBindVertexBuffers(c.handle, C.uint32_t(firstBinding), C.uint32_t(bufferCount), bufferArrayPtr, offsetArrayPtr)
}

func (c *CommandBuffer) CmdBindIndexBuffer(buffer *resource.Buffer, offset int, indexType core.IndexType) {
	bufferHandle := C.VkBuffer(unsafe.Pointer(buffer.Handle()))

	C.vkCmdBindIndexBuffer(c.handle, bufferHandle, C.VkDeviceSize(offset), C.VkIndexType(indexType))
}
