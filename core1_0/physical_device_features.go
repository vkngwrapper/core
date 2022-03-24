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

type PhysicalDeviceFeatures struct {
	RobustBufferAccess                      bool
	FullDrawIndexUint32                     bool
	ImageCubeArray                          bool
	IndependentBlend                        bool
	GeometryShader                          bool
	TessellationShader                      bool
	SampleRateShading                       bool
	DualSrcBlend                            bool
	LogicOp                                 bool
	MultiDrawIndirect                       bool
	DrawIndirectFirstInstance               bool
	DepthClamp                              bool
	DepthBiasClamp                          bool
	FillModeNonSolid                        bool
	DepthBounds                             bool
	WideLines                               bool
	LargePoints                             bool
	AlphaToOne                              bool
	MultiViewport                           bool
	SamplerAnisotropy                       bool
	TextureCompressionEtc2                  bool
	TextureCompressionAstcLdc               bool
	TextureCompressionBc                    bool
	OcclusionQueryPrecise                   bool
	PipelineStatisticsQuery                 bool
	VertexPipelineStoresAndAtomics          bool
	FragmentStoresAndAtomics                bool
	ShaderTessellationAndGeometryPointSize  bool
	ShaderImageGatherExtended               bool
	ShaderStorageImageExtendedFormats       bool
	ShaderStorageImageMultisample           bool
	ShaderStorageImageReadWithoutFormat     bool
	ShaderStorageImageWriteWithoutFormat    bool
	ShaderUniformBufferArrayDynamicIndexing bool
	ShaderSampledImageArrayDynamicIndexing  bool
	ShaderStorageBufferArrayDynamicIndexing bool
	ShaderStorageImageArrayDynamicIndexing  bool
	ShaderClipDistance                      bool
	ShaderCullDistance                      bool
	ShaderFloat64                           bool
	ShaderInt64                             bool
	ShaderInt16                             bool
	ShaderResourceResidency                 bool
	ShaderResourceMinLod                    bool
	SparseBinding                           bool
	SparseResidencyBuffer                   bool
	SparseResidencyImage2D                  bool
	SparseResidencyImage3D                  bool
	SparseResidency2Samples                 bool
	SparseResidency4Samples                 bool
	SparseResidency8Samples                 bool
	SparseResidency16Samples                bool
	SparseResidencyAliased                  bool
	VariableMultisampleRate                 bool
	InheritedQueries                        bool
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
