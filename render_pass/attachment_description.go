package render_pass

import (
	"github.com/CannibalVox/VKng/core"
)

type AttachmentDescription struct {
	Flags   core.AttachmentDescriptionFlags
	Format  core.DataFormat
	Samples core.SampleCounts

	LoadOp         core.AttachmentLoadOp
	StoreOp        core.AttachmentStoreOp
	StencilLoadOp  core.AttachmentLoadOp
	StencilStoreOp core.AttachmentStoreOp

	InitialLayout core.ImageLayout
	FinalLayout   core.ImageLayout
}
