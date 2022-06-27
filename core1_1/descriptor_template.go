package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

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
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoUnsafe := arena.Malloc(C.sizeof_struct_VkDescriptorImageInfo)
	info := (*C.VkDescriptorImageInfo)(infoUnsafe)
	info.sampler = C.VkSampler(unsafe.Pointer(data.Sampler.Handle()))
	info.imageView = C.VkImageView(unsafe.Pointer(data.ImageView.Handle()))
	info.imageLayout = C.VkImageLayout(data.ImageLayout)

	t.DeviceDriver.VkUpdateDescriptorSetWithTemplate(
		t.Device,
		descriptorSet.Handle(),
		t.DescriptorTemplateHandle,
		infoUnsafe,
	)
}

func (t *VulkanDescriptorUpdateTemplate) UpdateDescriptorSetFromBuffer(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorBufferInfo) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoUnsafe := arena.Malloc(C.sizeof_struct_VkDescriptorBufferInfo)
	info := (*C.VkDescriptorBufferInfo)(infoUnsafe)
	info.buffer = C.VkBuffer(unsafe.Pointer(data.Buffer.Handle()))
	info.offset = C.VkDeviceSize(data.Offset)
	info._range = C.VkDeviceSize(data.Range)

	t.DeviceDriver.VkUpdateDescriptorSetWithTemplate(
		t.Device,
		descriptorSet.Handle(),
		t.DescriptorTemplateHandle,
		infoUnsafe,
	)
}

func (t *VulkanDescriptorUpdateTemplate) UpdateDescriptorSetFromObjectHandle(descriptorSet core1_0.DescriptorSet, data driver.VulkanHandle) {
	t.DeviceDriver.VkUpdateDescriptorSetWithTemplate(
		t.Device,
		descriptorSet.Handle(),
		t.DescriptorTemplateHandle,
		unsafe.Pointer(data),
	)
}
