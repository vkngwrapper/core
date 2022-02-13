package core1_0

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

type FormatFeatures int32

const (
	FormatFeatureSampledImage                                                     FormatFeatures = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT
	FormatFeatureStorageImage                                                     FormatFeatures = C.VK_FORMAT_FEATURE_STORAGE_IMAGE_BIT
	FormatFeatureStorageImageAtomic                                               FormatFeatures = C.VK_FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT
	FormatFeatureUniformTexelBuffer                                               FormatFeatures = C.VK_FORMAT_FEATURE_UNIFORM_TEXEL_BUFFER_BIT
	FormatFeatureStorageTexelBuffer                                               FormatFeatures = C.VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT
	FormatFeatureStorageTexelBufferAtomic                                         FormatFeatures = C.VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT
	FormatFeatureVertexBuffer                                                     FormatFeatures = C.VK_FORMAT_FEATURE_VERTEX_BUFFER_BIT
	FormatFeatureColorAttachment                                                  FormatFeatures = C.VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT
	FormatFeatureColorAttachmentBlend                                             FormatFeatures = C.VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT
	FormatFeatureDepthStencilAttachment                                           FormatFeatures = C.VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT
	FormatFeatureBlitSource                                                       FormatFeatures = C.VK_FORMAT_FEATURE_BLIT_SRC_BIT
	FormatFeatureBlitDestination                                                  FormatFeatures = C.VK_FORMAT_FEATURE_BLIT_DST_BIT
	FormatFeatureSampledImageFilterLinear                                         FormatFeatures = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT
	FormatFeatureTransferSource                                                   FormatFeatures = C.VK_FORMAT_FEATURE_TRANSFER_SRC_BIT
	FormatFeatureTransferDestination                                              FormatFeatures = C.VK_FORMAT_FEATURE_TRANSFER_DST_BIT
	FormatFeatureMidpointChromaSamples                                            FormatFeatures = C.VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT
	FormatFeatureSampledImageYcbcrConversionLinearFilter                          FormatFeatures = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT
	FormatFeatureSampledImageYcbcrConversionSeparateReconstructionFilter          FormatFeatures = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicit          FormatFeatures = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitForceable FormatFeatures = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT
	FormatFeatureDisjoint                                                         FormatFeatures = C.VK_FORMAT_FEATURE_DISJOINT_BIT
	FormatFeatureCositedChromaSamples                                             FormatFeatures = C.VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT
	FormatFeatureSampledImageFilterMinMax                                         FormatFeatures = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT
	FormatFeatureSampledImageFilterCubic                                          FormatFeatures = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_IMG
	FormatFeatureAccelerationStructureVertexBufferKHR                             FormatFeatures = C.VK_FORMAT_FEATURE_ACCELERATION_STRUCTURE_VERTEX_BUFFER_BIT_KHR
	FormatFeatureFragmentDensityMapEXT                                            FormatFeatures = C.VK_FORMAT_FEATURE_FRAGMENT_DENSITY_MAP_BIT_EXT
	FormatFeatureFragmentShadingRateAttachmentKHR                                 FormatFeatures = C.VK_FORMAT_FEATURE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR
)

var formatFeatureToString = map[FormatFeatures]string{
	FormatFeatureSampledImage:                                                     "Sampled Image",
	FormatFeatureStorageImage:                                                     "Storage Image",
	FormatFeatureStorageImageAtomic:                                               "Storage Image, Atomic",
	FormatFeatureUniformTexelBuffer:                                               "Uniform Texel Buffer",
	FormatFeatureStorageTexelBuffer:                                               "Storage Texel Buffer",
	FormatFeatureStorageTexelBufferAtomic:                                         "Storage Texel Buffer, Atomic",
	FormatFeatureVertexBuffer:                                                     "Vertex Buffer",
	FormatFeatureColorAttachment:                                                  "Color Attachment",
	FormatFeatureColorAttachmentBlend:                                             "Color Attachment Blend",
	FormatFeatureDepthStencilAttachment:                                           "Depth Stencil Attachment",
	FormatFeatureBlitSource:                                                       "Blit Source",
	FormatFeatureBlitDestination:                                                  "Blit Destination",
	FormatFeatureSampledImageFilterLinear:                                         "Sampled Image, Linear Filter",
	FormatFeatureTransferSource:                                                   "Transfer Source",
	FormatFeatureTransferDestination:                                              "Transfer Destination",
	FormatFeatureMidpointChromaSamples:                                            "Midpoint Chroma Samples",
	FormatFeatureSampledImageYcbcrConversionLinearFilter:                          "Sampled Image, YCbCr Conversion, Linear Filter",
	FormatFeatureSampledImageYcbcrConversionSeparateReconstructionFilter:          "Sampled Image, YCbCr Conversion, Separate Reconstruction Filter",
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicit:          "Sampled Image, YCbCr Conversion, Chroma Reconstruction: Explicit",
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitForceable: "Sampled Image, YCbCr Conversion, Chroma Reconstruction: Explicit & Forceable",
	FormatFeatureDisjoint:                                                         "Disjoint",
	FormatFeatureCositedChromaSamples:                                             "Cosited Chroma Samples",
	FormatFeatureSampledImageFilterMinMax:                                         "Sampled Image, Min/Max Filter",
	FormatFeatureSampledImageFilterCubic:                                          "Sampled Image, Cubic Filter",
	FormatFeatureAccelerationStructureVertexBufferKHR:                             "Acceleration Structure: Vertex Buffer (Khronos Extension)",
	FormatFeatureFragmentDensityMapEXT:                                            "Fragment Density Map (Extension)",
	FormatFeatureFragmentShadingRateAttachmentKHR:                                 "Fragment Shading Rate Attachment (Khronos Extension)",
}

func (f FormatFeatures) String() string {
	return common.FlagsToString(f, formatFeatureToString)
}

type FormatProperties struct {
	LinearTilingFeatures  FormatFeatures
	OptimalTilingFeatures FormatFeatures
	BufferFeatures        FormatFeatures
}
