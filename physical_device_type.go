package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type PhysicalDeviceType int32

const (
	Other         PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_OTHER
	IntegratedGPU PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU
	DiscreteGPU   PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU
	VirtualGPU    PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU
	CPU           PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_CPU
)

var physicalDeviceTypeToString = map[PhysicalDeviceType]string{
	Other:         "Other",
	IntegratedGPU: "Integrated GPU",
	DiscreteGPU:   "Discrete GPU",
	VirtualGPU:    "Virtual GPU",
	CPU:           "CPU",
}

func (t PhysicalDeviceType) String() string {
	return physicalDeviceTypeToString[t]
}
