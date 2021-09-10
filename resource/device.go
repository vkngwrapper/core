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
type Device struct {
	loader *loader.Loader
	handle loader.VkDevice
}

func (d *Device) Loader() *loader.Loader {
	return d.loader
}

func (d *Device) Handle() loader.VkDevice {
	return d.handle
}

func (d *Device) Destroy() error {
	return d.loader.VkDestroyDevice(d.handle, nil)
}

func (d *Device) GetQueue(queueFamilyIndex int, queueIndex int) (*Queue, error) {
	var queueHandle loader.VkQueue

	err := d.loader.VkGetDeviceQueue(d.handle, loader.Uint32(queueFamilyIndex), loader.Uint32(queueIndex), &queueHandle)
	if err != nil {
		return nil, err
	}

	return &Queue{loader: d.loader, handle: queueHandle}, nil
}

func (d *Device) CreateShaderModule(allocator cgoalloc.Allocator, o *ShaderModuleOptions) (*ShaderModule, loader.VkResult, error) {
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

	return &ShaderModule{loader: d.loader, handle: shaderModule, device: d.handle}, res, nil
}

func (d *Device) CreateImageView(allocator cgoalloc.Allocator, o *ImageViewOptions) (*ImageView, loader.VkResult, error) {
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

	return &ImageView{loader: d.loader, handle: imageViewHandle, device: d.handle}, res, nil
}

func (d *Device) CreateSemaphore(allocator cgoalloc.Allocator, o *SemaphoreOptions) (*Semaphore, loader.VkResult, error) {
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

	return &Semaphore{loader: d.loader, device: d.handle, handle: semaphoreHandle}, res, nil
}

func (d *Device) CreateFence(allocator cgoalloc.Allocator, o *FenceOptions) (*Fence, loader.VkResult, error) {
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

	return &Fence{loader: d.loader, device: d.handle, handle: fenceHandle}, res, nil
}

func (d *Device) WaitForIdle() (loader.VkResult, error) {
	return d.loader.VkDeviceWaitIdle(d.handle)
}

func (d *Device) WaitForFences(allocator cgoalloc.Allocator, waitForAll bool, timeout time.Duration, fences []*Fence) (loader.VkResult, error) {
	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))
	defer allocator.Free(fenceUnsafePtr)

	fencePtr := (*loader.VkFence)(fenceUnsafePtr)

	fenceSlice := ([]loader.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].handle
	}

	waitAllConst := C.VK_FALSE
	if waitForAll {
		waitAllConst = C.VK_TRUE
	}

	return d.loader.VkWaitForFences(d.handle, loader.Uint32(fenceCount), fencePtr, loader.VkBool32(waitAllConst), loader.Uint64(core.TimeoutNanoseconds(timeout)))
}

func (d *Device) ResetFences(allocator cgoalloc.Allocator, fences []*Fence) (loader.VkResult, error) {
	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))
	defer allocator.Free(fenceUnsafePtr)

	fencePtr := (*loader.VkFence)(fenceUnsafePtr)
	fenceSlice := ([]loader.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].handle
	}

	return d.loader.VkResetFences(d.handle, loader.Uint32(fenceCount), fencePtr)
}

func (d *Device) CreateBuffer(allocator cgoalloc.Allocator, o *BufferOptions) (*Buffer, loader.VkResult, error) {
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

	return &Buffer{loader: d.loader, handle: buffer, device: d.handle}, res, nil
}

func (d *Device) AllocateMemory(allocator cgoalloc.Allocator, o *DeviceMemoryOptions) (*DeviceMemory, loader.VkResult, error) {
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

	return &DeviceMemory{
		loader: d.loader,
		device: d.handle,
		handle: deviceMemory,
	}, res, nil
}
