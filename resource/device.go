package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
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

func (d *Device) CreateShaderModule(allocator cgoalloc.Allocator, o *ShaderModuleOptions) (*ShaderModule, core.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	var shaderModule C.VkShaderModule
	res := core.Result(C.vkCreateShaderModule(d.handle, (*C.VkShaderModuleCreateInfo)(createInfo), nil, &shaderModule))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &ShaderModule{handle: shaderModule, device: d.handle}, res, nil
}

func (d *Device) CreateImageView(allocator cgoalloc.Allocator, o *ImageViewOptions) (*ImageView, core.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	var imageViewHandle C.VkImageView

	res := core.Result(C.vkCreateImageView(d.handle, (*C.VkImageViewCreateInfo)(createInfo), nil, &imageViewHandle))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &ImageView{handle: imageViewHandle, device: d.handle}, res, nil
}

func (d *Device) CreateSemaphore(allocator cgoalloc.Allocator, o *SemaphoreOptions) (*Semaphore, core.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	var semaphoreHandle C.VkSemaphore

	res := core.Result(C.vkCreateSemaphore(d.handle, (*C.VkSemaphoreCreateInfo)(createInfo), nil, &semaphoreHandle))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &Semaphore{device: d.handle, handle: semaphoreHandle}, res, nil
}

func (d *Device) CreateFence(allocator cgoalloc.Allocator, o *FenceOptions) (*Fence, core.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	var fenceHandle C.VkFence

	res := core.Result(C.vkCreateFence(d.handle, (*C.VkFenceCreateInfo)(createInfo), nil, &fenceHandle))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &Fence{device: d.handle, handle: fenceHandle}, res, nil
}

func (d *Device) WaitForIdle() (core.Result, error) {
	res := core.Result(C.vkDeviceWaitIdle(d.handle))
	return res, res.ToError()
}

func (d *Device) WaitForFences(allocator cgoalloc.Allocator, waitForAll bool, timeout time.Duration, fences []*Fence) (core.Result, error) {
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

	res := core.Result(C.vkWaitForFences(d.handle, C.uint32_t(fenceCount), fencePtr, C.uint(waitAllConst), C.uint64_t(core.TimeoutNanoseconds(timeout))))
	return res, res.ToError()
}

func (d *Device) ResetFences(allocator cgoalloc.Allocator, fences []*Fence) (core.Result, error) {
	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))
	defer allocator.Free(fenceUnsafePtr)

	fencePtr := (*C.VkFence)(fenceUnsafePtr)
	fenceSlice := ([]C.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].handle
	}

	res := core.Result(C.vkResetFences(d.handle, C.uint32_t(fenceCount), fencePtr))
	return res, res.ToError()
}

func (d *Device) CreateBuffer(allocator cgoalloc.Allocator, o *BufferOptions) (*Buffer, core.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	var buffer C.VkBuffer

	res := core.Result(C.vkCreateBuffer(d.handle, (*C.VkBufferCreateInfo)(createInfo), nil, &buffer))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &Buffer{handle: buffer, device: d.handle}, res, nil
}

func (d *Device) AllocateMemory(allocator cgoalloc.Allocator, o *DeviceMemoryOptions) (*DeviceMemory, core.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	var deviceMemory C.VkDeviceMemory

	res := core.Result(C.vkAllocateMemory(d.handle, (*C.VkMemoryAllocateInfo)(createInfo), nil, &deviceMemory))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &DeviceMemory{
		device: d.handle,
		handle: deviceMemory,
	}, res, nil
}
