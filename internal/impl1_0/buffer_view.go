package impl1_0

import (
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroyBufferView(bufferView types.BufferView, callbacks *driver.AllocationCallbacks) {
	if bufferView.Handle() == 0 {
		panic("bufferView cannot be uninitialized")
	}

	v.Driver.VkDestroyBufferView(bufferView.DeviceHandle(), bufferView.Handle(), callbacks.Handle())
}
