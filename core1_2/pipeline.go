package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanPipeline struct {
	core1_1.Pipeline
}

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
