package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type PhysicalDeviceLimits struct {
	MaxImageDimension1D    uint32
	MaxImageDimension2D    uint32
	MaxImageDimension3D    uint32
	MaxImageDimensionCube  uint32
	MaxImageArrayLayers    uint32
	MaxTexelBufferElements uint32

	MaxUniformBufferRange uint32
	MaxStorageBufferRange uint32
	MaxPushConstantsSize  uint32

	MaxMemoryAllocationCount  uint32
	MaxSamplerAllocationCount uint32

	BufferImageGranularity uint64
	SparseAddressSpaceSize uint64

	MaxBoundDescriptorSets                uint32
	MaxPerStageDescriptorSamplers         uint32
	MaxPerStageDescriptorUniformBuffers   uint32
	MaxPerStageDescriptorStorageBuffers   uint32
	MaxPerStageDescriptorSampledImages    uint32
	MaxPerStageDescriptorStorageImages    uint32
	MaxPerStageDescriptorInputAttachments uint32
	MaxPerStageResources                  uint32

	MaxDescriptorSetSamplers              uint32
	MaxDescriptorSetUniformBuffers        uint32
	MaxDescriptorSetUniformBuffersDynamic uint32
	MaxDescriptorSetStorageBuffers        uint32
	MaxDescriptorSetStorageBuffersDynamic uint32
	MaxDescriptorSetSampledImages         uint32
	MaxDescriptorSetStorageImages         uint32
	MaxDescriptorSetInputAttachments      uint32

	MaxVertexInputAttributes      uint32
	MaxVertexInputBindings        uint32
	MaxVertexInputAttributeOffset uint32
	MaxVertexInputBindingStride   uint32
	MaxVertexOutputComponents     uint32

	MaxTessellationGenerationLevel                  uint32
	MaxTessellationPatchSize                        uint32
	MaxTessellationControlPerVertexInputComponents  uint32
	MaxTessellationControlPerVertexOutputComponents uint32
	MaxTessellationControlPerPatchOutputComponents  uint32
	MaxTessellationControlTotalOutputComponents     uint32
	MaxTessellationEvaluationInputComponents        uint32
	MaxTessellationEvaluationOutputComponents       uint32

	MaxGeometryShaderInvocations     uint32
	MaxGeometryInputComponents       uint32
	MaxGeometryOutputComponents      uint32
	MaxGeometryOutputVertices        uint32
	MaxGeometryTotalOutputComponents uint32

	MaxFragmentInputComponents         uint32
	MaxFragmentOutputAttachments       uint32
	MaxFragmentDualSrcAttachments      uint32
	MaxFragmentCombinedOutputResources uint32

	MaxComputeSharedMemorySize     uint32
	MaxComputeWorkGroupCount       [3]uint32
	MaxComputeWorkGroupInvocations uint32
	MaxComputeWorkGroupSize        [3]uint32

	SubPixelPrecisionBits uint32
	SubTexelPrecisionBits uint32
	MipmapPrecisionBits   uint32

	MaxDrawIndexedIndexValue uint32
	MaxDrawIndirectCount     uint32

	MaxSamplerLodBias    float32
	MaxSamplerAnisotropy float32

	MaxViewports          uint32
	MaxViewportDimensions [2]uint32
	ViewportBoundsRange   [2]float32
	ViewportSubPixelBits  uint32

	MinMemoryMapAlignment           uint
	MinTexelBufferOffsetAlignment   uint64
	MinUniformBufferOffsetAlignment uint64
	MinStorageBufferOffsetAlignment uint64

	MinTexelOffset                  int32
	MaxTexelOffset                  uint32
	MinTexelGatherOffset            int32
	MaxTexelGatherOffset            uint32
	MinInterpolationOffset          float32
	MaxInterpolationOffset          float32
	SubPixelInterpolationOffsetBits uint32

	MaxFramebufferWidth  uint32
	MaxFramebufferHeight uint32
	MaxFramebufferLayers uint32

	FramebufferColorSampleCounts         SampleCounts
	FramebufferDepthSampleCounts         SampleCounts
	FramebufferStencilSampleCounts       SampleCounts
	FramebufferNoAttachmentsSampleCounts SampleCounts

	MaxColorAttachments             uint32
	SampledImageColorSampleCounts   SampleCounts
	SampledImageIntegerSampleCounts SampleCounts
	SampledImageDepthSampleCounts   SampleCounts
	SampledImageStencilSampleCounts SampleCounts
	StorageImageSampleCounts        SampleCounts
	MaxSampleMaskWords              uint32

	TimestampComputeAndGraphics bool
	TimestampPeriod             float32

	MaxClipDistances                uint32
	MaxCullDistances                uint32
	MaxCombinedClipAndCullDistances uint32

	DiscreteQueuePriorities uint32

	PointSizeRange       [2]float32
	LineWidthRange       [2]float32
	PointSizeGranularity float32
	LineWidthGranularity float32

	StrictLines             bool
	StandardSampleLocations bool

	OptimalBufferCopyOffsetAlignment   uint64
	OptimalBufferCopyRowPitchAlignment uint64
	NonCoherentAtomSize                uint64
}
