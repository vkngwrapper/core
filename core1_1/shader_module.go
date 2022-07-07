package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
)

type VulkanShaderModule struct {
	core1_0.ShaderModule
}

func PromoteShaderModule(shaderModule core1_0.ShaderModule) ShaderModule {
	if !shaderModule.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return shaderModule.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(shaderModule.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanShaderModule{shaderModule}
		}).(ShaderModule)
}
