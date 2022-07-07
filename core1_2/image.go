package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

type VulkanImage struct {
	core1_1.Image
}

func PromoteImage(image core1_0.Image) Image {
	if !image.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedImage := core1_1.PromoteImage(image)

	return image.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(image.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanImage{promotedImage}
		}).(Image)
}
