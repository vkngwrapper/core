package resources

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type BufferOptions struct {
	BufferSize         int
	Usages             core.BufferUsages
	SharingMode        core.SharingMode
	QueueFamilyIndices []int

	core.HaveNext
}

func (o *BufferOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkBufferCreateInfo)(allocator.Malloc(C.sizeof_struct_VkBufferCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next
	createInfo.size = C.VkDeviceSize(o.BufferSize)
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

	return unsafe.Pointer(createInfo), nil
}

type vulkanBuffer struct {
	loader loader.Loader
	device loader.VkDevice
	handle loader.VkBuffer
}

func (b *vulkanBuffer) Handle() loader.VkBuffer {
	return b.handle
}

func (b *vulkanBuffer) Destroy() error {
	return b.loader.VkDestroyBuffer(b.device, b.handle, nil)
}

func (b *vulkanBuffer) MemoryRequirements() (*core.MemoryRequirements, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	requirementsUnsafe := allocator.Malloc(C.sizeof_struct_VkMemoryRequirements)

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

func (b *vulkanBuffer) BindBufferMemory(memory DeviceMemory, offset int) (loader.VkResult, error) {
	return b.loader.VkBindBufferMemory(b.device, b.handle, memory.Handle(), loader.VkDeviceSize(offset))
}
