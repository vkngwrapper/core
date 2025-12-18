package core1_1

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanPipeline is an implementation of the Pipeline interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanPipeline struct {
	core1_0.Pipeline
}

// PromotePipeline accepts a Pipeline object from any core version. If provided a pipeline that supports
// at least core 1.1, it will return a core1_1.Pipeline. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanPipeline, even if it is provided a VulkanPipeline from a higher
// core version. Two Vulkan 1.1 compatible Pipeline objects with the same Pipeline.Handle will
// return the same interface value when passed to this method.
func PromotePipeline(pipeline core1_0.Pipeline) Pipeline {
	if pipeline == nil {
		return nil
	}
	if !pipeline.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	promoted, alreadyPromoted := pipeline.(Pipeline)
	if alreadyPromoted {
		return promoted
	}

	return pipeline.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(pipeline.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanPipeline{pipeline}
		}).(Pipeline)
}

// PromotePipelineSlice accepts a slice of Pipeline objects from any core version.
// If provided a pipeline that supports at least core 1.1, it will return a core1_1.Pipeline.
// Otherwise, it will left out of the returned slice. This method will always return a
// core1_1.VulkanPipeline, even if it is provided a VulkanPipeline from a higher core version. Two
// Vulkan 1.1 compatible Pipeline objects with the same Pipeline.Handle will return the same interface
// value when passed to this method.
func PromotePipelineSlice(pipelines []core1_0.Pipeline) []Pipeline {
	for i := 0; i < len(pipelines); i++ {
		if pipelines[i].APIVersion() < common.Vulkan1_1 {
			return nil
		}
	}

	outPipelines := make([]Pipeline, len(pipelines))
	for i := 0; i < len(pipelines); i++ {
		outPipelines[i] = PromotePipeline(pipelines[i])
	}

	return outPipelines
}
