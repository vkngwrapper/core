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

type vulkanCommandPool struct {
	driver Driver
	handle VkCommandPool
	device VkDevice
}

func (p *vulkanCommandPool) Handle() VkCommandPool {
	return p.handle
}

func (p *vulkanCommandPool) Destroy(callbacks *AllocationCallbacks) {
	p.driver.VkDestroyCommandPool(p.device, p.handle, nil)
}

func (p *vulkanCommandPool) FreeCommandBuffers(buffers []CommandBuffer) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	bufferCount := len(buffers)
	if bufferCount == 0 {
		return
	}

	destroyPtr := allocator.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{})))
	destroySlice := ([]VkCommandBuffer)(unsafe.Slice((*VkCommandBuffer)(destroyPtr), bufferCount))
	for i := 0; i < bufferCount; i++ {
		destroySlice[i] = buffers[i].Handle()
	}

	p.driver.VkFreeCommandBuffers(p.device, p.handle, Uint32(bufferCount), (*VkCommandBuffer)(destroyPtr))
}

func (p *vulkanCommandPool) AllocateCommandBuffers(o *CommandBufferOptions) ([]CommandBuffer, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	o.commandPool = p

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	commandBufferPtr := (*VkCommandBuffer)(arena.Malloc(o.BufferCount * int(unsafe.Sizeof([1]VkCommandBuffer{}))))

	res, err := p.driver.VkAllocateCommandBuffers(p.device, (*VkCommandBufferAllocateInfo)(createInfo), commandBufferPtr)
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	commandBufferArray := ([]VkCommandBuffer)(unsafe.Slice(commandBufferPtr, o.BufferCount))
	var result []CommandBuffer
	for i := 0; i < o.BufferCount; i++ {
		result = append(result, &vulkanCommandBuffer{driver: p.driver, pool: p.handle, device: p.device, handle: commandBufferArray[i]})
	}

	return result, res, nil
}

type CommandPoolResetFlags int32

const (
	CommandPoolResetReleaseResources CommandPoolResetFlags = C.VK_COMMAND_POOL_RESET_RELEASE_RESOURCES_BIT
)

var commandPoolResetFlagsToString = map[CommandPoolResetFlags]string{
	CommandPoolResetReleaseResources: "Release Resources",
}

func (f CommandPoolResetFlags) String() string {
	return common.FlagsToString(f, commandPoolResetFlagsToString)
}

func (p *vulkanCommandPool) Reset(flags CommandPoolResetFlags) (VkResult, error) {
	return p.driver.VkResetCommandPool(p.device, p.handle, VkCommandPoolResetFlags(flags))
}
