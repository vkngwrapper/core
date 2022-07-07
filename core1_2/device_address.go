package core1_2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"unsafe"
)

const (
	BufferCreateDeviceAddressCaptureReplay core1_0.BufferCreateFlags = C.VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT

	BufferUsageShaderDeviceAddress core1_0.BufferUsageFlags = C.VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT

	MemoryAllocateDeviceAddress              core1_1.MemoryAllocateFlags = C.VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT
	MemoryAllocateDeviceAddressCaptureReplay core1_1.MemoryAllocateFlags = C.VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT

	VkErrorInvalidOpaqueCaptureAddress common.VkResult = C.VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS
)

func init() {
	BufferCreateDeviceAddressCaptureReplay.Register("Device Address (Capture/Replay)")

	BufferUsageShaderDeviceAddress.Register("Shader Device Address")

	MemoryAllocateDeviceAddress.Register("Device Address")
	MemoryAllocateDeviceAddressCaptureReplay.Register("Device Address (Capture/Replay)")

	VkErrorInvalidOpaqueCaptureAddress.Register("invalid opaque capture address")
}

////

type BufferDeviceAddressInfo struct {
	Buffer core1_0.Buffer

	common.NextOptions
}

func (o BufferDeviceAddressInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBufferDeviceAddressInfo{})))
	}

	if o.Buffer == nil {
		return nil, errors.New("core1_2.DeviceMemoryAddressOptions.Buffer cannot be nil")
	}

	info := (*C.VkBufferDeviceAddressInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO
	info.pNext = next
	info.buffer = C.VkBuffer(unsafe.Pointer(o.Buffer.Handle()))

	return preallocatedPointer, nil
}

func (o BufferDeviceAddressInfo) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkBufferDeviceAddressInfo)(cDataPointer)
	return info.pNext, nil
}

////

type DeviceMemoryOpaqueCaptureAddressInfo struct {
	Memory core1_0.DeviceMemory

	common.NextOptions
}

func (o DeviceMemoryOpaqueCaptureAddressInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceMemoryOpaqueCaptureAddressInfo{})))
	}

	info := (*C.VkDeviceMemoryOpaqueCaptureAddressInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_MEMORY_OPAQUE_CAPTURE_ADDRESS_INFO
	info.pNext = next
	info.memory = C.VkDeviceMemory(unsafe.Pointer(o.Memory.Handle()))

	return preallocatedPointer, nil
}

////

type BufferOpaqueCaptureAddressCreateInfo struct {
	OpaqueCaptureAddress uint64

	common.NextOptions
}

func (o BufferOpaqueCaptureAddressCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBufferOpaqueCaptureAddressCreateInfo{})))
	}

	info := (*C.VkBufferOpaqueCaptureAddressCreateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_BUFFER_OPAQUE_CAPTURE_ADDRESS_CREATE_INFO
	info.pNext = next
	info.opaqueCaptureAddress = C.uint64_t(o.OpaqueCaptureAddress)

	return preallocatedPointer, nil
}

////

type MemoryOpaqueCaptureAddressAllocateInfo struct {
	OpaqueCaptureAddress uint64

	common.NextOptions
}

func (o MemoryOpaqueCaptureAddressAllocateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryOpaqueCaptureAddressAllocateInfo{})))
	}

	info := (*C.VkMemoryOpaqueCaptureAddressAllocateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_MEMORY_OPAQUE_CAPTURE_ADDRESS_ALLOCATE_INFO
	info.pNext = next
	info.opaqueCaptureAddress = C.uint64_t(o.OpaqueCaptureAddress)

	return preallocatedPointer, nil
}
