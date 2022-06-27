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

type VulkanDevice struct {
	core1_0.Device

	DeviceDriver      driver.Driver
	DeviceHandle      driver.VkDevice
	MaximumAPIVersion common.APIVersion
}

func PromoteDevice(device core1_0.Device) Device {
	if !device.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return device.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(device.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanDevice{
				Device:            device,
				DeviceDriver:      device.Driver(),
				DeviceHandle:      device.Handle(),
				MaximumAPIVersion: device.APIVersion(),
			}
		}).(Device)
}

func (d *VulkanDevice) BindBufferMemory(o []BindBufferMemoryOptions) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptionSlice[C.VkBindBufferMemoryInfo, BindBufferMemoryOptions](arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return d.DeviceDriver.VkBindBufferMemory2(d.DeviceHandle,
		driver.Uint32(len(o)),
		(*driver.VkBindBufferMemoryInfo)(unsafe.Pointer(optionPtr)),
	)
}

func (d *VulkanDevice) BindImageMemory(o []BindImageMemoryOptions) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptionSlice[C.VkBindImageMemoryInfo, BindImageMemoryOptions](arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return d.DeviceDriver.VkBindImageMemory2(
		d.DeviceHandle,
		driver.Uint32(len(o)),
		(*driver.VkBindImageMemoryInfo)(unsafe.Pointer(optionPtr)),
	)
}

func (d *VulkanDevice) BufferMemoryRequirements(o BufferMemoryRequirementsOptions, out *MemoryRequirementsOutData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOptions(arena, out)
	if err != nil {
		return err
	}

	d.DeviceDriver.VkGetBufferMemoryRequirements2(d.DeviceHandle,
		(*driver.VkBufferMemoryRequirementsInfo2)(optionPtr),
		(*driver.VkMemoryRequirements2)(outDataPtr),
	)

	return common.PopulateOutData(out, outDataPtr)
}

func (d *VulkanDevice) ImageMemoryRequirements(o ImageMemoryRequirementsOptions, out *MemoryRequirementsOutData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOptions(arena, out)
	if err != nil {
		return err
	}

	d.DeviceDriver.VkGetImageMemoryRequirements2(d.DeviceHandle,
		(*driver.VkImageMemoryRequirementsInfo2)(optionPtr),
		(*driver.VkMemoryRequirements2)(outDataPtr),
	)

	return common.PopulateOutData(out, outDataPtr)
}

func (d *VulkanDevice) SparseImageMemoryRequirements(o ImageSparseMemoryRequirementsOptions, outDataFactory func() *SparseImageMemoryRequirementsOutData) ([]*SparseImageMemoryRequirementsOutData, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, err
	}

	requirementCountPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	d.DeviceDriver.VkGetImageSparseMemoryRequirements2(d.DeviceHandle,
		(*driver.VkImageSparseMemoryRequirementsInfo2)(optionPtr),
		requirementCountPtr,
		nil,
	)

	count := int(*requirementCountPtr)
	if count == 0 {
		return nil, nil
	}

	outDataSlice := make([]*SparseImageMemoryRequirementsOutData, count)
	for i := 0; i < count; i++ {
		if outDataFactory != nil {
			outDataSlice[i] = outDataFactory()
		} else {
			outDataSlice[i] = &SparseImageMemoryRequirementsOutData{}
		}
	}

	outDataPtr, err := common.AllocOptionSlice[C.VkSparseImageMemoryRequirements2, *SparseImageMemoryRequirementsOutData](arena, outDataSlice)
	if err != nil {
		return nil, err
	}

	castOutDataPtr := (*C.VkSparseImageMemoryRequirements2)(outDataPtr)

	d.DeviceDriver.VkGetImageSparseMemoryRequirements2(d.DeviceHandle,
		(*driver.VkImageSparseMemoryRequirementsInfo2)(optionPtr),
		requirementCountPtr,
		(*driver.VkSparseImageMemoryRequirements2)(unsafe.Pointer(castOutDataPtr)),
	)

	err = common.PopulateOutDataSlice[C.VkSparseImageMemoryRequirements2, *SparseImageMemoryRequirementsOutData](outDataSlice, unsafe.Pointer(outDataPtr))
	if err != nil {
		return nil, err
	}

	return outDataSlice, nil
}

func (d *VulkanDevice) DescriptorSetLayoutSupport(o core1_0.DescriptorSetLayoutCreateOptions, outData *DescriptorSetLayoutSupportOutData) error {
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

	d.DeviceDriver.VkGetDescriptorSetLayoutSupport(d.DeviceHandle,
		(*driver.VkDescriptorSetLayoutCreateInfo)(optionsPtr),
		(*driver.VkDescriptorSetLayoutSupport)(outDataPtr))

	return common.PopulateOutData(outData, outDataPtr)
}

func (d *VulkanDevice) DeviceGroupPeerMemoryFeatures(heapIndex, localDeviceIndex, remoteDeviceIndex int) PeerMemoryFeatures {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	featuresPtr := (*driver.VkPeerMemoryFeatureFlags)(arena.Malloc(int(unsafe.Sizeof(C.VkPeerMemoryFeatureFlags(0)))))

	d.DeviceDriver.VkGetDeviceGroupPeerMemoryFeatures(
		d.DeviceHandle,
		driver.Uint32(heapIndex),
		driver.Uint32(localDeviceIndex),
		driver.Uint32(remoteDeviceIndex),
		featuresPtr,
	)

	return PeerMemoryFeatures(*featuresPtr)
}

func (d *VulkanDevice) CreateDescriptorUpdateTemplate(o DescriptorUpdateTemplateCreateOptions, allocator *driver.AllocationCallbacks) (DescriptorUpdateTemplate, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfoPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var templateHandle driver.VkDescriptorUpdateTemplate
	res, err := d.DeviceDriver.VkCreateDescriptorUpdateTemplate(d.DeviceHandle,
		(*driver.VkDescriptorUpdateTemplateCreateInfo)(createInfoPtr),
		allocator.Handle(),
		&templateHandle,
	)
	if err != nil {
		return nil, res, err
	}

	return CreateDescriptorUpdateTemplate(d.DeviceDriver, d.DeviceHandle, templateHandle, d.MaximumAPIVersion), res, nil
}

func (d *VulkanDevice) CreateSamplerYcbcrConversion(o SamplerYcbcrConversionCreateOptions, allocator *driver.AllocationCallbacks) (SamplerYcbcrConversion, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var ycbcrHandle driver.VkSamplerYcbcrConversion
	res, err := d.DeviceDriver.VkCreateSamplerYcbcrConversion(
		d.DeviceHandle,
		(*driver.VkSamplerYcbcrConversionCreateInfo)(optionPtr),
		allocator.Handle(),
		&ycbcrHandle,
	)
	if err != nil {
		return nil, res, err
	}

	return CreateSamplerYcbcrConversion(d.DeviceDriver, d.DeviceHandle, ycbcrHandle, d.MaximumAPIVersion), res, nil
}

//go:linkname createQueueObject github.com/CannibalVox/VKng/core/core1_0.createQueueObject
func createQueueObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueue, version common.APIVersion) *core1_0.VulkanQueue

func (d *VulkanDevice) GetQueue2(o DeviceQueueOptions) (core1_0.Queue, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, err
	}

	var queue driver.VkQueue
	d.DeviceDriver.VkGetDeviceQueue2(
		d.DeviceHandle,
		(*driver.VkDeviceQueueInfo2)(optionPtr),
		&queue,
	)

	return createQueueObject(d.DeviceDriver, d.DeviceHandle, queue, d.MaximumAPIVersion), nil
}
