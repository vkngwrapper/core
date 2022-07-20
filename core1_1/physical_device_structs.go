package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

// PointClippingBehavior specifies the point clipping behavior
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPointClippingBehavior.html
type PointClippingBehavior int32

var pointClippingBehaviorMapping = make(map[PointClippingBehavior]string)

func (e PointClippingBehavior) Register(str string) {
	pointClippingBehaviorMapping[e] = str
}

func (e PointClippingBehavior) String() string {
	return pointClippingBehaviorMapping[e]
}

////

// SubgroupFeatureFlags describes what group operations are supported with subgroup scope
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubgroupFeatureFlagBits.html
type SubgroupFeatureFlags int32

var subgroupFeaturesMapping = common.NewFlagStringMapping[SubgroupFeatureFlags]()

func (f SubgroupFeatureFlags) Register(str string) {
	subgroupFeaturesMapping.Register(f, str)
}
func (f SubgroupFeatureFlags) String() string {
	return subgroupFeaturesMapping.FlagsToString(f)
}

////

const (
	// LUIDSize is the length of a locally unique Device identifier
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_LUID_SIZE.html
	LUIDSize int = C.VK_LUID_SIZE
	// MaxGroupSize is the length of a PhysicalDevice handle array
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_MAX_DEVICE_GROUP_SIZE_KHR.html
	MaxGroupSize int = C.VK_MAX_DEVICE_GROUP_SIZE

	// PointClippingAllClipPlanes specifies that the primitive is discarded if the vertex lies
	// outside any clip plane, including the planes bounding the view volume
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPointClippingBehavior.html
	PointClippingAllClipPlanes PointClippingBehavior = C.VK_POINT_CLIPPING_BEHAVIOR_ALL_CLIP_PLANES
	// PointClippingUserClipPlanesOnly specifies that the primitive is discarded only if the vertex
	// lies outside any user clip plane
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPointClippingBehavior.html
	PointClippingUserClipPlanesOnly PointClippingBehavior = C.VK_POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY

	// SubgroupFeatureBasic specifies the Device will accept SPIR-V shader modules containing
	// the GroupNonUniform capability
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubgroupFeatureFlagBits.html
	SubgroupFeatureBasic SubgroupFeatureFlags = C.VK_SUBGROUP_FEATURE_BASIC_BIT
	// SubgroupFeatureVote specifies the Device will accept SPIR-V shader modules containing
	// the GroupNonUniformVote capability
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubgroupFeatureFlagBits.html
	SubgroupFeatureVote SubgroupFeatureFlags = C.VK_SUBGROUP_FEATURE_VOTE_BIT
	// SubgroupFeatureArithmetic specifies the Device will accept SPIR-V shader modules containing
	// the GroupNonUniformArithmetic capability
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubgroupFeatureFlagBits.html
	SubgroupFeatureArithmetic SubgroupFeatureFlags = C.VK_SUBGROUP_FEATURE_ARITHMETIC_BIT
	// SubgroupFeatureBallot specifies the Device will accept SPIR-V shader modules containing
	// the GroupNonUniformBallot capability
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubgroupFeatureFlagBits.html
	SubgroupFeatureBallot SubgroupFeatureFlags = C.VK_SUBGROUP_FEATURE_BALLOT_BIT
	// SubgroupFeatureShuffle specifies the Device will accept SPIR-V shader modules containing
	// the GroupNonUniformShuffle capability
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubgroupFeatureFlagBits.html
	SubgroupFeatureShuffle SubgroupFeatureFlags = C.VK_SUBGROUP_FEATURE_SHUFFLE_BIT
	// SubgroupFeatureShuffleRelative specifies the Device will accept SPIR-V shader modules
	// containing the GroupNonUniformShuffleRelative capability
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubgroupFeatureFlagBits.html
	SubgroupFeatureShuffleRelative SubgroupFeatureFlags = C.VK_SUBGROUP_FEATURE_SHUFFLE_RELATIVE_BIT
	// SubgroupFeatureClustered specifies the Device will accept SPIR-V shader modules containing
	// the GroupNonUniformClustered capability
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubgroupFeatureFlagBits.html
	SubgroupFeatureClustered SubgroupFeatureFlags = C.VK_SUBGROUP_FEATURE_CLUSTERED_BIT
	// SubgroupFeatureQuad specifies the Device will accept SPIR-V shader modules containing
	// the GroupNonUniformQuad capability
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubgroupFeatureFlagBits.html
	SubgroupFeatureQuad SubgroupFeatureFlags = C.VK_SUBGROUP_FEATURE_QUAD_BIT
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

// FormatProperties2 specifies the Image format properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2KHR.html
type FormatProperties2 struct {
	// FormatProperties describes features supported by the requested format
	FormatProperties core1_0.FormatProperties
	common.NextOutData
}

func (o *FormatProperties2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkFormatProperties2{})))
	}

	data := (*C.VkFormatProperties2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *FormatProperties2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkFormatProperties2)(cDataPointer)
	o.FormatProperties.LinearTilingFeatures = core1_0.FormatFeatureFlags(data.formatProperties.linearTilingFeatures)
	o.FormatProperties.OptimalTilingFeatures = core1_0.FormatFeatureFlags(data.formatProperties.optimalTilingFeatures)
	o.FormatProperties.BufferFeatures = core1_0.FormatFeatureFlags(data.formatProperties.bufferFeatures)

	return data.pNext, nil
}

////

// PhysicalDeviceImageFormatInfo2 specifies Image creation parameters
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2KHR.html
type PhysicalDeviceImageFormatInfo2 struct {
	// Format indicates the Image format, corresponding to ImageCreateInfo.Format
	Format core1_0.Format
	// Type indicates the ImageType, corresponding to ImageCreateInfo.ImageType
	Type core1_0.ImageType
	// Tiling indicates the Image tiling, corresponding to ImageCreateInfo.Tiling
	Tiling core1_0.ImageTiling
	// Usage indicates the intended usage of the Image, corresponding to ImageCreateInfo.Usage
	Usage core1_0.ImageUsageFlags
	// Flags indicates additional parameters of the Image, corresponding to ImageCreateInfo.Flags
	Flags core1_0.ImageCreateFlags

	common.NextOptions
}

func (o PhysicalDeviceImageFormatInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

// ImageFormatProperties2 specifies image format properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2KHR.html
type ImageFormatProperties2 struct {
	// ImageFormatProperties is a structure in which capabilities are returned
	ImageFormatProperties core1_0.ImageFormatProperties

	common.NextOutData
}

func (o *ImageFormatProperties2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageFormatProperties2{})))
	}

	data := (*C.VkImageFormatProperties2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *ImageFormatProperties2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkImageFormatProperties2)(cDataPointer)
	o.ImageFormatProperties.MaxExtent = core1_0.Extent3D{
		Width:  int(data.imageFormatProperties.maxExtent.width),
		Height: int(data.imageFormatProperties.maxExtent.height),
		Depth:  int(data.imageFormatProperties.maxExtent.depth),
	}
	o.ImageFormatProperties.MaxMipLevels = int(data.imageFormatProperties.maxMipLevels)
	o.ImageFormatProperties.MaxArrayLayers = int(data.imageFormatProperties.maxArrayLayers)
	o.ImageFormatProperties.SampleCounts = core1_0.SampleCountFlags(data.imageFormatProperties.sampleCounts)
	o.ImageFormatProperties.MaxResourceSize = int(data.imageFormatProperties.maxResourceSize)

	return data.pNext, nil
}

////

// PhysicalDeviceMemoryProperties2 specifies PhysicalDevice memory properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties2.html
type PhysicalDeviceMemoryProperties2 struct {
	// MemoryProperties is a structure which is populated with the same values as in
	// PhysicalDevice.MemoryProperties
	MemoryProperties core1_0.PhysicalDeviceMemoryProperties

	common.NextOutData
}

func (o *PhysicalDeviceMemoryProperties2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMemoryProperties2{})))
	}
	data := (*C.VkPhysicalDeviceMemoryProperties2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceMemoryProperties2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDeviceMemoryProperties2)(cDataPointer)

	memoryTypeCount := int(data.memoryProperties.memoryTypeCount)
	o.MemoryProperties.MemoryTypes = make([]core1_0.MemoryType, memoryTypeCount)

	for i := 0; i < memoryTypeCount; i++ {
		o.MemoryProperties.MemoryTypes[i].PropertyFlags = core1_0.MemoryPropertyFlags(data.memoryProperties.memoryTypes[i].propertyFlags)
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

// PhysicalDeviceProperties2 specifies PhysicalDevice properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html
type PhysicalDeviceProperties2 struct {
	// Properties describes properties of the PhysicalDevice
	Properties core1_0.PhysicalDeviceProperties

	common.NextOutData
}

func (o *PhysicalDeviceProperties2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceProperties2{})))
	}

	data := (*C.VkPhysicalDeviceProperties2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceProperties2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDeviceProperties2)(cDataPointer)

	err = (&o.Properties).PopulateFromCPointer(unsafe.Pointer(&data.properties))
	return data.pNext, err
}

////

// QueueFamilyProperties2 provides information about a Queue family
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties2KHR.html
type QueueFamilyProperties2 struct {
	// QueueFamilyProperties is populated with the same values as in
	// PhysicalDevice.QueueFamilyProperties
	QueueFamilyProperties core1_0.QueueFamilyProperties

	common.NextOutData
}

func (o *QueueFamilyProperties2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkQueueFamilyProperties2{})))
	}

	data := (*C.VkQueueFamilyProperties2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *QueueFamilyProperties2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkQueueFamilyProperties2)(cDataPointer)

	o.QueueFamilyProperties.QueueFlags = core1_0.QueueFlags(data.queueFamilyProperties.queueFlags)
	o.QueueFamilyProperties.QueueCount = int(data.queueFamilyProperties.queueCount)
	o.QueueFamilyProperties.TimestampValidBits = uint32(data.queueFamilyProperties.timestampValidBits)
	o.QueueFamilyProperties.MinImageTransferGranularity = core1_0.Extent3D{
		Width:  int(data.queueFamilyProperties.minImageTransferGranularity.width),
		Height: int(data.queueFamilyProperties.minImageTransferGranularity.height),
		Depth:  int(data.queueFamilyProperties.minImageTransferGranularity.depth),
	}

	return data.pNext, nil
}

////

// PhysicalDeviceSparseImageFormatInfo2 specifies sparse Image format inputs
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSparseImageFormatInfo2KHR.html
type PhysicalDeviceSparseImageFormatInfo2 struct {
	// Format is the Image format
	Format core1_0.Format
	// Type is the dimensionality of the Image
	Type core1_0.ImageType
	// Samples specifies the number of samples per texel
	Samples core1_0.SampleCountFlags
	// Usage describes the intended usage of the Image
	Usage core1_0.ImageUsageFlags
	// Tiling is the tiling arrangement of the texel blocks in memory
	Tiling core1_0.ImageTiling

	common.NextOptions
}

func (o PhysicalDeviceSparseImageFormatInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

// SparseImageFormatProperties2 specifies sparse Image format properties
type SparseImageFormatProperties2 struct {
	// Properties is populated with the same values as in PhysicalDevice.SparseImageFormatProperties
	Properties core1_0.SparseImageFormatProperties
	common.NextOutData
}

func (o *SparseImageFormatProperties2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSparseImageFormatProperties2{})))
	}

	data := (*C.VkSparseImageFormatProperties2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *SparseImageFormatProperties2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkSparseImageFormatProperties2)(cDataPointer)

	o.Properties.AspectMask = core1_0.ImageAspectFlags(data.properties.aspectMask)
	o.Properties.Flags = core1_0.SparseImageFormatFlags(data.properties.flags)
	o.Properties.ImageGranularity = core1_0.Extent3D{
		Width:  int(data.properties.imageGranularity.width),
		Height: int(data.properties.imageGranularity.height),
		Depth:  int(data.properties.imageGranularity.depth),
	}

	return data.pNext, nil
}

////

// PhysicalDeviceIDProperties speicifes IDs related to the PhysicalDevice
type PhysicalDeviceIDProperties struct {
	// DeviceUUID represents a universally-unique identifier for the device
	DeviceUUID uuid.UUID
	// DriverUUID represents a universally-unique identifier for the driver build
	// in use by the device
	DriverUUID uuid.UUID
	// DeviceLUID represents a locally-unique identifier for the device
	DeviceLUID uint64
	// DeviceNodeMask identifies the node within a linked device adapter corresponding to the
	// Device
	DeviceNodeMask uint32
	// DeviceLUIDValid is true if DeviceLUID contains a valid LUID and DeviceNodeMask contains
	// a valid node mask
	DeviceLUIDValid bool

	common.NextOutData
}

func (o *PhysicalDeviceIDProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceIDProperties{})))
	}
	info := (*C.VkPhysicalDeviceIDProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceIDProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
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

// PhysicalDeviceMaintenance3Properties describes DescriptorSet properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance3Properties.html
type PhysicalDeviceMaintenance3Properties struct {
	// MaxPerSetDescriptors is a maximum number of descriptors in a single DescriptorSet that is
	// guaranteed to satisfy any implementation-dependent constraints on the size of a
	// DescriptorSet itself
	MaxPerSetDescriptors int
	// MaxMemoryAllocationSize is the maximum size of a memory allocation that can be created,
	// even if the is more space available in the heap
	MaxMemoryAllocationSize int

	common.NextOutData
}

func (o *PhysicalDeviceMaintenance3Properties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMaintenance3Properties{})))
	}

	outData := (*C.VkPhysicalDeviceMaintenance3Properties)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceMaintenance3Properties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDeviceMaintenance3Properties)(cDataPointer)

	o.MaxMemoryAllocationSize = int(outData.maxMemoryAllocationSize)
	o.MaxPerSetDescriptors = int(outData.maxPerSetDescriptors)

	return outData.pNext, nil
}

////

// PhysicalDeviceMultiviewProperties describes multiview limits that can be supported by an
// implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiviewProperties.html
type PhysicalDeviceMultiviewProperties struct {
	// MaxMultiviewViewCount is one greater than the maximum view index that can be used in
	// a subpass
	MaxMultiviewViewCount int
	// MaxMultiviewInstanceIndex is the maximum
	MaxMultiviewInstanceIndex int

	common.NextOutData
}

func (o *PhysicalDeviceMultiviewProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMultiviewProperties{})))
	}

	info := (*C.VkPhysicalDeviceMultiviewProperties)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceMultiviewProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceMultiviewProperties)(cDataPointer)
	o.MaxMultiviewViewCount = int(info.maxMultiviewViewCount)
	o.MaxMultiviewInstanceIndex = int(info.maxMultiviewInstanceIndex)

	return info.pNext, nil
}

////

// PhysicalDevicePointClippingProperties describes the point clipping behavior supported
// by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePointClippingProperties.html
type PhysicalDevicePointClippingProperties struct {
	// PointClippingBehavior specifies the point clipping behavior supported by the implementation
	PointClippingBehavior PointClippingBehavior

	common.NextOutData
}

func (o *PhysicalDevicePointClippingProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevicePointClippingProperties{})))
	}

	properties := (*C.VkPhysicalDevicePointClippingProperties)(preallocatedPointer)
	properties.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES
	properties.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDevicePointClippingProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	properties := (*C.VkPhysicalDevicePointClippingProperties)(cDataPointer)
	o.PointClippingBehavior = PointClippingBehavior(properties.pointClippingBehavior)

	return properties.pNext, nil
}

////

// PhysicalDeviceProtectedMemoryProperties describes protected memory properties that can
// be supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProtectedMemoryProperties.html
type PhysicalDeviceProtectedMemoryProperties struct {
	// ProtectedNoFault specifies how an implementation behaves when an application attempts
	// to write to unprotected memory in a protected Queue operation, or perform a query in a
	// protected Queue operation
	ProtectedNoFault bool

	common.NextOutData
}

func (o *PhysicalDeviceProtectedMemoryProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceProtectedMemoryProperties)
	}

	properties := (*C.VkPhysicalDeviceProtectedMemoryProperties)(preallocatedPointer)
	properties.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_PROPERTIES
	properties.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceProtectedMemoryProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	properties := (*C.VkPhysicalDeviceProtectedMemoryProperties)(cDataPointer)

	o.ProtectedNoFault = properties.protectedNoFault != C.VkBool32(0)

	return properties.pNext, nil
}

////

// PhysicalDeviceSubgroupProperties describes subgroup support for an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSubgroupProperties.html
type PhysicalDeviceSubgroupProperties struct {
	// SubgroupSize is the default number of invocations in each subgroup
	SubgroupSize int
	// SupportedStages describes the shader stages that group operations with subgroup scope
	// are supported in
	SupportedStages core1_0.ShaderStageFlags
	// SupportedOperations specifies the sets of group operations with subgroup scope supported
	// on this Device
	SupportedOperations SubgroupFeatureFlags
	// QuadOperationsInAllStages specifies whether quad group operations are available in all
	// stages, or are restricted to fragment and compute stages
	QuadOperationsInAllStages bool

	common.NextOutData
}

func (o *PhysicalDeviceSubgroupProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDeviceSubgroupProperties)
	}

	properties := (*C.VkPhysicalDeviceSubgroupProperties)(preallocatedPointer)
	properties.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_PROPERTIES
	properties.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceSubgroupProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	properties := (*C.VkPhysicalDeviceSubgroupProperties)(cDataPointer)

	o.SubgroupSize = int(properties.subgroupSize)
	o.SupportedStages = core1_0.ShaderStageFlags(properties.supportedStages)
	o.SupportedOperations = SubgroupFeatureFlags(properties.supportedOperations)
	o.QuadOperationsInAllStages = properties.quadOperationsInAllStages != C.VkBool32(0)

	return properties.pNext, nil
}
