package core1_1

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanImage is an implementation of the Image interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanImage struct {
	core1_0.Image
}

// PromoteImage accepts a Image object from any core version. If provided a image that supports
// at least core 1.1, it will return a core1_1.Image. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanImage, even if it is provided a VulkanImage from a higher
// core version. Two Vulkan 1.1 compatible Image objects with the same Image.Handle will
// return the same interface value when passed to this method.
func PromoteImage(image core1_0.Image) Image {
	if !image.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return image.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(image.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanImage{image}
		}).(Image)
}
