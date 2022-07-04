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
	common.NextOutData
}

func (o *FormatPropertiesOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
	o.FormatProperties.LinearTilingFeatures = core1_0.FormatFeatures(data.formatProperties.linearTilingFeatures)
	o.FormatProperties.OptimalTilingFeatures = core1_0.FormatFeatures(data.formatProperties.optimalTilingFeatures)
	o.FormatProperties.BufferFeatures = core1_0.FormatFeatures(data.formatProperties.bufferFeatures)

	return data.pNext, nil
}

////

type ImageFormatOptions struct {
	Format core1_0.DataFormat
	Type   core1_0.ImageType
	Tiling core1_0.ImageTiling
	Usage  core1_0.ImageUsages
	Flags  core1_0.ImageCreateFlags

	common.NextOptions
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

////

type ImageFormatPropertiesOutData struct {
	ImageFormatProperties core1_0.ImageFormatProperties

	common.NextOutData
}

func (o *ImageFormatPropertiesOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageFormatProperties2{})))
	}

	data := (*C.VkImageFormatProperties2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *ImageFormatPropertiesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkImageFormatProperties2)(cDataPointer)
	o.ImageFormatProperties.MaxExtent = core1_0.Extent3D{
		Width:  int(data.imageFormatProperties.maxExtent.width),
		Height: int(data.imageFormatProperties.maxExtent.height),
		Depth:  int(data.imageFormatProperties.maxExtent.depth),
	}
	o.ImageFormatProperties.MaxMipLevels = int(data.imageFormatProperties.maxMipLevels)
	o.ImageFormatProperties.MaxArrayLayers = int(data.imageFormatProperties.maxArrayLayers)
	o.ImageFormatProperties.SampleCounts = core1_0.SampleCounts(data.imageFormatProperties.sampleCounts)
	o.ImageFormatProperties.MaxResourceSize = int(data.imageFormatProperties.maxResourceSize)

	return data.pNext, nil
}

////

type MemoryPropertiesOutData struct {
	MemoryProperties core1_0.PhysicalDeviceMemoryProperties

	common.NextOutData
}

func (o *MemoryPropertiesOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
	o.MemoryProperties.MemoryTypes = make([]core1_0.MemoryType, memoryTypeCount)

	for i := 0; i < memoryTypeCount; i++ {
		o.MemoryProperties.MemoryTypes[i].Properties = core1_0.MemoryProperties(data.memoryProperties.memoryTypes[i].propertyFlags)
		o.MemoryProperties.MemoryTypes[i].HeapIndex = int(data.memoryProperties.memoryTypes[i].heapIndex)
	}

	memoryHeapCount := int(data.memoryProperties.memoryHeapCount)
	o.MemoryProperties.MemoryHeaps = make([]core1_0.MemoryHeap, memoryHeapCount)

	for i := 0; i < memoryHeapCount; i++ {
		o.MemoryProperties.MemoryHeaps[i].Size = int(data.memoryProperties.memoryHeaps[i].size)
		o.MemoryProperties.MemoryHeaps[i].Flags = core1_0.MemoryHeapFlags(data.memoryProperties.memoryHeaps[i].flags)
	}

	return data.pNext, nil
}

////

type DevicePropertiesOutData struct {
	Properties core1_0.PhysicalDeviceProperties

	common.NextOutData
}

func (o *DevicePropertiesOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

	common.NextOutData
}

func (o *QueueFamilyOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

	o.QueueFamily.Flags = core1_0.QueueFlags(data.queueFamilyProperties.queueFlags)
	o.QueueFamily.QueueCount = int(data.queueFamilyProperties.queueCount)
	o.QueueFamily.TimestampValidBits = uint32(data.queueFamilyProperties.timestampValidBits)
	o.QueueFamily.MinImageTransferGranularity = core1_0.Extent3D{
		Width:  int(data.queueFamilyProperties.minImageTransferGranularity.width),
		Height: int(data.queueFamilyProperties.minImageTransferGranularity.height),
		Depth:  int(data.queueFamilyProperties.minImageTransferGranularity.depth),
	}

	return data.pNext, nil
}

////

type SparseImageFormatOptions struct {
	Format  core1_0.DataFormat
	Type    core1_0.ImageType
	Samples core1_0.SampleCounts
	Usage   core1_0.ImageUsages
	Tiling  core1_0.ImageTiling

	common.NextOptions
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

////

type SparseImageFormatPropertiesOutData struct {
	SparseImageFormatProperties core1_0.SparseImageFormatProperties
	common.NextOutData
}

func (o *SparseImageFormatPropertiesOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

	o.SparseImageFormatProperties.AspectMask = core1_0.ImageAspectFlags(data.properties.aspectMask)
	o.SparseImageFormatProperties.Flags = core1_0.SparseImageFormatFlags(data.properties.flags)
	o.SparseImageFormatProperties.ImageGranularity = core1_0.Extent3D{
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

	common.NextOutData
}

func (o *PhysicalDeviceIDOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

type PhysicalDeviceMaintenance3OutData struct {
	MaxPerSetDescriptors    int
	MaxMemoryAllocationSize int

	common.NextOutData
}

func (o *PhysicalDeviceMaintenance3OutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMaintenance3Properties{})))
	}

	outData := (*C.VkPhysicalDeviceMaintenance3Properties)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceMaintenance3OutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDeviceMaintenance3Properties)(cDataPointer)

	o.MaxMemoryAllocationSize = int(outData.maxMemoryAllocationSize)
	o.MaxPerSetDescriptors = int(outData.maxPerSetDescriptors)

	return outData.pNext, nil
}

////

type PhysicalDeviceMultiviewOutData struct {
	MaxMultiviewViewCount     int
	MaxMultiviewInstanceIndex int

	common.NextOutData
}

func (o *PhysicalDeviceMultiviewOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMultiviewProperties{})))
	}

	info := (*C.VkPhysicalDeviceMultiviewProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceMultiviewOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceMultiviewProperties)(cDataPointer)
	o.MaxMultiviewViewCount = int(info.maxMultiviewViewCount)
	o.MaxMultiviewInstanceIndex = int(info.maxMultiviewInstanceIndex)

	return info.pNext, nil
}

////

type PhysicalDevicePointClippingOutData struct {
	PointClippingBehavior PointClippingBehavior

	common.NextOutData
}

func (o *PhysicalDevicePointClippingOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevicePointClippingProperties{})))
	}

	properties := (*C.VkPhysicalDevicePointClippingProperties)(preallocatedPointer)
	properties.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES
	properties.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDevicePointClippingOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	properties := (*C.VkPhysicalDevicePointClippingProperties)(cDataPointer)
	o.PointClippingBehavior = PointClippingBehavior(properties.pointClippingBehavior)

	return properties.pNext, nil
}

////

type PhysicalDeviceProtectedMemoryOutData struct {
	ProtectedNoFault bool

	common.NextOutData
}

func (o *PhysicalDeviceProtectedMemoryOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceProtectedMemoryProperties)
	}

	properties := (*C.VkPhysicalDeviceProtectedMemoryProperties)(preallocatedPointer)
	properties.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_PROPERTIES
	properties.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceProtectedMemoryOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	properties := (*C.VkPhysicalDeviceProtectedMemoryProperties)(cDataPointer)

	o.ProtectedNoFault = properties.protectedNoFault != C.VkBool32(0)

	return properties.pNext, nil
}

////

type PhysicalDeviceSubgroupOutData struct {
	SubgroupSize              int
	SupportedStages           core1_0.ShaderStages
	SupportedOperations       SubgroupFeatures
	QuadOperationsInAllStages bool

	common.NextOutData
}

func (o *PhysicalDeviceSubgroupOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceSubgroupProperties)
	}

	properties := (*C.VkPhysicalDeviceSubgroupProperties)(preallocatedPointer)
	properties.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_PROPERTIES
	properties.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceSubgroupOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	properties := (*C.VkPhysicalDeviceSubgroupProperties)(cDataPointer)

	o.SubgroupSize = int(properties.subgroupSize)
	o.SupportedStages = core1_0.ShaderStages(properties.supportedStages)
	o.SupportedOperations = SubgroupFeatures(properties.supportedOperations)
	o.QuadOperationsInAllStages = properties.quadOperationsInAllStages != C.VkBool32(0)

	return properties.pNext, nil
}
