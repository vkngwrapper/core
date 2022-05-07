package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanInstance struct {
	InstanceDriver driver.Driver
	InstanceHandle driver.VkInstance

	MaximumVersion common.APIVersion
}

func (i *VulkanInstance) attemptEnumeratePhysicalDeviceGroups(outDataFactory func() *core1_1.DeviceGroupOutData) ([]*core1_1.DeviceGroupOutData, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	countPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := i.InstanceDriver.VkEnumeratePhysicalDeviceGroups(
		i.InstanceHandle,
		countPtr,
		nil,
	)
	if err != nil {
		return nil, res, err
	}

	count := int(*countPtr)
	if count == 0 {
		return nil, core1_0.VKSuccess, nil
	}

	outDataSlice := make([]*core1_1.DeviceGroupOutData, count)
	for i := 0; i < count; i++ {
		if outDataFactory != nil {
			outDataSlice[i] = outDataFactory()
		} else {
			outDataSlice[i] = &core1_1.DeviceGroupOutData{}
		}
	}

	outData, err := common.AllocOptionSlice[C.VkPhysicalDeviceGroupProperties, *core1_1.DeviceGroupOutData](arena, outDataSlice)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	res, err = i.InstanceDriver.VkEnumeratePhysicalDeviceGroups(
		i.InstanceHandle,
		countPtr,
		(*driver.VkPhysicalDeviceGroupProperties)(unsafe.Pointer(outData)),
	)
	if err != nil {
		return nil, res, err
	}

	err = common.PopulateOutDataSlice[C.VkPhysicalDeviceGroupProperties, *core1_1.DeviceGroupOutData](outDataSlice, unsafe.Pointer(outData),
		i.InstanceDriver, i.InstanceHandle, i.MaximumVersion)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	return outDataSlice, res, nil
}

func (i *VulkanInstance) PhysicalDeviceGroups(outDataFactory func() *core1_1.DeviceGroupOutData) ([]*core1_1.DeviceGroupOutData, common.VkResult, error) {
	var outData []*core1_1.DeviceGroupOutData
	var result common.VkResult
	var err error

	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		outData, result, err = i.attemptEnumeratePhysicalDeviceGroups(outDataFactory)
	}
	return outData, result, err
}
