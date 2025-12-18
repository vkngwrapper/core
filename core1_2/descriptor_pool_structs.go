package core1_2

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

const (
	// DescriptorPoolCreateUpdateAfterBind specifies that DescriptorSet objects allocated from this
	// pool can include bindings with DescriptorBindingUpdateAfterBind
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateFlagBits.html
	DescriptorPoolCreateUpdateAfterBind core1_0.DescriptorPoolCreateFlags = C.VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT

	// VkErrorFragmentation indicates a DescriptorPool creation has failed due to fragmentation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VkErrorFragmentation common.VkResult = C.VK_ERROR_FRAGMENTATION
)

func init() {
	DescriptorPoolCreateUpdateAfterBind.Register("Update After Bind")

	VkErrorFragmentation.Register("fragmentation")
}
