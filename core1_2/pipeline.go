package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

// VulkanPipeline is an implementation of the Pipeline interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanPipeline struct {
	core1_1.Pipeline
}

// PromotePipeline accepts a Pipeline object from any core version. If provided a pipeline that supports
// at least core 1.2, it will return a core1_2.Pipeline. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanPipeline, even if it is provided a VulkanPipeline from a higher
// core version. Two Vulkan 1.2 compatible Pipeline objects with the same Pipeline.Handle will
// return the same interface value when passed to this method.
func PromotePipeline(pipeline core1_0.Pipeline) Pipeline {
	if !pipeline.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedPipeline := core1_1.PromotePipeline(pipeline)
	return pipeline.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(pipeline.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanPipeline{promotedPipeline}
		}).(Pipeline)
}

func PromotePipelineSlice(pipelines []core1_0.Pipeline) []Pipeline {
	outPipelines := make([]Pipeline, len(pipelines))
	for i := 0; i < len(pipelines); i++ {
		outPipelines[i] = PromotePipeline(pipelines[i])
	}

	return outPipelines
}
