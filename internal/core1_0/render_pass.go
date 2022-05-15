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
	DeviceDriver     driver.Driver
	Device           driver.VkDevice
	RenderPassHandle driver.VkRenderPass

	MaximumAPIVersion common.APIVersion
}

func (p *VulkanRenderPass) Handle() driver.VkRenderPass {
	return p.RenderPassHandle
}

func (p *VulkanRenderPass) Driver() driver.Driver {
	return p.DeviceDriver
}

func (p *VulkanRenderPass) APIVersion() common.APIVersion {
	return p.MaximumAPIVersion
}

func (p *VulkanRenderPass) Destroy(callbacks *driver.AllocationCallbacks) {
	p.DeviceDriver.VkDestroyRenderPass(p.Device, p.RenderPassHandle, callbacks.Handle())
	p.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(p.RenderPassHandle))
}

func (p *VulkanRenderPass) RenderAreaGranularity() common.Extent2D {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	extentPtr := (*C.VkExtent2D)(arena.Malloc(C.sizeof_struct_VkExtent2D))

	p.DeviceDriver.VkGetRenderAreaGranularity(p.Device, p.RenderPassHandle, (*driver.VkExtent2D)(unsafe.Pointer(extentPtr)))

	return common.Extent2D{
		Width:  int(extentPtr.width),
		Height: int(extentPtr.height),
	}
}
