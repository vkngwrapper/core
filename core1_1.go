package core

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/internal/objects"
	"github.com/CannibalVox/cgoparam"
)

type VulkanLoader1_1 struct {
	driver driver.Driver
}

func (l *VulkanLoader1_1) CreateDescriptorUpdateTemplate(device Device, o core1_1.DescriptorUpdateTemplateCreateOptions, allocator *driver.AllocationCallbacks) (DescriptorUpdateTemplate, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfoPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var templateHandle driver.VkDescriptorUpdateTemplate
	res, err := l.driver.VkCreateDescriptorUpdateTemplate(device.Handle(),
		(*driver.VkDescriptorUpdateTemplateCreateInfo)(createInfoPtr),
		allocator.Handle(),
		&templateHandle,
	)
	if err != nil {
		return nil, res, err
	}

	return objects.CreateDescriptorUpdateTemplate(device.Driver(), device.Handle(), templateHandle, device.APIVersion()), res, nil
}

func (l *VulkanLoader1_1) CreateSamplerYcbcrConversion(device Device, o core1_1.SamplerYcbcrConversionCreateOptions, allocator *driver.AllocationCallbacks) (SamplerYcbcrConversion, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var ycbcrHandle driver.VkSamplerYcbcrConversion
	res, err := l.driver.VkCreateSamplerYcbcrConversion(
		device.Handle(),
		(*driver.VkSamplerYcbcrConversionCreateInfo)(optionPtr),
		allocator.Handle(),
		&ycbcrHandle,
	)
	if err != nil {
		return nil, res, err
	}

	return objects.CreateSamplerYcbcrConversion(device.Driver(), device.Handle(), ycbcrHandle, device.APIVersion()), res, nil
}

func (l *VulkanLoader1_1) GetQueue(device Device, o core1_1.DeviceQueueOptions) (Queue, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, err
	}

	var queue driver.VkQueue
	l.driver.VkGetDeviceQueue2(
		device.Handle(),
		(*driver.VkDeviceQueueInfo2)(optionPtr),
		&queue,
	)

	return objects.CreateQueue(device.Driver(), device.Handle(), queue, device.APIVersion()), nil
}
