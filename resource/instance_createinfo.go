package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type InstanceOptions struct {
	ApplicationName    string
	ApplicationVersion core.Version
	EngineName         string
	EngineVersion      core.Version
	VulkanVersion      core.APIVersion

	ExtensionNames []string
	LayerNames     []string

	Next core.Options
}

func (o *InstanceOptions) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	cApplication := cgoalloc.CString(allocator, o.ApplicationName)
	cEngine := cgoalloc.CString(allocator, o.EngineName)

	appInfo := (*C.VkApplicationInfo)(allocator.Malloc(int(unsafe.Sizeof(C.VkApplicationInfo{}))))

	appInfo.sType = C.VK_STRUCTURE_TYPE_APPLICATION_INFO
	appInfo.pApplicationName = (*C.char)(cApplication)
	appInfo.pEngineName = (*C.char)(cEngine)
	appInfo.applicationVersion = C.uint32_t(o.ApplicationVersion)
	appInfo.engineVersion = C.uint32_t(o.EngineVersion)
	appInfo.apiVersion = C.uint32_t(o.VulkanVersion)

	createInfo := (*C.VkInstanceCreateInfo)(allocator.Malloc(int(unsafe.Sizeof(C.VkInstanceCreateInfo{}))))

	// Alloc array of extension names
	numExtensions := len(o.ExtensionNames)
	extNamePtr := allocator.Malloc(numExtensions * int(unsafe.Sizeof(uintptr(0))))
	extNames := ([]*C.char)(unsafe.Slice((**C.char)(extNamePtr), numExtensions))
	for i := 0; i < numExtensions; i++ {
		extNames[i] = (*C.char)(cgoalloc.CString(allocator, o.ExtensionNames[i]))
	}

	// Alloc array of layer names
	numLayers := len(o.LayerNames)
	layerNamePtr := allocator.Malloc(numLayers * int(unsafe.Sizeof(uintptr(0))))
	layerNames := ([]*C.char)(unsafe.Slice((**C.char)(layerNamePtr), numLayers))
	for i := 0; i < numLayers; i++ {
		layerNames[i] = (*C.char)(cgoalloc.CString(allocator, o.LayerNames[i]))
	}

	createInfo.sType = C.VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO
	createInfo.flags = 0
	createInfo.pApplicationInfo = appInfo
	createInfo.enabledExtensionCount = C.uint32_t(numExtensions)
	createInfo.ppEnabledExtensionNames = (**C.char)(extNamePtr)
	createInfo.enabledLayerCount = C.uint32_t(numLayers)
	createInfo.ppEnabledLayerNames = (**C.char)(layerNamePtr)

	var err error
	var next unsafe.Pointer
	if o.Next != nil {
		next, err = o.Next.AllocForC(allocator)
	}

	if err != nil {
		return nil, err
	}
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}
