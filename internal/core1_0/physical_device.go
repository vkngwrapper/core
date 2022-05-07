package internal1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanPhysicalDevice struct {
	InstanceDriver       driver.Driver
	PhysicalDeviceHandle driver.VkPhysicalDevice
	MaximumVersion       common.APIVersion

	PhysicalDevice1_1 core1_1.PhysicalDevice
}

func (d *VulkanPhysicalDevice) Handle() driver.VkPhysicalDevice {
	return d.PhysicalDeviceHandle
}

func (d *VulkanPhysicalDevice) Driver() driver.Driver {
	return d.InstanceDriver
}

func (d *VulkanPhysicalDevice) APIVersion() common.APIVersion {
	return d.MaximumVersion
}

func (d *VulkanPhysicalDevice) Core1_1() core1_1.PhysicalDevice {
	return d.PhysicalDevice1_1
}

func (d *VulkanPhysicalDevice) QueueFamilyProperties() []*core1_0.QueueFamily {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	count := (*driver.Uint32)(allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	d.InstanceDriver.VkGetPhysicalDeviceQueueFamilyProperties(d.PhysicalDeviceHandle, count, nil)

	if *count == 0 {
		return nil
	}

	goCount := int(*count)

	allocatedHandles := allocator.Malloc(goCount * int(unsafe.Sizeof(C.VkQueueFamilyProperties{})))
	familyProperties := ([]C.VkQueueFamilyProperties)(unsafe.Slice((*C.VkQueueFamilyProperties)(allocatedHandles), int(*count)))

	d.InstanceDriver.VkGetPhysicalDeviceQueueFamilyProperties(d.PhysicalDeviceHandle, count, (*driver.VkQueueFamilyProperties)(allocatedHandles))

	var queueFamilies []*core1_0.QueueFamily
	for i := 0; i < goCount; i++ {
		queueFamilies = append(queueFamilies, &core1_0.QueueFamily{
			Flags:              common.QueueFlags(familyProperties[i].queueFlags),
			QueueCount:         int(familyProperties[i].queueCount),
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

func (d *VulkanPhysicalDevice) Properties() (*core1_0.PhysicalDeviceProperties, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	propertiesUnsafe := allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceProperties{})))

	d.InstanceDriver.VkGetPhysicalDeviceProperties(d.PhysicalDeviceHandle, (*driver.VkPhysicalDeviceProperties)(propertiesUnsafe))

	properties := &core1_0.PhysicalDeviceProperties{}
	err := properties.PopulateFromCPointer(propertiesUnsafe)
	return properties, err
}

func (d *VulkanPhysicalDevice) Features() *core1_0.PhysicalDeviceFeatures {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	featuresUnsafe := allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceFeatures{})))

	d.InstanceDriver.VkGetPhysicalDeviceFeatures(d.PhysicalDeviceHandle, (*driver.VkPhysicalDeviceFeatures)(featuresUnsafe))

	features := &core1_0.PhysicalDeviceFeatures{}
	features.PopulateFromCPointer(featuresUnsafe)
	return features
}

func (d *VulkanPhysicalDevice) attemptAvailableExtensions(layerNamePtr *driver.Char) (map[string]*common.ExtensionProperties, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	extensionCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	extensionCount := (*driver.Uint32)(extensionCountPtr)

	res, err := d.InstanceDriver.VkEnumerateDeviceExtensionProperties(d.PhysicalDeviceHandle, layerNamePtr, extensionCount, nil)

	if err != nil || *extensionCount == 0 {
		return nil, res, err
	}

	extensionTotal := int(*extensionCount)
	extensionsPtr := allocator.Malloc(extensionTotal * C.sizeof_struct_VkExtensionProperties)

	res, err = d.InstanceDriver.VkEnumerateDeviceExtensionProperties(d.PhysicalDeviceHandle, layerNamePtr, extensionCount, (*driver.VkExtensionProperties)(extensionsPtr))
	if err != nil || res == core1_0.VKIncomplete {
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

func (d *VulkanPhysicalDevice) AvailableExtensions() (map[string]*common.ExtensionProperties, common.VkResult, error) {
	// There may be a race condition that adds new available extensions between getting the
	// extension count & pulling the extensions, in which case, attemptAvailableExtensions will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*common.ExtensionProperties
	var result common.VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		layers, result, err = d.attemptAvailableExtensions(nil)
	}
	return layers, result, err
}

func (d *VulkanPhysicalDevice) AvailableExtensionsForLayer(layerName string) (map[string]*common.ExtensionProperties, common.VkResult, error) {
	// There may be a race condition that adds new available extensions between getting the
	// extension count & pulling the extensions, in which case, attemptAvailableExtensions will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*common.ExtensionProperties
	var result common.VkResult
	var err error
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	layerNamePtr := (*driver.Char)(allocator.CString(layerName))
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		layers, result, err = d.attemptAvailableExtensions(layerNamePtr)
	}
	return layers, result, err
}

func (d *VulkanPhysicalDevice) attemptAvailableLayers() (map[string]*common.LayerProperties, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	layerCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	layerCount := (*driver.Uint32)(layerCountPtr)

	res, err := d.InstanceDriver.VkEnumerateDeviceLayerProperties(d.PhysicalDeviceHandle, layerCount, nil)

	if err != nil || *layerCount == 0 {
		return nil, res, err
	}

	layerTotal := int(*layerCount)
	layersPtr := allocator.Malloc(layerTotal * C.sizeof_struct_VkLayerProperties)

	res, err = d.InstanceDriver.VkEnumerateDeviceLayerProperties(d.PhysicalDeviceHandle, layerCount, (*driver.VkLayerProperties)(layersPtr))
	if err != nil || res == core1_0.VKIncomplete {
		return nil, res, err
	}

	retVal := make(map[string]*common.LayerProperties)
	layerSlice := ([]C.VkLayerProperties)(unsafe.Slice((*C.VkLayerProperties)(layersPtr), layerTotal))

	for i := 0; i < layerTotal; i++ {
		layer := layerSlice[i]

		outLayer := &common.LayerProperties{
			LayerName:             C.GoString((*C.char)(&layer.layerName[0])),
			Description:           C.GoString((*C.char)(&layer.description[0])),
			SpecVersion:           common.Version(layer.specVersion),
			ImplementationVersion: common.Version(layer.implementationVersion),
		}

		existingLayer, ok := retVal[outLayer.LayerName]
		if ok && existingLayer.SpecVersion >= outLayer.SpecVersion {
			continue
		} else if ok && existingLayer.SpecVersion == outLayer.SpecVersion && existingLayer.ImplementationVersion >= outLayer.ImplementationVersion {
			continue
		}
		retVal[outLayer.LayerName] = outLayer
	}

	return retVal, res, nil
}

func (d *VulkanPhysicalDevice) AvailableLayers() (map[string]*common.LayerProperties, common.VkResult, error) {
	// There may be a race condition that adds new available extensions between getting the
	// extension count & pulling the extensions, in which case, attemptAvailableExtensions will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*common.LayerProperties
	var result common.VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		layers, result, err = d.attemptAvailableLayers()
	}
	return layers, result, err
}

func (d *VulkanPhysicalDevice) FormatProperties(format common.DataFormat) *core1_0.FormatProperties {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	properties := (*C.VkFormatProperties)(allocator.Malloc(C.sizeof_struct_VkFormatProperties))

	d.InstanceDriver.VkGetPhysicalDeviceFormatProperties(d.PhysicalDeviceHandle, driver.VkFormat(format), (*driver.VkFormatProperties)(unsafe.Pointer(properties)))

	return &core1_0.FormatProperties{
		LinearTilingFeatures:  common.FormatFeatures(properties.linearTilingFeatures),
		OptimalTilingFeatures: common.FormatFeatures(properties.optimalTilingFeatures),
		BufferFeatures:        common.FormatFeatures(properties.bufferFeatures),
	}
}

func (d *VulkanPhysicalDevice) ImageFormatProperties(format common.DataFormat, imageType common.ImageType, tiling common.ImageTiling, usages common.ImageUsages, flags common.ImageCreateFlags) (*core1_0.ImageFormatProperties, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	properties := (*C.VkImageFormatProperties)(allocator.Malloc(C.sizeof_struct_VkImageFormatProperties))

	res, err := d.InstanceDriver.VkGetPhysicalDeviceImageFormatProperties(d.PhysicalDeviceHandle, driver.VkFormat(format), driver.VkImageType(imageType), driver.VkImageTiling(tiling), driver.VkImageUsageFlags(usages), driver.VkImageCreateFlags(flags), (*driver.VkImageFormatProperties)(unsafe.Pointer(properties)))
	if err != nil {
		return nil, res, err
	}

	return &core1_0.ImageFormatProperties{
		MaxExtent: common.Extent3D{
			Width:  int(properties.maxExtent.width),
			Height: int(properties.maxExtent.height),
			Depth:  int(properties.maxExtent.depth),
		},
		MaxMipLevels:    int(properties.maxMipLevels),
		MaxArrayLayers:  int(properties.maxArrayLayers),
		SampleCounts:    common.SampleCounts(properties.sampleCounts),
		MaxResourceSize: int(properties.maxResourceSize),
	}, res, nil
}

func (d *VulkanPhysicalDevice) SparseImageFormatProperties(format common.DataFormat, imageType common.ImageType, samples common.SampleCounts, usages common.ImageUsages, tiling common.ImageTiling) []core1_0.SparseImageFormatProperties {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	propertiesCount := (*C.uint32_t)(arena.Malloc(4))

	d.InstanceDriver.VkGetPhysicalDeviceSparseImageFormatProperties(d.PhysicalDeviceHandle, driver.VkFormat(format), driver.VkImageType(imageType), driver.VkSampleCountFlagBits(samples), driver.VkImageUsageFlags(usages), driver.VkImageTiling(tiling), (*driver.Uint32)(propertiesCount), nil)

	if *propertiesCount == 0 {
		return nil
	}

	propertiesPtr := (*C.VkSparseImageFormatProperties)(arena.Malloc(int(*propertiesCount) * C.sizeof_struct_VkSparseImageFormatProperties))

	d.InstanceDriver.VkGetPhysicalDeviceSparseImageFormatProperties(d.PhysicalDeviceHandle, driver.VkFormat(format), driver.VkImageType(imageType), driver.VkSampleCountFlagBits(samples), driver.VkImageUsageFlags(usages), driver.VkImageTiling(tiling), (*driver.Uint32)(unsafe.Pointer(propertiesCount)), (*driver.VkSparseImageFormatProperties)(unsafe.Pointer(propertiesPtr)))

	propertiesSlice := ([]C.VkSparseImageFormatProperties)(unsafe.Slice(propertiesPtr, int(*propertiesCount)))

	outReqs := make([]core1_0.SparseImageFormatProperties, *propertiesCount)
	for j := 0; j < int(*propertiesCount); j++ {
		inProps := propertiesSlice[j]
		outReqs[j].Flags = common.SparseImageFormatFlags(inProps.flags)
		outReqs[j].ImageGranularity = common.Extent3D{
			Width:  int(inProps.imageGranularity.width),
			Height: int(inProps.imageGranularity.height),
			Depth:  int(inProps.imageGranularity.depth),
		}
		outReqs[j].AspectMask = common.ImageAspectFlags(inProps.aspectMask)
	}

	return outReqs
}
