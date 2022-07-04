package core1_2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"unsafe"
)

type DriverID int32

var driverIDMapping = make(map[DriverID]string)

func (e DriverID) Register(str string) {
	driverIDMapping[e] = str
}

func (e DriverID) String() string {
	return driverIDMapping[e]
}

type ResolveModeFlags int32

var resolveModeFlagsMapping = common.NewFlagStringMapping[ResolveModeFlags]()

func (f ResolveModeFlags) Register(str string) {
	resolveModeFlagsMapping.Register(f, str)
}

func (f ResolveModeFlags) String() string {
	return resolveModeFlagsMapping.FlagsToString(f)
}

////

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
	MaxDriverInfoSize int = C.VK_MAX_DRIVER_INFO_SIZE
	MaxDriverNameSize int = C.VK_MAX_DRIVER_NAME_SIZE

	DriverIDAmdOpenSource           DriverID = C.VK_DRIVER_ID_AMD_OPEN_SOURCE
	DriverIDAmdProprietary          DriverID = C.VK_DRIVER_ID_AMD_PROPRIETARY
	DriverIDArmProprietary          DriverID = C.VK_DRIVER_ID_ARM_PROPRIETARY
	DriverIDBroadcomProprietary     DriverID = C.VK_DRIVER_ID_BROADCOM_PROPRIETARY
	DriverIDGgpProprietary          DriverID = C.VK_DRIVER_ID_GGP_PROPRIETARY
	DriverIDGoogleSwiftshader       DriverID = C.VK_DRIVER_ID_GOOGLE_SWIFTSHADER
	DriverIDImaginationProprietary  DriverID = C.VK_DRIVER_ID_IMAGINATION_PROPRIETARY
	DriverIDIntelOpenSourceMesa     DriverID = C.VK_DRIVER_ID_INTEL_OPEN_SOURCE_MESA
	DriverIDIntelProprietaryWindows DriverID = C.VK_DRIVER_ID_INTEL_PROPRIETARY_WINDOWS
	DriverIDMesaRadV                DriverID = C.VK_DRIVER_ID_MESA_RADV
	DriverIDNvidiaProprietary       DriverID = C.VK_DRIVER_ID_NVIDIA_PROPRIETARY
	DriverIDQualcommProprietary     DriverID = C.VK_DRIVER_ID_QUALCOMM_PROPRIETARY

	ResolveModeAverage    ResolveModeFlags = C.VK_RESOLVE_MODE_AVERAGE_BIT
	ResolveModeMax        ResolveModeFlags = C.VK_RESOLVE_MODE_MAX_BIT
	ResolveModeMin        ResolveModeFlags = C.VK_RESOLVE_MODE_MIN_BIT
	ResolveModeNone       ResolveModeFlags = C.VK_RESOLVE_MODE_NONE
	ResolveModeSampleZero ResolveModeFlags = C.VK_RESOLVE_MODE_SAMPLE_ZERO_BIT

	ShaderFloatControlsIndependence32BitOnly ShaderFloatControlsIndependence = C.VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY
	ShaderFloatControlsIndependenceAll       ShaderFloatControlsIndependence = C.VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL
	ShaderFloatControlsIndependenceNone      ShaderFloatControlsIndependence = C.VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE
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

type ConformanceVersion struct {
	Major    uint8
	Minor    uint8
	Subminor uint8
	Patch    uint8
}

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

type PhysicalDeviceDriverOutData struct {
	DriverID           DriverID
	DriverName         string
	DriverInfo         string
	ConformanceVersion ConformanceVersion

	common.NextOutData
}

func (o *PhysicalDeviceDriverOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDriverProperties{})))
	}

	outData := (*C.VkPhysicalDeviceDriverProperties)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceDriverOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
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

type PhysicalDeviceDepthStencilResolveOutData struct {
	SupportedDepthResolveModes   ResolveModeFlags
	SupportedStencilResolveModes ResolveModeFlags
	IndependentResolveNone       bool
	IndependentResolve           bool

	common.NextOutData
}

func (o *PhysicalDeviceDepthStencilResolveOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDepthStencilResolveProperties{})))
	}

	info := (*C.VkPhysicalDeviceDepthStencilResolveProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceDepthStencilResolveOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceDepthStencilResolveProperties)(cDataPointer)
	o.SupportedStencilResolveModes = ResolveModeFlags(info.supportedStencilResolveModes)
	o.SupportedDepthResolveModes = ResolveModeFlags(info.supportedDepthResolveModes)
	o.IndependentResolveNone = info.independentResolveNone != C.VkBool32(0)
	o.IndependentResolve = info.independentResolve != C.VkBool32(0)

	return info.pNext, nil
}

////

type PhysicalDeviceDescriptorIndexingOutData struct {
	MaxUpdateAfterBindDescriptorsInAllPools            int
	ShaderUniformBufferArrayNonUniformIndexingNative   bool
	ShaderSampledImageArrayNonUniformIndexingNative    bool
	ShaderStorageBufferArrayNonUniformIndexingNative   bool
	ShaderStorageImageArrayNonUniformIndexingNative    bool
	ShaderInputAttachmentArrayNonUniformIndexingNative bool
	RobustBufferAccessUpdateAfterBind                  bool
	QuadDivergentImplicitLod                           bool

	MaxPerStageDescriptorUpdateAfterBindSamplers         int
	MaxPerStageDescriptorUpdateAfterBindUniformBuffers   int
	MaxPerStageDescriptorUpdateAfterBindStorageBuffers   int
	MaxPerStageDescriptorUpdateAfterBindSampledImages    int
	MaxPerStageDescriptorUpdateAfterBindStorageImages    int
	MaxPerStageDescriptorUpdateAfterBindInputAttachments int
	MaxPerStageUpdateAfterBindResources                  int

	MaxDescriptorSetUpdateAfterBindSamplers              int
	MaxDescriptorSetUpdateAfterBindUniformBuffers        int
	MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic int
	MaxDescriptorSetUpdateAfterBindStorageBuffers        int
	MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic int
	MaxDescriptorSetUpdateAfterBindSampledImages         int
	MaxDescriptorSetUpdateAfterBindStorageImages         int
	MaxDescriptorSetUpdateAfterBindInputAttachments      int

	common.NextOutData
}

func (o *PhysicalDeviceDescriptorIndexingOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDescriptorIndexingProperties{})))
	}

	info := (*C.VkPhysicalDeviceDescriptorIndexingProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceDescriptorIndexingOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
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

type PhysicalDeviceFloatControlsOutData struct {
	DenormBehaviorIndependence ShaderFloatControlsIndependence
	RoundingMoundIndependence  ShaderFloatControlsIndependence

	ShaderSignedZeroInfNanPreserveFloat16 bool
	ShaderSignedZeroInfNanPreserveFloat32 bool
	ShaderSignedZeroInfNanPreserveFloat64 bool
	ShaderDenormPreserveFloat16           bool
	ShaderDenormPreserveFloat32           bool
	ShaderDenormPreserveFloat64           bool
	ShaderDenormFlushToZeroFloat16        bool
	ShaderDenormFlushToZeroFloat32        bool
	ShaderDenormFlushToZeroFloat64        bool
	ShaderRoundingModeRTEFloat16          bool
	ShaderRoundingModeRTEFloat32          bool
	ShaderRoundingModeRTEFloat64          bool
	ShaderRoundingModeRTZFloat16          bool
	ShaderRoundingModeRTZFloat32          bool
	ShaderRoundingModeRTZFloat64          bool

	common.NextOutData
}

func (o *PhysicalDeviceFloatControlsOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceFloatControlsProperties{})))
	}

	info := (*C.VkPhysicalDeviceFloatControlsProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceFloatControlsOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceFloatControlsProperties)(cDataPointer)

	o.DenormBehaviorIndependence = ShaderFloatControlsIndependence(info.denormBehaviorIndependence)
	o.RoundingMoundIndependence = ShaderFloatControlsIndependence(info.roundingModeIndependence)
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

type PhysicalDeviceSamplerFilterMinmaxOutData struct {
	FilterMinmaxSingleComponentFormats bool
	FilterMinmaxImageComponentMapping  bool

	common.NextOutData
}

func (o *PhysicalDeviceSamplerFilterMinmaxOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceSamplerFilterMinmaxProperties{})))
	}

	info := (*C.VkPhysicalDeviceSamplerFilterMinmaxProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceSamplerFilterMinmaxOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceSamplerFilterMinmaxProperties)(cDataPointer)

	o.FilterMinmaxSingleComponentFormats = info.filterMinmaxSingleComponentFormats != C.VkBool32(0)
	o.FilterMinmaxImageComponentMapping = info.filterMinmaxImageComponentMapping != C.VkBool32(0)

	return info.pNext, nil
}

////

type PhysicalDeviceTimelineSemaphoreOutData struct {
	MaxTimelineSemaphoreValueDifference uint64

	common.NextOutData
}

func (o *PhysicalDeviceTimelineSemaphoreOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceTimelineSemaphoreProperties{})))
	}

	info := (*C.VkPhysicalDeviceTimelineSemaphoreProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceTimelineSemaphoreOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceTimelineSemaphoreProperties)(cDataPointer)

	o.MaxTimelineSemaphoreValueDifference = uint64(info.maxTimelineSemaphoreValueDifference)

	return info.pNext, nil
}

////

type PhysicalDeviceVulkan11OutData struct {
	DeviceUUID uuid.UUID
	DriverUUID uuid.UUID
	DeviceLUID uint64

	DeviceNodeMask  uint32
	DeviceLUIDValid bool

	SubgroupSize                      int
	SubgroupSupportedStages           core1_0.ShaderStages
	SubgroupSupportedOperations       core1_1.SubgroupFeatures
	SubgroupQuadOperationsInAllStages bool

	PointClippingBehavior     core1_1.PointClippingBehavior
	MaxMultiviewViewCount     int
	MaxMultiviewInstanceIndex int
	ProtectedNoFault          bool
	MaxPerSetDescriptors      int
	MaxMemoryAllocationSize   int

	common.NextOutData
}

func (o *PhysicalDeviceVulkan11OutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceVulkan11Properties)
	}

	info := (*C.VkPhysicalDeviceVulkan11Properties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_1_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceVulkan11OutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
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
	o.SubgroupSupportedStages = core1_0.ShaderStages(info.subgroupSupportedStages)
	o.SubgroupSupportedOperations = core1_1.SubgroupFeatures(info.subgroupSupportedOperations)
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

type PhysicalDeviceVulkan12OutData struct {
	DriverID           DriverID
	DriverName         string
	DriverInfo         string
	ConformanceVersion ConformanceVersion

	DenormBehaviorIndependence ShaderFloatControlsIndependence
	RoundingModeIndependence   ShaderFloatControlsIndependence

	ShaderSignedZeroInfNanPreserveFloat16 bool
	ShaderSignedZeroInfNanPreserveFloat32 bool
	ShaderSignedZeroInfNanPreserveFloat64 bool
	ShaderDenormPreserveFloat16           bool
	ShaderDenormPreserveFloat32           bool
	ShaderDenormPreserveFloat64           bool
	ShaderDenormFlushToZeroFloat16        bool
	ShaderDenormFlushToZeroFloat32        bool
	ShaderDenormFlushToZeroFloat64        bool
	ShaderRoundingModeRTEFloat16          bool
	ShaderRoundingModeRTEFloat32          bool
	ShaderRoundingModeRTEFloat64          bool
	ShaderRoundingModeRTZFloat16          bool
	ShaderRoundingModeRTZFloat32          bool
	ShaderRoundingModeRTZFloat64          bool

	MaxUpdateAfterBindDescriptorsInAllPools            int
	ShaderUniformBufferArrayNonUniformIndexingNative   bool
	ShaderSampledImageArrayNonUniformIndexingNative    bool
	ShaderStorageBufferArrayNonUniformIndexingNative   bool
	ShaderStorageImageArrayNonUniformIndexingNative    bool
	ShaderInputAttachmentArrayNonUniformIndexingNative bool

	RobustBufferAccessUpdateAfterBind bool
	QuadDivergentImplicitLod          bool

	MaxPerStageDescriptorUpdateAfterBindSamplers         int
	MaxPerStageDescriptorUpdateAfterBindUniformBuffers   int
	MaxPerStageDescriptorUpdateAfterBindStorageBuffers   int
	MaxPerStageDescriptorUpdateAfterBindSampledImages    int
	MaxPerStageDescriptorUpdateAfterBindStorageImages    int
	MaxPerStageDescriptorUpdateAfterBindInputAttachments int
	MaxPerStageUpdateAfterBindResources                  int

	MaxDescriptorSetUpdateAfterBindSamplers              int
	MaxDescriptorSetUpdateAfterBindUniformBuffers        int
	MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic int
	MaxDescriptorSetUpdateAfterBindStorageBuffers        int
	MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic int
	MaxDescriptorSetUpdateAfterBindSampledImages         int
	MaxDescriptorSetUpdateAfterBindStorageImages         int
	MaxDescriptorSetUpdateAfterBindInputAttachments      int

	SupportedDepthResolveModes   ResolveModeFlags
	SupportedStencilResolveModes ResolveModeFlags
	IndependentResolveNone       bool
	IndependentResolve           bool

	FilterMinmaxSingleComponentFormats bool
	FilterMinmaxImageComponentMapping  bool

	MaxTimelineSemaphoreValueDifference uint64
	FramebufferIntegerColorSampleCounts core1_0.SampleCounts

	common.NextOutData
}

func (o *PhysicalDeviceVulkan12OutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceVulkan12Properties)
	}

	outData := (*C.VkPhysicalDeviceVulkan12Properties)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_2_PROPERTIES
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceVulkan12OutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
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
	o.FramebufferIntegerColorSampleCounts = core1_0.SampleCounts(outData.framebufferIntegerColorSampleCounts)

	return outData.pNext, nil
}
