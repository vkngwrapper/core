package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type AttachmentDescriptionFlags int32

const (
	AttachmentMayAlias AttachmentDescriptionFlags = C.VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT
)

var attachmentDescriptionFlagsToString = map[AttachmentDescriptionFlags]string{
	AttachmentMayAlias: "May Alias",
}

func (f AttachmentDescriptionFlags) String() string {
	return FlagsToString(f, attachmentDescriptionFlagsToString)
}
