package resources

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/cgoparam"
	"time"
	"unsafe"
)

type DeviceHandle C.VkDevice
type vulkanDevice struct {
	loader loader.Loader
	handle loader.VkDevice
}

func (d *vulkanDevice) Loader() loader.Loader {
	return d.loader
}

func (d *vulkanDevice) Handle() loader.VkDevice {
	return d.handle
}

func (d *vulkanDevice) Destroy() error {
	return d.loader.VkDestroyDevice(d.handle, nil)
}

func (d *vulkanDevice) GetQueue(queueFamilyIndex int, queueIndex int) (Queue, error) {
	var queueHandle loader.VkQueue

	err := d.loader.VkGetDeviceQueue(d.handle, loader.Uint32(queueFamilyIndex), loader.Uint32(queueIndex), &queueHandle)
	if err != nil {
		return nil, err
	}

	return &vulkanQueue{loader: d.loader, handle: queueHandle}, nil
}

func (d *vulkanDevice) CreateShaderModule(o *ShaderModuleOptions) (ShaderModule, loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	var shaderModule loader.VkShaderModule
	res, err := d.loader.VkCreateShaderModule(d.handle, (*loader.VkShaderModuleCreateInfo)(createInfo), nil, &shaderModule)
	if err != nil {
		return nil, res, err
	}

	return &vulkanShaderModule{loader: d.loader, handle: shaderModule, device: d.handle}, res, nil
}

func (d *vulkanDevice) CreateImageView(o *ImageViewOptions) (ImageView, loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	var imageViewHandle loader.VkImageView

	res, err := d.loader.VkCreateImageView(d.handle, (*loader.VkImageViewCreateInfo)(createInfo), nil, &imageViewHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanImageView{loader: d.loader, handle: imageViewHandle, device: d.handle}, res, nil
}

func (d *vulkanDevice) CreateSemaphore(o *SemaphoreOptions) (Semaphore, loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	var semaphoreHandle loader.VkSemaphore

	res, err := d.loader.VkCreateSemaphore(d.handle, (*loader.VkSemaphoreCreateInfo)(createInfo), nil, &semaphoreHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanSemaphore{loader: d.loader, device: d.handle, handle: semaphoreHandle}, res, nil
}

func (d *vulkanDevice) CreateFence(o *FenceOptions) (Fence, loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	var fenceHandle loader.VkFence

	res, err := d.loader.VkCreateFence(d.handle, (*loader.VkFenceCreateInfo)(createInfo), nil, &fenceHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanFence{loader: d.loader, device: d.handle, handle: fenceHandle}, res, nil
}

func (d *vulkanDevice) WaitForIdle() (loader.VkResult, error) {
	return d.loader.VkDeviceWaitIdle(d.handle)
}

func (d *vulkanDevice) WaitForFences(waitForAll bool, timeout time.Duration, fences []Fence) (loader.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*loader.VkFence)(fenceUnsafePtr)

	fenceSlice := ([]loader.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].Handle()
	}

	waitAllConst := C.VK_FALSE
	if waitForAll {
		waitAllConst = C.VK_TRUE
	}

	return d.loader.VkWaitForFences(d.handle, loader.Uint32(fenceCount), fencePtr, loader.VkBool32(waitAllConst), loader.Uint64(core.TimeoutNanoseconds(timeout)))
}

func (d *vulkanDevice) ResetFences(fences []Fence) (loader.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*loader.VkFence)(fenceUnsafePtr)
	fenceSlice := ([]loader.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].Handle()
	}

	return d.loader.VkResetFences(d.handle, loader.Uint32(fenceCount), fencePtr)
}

func (d *vulkanDevice) CreateBuffer(o *BufferOptions) (Buffer, loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	var buffer loader.VkBuffer

	res, err := d.loader.VkCreateBuffer(d.handle, (*loader.VkBufferCreateInfo)(createInfo), nil, &buffer)
	if err != nil {
		return nil, res, err
	}

	return &vulkanBuffer{loader: d.loader, handle: buffer, device: d.handle}, res, nil
}

func (d *vulkanDevice) AllocateMemory(o *DeviceMemoryOptions) (DeviceMemory, loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	var deviceMemory loader.VkDeviceMemory

	res, err := d.loader.VkAllocateMemory(d.handle, (*loader.VkMemoryAllocateInfo)(createInfo), nil, &deviceMemory)
	if err != nil {
		return nil, res, err
	}

	return &vulkanDeviceMemory{
		loader: d.loader,
		device: d.handle,
		handle: deviceMemory,
	}, res, nil
}

func (d *vulkanDevice) CreateDescriptorSetLayout(o *DescriptorSetLayoutOptions) (DescriptorSetLayout, loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	var descriptorSetLayout loader.VkDescriptorSetLayout

	res, err := d.loader.VkCreateDescriptorSetLayout(d.handle, (*loader.VkDescriptorSetLayoutCreateInfo)(createInfo), nil, &descriptorSetLayout)
	if err != nil {
		return nil, res, err
	}

	return &vulkanDescriptorSetLayout{
		loader: d.loader,
		device: d.handle,
		handle: descriptorSetLayout,
	}, res, nil
}

func (d *vulkanDevice) CreateDescriptorPool(o *DescriptorPoolOptions) (DescriptorPool, loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	var descriptorPool loader.VkDescriptorPool

	res, err := d.loader.VkCreateDescriptorPool(d.handle, (*loader.VkDescriptorPoolCreateInfo)(createInfo), nil, &descriptorPool)
	if err != nil {
		return nil, res, err
	}

	return &vulkanDescriptorPool{
		loader: d.loader,
		handle: descriptorPool,
		device: d.handle,
	}, res, nil
}

func (d *vulkanDevice) AllocateDescriptorSet(o *DescriptorSetOptions) ([]DescriptorSet, loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	setCount := len(o.AllocationLayouts)
	descriptorSets := (*loader.VkDescriptorSet)(arena.Malloc(setCount * int(unsafe.Sizeof([1]C.VkDescriptorSet{}))))

	res, err := d.loader.VkAllocateDescriptorSets(d.handle, (*loader.VkDescriptorSetAllocateInfo)(createInfo), descriptorSets)
	if err != nil {
		return nil, res, err
	}

	var sets []DescriptorSet
	descriptorSetSlice := ([]loader.VkDescriptorSet)(unsafe.Slice(descriptorSets, setCount))
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
			next, err := core.AllocNext(arena, writes[i])
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
			next, err := core.AllocNext(arena, copies[i])
			if err != nil {
				return err
			}

			copies[i].populate(arena, &(copySlice[i]), next)
		}
	}

	return d.loader.VkUpdateDescriptorSets(d.handle, loader.Uint32(writeCount), (*loader.VkWriteDescriptorSet)(writePtr), loader.Uint32(copyCount), (*loader.VkCopyDescriptorSet)(copyPtr))
}
