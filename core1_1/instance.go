package core1_1

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

type VulkanInstance struct {
	core1_0.Instance

	InstanceDriver driver.Driver
	InstanceHandle driver.VkInstance

	MaximumVersion common.APIVersion
}

func PromoteInstance(instance core1_0.Instance) Instance {
	if !instance.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return instance.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(instance.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanInstance{
				Instance: instance,

				InstanceDriver: instance.Driver(),
				InstanceHandle: instance.Handle(),

				MaximumVersion: instance.APIVersion(),
			}
		}).(Instance)
}

func (i *VulkanInstance) attemptEnumeratePhysicalDeviceGroups(outDataFactory func() *PhysicalDeviceGroupProperties) ([]*PhysicalDeviceGroupProperties, common.VkResult, error) {
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

	outDataSlice := make([]*PhysicalDeviceGroupProperties, count)
	for i := 0; i < count; i++ {
		if outDataFactory != nil {
			outDataSlice[i] = outDataFactory()
		} else {
			outDataSlice[i] = &PhysicalDeviceGroupProperties{}
		}
	}

	outData, err := common.AllocOutDataHeaderSlice[C.VkPhysicalDeviceGroupProperties, *PhysicalDeviceGroupProperties](arena, outDataSlice)
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

	err = common.PopulateOutDataSlice[C.VkPhysicalDeviceGroupProperties, *PhysicalDeviceGroupProperties](outDataSlice, unsafe.Pointer(outData),
		i.InstanceDriver, i.InstanceHandle, i.MaximumVersion)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	return outDataSlice, res, nil
}

func (i *VulkanInstance) EnumeratePhysicalDeviceGroups(outDataFactory func() *PhysicalDeviceGroupProperties) ([]*PhysicalDeviceGroupProperties, common.VkResult, error) {
	var outData []*PhysicalDeviceGroupProperties
	var result common.VkResult
	var err error

	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		outData, result, err = i.attemptEnumeratePhysicalDeviceGroups(outDataFactory)
	}
	return outData, result, err
}
