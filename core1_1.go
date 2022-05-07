package core

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	core_internal "github.com/CannibalVox/VKng/core/internal/core1_1"
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

	descriptorTemplate := device.Driver().ObjectStore().GetOrCreate(driver.VulkanHandle(templateHandle),
		func() interface{} {
			template := &core_internal.VulkanDescriptorUpdateTemplate{
				Driver:                   device.Driver(),
				Device:                   device.Handle(),
				DescriptorTemplateHandle: templateHandle,
				MaximumAPIVersion:        device.APIVersion(),
			}

			return template
		}).(*core_internal.VulkanDescriptorUpdateTemplate)

	return descriptorTemplate, res, nil
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

	ycbcr := device.Driver().ObjectStore().GetOrCreate(driver.VulkanHandle(ycbcrHandle),
		func() interface{} {
			return &core_internal.VulkanSamplerYcbcrConversion{
				Driver:            l.driver,
				Device:            device.Handle(),
				YcbcrHandle:       ycbcrHandle,
				MaximumAPIVersion: device.APIVersion(),
			}
		}).(*core_internal.VulkanSamplerYcbcrConversion)

	return ycbcr, res, nil
}
