package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type SubpassContents int32

const (
	ContentsInline                  SubpassContents = C.VK_SUBPASS_CONTENTS_INLINE
	ContentsSecondaryCommandBuffers SubpassContents = C.VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS
)

var subpassContentsToString = map[SubpassContents]string{
	ContentsInline:                  "Inline",
	ContentsSecondaryCommandBuffers: "Secondary Command Buffers",
}

func (c SubpassContents) String() string {
	return subpassContentsToString[c]
}
