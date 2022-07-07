package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

type VulkanShaderModule struct {
	core1_1.ShaderModule
}

func PromoteShaderModule(shaderModule core1_0.ShaderModule) ShaderModule {
	if !shaderModule.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedShaderModule := core1_1.PromoteShaderModule(shaderModule)
	return shaderModule.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(shaderModule.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanShaderModule{promotedShaderModule}
		}).(ShaderModule)
}
