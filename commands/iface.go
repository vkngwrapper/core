package commands

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/VKng/core/pipeline"
	"github.com/CannibalVox/VKng/core/resources"
)

//go:generate mockgen -source iface.go -destination ../mocks/commands.go -package=mocks

type CommandBuffer interface {
	Handle() loader.VkCommandBuffer
	Destroy() error

	Begin(o *BeginOptions) (loader.VkResult, error)
	End() (loader.VkResult, error)

	CmdBeginRenderPass(contents SubpassContents, o *RenderPassBeginOptions) error
	CmdEndRenderPass() error
	CmdBindPipeline(bindPoint core.PipelineBindPoint, pipeline pipeline.Pipeline) error
	CmdDraw(vertexCount, instanceCount int, firstVertex, firstInstance uint32) error
	CmdDrawIndexed(indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32) error
	CmdBindVertexBuffers(firstBinding uint32, buffers []resources.Buffer, bufferOffsets []int) error
	CmdBindIndexBuffer(buffer resources.Buffer, offset int, indexType core.IndexType) error
	CmdCopyBuffer(srcBuffer resources.Buffer, dstBuffer resources.Buffer, copyRegions []BufferCopy) error
	CmdBindDescriptorSets(bindPoint core.PipelineBindPoint, layout pipeline.PipelineLayout, firstSet int, sets []resources.DescriptorSet, dynamicOffsets []int) error
}

type CommandPool interface {
	Handle() loader.VkCommandPool
	Destroy() error
	DestroyBuffers(buffers []CommandBuffer) error
}
