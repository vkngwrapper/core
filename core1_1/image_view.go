package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
)

type VulkanImageView struct {
	core1_0.ImageView
}

func PromoteImageView(imageView core1_0.ImageView) ImageView {
	if !imageView.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return imageView.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(imageView.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanImageView{imageView}
		}).(ImageView)
}
