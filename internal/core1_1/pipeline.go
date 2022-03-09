package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanPipeline struct {
	Driver         driver.Driver
	Device         driver.VkDevice
	PipelineHandle driver.VkPipeline
}
