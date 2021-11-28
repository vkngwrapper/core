package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"strings"
	"unsafe"
)

type vulkanSampler struct {
	device VkDevice
	driver Driver
	handle VkSampler
}

func (s *vulkanSampler) Handle() VkSampler {
	return s.handle
}

func (s *vulkanSampler) Destroy() {
	s.driver.VkDestroySampler(s.device, s.handle, nil)
}

type SamplerFlags int32

const (
	SamplerSubsampledEXT                     = C.VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT
	SamplerSubsampledCoarseReconstructionEXT = C.VK_SAMPLER_CREATE_SUBSAMPLED_COARSE_RECONSTRUCTION_BIT_EXT
)

var samplerFlagsToString = map[SamplerFlags]string{
	SamplerSubsampledEXT:                     "Subsampled (Extension)",
	SamplerSubsampledCoarseReconstructionEXT: "Subsampled Coarse Reconstruction (Extension)",
}

func (f SamplerFlags) String() string {
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := SamplerFlags(1 << i)
		if (f & checkBit) != 0 {
			if hasOne {
				sb.WriteString("|")
			}
			sb.WriteString(samplerFlagsToString[checkBit])
			hasOne = true
		}
	}

	return sb.String()
}

type SamplerOptions struct {
	Flags        SamplerFlags
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

func (o *SamplerOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkSamplerCreateInfo)(allocator.Malloc(C.sizeof_struct_VkSamplerCreateInfo))
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

	return unsafe.Pointer(createInfo), nil
}
