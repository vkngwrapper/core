package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/cgoalloc"
	"time"
	"unsafe"
)

type DeviceHandle C.VkDevice
type vulkanDevice struct {
	loader *loader.Loader
	handle loader.VkDevice
}

func (d *vulkanDevice) Loader() *loader.Loader {
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

func (d *vulkanDevice) CreateShaderModule(allocator cgoalloc.Allocator, o *ShaderModuleOptions) (ShaderModule, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
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

func (d *vulkanDevice) CreateImageView(allocator cgoalloc.Allocator, o *ImageViewOptions) (ImageView, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
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

func (d *vulkanDevice) CreateSemaphore(allocator cgoalloc.Allocator, o *SemaphoreOptions) (Semaphore, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
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

func (d *vulkanDevice) CreateFence(allocator cgoalloc.Allocator, o *FenceOptions) (Fence, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
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

func (d *vulkanDevice) WaitForFences(allocator cgoalloc.Allocator, waitForAll bool, timeout time.Duration, fences []Fence) (loader.VkResult, error) {
	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))
	defer allocator.Free(fenceUnsafePtr)

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

func (d *vulkanDevice) ResetFences(allocator cgoalloc.Allocator, fences []Fence) (loader.VkResult, error) {
	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))
	defer allocator.Free(fenceUnsafePtr)

	fencePtr := (*loader.VkFence)(fenceUnsafePtr)
	fenceSlice := ([]loader.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].Handle()
	}

	return d.loader.VkResetFences(d.handle, loader.Uint32(fenceCount), fencePtr)
}

func (d *vulkanDevice) CreateBuffer(allocator cgoalloc.Allocator, o *BufferOptions) (Buffer, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
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

func (d *vulkanDevice) AllocateMemory(allocator cgoalloc.Allocator, o *DeviceMemoryOptions) (DeviceMemory, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
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
