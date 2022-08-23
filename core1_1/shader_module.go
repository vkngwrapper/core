package core1_1

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanShaderModule is an implementation of the ShaderModule interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanShaderModule struct {
	core1_0.ShaderModule
}

// PromoteShaderModule accepts a ShaderModule object from any core version. If provided a shader module that supports
// at least core 1.1, it will return a core1_1.ShaderModule. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanShaderModule, even if it is provided a VulkanShaderModule from a higher
// core version. Two Vulkan 1.1 compatible ShaderModule objects with the same ShaderModule.Handle will
// return the same interface value when passed to this method.
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
