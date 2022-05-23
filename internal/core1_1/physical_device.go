package internal1_1

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

type VulkanPhysicalDevice struct {
	core1_0.PhysicalDevice

	InstanceScoped1_1 core1_1.InstanceScopedPhysicalDevice
}

func (p *VulkanPhysicalDevice) InstanceScopedPhysicalDevice1_1() core1_1.InstanceScopedPhysicalDevice {
	return p.InstanceScoped1_1
}

func PromotePhysicalDevice(physicalDevice core1_0.PhysicalDevice) core1_1.PhysicalDevice {
	if !physicalDevice.DeviceAPIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	instanceScoped := PromoteInstanceScopePhysicalDevice(physicalDevice)

	return physicalDevice.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(physicalDevice.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanPhysicalDevice{
				PhysicalDevice: physicalDevice,

				InstanceScoped1_1: instanceScoped,
			}
		}).(core1_1.PhysicalDevice)
}

type VulkanInstanceScopedPhysicalDevice struct {
	core1_0.PhysicalDevice

	InstanceDriver       driver.Driver
	PhysicalDeviceHandle driver.VkPhysicalDevice
}

func PromoteInstanceScopePhysicalDevice(physicalDevice core1_0.PhysicalDevice) core1_1.InstanceScopedPhysicalDevice {
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
		}).(core1_1.InstanceScopedPhysicalDevice)
}

func (p *VulkanInstanceScopedPhysicalDevice) ExternalFenceProperties(o core1_1.ExternalFenceOptions, outData *core1_1.ExternalFenceOutData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOptions(arena, outData)
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

func (p *VulkanInstanceScopedPhysicalDevice) ExternalBufferProperties(o core1_1.ExternalBufferOptions, outData *core1_1.ExternalBufferOutData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionsPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOptions(arena, outData)
	if err != nil {
		return err
	}

	p.InstanceDriver.VkGetPhysicalDeviceExternalBufferProperties(p.PhysicalDeviceHandle,
		(*driver.VkPhysicalDeviceExternalBufferInfo)(optionsPtr),
		(*driver.VkExternalBufferProperties)(outDataPtr))

	return common.PopulateOutData(outData, outDataPtr)
}

func (p *VulkanInstanceScopedPhysicalDevice) ExternalSemaphoreProperties(o core1_1.ExternalSemaphoreOptions, outData *core1_1.ExternalSemaphoreOutData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOptions(arena, outData)
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

func (p *VulkanInstanceScopedPhysicalDevice) Features2(out *core1_1.DeviceFeaturesOutData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOptions(arena, out)
	if err != nil {
		return err
	}

	p.InstanceDriver.VkGetPhysicalDeviceFeatures2(p.PhysicalDeviceHandle, (*driver.VkPhysicalDeviceFeatures2)(outData))

	return common.PopulateOutData(out, outData)
}

func (p *VulkanInstanceScopedPhysicalDevice) FormatProperties2(format common.DataFormat, out *core1_1.FormatPropertiesOutData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOptions(arena, out)
	if err != nil {
		return err
	}

	p.InstanceDriver.VkGetPhysicalDeviceFormatProperties2(p.PhysicalDeviceHandle, driver.VkFormat(format), (*driver.VkFormatProperties2)(outData))

	return common.PopulateOutData(out, outData)
}

func (p *VulkanInstanceScopedPhysicalDevice) ImageFormatProperties2(o core1_1.ImageFormatOptions, out *core1_1.ImageFormatPropertiesOutData) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionData, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	outData, err := common.AllocOptions(arena, out)
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

func (p *VulkanInstanceScopedPhysicalDevice) MemoryProperties2(out *core1_1.MemoryPropertiesOutData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOptions(arena, out)
	if err != nil {
		return err
	}

	p.InstanceDriver.VkGetPhysicalDeviceMemoryProperties2(p.PhysicalDeviceHandle, (*driver.VkPhysicalDeviceMemoryProperties2)(outData))

	return common.PopulateOutData(out, outData)
}

func (p *VulkanInstanceScopedPhysicalDevice) Properties2(out *core1_1.DevicePropertiesOutData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOptions(arena, out)
	if err != nil {
		return err
	}

	p.InstanceDriver.VkGetPhysicalDeviceProperties2(p.PhysicalDeviceHandle,
		(*driver.VkPhysicalDeviceProperties2)(outData))

	return common.PopulateOutData(out, outData)
}

func (p *VulkanInstanceScopedPhysicalDevice) QueueFamilyProperties2(outDataFactory func() *core1_1.QueueFamilyOutData) ([]*core1_1.QueueFamilyOutData, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outDataCountPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	p.InstanceDriver.VkGetPhysicalDeviceQueueFamilyProperties2(p.PhysicalDeviceHandle, outDataCountPtr, nil)

	outDataCount := int(*outDataCountPtr)
	if outDataCount == 0 {
		return nil, nil
	}

	out := make([]*core1_1.QueueFamilyOutData, outDataCount)
	for i := 0; i < outDataCount; i++ {
		if outDataFactory != nil {
			out[i] = outDataFactory()
		} else {
			out[i] = &core1_1.QueueFamilyOutData{}
		}
	}

	outData, err := common.AllocOptionSlice[C.VkQueueFamilyProperties2, *core1_1.QueueFamilyOutData](arena, out)
	if err != nil {
		return nil, err
	}

	p.InstanceDriver.VkGetPhysicalDeviceQueueFamilyProperties2(p.PhysicalDeviceHandle, outDataCountPtr, (*driver.VkQueueFamilyProperties2)(unsafe.Pointer(outData)))

	err = common.PopulateOutDataSlice[C.VkQueueFamilyProperties2, *core1_1.QueueFamilyOutData](out, unsafe.Pointer(outData))
	return out, err
}

func (p *VulkanInstanceScopedPhysicalDevice) SparseImageFormatProperties2(o core1_1.SparseImageFormatOptions, outDataFactory func() *core1_1.SparseImageFormatPropertiesOutData) ([]*core1_1.SparseImageFormatPropertiesOutData, error) {
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

	out := make([]*core1_1.SparseImageFormatPropertiesOutData, outDataCount)
	for i := 0; i < outDataCount; i++ {
		if outDataFactory != nil {
			out[i] = outDataFactory()
		} else {
			out[i] = &core1_1.SparseImageFormatPropertiesOutData{}
		}
	}

	outData, err := common.AllocOptionSlice[C.VkSparseImageFormatProperties2, *core1_1.SparseImageFormatPropertiesOutData](arena, out)
	if err != nil {
		return nil, err
	}

	p.InstanceDriver.VkGetPhysicalDeviceSparseImageFormatProperties2(p.PhysicalDeviceHandle,
		(*driver.VkPhysicalDeviceSparseImageFormatInfo2)(optionData),
		outDataCountPtr,
		(*driver.VkSparseImageFormatProperties2)(unsafe.Pointer(outData)),
	)

	err = common.PopulateOutDataSlice[C.VkSparseImageFormatProperties2, *core1_1.SparseImageFormatPropertiesOutData](out, unsafe.Pointer(outData))

	return out, err
}
