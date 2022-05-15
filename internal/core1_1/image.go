package internal1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanImage struct {
	core1_0.Image
}

func PromoteImage(image core1_0.Image) core1_1.Image {
	if !image.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return image.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(image.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanImage{image}
		}).(core1_1.Image)
}
