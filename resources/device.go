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
