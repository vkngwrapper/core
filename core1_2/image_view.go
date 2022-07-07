package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

type VulkanImageView struct {
	core1_1.ImageView
}

func PromoteImageView(imageView core1_0.ImageView) ImageView {
	if !imageView.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedImageView := core1_1.PromoteImageView(imageView)
	return imageView.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(imageView.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanImageView{promotedImageView}
		}).(ImageView)
}
