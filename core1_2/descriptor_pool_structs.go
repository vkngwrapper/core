package core1_2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

const (
	DescriptorPoolCreateUpdateAfterBind common.DescriptorPoolCreateFlags = C.VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT

	VkErrorFragmentation common.VkResult = C.VK_ERROR_FRAGMENTATION
)

func init() {
	DescriptorPoolCreateUpdateAfterBind.Register("Update After Bind")

	VkErrorFragmentation.Register("fragmentation")
}
