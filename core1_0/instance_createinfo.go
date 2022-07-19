package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// InstanceCreateFlags specifies behavior of the Instance
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateFlagBits.html
type InstanceCreateFlags int32

var instanceCreateFlagsMapping = common.NewFlagStringMapping[InstanceCreateFlags]()

func (f InstanceCreateFlags) Register(str string) {
	instanceCreateFlagsMapping.Register(f, str)
}

func (f InstanceCreateFlags) String() string {
	return instanceCreateFlagsMapping.FlagsToString(f)
}

////

// InstanceCreateInfo specifies parameters of a newly-created Instance
type InstanceCreateInfo struct {
	// ApplicationName is a string containing the name of the application
	ApplicationName string
	// ApplicationVersion contains the developer-supplied verison number of the application
	ApplicationVersion common.Version
	// EngineName is a string containing the name of the engine, if any, used to create
	// the application
	EngineName string
	// EngineVersion contains the developer-supplied version number of the engine used to
	// create the application
	EngineVersion common.Version
	// APIVersion must be the highest version of Vulkan that the application is designed to use
	APIVersion common.APIVersion

	// Flags indicates the behavior of the Instance
	Flags InstanceCreateFlags

	// EnabledExtensionNames is a slice of strings containing the names of extensions to enable
	EnabledExtensionNames []string
	// EnabledLayerNames is a slice of strings containing the names of layers to enable for the
	// created Instance
	EnabledLayerNames []string

	common.NextOptions
}

func (o InstanceCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkInstanceCreateInfo{})))
	}

	appInfo := (*C.VkApplicationInfo)(allocator.Malloc(int(unsafe.Sizeof(C.VkApplicationInfo{}))))

	appInfo.sType = C.VK_STRUCTURE_TYPE_APPLICATION_INFO
	appInfo.pNext = nil
	appInfo.pApplicationName = nil
	appInfo.pEngineName = nil

	if o.ApplicationName != "" {
		cApplication := allocator.CString(o.ApplicationName)
		appInfo.pApplicationName = (*C.char)(cApplication)
	}

	if o.EngineName != "" {
		cEngine := allocator.CString(o.EngineName)
		appInfo.pEngineName = (*C.char)(cEngine)
	}

	appInfo.applicationVersion = C.uint32_t(o.ApplicationVersion)
	appInfo.engineVersion = C.uint32_t(o.EngineVersion)
	appInfo.apiVersion = C.uint32_t(o.APIVersion)

	createInfo := (*C.VkInstanceCreateInfo)(preallocatedPointer)

	// Alloc array of extension names
	numExtensions := len(o.EnabledExtensionNames)
	extNamePtr := allocator.Malloc(numExtensions * int(unsafe.Sizeof(uintptr(0))))
	extNames := ([]*C.char)(unsafe.Slice((**C.char)(extNamePtr), numExtensions))
	for i := 0; i < numExtensions; i++ {
		extNames[i] = (*C.char)(allocator.CString(o.EnabledExtensionNames[i]))
	}

	// Alloc array of layer names
	numLayers := len(o.EnabledLayerNames)
	layerNamePtr := allocator.Malloc(numLayers * int(unsafe.Sizeof(uintptr(0))))
	layerNames := ([]*C.char)(unsafe.Slice((**C.char)(layerNamePtr), numLayers))
	for i := 0; i < numLayers; i++ {
		layerNames[i] = (*C.char)(allocator.CString(o.EnabledLayerNames[i]))
	}

	createInfo.sType = C.VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO
	createInfo.flags = C.VkInstanceCreateFlags(o.Flags)
	createInfo.pNext = next
	createInfo.pApplicationInfo = appInfo
	createInfo.enabledExtensionCount = C.uint32_t(numExtensions)
	createInfo.ppEnabledExtensionNames = (**C.char)(extNamePtr)
	createInfo.enabledLayerCount = C.uint32_t(numLayers)
	createInfo.ppEnabledLayerNames = (**C.char)(layerNamePtr)

	return preallocatedPointer, nil
}
