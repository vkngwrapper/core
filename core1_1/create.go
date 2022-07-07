package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/driver"
)

func CreateDescriptorUpdateTemplate(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorUpdateTemplate, version common.APIVersion) *VulkanDescriptorUpdateTemplate {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_1,
		func() any {
			template := &VulkanDescriptorUpdateTemplate{
				DeviceDriver:             coreDriver,
				Device:                   device,
				DescriptorTemplateHandle: handle,
				MaximumAPIVersion:        version,
			}

			return template
		}).(*VulkanDescriptorUpdateTemplate)
}

func CreateSamplerYcbcrConversion(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSamplerYcbcrConversion, version common.APIVersion) *VulkanSamplerYcbcrConversion {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_1,
		func() any {
			return &VulkanSamplerYcbcrConversion{
				DeviceDriver:      coreDriver,
				Device:            device,
				YcbcrHandle:       handle,
				MaximumAPIVersion: version,
			}
		}).(*VulkanSamplerYcbcrConversion)
}
