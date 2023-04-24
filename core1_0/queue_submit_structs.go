package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/driver"
)

const (
	// PipelineStageTopOfPipe is equivalent to PipelineStageAllCommands with AccessFlags set to 0
	// when specified in the second synchronization scope, but specifies no stage of execution when
	// specified in the first scope
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
	PipelineStageTopOfPipe PipelineStageFlags = C.VK_PIPELINE_STAGE_TOP_OF_PIPE_BIT
	// PipelineStageDrawIndirect specifies the stage of the Pipeline where DrawIndirect...
	// and DispatchIndirect... data structures are consumed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
	PipelineStageDrawIndirect PipelineStageFlags = C.VK_PIPELINE_STAGE_DRAW_INDIRECT_BIT
	// PipelineStageVertexInput specifies the stage of the Pipeline where vertex and index buffers are
	// consumed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
	PipelineStageVertexInput PipelineStageFlags = C.VK_PIPELINE_STAGE_VERTEX_INPUT_BIT
	// PipelineStageVertexShader specifies the vertex shader stage
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
	PipelineStageVertexShader PipelineStageFlags = C.VK_PIPELINE_STAGE_VERTEX_SHADER_BIT
	// PipelineStageTessellationControlShader specifies the tessellation control shader stage
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
	PipelineStageTessellationControlShader PipelineStageFlags = C.VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT
	// PipelineStageTessellationEvaluationShader specifies the tessellation evaluation shader stage
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
	PipelineStageTessellationEvaluationShader PipelineStageFlags = C.VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT
	// PipelineStageGeometryShader specifies the geometry shader stage
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
	PipelineStageGeometryShader PipelineStageFlags = C.VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT
	// PipelineStageFragmentShader specifies the fragment shader stage
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
	PipelineStageFragmentShader PipelineStageFlags = C.VK_PIPELINE_STAGE_FRAGMENT_SHADER_BIT
	// PipelineStageEarlyFragmentTests specifies the stage of the Pipeline where early fragment tests
	// (depth and stencil tests before fragment shading) are performed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
	PipelineStageEarlyFragmentTests PipelineStageFlags = C.VK_PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT
	// PipelineStageLateFragmentTests specifies the stage of the Pipeline where late fragment tests
	// (depth and stencil tests after fragment shading) are performed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
	PipelineStageLateFragmentTests PipelineStageFlags = C.VK_PIPELINE_STAGE_LATE_FRAGMENT_TESTS_BIT
	// PipelineStageColorAttachmentOutput specifies the stage of the Pipeline after blending where
	// the final color values are output from the Pipeline
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
	PipelineStageColorAttachmentOutput PipelineStageFlags = C.VK_PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT
	// PipelineStageComputeShader specifies the execution of a compute shader
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
	PipelineStageComputeShader PipelineStageFlags = C.VK_PIPELINE_STAGE_COMPUTE_SHADER_BIT
	// PipelineStageTransfer specifies the following commands:
	// * All copy commands including CommandBuffer.CmdCopyQueryPoolResults
	// * CommandBuffer.CmdBlitImage
	// * CommandBuffer.CmdResolveImage
	// * All clear commands, with the exception of CommandBuffer.CmdClearAttachments
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
	PipelineStageTransfer PipelineStageFlags = C.VK_PIPELINE_STAGE_TRANSFER_BIT
	// PipelineStageBottomOfPipe is equivalent to PipelineStageAllCommands with AccessFlags set to 0
	// when specified in the first synchronization scope, but specifies no stage of execution when
	// specified in the second scope
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
	PipelineStageBottomOfPipe PipelineStageFlags = C.VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
	// PipelineStageHost specifies a pseudo-stage indicating execution on the host of reads/writes
	// of DeviceMemory. This stage is not invoked by any commands recorded in a CommandBuffer
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
	PipelineStageHost PipelineStageFlags = C.VK_PIPELINE_STAGE_HOST_BIT
	// PipelineStageAllGraphics specifies the execution of all graphics Pipeline stages, and is
	// equivalent to the logical OR of:
	// * PipelineStageDrawIndirect
	// * PipelineStageVertexInput
	// * PipelineStageVertexShader
	// * PipelineStageTessellationControlShader
	// * PipelineStageTessellationEvaluationShader
	// * PipelineStageGeometryShader
	// * PipelineStageFragmentShader
	// * PipelineStageEarlyFragmentTests
	// * PipelineStageLateFragmentTests
	// * PipelineColorAttachmentOutput
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
	PipelineStageAllGraphics PipelineStageFlags = C.VK_PIPELINE_STAGE_ALL_GRAPHICS_BIT
	// PipelineStageAllCommands specifies all operations performed by all commands supported on the
	// queue it is used with
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
	PipelineStageAllCommands PipelineStageFlags = C.VK_PIPELINE_STAGE_ALL_COMMANDS_BIT
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

// SubmitInfo specifies a Queue submit operation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html
type SubmitInfo struct {
	// CommandBuffers is a slice of CommandBuffer objects to execute in the batch
	CommandBuffers []CommandBuffer
	// WaitSemaphores is a slice of Semaphore objects upon which to wait before the CommandBuffer
	// objects for this batch begin execution
	WaitSemaphores []Semaphore
	// WaitDstStageMask is a slice of PipelineStageFlags at which each corresponding semaphore
	// wait will occur
	WaitDstStageMask []PipelineStageFlags
	// SignalSemaphores is a slice of Semaphore objects which will be signaled when the
	// CommandBuffer objects for this batch have completed execution
	SignalSemaphores []Semaphore

	common.NextOptions
}

func (o SubmitInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if len(o.WaitSemaphores) != len(o.WaitDstStageMask) {
		return nil, errors.Errorf("attempted to submit with %d wait semaphores but %d dst stages- these should match", len(o.WaitSemaphores), len(o.WaitDstStageMask))
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
			if o.WaitSemaphores[i] == nil {
				return nil, errors.Errorf("core1_0.SubmitInfo.WaitSemaphores cannot contain nil elements, but "+
					"element %d is nil", i)
			}
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
			if o.SignalSemaphores[i] == nil {
				return nil, errors.Errorf("core1_0.SubmitInfo.SignalSemaphores cannot contain nil elements, but "+
					"element %d is nil", i)
			}
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
			if o.CommandBuffers[i] == nil {
				return nil, errors.Errorf("core1_0.SubmitInfo.CommandBuffers cannot contain nil elements, but "+
					"element %d is nil", i)
			}
			commandBufferSlice[i] = o.CommandBuffers[i].Handle()
		}

		createInfo.pCommandBuffers = (*C.VkCommandBuffer)(commandBufferPtrUnsafe)
	}

	return preallocatedPointer, nil
}
