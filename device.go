package core

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/cgoalloc"
	"time"
	"unsafe"
)

type DeviceHandle C.VkDevice
type Device struct {
	handle C.VkDevice
}

func (d *Device) Handle() DeviceHandle {
	return DeviceHandle(d.handle)
}

func (d *Device) Destroy() {
	C.vkDestroyDevice(d.handle, nil)
}

func (d *Device) GetQueue(queueFamilyIndex int, queueIndex int) (*Queue, error) {
	var queueHandle C.VkQueue

	C.vkGetDeviceQueue(d.handle, C.uint32_t(queueFamilyIndex), C.uint32_t(queueIndex), &queueHandle)

	return &Queue{handle: QueueHandle(queueHandle)}, nil
}

func (d *Device) CreateShaderModule(allocator cgoalloc.Allocator, o *ShaderModuleOptions) (*ShaderModule, VKng.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, VKng.VKErrorUnknown, err
	}

	var shaderModule C.VkShaderModule
	res := VKng.Result(C.vkCreateShaderModule(d.handle, (*C.VkShaderModuleCreateInfo)(createInfo), nil, &shaderModule))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &ShaderModule{handle: shaderModule, device: d.handle}, res, nil
}

func (d *Device) CreateImageView(allocator cgoalloc.Allocator, o *ImageViewOptions) (*ImageView, VKng.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, VKng.VKErrorUnknown, err
	}

	var imageViewHandle C.VkImageView

	res := VKng.Result(C.vkCreateImageView(d.handle, (*C.VkImageViewCreateInfo)(createInfo), nil, &imageViewHandle))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &ImageView{handle: imageViewHandle, device: d.handle}, res, nil
}

func (d *Device) CreateSemaphore(allocator cgoalloc.Allocator, o *SemaphoreOptions) (*Semaphore, VKng.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, VKng.VKErrorUnknown, err
	}

	var semaphoreHandle C.VkSemaphore

	res := VKng.Result(C.vkCreateSemaphore(d.handle, (*C.VkSemaphoreCreateInfo)(createInfo), nil, &semaphoreHandle))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &Semaphore{device: d.handle, handle: semaphoreHandle}, res, nil
}

func (d *Device) CreateFence(allocator cgoalloc.Allocator, o *FenceOptions) (*Fence, VKng.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, VKng.VKErrorUnknown, err
	}

	var fenceHandle C.VkFence

	res := VKng.Result(C.vkCreateFence(d.handle, (*C.VkFenceCreateInfo)(createInfo), nil, &fenceHandle))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &Fence{device: d.handle, handle: fenceHandle}, res, nil
}

func (d *Device) WaitForIdle() (VKng.Result, error) {
	res := VKng.Result(C.vkDeviceWaitIdle(d.handle))
	return res, res.ToError()
}

func (d *Device) WaitForFences(allocator cgoalloc.Allocator, waitForAll bool, timeout time.Duration, fences []*Fence) (VKng.Result, error) {
	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))
	defer allocator.Free(fenceUnsafePtr)

	fencePtr := (*C.VkFence)(fenceUnsafePtr)

	fenceSlice := ([]C.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].handle
	}

	waitAllConst := C.VK_FALSE
	if waitForAll {
		waitAllConst = C.VK_TRUE
	}

	res := VKng.Result(C.vkWaitForFences(d.handle, C.uint32_t(fenceCount), fencePtr, C.uint(waitAllConst), C.uint64_t(VKng.TimeoutNanoseconds(timeout))))
	return res, res.ToError()
}

func (d *Device) ResetFences(allocator cgoalloc.Allocator, fences []*Fence) (VKng.Result, error) {
	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))
	defer allocator.Free(fenceUnsafePtr)

	fencePtr := (*C.VkFence)(fenceUnsafePtr)
	fenceSlice := ([]C.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].handle
	}

	res := VKng.Result(C.vkResetFences(d.handle, C.uint32_t(fenceCount), fencePtr))
	return res, res.ToError()
}

func (d *Device) CreateBuffer(allocator cgoalloc.Allocator, o *BufferOptions) (*Buffer, VKng.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, VKng.VKErrorUnknown, err
	}

	var buffer C.VkBuffer

	res := VKng.Result(C.vkCreateBuffer(d.handle, (*C.VkBufferCreateInfo)(createInfo), nil, &buffer))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &Buffer{handle: buffer, device: d.handle}, res, nil
}

func (d *Device) AllocateMemory(allocator cgoalloc.Allocator, o *DeviceMemoryOptions) (*DeviceMemory, VKng.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, VKng.VKErrorUnknown, err
	}

	var deviceMemory C.VkDeviceMemory

	res := VKng.Result(C.vkAllocateMemory(d.handle, (*C.VkMemoryAllocateInfo)(createInfo), nil, &deviceMemory))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &DeviceMemory{
		device: d.handle,
		handle: deviceMemory,
	}, res, nil
}
