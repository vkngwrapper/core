package impl1_1

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
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
)

// VulkanDevice is an implementation of the Device interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDevice struct {
	impl1_0.VulkanDevice

	DeviceObjectBuilder core1_1.DeviceObjectBuilder
}

func (d *VulkanDevice) BindBufferMemory2(o []core1_1.BindBufferMemoryInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptionSlice[C.VkBindBufferMemoryInfo, core1_1.BindBufferMemoryInfo](arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return d.Driver().VkBindBufferMemory2(d.Handle(),
		driver.Uint32(len(o)),
		(*driver.VkBindBufferMemoryInfo)(unsafe.Pointer(optionPtr)),
	)
}

func (d *VulkanDevice) BindImageMemory2(o []core1_1.BindImageMemoryInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptionSlice[C.VkBindImageMemoryInfo, core1_1.BindImageMemoryInfo](arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return d.Driver().VkBindImageMemory2(
		d.Handle(),
		driver.Uint32(len(o)),
		(*driver.VkBindImageMemoryInfo)(unsafe.Pointer(optionPtr)),
	)
}

func (d *VulkanDevice) BufferMemoryRequirements2(o core1_1.BufferMemoryRequirementsInfo2, out *core1_1.MemoryRequirements2) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	d.Driver().VkGetBufferMemoryRequirements2(d.Handle(),
		(*driver.VkBufferMemoryRequirementsInfo2)(optionPtr),
		(*driver.VkMemoryRequirements2)(outDataPtr),
	)

	return common.PopulateOutData(out, outDataPtr)
}

func (d *VulkanDevice) ImageMemoryRequirements2(o core1_1.ImageMemoryRequirementsInfo2, out *core1_1.MemoryRequirements2) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	d.Driver().VkGetImageMemoryRequirements2(d.Handle(),
		(*driver.VkImageMemoryRequirementsInfo2)(optionPtr),
		(*driver.VkMemoryRequirements2)(outDataPtr),
	)

	return common.PopulateOutData(out, outDataPtr)
}

func (d *VulkanDevice) ImageSparseMemoryRequirements2(o core1_1.ImageSparseMemoryRequirementsInfo2, outDataFactory func() *core1_1.SparseImageMemoryRequirements2) ([]*core1_1.SparseImageMemoryRequirements2, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, err
	}

	requirementCountPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	d.Driver().VkGetImageSparseMemoryRequirements2(d.Handle(),
		(*driver.VkImageSparseMemoryRequirementsInfo2)(optionPtr),
		requirementCountPtr,
		nil,
	)

	count := int(*requirementCountPtr)
	if count == 0 {
		return nil, nil
	}

	outDataSlice := make([]*core1_1.SparseImageMemoryRequirements2, count)
	for i := 0; i < count; i++ {
		if outDataFactory != nil {
			outDataSlice[i] = outDataFactory()
		} else {
			outDataSlice[i] = &core1_1.SparseImageMemoryRequirements2{}
		}
	}

	outDataPtr, err := common.AllocOutDataHeaderSlice[C.VkSparseImageMemoryRequirements2, *core1_1.SparseImageMemoryRequirements2](arena, outDataSlice)
	if err != nil {
		return nil, err
	}

	castOutDataPtr := (*C.VkSparseImageMemoryRequirements2)(outDataPtr)

	d.Driver().VkGetImageSparseMemoryRequirements2(d.Handle(),
		(*driver.VkImageSparseMemoryRequirementsInfo2)(optionPtr),
		requirementCountPtr,
		(*driver.VkSparseImageMemoryRequirements2)(unsafe.Pointer(castOutDataPtr)),
	)

	err = common.PopulateOutDataSlice[C.VkSparseImageMemoryRequirements2, *core1_1.SparseImageMemoryRequirements2](outDataSlice, unsafe.Pointer(outDataPtr))
	if err != nil {
		return nil, err
	}

	return outDataSlice, nil
}

func (d *VulkanDevice) DescriptorSetLayoutSupport(o core1_0.DescriptorSetLayoutCreateInfo, outData *core1_1.DescriptorSetLayoutSupport) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionsPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOutDataHeader(arena, outData)
	if err != nil {
		return err
	}

	d.Driver().VkGetDescriptorSetLayoutSupport(d.Handle(),
		(*driver.VkDescriptorSetLayoutCreateInfo)(optionsPtr),
		(*driver.VkDescriptorSetLayoutSupport)(outDataPtr))

	return common.PopulateOutData(outData, outDataPtr)
}

func (d *VulkanDevice) DeviceGroupPeerMemoryFeatures(heapIndex, localDeviceIndex, remoteDeviceIndex int) core1_1.PeerMemoryFeatureFlags {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	featuresPtr := (*driver.VkPeerMemoryFeatureFlags)(arena.Malloc(int(unsafe.Sizeof(C.VkPeerMemoryFeatureFlags(0)))))

	d.Driver().VkGetDeviceGroupPeerMemoryFeatures(
		d.Handle(),
		driver.Uint32(heapIndex),
		driver.Uint32(localDeviceIndex),
		driver.Uint32(remoteDeviceIndex),
		featuresPtr,
	)

	return core1_1.PeerMemoryFeatureFlags(*featuresPtr)
}

func (d *VulkanDevice) CreateDescriptorUpdateTemplate(o core1_1.DescriptorUpdateTemplateCreateInfo, allocator *driver.AllocationCallbacks) (core1_1.DescriptorUpdateTemplate, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfoPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var templateHandle driver.VkDescriptorUpdateTemplate
	res, err := d.Driver().VkCreateDescriptorUpdateTemplate(d.Handle(),
		(*driver.VkDescriptorUpdateTemplateCreateInfo)(createInfoPtr),
		allocator.Handle(),
		&templateHandle,
	)
	if err != nil {
		return nil, res, err
	}

	return d.DeviceObjectBuilder.CreateDescriptorUpdateTemplate(d.Driver(), d.Handle(), templateHandle, d.APIVersion()), res, nil
}

func (d *VulkanDevice) CreateSamplerYcbcrConversion(o core1_1.SamplerYcbcrConversionCreateInfo, allocator *driver.AllocationCallbacks) (core1_1.SamplerYcbcrConversion, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var ycbcrHandle driver.VkSamplerYcbcrConversion
	res, err := d.Driver().VkCreateSamplerYcbcrConversion(
		d.Handle(),
		(*driver.VkSamplerYcbcrConversionCreateInfo)(optionPtr),
		allocator.Handle(),
		&ycbcrHandle,
	)
	if err != nil {
		return nil, res, err
	}

	return d.DeviceObjectBuilder.CreateSamplerYcbcrConversion(d.Driver(), d.Handle(), ycbcrHandle, d.APIVersion()), res, nil
}

func (d *VulkanDevice) GetQueue2(o core1_1.DeviceQueueInfo2) (core1_0.Queue, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, err
	}

	var queue driver.VkQueue
	d.Driver().VkGetDeviceQueue2(
		d.Handle(),
		(*driver.VkDeviceQueueInfo2)(optionPtr),
		&queue,
	)

	return d.VulkanDevice.DeviceObjectBuilder.CreateQueueObject(d.Driver(), d.Handle(), queue, d.APIVersion()), nil
}
