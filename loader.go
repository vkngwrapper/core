package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanLoader struct {
	driver driver.Driver
}

func NewLoader(
	driver driver.Driver,
) *VulkanLoader {
	return &VulkanLoader{
		driver: driver,
	}
}

var _ Loader = &VulkanLoader{}

func (l *VulkanLoader) APIVersion() common.APIVersion {
	return l.driver.Version()
}

func (l *VulkanLoader) Driver() driver.Driver {
	return l.driver
}

func (l *VulkanLoader) attemptAvailableExtensions(layerName *driver.Char) (map[string]*core1_0.ExtensionProperties, common.VkResult, error) {
	alloc := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(alloc)

	extensionCount := (*driver.Uint32)(alloc.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := l.driver.VkEnumerateInstanceExtensionProperties(layerName, extensionCount, nil)
	if err != nil || *extensionCount == 0 {
		return nil, res, err
	}

	extensionsUnsafe := alloc.Malloc(int(*extensionCount) * C.sizeof_struct_VkExtensionProperties)

	res, err = l.driver.VkEnumerateInstanceExtensionProperties(layerName, extensionCount, (*driver.VkExtensionProperties)(extensionsUnsafe))
	if err != nil || res == core1_0.VKIncomplete {
		return nil, res, err
	}

	intExtensionCount := int(*extensionCount)
	extensionArray := ([]C.VkExtensionProperties)(unsafe.Slice((*C.VkExtensionProperties)(extensionsUnsafe), intExtensionCount))
	outExtensions := make(map[string]*core1_0.ExtensionProperties)
	for i := 0; i < intExtensionCount; i++ {
		extension := extensionArray[i]

		outExtension := &core1_0.ExtensionProperties{
			ExtensionName: C.GoString((*C.char)(&extension.extensionName[0])),
			SpecVersion:   common.Version(extension.specVersion),
		}

		existingExtension, ok := outExtensions[outExtension.ExtensionName]
		if ok && existingExtension.SpecVersion >= outExtension.SpecVersion {
			continue
		}
		outExtensions[outExtension.ExtensionName] = outExtension
	}

	return outExtensions, res, nil
}

func (l *VulkanLoader) AvailableExtensions() (map[string]*core1_0.ExtensionProperties, common.VkResult, error) {
	// There may be a race condition that adds new available extensions between getting the
	// extension count & pulling the extensions, in which case, attemptAvailableExtensions will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*core1_0.ExtensionProperties
	var result common.VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		layers, result, err = l.attemptAvailableExtensions(nil)
	}
	return layers, result, err
}

func (l *VulkanLoader) AvailableExtensionsForLayer(layerName string) (map[string]*core1_0.ExtensionProperties, common.VkResult, error) {
	// There may be a race condition that adds new available extensions between getting the
	// extension count & pulling the extensions, in which case, attemptAvailableExtensions will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*core1_0.ExtensionProperties
	var result common.VkResult
	var err error
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	layerNamePtr := (*driver.Char)(allocator.CString(layerName))
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		layers, result, err = l.attemptAvailableExtensions(layerNamePtr)
	}
	return layers, result, err
}

func (l *VulkanLoader) attemptAvailableLayers() (map[string]*core1_0.LayerProperties, common.VkResult, error) {
	alloc := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(alloc)

	layerCount := (*driver.Uint32)(alloc.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := l.driver.VkEnumerateInstanceLayerProperties(layerCount, nil)
	if err != nil || *layerCount == 0 {
		return nil, res, err
	}

	layersUnsafe := alloc.Malloc(int(*layerCount) * C.sizeof_struct_VkLayerProperties)

	res, err = l.driver.VkEnumerateInstanceLayerProperties(layerCount, (*driver.VkLayerProperties)(layersUnsafe))
	if err != nil || res == core1_0.VKIncomplete {
		return nil, res, err
	}

	intLayerCount := int(*layerCount)
	layerArray := ([]C.VkLayerProperties)(unsafe.Slice((*C.VkLayerProperties)(layersUnsafe), intLayerCount))
	outLayers := make(map[string]*core1_0.LayerProperties)
	for i := 0; i < intLayerCount; i++ {
		layer := layerArray[i]

		outLayer := &core1_0.LayerProperties{
			LayerName:             C.GoString((*C.char)(&layer.layerName[0])),
			SpecVersion:           common.Version(layer.specVersion),
			ImplementationVersion: common.Version(layer.implementationVersion),
			Description:           C.GoString((*C.char)(&layer.description[0])),
		}

		existingLayer, ok := outLayers[outLayer.LayerName]
		if ok && existingLayer.SpecVersion >= outLayer.SpecVersion {
			continue
		}
		outLayers[outLayer.LayerName] = outLayer
	}

	return outLayers, res, nil
}

func (l *VulkanLoader) AvailableLayers() (map[string]*core1_0.LayerProperties, common.VkResult, error) {
	// There may be a race condition that adds new available layers between getting the
	// layer count & pulling the layers, in which case, attemptAvailableLayers will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*core1_0.LayerProperties
	var result common.VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		layers, result, err = l.attemptAvailableLayers()
	}
	return layers, result, err
}

//go:linkname createInstanceObject github.com/CannibalVox/VKng/core/core1_0.createInstanceObject
func createInstanceObject(instanceDriver driver.Driver, handle driver.VkInstance, version common.APIVersion) *core1_0.VulkanInstance

func (l *VulkanLoader) CreateInstance(allocationCallbacks *driver.AllocationCallbacks, options core1_0.InstanceCreateOptions) (core1_0.Instance, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var instanceHandle driver.VkInstance

	res, err := l.driver.VkCreateInstance((*driver.VkInstanceCreateInfo)(createInfo), allocationCallbacks.Handle(), &instanceHandle)
	if err != nil {
		return nil, res, err
	}

	instanceDriver, err := l.driver.CreateInstanceDriver((driver.VkInstance)(instanceHandle))
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	version := l.APIVersion().Min(options.VulkanVersion)
	instance := createInstanceObject(instanceDriver, instanceHandle, version)

	for _, extension := range options.ExtensionNames {
		instance.ActiveInstanceExtensions[extension] = struct{}{}
	}

	return instance, res, nil
}
