package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/core1_0"

const (
	ImageAspectPlane0 core1_0.ImageAspectFlags = C.VK_IMAGE_ASPECT_PLANE_0_BIT
	ImageAspectPlane1 core1_0.ImageAspectFlags = C.VK_IMAGE_ASPECT_PLANE_1_BIT
	ImageAspectPlane2 core1_0.ImageAspectFlags = C.VK_IMAGE_ASPECT_PLANE_2_BIT

	ImageCreate2DArrayCompatible        core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT
	ImageCreateAlias                    core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_ALIAS_BIT
	ImageCreateBlockTexelViewCompatible core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT
	ImageCreateDisjoint                 core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_DISJOINT_BIT
	ImageCreateExtendedUsage            core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_EXTENDED_USAGE_BIT
	ImageCreateProtected                core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_PROTECTED_BIT
	ImageCreateSplitInstanceBindRegions core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT

	ImageLayoutDepthAttachmentStencilReadOnlyOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL
	ImageLayoutDepthReadOnlyStencilAttachmentOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL
)

func init() {
	ImageAspectPlane0.Register("Plane 0")
	ImageAspectPlane1.Register("Plane 1")
	ImageAspectPlane2.Register("Plane 2")

	ImageCreate2DArrayCompatible.Register("2D Array Compatible")
	ImageCreateAlias.Register("Alias")
	ImageCreateBlockTexelViewCompatible.Register("Block Texel View Compatible")
	ImageCreateDisjoint.Register("Disjoint")
	ImageCreateExtendedUsage.Register("Extended Usage")
	ImageCreateProtected.Register("Protected")
	ImageCreateSplitInstanceBindRegions.Register("Split Instance Bind Regions")

	ImageLayoutDepthReadOnlyStencilAttachmentOptimal.Register("Depth Read-Only Stencil Attachment Optimal")
	ImageLayoutDepthAttachmentStencilReadOnlyOptimal.Register("Depth Attachment Stencil Read-Only Optimal")

}
