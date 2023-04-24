package core1_2

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
)

// DriverID specifies khronos driver id's
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
type DriverID int32

var driverIDMapping = make(map[DriverID]string)

func (e DriverID) Register(str string) {
	driverIDMapping[e] = str
}

func (e DriverID) String() string {
	return driverIDMapping[e]
}

// ResolveModeFlags indicates supported depth and stencil resolve modes
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html
type ResolveModeFlags int32

var resolveModeFlagsMapping = common.NewFlagStringMapping[ResolveModeFlags]()

func (f ResolveModeFlags) Register(str string) {
	resolveModeFlagsMapping.Register(f, str)
}

func (f ResolveModeFlags) String() string {
	return resolveModeFlagsMapping.FlagsToString(f)
}

////

// ShaderFloatControlsIndependence specifies whether, and how, shader float controls
// can be set separately
type ShaderFloatControlsIndependence int32

var shaderFloatControlsIndependenceMapping = make(map[ShaderFloatControlsIndependence]string)

func (e ShaderFloatControlsIndependence) Register(str string) {
	shaderFloatControlsIndependenceMapping[e] = str
}

func (e ShaderFloatControlsIndependence) String() string {
	return shaderFloatControlsIndependenceMapping[e]
}

////

const (
	// MaxDriverInfoSize is the length of a PhysicalDevice driver information string
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_MAX_DRIVER_INFO_SIZE.html
	MaxDriverInfoSize int = C.VK_MAX_DRIVER_INFO_SIZE
	// MaxDriverNameSize is the maximum length of a PhysicalDevice driver name string
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_MAX_DRIVER_NAME_SIZE.html
	MaxDriverNameSize int = C.VK_MAX_DRIVER_NAME_SIZE

	// DriverIDAmdOpenSource indicates open-source AMD drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDAmdOpenSource DriverID = C.VK_DRIVER_ID_AMD_OPEN_SOURCE
	// DriverIDAmdProprietary indicates proprietary AMD drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDAmdProprietary DriverID = C.VK_DRIVER_ID_AMD_PROPRIETARY
	// DriverIDArmProprietary indicates proprietary ARM drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDArmProprietary DriverID = C.VK_DRIVER_ID_ARM_PROPRIETARY
	// DriverIDBroadcomProprietary indicates proprietary Broadcom drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDBroadcomProprietary DriverID = C.VK_DRIVER_ID_BROADCOM_PROPRIETARY
	// DriverIDGgpProprietary indicates proprietary GGP drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDGgpProprietary DriverID = C.VK_DRIVER_ID_GGP_PROPRIETARY
	// DriverIDGoogleSwiftshader indicates Google Swiftshader drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDGoogleSwiftshader DriverID = C.VK_DRIVER_ID_GOOGLE_SWIFTSHADER
	// DriverIDImaginationProprietary indicates proprietary Imagination drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDImaginationProprietary DriverID = C.VK_DRIVER_ID_IMAGINATION_PROPRIETARY
	// DriverIDIntelOpenSourceMesa indicates open-source Mesa drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDIntelOpenSourceMesa DriverID = C.VK_DRIVER_ID_INTEL_OPEN_SOURCE_MESA
	// DriverIDIntelProprietaryWindows indicates proprietary Intel drivers for Windows
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDIntelProprietaryWindows DriverID = C.VK_DRIVER_ID_INTEL_PROPRIETARY_WINDOWS
	// DriverIDMesaRadV indicates Mesa Rad-V drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDMesaRadV DriverID = C.VK_DRIVER_ID_MESA_RADV
	// DriverIDNvidiaProprietary indicates proprietary NVidia drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDNvidiaProprietary DriverID = C.VK_DRIVER_ID_NVIDIA_PROPRIETARY
	// DriverIDQualcommProprietary indicates proprietary Qualcomm drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDQualcommProprietary DriverID = C.VK_DRIVER_ID_QUALCOMM_PROPRIETARY

	// ResolveModeAverage indicates that the result of the resolve operation is the average
	// of the sample values
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html
	ResolveModeAverage ResolveModeFlags = C.VK_RESOLVE_MODE_AVERAGE_BIT
	// ResolveModeMax indicates that the result of the resolve operation is the maximum of the
	// sample values
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html
	ResolveModeMax ResolveModeFlags = C.VK_RESOLVE_MODE_MAX_BIT
	// ResolveModeMin indicates that the result of the resolve operation is the minimum of the
	// sample values
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html
	ResolveModeMin ResolveModeFlags = C.VK_RESOLVE_MODE_MIN_BIT
	// ResolveModeNone indicates that no resolve operation is performed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html
	ResolveModeNone ResolveModeFlags = C.VK_RESOLVE_MODE_NONE
	// ResolveModeSampleZero indicates that the result of the resolve operation is equal to
	// the value of sample 0
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html
	ResolveModeSampleZero ResolveModeFlags = C.VK_RESOLVE_MODE_SAMPLE_ZERO_BIT

	// ShaderFloatControlsIndependence32BitOnly specifies that shader float controls for 32-bit
	// floating point can be set independently
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderFloatControlsIndependence.html
	ShaderFloatControlsIndependence32BitOnly ShaderFloatControlsIndependence = C.VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY
	// ShaderFloatControlsIndependenceAll specifies that shader float controls for all
	// bit widths can be set independently
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderFloatControlsIndependence.html
	ShaderFloatControlsIndependenceAll ShaderFloatControlsIndependence = C.VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL
	// ShaderFloatControlsIndependenceNone specifies that shader float controls for all bit widths
	// must be set identically
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderFloatControlsIndependence.html
	ShaderFloatControlsIndependenceNone ShaderFloatControlsIndependence = C.VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE
)

func init() {
	DriverIDAmdOpenSource.Register("AMD Open-Source")
	DriverIDAmdProprietary.Register("AMD Proprietary")
	DriverIDArmProprietary.Register("ARM Proprietary")
	DriverIDBroadcomProprietary.Register("Broadcom Proprietary")
	DriverIDGgpProprietary.Register("GGP Proprietary")
	DriverIDGoogleSwiftshader.Register("Google Swiftshader")
	DriverIDImaginationProprietary.Register("Imagination Proprietary")
	DriverIDIntelOpenSourceMesa.Register("Intel Open-Source (Mesa)")
	DriverIDIntelProprietaryWindows.Register("Intel Proprietary (Windows)")
	DriverIDMesaRadV.Register("Mesa RADV")
	DriverIDNvidiaProprietary.Register("Nvidia Proprietary")
	DriverIDQualcommProprietary.Register("Qualcomm Proprietary")

	ResolveModeAverage.Register("Average")
	ResolveModeMax.Register("Max")
	ResolveModeMin.Register("Min")
	ResolveModeNone.Register("None")
	ResolveModeSampleZero.Register("Sample Zero")

	ShaderFloatControlsIndependenceAll.Register("All")
	ShaderFloatControlsIndependenceNone.Register("None")
	ShaderFloatControlsIndependence32BitOnly.Register("32-Bit Only")
}

////

// ConformanceVersion contains the comformance test suite version the implementation is
// compliant with
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConformanceVersionKHR.html
type ConformanceVersion struct {
	// Major is the major version number of the conformance test suite
	Major uint8
	// Minor is the minor version number of the conformance test suite
	Minor uint8
	// Subminor is the subminor version number of the conformance test suite
	Subminor uint8
	// Patch is the patch version number of the conformance test suite
	Patch uint8
}

// IsAtLeast returns true if the other ConformanceVersion is at least as high as this one
func (v ConformanceVersion) IsAtLeast(other ConformanceVersion) bool {
	if v.Major > other.Major {
		return true
	} else if v.Major < other.Major {
		return false
	}

	if v.Minor > other.Minor {
		return true
	} else if v.Minor < other.Minor {
		return false
	}

	if v.Subminor > other.Subminor {
		return true
	} else if v.Subminor < other.Subminor {
		return false
	}

	return v.Patch >= other.Patch
}

////

// PhysicalDeviceDriverProperties contains driver identification information
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDriverProperties.html
type PhysicalDeviceDriverProperties struct {
	// DriverID is a unique identifier for the driver of the PhysicalDevice
	DriverID DriverID
	// DriverName is a string which is the name of the driver
	DriverName string
	// DriverInfo is a string with additional information about the driver
	DriverInfo string
	// ConformanceVersion is the version of the Vulkan conformance test thsi driver is conformant
	// against
	ConformanceVersion ConformanceVersion

	common.NextOutData
}

func (o *PhysicalDeviceDriverProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDriverProperties{})))
	}

	outData := (*C.VkPhysicalDeviceDriverProperties)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceDriverProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDeviceDriverProperties)(cDataPointer)
	o.DriverID = DriverID(outData.driverID)
	o.ConformanceVersion.Major = uint8(outData.conformanceVersion.major)
	o.ConformanceVersion.Minor = uint8(outData.conformanceVersion.minor)
	o.ConformanceVersion.Subminor = uint8(outData.conformanceVersion.subminor)
	o.ConformanceVersion.Patch = uint8(outData.conformanceVersion.patch)
	o.DriverName = C.GoString(&outData.driverName[0])
	o.DriverInfo = C.GoString(&outData.driverInfo[0])

	return outData.pNext, nil
}

////

// PhysicalDeviceDepthStencilResolveProperties describes depth/stencil resolve properties that can
// be supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html
type PhysicalDeviceDepthStencilResolveProperties struct {
	// SupportedDepthResolveModes indicates the set of supported depth resolve modes
	SupportedDepthResolveModes ResolveModeFlags
	// SupportedStencilResolveModes indicates the set of supported stencil resolve modes
	SupportedStencilResolveModes ResolveModeFlags
	// IndependentResolveNone is true if the implementation supports setting the depth
	// and stencil resolve modes to different values when one of those modes is ResolveModeNone
	IndependentResolveNone bool
	// IndependentResolve is true if the implementation supports all combinations of the supported
	// depth and stencil resolve modes, including setting either depth or stencil resolve mode to
	// ResolveModeNone
	IndependentResolve bool

	common.NextOutData
}

func (o *PhysicalDeviceDepthStencilResolveProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDepthStencilResolveProperties{})))
	}

	info := (*C.VkPhysicalDeviceDepthStencilResolveProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceDepthStencilResolveProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceDepthStencilResolveProperties)(cDataPointer)
	o.SupportedStencilResolveModes = ResolveModeFlags(info.supportedStencilResolveModes)
	o.SupportedDepthResolveModes = ResolveModeFlags(info.supportedDepthResolveModes)
	o.IndependentResolveNone = info.independentResolveNone != C.VkBool32(0)
	o.IndependentResolve = info.independentResolve != C.VkBool32(0)

	return info.pNext, nil
}

////

// PhysicalDeviceDescriptorIndexingProperties describes descriptor indexing properties
// that can be supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingProperties.html
type PhysicalDeviceDescriptorIndexingProperties struct {
	// MaxUpdateAfterBindDescriptorsInAllPools is the maximum number of descriptors (summed over
	// all descriptor types) that can be created across all pools that are created with
	// DescriptorPoolCreateUpdateAfterBind
	MaxUpdateAfterBindDescriptorsInAllPools int
	// ShaderUniformBufferArrayNonUniformIndexingNative is a boolean value indicating whether
	// uniform Buffer descriptors natively support nonuniform indexing
	ShaderUniformBufferArrayNonUniformIndexingNative bool
	// ShaderSampledImageArrayNonUniformIndexingNative is a boolean value indicating whether
	// Sampler and Image descriptors natively support nonuniform indexing
	ShaderSampledImageArrayNonUniformIndexingNative bool
	// ShaderStorageBufferArrayNonUniformIndexingNative is a boolean value indicating whether
	// storage Buffer descriptors natively support nonuniform indexing
	ShaderStorageBufferArrayNonUniformIndexingNative bool
	// ShaderStorageImageArrayNonUniformIndexingNative is a boolean value indicating whether storage
	// Image descriptors natively support nonuniform indexing
	ShaderStorageImageArrayNonUniformIndexingNative bool
	// ShaderInputAttachmentArrayNonUniformIndexingNative is a boolean value indicating whether
	// input attachment descriptors natively support nonuniform indexing
	ShaderInputAttachmentArrayNonUniformIndexingNative bool
	// RobustBufferAccessUpdateAfterBind is a boolean value indicating whether RobustBufferAccess
	// can be enabled in a Device simultaneously with DescriptorBindingUniformBufferUpdateAfterBind,
	// DescriptorBindingStorageBufferUpdateAfterBind,
	// DescriptorBindingUniformTexelBufferUpdateAfterBind, and/or
	// DescriptorBindingStorageTexelBufferUpdateAfterBind
	RobustBufferAccessUpdateAfterBind bool
	// QuadDivergentImplicitLod is a boolean value indicating whether implicit level of detail
	// calculations for Image operations have well-defined results when the Image and/or Sampler
	// objects used for the instruction are not uniform within a quad
	QuadDivergentImplicitLod bool

	// MaxPerStageDescriptorUpdateAfterBindSamplers is similar to <axPerStageDescriptorSamplers
	// but counts descriptors from descriptor sets created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindSamplers int
	// MaxPerStageDescriptorUpdateAfterBindUniformBuffers is similar to
	// MaxPerStageDescriptorUniformBuffers but counts descriptors from DescriptorSet objects
	// created with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindUniformBuffers int
	// MaxPerStageDescriptorUpdateAfterBindStorageBuffers is similar to
	// MaxPerStageDescriptorStorageBuffers but counts descriptors from DescriptorSet created with
	// or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindStorageBuffers int
	// MaxPerStageDescriptorUpdateAfterBindSampledImages is similar to
	// MaxPerStageDescriptorSampledImages but counts descriptors from DescriptorSets created with
	// or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindSampledImages int
	// MaxPerStageDescriptorUpdateAfterBindStorageImages is similar to
	// MaxPerStageDescriptorStorageImages but counts descriptors from DescriptorSet objects created
	// with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindStorageImages int
	// MaxPerStageDescriptorUpdateAfterBindInputAttachments  is similar to
	// MaxPerStageDescriptorInputAttachments but counts descriptors from DescriptorSet objects
	// created with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindInputAttachments int
	// MaxPerStageUpdateAfterBindResources is similar to MaxPerStageResources but counts
	// descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageUpdateAfterBindResources int

	// MaxDescriptorSetUpdateAfterBindSamplers is similar to MaxDescriptorSetSamplers but counts
	// descriptors from DescriptorSet created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindSamplers int
	// MaxDescriptorSetUpdateAfterBindUniformBuffers is similar to MaxDescriptorSetUniformBuffers
	// but counts descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindUniformBuffers int
	// MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic is similar to
	// MaxDescriptorSetUniformBuffersDynamic but counts descriptors from DescriptorSet objects
	// created with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic int
	// MaxDescriptorSetUpdateAfterBindStorageBuffers is similar to MaxDescriptorSetStorageBuffers
	// but counts descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindStorageBuffers int
	// MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic is similar to
	// MaxDescriptorSetStorageBuffersDynamic but counts descriptors from DescriptorSet objects
	// created with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic int
	// MaxDescriptorSetUpdateAfterBindSampledImages is similar to MaxDescriptorSetSampledImages
	// but counts descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindSampledImages int
	// MaxDescriptorSetUpdateAfterBindStorageImages is similar to MaxDescriptorSetStorageImages
	// but counts descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindStorageImages int
	// MaxDescriptorSetUpdateAfterBindInputAttachments is similar to MaxDescriptorSetInputAttachments
	// but counts descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindInputAttachments int

	common.NextOutData
}

func (o *PhysicalDeviceDescriptorIndexingProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDescriptorIndexingProperties{})))
	}

	info := (*C.VkPhysicalDeviceDescriptorIndexingProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceDescriptorIndexingProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceDescriptorIndexingProperties)(cDataPointer)

	o.MaxUpdateAfterBindDescriptorsInAllPools = int(info.maxUpdateAfterBindDescriptorsInAllPools)
	o.ShaderUniformBufferArrayNonUniformIndexingNative = info.shaderUniformBufferArrayNonUniformIndexingNative != C.VkBool32(0)
	o.ShaderSampledImageArrayNonUniformIndexingNative = info.shaderSampledImageArrayNonUniformIndexingNative != C.VkBool32(0)
	o.ShaderStorageBufferArrayNonUniformIndexingNative = info.shaderStorageBufferArrayNonUniformIndexingNative != C.VkBool32(0)
	o.ShaderStorageImageArrayNonUniformIndexingNative = info.shaderStorageImageArrayNonUniformIndexingNative != C.VkBool32(0)
	o.ShaderInputAttachmentArrayNonUniformIndexingNative = info.shaderInputAttachmentArrayNonUniformIndexingNative != C.VkBool32(0)
	o.RobustBufferAccessUpdateAfterBind = info.robustBufferAccessUpdateAfterBind != C.VkBool32(0)
	o.QuadDivergentImplicitLod = info.quadDivergentImplicitLod != C.VkBool32(0)

	o.MaxPerStageDescriptorUpdateAfterBindSamplers = int(info.maxPerStageDescriptorUpdateAfterBindSamplers)
	o.MaxPerStageDescriptorUpdateAfterBindUniformBuffers = int(info.maxPerStageDescriptorUpdateAfterBindUniformBuffers)
	o.MaxPerStageDescriptorUpdateAfterBindStorageBuffers = int(info.maxPerStageDescriptorUpdateAfterBindStorageBuffers)
	o.MaxPerStageDescriptorUpdateAfterBindSampledImages = int(info.maxPerStageDescriptorUpdateAfterBindSampledImages)
	o.MaxPerStageDescriptorUpdateAfterBindStorageImages = int(info.maxPerStageDescriptorUpdateAfterBindStorageImages)
	o.MaxPerStageDescriptorUpdateAfterBindInputAttachments = int(info.maxPerStageDescriptorUpdateAfterBindInputAttachments)
	o.MaxPerStageUpdateAfterBindResources = int(info.maxPerStageUpdateAfterBindResources)

	o.MaxDescriptorSetUpdateAfterBindSamplers = int(info.maxDescriptorSetUpdateAfterBindSamplers)
	o.MaxDescriptorSetUpdateAfterBindUniformBuffers = int(info.maxDescriptorSetUpdateAfterBindUniformBuffers)
	o.MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic = int(info.maxDescriptorSetUpdateAfterBindUniformBuffersDynamic)
	o.MaxDescriptorSetUpdateAfterBindStorageBuffers = int(info.maxDescriptorSetUpdateAfterBindStorageBuffers)
	o.MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic = int(info.maxDescriptorSetUpdateAfterBindStorageBuffersDynamic)
	o.MaxDescriptorSetUpdateAfterBindSampledImages = int(info.maxDescriptorSetUpdateAfterBindSampledImages)
	o.MaxDescriptorSetUpdateAfterBindStorageImages = int(info.maxDescriptorSetUpdateAfterBindStorageImages)
	o.MaxDescriptorSetUpdateAfterBindInputAttachments = int(info.maxDescriptorSetUpdateAfterBindInputAttachments)

	return info.pNext, nil
}

////

// PhysicalDeviceFloatControlsProperties describes properties supported by khr_shader_float_controls
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFloatControlsProperties.html
type PhysicalDeviceFloatControlsProperties struct {
	// DenormBehaviorIndependence indicates whether, and how, denorm behavior can be
	// set independently for different bit widths
	DenormBehaviorIndependence ShaderFloatControlsIndependence
	// RoundingModeIndependence indicates whether, and how, rounding modes can be set indpendently
	// for different bit widths
	RoundingModeIndependence ShaderFloatControlsIndependence

	// ShaderSignedZeroInfNanPreserveFloat16 indicates whether the sign of zero, NaN, and +/- infinity
	// can be preserved in 16-bit floating-point computations
	ShaderSignedZeroInfNanPreserveFloat16 bool
	// ShaderSignedZeroInfNanPreserveFloat32 indicates whether the sign of zero, NaN, and +/- infinity
	// can be preserved in 32-bit floating-point computations
	ShaderSignedZeroInfNanPreserveFloat32 bool
	// ShaderSignedZeroInfNanPreserveFloat64 indicates whether the sign of zero, NaN, and +/- infinity
	// can be preserved in 64-bit floating-point computations
	ShaderSignedZeroInfNanPreserveFloat64 bool
	// ShaderDenormPreserveFloat16 indicates whether denormals can be preserved in 16-bit floating-point
	// computations
	ShaderDenormPreserveFloat16 bool
	// ShaderDenormPreserveFloat32 indicates whether denormals can be preserved in 32-bit floating-point
	// computations
	ShaderDenormPreserveFloat32 bool
	// ShaderDenormPreserveFloat64 indicates whether denormals can be preserved in 64-bit floating-point
	// computations
	ShaderDenormPreserveFloat64 bool
	// ShaderDenormFlushToZeroFloat16 indicates whether denormals can be flushed to zero in 16-bit
	// floating-point computations
	ShaderDenormFlushToZeroFloat16 bool
	// ShaderDenormFlushToZeroFloat32 indicates whether denormals can be flushed to zero in 32-bit
	// floating-point computations
	ShaderDenormFlushToZeroFloat32 bool
	// ShaderDenormFlushToZeroFloat64 indicates whether denormals can be flushed to zero in 64-bit
	// floating-point computations
	ShaderDenormFlushToZeroFloat64 bool
	// ShaderRoundingModeRTEFloat16 indicates whether an implementation supports the round-to-nearest-even
	// rounding mode for 16-bit floating-point arithmetic and conversion instructions
	ShaderRoundingModeRTEFloat16 bool
	// ShaderRoundingModeRTEFloat32 indicates whether an implementation supports the round-to-nearest-even
	// rounding mode for 32-bit floating-point arithmetic and conversion instructions
	ShaderRoundingModeRTEFloat32 bool
	// ShaderRoundingModeRTEFloat64 indicates whether an implementation supports the round-to-nearest-even
	// rounding mode for 64-bit floating-point arithmetic and conversion instructions
	ShaderRoundingModeRTEFloat64 bool
	// ShaderRoundingModeRTZFloat16 indicates whether an implementation supports the round-toward-zero
	// rounding mode for 16-bit floating-point arithmetic and conversion instructions
	ShaderRoundingModeRTZFloat16 bool
	// ShaderRoundingModeRTZFloat32 indicates whether an implementation supports the round-toward-zero
	// rounding mode for 32-bit floating-point arithmetic and conversion instructions
	ShaderRoundingModeRTZFloat32 bool
	// ShaderRoundingModeRTZFloat64 indicates whether an implementation supports the round-toward-zero
	// rounding mode for 64-bit floating-point arithmetic and conversion instructions
	ShaderRoundingModeRTZFloat64 bool

	common.NextOutData
}

func (o *PhysicalDeviceFloatControlsProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceFloatControlsProperties{})))
	}

	info := (*C.VkPhysicalDeviceFloatControlsProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceFloatControlsProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceFloatControlsProperties)(cDataPointer)

	o.DenormBehaviorIndependence = ShaderFloatControlsIndependence(info.denormBehaviorIndependence)
	o.RoundingModeIndependence = ShaderFloatControlsIndependence(info.roundingModeIndependence)
	o.ShaderSignedZeroInfNanPreserveFloat16 = info.shaderSignedZeroInfNanPreserveFloat16 != C.VkBool32(0)
	o.ShaderSignedZeroInfNanPreserveFloat32 = info.shaderSignedZeroInfNanPreserveFloat32 != C.VkBool32(0)
	o.ShaderSignedZeroInfNanPreserveFloat64 = info.shaderSignedZeroInfNanPreserveFloat64 != C.VkBool32(0)
	o.ShaderDenormPreserveFloat16 = info.shaderDenormPreserveFloat16 != C.VkBool32(0)
	o.ShaderDenormPreserveFloat32 = info.shaderDenormPreserveFloat32 != C.VkBool32(0)
	o.ShaderDenormPreserveFloat64 = info.shaderDenormPreserveFloat64 != C.VkBool32(0)
	o.ShaderDenormFlushToZeroFloat16 = info.shaderDenormFlushToZeroFloat16 != C.VkBool32(0)
	o.ShaderDenormFlushToZeroFloat32 = info.shaderDenormFlushToZeroFloat32 != C.VkBool32(0)
	o.ShaderDenormFlushToZeroFloat64 = info.shaderDenormFlushToZeroFloat64 != C.VkBool32(0)
	o.ShaderRoundingModeRTEFloat16 = info.shaderRoundingModeRTEFloat16 != C.VkBool32(0)
	o.ShaderRoundingModeRTEFloat32 = info.shaderRoundingModeRTEFloat32 != C.VkBool32(0)
	o.ShaderRoundingModeRTEFloat64 = info.shaderRoundingModeRTEFloat64 != C.VkBool32(0)
	o.ShaderRoundingModeRTZFloat16 = info.shaderRoundingModeRTZFloat16 != C.VkBool32(0)
	o.ShaderRoundingModeRTZFloat32 = info.shaderRoundingModeRTZFloat32 != C.VkBool32(0)
	o.ShaderRoundingModeRTZFloat64 = info.shaderRoundingModeRTZFloat64 != C.VkBool32(0)

	return info.pNext, nil
}

////

// PhysicalDeviceSamplerFilterMinmaxProperties describes Sampler filter minmax limits that can
// be supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSamplerFilterMinmaxProperties.html
type PhysicalDeviceSamplerFilterMinmaxProperties struct {
	// FilterMinmaxSingleComponentFormats indicates whether a minimum set of required formats
	// support min/max filtering
	FilterMinmaxSingleComponentFormats bool
	// FilterMinmaxImageComponentMapping indicates whether the implementation support non-identity
	// component mapping of the Image when doing min/max filtering
	FilterMinmaxImageComponentMapping bool

	common.NextOutData
}

func (o *PhysicalDeviceSamplerFilterMinmaxProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceSamplerFilterMinmaxProperties{})))
	}

	info := (*C.VkPhysicalDeviceSamplerFilterMinmaxProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceSamplerFilterMinmaxProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceSamplerFilterMinmaxProperties)(cDataPointer)

	o.FilterMinmaxSingleComponentFormats = info.filterMinmaxSingleComponentFormats != C.VkBool32(0)
	o.FilterMinmaxImageComponentMapping = info.filterMinmaxImageComponentMapping != C.VkBool32(0)

	return info.pNext, nil
}

////

// PhysicalDeviceTimelineSemaphoreProperties describes timeline Semaphore properties that
// can be supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTimelineSemaphoreProperties.html
type PhysicalDeviceTimelineSemaphoreProperties struct {
	// MaxTimelineSemaphoreValueDifference indicates the maximum difference allowed by the
	// implementation between the current value of a timeline Semaphore and any pending signal or
	// wait operations
	MaxTimelineSemaphoreValueDifference uint64

	common.NextOutData
}

func (o *PhysicalDeviceTimelineSemaphoreProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceTimelineSemaphoreProperties{})))
	}

	info := (*C.VkPhysicalDeviceTimelineSemaphoreProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceTimelineSemaphoreProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceTimelineSemaphoreProperties)(cDataPointer)

	o.MaxTimelineSemaphoreValueDifference = uint64(info.maxTimelineSemaphoreValueDifference)

	return info.pNext, nil
}

////

// PhysicalDeviceVulkan11Properties specifies PhysicalDevice properties for functionality
// promoted to Vulkan 1.1
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan11Properties.html
type PhysicalDeviceVulkan11Properties struct {
	// DeviceUUID represents a universally-unique identifier for the Device
	DeviceUUID uuid.UUID
	// DriverUUID represents a universally-unique identifier for the driver build in use by
	// the Device
	DriverUUID uuid.UUID
	// DeviceLUID represents a locally-unique identifier for the Device
	DeviceLUID uint64

	// DeviceNodeMask identifies the node within a linked Device adapter corresponding to the
	// Device
	DeviceNodeMask uint32
	// DeviceLUIDValid is true if DeviceLUID contains a valid LUID and DeviceNodeMask contains
	// a valid node mask
	DeviceLUIDValid bool

	// SubgroupSize is the default number of invocations in each subgroup
	SubgroupSize int
	// SubgroupSupportedStages describes the shader stages that group operations with
	// subgroup scope are supported in
	SubgroupSupportedStages core1_0.ShaderStageFlags
	// SubgroupSupportedOperations specifies the sets of group operations with subgroup
	// scope supported on this Device
	SubgroupSupportedOperations core1_1.SubgroupFeatureFlags
	// SubgroupQuadOperationsInAllStages specifies whether quad group operations are available
	// in all stages, or are restricted to fragment and compute stages
	SubgroupQuadOperationsInAllStages bool

	// PointClippingBehavior specifies the point clipping behavior supported by the implementation
	PointClippingBehavior core1_1.PointClippingBehavior
	// MaxMultiviewViewCount is one greater than the maximum view index that can be used in a
	// subpass
	MaxMultiviewViewCount int
	// MaxMultiviewInstanceIndex is the maximum valid value of instance index allowed to be
	// generated by a drawing command recorded within a subpass of a multiview RenderPass instance
	MaxMultiviewInstanceIndex int
	// ProtectedNoFault specifies how an implementation behaves when an application attempts to write
	// to unprotected memory in a protected Queue operation, read from protected memory in an
	// unprotected Queue operation, or perform a query in a protected Queue operation
	ProtectedNoFault bool
	// MaxPerSetDescriptors is a maximum number of descriptors (summed over all descriptor types)
	// in a single DescriptorSet that is guaranteed to satisfy any implementation-dependent contraints
	// on the size of a DescriptorSet itself
	MaxPerSetDescriptors int
	// MaxMemoryAllocationSize is the maximum size of a memory allocation that can be created,
	// even if there is more space available in the heap
	MaxMemoryAllocationSize int

	common.NextOutData
}

func (o *PhysicalDeviceVulkan11Properties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceVulkan11Properties)
	}

	info := (*C.VkPhysicalDeviceVulkan11Properties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_1_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceVulkan11Properties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceVulkan11Properties)(cDataPointer)

	deviceUUIDBytes := C.GoBytes(unsafe.Pointer(&info.deviceUUID[0]), C.VK_UUID_SIZE)
	o.DeviceUUID, err = uuid.FromBytes(deviceUUIDBytes)
	if err != nil {
		return nil, errors.Wrap(err, "vulkan provided invalid device uuid")
	}

	driverUUIDBytes := C.GoBytes(unsafe.Pointer(&info.driverUUID[0]), C.VK_UUID_SIZE)
	o.DriverUUID, err = uuid.FromBytes(driverUUIDBytes)
	if err != nil {
		return nil, errors.Wrap(err, "vulkan provided invalid driver uuid")
	}

	o.DeviceLUID = *(*uint64)(unsafe.Pointer(&info.deviceLUID[0]))
	o.DeviceNodeMask = uint32(info.deviceNodeMask)
	o.DeviceLUIDValid = info.deviceLUIDValid != C.VkBool32(0)
	o.SubgroupSize = int(info.subgroupSize)
	o.SubgroupSupportedStages = core1_0.ShaderStageFlags(info.subgroupSupportedStages)
	o.SubgroupSupportedOperations = core1_1.SubgroupFeatureFlags(info.subgroupSupportedOperations)
	o.SubgroupQuadOperationsInAllStages = info.subgroupQuadOperationsInAllStages != C.VkBool32(0)
	o.PointClippingBehavior = core1_1.PointClippingBehavior(info.pointClippingBehavior)
	o.MaxMultiviewViewCount = int(info.maxMultiviewViewCount)
	o.MaxMultiviewInstanceIndex = int(info.maxMultiviewInstanceIndex)
	o.ProtectedNoFault = info.protectedNoFault != C.VkBool32(0)
	o.MaxPerSetDescriptors = int(info.maxPerSetDescriptors)
	o.MaxMemoryAllocationSize = int(info.maxMemoryAllocationSize)

	return info.pNext, nil
}

////

// PhysicalDeviceVulkan12Properties specifies PhysicalDevice properties for functionality
// promoted to Vulkan 1.2
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan12Properties.html
type PhysicalDeviceVulkan12Properties struct {
	// DriverID is a unique identifier for the driver of the PhysicalDevice
	DriverID DriverID
	// DriverName is a string which is the name of the driver
	DriverName string
	// DriverInfo is a string with additional information about the driver
	DriverInfo string
	// ConformanceVersion is the version of the Vulkan conformance test this driver is
	// comformant against
	ConformanceVersion ConformanceVersion

	// DenormBehaviorIndependence indicates whether, and how, denorm behavior can be set
	// independently for different bit widths
	DenormBehaviorIndependence ShaderFloatControlsIndependence
	// RoundingModeIndependence indicates whether, and how, rounding modes can be set
	// independently for different bit widths
	RoundingModeIndependence ShaderFloatControlsIndependence

	// ShaderSignedZeroInfNanPreserveFloat16 indicates whether the sign of zero, NaN, and
	// +/- infinity can be preserved in 16-bit floating-point computations
	ShaderSignedZeroInfNanPreserveFloat16 bool
	// ShaderSignedZeroInfNanPreserveFloat32 indicates whether the sign of zero, NaN, and
	// +/- infinity can be preserved in 32-bit floating-point computations
	ShaderSignedZeroInfNanPreserveFloat32 bool
	// ShaderSignedZeroInfNanPreserveFloat64 indicates whether the sign of zero, NaN, and
	// +/- infinity can be preserved in 64-bit floating-point computations
	ShaderSignedZeroInfNanPreserveFloat64 bool
	// ShaderDenormPreserveFloat16 indicates whether denormals can be preserved in 16-bit
	// floating-point computations
	ShaderDenormPreserveFloat16 bool
	// ShaderDenormPreserveFloat32 indicates whether denormals can be preserved in 32-bit
	// floating-point computations
	ShaderDenormPreserveFloat32 bool
	// ShaderDenormPreserveFloat64 indicates whether denormals can be preserved in 64-bit
	// floating-point computations
	ShaderDenormPreserveFloat64 bool
	// ShaderDenormFlushToZeroFloat16 indicates whether denormals can be flushed to zero
	// in 16-bit floating-point computations
	ShaderDenormFlushToZeroFloat16 bool
	// ShaderDenormFlushToZeroFloat32 indicates whether denormals can be flushed to zero
	// in 32-bit floating-point computations
	ShaderDenormFlushToZeroFloat32 bool
	// ShaderDenormFlushToZeroFloat64 indicates whether denormals can be flushed to zero
	// in 64-bit floating-point computations
	ShaderDenormFlushToZeroFloat64 bool
	// ShaderRoundingModeRTEFloat16 indicates whether an implementation supports the
	// round-to-nearest-even rounding mode for 16-bit floating-point arithmetic and conversion
	// instructions
	ShaderRoundingModeRTEFloat16 bool
	// ShaderRoundingModeRTEFloat32 indicates whether an implementation supports the
	// round-to-nearest-even rounding mode for 32-bit floating-point arithmetic and conversion
	// instructions
	ShaderRoundingModeRTEFloat32 bool
	// ShaderRoundingModeRTEFloat64 indicates whether an implementation supports the
	// round-to-nearest-even rounding mode for 64-bit floating-point arithmetic and conversion
	// instructions
	ShaderRoundingModeRTEFloat64 bool
	// ShaderRoundingModeRTZFloat16 indicates whether an implementation supports the
	// round-towards-zero rounding mode for 16-bit floating-point arithmetic and conversion
	// instructions
	ShaderRoundingModeRTZFloat16 bool
	// ShaderRoundingModeRTZFloat32 indicates whether an implementation supports the
	// round-towards-zero rounding mode for 32-bit floating-point arithmetic and conversion
	// instructions
	ShaderRoundingModeRTZFloat32 bool
	// ShaderRoundingModeRTZFloat64 indicates whether an implementation supports the
	// round-towards-zero rounding mode for 64-bit floating-point arithmetic and conversion
	// instructions
	ShaderRoundingModeRTZFloat64 bool

	// MaxUpdateAfterBindDescriptorsInAllPools is the maximum number of descriptors
	// (summed over all descriptor types) that can be created across all pools that are
	// created with DescriptorPoolCreateUpdateAfterBind
	MaxUpdateAfterBindDescriptorsInAllPools int
	// ShaderUniformBufferArrayNonUniformIndexingNative indicates whether uniform Buffer
	// descriptors natively support nonuniform indexing
	ShaderUniformBufferArrayNonUniformIndexingNative bool
	// ShaderSampledImageArrayNonUniformIndexingNative indicates whether Sampler and Image
	// descriptors natively support nonuniform indexing
	ShaderSampledImageArrayNonUniformIndexingNative bool
	// ShaderStorageBufferArrayNonUniformIndexingNative indicates whether storage Buffer
	// descriptors natively support nonuniform indexing
	ShaderStorageBufferArrayNonUniformIndexingNative bool
	// ShaderStorageImageArrayNonUniformIndexingNative indicates whether storage Image
	// descriptors natively support nonuniform indexing
	ShaderStorageImageArrayNonUniformIndexingNative bool
	// ShaderInputAttachmentArrayNonUniformIndexingNative indicates whether input attachment
	// descriptors natively support nonuniform indexing
	ShaderInputAttachmentArrayNonUniformIndexingNative bool

	// RobustBufferAccessUpdateAfterBind indicates whether RobustBufferAccess can be enabled
	// in a Device simultaneously with DescriptorBindingUniformBufferUpdateAfterBind,
	// DescriptorBindingStorageBufferUpdateAfterBind,
	// DescriptorBindingUniformTexelBufferUpdateAfterBind, and/or
	// DescriptorBindingStorageTexelBufferUpdateAfterBind
	RobustBufferAccessUpdateAfterBind bool
	// QuadDivergentImplicitLod indicates whether imlicit level of detail calculations for Image
	// operations have well-defined results when the Image and/or Sampler objects used for
	// the instructions are not uniform within a quad
	QuadDivergentImplicitLod bool

	// MaxPerStageDescriptorUpdateAfterBindSamplers is similar to MaxPerStageDescriptorSamplers
	// but counts descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindSamplers int
	// MaxPerStageDescriptorUpdateAfterBindUniformBuffers is similar to
	// MaxPerStageDescriptorUniformBuffers but counts descriptors from DescriptorSet objects
	// created with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindUniformBuffers int
	// MaxPerStageDescriptorUpdateAfterBindStorageBuffers is similar to
	// MaxPerStageDescriptorStorageBuffers but counts descriptors from DescriptorSet objects
	// created with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindStorageBuffers int
	// MaxPerStageDescriptorUpdateAfterBindSampledImages is similar to
	// MaxPerStageDescriptorSampledImages but counts descriptors from DescriptorSet objects
	// created with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindSampledImages int
	// MaxPerStageDescriptorUpdateAfterBindStorageImages is similar to
	// MaxPerStageDescriptorStorageImages but counts descriptors from DescriptorSet objects
	// created with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindStorageImages int
	// MaxPerStageDescriptorUpdateAfterBindInputAttachments is similar to
	// MaxPerStageDescriptorInputAttachments but counts descriptors from DescriptorSet objects
	// created with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindInputAttachments int
	// MaxPerStageUpdateAfterBindResources is similar to MaxPerStageResources but counts
	// descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageUpdateAfterBindResources int

	// MaxDescriptorSetUpdateAfterBindSamplers is similar to MaxDescriptorSetSamplers but counts
	// descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindSamplers int
	// MaxDescriptorSetUpdateAfterBindUniformBuffers is similar to MaxDescriptorSetUniformBuffers
	// but counts descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindUniformBuffers int
	// MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic is similar to
	// MaxDescriptorSetUniformBuffersDynamic but counts descriptors from DescriptorSet objects
	// created with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic int
	// MaxDescriptorSetUpdateAfterBindStorageBuffers is similar to MaxDescriptorSetStorageBuffers
	// but counts descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindStorageBuffers int
	// MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic is similar to
	// MaxDescriptorSetStorageBuffersDynamic but counts descriptors from DescriptorSet objects
	// created with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic int
	// MaxDescriptorSetUpdateAfterBindSampledImages is similar to MaxDescriptorSetSampledImages
	// but counts descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindSampledImages int
	// MaxDescriptorSetUpdateAfterBindStorageImages is similar to MaxDescriptorSetStorageImages
	// but counts descriptors from descriptor sets created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindStorageImages int
	// MaxDescriptorSetUpdateAfterBindInputAttachments is similar to
	// MaxDescriptorSetInputAttachments but counts descriptors from DescriptorSet objects
	// created with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindInputAttachments int

	// SupportedDepthResolveModes indicates the set of supported depth resolve modes
	SupportedDepthResolveModes ResolveModeFlags
	// SupportedStencilResolveModes idnicates the set of supported stencil resolve modes
	SupportedStencilResolveModes ResolveModeFlags
	// IndependentResolveNone is true if the implementation supports setting the depth and
	// stencil resolve modes to different values when one of those modes is ResolveModeNone
	IndependentResolveNone bool
	// IndependentResolve is true if the implementation supports all combinations of the supported
	// depth and stencil resolve modes, including setting either depth or stencil resolve mode
	// to ResolveModeNone
	IndependentResolve bool

	// FilterMinmaxSingleComponentFormats indicates whether a minimum set of required formats
	// support min/max filtering
	FilterMinmaxSingleComponentFormats bool
	// FilterMinmaxImageComponentMapping indicates whether the implementation supports non-identity
	// component mapping of the Image when doing min/max filtering
	FilterMinmaxImageComponentMapping bool

	// MaxTimelineSemaphoreValueDifference indicates the maximum difference allowed by the
	// implementation between the current value of a timeline Semaphore and any pending
	// signal or wait operations
	MaxTimelineSemaphoreValueDifference uint64
	// FramebufferIntegerColorSampleCounts indicates the color sample counts that are supported
	// for all Framebuffer color attachments with integer formats
	FramebufferIntegerColorSampleCounts core1_0.SampleCountFlags

	common.NextOutData
}

func (o *PhysicalDeviceVulkan12Properties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceVulkan12Properties)
	}

	outData := (*C.VkPhysicalDeviceVulkan12Properties)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_2_PROPERTIES
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceVulkan12Properties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDeviceVulkan12Properties)(cDataPointer)

	o.DriverID = DriverID(outData.driverID)
	o.ConformanceVersion.Major = uint8(outData.conformanceVersion.major)
	o.ConformanceVersion.Minor = uint8(outData.conformanceVersion.minor)
	o.ConformanceVersion.Subminor = uint8(outData.conformanceVersion.subminor)
	o.ConformanceVersion.Patch = uint8(outData.conformanceVersion.patch)
	o.DriverName = C.GoString(&outData.driverName[0])
	o.DriverInfo = C.GoString(&outData.driverInfo[0])

	o.DenormBehaviorIndependence = ShaderFloatControlsIndependence(outData.denormBehaviorIndependence)
	o.RoundingModeIndependence = ShaderFloatControlsIndependence(outData.roundingModeIndependence)

	o.ShaderSignedZeroInfNanPreserveFloat16 = outData.shaderSignedZeroInfNanPreserveFloat16 != C.VkBool32(0)
	o.ShaderSignedZeroInfNanPreserveFloat32 = outData.shaderSignedZeroInfNanPreserveFloat32 != C.VkBool32(0)
	o.ShaderSignedZeroInfNanPreserveFloat64 = outData.shaderSignedZeroInfNanPreserveFloat64 != C.VkBool32(0)
	o.ShaderDenormPreserveFloat16 = outData.shaderDenormPreserveFloat16 != C.VkBool32(0)
	o.ShaderDenormPreserveFloat32 = outData.shaderDenormPreserveFloat32 != C.VkBool32(0)
	o.ShaderDenormPreserveFloat64 = outData.shaderDenormPreserveFloat64 != C.VkBool32(0)
	o.ShaderDenormFlushToZeroFloat16 = outData.shaderDenormFlushToZeroFloat16 != C.VkBool32(0)
	o.ShaderDenormFlushToZeroFloat32 = outData.shaderDenormFlushToZeroFloat32 != C.VkBool32(0)
	o.ShaderDenormFlushToZeroFloat64 = outData.shaderDenormFlushToZeroFloat64 != C.VkBool32(0)
	o.ShaderRoundingModeRTEFloat16 = outData.shaderRoundingModeRTEFloat16 != C.VkBool32(0)
	o.ShaderRoundingModeRTEFloat32 = outData.shaderRoundingModeRTEFloat32 != C.VkBool32(0)
	o.ShaderRoundingModeRTEFloat64 = outData.shaderRoundingModeRTEFloat64 != C.VkBool32(0)
	o.ShaderRoundingModeRTZFloat16 = outData.shaderRoundingModeRTZFloat16 != C.VkBool32(0)
	o.ShaderRoundingModeRTZFloat32 = outData.shaderRoundingModeRTZFloat32 != C.VkBool32(0)
	o.ShaderRoundingModeRTZFloat64 = outData.shaderRoundingModeRTZFloat64 != C.VkBool32(0)

	o.MaxUpdateAfterBindDescriptorsInAllPools = int(outData.maxUpdateAfterBindDescriptorsInAllPools)

	o.ShaderUniformBufferArrayNonUniformIndexingNative = outData.shaderUniformBufferArrayNonUniformIndexingNative != C.VkBool32(0)
	o.ShaderSampledImageArrayNonUniformIndexingNative = outData.shaderSampledImageArrayNonUniformIndexingNative != C.VkBool32(0)
	o.ShaderStorageBufferArrayNonUniformIndexingNative = outData.shaderStorageBufferArrayNonUniformIndexingNative != C.VkBool32(0)
	o.ShaderStorageImageArrayNonUniformIndexingNative = outData.shaderStorageImageArrayNonUniformIndexingNative != C.VkBool32(0)
	o.ShaderInputAttachmentArrayNonUniformIndexingNative = outData.shaderInputAttachmentArrayNonUniformIndexingNative != C.VkBool32(0)

	o.RobustBufferAccessUpdateAfterBind = outData.robustBufferAccessUpdateAfterBind != C.VkBool32(0)
	o.QuadDivergentImplicitLod = outData.quadDivergentImplicitLod != C.VkBool32(0)

	o.MaxPerStageDescriptorUpdateAfterBindSamplers = int(outData.maxPerStageDescriptorUpdateAfterBindSamplers)
	o.MaxPerStageDescriptorUpdateAfterBindUniformBuffers = int(outData.maxPerStageDescriptorUpdateAfterBindUniformBuffers)
	o.MaxPerStageDescriptorUpdateAfterBindStorageBuffers = int(outData.maxPerStageDescriptorUpdateAfterBindStorageBuffers)
	o.MaxPerStageDescriptorUpdateAfterBindSampledImages = int(outData.maxPerStageDescriptorUpdateAfterBindSampledImages)
	o.MaxPerStageDescriptorUpdateAfterBindStorageImages = int(outData.maxPerStageDescriptorUpdateAfterBindStorageImages)
	o.MaxPerStageDescriptorUpdateAfterBindInputAttachments = int(outData.maxPerStageDescriptorUpdateAfterBindInputAttachments)
	o.MaxPerStageUpdateAfterBindResources = int(outData.maxPerStageUpdateAfterBindResources)

	o.MaxDescriptorSetUpdateAfterBindSamplers = int(outData.maxDescriptorSetUpdateAfterBindSamplers)
	o.MaxDescriptorSetUpdateAfterBindUniformBuffers = int(outData.maxDescriptorSetUpdateAfterBindUniformBuffers)
	o.MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic = int(outData.maxDescriptorSetUpdateAfterBindUniformBuffersDynamic)
	o.MaxDescriptorSetUpdateAfterBindStorageBuffers = int(outData.maxDescriptorSetUpdateAfterBindStorageBuffers)
	o.MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic = int(outData.maxDescriptorSetUpdateAfterBindStorageBuffersDynamic)
	o.MaxDescriptorSetUpdateAfterBindSampledImages = int(outData.maxDescriptorSetUpdateAfterBindSampledImages)
	o.MaxDescriptorSetUpdateAfterBindStorageImages = int(outData.maxDescriptorSetUpdateAfterBindStorageImages)
	o.MaxDescriptorSetUpdateAfterBindInputAttachments = int(outData.maxDescriptorSetUpdateAfterBindInputAttachments)

	o.SupportedDepthResolveModes = ResolveModeFlags(outData.supportedDepthResolveModes)
	o.SupportedStencilResolveModes = ResolveModeFlags(outData.supportedStencilResolveModes)
	o.IndependentResolveNone = outData.independentResolveNone != C.VkBool32(0)
	o.IndependentResolve = outData.independentResolve != C.VkBool32(0)

	o.FilterMinmaxSingleComponentFormats = outData.filterMinmaxSingleComponentFormats != C.VkBool32(0)
	o.FilterMinmaxImageComponentMapping = outData.filterMinmaxImageComponentMapping != C.VkBool32(0)

	o.MaxTimelineSemaphoreValueDifference = uint64(outData.maxTimelineSemaphoreValueDifference)
	o.FramebufferIntegerColorSampleCounts = core1_0.SampleCountFlags(outData.framebufferIntegerColorSampleCounts)

	return outData.pNext, nil
}
