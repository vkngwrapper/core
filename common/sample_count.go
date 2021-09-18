package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
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

var sampleCountsToString = map[SampleCounts]string{
	Samples1:  "1 Samples",
	Samples2:  "2 Samples",
	Samples4:  "4 Samples",
	Samples8:  "8 Samples",
	Samples16: "16 Samples",
	Samples32: "32 Samples",
	Samples64: "64 Samples",
}

func (c SampleCounts) String() string {
	if c == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := SampleCounts(1 << i)
		if (c & checkBit) != 0 {
			str, hasStr := sampleCountsToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteRune('|')
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}

var sampleCountsToCount = map[SampleCounts]int{
	Samples1:  1,
	Samples2:  2,
	Samples4:  4,
	Samples8:  8,
	Samples16: 16,
	Samples32: 32,
	Samples64: 64,
}

func (c SampleCounts) Count() int {
	var outCount int
	for i := 0; i < 32; i++ {
		checkBit := SampleCounts(1 << i)
		if (c & checkBit) != 0 {
			count, hasCount := sampleCountsToCount[checkBit]
			if hasCount {
				outCount = count
			}
		}
	}

	return outCount
}
