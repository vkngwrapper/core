package core

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "strings"

type AttachmentDescriptionFlags int32

const (
	AttachmentMayAlias AttachmentDescriptionFlags = C.VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT
)

var attachmentDescriptionFlagsToString = map[AttachmentDescriptionFlags]string{
	AttachmentMayAlias: "May Alias",
}

func (f AttachmentDescriptionFlags) String() string {
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder
	for i := 0; i < 32; i++ {
		checkBit := AttachmentDescriptionFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := attachmentDescriptionFlagsToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteString("|")
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}
