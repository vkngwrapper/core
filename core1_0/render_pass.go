package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/driver"
	"unsafe"
)

type VulkanRenderPass struct {
	deviceDriver     driver.Driver
	device           driver.VkDevice
	renderPassHandle driver.VkRenderPass

	maximumAPIVersion common.APIVersion
}

func (p *VulkanRenderPass) Handle() driver.VkRenderPass {
	return p.renderPassHandle
}

func (p *VulkanRenderPass) Driver() driver.Driver {
	return p.deviceDriver
}

func (p *VulkanRenderPass) DeviceHandle() driver.VkDevice {
	return p.device
}

func (p *VulkanRenderPass) APIVersion() common.APIVersion {
	return p.maximumAPIVersion
}

func (p *VulkanRenderPass) Destroy(callbacks *driver.AllocationCallbacks) {
	p.deviceDriver.VkDestroyRenderPass(p.device, p.renderPassHandle, callbacks.Handle())
	p.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(p.renderPassHandle))
}

func (p *VulkanRenderPass) RenderAreaGranularity() Extent2D {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	extentPtr := (*C.VkExtent2D)(arena.Malloc(C.sizeof_struct_VkExtent2D))

	p.deviceDriver.VkGetRenderAreaGranularity(p.device, p.renderPassHandle, (*driver.VkExtent2D)(unsafe.Pointer(extentPtr)))

	return Extent2D{
		Width:  int(extentPtr.width),
		Height: int(extentPtr.height),
	}
}
