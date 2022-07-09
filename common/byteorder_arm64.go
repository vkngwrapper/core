package common

import "encoding/binary"

// ByteOrder provides binary.LittleEndian or binary.BigEndian, depending on the current platform
var ByteOrder = binary.LittleEndian
