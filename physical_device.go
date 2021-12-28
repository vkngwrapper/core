package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"unsafe"
)

type vulkanPhysicalDevice struct {
	driver Driver
	handle VkPhysicalDevice
}

func (d *vulkanPhysicalDevice) Handle() VkPhysicalDevice {
	return d.handle
}

func (d *vulkanPhysicalDevice) Driver() Driver {
	return d.driver
}

func (d *vulkanPhysicalDevice) QueueFamilyProperties() []*common.QueueFamily {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	count := (*Uint32)(allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	d.driver.VkGetPhysicalDeviceQueueFamilyProperties(d.handle, count, nil)

	if *count == 0 {
		return nil
	}

	goCount := int(*count)

	allocatedHandles := allocator.Malloc(goCount * int(unsafe.Sizeof(C.VkQueueFamilyProperties{})))
	familyProperties := ([]C.VkQueueFamilyProperties)(unsafe.Slice((*C.VkQueueFamilyProperties)(allocatedHandles), int(*count)))

	d.driver.VkGetPhysicalDeviceQueueFamilyProperties(d.handle, count, (*VkQueueFamilyProperties)(allocatedHandles))

	var queueFamilies []*common.QueueFamily
	for i := 0; i < goCount; i++ {
		queueFamilies = append(queueFamilies, &common.QueueFamily{
			Flags:              common.QueueFlags(familyProperties[i].queueFlags),
			QueueCount:         uint32(familyProperties[i].queueCount),
			TimestampValidBits: uint32(familyProperties[i].timestampValidBits),
			MinImageTransferGranularity: common.Extent3D{
				Width:  int(familyProperties[i].minImageTransferGranularity.width),
				Height: int(familyProperties[i].minImageTransferGranularity.height),
				Depth:  int(familyProperties[i].minImageTransferGranularity.depth),
			},
		})
	}

	return queueFamilies
}

func (d *vulkanPhysicalDevice) Properties() *common.PhysicalDeviceProperties {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	propertiesUnsafe := allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceProperties{})))

	d.driver.VkGetPhysicalDeviceProperties(d.handle, (*VkPhysicalDeviceProperties)(propertiesUnsafe))

	return createPhysicalDeviceProperties((*C.VkPhysicalDeviceProperties)(propertiesUnsafe))
}

func (d *vulkanPhysicalDevice) Features() *common.PhysicalDeviceFeatures {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	featuresUnsafe := allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceFeatures{})))

	d.driver.VkGetPhysicalDeviceFeatures(d.handle, (*VkPhysicalDeviceFeatures)(featuresUnsafe))

	return createPhysicalDeviceFeatures((*C.VkPhysicalDeviceFeatures)(featuresUnsafe))
}

func (d *vulkanPhysicalDevice) attemptAvailableExtensions(layerNamePtr *Char) (map[string]*common.ExtensionProperties, VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	extensionCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	extensionCount := (*Uint32)(extensionCountPtr)

	res, err := d.driver.VkEnumerateDeviceExtensionProperties(d.handle, nil, extensionCount, nil)

	if err != nil || *extensionCount == 0 {
		return nil, res, err
	}

	extensionTotal := int(*extensionCount)
	extensionsPtr := allocator.Malloc(extensionTotal * int(unsafe.Sizeof([1]C.VkExtensionProperties{})))

	res, err = d.driver.VkEnumerateDeviceExtensionProperties(d.handle, nil, extensionCount, (*VkExtensionProperties)(extensionsPtr))
	if err != nil {
		return nil, res, err
	}

	retVal := make(map[string]*common.ExtensionProperties)
	extensionSlice := ([]C.VkExtensionProperties)(unsafe.Slice((*C.VkExtensionProperties)(extensionsPtr), extensionTotal))

	for i := 0; i < extensionTotal; i++ {
		extension := extensionSlice[i]

		outExtension := &common.ExtensionProperties{
			ExtensionName: C.GoString((*C.char)(&extension.extensionName[0])),
			SpecVersion:   common.Version(extension.specVersion),
		}

		existingExtension, ok := retVal[outExtension.ExtensionName]
		if ok && existingExtension.SpecVersion >= outExtension.SpecVersion {
			continue
		}
		retVal[outExtension.ExtensionName] = outExtension
	}

	return retVal, res, nil
}

func (d *vulkanPhysicalDevice) AvailableExtensions() (map[string]*common.ExtensionProperties, VkResult, error) {
	// There may be a race condition that adds new available extensions between getting the
	// extension count & pulling the extensions, in which case, attemptAvailableExtensions will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*common.ExtensionProperties
	var result VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == VKIncomplete) {
		layers, result, err = d.attemptAvailableExtensions(nil)
	}
	return layers, result, err
}

func (d *vulkanPhysicalDevice) AvailableExtensionsForLayer(layerName string) (map[string]*common.ExtensionProperties, VkResult, error) {
	// There may be a race condition that adds new available extensions between getting the
	// extension count & pulling the extensions, in which case, attemptAvailableExtensions will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*common.ExtensionProperties
	var result VkResult
	var err error
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	layerNamePtr := (*Char)(allocator.CString(layerName))
	for doWhile := true; doWhile; doWhile = (result == VKIncomplete) {
		layers, result, err = d.attemptAvailableExtensions(layerNamePtr)
	}
	return layers, result, err
}

func (d *vulkanPhysicalDevice) FormatProperties(format common.DataFormat) *common.FormatProperties {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	properties := (*C.VkFormatProperties)(allocator.Malloc(C.sizeof_struct_VkFormatProperties))

	d.driver.VkGetPhysicalDeviceFormatProperties(d.handle, VkFormat(format), (*VkFormatProperties)(properties))

	return &common.FormatProperties{
		LinearTilingFeatures:  common.FormatFeatures(properties.linearTilingFeatures),
		OptimalTilingFeatures: common.FormatFeatures(properties.optimalTilingFeatures),
		BufferFeatures:        common.FormatFeatures(properties.bufferFeatures),
	}
}

func createPhysicalDeviceFeatures(f *C.VkPhysicalDeviceFeatures) *common.PhysicalDeviceFeatures {
	return &common.PhysicalDeviceFeatures{
		RobustBufferAccess:                      f.robustBufferAccess != C.VK_FALSE,
		FullDrawIndexUint32:                     f.fullDrawIndexUint32 != C.VK_FALSE,
		ImageCubeArray:                          f.imageCubeArray != C.VK_FALSE,
		IndependentBlend:                        f.independentBlend != C.VK_FALSE,
		GeometryShader:                          f.geometryShader != C.VK_FALSE,
		TessellationShader:                      f.tessellationShader != C.VK_FALSE,
		SampleRateShading:                       f.sampleRateShading != C.VK_FALSE,
		DualSrcBlend:                            f.dualSrcBlend != C.VK_FALSE,
		LogicOp:                                 f.logicOp != C.VK_FALSE,
		MultiDrawIndirect:                       f.multiDrawIndirect != C.VK_FALSE,
		DrawIndirectFirstInstance:               f.drawIndirectFirstInstance != C.VK_FALSE,
		DepthClamp:                              f.depthClamp != C.VK_FALSE,
		DepthBiasClamp:                          f.depthBiasClamp != C.VK_FALSE,
		FillModeNonSolid:                        f.fillModeNonSolid != C.VK_FALSE,
		DepthBounds:                             f.depthBounds != C.VK_FALSE,
		WideLines:                               f.wideLines != C.VK_FALSE,
		LargePoints:                             f.largePoints != C.VK_FALSE,
		AlphaToOne:                              f.alphaToOne != C.VK_FALSE,
		MultiViewport:                           f.multiViewport != C.VK_FALSE,
		SamplerAnisotropy:                       f.samplerAnisotropy != C.VK_FALSE,
		TextureCompressionEtc2:                  f.textureCompressionETC2 != C.VK_FALSE,
		TextureCompressionAstcLdc:               f.textureCompressionASTC_LDR != C.VK_FALSE,
		TextureCompressionBc:                    f.textureCompressionBC != C.VK_FALSE,
		OcclusionQueryPrecise:                   f.occlusionQueryPrecise != C.VK_FALSE,
		PipelineStatisticsQuery:                 f.pipelineStatisticsQuery != C.VK_FALSE,
		VertexPipelineStoresAndAtomics:          f.vertexPipelineStoresAndAtomics != C.VK_FALSE,
		FragmentStoresAndAtomics:                f.fragmentStoresAndAtomics != C.VK_FALSE,
		ShaderTessellationAndGeometryPointSize:  f.shaderTessellationAndGeometryPointSize != C.VK_FALSE,
		ShaderImageGatherExtended:               f.shaderImageGatherExtended != C.VK_FALSE,
		ShaderStorageImageExtendedFormats:       f.shaderStorageImageExtendedFormats != C.VK_FALSE,
		ShaderStorageImageMultisample:           f.shaderStorageImageMultisample != C.VK_FALSE,
		ShaderStorageImageReadWithoutFormat:     f.shaderStorageImageReadWithoutFormat != C.VK_FALSE,
		ShaderStorageImageWriteWithoutFormat:    f.shaderStorageImageWriteWithoutFormat != C.VK_FALSE,
		ShaderUniformBufferArrayDynamicIndexing: f.shaderUniformBufferArrayDynamicIndexing != C.VK_FALSE,
		ShaderSampledImageArrayDynamicIndexing:  f.shaderSampledImageArrayDynamicIndexing != C.VK_FALSE,
		ShaderStorageBufferArrayDynamicIndexing: f.shaderStorageBufferArrayDynamicIndexing != C.VK_FALSE,
		ShaderStorageImageArrayDynamicIndexing:  f.shaderStorageImageArrayDynamicIndexing != C.VK_FALSE,
		ShaderClipDistance:                      f.shaderClipDistance != C.VK_FALSE,
		ShaderCullDistance:                      f.shaderCullDistance != C.VK_FALSE,
		ShaderFloat64:                           f.shaderFloat64 != C.VK_FALSE,
		ShaderInt64:                             f.shaderInt64 != C.VK_FALSE,
		ShaderInt16:                             f.shaderInt16 != C.VK_FALSE,
		ShaderResourceResidency:                 f.shaderResourceResidency != C.VK_FALSE,
		ShaderResourceMinLod:                    f.shaderResourceMinLod != C.VK_FALSE,
		SparseBinding:                           f.sparseBinding != C.VK_FALSE,
		SparseResidencyBuffer:                   f.sparseResidencyBuffer != C.VK_FALSE,
		SparseResidencyImage2D:                  f.sparseResidencyImage2D != C.VK_FALSE,
		SparseResidencyImage3D:                  f.sparseResidencyImage3D != C.VK_FALSE,
		SparseResidency2Samples:                 f.sparseResidency2Samples != C.VK_FALSE,
		SparseResidency4Samples:                 f.sparseResidency4Samples != C.VK_FALSE,
		SparseResidency8Samples:                 f.sparseResidency8Samples != C.VK_FALSE,
		SparseResidency16Samples:                f.sparseResidency16Samples != C.VK_FALSE,
		SparseResidencyAliased:                  f.sparseResidencyAliased != C.VK_FALSE,
		VariableMultisampleRate:                 f.variableMultisampleRate != C.VK_FALSE,
		InheritedQueries:                        f.inheritedQueries != C.VK_FALSE,
	}
}

func createPhysicalDeviceLimits(l *C.VkPhysicalDeviceLimits) *common.PhysicalDeviceLimits {
	return &common.PhysicalDeviceLimits{
		MaxImageDimension1D:                             int(l.maxImageDimension1D),
		MaxImageDimension2D:                             int(l.maxImageDimension2D),
		MaxImageDimension3D:                             int(l.maxImageDimension3D),
		MaxImageDimensionCube:                           int(l.maxImageDimensionCube),
		MaxImageArrayLayers:                             int(l.maxImageArrayLayers),
		MaxTexelBufferElements:                          int(l.maxTexelBufferElements),
		MaxUniformBufferRange:                           int(l.maxUniformBufferRange),
		MaxStorageBufferRange:                           int(l.maxStorageBufferRange),
		MaxPushConstantsSize:                            int(l.maxPushConstantsSize),
		MaxMemoryAllocationCount:                        int(l.maxMemoryAllocationCount),
		MaxSamplerAllocationCount:                       int(l.maxSamplerAllocationCount),
		BufferImageGranularity:                          int(l.bufferImageGranularity),
		SparseAddressSpaceSize:                          int(l.sparseAddressSpaceSize),
		MaxBoundDescriptorSets:                          int(l.maxBoundDescriptorSets),
		MaxPerStageDescriptorSamplers:                   int(l.maxPerStageDescriptorSamplers),
		MaxPerStageDescriptorUniformBuffers:             int(l.maxPerStageDescriptorUniformBuffers),
		MaxPerStageDescriptorStorageBuffers:             int(l.maxPerStageDescriptorStorageBuffers),
		MaxPerStageDescriptorSampledImages:              int(l.maxPerStageDescriptorSampledImages),
		MaxPerStageDescriptorStorageImages:              int(l.maxPerStageDescriptorStorageImages),
		MaxPerStageDescriptorInputAttachments:           int(l.maxPerStageDescriptorInputAttachments),
		MaxPerStageResources:                            int(l.maxPerStageResources),
		MaxDescriptorSetSamplers:                        int(l.maxDescriptorSetSamplers),
		MaxDescriptorSetUniformBuffers:                  int(l.maxDescriptorSetUniformBuffers),
		MaxDescriptorSetUniformBuffersDynamic:           int(l.maxDescriptorSetUniformBuffersDynamic),
		MaxDescriptorSetStorageBuffers:                  int(l.maxDescriptorSetStorageBuffers),
		MaxDescriptorSetStorageBuffersDynamic:           int(l.maxDescriptorSetStorageBuffersDynamic),
		MaxDescriptorSetSampledImages:                   int(l.maxDescriptorSetSampledImages),
		MaxDescriptorSetStorageImages:                   int(l.maxDescriptorSetStorageImages),
		MaxDescriptorSetInputAttachments:                int(l.maxDescriptorSetInputAttachments),
		MaxVertexInputAttributes:                        int(l.maxVertexInputAttributes),
		MaxVertexInputBindings:                          int(l.maxVertexInputBindings),
		MaxVertexInputAttributeOffset:                   int(l.maxVertexInputAttributeOffset),
		MaxVertexInputBindingStride:                     int(l.maxVertexInputBindingStride),
		MaxVertexOutputComponents:                       int(l.maxVertexOutputComponents),
		MaxTessellationGenerationLevel:                  int(l.maxTessellationGenerationLevel),
		MaxTessellationPatchSize:                        int(l.maxTessellationPatchSize),
		MaxTessellationControlPerVertexInputComponents:  int(l.maxTessellationControlPerVertexInputComponents),
		MaxTessellationControlPerVertexOutputComponents: int(l.maxTessellationControlPerVertexOutputComponents),
		MaxTessellationControlPerPatchOutputComponents:  int(l.maxTessellationControlPerPatchOutputComponents),
		MaxTessellationControlTotalOutputComponents:     int(l.maxTessellationControlTotalOutputComponents),
		MaxTessellationEvaluationInputComponents:        int(l.maxTessellationEvaluationInputComponents),
		MaxTessellationEvaluationOutputComponents:       int(l.maxTessellationEvaluationOutputComponents),
		MaxGeometryShaderInvocations:                    int(l.maxGeometryShaderInvocations),
		MaxGeometryInputComponents:                      int(l.maxGeometryInputComponents),
		MaxGeometryOutputComponents:                     int(l.maxGeometryOutputComponents),
		MaxGeometryOutputVertices:                       int(l.maxGeometryOutputVertices),
		MaxGeometryTotalOutputComponents:                int(l.maxGeometryTotalOutputComponents),
		MaxFragmentInputComponents:                      int(l.maxFragmentInputComponents),
		MaxFragmentOutputAttachments:                    int(l.maxFragmentOutputAttachments),
		MaxFragmentDualSrcAttachments:                   int(l.maxFragmentDualSrcAttachments),
		MaxFragmentCombinedOutputResources:              int(l.maxFragmentCombinedOutputResources),
		MaxComputeSharedMemorySize:                      int(l.maxComputeSharedMemorySize),
		MaxComputeWorkGroupInvocations:                  int(l.maxComputeWorkGroupInvocations),
		SubPixelPrecisionBits:                           int(l.subPixelPrecisionBits),
		SubTexelPrecisionBits:                           int(l.subTexelPrecisionBits),
		MipmapPrecisionBits:                             int(l.mipmapPrecisionBits),
		MaxDrawIndexedIndexValue:                        int(l.maxDrawIndexedIndexValue),
		MaxDrawIndirectCount:                            int(l.maxDrawIndirectCount),
		MaxSamplerLodBias:                               float32(l.maxSamplerLodBias),
		MaxSamplerAnisotropy:                            float32(l.maxSamplerAnisotropy),
		MaxViewports:                                    int(l.maxViewports),
		ViewportSubPixelBits:                            int(l.viewportSubPixelBits),
		MinMemoryMapAlignment:                           int(l.minMemoryMapAlignment),
		MinTexelBufferOffsetAlignment:                   int(l.minTexelBufferOffsetAlignment),
		MinUniformBufferOffsetAlignment:                 int(l.minUniformBufferOffsetAlignment),
		MinStorageBufferOffsetAlignment:                 int(l.minStorageBufferOffsetAlignment),
		MinTexelOffset:                                  int(l.minTexelOffset),
		MaxTexelOffset:                                  int(l.maxTexelOffset),
		MinTexelGatherOffset:                            int(l.minTexelGatherOffset),
		MaxTexelGatherOffset:                            int(l.maxTexelGatherOffset),
		MinInterpolationOffset:                          float32(l.minInterpolationOffset),
		MaxInterpolationOffset:                          float32(l.maxInterpolationOffset),
		SubPixelInterpolationOffsetBits:                 int(l.subPixelInterpolationOffsetBits),
		MaxFramebufferWidth:                             int(l.maxFramebufferWidth),
		MaxFramebufferHeight:                            int(l.maxFramebufferHeight),
		MaxFramebufferLayers:                            int(l.maxFramebufferLayers),
		FramebufferColorSampleCounts:                    common.SampleCounts(l.framebufferColorSampleCounts),
		FramebufferDepthSampleCounts:                    common.SampleCounts(l.framebufferDepthSampleCounts),
		FramebufferStencilSampleCounts:                  common.SampleCounts(l.framebufferStencilSampleCounts),
		FramebufferNoAttachmentsSampleCounts:            common.SampleCounts(l.framebufferNoAttachmentsSampleCounts),
		MaxColorAttachments:                             int(l.maxColorAttachments),
		SampledImageColorSampleCounts:                   common.SampleCounts(l.sampledImageColorSampleCounts),
		SampledImageIntegerSampleCounts:                 common.SampleCounts(l.sampledImageIntegerSampleCounts),
		SampledImageDepthSampleCounts:                   common.SampleCounts(l.sampledImageDepthSampleCounts),
		SampledImageStencilSampleCounts:                 common.SampleCounts(l.sampledImageStencilSampleCounts),
		StorageImageSampleCounts:                        common.SampleCounts(l.storageImageSampleCounts),
		MaxSampleMaskWords:                              int(l.maxSampleMaskWords),
		TimestampComputeAndGraphics:                     l.timestampComputeAndGraphics != C.VK_FALSE,
		TimestampPeriod:                                 float32(l.timestampPeriod),
		MaxClipDistances:                                int(l.maxClipDistances),
		MaxCullDistances:                                int(l.maxCullDistances),
		MaxCombinedClipAndCullDistances:                 int(l.maxCombinedClipAndCullDistances),
		DiscreteQueuePriorities:                         int(l.discreteQueuePriorities),
		PointSizeGranularity:                            float32(l.pointSizeGranularity),
		LineWidthGranularity:                            float32(l.lineWidthGranularity),
		StrictLines:                                     l.strictLines != C.VK_FALSE,
		StandardSampleLocations:                         l.standardSampleLocations != C.VK_FALSE,
		OptimalBufferCopyOffsetAlignment:                int(l.optimalBufferCopyOffsetAlignment),
		OptimalBufferCopyRowPitchAlignment:              int(l.optimalBufferCopyRowPitchAlignment),
		NonCoherentAtomSize:                             int(l.nonCoherentAtomSize),
		MaxComputeWorkGroupCount: [3]int{
			int(l.maxComputeWorkGroupCount[0]),
			int(l.maxComputeWorkGroupCount[1]),
			int(l.maxComputeWorkGroupCount[2]),
		},
		MaxComputeWorkGroupSize: [3]int{
			int(l.maxComputeWorkGroupSize[0]),
			int(l.maxComputeWorkGroupSize[1]),
			int(l.maxComputeWorkGroupSize[2]),
		},
		MaxViewportDimensions: [2]int{
			int(l.maxViewportDimensions[0]),
			int(l.maxViewportDimensions[1]),
		},
		ViewportBoundsRange: [2]float32{
			float32(l.viewportBoundsRange[0]),
			float32(l.viewportBoundsRange[1]),
		},
		PointSizeRange: [2]float32{
			float32(l.pointSizeRange[0]),
			float32(l.pointSizeRange[1]),
		},
		LineWidthRange: [2]float32{
			float32(l.lineWidthRange[0]),
			float32(l.lineWidthRange[1]),
		},
	}
}

func createSparseProperties(p *C.VkPhysicalDeviceSparseProperties) *common.PhysicalDeviceSparseProperties {
	return &common.PhysicalDeviceSparseProperties{
		ResidencyStandard2DBlockShape:            p.residencyStandard2DBlockShape != C.VK_FALSE,
		ResidencyStandard2DMultisampleBlockShape: p.residencyStandard2DMultisampleBlockShape != C.VK_FALSE,
		ResidencyStandard3DBlockShape:            p.residencyStandard3DBlockShape != C.VK_FALSE,
		ResidencyNonResidentStrict:               p.residencyNonResidentStrict != C.VK_FALSE,
		ResidencyAlignedMipSize:                  p.residencyAlignedMipSize != C.VK_FALSE,
	}
}

func createPhysicalDeviceProperties(p *C.VkPhysicalDeviceProperties) *common.PhysicalDeviceProperties {
	uuidBytes := C.GoBytes(unsafe.Pointer(&p.pipelineCacheUUID[0]), C.VK_UUID_SIZE)
	uuid, err := uuid.FromBytes(uuidBytes)
	if err != nil {
		panic(errors.Wrap(err, "vulkan provided invalid pipeline cache uuid"))
	}

	return &common.PhysicalDeviceProperties{
		Type: common.PhysicalDeviceType(p.deviceType),
		Name: C.GoString((*C.char)(&p.deviceName[0])),

		APIVersion:    common.APIVersion(p.apiVersion),
		DriverVersion: common.Version(p.driverVersion),

		VendorID: uint32(p.vendorID),
		DeviceID: uint32(p.deviceID),

		PipelineCacheUUID: uuid,
		Limits:            createPhysicalDeviceLimits(&p.limits),
		SparseProperties:  createSparseProperties(&p.sparseProperties),
	}
}
