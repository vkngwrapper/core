package core

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

type BeginInfoFlags int32

const (
	BeginInfoOneTimeSubmit      BeginInfoFlags = C.VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT
	BeginInfoRenderPassContinue BeginInfoFlags = C.VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT
	BeginInfoSimultaneousUse    BeginInfoFlags = C.VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT
)

var beginInfoFlagsToString = map[BeginInfoFlags]string{
	BeginInfoOneTimeSubmit:      "One Time Submit",
	BeginInfoRenderPassContinue: "Render Pass Continue",
	BeginInfoSimultaneousUse:    "Simultaneous Use",
}

func (f BeginInfoFlags) String() string {
	return common.FlagsToString(f, beginInfoFlagsToString)
}

type CommandBufferInheritanceOptions struct {
	Framebuffer Framebuffer
	RenderPass  RenderPass
	SubPass     int

	OcclusionQueryEnable bool
	QueryFlags           common.QueryControlFlags
	PipelineStatistics   common.QueryPipelineStatisticFlags

	common.HaveNext
}

func (o *CommandBufferInheritanceOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkCommandBufferInheritanceInfo)(allocator.Malloc(C.sizeof_struct_VkCommandBufferInheritanceInfo))

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

type BeginOptions struct {
	Flags           BeginInfoFlags
	InheritanceInfo *CommandBufferInheritanceOptions

	common.HaveNext
}

func (o *BeginOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkCommandBufferBeginInfo)(allocator.Malloc(C.sizeof_struct_VkCommandBufferBeginInfo))

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
