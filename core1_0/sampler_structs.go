package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

const (
	// LodClampNone is a special constant value used for SamplerCreateInfo.MaxLod to indicate
	// that maximum LOD clamping should not be performed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_LOD_CLAMP_NONE.html
	LodClampNone float32 = C.VK_LOD_CLAMP_NONE

	// BorderColorFloatTransparentBlack specifies a transparent, floating-point format,
	// black color
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBorderColor.html
	BorderColorFloatTransparentBlack BorderColor = C.VK_BORDER_COLOR_FLOAT_TRANSPARENT_BLACK
	// BorderColorIntTransparentBlack specifies a transparent, integer format, black color
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBorderColor.html
	BorderColorIntTransparentBlack BorderColor = C.VK_BORDER_COLOR_INT_TRANSPARENT_BLACK
	// BorderColorFloatOpaqueBlack specifies an opaque, floating-point format, black color
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBorderColor.html
	BorderColorFloatOpaqueBlack BorderColor = C.VK_BORDER_COLOR_FLOAT_OPAQUE_BLACK
	// BorderColorIntOpaqueBlack specifies an opaque, integer format, black color
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBorderColor.html
	BorderColorIntOpaqueBlack BorderColor = C.VK_BORDER_COLOR_INT_OPAQUE_BLACK
	// BorderColorFloatOpaqueWhite specifies an opaque, floating-point format, white color
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBorderColor.html
	BorderColorFloatOpaqueWhite BorderColor = C.VK_BORDER_COLOR_FLOAT_OPAQUE_WHITE
	// BorderColorIntOpaqueWhite specifies an opaque, integer format, white color
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBorderColor.html
	BorderColorIntOpaqueWhite BorderColor = C.VK_BORDER_COLOR_INT_OPAQUE_WHITE

	// CompareOpNever specifies that the comparison always evaluates false
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html
	CompareOpNever CompareOp = C.VK_COMPARE_OP_NEVER
	// CompareOpLess specifies that the comparison evaluates reference < test
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html
	CompareOpLess CompareOp = C.VK_COMPARE_OP_LESS
	// CompareOpEqual specifies that the comparison evaluates reference == test
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html
	CompareOpEqual CompareOp = C.VK_COMPARE_OP_EQUAL
	// CompareOpLessOrEqual specifies that the comparison evaluates reference <= test
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html
	CompareOpLessOrEqual CompareOp = C.VK_COMPARE_OP_LESS_OR_EQUAL
	// CompareOpGreater specifies that the comparison evaluates reference > test
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html
	CompareOpGreater CompareOp = C.VK_COMPARE_OP_GREATER
	// CompareOpNotEqual specifies that the comparison evaluates reference != test
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html
	CompareOpNotEqual CompareOp = C.VK_COMPARE_OP_NOT_EQUAL
	// CompareOpGreaterOrEqual specifies that the comparison evaluates reference >= test
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html
	CompareOpGreaterOrEqual CompareOp = C.VK_COMPARE_OP_GREATER_OR_EQUAL
	// CompareOpAlways specifies that the comparison always evaluates true
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html
	CompareOpAlways CompareOp = C.VK_COMPARE_OP_ALWAYS

	// FilterNearest specifies nearest filtering
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilter.html
	FilterNearest Filter = C.VK_FILTER_NEAREST
	// FilterLinear specifies linear filtering
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilter.html
	FilterLinear Filter = C.VK_FILTER_LINEAR

	// SamplerMipmapModeNearest specifies nearest filtering
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerMipmapMode.html
	SamplerMipmapModeNearest SamplerMipmapMode = C.VK_SAMPLER_MIPMAP_MODE_NEAREST
	// SamplerMipmapModeLinear specifiest linear filtering
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerMipmapMode.html
	SamplerMipmapModeLinear SamplerMipmapMode = C.VK_SAMPLER_MIPMAP_MODE_LINEAR

	// SamplerAddressModeRepeat specifies that the repeat wrap mode will be used
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html
	SamplerAddressModeRepeat SamplerAddressMode = C.VK_SAMPLER_ADDRESS_MODE_REPEAT
	// SamplerAddressModeMirroredRepeat specifies that the mirrored repeat wrap mode will be used
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html
	SamplerAddressModeMirroredRepeat SamplerAddressMode = C.VK_SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT
	// SamplerAddressModeClampToEdge specifies that the clamp-to-edge wrap mode will be used
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html
	SamplerAddressModeClampToEdge SamplerAddressMode = C.VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE
	// SamplerAddressModeClampToBorder specifies that the clamp-to-border wrap mode will be used
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html
	SamplerAddressModeClampToBorder SamplerAddressMode = C.VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER
)

func init() {
	BorderColorFloatTransparentBlack.Register("Transparent Black - Float")
	BorderColorIntTransparentBlack.Register("Transparent Black - Int")
	BorderColorFloatOpaqueBlack.Register("Opaque Black - Float")
	BorderColorIntOpaqueBlack.Register("Opaque Black - Int")
	BorderColorFloatOpaqueWhite.Register("Opaque White - Float")
	BorderColorIntOpaqueWhite.Register("Opaque White - Int")

	CompareOpNever.Register("Never")
	CompareOpLess.Register("Less Than")
	CompareOpEqual.Register("Equal")
	CompareOpLessOrEqual.Register("Less Than Or Equal")
	CompareOpGreater.Register("Greater Than")
	CompareOpNotEqual.Register("Not Equal")
	CompareOpGreaterOrEqual.Register("Greater Than Or Equal")
	CompareOpAlways.Register("Always")

	FilterNearest.Register("Nearest")
	FilterLinear.Register("Linear")

	SamplerMipmapModeNearest.Register("Nearest")
	SamplerMipmapModeLinear.Register("Linear")

	SamplerAddressModeRepeat.Register("Repeat")
	SamplerAddressModeMirroredRepeat.Register("Mirrored Repeat")
	SamplerAddressModeClampToEdge.Register("Clamp to Edge")
	SamplerAddressModeClampToBorder.Register("Clamp to Border")
}

// SamplerCreateInfo specifies parameters of a newly-created Sampler
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html
type SamplerCreateInfo struct {
	// Flags describes additional parameters of the Sampler
	Flags SamplerCreateFlags
	// MagFilter specifies the magnification filter to apply to lookups
	MagFilter Filter
	// MinFilter specifies the minification filter to apply to lookups
	MinFilter Filter
	// MipmapMode specifies the mipmap filter to apply to lookups
	MipmapMode SamplerMipmapMode
	// AddressModeU specifies the addressing mode for U coordinates outside [0,1)
	AddressModeU SamplerAddressMode
	// AddressModeV specifies the addressing mode for V coordinates outside [0,1)
	AddressModeV SamplerAddressMode
	// AddressModeW specifies the addressing mode for W coordinates outside [0,1)
	AddressModeW SamplerAddressMode

	// MipLodBias is the bias to be added to mipmap level-of-detail calculations and bias provided
	// by Image sampling functions
	MipLodBias float32
	// MinLod is used to clamp the minimum of the computed level-of-detail value
	MinLod float32
	// MaxLod is used to clamp the maximum of the computed level-of-detail value- to avoid
	// clamping the maximum value, set MaxLod to the constant LodClampNone
	MaxLod float32

	// AnisotropyEnable is true to enable anisotropic filtering
	AnisotropyEnable bool
	// MaxAnisotropy is the anisotropy value clamp used by the Sampler when AnisotropyEnable is true
	MaxAnisotropy float32

	// CompareEnable is true to enable comparison against a reference value during lookups, or
	// false otherwise
	CompareEnable bool
	// CompareOp specifies the comparison operator to apply to fetched data before filtering
	CompareOp CompareOp

	// BorderColor specifies the predefined border color to use
	BorderColor BorderColor
	// UnnormalizedCoordinates controls whether to use unnormalized or normalized texel coordinates
	// to address texels of the image. When set to true, the range of the Image coordinates used
	// to lookup the texel is in the range of 0 to the Image size in each dimension. When set to
	// false, the range of Image coordinates is 0..1
	UnnormalizedCoordinates bool

	common.NextOptions
}

func (o SamplerCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkSamplerCreateInfo)
	}
	createInfo := (*C.VkSamplerCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_SAMPLER_CREATE_INFO
	createInfo.pNext = next
	createInfo.flags = C.VkSamplerCreateFlags(o.Flags)
	createInfo.magFilter = C.VkFilter(o.MagFilter)
	createInfo.minFilter = C.VkFilter(o.MinFilter)
	createInfo.mipmapMode = C.VkSamplerMipmapMode(o.MipmapMode)
	createInfo.addressModeU = C.VkSamplerAddressMode(o.AddressModeU)
	createInfo.addressModeV = C.VkSamplerAddressMode(o.AddressModeV)
	createInfo.addressModeW = C.VkSamplerAddressMode(o.AddressModeW)
	createInfo.mipLodBias = C.float(o.MipLodBias)
	createInfo.anisotropyEnable = C.VK_FALSE
	if o.AnisotropyEnable {
		createInfo.anisotropyEnable = C.VK_TRUE
	}
	createInfo.maxAnisotropy = C.float(o.MaxAnisotropy)
	createInfo.compareEnable = C.VK_FALSE
	if o.CompareEnable {
		createInfo.compareEnable = C.VK_TRUE
	}
	createInfo.compareOp = C.VkCompareOp(o.CompareOp)
	createInfo.minLod = C.float(o.MinLod)
	createInfo.maxLod = C.float(o.MaxLod)
	createInfo.borderColor = C.VkBorderColor(o.BorderColor)
	createInfo.unnormalizedCoordinates = C.VK_FALSE
	if o.UnnormalizedCoordinates {
		createInfo.unnormalizedCoordinates = C.VK_TRUE
	}

	return preallocatedPointer, nil
}
