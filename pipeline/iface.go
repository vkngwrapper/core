package pipeline

import "github.com/CannibalVox/VKng/core/loader"

//go:generate mockgen -source iface.go -destination ../mocks/pipeline.go -package=mocks

type Pipeline interface {
	Handle() loader.VkPipeline
	Destroy() error
}

type PipelineLayout interface {
	Handle() loader.VkPipelineLayout
	Destroy() error
}
