package resources

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"time"
	"unsafe"
)

//go:generate mockgen -source iface.go -destination ../mocks/resource.go -package=mocks

type Buffer interface {
	Handle() loader.VkBuffer
	Destroy() error
	MemoryRequirements() (*core.MemoryRequirements, error)
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
	Loader() loader.Loader
	GetQueue(queueFamilyIndex int, queueIndex int) (Queue, error)
	CreateShaderModule(o *ShaderModuleOptions) (ShaderModule, loader.VkResult, error)
	CreateImageView(o *ImageViewOptions) (ImageView, loader.VkResult, error)
	CreateSemaphore(o *SemaphoreOptions) (Semaphore, loader.VkResult, error)
	CreateFence(o *FenceOptions) (Fence, loader.VkResult, error)
	WaitForIdle() (loader.VkResult, error)
	WaitForFences(waitForAll bool, timeout time.Duration, fences []Fence) (loader.VkResult, error)
	ResetFences(fences []Fence) (loader.VkResult, error)
	CreateBuffer(o *BufferOptions) (Buffer, loader.VkResult, error)
	AllocateMemory(o *DeviceMemoryOptions) (DeviceMemory, loader.VkResult, error)
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
	Loader() loader.Loader
	PhysicalDevices() ([]PhysicalDevice, loader.VkResult, error)
}

type PhysicalDevice interface {
	Handle() loader.VkPhysicalDevice
	QueueFamilyProperties() ([]*core.QueueFamily, error)
	CreateDevice(options *DeviceOptions) (Device, loader.VkResult, error)
	Properties() (*core.PhysicalDeviceProperties, error)
	Features() (*core.PhysicalDeviceFeatures, error)
	AvailableExtensions() (map[string]*core.ExtensionProperties, loader.VkResult, error)
	MemoryProperties() *PhysicalDeviceMemoryProperties
}

type Queue interface {
	Handle() loader.VkQueue
	Loader() loader.Loader
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
