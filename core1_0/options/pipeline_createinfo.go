package options

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/iface"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type PipelineFlags int32

const (
	PipelineDisableOptimization                         = C.VK_PIPELINE_CREATE_DISABLE_OPTIMIZATION_BIT
	PipelineAllowDerivatives                            = C.VK_PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT
	PipelineDerivative                                  = C.VK_PIPELINE_CREATE_DERIVATIVE_BIT
	PipelineViewIndexFromDeviceIndex                    = C.VK_PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT
	PipelineDispatchBase                                = C.VK_PIPELINE_CREATE_DISPATCH_BASE_BIT
	PipelineRayTracingNoNullAnyHitShadersKHR            = C.VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR
	PipelineRayTracingNoNullClosestHitShadersKHR        = C.VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR
	PipelineRayTracingNoNullMissShadersKHR              = C.VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR
	PipelineRayTracingNoNullIntersectionShadersKHR      = C.VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR
	PipelineRayTracingSkipTrianglesKHR                  = C.VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR
	PipelineRayTracingSkipAABBsKHR                      = C.VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR
	PipelineRayTracingShaderGroupHandleCaptureReplayKHR = C.VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR
	PipelineDeferCompileNV                              = C.VK_PIPELINE_CREATE_DEFER_COMPILE_BIT_NV
	PipelineCaptureStatisticsKHR                        = C.VK_PIPELINE_CREATE_CAPTURE_STATISTICS_BIT_KHR
	PipelineCaptureInternalRepresentationsKHR           = C.VK_PIPELINE_CREATE_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR
	PipelineIndirectBindableNV                          = C.VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV
	PipelineLibraryKHR                                  = C.VK_PIPELINE_CREATE_LIBRARY_BIT_KHR
	PipelineFailOnPipelineRequiredEXT                   = C.VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT_EXT
	PipelineEarlyReturnOnFailureEXT                     = C.VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT_EXT
	PipelineRayTracingAllowMotionNV                     = C.VK_PIPELINE_CREATE_RAY_TRACING_ALLOW_MOTION_BIT_NV
)

var pipelineFlagsToString = map[PipelineFlags]string{
	PipelineDisableOptimization:                         "Disable Optimization",
	PipelineAllowDerivatives:                            "Allow Derivatives",
	PipelineDerivative:                                  "Derivative",
	PipelineViewIndexFromDeviceIndex:                    "View Index From Device Index",
	PipelineDispatchBase:                                "Dispatch Base",
	PipelineRayTracingNoNullAnyHitShadersKHR:            "Ray Tracing: No Null, Any Hit Shaders (Khronos Extension)",
	PipelineRayTracingNoNullClosestHitShadersKHR:        "Ray Tracing: No Null, Closest Hit Shaders (Khronos Extension)",
	PipelineRayTracingNoNullMissShadersKHR:              "Ray Tracing: No Null, Miss Shaders (Khronos Extension)",
	PipelineRayTracingNoNullIntersectionShadersKHR:      "Ray Tracing: No Null, Intersection Shaders (Khronos Extension)",
	PipelineRayTracingSkipTrianglesKHR:                  "Ray Tracing: Skip Triangles (Khronos Extension)",
	PipelineRayTracingSkipAABBsKHR:                      "Ray Tracing: Skip AABBs (Khronos Extension)",
	PipelineRayTracingShaderGroupHandleCaptureReplayKHR: "Ray Tracing: Shader Group Handle Capture/Replay (Khronos Extension)",
	PipelineDeferCompileNV:                              "Defer Compile (Nvidia Extension)",
	PipelineCaptureStatisticsKHR:                        "Capture Statistics (Khronos Extension)",
	PipelineCaptureInternalRepresentationsKHR:           "Capture Internal Representations (Khronos Extension)",
	PipelineIndirectBindableNV:                          "Indirect Bindable (Nvidia Extension)",
	PipelineLibraryKHR:                                  "Library (Khronos Extension)",
	PipelineFailOnPipelineRequiredEXT:                   "Fail On Pipeline Required (Extension)",
	PipelineEarlyReturnOnFailureEXT:                     "Early Return On Failure (Extension)",
	PipelineRayTracingAllowMotionNV:                     "Ray Tracing: Allow Motion (Nvidia Extension)",
}

func (f PipelineFlags) String() string {
	return common.FlagsToString(f, pipelineFlagsToString)
}

type GraphicsPipelineOptions struct {
	Flags PipelineFlags

	ShaderStages  []ShaderStage
	VertexInput   *VertexInputOptions
	InputAssembly *InputAssemblyOptions
	Tessellation  *TessellationOptions
	Viewport      *ViewportOptions
	Rasterization *RasterizationOptions
	Multisample   *MultisampleOptions
	DepthStencil  *DepthStencilOptions
	ColorBlend    *ColorBlendOptions
	DynamicState  *DynamicStateOptions

	Layout     iface.PipelineLayout
	RenderPass iface.RenderPass

	SubPass           int
	BasePipeline      iface.Pipeline
	BasePipelineIndex int

	core.HaveNext
}

func (o GraphicsPipelineOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkGraphicsPipelineCreateInfo)
	}
	createInfo := (*C.VkGraphicsPipelineCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO
	createInfo.flags = C.VkPipelineCreateFlags(o.Flags)
	createInfo.pNext = next

	stageCount := len(o.ShaderStages)
	createInfo.stageCount = C.uint32_t(stageCount)
	createInfo.pStages = nil
	createInfo.pVertexInputState = nil
	createInfo.pInputAssemblyState = nil
	createInfo.pTessellationState = nil
	createInfo.pViewportState = nil
	createInfo.pRasterizationState = nil
	createInfo.pMultisampleState = nil
	createInfo.pDepthStencilState = nil
	createInfo.pColorBlendState = nil
	createInfo.pDynamicState = nil
	createInfo.layout = nil
	createInfo.renderPass = nil
	createInfo.subpass = C.uint32_t(o.SubPass)
	createInfo.basePipelineHandle = (C.VkPipeline)(C.VK_NULL_HANDLE)
	createInfo.basePipelineIndex = C.int32_t(o.BasePipelineIndex)

	if o.Layout != nil {
		createInfo.layout = (C.VkPipelineLayout)(unsafe.Pointer(o.Layout.Handle()))
	}

	if o.RenderPass != nil {
		createInfo.renderPass = (C.VkRenderPass)(unsafe.Pointer(o.RenderPass.Handle()))
	}

	if o.BasePipeline != nil {
		createInfo.basePipelineHandle = (C.VkPipeline)(unsafe.Pointer(o.BasePipeline.Handle()))
	}

	if stageCount > 0 {
		var err error
		createInfo.pStages, err = core.AllocOptionSlice[C.VkPipelineShaderStageCreateInfo, ShaderStage](allocator, o.ShaderStages)
		if err != nil {
			return nil, err
		}
	}

	if o.VertexInput != nil {
		vertInput, err := core.AllocOptions(allocator, o.VertexInput)
		if err != nil {
			return nil, err
		}

		createInfo.pVertexInputState = (*C.VkPipelineVertexInputStateCreateInfo)(vertInput)
	}

	if o.InputAssembly != nil {
		inputAss, err := core.AllocOptions(allocator, o.InputAssembly)
		if err != nil {
			return nil, err
		}

		createInfo.pInputAssemblyState = (*C.VkPipelineInputAssemblyStateCreateInfo)(inputAss)
	}

	if o.Tessellation != nil {
		tessellation, err := core.AllocOptions(allocator, o.Tessellation)
		if err != nil {
			return nil, err
		}

		createInfo.pTessellationState = (*C.VkPipelineTessellationStateCreateInfo)(tessellation)
	}

	if o.Viewport != nil {
		viewport, err := core.AllocOptions(allocator, o.Viewport)
		if err != nil {
			return nil, err
		}

		createInfo.pViewportState = (*C.VkPipelineViewportStateCreateInfo)(viewport)
	}

	if o.Rasterization != nil {
		rasterization, err := core.AllocOptions(allocator, o.Rasterization)
		if err != nil {
			return nil, err
		}

		createInfo.pRasterizationState = (*C.VkPipelineRasterizationStateCreateInfo)(rasterization)
	}

	if o.Multisample != nil {
		multisample, err := core.AllocOptions(allocator, o.Multisample)
		if err != nil {
			return nil, err
		}

		createInfo.pMultisampleState = (*C.VkPipelineMultisampleStateCreateInfo)(multisample)
	}

	if o.DepthStencil != nil {
		depthStencil, err := core.AllocOptions(allocator, o.DepthStencil)
		if err != nil {
			return nil, err
		}

		createInfo.pDepthStencilState = (*C.VkPipelineDepthStencilStateCreateInfo)(depthStencil)
	}

	if o.ColorBlend != nil {
		colorBlend, err := core.AllocOptions(allocator, o.ColorBlend)
		if err != nil {
			return nil, err
		}

		createInfo.pColorBlendState = (*C.VkPipelineColorBlendStateCreateInfo)(colorBlend)
	}

	if o.DynamicState != nil {
		dynamicState, err := core.AllocOptions(allocator, o.DynamicState)
		if err != nil {
			return nil, err
		}

		createInfo.pDynamicState = (*C.VkPipelineDynamicStateCreateInfo)(dynamicState)
	}

	return preallocatedPointer, nil
}

type ComputePipelineOptions struct {
	Flags  PipelineFlags
	Shader ShaderStage
	Layout iface.PipelineLayout

	BasePipeline      iface.Pipeline
	BasePipelineIndex int

	core.HaveNext
}

func (o ComputePipelineOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkComputePipelineCreateInfo)
	}

	createInfo := (*C.VkComputePipelineCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO
	createInfo.pNext = next
	createInfo.flags = C.VkPipelineCreateFlags(o.Flags)

	_, err := core.AllocOptions(allocator, &o.Shader, unsafe.Pointer(&createInfo.stage))
	if err != nil {
		return nil, err
	}

	createInfo.layout = C.VkPipelineLayout(unsafe.Pointer(o.Layout.Handle()))
	createInfo.basePipelineHandle = C.VkPipeline(unsafe.Pointer(o.BasePipeline.Handle()))
	createInfo.basePipelineIndex = C.int32_t(o.BasePipelineIndex)

	return preallocatedPointer, nil
}
