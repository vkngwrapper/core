package resources

import "github.com/CannibalVox/VKng/core/loader"

type vulkanSampler struct {
	handle loader.VkSampler
}

func (s *vulkanSampler) Handle() loader.VkSampler {
	return s.handle
}
