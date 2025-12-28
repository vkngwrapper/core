package impl1_1

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *InstanceVulkanDriver) attemptEnumeratePhysicalDeviceGroups(outDataFactory func() *core1_1.PhysicalDeviceGroupProperties) ([]*core1_1.PhysicalDeviceGroupProperties, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	countPtr := (*loader.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := v.LoaderObj.VkEnumeratePhysicalDeviceGroups(
		v.InstanceObj.Handle(),
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

	outDataSlice := make([]*core1_1.PhysicalDeviceGroupProperties, count)
	for i := 0; i < count; i++ {
		if outDataFactory != nil {
			outDataSlice[i] = outDataFactory()
		} else {
			outDataSlice[i] = &core1_1.PhysicalDeviceGroupProperties{}
		}
	}

	outData, err := common.AllocOutDataHeaderSlice[C.VkPhysicalDeviceGroupProperties, *core1_1.PhysicalDeviceGroupProperties](arena, outDataSlice)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	res, err = v.LoaderObj.VkEnumeratePhysicalDeviceGroups(
		v.InstanceObj.Handle(),
		countPtr,
		(*loader.VkPhysicalDeviceGroupProperties)(unsafe.Pointer(outData)),
	)
	if err != nil {
		return nil, res, err
	}

	err = common.PopulateOutDataSlice[C.VkPhysicalDeviceGroupProperties, *core1_1.PhysicalDeviceGroupProperties](outDataSlice, unsafe.Pointer(outData),
		v.LoaderObj, v.InstanceObj.APIVersion())
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	return outDataSlice, res, nil
}

func (v *InstanceVulkanDriver) EnumeratePhysicalDeviceGroups(outDataFactory func() *core1_1.PhysicalDeviceGroupProperties) ([]*core1_1.PhysicalDeviceGroupProperties, common.VkResult, error) {
	var outData []*core1_1.PhysicalDeviceGroupProperties
	var result common.VkResult
	var err error

	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		outData, result, err = v.attemptEnumeratePhysicalDeviceGroups(outDataFactory)
	}
	return outData, result, err
}
