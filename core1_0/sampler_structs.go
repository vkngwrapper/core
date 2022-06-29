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
	BorderColorFloatTransparentBlack BorderColor = C.VK_BORDER_COLOR_FLOAT_TRANSPARENT_BLACK
	BorderColorIntTransparentBlack   BorderColor = C.VK_BORDER_COLOR_INT_TRANSPARENT_BLACK
	BorderColorFloatOpaqueBlack      BorderColor = C.VK_BORDER_COLOR_FLOAT_OPAQUE_BLACK
	BorderColorIntOpaqueBlack        BorderColor = C.VK_BORDER_COLOR_INT_OPAQUE_BLACK
	BorderColorFloatOpaqueWhite      BorderColor = C.VK_BORDER_COLOR_FLOAT_OPAQUE_WHITE
	BorderColorIntOpaqueWhite        BorderColor = C.VK_BORDER_COLOR_INT_OPAQUE_WHITE

	CompareNever          CompareOp = C.VK_COMPARE_OP_NEVER
	CompareLess           CompareOp = C.VK_COMPARE_OP_LESS
	CompareEqual          CompareOp = C.VK_COMPARE_OP_EQUAL
	CompareLessOrEqual    CompareOp = C.VK_COMPARE_OP_LESS_OR_EQUAL
	CompareGreater        CompareOp = C.VK_COMPARE_OP_GREATER
	CompareNotEqual       CompareOp = C.VK_COMPARE_OP_NOT_EQUAL
	CompareGreaterOrEqual CompareOp = C.VK_COMPARE_OP_GREATER_OR_EQUAL
	CompareAlways         CompareOp = C.VK_COMPARE_OP_ALWAYS

	FilterNearest Filter = C.VK_FILTER_NEAREST
	FilterLinear  Filter = C.VK_FILTER_LINEAR

	MipmapNearest MipmapMode = C.VK_SAMPLER_MIPMAP_MODE_NEAREST
	MipmapLinear  MipmapMode = C.VK_SAMPLER_MIPMAP_MODE_LINEAR

	SamplerAddressModeRepeat         SamplerAddressMode = C.VK_SAMPLER_ADDRESS_MODE_REPEAT
	SamplerAddressModeMirroredRepeat SamplerAddressMode = C.VK_SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT
	SamplerAddressModeClampToEdge    SamplerAddressMode = C.VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE
	SamplerAddressModeClampToBorder  SamplerAddressMode = C.VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER
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

type SamplerCreateOptions struct {
	Flags        SamplerCreateFlags
	MagFilter    Filter
	MinFilter    Filter
	MipmapMode   MipmapMode
	AddressModeU SamplerAddressMode
	AddressModeV SamplerAddressMode
	AddressModeW SamplerAddressMode

	MipLodBias float32
	MinLod     float32
	MaxLod     float32

	AnisotropyEnable bool
	MaxAnisotropy    float32

	CompareEnable bool
	CompareOp     CompareOp

	BorderColor             BorderColor
	UnnormalizedCoordinates bool

	common.HaveNext
}

func (o SamplerCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

func (o SamplerCreateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkSamplerCreateInfo)(cDataPointer)
	return createInfo.pNext, nil
}
