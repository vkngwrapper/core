package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
)

type VulkanBufferView struct {
	core1_0.BufferView
}

func PromoteBufferView(bufferView core1_0.BufferView) BufferView {
	if !bufferView.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return bufferView.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(bufferView.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanBufferView{bufferView}
		}).(BufferView)
}
