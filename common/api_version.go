package common

/*
#include <stdlib.h>
#include "vulkan.h"
*/
import "C"
import (
	"fmt"
	"strings"
)

// APIVersion is an API version number, represents a single release of the Vulkan API
type APIVersion uint32

const (
	// Vulkan1_0 indicates core 1.0
	Vulkan1_0 APIVersion = C.VK_API_VERSION_1_0
	// Vulkan1_1 indicates core 1.1
	Vulkan1_1 APIVersion = C.VK_API_VERSION_1_1
	// Vulkan1_2 indicates core 1.2
	Vulkan1_2 APIVersion = C.VK_API_VERSION_1_2
)

// Variant is the variant number of the APIVersion number- this number is rarely included in
// vulkan version numbering, and vkngwrapper only supports variant 0. In API version
// 1.2.189, the variant is 0.
func (v APIVersion) Variant() uint32 {
	return uint32(v >> 29)
}

// Major is the major version number of the APIVersion number- this is the first number
// in the vulkan version number. In API version 1.2.189, the major version is 1.
func (v APIVersion) Major() uint32 {
	return uint32((v >> 22) & 0x7f)
}

// Minor is the minor version number of the APIVersion number- this is the second number
// in the vulkan version number. In API version 1.2.189, the minor version is 2.
func (v APIVersion) Minor() uint32 {
	return uint32((v >> 12) & 0x3ff)
}

// Patch is the patch version number of the APIVersion number- this is the third number
// in the vulkan version number. In API version 1.2.189, the patch version is 189.
func (v APIVersion) Patch() uint32 {
	return uint32(v & 0xfff)
}

// IsAtLeast returns true if the receiver v represents an API version greater than
// or equal to the argument otherVersion, or false otherwise.
func (v APIVersion) IsAtLeast(otherVersion APIVersion) bool {
	if v.Variant() != otherVersion.Variant() {
		return false
	}

	return v >= otherVersion
}

// Min returns the lower of the two APIVersions: the receiver v, or the
// argument otherVersion.
func (v APIVersion) Min(otherVersion APIVersion) APIVersion {
	if otherVersion < v {
		return otherVersion
	}

	return v
}

// Max returns the greater of the two APIVersions: the receiver v, or the
// argument otherVersion.
func (v APIVersion) Max(otherVersion APIVersion) APIVersion {
	if otherVersion > v {
		return otherVersion
	}

	return v
}

// MatchesMajorVersion will return true if the receiver v and the argument
// otherVersion have the same Variant and Major versions, or false otherwise.
func (v APIVersion) MatchesMajorVersion(otherVersion APIVersion) bool {
	return v.Variant() == otherVersion.Variant() && v.Major() == otherVersion.Major()
}

// MatchesMinorVersion will return true if the receiver v and the argument
// otherVersion have the same Variant, Major, and Minor versions, or false otherwise.
func (v APIVersion) MatchesMinorVersion(otherVersion APIVersion) bool {
	return v.MatchesMajorVersion(otherVersion) && v.Minor() == otherVersion.Minor()
}

// String provides an attractively-formatted Vulkan api version
func (v APIVersion) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d.%d.%d", v.Major(), v.Minor(), v.Patch()))

	variant := v.Variant()
	if variant != 0 {
		sb.WriteString(fmt.Sprintf(" (variant %d)", variant))
	}

	return sb.String()
}
