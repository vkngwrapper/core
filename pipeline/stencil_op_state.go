package pipeline

import (
	"github.com/CannibalVox/VKng/core"
)

type StencilOpState struct {
	FailOp      core.StencilOp
	PassOp      core.StencilOp
	DepthFailOp core.StencilOp

	CompareOp   core.CompareOp
	CompareMask uint32
	WriteMask   uint32

	Reference uint32
}
