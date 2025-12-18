package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

const (
	// BufferCreateSparseBinding specifies that the buffer will be backed using sparse memory
	// binding.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferCreateFlagBits.html
	BufferCreateSparseBinding BufferCreateFlags = C.VK_BUFFER_CREATE_SPARSE_BINDING_BIT
	// BufferCreateSparseResidency  specifies that the buffer can be partially backed using
	// sparse memory binding. Buffers created with this flag must also be created with the
	// BufferCreateSparseBinding flag.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferCreateFlagBits.html
	BufferCreateSparseResidency BufferCreateFlags = C.VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT
	// BufferCreateSparseAliased specifies that the buffer will be backed using sparse memory
	// binding with memory ranges that might also simultaneously be backing another buffer
	// (or another portion of the same buffer). Buffers created with this flag must also be
	// created with the BufferCreateSparseBinding flag.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferCreateFlagBits.html
	BufferCreateSparseAliased BufferCreateFlags = C.VK_BUFFER_CREATE_SPARSE_ALIASED_BIT

	// BufferUsageTransferSrc specifies that the buffer can be used as the source of a transfer command
	// (see the definition of PipelineStageTransfer).
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html
	BufferUsageTransferSrc BufferUsageFlags = C.VK_BUFFER_USAGE_TRANSFER_SRC_BIT
	// BufferUsageTransferDst specifies that the buffer can be used as the destination of a transfer
	// command.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html
	BufferUsageTransferDst BufferUsageFlags = C.VK_BUFFER_USAGE_TRANSFER_DST_BIT
	// BufferUsageUniformTexelBuffer  specifies that the buffer can be used to create a BufferView
	// suitable for occupying a DescriptorSet slot of type DescriptorTypeUniformTexelBuffer.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html
	BufferUsageUniformTexelBuffer BufferUsageFlags = C.VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT
	// BufferUsageStorageTexelBuffer specifies that the buffer can be used to create a BufferView
	// suitable for occupying a DescriptorSet slot of type DescriptorTypeStorageTexelBuffer.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html
	BufferUsageStorageTexelBuffer BufferUsageFlags = C.VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT
	// BufferUsageUniformBuffer specifies that the buffer can be used in a DescriptorBufferInfo
	// suitable for occupying a DescriptorSet slot either of type DescriptorTypeUniformBuffer
	// or DescriptorTypeUniformBufferDynamic
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html
	BufferUsageUniformBuffer BufferUsageFlags = C.VK_BUFFER_USAGE_UNIFORM_BUFFER_BIT
	// BufferUsageStorageBuffer specifies that the buffer can be used in a DescriptorBufferInfo
	// suitable for occupying a DescriptorSet slot either of type DescriptorTypeStorageBuffer or
	// DescriptorTypeStorageBufferDynamic.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html
	BufferUsageStorageBuffer BufferUsageFlags = C.VK_BUFFER_USAGE_STORAGE_BUFFER_BIT
	// BufferUsageIndexBuffer specifies that the buffer is suitable for passing as the buffer parameter
	// to CommandBuffer.CmdBindIndexBuffer.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html
	BufferUsageIndexBuffer BufferUsageFlags = C.VK_BUFFER_USAGE_INDEX_BUFFER_BIT
	// BufferUsageVertexBuffer specifies that the buffer is suitable for passing as an element of the
	// buffers slice to CommandBuffer.CmdBindVertexBuffers.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html
	BufferUsageVertexBuffer BufferUsageFlags = C.VK_BUFFER_USAGE_VERTEX_BUFFER_BIT
	// BufferUsageIndirectBuffer specifies that the buffer is suitable for passing as the buffer parameter
	// to CommandBuffer.CmdDrawIndirect, CommandBuffer.CmdDrawIndexedIndirect,
	// or CommandBuffer.CmdDispatchIndirect.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html
	BufferUsageIndirectBuffer BufferUsageFlags = C.VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT

	// SharingModeExclusive specifies that access to any range or image subresource of the object will be
	// exclusive to a single queue family at a time.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSharingMode.html
	SharingModeExclusive SharingMode = C.VK_SHARING_MODE_EXCLUSIVE
	// SharingModeConcurrent  specifies that concurrent access to any range or image subresource of the
	// object from multiple queue families is supported.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSharingMode.html
	SharingModeConcurrent SharingMode = C.VK_SHARING_MODE_CONCURRENT
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

	SharingModeExclusive.Register("Exclusive")
	SharingModeConcurrent.Register("Concurrent")
}

// BufferCreateInfo specifies the parameters of a newly-created buffer object
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html
type BufferCreateInfo struct {
	// Flags specifies additional parameters of the buffer
	Flags BufferCreateFlags
	// Size is the size in bytes of the buffer to be created
	Size int
	// Usage spcifies allowed usages of the buffer
	Usage BufferUsageFlags
	// SharingMode specifies the sharing mode of the buffer when it will be accessed by multiple
	// queue families
	SharingMode SharingMode
	// QueueFamilyIndices is a slice of queue families that will access this buffer. It is ignored
	// if SharingMode is not SharingModeConcurrent.
	QueueFamilyIndices []int

	// NextOptions allows additional creation option structures to be chained
	common.NextOptions
}

func (o BufferCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkBufferCreateInfo)
	}
	createInfo := (*C.VkBufferCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
	createInfo.flags = C.VkBufferCreateFlags(o.Flags)
	createInfo.pNext = next
	createInfo.size = C.VkDeviceSize(o.Size)
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
