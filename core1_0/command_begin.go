package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

const (
	CommandBufferUsageOneTimeSubmit      CommandBufferUsageFlags = C.VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT
	CommandBufferUsageRenderPassContinue CommandBufferUsageFlags = C.VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT
	CommandBufferUsageSimultaneousUse    CommandBufferUsageFlags = C.VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT

	QueryPrecise QueryControlFlags = C.VK_QUERY_CONTROL_PRECISE_BIT

	QueryPipelineStatisticInputAssemblyVertices                   QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_VERTICES_BIT
	QueryPipelineStatisticInputAssemblyPrimitives                 QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_PRIMITIVES_BIT
	QueryPipelineStatisticVertexShaderInvocations                 QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_VERTEX_SHADER_INVOCATIONS_BIT
	QueryPipelineStatisticGeometryShaderInvocations               QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_INVOCATIONS_BIT
	QueryPipelineStatisticGeometryShaderPrimitives                QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_PRIMITIVES_BIT
	QueryPipelineStatisticClippingInvocations                     QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT
	QueryPipelineStatisticClippingPrimitives                      QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_CLIPPING_PRIMITIVES_BIT
	QueryPipelineStatisticFragmentShaderInvocations               QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_FRAGMENT_SHADER_INVOCATIONS_BIT
	QueryPipelineStatisticTessellationControlShaderPatches        QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_TESSELLATION_CONTROL_SHADER_PATCHES_BIT
	QueryPipelineStatisticTessellationEvaluationShaderInvocations QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_TESSELLATION_EVALUATION_SHADER_INVOCATIONS_BIT
	QueryPipelineStatisticComputeShaderInvocations                QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_COMPUTE_SHADER_INVOCATIONS_BIT
)

func init() {
	CommandBufferUsageOneTimeSubmit.Register("One Time Submit")
	CommandBufferUsageRenderPassContinue.Register("Render Pass Continue")
	CommandBufferUsageSimultaneousUse.Register("Simultaneous Use")

	QueryPrecise.Register("Precise")

	QueryPipelineStatisticInputAssemblyVertices.Register("Input Assembly Vertices")
	QueryPipelineStatisticInputAssemblyPrimitives.Register("Input Assembly Primitives")
	QueryPipelineStatisticVertexShaderInvocations.Register("Vertex Shader Invocations")
	QueryPipelineStatisticGeometryShaderInvocations.Register("Geometry Shader Invocations")
	QueryPipelineStatisticGeometryShaderPrimitives.Register("Geometry Shader Primitives")
	QueryPipelineStatisticClippingInvocations.Register("Clipping Invocations")
	QueryPipelineStatisticClippingPrimitives.Register("Clipping Primitives")
	QueryPipelineStatisticFragmentShaderInvocations.Register("Fragment Shader Invocations")
	QueryPipelineStatisticTessellationControlShaderPatches.Register("Tessellation Control Shader Patches")
	QueryPipelineStatisticTessellationEvaluationShaderInvocations.Register("Tessellation Evaluation Shader Invocations")
	QueryPipelineStatisticComputeShaderInvocations.Register("Compute Shader Invocations")
}

type CommandBufferInheritanceInfo struct {
	Framebuffer Framebuffer
	RenderPass  RenderPass
	Subpass     int

	OcclusionQueryEnable bool
	QueryFlags           QueryControlFlags
	PipelineStatistics   QueryPipelineStatisticFlags

	common.NextOptions
}

func (o CommandBufferInheritanceInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkCommandBufferInheritanceInfo)
	}
	createInfo := (*C.VkCommandBufferInheritanceInfo)(preallocatedPointer)

	createInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO
	createInfo.pNext = next

	createInfo.renderPass = nil
	createInfo.framebuffer = nil

	if o.Framebuffer != nil {
		createInfo.framebuffer = (C.VkFramebuffer)(unsafe.Pointer(o.Framebuffer.Handle()))
	}

	if o.RenderPass != nil {
		createInfo.renderPass = (C.VkRenderPass)(unsafe.Pointer(o.RenderPass.Handle()))
	}

	createInfo.subpass = C.uint32_t(o.Subpass)
	createInfo.occlusionQueryEnable = C.VK_FALSE

	if o.OcclusionQueryEnable {
		createInfo.occlusionQueryEnable = C.VK_TRUE
	}

	createInfo.queryFlags = (C.VkQueryControlFlags)(o.QueryFlags)
	createInfo.pipelineStatistics = (C.VkQueryPipelineStatisticFlags)(o.PipelineStatistics)

	return unsafe.Pointer(createInfo), nil
}

type CommandBufferBeginInfo struct {
	Flags           CommandBufferUsageFlags
	InheritanceInfo *CommandBufferInheritanceInfo

	common.NextOptions
}

func (o CommandBufferBeginInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkCommandBufferBeginInfo)
	}
	createInfo := (*C.VkCommandBufferBeginInfo)(preallocatedPointer)

	createInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
	createInfo.flags = C.VkCommandBufferUsageFlags(o.Flags)
	createInfo.pNext = next

	createInfo.pInheritanceInfo = nil

	if o.InheritanceInfo != nil {
		info, err := common.AllocOptions(allocator, *o.InheritanceInfo)
		if err != nil {
			return nil, err
		}
		createInfo.pInheritanceInfo = (*C.VkCommandBufferInheritanceInfo)(info)
	}

	return unsafe.Pointer(createInfo), nil
}
