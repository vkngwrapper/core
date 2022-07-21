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
	// BufferCreateDeviceAddressCaptureReplay specifies that the Buffer object's address can
	// be saved and reused on a subsequent run (e.g. for trace capture and replay)
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateFlagBits.html
	BufferCreateDeviceAddressCaptureReplay core1_0.BufferCreateFlags = C.VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT

	// BufferUsageShaderDeviceAddress specifies that the Buffer can be used to retrieve a
	// Buffer device address via Device.GetBufferDeviceAddress and use that address to
	// access the Buffer object's memory from a shader
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html
	BufferUsageShaderDeviceAddress core1_0.BufferUsageFlags = C.VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT

	// MemoryAllocateDeviceAddress specifies that the memory can be attached to a Buffer object
	// created with BufferUsageShaderDeviceAddress set in Usage, and that the DeviceMemory object
	// can be used to retrieve an opaque address via Device.GetDeviceMemoryOpaqueCaptureAddress
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagBits.html
	MemoryAllocateDeviceAddress core1_1.MemoryAllocateFlags = C.VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT
	// MemoryAllocateDeviceAddressCaptureReplay specifies that the memory's address can be saved
	// and reused on a subsequent run (e.g. for trace capture and replay)
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagBits.html
	MemoryAllocateDeviceAddressCaptureReplay core1_1.MemoryAllocateFlags = C.VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT

	// VkErrorInvalidOpaqueCaptureAddress indicates a Buffer creation or memory allocation failed
	// because the requested address is not available
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

// BufferDeviceAddressInfo specifies the Buffer to query an address for
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferDeviceAddressInfo.html
type BufferDeviceAddressInfo struct {
	// Buffer specifies the Buffer whose address is beign queried
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

// DeviceMemoryOpaqueCaptureAddressInfo specifies the DeviceMemory object to query an address for
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryOpaqueCaptureAddressInfo.html
type DeviceMemoryOpaqueCaptureAddressInfo struct {
	// Memory specifies the DeviceMemory whose address is being queried
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

// BufferOpaqueCaptureAddressCreateInfo requests a specific address for a Buffer
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html
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

// MemoryOpaqueCaptureAddressAllocateInfo requests a specific address for a memory allocation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryOpaqueCaptureAddressAllocateInfoKHR.html
type MemoryOpaqueCaptureAddressAllocateInfo struct {
	// OpaqueCaptureAddress is the opaque capture address requested for the memory allocation
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
