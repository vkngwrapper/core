package pipeline

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type DescriptorSetLayoutHandle C.VkDescriptorSetLayout
type DescriptorSetLayout struct {
	handle C.VkDescriptorSetLayout
}

func (h *DescriptorSetLayout) Handle() DescriptorSetLayoutHandle {
	return DescriptorSetLayoutHandle(h.handle)
}
