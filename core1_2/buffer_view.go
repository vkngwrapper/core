package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

type VulkanBufferView struct {
	core1_1.BufferView
}

func PromoteBufferView(bufferView core1_0.BufferView) BufferView {
	if !bufferView.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedBufferView := core1_1.PromoteBufferView(bufferView)

	return bufferView.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(bufferView.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanBufferView{promotedBufferView}
		}).(BufferView)
}
