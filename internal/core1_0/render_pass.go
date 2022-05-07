package internal1_0

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
	Driver           driver.Driver
	Device           driver.VkDevice
	RenderPassHandle driver.VkRenderPass

	MaximumAPIVersion common.APIVersion
}

func (p *VulkanRenderPass) Handle() driver.VkRenderPass {
	return p.RenderPassHandle
}

func (p *VulkanRenderPass) Destroy(callbacks *driver.AllocationCallbacks) {
	p.Driver.VkDestroyRenderPass(p.Device, p.RenderPassHandle, callbacks.Handle())
	p.Driver.ObjectStore().Delete(driver.VulkanHandle(p.RenderPassHandle), p)
}

func (p *VulkanRenderPass) RenderAreaGranularity() common.Extent2D {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	extentPtr := (*C.VkExtent2D)(arena.Malloc(C.sizeof_struct_VkExtent2D))

	p.Driver.VkGetRenderAreaGranularity(p.Device, p.RenderPassHandle, (*driver.VkExtent2D)(unsafe.Pointer(extentPtr)))

	return common.Extent2D{
		Width:  int(extentPtr.width),
		Height: int(extentPtr.height),
	}
}
