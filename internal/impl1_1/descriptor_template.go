package impl1_1

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
	"github.com/vkngwrapper/core/v3/types"
)

func (v *DeviceVulkanDriver) DestroyDescriptorUpdateTemplate(template types.DescriptorUpdateTemplate, allocator *loader.AllocationCallbacks) {
	v.LoaderObj.VkDestroyDescriptorUpdateTemplate(template.DeviceHandle(), template.Handle(), allocator.Handle())
}

func (v *DeviceVulkanDriver) UpdateDescriptorSetWithTemplateFromImage(descriptorSet types.DescriptorSet, template types.DescriptorUpdateTemplate, data core1_0.DescriptorImageInfo) {
	if descriptorSet.Handle() == 0 {
		panic("descriptorSet cannot be uninitialized")
	}
	if template.Handle() == 0 {
		panic("template cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoUnsafe := arena.Malloc(C.sizeof_struct_VkDescriptorImageInfo)
	info := (*C.VkDescriptorImageInfo)(infoUnsafe)
	info.sampler = nil
	info.imageView = nil
	info.imageLayout = C.VkImageLayout(data.ImageLayout)

	if data.Sampler.Handle() != 0 {
		info.sampler = C.VkSampler(unsafe.Pointer(data.Sampler.Handle()))
	}

	if data.ImageView.Handle() != 0 {
		info.imageView = C.VkImageView(unsafe.Pointer(data.ImageView.Handle()))
	}

	v.LoaderObj.VkUpdateDescriptorSetWithTemplate(
		descriptorSet.DeviceHandle(),
		descriptorSet.Handle(),
		template.Handle(),
		infoUnsafe,
	)
}

func (v *DeviceVulkanDriver) UpdateDescriptorSetWithTemplateFromBuffer(descriptorSet types.DescriptorSet, template types.DescriptorUpdateTemplate, data core1_0.DescriptorBufferInfo) {
	if descriptorSet.Handle() == 0 {
		panic("descriptorSet cannot be uninitialized")
	}
	if template.Handle() == 0 {
		panic("template cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoUnsafe := arena.Malloc(C.sizeof_struct_VkDescriptorBufferInfo)
	info := (*C.VkDescriptorBufferInfo)(infoUnsafe)
	info.buffer = nil
	info.offset = C.VkDeviceSize(data.Offset)
	info._range = C.VkDeviceSize(data.Range)

	if data.Buffer.Handle() != 0 {
		info.buffer = C.VkBuffer(unsafe.Pointer(data.Buffer.Handle()))
	}

	v.LoaderObj.VkUpdateDescriptorSetWithTemplate(
		descriptorSet.DeviceHandle(),
		descriptorSet.Handle(),
		template.Handle(),
		infoUnsafe,
	)
}

func (v *DeviceVulkanDriver) UpdateDescriptorSetWithTemplateFromObjectHandle(descriptorSet types.DescriptorSet, template types.DescriptorUpdateTemplate, data loader.VulkanHandle) {
	if descriptorSet.Handle() == 0 {
		panic("descriptorSet cannot be uninitialized")
	}
	if template.Handle() == 0 {
		panic("template cannot be uninitialized")
	}

	v.LoaderObj.VkUpdateDescriptorSetWithTemplate(
		descriptorSet.DeviceHandle(),
		descriptorSet.Handle(),
		template.Handle(),
		unsafe.Pointer(data),
	)
}
