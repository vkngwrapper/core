package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type ChromaLocation int32

var chromaLocationMapping = make(map[ChromaLocation]string)

func (e ChromaLocation) Register(str string) {
	chromaLocationMapping[e] = str
}

func (e ChromaLocation) String() string {
	return chromaLocationMapping[e]
}

////

type SamplerYcbcrModelConversion int32

var samplerModelConversionMapping = make(map[SamplerYcbcrModelConversion]string)

func (e SamplerYcbcrModelConversion) Register(str string) {
	samplerModelConversionMapping[e] = str
}

func (e SamplerYcbcrModelConversion) String() string {
	return samplerModelConversionMapping[e]
}

////

type SamplerYcbcrRange int32

var samplerRangeMapping = make(map[SamplerYcbcrRange]string)

func (e SamplerYcbcrRange) Register(str string) {
	samplerRangeMapping[e] = str
}

func (e SamplerYcbcrRange) String() string {
	return samplerRangeMapping[e]
}

////

const (
	ChromaLocationCositedEven ChromaLocation = C.VK_CHROMA_LOCATION_COSITED_EVEN
	ChromaLocationMidpoint    ChromaLocation = C.VK_CHROMA_LOCATION_MIDPOINT

	FormatFeatureCositedChromaSamples                                             core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT
	FormatFeatureDisjoint                                                         core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_DISJOINT_BIT
	FormatFeatureMidpointChromaSamples                                            core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicit          core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitForceable core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT
	FormatFeatureSampledImageYcbcrConversionLinearFilter                          core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT
	FormatFeatureSampledImageYcbcrConversionSeparateReconstructionFilter          core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT

	SamplerYcbcrModelConversionRGBIdentity   SamplerYcbcrModelConversion = C.VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY
	SamplerYcbcrModelConversionYcbcr2020     SamplerYcbcrModelConversion = C.VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020
	SamplerYcbcrModelConversionYcbcr601      SamplerYcbcrModelConversion = C.VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601
	SamplerYcbcrModelConversionYcbcr709      SamplerYcbcrModelConversion = C.VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709
	SamplerYcbcrModelConversionYcbcrIdentity SamplerYcbcrModelConversion = C.VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY

	SamplerYcbcrRangeITUFull   SamplerYcbcrRange = C.VK_SAMPLER_YCBCR_RANGE_ITU_FULL
	SamplerYcbcrRangeITUNarrow SamplerYcbcrRange = C.VK_SAMPLER_YCBCR_RANGE_ITU_NARROW
)

func init() {
	ChromaLocationCositedEven.Register("Cosited Even")
	ChromaLocationMidpoint.Register("Midpoint")

	FormatFeatureCositedChromaSamples.Register("Cosited Chroma Samples")
	FormatFeatureDisjoint.Register("Disjoint")
	FormatFeatureMidpointChromaSamples.Register("Midpoint Chroma Samples")
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicit.Register("Sampled Image Ycbcr Conversion - Chroma Reconstruction (Explicit)")
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitForceable.Register("Sampled Image Ycbcr Conversion - Chroma Reconstruction (Explicit, Forceable)")
	FormatFeatureSampledImageYcbcrConversionLinearFilter.Register("Sampled Image Ycbcr Conversion - Linear Filter")
	FormatFeatureSampledImageYcbcrConversionSeparateReconstructionFilter.Register("Sampled Image Ycbcr Conversion - Separate Reconstruction Filter")

	SamplerYcbcrModelConversionRGBIdentity.Register("RGB Identity")
	SamplerYcbcrModelConversionYcbcr2020.Register("Ycbcr 2020")
	SamplerYcbcrModelConversionYcbcr601.Register("Ycbcr 601")
	SamplerYcbcrModelConversionYcbcr709.Register("Ycbcr 709")
	SamplerYcbcrModelConversionYcbcrIdentity.Register("Ycbcr Identity")

	SamplerYcbcrRangeITUFull.Register("ITU Full")
	SamplerYcbcrRangeITUNarrow.Register("ITU Narrow")
}

type SamplerYcbcrConversionCreateInfo struct {
	Format                      core1_0.Format
	YcbcrModel                  SamplerYcbcrModelConversion
	YcbcrRange                  SamplerYcbcrRange
	Components                  core1_0.ComponentMapping
	XChromaOffset               ChromaLocation
	YChromaOffset               ChromaLocation
	ChromaFilter                core1_0.Filter
	ForceExplicitReconstruction bool

	common.NextOptions
}

func (o SamplerYcbcrConversionCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSamplerYcbcrConversionCreateInfo{})))
	}

	info := (*C.VkSamplerYcbcrConversionCreateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO
	info.pNext = next
	info.format = C.VkFormat(o.Format)
	info.ycbcrModel = C.VkSamplerYcbcrModelConversion(o.YcbcrModel)
	info.ycbcrRange = C.VkSamplerYcbcrRange(o.YcbcrRange)
	info.components.r = C.VkComponentSwizzle(o.Components.R)
	info.components.g = C.VkComponentSwizzle(o.Components.G)
	info.components.b = C.VkComponentSwizzle(o.Components.B)
	info.components.a = C.VkComponentSwizzle(o.Components.A)
	info.xChromaOffset = C.VkChromaLocation(o.XChromaOffset)
	info.yChromaOffset = C.VkChromaLocation(o.YChromaOffset)
	info.chromaFilter = C.VkFilter(o.ChromaFilter)
	info.forceExplicitReconstruction = C.VkBool32(0)

	if o.ForceExplicitReconstruction {
		info.forceExplicitReconstruction = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

////

type SamplerYcbcrConversionImageFormatProperties struct {
	CombinedImageSamplerDescriptorCount int

	common.NextOutData
}

func (o *SamplerYcbcrConversionImageFormatProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSamplerYcbcrConversionImageFormatProperties{})))
	}

	info := (*C.VkSamplerYcbcrConversionImageFormatProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *SamplerYcbcrConversionImageFormatProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkSamplerYcbcrConversionImageFormatProperties)(cDataPointer)

	o.CombinedImageSamplerDescriptorCount = int(info.combinedImageSamplerDescriptorCount)

	return info.pNext, nil
}

////

type ImagePlaneMemoryRequirementsInfo struct {
	PlaneAspect core1_0.ImageAspectFlags

	common.NextOptions
}

func (o ImagePlaneMemoryRequirementsInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImagePlaneMemoryRequirementsInfo{})))
	}

	info := (*C.VkImagePlaneMemoryRequirementsInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO
	info.pNext = next
	info.planeAspect = C.VkImageAspectFlagBits(o.PlaneAspect)

	return preallocatedPointer, nil
}

////

type SamplerYcbcrConversionInfo struct {
	Conversion SamplerYcbcrConversion

	common.NextOptions
}

func (o SamplerYcbcrConversionInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSamplerYcbcrConversionInfo{})))
	}

	info := (*C.VkSamplerYcbcrConversionInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO
	info.pNext = next
	info.conversion = C.VkSamplerYcbcrConversion(unsafe.Pointer(o.Conversion.Handle()))

	return preallocatedPointer, nil
}
