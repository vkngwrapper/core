package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type ImageLayout int32

const (
	LayoutUndefined                               ImageLayout = C.VK_IMAGE_LAYOUT_UNDEFINED
	LayoutGeneral                                 ImageLayout = C.VK_IMAGE_LAYOUT_GENERAL
	LayoutColorAttachmentOptimal                  ImageLayout = C.VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
	LayoutDepthStencilAttachmentOptimal           ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL
	LayoutDepthStencilReadOnlyOptimal             ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL
	LayoutShaderReadOnlyOptimal                   ImageLayout = C.VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL
	LayoutTransferSrcOptimal                      ImageLayout = C.VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL
	LayoutTransferDstOptimal                      ImageLayout = C.VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL
	LayoutPreInitialized                          ImageLayout = C.VK_IMAGE_LAYOUT_PREINITIALIZED
	LayoutDepthReadOnlyStencilAttachmentOptimal   ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL
	LayoutDepthAttachmentStencilReadOnlyOptimal   ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL
	LayoutDepthAttachmentOptimal                  ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL
	LayoutDepthReadOnlyOptimal                    ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL
	LayoutStencilAttachmentOptimal                ImageLayout = C.VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL
	LayoutStencilReadOnlyOptimal                  ImageLayout = C.VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL
	LayoutPresentSrcKHR                           ImageLayout = C.VK_IMAGE_LAYOUT_PRESENT_SRC_KHR
	LayoutSharedPresentKHR                        ImageLayout = C.VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR
	LayoutFragmentDensityMapOptimalEXT            ImageLayout = C.VK_IMAGE_LAYOUT_FRAGMENT_DENSITY_MAP_OPTIMAL_EXT
	LayoutFragmentShadingRateAttachmentOptimalKHR ImageLayout = C.VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR
	LayoutReadOnlyOptimalKHR                      ImageLayout = C.VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR
	LayoutAttachmentOptimalKHR                    ImageLayout = C.VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR
)

var imageLayoutToString = map[ImageLayout]string{
	LayoutUndefined:                               "Undefined",
	LayoutGeneral:                                 "General",
	LayoutColorAttachmentOptimal:                  "Color Attachment",
	LayoutDepthStencilAttachmentOptimal:           "Depth & Stencil Attachment",
	LayoutDepthStencilReadOnlyOptimal:             "Depth & Stencil Read-Only",
	LayoutShaderReadOnlyOptimal:                   "Shader Read-Only",
	LayoutTransferSrcOptimal:                      "Transfer Source",
	LayoutTransferDstOptimal:                      "Transfer Destination",
	LayoutPreInitialized:                          "Pre-Initialized",
	LayoutDepthReadOnlyStencilAttachmentOptimal:   "Depth Read-Only & Stencil Attachment",
	LayoutDepthAttachmentStencilReadOnlyOptimal:   "Depth Attachment & Stencil Read-Only",
	LayoutDepthAttachmentOptimal:                  "Depth Attachment",
	LayoutDepthReadOnlyOptimal:                    "Depth Read-Only",
	LayoutStencilAttachmentOptimal:                "Stencil Attachment",
	LayoutStencilReadOnlyOptimal:                  "Stencil Read-Only",
	LayoutPresentSrcKHR:                           "Present Source (Khronos Extension)",
	LayoutSharedPresentKHR:                        "Shared Present (Khronos Extension)",
	LayoutFragmentDensityMapOptimalEXT:            "Fragment Density Map (Extension)",
	LayoutFragmentShadingRateAttachmentOptimalKHR: "Fragment Shading Rate Attachment (Khronos Extension)",
	LayoutReadOnlyOptimalKHR:                      "Read-Only (Khronos Extension)",
	LayoutAttachmentOptimalKHR:                    "Attachment (Khronos Extension)",
}

func (l ImageLayout) String() string {
	return imageLayoutToString[l]
}
