package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type VulkanImage struct {
	deviceDriver driver.Driver
	imageHandle  driver.VkImage
	device       driver.VkDevice

	maximumAPIVersion common.APIVersion
}

func (i *VulkanImage) Handle() driver.VkImage {
	return i.imageHandle
}

func (i *VulkanImage) DeviceHandle() driver.VkDevice {
	return i.device
}

func (i *VulkanImage) APIVersion() common.APIVersion {
	return i.maximumAPIVersion
}

func (i *VulkanImage) Driver() driver.Driver {
	return i.deviceDriver
}

func (i *VulkanImage) Destroy(callbacks *driver.AllocationCallbacks) {
	i.deviceDriver.VkDestroyImage(i.device, i.imageHandle, callbacks.Handle())
	i.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(i.imageHandle))
}

func (i *VulkanImage) MemoryRequirements() *MemoryRequirements {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	memRequirementsUnsafe := arena.Malloc(C.sizeof_struct_VkMemoryRequirements)

	i.deviceDriver.VkGetImageMemoryRequirements(i.device, i.imageHandle, (*driver.VkMemoryRequirements)(memRequirementsUnsafe))

	memRequirements := (*C.VkMemoryRequirements)(memRequirementsUnsafe)

	return &MemoryRequirements{
		Size:       int(memRequirements.size),
		Alignment:  int(memRequirements.alignment),
		MemoryType: uint32(memRequirements.memoryTypeBits),
	}
}

func (i *VulkanImage) BindImageMemory(memory DeviceMemory, offset int) (common.VkResult, error) {
	if memory == nil {
		return VKErrorUnknown, errors.New("received nil DeviceMemory")
	}
	if offset < 0 {
		return VKErrorUnknown, errors.New("received negative offset")
	}

	return i.deviceDriver.VkBindImageMemory(i.device, i.imageHandle, memory.Handle(), driver.VkDeviceSize(offset))
}

func (i *VulkanImage) SubresourceLayout(subresource *ImageSubresource) *SubresourceLayout {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	subresourcePtr := (*C.VkImageSubresource)(arena.Malloc(C.sizeof_struct_VkImageSubresource))
	subresourceLayoutUnsafe := arena.Malloc(C.sizeof_struct_VkSubresourceLayout)

	subresourcePtr.aspectMask = C.VkImageAspectFlags(subresource.AspectMask)
	subresourcePtr.mipLevel = C.uint32_t(subresource.MipLevel)
	subresourcePtr.arrayLayer = C.uint32_t(subresource.ArrayLayer)

	i.deviceDriver.VkGetImageSubresourceLayout(i.device, i.imageHandle, (*driver.VkImageSubresource)(unsafe.Pointer(subresourcePtr)), (*driver.VkSubresourceLayout)(subresourceLayoutUnsafe))

	subresourceLayout := (*C.VkSubresourceLayout)(subresourceLayoutUnsafe)
	return &SubresourceLayout{
		Offset:     int(subresourceLayout.offset),
		Size:       int(subresourceLayout.size),
		RowPitch:   int(subresourceLayout.rowPitch),
		ArrayPitch: int(subresourceLayout.arrayPitch),
		DepthPitch: int(subresourceLayout.depthPitch),
	}
}
