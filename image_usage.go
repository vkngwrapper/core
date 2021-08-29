package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "strings"

type ImageUsage int32

const (
	UsageTransferSrc                   ImageUsage = C.VK_IMAGE_USAGE_TRANSFER_SRC_BIT
	UsageTransferDest                  ImageUsage = C.VK_IMAGE_USAGE_TRANSFER_DST_BIT
	UsageSampled                       ImageUsage = C.VK_IMAGE_USAGE_SAMPLED_BIT
	UsageStorage                       ImageUsage = C.VK_IMAGE_USAGE_STORAGE_BIT
	UsageColorAttachment               ImageUsage = C.VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT
	UsageDepthStencilAttachment        ImageUsage = C.VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT
	UsageTransientAttachment           ImageUsage = C.VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT
	UsageInputAttachment               ImageUsage = C.VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT
	UsageFragmentDensityMap            ImageUsage = C.VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT
	UsageFragmentShadingRateAttachment ImageUsage = C.VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR
	UsageAllUsages                     ImageUsage = UsageTransferSrc | UsageTransferDest | UsageSampled | UsageStorage | UsageColorAttachment |
		UsageDepthStencilAttachment | UsageTransientAttachment | UsageInputAttachment |
		UsageFragmentDensityMap | UsageFragmentShadingRateAttachment
)

var imageUsageToString = map[ImageUsage]string{
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
