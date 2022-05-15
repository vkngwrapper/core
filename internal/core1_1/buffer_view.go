package internal1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanBufferView struct {
	core1_0.BufferView
}

func PromoteBufferView(bufferView core1_0.BufferView) core1_1.BufferView {
	if !bufferView.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return bufferView.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(bufferView.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanBufferView{bufferView}
		}).(core1_1.BufferView)
}
