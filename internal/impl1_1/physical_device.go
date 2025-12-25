package impl1_1

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *InstanceVulkanDriver) GetPhysicalDeviceExternalFenceProperties(physicalDevice core.PhysicalDevice, o core1_1.PhysicalDeviceExternalFenceInfo, outData *core1_1.ExternalFenceProperties) error {
	if physicalDevice.Handle() == 0 {
		return fmt.Errorf("physicalDevice cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOutDataHeader(arena, outData)
	if err != nil {
		return err
	}

	v.LoaderObj.VkGetPhysicalDeviceExternalFenceProperties(
		physicalDevice.Handle(),
		(*loader.VkPhysicalDeviceExternalFenceInfo)(infoPtr),
		(*loader.VkExternalFenceProperties)(outDataPtr),
	)

	return common.PopulateOutData(outData, outDataPtr)
}

func (v *InstanceVulkanDriver) GetPhysicalDeviceExternalBufferProperties(physicalDevice core.PhysicalDevice, o core1_1.PhysicalDeviceExternalBufferInfo, outData *core1_1.ExternalBufferProperties) error {
	if physicalDevice.Handle() == 0 {
		return fmt.Errorf("physicalDevice cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionsPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOutDataHeader(arena, outData)
	if err != nil {
		return err
	}

	v.LoaderObj.VkGetPhysicalDeviceExternalBufferProperties(physicalDevice.Handle(),
		(*loader.VkPhysicalDeviceExternalBufferInfo)(optionsPtr),
		(*loader.VkExternalBufferProperties)(outDataPtr))

	return common.PopulateOutData(outData, outDataPtr)
}

func (v *InstanceVulkanDriver) GetPhysicalDeviceExternalSemaphoreProperties(physicalDevice core.PhysicalDevice, o core1_1.PhysicalDeviceExternalSemaphoreInfo, outData *core1_1.ExternalSemaphoreProperties) error {
	if physicalDevice.Handle() == 0 {
		return fmt.Errorf("physicalDevice cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOutDataHeader(arena, outData)
	if err != nil {
		return err
	}

	v.LoaderObj.VkGetPhysicalDeviceExternalSemaphoreProperties(
		physicalDevice.Handle(),
		(*loader.VkPhysicalDeviceExternalSemaphoreInfo)(infoPtr),
		(*loader.VkExternalSemaphoreProperties)(outDataPtr),
	)

	return common.PopulateOutData(outData, outDataPtr)
}

func (v *InstanceVulkanDriver) GetPhysicalDeviceFeatures2(physicalDevice core.PhysicalDevice, out *core1_1.PhysicalDeviceFeatures2) error {
	if physicalDevice.Handle() == 0 {
		return fmt.Errorf("physicalDevice cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	v.LoaderObj.VkGetPhysicalDeviceFeatures2(physicalDevice.Handle(), (*loader.VkPhysicalDeviceFeatures2)(outData))

	return common.PopulateOutData(out, outData)
}

func (v *InstanceVulkanDriver) GetPhysicalDeviceFormatProperties2(physicalDevice core.PhysicalDevice, format core1_0.Format, out *core1_1.FormatProperties2) error {
	if physicalDevice.Handle() == 0 {
		return fmt.Errorf("physicalDevice cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	v.LoaderObj.VkGetPhysicalDeviceFormatProperties2(physicalDevice.Handle(), loader.VkFormat(format), (*loader.VkFormatProperties2)(outData))

	return common.PopulateOutData(out, outData)
}

func (v *InstanceVulkanDriver) GetPhysicalDeviceImageFormatProperties2(physicalDevice core.PhysicalDevice, o core1_1.PhysicalDeviceImageFormatInfo2, out *core1_1.ImageFormatProperties2) (common.VkResult, error) {
	if physicalDevice.Handle() == 0 {
		return core1_0.VKErrorUnknown, fmt.Errorf("physicalDevice cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionData, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	res, err := v.LoaderObj.VkGetPhysicalDeviceImageFormatProperties2(physicalDevice.Handle(),
		(*loader.VkPhysicalDeviceImageFormatInfo2)(optionData),
		(*loader.VkImageFormatProperties2)(outData),
	)
	if err != nil {
		return res, err
	}

	err = common.PopulateOutData(out, outData)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return res, nil
}

func (v *InstanceVulkanDriver) GetPhysicalDeviceMemoryProperties2(physicalDevice core.PhysicalDevice, out *core1_1.PhysicalDeviceMemoryProperties2) error {
	if physicalDevice.Handle() == 0 {
		return fmt.Errorf("physicalDevice cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	v.LoaderObj.VkGetPhysicalDeviceMemoryProperties2(physicalDevice.Handle(), (*loader.VkPhysicalDeviceMemoryProperties2)(outData))

	return common.PopulateOutData(out, outData)
}

func (v *InstanceVulkanDriver) GetPhysicalDeviceProperties2(physicalDevice core.PhysicalDevice, out *core1_1.PhysicalDeviceProperties2) error {
	if physicalDevice.Handle() == 0 {
		return fmt.Errorf("physicalDevice cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	v.LoaderObj.VkGetPhysicalDeviceProperties2(physicalDevice.Handle(),
		(*loader.VkPhysicalDeviceProperties2)(outData))

	return common.PopulateOutData(out, outData)
}

func (v *InstanceVulkanDriver) GetPhysicalDeviceQueueFamilyProperties2(physicalDevice core.PhysicalDevice, outDataFactory func() *core1_1.QueueFamilyProperties2) ([]*core1_1.QueueFamilyProperties2, error) {
	if physicalDevice.Handle() == 0 {
		return nil, fmt.Errorf("physicalDevice cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outDataCountPtr := (*loader.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	v.LoaderObj.VkGetPhysicalDeviceQueueFamilyProperties2(physicalDevice.Handle(), outDataCountPtr, nil)

	outDataCount := int(*outDataCountPtr)
	if outDataCount == 0 {
		return nil, nil
	}

	out := make([]*core1_1.QueueFamilyProperties2, outDataCount)
	for i := 0; i < outDataCount; i++ {
		if outDataFactory != nil {
			out[i] = outDataFactory()
		} else {
			out[i] = &core1_1.QueueFamilyProperties2{}
		}
	}

	outData, err := common.AllocOutDataHeaderSlice[C.VkQueueFamilyProperties2, *core1_1.QueueFamilyProperties2](arena, out)
	if err != nil {
		return nil, err
	}

	v.LoaderObj.VkGetPhysicalDeviceQueueFamilyProperties2(physicalDevice.Handle(), outDataCountPtr, (*loader.VkQueueFamilyProperties2)(unsafe.Pointer(outData)))

	err = common.PopulateOutDataSlice[C.VkQueueFamilyProperties2, *core1_1.QueueFamilyProperties2](out, unsafe.Pointer(outData))
	return out, err
}

func (v *InstanceVulkanDriver) GetPhysicalDeviceSparseImageFormatProperties2(physicalDevice core.PhysicalDevice, o core1_1.PhysicalDeviceSparseImageFormatInfo2, outDataFactory func() *core1_1.SparseImageFormatProperties2) ([]*core1_1.SparseImageFormatProperties2, error) {
	if physicalDevice.Handle() == 0 {
		return nil, fmt.Errorf("physicalDevice cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outDataCountPtr := (*loader.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	optionData, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, err
	}

	v.LoaderObj.VkGetPhysicalDeviceSparseImageFormatProperties2(physicalDevice.Handle(), (*loader.VkPhysicalDeviceSparseImageFormatInfo2)(optionData), outDataCountPtr, nil)

	outDataCount := int(*outDataCountPtr)
	if outDataCount == 0 {
		return nil, nil
	}

	out := make([]*core1_1.SparseImageFormatProperties2, outDataCount)
	for i := 0; i < outDataCount; i++ {
		if outDataFactory != nil {
			out[i] = outDataFactory()
		} else {
			out[i] = &core1_1.SparseImageFormatProperties2{}
		}
	}

	outData, err := common.AllocOutDataHeaderSlice[C.VkSparseImageFormatProperties2, *core1_1.SparseImageFormatProperties2](arena, out)
	if err != nil {
		return nil, err
	}

	v.LoaderObj.VkGetPhysicalDeviceSparseImageFormatProperties2(physicalDevice.Handle(),
		(*loader.VkPhysicalDeviceSparseImageFormatInfo2)(optionData),
		outDataCountPtr,
		(*loader.VkSparseImageFormatProperties2)(unsafe.Pointer(outData)),
	)

	err = common.PopulateOutDataSlice[C.VkSparseImageFormatProperties2, *core1_1.SparseImageFormatProperties2](out, unsafe.Pointer(outData))

	return out, err
}
