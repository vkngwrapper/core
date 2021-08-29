package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "strings"

type ImageUsages int32

const (
	UsageTransferSrc                   ImageUsages = C.VK_IMAGE_USAGE_TRANSFER_SRC_BIT
	UsageTransferDest                  ImageUsages = C.VK_IMAGE_USAGE_TRANSFER_DST_BIT
	UsageSampled                       ImageUsages = C.VK_IMAGE_USAGE_SAMPLED_BIT
	UsageStorage                       ImageUsages = C.VK_IMAGE_USAGE_STORAGE_BIT
	UsageColorAttachment               ImageUsages = C.VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT
	UsageDepthStencilAttachment        ImageUsages = C.VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT
	UsageTransientAttachment           ImageUsages = C.VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT
	UsageInputAttachment               ImageUsages = C.VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT
	UsageFragmentDensityMap            ImageUsages = C.VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT
	UsageFragmentShadingRateAttachment ImageUsages = C.VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR
	UsageAllUsages                     ImageUsages = UsageTransferSrc | UsageTransferDest | UsageSampled | UsageStorage | UsageColorAttachment |
		UsageDepthStencilAttachment | UsageTransientAttachment | UsageInputAttachment |
		UsageFragmentDensityMap | UsageFragmentShadingRateAttachment
)

var imageUsageToString = map[ImageUsages]string{
	UsageTransferSrc:                   "Transfer Source",
	UsageTransferDest:                  "Transfer Destination",
	UsageSampled:                       "Sampled",
	UsageStorage:                       "Storage",
	UsageColorAttachment:               "Color Attachment",
	UsageDepthStencilAttachment:        "Depth Stencil Attachment",
	UsageTransientAttachment:           "Transient Attachment",
	UsageInputAttachment:               "Input Attachment",
	UsageFragmentDensityMap:            "Fragment Density Map",
	UsageFragmentShadingRateAttachment: "Fragment Shading Rate Attachment",
}

func (u ImageUsages) String() string {
	if u == 0 {
		return "None"
	}

	hasOne := false
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		shiftedBit := ImageUsages(1 << i)
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
