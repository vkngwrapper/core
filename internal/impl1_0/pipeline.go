package impl1_0

import (
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroyPipeline(pipeline types.Pipeline, callbacks *driver.AllocationCallbacks) {
	if pipeline.Handle() == 0 {
		panic("pipeline was uninitialized")
	}
	v.Driver.VkDestroyPipeline(pipeline.DeviceHandle(), pipeline.Handle(), callbacks.Handle())
}
