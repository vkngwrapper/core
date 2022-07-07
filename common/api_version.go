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

type APIVersion uint32

const (
	Vulkan1_0 APIVersion = C.VK_API_VERSION_1_0
	Vulkan1_1 APIVersion = C.VK_API_VERSION_1_1
	Vulkan1_2 APIVersion = C.VK_API_VERSION_1_2
)

func (v APIVersion) Variant() uint32 {
	return uint32(v >> 29)
}

func (v APIVersion) Major() uint32 {
	return uint32((v >> 22) & 0x7f)
}

func (v APIVersion) Minor() uint32 {
	return uint32((v >> 12) & 0x3ff)
}

func (v APIVersion) Patch() uint32 {
	return uint32(v & 0xfff)
}

func (v APIVersion) IsAtLeast(otherVersion APIVersion) bool {
	if v.Variant() != otherVersion.Variant() {
		return false
	}

	return v >= otherVersion
}

func (v APIVersion) Min(otherVersion APIVersion) APIVersion {
	if otherVersion < v {
		return otherVersion
	}

	return v
}

func (v APIVersion) Max(otherVersion APIVersion) APIVersion {
	if otherVersion > v {
		return otherVersion
	}

	return v
}

func (v APIVersion) MatchesMajorVersion(otherVersion APIVersion) bool {
	return v.Variant() == otherVersion.Variant() && v.Major() == otherVersion.Major()
}

func (v APIVersion) MatchesMinorVersion(otherVersion APIVersion) bool {
	return v.MatchesMajorVersion(otherVersion) && v.Minor() == otherVersion.Minor()
}

func (v APIVersion) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d.%d.%d", v.Major(), v.Minor(), v.Patch()))

	variant := v.Variant()
	if variant != 0 {
		sb.WriteString(fmt.Sprintf(" (variant %d)", variant))
	}

	return sb.String()
}
