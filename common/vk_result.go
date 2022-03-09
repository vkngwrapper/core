package common

import (
	"fmt"
)

type VkResultError struct {
	code VkResult
}

func (err *VkResultError) Error() string {
	return fmt.Sprintf("vulkan error: %s", err.code.String())
}
