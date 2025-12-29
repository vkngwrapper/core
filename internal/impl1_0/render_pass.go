package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) DestroyRenderPass(renderPass core1_0.RenderPass, callbacks *loader.AllocationCallbacks) {
	if !renderPass.Initialized() {
		panic("renderPass was uninitialized")
	}
	v.LoaderObj.VkDestroyRenderPass(renderPass.DeviceHandle(), renderPass.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) GetRenderAreaGranularity(renderPass core1_0.RenderPass) core1_0.Extent2D {
	if !renderPass.Initialized() {
		panic("renderPass was uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	extentPtr := (*C.VkExtent2D)(arena.Malloc(C.sizeof_struct_VkExtent2D))

	v.LoaderObj.VkGetRenderAreaGranularity(renderPass.DeviceHandle(), renderPass.Handle(), (*loader.VkExtent2D)(unsafe.Pointer(extentPtr)))

	return core1_0.Extent2D{
		Width:  int(extentPtr.width),
		Height: int(extentPtr.height),
	}
}
