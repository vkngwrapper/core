package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	internal_core1_0 "github.com/CannibalVox/VKng/core/internal/core1_0"
)

type VulkanBuffer struct {
	internal_core1_0.VulkanBuffer
}

func PromoteBuffer(buffer core1_0.Buffer) core1_1.Buffer {
	goodBuffer, ok := buffer.(core1_1.Buffer)
	if ok {
		return goodBuffer
	}

	oldVulkanBuffer, ok := buffer.(*internal_core1_0.VulkanBuffer)
	if ok && oldVulkanBuffer.MaximumAPIVersion.IsAtLeast(common.Vulkan1_1) {
		return &VulkanBuffer{*oldVulkanBuffer}
	}

	return nil
}
