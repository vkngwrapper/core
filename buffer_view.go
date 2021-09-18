package core

type vulkanBufferView struct {
	handle VkBufferView
}

func (v *vulkanBufferView) Handle() VkBufferView {
	return v.handle
}
