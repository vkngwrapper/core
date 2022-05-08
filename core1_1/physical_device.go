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
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"unsafe"
)

type PointClippingBehavior int32

var pointClippingBehaviorMapping = make(map[PointClippingBehavior]string)

func (e PointClippingBehavior) Register(str string) {
	pointClippingBehaviorMapping[e] = str
}

func (e PointClippingBehavior) String() string {
	return pointClippingBehaviorMapping[e]
}

////

type SubgroupFeatures int32

var subgroupFeaturesMapping = common.NewFlagStringMapping[SubgroupFeatures]()

func (f SubgroupFeatures) Register(str string) {
	subgroupFeaturesMapping.Register(f, str)
}
func (f SubgroupFeatures) String() string {
	return subgroupFeaturesMapping.FlagsToString(f)
}

////

const (
	LUIDSize     int = C.VK_LUID_SIZE
	MaxGroupSize int = C.VK_MAX_DEVICE_GROUP_SIZE

	PointClippingAllClipPlanes      PointClippingBehavior = C.VK_POINT_CLIPPING_BEHAVIOR_ALL_CLIP_PLANES
	PointClippingUserClipPlanesOnly PointClippingBehavior = C.VK_POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY

	SubgroupFeatureBasic           SubgroupFeatures = C.VK_SUBGROUP_FEATURE_BASIC_BIT
	SubgroupFeatureVote            SubgroupFeatures = C.VK_SUBGROUP_FEATURE_VOTE_BIT
	SubgroupFeatureArithmetic      SubgroupFeatures = C.VK_SUBGROUP_FEATURE_ARITHMETIC_BIT
	SubgroupFeatureBallot          SubgroupFeatures = C.VK_SUBGROUP_FEATURE_BALLOT_BIT
	SubgroupFeatureShuffle         SubgroupFeatures = C.VK_SUBGROUP_FEATURE_SHUFFLE_BIT
	SubgroupFeatureShuffleRelative SubgroupFeatures = C.VK_SUBGROUP_FEATURE_SHUFFLE_RELATIVE_BIT
	SubgroupFeatureClustered       SubgroupFeatures = C.VK_SUBGROUP_FEATURE_CLUSTERED_BIT
	SubgroupFeatureQuad            SubgroupFeatures = C.VK_SUBGROUP_FEATURE_QUAD_BIT
)

func init() {
	PointClippingAllClipPlanes.Register("All Clip Planes")
	PointClippingUserClipPlanesOnly.Register("User Clip Planes Only")

	SubgroupFeatureBasic.Register("Basic")
	SubgroupFeatureVote.Register("Vote")
	SubgroupFeatureArithmetic.Register("Arithmetic")
	SubgroupFeatureBallot.Register("Ballot")
	SubgroupFeatureShuffle.Register("Shuffle")
	SubgroupFeatureShuffleRelative.Register("Shuffle (Relative)")
	SubgroupFeatureClustered.Register("Clustered")
	SubgroupFeatureQuad.Register("Quad")
}

////

type FormatPropertiesOutData struct {
	FormatProperties core1_0.FormatProperties
	common.HaveNext
}

func (o *FormatPropertiesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkFormatProperties2{})))
	}

	data := (*C.VkFormatProperties2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *FormatPropertiesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkFormatProperties2)(cDataPointer)
	o.FormatProperties.LinearTilingFeatures = common.FormatFeatures(data.formatProperties.linearTilingFeatures)
	o.FormatProperties.OptimalTilingFeatures = common.FormatFeatures(data.formatProperties.optimalTilingFeatures)
	o.FormatProperties.BufferFeatures = common.FormatFeatures(data.formatProperties.bufferFeatures)

	return data.pNext, nil
}

////

type ImageFormatOptions struct {
	Format common.DataFormat
	Type   common.ImageType
	Tiling common.ImageTiling
	Usage  common.ImageUsages
	Flags  common.ImageCreateFlags

	common.HaveNext
}

func (o ImageFormatOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceImageFormatInfo2{})))
	}
	info := (*C.VkPhysicalDeviceImageFormatInfo2)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2
	info.pNext = next
	info.format = C.VkFormat(o.Format)
	info._type = C.VkImageType(o.Type)
	info.tiling = C.VkImageTiling(o.Tiling)
	info.usage = C.VkImageUsageFlags(o.Usage)
	info.flags = C.VkImageCreateFlags(o.Flags)

	return preallocatedPointer, nil
}

func (o ImageFormatOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceImageFormatInfo2)(cDataPointer)
	return info.pNext, nil
}

////

type ImageFormatOutData struct {
	ImageFormatProperties core1_0.ImageFormatProperties

	common.HaveNext
}

func (o *ImageFormatOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageFormatProperties2{})))
	}

	data := (*C.VkImageFormatProperties2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *ImageFormatOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkImageFormatProperties2)(cDataPointer)
	o.ImageFormatProperties.MaxExtent = common.Extent3D{
		Width:  int(data.imageFormatProperties.maxExtent.width),
		Height: int(data.imageFormatProperties.maxExtent.height),
		Depth:  int(data.imageFormatProperties.maxExtent.depth),
	}
	o.ImageFormatProperties.MaxMipLevels = int(data.imageFormatProperties.maxMipLevels)
	o.ImageFormatProperties.MaxArrayLayers = int(data.imageFormatProperties.maxArrayLayers)
	o.ImageFormatProperties.SampleCounts = common.SampleCounts(data.imageFormatProperties.sampleCounts)
	o.ImageFormatProperties.MaxResourceSize = int(data.imageFormatProperties.maxResourceSize)

	return data.pNext, nil
}

////

type MemoryPropertiesOutData struct {
	MemoryProperties core1_0.PhysicalDeviceMemoryProperties

	common.HaveNext
}

func (o *MemoryPropertiesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMemoryProperties2{})))
	}
	data := (*C.VkPhysicalDeviceMemoryProperties2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *MemoryPropertiesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDeviceMemoryProperties2)(cDataPointer)

	memoryTypeCount := int(data.memoryProperties.memoryTypeCount)
	o.MemoryProperties.MemoryTypes = make([]common.MemoryType, memoryTypeCount)

	for i := 0; i < memoryTypeCount; i++ {
		o.MemoryProperties.MemoryTypes[i].Properties = common.MemoryProperties(data.memoryProperties.memoryTypes[i].propertyFlags)
		o.MemoryProperties.MemoryTypes[i].HeapIndex = int(data.memoryProperties.memoryTypes[i].heapIndex)
	}

	memoryHeapCount := int(data.memoryProperties.memoryHeapCount)
	o.MemoryProperties.MemoryHeaps = make([]common.MemoryHeap, memoryHeapCount)

	for i := 0; i < memoryHeapCount; i++ {
		o.MemoryProperties.MemoryHeaps[i].Size = int(data.memoryProperties.memoryHeaps[i].size)
		o.MemoryProperties.MemoryHeaps[i].Flags = common.MemoryHeapFlags(data.memoryProperties.memoryHeaps[i].flags)
	}

	return data.pNext, nil
}

////

type DevicePropertiesOutData struct {
	Properties core1_0.PhysicalDeviceProperties

	common.HaveNext
}

func (o *DevicePropertiesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceProperties2{})))
	}

	data := (*C.VkPhysicalDeviceProperties2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *DevicePropertiesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDeviceProperties2)(cDataPointer)

	err = (&o.Properties).PopulateFromCPointer(unsafe.Pointer(&data.properties))
	return data.pNext, err
}

////

type QueueFamilyOutData struct {
	QueueFamily core1_0.QueueFamily

	common.HaveNext
}

func (o *QueueFamilyOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkQueueFamilyProperties2{})))
	}

	data := (*C.VkQueueFamilyProperties2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *QueueFamilyOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkQueueFamilyProperties2)(cDataPointer)

	o.QueueFamily.Flags = common.QueueFlags(data.queueFamilyProperties.queueFlags)
	o.QueueFamily.QueueCount = int(data.queueFamilyProperties.queueCount)
	o.QueueFamily.TimestampValidBits = uint32(data.queueFamilyProperties.timestampValidBits)
	o.QueueFamily.MinImageTransferGranularity = common.Extent3D{
		Width:  int(data.queueFamilyProperties.minImageTransferGranularity.width),
		Height: int(data.queueFamilyProperties.minImageTransferGranularity.height),
		Depth:  int(data.queueFamilyProperties.minImageTransferGranularity.depth),
	}

	return data.pNext, nil
}

////

type SparseImageFormatOptions struct {
	Format  common.DataFormat
	Type    common.ImageType
	Samples common.SampleCounts
	Usage   common.ImageUsages
	Tiling  common.ImageTiling

	common.HaveNext
}

func (o SparseImageFormatOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceSparseImageFormatInfo2{})))
	}

	createInfo := (*C.VkPhysicalDeviceSparseImageFormatInfo2)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2
	createInfo.pNext = next
	createInfo.format = C.VkFormat(o.Format)
	createInfo._type = C.VkImageType(o.Type)
	createInfo.samples = C.VkSampleCountFlagBits(o.Samples)
	createInfo.usage = C.VkImageUsageFlags(o.Usage)
	createInfo.tiling = C.VkImageTiling(o.Tiling)

	return preallocatedPointer, nil
}

func (o SparseImageFormatOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkPhysicalDeviceSparseImageFormatInfo2)(cDataPointer)

	return unsafe.Pointer(createInfo.pNext), nil
}

////

type SparseImageFormatPropertiesOutData struct {
	SparseImageFormatProperties core1_0.SparseImageFormatProperties
	common.HaveNext
}

func (o *SparseImageFormatPropertiesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSparseImageFormatProperties2{})))
	}

	data := (*C.VkSparseImageFormatProperties2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *SparseImageFormatPropertiesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkSparseImageFormatProperties2)(cDataPointer)

	o.SparseImageFormatProperties.AspectMask = common.ImageAspectFlags(data.properties.aspectMask)
	o.SparseImageFormatProperties.Flags = common.SparseImageFormatFlags(data.properties.flags)
	o.SparseImageFormatProperties.ImageGranularity = common.Extent3D{
		Width:  int(data.properties.imageGranularity.width),
		Height: int(data.properties.imageGranularity.height),
		Depth:  int(data.properties.imageGranularity.depth),
	}

	return data.pNext, nil
}

////

type PhysicalDeviceIDOutData struct {
	DeviceUUID      uuid.UUID
	DriverUUID      uuid.UUID
	DeviceLUID      uint64
	DeviceNodeMask  uint32
	DeviceLUIDValid bool

	common.HaveNext
}

func (o *PhysicalDeviceIDOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceIDProperties{})))
	}
	info := (*C.VkPhysicalDeviceIDProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceIDOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceIDProperties)(cDataPointer)

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

	return info.pNext, nil
}

////

type Maintenance3OutData struct {
	MaxPerSetDescriptors    int
	MaxMemoryAllocationSize int

	common.HaveNext
}

func (o *Maintenance3OutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMaintenance3Properties{})))
	}

	outData := (*C.VkPhysicalDeviceMaintenance3Properties)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *Maintenance3OutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDeviceMaintenance3Properties)(cDataPointer)

	o.MaxMemoryAllocationSize = int(outData.maxMemoryAllocationSize)
	o.MaxPerSetDescriptors = int(outData.maxPerSetDescriptors)

	return outData.pNext, nil
}

////

type MultiviewPropertiesOutData struct {
	MaxMultiviewViewCount     int
	MaxMultiviewInstanceIndex int

	common.HaveNext
}

func (o *MultiviewPropertiesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMultiviewProperties{})))
	}

	info := (*C.VkPhysicalDeviceMultiviewProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *MultiviewPropertiesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceMultiviewProperties)(cDataPointer)
	o.MaxMultiviewViewCount = int(info.maxMultiviewViewCount)
	o.MaxMultiviewInstanceIndex = int(info.maxMultiviewInstanceIndex)

	return info.pNext, nil
}

////

type PointClippingOutData struct {
	PointClippingBehavior PointClippingBehavior

	common.HaveNext
}

func (o *PointClippingOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevicePointClippingProperties{})))
	}

	properties := (*C.VkPhysicalDevicePointClippingProperties)(preallocatedPointer)
	properties.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES
	properties.pNext = next

	return preallocatedPointer, nil
}

func (o *PointClippingOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	properties := (*C.VkPhysicalDevicePointClippingProperties)(cDataPointer)
	o.PointClippingBehavior = PointClippingBehavior(properties.pointClippingBehavior)

	return properties.pNext, nil
}

////

type ProtectedMemoryOutData struct {
	ProtectedNoFault bool

	common.HaveNext
}

func (o *ProtectedMemoryOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceProtectedMemoryProperties)
	}

	properties := (*C.VkPhysicalDeviceProtectedMemoryProperties)(preallocatedPointer)
	properties.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_PROPERTIES
	properties.pNext = next

	return preallocatedPointer, nil
}

func (o *ProtectedMemoryOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	properties := (*C.VkPhysicalDeviceProtectedMemoryProperties)(cDataPointer)

	o.ProtectedNoFault = properties.protectedNoFault != C.VkBool32(0)

	return properties.pNext, nil
}

////

type SubgroupOutData struct {
	SubgroupSize              int
	SupportedStages           common.ShaderStages
	SupportedOperations       SubgroupFeatures
	QuadOperationsInAllStages bool

	common.HaveNext
}

func (o *SubgroupOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceSubgroupProperties)
	}

	properties := (*C.VkPhysicalDeviceSubgroupProperties)(preallocatedPointer)
	properties.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_PROPERTIES
	properties.pNext = next

	return preallocatedPointer, nil
}

func (o *SubgroupOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	properties := (*C.VkPhysicalDeviceSubgroupProperties)(cDataPointer)

	o.SubgroupSize = int(properties.subgroupSize)
	o.SupportedStages = common.ShaderStages(properties.supportedStages)
	o.SupportedOperations = SubgroupFeatures(properties.supportedOperations)
	o.QuadOperationsInAllStages = properties.quadOperationsInAllStages != C.VkBool32(0)

	return properties.pNext, nil
}
