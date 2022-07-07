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

type GraphicsPipelineCreateInfo struct {
	Flags PipelineCreateFlags

	Stages             []PipelineShaderStageCreateInfo
	VertexInputState   *PipelineVertexInputStateCreateInfo
	InputAssemblyState *PipelineInputAssemblyStateCreateInfo
	TessellationState  *PipelineTessellationStateCreateInfo
	ViewportState      *PipelineViewportStateCreateInfo
	RasterizationState *PipelineRasterizationStateCreateInfo
	MultisampleState   *PipelineMultisampleStateCreateInfo
	DepthStencilState  *PipelineDepthStencilStateCreateInfo
	ColorBlendState    *PipelineColorBlendStateCreateInfo
	DynamicState       *PipelineDynamicStateCreateInfo

	Layout     PipelineLayout
	RenderPass RenderPass

	Subpass           int
	BasePipeline      Pipeline
	BasePipelineIndex int

	common.NextOptions
}

func (o GraphicsPipelineCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkGraphicsPipelineCreateInfo)
	}
	createInfo := (*C.VkGraphicsPipelineCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO
	createInfo.flags = C.VkPipelineCreateFlags(o.Flags)
	createInfo.pNext = next

	stageCount := len(o.Stages)
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
	createInfo.subpass = C.uint32_t(o.Subpass)
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
		createInfo.pStages, err = common.AllocOptionSlice[C.VkPipelineShaderStageCreateInfo, PipelineShaderStageCreateInfo](allocator, o.Stages)
		if err != nil {
			return nil, err
		}
	}

	if o.VertexInputState != nil {
		vertInput, err := common.AllocOptions(allocator, o.VertexInputState)
		if err != nil {
			return nil, err
		}

		createInfo.pVertexInputState = (*C.VkPipelineVertexInputStateCreateInfo)(vertInput)
	}

	if o.InputAssemblyState != nil {
		inputAss, err := common.AllocOptions(allocator, o.InputAssemblyState)
		if err != nil {
			return nil, err
		}

		createInfo.pInputAssemblyState = (*C.VkPipelineInputAssemblyStateCreateInfo)(inputAss)
	}

	if o.TessellationState != nil {
		tessellation, err := common.AllocOptions(allocator, o.TessellationState)
		if err != nil {
			return nil, err
		}

		createInfo.pTessellationState = (*C.VkPipelineTessellationStateCreateInfo)(tessellation)
	}

	if o.ViewportState != nil {
		viewport, err := common.AllocOptions(allocator, o.ViewportState)
		if err != nil {
			return nil, err
		}

		createInfo.pViewportState = (*C.VkPipelineViewportStateCreateInfo)(viewport)
	}

	if o.RasterizationState != nil {
		rasterization, err := common.AllocOptions(allocator, o.RasterizationState)
		if err != nil {
			return nil, err
		}

		createInfo.pRasterizationState = (*C.VkPipelineRasterizationStateCreateInfo)(rasterization)
	}

	if o.MultisampleState != nil {
		multisample, err := common.AllocOptions(allocator, o.MultisampleState)
		if err != nil {
			return nil, err
		}

		createInfo.pMultisampleState = (*C.VkPipelineMultisampleStateCreateInfo)(multisample)
	}

	if o.DepthStencilState != nil {
		depthStencil, err := common.AllocOptions(allocator, o.DepthStencilState)
		if err != nil {
			return nil, err
		}

		createInfo.pDepthStencilState = (*C.VkPipelineDepthStencilStateCreateInfo)(depthStencil)
	}

	if o.ColorBlendState != nil {
		colorBlend, err := common.AllocOptions(allocator, o.ColorBlendState)
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

type ComputePipelineCreateInfo struct {
	Flags  PipelineCreateFlags
	Stage  PipelineShaderStageCreateInfo
	Layout PipelineLayout

	BasePipeline      Pipeline
	BasePipelineIndex int

	common.NextOptions
}

func (o ComputePipelineCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkComputePipelineCreateInfo)
	}

	createInfo := (*C.VkComputePipelineCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO
	createInfo.pNext = next
	createInfo.flags = C.VkPipelineCreateFlags(o.Flags)

	_, err := common.AllocOptions(allocator, &o.Stage, unsafe.Pointer(&createInfo.stage))
	if err != nil {
		return nil, err
	}

	createInfo.layout = C.VkPipelineLayout(unsafe.Pointer(o.Layout.Handle()))
	createInfo.basePipelineHandle = C.VkPipeline(unsafe.Pointer(o.BasePipeline.Handle()))
	createInfo.basePipelineIndex = C.int32_t(o.BasePipelineIndex)

	return preallocatedPointer, nil
}
