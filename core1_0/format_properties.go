package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

const (
	FormatFeatureSampledImage             FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT
	FormatFeatureStorageImage             FormatFeatureFlags = C.VK_FORMAT_FEATURE_STORAGE_IMAGE_BIT
	FormatFeatureStorageImageAtomic       FormatFeatureFlags = C.VK_FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT
	FormatFeatureUniformTexelBuffer       FormatFeatureFlags = C.VK_FORMAT_FEATURE_UNIFORM_TEXEL_BUFFER_BIT
	FormatFeatureStorageTexelBuffer       FormatFeatureFlags = C.VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT
	FormatFeatureStorageTexelBufferAtomic FormatFeatureFlags = C.VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT
	FormatFeatureVertexBuffer             FormatFeatureFlags = C.VK_FORMAT_FEATURE_VERTEX_BUFFER_BIT
	FormatFeatureColorAttachment          FormatFeatureFlags = C.VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT
	FormatFeatureColorAttachmentBlend     FormatFeatureFlags = C.VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT
	FormatFeatureDepthStencilAttachment   FormatFeatureFlags = C.VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT
	FormatFeatureBlitSource               FormatFeatureFlags = C.VK_FORMAT_FEATURE_BLIT_SRC_BIT
	FormatFeatureBlitDestination          FormatFeatureFlags = C.VK_FORMAT_FEATURE_BLIT_DST_BIT
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

type FormatProperties struct {
	LinearTilingFeatures  FormatFeatureFlags
	OptimalTilingFeatures FormatFeatureFlags
	BufferFeatures        FormatFeatureFlags
}
