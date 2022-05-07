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

type DeviceFeaturesOutData struct {
	Features core1_0.PhysicalDeviceFeatures

	common.HaveNext
}

func (o *DeviceFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceFeatures2{})))
	}

	data := (*C.VkPhysicalDeviceFeatures2)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *DeviceFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDeviceFeatures2)(cDataPointer)

	(&o.Features).PopulateFromCPointer(unsafe.Pointer(&data.features))

	return data.pNext, nil
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
