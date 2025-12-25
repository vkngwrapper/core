package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) DestroyImage(image core.Image, callbacks *loader.AllocationCallbacks) {
	if image.Handle() == 0 {
		panic("image was uninitialized")
	}

	v.LoaderObj.VkDestroyImage(image.DeviceHandle(), image.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) GetImageMemoryRequirements(image core.Image) *core1_0.MemoryRequirements {
	if image.Handle() == 0 {
		panic("image was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	memRequirementsUnsafe := arena.Malloc(C.sizeof_struct_VkMemoryRequirements)

	v.LoaderObj.VkGetImageMemoryRequirements(image.DeviceHandle(), image.Handle(), (*loader.VkMemoryRequirements)(memRequirementsUnsafe))

	memRequirements := (*C.VkMemoryRequirements)(memRequirementsUnsafe)

	return &core1_0.MemoryRequirements{
		Size:           int(memRequirements.size),
		Alignment:      int(memRequirements.alignment),
		MemoryTypeBits: uint32(memRequirements.memoryTypeBits),
	}
}

func (v *DeviceVulkanDriver) BindImageMemory(image core.Image, memory core.DeviceMemory, offset int) (common.VkResult, error) {
	if image.Handle() == 0 {
		return core1_0.VKErrorUnknown, errors.New("received uninitialized Image")
	}
	if memory.Handle() == 0 {
		return core1_0.VKErrorUnknown, errors.New("received uninitialized DeviceMemory")
	}
	if offset < 0 {
		return core1_0.VKErrorUnknown, errors.New("received negative offset")
	}

	return v.LoaderObj.VkBindImageMemory(image.DeviceHandle(), image.Handle(), memory.Handle(), loader.VkDeviceSize(offset))
}

func (v *DeviceVulkanDriver) GetImageSubresourceLayout(image core.Image, subresource *core1_0.ImageSubresource) *core1_0.SubresourceLayout {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	subresourcePtr := (*C.VkImageSubresource)(arena.Malloc(C.sizeof_struct_VkImageSubresource))
	subresourceLayoutUnsafe := arena.Malloc(C.sizeof_struct_VkSubresourceLayout)

	subresourcePtr.aspectMask = C.VkImageAspectFlags(subresource.AspectMask)
	subresourcePtr.mipLevel = C.uint32_t(subresource.MipLevel)
	subresourcePtr.arrayLayer = C.uint32_t(subresource.ArrayLayer)

	v.LoaderObj.VkGetImageSubresourceLayout(image.DeviceHandle(), image.Handle(), (*loader.VkImageSubresource)(unsafe.Pointer(subresourcePtr)), (*loader.VkSubresourceLayout)(subresourceLayoutUnsafe))

	subresourceLayout := (*C.VkSubresourceLayout)(subresourceLayoutUnsafe)
	return &core1_0.SubresourceLayout{
		Offset:     int(subresourceLayout.offset),
		Size:       int(subresourceLayout.size),
		RowPitch:   int(subresourceLayout.rowPitch),
		ArrayPitch: int(subresourceLayout.arrayPitch),
		DepthPitch: int(subresourceLayout.depthPitch),
	}
}
