package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

const (
	// SparseMemoryBindMetadata specifies that the memory being bound is only for the
	// metadata aspect
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseMemoryBindFlagBits.html
	SparseMemoryBindMetadata SparseMemoryBindFlags = C.VK_SPARSE_MEMORY_BIND_METADATA_BIT
)

func init() {
	SparseMemoryBindMetadata.Register("Metadata")
}

// SparseMemoryBind specifies a sparse memory bind operation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseMemoryBind.html
type SparseMemoryBind struct {
	// ResourceOffset is the offset into the resource
	ResourceOffset int
	// Size is the size of the memory region to be bound
	Size int

	// Memory is the DeviceMemory object that the range of the resource is bound to
	Memory DeviceMemory
	// MemoryOffset is the offset into the DeviceMemory object to bind the resource range to
	MemoryOffset int

	// Flags specifies usage of the binding operation
	Flags SparseMemoryBindFlags
}

func (b SparseMemoryBind) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkSparseMemoryBind)
	}

	bind := (*C.VkSparseMemoryBind)(preallocatedPointer)
	bind.resourceOffset = C.VkDeviceSize(b.ResourceOffset)
	bind.size = C.VkDeviceSize(b.Size)
	bind.memory = C.VkDeviceMemory(unsafe.Pointer(b.Memory.Handle()))
	bind.memoryOffset = C.VkDeviceSize(b.MemoryOffset)
	bind.flags = C.VkSparseMemoryBindFlags(b.Flags)

	return preallocatedPointer, nil
}

// SparseBufferMemoryBindInfo specifies a sparse buffer memory bind operation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseBufferMemoryBindInfo.html
type SparseBufferMemoryBindInfo struct {
	// Buffer is the Buffer object to be bound
	Buffer Buffer
	// Binds is a slice of SparseMemoryBind structures
	Binds []SparseMemoryBind
}

// SparseImageOpaqueMemoryBindInfo specifies sparse Image opaque memory bind information
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageOpaqueMemoryBindInfo.html
type SparseImageOpaqueMemoryBindInfo struct {
	// Image is the Image object to be bound
	Image Image
	// Binds is a slice of SparseMemoryBind structures
	Binds []SparseMemoryBind
}

// SparseImageMemoryBind specifies sparse Image memory bind
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryBind.html
type SparseImageMemoryBind struct {
	// Subresource is the Image aspect and region of interest in the Image
	Subresource ImageSubresource
	// Offset are the coordinates of the first texel within the Image subresource to bind
	Offset Offset3D
	// Extent is the size in texels of the region within the Image subresource to bind
	Extent Extent3D

	// Memory is the DeviceMemory object that the sparse Image blocks of the Image are bound to
	Memory DeviceMemory
	// MemoryOffset is an offset into the DeviceMemory object
	MemoryOffset int

	// Flags are sparse memory binding flags
	Flags SparseMemoryBindFlags
}

// SparseImageMemoryBindInfo specifies sparse Image memory bind information
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryBindInfo.html
type SparseImageMemoryBindInfo struct {
	// Image is the Image object to be bound
	Image Image
	// Binds is a slice of SparseImageMemoryBind structures
	Binds []SparseImageMemoryBind
}

// BindSparseInfo specifies a sparse binding operation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindSparseInfo.html
type BindSparseInfo struct {
	// WaitSemaphores is a slice of Semaphore objects upon which to wait before the sparse
	// binding operations for this batch begin execution
	WaitSemaphores []Semaphore
	// SignalSemaphores a slice of Semaphore objects which will be signaled when the sparse binding
	// operations for this batch have completed execution
	SignalSemaphores []Semaphore

	// BufferBinds is a slice of SparseBufferMemoryBindInfo structures
	BufferBinds []SparseBufferMemoryBindInfo
	// ImageOpaqueBinds is a slice of SparseImageOpaqueBindInfo structures, indicating opaque
	// sparse Image bindings to perform
	ImageOpaqueBinds []SparseImageOpaqueMemoryBindInfo
	// ImageBinds is a slice of SparseImageMemoryBindInfo structures, indicating sparse Image
	// bindings to perform
	ImageBinds []SparseImageMemoryBindInfo

	common.NextOptions
}

func (b BindSparseInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkBindSparseInfo)
	}
	var err error
	createInfo := (*C.VkBindSparseInfo)(preallocatedPointer)
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
			waitSemaphoreSlice[i] = C.VkSemaphore(unsafe.Pointer(b.WaitSemaphores[i].Handle()))
		}

		createInfo.pWaitSemaphores = waitSemaphorePtr
	}

	if bufferBindCount > 0 {
		bufferBindPtr := (*C.VkSparseBufferMemoryBindInfo)(allocator.Malloc(bufferBindCount * C.sizeof_struct_VkSparseBufferMemoryBindInfo))
		bufferBindSlice := ([]C.VkSparseBufferMemoryBindInfo)(unsafe.Slice(bufferBindPtr, bufferBindCount))

		for i := 0; i < bufferBindCount; i++ {
			bufferBindSlice[i].buffer = C.VkBuffer(unsafe.Pointer(b.BufferBinds[i].Buffer.Handle()))
			bindCount := len(b.BufferBinds[i].Binds)
			bufferBindSlice[i].bindCount = C.uint32_t(bindCount)
			bufferBindSlice[i].pBinds = nil

			if bindCount > 0 {
				bufferBindSlice[i].pBinds, err = common.AllocSlice[C.VkSparseMemoryBind, SparseMemoryBind](allocator, b.BufferBinds[i].Binds)
				if err != nil {
					return nil, err
				}
			}
		}

		createInfo.pBufferBinds = bufferBindPtr
	}

	if imageOpaqueBindCount > 0 {
		imageOpaqueBindPtr := (*C.VkSparseImageOpaqueMemoryBindInfo)(allocator.Malloc(imageOpaqueBindCount * C.sizeof_struct_VkSparseImageOpaqueMemoryBindInfo))
		imageOpaqueBindSlice := ([]C.VkSparseImageOpaqueMemoryBindInfo)(unsafe.Slice(imageOpaqueBindPtr, imageOpaqueBindCount))

		for i := 0; i < imageOpaqueBindCount; i++ {
			imageOpaqueBindSlice[i].image = C.VkImage(unsafe.Pointer(b.ImageOpaqueBinds[i].Image.Handle()))
			bindCount := len(b.ImageOpaqueBinds[i].Binds)
			imageOpaqueBindSlice[i].bindCount = C.uint32_t(bindCount)
			imageOpaqueBindSlice[i].pBinds = nil

			if bindCount > 0 {
				imageOpaqueBindSlice[i].pBinds, err = common.AllocSlice[C.VkSparseMemoryBind, SparseMemoryBind](allocator, b.ImageOpaqueBinds[i].Binds)
				if err != nil {
					return nil, err
				}
			}
		}

		createInfo.pImageOpaqueBinds = imageOpaqueBindPtr
	}

	if imageBindCount > 0 {
		imageBindPtr := (*C.VkSparseImageMemoryBindInfo)(allocator.Malloc(imageBindCount * C.sizeof_struct_VkSparseImageMemoryBindInfo))
		imageBindSlice := ([]C.VkSparseImageMemoryBindInfo)(unsafe.Slice(imageBindPtr, imageBindCount))

		for i := 0; i < imageBindCount; i++ {
			imageBindSlice[i].image = C.VkImage(unsafe.Pointer(b.ImageBinds[i].Image.Handle()))
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
					bindSlice[j].memory = C.VkDeviceMemory(unsafe.Pointer(b.ImageBinds[i].Binds[j].Memory.Handle()))
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
			signalSemaphoreSlice[i] = C.VkSemaphore(unsafe.Pointer(b.SignalSemaphores[i].Handle()))
		}

		createInfo.pSignalSemaphores = signalSemaphorePtr
	}

	return preallocatedPointer, nil
}
