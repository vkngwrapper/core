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
)
import internal_core1_0 "github.com/CannibalVox/VKng/core/internal/core1_0"

type VulkanCommandBuffer struct {
	internal_core1_0.VulkanCommandBuffer
}

func PromoteCommandBuffer(buffer core1_0.CommandBuffer) core1_1.CommandBuffer {
	goodBuffer, ok := buffer.(core1_1.CommandBuffer)
	if ok {
		return goodBuffer
	}

	oldVulkanBuffer, ok := buffer.(*internal_core1_0.VulkanCommandBuffer)
	if ok && oldVulkanBuffer.MaximumAPIVersion.IsAtLeast(common.Vulkan1_1) {
		return &VulkanCommandBuffer{*oldVulkanBuffer}
	}

	return nil
}


