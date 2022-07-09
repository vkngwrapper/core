package common

import (
	"fmt"
)

// VkResultError is an error object wrapping a VkResult return code. Any
// Vulkan command which returns a VkResult will return a VkResultError when
// returning a VkResult that indicates an error. Be aware: there are many
// non-VK_SUCCESS, non-error return codes!
type VkResultError struct {
	code VkResult
}

func (err *VkResultError) Error() string {
	return fmt.Sprintf("vulkan error: %s", err.code.String())
}
