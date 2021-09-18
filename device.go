package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"time"
	"unsafe"
)

type DeviceHandle C.VkDevice
type vulkanDevice struct {
	driver Driver
	handle VkDevice
}

func (d *vulkanDevice) Driver() Driver {
	return d.driver
}

func (d *vulkanDevice) Handle() VkDevice {
	return d.handle
}

func (d *vulkanDevice) Destroy() error {
	return d.driver.VkDestroyDevice(d.handle, nil)
}

func (d *vulkanDevice) GetQueue(queueFamilyIndex int, queueIndex int) (Queue, error) {
	var queueHandle VkQueue

	err := d.driver.VkGetDeviceQueue(d.handle, Uint32(queueFamilyIndex), Uint32(queueIndex), &queueHandle)
	if err != nil {
		return nil, err
	}

	return &vulkanQueue{driver: d.driver, handle: queueHandle}, nil
}

func (d *vulkanDevice) CreateShaderModule(o *ShaderModuleOptions) (ShaderModule, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var shaderModule VkShaderModule
	res, err := d.driver.VkCreateShaderModule(d.handle, (*VkShaderModuleCreateInfo)(createInfo), nil, &shaderModule)
	if err != nil {
		return nil, res, err
	}

	return &vulkanShaderModule{driver: d.driver, handle: shaderModule, device: d.handle}, res, nil
}

func (d *vulkanDevice) CreateImageView(o *ImageViewOptions) (ImageView, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var imageViewHandle VkImageView

	res, err := d.driver.VkCreateImageView(d.handle, (*VkImageViewCreateInfo)(createInfo), nil, &imageViewHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanImageView{driver: d.driver, handle: imageViewHandle, device: d.handle}, res, nil
}

func (d *vulkanDevice) CreateSemaphore(o *SemaphoreOptions) (Semaphore, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var semaphoreHandle VkSemaphore

	res, err := d.driver.VkCreateSemaphore(d.handle, (*VkSemaphoreCreateInfo)(createInfo), nil, &semaphoreHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanSemaphore{driver: d.driver, device: d.handle, handle: semaphoreHandle}, res, nil
}

func (d *vulkanDevice) CreateFence(o *FenceOptions) (Fence, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var fenceHandle VkFence

	res, err := d.driver.VkCreateFence(d.handle, (*VkFenceCreateInfo)(createInfo), nil, &fenceHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanFence{driver: d.driver, device: d.handle, handle: fenceHandle}, res, nil
}

func (d *vulkanDevice) WaitForIdle() (VkResult, error) {
	return d.driver.VkDeviceWaitIdle(d.handle)
}

func (d *vulkanDevice) WaitForFences(waitForAll bool, timeout time.Duration, fences []Fence) (VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*VkFence)(fenceUnsafePtr)

	fenceSlice := ([]VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].Handle()
	}

	waitAllConst := C.VK_FALSE
	if waitForAll {
		waitAllConst = C.VK_TRUE
	}

	return d.driver.VkWaitForFences(d.handle, Uint32(fenceCount), fencePtr, VkBool32(waitAllConst), Uint64(common.TimeoutNanoseconds(timeout)))
}

func (d *vulkanDevice) ResetFences(fences []Fence) (VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*VkFence)(fenceUnsafePtr)
	fenceSlice := ([]VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].Handle()
	}

	return d.driver.VkResetFences(d.handle, Uint32(fenceCount), fencePtr)
}

func (d *vulkanDevice) CreateBuffer(o *BufferOptions) (Buffer, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var buffer VkBuffer

	res, err := d.driver.VkCreateBuffer(d.handle, (*VkBufferCreateInfo)(createInfo), nil, &buffer)
	if err != nil {
		return nil, res, err
	}

	return &vulkanBuffer{driver: d.driver, handle: buffer, device: d.handle}, res, nil
}

func (d *vulkanDevice) AllocateMemory(o *DeviceMemoryOptions) (DeviceMemory, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var deviceMemory VkDeviceMemory

	res, err := d.driver.VkAllocateMemory(d.handle, (*VkMemoryAllocateInfo)(createInfo), nil, &deviceMemory)
	if err != nil {
		return nil, res, err
	}

	return &vulkanDeviceMemory{
		driver: d.driver,
		device: d.handle,
		handle: deviceMemory,
	}, res, nil
}

func (d *vulkanDevice) CreateDescriptorSetLayout(o *DescriptorSetLayoutOptions) (DescriptorSetLayout, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var descriptorSetLayout VkDescriptorSetLayout

	res, err := d.driver.VkCreateDescriptorSetLayout(d.handle, (*VkDescriptorSetLayoutCreateInfo)(createInfo), nil, &descriptorSetLayout)
	if err != nil {
		return nil, res, err
	}

	return &vulkanDescriptorSetLayout{
		driver: d.driver,
		device: d.handle,
		handle: descriptorSetLayout,
	}, res, nil
}

func (d *vulkanDevice) CreateDescriptorPool(o *DescriptorPoolOptions) (DescriptorPool, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var descriptorPool VkDescriptorPool

	res, err := d.driver.VkCreateDescriptorPool(d.handle, (*VkDescriptorPoolCreateInfo)(createInfo), nil, &descriptorPool)
	if err != nil {
		return nil, res, err
	}

	return &vulkanDescriptorPool{
		driver: d.driver,
		handle: descriptorPool,
		device: d.handle,
	}, res, nil
}

func (d *vulkanDevice) AllocateDescriptorSet(o *DescriptorSetOptions) ([]DescriptorSet, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	setCount := len(o.AllocationLayouts)
	descriptorSets := (*VkDescriptorSet)(arena.Malloc(setCount * int(unsafe.Sizeof([1]C.VkDescriptorSet{}))))

	res, err := d.driver.VkAllocateDescriptorSets(d.handle, (*VkDescriptorSetAllocateInfo)(createInfo), descriptorSets)
	if err != nil {
		return nil, res, err
	}

	var sets []DescriptorSet
	descriptorSetSlice := ([]VkDescriptorSet)(unsafe.Slice(descriptorSets, setCount))
	for i := 0; i < setCount; i++ {
		sets = append(sets, &vulkanDescriptorSet{handle: descriptorSetSlice[i]})
	}

	return sets, res, nil
}

func (d *vulkanDevice) UpdateDescriptorSets(writes []*WriteDescriptorSetOptions, copies []*CopyDescriptorSetOptions) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	writeCount := len(writes)
	copyCount := len(copies)

	var writePtr unsafe.Pointer
	var copyPtr unsafe.Pointer

	if writeCount > 0 {
		writePtr = arena.Malloc(writeCount * C.sizeof_struct_VkWriteDescriptorSet)
		writeSlice := ([]C.VkWriteDescriptorSet)(unsafe.Slice((*C.VkWriteDescriptorSet)(writePtr), writeCount))
		for i := 0; i < writeCount; i++ {
			next, err := common.AllocNext(arena, writes[i])
			if err != nil {
				return err
			}

			writes[i].populate(arena, &(writeSlice[i]), next)
		}
	}

	if copyCount > 0 {
		copyPtr = arena.Malloc(copyCount * C.sizeof_struct_VkCopyDescriptorSet)
		copySlice := ([]C.VkCopyDescriptorSet)(unsafe.Slice((*C.VkCopyDescriptorSet)(copyPtr), copyCount))
		for i := 0; i < copyCount; i++ {
			next, err := common.AllocNext(arena, copies[i])
			if err != nil {
				return err
			}

			copies[i].populate(arena, &(copySlice[i]), next)
		}
	}

	return d.driver.VkUpdateDescriptorSets(d.handle, Uint32(writeCount), (*VkWriteDescriptorSet)(writePtr), Uint32(copyCount), (*VkCopyDescriptorSet)(copyPtr))
}
