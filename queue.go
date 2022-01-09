package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"strings"
	"unsafe"
)

type vulkanQueue struct {
	driver Driver
	handle VkQueue
}

func (q *vulkanQueue) Handle() VkQueue {
	return q.handle
}

func (q *vulkanQueue) Driver() Driver {
	return q.driver
}

func (q *vulkanQueue) WaitForIdle() (VkResult, error) {
	return q.driver.VkQueueWaitIdle(q.handle)
}

type SparseMemoryBindFlags int32

const (
	SparseMemoryBindMetadata SparseMemoryBindFlags = C.VK_SPARSE_MEMORY_BIND_METADATA_BIT
)

var sparseMemoryBindFlagsToString = map[SparseMemoryBindFlags]string{
	SparseMemoryBindMetadata: "Metadata",
}

func (f SparseMemoryBindFlags) String() string {
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := SparseMemoryBindFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := sparseMemoryBindFlagsToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteRune('|')
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}

type SparseMemoryBind struct {
	ResourceOffset int
	Size           int

	Memory       DeviceMemory
	MemoryOffset int

	Flags SparseMemoryBindFlags
}

func (b *SparseMemoryBind) populate(bind *C.VkSparseMemoryBind) {
	bind.resourceOffset = C.VkDeviceSize(b.ResourceOffset)
	bind.size = C.VkDeviceSize(b.Size)
	bind.memory = C.VkDeviceMemory(b.Memory.Handle())
	bind.memoryOffset = C.VkDeviceSize(b.MemoryOffset)
	bind.flags = C.VkSparseMemoryBindFlags(b.Flags)
}

type SparseBufferMemoryBindInfo struct {
	Buffer Buffer
	Binds  []SparseMemoryBind
}

type SparseImageOpaqueMemoryBindInfo struct {
	Image Image
	Binds []SparseMemoryBind
}

type SparseImageMemoryBind struct {
	Subresource common.ImageSubresource
	Offset      common.Offset3D
	Extent      common.Extent3D

	Memory       DeviceMemory
	MemoryOffset int

	Flags SparseMemoryBindFlags
}

type SparseImageMemoryBindInfo struct {
	Image Image
	Binds []SparseImageMemoryBind
}

type BindSparseOptions struct {
	WaitSemaphores   []Semaphore
	SignalSemaphores []Semaphore

	BufferBinds      []SparseBufferMemoryBindInfo
	ImageOpaqueBinds []SparseImageOpaqueMemoryBindInfo
	ImageBinds       []SparseImageMemoryBindInfo

	common.HaveNext
}

func (b *BindSparseOptions) populate(allocator *cgoparam.Allocator, createInfo *C.VkBindSparseInfo, next unsafe.Pointer) error {
	createInfo.sType = C.VK_STRUCTURE_TYPE_BIND_SPARSE_INFO
	createInfo.pNext = next

	waitSemaphoreCount := len(b.WaitSemaphores)
	signalSemaphoreCount := len(b.SignalSemaphores)
	bufferBindCount := len(b.BufferBinds)
	imageOpaqueBindCount := len(b.ImageOpaqueBinds)
	imageBindCount := len(b.ImageBinds)

	createInfo.waitSemaphoreCount = C.uint32_t(waitSemaphoreCount)
	createInfo.pWaitSemaphores = nil
	createInfo.bufferBindCount = C.uint32_t(bufferBindCount)
	createInfo.pBufferBinds = nil
	createInfo.imageOpaqueBindCount = C.uint32_t(imageOpaqueBindCount)
	createInfo.pImageOpaqueBinds = nil
	createInfo.imageBindCount = C.uint32_t(imageBindCount)
	createInfo.pImageBinds = nil
	createInfo.signalSemaphoreCount = C.uint32_t(signalSemaphoreCount)
	createInfo.pSignalSemaphores = nil

	if waitSemaphoreCount > 0 {
		waitSemaphorePtr := (*C.VkSemaphore)(allocator.Malloc(waitSemaphoreCount * int(unsafe.Sizeof([1]C.VkSemaphore{}))))
		waitSemaphoreSlice := ([]C.VkSemaphore)(unsafe.Slice(waitSemaphorePtr, waitSemaphoreCount))

		for i := 0; i < waitSemaphoreCount; i++ {
			waitSemaphoreSlice[i] = C.VkSemaphore(b.WaitSemaphores[i].Handle())
		}

		createInfo.pWaitSemaphores = waitSemaphorePtr
	}

	if bufferBindCount > 0 {
		bufferBindPtr := (*C.VkSparseBufferMemoryBindInfo)(allocator.Malloc(bufferBindCount * C.sizeof_struct_VkSparseBufferMemoryBindInfo))
		bufferBindSlice := ([]C.VkSparseBufferMemoryBindInfo)(unsafe.Slice(bufferBindPtr, bufferBindCount))

		for i := 0; i < bufferBindCount; i++ {
			bufferBindSlice[i].buffer = C.VkBuffer(b.BufferBinds[i].Buffer.Handle())
			bindCount := len(b.BufferBinds[i].Binds)
			bufferBindSlice[i].bindCount = C.uint32_t(bindCount)
			bufferBindSlice[i].pBinds = nil

			if bindCount > 0 {
				bindPtr := (*C.VkSparseMemoryBind)(allocator.Malloc(bindCount * C.sizeof_struct_VkSparseMemoryBind))
				bindSlice := ([]C.VkSparseMemoryBind)(unsafe.Slice(bindPtr, bindCount))

				for j := 0; j < bindCount; j++ {
					b.BufferBinds[i].Binds[j].populate(&bindSlice[j])
				}
				bufferBindSlice[i].pBinds = bindPtr
			}
		}

		createInfo.pBufferBinds = bufferBindPtr
	}

	if imageOpaqueBindCount > 0 {
		imageOpaqueBindPtr := (*C.VkSparseImageOpaqueMemoryBindInfo)(allocator.Malloc(imageOpaqueBindCount * C.sizeof_struct_VkSparseImageOpaqueMemoryBindInfo))
		imageOpaqueBindSlice := ([]C.VkSparseImageOpaqueMemoryBindInfo)(unsafe.Slice(imageOpaqueBindPtr, imageOpaqueBindCount))

		for i := 0; i < imageOpaqueBindCount; i++ {
			imageOpaqueBindSlice[i].image = C.VkImage(b.ImageOpaqueBinds[i].Image.Handle())
			bindCount := len(b.ImageOpaqueBinds[i].Binds)
			imageOpaqueBindSlice[i].bindCount = C.uint32_t(bindCount)
			imageOpaqueBindSlice[i].pBinds = nil

			if bindCount > 0 {
				bindPtr := (*C.VkSparseMemoryBind)(allocator.Malloc(bindCount * C.sizeof_struct_VkSparseMemoryBind))
				bindSlice := ([]C.VkSparseMemoryBind)(unsafe.Slice(bindPtr, bindCount))

				for j := 0; j < bindCount; j++ {
					b.ImageOpaqueBinds[i].Binds[j].populate(&bindSlice[j])
				}
				imageOpaqueBindSlice[i].pBinds = bindPtr
			}
		}

		createInfo.pImageOpaqueBinds = imageOpaqueBindPtr
	}

	if imageBindCount > 0 {
		imageBindPtr := (*C.VkSparseImageMemoryBindInfo)(allocator.Malloc(imageBindCount * C.sizeof_struct_VkSparseImageMemoryBindInfo))
		imageBindSlice := ([]C.VkSparseImageMemoryBindInfo)(unsafe.Slice(imageBindPtr, imageBindCount))

		for i := 0; i < imageBindCount; i++ {
			imageBindSlice[i].image = C.VkImage(b.ImageBinds[i].Image.Handle())
			bindCount := len(b.ImageBinds[i].Binds)
			imageBindSlice[i].bindCount = C.uint32_t(bindCount)
			imageBindSlice[i].pBinds = nil

			if bindCount > 0 {
				bindPtr := (*C.VkSparseImageMemoryBind)(allocator.Malloc(bindCount * C.sizeof_struct_VkSparseImageMemoryBind))
				bindSlice := ([]C.VkSparseImageMemoryBind)(unsafe.Slice(bindPtr, bindCount))

				for j := 0; j < bindCount; j++ {
					bindSlice[j].subresource.aspectMask = C.VkImageAspectFlags(b.ImageBinds[i].Binds[j].Subresource.AspectMask)
					bindSlice[j].subresource.mipLevel = C.uint32_t(b.ImageBinds[i].Binds[j].Subresource.MipLevel)
					bindSlice[j].subresource.arrayLayer = C.uint32_t(b.ImageBinds[i].Binds[j].Subresource.ArrayLayer)
					bindSlice[j].offset.x = C.int32_t(b.ImageBinds[i].Binds[j].Offset.X)
					bindSlice[j].offset.y = C.int32_t(b.ImageBinds[i].Binds[j].Offset.Y)
					bindSlice[j].offset.z = C.int32_t(b.ImageBinds[i].Binds[j].Offset.Z)
					bindSlice[j].extent.width = C.uint32_t(b.ImageBinds[i].Binds[j].Extent.Width)
					bindSlice[j].extent.height = C.uint32_t(b.ImageBinds[i].Binds[j].Extent.Height)
					bindSlice[j].extent.depth = C.uint32_t(b.ImageBinds[i].Binds[j].Extent.Depth)
					bindSlice[j].memory = C.VkDeviceMemory(b.ImageBinds[i].Binds[j].Memory.Handle())
					bindSlice[j].memoryOffset = C.VkDeviceSize(b.ImageBinds[i].Binds[j].MemoryOffset)
					bindSlice[j].flags = C.VkSparseMemoryBindFlags(b.ImageBinds[i].Binds[j].Flags)
				}

				imageBindSlice[i].pBinds = bindPtr
			}
		}

		createInfo.pImageBinds = imageBindPtr
	}

	if signalSemaphoreCount > 0 {
		signalSemaphorePtr := (*C.VkSemaphore)(allocator.Malloc(signalSemaphoreCount * int(unsafe.Sizeof([1]C.VkSemaphore{}))))
		signalSemaphoreSlice := ([]C.VkSemaphore)(unsafe.Slice(signalSemaphorePtr, signalSemaphoreCount))

		for i := 0; i < signalSemaphoreCount; i++ {
			signalSemaphoreSlice[i] = C.VkSemaphore(b.SignalSemaphores[i].Handle())
		}

		createInfo.pSignalSemaphores = signalSemaphorePtr
	}

	return nil
}

func (b *BindSparseOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkBindSparseInfo)(allocator.Malloc(C.sizeof_struct_VkBindSparseInfo))

	err := b.populate(allocator, createInfo, next)
	if err != nil {
		return nil, err
	}

	return unsafe.Pointer(createInfo), nil
}

func (q *vulkanQueue) BindSparse(fence Fence, bindInfos []*BindSparseOptions) (VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	var fenceHandle VkFence
	if fence != nil {
		fenceHandle = fence.Handle()
	}

	bindInfoCount := len(bindInfos)
	bindInfoPtr := (*C.VkBindSparseInfo)(arena.Malloc(bindInfoCount * C.sizeof_struct_VkBindSparseInfo))
	bindInfoSlice := ([]C.VkBindSparseInfo)(unsafe.Slice(bindInfoPtr, bindInfoCount))

	for i := 0; i < bindInfoCount; i++ {
		next, err := common.AllocNext(arena, bindInfos[i])
		if err != nil {
			return VKErrorUnknown, err
		}

		err = bindInfos[i].populate(arena, &bindInfoSlice[i], next)
		if err != nil {
			return VKErrorUnknown, err
		}
	}

	return q.driver.VkQueueBindSparse(q.handle, Uint32(bindInfoCount), (*VkBindSparseInfo)(bindInfoPtr), fenceHandle)
}
