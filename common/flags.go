package common

import "strings"

type flags interface {
	~int32 | ~uint32
}

// FlagStringMapping is used as a base type for many bitflag enums in vkngwrapper.
// It has the capability to Register flags with a descriptive string in an init() method.
// Once that has done, the flag type can be stringified into a pipe-separated list of
// flags.
type FlagStringMapping[T flags] struct {
	stringValues map[T]string
}

// NewFlagStringMapping creates a FlagStringMapping for use in flag types
func NewFlagStringMapping[T flags]() FlagStringMapping[T] {
	return FlagStringMapping[T]{make(map[T]string)}
}

// Register maps a flag value to a string that will represent that flag
// when using String(). It is best to call this from init()
func (m FlagStringMapping[T]) Register(value T, str string) {
	m.stringValues[value] = str
}

// FlagsToString returns a formatted string representing the flag value
// passed in. It is a pipe-separated list of descriptive strings for each
// flag active in the value.
func (m FlagStringMapping[T]) FlagsToString(value T) string {
	if value == 0 {
		return "None"
	}

	hasOne := false
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		shiftedBit := T(1 << i)
		if value&shiftedBit != 0 {
			strVal, exists := m.stringValues[shiftedBit]
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
