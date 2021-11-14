package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "strings"

type ImageUsages int32

const (
	ImageTransferSrc                      ImageUsages = C.VK_IMAGE_USAGE_TRANSFER_SRC_BIT
	ImageTransferDest                     ImageUsages = C.VK_IMAGE_USAGE_TRANSFER_DST_BIT
	ImageSampled                          ImageUsages = C.VK_IMAGE_USAGE_SAMPLED_BIT
	ImageStorage                          ImageUsages = C.VK_IMAGE_USAGE_STORAGE_BIT
	ImageColorAttachment                  ImageUsages = C.VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT
	ImageDepthStencilAttachment           ImageUsages = C.VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT
	ImageTransientAttachment              ImageUsages = C.VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT
	ImageInputAttachment                  ImageUsages = C.VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT
	ImageFragmentDensityMapEXT            ImageUsages = C.VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT
	ImageFragmentShadingRateAttachmentKHR ImageUsages = C.VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR
)

var imageUsageToString = map[ImageUsages]string{
	ImageTransferSrc:                      "Transfer Source",
	ImageTransferDest:                     "Transfer Destination",
	ImageSampled:                          "Sampled",
	ImageStorage:                          "Storage",
	ImageColorAttachment:                  "Color Attachment",
	ImageDepthStencilAttachment:           "Depth Stencil Attachment",
	ImageTransientAttachment:              "Transient Attachment",
	ImageInputAttachment:                  "Input Attachment",
	ImageFragmentDensityMapEXT:            "Fragment Density Map (Extension)",
	ImageFragmentShadingRateAttachmentKHR: "Fragment Shading Rate Attachment (Khronos Extension)",
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
