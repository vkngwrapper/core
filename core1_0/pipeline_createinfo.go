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
	PipelineCreateDisableOptimization PipelineCreateFlags = C.VK_PIPELINE_CREATE_DISABLE_OPTIMIZATION_BIT
	PipelineCreateAllowDerivatives    PipelineCreateFlags = C.VK_PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT
	PipelineCreateDerivative          PipelineCreateFlags = C.VK_PIPELINE_CREATE_DERIVATIVE_BIT
)

func init() {
	PipelineCreateDisableOptimization.Register("Disable Optimization")
	PipelineCreateAllowDerivatives.Register("Allow Derivatives")
	PipelineCreateDerivative.Register("Derivative")
}

type GraphicsPipelineCreateOptions struct {
	Flags PipelineCreateFlags

	ShaderStages  []ShaderStageOptions
	VertexInput   *VertexInputStateOptions
	InputAssembly *InputAssemblyStateOptions
	Tessellation  *TessellationStateOptions
	Viewport      *ViewportStateOptions
	Rasterization *RasterizationStateOptions
	Multisample   *MultisampleStateOptions
	DepthStencil  *DepthStencilStateOptions
	ColorBlend    *ColorBlendStateOptions
	DynamicState  *DynamicStateOptions

	Layout     PipelineLayout
	RenderPass RenderPass

	SubPass           int
	BasePipeline      Pipeline
	BasePipelineIndex int

	common.HaveNext
}

func (o GraphicsPipelineCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
		createInfo.pStages, err = common.AllocOptionSlice[C.VkPipelineShaderStageCreateInfo, ShaderStageOptions](allocator, o.ShaderStages)
		if err != nil {
			return nil, err
		}
	}

	if o.VertexInput != nil {
		vertInput, err := common.AllocOptions(allocator, o.VertexInput)
		if err != nil {
			return nil, err
		}

		createInfo.pVertexInputState = (*C.VkPipelineVertexInputStateCreateInfo)(vertInput)
	}

	if o.InputAssembly != nil {
		inputAss, err := common.AllocOptions(allocator, o.InputAssembly)
		if err != nil {
			return nil, err
		}

		createInfo.pInputAssemblyState = (*C.VkPipelineInputAssemblyStateCreateInfo)(inputAss)
	}

	if o.Tessellation != nil {
		tessellation, err := common.AllocOptions(allocator, o.Tessellation)
		if err != nil {
			return nil, err
		}

		createInfo.pTessellationState = (*C.VkPipelineTessellationStateCreateInfo)(tessellation)
	}

	if o.Viewport != nil {
		viewport, err := common.AllocOptions(allocator, o.Viewport)
		if err != nil {
			return nil, err
		}

		createInfo.pViewportState = (*C.VkPipelineViewportStateCreateInfo)(viewport)
	}

	if o.Rasterization != nil {
		rasterization, err := common.AllocOptions(allocator, o.Rasterization)
		if err != nil {
			return nil, err
		}

		createInfo.pRasterizationState = (*C.VkPipelineRasterizationStateCreateInfo)(rasterization)
	}

	if o.Multisample != nil {
		multisample, err := common.AllocOptions(allocator, o.Multisample)
		if err != nil {
			return nil, err
		}

		createInfo.pMultisampleState = (*C.VkPipelineMultisampleStateCreateInfo)(multisample)
	}

	if o.DepthStencil != nil {
		depthStencil, err := common.AllocOptions(allocator, o.DepthStencil)
		if err != nil {
			return nil, err
		}

		createInfo.pDepthStencilState = (*C.VkPipelineDepthStencilStateCreateInfo)(depthStencil)
	}

	if o.ColorBlend != nil {
		colorBlend, err := common.AllocOptions(allocator, o.ColorBlend)
		if err != nil {
			return nil, err
		}

		createInfo.pColorBlendState = (*C.VkPipelineColorBlendStateCreateInfo)(colorBlend)
	}

	if o.DynamicState != nil {
		dynamicState, err := common.AllocOptions(allocator, o.DynamicState)
		if err != nil {
			return nil, err
		}

		createInfo.pDynamicState = (*C.VkPipelineDynamicStateCreateInfo)(dynamicState)
	}

	return preallocatedPointer, nil
}

func (o GraphicsPipelineCreateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkGraphicsPipelineCreateInfo)(cDataPointer)
	return createInfo.pNext, nil
}

type ComputePipelineCreateOptions struct {
	Flags  PipelineCreateFlags
	Shader ShaderStageOptions
	Layout PipelineLayout

	BasePipeline      Pipeline
	BasePipelineIndex int

	common.HaveNext
}

func (o ComputePipelineCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkComputePipelineCreateInfo)
	}

	createInfo := (*C.VkComputePipelineCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO
	createInfo.pNext = next
	createInfo.flags = C.VkPipelineCreateFlags(o.Flags)

	_, err := common.AllocOptions(allocator, &o.Shader, unsafe.Pointer(&createInfo.stage))
	if err != nil {
		return nil, err
	}

	createInfo.layout = C.VkPipelineLayout(unsafe.Pointer(o.Layout.Handle()))
	createInfo.basePipelineHandle = C.VkPipeline(unsafe.Pointer(o.BasePipeline.Handle()))
	createInfo.basePipelineIndex = C.int32_t(o.BasePipelineIndex)

	return preallocatedPointer, nil
}

func (o ComputePipelineCreateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkComputePipelineCreateInfo)(cDataPointer)
	return createInfo.pNext, nil
}
