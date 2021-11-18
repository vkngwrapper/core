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

type SubmitOptions struct {
	CommandBuffers   []CommandBuffer
	WaitSemaphores   []Semaphore
	WaitDstStages    []common.PipelineStages
	SignalSemaphores []Semaphore

	common.HaveNext
}

func (o *SubmitOptions) populate(allocator *cgoparam.Allocator, createInfo *C.VkSubmitInfo, next unsafe.Pointer) error {
	if len(o.WaitSemaphores) != len(o.WaitDstStages) {
		return errors.Newf("attempted to submit with %d wait semaphores but %d dst stages- these should match", len(o.WaitSemaphores), len(o.WaitDstStages))
	}

	createInfo.sType = C.VK_STRUCTURE_TYPE_SUBMIT_INFO
	createInfo.pNext = next

	waitSemaphoreCount := len(o.WaitSemaphores)
	createInfo.waitSemaphoreCount = C.uint32_t(waitSemaphoreCount)
	createInfo.pWaitSemaphores = nil
	createInfo.pWaitDstStageMask = nil
	if waitSemaphoreCount > 0 {
		semaphorePtr := (*C.VkSemaphore)(allocator.Malloc(waitSemaphoreCount * int(unsafe.Sizeof([1]C.VkSemaphore{}))))
		semaphoreSlice := ([]C.VkSemaphore)(unsafe.Slice(semaphorePtr, waitSemaphoreCount))

		stagePtr := (*C.VkPipelineStageFlags)(allocator.Malloc(waitSemaphoreCount * int(unsafe.Sizeof(C.VkPipelineStageFlags(0)))))
		stageSlice := ([]C.VkPipelineStageFlags)(unsafe.Slice(stagePtr, waitSemaphoreCount))

		for i := 0; i < waitSemaphoreCount; i++ {
			semaphoreSlice[i] = (C.VkSemaphore)(unsafe.Pointer(o.WaitSemaphores[i].Handle()))
			stageSlice[i] = (C.VkPipelineStageFlags)(o.WaitDstStages[i])
		}

		createInfo.pWaitSemaphores = semaphorePtr
		createInfo.pWaitDstStageMask = stagePtr
	}

	signalSemaphoreCount := len(o.SignalSemaphores)
	createInfo.signalSemaphoreCount = C.uint32_t(signalSemaphoreCount)
	createInfo.pSignalSemaphores = nil
	if signalSemaphoreCount > 0 {
		semaphorePtr := (*C.VkSemaphore)(allocator.Malloc(signalSemaphoreCount * int(unsafe.Sizeof([1]C.VkSemaphore{}))))
		semaphoreSlice := ([]C.VkSemaphore)(unsafe.Slice(semaphorePtr, signalSemaphoreCount))

		for i := 0; i < signalSemaphoreCount; i++ {
			semaphoreSlice[i] = (C.VkSemaphore)(unsafe.Pointer(o.SignalSemaphores[i].Handle()))
		}

		createInfo.pSignalSemaphores = semaphorePtr
	}

	commandBufferCount := len(o.CommandBuffers)
	createInfo.commandBufferCount = C.uint32_t(commandBufferCount)
	createInfo.pCommandBuffers = nil
	if commandBufferCount > 0 {
		commandBufferPtrUnsafe := allocator.Malloc(commandBufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{})))
		commandBufferSlice := ([]VkCommandBuffer)(unsafe.Slice((*VkCommandBuffer)(commandBufferPtrUnsafe), commandBufferCount))

		for i := 0; i < commandBufferCount; i++ {
			commandBufferSlice[i] = o.CommandBuffers[i].Handle()
		}

		createInfo.pCommandBuffers = (*C.VkCommandBuffer)(commandBufferPtrUnsafe)
	}

	return nil
}

func (o *SubmitOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkSubmitInfo)(allocator.Malloc(C.sizeof_struct_VkSubmitInfo))
	err := o.populate(allocator, createInfo, next)
	if err != nil {
		return nil, err
	}

	return unsafe.Pointer(createInfo), nil
}

func (q *vulkanQueue) SubmitToQueue(fence Fence, o []*SubmitOptions) (VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	submitCount := len(o)
	createInfoPtrUnsafe := arena.Malloc(submitCount * C.sizeof_struct_VkSubmitInfo)
	createInfoSlice := ([]C.VkSubmitInfo)(unsafe.Slice((*C.VkSubmitInfo)(createInfoPtrUnsafe), submitCount))

	for i := 0; i < submitCount; i++ {
		next, err := common.AllocNext(arena, o[i])
		if err != nil {
			return VKErrorUnknown, err
		}

		err = o[i].populate(arena, &(createInfoSlice[i]), next)
		if err != nil {
			return VKErrorUnknown, err
		}
	}

	var fenceHandle VkFence = nil
	if fence != nil {
		fenceHandle = fence.Handle()
	}

	return q.Driver().VkQueueSubmit(q.Handle(), Uint32(submitCount), (*VkSubmitInfo)(createInfoPtrUnsafe), fenceHandle)
}
