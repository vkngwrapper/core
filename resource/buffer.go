package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type BufferOptions struct {
	BufferSize         int
	Usages             core.BufferUsages
	SharingMode        core.SharingMode
	QueueFamilyIndices []int

	Next core.Options
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

type Buffer struct {
	loader *loader.Loader
	device loader.VkDevice
	handle loader.VkBuffer
}

func (b *Buffer) Handle() loader.VkBuffer {
	return b.handle
}

func (b *Buffer) Destroy() error {
	return b.loader.VkDestroyBuffer(b.device, b.handle, nil)
}

func (b *Buffer) MemoryRequirements(allocator cgoalloc.Allocator) (*core.MemoryRequirements, error) {
	requirementsUnsafe := allocator.Malloc(C.sizeof_struct_VkMemoryRequirements)
	defer allocator.Free(requirementsUnsafe)

	err := b.loader.VkGetBufferMemoryRequirements(b.device, b.handle, (*loader.VkMemoryRequirements)(requirementsUnsafe))
	if err != nil {
		return nil, err
	}

	requirements := (*C.VkMemoryRequirements)(requirementsUnsafe)

	return &core.MemoryRequirements{
		Size:       int(requirements.size),
		Alignment:  int(requirements.alignment),
		MemoryType: uint32(requirements.memoryTypeBits),
	}, nil
}

func (b *Buffer) BindBufferMemory(memory *DeviceMemory, offset int) (loader.VkResult, error) {
	return b.loader.VkBindBufferMemory(b.device, b.handle, memory.handle, loader.VkDeviceSize(offset))
}
