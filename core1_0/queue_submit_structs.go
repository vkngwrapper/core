package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

const (
	PipelineStageTopOfPipe                    PipelineStageFlags = C.VK_PIPELINE_STAGE_TOP_OF_PIPE_BIT
	PipelineStageDrawIndirect                 PipelineStageFlags = C.VK_PIPELINE_STAGE_DRAW_INDIRECT_BIT
	PipelineStageVertexInput                  PipelineStageFlags = C.VK_PIPELINE_STAGE_VERTEX_INPUT_BIT
	PipelineStageVertexShader                 PipelineStageFlags = C.VK_PIPELINE_STAGE_VERTEX_SHADER_BIT
	PipelineStageTessellationControlShader    PipelineStageFlags = C.VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT
	PipelineStageTessellationEvaluationShader PipelineStageFlags = C.VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT
	PipelineStageGeometryShader               PipelineStageFlags = C.VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT
	PipelineStageFragmentShader               PipelineStageFlags = C.VK_PIPELINE_STAGE_FRAGMENT_SHADER_BIT
	PipelineStageEarlyFragmentTests           PipelineStageFlags = C.VK_PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT
	PipelineStageLateFragmentTests            PipelineStageFlags = C.VK_PIPELINE_STAGE_LATE_FRAGMENT_TESTS_BIT
	PipelineStageColorAttachmentOutput        PipelineStageFlags = C.VK_PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT
	PipelineStageComputeShader                PipelineStageFlags = C.VK_PIPELINE_STAGE_COMPUTE_SHADER_BIT
	PipelineStageTransfer                     PipelineStageFlags = C.VK_PIPELINE_STAGE_TRANSFER_BIT
	PipelineStageBottomOfPipe                 PipelineStageFlags = C.VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
	PipelineStageHost                         PipelineStageFlags = C.VK_PIPELINE_STAGE_HOST_BIT
	PipelineStageAllGraphics                  PipelineStageFlags = C.VK_PIPELINE_STAGE_ALL_GRAPHICS_BIT
	PipelineStageAllCommands                  PipelineStageFlags = C.VK_PIPELINE_STAGE_ALL_COMMANDS_BIT
)

func init() {
	PipelineStageTopOfPipe.Register("Top Of Pipe")
	PipelineStageDrawIndirect.Register("Draw Indirect")
	PipelineStageVertexInput.Register("Vertex Input")
	PipelineStageVertexShader.Register("Vertex Shader")
	PipelineStageTessellationControlShader.Register("Tessellation Control Shader")
	PipelineStageTessellationEvaluationShader.Register("Tessellation Evaluation Shader")
	PipelineStageGeometryShader.Register("Geometry Shader")
	PipelineStageFragmentShader.Register("Fragment Shader")
	PipelineStageEarlyFragmentTests.Register("Early Fragment Tests")
	PipelineStageLateFragmentTests.Register("Late Fragment Tests")
	PipelineStageColorAttachmentOutput.Register("Color Attachment Output")
	PipelineStageComputeShader.Register("Compute Shader")
	PipelineStageTransfer.Register("Transfer")
	PipelineStageBottomOfPipe.Register("Bottom Of Pipe")
	PipelineStageHost.Register("Host")
	PipelineStageAllGraphics.Register("All Graphics")
	PipelineStageAllCommands.Register("All Commands")
}

type SubmitInfo struct {
	CommandBuffers   []CommandBuffer
	WaitSemaphores   []Semaphore
	WaitDstStageMask []PipelineStageFlags
	SignalSemaphores []Semaphore

	common.NextOptions
}

func (o SubmitInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if len(o.WaitSemaphores) != len(o.WaitDstStageMask) {
		return nil, errors.Newf("attempted to submit with %d wait semaphores but %d dst stages- these should match", len(o.WaitSemaphores), len(o.WaitDstStageMask))
	}
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkSubmitInfo)
	}

	createInfo := (*C.VkSubmitInfo)(preallocatedPointer)
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
			stageSlice[i] = (C.VkPipelineStageFlags)(o.WaitDstStageMask[i])
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
		commandBufferSlice := ([]driver.VkCommandBuffer)(unsafe.Slice((*driver.VkCommandBuffer)(commandBufferPtrUnsafe), commandBufferCount))

		for i := 0; i < commandBufferCount; i++ {
			commandBufferSlice[i] = o.CommandBuffers[i].Handle()
		}

		createInfo.pCommandBuffers = (*C.VkCommandBuffer)(commandBufferPtrUnsafe)
	}

	return preallocatedPointer, nil
}
