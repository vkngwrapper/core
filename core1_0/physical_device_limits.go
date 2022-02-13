package core1_0

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

type PhysicalDeviceLimits struct {
	MaxImageDimension1D    int
	MaxImageDimension2D    int
	MaxImageDimension3D    int
	MaxImageDimensionCube  int
	MaxImageArrayLayers    int
	MaxTexelBufferElements int

	MaxUniformBufferRange int
	MaxStorageBufferRange int
	MaxPushConstantsSize  int

	MaxMemoryAllocationCount  int
	MaxSamplerAllocationCount int

	BufferImageGranularity int
	SparseAddressSpaceSize int

	MaxBoundDescriptorSets                int
	MaxPerStageDescriptorSamplers         int
	MaxPerStageDescriptorUniformBuffers   int
	MaxPerStageDescriptorStorageBuffers   int
	MaxPerStageDescriptorSampledImages    int
	MaxPerStageDescriptorStorageImages    int
	MaxPerStageDescriptorInputAttachments int
	MaxPerStageResources                  int

	MaxDescriptorSetSamplers              int
	MaxDescriptorSetUniformBuffers        int
	MaxDescriptorSetUniformBuffersDynamic int
	MaxDescriptorSetStorageBuffers        int
	MaxDescriptorSetStorageBuffersDynamic int
	MaxDescriptorSetSampledImages         int
	MaxDescriptorSetStorageImages         int
	MaxDescriptorSetInputAttachments      int

	MaxVertexInputAttributes      int
	MaxVertexInputBindings        int
	MaxVertexInputAttributeOffset int
	MaxVertexInputBindingStride   int
	MaxVertexOutputComponents     int

	MaxTessellationGenerationLevel                  int
	MaxTessellationPatchSize                        int
	MaxTessellationControlPerVertexInputComponents  int
	MaxTessellationControlPerVertexOutputComponents int
	MaxTessellationControlPerPatchOutputComponents  int
	MaxTessellationControlTotalOutputComponents     int
	MaxTessellationEvaluationInputComponents        int
	MaxTessellationEvaluationOutputComponents       int

	MaxGeometryShaderInvocations     int
	MaxGeometryInputComponents       int
	MaxGeometryOutputComponents      int
	MaxGeometryOutputVertices        int
	MaxGeometryTotalOutputComponents int

	MaxFragmentInputComponents         int
	MaxFragmentOutputAttachments       int
	MaxFragmentDualSrcAttachments      int
	MaxFragmentCombinedOutputResources int

	MaxComputeSharedMemorySize     int
	MaxComputeWorkGroupCount       [3]int
	MaxComputeWorkGroupInvocations int
	MaxComputeWorkGroupSize        [3]int

	SubPixelPrecisionBits int
	SubTexelPrecisionBits int
	MipmapPrecisionBits   int

	MaxDrawIndexedIndexValue int
	MaxDrawIndirectCount     int

	MaxSamplerLodBias    float32
	MaxSamplerAnisotropy float32

	MaxViewports          int
	MaxViewportDimensions [2]int
	ViewportBoundsRange   [2]float32
	ViewportSubPixelBits  int

	MinMemoryMapAlignment           int
	MinTexelBufferOffsetAlignment   int
	MinUniformBufferOffsetAlignment int
	MinStorageBufferOffsetAlignment int

	MinTexelOffset                  int
	MaxTexelOffset                  int
	MinTexelGatherOffset            int
	MaxTexelGatherOffset            int
	MinInterpolationOffset          float32
	MaxInterpolationOffset          float32
	SubPixelInterpolationOffsetBits int

	MaxFramebufferWidth  int
	MaxFramebufferHeight int
	MaxFramebufferLayers int

	FramebufferColorSampleCounts         common.SampleCounts
	FramebufferDepthSampleCounts         common.SampleCounts
	FramebufferStencilSampleCounts       common.SampleCounts
	FramebufferNoAttachmentsSampleCounts common.SampleCounts

	MaxColorAttachments             int
	SampledImageColorSampleCounts   common.SampleCounts
	SampledImageIntegerSampleCounts common.SampleCounts
	SampledImageDepthSampleCounts   common.SampleCounts
	SampledImageStencilSampleCounts common.SampleCounts
	StorageImageSampleCounts        common.SampleCounts
	MaxSampleMaskWords              int

	TimestampComputeAndGraphics bool
	TimestampPeriod             float32

	MaxClipDistances                int
	MaxCullDistances                int
	MaxCombinedClipAndCullDistances int

	DiscreteQueuePriorities int

	PointSizeRange       [2]float32
	LineWidthRange       [2]float32
	PointSizeGranularity float32
	LineWidthGranularity float32

	StrictLines             bool
	StandardSampleLocations bool

	OptimalBufferCopyOffsetAlignment   int
	OptimalBufferCopyRowPitchAlignment int
	NonCoherentAtomSize                int
}
