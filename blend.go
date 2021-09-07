package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "strings"

type BlendFactor int32

const (
	BlendZero                  BlendFactor = C.VK_BLEND_FACTOR_ZERO
	BlendOne                   BlendFactor = C.VK_BLEND_FACTOR_ONE
	BlendSrcColor              BlendFactor = C.VK_BLEND_FACTOR_SRC_COLOR
	BlendOneMinusSrcColor      BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC_COLOR
	BlendDstColor              BlendFactor = C.VK_BLEND_FACTOR_DST_COLOR
	BlendOneMinusDstColor      BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_DST_COLOR
	BlendSrcAlpha              BlendFactor = C.VK_BLEND_FACTOR_SRC_ALPHA
	BlendOneMinusSrcAlpha      BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC_ALPHA
	BlendDstAlpha              BlendFactor = C.VK_BLEND_FACTOR_DST_ALPHA
	BlendOneMinusDstAlpha      BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_DST_ALPHA
	BlendConstantColor         BlendFactor = C.VK_BLEND_FACTOR_CONSTANT_COLOR
	BlendOneMinusConstantColor BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_COLOR
	BlendConstantAlpha         BlendFactor = C.VK_BLEND_FACTOR_CONSTANT_ALPHA
	BlendOneMinusConstantAlpha BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA
	BlendSrcAlphaSaturate      BlendFactor = C.VK_BLEND_FACTOR_SRC_ALPHA_SATURATE
	BlendSrc1Color             BlendFactor = C.VK_BLEND_FACTOR_SRC1_COLOR
	BlendOneMinusSrc1Color     BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR
	BlendSrc1Alpha             BlendFactor = C.VK_BLEND_FACTOR_SRC1_ALPHA
	BlendOneMinusSrc1Alpha     BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA
)

var blendFactorToString = map[BlendFactor]string{
	BlendZero:                  "0",
	BlendOne:                   "1",
	BlendSrcColor:              "Src Color",
	BlendOneMinusSrcColor:      "1 - Src Color",
	BlendDstColor:              "Dst Color",
	BlendOneMinusDstColor:      "1 - Dst Color",
	BlendSrcAlpha:              "Src Alpha",
	BlendOneMinusSrcAlpha:      "1 - Src Alpha",
	BlendDstAlpha:              "Dst Alpha",
	BlendOneMinusDstAlpha:      "1 - Dst Alpha",
	BlendConstantColor:         "Constant Color",
	BlendOneMinusConstantColor: "1 - Constant Color",
	BlendConstantAlpha:         "Constant Alpha",
	BlendOneMinusConstantAlpha: "1 - Constant Alpha",
	BlendSrcAlphaSaturate:      "Alpha Saturate",
	BlendSrc1Color:             "Src1 Color",
	BlendOneMinusSrc1Color:     "1 - Src1 Color",
	BlendSrc1Alpha:             "Src1 Alpha",
	BlendOneMinusSrc1Alpha:     "1 - Src1 Alpha",
}

func (f BlendFactor) String() string {
	return blendFactorToString[f]
}

type BlendOp int32

const (
	BlendOpAdd              BlendOp = C.VK_BLEND_OP_ADD
	BlendOpSubtract         BlendOp = C.VK_BLEND_OP_SUBTRACT
	BlendOpMin              BlendOp = C.VK_BLEND_OP_MIN
	BlendOpMax              BlendOp = C.VK_BLEND_OP_MAX
	BlendOpZero             BlendOp = C.VK_BLEND_OP_ZERO_EXT
	BlendOpSrc              BlendOp = C.VK_BLEND_OP_SRC_EXT
	BlendOpDst              BlendOp = C.VK_BLEND_OP_DST_EXT
	BlendOpSrcOver          BlendOp = C.VK_BLEND_OP_SRC_OVER_EXT
	BlendOpDstOver          BlendOp = C.VK_BLEND_OP_DST_OVER_EXT
	BlendOpSrcIn            BlendOp = C.VK_BLEND_OP_SRC_IN_EXT
	BlendOpDstIn            BlendOp = C.VK_BLEND_OP_DST_IN_EXT
	BlendOpSrcOut           BlendOp = C.VK_BLEND_OP_SRC_OUT_EXT
	BlendOpDstOut           BlendOp = C.VK_BLEND_OP_DST_OUT_EXT
	BlendOpSrcAtop          BlendOp = C.VK_BLEND_OP_SRC_ATOP_EXT
	BlendOpDstAtop          BlendOp = C.VK_BLEND_OP_DST_ATOP_EXT
	BlendOpXor              BlendOp = C.VK_BLEND_OP_XOR_EXT
	BlendOpMultiply         BlendOp = C.VK_BLEND_OP_MULTIPLY_EXT
	BlendOpScreen           BlendOp = C.VK_BLEND_OP_SCREEN_EXT
	BlendOpOverlay          BlendOp = C.VK_BLEND_OP_OVERLAY_EXT
	BlendOpDarken           BlendOp = C.VK_BLEND_OP_DARKEN_EXT
	BlendOpLighten          BlendOp = C.VK_BLEND_OP_LIGHTEN_EXT
	BlendOpColorDodge       BlendOp = C.VK_BLEND_OP_COLORDODGE_EXT
	BlendOpColorBurn        BlendOp = C.VK_BLEND_OP_COLORBURN_EXT
	BlendOpHardLight        BlendOp = C.VK_BLEND_OP_HARDLIGHT_EXT
	BlendOpSoftLight        BlendOp = C.VK_BLEND_OP_SOFTLIGHT_EXT
	BlendOpDifference       BlendOp = C.VK_BLEND_OP_DIFFERENCE_EXT
	BlendOpExclusion        BlendOp = C.VK_BLEND_OP_EXCLUSION_EXT
	BlendOpInvert           BlendOp = C.VK_BLEND_OP_INVERT_EXT
	BlendOpInvertRGB        BlendOp = C.VK_BLEND_OP_INVERT_RGB_EXT
	BlendOpLinearDodge      BlendOp = C.VK_BLEND_OP_LINEARDODGE_EXT
	BlendOpLinearBurn       BlendOp = C.VK_BLEND_OP_LINEARBURN_EXT
	BlendOpVividLight       BlendOp = C.VK_BLEND_OP_VIVIDLIGHT_EXT
	BlendOpLinearLight      BlendOp = C.VK_BLEND_OP_LINEARLIGHT_EXT
	BlendOpPinLight         BlendOp = C.VK_BLEND_OP_PINLIGHT_EXT
	BlendOpHardMix          BlendOp = C.VK_BLEND_OP_HARDMIX_EXT
	BlendOpHSLHue           BlendOp = C.VK_BLEND_OP_HSL_HUE_EXT
	BlendOpHSLSaturation    BlendOp = C.VK_BLEND_OP_HSL_SATURATION_EXT
	BlendOpHSLColor         BlendOp = C.VK_BLEND_OP_HSL_COLOR_EXT
	BlendOpHSLLuminosity    BlendOp = C.VK_BLEND_OP_HSL_LUMINOSITY_EXT
	BlendOpPlus             BlendOp = C.VK_BLEND_OP_PLUS_EXT
	BlendOpPlusClamped      BlendOp = C.VK_BLEND_OP_PLUS_CLAMPED_EXT
	BlendOpPlusClampedAlpha BlendOp = C.VK_BLEND_OP_PLUS_CLAMPED_ALPHA_EXT
	BlendOpPlusDarker       BlendOp = C.VK_BLEND_OP_PLUS_DARKER_EXT
	BlendOpMinus            BlendOp = C.VK_BLEND_OP_MINUS_EXT
	BlendOpMinusClamped     BlendOp = C.VK_BLEND_OP_MINUS_CLAMPED_EXT
	BlendOpContrast         BlendOp = C.VK_BLEND_OP_CONTRAST_EXT
	BlendOpInvertOVG        BlendOp = C.VK_BLEND_OP_INVERT_OVG_EXT
	BlendOpRed              BlendOp = C.VK_BLEND_OP_RED_EXT
	BlendOpGreen            BlendOp = C.VK_BLEND_OP_GREEN_EXT
	BlendOpBlue             BlendOp = C.VK_BLEND_OP_BLUE_EXT
)

var blendOpToString = map[BlendOp]string{
	BlendOpAdd:              "Add",
	BlendOpSubtract:         "Subtract",
	BlendOpMin:              "Min",
	BlendOpMax:              "Max",
	BlendOpZero:             "Zero",
	BlendOpSrc:              "Src",
	BlendOpDst:              "Dst",
	BlendOpSrcOver:          "Src Over",
	BlendOpDstOver:          "Dst Over",
	BlendOpSrcIn:            "Src In",
	BlendOpDstIn:            "Dst In",
	BlendOpSrcOut:           "Src Out",
	BlendOpDstOut:           "Dst Out",
	BlendOpSrcAtop:          "Src Atop",
	BlendOpDstAtop:          "Dst Atop",
	BlendOpXor:              "Xor",
	BlendOpMultiply:         "Multiply",
	BlendOpScreen:           "Screen",
	BlendOpOverlay:          "Overlay",
	BlendOpDarken:           "Darken",
	BlendOpLighten:          "Lighten",
	BlendOpColorDodge:       "Color Dodge",
	BlendOpColorBurn:        "Color Burn",
	BlendOpHardLight:        "Hard Light",
	BlendOpSoftLight:        "Soft Light",
	BlendOpDifference:       "Difference",
	BlendOpExclusion:        "Exclusion",
	BlendOpInvert:           "Invert",
	BlendOpInvertRGB:        "Invert RGB",
	BlendOpLinearDodge:      "Linear Dodge",
	BlendOpLinearBurn:       "Linear Burn",
	BlendOpVividLight:       "Vivid Light",
	BlendOpLinearLight:      "Linear Light",
	BlendOpPinLight:         "Pin Light",
	BlendOpHardMix:          "Hard Mix",
	BlendOpHSLHue:           "Hue (HSL)",
	BlendOpHSLSaturation:    "Saturation (HSL)",
	BlendOpHSLColor:         "Color (HSL)",
	BlendOpHSLLuminosity:    "Luminosity (HSL)",
	BlendOpPlus:             "Plus",
	BlendOpPlusClamped:      "Plus Clamped",
	BlendOpPlusClampedAlpha: "Plus Clamped Alpha",
	BlendOpPlusDarker:       "Plus Darker",
	BlendOpMinus:            "Minus",
	BlendOpMinusClamped:     "Minus Clamped",
	BlendOpContrast:         "Contrast",
	BlendOpInvertOVG:        "Invert OVG",
	BlendOpRed:              "Red",
	BlendOpGreen:            "Green",
	BlendOpBlue:             "Blue",
}

func (o BlendOp) String() string {
	return blendOpToString[o]
}

type ColorComponents int32

const (
	ComponentRed   ColorComponents = C.VK_COLOR_COMPONENT_R_BIT
	ComponentGreen ColorComponents = C.VK_COLOR_COMPONENT_G_BIT
	ComponentBlue  ColorComponents = C.VK_COLOR_COMPONENT_B_BIT
	ComponentAlpha ColorComponents = C.VK_COLOR_COMPONENT_A_BIT
)

func (c ColorComponents) String() string {
	if c == 0 {
		return "None"
	}

	var sb strings.Builder
	if (c & ComponentRed) != 0 {
		sb.WriteRune('R')
	}
	if (c & ComponentGreen) != 0 {
		sb.WriteRune('G')
	}
	if (c & ComponentBlue) != 0 {
		sb.WriteRune('B')
	}
	if (c & ComponentAlpha) != 0 {
		sb.WriteRune('A')
	}

	return sb.String()
}
