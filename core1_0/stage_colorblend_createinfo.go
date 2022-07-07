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

type PipelineColorBlendStateCreateFlags uint32

var pipelineColorBlendStateCreateFLagsMapping = common.NewFlagStringMapping[PipelineColorBlendStateCreateFlags]()

func (f PipelineColorBlendStateCreateFlags) Register(str string) {
	pipelineColorBlendStateCreateFLagsMapping.Register(f, str)
}

func (f PipelineColorBlendStateCreateFlags) String() string {
	return pipelineColorBlendStateCreateFLagsMapping.FlagsToString(f)
}

////

const (
	BlendFactorZero                  BlendFactor = C.VK_BLEND_FACTOR_ZERO
	BlendFactorOne                   BlendFactor = C.VK_BLEND_FACTOR_ONE
	BlendFactorSrcColor              BlendFactor = C.VK_BLEND_FACTOR_SRC_COLOR
	BlendFactorOneMinusSrcColor      BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC_COLOR
	BlendFactorDstColor              BlendFactor = C.VK_BLEND_FACTOR_DST_COLOR
	BlendFactorOneMinusDstColor      BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_DST_COLOR
	BlendFactorSrcAlpha              BlendFactor = C.VK_BLEND_FACTOR_SRC_ALPHA
	BlendFactorOneMinusSrcAlpha      BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC_ALPHA
	BlendFactorDstAlpha              BlendFactor = C.VK_BLEND_FACTOR_DST_ALPHA
	BlendFactorOneMinusDstAlpha      BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_DST_ALPHA
	BlendFactorConstantColor         BlendFactor = C.VK_BLEND_FACTOR_CONSTANT_COLOR
	BlendFactorOneMinusConstantColor BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_COLOR
	BlendFactorConstantAlpha         BlendFactor = C.VK_BLEND_FACTOR_CONSTANT_ALPHA
	BlendFactorOneMinusConstantAlpha BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA
	BlendFactorSrcAlphaSaturate      BlendFactor = C.VK_BLEND_FACTOR_SRC_ALPHA_SATURATE
	BlendFactorSrc1Color             BlendFactor = C.VK_BLEND_FACTOR_SRC1_COLOR
	BlendFactorOneMinusSrc1Color     BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR
	BlendFactorSrc1Alpha             BlendFactor = C.VK_BLEND_FACTOR_SRC1_ALPHA
	BlendFactorOneMinusSrc1Alpha     BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA

	BlendOpAdd      BlendOp = C.VK_BLEND_OP_ADD
	BlendOpSubtract BlendOp = C.VK_BLEND_OP_SUBTRACT
	BlendOpMin      BlendOp = C.VK_BLEND_OP_MIN
	BlendOpMax      BlendOp = C.VK_BLEND_OP_MAX

	LogicOpClear        LogicOp = C.VK_LOGIC_OP_CLEAR
	LogicOpAnd          LogicOp = C.VK_LOGIC_OP_AND
	LogicOpAndReverse   LogicOp = C.VK_LOGIC_OP_AND_REVERSE
	LogicOpCopy         LogicOp = C.VK_LOGIC_OP_COPY
	LogicOpAndInverted  LogicOp = C.VK_LOGIC_OP_AND_INVERTED
	LogicOpNoop         LogicOp = C.VK_LOGIC_OP_NO_OP
	LogicOpXor          LogicOp = C.VK_LOGIC_OP_XOR
	LogicOpOr           LogicOp = C.VK_LOGIC_OP_OR
	LogicOpNor          LogicOp = C.VK_LOGIC_OP_NOR
	LogicOpEquivalent   LogicOp = C.VK_LOGIC_OP_EQUIVALENT
	LogicOpInvert       LogicOp = C.VK_LOGIC_OP_INVERT
	LogicOpOrReverse    LogicOp = C.VK_LOGIC_OP_OR_REVERSE
	LogicOpCopyInverted LogicOp = C.VK_LOGIC_OP_COPY_INVERTED
	LogicOpOrInverted   LogicOp = C.VK_LOGIC_OP_OR_INVERTED
	LogicOpNand         LogicOp = C.VK_LOGIC_OP_NAND
	LogicOpSet          LogicOp = C.VK_LOGIC_OP_SET
)

func init() {
	BlendFactorZero.Register("0")
	BlendFactorOne.Register("1")
	BlendFactorSrcColor.Register("Src Color")
	BlendFactorOneMinusSrcColor.Register("1 - Src Color")
	BlendFactorDstColor.Register("Dst Color")
	BlendFactorOneMinusDstColor.Register("1 - Dst Color")
	BlendFactorSrcAlpha.Register("Src Alpha")
	BlendFactorOneMinusSrcAlpha.Register("1 - Src Alpha")
	BlendFactorDstAlpha.Register("Dst Alpha")
	BlendFactorOneMinusDstAlpha.Register("1 - Dst Alpha")
	BlendFactorConstantColor.Register("Constant Color")
	BlendFactorOneMinusConstantColor.Register("1 - Constant Color")
	BlendFactorConstantAlpha.Register("Constant Alpha")
	BlendFactorOneMinusConstantAlpha.Register("1 - Constant Alpha")
	BlendFactorSrcAlphaSaturate.Register("Alpha Saturate")
	BlendFactorSrc1Color.Register("Src1 Color")
	BlendFactorOneMinusSrc1Color.Register("1 - Src1 Color")
	BlendFactorSrc1Alpha.Register("Src1 Alpha")
	BlendFactorOneMinusSrc1Alpha.Register("1 - Src1 Alpha")

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

type PipelineColorBlendAttachmentState struct {
	BlendEnabled bool

	SrcColorBlendFactor BlendFactor
	DstColorBlendFactor BlendFactor
	ColorBlendOp        BlendOp

	SrcAlphaBlendFactor BlendFactor
	DstAlphaBlendFactor BlendFactor
	AlphaBlendOp        BlendOp

	ColorWriteMask ColorComponentFlags
}

type PipelineColorBlendStateCreateInfo struct {
	Flags          PipelineColorBlendStateCreateFlags
	LogicOpEnabled bool
	LogicOp        LogicOp

	BlendConstants [4]float32
	Attachments    []PipelineColorBlendAttachmentState

	common.NextOptions
}

func (o PipelineColorBlendStateCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineColorBlendStateCreateInfo)
	}
	createInfo := (*C.VkPipelineColorBlendStateCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO
	createInfo.flags = C.VkPipelineColorBlendStateCreateFlags(o.Flags)
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

			attachmentSlice[i].srcColorBlendFactor = C.VkBlendFactor(o.Attachments[i].SrcColorBlendFactor)
			attachmentSlice[i].dstColorBlendFactor = C.VkBlendFactor(o.Attachments[i].DstColorBlendFactor)
			attachmentSlice[i].colorBlendOp = C.VkBlendOp(o.Attachments[i].ColorBlendOp)
			attachmentSlice[i].srcAlphaBlendFactor = C.VkBlendFactor(o.Attachments[i].SrcAlphaBlendFactor)
			attachmentSlice[i].dstAlphaBlendFactor = C.VkBlendFactor(o.Attachments[i].DstAlphaBlendFactor)
			attachmentSlice[i].alphaBlendOp = C.VkBlendOp(o.Attachments[i].AlphaBlendOp)
			attachmentSlice[i].colorWriteMask = C.VkColorComponentFlags(o.Attachments[i].ColorWriteMask)
		}

		createInfo.pAttachments = attachmentsPtr
	}

	return preallocatedPointer, nil
}
