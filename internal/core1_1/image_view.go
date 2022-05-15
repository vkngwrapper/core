package internal1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanImageView struct {
	core1_0.ImageView
}

func PromoteImageView(imageView core1_0.ImageView) core1_1.ImageView {
	if !imageView.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return imageView.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(imageView.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanImageView{imageView}
		}).(core1_1.ImageView)
}
