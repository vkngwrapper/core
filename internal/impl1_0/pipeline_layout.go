package impl1_0

import (
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroyPipelineLayout(pipelineLayout types.PipelineLayout, callbacks *driver.AllocationCallbacks) {
	if pipelineLayout.Handle() == 0 {
		panic("pipelineLayout was uninitialized")
	}
	v.Driver.VkDestroyPipelineLayout(pipelineLayout.DeviceHandle(), pipelineLayout.Handle(), callbacks.Handle())
}
