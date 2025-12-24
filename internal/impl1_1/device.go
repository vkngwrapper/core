package impl1_1

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *DeviceVulkanDriver) BindBufferMemory2(o ...core1_1.BindBufferMemoryInfo) (common.VkResult, error) {
	if len(o) == 0 {
		return core1_0.VKSuccess, nil
	}

	for i, info := range o {
		if info.Buffer.Handle() == 0 {
			return core1_0.VKErrorUnknown, fmt.Errorf("buffers in the list of info objects cannot be uninitialized, but buffer %d is uninitialized", i)
		}
		if info.Buffer.DeviceHandle() != o[0].Buffer.DeviceHandle() {
			return core1_0.VKErrorUnknown, fmt.Errorf("buffers in the list of info objects must belong to the same Device, but Buffer %d belongs to a different device from Buffer 0", i)
		}
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptionSlice[C.VkBindBufferMemoryInfo, core1_1.BindBufferMemoryInfo](arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return v.LoaderObj.VkBindBufferMemory2(o[0].Buffer.DeviceHandle(),
		loader.Uint32(len(o)),
		(*loader.VkBindBufferMemoryInfo)(unsafe.Pointer(optionPtr)),
	)
}

func (v *DeviceVulkanDriver) BindImageMemory2(o ...core1_1.BindImageMemoryInfo) (common.VkResult, error) {
	if len(o) == 0 {
		return core1_0.VKSuccess, nil
	}

	for i, info := range o {
		if info.Image.Handle() == 0 {
			return core1_0.VKErrorUnknown, fmt.Errorf("images in the list of info objects cannot be uninitialized but Image %d is uninitialized", i)
		}
		if info.Image.DeviceHandle() != o[0].Image.DeviceHandle() {
			return core1_0.VKErrorUnknown, fmt.Errorf("images in the list of info objects must belong to the same Device, but Image %d belongs to a different device from Image 0", i)
		}
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptionSlice[C.VkBindImageMemoryInfo, core1_1.BindImageMemoryInfo](arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return v.LoaderObj.VkBindImageMemory2(
		o[0].Image.DeviceHandle(),
		loader.Uint32(len(o)),
		(*loader.VkBindImageMemoryInfo)(unsafe.Pointer(optionPtr)),
	)
}

func (v *DeviceVulkanDriver) GetBufferMemoryRequirements2(o core1_1.BufferMemoryRequirementsInfo2, out *core1_1.MemoryRequirements2) error {
	if o.Buffer.Handle() == 0 {
		return fmt.Errorf("o.Buffer cannot be uninitialized")
	}

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

	v.LoaderObj.VkGetBufferMemoryRequirements2(o.Buffer.DeviceHandle(),
		(*loader.VkBufferMemoryRequirementsInfo2)(optionPtr),
		(*loader.VkMemoryRequirements2)(outDataPtr),
	)

	return common.PopulateOutData(out, outDataPtr)
}

func (v *DeviceVulkanDriver) GetImageMemoryRequirements2(o core1_1.ImageMemoryRequirementsInfo2, out *core1_1.MemoryRequirements2) error {
	if o.Image.Handle() == 0 {
		return fmt.Errorf("o.Image cannot be uninitialized")
	}

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

	v.LoaderObj.VkGetImageMemoryRequirements2(o.Image.DeviceHandle(),
		(*loader.VkImageMemoryRequirementsInfo2)(optionPtr),
		(*loader.VkMemoryRequirements2)(outDataPtr),
	)

	return common.PopulateOutData(out, outDataPtr)
}

func (v *DeviceVulkanDriver) GetImageSparseMemoryRequirements2(o core1_1.ImageSparseMemoryRequirementsInfo2, outDataFactory func() *core1_1.SparseImageMemoryRequirements2) ([]*core1_1.SparseImageMemoryRequirements2, error) {
	if o.Image.Handle() == 0 {
		return nil, fmt.Errorf("o.Image cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, err
	}

	requirementCountPtr := (*loader.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	v.LoaderObj.VkGetImageSparseMemoryRequirements2(o.Image.DeviceHandle(),
		(*loader.VkImageSparseMemoryRequirementsInfo2)(optionPtr),
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

	v.LoaderObj.VkGetImageSparseMemoryRequirements2(o.Image.DeviceHandle(),
		(*loader.VkImageSparseMemoryRequirementsInfo2)(optionPtr),
		requirementCountPtr,
		(*loader.VkSparseImageMemoryRequirements2)(unsafe.Pointer(castOutDataPtr)),
	)

	err = common.PopulateOutDataSlice[C.VkSparseImageMemoryRequirements2, *core1_1.SparseImageMemoryRequirements2](outDataSlice, unsafe.Pointer(outDataPtr))
	if err != nil {
		return nil, err
	}

	return outDataSlice, nil
}

func (v *DeviceVulkanDriver) GetDescriptorSetLayoutSupport(device types.Device, o core1_0.DescriptorSetLayoutCreateInfo, outData *core1_1.DescriptorSetLayoutSupport) error {
	if device.Handle() == 0 {
		return fmt.Errorf("device cannot be uninitialized")
	}
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

	v.LoaderObj.VkGetDescriptorSetLayoutSupport(device.Handle(),
		(*loader.VkDescriptorSetLayoutCreateInfo)(optionsPtr),
		(*loader.VkDescriptorSetLayoutSupport)(outDataPtr))

	return common.PopulateOutData(outData, outDataPtr)
}

func (v *DeviceVulkanDriver) GetDeviceGroupPeerMemoryFeatures(device types.Device, heapIndex, localDeviceIndex, remoteDeviceIndex int) core1_1.PeerMemoryFeatureFlags {
	if device.Handle() == 0 {
		panic("device cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	featuresPtr := (*loader.VkPeerMemoryFeatureFlags)(arena.Malloc(int(unsafe.Sizeof(C.VkPeerMemoryFeatureFlags(0)))))

	v.LoaderObj.VkGetDeviceGroupPeerMemoryFeatures(
		device.Handle(),
		loader.Uint32(heapIndex),
		loader.Uint32(localDeviceIndex),
		loader.Uint32(remoteDeviceIndex),
		featuresPtr,
	)

	return core1_1.PeerMemoryFeatureFlags(*featuresPtr)
}

func (v *DeviceVulkanDriver) CreateDescriptorUpdateTemplate(device types.Device, o core1_1.DescriptorUpdateTemplateCreateInfo, allocator *loader.AllocationCallbacks) (types.DescriptorUpdateTemplate, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.DescriptorUpdateTemplate{}, core1_0.VKErrorUnknown, fmt.Errorf("device cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfoPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.DescriptorUpdateTemplate{}, core1_0.VKErrorUnknown, err
	}

	var templateHandle loader.VkDescriptorUpdateTemplate
	res, err := v.LoaderObj.VkCreateDescriptorUpdateTemplate(device.Handle(),
		(*loader.VkDescriptorUpdateTemplateCreateInfo)(createInfoPtr),
		allocator.Handle(),
		&templateHandle,
	)
	if err != nil {
		return types.DescriptorUpdateTemplate{}, res, err
	}

	return types.InternalDescriptorUpdateTemplate(device.Handle(), templateHandle, device.APIVersion()), res, nil
}

func (v *DeviceVulkanDriver) CreateSamplerYcbcrConversion(device types.Device, o core1_1.SamplerYcbcrConversionCreateInfo, allocator *loader.AllocationCallbacks) (types.SamplerYcbcrConversion, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.SamplerYcbcrConversion{}, core1_0.VKErrorUnknown, fmt.Errorf("device cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.SamplerYcbcrConversion{}, core1_0.VKErrorUnknown, err
	}

	var ycbcrHandle loader.VkSamplerYcbcrConversion
	res, err := v.LoaderObj.VkCreateSamplerYcbcrConversion(
		device.Handle(),
		(*loader.VkSamplerYcbcrConversionCreateInfo)(optionPtr),
		allocator.Handle(),
		&ycbcrHandle,
	)
	if err != nil {
		return types.SamplerYcbcrConversion{}, res, err
	}

	return types.InternalSamplerYcbcrConversion(device.Handle(), ycbcrHandle, device.APIVersion()), res, nil
}

func (v *DeviceVulkanDriver) GetDeviceQueue2(device types.Device, o core1_1.DeviceQueueInfo2) (types.Queue, error) {
	if device.Handle() == 0 {
		return types.Queue{}, fmt.Errorf("device cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.Queue{}, err
	}

	var queue loader.VkQueue
	v.LoaderObj.VkGetDeviceQueue2(
		device.Handle(),
		(*loader.VkDeviceQueueInfo2)(optionPtr),
		&queue,
	)

	return types.InternalQueue(device.Handle(), queue, device.APIVersion()), nil
}
