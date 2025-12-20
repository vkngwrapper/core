package impl1_2

import (
	"time"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
)

// VulkanDevice is an implementation of the Device interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDevice struct {
	impl1_1.VulkanDevice
}

func (d *VulkanDevice) CreateRenderPass2(allocator *driver.AllocationCallbacks, options core1_2.RenderPassCreateInfo2) (core1_0.RenderPass, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var renderPassHandle driver.VkRenderPass
	res, err := d.Driver().VkCreateRenderPass2(
		d.Handle(),
		(*driver.VkRenderPassCreateInfo2)(infoPtr),
		allocator.Handle(),
		&renderPassHandle,
	)
	if err != nil {
		return nil, res, err
	}

	renderPass := d.VulkanDevice.VulkanDevice.DeviceObjectBuilder.CreateRenderPassObject(
		d.Driver(),
		d.Handle(),
		renderPassHandle,
		d.APIVersion(),
	)

	return renderPass, res, nil
}

func (d *VulkanDevice) GetBufferDeviceAddress(o core1_2.BufferDeviceAddressInfo) (uint64, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := d.Driver().VkGetBufferDeviceAddress(
		d.Handle(),
		(*driver.VkBufferDeviceAddressInfo)(info),
	)
	return uint64(address), nil
}

func (d *VulkanDevice) GetBufferOpaqueCaptureAddress(o core1_2.BufferDeviceAddressInfo) (uint64, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := d.Driver().VkGetBufferOpaqueCaptureAddress(
		d.Handle(),
		(*driver.VkBufferDeviceAddressInfo)(info),
	)
	return uint64(address), nil
}

func (d *VulkanDevice) GetDeviceMemoryOpaqueCaptureAddress(o core1_2.DeviceMemoryOpaqueCaptureAddressInfo) (uint64, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := d.Driver().VkGetDeviceMemoryOpaqueCaptureAddress(
		d.Handle(),
		(*driver.VkDeviceMemoryOpaqueCaptureAddressInfo)(info),
	)
	return uint64(address), nil
}

func (d *VulkanDevice) SignalSemaphore(o core1_2.SemaphoreSignalInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	signalPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return d.Driver().VkSignalSemaphore(
		d.Handle(),
		(*driver.VkSemaphoreSignalInfo)(signalPtr),
	)
}

func (d *VulkanDevice) WaitSemaphores(timeout time.Duration, o core1_2.SemaphoreWaitInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	waitPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return d.Driver().VkWaitSemaphores(
		d.Handle(),
		(*driver.VkSemaphoreWaitInfo)(waitPtr),
		driver.Uint64(common.TimeoutNanoseconds(timeout)),
	)
}
