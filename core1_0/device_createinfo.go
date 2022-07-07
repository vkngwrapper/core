package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type DeviceCreateFlags int32

var deviceCreateFlags = common.NewFlagStringMapping[DeviceCreateFlags]()

func (f DeviceCreateFlags) Register(str string) {
	deviceCreateFlags.Register(f, str)
}

func (f DeviceCreateFlags) String() string {
	return deviceCreateFlags.FlagsToString(f)
}

////

type DeviceCreateInfo struct {
	Flags                 DeviceCreateFlags
	QueueCreateInfos      []DeviceQueueCreateInfo
	EnabledFeatures       *PhysicalDeviceFeatures
	EnabledExtensionNames []string
	EnabledLayerNames     []string

	common.NextOptions
}

func (o DeviceCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if len(o.QueueCreateInfos) == 0 {
		return nil, errors.New("alloc DeviceCreateInfo: no queue families added")
	}
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceCreateInfo{})))
	}

	// Alloc queue families
	queueFamilyPtr, err := common.AllocOptionSlice[C.VkDeviceQueueCreateInfo, DeviceQueueCreateInfo](allocator, o.QueueCreateInfos)
	if err != nil {
		return nil, err
	}

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

	createInfo := (*C.VkDeviceCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
	createInfo.flags = C.VkDeviceCreateFlags(o.Flags)
	createInfo.pNext = next
	createInfo.queueCreateInfoCount = C.uint32_t(len(o.QueueCreateInfos))
	createInfo.pQueueCreateInfos = (*C.VkDeviceQueueCreateInfo)(queueFamilyPtr)
	createInfo.enabledLayerCount = C.uint(numLayers)
	createInfo.ppEnabledLayerNames = (**C.char)(layerNamePtr)
	createInfo.enabledExtensionCount = C.uint(numExtensions)
	createInfo.ppEnabledExtensionNames = (**C.char)(extNamePtr)

	// Init feature list
	if o.EnabledFeatures != nil {
		featuresPtr, err := o.EnabledFeatures.PopulateCPointer(allocator, nil)
		if err != nil {
			return nil, err
		}

		createInfo.pEnabledFeatures = (*C.VkPhysicalDeviceFeatures)(featuresPtr)
	} else {
		createInfo.pEnabledFeatures = nil
	}

	return unsafe.Pointer(createInfo), nil
}
