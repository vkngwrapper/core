package internal1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanPipeline struct {
	core1_0.Pipeline
}

func PromotePipeline(pipeline core1_0.Pipeline) core1_1.Pipeline {
	if !pipeline.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return pipeline.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(pipeline.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanPipeline{pipeline}
		}).(core1_1.Pipeline)
}

func PromotePipelineSlice(pipelines []core1_0.Pipeline) []core1_1.Pipeline {
	outPipelines := make([]core1_1.Pipeline, len(pipelines))
	for i := 0; i < len(pipelines); i++ {
		outPipelines[i] = PromotePipeline(pipelines[i])
	}

	return outPipelines
}
