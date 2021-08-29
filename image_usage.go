package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "strings"

type ImageUsage int32

const (
	TransferSrc                   ImageUsage = C.VK_IMAGE_USAGE_TRANSFER_SRC_BIT
	TransferDest                  ImageUsage = C.VK_IMAGE_USAGE_TRANSFER_DST_BIT
	Sampled                       ImageUsage = C.VK_IMAGE_USAGE_SAMPLED_BIT
	Storage                       ImageUsage = C.VK_IMAGE_USAGE_STORAGE_BIT
	ColorAttachment               ImageUsage = C.VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT
	DepthStencilAttachment        ImageUsage = C.VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT
	TransientAttachment           ImageUsage = C.VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT
	InputAttachment               ImageUsage = C.VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT
	FragmentDensityMap            ImageUsage = C.VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT
	FragmentShadingRateAttachment ImageUsage = C.VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR
	AllUsages                     ImageUsage = TransferSrc | TransferDest | Sampled | Storage | ColorAttachment |
		DepthStencilAttachment | TransientAttachment | InputAttachment |
		FragmentDensityMap | FragmentShadingRateAttachment
)

var imageUsageToString = map[ImageUsage]string{
	TransferSrc:                   "Transfer Source",
	TransferDest:                  "Transfer Destination",
	Sampled:                       "Sampled",
	Storage:                       "Storage",
	ColorAttachment:               "Color Attachment",
	DepthStencilAttachment:        "Depth Stencil Attachment",
	TransientAttachment:           "Transient Attachment",
	InputAttachment:               "Input Attachment",
	FragmentDensityMap:            "Fragment Density Map",
	FragmentShadingRateAttachment: "Fragment Shading Rate Attachment",
}

func (u ImageUsage) String() string {
	hasOne := false
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		shiftedBit := ImageUsage(1 << i)
		if u&shiftedBit != 0 {
			strVal, exists := imageUsageToString[shiftedBit]
			if exists {
				if hasOne {
					sb.WriteString("|")
				}
				sb.WriteString(strVal)
				hasOne = true
			}
		}
	}

	return sb.String()
}
