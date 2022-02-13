package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/iface"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type VulkanImage struct {
	driver driver.Driver
	handle driver.VkImage
	device driver.VkDevice
}

func CreateImageFromHandles(handle driver.VkImage, device driver.VkDevice, driver driver.Driver) *VulkanImage {
	return &VulkanImage{handle: handle, device: device, driver: driver}
}

func (i *VulkanImage) Handle() driver.VkImage {
	return i.handle
}

func (i *VulkanImage) Destroy(callbacks *driver.AllocationCallbacks) {
	i.driver.VkDestroyImage(i.device, i.handle, callbacks.Handle())
}

func (i *VulkanImage) MemoryRequirements() *core1_0.MemoryRequirements {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	memRequirementsUnsafe := arena.Malloc(C.sizeof_struct_VkMemoryRequirements)

	i.driver.VkGetImageMemoryRequirements(i.device, i.handle, (*driver.VkMemoryRequirements)(memRequirementsUnsafe))

	memRequirements := (*C.VkMemoryRequirements)(memRequirementsUnsafe)

	return &core1_0.MemoryRequirements{
		Size:       int(memRequirements.size),
		Alignment:  int(memRequirements.alignment),
		MemoryType: uint32(memRequirements.memoryTypeBits),
	}
}

func (i *VulkanImage) BindImageMemory(memory iface.DeviceMemory, offset int) (common.VkResult, error) {
	if memory == nil {
		return common.VKErrorUnknown, errors.New("received nil DeviceMemory")
	}
	if offset < 0 {
		return common.VKErrorUnknown, errors.New("received negative offset")
	}

	return i.driver.VkBindImageMemory(i.device, i.handle, memory.Handle(), driver.VkDeviceSize(offset))
}

func (i *VulkanImage) SubresourceLayout(subresource *common.ImageSubresource) *common.SubresourceLayout {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	subresourcePtr := (*C.VkImageSubresource)(arena.Malloc(C.sizeof_struct_VkImageSubresource))
	subresourceLayoutUnsafe := arena.Malloc(C.sizeof_struct_VkSubresourceLayout)

	subresourcePtr.aspectMask = C.VkImageAspectFlags(subresource.AspectMask)
	subresourcePtr.mipLevel = C.uint32_t(subresource.MipLevel)
	subresourcePtr.arrayLayer = C.uint32_t(subresource.ArrayLayer)

	i.driver.VkGetImageSubresourceLayout(i.device, i.handle, (*driver.VkImageSubresource)(unsafe.Pointer(subresourcePtr)), (*driver.VkSubresourceLayout)(subresourceLayoutUnsafe))

	subresourceLayout := (*C.VkSubresourceLayout)(subresourceLayoutUnsafe)
	return &common.SubresourceLayout{
		Offset:     int(subresourceLayout.offset),
		Size:       int(subresourceLayout.size),
		RowPitch:   int(subresourceLayout.rowPitch),
		ArrayPitch: int(subresourceLayout.arrayPitch),
		DepthPitch: int(subresourceLayout.depthPitch),
	}
}
