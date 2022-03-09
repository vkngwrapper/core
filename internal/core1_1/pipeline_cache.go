package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanPipelineCache struct {
	Driver              driver.Driver
	Device              driver.VkDevice
	PipelineCacheHandle driver.VkPipelineCache
}
