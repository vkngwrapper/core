package core1_1

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanImageView is an implementation of the ImageView interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanImageView struct {
	core1_0.ImageView
}

// PromoteImageView accepts a ImageView object from any core version. If provided a image view that supports
// at least core 1.1, it will return a core1_1.ImageView. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanImageView, even if it is provided a VulkanImageView from a higher
// core version. Two Vulkan 1.1 compatible ImageView objects with the same ImageView.Handle will
// return the same interface value when passed to this method.
func PromoteImageView(imageView core1_0.ImageView) ImageView {
	if imageView == nil {
		return nil
	}
	if !imageView.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	promoted, alreadyPromoted := imageView.(ImageView)
	if alreadyPromoted {
		return promoted
	}

	return imageView.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(imageView.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanImageView{imageView}
		}).(ImageView)
}
