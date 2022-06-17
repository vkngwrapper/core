package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanImageView struct {
	core1_1.ImageView
}

func PromoteImageView(imageView core1_0.ImageView) core1_2.ImageView {
	if !imageView.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return imageView.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(imageView.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanImageView{core1_1.PromoteImageView(imageView)}
		}).(core1_2.ImageView)
}
