package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *InstanceVulkanDriver) GetPhysicalDeviceQueueFamilyProperties(physicalDevice core.PhysicalDevice) []*core1_0.QueueFamilyProperties {
	if physicalDevice.Handle() == 0 {
		panic("physicalDevice was uninitialized")
	}

	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	count := (*loader.Uint32)(allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	v.LoaderObj.VkGetPhysicalDeviceQueueFamilyProperties(physicalDevice.Handle(), count, nil)

	if *count == 0 {
		return nil
	}

	goCount := int(*count)

	allocatedHandles := allocator.Malloc(goCount * int(unsafe.Sizeof(C.VkQueueFamilyProperties{})))
	familyProperties := ([]C.VkQueueFamilyProperties)(unsafe.Slice((*C.VkQueueFamilyProperties)(allocatedHandles), int(*count)))

	v.LoaderObj.VkGetPhysicalDeviceQueueFamilyProperties(physicalDevice.Handle(), count, (*loader.VkQueueFamilyProperties)(allocatedHandles))

	var queueFamilies []*core1_0.QueueFamilyProperties
	for i := 0; i < goCount; i++ {
		queueFamilies = append(queueFamilies, &core1_0.QueueFamilyProperties{
			QueueFlags:         core1_0.QueueFlags(familyProperties[i].queueFlags),
			QueueCount:         int(familyProperties[i].queueCount),
			TimestampValidBits: uint32(familyProperties[i].timestampValidBits),
			MinImageTransferGranularity: core1_0.Extent3D{
				Width:  int(familyProperties[i].minImageTransferGranularity.width),
				Height: int(familyProperties[i].minImageTransferGranularity.height),
				Depth:  int(familyProperties[i].minImageTransferGranularity.depth),
			},
		})
	}

	return queueFamilies
}

func (v *InstanceVulkanDriver) GetPhysicalDeviceProperties(physicalDevice core.PhysicalDevice) (*core1_0.PhysicalDeviceProperties, error) {
	if physicalDevice.Handle() == 0 {
		return nil, fmt.Errorf("physicalDevice was uninitialized")
	}
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	propertiesUnsafe := allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceProperties{})))

	v.LoaderObj.VkGetPhysicalDeviceProperties(physicalDevice.Handle(), (*loader.VkPhysicalDeviceProperties)(propertiesUnsafe))

	properties := &core1_0.PhysicalDeviceProperties{}
	err := properties.PopulateFromCPointer(propertiesUnsafe)
	return properties, err
}

func (v *InstanceVulkanDriver) GetPhysicalDeviceFeatures(physicalDevice core.PhysicalDevice) *core1_0.PhysicalDeviceFeatures {
	if physicalDevice.Handle() == 0 {
		panic("physicalDevice was uninitialized")
	}

	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	featuresUnsafe := allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceFeatures{})))

	v.LoaderObj.VkGetPhysicalDeviceFeatures(physicalDevice.Handle(), (*loader.VkPhysicalDeviceFeatures)(featuresUnsafe))

	features := &core1_0.PhysicalDeviceFeatures{}
	features.PopulateFromCPointer(featuresUnsafe)
	return features
}

func (v *InstanceVulkanDriver) attemptAvailableExtensions(physicalDevice core.PhysicalDevice, layerNamePtr *loader.Char) (map[string]*core1_0.ExtensionProperties, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	extensionCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	extensionCount := (*loader.Uint32)(extensionCountPtr)

	res, err := v.LoaderObj.VkEnumerateDeviceExtensionProperties(physicalDevice.Handle(), layerNamePtr, extensionCount, nil)

	if err != nil || *extensionCount == 0 {
		return nil, res, err
	}

	extensionTotal := int(*extensionCount)
	extensionsPtr := allocator.Malloc(extensionTotal * C.sizeof_struct_VkExtensionProperties)

	res, err = v.LoaderObj.VkEnumerateDeviceExtensionProperties(physicalDevice.Handle(), layerNamePtr, extensionCount, (*loader.VkExtensionProperties)(extensionsPtr))
	if err != nil || res == core1_0.VKIncomplete {
		return nil, res, err
	}

	retVal := make(map[string]*core1_0.ExtensionProperties)
	extensionSlice := ([]C.VkExtensionProperties)(unsafe.Slice((*C.VkExtensionProperties)(extensionsPtr), extensionTotal))

	for i := 0; i < extensionTotal; i++ {
		extension := extensionSlice[i]

		outExtension := &core1_0.ExtensionProperties{
			ExtensionName: C.GoString((*C.char)(&extension.extensionName[0])),
			SpecVersion:   uint(extension.specVersion),
		}

		existingExtension, ok := retVal[outExtension.ExtensionName]
		if ok && existingExtension.SpecVersion >= outExtension.SpecVersion {
			continue
		}
		retVal[outExtension.ExtensionName] = outExtension
	}

	return retVal, res, nil
}

func (v *InstanceVulkanDriver) EnumerateDeviceExtensionProperties(physicalDevice core.PhysicalDevice) (map[string]*core1_0.ExtensionProperties, common.VkResult, error) {
	if physicalDevice.Handle() == 0 {
		return nil, core1_0.VKErrorUnknown, fmt.Errorf("physicalDevice is uninitialized")
	}

	// There may be a race condition that adds new available extensions between getting the
	// extension count & pulling the extensions, in which case, attemptAvailableExtensions will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*core1_0.ExtensionProperties
	var result common.VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		layers, result, err = v.attemptAvailableExtensions(physicalDevice, nil)
	}
	return layers, result, err
}

func (v *InstanceVulkanDriver) EnumerateDeviceExtensionPropertiesForLayer(physicalDevice core.PhysicalDevice, layerName string) (map[string]*core1_0.ExtensionProperties, common.VkResult, error) {
	if physicalDevice.Handle() == 0 {
		return nil, core1_0.VKErrorUnknown, fmt.Errorf("physicalDevice is uninitialized")
	}

	// There may be a race condition that adds new available extensions between getting the
	// extension count & pulling the extensions, in which case, attemptAvailableExtensions will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*core1_0.ExtensionProperties
	var result common.VkResult
	var err error
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	layerNamePtr := (*loader.Char)(allocator.CString(layerName))
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		layers, result, err = v.attemptAvailableExtensions(physicalDevice, layerNamePtr)
	}
	return layers, result, err
}

func (v *InstanceVulkanDriver) attemptAvailableLayers(physicalDevice core.PhysicalDevice) (map[string]*core1_0.LayerProperties, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	layerCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	layerCount := (*loader.Uint32)(layerCountPtr)

	res, err := v.LoaderObj.VkEnumerateDeviceLayerProperties(physicalDevice.Handle(), layerCount, nil)

	if err != nil || *layerCount == 0 {
		return nil, res, err
	}

	layerTotal := int(*layerCount)
	layersPtr := allocator.Malloc(layerTotal * C.sizeof_struct_VkLayerProperties)

	res, err = v.LoaderObj.VkEnumerateDeviceLayerProperties(physicalDevice.Handle(), layerCount, (*loader.VkLayerProperties)(layersPtr))
	if err != nil || res == core1_0.VKIncomplete {
		return nil, res, err
	}

	retVal := make(map[string]*core1_0.LayerProperties)
	layerSlice := ([]C.VkLayerProperties)(unsafe.Slice((*C.VkLayerProperties)(layersPtr), layerTotal))

	for i := 0; i < layerTotal; i++ {
		layer := layerSlice[i]

		outLayer := &core1_0.LayerProperties{
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

func (v *InstanceVulkanDriver) EnumerateDeviceLayerProperties(physicalDevice core.PhysicalDevice) (map[string]*core1_0.LayerProperties, common.VkResult, error) {
	if physicalDevice.Handle() == 0 {
		return nil, core1_0.VKErrorUnknown, fmt.Errorf("physicalDevice is uninitialized")
	}

	// There may be a race condition that adds new available extensions between getting the
	// extension count & pulling the extensions, in which case, attemptAvailableExtensions will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*core1_0.LayerProperties
	var result common.VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		layers, result, err = v.attemptAvailableLayers(physicalDevice)
	}
	return layers, result, err
}

func (v *InstanceVulkanDriver) GetPhysicalDeviceFormatProperties(physicalDevice core.PhysicalDevice, format core1_0.Format) *core1_0.FormatProperties {
	if physicalDevice.Handle() == 0 {
		panic("physicalDevice is uninitialized")
	}
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	properties := (*C.VkFormatProperties)(allocator.Malloc(C.sizeof_struct_VkFormatProperties))

	v.LoaderObj.VkGetPhysicalDeviceFormatProperties(physicalDevice.Handle(), loader.VkFormat(format), (*loader.VkFormatProperties)(unsafe.Pointer(properties)))

	return &core1_0.FormatProperties{
		LinearTilingFeatures:  core1_0.FormatFeatureFlags(properties.linearTilingFeatures),
		OptimalTilingFeatures: core1_0.FormatFeatureFlags(properties.optimalTilingFeatures),
		BufferFeatures:        core1_0.FormatFeatureFlags(properties.bufferFeatures),
	}
}

func (v *InstanceVulkanDriver) GetPhysicalDeviceImageFormatProperties(physicalDevice core.PhysicalDevice, format core1_0.Format, imageType core1_0.ImageType, tiling core1_0.ImageTiling, usages core1_0.ImageUsageFlags, flags core1_0.ImageCreateFlags) (*core1_0.ImageFormatProperties, common.VkResult, error) {
	if physicalDevice.Handle() == 0 {
		return nil, core1_0.VKErrorUnknown, fmt.Errorf("physicalDevice is uninitialized")
	}
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	properties := (*C.VkImageFormatProperties)(allocator.Malloc(C.sizeof_struct_VkImageFormatProperties))

	res, err := v.LoaderObj.VkGetPhysicalDeviceImageFormatProperties(physicalDevice.Handle(), loader.VkFormat(format), loader.VkImageType(imageType), loader.VkImageTiling(tiling), loader.VkImageUsageFlags(usages), loader.VkImageCreateFlags(flags), (*loader.VkImageFormatProperties)(unsafe.Pointer(properties)))
	if err != nil {
		return nil, res, err
	}

	return &core1_0.ImageFormatProperties{
		MaxExtent: core1_0.Extent3D{
			Width:  int(properties.maxExtent.width),
			Height: int(properties.maxExtent.height),
			Depth:  int(properties.maxExtent.depth),
		},
		MaxMipLevels:    int(properties.maxMipLevels),
		MaxArrayLayers:  int(properties.maxArrayLayers),
		SampleCounts:    core1_0.SampleCountFlags(properties.sampleCounts),
		MaxResourceSize: int(properties.maxResourceSize),
	}, res, nil
}

func (v *InstanceVulkanDriver) GetPhysicalDeviceSparseImageFormatProperties(physicalDevice core.PhysicalDevice, format core1_0.Format, imageType core1_0.ImageType, samples core1_0.SampleCountFlags, usages core1_0.ImageUsageFlags, tiling core1_0.ImageTiling) []core1_0.SparseImageFormatProperties {
	if physicalDevice.Handle() == 0 {
		panic("physicalDevice is uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	propertiesCount := (*C.uint32_t)(arena.Malloc(4))

	v.LoaderObj.VkGetPhysicalDeviceSparseImageFormatProperties(physicalDevice.Handle(), loader.VkFormat(format), loader.VkImageType(imageType), loader.VkSampleCountFlagBits(samples), loader.VkImageUsageFlags(usages), loader.VkImageTiling(tiling), (*loader.Uint32)(propertiesCount), nil)

	if *propertiesCount == 0 {
		return nil
	}

	propertiesPtr := (*C.VkSparseImageFormatProperties)(arena.Malloc(int(*propertiesCount) * C.sizeof_struct_VkSparseImageFormatProperties))

	v.LoaderObj.VkGetPhysicalDeviceSparseImageFormatProperties(physicalDevice.Handle(), loader.VkFormat(format), loader.VkImageType(imageType), loader.VkSampleCountFlagBits(samples), loader.VkImageUsageFlags(usages), loader.VkImageTiling(tiling), (*loader.Uint32)(unsafe.Pointer(propertiesCount)), (*loader.VkSparseImageFormatProperties)(unsafe.Pointer(propertiesPtr)))

	propertiesSlice := ([]C.VkSparseImageFormatProperties)(unsafe.Slice(propertiesPtr, int(*propertiesCount)))

	outReqs := make([]core1_0.SparseImageFormatProperties, *propertiesCount)
	for j := 0; j < int(*propertiesCount); j++ {
		inProps := propertiesSlice[j]
		outReqs[j].Flags = core1_0.SparseImageFormatFlags(inProps.flags)
		outReqs[j].ImageGranularity = core1_0.Extent3D{
			Width:  int(inProps.imageGranularity.width),
			Height: int(inProps.imageGranularity.height),
			Depth:  int(inProps.imageGranularity.depth),
		}
		outReqs[j].AspectMask = core1_0.ImageAspectFlags(inProps.aspectMask)
	}

	return outReqs
}

func (v *InstanceVulkanDriver) CreateDevice(physicalDevice core.PhysicalDevice, allocationCallbacks *loader.AllocationCallbacks, options core1_0.DeviceCreateInfo) (core.Device, common.VkResult, error) {
	if physicalDevice.Handle() == 0 {
		return core.Device{}, core1_0.VKErrorUnknown, fmt.Errorf("physicalDevice is uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return core.Device{}, core1_0.VKErrorUnknown, err
	}

	var deviceHandle loader.VkDevice
	res, err := v.LoaderObj.VkCreateDevice(physicalDevice.Handle(), (*loader.VkDeviceCreateInfo)(createInfo), allocationCallbacks.Handle(), &deviceHandle)
	if err != nil {
		return core.Device{}, res, err
	}

	// deviceDriver, err := v.LoaderObj.CreateDeviceLoader(deviceHandle)
	// if err != nil {
	// 	return core.Device{}, core1_0.VKErrorUnknown, err
	// }

	device := core.InternalDevice(deviceHandle, physicalDevice.DeviceAPIVersion(), options.EnabledExtensionNames)

	return device, res, nil
}
