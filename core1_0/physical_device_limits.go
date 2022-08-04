package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"

// PhysicalDeviceLimits reports implementation-dependent PhysicalDevice limits
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html
type PhysicalDeviceLimits struct {
	// MaxImageDimension1D is the largest dimension (width) that is guaranteed to be supported
	// for all Image objects created with an ImageType of ImageType1D
	MaxImageDimension1D int
	// MaxImageDimension2D is the largest dimension (width or height) that is guaranteed to be
	// supported for all Image objects created with an ImageType of ImageType2D and without
	// ImageCreateCubeCompatible set in flags
	MaxImageDimension2D int
	// MaxImageDimension3D is the largest dimension (width, height, or depth) that is guaranteed
	// to be supported for all Image objects created with an ImageType of ImageType3D
	MaxImageDimension3D int
	// MaxImageDimensionCube is the largest dimension (width or height) that is guaranteed to
	// be supported for all Image objects created with an ImageType of ImageType2D and with
	// ImageCreateCubeCompatible set in flags
	MaxImageDimensionCube int
	// MaxImageArrayLayers is the maximum number of layers for an Image
	MaxImageArrayLayers int
	// MaxTexelBufferElements is the maximum number of addressable texels for a BufferView
	// created on a Buffer which was created with the BufferUsageUniformTexelBuffer or
	// BufferUsageStorageTexelBuffer usages set in BufferCreateInfo
	MaxTexelBufferElements int

	// MaxUniformBufferRange is the maximum value that can be specified in the range member
	// of a DescriptorBufferInfo structure passed to Device.UpdateDescriptorSets for
	// descriptors of type DescriptorTypeUniformBuffer or DescriptorTypeUniformBufferDynamic
	MaxUniformBufferRange int
	// MaxStorageBufferRange is the maximum value that can be specified in the range member of
	// a DescriptorBufferInfo structure passed to Device.UpdateDescriptorSets for
	// descriptors of type DescriptorTypeStorageBuffer or DescriptorTypeStorageBufferDynamic
	MaxStorageBufferRange int
	// MaxPushConstantsSize is the maximum size, in bytes, of the pool of push constant memory
	MaxPushConstantsSize int

	// MaxMemoryAllocationCount is the maximum number of DeviceMemory allocations, as created
	// by Device.AllocateMemory, which can simultaneously exist
	MaxMemoryAllocationCount int
	// MaxSamplerAllocationCount is the maximum number of Sampler objects, as created by
	// Device.CreateSampler, which can simultaneously exist on a device
	MaxSamplerAllocationCount int

	// BufferImageGranularity si the granularity, in bytes, at which Buffer or linear
	// Image resources, and optimal Image resources can be bound to adjacent offsets in the
	// same DeviceMemory object without aliasing
	BufferImageGranularity int
	// SparseAddressSpaceSize is the total amount of address space available, in bytes,
	// for sparse memory resources
	SparseAddressSpaceSize int

	// MaxBoundDescriptorSets is the maximum number of DescriptorSet objects that can be
	// simultaneously used by a Pipeline
	MaxBoundDescriptorSets int
	// MaxPerStageDescriptorSamplers is the maximum number of Sample objects that can be
	// accessible to a single shader stage in a PipelineLayout
	MaxPerStageDescriptorSamplers int
	// MaxPerStageDescriptorUniformBuffers is the maximum number of uniform Buffer objects
	// that can be accessible to a single shader stage in a PipelineLayout
	MaxPerStageDescriptorUniformBuffers int
	// MaxPerStageDescriptorStorageBuffers is the maximum number of storage Buffer objects
	// that can be accessible to a single shader stage in a PipelineLayout
	MaxPerStageDescriptorStorageBuffers int
	// MaxPerStageDescriptorSampledImages is the maximum number of sampled Image objects that
	// can be accessible to a single shader stage in a PipelineLayout
	MaxPerStageDescriptorSampledImages int
	// MaxPerStageDescriptorStorageImages is the maximum number of storage Image objects that
	// can be accessible to a single shader stage in a PipelineLayout
	MaxPerStageDescriptorStorageImages int
	// MaxPerStageDescriptorInputAttachments is the maximum number of input attachments that
	// can be accessible to a single shader stage in a PipelineLayout
	MaxPerStageDescriptorInputAttachments int
	// MaxPerStageResources is the maximum number of resources that can be accessible to a single
	// shader stage in a PipelineLayout. Descriptors with a type of DescriptorTypeCombinedImageSampler,
	// DescriptorTypeSampledImage, DescriptorTypeStorageImage, DescriptorTypeUniformTexelBuffer
	// DescriptorTypeStorageTexelBuffer, DescriptorTypeUniformBuffer, DescriptorTypeStorageBuffer,
	// DescriptorTypeUniformBufferDynamic, DescriptorTypeStorageBufferDynamic, and
	// DescriptorTypeInputAttachment all count against this limit
	MaxPerStageResources int

	// MaxDescriptorSetSamplers is the maximum number of Sampler objects that can be included in
	// a PipelineLayout
	MaxDescriptorSetSamplers int
	// MaxDescriptorSetUniformBuffers ist he maximum number of uniform Buffer objects that can
	// be included in a PipelineLayout
	MaxDescriptorSetUniformBuffers int
	// MaxDescriptorSetUniformBuffersDynamic is the maximum number of dynamic uniform Buffer
	// objects that can be included in a PipelineLayout
	MaxDescriptorSetUniformBuffersDynamic int
	// MaxDescriptorSetStorageBuffers is the maximum number of storage Buffer objects that can be
	// included in a PipelineLayout
	MaxDescriptorSetStorageBuffers int
	// MaxDescriptorSetStorageBuffersDynamic is the maximum number of dynamic storage Buffer
	// objects that can be included in a PipelineLayout
	MaxDescriptorSetStorageBuffersDynamic int
	// MaxDescriptorSetSampledImages is the maximum number of sampled Image objects that can
	// be included in a PipelineLayout
	MaxDescriptorSetSampledImages int
	// MaxDescriptorSetStorageImages is the maximum number of storage Image objects that can
	// be included in a PipelineLayout
	MaxDescriptorSetStorageImages int
	// MaxDescriptorSetInputAttachments is the maximum number of input attachments that can be
	// included in a PipelineLayout
	MaxDescriptorSetInputAttachments int

	// MaxVertexInputAttributes is the maximum number of vertex input attributes that can
	// be specified for a graphics Pipeline
	MaxVertexInputAttributes int
	// MaxVertexInputBindings is the maximum number of vertex Buffer objects that can be specified
	// for providing vertex attributes to a graphics Pipeline
	MaxVertexInputBindings int
	// MaxVertexInputAttributeOffset is the maximum vertex input attribute offset that can be added
	// to the vertex input binding stride
	MaxVertexInputAttributeOffset int
	// MaxVertexInputBindingStride is the maximum vertex input binding stride that can be specified
	// in a vertex input binding
	MaxVertexInputBindingStride int
	// MaxVertexOutputComponents is the maximum number of components of output variables which
	// can be output by a vertex shader
	MaxVertexOutputComponents int

	// MaxTessellationGenerationLevel is the maximum tessellation generation level supported
	// by the fixed-function tessellation primitive generator
	MaxTessellationGenerationLevel int
	// MaxTessellationPatchSize is the maximum patch size, in vertices, of patches that can
	// be processed by the tessellation control shader and tessellation primitive generator
	MaxTessellationPatchSize int
	// MaxTessellationControlPerVertexInputComponents is the maximum number of components
	// of input variables which can be provided as per-vertex inputs to the tessellation
	// control shader
	MaxTessellationControlPerVertexInputComponents int
	// MaxTessellationControlPerVertexOutputComponents is the maximum number of components of
	// per-vertex output variables which can be output from the tessellation control shader
	// stage
	MaxTessellationControlPerVertexOutputComponents int
	// MaxTessellationControlPerPatchOutputComponents is the maximum number of comonents of
	// per-patch output variables which can be output from the tessellation control shader
	// stage
	MaxTessellationControlPerPatchOutputComponents int
	// MaxTessellationControlTotalOutputComponents is the maximum total number of components
	// of per-vertex and per-patch output variables which can be output from the tessellation
	// control shader stage
	MaxTessellationControlTotalOutputComponents int
	// MaxTessellationEvaluationInputComponents is the maximum number of components of input
	// variables which can be provided as per-vertex inputs to the tessellation evaluation shader
	// stage
	MaxTessellationEvaluationInputComponents int
	// MaxTessellationEvaluationOutputComponents is the maximum number of components of per-vertex
	// output variables which can be output from the tessellation evaluation shader stage
	MaxTessellationEvaluationOutputComponents int

	// MaxGeometryShaderInvocations is the maximum invocation count supported for instanced
	// geometry shaders
	MaxGeometryShaderInvocations int
	// MaxGeometryInputComponents is the maximum number of components of input variables which
	// can be provided as inputs to the geometry shader stage
	MaxGeometryInputComponents int
	// MaxGeometryOutputComponents is the maximum number of components of output variables
	// which can be output from the geometry shader stage
	MaxGeometryOutputComponents int
	// MaxGeometryOutputVertices is the maximum number of vertices which can be emitted by any
	// geometry shader
	MaxGeometryOutputVertices int
	// MaxGeometryTotalOutputComponents is the maximum total number of components of output variables,
	// across all emitted vertices, which can be output from the geometry shader stage
	MaxGeometryTotalOutputComponents int

	// MaxFragmentInputComponents is the maximum number of components of input variables which can
	// be provided as inputs to the fragment shader stage
	MaxFragmentInputComponents int
	// MaxFragmentOutputAttachments is the maximum number of output attachments which can be
	// written to by the fragment shader stage
	MaxFragmentOutputAttachments int
	// MaxFragmentDualSrcAttachments is the maximum number of output attachments which can be
	// written to by the fragment shader stage when blending is enabled and one of the dual source
	// blend modes is in use
	MaxFragmentDualSrcAttachments int
	// MaxFragmentCombinedOutputResources is the total number of storage Buffer objects, storage
	// Image objects, and output Location decorated color attachments which can be used in the
	// fragment shader stage
	MaxFragmentCombinedOutputResources int

	// MaxComputeSharedMemorySize is the maximum total storage size, in bytes, available for
	// variables declared with the Workgroup storage class in shader modules in the compute shader
	// stage
	MaxComputeSharedMemorySize int
	// MaxComputeWorkGroupCount is the maximum number of local workgroups that can be dispatched
	// by a single dispatching command
	MaxComputeWorkGroupCount [3]int
	// MaxComputeWorkGroupInvocations is the maximum total number of compute shader invocations in
	// a single local workgroup
	MaxComputeWorkGroupInvocations int
	// MaxComputeWorkGroupSize is the maximum size of a local compute workgroup, per dimension
	MaxComputeWorkGroupSize [3]int

	// SubPixelPrecisionBits is the number of bits of subpixel precision in Framebuffer coordinates
	// xf and yf
	SubPixelPrecisionBits int
	// SubTexelPrecisionBits is the number of bits of precision in the division along an axis of
	// an Image used for minification and magnification filters
	SubTexelPrecisionBits int
	// MipmapPrecisionBits is the number of bits of division that the LOD calculation for mipmap
	// fetching get snapped to when determining the contribution from each mip level to the mip
	// filtered results
	MipmapPrecisionBits int

	// MaxDrawIndexedIndexValue is the maximum index value that can be used for indexed draw
	// calls when using 32-bit indices
	MaxDrawIndexedIndexValue int
	// MaxDrawIndirectCount is the maximum draw count that is supported for indirect drawing calls
	MaxDrawIndirectCount int

	// MaxSamplerLodBias is the maximum absolute sampler LOD bias
	MaxSamplerLodBias float32
	// MaxSamplerAnisotropy is the maximum degree of sampler anisotropy
	MaxSamplerAnisotropy float32

	// MaxViewports is the maximum number of active viewports
	MaxViewports int
	// MaxViewportDimensions are the maximum viewport dimensions in the X (width) and Y (height)
	// dimensions, respectively
	MaxViewportDimensions [2]int
	// ViewportBoundsRange is the [minimum, maximum] range that the corners of a viewport must be
	// contained in
	ViewportBoundsRange [2]float32
	// ViewportSubPixelBits is the number of bits of subpixel precision for viewport bounds
	ViewportSubPixelBits int

	// MinMemoryMapAlignment is the minimum required alignment, in bytes, of host visible memory
	// allocations within the host address space
	MinMemoryMapAlignment int
	// MinTexelBufferOffsetAlignment is the minimum required alignment, in bytes, for the offset
	// member of the BufferViewCreateInfo structure for texel Buffer objects
	MinTexelBufferOffsetAlignment int
	// MinUniformBufferOffsetAlignment is the minimum required alignment, in bytes, for the offset
	// member of the DescriptorBufferInfo structure for uniform Buffer objects
	MinUniformBufferOffsetAlignment int
	// MinStorageBufferOffsetAlignment is the minimum required alignment, in bytes, for the offset
	// member of the DescriptorBufferInfo structure for storage Buffer objects
	MinStorageBufferOffsetAlignment int

	// MinTexelOffset is the minimum offset value for the ConstOffset Image operand and any of the
	// OpImageSample... or OpImageFetch... Image instructions
	MinTexelOffset int
	// MaxTexelOffset is the maximum offset value for the ConstOffset Image operand and any of the
	// OpImageSample... or OpImageFetch... Image instructions
	MaxTexelOffset int
	// MinTexelGatherOffset is the minimum offset value for the Offset, ConstOffset, or ConstOffsets
	// image operands of any of the OpImage...Gather Image instructions
	MinTexelGatherOffset int
	// MaxTexelGatherOffset is the maximum offset value for the Offset, ConstOffset, or ConstOffsets
	// image operands of any of the OpImage...Gather Image instructions
	MaxTexelGatherOffset int
	// MinInterpolationOffset is the base minimum (inclusive) negative offset value for the Offset
	// operand of the InterpolateAtOffset extended instruction.
	MinInterpolationOffset float32
	// MaxInterpolationOffset is the base maximum (inclusive) negative offset value for the Offset
	// operand of the InterpolateAtOffset extended instruction.
	MaxInterpolationOffset float32
	// SubPixelInterpolationOffsetBits is the number of fractional bits that the x and y offsets
	// to the InterpolateAtOffset extended instruction may be rounded to as fixed-point values.
	SubPixelInterpolationOffsetBits int

	// MaxFramebufferWidth is the maximum width for a Framebuffer
	MaxFramebufferWidth int
	// MaxFramebufferHeight is the maximum height for a Framebuffer
	MaxFramebufferHeight int
	// MaxFramebufferLayers is the maximum layer count for a layered Framebuffer
	MaxFramebufferLayers int

	// FramebufferColorSampleCounts indicates the color sample counts that are supported for all
	// Framebuffer color atttachments with floating- or fixed-point formats
	FramebufferColorSampleCounts SampleCountFlags
	// FramebufferDepthSampleCounts indicates the supported depth sample counts for all Framebuffer
	// depth/stencil attachments, when the format includes a depth component
	FramebufferDepthSampleCounts SampleCountFlags
	// FramebufferStencilSampleCounts indicates the supported stencil sample counts for all
	// Framebuffer depth/stencil attachments, when the format includes a stencil component
	FramebufferStencilSampleCounts SampleCountFlags
	// FramebufferNoAttachmentsSampleCount indicates the supported sample counts for a subpass which
	// uses no attachments
	FramebufferNoAttachmentsSampleCounts SampleCountFlags

	// MaxColorAttachments is the maximum number of color attachments that can be used by a
	// subpass in a RenderPass
	MaxColorAttachments int
	// SampledImageColorSampleCounts indicates the sample counts supported for all 2D Image objects
	// created with ImageTilingOptimal, usage containing ImageUsageSampled, and a non-integer color
	// format
	SampledImageColorSampleCounts SampleCountFlags
	// SampledImageIntegerSampleCounts indicates the sample counts supported for all 2D Image objects
	// created with ImageTilingOptimal, usage containing UsageSampled, and an integer color format
	SampledImageIntegerSampleCounts SampleCountFlags
	// SampledImageDepthSampleCounts indicates the sample counts supported for all 2D Image objects
	// created with ImageTilingOptimal, usage containing ImageUsageSampled, and a depth format
	SampledImageDepthSampleCounts SampleCountFlags
	// SampledImageStencilSampleCounts indicates the sample counts supported for all 2D Image objects
	// created with ImageTilingOptimal, usage containing ImageUsageSampled, and a stencil format
	SampledImageStencilSampleCounts SampleCountFlags
	// StorageImageSampleCounts indicates the Sample counts supported for all 2D images created
	// with ImageTilingOptimal, and usage containing ImageUsageStorage
	StorageImageSampleCounts SampleCountFlags
	// MaxSampleMaskWords is the maximum number of array elements in a variable decorated with
	// the SampleMask built-in decoration
	MaxSampleMaskWords int

	// TimestampComputeAndGraphics specifies support for timestamps on all graphics and compute
	// queues
	TimestampComputeAndGraphics bool
	// TimestampPeriod is the number of nanoseconds required for a timestamp query to be incremented
	// by 1
	TimestampPeriod float32

	// MaxClipDistances is the maximum number of clip distances that can be used in a single shader
	// stage
	MaxClipDistances int
	// MaxCullDistances is the maximum number of cull distances that can be used in a single shader
	// stage
	MaxCullDistances int
	// MaxCombinedClipAndCullDistances is the maximum combined number of clip and cull distances
	// that can be used in a single shader stage
	MaxCombinedClipAndCullDistances int

	// DiscreteQueuePriorities is the number of discrete priorities that can be assigned to a Queue
	// based on the value of each member of DeviceQueueCreateInfo.QueuePriorities
	DiscreteQueuePriorities int

	// PointSizeRange is the range [minimum,maximum] of supported sizes for points
	PointSizeRange [2]float32
	// LineWidthRange is the range [minimum,maximum] of supported widths for lines
	LineWidthRange [2]float32
	// PointSizeGranularity is the granularity of supported point sizes
	PointSizeGranularity float32
	// LineWidthGranularity is the granularity of supported line widths
	LineWidthGranularity float32

	// StrictLines specifies whether lines are rasterized according to the preferred method of
	// rasterization
	StrictLines bool
	// StandardSampleLocations specifies whether rasterization uses the standard sample locations
	StandardSampleLocations bool

	// OptimalBufferCopyOffsetAlignment is the optimal Buffer offset alignment in bytes for
	// CommandBuffer.CmdCopyBufferToImage and CommandBuffer.CmdCopyImageToBuffer
	OptimalBufferCopyOffsetAlignment int
	// OptimalBufferCopyRowPitchAlignment is the optimal Buffer row pitch alignment in bytes
	// for CommandBuffer.CmdCopyBufferToImage and CommandBuffer.CmdCopyImageToBuffer
	OptimalBufferCopyRowPitchAlignment int
	// NonCoherentAtomSize is the size and alignment in bytes that bounds concurrent access
	// to host-mapped device memory
	NonCoherentAtomSize int
}
