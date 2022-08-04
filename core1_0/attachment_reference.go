package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"

// AttachmentUnused indicates that a render pass attachment is not used
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VK_ATTACHMENT_UNUSED.html
const AttachmentUnused int = C.VK_ATTACHMENT_UNUSED

// AttachmentReference specifies an attachment reference
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference.html
type AttachmentReference struct {
	// Attachment is either an integer value identifying an attachment at the corresponding
	// index or AttachmentUnused to signify that it is not used
	Attachment int
	// Layout specifies the layout the attachment uses during the subpass
	Layout ImageLayout
}
