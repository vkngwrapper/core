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
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
)

// VulkanPhysicalDevice is an implementation of the PhysicalDevice interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanPhysicalDevice struct {
	impl1_0.VulkanPhysicalDevice
}

func (p *VulkanPhysicalDevice) ExternalFenceProperties(o core1_1.PhysicalDeviceExternalFenceInfo, outData *core1_1.ExternalFenceProperties) error {
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

	p.InstanceDriver.VkGetPhysicalDeviceExternalFenceProperties(
		p.PhysicalDeviceHandle,
		(*driver.VkPhysicalDeviceExternalFenceInfo)(infoPtr),
		(*driver.VkExternalFenceProperties)(outDataPtr),
	)

	return common.PopulateOutData(outData, outDataPtr)
}

func (p *VulkanPhysicalDevice) ExternalBufferProperties(o core1_1.PhysicalDeviceExternalBufferInfo, outData *core1_1.ExternalBufferProperties) error {
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

	p.InstanceDriver.VkGetPhysicalDeviceExternalBufferProperties(p.PhysicalDeviceHandle,
		(*driver.VkPhysicalDeviceExternalBufferInfo)(optionsPtr),
		(*driver.VkExternalBufferProperties)(outDataPtr))

	return common.PopulateOutData(outData, outDataPtr)
}

func (p *VulkanPhysicalDevice) ExternalSemaphoreProperties(o core1_1.PhysicalDeviceExternalSemaphoreInfo, outData *core1_1.ExternalSemaphoreProperties) error {
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

	p.InstanceDriver.VkGetPhysicalDeviceExternalSemaphoreProperties(
		p.PhysicalDeviceHandle,
		(*driver.VkPhysicalDeviceExternalSemaphoreInfo)(infoPtr),
		(*driver.VkExternalSemaphoreProperties)(outDataPtr),
	)

	return common.PopulateOutData(outData, outDataPtr)
}

func (p *VulkanPhysicalDevice) Features2(out *core1_1.PhysicalDeviceFeatures2) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	p.InstanceDriver.VkGetPhysicalDeviceFeatures2(p.PhysicalDeviceHandle, (*driver.VkPhysicalDeviceFeatures2)(outData))

	return common.PopulateOutData(out, outData)
}

func (p *VulkanPhysicalDevice) FormatProperties2(format core1_0.Format, out *core1_1.FormatProperties2) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	p.InstanceDriver.VkGetPhysicalDeviceFormatProperties2(p.PhysicalDeviceHandle, driver.VkFormat(format), (*driver.VkFormatProperties2)(outData))

	return common.PopulateOutData(out, outData)
}

func (p *VulkanPhysicalDevice) ImageFormatProperties2(o core1_1.PhysicalDeviceImageFormatInfo2, out *core1_1.ImageFormatProperties2) (common.VkResult, error) {
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

	res, err := p.InstanceDriver.VkGetPhysicalDeviceImageFormatProperties2(p.PhysicalDeviceHandle,
		(*driver.VkPhysicalDeviceImageFormatInfo2)(optionData),
		(*driver.VkImageFormatProperties2)(outData),
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

func (p *VulkanPhysicalDevice) MemoryProperties2(out *core1_1.PhysicalDeviceMemoryProperties2) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	p.InstanceDriver.VkGetPhysicalDeviceMemoryProperties2(p.PhysicalDeviceHandle, (*driver.VkPhysicalDeviceMemoryProperties2)(outData))

	return common.PopulateOutData(out, outData)
}

func (p *VulkanPhysicalDevice) Properties2(out *core1_1.PhysicalDeviceProperties2) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	p.InstanceDriver.VkGetPhysicalDeviceProperties2(p.PhysicalDeviceHandle,
		(*driver.VkPhysicalDeviceProperties2)(outData))

	return common.PopulateOutData(out, outData)
}

func (p *VulkanPhysicalDevice) QueueFamilyProperties2(outDataFactory func() *core1_1.QueueFamilyProperties2) ([]*core1_1.QueueFamilyProperties2, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outDataCountPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	p.InstanceDriver.VkGetPhysicalDeviceQueueFamilyProperties2(p.PhysicalDeviceHandle, outDataCountPtr, nil)

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

	p.InstanceDriver.VkGetPhysicalDeviceQueueFamilyProperties2(p.PhysicalDeviceHandle, outDataCountPtr, (*driver.VkQueueFamilyProperties2)(unsafe.Pointer(outData)))

	err = common.PopulateOutDataSlice[C.VkQueueFamilyProperties2, *core1_1.QueueFamilyProperties2](out, unsafe.Pointer(outData))
	return out, err
}

func (p *VulkanPhysicalDevice) SparseImageFormatProperties2(o core1_1.PhysicalDeviceSparseImageFormatInfo2, outDataFactory func() *core1_1.SparseImageFormatProperties2) ([]*core1_1.SparseImageFormatProperties2, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outDataCountPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	optionData, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, err
	}

	p.InstanceDriver.VkGetPhysicalDeviceSparseImageFormatProperties2(p.PhysicalDeviceHandle, (*driver.VkPhysicalDeviceSparseImageFormatInfo2)(optionData), outDataCountPtr, nil)

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

	p.InstanceDriver.VkGetPhysicalDeviceSparseImageFormatProperties2(p.PhysicalDeviceHandle,
		(*driver.VkPhysicalDeviceSparseImageFormatInfo2)(optionData),
		outDataCountPtr,
		(*driver.VkSparseImageFormatProperties2)(unsafe.Pointer(outData)),
	)

	err = common.PopulateOutDataSlice[C.VkSparseImageFormatProperties2, *core1_1.SparseImageFormatProperties2](out, unsafe.Pointer(outData))

	return out, err
}
