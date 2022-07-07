package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanPhysicalDevice struct {
	instanceDriver       driver.Driver
	physicalDeviceHandle driver.VkPhysicalDevice

	instanceVersion      common.APIVersion
	maximumDeviceVersion common.APIVersion
}

func (d *VulkanPhysicalDevice) Handle() driver.VkPhysicalDevice {
	return d.physicalDeviceHandle
}

func (d *VulkanPhysicalDevice) Driver() driver.Driver {
	return d.instanceDriver
}

func (d *VulkanPhysicalDevice) DeviceAPIVersion() common.APIVersion {
	return d.maximumDeviceVersion
}

func (d *VulkanPhysicalDevice) InstanceAPIVersion() common.APIVersion {
	return d.instanceVersion
}

func (d *VulkanPhysicalDevice) QueueFamilyProperties() []*QueueFamily {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	count := (*driver.Uint32)(allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	d.instanceDriver.VkGetPhysicalDeviceQueueFamilyProperties(d.physicalDeviceHandle, count, nil)

	if *count == 0 {
		return nil
	}

	goCount := int(*count)

	allocatedHandles := allocator.Malloc(goCount * int(unsafe.Sizeof(C.VkQueueFamilyProperties{})))
	familyProperties := ([]C.VkQueueFamilyProperties)(unsafe.Slice((*C.VkQueueFamilyProperties)(allocatedHandles), int(*count)))

	d.instanceDriver.VkGetPhysicalDeviceQueueFamilyProperties(d.physicalDeviceHandle, count, (*driver.VkQueueFamilyProperties)(allocatedHandles))

	var queueFamilies []*QueueFamily
	for i := 0; i < goCount; i++ {
		queueFamilies = append(queueFamilies, &QueueFamily{
			QueueFlags:         QueueFlags(familyProperties[i].queueFlags),
			QueueCount:         int(familyProperties[i].queueCount),
			TimestampValidBits: uint32(familyProperties[i].timestampValidBits),
			MinImageTransferGranularity: Extent3D{
				Width:  int(familyProperties[i].minImageTransferGranularity.width),
				Height: int(familyProperties[i].minImageTransferGranularity.height),
				Depth:  int(familyProperties[i].minImageTransferGranularity.depth),
			},
		})
	}

	return queueFamilies
}

func (d *VulkanPhysicalDevice) Properties() (*PhysicalDeviceProperties, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	propertiesUnsafe := allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceProperties{})))

	d.instanceDriver.VkGetPhysicalDeviceProperties(d.physicalDeviceHandle, (*driver.VkPhysicalDeviceProperties)(propertiesUnsafe))

	properties := &PhysicalDeviceProperties{}
	err := properties.PopulateFromCPointer(propertiesUnsafe)
	return properties, err
}

func (d *VulkanPhysicalDevice) Features() *PhysicalDeviceFeatures {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	featuresUnsafe := allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceFeatures{})))

	d.instanceDriver.VkGetPhysicalDeviceFeatures(d.physicalDeviceHandle, (*driver.VkPhysicalDeviceFeatures)(featuresUnsafe))

	features := &PhysicalDeviceFeatures{}
	features.PopulateFromCPointer(featuresUnsafe)
	return features
}

func (d *VulkanPhysicalDevice) attemptAvailableExtensions(layerNamePtr *driver.Char) (map[string]*ExtensionProperties, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	extensionCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	extensionCount := (*driver.Uint32)(extensionCountPtr)

	res, err := d.instanceDriver.VkEnumerateDeviceExtensionProperties(d.physicalDeviceHandle, layerNamePtr, extensionCount, nil)

	if err != nil || *extensionCount == 0 {
		return nil, res, err
	}

	extensionTotal := int(*extensionCount)
	extensionsPtr := allocator.Malloc(extensionTotal * C.sizeof_struct_VkExtensionProperties)

	res, err = d.instanceDriver.VkEnumerateDeviceExtensionProperties(d.physicalDeviceHandle, layerNamePtr, extensionCount, (*driver.VkExtensionProperties)(extensionsPtr))
	if err != nil || res == VKIncomplete {
		return nil, res, err
	}

	retVal := make(map[string]*ExtensionProperties)
	extensionSlice := ([]C.VkExtensionProperties)(unsafe.Slice((*C.VkExtensionProperties)(extensionsPtr), extensionTotal))

	for i := 0; i < extensionTotal; i++ {
		extension := extensionSlice[i]

		outExtension := &ExtensionProperties{
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

func (d *VulkanPhysicalDevice) EnumerateDeviceExtensionProperties() (map[string]*ExtensionProperties, common.VkResult, error) {
	// There may be a race condition that adds new available extensions between getting the
	// extension count & pulling the extensions, in which case, attemptAvailableExtensions will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*ExtensionProperties
	var result common.VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == VKIncomplete) {
		layers, result, err = d.attemptAvailableExtensions(nil)
	}
	return layers, result, err
}

func (d *VulkanPhysicalDevice) EnumerateDeviceExtensionPropertiesForLayer(layerName string) (map[string]*ExtensionProperties, common.VkResult, error) {
	// There may be a race condition that adds new available extensions between getting the
	// extension count & pulling the extensions, in which case, attemptAvailableExtensions will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*ExtensionProperties
	var result common.VkResult
	var err error
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	layerNamePtr := (*driver.Char)(allocator.CString(layerName))
	for doWhile := true; doWhile; doWhile = (result == VKIncomplete) {
		layers, result, err = d.attemptAvailableExtensions(layerNamePtr)
	}
	return layers, result, err
}

func (d *VulkanPhysicalDevice) attemptAvailableLayers() (map[string]*LayerProperties, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	layerCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	layerCount := (*driver.Uint32)(layerCountPtr)

	res, err := d.instanceDriver.VkEnumerateDeviceLayerProperties(d.physicalDeviceHandle, layerCount, nil)

	if err != nil || *layerCount == 0 {
		return nil, res, err
	}

	layerTotal := int(*layerCount)
	layersPtr := allocator.Malloc(layerTotal * C.sizeof_struct_VkLayerProperties)

	res, err = d.instanceDriver.VkEnumerateDeviceLayerProperties(d.physicalDeviceHandle, layerCount, (*driver.VkLayerProperties)(layersPtr))
	if err != nil || res == VKIncomplete {
		return nil, res, err
	}

	retVal := make(map[string]*LayerProperties)
	layerSlice := ([]C.VkLayerProperties)(unsafe.Slice((*C.VkLayerProperties)(layersPtr), layerTotal))

	for i := 0; i < layerTotal; i++ {
		layer := layerSlice[i]

		outLayer := &LayerProperties{
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

func (d *VulkanPhysicalDevice) EnumerateDeviceLayerProperties() (map[string]*LayerProperties, common.VkResult, error) {
	// There may be a race condition that adds new available extensions between getting the
	// extension count & pulling the extensions, in which case, attemptAvailableExtensions will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*LayerProperties
	var result common.VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == VKIncomplete) {
		layers, result, err = d.attemptAvailableLayers()
	}
	return layers, result, err
}

func (d *VulkanPhysicalDevice) FormatProperties(format Format) *FormatProperties {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	properties := (*C.VkFormatProperties)(allocator.Malloc(C.sizeof_struct_VkFormatProperties))

	d.instanceDriver.VkGetPhysicalDeviceFormatProperties(d.physicalDeviceHandle, driver.VkFormat(format), (*driver.VkFormatProperties)(unsafe.Pointer(properties)))

	return &FormatProperties{
		LinearTilingFeatures:  FormatFeatureFlags(properties.linearTilingFeatures),
		OptimalTilingFeatures: FormatFeatureFlags(properties.optimalTilingFeatures),
		BufferFeatures:        FormatFeatureFlags(properties.bufferFeatures),
	}
}

func (d *VulkanPhysicalDevice) ImageFormatProperties(format Format, imageType ImageType, tiling ImageTiling, usages ImageUsageFlags, flags ImageCreateFlags) (*ImageFormatProperties, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	properties := (*C.VkImageFormatProperties)(allocator.Malloc(C.sizeof_struct_VkImageFormatProperties))

	res, err := d.instanceDriver.VkGetPhysicalDeviceImageFormatProperties(d.physicalDeviceHandle, driver.VkFormat(format), driver.VkImageType(imageType), driver.VkImageTiling(tiling), driver.VkImageUsageFlags(usages), driver.VkImageCreateFlags(flags), (*driver.VkImageFormatProperties)(unsafe.Pointer(properties)))
	if err != nil {
		return nil, res, err
	}

	return &ImageFormatProperties{
		MaxExtent: Extent3D{
			Width:  int(properties.maxExtent.width),
			Height: int(properties.maxExtent.height),
			Depth:  int(properties.maxExtent.depth),
		},
		MaxMipLevels:    int(properties.maxMipLevels),
		MaxArrayLayers:  int(properties.maxArrayLayers),
		SampleCounts:    SampleCountFlags(properties.sampleCounts),
		MaxResourceSize: int(properties.maxResourceSize),
	}, res, nil
}

func (d *VulkanPhysicalDevice) SparseImageFormatProperties(format Format, imageType ImageType, samples SampleCountFlags, usages ImageUsageFlags, tiling ImageTiling) []SparseImageFormatProperties {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	propertiesCount := (*C.uint32_t)(arena.Malloc(4))

	d.instanceDriver.VkGetPhysicalDeviceSparseImageFormatProperties(d.physicalDeviceHandle, driver.VkFormat(format), driver.VkImageType(imageType), driver.VkSampleCountFlagBits(samples), driver.VkImageUsageFlags(usages), driver.VkImageTiling(tiling), (*driver.Uint32)(propertiesCount), nil)

	if *propertiesCount == 0 {
		return nil
	}

	propertiesPtr := (*C.VkSparseImageFormatProperties)(arena.Malloc(int(*propertiesCount) * C.sizeof_struct_VkSparseImageFormatProperties))

	d.instanceDriver.VkGetPhysicalDeviceSparseImageFormatProperties(d.physicalDeviceHandle, driver.VkFormat(format), driver.VkImageType(imageType), driver.VkSampleCountFlagBits(samples), driver.VkImageUsageFlags(usages), driver.VkImageTiling(tiling), (*driver.Uint32)(unsafe.Pointer(propertiesCount)), (*driver.VkSparseImageFormatProperties)(unsafe.Pointer(propertiesPtr)))

	propertiesSlice := ([]C.VkSparseImageFormatProperties)(unsafe.Slice(propertiesPtr, int(*propertiesCount)))

	outReqs := make([]SparseImageFormatProperties, *propertiesCount)
	for j := 0; j < int(*propertiesCount); j++ {
		inProps := propertiesSlice[j]
		outReqs[j].Flags = SparseImageFormatFlags(inProps.flags)
		outReqs[j].ImageGranularity = Extent3D{
			Width:  int(inProps.imageGranularity.width),
			Height: int(inProps.imageGranularity.height),
			Depth:  int(inProps.imageGranularity.depth),
		}
		outReqs[j].AspectMask = ImageAspectFlags(inProps.aspectMask)
	}

	return outReqs
}

func (d *VulkanPhysicalDevice) CreateDevice(allocationCallbacks *driver.AllocationCallbacks, options DeviceCreateInfo) (Device, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var deviceHandle driver.VkDevice
	res, err := d.instanceDriver.VkCreateDevice(d.physicalDeviceHandle, (*driver.VkDeviceCreateInfo)(createInfo), allocationCallbacks.Handle(), &deviceHandle)
	if err != nil {
		return nil, res, err
	}

	deviceDriver, err := d.instanceDriver.CreateDeviceDriver(deviceHandle)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	device := createDeviceObject(deviceDriver, deviceHandle, d.maximumDeviceVersion)

	for _, extension := range options.EnabledExtensionNames {
		device.activeDeviceExtensions[extension] = struct{}{}
	}

	return device, res, nil
}
