package common

import "fmt"

// Version is an integer intended to represent a semantic version, with
// major, minor and patch versions, like 1.2.5
type Version uint32

// CreateVersion generates a Version value for the given semantic version
func CreateVersion(major, minor, patch uint32) Version {
	return Version((major << 22) | (minor << 12) | patch)
}

// Major retrieves the major version value for a Version. In version 1.2.5,
// 1 is the major version.
func (v Version) Major() uint32 {
	return uint32(v >> 22)
}

// Minor retrieves the minor version value for a Version. In version 1.2.5,
// 2 is the minor version.
func (v Version) Minor() uint32 {
	return uint32((v >> 12) & 0x3FF)
}

// Patch retrieves the patch version value for a Version. In version 1.2.5,
// 5 is the patch version.
func (v Version) Patch() uint32 {
	return uint32(v & 0xFFF)
}

// String generates an aesthetically-pleasing semantic version tag for a Version.
func (v Version) String() string {
	return fmt.Sprintf("v%d.%d.%d.", v.Major(), v.Minor(), v.Patch())
}
