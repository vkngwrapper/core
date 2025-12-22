package impl1_0

import (
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroyImageView(imageView types.ImageView, callbacks *driver.AllocationCallbacks) {
	if imageView.Handle() == 0 {
		panic("imageView was uninitialized")
	}
	v.Driver.VkDestroyImageView(imageView.DeviceHandle(), imageView.Handle(), callbacks.Handle())
}
