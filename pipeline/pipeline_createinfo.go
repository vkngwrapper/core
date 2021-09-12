package pipeline

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/render_pass"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type Options struct {
	ShaderStages  []*ShaderStage
	VertexInput   *VertexInputOptions
	InputAssembly *InputAssemblyOptions
	Tessellation  *TessellationOptions
	Viewport      *ViewportOptions
	Rasterization *RasterizationOptions
	Multisample   *MultisampleOptions
	DepthStencil  *DepthStencilOptions
	ColorBlend    *ColorBlendOptions
	DynamicState  *DynamicStateOptions

	Layout     PipelineLayout
	RenderPass render_pass.RenderPass

	SubPass           int
	BasePipeline      Pipeline
	BasePipelineIndex int

	core.HaveNext
}

func (o *Options) populate(allocator *cgoparam.Allocator, createInfo *C.VkGraphicsPipelineCreateInfo, next unsafe.Pointer) error {
	createInfo.sType = C.VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO
	createInfo.flags = 0
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
		stagesPtr := (*C.VkPipelineShaderStageCreateInfo)(allocator.Malloc(stageCount * C.sizeof_struct_VkPipelineShaderStageCreateInfo))
		stagesSlice := ([]C.VkPipelineShaderStageCreateInfo)(unsafe.Slice(stagesPtr, stageCount))
		for i := 0; i < stageCount; i++ {
			stageNext, err := core.AllocNext(allocator, o.ShaderStages[i])
			if err != nil {
				return err
			}

			err = o.ShaderStages[i].populate(allocator, &stagesSlice[i], stageNext)
			if err != nil {
				return err
			}
		}
		createInfo.pStages = stagesPtr
	}

	if o.VertexInput != nil {
		vertInput, err := core.AllocOptions(allocator, o.VertexInput)
		if err != nil {
			return err
		}

		createInfo.pVertexInputState = (*C.VkPipelineVertexInputStateCreateInfo)(vertInput)
	}

	if o.InputAssembly != nil {
		inputAss, err := core.AllocOptions(allocator, o.InputAssembly)
		if err != nil {
			return err
		}

		createInfo.pInputAssemblyState = (*C.VkPipelineInputAssemblyStateCreateInfo)(inputAss)
	}

	if o.Tessellation != nil {
		tessellation, err := core.AllocOptions(allocator, o.Tessellation)
		if err != nil {
			return err
		}

		createInfo.pTessellationState = (*C.VkPipelineTessellationStateCreateInfo)(tessellation)
	}

	if o.Viewport != nil {
		viewport, err := core.AllocOptions(allocator, o.Viewport)
		if err != nil {
			return err
		}

		createInfo.pViewportState = (*C.VkPipelineViewportStateCreateInfo)(viewport)
	}

	if o.Rasterization != nil {
		rasterization, err := core.AllocOptions(allocator, o.Rasterization)
		if err != nil {
			return err
		}

		createInfo.pRasterizationState = (*C.VkPipelineRasterizationStateCreateInfo)(rasterization)
	}

	if o.Multisample != nil {
		multisample, err := core.AllocOptions(allocator, o.Multisample)
		if err != nil {
			return err
		}

		createInfo.pMultisampleState = (*C.VkPipelineMultisampleStateCreateInfo)(multisample)
	}

	if o.DepthStencil != nil {
		depthStencil, err := core.AllocOptions(allocator, o.DepthStencil)
		if err != nil {
			return err
		}

		createInfo.pDepthStencilState = (*C.VkPipelineDepthStencilStateCreateInfo)(depthStencil)
	}

	if o.ColorBlend != nil {
		colorBlend, err := core.AllocOptions(allocator, o.ColorBlend)
		if err != nil {
			return err
		}

		createInfo.pColorBlendState = (*C.VkPipelineColorBlendStateCreateInfo)(colorBlend)
	}

	if o.DynamicState != nil {
		dynamicState, err := core.AllocOptions(allocator, o.DynamicState)
		if err != nil {
			return err
		}

		createInfo.pDynamicState = (*C.VkPipelineDynamicStateCreateInfo)(dynamicState)
	}

	return nil
}

func (o *Options) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkGraphicsPipelineCreateInfo)(allocator.Malloc(C.sizeof_struct_VkGraphicsPipelineCreateInfo))

	err := o.populate(allocator, createInfo, next)
	return unsafe.Pointer(createInfo), err
}
