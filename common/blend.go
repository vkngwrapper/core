package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
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
	BlendOpAdd                 BlendOp = C.VK_BLEND_OP_ADD
	BlendOpSubtract            BlendOp = C.VK_BLEND_OP_SUBTRACT
	BlendOpMin                 BlendOp = C.VK_BLEND_OP_MIN
	BlendOpMax                 BlendOp = C.VK_BLEND_OP_MAX
	BlendOpZeroEXT             BlendOp = C.VK_BLEND_OP_ZERO_EXT
	BlendOpSrcEXT              BlendOp = C.VK_BLEND_OP_SRC_EXT
	BlendOpDstEXT              BlendOp = C.VK_BLEND_OP_DST_EXT
	BlendOpSrcOverEXT          BlendOp = C.VK_BLEND_OP_SRC_OVER_EXT
	BlendOpDstOverEXT          BlendOp = C.VK_BLEND_OP_DST_OVER_EXT
	BlendOpSrcInEXT            BlendOp = C.VK_BLEND_OP_SRC_IN_EXT
	BlendOpDstInEXT            BlendOp = C.VK_BLEND_OP_DST_IN_EXT
	BlendOpSrcOutEXT           BlendOp = C.VK_BLEND_OP_SRC_OUT_EXT
	BlendOpDstOutEXT           BlendOp = C.VK_BLEND_OP_DST_OUT_EXT
	BlendOpSrcAtopEXT          BlendOp = C.VK_BLEND_OP_SRC_ATOP_EXT
	BlendOpDstAtopEXT          BlendOp = C.VK_BLEND_OP_DST_ATOP_EXT
	BlendOpXorEXT              BlendOp = C.VK_BLEND_OP_XOR_EXT
	BlendOpMultiplyEXT         BlendOp = C.VK_BLEND_OP_MULTIPLY_EXT
	BlendOpScreenEXT           BlendOp = C.VK_BLEND_OP_SCREEN_EXT
	BlendOpOverlayEXT          BlendOp = C.VK_BLEND_OP_OVERLAY_EXT
	BlendOpDarkenEXT           BlendOp = C.VK_BLEND_OP_DARKEN_EXT
	BlendOpLightenEXT          BlendOp = C.VK_BLEND_OP_LIGHTEN_EXT
	BlendOpColorDodgeEXT       BlendOp = C.VK_BLEND_OP_COLORDODGE_EXT
	BlendOpColorBurnEXT        BlendOp = C.VK_BLEND_OP_COLORBURN_EXT
	BlendOpHardLightEXT        BlendOp = C.VK_BLEND_OP_HARDLIGHT_EXT
	BlendOpSoftLightEXT        BlendOp = C.VK_BLEND_OP_SOFTLIGHT_EXT
	BlendOpDifferenceEXT       BlendOp = C.VK_BLEND_OP_DIFFERENCE_EXT
	BlendOpExclusionEXT        BlendOp = C.VK_BLEND_OP_EXCLUSION_EXT
	BlendOpInvertEXT           BlendOp = C.VK_BLEND_OP_INVERT_EXT
	BlendOpInvertRGBEXT        BlendOp = C.VK_BLEND_OP_INVERT_RGB_EXT
	BlendOpLinearDodgeEXT      BlendOp = C.VK_BLEND_OP_LINEARDODGE_EXT
	BlendOpLinearBurnEXT       BlendOp = C.VK_BLEND_OP_LINEARBURN_EXT
	BlendOpVividLightEXT       BlendOp = C.VK_BLEND_OP_VIVIDLIGHT_EXT
	BlendOpLinearLightEXT      BlendOp = C.VK_BLEND_OP_LINEARLIGHT_EXT
	BlendOpPinLightEXT         BlendOp = C.VK_BLEND_OP_PINLIGHT_EXT
	BlendOpHardMixEXT          BlendOp = C.VK_BLEND_OP_HARDMIX_EXT
	BlendOpHSLHueEXT           BlendOp = C.VK_BLEND_OP_HSL_HUE_EXT
	BlendOpHSLSaturationEXT    BlendOp = C.VK_BLEND_OP_HSL_SATURATION_EXT
	BlendOpHSLColorEXT         BlendOp = C.VK_BLEND_OP_HSL_COLOR_EXT
	BlendOpHSLLuminosityEXT    BlendOp = C.VK_BLEND_OP_HSL_LUMINOSITY_EXT
	BlendOpPlusEXT             BlendOp = C.VK_BLEND_OP_PLUS_EXT
	BlendOpPlusClampedEXT      BlendOp = C.VK_BLEND_OP_PLUS_CLAMPED_EXT
	BlendOpPlusClampedAlphaEXT BlendOp = C.VK_BLEND_OP_PLUS_CLAMPED_ALPHA_EXT
	BlendOpPlusDarkerEXT       BlendOp = C.VK_BLEND_OP_PLUS_DARKER_EXT
	BlendOpMinusEXT            BlendOp = C.VK_BLEND_OP_MINUS_EXT
	BlendOpMinusClampedEXT     BlendOp = C.VK_BLEND_OP_MINUS_CLAMPED_EXT
	BlendOpContrastEXT         BlendOp = C.VK_BLEND_OP_CONTRAST_EXT
	BlendOpInvertOVGEXT        BlendOp = C.VK_BLEND_OP_INVERT_OVG_EXT
	BlendOpRedEXT              BlendOp = C.VK_BLEND_OP_RED_EXT
	BlendOpGreenEXT            BlendOp = C.VK_BLEND_OP_GREEN_EXT
	BlendOpBlueEXT             BlendOp = C.VK_BLEND_OP_BLUE_EXT
)

var blendOpToString = map[BlendOp]string{
	BlendOpAdd:                 "Add",
	BlendOpSubtract:            "Subtract",
	BlendOpMin:                 "Min",
	BlendOpMax:                 "Max",
	BlendOpZeroEXT:             "Zero (Extension)",
	BlendOpSrcEXT:              "Src (Extension)",
	BlendOpDstEXT:              "Dst (Extension)",
	BlendOpSrcOverEXT:          "Src Over (Extension)",
	BlendOpDstOverEXT:          "Dst Over (Extension)",
	BlendOpSrcInEXT:            "Src In (Extension)",
	BlendOpDstInEXT:            "Dst In (Extension)",
	BlendOpSrcOutEXT:           "Src Out (Extension)",
	BlendOpDstOutEXT:           "Dst Out (Extension)",
	BlendOpSrcAtopEXT:          "Src Atop (Extension)",
	BlendOpDstAtopEXT:          "Dst Atop (Extension)",
	BlendOpXorEXT:              "Xor (Extension)",
	BlendOpMultiplyEXT:         "Multiply (Extension)",
	BlendOpScreenEXT:           "Screen (Extension)",
	BlendOpOverlayEXT:          "Overlay (Extension)",
	BlendOpDarkenEXT:           "Darken (Extension)",
	BlendOpLightenEXT:          "Lighten (Extension)",
	BlendOpColorDodgeEXT:       "Color Dodge (Extension)",
	BlendOpColorBurnEXT:        "Color Burn (Extension)",
	BlendOpHardLightEXT:        "Hard Light (Extension)",
	BlendOpSoftLightEXT:        "Soft Light (Extension)",
	BlendOpDifferenceEXT:       "Difference (Extension)",
	BlendOpExclusionEXT:        "Exclusion (Extension)",
	BlendOpInvertEXT:           "Invert (Extension)",
	BlendOpInvertRGBEXT:        "Invert RGB (Extension)",
	BlendOpLinearDodgeEXT:      "Linear Dodge (Extension)",
	BlendOpLinearBurnEXT:       "Linear Burn (Extension)",
	BlendOpVividLightEXT:       "Vivid Light (Extension)",
	BlendOpLinearLightEXT:      "Linear Light (Extension)",
	BlendOpPinLightEXT:         "Pin Light (Extension)",
	BlendOpHardMixEXT:          "Hard Mix (Extension)",
	BlendOpHSLHueEXT:           "Hue (HSL) (Extension)",
	BlendOpHSLSaturationEXT:    "Saturation (HSL) (Extension)",
	BlendOpHSLColorEXT:         "Color (HSL) (Extension)",
	BlendOpHSLLuminosityEXT:    "Luminosity (HSL) (Extension)",
	BlendOpPlusEXT:             "Plus (Extension)",
	BlendOpPlusClampedEXT:      "Plus Clamped (Extension)",
	BlendOpPlusClampedAlphaEXT: "Plus Clamped Alpha (Extension)",
	BlendOpPlusDarkerEXT:       "Plus Darker (Extension)",
	BlendOpMinusEXT:            "Minus (Extension)",
	BlendOpMinusClampedEXT:     "Minus Clamped (Extension)",
	BlendOpContrastEXT:         "Contrast (Extension)",
	BlendOpInvertOVGEXT:        "Invert OVG (Extension)",
	BlendOpRedEXT:              "Red (Extension)",
	BlendOpGreenEXT:            "Green (Extension)",
	BlendOpBlueEXT:             "Blue (Extension)",
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
