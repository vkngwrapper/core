package core

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

type vulkanSampler struct {
	device VkDevice
	driver Driver
	handle VkSampler
}

func (s *vulkanSampler) Handle() VkSampler {
	return s.handle
}

func (s *vulkanSampler) Destroy(callbacks *AllocationCallbacks) {
	s.driver.VkDestroySampler(s.device, s.handle, callbacks.Handle())
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
	return common.FlagsToString(f, samplerFlagsToString)
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
