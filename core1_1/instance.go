package core1_1

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	"unsafe"
)

// VulkanInstance is an implementation of the Instance interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanInstance struct {
	core1_0.Instance

	InstanceDriver driver.Driver
	InstanceHandle driver.VkInstance

	MaximumVersion common.APIVersion
}

// PromoteInstance accepts a Instance object from any core version. If provided a instance that supports
// at least core 1.1, it will return a core1_1.Instance. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanInstance, even if it is provided a VulkanInstance from a higher
// core version. Two Vulkan 1.1 compatible Instance objects with the same Instance.Handle will
// return the same interface value when passed to this method.
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
