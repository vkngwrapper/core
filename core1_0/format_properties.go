package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

const (
	FormatFeatureSampledImage             common.FormatFeatures = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT
	FormatFeatureStorageImage             common.FormatFeatures = C.VK_FORMAT_FEATURE_STORAGE_IMAGE_BIT
	FormatFeatureStorageImageAtomic       common.FormatFeatures = C.VK_FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT
	FormatFeatureUniformTexelBuffer       common.FormatFeatures = C.VK_FORMAT_FEATURE_UNIFORM_TEXEL_BUFFER_BIT
	FormatFeatureStorageTexelBuffer       common.FormatFeatures = C.VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT
	FormatFeatureStorageTexelBufferAtomic common.FormatFeatures = C.VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT
	FormatFeatureVertexBuffer             common.FormatFeatures = C.VK_FORMAT_FEATURE_VERTEX_BUFFER_BIT
	FormatFeatureColorAttachment          common.FormatFeatures = C.VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT
	FormatFeatureColorAttachmentBlend     common.FormatFeatures = C.VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT
	FormatFeatureDepthStencilAttachment   common.FormatFeatures = C.VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT
	FormatFeatureBlitSource               common.FormatFeatures = C.VK_FORMAT_FEATURE_BLIT_SRC_BIT
	FormatFeatureBlitDestination          common.FormatFeatures = C.VK_FORMAT_FEATURE_BLIT_DST_BIT
	FormatFeatureSampledImageFilterLinear common.FormatFeatures = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT
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
	LinearTilingFeatures  common.FormatFeatures
	OptimalTilingFeatures common.FormatFeatures
	BufferFeatures        common.FormatFeatures
}
