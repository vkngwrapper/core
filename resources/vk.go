package resources

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

func AvailableExtensions(load loader.Loader) (map[string]*core.ExtensionProperties, loader.VkResult, error) {
	alloc := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(alloc)

	extensionCount := (*loader.Uint32)(alloc.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := load.VkEnumerateInstanceExtensionProperties(nil, extensionCount, nil)
	if err != nil || *extensionCount == 0 {
		return nil, res, err
	}

	extensionsUnsafe := alloc.Malloc(int(*extensionCount) * int(unsafe.Sizeof(C.VkExtensionProperties{})))

	res, err = load.VkEnumerateInstanceExtensionProperties(nil, extensionCount, (*loader.VkExtensionProperties)(extensionsUnsafe))
	if err != nil {
		return nil, res, err
	}

	intExtensionCount := int(*extensionCount)
	extensionArray := ([]C.VkExtensionProperties)(unsafe.Slice((*C.VkExtensionProperties)(extensionsUnsafe), intExtensionCount))
	outExtensions := make(map[string]*core.ExtensionProperties)
	for i := 0; i < intExtensionCount; i++ {
		extension := extensionArray[i]

		outExtension := &core.ExtensionProperties{
			ExtensionName: C.GoString((*C.char)(&extension.extensionName[0])),
			SpecVersion:   core.Version(extension.specVersion),
		}

		existingExtension, ok := outExtensions[outExtension.ExtensionName]
		if ok && existingExtension.SpecVersion >= outExtension.SpecVersion {
			continue
		}
		outExtensions[outExtension.ExtensionName] = outExtension
	}

	return outExtensions, res, nil
}

func AvailableLayers(load loader.Loader) (map[string]*core.LayerProperties, loader.VkResult, error) {
	alloc := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(alloc)

	layerCount := (*loader.Uint32)(alloc.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := load.VkEnumerateInstanceLayerProperties(layerCount, nil)
	if err != nil || *layerCount == 0 {
		return nil, res, err
	}

	layersUnsafe := alloc.Malloc(int(*layerCount) * int(unsafe.Sizeof(C.VkLayerProperties{})))

	res, err = load.VkEnumerateInstanceLayerProperties(layerCount, (*loader.VkLayerProperties)(layersUnsafe))
	if err != nil {
		return nil, res, err
	}

	intLayerCount := int(*layerCount)
	layerArray := ([]C.VkLayerProperties)(unsafe.Slice((*C.VkLayerProperties)(layersUnsafe), intLayerCount))
	outLayers := make(map[string]*core.LayerProperties)
	for i := 0; i < intLayerCount; i++ {
		layer := layerArray[i]

		outLayer := &core.LayerProperties{
			LayerName:             C.GoString((*C.char)(&layer.layerName[0])),
			SpecVersion:           core.Version(layer.specVersion),
			ImplementationVersion: core.Version(layer.implementationVersion),
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
