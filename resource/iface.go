package resource

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/cgoalloc"
	"time"
	"unsafe"
)

type Buffer interface {
	Handle() loader.VkBuffer
	Destroy() error
	MemoryRequirements(allocator cgoalloc.Allocator) (*core.MemoryRequirements, error)
	BindBufferMemory(memory DeviceMemory, offset int) (loader.VkResult, error)
}

type DeviceMemory interface {
	Handle() loader.VkDeviceMemory
	Free() error
	MapMemory(offset int, size int) (unsafe.Pointer, loader.VkResult, error)
	UnmapMemory() error
	WriteData(offset int, data interface{}) (loader.VkResult, error)
}

type Device interface {
	Handle() loader.VkDevice
	Destroy() error
	Loader() *loader.Loader
	GetQueue(queueFamilyIndex int, queueIndex int) (Queue, error)
	CreateShaderModule(allocator cgoalloc.Allocator, o *ShaderModuleOptions) (ShaderModule, loader.VkResult, error)
	CreateImageView(allocator cgoalloc.Allocator, o *ImageViewOptions) (ImageView, loader.VkResult, error)
	CreateSemaphore(allocator cgoalloc.Allocator, o *SemaphoreOptions) (Semaphore, loader.VkResult, error)
	CreateFence(allocator cgoalloc.Allocator, o *FenceOptions) (Fence, loader.VkResult, error)
	WaitForIdle() (loader.VkResult, error)
	WaitForFences(allocator cgoalloc.Allocator, waitForAll bool, timeout time.Duration, fences []Fence) (loader.VkResult, error)
	ResetFences(allocator cgoalloc.Allocator, fences []Fence) (loader.VkResult, error)
	CreateBuffer(allocator cgoalloc.Allocator, o *BufferOptions) (Buffer, loader.VkResult, error)
	AllocateMemory(allocator cgoalloc.Allocator, o *DeviceMemoryOptions) (DeviceMemory, loader.VkResult, error)
}

type Fence interface {
	Handle() loader.VkFence
	Destroy() error
}

type Image interface {
	Handle() loader.VkImage
}

type ImageView interface {
	Handle() loader.VkImageView
	Destroy() error
}

type Instance interface {
	Handle() loader.VkInstance
	Destroy() error
	Loader() *loader.Loader
	PhysicalDevices(allocator cgoalloc.Allocator) ([]PhysicalDevice, loader.VkResult, error)
}

type PhysicalDevice interface {
	Handle() loader.VkPhysicalDevice
	QueueFamilyProperties(allocator cgoalloc.Allocator) ([]*core.QueueFamily, error)
	CreateDevice(allocator cgoalloc.Allocator, options *DeviceOptions) (Device, loader.VkResult, error)
	Properties(allocator cgoalloc.Allocator) (*core.PhysicalDeviceProperties, error)
	Features(allocator cgoalloc.Allocator) (*core.PhysicalDeviceFeatures, error)
	AvailableExtensions(allocator cgoalloc.Allocator) (map[string]*core.ExtensionProperties, loader.VkResult, error)
	MemoryProperties(allocator cgoalloc.Allocator) *PhysicalDeviceMemoryProperties
}

type Queue interface {
	Handle() loader.VkQueue
	Loader() *loader.Loader
	WaitForIdle() (loader.VkResult, error)
}

type Semaphore interface {
	Handle() loader.VkSemaphore
	Destroy() error
}

type ShaderModule interface {
	Handle() loader.VkShaderModule
	Destroy() error
}
