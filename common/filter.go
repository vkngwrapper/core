package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type Filter int32

const (
	FilterNearest = C.VK_FILTER_NEAREST
	FilterLinear  = C.VK_FILTER_LINEAR
	FilterCubic   = C.VK_FILTER_CUBIC_IMG
)

var filterToString = map[Filter]string{
	FilterNearest: "Nearest",
	FilterLinear:  "Linear",
	FilterCubic:   "Cubic",
}

func (f Filter) String() string {
	return filterToString[f]
}

type MipmapMode int32

const (
	MipmapNearest = C.VK_SAMPLER_MIPMAP_MODE_NEAREST
	MipmapLinear  = C.VK_SAMPLER_MIPMAP_MODE_LINEAR
)

var mipmapModeToString = map[MipmapMode]string{
	MipmapNearest: "Nearest",
	MipmapLinear:  "Linear",
}

func (m MipmapMode) String() string {
	return mipmapModeToString[m]
}

type SamplerAddressMode int32

const (
	AddressModeRepeat            = C.VK_SAMPLER_ADDRESS_MODE_REPEAT
	AddressModeMirroredRepeat    = C.VK_SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT
	AddressModeClampToEdge       = C.VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE
	AddressModeClampToBorder     = C.VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER
	AddressModeMirrorClampToEdge = C.VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE
)

var samplerAddressModeToString = map[SamplerAddressMode]string{
	AddressModeRepeat:            "Repeat",
	AddressModeMirroredRepeat:    "Mirrored Repeat",
	AddressModeClampToEdge:       "Clamp to Edge",
	AddressModeClampToBorder:     "Clamp to Border",
	AddressModeMirrorClampToEdge: "Mirrored Clamp To Edge",
}

func (m SamplerAddressMode) String() string {
	return samplerAddressModeToString[m]
}
