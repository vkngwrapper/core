package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/core1_0"

const (
	// ImageAspectPlane0 specifies memory plane 0
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html
	ImageAspectPlane0 core1_0.ImageAspectFlags = C.VK_IMAGE_ASPECT_PLANE_0_BIT
	// ImageAspectPlane1 specifies memory plane 1
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html
	ImageAspectPlane1 core1_0.ImageAspectFlags = C.VK_IMAGE_ASPECT_PLANE_1_BIT
	// ImageAspectPlane2 specifies memory plane 2
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html
	ImageAspectPlane2 core1_0.ImageAspectFlags = C.VK_IMAGE_ASPECT_PLANE_2_BIT

	// ImageCreate2DArrayCompatible specifies that the Image can be used to create an ImageView of
	// type core1_0.ImageViewType2D or core1_0.ImageViewType2DArray
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
	ImageCreate2DArrayCompatible core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT
	// ImageCreateAlias specifies that two Image objects created with the same creation parameters
	// and aliased to the same memory can interpret the contents of the memory consistently each
	// other
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
	ImageCreateAlias core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_ALIAS_BIT
	// ImageCreateBlockTexelViewCompatible specifies that the Image having a compressed format can be
	// used to create an ImageView with an uncompressed format where each texel in the ImageView
	// corresponds to a compressed texel block of the Image
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
	ImageCreateBlockTexelViewCompatible core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT
	// ImageCreateDisjoint specifies that an Image with a multi-planar format must have each plane
	// separately bound to memory, rather than having a single memory binding for the whole Image
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
	ImageCreateDisjoint core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_DISJOINT_BIT
	// ImageCreateExtendedUsage specifies that the Image can be created with usage flags that are not
	// supported for the format the Image is created with but are supported for at least one format
	// an ImageView created from this Image can have
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
	ImageCreateExtendedUsage core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_EXTENDED_USAGE_BIT
	// ImageCreateProtected specifies that the Image is a protected Image
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
	ImageCreateProtected core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_PROTECTED_BIT
	// ImageCreateSplitInstanceBindRegions specifies that the Image can be used with a non-empty
	// BindImageMemoryDeviceGroupInfo.SplitInstanceBindRegions
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
	ImageCreateSplitInstanceBindRegions core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT

	// ImageLayoutDepthAttachmentStencilReadOnlyOptimal specifies a layout for depth/stencil format
	// Image objects allowing read and write access to the depth aspect as a depth attachment, and read-only
	// access to the stencil aspect as a stencil attachment or in shaders as a sampled Image, combined
	// Image/Sampler, or input attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutDepthAttachmentStencilReadOnlyOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL
	// ImageLayoutDepthReadOnlyStencilAttachmentOptimal specifies a layout for depth/stencil format Image objects
	// allowing read and write access to the stencil aspect as a stencil attachment, and read-only access
	// to the depth aspect as a depth attachment or in shaders as a sampled Image, combined Image/Sampler,
	// or input attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
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
