package resources

import "github.com/CannibalVox/VKng/core/loader"

type vulkanBufferView struct {
	handle loader.VkBufferView
}

func (v *vulkanBufferView) Handle() loader.VkBufferView {
	return v.handle
}
