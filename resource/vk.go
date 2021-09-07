package resource

/*
#cgo windows LDFLAGS: -lvulkan
#cgo linux freebsd darwin openbsd pkg-config: vulkan
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

func AvailableExtensions(alloc cgoalloc.Allocator) (map[string]*core.ExtensionProperties, core.Result, error) {
	extensionCount := (*C.uint32_t)(alloc.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	defer alloc.Free(unsafe.Pointer(extensionCount))

	res := core.Result(C.vkEnumerateInstanceExtensionProperties(nil, extensionCount, nil))
	err := res.ToError()
	if err != nil || *extensionCount == 0 {
		return nil, res, err
	}

	extensions := (*C.VkExtensionProperties)(alloc.Malloc(int(*extensionCount) * int(unsafe.Sizeof(C.VkExtensionProperties{}))))
	defer alloc.Free(unsafe.Pointer(extensions))

	res = core.Result(C.vkEnumerateInstanceExtensionProperties(nil, extensionCount, extensions))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	intExtensionCount := int(*extensionCount)
	extensionArray := ([]C.VkExtensionProperties)(unsafe.Slice(extensions, intExtensionCount))
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

func AvailableLayers(alloc cgoalloc.Allocator) (map[string]*core.LayerProperties, core.Result, error) {
	layerCount := (*C.uint32_t)(alloc.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	defer alloc.Free(unsafe.Pointer(layerCount))

	res := core.Result(C.vkEnumerateInstanceLayerProperties(layerCount, nil))
	err := res.ToError()
	if err != nil || *layerCount == 0 {
		return nil, res, err
	}

	layers := (*C.VkLayerProperties)(alloc.Malloc(int(*layerCount) * int(unsafe.Sizeof(C.VkLayerProperties{}))))
	defer alloc.Free(unsafe.Pointer(layers))

	res = core.Result(C.vkEnumerateInstanceLayerProperties(layerCount, layers))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	intLayerCount := int(*layerCount)
	layerArray := ([]C.VkLayerProperties)(unsafe.Slice(layers, intLayerCount))
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
