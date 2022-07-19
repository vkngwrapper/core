package core1_2

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/common/extensions"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
	"time"
)

// VulkanDevice is an implementation of the Device interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDevice struct {
	core1_1.Device

	DeviceDriver      driver.Driver
	DeviceHandle      driver.VkDevice
	MaximumAPIVersion common.APIVersion
}

// PromoteDevice accepts a Device object from any core version. If provided a device that supports
// at least core 1.2, it will return a core1_2.Device. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanDevice, even if it is provided a VulkanDevice from a higher
// core version. Two Vulkan 1.2 compatible Device objects with the same Device.Handle will
// return the same interface value when passed to this method.
func PromoteDevice(device core1_0.Device) Device {
	if device == nil {
		return nil
	}
	if !device.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedDevice := core1_1.PromoteDevice(device)

	return device.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(device.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanDevice{
				Device: promotedDevice,

				DeviceDriver:      device.Driver(),
				DeviceHandle:      device.Handle(),
				MaximumAPIVersion: device.APIVersion(),
			}
		}).(Device)
}

var _ = PromoteDevice(nil)

func (d *VulkanDevice) CreateRenderPass2(allocator *driver.AllocationCallbacks, options RenderPassCreateInfo2) (core1_0.RenderPass, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var renderPassHandle driver.VkRenderPass
	res, err := d.DeviceDriver.VkCreateRenderPass2(
		d.DeviceHandle,
		(*driver.VkRenderPassCreateInfo2)(infoPtr),
		allocator.Handle(),
		&renderPassHandle,
	)
	if err != nil {
		return nil, res, err
	}

	renderPass := extensions.CreateRenderPassObject(
		d.DeviceDriver,
		d.DeviceHandle,
		renderPassHandle,
		d.MaximumAPIVersion,
	)

	return renderPass, res, nil
}

func (d *VulkanDevice) GetBufferDeviceAddress(o BufferDeviceAddressInfo) (uint64, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := d.DeviceDriver.VkGetBufferDeviceAddress(
		d.DeviceHandle,
		(*driver.VkBufferDeviceAddressInfo)(info),
	)
	return uint64(address), nil
}

func (d *VulkanDevice) GetBufferOpaqueCaptureAddress(o BufferDeviceAddressInfo) (uint64, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := d.DeviceDriver.VkGetBufferOpaqueCaptureAddress(
		d.DeviceHandle,
		(*driver.VkBufferDeviceAddressInfo)(info),
	)
	return uint64(address), nil
}

func (d *VulkanDevice) GetDeviceMemoryOpaqueCaptureAddress(o DeviceMemoryOpaqueCaptureAddressInfo) (uint64, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := d.DeviceDriver.VkGetDeviceMemoryOpaqueCaptureAddress(
		d.DeviceHandle,
		(*driver.VkDeviceMemoryOpaqueCaptureAddressInfo)(info),
	)
	return uint64(address), nil
}

func (d *VulkanDevice) SignalSemaphore(o SemaphoreSignalInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	signalPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return d.DeviceDriver.VkSignalSemaphore(
		d.DeviceHandle,
		(*driver.VkSemaphoreSignalInfo)(signalPtr),
	)
}

func (d *VulkanDevice) WaitSemaphores(timeout time.Duration, o SemaphoreWaitInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	waitPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return d.DeviceDriver.VkWaitSemaphores(
		d.DeviceHandle,
		(*driver.VkSemaphoreWaitInfo)(waitPtr),
		driver.Uint64(common.TimeoutNanoseconds(timeout)),
	)
}
