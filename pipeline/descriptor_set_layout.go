package pipeline

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/loader"

type DescriptorSetLayout struct {
	loader loader.Loader
	handle loader.VkDescriptorSetLayout
}

func (h *DescriptorSetLayout) Handle() loader.VkDescriptorSetLayout {
	return h.handle
}
