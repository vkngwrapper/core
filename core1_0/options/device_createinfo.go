package options

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type QueueFamilyOptions struct {
	QueueFamilyIndex int
	QueuePriorities  []float32
}

type DeviceOptions struct {
	QueueFamilies   []*QueueFamilyOptions
	EnabledFeatures *PhysicalDeviceFeatures
	ExtensionNames  []string
	LayerNames      []string

	core.HaveNext
}

func (o *DeviceOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if len(o.QueueFamilies) == 0 {
		return nil, errors.New("alloc DeviceOptions: no queue families added")
	}
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceCreateInfo{})))
	}

	// Alloc queue families
	queueFamilyPtr := allocator.Malloc(len(o.QueueFamilies) * int(unsafe.Sizeof([1]C.VkDeviceQueueCreateInfo{})))
	queueFamilyArray := ([]C.VkDeviceQueueCreateInfo)(unsafe.Slice((*C.VkDeviceQueueCreateInfo)(queueFamilyPtr), len(o.QueueFamilies)))

	for idx, queueFamily := range o.QueueFamilies {
		if len(queueFamily.QueuePriorities) == 0 {
			return nil, errors.Newf("alloc DeviceOptions: queue family %d had no queue priorities", queueFamily.QueueFamilyIndex)
		}

		prioritiesPtr := allocator.Malloc(len(queueFamily.QueuePriorities) * int(unsafe.Sizeof(C.float(0))))
		prioritiesArray := ([]C.float)(unsafe.Slice((*C.float)(prioritiesPtr), len(queueFamily.QueuePriorities)))
		for idx, priority := range queueFamily.QueuePriorities {
			prioritiesArray[idx] = C.float(priority)
		}

		queueFamilyArray[idx].sType = C.VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
		queueFamilyArray[idx].flags = 0
		queueFamilyArray[idx].pNext = nil
		queueFamilyArray[idx].queueCount = C.uint32_t(len(queueFamily.QueuePriorities))
		queueFamilyArray[idx].queueFamilyIndex = C.uint32_t(queueFamily.QueueFamilyIndex)
		queueFamilyArray[idx].pQueuePriorities = (*C.float)(prioritiesPtr)
	}

	// Alloc array of extension names
	numExtensions := len(o.ExtensionNames)
	extNamePtr := allocator.Malloc(numExtensions * int(unsafe.Sizeof(uintptr(0))))
	extNames := ([]*C.char)(unsafe.Slice((**C.char)(extNamePtr), numExtensions))
	for i := 0; i < numExtensions; i++ {
		extNames[i] = (*C.char)(allocator.CString(o.ExtensionNames[i]))
	}

	// Alloc array of layer names
	numLayers := len(o.LayerNames)
	layerNamePtr := allocator.Malloc(numLayers * int(unsafe.Sizeof(uintptr(0))))
	layerNames := ([]*C.char)(unsafe.Slice((**C.char)(layerNamePtr), numLayers))
	for i := 0; i < numLayers; i++ {
		layerNames[i] = (*C.char)(allocator.CString(o.LayerNames[i]))
	}

	createInfo := (*C.VkDeviceCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next
	createInfo.queueCreateInfoCount = C.uint32_t(len(o.QueueFamilies))
	createInfo.pQueueCreateInfos = (*C.VkDeviceQueueCreateInfo)(queueFamilyPtr)
	createInfo.enabledLayerCount = C.uint(numLayers)
	createInfo.ppEnabledLayerNames = (**C.char)(layerNamePtr)
	createInfo.enabledExtensionCount = C.uint(numExtensions)
	createInfo.ppEnabledExtensionNames = (**C.char)(extNamePtr)

	// Init feature list
	if o.EnabledFeatures != nil {
		features := (*C.VkPhysicalDeviceFeatures)(allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceFeatures{}))))
		populateFeatures(features, o.EnabledFeatures)
		createInfo.pEnabledFeatures = features
	} else {
		createInfo.pEnabledFeatures = nil
	}

	return unsafe.Pointer(createInfo), nil
}

func populateFeatures(f *C.VkPhysicalDeviceFeatures, enabledFeatures *PhysicalDeviceFeatures) {
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

	if enabledFeatures.RobustBufferAccess {
		f.robustBufferAccess = C.VK_TRUE
	}
	if enabledFeatures.FullDrawIndexUint32 {
		f.fullDrawIndexUint32 = C.VK_TRUE
	}
	if enabledFeatures.ImageCubeArray {
		f.imageCubeArray = C.VK_TRUE
	}
	if enabledFeatures.IndependentBlend {
		f.independentBlend = C.VK_TRUE
	}
	if enabledFeatures.GeometryShader {
		f.geometryShader = C.VK_TRUE
	}
	if enabledFeatures.TessellationShader {
		f.tessellationShader = C.VK_TRUE
	}
	if enabledFeatures.SampleRateShading {
		f.sampleRateShading = C.VK_TRUE
	}
	if enabledFeatures.DualSrcBlend {
		f.dualSrcBlend = C.VK_TRUE
	}
	if enabledFeatures.LogicOp {
		f.logicOp = C.VK_TRUE
	}
	if enabledFeatures.MultiDrawIndirect {
		f.multiDrawIndirect = C.VK_TRUE
	}
	if enabledFeatures.DrawIndirectFirstInstance {
		f.drawIndirectFirstInstance = C.VK_TRUE
	}
	if enabledFeatures.DepthClamp {
		f.depthClamp = C.VK_TRUE
	}
	if enabledFeatures.DepthBiasClamp {
		f.depthBiasClamp = C.VK_TRUE
	}
	if enabledFeatures.FillModeNonSolid {
		f.fillModeNonSolid = C.VK_TRUE
	}
	if enabledFeatures.DepthBounds {
		f.depthBounds = C.VK_TRUE
	}
	if enabledFeatures.WideLines {
		f.wideLines = C.VK_TRUE
	}
	if enabledFeatures.LargePoints {
		f.largePoints = C.VK_TRUE
	}
	if enabledFeatures.AlphaToOne {
		f.alphaToOne = C.VK_TRUE
	}
	if enabledFeatures.MultiViewport {
		f.multiViewport = C.VK_TRUE
	}
	if enabledFeatures.SamplerAnisotropy {
		f.samplerAnisotropy = C.VK_TRUE
	}
	if enabledFeatures.TextureCompressionEtc2 {
		f.textureCompressionETC2 = C.VK_TRUE
	}
	if enabledFeatures.TextureCompressionAstcLdc {
		f.textureCompressionASTC_LDR = C.VK_TRUE
	}
	if enabledFeatures.TextureCompressionBc {
		f.textureCompressionBC = C.VK_TRUE
	}
	if enabledFeatures.OcclusionQueryPrecise {
		f.occlusionQueryPrecise = C.VK_TRUE
	}
	if enabledFeatures.PipelineStatisticsQuery {
		f.pipelineStatisticsQuery = C.VK_TRUE
	}
	if enabledFeatures.VertexPipelineStoresAndAtomics {
		f.vertexPipelineStoresAndAtomics = C.VK_TRUE
	}
	if enabledFeatures.FragmentStoresAndAtomics {
		f.fragmentStoresAndAtomics = C.VK_TRUE
	}
	if enabledFeatures.ShaderTessellationAndGeometryPointSize {
		f.shaderTessellationAndGeometryPointSize = C.VK_TRUE
	}
	if enabledFeatures.ShaderImageGatherExtended {
		f.shaderImageGatherExtended = C.VK_TRUE
	}
	if enabledFeatures.ShaderStorageImageExtendedFormats {
		f.shaderStorageImageExtendedFormats = C.VK_TRUE
	}
	if enabledFeatures.ShaderStorageImageMultisample {
		f.shaderStorageImageMultisample = C.VK_TRUE
	}
	if enabledFeatures.ShaderStorageImageReadWithoutFormat {
		f.shaderStorageImageReadWithoutFormat = C.VK_TRUE
	}
	if enabledFeatures.ShaderStorageImageWriteWithoutFormat {
		f.shaderStorageImageWriteWithoutFormat = C.VK_TRUE
	}
	if enabledFeatures.ShaderUniformBufferArrayDynamicIndexing {
		f.shaderUniformBufferArrayDynamicIndexing = C.VK_TRUE
	}
	if enabledFeatures.ShaderSampledImageArrayDynamicIndexing {
		f.shaderSampledImageArrayDynamicIndexing = C.VK_TRUE
	}
	if enabledFeatures.ShaderStorageBufferArrayDynamicIndexing {
		f.shaderStorageBufferArrayDynamicIndexing = C.VK_TRUE
	}
	if enabledFeatures.ShaderStorageImageArrayDynamicIndexing {
		f.shaderStorageImageArrayDynamicIndexing = C.VK_TRUE
	}
	if enabledFeatures.ShaderClipDistance {
		f.shaderClipDistance = C.VK_TRUE
	}
	if enabledFeatures.ShaderCullDistance {
		f.shaderCullDistance = C.VK_TRUE
	}
	if enabledFeatures.ShaderFloat64 {
		f.shaderFloat64 = C.VK_TRUE
	}
	if enabledFeatures.ShaderInt64 {
		f.shaderInt64 = C.VK_TRUE
	}
	if enabledFeatures.ShaderInt16 {
		f.shaderInt16 = C.VK_TRUE
	}
	if enabledFeatures.ShaderResourceResidency {
		f.shaderResourceResidency = C.VK_TRUE
	}
	if enabledFeatures.ShaderResourceMinLod {
		f.shaderResourceMinLod = C.VK_TRUE
	}
	if enabledFeatures.SparseBinding {
		f.sparseBinding = C.VK_TRUE
	}
	if enabledFeatures.SparseResidencyBuffer {
		f.sparseResidencyBuffer = C.VK_TRUE
	}
	if enabledFeatures.SparseResidencyImage2D {
		f.sparseResidencyImage2D = C.VK_TRUE
	}
	if enabledFeatures.SparseResidencyImage3D {
		f.sparseResidencyImage3D = C.VK_TRUE
	}
	if enabledFeatures.SparseResidency2Samples {
		f.sparseResidency2Samples = C.VK_TRUE
	}
	if enabledFeatures.SparseResidency4Samples {
		f.sparseResidency4Samples = C.VK_TRUE
	}
	if enabledFeatures.SparseResidency8Samples {
		f.sparseResidency8Samples = C.VK_TRUE
	}
	if enabledFeatures.SparseResidency16Samples {
		f.sparseResidency16Samples = C.VK_TRUE
	}
	if enabledFeatures.SparseResidencyAliased {
		f.sparseResidencyAliased = C.VK_TRUE
	}
	if enabledFeatures.VariableMultisampleRate {
		f.variableMultisampleRate = C.VK_TRUE
	}
	if enabledFeatures.InheritedQueries {
		f.inheritedQueries = C.VK_TRUE
	}
}
