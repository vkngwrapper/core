package core1_2

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanImage is an implementation of the Image interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanImage struct {
	core1_1.Image
}

// PromoteImage accepts a Image object from any core version. If provided an image that supports
// at least core 1.2, it will return a core1_2.Image. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanImage, even if it is provided a VulkanImage from a higher
// core version. Two Vulkan 1.2 compatible Image objects with the same Image.Handle will
// return the same interface value when passed to this method.
func PromoteImage(image core1_0.Image) Image {
	if image == nil {
		return nil
	}
	if !image.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promoted, alreadyPromoted := image.(Image)
	if alreadyPromoted {
		return promoted
	}

	promotedImage := core1_1.PromoteImage(image)

	return image.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(image.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanImage{promotedImage}
		}).(Image)
}
