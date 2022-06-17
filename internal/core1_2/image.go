package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanImage struct {
	core1_1.Image
}

func PromoteImage(image core1_0.Image) core1_2.Image {
	if !image.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return image.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(image.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanImage{core1_1.PromoteImage(image)}
		}).(core1_2.Image)
}
