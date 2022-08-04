package core1_1

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

// ChromaLocation is the position of downsampled chroma samples
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocation.html
type ChromaLocation int32

var chromaLocationMapping = make(map[ChromaLocation]string)

func (e ChromaLocation) Register(str string) {
	chromaLocationMapping[e] = str
}

func (e ChromaLocation) String() string {
	return chromaLocationMapping[e]
}

////

// SamplerYcbcrModelConversion is the color model component of a color space
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrModelConversion.html
type SamplerYcbcrModelConversion int32

var samplerModelConversionMapping = make(map[SamplerYcbcrModelConversion]string)

func (e SamplerYcbcrModelConversion) Register(str string) {
	samplerModelConversionMapping[e] = str
}

func (e SamplerYcbcrModelConversion) String() string {
	return samplerModelConversionMapping[e]
}

////

// SamplerYcbcrRange is a range of encoded values in a color space
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrRange.html
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
	// ChromaLocationCositedEven specifies that downsampled chroma samples are aligned
	// with luma samples with even coordinates
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocation.html
	ChromaLocationCositedEven ChromaLocation = C.VK_CHROMA_LOCATION_COSITED_EVEN
	// ChromaLocationMidpoint specifies that downsampled chroma samples are located halfway
	// between each even luma sample and the nearest higher odd luma sample
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocation.html
	ChromaLocationMidpoint ChromaLocation = C.VK_CHROMA_LOCATION_MIDPOINT

	// FormatFeatureCositedChromaSamples specifies that an application can define a
	// SamplerYcbcrConversion using this format as a source, and that an Image of this format
	// can be used with a SamplerYcbcrConversionCreateInfo.XChromaOffset and/or a
	// SamplerYcbcrConversionCreateInfo.YChromaOffset of ChromaLocationCositedEven
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureCositedChromaSamples core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT
	// FormatFeatureDisjoint specifies that a multi-planar Image can have ImageCreateDisjoint
	// set during Image creation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureDisjoint core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_DISJOINT_BIT
	// FormatFeatureMidpointChromaSamples specifies that an application can define a
	// SamplerYcbcrConversion using this format as a source, and that an Image of this format
	// can be used with a SamplerYcbcrConversionCreateInfo.XChromaOffset and/or a
	// SamplerYcbcrConversionCreateInfo.YChromaOffset of ChromaLocationMidpoint
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureMidpointChromaSamples core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT
	// FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicit specifies that
	// reconstruction is explicit
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicit core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT
	// FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitForceable specifies
	// that reconstruction can be forcibly made explicit by setting
	// SamplerYcbcrConversionCreateInfo.ForceExplicitReconstruction to true
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitForceable core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT
	// FormatFeatureSampledImageYcbcrConversionLinearFilter specifies that an application can
	// define a SamplerYcbcrConversion using this format as a source with ChromaFilter set
	// to core1_0.FilterLinear
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureSampledImageYcbcrConversionLinearFilter core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT
	// FormatFeatureSampledImageYcbcrConversionSeparateReconstructionFilter specifies that
	// the format can have different chroma, min, and mag filters
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureSampledImageYcbcrConversionSeparateReconstructionFilter core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT

	// SamplerYcbcrModelConversionRGBIdentity specifies that the input values to the conversion
	// are unmodified
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrModelConversion.html
	SamplerYcbcrModelConversionRGBIdentity SamplerYcbcrModelConversion = C.VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY
	// SamplerYcbcrModelConversionYcbcr2020 specifies the color model conversion from
	// Y'CbCr to R'G'B' defined in BT.2020
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrModelConversion.html
	SamplerYcbcrModelConversionYcbcr2020 SamplerYcbcrModelConversion = C.VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020
	// SamplerYcbcrModelConversionYcbcr601 specifies the color model conversion from Y'CbCr to
	// R'G'B' defined in BT.601
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrModelConversion.html
	SamplerYcbcrModelConversionYcbcr601 SamplerYcbcrModelConversion = C.VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601
	// SamplerYcbcrModelConversionYcbcr709 specifies the color model conversion from Y'CbCr to
	// R'G'B' defined in BT.709
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrModelConversion.html
	SamplerYcbcrModelConversionYcbcr709 SamplerYcbcrModelConversion = C.VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709
	// SamplerYcbcrModelConversionYcbcrIdentity specifies no model conversion but the inputs
	// are range expanded as for Y'CbCr
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrModelConversion.html
	SamplerYcbcrModelConversionYcbcrIdentity SamplerYcbcrModelConversion = C.VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY

	// SamplerYcbcrRangeITUFull specifies that the full range of the encoded values are valid and
	// interpreted according to the ITU "full range" quantization rules
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrRange.html
	SamplerYcbcrRangeITUFull SamplerYcbcrRange = C.VK_SAMPLER_YCBCR_RANGE_ITU_FULL
	// SamplerYcbcrRangeITUNarrow specifies that headroom and foot room are reserved in the
	// numerical range of encoded values, and the remaining values are expanded according to
	// the ITU "narrow range" quantization rules
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrRange.html
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

// SamplerYcbcrConversionCreateInfo specifies the parameters of the newly-created conversion
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html
type SamplerYcbcrConversionCreateInfo struct {
	// Format is the format of the Image from which color information will be retrieved
	Format core1_0.Format
	// YcbcrModel describes the color matrix for conversion between color models
	YcbcrModel SamplerYcbcrModelConversion
	// YcbcrRange describes whether the encoded values have headroom and foot room, or whether
	// the encoding uses the full numerical range
	YcbcrRange SamplerYcbcrRange
	// Components applies a swizzle based on core1_0.ComponentSwizzle enums prior to range
	// expansion and color model conversion
	Components core1_0.ComponentMapping
	// XChromaOffset describes the sample location associated with downsampled chroma components
	// in the x dimension
	XChromaOffset ChromaLocation
	// YChromaOffset describes the sample location associated with downsampled chroma components
	// in the y dimension
	YChromaOffset ChromaLocation
	// ChromaFilter is the filter for chroma reconstruction
	ChromaFilter core1_0.Filter
	// ForceExplicitReconstruction can be used to ensure that reconstruction is done explicitly,
	// if supported
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

// SamplerYcbcrConversionImageFormatProperties specifies combined Image sampler descriptor
// count for multi-planar images
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionImageFormatProperties.html
type SamplerYcbcrConversionImageFormatProperties struct {
	// CombinedImageSamplerDescriptorCount is the number of combined Image sampler descriptors that
	// the implementation uses to access the format
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

// ImagePlaneMemoryRequirementsInfo specifies Image plane for memory requirements
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePlaneMemoryRequirementsInfo.html
type ImagePlaneMemoryRequirementsInfo struct {
	// PlaneAspect specifies the aspect corresponding to the Image plane
	// to query
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

// SamplerYcbcrConversionInfo specifies a Y'CbCr conversion to a Sampler or ImageView
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionInfo.html
type SamplerYcbcrConversionInfo struct {
	// Conversion is a SamplerYcbcrConversion object created from the Device
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
