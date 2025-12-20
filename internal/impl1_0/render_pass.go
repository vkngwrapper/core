package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanRenderPass is an implementation of the RenderPass interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
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

func (p *VulkanRenderPass) DeviceHandle() driver.VkDevice {
	return p.Device
}

func (p *VulkanRenderPass) APIVersion() common.APIVersion {
	return p.MaximumAPIVersion
}

func (p *VulkanRenderPass) Destroy(callbacks *driver.AllocationCallbacks) {
	p.DeviceDriver.VkDestroyRenderPass(p.Device, p.RenderPassHandle, callbacks.Handle())
}

func (p *VulkanRenderPass) RenderAreaGranularity() core1_0.Extent2D {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	extentPtr := (*C.VkExtent2D)(arena.Malloc(C.sizeof_struct_VkExtent2D))

	p.DeviceDriver.VkGetRenderAreaGranularity(p.Device, p.RenderPassHandle, (*driver.VkExtent2D)(unsafe.Pointer(extentPtr)))

	return core1_0.Extent2D{
		Width:  int(extentPtr.width),
		Height: int(extentPtr.height),
	}
}
