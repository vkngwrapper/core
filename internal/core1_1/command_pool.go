package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import internal_core1_0 "github.com/CannibalVox/VKng/core/internal/core1_0"

type VulkanCommandPool struct {
	internal_core1_0.VulkanCommandPool
}

func PromoteCommandPool()
