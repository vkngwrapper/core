package core1_1

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
)

// SamplerYcbcrConversion is an opaque representation of a device-specific sampler YCbCr conversion
// description.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html
type SamplerYcbcrConversion struct {
	device      loader.VkDevice
	ycbcrHandle loader.VkSamplerYcbcrConversion

	apiVersion common.APIVersion
}

func (y SamplerYcbcrConversion) Handle() loader.VkSamplerYcbcrConversion {
	return y.ycbcrHandle
}

func (y SamplerYcbcrConversion) DeviceHandle() loader.VkDevice {
	return y.device
}

func (y SamplerYcbcrConversion) APIVersion() common.APIVersion {
	return y.apiVersion
}

func (y SamplerYcbcrConversion) Initialized() bool {
	return y.ycbcrHandle != 0
}

func InternalSamplerYcbcrConversion(device loader.VkDevice, handle loader.VkSamplerYcbcrConversion, version common.APIVersion) SamplerYcbcrConversion {
	return SamplerYcbcrConversion{
		device:      device,
		ycbcrHandle: handle,
		apiVersion:  version,
	}
}

// DescriptorUpdateTemplate specifies a mapping from descriptor update information in host memory to
// descriptors in a DescriptorSet
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html
type DescriptorUpdateTemplate struct {
	device                   loader.VkDevice
	descriptorTemplateHandle loader.VkDescriptorUpdateTemplate

	apiVersion common.APIVersion
}

func (t DescriptorUpdateTemplate) Handle() loader.VkDescriptorUpdateTemplate {
	return t.descriptorTemplateHandle
}

func (t DescriptorUpdateTemplate) DeviceHandle() loader.VkDevice {
	return t.device
}

func (t DescriptorUpdateTemplate) APIVersion() common.APIVersion {
	return t.apiVersion
}

func (t DescriptorUpdateTemplate) Initialized() bool {
	return t.descriptorTemplateHandle != 0
}

func InternalDescriptorUpdateTemplate(device loader.VkDevice, handle loader.VkDescriptorUpdateTemplate, version common.APIVersion) DescriptorUpdateTemplate {
	return DescriptorUpdateTemplate{
		device:                   device,
		descriptorTemplateHandle: handle,
		apiVersion:               version,
	}
}
