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
	BorderColorFloatTransparentBlack common.BorderColor = C.VK_BORDER_COLOR_FLOAT_TRANSPARENT_BLACK
	BorderColorIntTransparentBlack   common.BorderColor = C.VK_BORDER_COLOR_INT_TRANSPARENT_BLACK
	BorderColorFloatOpaqueBlack      common.BorderColor = C.VK_BORDER_COLOR_FLOAT_OPAQUE_BLACK
	BorderColorIntOpaqueBlack        common.BorderColor = C.VK_BORDER_COLOR_INT_OPAQUE_BLACK
	BorderColorFloatOpaqueWhite      common.BorderColor = C.VK_BORDER_COLOR_FLOAT_OPAQUE_WHITE
	BorderColorIntOpaqueWhite        common.BorderColor = C.VK_BORDER_COLOR_INT_OPAQUE_WHITE

	CompareNever          common.CompareOp = C.VK_COMPARE_OP_NEVER
	CompareLess           common.CompareOp = C.VK_COMPARE_OP_LESS
	CompareEqual          common.CompareOp = C.VK_COMPARE_OP_EQUAL
	CompareLessOrEqual    common.CompareOp = C.VK_COMPARE_OP_LESS_OR_EQUAL
	CompareGreater        common.CompareOp = C.VK_COMPARE_OP_GREATER
	CompareNotEqual       common.CompareOp = C.VK_COMPARE_OP_NOT_EQUAL
	CompareGreaterOrEqual common.CompareOp = C.VK_COMPARE_OP_GREATER_OR_EQUAL
	CompareAlways         common.CompareOp = C.VK_COMPARE_OP_ALWAYS

	FilterNearest common.Filter = C.VK_FILTER_NEAREST
	FilterLinear  common.Filter = C.VK_FILTER_LINEAR

	MipmapNearest common.MipmapMode = C.VK_SAMPLER_MIPMAP_MODE_NEAREST
	MipmapLinear  common.MipmapMode = C.VK_SAMPLER_MIPMAP_MODE_LINEAR

	SamplerAddressModeRepeat         common.SamplerAddressMode = C.VK_SAMPLER_ADDRESS_MODE_REPEAT
	SamplerAddressModeMirroredRepeat common.SamplerAddressMode = C.VK_SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT
	SamplerAddressModeClampToEdge    common.SamplerAddressMode = C.VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE
	SamplerAddressModeClampToBorder  common.SamplerAddressMode = C.VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER
)

func init() {
	BorderColorFloatTransparentBlack.Register("Transparent Black - Float")
	BorderColorIntTransparentBlack.Register("Transparent Black - Int")
	BorderColorFloatOpaqueBlack.Register("Opaque Black - Float")
	BorderColorIntOpaqueBlack.Register("Opaque Black - Int")
	BorderColorFloatOpaqueWhite.Register("Opaque White - Float")
	BorderColorIntOpaqueWhite.Register("Opaque White - Int")

	CompareNever.Register("Never")
	CompareLess.Register("Less Than")
	CompareEqual.Register("Equal")
	CompareLessOrEqual.Register("Less Than Or Equal")
	CompareGreater.Register("Greater Than")
	CompareNotEqual.Register("Not Equal")
	CompareGreaterOrEqual.Register("Greater Than Or Equal")
	CompareAlways.Register("Always")

	FilterNearest.Register("Nearest")
	FilterLinear.Register("Linear")

	MipmapNearest.Register("Nearest")
	MipmapLinear.Register("Linear")

	SamplerAddressModeRepeat.Register("Repeat")
	SamplerAddressModeMirroredRepeat.Register("Mirrored Repeat")
	SamplerAddressModeClampToEdge.Register("Clamp to Edge")
	SamplerAddressModeClampToBorder.Register("Clamp to Border")
}

type SamplerOptions struct {
	Flags        common.SamplerCreateFlags
	MagFilter    common.Filter
	MinFilter    common.Filter
	MipmapMode   common.MipmapMode
	AddressModeU common.SamplerAddressMode
	AddressModeV common.SamplerAddressMode
	AddressModeW common.SamplerAddressMode

	MipLodBias float32
	MinLod     float32
	MaxLod     float32

	AnisotropyEnable bool
	MaxAnisotropy    float32

	CompareEnable bool
	CompareOp     common.CompareOp

	BorderColor             common.BorderColor
	UnnormalizedCoordinates bool

	common.HaveNext
}

func (o *SamplerOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
