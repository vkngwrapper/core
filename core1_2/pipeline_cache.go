package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanPipelineCache struct {
	core1_1.PipelineCache
}

func PromotePipelineCache(pipelineCache core1_0.PipelineCache) PipelineCache {
	if !pipelineCache.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedPipelineCache := core1_1.PromotePipelineCache(pipelineCache)
	return pipelineCache.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(pipelineCache.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanPipelineCache{promotedPipelineCache}
		}).(PipelineCache)
}
