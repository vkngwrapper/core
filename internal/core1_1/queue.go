package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanQueue struct {
	DeviceDriver driver.Driver
	QueueHandle  driver.VkQueue
}
