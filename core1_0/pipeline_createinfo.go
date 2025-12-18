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
	"github.com/vkngwrapper/core/v3/common"
)

const (
	// PipelineCreateDisableOptimization specifies that the created Pipeline will not be optimized
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html
	PipelineCreateDisableOptimization PipelineCreateFlags = C.VK_PIPELINE_CREATE_DISABLE_OPTIMIZATION_BIT
	// PipelineCreateAllowDerivatives specifies that the Pipeline to be created is allowed to
	// be the parent of a Pipeline that will be created in a subsequent Pipeline creation call
	PipelineCreateAllowDerivatives PipelineCreateFlags = C.VK_PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT
	// PipelineCreateDerivative specifies that the Pipeline to be created will be a child of a
	// previously-created parent Pipeline
	PipelineCreateDerivative PipelineCreateFlags = C.VK_PIPELINE_CREATE_DERIVATIVE_BIT
)

func init() {
	PipelineCreateDisableOptimization.Register("Disable Optimization")
	PipelineCreateAllowDerivatives.Register("Allow Derivatives")
	PipelineCreateDerivative.Register("Derivative")
}

// GraphicsPipelineCreateInfo specifies parameters of a newly-created graphics Pipeline
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html
type GraphicsPipelineCreateInfo struct {
	// Flags specifies how the Pipeline will be generated
	Flags PipelineCreateFlags

	// Stages is a slice of PipelineShaderStageCreateInfo structures describing the set of shader
	// stages to be included in the graphics Pipeline
	Stages []PipelineShaderStageCreateInfo
	// VertexInputState defines vertex input state for use with vertex shading
	VertexInputState *PipelineVertexInputStateCreateInfo
	// InputAssemblyState determines input assembly behavior for vertex shading
	InputAssemblyState *PipelineInputAssemblyStateCreateInfo
	// TessellationState defines tessellation state used by tessellation shaders
	TessellationState *PipelineTessellationStateCreateInfo
	// ViewportState defines viewport state used when rasterization is enabled
	ViewportState *PipelineViewportStateCreateInfo
	// RasterizationState defines rasterization state
	RasterizationState *PipelineRasterizationStateCreateInfo
	// MultisampleState defines multisample state used when rasterization is enabled
	MultisampleState *PipelineMultisampleStateCreateInfo
	// DepthStencilState defines depth/stencil state used when rasterization is enabled for depth
	// or stencil attachments accessed during rendering
	DepthStencilState *PipelineDepthStencilStateCreateInfo
	// ColorBlendState defines color blend state used when rasterization is enabled for any
	// color attachments accessed during rendering
	ColorBlendState *PipelineColorBlendStateCreateInfo
	// DynamicState defines which properties of the Pipeline state object are dynamic and can
	// be changed independently of the Pipeline state
	DynamicState *PipelineDynamicStateCreateInfo

	// Layout is the description of binding locations used by both the Pipeline and DescriptorSet
	// objects used with the Pipeline
	Layout PipelineLayout
	// RenderPass is a RenderPass object describing the environment in which the Pipeline will be used
	RenderPass RenderPass

	// Subpass is the index of the subpass in the RenderPass where this Pipeline will be used
	Subpass int
	// BasePipeline is a Pipeline object to derive from
	BasePipeline Pipeline
	// BasePipelineIndex is an index into the createInfos parameter to use as a Pipeline to derive from
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
	createInfo.basePipelineHandle = (C.VkPipeline)(nil)
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

// ComputePipelineCreateInfo specifies parameters of a newly-created compute Pipeline
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html
type ComputePipelineCreateInfo struct {
	// Flags specifies how the Pipeline will be generated
	Flags PipelineCreateFlags
	// Stage describes the compute shader
	Stage PipelineShaderStageCreateInfo
	// Layout is the description of binding locations used by both the Pipeline and DescriptorSet
	// objects used with the Pipeline
	Layout PipelineLayout

	// BasePipeline is a Pipeline to derive from
	BasePipeline Pipeline
	// BasePipelineIndex is an index into the createInfos parameters to use as a Pipeline to derive from
	BasePipelineIndex int

	common.NextOptions
}

func (o ComputePipelineCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Layout == nil {
		return nil, errors.New("core1_0.ComputePipelineCreateInfo.Layout cannot be nil")
	}
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkComputePipelineCreateInfo)
	}

	createInfo := (*C.VkComputePipelineCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO
	createInfo.pNext = next
	createInfo.flags = C.VkPipelineCreateFlags(o.Flags)
	createInfo.basePipelineHandle = (C.VkPipeline)(nil)

	_, err := common.AllocOptions(allocator, &o.Stage, unsafe.Pointer(&createInfo.stage))
	if err != nil {
		return nil, err
	}

	if o.BasePipeline != nil {
		createInfo.basePipelineHandle = C.VkPipeline(unsafe.Pointer(o.BasePipeline.Handle()))
	}

	createInfo.layout = C.VkPipelineLayout(unsafe.Pointer(o.Layout.Handle()))
	createInfo.basePipelineIndex = C.int32_t(o.BasePipelineIndex)

	return preallocatedPointer, nil
}
