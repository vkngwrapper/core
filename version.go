package core

import "fmt"

type Version uint32

func CreateVersion(major, minor, patch uint32) Version {
	return Version((major << 22) | (minor << 12) | patch)
}

func (v Version) Major() uint32 {
	return uint32(v >> 22)
}

func (v Version) Minor() uint32 {
	return uint32((v >> 12) & 0x3FF)
}

func (v Version) Patch() uint32 {
	return uint32(v & 0xFFF)
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d.", v.Major(), v.Minor(), v.Patch())
}
