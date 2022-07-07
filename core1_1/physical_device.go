package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
	"unsafe"
)

type VulkanPhysicalDevice struct {
	core1_0.PhysicalDevice

	InstanceScoped1_1 InstanceScopedPhysicalDevice
}

func (p *VulkanPhysicalDevice) InstanceScopedPhysicalDevice1_1() InstanceScopedPhysicalDevice {
	return p.InstanceScoped1_1
}

func PromotePhysicalDevice(physicalDevice core1_0.PhysicalDevice) PhysicalDevice {
	if !physicalDevice.DeviceAPIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	instanceScoped := PromoteInstanceScopedPhysicalDevice(physicalDevice)

	return physicalDevice.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(physicalDevice.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanPhysicalDevice{
				PhysicalDevice: physicalDevice,

				InstanceScoped1_1: instanceScoped,
			}
		}).(PhysicalDevice)
}

type VulkanInstanceScopedPhysicalDevice struct {
	core1_0.PhysicalDevice

	InstanceDriver       driver.Driver
	PhysicalDeviceHandle driver.VkPhysicalDevice
}

func PromoteInstanceScopedPhysicalDevice(physicalDevice core1_0.PhysicalDevice) InstanceScopedPhysicalDevice {
	if !physicalDevice.InstanceAPIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return physicalDevice.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(physicalDevice.Handle()),
		driver.Core1_1InstanceScope,
		func() any {
			return &VulkanInstanceScopedPhysicalDevice{
				PhysicalDevice: physicalDevice,

				InstanceDriver:       physicalDevice.Driver(),
				PhysicalDeviceHandle: physicalDevice.Handle(),
			}
		}).(InstanceScopedPhysicalDevice)
}

func (p *VulkanInstanceScopedPhysicalDevice) ExternalFenceProperties(o PhysicalDeviceExternalFenceInfo, outData *ExternalFenceProperties) error {
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

func (p *VulkanInstanceScopedPhysicalDevice) ExternalBufferProperties(o PhysicalDeviceExternalBufferInfo, outData *ExternalBufferProperties) error {
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

func (p *VulkanInstanceScopedPhysicalDevice) ExternalSemaphoreProperties(o PhysicalDeviceExternalSemaphoreInfo, outData *ExternalSemaphoreProperties) error {
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

func (p *VulkanInstanceScopedPhysicalDevice) Features2(out *PhysicalDeviceFeatures2) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	p.InstanceDriver.VkGetPhysicalDeviceFeatures2(p.PhysicalDeviceHandle, (*driver.VkPhysicalDeviceFeatures2)(outData))

	return common.PopulateOutData(out, outData)
}

func (p *VulkanInstanceScopedPhysicalDevice) FormatProperties2(format core1_0.Format, out *FormatProperties2) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	p.InstanceDriver.VkGetPhysicalDeviceFormatProperties2(p.PhysicalDeviceHandle, driver.VkFormat(format), (*driver.VkFormatProperties2)(outData))

	return common.PopulateOutData(out, outData)
}

func (p *VulkanInstanceScopedPhysicalDevice) ImageFormatProperties2(o PhysicalDeviceImageFormatInfo2, out *ImageFormatProperties2) (common.VkResult, error) {
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

func (p *VulkanInstanceScopedPhysicalDevice) MemoryProperties2(out *PhysicalDeviceMemoryProperties2) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	p.InstanceDriver.VkGetPhysicalDeviceMemoryProperties2(p.PhysicalDeviceHandle, (*driver.VkPhysicalDeviceMemoryProperties2)(outData))

	return common.PopulateOutData(out, outData)
}

func (p *VulkanInstanceScopedPhysicalDevice) Properties2(out *PhysicalDeviceProperties2) error {
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

func (p *VulkanInstanceScopedPhysicalDevice) QueueFamilyProperties2(outDataFactory func() *QueueFamilyProperties2) ([]*QueueFamilyProperties2, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outDataCountPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	p.InstanceDriver.VkGetPhysicalDeviceQueueFamilyProperties2(p.PhysicalDeviceHandle, outDataCountPtr, nil)

	outDataCount := int(*outDataCountPtr)
	if outDataCount == 0 {
		return nil, nil
	}

	out := make([]*QueueFamilyProperties2, outDataCount)
	for i := 0; i < outDataCount; i++ {
		if outDataFactory != nil {
			out[i] = outDataFactory()
		} else {
			out[i] = &QueueFamilyProperties2{}
		}
	}

	outData, err := common.AllocOutDataHeaderSlice[C.VkQueueFamilyProperties2, *QueueFamilyProperties2](arena, out)
	if err != nil {
		return nil, err
	}

	p.InstanceDriver.VkGetPhysicalDeviceQueueFamilyProperties2(p.PhysicalDeviceHandle, outDataCountPtr, (*driver.VkQueueFamilyProperties2)(unsafe.Pointer(outData)))

	err = common.PopulateOutDataSlice[C.VkQueueFamilyProperties2, *QueueFamilyProperties2](out, unsafe.Pointer(outData))
	return out, err
}

func (p *VulkanInstanceScopedPhysicalDevice) SparseImageFormatProperties2(o PhysicalDeviceSparseImageFormatInfo2, outDataFactory func() *SparseImageFormatProperties2) ([]*SparseImageFormatProperties2, error) {
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

	out := make([]*SparseImageFormatProperties2, outDataCount)
	for i := 0; i < outDataCount; i++ {
		if outDataFactory != nil {
			out[i] = outDataFactory()
		} else {
			out[i] = &SparseImageFormatProperties2{}
		}
	}

	outData, err := common.AllocOutDataHeaderSlice[C.VkSparseImageFormatProperties2, *SparseImageFormatProperties2](arena, out)
	if err != nil {
		return nil, err
	}

	p.InstanceDriver.VkGetPhysicalDeviceSparseImageFormatProperties2(p.PhysicalDeviceHandle,
		(*driver.VkPhysicalDeviceSparseImageFormatInfo2)(optionData),
		outDataCountPtr,
		(*driver.VkSparseImageFormatProperties2)(unsafe.Pointer(outData)),
	)

	err = common.PopulateOutDataSlice[C.VkSparseImageFormatProperties2, *SparseImageFormatProperties2](out, unsafe.Pointer(outData))

	return out, err
}
