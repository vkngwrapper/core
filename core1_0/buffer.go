package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

const (
	UsageTransferSrc        common.BufferUsages = C.VK_BUFFER_USAGE_TRANSFER_SRC_BIT
	UsageTransferDst        common.BufferUsages = C.VK_BUFFER_USAGE_TRANSFER_DST_BIT
	UsageUniformTexelBuffer common.BufferUsages = C.VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT
	UsageStorageTexelBuffer common.BufferUsages = C.VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT
	UsageUniformBuffer      common.BufferUsages = C.VK_BUFFER_USAGE_UNIFORM_BUFFER_BIT
	UsageStorageBuffer      common.BufferUsages = C.VK_BUFFER_USAGE_STORAGE_BUFFER_BIT
	UsageIndexBuffer        common.BufferUsages = C.VK_BUFFER_USAGE_INDEX_BUFFER_BIT
	UsageVertexBuffer       common.BufferUsages = C.VK_BUFFER_USAGE_VERTEX_BUFFER_BIT
	UsageIndirectBuffer     common.BufferUsages = C.VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT

	SharingExclusive  common.SharingMode = C.VK_SHARING_MODE_EXCLUSIVE
	SharingConcurrent common.SharingMode = C.VK_SHARING_MODE_CONCURRENT
)

func init() {
	UsageTransferSrc.Register("Transfer Source")
	UsageTransferDst.Register("Transfer Destination")
	UsageUniformTexelBuffer.Register("Uniform Texel Buffer")
	UsageStorageTexelBuffer.Register("Storage Texel Buffer")
	UsageUniformBuffer.Register("Uniform Buffer")
	UsageStorageBuffer.Register("Storage Buffer")
	UsageIndexBuffer.Register("Index Buffer")
	UsageVertexBuffer.Register("Vertex Buffer")
	UsageIndirectBuffer.Register("Indirect Buffer")

	SharingExclusive.Register("Exclusive")
	SharingConcurrent.Register("Concurrent")
}

type BufferOptions struct {
	BufferSize         int
	Usage              common.BufferUsages
	SharingMode        common.SharingMode
	QueueFamilyIndices []int

	common.HaveNext
}

func (o BufferOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkBufferCreateInfo)
	}
	createInfo := (*C.VkBufferCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next
	createInfo.size = C.VkDeviceSize(o.BufferSize)
	createInfo.usage = C.VkBufferUsageFlags(o.Usage)
	createInfo.sharingMode = C.VkSharingMode(o.SharingMode)

	queueFamilyCount := len(o.QueueFamilyIndices)
	createInfo.queueFamilyIndexCount = C.uint32_t(queueFamilyCount)
	createInfo.pQueueFamilyIndices = nil

	if queueFamilyCount > 0 {
		indicesPtr := (*C.uint32_t)(allocator.Malloc(queueFamilyCount * int(unsafe.Sizeof(C.uint32_t(0)))))
		indicesSlice := ([]C.uint32_t)(unsafe.Slice(indicesPtr, queueFamilyCount))

		for i := 0; i < queueFamilyCount; i++ {
			indicesSlice[i] = C.uint32_t(o.QueueFamilyIndices[i])
		}

		createInfo.pQueueFamilyIndices = indicesPtr
	}

	return preallocatedPointer, nil
}
