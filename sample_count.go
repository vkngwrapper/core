package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "strings"

type SampleCounts int32

const (
	Samples1  = C.VK_SAMPLE_COUNT_1_BIT
	Samples2  = C.VK_SAMPLE_COUNT_2_BIT
	Samples4  = C.VK_SAMPLE_COUNT_4_BIT
	Samples8  = C.VK_SAMPLE_COUNT_8_BIT
	Samples16 = C.VK_SAMPLE_COUNT_16_BIT
	Samples32 = C.VK_SAMPLE_COUNT_32_BIT
	Samples64 = C.VK_SAMPLE_COUNT_64_BIT
)

func (c SampleCounts) String() string {
	hasOne := false
	var sb strings.Builder

	if c&Samples1 != 0 {
		sb.WriteString("1 Sample")
		hasOne = true
	}

	if c&Samples2 != 0 {
		if hasOne {
			sb.WriteString("|")
		}

		sb.WriteString("2 Samples")
		hasOne = true
	}

	if c&Samples4 != 0 {
		if hasOne {
			sb.WriteString("|")
		}

		sb.WriteString("4 Samples")
		hasOne = true
	}

	if c&Samples8 != 0 {
		if hasOne {
			sb.WriteString("|")
		}

		sb.WriteString("8 Samples")
		hasOne = true
	}

	if c&Samples16 != 0 {
		if hasOne {
			sb.WriteString("|")
		}

		sb.WriteString("16 Samples")
		hasOne = true
	}

	if c&Samples32 != 0 {
		if hasOne {
			sb.WriteString("|")
		}

		sb.WriteString("32 Samples")
		hasOne = true
	}

	if c&Samples64 != 0 {
		if hasOne {
			sb.WriteString("|")
		}

		sb.WriteString("64 Samples")
	}

	return sb.String()
}
