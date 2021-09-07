package render_pass

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
)

const SubpassExternal = int(C.VK_SUBPASS_EXTERNAL)

type SubPassDependency struct {
	Flags core.DependencyFlags

	SrcSubPassIndex int
	DstSubPassIndex int

	SrcStageMask core.PipelineStages
	DstStageMask core.PipelineStages

	SrcAccess core.AccessFlags
	DstAccess core.AccessFlags
}
