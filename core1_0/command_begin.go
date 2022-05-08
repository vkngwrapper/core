package core1_0

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

const (
	BeginInfoOneTimeSubmit      common.BeginInfoFlags = C.VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT
	BeginInfoRenderPassContinue common.BeginInfoFlags = C.VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT
	BeginInfoSimultaneousUse    common.BeginInfoFlags = C.VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT

	QueryPrecise common.QueryControlFlags = C.VK_QUERY_CONTROL_PRECISE_BIT

	QueryStatisticInputAssemblyVertices                   common.QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_VERTICES_BIT
	QueryStatisticInputAssemblyPrimitives                 common.QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_PRIMITIVES_BIT
	QueryStatisticVertexShaderInvocations                 common.QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_VERTEX_SHADER_INVOCATIONS_BIT
	QueryStatisticGeometryShaderInvocations               common.QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_INVOCATIONS_BIT
	QueryStatisticGeometryShaderPrimitives                common.QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_PRIMITIVES_BIT
	QueryStatisticClippingInvocations                     common.QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT
	QueryStatisticClippingPrimitives                      common.QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_CLIPPING_PRIMITIVES_BIT
	QueryStatisticFragmentShaderInvocations               common.QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_FRAGMENT_SHADER_INVOCATIONS_BIT
	QueryStatisticTessellationControlShaderPatches        common.QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_TESSELLATION_CONTROL_SHADER_PATCHES_BIT
	QueryStatisticTessellationEvaluationShaderInvocations common.QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_TESSELLATION_EVALUATION_SHADER_INVOCATIONS_BIT
	QueryStatisticComputeShaderInvocations                common.QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_COMPUTE_SHADER_INVOCATIONS_BIT
)

func init() {
	BeginInfoOneTimeSubmit.Register("One Time Submit")
	BeginInfoRenderPassContinue.Register("Render Pass Continue")
	BeginInfoSimultaneousUse.Register("Simultaneous Use")

	QueryPrecise.Register("Precise")

	QueryStatisticInputAssemblyVertices.Register("Input Assembly Vertices")
	QueryStatisticInputAssemblyPrimitives.Register("Input Assembly Primitives")
	QueryStatisticVertexShaderInvocations.Register("Vertex Shader Invocations")
	QueryStatisticGeometryShaderInvocations.Register("Geometry Shader Invocations")
	QueryStatisticGeometryShaderPrimitives.Register("Geometry Shader Primitives")
	QueryStatisticClippingInvocations.Register("Clipping Invocations")
	QueryStatisticClippingPrimitives.Register("Clipping Primitives")
	QueryStatisticFragmentShaderInvocations.Register("Fragment Shader Invocations")
	QueryStatisticTessellationControlShaderPatches.Register("Tessellation Control Shader Patches")
	QueryStatisticTessellationEvaluationShaderInvocations.Register("Tessellation Evaluation Shader Invocations")
	QueryStatisticComputeShaderInvocations.Register("Compute Shader Invocations")
}

type InheritanceOptions struct {
	Framebuffer Framebuffer
	RenderPass  RenderPass
	SubPass     int

	OcclusionQueryEnable bool
	QueryFlags           common.QueryControlFlags
	PipelineStatistics   common.QueryPipelineStatisticFlags

	common.HaveNext
}

func (o InheritanceOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

	createInfo.subpass = C.uint32_t(o.SubPass)
	createInfo.occlusionQueryEnable = C.VK_FALSE

	if o.OcclusionQueryEnable {
		createInfo.occlusionQueryEnable = C.VK_TRUE
	}

	createInfo.queryFlags = (C.VkQueryControlFlags)(o.QueryFlags)
	createInfo.pipelineStatistics = (C.VkQueryPipelineStatisticFlags)(o.PipelineStatistics)

	return unsafe.Pointer(createInfo), nil
}

func (o InheritanceOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkCommandBufferInheritanceInfo)(cDataPointer)
	return createInfo.pNext, nil
}

type BeginOptions struct {
	Flags           common.BeginInfoFlags
	InheritanceInfo *InheritanceOptions

	common.HaveNext
}

func (o BeginOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkCommandBufferBeginInfo)
	}
	createInfo := (*C.VkCommandBufferBeginInfo)(preallocatedPointer)

	createInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
	createInfo.flags = C.VkCommandBufferUsageFlags(o.Flags)
	createInfo.pNext = next

	createInfo.pInheritanceInfo = nil

	if o.InheritanceInfo != nil {
		info, err := common.AllocOptions(allocator, o.InheritanceInfo)
		if err != nil {
			return nil, err
		}
		createInfo.pInheritanceInfo = (*C.VkCommandBufferInheritanceInfo)(info)
	}

	return unsafe.Pointer(createInfo), nil
}

func (o BeginOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkCommandBufferBeginInfo)(cDataPointer)
	return createInfo.pNext, nil
}
