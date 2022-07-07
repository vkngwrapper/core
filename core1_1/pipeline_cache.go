package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
)

type VulkanPipelineCache struct {
	core1_0.PipelineCache
}

func PromotePipelineCache(pipelineCache core1_0.PipelineCache) PipelineCache {
	if !pipelineCache.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return pipelineCache.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(pipelineCache.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanPipelineCache{pipelineCache}
		}).(PipelineCache)
}
