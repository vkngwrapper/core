package commands

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/VKng/core/pipeline"
	"github.com/CannibalVox/VKng/core/resources"
	"github.com/CannibalVox/cgoalloc"
)

//go:generate mockgen -source iface.go -destination ../mocks/commands.go -package=mocks

type CommandBuffer interface {
	Handle() loader.VkCommandBuffer
	Destroy() error

	Begin(allocator cgoalloc.Allocator, o *BeginOptions) (loader.VkResult, error)
	End() (loader.VkResult, error)

	CmdBeginRenderPass(allocator cgoalloc.Allocator, contents SubpassContents, o *RenderPassBeginOptions) error
	CmdEndRenderPass() error
	CmdBindPipeline(bindPoint core.PipelineBindPoint, pipeline pipeline.Pipeline) error
	CmdDraw(vertexCount, instanceCount int, firstVertex, firstInstance uint32) error
	CmdDrawIndexed(indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32) error
	CmdBindVertexBuffers(allocator cgoalloc.Allocator, firstBinding uint32, buffers []resources.Buffer, bufferOffsets []int) error
	CmdBindIndexBuffer(buffer resources.Buffer, offset int, indexType core.IndexType) error
	CmdCopyBuffer(allocator cgoalloc.Allocator, srcBuffer resources.Buffer, dstBuffer resources.Buffer, copyRegions []BufferCopy) error
}

type CommandPool interface {
	Handle() loader.VkCommandPool
	Destroy() error
	DestroyBuffers(allocator cgoalloc.Allocator, buffers []CommandBuffer) error
}
