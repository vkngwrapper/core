package core

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type BufferOptions struct {
	BufferSize         int
	Usages             VKng.BufferUsages
	SharingMode        VKng.SharingMode
	QueueFamilyIndices []int

	Next VKng.Options
}

func (o *BufferOptions) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	createInfo := (*C.VkBufferCreateInfo)(allocator.Malloc(C.sizeof_struct_VkBufferCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
	createInfo.flags = 0
	createInfo.size = C.size_t(o.BufferSize)
	createInfo.usage = C.VkBufferUsageFlags(o.Usages)
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

	var next unsafe.Pointer
	var err error

	if o.Next != nil {
		next, err = o.Next.AllocForC(allocator)
	}
	if err != nil {
		return nil, err
	}
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}

type BufferHandle C.VkBuffer
type Buffer struct {
	device C.VkDevice
	handle C.VkBuffer
}

func (b *Buffer) Handle() BufferHandle {
	return BufferHandle(b.handle)
}

func (b *Buffer) Destroy() {
	C.vkDestroyBuffer(b.device, b.handle, nil)
}

func (b *Buffer) MemoryRequirements(allocator cgoalloc.Allocator) *VKng.MemoryRequirements {
	requirementsUnsafe := allocator.Malloc(C.sizeof_struct_VkMemoryRequirements)
	defer allocator.Free(requirementsUnsafe)

	requirements := (*C.VkMemoryRequirements)(requirementsUnsafe)

	C.vkGetBufferMemoryRequirements(b.device, b.handle, requirements)

	return &VKng.MemoryRequirements{
		Size:       int(requirements.size),
		Alignment:  int(requirements.alignment),
		MemoryType: uint32(requirements.memoryTypeBits),
	}
}

func (b *Buffer) BindBufferMemory(memory *DeviceMemory, offset int) (VKng.Result, error) {
	res := VKng.Result(C.vkBindBufferMemory(b.device, b.handle, memory.handle, C.VkDeviceSize(offset)))
	return res, res.ToError()
}
