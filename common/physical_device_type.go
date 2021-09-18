package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type PhysicalDeviceType int32

const (
	DeviceOther         PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_OTHER
	DeviceIntegratedGPU PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU
	DeviceDiscreteGPU   PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU
	DeviceVirtualGPU    PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU
	DeviceCPU           PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_CPU
)

var physicalDeviceTypeToString = map[PhysicalDeviceType]string{
	DeviceOther:         "Other",
	DeviceIntegratedGPU: "Integrated GPU",
	DeviceDiscreteGPU:   "Discrete GPU",
	DeviceVirtualGPU:    "Virtual GPU",
	DeviceCPU:           "CPU",
}

func (t PhysicalDeviceType) String() string {
	return physicalDeviceTypeToString[t]
}
