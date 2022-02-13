package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanRenderPass struct {
	driver driver.Driver
	device driver.VkDevice
	handle driver.VkRenderPass
}

func (p *VulkanRenderPass) Handle() driver.VkRenderPass {
	return p.handle
}

func (p *VulkanRenderPass) Destroy(callbacks *driver.AllocationCallbacks) {
	p.driver.VkDestroyRenderPass(p.device, p.handle, callbacks.Handle())
}

func (p *VulkanRenderPass) RenderAreaGranularity() common.Extent2D {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	extentPtr := (*C.VkExtent2D)(arena.Malloc(C.sizeof_struct_VkExtent2D))

	p.driver.VkGetRenderAreaGranularity(p.device, p.handle, (*driver.VkExtent2D)(unsafe.Pointer(extentPtr)))

	return common.Extent2D{
		Width:  int(extentPtr.width),
		Height: int(extentPtr.height),
	}
}
