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
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroyRenderPass(renderPass types.RenderPass, callbacks *driver.AllocationCallbacks) {
	if renderPass.Handle() == 0 {
		panic("renderPass was uninitialized")
	}
	v.Driver.VkDestroyRenderPass(renderPass.DeviceHandle(), renderPass.Handle(), callbacks.Handle())
}

func (v *Vulkan) GetRenderAreaGranularity(renderPass types.RenderPass) core1_0.Extent2D {
	if renderPass.Handle() == 0 {
		panic("renderPass was uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	extentPtr := (*C.VkExtent2D)(arena.Malloc(C.sizeof_struct_VkExtent2D))

	v.Driver.VkGetRenderAreaGranularity(renderPass.DeviceHandle(), renderPass.Handle(), (*driver.VkExtent2D)(unsafe.Pointer(extentPtr)))

	return core1_0.Extent2D{
		Width:  int(extentPtr.width),
		Height: int(extentPtr.height),
	}
}
