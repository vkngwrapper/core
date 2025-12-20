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
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanImage is an implementation of the Image interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanImage struct {
	DeviceDriver driver.Driver
	ImageHandle  driver.VkImage
	Device       driver.VkDevice

	MaximumAPIVersion common.APIVersion
}

func (i *VulkanImage) Handle() driver.VkImage {
	return i.ImageHandle
}

func (i *VulkanImage) DeviceHandle() driver.VkDevice {
	return i.Device
}

func (i *VulkanImage) APIVersion() common.APIVersion {
	return i.MaximumAPIVersion
}

func (i *VulkanImage) Driver() driver.Driver {
	return i.DeviceDriver
}

func (i *VulkanImage) Destroy(callbacks *driver.AllocationCallbacks) {
	i.DeviceDriver.VkDestroyImage(i.Device, i.ImageHandle, callbacks.Handle())
}

func (i *VulkanImage) MemoryRequirements() *core1_0.MemoryRequirements {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	memRequirementsUnsafe := arena.Malloc(C.sizeof_struct_VkMemoryRequirements)

	i.DeviceDriver.VkGetImageMemoryRequirements(i.Device, i.ImageHandle, (*driver.VkMemoryRequirements)(memRequirementsUnsafe))

	memRequirements := (*C.VkMemoryRequirements)(memRequirementsUnsafe)

	return &core1_0.MemoryRequirements{
		Size:           int(memRequirements.size),
		Alignment:      int(memRequirements.alignment),
		MemoryTypeBits: uint32(memRequirements.memoryTypeBits),
	}
}

func (i *VulkanImage) BindImageMemory(memory core1_0.DeviceMemory, offset int) (common.VkResult, error) {
	if memory == nil {
		return core1_0.VKErrorUnknown, errors.New("received nil DeviceMemory")
	}
	if offset < 0 {
		return core1_0.VKErrorUnknown, errors.New("received negative offset")
	}

	return i.DeviceDriver.VkBindImageMemory(i.Device, i.ImageHandle, memory.Handle(), driver.VkDeviceSize(offset))
}

func (i *VulkanImage) SubresourceLayout(subresource *core1_0.ImageSubresource) *core1_0.SubresourceLayout {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	subresourcePtr := (*C.VkImageSubresource)(arena.Malloc(C.sizeof_struct_VkImageSubresource))
	subresourceLayoutUnsafe := arena.Malloc(C.sizeof_struct_VkSubresourceLayout)

	subresourcePtr.aspectMask = C.VkImageAspectFlags(subresource.AspectMask)
	subresourcePtr.mipLevel = C.uint32_t(subresource.MipLevel)
	subresourcePtr.arrayLayer = C.uint32_t(subresource.ArrayLayer)

	i.DeviceDriver.VkGetImageSubresourceLayout(i.Device, i.ImageHandle, (*driver.VkImageSubresource)(unsafe.Pointer(subresourcePtr)), (*driver.VkSubresourceLayout)(subresourceLayoutUnsafe))

	subresourceLayout := (*C.VkSubresourceLayout)(subresourceLayoutUnsafe)
	return &core1_0.SubresourceLayout{
		Offset:     int(subresourceLayout.offset),
		Size:       int(subresourceLayout.size),
		RowPitch:   int(subresourceLayout.rowPitch),
		ArrayPitch: int(subresourceLayout.arrayPitch),
		DepthPitch: int(subresourceLayout.depthPitch),
	}
}
