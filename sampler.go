package core

type vulkanSampler struct {
	handle VkSampler
}

func (s *vulkanSampler) Handle() VkSampler {
	return s.handle
}
