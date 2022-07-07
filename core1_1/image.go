package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
)

type VulkanImage struct {
	core1_0.Image
}

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
