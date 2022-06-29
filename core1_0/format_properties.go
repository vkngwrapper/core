package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

const (
	FormatFeatureSampledImage             FormatFeatures = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT
	FormatFeatureStorageImage             FormatFeatures = C.VK_FORMAT_FEATURE_STORAGE_IMAGE_BIT
	FormatFeatureStorageImageAtomic       FormatFeatures = C.VK_FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT
	FormatFeatureUniformTexelBuffer       FormatFeatures = C.VK_FORMAT_FEATURE_UNIFORM_TEXEL_BUFFER_BIT
	FormatFeatureStorageTexelBuffer       FormatFeatures = C.VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT
	FormatFeatureStorageTexelBufferAtomic FormatFeatures = C.VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT
	FormatFeatureVertexBuffer             FormatFeatures = C.VK_FORMAT_FEATURE_VERTEX_BUFFER_BIT
	FormatFeatureColorAttachment          FormatFeatures = C.VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT
	FormatFeatureColorAttachmentBlend     FormatFeatures = C.VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT
	FormatFeatureDepthStencilAttachment   FormatFeatures = C.VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT
	FormatFeatureBlitSource               FormatFeatures = C.VK_FORMAT_FEATURE_BLIT_SRC_BIT
	FormatFeatureBlitDestination          FormatFeatures = C.VK_FORMAT_FEATURE_BLIT_DST_BIT
	FormatFeatureSampledImageFilterLinear FormatFeatures = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT
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
	LinearTilingFeatures  FormatFeatures
	OptimalTilingFeatures FormatFeatures
	BufferFeatures        FormatFeatures
}
