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

// VulkanDevice is an implementation of the Device interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDevice struct {
	core1_0.Device

	DeviceDriver      driver.Driver
	DeviceHandle      driver.VkDevice
	MaximumAPIVersion common.APIVersion
}

// PromoteDevice accepts a Device object from any core version. If provided a device that supports
// at least core 1.1, it will return a core1_1.Device. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanDevice, even if it is provided a VulkanDevice from a higher
// core version. Two Vulkan 1.1 compatible Device objects with the same Device.Handle will
// return the same interface value when passed to this method.
func PromoteDevice(device core1_0.Device) Device {
	if device == nil {
		return nil
	}
	if !device.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}
	promoted, alreadyPromoted := device.(Device)
	if alreadyPromoted {
		return promoted
	}

	return device.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(device.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanDevice{
				Device:            device,
				DeviceDriver:      device.Driver(),
				DeviceHandle:      device.Handle(),
				MaximumAPIVersion: device.APIVersion(),
			}
		}).(Device)
}

func (d *VulkanDevice) BindBufferMemory2(o []BindBufferMemoryInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptionSlice[C.VkBindBufferMemoryInfo, BindBufferMemoryInfo](arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return d.DeviceDriver.VkBindBufferMemory2(d.DeviceHandle,
		driver.Uint32(len(o)),
		(*driver.VkBindBufferMemoryInfo)(unsafe.Pointer(optionPtr)),
	)
}

func (d *VulkanDevice) BindImageMemory2(o []BindImageMemoryInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptionSlice[C.VkBindImageMemoryInfo, BindImageMemoryInfo](arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return d.DeviceDriver.VkBindImageMemory2(
		d.DeviceHandle,
		driver.Uint32(len(o)),
		(*driver.VkBindImageMemoryInfo)(unsafe.Pointer(optionPtr)),
	)
}

func (d *VulkanDevice) BufferMemoryRequirements2(o BufferMemoryRequirementsInfo2, out *MemoryRequirements2) error {
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

	d.DeviceDriver.VkGetBufferMemoryRequirements2(d.DeviceHandle,
		(*driver.VkBufferMemoryRequirementsInfo2)(optionPtr),
		(*driver.VkMemoryRequirements2)(outDataPtr),
	)

	return common.PopulateOutData(out, outDataPtr)
}

func (d *VulkanDevice) ImageMemoryRequirements2(o ImageMemoryRequirementsInfo2, out *MemoryRequirements2) error {
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

	d.DeviceDriver.VkGetImageMemoryRequirements2(d.DeviceHandle,
		(*driver.VkImageMemoryRequirementsInfo2)(optionPtr),
		(*driver.VkMemoryRequirements2)(outDataPtr),
	)

	return common.PopulateOutData(out, outDataPtr)
}

func (d *VulkanDevice) ImageSparseMemoryRequirements2(o ImageSparseMemoryRequirementsInfo2, outDataFactory func() *SparseImageMemoryRequirements2) ([]*SparseImageMemoryRequirements2, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, err
	}

	requirementCountPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	d.DeviceDriver.VkGetImageSparseMemoryRequirements2(d.DeviceHandle,
		(*driver.VkImageSparseMemoryRequirementsInfo2)(optionPtr),
		requirementCountPtr,
		nil,
	)

	count := int(*requirementCountPtr)
	if count == 0 {
		return nil, nil
	}

	outDataSlice := make([]*SparseImageMemoryRequirements2, count)
	for i := 0; i < count; i++ {
		if outDataFactory != nil {
			outDataSlice[i] = outDataFactory()
		} else {
			outDataSlice[i] = &SparseImageMemoryRequirements2{}
		}
	}

	outDataPtr, err := common.AllocOutDataHeaderSlice[C.VkSparseImageMemoryRequirements2, *SparseImageMemoryRequirements2](arena, outDataSlice)
	if err != nil {
		return nil, err
	}

	castOutDataPtr := (*C.VkSparseImageMemoryRequirements2)(outDataPtr)

	d.DeviceDriver.VkGetImageSparseMemoryRequirements2(d.DeviceHandle,
		(*driver.VkImageSparseMemoryRequirementsInfo2)(optionPtr),
		requirementCountPtr,
		(*driver.VkSparseImageMemoryRequirements2)(unsafe.Pointer(castOutDataPtr)),
	)

	err = common.PopulateOutDataSlice[C.VkSparseImageMemoryRequirements2, *SparseImageMemoryRequirements2](outDataSlice, unsafe.Pointer(outDataPtr))
	if err != nil {
		return nil, err
	}

	return outDataSlice, nil
}

func (d *VulkanDevice) DescriptorSetLayoutSupport(o core1_0.DescriptorSetLayoutCreateInfo, outData *DescriptorSetLayoutSupport) error {
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

	d.DeviceDriver.VkGetDescriptorSetLayoutSupport(d.DeviceHandle,
		(*driver.VkDescriptorSetLayoutCreateInfo)(optionsPtr),
		(*driver.VkDescriptorSetLayoutSupport)(outDataPtr))

	return common.PopulateOutData(outData, outDataPtr)
}

func (d *VulkanDevice) DeviceGroupPeerMemoryFeatures(heapIndex, localDeviceIndex, remoteDeviceIndex int) PeerMemoryFeatureFlags {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	featuresPtr := (*driver.VkPeerMemoryFeatureFlags)(arena.Malloc(int(unsafe.Sizeof(C.VkPeerMemoryFeatureFlags(0)))))

	d.DeviceDriver.VkGetDeviceGroupPeerMemoryFeatures(
		d.DeviceHandle,
		driver.Uint32(heapIndex),
		driver.Uint32(localDeviceIndex),
		driver.Uint32(remoteDeviceIndex),
		featuresPtr,
	)

	return PeerMemoryFeatureFlags(*featuresPtr)
}

func (d *VulkanDevice) CreateDescriptorUpdateTemplate(o DescriptorUpdateTemplateCreateInfo, allocator *driver.AllocationCallbacks) (DescriptorUpdateTemplate, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfoPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var templateHandle driver.VkDescriptorUpdateTemplate
	res, err := d.DeviceDriver.VkCreateDescriptorUpdateTemplate(d.DeviceHandle,
		(*driver.VkDescriptorUpdateTemplateCreateInfo)(createInfoPtr),
		allocator.Handle(),
		&templateHandle,
	)
	if err != nil {
		return nil, res, err
	}

	return CreateDescriptorUpdateTemplate(d.DeviceDriver, d.DeviceHandle, templateHandle, d.MaximumAPIVersion), res, nil
}

func (d *VulkanDevice) CreateSamplerYcbcrConversion(o SamplerYcbcrConversionCreateInfo, allocator *driver.AllocationCallbacks) (SamplerYcbcrConversion, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var ycbcrHandle driver.VkSamplerYcbcrConversion
	res, err := d.DeviceDriver.VkCreateSamplerYcbcrConversion(
		d.DeviceHandle,
		(*driver.VkSamplerYcbcrConversionCreateInfo)(optionPtr),
		allocator.Handle(),
		&ycbcrHandle,
	)
	if err != nil {
		return nil, res, err
	}

	return CreateSamplerYcbcrConversion(d.DeviceDriver, d.DeviceHandle, ycbcrHandle, d.MaximumAPIVersion), res, nil
}

//go:linkname createQueueObject github.com/vkngwrapper/core/v3/core1_0.createQueueObject
func createQueueObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueue, version common.APIVersion) *core1_0.VulkanQueue

func (d *VulkanDevice) GetQueue2(o DeviceQueueInfo2) (core1_0.Queue, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, err
	}

	var queue driver.VkQueue
	d.DeviceDriver.VkGetDeviceQueue2(
		d.DeviceHandle,
		(*driver.VkDeviceQueueInfo2)(optionPtr),
		&queue,
	)

	return createQueueObject(d.DeviceDriver, d.DeviceHandle, queue, d.MaximumAPIVersion), nil
}
