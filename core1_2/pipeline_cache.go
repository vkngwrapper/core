package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

// VulkanPipelineCache is an implementation of the PipelineCache interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanPipelineCache struct {
	core1_1.PipelineCache
}

// PromotePipelineCache accepts a PipelineCache object from any core version. If provided a pipeline cache that supports
// at least core 1.2, it will return a core1_2.PipelineCache. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanPipelineCache, even if it is provided a VulkanPipelineCache from a higher
// core version. Two Vulkan 1.2 compatible PipelineCache objects with the same PipelineCache.Handle will
// return the same interface value when passed to this method.
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
