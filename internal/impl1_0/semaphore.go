package impl1_0

import (
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroySemaphore(semaphore types.Semaphore, callbacks *driver.AllocationCallbacks) {
	if semaphore.Handle() == 0 {
		panic("semaphore was uninitialized")
	}
	v.Driver.VkDestroySemaphore(semaphore.DeviceHandle(), semaphore.Handle(), callbacks.Handle())
}
