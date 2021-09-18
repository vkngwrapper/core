package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type AttachmentLoadOp int32

const (
	LoadOpLoad     AttachmentLoadOp = C.VK_ATTACHMENT_LOAD_OP_LOAD
	LoadOpClear    AttachmentLoadOp = C.VK_ATTACHMENT_LOAD_OP_CLEAR
	LoadOpDontCare AttachmentLoadOp = C.VK_ATTACHMENT_LOAD_OP_DONT_CARE
	LoadOpNone     AttachmentLoadOp = C.VK_ATTACHMENT_LOAD_OP_NONE_EXT
)

var attachmentLoadOpToString = map[AttachmentLoadOp]string{
	LoadOpLoad:     "Load",
	LoadOpClear:    "Clear",
	LoadOpDontCare: "Don't Care",
	LoadOpNone:     "None",
}

func (o AttachmentLoadOp) String() string {
	return attachmentLoadOpToString[o]
}
