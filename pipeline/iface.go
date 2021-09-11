package pipeline

import "github.com/CannibalVox/VKng/core/loader"

type Pipeline interface {
	Handle() loader.VkPipeline
	Destroy() error
}

type PipelineLayout interface {
	Handle() loader.VkPipelineLayout
	Destroy() error
}
