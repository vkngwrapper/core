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
	BufferCreateSparseBinding   common.BufferCreateFlags = C.VK_BUFFER_CREATE_SPARSE_BINDING_BIT
	BufferCreateSparseResidency common.BufferCreateFlags = C.VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT
	BufferCreateSparseAliased   common.BufferCreateFlags = C.VK_BUFFER_CREATE_SPARSE_ALIASED_BIT

	BufferUsageTransferSrc        common.BufferUsages = C.VK_BUFFER_USAGE_TRANSFER_SRC_BIT
	BufferUsageTransferDst        common.BufferUsages = C.VK_BUFFER_USAGE_TRANSFER_DST_BIT
	BufferUsageUniformTexelBuffer common.BufferUsages = C.VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT
	BufferUsageStorageTexelBuffer common.BufferUsages = C.VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT
	BufferUsageUniformBuffer      common.BufferUsages = C.VK_BUFFER_USAGE_UNIFORM_BUFFER_BIT
	BufferUsageStorageBuffer      common.BufferUsages = C.VK_BUFFER_USAGE_STORAGE_BUFFER_BIT
	BufferUsageIndexBuffer        common.BufferUsages = C.VK_BUFFER_USAGE_INDEX_BUFFER_BIT
	BufferUsageVertexBuffer       common.BufferUsages = C.VK_BUFFER_USAGE_VERTEX_BUFFER_BIT
	BufferUsageIndirectBuffer     common.BufferUsages = C.VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT

	SharingExclusive  common.SharingMode = C.VK_SHARING_MODE_EXCLUSIVE
	SharingConcurrent common.SharingMode = C.VK_SHARING_MODE_CONCURRENT
)

func init() {
	BufferCreateSparseBinding.Register("Sparse Binding")
	BufferCreateSparseResidency.Register("Sparse Residency")
	BufferCreateSparseAliased.Register("Sparse Aliased")

	BufferUsageTransferSrc.Register("Transfer Source")
	BufferUsageTransferDst.Register("Transfer Destination")
	BufferUsageUniformTexelBuffer.Register("Uniform Texel Buffer")
	BufferUsageStorageTexelBuffer.Register("Storage Texel Buffer")
	BufferUsageUniformBuffer.Register("Uniform Buffer")
	BufferUsageStorageBuffer.Register("Storage Buffer")
	BufferUsageIndexBuffer.Register("Index Buffer")
	BufferUsageVertexBuffer.Register("Vertex Buffer")
	BufferUsageIndirectBuffer.Register("Indirect Buffer")

	SharingExclusive.Register("Exclusive")
	SharingConcurrent.Register("Concurrent")
}

type BufferOptions struct {
	Flags              common.BufferCreateFlags
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
	createInfo.flags = C.VkBufferCreateFlags(o.Flags)
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

func (o BufferOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkBufferCreateInfo)(cDataPointer)
	return createInfo.pNext, nil
}
