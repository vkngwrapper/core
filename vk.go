package core

/*
#cgo windows LDFLAGS: -lvulkan
#cgo linux freebsd darwin openbsd pkg-config: vulkan
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

func AvailableExtensions(alloc cgoalloc.Allocator) (map[string]*VKng.ExtensionProperties, VKng.Result, error) {
	extensionCount := (*C.uint32_t)(alloc.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	defer alloc.Free(unsafe.Pointer(extensionCount))

	res := VKng.Result(C.vkEnumerateInstanceExtensionProperties(nil, extensionCount, nil))
	err := res.ToError()
	if err != nil || *extensionCount == 0 {
		return nil, res, err
	}

	extensions := (*C.VkExtensionProperties)(alloc.Malloc(int(*extensionCount) * int(unsafe.Sizeof(C.VkExtensionProperties{}))))
	defer alloc.Free(unsafe.Pointer(extensions))

	res = VKng.Result(C.vkEnumerateInstanceExtensionProperties(nil, extensionCount, extensions))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	intExtensionCount := int(*extensionCount)
	extensionArray := ([]C.VkExtensionProperties)(unsafe.Slice(extensions, intExtensionCount))
	outExtensions := make(map[string]*VKng.ExtensionProperties)
	for i := 0; i < intExtensionCount; i++ {
		extension := extensionArray[i]

		outExtension := &VKng.ExtensionProperties{
			ExtensionName: C.GoString((*C.char)(&extension.extensionName[0])),
			SpecVersion:   VKng.Version(extension.specVersion),
		}

		existingExtension, ok := outExtensions[outExtension.ExtensionName]
		if ok && existingExtension.SpecVersion >= outExtension.SpecVersion {
			continue
		}
		outExtensions[outExtension.ExtensionName] = outExtension
	}

	return outExtensions, res, nil
}

func AvailableLayers(alloc cgoalloc.Allocator) (map[string]*VKng.LayerProperties, VKng.Result, error) {
	layerCount := (*C.uint32_t)(alloc.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	defer alloc.Free(unsafe.Pointer(layerCount))

	res := VKng.Result(C.vkEnumerateInstanceLayerProperties(layerCount, nil))
	err := res.ToError()
	if err != nil || *layerCount == 0 {
		return nil, res, err
	}

	layers := (*C.VkLayerProperties)(alloc.Malloc(int(*layerCount) * int(unsafe.Sizeof(C.VkLayerProperties{}))))
	defer alloc.Free(unsafe.Pointer(layers))

	res = VKng.Result(C.vkEnumerateInstanceLayerProperties(layerCount, layers))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	intLayerCount := int(*layerCount)
	layerArray := ([]C.VkLayerProperties)(unsafe.Slice(layers, intLayerCount))
	outLayers := make(map[string]*VKng.LayerProperties)
	for i := 0; i < intLayerCount; i++ {
		layer := layerArray[i]

		outLayer := &VKng.LayerProperties{
			LayerName:             C.GoString((*C.char)(&layer.layerName[0])),
			SpecVersion:           VKng.Version(layer.specVersion),
			ImplementationVersion: VKng.Version(layer.implementationVersion),
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
