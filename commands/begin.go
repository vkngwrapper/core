package commands

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/render_pass"
	"github.com/CannibalVox/cgoparam"
	"strings"
	"unsafe"
)

type BeginInfoFlags int32

const (
	OneTimeSubmit      BeginInfoFlags = C.VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT
	RenderPassContinue BeginInfoFlags = C.VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT
	SimultaneousUse    BeginInfoFlags = C.VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT
)

var beginInfoFlagsToString = map[BeginInfoFlags]string{
	OneTimeSubmit:      "One Time Submit",
	RenderPassContinue: "Render Pass Continue",
	SimultaneousUse:    "Simultaneous Use",
}

func (f BeginInfoFlags) String() string {
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder
	for i := 0; i < 32; i++ {
		checkBit := BeginInfoFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := beginInfoFlagsToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteRune('|')
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}

type CommandBufferInheritanceOptions struct {
	Framebuffer render_pass.Framebuffer
	RenderPass  render_pass.RenderPass
	SubPass     int

	OcclusionQueryEnable bool
	QueryFlags           core.QueryControlFlags
	PipelineStatistics   core.QueryPipelineStatisticFlags

	core.HaveNext
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

	core.HaveNext
}

func (o *BeginOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkCommandBufferBeginInfo)(allocator.Malloc(C.sizeof_struct_VkCommandBufferBeginInfo))

	createInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
	createInfo.flags = C.VkCommandBufferUsageFlags(o.Flags)
	createInfo.pNext = next

	createInfo.pInheritanceInfo = nil

	if o.InheritanceInfo != nil {
		info, err := core.AllocOptions(allocator, o.InheritanceInfo)
		if err != nil {
			return nil, err
		}
		createInfo.pInheritanceInfo = (*C.VkCommandBufferInheritanceInfo)(info)
	}

	return unsafe.Pointer(createInfo), nil
}
