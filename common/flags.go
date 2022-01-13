package common

import "strings"

type flags interface {
	~int32 | ~uint32
}

func FlagsToString[T flags](value T, stringValues map[T]string) string {
	if value == 0 {
		return "None"
	}

	hasOne := false
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		shiftedBit := T(1 << i)
		if value&shiftedBit != 0 {
			strVal, exists := stringValues[shiftedBit]
			if exists {
				if hasOne {
					sb.WriteString("|")
				}
				sb.WriteString(strVal)
				hasOne = true
			}
		}
	}

	return sb.String()
}
