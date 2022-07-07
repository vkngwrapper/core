package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"unsafe"
)

const (
	QueueGraphics      QueueFlags = C.VK_QUEUE_GRAPHICS_BIT
	QueueCompute       QueueFlags = C.VK_QUEUE_COMPUTE_BIT
	QueueTransfer      QueueFlags = C.VK_QUEUE_TRANSFER_BIT
	QueueSparseBinding QueueFlags = C.VK_QUEUE_SPARSE_BINDING_BIT

	MemoryPropertyDeviceLocal     MemoryPropertyFlags = C.VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT
	MemoryPropertyHostVisible     MemoryPropertyFlags = C.VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT
	MemoryPropertyHostCoherent    MemoryPropertyFlags = C.VK_MEMORY_PROPERTY_HOST_COHERENT_BIT
	MemoryPropertyLazilyAllocated MemoryPropertyFlags = C.VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT

	MemoryHeapDeviceLocal MemoryHeapFlags = C.VK_MEMORY_HEAP_DEVICE_LOCAL_BIT

	PhysicalDeviceTypeOther         PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_OTHER
	PhysicalDeviceTypeIntegratedGPU PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU
	PhysicalDeviceTypeDiscreteGPU   PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU
	PhysicalDeviceTypeVirtualGPU    PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU
	PhysicalDeviceTypeCPU           PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_CPU
)

func init() {
	QueueGraphics.Register("Graphics")
	QueueCompute.Register("Compute")
	QueueTransfer.Register("Transfer")
	QueueSparseBinding.Register("Sparse Binding")

	MemoryPropertyDeviceLocal.Register("Device Local")
	MemoryPropertyHostVisible.Register("Host Visible")
	MemoryPropertyHostCoherent.Register("Host Coherent")
	MemoryPropertyLazilyAllocated.Register("Lazily Allocated")

	MemoryHeapDeviceLocal.Register("Device Local")

	PhysicalDeviceTypeOther.Register("Other")
	PhysicalDeviceTypeIntegratedGPU.Register("Integrated GPU")
	PhysicalDeviceTypeDiscreteGPU.Register("Discrete GPU")
	PhysicalDeviceTypeVirtualGPU.Register("Virtual GPU")
	PhysicalDeviceTypeCPU.Register("CPU")
}

type PhysicalDeviceSparseProperties struct {
	ResidencyStandard2DBlockShape            bool
	ResidencyStandard2DMultisampleBlockShape bool
	ResidencyStandard3DBlockShape            bool
	ResidencyAlignedMipSize                  bool
	ResidencyNonResidentStrict               bool
}

type PhysicalDeviceProperties struct {
	DriverType PhysicalDeviceType
	DriverName string

	APIVersion    common.APIVersion
	DriverVersion common.Version
	VendorID      uint32
	DeviceID      uint32

	PipelineCacheUUID uuid.UUID
	Limits            *PhysicalDeviceLimits
	SparseProperties  *PhysicalDeviceSparseProperties
}

func (p *PhysicalDeviceProperties) PopulateFromCPointer(cPointer unsafe.Pointer) error {
	pData := (*C.VkPhysicalDeviceProperties)(cPointer)

	uuidBytes := C.GoBytes(unsafe.Pointer(&pData.pipelineCacheUUID[0]), C.VK_UUID_SIZE)
	uuid, err := uuid.FromBytes(uuidBytes)
	if err != nil {
		return errors.Wrap(err, "vulkan provided invalid pipeline cache uuid")
	}

	p.DriverType = PhysicalDeviceType(pData.deviceType)
	p.DriverName = C.GoString((*C.char)(&pData.deviceName[0]))
	p.APIVersion = common.APIVersion(pData.apiVersion)
	p.DriverVersion = common.Version(pData.driverVersion)
	p.VendorID = uint32(pData.vendorID)
	p.DeviceID = uint32(pData.deviceID)
	p.PipelineCacheUUID = uuid
	p.Limits = createPhysicalDeviceLimits(&pData.limits)
	p.SparseProperties = createSparseProperties(&pData.sparseProperties)

	return nil
}

type QueueFamily struct {
	QueueFlags                  QueueFlags
	QueueCount                  int
	TimestampValidBits          uint32
	MinImageTransferGranularity Extent3D
}

type PhysicalDeviceMemoryProperties struct {
	MemoryTypes []MemoryType
	MemoryHeaps []MemoryHeap
}

func createPhysicalDeviceLimits(l *C.VkPhysicalDeviceLimits) *PhysicalDeviceLimits {
	return &PhysicalDeviceLimits{
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
		FramebufferColorSampleCounts:                    SampleCountFlags(l.framebufferColorSampleCounts),
		FramebufferDepthSampleCounts:                    SampleCountFlags(l.framebufferDepthSampleCounts),
		FramebufferStencilSampleCounts:                  SampleCountFlags(l.framebufferStencilSampleCounts),
		FramebufferNoAttachmentsSampleCounts:            SampleCountFlags(l.framebufferNoAttachmentsSampleCounts),
		MaxColorAttachments:                             int(l.maxColorAttachments),
		SampledImageColorSampleCounts:                   SampleCountFlags(l.sampledImageColorSampleCounts),
		SampledImageIntegerSampleCounts:                 SampleCountFlags(l.sampledImageIntegerSampleCounts),
		SampledImageDepthSampleCounts:                   SampleCountFlags(l.sampledImageDepthSampleCounts),
		SampledImageStencilSampleCounts:                 SampleCountFlags(l.sampledImageStencilSampleCounts),
		StorageImageSampleCounts:                        SampleCountFlags(l.storageImageSampleCounts),
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

func createSparseProperties(p *C.VkPhysicalDeviceSparseProperties) *PhysicalDeviceSparseProperties {
	return &PhysicalDeviceSparseProperties{
		ResidencyStandard2DBlockShape:            p.residencyStandard2DBlockShape != C.VK_FALSE,
		ResidencyStandard2DMultisampleBlockShape: p.residencyStandard2DMultisampleBlockShape != C.VK_FALSE,
		ResidencyStandard3DBlockShape:            p.residencyStandard3DBlockShape != C.VK_FALSE,
		ResidencyNonResidentStrict:               p.residencyNonResidentStrict != C.VK_FALSE,
		ResidencyAlignedMipSize:                  p.residencyAlignedMipSize != C.VK_FALSE,
	}
}
