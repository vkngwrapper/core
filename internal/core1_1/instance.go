package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/core1_0"
)

type VulkanInstance struct {
	core1_0.Instance
}
