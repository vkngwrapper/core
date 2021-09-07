package render_pass

import (
	"github.com/CannibalVox/VKng/core"
)

type SubPass struct {
	BindPoint core.PipelineBindPoint

	InputAttachments           []core.AttachmentReference
	ColorAttachments           []core.AttachmentReference
	ResolveAttachments         []core.AttachmentReference
	DepthStencilAttachments    []core.AttachmentReference
	PreservedAttachmentIndices []int
}
