package core1_1

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanDescriptorUpdateTemplate is an implementation of the DescriptorUpdateTemplate interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDescriptorUpdateTemplate struct {
	DeviceDriver             driver.Driver
	Device                   driver.VkDevice
	DescriptorTemplateHandle driver.VkDescriptorUpdateTemplate

	MaximumAPIVersion common.APIVersion
}

func (t *VulkanDescriptorUpdateTemplate) Handle() driver.VkDescriptorUpdateTemplate {
	return t.DescriptorTemplateHandle
}

func (t *VulkanDescriptorUpdateTemplate) DeviceHandle() driver.VkDevice {
	return t.Device
}

func (t *VulkanDescriptorUpdateTemplate) Driver() driver.Driver {
	return t.DeviceDriver
}

func (t *VulkanDescriptorUpdateTemplate) APIVersion() common.APIVersion {
	return t.MaximumAPIVersion
}

func (t *VulkanDescriptorUpdateTemplate) Destroy(allocator *driver.AllocationCallbacks) {
	t.DeviceDriver.VkDestroyDescriptorUpdateTemplate(t.Device, t.DescriptorTemplateHandle, allocator.Handle())
}

func (t *VulkanDescriptorUpdateTemplate) UpdateDescriptorSetFromImage(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorImageInfo) {
	if descriptorSet == nil {
		panic("descriptorSet cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoUnsafe := arena.Malloc(C.sizeof_struct_VkDescriptorImageInfo)
	info := (*C.VkDescriptorImageInfo)(infoUnsafe)
	info.sampler = nil
	info.imageView = nil
	info.imageLayout = C.VkImageLayout(data.ImageLayout)

	if data.Sampler != nil {
		info.sampler = C.VkSampler(unsafe.Pointer(data.Sampler.Handle()))
	}

	if data.ImageView != nil {
		info.imageView = C.VkImageView(unsafe.Pointer(data.ImageView.Handle()))
	}

	t.DeviceDriver.VkUpdateDescriptorSetWithTemplate(
		t.Device,
		descriptorSet.Handle(),
		t.DescriptorTemplateHandle,
		infoUnsafe,
	)
}

func (t *VulkanDescriptorUpdateTemplate) UpdateDescriptorSetFromBuffer(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorBufferInfo) {
	if descriptorSet == nil {
		panic("descriptorSet cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoUnsafe := arena.Malloc(C.sizeof_struct_VkDescriptorBufferInfo)
	info := (*C.VkDescriptorBufferInfo)(infoUnsafe)
	info.buffer = nil
	info.offset = C.VkDeviceSize(data.Offset)
	info._range = C.VkDeviceSize(data.Range)

	if data.Buffer != nil {
		info.buffer = C.VkBuffer(unsafe.Pointer(data.Buffer.Handle()))
	}

	t.DeviceDriver.VkUpdateDescriptorSetWithTemplate(
		t.Device,
		descriptorSet.Handle(),
		t.DescriptorTemplateHandle,
		infoUnsafe,
	)
}

func (t *VulkanDescriptorUpdateTemplate) UpdateDescriptorSetFromObjectHandle(descriptorSet core1_0.DescriptorSet, data driver.VulkanHandle) {
	if descriptorSet == nil {
		panic("descriptorSet cannot be nil")
	}

	t.DeviceDriver.VkUpdateDescriptorSetWithTemplate(
		t.Device,
		descriptorSet.Handle(),
		t.DescriptorTemplateHandle,
		unsafe.Pointer(data),
	)
}
