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
	BlendZero                  common.BlendFactor = C.VK_BLEND_FACTOR_ZERO
	BlendOne                   common.BlendFactor = C.VK_BLEND_FACTOR_ONE
	BlendSrcColor              common.BlendFactor = C.VK_BLEND_FACTOR_SRC_COLOR
	BlendOneMinusSrcColor      common.BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC_COLOR
	BlendDstColor              common.BlendFactor = C.VK_BLEND_FACTOR_DST_COLOR
	BlendOneMinusDstColor      common.BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_DST_COLOR
	BlendSrcAlpha              common.BlendFactor = C.VK_BLEND_FACTOR_SRC_ALPHA
	BlendOneMinusSrcAlpha      common.BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC_ALPHA
	BlendDstAlpha              common.BlendFactor = C.VK_BLEND_FACTOR_DST_ALPHA
	BlendOneMinusDstAlpha      common.BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_DST_ALPHA
	BlendConstantColor         common.BlendFactor = C.VK_BLEND_FACTOR_CONSTANT_COLOR
	BlendOneMinusConstantColor common.BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_COLOR
	BlendConstantAlpha         common.BlendFactor = C.VK_BLEND_FACTOR_CONSTANT_ALPHA
	BlendOneMinusConstantAlpha common.BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA
	BlendSrcAlphaSaturate      common.BlendFactor = C.VK_BLEND_FACTOR_SRC_ALPHA_SATURATE
	BlendSrc1Color             common.BlendFactor = C.VK_BLEND_FACTOR_SRC1_COLOR
	BlendOneMinusSrc1Color     common.BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR
	BlendSrc1Alpha             common.BlendFactor = C.VK_BLEND_FACTOR_SRC1_ALPHA
	BlendOneMinusSrc1Alpha     common.BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA

	BlendOpAdd      common.BlendOp = C.VK_BLEND_OP_ADD
	BlendOpSubtract common.BlendOp = C.VK_BLEND_OP_SUBTRACT
	BlendOpMin      common.BlendOp = C.VK_BLEND_OP_MIN
	BlendOpMax      common.BlendOp = C.VK_BLEND_OP_MAX

	LogicOpClear        common.LogicOp = C.VK_LOGIC_OP_CLEAR
	LogicOpAnd          common.LogicOp = C.VK_LOGIC_OP_AND
	LogicOpAndReverse   common.LogicOp = C.VK_LOGIC_OP_AND_REVERSE
	LogicOpCopy         common.LogicOp = C.VK_LOGIC_OP_COPY
	LogicOpAndInverted  common.LogicOp = C.VK_LOGIC_OP_AND_INVERTED
	LogicOpNoop         common.LogicOp = C.VK_LOGIC_OP_NO_OP
	LogicOpXor          common.LogicOp = C.VK_LOGIC_OP_XOR
	LogicOpOr           common.LogicOp = C.VK_LOGIC_OP_OR
	LogicOpNor          common.LogicOp = C.VK_LOGIC_OP_NOR
	LogicOpEquivalent   common.LogicOp = C.VK_LOGIC_OP_EQUIVALENT
	LogicOpInvert       common.LogicOp = C.VK_LOGIC_OP_INVERT
	LogicOpOrReverse    common.LogicOp = C.VK_LOGIC_OP_OR_REVERSE
	LogicOpCopyInverted common.LogicOp = C.VK_LOGIC_OP_COPY_INVERTED
	LogicOpOrInverted   common.LogicOp = C.VK_LOGIC_OP_OR_INVERTED
	LogicOpNand         common.LogicOp = C.VK_LOGIC_OP_NAND
	LogicOpSet          common.LogicOp = C.VK_LOGIC_OP_SET
)

func init() {
	BlendZero.Register("0")
	BlendOne.Register("1")
	BlendSrcColor.Register("Src Color")
	BlendOneMinusSrcColor.Register("1 - Src Color")
	BlendDstColor.Register("Dst Color")
	BlendOneMinusDstColor.Register("1 - Dst Color")
	BlendSrcAlpha.Register("Src Alpha")
	BlendOneMinusSrcAlpha.Register("1 - Src Alpha")
	BlendDstAlpha.Register("Dst Alpha")
	BlendOneMinusDstAlpha.Register("1 - Dst Alpha")
	BlendConstantColor.Register("Constant Color")
	BlendOneMinusConstantColor.Register("1 - Constant Color")
	BlendConstantAlpha.Register("Constant Alpha")
	BlendOneMinusConstantAlpha.Register("1 - Constant Alpha")
	BlendSrcAlphaSaturate.Register("Alpha Saturate")
	BlendSrc1Color.Register("Src1 Color")
	BlendOneMinusSrc1Color.Register("1 - Src1 Color")
	BlendSrc1Alpha.Register("Src1 Alpha")
	BlendOneMinusSrc1Alpha.Register("1 - Src1 Alpha")

	BlendOpAdd.Register("Add")
	BlendOpSubtract.Register("Subtract")
	BlendOpMin.Register("Min")
	BlendOpMax.Register("Max")

	LogicOpClear.Register("Clear")
	LogicOpAnd.Register("And")
	LogicOpAndReverse.Register("And Reverse")
	LogicOpCopy.Register("Copy")
	LogicOpAndInverted.Register("And Inverted")
	LogicOpNoop.Register("No-Op")
	LogicOpXor.Register("Xor")
	LogicOpOr.Register("Or")
	LogicOpNor.Register("Nor")
	LogicOpEquivalent.Register("Equivalent")
	LogicOpInvert.Register("Invert")
	LogicOpOrReverse.Register("Or Reverse")
	LogicOpCopyInverted.Register("Copy Inverted")
	LogicOpOrInverted.Register("Or Inverted")
	LogicOpNand.Register("Nand")
	LogicOpSet.Register("Set")
}

type ColorBlendAttachment struct {
	BlendEnabled bool

	SrcColor     common.BlendFactor
	DstColor     common.BlendFactor
	ColorBlendOp common.BlendOp

	SrcAlpha     common.BlendFactor
	DstAlpha     common.BlendFactor
	AlphaBlendOp common.BlendOp

	WriteMask common.ColorComponents
}

type ColorBlendStateOptions struct {
	LogicOpEnabled bool
	LogicOp        common.LogicOp

	BlendConstants [4]float32
	Attachments    []ColorBlendAttachment

	common.HaveNext
}

func (o ColorBlendStateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineColorBlendStateCreateInfo)
	}
	createInfo := (*C.VkPipelineColorBlendStateCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next
	createInfo.logicOpEnable = C.VK_FALSE

	if o.LogicOpEnabled {
		createInfo.logicOpEnable = C.VK_TRUE
	}

	createInfo.logicOp = C.VkLogicOp(o.LogicOp)

	for i := 0; i < 4; i++ {
		createInfo.blendConstants[i] = C.float(o.BlendConstants[i])
	}

	attachmentCount := len(o.Attachments)
	createInfo.attachmentCount = C.uint32_t(attachmentCount)
	createInfo.pAttachments = nil

	if attachmentCount > 0 {
		attachmentsPtr := (*C.VkPipelineColorBlendAttachmentState)(allocator.Malloc(attachmentCount * C.sizeof_struct_VkPipelineColorBlendAttachmentState))
		attachmentSlice := ([]C.VkPipelineColorBlendAttachmentState)(unsafe.Slice(attachmentsPtr, attachmentCount))

		for i := 0; i < attachmentCount; i++ {
			attachmentSlice[i].blendEnable = C.VK_FALSE
			if o.Attachments[i].BlendEnabled {
				attachmentSlice[i].blendEnable = C.VK_TRUE
			}

			attachmentSlice[i].srcColorBlendFactor = C.VkBlendFactor(o.Attachments[i].SrcColor)
			attachmentSlice[i].dstColorBlendFactor = C.VkBlendFactor(o.Attachments[i].DstColor)
			attachmentSlice[i].colorBlendOp = C.VkBlendOp(o.Attachments[i].ColorBlendOp)
			attachmentSlice[i].srcAlphaBlendFactor = C.VkBlendFactor(o.Attachments[i].SrcAlpha)
			attachmentSlice[i].dstAlphaBlendFactor = C.VkBlendFactor(o.Attachments[i].DstAlpha)
			attachmentSlice[i].alphaBlendOp = C.VkBlendOp(o.Attachments[i].AlphaBlendOp)
			attachmentSlice[i].colorWriteMask = C.VkColorComponentFlags(o.Attachments[i].WriteMask)
		}

		createInfo.pAttachments = attachmentsPtr
	}

	return preallocatedPointer, nil
}

func (o ColorBlendStateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkPipelineColorBlendStateCreateInfo)(cDataPointer)
	return createInfo.pNext, nil
}
