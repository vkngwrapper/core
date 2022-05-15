package internal1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanShaderModule struct {
	core1_0.ShaderModule
}

func PromoteShaderModule(shaderModule core1_0.ShaderModule) core1_1.ShaderModule {
	if !shaderModule.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return shaderModule.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(shaderModule.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanShaderModule{shaderModule}
		}).(core1_1.ShaderModule)
}
