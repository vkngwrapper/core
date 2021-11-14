package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type AttachmentStoreOp int32

const (
	StoreOpStore    AttachmentStoreOp = C.VK_ATTACHMENT_STORE_OP_STORE
	StoreOpDontCare AttachmentStoreOp = C.VK_ATTACHMENT_STORE_OP_DONT_CARE
	StoreOpNoneEXT  AttachmentStoreOp = C.VK_ATTACHMENT_STORE_OP_NONE_EXT
)

var attachmentStoreOpToString = map[AttachmentStoreOp]string{
	StoreOpStore:    "Store",
	StoreOpDontCare: "Don't Care",
	StoreOpNoneEXT:  "None (Extension)",
}

func (o AttachmentStoreOp) String() string {
	return attachmentStoreOpToString[o]
}
