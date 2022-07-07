package core1_2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
)

const (
	DescriptorPoolCreateUpdateAfterBind core1_0.DescriptorPoolCreateFlags = C.VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT

	VkErrorFragmentation common.VkResult = C.VK_ERROR_FRAGMENTATION
)

func init() {
	DescriptorPoolCreateUpdateAfterBind.Register("Update After Bind")

	VkErrorFragmentation.Register("fragmentation")
}
