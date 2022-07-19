package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

// PhysicalDeviceFeatures describes the fine-grained features that can be supported by an
// implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures.html
type PhysicalDeviceFeatures struct {
	// RobustBufferAccess specifies that access to Buffer objects are bounds-checked against the
	// range of the Buffer descriptor
	RobustBufferAccess bool
	// FullDrawIndexUint32 specifies the full 32-bit range of indices is supported for indexed
	// draw calls when using an IndexType of IndexTypeUInt32
	FullDrawIndexUint32 bool
	// ImageCubeArray specifies whether ImageView objects with an ImageViewType of ImageViewTypeCubeArray
	// can be created
	ImageCubeArray bool
	// IndependentBlend specifies whether the PipelineColorBlendAttachmentState settings are controlled
	// independently per-attachment. If this feature is not enabled, the PipelineColorBlendAttachmentState
	// settings for all color attachments must be identical
	IndependentBlend bool
	// GeometryShader specifies whether geometry shaders are supported
	GeometryShader bool
	// TessellationShader specifies whether tessellation control and evaluation shaders are supported
	TessellationShader bool
	// SampleRateShading specifies whether sample shading and multisample interpolation are supported
	SampleRateShading bool
	// DualSrcBlend specifies whether blend operations which take two sources are supported
	DualSrcBlend bool
	// LogicOp specifies whether logic operations are supported
	LogicOp bool
	// MultiDrawIndirect specifies whether multiple draw indirect is supported. If this feature is not
	// enabled, the drawCount parameter to CommandBuffer.CmdDrawIndirect and CommandBuffer.CmdDrawIndexedIndirect
	// must be 0 or 1
	MultiDrawIndirect bool
	// DrawIndirectFirstInstance specifies whether indirect drawing calls support the firstInstance
	// parameter
	DrawIndirectFirstInstance bool
	// DepthClamp specifies whether depth clamping is supported
	DepthClamp bool
	// DepthBiasClamp specifies whether depth bias clamping is supported
	DepthBiasClamp bool
	// FillModeNonSolid specifies whether point and wireframe fill modes are supported
	FillModeNonSolid bool
	// DepthBounds specifies whether depth bounds tests are supported
	DepthBounds bool
	// WideLines specifies whether lines with width other than 1.0 are supported
	WideLines bool
	// LargePoints specifies whether points with size greater than 1.0 are supported
	LargePoints bool
	// AlphaToOne specifies whether the implementation is able to replace the alpha
	// value of the fragment shader color output in the multisample coverage fragment
	// operation. If this feature is not enabled, then the alphaToOneEnable member
	// of PipelineMultisampleStateCreateInfo must be set to false
	AlphaToOne bool
	// MultiViewport specifies whether more than one viewport is supported
	MultiViewport bool
	// SamplerAnisotropy specifies whether anisotropic filtering is supported
	SamplerAnisotropy bool
	// TextureCompressionEtc2 specifies whether all of the ETC2 and EAC compressed texture
	// formats are supported. If the feature is not enabled, PhysicalDevice.FormatProperties
	// and PhysicalDevice.ImageFormatProperties can be used to check for supported properties
	// of individual formats as normal
	TextureCompressionEtc2 bool
	// TextureCompressionAstcLdc specifies whether all of the ASTC LDR compressed texture
	// formats are supported. If the feature is not enabled, PhysicalDevice.FormatProperties
	// and PhysicalDevice.ImageFormatProperties can be used to check for supported properties
	// of individual formats as normal
	TextureCompressionAstcLdc bool
	// TextureCompressionBc specifies whether all of the BC compressed texture formats are supported
	// If the feature is not enabled, PhysicalDevice.FormatProperties and
	// PhysicalDevice.ImageFormatProperties can be used to check for supported properties of
	//individual formats as normal
	TextureCompressionBc bool
	// OcclusionQueryPrecise specifies whether occlusion queries returning actual sample counts
	// are supported
	OcclusionQueryPrecise bool
	// PipelineStatisticsQuery specifies whether the Pipeline statistics queries are supported
	PipelineStatisticsQuery bool
	// VertexPipelineStoresAndAtomics specifies whether storage Buffer objects and Image objects
	// support stores and atomic operations in the vertex, tessellation, and geometry shader stages
	VertexPipelineStoresAndAtomics bool
	// FragmentStoresAndAtomics specifies whether storage Buffer objects and Image objects support
	// stores and atomic operations in the fragment shader stages
	FragmentStoresAndAtomics bool
	// ShaderTessellationAndGeometryPointSize specifies whether the PointSize built-in decoration
	// is available in the tessellation control, tessellation evaluation, and geometry shader stages
	ShaderTessellationAndGeometryPointSize bool
	// ShaderImageGatherExtended specifies whether the extended set of Image gather instructions
	// are available in shader code
	ShaderImageGatherExtended bool
	// ShaderStorageImageExtendedFormats specifies whether all the "storage Image extended formats"
	// are supported
	ShaderStorageImageExtendedFormats bool
	// ShaderStorageImageMultisample specifies whether multisampled storage Image objects are supported
	ShaderStorageImageMultisample bool
	// ShaderStorageImageReadWithoutFormat specifies whether storage Image objects require a format
	// qualifier to be specified when reading
	ShaderStorageImageReadWithoutFormat bool
	// ShaderStorageImageWriteWithoutFormat specifies whether storage Image objects require a format
	// qualifier to be specified when writing
	ShaderStorageImageWriteWithoutFormat bool
	// ShaderUniformBufferArrayDynamicIndexing specifies whether arrays of uniform Buffrer objects can
	// be indexed by dynamically uniform integer expressions in shader code
	ShaderUniformBufferArrayDynamicIndexing bool
	// ShaderSampledImageArrayDynamicIndexing specifies whether arrays of Sampler objects or sampled Image
	// objects can be indexed by dynamically uniform expressions in shader code
	ShaderSampledImageArrayDynamicIndexing bool
	// ShaderStorageBufferArrayDynamicIndexing specifies whether arrays of storage Buffer objects
	// can be indexed by dynamically uniform integer expressions in shader code
	ShaderStorageBufferArrayDynamicIndexing bool
	// ShaderStorageImageArrayDynamicIndexing specifies arrays of storage Image objects can be
	// indexed by dynamically uniform integer expressions in shader code
	ShaderStorageImageArrayDynamicIndexing bool
	// ShaderClipDistance specifies whether clip distances are supported in shader code
	ShaderClipDistance bool
	// ShaderCullDistance specifies whether cull distances are supported in shader code
	ShaderCullDistance bool
	// ShaderFloat64 specifies whether 64-bit floats (doubles) are supported in shader code
	ShaderFloat64 bool
	// ShaderInt64 specifies whether 64-bit integer (signed and unsigned) are supported in shader
	// code
	ShaderInt64 bool
	// ShaderInt16 specifies whether 16-bit integers (signed and unsigned) are supported in shader
	// code
	ShaderInt16 bool
	// ShaderResourceResidency specifies whether Image operations that return resource residency
	// information are supported in shader code
	ShaderResourceResidency bool
	// ShaderResourceMinLod specifies whether Image operations specifying the minimum resource LOD
	// are supported in shader code
	ShaderResourceMinLod bool
	// SparseBinding specifies whether resource memory can be managed at opaque sparse block level
	// instead of at object level
	SparseBinding bool
	// SparseResidencyBuffer specifies whether the Device can access partially resident Buffer objects
	SparseResidencyBuffer bool
	// SparseResidencyImage2D specifies whether the Device can access partially resident 2D Image
	// objects with 1 sample per pixel
	SparseResidencyImage2D bool
	// SparseResidencyImage3D specifies whether the Device can access partially resident 3D Image
	// objects
	SparseResidencyImage3D bool
	// SparseResidency2Samples specifies whether the PhysicalDevice can access partially resident
	// 2D Image objects with 2 samples per pixel
	SparseResidency2Samples bool
	// SparseResidency4Samples specifies whether the PhysicalDevice can access partially resident
	// 2D Image objects with 4 samples per pixel
	SparseResidency4Samples bool
	// SparseResidency8Samples specifies whether the PhysicalDevice can access partially resident
	// 2D Image objects with 8 samples per pixel
	SparseResidency8Samples bool
	// SparseResidency16Samples specifies whether the PhysicalDevice can access partially resident
	// 2D Image objects with 16 samples per pixel
	SparseResidency16Samples bool
	// SparseResidencyAliased specifies whether the PhysicalDevice can correctly access data aliased
	// into multiple locations
	SparseResidencyAliased bool
	// VariableMultisampleRate specifies whether all Pipeline objects that will be bound to a
	// CommandBuffer during a subpass which uses no attachments must have the same value for
	// PipelineMultisampleStateCreateInfo.rasterizationSamples
	VariableMultisampleRate bool
	// InheritedQueries specifies whether a secondary CommandBuffer may be executed while a
	// query is active
	InheritedQueries bool
}

func (p *PhysicalDeviceFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceFeatures)
	}

	f := (*C.VkPhysicalDeviceFeatures)(preallocatedPointer)
	f.robustBufferAccess = C.VK_FALSE
	f.fullDrawIndexUint32 = C.VK_FALSE
	f.imageCubeArray = C.VK_FALSE
	f.independentBlend = C.VK_FALSE
	f.geometryShader = C.VK_FALSE
	f.tessellationShader = C.VK_FALSE
	f.sampleRateShading = C.VK_FALSE
	f.dualSrcBlend = C.VK_FALSE
	f.logicOp = C.VK_FALSE
	f.multiDrawIndirect = C.VK_FALSE
	f.drawIndirectFirstInstance = C.VK_FALSE
	f.depthClamp = C.VK_FALSE
	f.depthBiasClamp = C.VK_FALSE
	f.fillModeNonSolid = C.VK_FALSE
	f.depthBounds = C.VK_FALSE
	f.wideLines = C.VK_FALSE
	f.largePoints = C.VK_FALSE
	f.alphaToOne = C.VK_FALSE
	f.multiViewport = C.VK_FALSE
	f.samplerAnisotropy = C.VK_FALSE
	f.textureCompressionETC2 = C.VK_FALSE
	f.textureCompressionASTC_LDR = C.VK_FALSE
	f.textureCompressionBC = C.VK_FALSE
	f.occlusionQueryPrecise = C.VK_FALSE
	f.pipelineStatisticsQuery = C.VK_FALSE
	f.vertexPipelineStoresAndAtomics = C.VK_FALSE
	f.fragmentStoresAndAtomics = C.VK_FALSE
	f.shaderTessellationAndGeometryPointSize = C.VK_FALSE
	f.shaderImageGatherExtended = C.VK_FALSE
	f.shaderStorageImageExtendedFormats = C.VK_FALSE
	f.shaderStorageImageMultisample = C.VK_FALSE
	f.shaderStorageImageReadWithoutFormat = C.VK_FALSE
	f.shaderStorageImageWriteWithoutFormat = C.VK_FALSE
	f.shaderUniformBufferArrayDynamicIndexing = C.VK_FALSE
	f.shaderSampledImageArrayDynamicIndexing = C.VK_FALSE
	f.shaderStorageBufferArrayDynamicIndexing = C.VK_FALSE
	f.shaderStorageImageArrayDynamicIndexing = C.VK_FALSE
	f.shaderClipDistance = C.VK_FALSE
	f.shaderCullDistance = C.VK_FALSE
	f.shaderFloat64 = C.VK_FALSE
	f.shaderInt64 = C.VK_FALSE
	f.shaderInt16 = C.VK_FALSE
	f.shaderResourceResidency = C.VK_FALSE
	f.shaderResourceMinLod = C.VK_FALSE
	f.sparseBinding = C.VK_FALSE
	f.sparseResidencyBuffer = C.VK_FALSE
	f.sparseResidencyImage2D = C.VK_FALSE
	f.sparseResidencyImage3D = C.VK_FALSE
	f.sparseResidency2Samples = C.VK_FALSE
	f.sparseResidency4Samples = C.VK_FALSE
	f.sparseResidency8Samples = C.VK_FALSE
	f.sparseResidency16Samples = C.VK_FALSE
	f.sparseResidencyAliased = C.VK_FALSE
	f.variableMultisampleRate = C.VK_FALSE
	f.inheritedQueries = C.VK_FALSE

	if p.RobustBufferAccess {
		f.robustBufferAccess = C.VK_TRUE
	}
	if p.FullDrawIndexUint32 {
		f.fullDrawIndexUint32 = C.VK_TRUE
	}
	if p.ImageCubeArray {
		f.imageCubeArray = C.VK_TRUE
	}
	if p.IndependentBlend {
		f.independentBlend = C.VK_TRUE
	}
	if p.GeometryShader {
		f.geometryShader = C.VK_TRUE
	}
	if p.TessellationShader {
		f.tessellationShader = C.VK_TRUE
	}
	if p.SampleRateShading {
		f.sampleRateShading = C.VK_TRUE
	}
	if p.DualSrcBlend {
		f.dualSrcBlend = C.VK_TRUE
	}
	if p.LogicOp {
		f.logicOp = C.VK_TRUE
	}
	if p.MultiDrawIndirect {
		f.multiDrawIndirect = C.VK_TRUE
	}
	if p.DrawIndirectFirstInstance {
		f.drawIndirectFirstInstance = C.VK_TRUE
	}
	if p.DepthClamp {
		f.depthClamp = C.VK_TRUE
	}
	if p.DepthBiasClamp {
		f.depthBiasClamp = C.VK_TRUE
	}
	if p.FillModeNonSolid {
		f.fillModeNonSolid = C.VK_TRUE
	}
	if p.DepthBounds {
		f.depthBounds = C.VK_TRUE
	}
	if p.WideLines {
		f.wideLines = C.VK_TRUE
	}
	if p.LargePoints {
		f.largePoints = C.VK_TRUE
	}
	if p.AlphaToOne {
		f.alphaToOne = C.VK_TRUE
	}
	if p.MultiViewport {
		f.multiViewport = C.VK_TRUE
	}
	if p.SamplerAnisotropy {
		f.samplerAnisotropy = C.VK_TRUE
	}
	if p.TextureCompressionEtc2 {
		f.textureCompressionETC2 = C.VK_TRUE
	}
	if p.TextureCompressionAstcLdc {
		f.textureCompressionASTC_LDR = C.VK_TRUE
	}
	if p.TextureCompressionBc {
		f.textureCompressionBC = C.VK_TRUE
	}
	if p.OcclusionQueryPrecise {
		f.occlusionQueryPrecise = C.VK_TRUE
	}
	if p.PipelineStatisticsQuery {
		f.pipelineStatisticsQuery = C.VK_TRUE
	}
	if p.VertexPipelineStoresAndAtomics {
		f.vertexPipelineStoresAndAtomics = C.VK_TRUE
	}
	if p.FragmentStoresAndAtomics {
		f.fragmentStoresAndAtomics = C.VK_TRUE
	}
	if p.ShaderTessellationAndGeometryPointSize {
		f.shaderTessellationAndGeometryPointSize = C.VK_TRUE
	}
	if p.ShaderImageGatherExtended {
		f.shaderImageGatherExtended = C.VK_TRUE
	}
	if p.ShaderStorageImageExtendedFormats {
		f.shaderStorageImageExtendedFormats = C.VK_TRUE
	}
	if p.ShaderStorageImageMultisample {
		f.shaderStorageImageMultisample = C.VK_TRUE
	}
	if p.ShaderStorageImageReadWithoutFormat {
		f.shaderStorageImageReadWithoutFormat = C.VK_TRUE
	}
	if p.ShaderStorageImageWriteWithoutFormat {
		f.shaderStorageImageWriteWithoutFormat = C.VK_TRUE
	}
	if p.ShaderUniformBufferArrayDynamicIndexing {
		f.shaderUniformBufferArrayDynamicIndexing = C.VK_TRUE
	}
	if p.ShaderSampledImageArrayDynamicIndexing {
		f.shaderSampledImageArrayDynamicIndexing = C.VK_TRUE
	}
	if p.ShaderStorageBufferArrayDynamicIndexing {
		f.shaderStorageBufferArrayDynamicIndexing = C.VK_TRUE
	}
	if p.ShaderStorageImageArrayDynamicIndexing {
		f.shaderStorageImageArrayDynamicIndexing = C.VK_TRUE
	}
	if p.ShaderClipDistance {
		f.shaderClipDistance = C.VK_TRUE
	}
	if p.ShaderCullDistance {
		f.shaderCullDistance = C.VK_TRUE
	}
	if p.ShaderFloat64 {
		f.shaderFloat64 = C.VK_TRUE
	}
	if p.ShaderInt64 {
		f.shaderInt64 = C.VK_TRUE
	}
	if p.ShaderInt16 {
		f.shaderInt16 = C.VK_TRUE
	}
	if p.ShaderResourceResidency {
		f.shaderResourceResidency = C.VK_TRUE
	}
	if p.ShaderResourceMinLod {
		f.shaderResourceMinLod = C.VK_TRUE
	}
	if p.SparseBinding {
		f.sparseBinding = C.VK_TRUE
	}
	if p.SparseResidencyBuffer {
		f.sparseResidencyBuffer = C.VK_TRUE
	}
	if p.SparseResidencyImage2D {
		f.sparseResidencyImage2D = C.VK_TRUE
	}
	if p.SparseResidencyImage3D {
		f.sparseResidencyImage3D = C.VK_TRUE
	}
	if p.SparseResidency2Samples {
		f.sparseResidency2Samples = C.VK_TRUE
	}
	if p.SparseResidency4Samples {
		f.sparseResidency4Samples = C.VK_TRUE
	}
	if p.SparseResidency8Samples {
		f.sparseResidency8Samples = C.VK_TRUE
	}
	if p.SparseResidency16Samples {
		f.sparseResidency16Samples = C.VK_TRUE
	}
	if p.SparseResidencyAliased {
		f.sparseResidencyAliased = C.VK_TRUE
	}
	if p.VariableMultisampleRate {
		f.variableMultisampleRate = C.VK_TRUE
	}
	if p.InheritedQueries {
		f.inheritedQueries = C.VK_TRUE
	}

	return preallocatedPointer, nil
}

func (p *PhysicalDeviceFeatures) PopulateFromCPointer(cPointer unsafe.Pointer) {
	f := (*C.VkPhysicalDeviceFeatures)(cPointer)

	p.RobustBufferAccess = f.robustBufferAccess != C.VK_FALSE
	p.FullDrawIndexUint32 = f.fullDrawIndexUint32 != C.VK_FALSE
	p.ImageCubeArray = f.imageCubeArray != C.VK_FALSE
	p.IndependentBlend = f.independentBlend != C.VK_FALSE
	p.GeometryShader = f.geometryShader != C.VK_FALSE
	p.TessellationShader = f.tessellationShader != C.VK_FALSE
	p.SampleRateShading = f.sampleRateShading != C.VK_FALSE
	p.DualSrcBlend = f.dualSrcBlend != C.VK_FALSE
	p.LogicOp = f.logicOp != C.VK_FALSE
	p.MultiDrawIndirect = f.multiDrawIndirect != C.VK_FALSE
	p.DrawIndirectFirstInstance = f.drawIndirectFirstInstance != C.VK_FALSE
	p.DepthClamp = f.depthClamp != C.VK_FALSE
	p.DepthBiasClamp = f.depthBiasClamp != C.VK_FALSE
	p.FillModeNonSolid = f.fillModeNonSolid != C.VK_FALSE
	p.DepthBounds = f.depthBounds != C.VK_FALSE
	p.WideLines = f.wideLines != C.VK_FALSE
	p.LargePoints = f.largePoints != C.VK_FALSE
	p.AlphaToOne = f.alphaToOne != C.VK_FALSE
	p.MultiViewport = f.multiViewport != C.VK_FALSE
	p.SamplerAnisotropy = f.samplerAnisotropy != C.VK_FALSE
	p.TextureCompressionEtc2 = f.textureCompressionETC2 != C.VK_FALSE
	p.TextureCompressionAstcLdc = f.textureCompressionASTC_LDR != C.VK_FALSE
	p.TextureCompressionBc = f.textureCompressionBC != C.VK_FALSE
	p.OcclusionQueryPrecise = f.occlusionQueryPrecise != C.VK_FALSE
	p.PipelineStatisticsQuery = f.pipelineStatisticsQuery != C.VK_FALSE
	p.VertexPipelineStoresAndAtomics = f.vertexPipelineStoresAndAtomics != C.VK_FALSE
	p.FragmentStoresAndAtomics = f.fragmentStoresAndAtomics != C.VK_FALSE
	p.ShaderTessellationAndGeometryPointSize = f.shaderTessellationAndGeometryPointSize != C.VK_FALSE
	p.ShaderImageGatherExtended = f.shaderImageGatherExtended != C.VK_FALSE
	p.ShaderStorageImageExtendedFormats = f.shaderStorageImageExtendedFormats != C.VK_FALSE
	p.ShaderStorageImageMultisample = f.shaderStorageImageMultisample != C.VK_FALSE
	p.ShaderStorageImageReadWithoutFormat = f.shaderStorageImageReadWithoutFormat != C.VK_FALSE
	p.ShaderStorageImageWriteWithoutFormat = f.shaderStorageImageWriteWithoutFormat != C.VK_FALSE
	p.ShaderUniformBufferArrayDynamicIndexing = f.shaderUniformBufferArrayDynamicIndexing != C.VK_FALSE
	p.ShaderSampledImageArrayDynamicIndexing = f.shaderSampledImageArrayDynamicIndexing != C.VK_FALSE
	p.ShaderStorageBufferArrayDynamicIndexing = f.shaderStorageBufferArrayDynamicIndexing != C.VK_FALSE
	p.ShaderStorageImageArrayDynamicIndexing = f.shaderStorageImageArrayDynamicIndexing != C.VK_FALSE
	p.ShaderClipDistance = f.shaderClipDistance != C.VK_FALSE
	p.ShaderCullDistance = f.shaderCullDistance != C.VK_FALSE
	p.ShaderFloat64 = f.shaderFloat64 != C.VK_FALSE
	p.ShaderInt64 = f.shaderInt64 != C.VK_FALSE
	p.ShaderInt16 = f.shaderInt16 != C.VK_FALSE
	p.ShaderResourceResidency = f.shaderResourceResidency != C.VK_FALSE
	p.ShaderResourceMinLod = f.shaderResourceMinLod != C.VK_FALSE
	p.SparseBinding = f.sparseBinding != C.VK_FALSE
	p.SparseResidencyBuffer = f.sparseResidencyBuffer != C.VK_FALSE
	p.SparseResidencyImage2D = f.sparseResidencyImage2D != C.VK_FALSE
	p.SparseResidencyImage3D = f.sparseResidencyImage3D != C.VK_FALSE
	p.SparseResidency2Samples = f.sparseResidency2Samples != C.VK_FALSE
	p.SparseResidency4Samples = f.sparseResidency4Samples != C.VK_FALSE
	p.SparseResidency8Samples = f.sparseResidency8Samples != C.VK_FALSE
	p.SparseResidency16Samples = f.sparseResidency16Samples != C.VK_FALSE
	p.SparseResidencyAliased = f.sparseResidencyAliased != C.VK_FALSE
	p.VariableMultisampleRate = f.variableMultisampleRate != C.VK_FALSE
	p.InheritedQueries = f.inheritedQueries != C.VK_FALSE
}
