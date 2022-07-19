package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// PipelineColorBlendStateCreateFlags specifies additional parameters of an Image
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateFlagBits.html
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
	// BlendFactorZero provides 0 for all channels to the blend operation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorZero BlendFactor = C.VK_BLEND_FACTOR_ZERO
	// BlendFactorOne provides 1 for all channels to the blend operation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorOne BlendFactor = C.VK_BLEND_FACTOR_ONE
	// BlendFactorSrcColor provides R(s0), G(s0), and B(s0) to color blend operations and
	// A(s0) to alpha blend operations
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorSrcColor BlendFactor = C.VK_BLEND_FACTOR_SRC_COLOR
	// BlendFactorOneMinusSrcColor provides 1-R(s0), 1-G(s0), and 1-B(s0) to color blend operations
	// and 1-A(s0) to alpha blend operations
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorOneMinusSrcColor BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC_COLOR
	// BlendFactorDstColor provides R(d), G(d), and B(d) to color blend operations and A(d)
	// to alpha blend operations
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorDstColor BlendFactor = C.VK_BLEND_FACTOR_DST_COLOR
	// BlendFactorOneMinusDstColor provides 1-R(d), 1-G(d), and 1-B(d) to color blend operations
	// and 1-A(d) to alpha blend operations
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorOneMinusDstColor BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_DST_COLOR
	// BlendFactorSrcAlpha provides A(s0) for all channels to the blend operation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorSrcAlpha BlendFactor = C.VK_BLEND_FACTOR_SRC_ALPHA
	// BlendFactorOneMinusSrcAlpha provides 1-A(s0) for all channels to the blend operation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorOneMinusSrcAlpha BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC_ALPHA
	// BlendFactorDstAlpha provides A(d) for all channels to the blend operation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorDstAlpha BlendFactor = C.VK_BLEND_FACTOR_DST_ALPHA
	// BlendFactorOneMinusDstAlpha provides 1-A(d) for all channels to the blend operation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorOneMinusDstAlpha BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_DST_ALPHA
	// BlendFactorConstantColor provides R(c), G(c), and B(c) to color blend operations and
	// A(c) to alpha blend operations
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorConstantColor BlendFactor = C.VK_BLEND_FACTOR_CONSTANT_COLOR
	// BlendFactorOneMinusConstantColor provides 1-R(c), 1-G(c), and 1-B(c) to color blend
	// operations and 1-A(c) to alpha blend operations
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorOneMinusConstantColor BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_COLOR
	// BlendFactorConstantAlpha provides A(c) for all channels to the blend operation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorConstantAlpha BlendFactor = C.VK_BLEND_FACTOR_CONSTANT_ALPHA
	// BlendFactorOneMinusConstantAlpha provides 1-A(c) for all channels to the blend operation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorOneMinusConstantAlpha BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA
	// BlendFactorSrcAlphaSaturate provides MIN(A(s0), 1-A(d)) for all channels to color blend
	// operations and 1 to alpha blend operations
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorSrcAlphaSaturate BlendFactor = C.VK_BLEND_FACTOR_SRC_ALPHA_SATURATE
	// BlendFactorSrc1Color provides R(s1), G(s1), and B(s1) to color blend operations and
	// A(s1) to alpha blend operations
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorSrc1Color BlendFactor = C.VK_BLEND_FACTOR_SRC1_COLOR
	// BlendFactorOneMinusSrc1Color provides 1-R(s1), 1-G(s1), and 1-B(s1) to color blend
	// operations and 1-A(s1) to alpha blend operations
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorOneMinusSrc1Color BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR
	// BlendFactorSrc1Alpha provides A(s1) for all channels to the blend operation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorSrc1Alpha BlendFactor = C.VK_BLEND_FACTOR_SRC1_ALPHA
	// BlendFactorOneMinusSrc1Alpha provides 1-A(s1) for all channels to the blend operation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
	BlendFactorOneMinusSrc1Alpha BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA

	// BlendOpAdd outputs the following, given s0 is the source color, d is the destination
	// color, sf is the source blend factor, and df is the destination blend factor:
	// * Color blend operation: [R(s0) * R(sf) + R(d) * R(df), G(s0) * G(sf) + G(d) * G(df),
	// B(s0) * B(sf) + B(d) * B(df)]
	// * Alpha blend operation: A(s0) * A(sf) + A(d) * A(df)
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOp.html
	BlendOpAdd BlendOp = C.VK_BLEND_OP_ADD
	// BlendOpSubtract  outputs the following, given s0 is the source color, d is the destination
	// color, sf is the source blend factor, and df is the destination blend factor:
	// * Color blend operation: [R(s0) * R(sf) - R(d) * R(df), G(s0) * G(sf) - G(d) * G(df),
	// B(s0) * B(sf) - B(d) * B(df)]
	// * Alpha blend operation: A(s0) * A(sf) - A(d) * A(df)
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOp.html
	BlendOpSubtract BlendOp = C.VK_BLEND_OP_SUBTRACT
	// BlendOpMin outputs the following, given s0 is the source color and d is the destination color:
	// * Color blend operation: [MIN(R(s0), R(d)), MIN(G(s0), G(d)), MIN(B(s0), B(d))]
	// * Alpha blend operation: MIN(A(s0), A(d))
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOp.html
	BlendOpMin BlendOp = C.VK_BLEND_OP_MIN
	// BlendOpMax outputs the following, given s0 is the source color and d is the destination color:
	// * Color blend operation: [MAX(R(s0), R(d)), MAX(G(s0), G(d)), MAX(B(s0), B(d))]
	// * Alpha blend operation: MAX(A(s0), A(d))
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOp.html
	BlendOpMax BlendOp = C.VK_BLEND_OP_MAX

	// LogicOpClear sets the output value to 0
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html
	LogicOpClear LogicOp = C.VK_LOGIC_OP_CLEAR
	// LogicOpAnd sets the output value to s0 & d
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html
	LogicOpAnd LogicOp = C.VK_LOGIC_OP_AND
	// LogicOpAndReverse sets the output value to s0 & ~d
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html
	LogicOpAndReverse LogicOp = C.VK_LOGIC_OP_AND_REVERSE
	// LogicOpCopy sets the output value to s0
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html
	LogicOpCopy LogicOp = C.VK_LOGIC_OP_COPY
	// LogicOpAndInverted sets the output value to ~s0 & d
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html
	LogicOpAndInverted LogicOp = C.VK_LOGIC_OP_AND_INVERTED
	// LogicOpNoop sets the output value to d
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html
	LogicOpNoop LogicOp = C.VK_LOGIC_OP_NO_OP
	// LogicOpXor sets the output value to s0 ^ d
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html
	LogicOpXor LogicOp = C.VK_LOGIC_OP_XOR
	// LogicOpOr sets the output value to s0 | d
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html
	LogicOpOr LogicOp = C.VK_LOGIC_OP_OR
	// LogicOpNor sets the output value to ~(s0 | d)
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html
	LogicOpNor LogicOp = C.VK_LOGIC_OP_NOR
	// LogicOpEquivalent sets the output value to ~(s0 ^ d)
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html
	LogicOpEquivalent LogicOp = C.VK_LOGIC_OP_EQUIVALENT
	// LogicOpInvert sets the output value to ~d
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html
	LogicOpInvert LogicOp = C.VK_LOGIC_OP_INVERT
	// LogicOpOrReverse sets the output value to s0 | ~d
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html
	LogicOpOrReverse LogicOp = C.VK_LOGIC_OP_OR_REVERSE
	// LogicOpCopyInverted sets the output value to ~s0
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html
	LogicOpCopyInverted LogicOp = C.VK_LOGIC_OP_COPY_INVERTED
	// LogicOpOrInverted sets the output value to ~s0 | d
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html
	LogicOpOrInverted LogicOp = C.VK_LOGIC_OP_OR_INVERTED
	// LogicOpNand sets the output value to ~(s0 & d)
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html
	LogicOpNand LogicOp = C.VK_LOGIC_OP_NAND
	// LogicOpSet sets the output value to 0xFFFFF...
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html
	LogicOpSet LogicOp = C.VK_LOGIC_OP_SET
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

// PipelineColorBlendAttachmentState specifies a Pipeline color blend attachment state
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html
type PipelineColorBlendAttachmentState struct {
	// BlendEnabled controls whether blending is enabled for the corresponding color attachment
	BlendEnabled bool

	// SrcColorBlendFactor selects which blend factor is used to determine the source factors
	// [R(sf), G(sf), B(sf)]
	SrcColorBlendFactor BlendFactor
	// DstColorBlendFactor selects which blend factor is used to determine the destination factors
	// [R(df), G(df), B(df)]
	DstColorBlendFactor BlendFactor
	// ColorBlendOp selects which blend operation is used to calculate the RGBG values to write to
	// the color attachment
	ColorBlendOp BlendOp

	// SrcAlphaBlendFactor selects which blend factor is used to determine the source factor A(sf)
	SrcAlphaBlendFactor BlendFactor
	// DstAlphaBlendFactor selects which blend factor is used to determine the detination factor A(sf)
	DstAlphaBlendFactor BlendFactor
	// AlphaBlendOp selects which blend operation is used to calculate the alpha values to write to the color
	// attachment
	AlphaBlendOp BlendOp

	// ColorWriteMask specifies which of the R, G, B, and/or A components are enabled for writing
	ColorWriteMask ColorComponentFlags
}

// PipelineColorBlendStateCreateInfo specifies parameters of a newly-created Pipeline color blend state
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html
type PipelineColorBlendStateCreateInfo struct {
	// Flags specifies additional color blending information
	Flags PipelineColorBlendStateCreateFlags
	// LogicOpEnabled controls whether to apply logical operations
	LogicOpEnabled bool
	// LogicOp selects which logical operation to apply
	LogicOp LogicOp

	// BlendConstants is an array of 4 values used as the R, G, B, and A components of the
	// blend constant that are used in blending, depending on the blend factor
	BlendConstants [4]float32
	// Attachments is a slice of PipelineColorBlendAttachmentState structures defining blend state
	// for each color attachment
	Attachments []PipelineColorBlendAttachmentState

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
