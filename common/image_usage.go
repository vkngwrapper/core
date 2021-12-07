package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "strings"

type ImageUsages int32

const (
	ImageUsageTransferSrc                      ImageUsages = C.VK_IMAGE_USAGE_TRANSFER_SRC_BIT
	ImageUsageTransferDst                      ImageUsages = C.VK_IMAGE_USAGE_TRANSFER_DST_BIT
	ImageUsageSampled                          ImageUsages = C.VK_IMAGE_USAGE_SAMPLED_BIT
	ImageUsageStorage                          ImageUsages = C.VK_IMAGE_USAGE_STORAGE_BIT
	ImageUsageColorAttachment                  ImageUsages = C.VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT
	ImageUsageDepthStencilAttachment           ImageUsages = C.VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT
	ImageUsageTransientAttachment              ImageUsages = C.VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT
	ImageUsageInputAttachment                  ImageUsages = C.VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT
	ImageUsageFragmentDensityMapEXT            ImageUsages = C.VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT
	ImageUsageFragmentShadingRateAttachmentKHR ImageUsages = C.VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR
)

var imageUsageToString = map[ImageUsages]string{
	ImageUsageTransferSrc:                      "Transfer Source",
	ImageUsageTransferDst:                      "Transfer Destination",
	ImageUsageSampled:                          "Sampled",
	ImageUsageStorage:                          "Storage",
	ImageUsageColorAttachment:                  "Color Attachment",
	ImageUsageDepthStencilAttachment:           "Depth Stencil Attachment",
	ImageUsageTransientAttachment:              "Transient Attachment",
	ImageUsageInputAttachment:                  "Input Attachment",
	ImageUsageFragmentDensityMapEXT:            "Fragment Density Map (Extension)",
	ImageUsageFragmentShadingRateAttachmentKHR: "Fragment Shading Rate Attachment (Khronos Extension)",
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
