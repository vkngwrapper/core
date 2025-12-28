package impl1_2

import (
	"fmt"
	"time"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) CreateRenderPass2(allocator *loader.AllocationCallbacks, options core1_2.RenderPassCreateInfo2) (core.RenderPass, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, options)
	if err != nil {
		return core.RenderPass{}, core1_0.VKErrorUnknown, err
	}

	var renderPassHandle loader.VkRenderPass
	res, err := v.LoaderObj.VkCreateRenderPass2(
		v.DeviceObj.Handle(),
		(*loader.VkRenderPassCreateInfo2)(infoPtr),
		allocator.Handle(),
		&renderPassHandle,
	)
	if err != nil {
		return core.RenderPass{}, res, err
	}

	renderPass := core.InternalRenderPass(
		v.DeviceObj.Handle(),
		renderPassHandle,
		v.DeviceObj.APIVersion(),
	)

	return renderPass, res, nil
}

func (v *DeviceVulkanDriver) GetBufferDeviceAddress(o core1_2.BufferDeviceAddressInfo) (uint64, error) {
	if o.Buffer.Handle() == 0 {
		return 0, fmt.Errorf("o.Buffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := v.LoaderObj.VkGetBufferDeviceAddress(
		o.Buffer.DeviceHandle(),
		(*loader.VkBufferDeviceAddressInfo)(info),
	)
	return uint64(address), nil
}

func (v *DeviceVulkanDriver) GetBufferOpaqueCaptureAddress(o core1_2.BufferDeviceAddressInfo) (uint64, error) {
	if o.Buffer.Handle() == 0 {
		return 0, fmt.Errorf("o.Buffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := v.LoaderObj.VkGetBufferOpaqueCaptureAddress(
		o.Buffer.DeviceHandle(),
		(*loader.VkBufferDeviceAddressInfo)(info),
	)
	return uint64(address), nil
}

func (v *DeviceVulkanDriver) GetDeviceMemoryOpaqueCaptureAddress(o core1_2.DeviceMemoryOpaqueCaptureAddressInfo) (uint64, error) {
	if o.Memory.Handle() == 0 {
		return 0, fmt.Errorf("o.Memory cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := v.LoaderObj.VkGetDeviceMemoryOpaqueCaptureAddress(
		o.Memory.DeviceHandle(),
		(*loader.VkDeviceMemoryOpaqueCaptureAddressInfo)(info),
	)
	return uint64(address), nil
}

func (v *DeviceVulkanDriver) SignalSemaphore(o core1_2.SemaphoreSignalInfo) (common.VkResult, error) {
	if o.Semaphore.Handle() == 0 {
		return core1_0.VKErrorUnknown, fmt.Errorf("o.Semaphore cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	signalPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return v.LoaderObj.VkSignalSemaphore(
		o.Semaphore.DeviceHandle(),
		(*loader.VkSemaphoreSignalInfo)(signalPtr),
	)
}

func (v *DeviceVulkanDriver) WaitSemaphores(timeout time.Duration, o core1_2.SemaphoreWaitInfo) (common.VkResult, error) {
	if len(o.Semaphores) == 0 {
		return core1_0.VKSuccess, nil
	}

	for i, semaphore := range o.Semaphores {
		if semaphore.Handle() == 0 {
			return core1_0.VKErrorUnknown, fmt.Errorf("semaphore values cannot be uninitialized but semaphore %d is uninitialized", i)
		}
		if semaphore.DeviceHandle() != o.Semaphores[0].DeviceHandle() {
			return core1_0.VKErrorUnknown, fmt.Errorf("all Semaphore values must be owned by the same Device, but Semaphore %d is owned by a different Device from Semaphore 0", i)
		}
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	waitPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return v.LoaderObj.VkWaitSemaphores(
		o.Semaphores[0].DeviceHandle(),
		(*loader.VkSemaphoreWaitInfo)(waitPtr),
		loader.Uint64(common.TimeoutNanoseconds(timeout)),
	)
}
