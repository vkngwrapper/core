package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

const (
	// FormatFeatureSampledImage specifies that an ImageView can be sampled from
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureSampledImage FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT
	// FormatFeatureStorageImage specifies that an ImageView can be used as a storage Image
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureStorageImage FormatFeatureFlags = C.VK_FORMAT_FEATURE_STORAGE_IMAGE_BIT
	// FormatFeatureStorageImageAtomic specifies that an ImageView can be used as a storage Image
	// that supports atomic operations
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureStorageImageAtomic FormatFeatureFlags = C.VK_FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT
	// FormatFeatureUniformTexelBuffer specifies that the format can be used to create a BufferView
	// that can be bound to a uniform texel Buffer descriptor
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureUniformTexelBuffer FormatFeatureFlags = C.VK_FORMAT_FEATURE_UNIFORM_TEXEL_BUFFER_BIT
	// FormatFeatureStorageTexelBuffer specifies that the format can be used to create a BufferView
	// that can be bound to a storage texel Buffer descriptor
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureStorageTexelBuffer FormatFeatureFlags = C.VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT
	// FormatFeatureStorageTexelBufferAtomic specifies that atomic operations are supported on
	// storage texel Buffer objects with this format
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureStorageTexelBufferAtomic FormatFeatureFlags = C.VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT
	// FormatFeatureVertexBuffer specifies that the format can be used as a vertex attribute
	// format
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureVertexBuffer FormatFeatureFlags = C.VK_FORMAT_FEATURE_VERTEX_BUFFER_BIT
	// FormatFeatureColorAttachment specifies that an ImageView can be used as a Framebuffer color
	// attachment and as an input attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureColorAttachment FormatFeatureFlags = C.VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT
	// FormatFeatureColorAttachmentBlend specifies that an ImageView can be used as a Framebuffer
	// color attachment that supports blending and as an input attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureColorAttachmentBlend FormatFeatureFlags = C.VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT
	// FormatFeatureDepthStencilAttachment specifies that an ImageView can be used as a Framebuffer
	// depth/stencil attachment and as an input attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureDepthStencilAttachment FormatFeatureFlags = C.VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT
	// FormatFeatureBlitSource specifies that an Image can be used as a source Image in a blit
	// operation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureBlitSource FormatFeatureFlags = C.VK_FORMAT_FEATURE_BLIT_SRC_BIT
	// FormatFeatureBlitDestination specifies that an Image can be used as a destination Image in a
	// blit operation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureBlitDestination FormatFeatureFlags = C.VK_FORMAT_FEATURE_BLIT_DST_BIT
	// FormatFeatureSampledImageFilterLinear specifies that an ImageView can be used with a Sampler
	// that has either of magFilter or minFilter set to FilterLinear, or mipmapMode set to
	// SamplerMipmapModeLinear
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureSampledImageFilterLinear FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT
)

func init() {
	FormatFeatureSampledImage.Register("Sampled Image")
	FormatFeatureStorageImage.Register("Storage Image")
	FormatFeatureStorageImageAtomic.Register("Storage Image, Atomic")
	FormatFeatureUniformTexelBuffer.Register("Uniform Texel Buffer")
	FormatFeatureStorageTexelBuffer.Register("Storage Texel Buffer")
	FormatFeatureStorageTexelBufferAtomic.Register("Storage Texel Buffer, Atomic")
	FormatFeatureVertexBuffer.Register("Vertex Buffer")
	FormatFeatureColorAttachment.Register("Color Attachment")
	FormatFeatureColorAttachmentBlend.Register("Color Attachment Blend")
	FormatFeatureDepthStencilAttachment.Register("Depth Stencil Attachment")
	FormatFeatureBlitSource.Register("Blit Source")
	FormatFeatureBlitDestination.Register("Blit Destination")
	FormatFeatureSampledImageFilterLinear.Register("Sampled Image, Linear Filter")
}

// FormatProperties specifies Image format properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties.html
type FormatProperties struct {
	// LinearTilingFeatures specifies features supported by Image objects created with a tiling
	// parameter of ImageTilingLinear
	LinearTilingFeatures FormatFeatureFlags
	// OptimalTilingFeatures specifies features supported by Image objects created with a tiling
	// parameter of ImageTilingOptimal
	OptimalTilingFeatures FormatFeatureFlags
	// BufferFeatures specifies features supported by Buffer objects
	BufferFeatures FormatFeatureFlags
}
