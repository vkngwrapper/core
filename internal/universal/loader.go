package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0/options"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/iface"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type VulkanLoader[Buffer iface.Buffer, BufferView iface.BufferView, CommandPool iface.CommandPool,
	DescriptorPool iface.DescriptorPool, DescriptorSet iface.DescriptorSet, Device iface.Device, Event iface.Event,
	Fence iface.Fence, Framebuffer iface.Framebuffer, Instance iface.Instance, Image iface.Image,
	ImageView iface.ImageView, PipelineCache iface.PipelineCache, PipelineLayout iface.PipelineLayout,
	QueryPool iface.QueryPool, RenderPass iface.RenderPass, Sampler iface.Sampler, Semaphore iface.Semaphore,
	ShaderModule iface.ShaderModule, CommandBuffer iface.CommandBuffer, DeviceMemory iface.DeviceMemory,
	DescriptorSetLayout iface.DescriptorSetLayout, Pipeline iface.Pipeline, PhysicalDevice iface.PhysicalDevice,
	Queue iface.Queue] struct {
	driver driver.Driver

	zeroBuffer              Buffer
	zeroBufferView          BufferView
	zeroCommandPool         CommandPool
	zeroDescriptorPool      DescriptorPool
	zeroDescriptorSet       DescriptorSet
	zeroDevice              Device
	zeroEvent               Event
	zeroFence               Fence
	zeroFramebuffer         Framebuffer
	zeroInstance            Instance
	zeroImage               Image
	zeroImageView           ImageView
	zeroPipelineCache       PipelineCache
	zeroPipelineLayout      PipelineLayout
	zeroQueryPool           QueryPool
	zeroRenderPass          RenderPass
	zeroSampler             Sampler
	zeroSemaphore           Semaphore
	zeroShaderModule        ShaderModule
	zeroCommandBuffer       CommandBuffer
	zeroDeviceMemory        DeviceMemory
	zeroDescriptorSetLayout DescriptorSetLayout
	zeroPipeline            Pipeline
	zeroPhysicalDevice      PhysicalDevice
}

func NewLoader[Buffer iface.Buffer, BufferView iface.BufferView, CommandPool iface.CommandPool,
	DescriptorPool iface.DescriptorPool, DescriptorSet iface.DescriptorSet, Device iface.Device, Event iface.Event,
	Fence iface.Fence, Framebuffer iface.Framebuffer, Instance iface.Instance, Image iface.Image,
	ImageView iface.ImageView, PipelineCache iface.PipelineCache, PipelineLayout iface.PipelineLayout,
	QueryPool iface.QueryPool, RenderPass iface.RenderPass, Sampler iface.Sampler, Semaphore iface.Semaphore,
	ShaderModule iface.ShaderModule, CommandBuffer iface.CommandBuffer, DeviceMemory iface.DeviceMemory,
	DescriptorSetLayout iface.DescriptorSetLayout, Pipeline iface.Pipeline, PhysicalDevice iface.PhysicalDevice,
	Queue iface.Queue](
	driver driver.Driver,
) *VulkanLoader[Buffer,
	BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence, Framebuffer, Instance, Image,
	ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler, Semaphore, ShaderModule,
	CommandBuffer, DeviceMemory, DescriptorSetLayout, Pipeline, PhysicalDevice, Queue] {

	return &VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event,
		Fence, Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass,
		Sampler, Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout, Pipeline, PhysicalDevice,
		Queue]{
		driver: driver,
	}
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) Version() common.APIVersion {
	return l.driver.Version()
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout, Pipeline, PhysicalDevice, Queue]) Driver() driver.Driver {
	return l.driver
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) attemptAvailableExtensions(layerName *driver.Char) (map[string]*common.ExtensionProperties, common.VkResult, error) {
	alloc := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(alloc)

	extensionCount := (*driver.Uint32)(alloc.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := l.driver.VkEnumerateInstanceExtensionProperties(layerName, extensionCount, nil)
	if err != nil || *extensionCount == 0 {
		return nil, res, err
	}

	extensionsUnsafe := alloc.Malloc(int(*extensionCount) * C.sizeof_struct_VkExtensionProperties)

	res, err = l.driver.VkEnumerateInstanceExtensionProperties(layerName, extensionCount, (*driver.VkExtensionProperties)(extensionsUnsafe))
	if err != nil {
		return nil, res, err
	}

	intExtensionCount := int(*extensionCount)
	extensionArray := ([]C.VkExtensionProperties)(unsafe.Slice((*C.VkExtensionProperties)(extensionsUnsafe), intExtensionCount))
	outExtensions := make(map[string]*common.ExtensionProperties)
	for i := 0; i < intExtensionCount; i++ {
		extension := extensionArray[i]

		outExtension := &common.ExtensionProperties{
			ExtensionName: C.GoString((*C.char)(&extension.extensionName[0])),
			SpecVersion:   common.Version(extension.specVersion),
		}

		existingExtension, ok := outExtensions[outExtension.ExtensionName]
		if ok && existingExtension.SpecVersion >= outExtension.SpecVersion {
			continue
		}
		outExtensions[outExtension.ExtensionName] = outExtension
	}

	return outExtensions, res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) AvailableExtensions() (map[string]*common.ExtensionProperties, common.VkResult, error) {
	// There may be a race condition that adds new available extensions between getting the
	// extension count & pulling the extensions, in which case, attemptAvailableExtensions will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*common.ExtensionProperties
	var result common.VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == common.VKIncomplete) {
		layers, result, err = l.attemptAvailableExtensions(nil)
	}
	return layers, result, err
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) AvailableExtensionsForLayer(layerName string) (map[string]*common.ExtensionProperties, common.VkResult, error) {
	// There may be a race condition that adds new available extensions between getting the
	// extension count & pulling the extensions, in which case, attemptAvailableExtensions will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*common.ExtensionProperties
	var result common.VkResult
	var err error
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	layerNamePtr := (*driver.Char)(allocator.CString(layerName))
	for doWhile := true; doWhile; doWhile = (result == common.VKIncomplete) {
		layers, result, err = l.attemptAvailableExtensions(layerNamePtr)
	}
	return layers, result, err
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) attemptAvailableLayers() (map[string]*common.LayerProperties, common.VkResult, error) {
	alloc := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(alloc)

	layerCount := (*driver.Uint32)(alloc.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := l.driver.VkEnumerateInstanceLayerProperties(layerCount, nil)
	if err != nil || *layerCount == 0 {
		return nil, res, err
	}

	layersUnsafe := alloc.Malloc(int(*layerCount) * C.sizeof_struct_VkLayerProperties)

	res, err = l.driver.VkEnumerateInstanceLayerProperties(layerCount, (*driver.VkLayerProperties)(layersUnsafe))
	if err != nil {
		return nil, res, err
	}

	intLayerCount := int(*layerCount)
	layerArray := ([]C.VkLayerProperties)(unsafe.Slice((*C.VkLayerProperties)(layersUnsafe), intLayerCount))
	outLayers := make(map[string]*common.LayerProperties)
	for i := 0; i < intLayerCount; i++ {
		layer := layerArray[i]

		outLayer := &common.LayerProperties{
			LayerName:             C.GoString((*C.char)(&layer.layerName[0])),
			SpecVersion:           common.Version(layer.specVersion),
			ImplementationVersion: common.Version(layer.implementationVersion),
			Description:           C.GoString((*C.char)(&layer.description[0])),
		}

		existingLayer, ok := outLayers[outLayer.LayerName]
		if ok && existingLayer.SpecVersion >= outLayer.SpecVersion {
			continue
		}
		outLayers[outLayer.LayerName] = outLayer
	}

	return outLayers, res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) AvailableLayers() (map[string]*common.LayerProperties, common.VkResult, error) {
	// There may be a race condition that adds new available layers between getting the
	// layer count & pulling the layers, in which case, attemptAvailableLayers will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*common.LayerProperties
	var result common.VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == common.VKIncomplete) {
		layers, result, err = l.attemptAvailableLayers()
	}
	return layers, result, err
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateBufferView(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, options *options.BufferViewOptions) (BufferView, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, options)
	if err != nil {
		return l.zeroBufferView, common.VKErrorUnknown, err
	}

	var bufferViewHandle driver.VkBufferView

	res, err := device.Driver().VkCreateBufferView(device.Handle(), (*driver.VkBufferViewCreateInfo)(createInfo), allocationCallbacks.Handle(), &bufferViewHandle)
	if err != nil {
		return l.zeroBufferView, res, err
	}

	buffer := iface.BufferView(&VulkanBufferView{
		driver: device.Driver(),
		device: device.Handle(),
		handle: bufferViewHandle,
	})
	return buffer.(BufferView), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateInstance(allocationCallbacks *driver.AllocationCallbacks, options *options.InstanceOptions) (Instance, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, options)
	if err != nil {
		return l.zeroInstance, common.VKErrorUnknown, err
	}

	var instanceHandle driver.VkInstance

	res, err := l.driver.VkCreateInstance((*driver.VkInstanceCreateInfo)(createInfo), allocationCallbacks.Handle(), &instanceHandle)
	if err != nil {
		return l.zeroInstance, res, err
	}

	instanceDriver, err := l.driver.CreateInstanceDriver((driver.VkInstance)(instanceHandle))
	if err != nil {
		return l.zeroInstance, common.VKErrorUnknown, err
	}

	instance := iface.Instance(&VulkanInstance{
		driver: instanceDriver,
		handle: instanceHandle,
	})
	return instance.(Instance), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateDevice(physicalDevice iface.PhysicalDevice, allocationCallbacks *driver.AllocationCallbacks, options *options.DeviceOptions) (Device, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, options)
	if err != nil {
		return l.zeroDevice, common.VKErrorUnknown, err
	}

	var deviceHandle driver.VkDevice
	res, err := physicalDevice.Driver().VkCreateDevice(physicalDevice.Handle(), (*driver.VkDeviceCreateInfo)(createInfo), allocationCallbacks.Handle(), &deviceHandle)
	if err != nil {
		return l.zeroDevice, res, err
	}

	deviceDriver, err := physicalDevice.Driver().CreateDeviceDriver(deviceHandle)
	if err != nil {
		return l.zeroDevice, common.VKErrorUnknown, err
	}

	device := iface.Device(&VulkanDevice{driver: deviceDriver, handle: deviceHandle})
	return device.(Device), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateShaderModule(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.ShaderModuleOptions) (ShaderModule, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return l.zeroShaderModule, common.VKErrorUnknown, err
	}

	var shaderModuleHandle driver.VkShaderModule
	res, err := device.Driver().VkCreateShaderModule(device.Handle(), (*driver.VkShaderModuleCreateInfo)(createInfo), allocationCallbacks.Handle(), &shaderModuleHandle)
	if err != nil {
		return l.zeroShaderModule, res, err
	}

	shaderModule := iface.ShaderModule(&VulkanShaderModule{driver: device.Driver(), handle: shaderModuleHandle, device: device.Handle()})
	return shaderModule.(ShaderModule), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateImageView(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.ImageViewOptions) (ImageView, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return l.zeroImageView, common.VKErrorUnknown, err
	}

	var imageViewHandle driver.VkImageView

	res, err := device.Driver().VkCreateImageView(device.Handle(), (*driver.VkImageViewCreateInfo)(createInfo), allocationCallbacks.Handle(), &imageViewHandle)
	if err != nil {
		return l.zeroImageView, res, err
	}

	imageView := iface.ImageView(&VulkanImageView{driver: device.Driver(), handle: imageViewHandle, device: device.Handle()})
	return imageView.(ImageView), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateSemaphore(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.SemaphoreOptions) (Semaphore, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return l.zeroSemaphore, common.VKErrorUnknown, err
	}

	var semaphoreHandle driver.VkSemaphore

	res, err := device.Driver().VkCreateSemaphore(device.Handle(), (*driver.VkSemaphoreCreateInfo)(createInfo), allocationCallbacks.Handle(), &semaphoreHandle)
	if err != nil {
		return l.zeroSemaphore, res, err
	}

	semaphore := iface.Semaphore(&VulkanSemaphore{driver: device.Driver(), device: device.Handle(), handle: semaphoreHandle})
	return semaphore.(Semaphore), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateFence(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.FenceOptions) (Fence, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return l.zeroFence, common.VKErrorUnknown, err
	}

	var fenceHandle driver.VkFence

	res, err := device.Driver().VkCreateFence(device.Handle(), (*driver.VkFenceCreateInfo)(createInfo), allocationCallbacks.Handle(), &fenceHandle)
	if err != nil {
		return l.zeroFence, res, err
	}

	fence := iface.Fence(&VulkanFence{driver: device.Driver(), device: device.Handle(), handle: fenceHandle})
	return fence.(Fence), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateBuffer(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.BufferOptions) (Buffer, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return l.zeroBuffer, common.VKErrorUnknown, err
	}

	var bufferHandle driver.VkBuffer

	res, err := device.Driver().VkCreateBuffer(device.Handle(), (*driver.VkBufferCreateInfo)(createInfo), allocationCallbacks.Handle(), &bufferHandle)
	if err != nil {
		return l.zeroBuffer, res, err
	}

	buffer := iface.Buffer(&VulkanBuffer{driver: device.Driver(), handle: bufferHandle, device: device.Handle()})
	return buffer.(Buffer), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateDescriptorSetLayout(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.DescriptorSetLayoutOptions) (DescriptorSetLayout, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return l.zeroDescriptorSetLayout, common.VKErrorUnknown, err
	}

	var descriptorSetLayoutHandle driver.VkDescriptorSetLayout

	res, err := device.Driver().VkCreateDescriptorSetLayout(device.Handle(), (*driver.VkDescriptorSetLayoutCreateInfo)(createInfo), allocationCallbacks.Handle(), &descriptorSetLayoutHandle)
	if err != nil {
		return l.zeroDescriptorSetLayout, res, err
	}

	descriptorSetLayout := iface.DescriptorSetLayout(&VulkanDescriptorSetLayout{
		driver: device.Driver(),
		device: device.Handle(),
		handle: descriptorSetLayoutHandle,
	})
	return descriptorSetLayout.(DescriptorSetLayout), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateDescriptorPool(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.DescriptorPoolOptions) (DescriptorPool, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return l.zeroDescriptorPool, common.VKErrorUnknown, err
	}

	var descriptorPoolHandle driver.VkDescriptorPool

	res, err := device.Driver().VkCreateDescriptorPool(device.Handle(), (*driver.VkDescriptorPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &descriptorPoolHandle)
	if err != nil {
		return l.zeroDescriptorPool, res, err
	}

	descriptorPool := iface.DescriptorPool(&VulkanDescriptorPool{
		driver: device.Driver(),
		handle: descriptorPoolHandle,
		device: device.Handle(),
	})
	return descriptorPool.(DescriptorPool), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateCommandPool(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.CommandPoolOptions) (CommandPool, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return l.zeroCommandPool, common.VKErrorUnknown, err
	}

	var cmdPoolHandle driver.VkCommandPool
	res, err := device.Driver().VkCreateCommandPool(device.Handle(), (*driver.VkCommandPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &cmdPoolHandle)
	if err != nil {
		return l.zeroCommandPool, res, err
	}

	cmdPool := iface.CommandPool(&VulkanCommandPool{driver: device.Driver(), handle: cmdPoolHandle, device: device.Handle()})
	return cmdPool.(CommandPool), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateEvent(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.EventOptions) (Event, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return l.zeroEvent, common.VKErrorUnknown, err
	}

	var eventHandle driver.VkEvent
	res, err := device.Driver().VkCreateEvent(device.Handle(), (*driver.VkEventCreateInfo)(createInfo), allocationCallbacks.Handle(), &eventHandle)
	if err != nil {
		return l.zeroEvent, res, err
	}

	event := iface.Event(&VulkanEvent{driver: device.Driver(), handle: eventHandle, device: device.Handle()})
	return event.(Event), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateFrameBuffer(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.FramebufferOptions) (Framebuffer, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return l.zeroFramebuffer, common.VKErrorUnknown, err
	}

	var framebufferHandle driver.VkFramebuffer

	res, err := device.Driver().VkCreateFramebuffer(device.Handle(), (*driver.VkFramebufferCreateInfo)(createInfo), allocationCallbacks.Handle(), &framebufferHandle)
	if err != nil {
		return l.zeroFramebuffer, res, err
	}

	framebuffer := iface.Framebuffer(&VulkanFramebuffer{driver: device.Driver(), device: device.Handle(), handle: framebufferHandle})
	return framebuffer.(Framebuffer), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateGraphicsPipelines(device iface.Device, pipelineCache iface.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []options.GraphicsPipelineOptions) ([]Pipeline, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	pipelineCount := len(o)

	pipelineCreateInfosPtr, err := core.AllocOptionSlice[C.VkGraphicsPipelineCreateInfo, options.GraphicsPipelineOptions](arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	pipelinePtr := (*driver.VkPipeline)(arena.Malloc(pipelineCount * int(unsafe.Sizeof([1]driver.VkPipeline{}))))

	var pipelineCacheHandle driver.VkPipelineCache
	if pipelineCache != nil {
		pipelineCacheHandle = pipelineCache.Handle()
	}

	res, err := device.Driver().VkCreateGraphicsPipelines(device.Handle(), pipelineCacheHandle, driver.Uint32(pipelineCount), (*driver.VkGraphicsPipelineCreateInfo)(unsafe.Pointer(pipelineCreateInfosPtr)), allocationCallbacks.Handle(), pipelinePtr)
	if err != nil {
		return nil, res, err
	}

	var output []Pipeline
	pipelineSlice := ([]driver.VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))
	for i := 0; i < pipelineCount; i++ {
		pipeline := iface.Pipeline(&VulkanPipeline{driver: device.Driver(), device: device.Handle(), handle: pipelineSlice[i]})
		output = append(output, pipeline.(Pipeline))
	}

	return output, res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateComputePipelines(device iface.Device, pipelineCache iface.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []options.ComputePipelineOptions) ([]Pipeline, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	pipelineCount := len(o)

	pipelineCreateInfosPtr, err := core.AllocOptionSlice[C.VkComputePipelineCreateInfo, options.ComputePipelineOptions](arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	pipelinePtr := (*driver.VkPipeline)(arena.Malloc(pipelineCount * int(unsafe.Sizeof([1]driver.VkPipeline{}))))

	var pipelineCacheHandle driver.VkPipelineCache
	if pipelineCache != nil {
		pipelineCacheHandle = pipelineCache.Handle()
	}

	res, err := device.Driver().VkCreateComputePipelines(device.Handle(), pipelineCacheHandle, driver.Uint32(pipelineCount), (*driver.VkComputePipelineCreateInfo)(unsafe.Pointer(pipelineCreateInfosPtr)), allocationCallbacks.Handle(), pipelinePtr)
	if err != nil {
		return nil, res, err
	}

	var output []Pipeline
	pipelineSlice := ([]driver.VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))
	for i := 0; i < pipelineCount; i++ {
		pipeline := iface.Pipeline(&VulkanPipeline{driver: device.Driver(), device: device.Handle(), handle: pipelineSlice[i]})
		output = append(output, pipeline.(Pipeline))
	}

	return output, res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateImage(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.ImageOptions) (Image, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return l.zeroImage, common.VKErrorUnknown, err
	}

	var imageHandle driver.VkImage
	res, err := device.Driver().VkCreateImage(device.Handle(), (*driver.VkImageCreateInfo)(createInfo), allocationCallbacks.Handle(), &imageHandle)
	if err != nil {
		return l.zeroImage, res, err
	}

	image := iface.Image(&VulkanImage{device: device.Handle(), handle: imageHandle, driver: device.Driver()})
	return image.(Image), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreatePipelineCache(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.PipelineCacheOptions) (PipelineCache, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return l.zeroPipelineCache, common.VKErrorUnknown, err
	}

	var pipelineCacheHandle driver.VkPipelineCache
	res, err := device.Driver().VkCreatePipelineCache(device.Handle(), (*driver.VkPipelineCacheCreateInfo)(createInfo), allocationCallbacks.Handle(), &pipelineCacheHandle)
	if err != nil {
		return l.zeroPipelineCache, res, err
	}

	pipelineCache := iface.PipelineCache(&VulkanPipelineCache{driver: device.Driver(), handle: pipelineCacheHandle, device: device.Handle()})
	return pipelineCache.(PipelineCache), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreatePipelineLayout(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.PipelineLayoutOptions) (PipelineLayout, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return l.zeroPipelineLayout, common.VKErrorUnknown, err
	}

	var pipelineLayoutHandle driver.VkPipelineLayout
	res, err := device.Driver().VkCreatePipelineLayout(device.Handle(), (*driver.VkPipelineLayoutCreateInfo)(createInfo), allocationCallbacks.Handle(), &pipelineLayoutHandle)
	if err != nil {
		return l.zeroPipelineLayout, res, err
	}

	pipelineLayout := iface.PipelineLayout(&VulkanPipelineLayout{driver: device.Driver(), handle: pipelineLayoutHandle, device: device.Handle()})
	return pipelineLayout.(PipelineLayout), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateQueryPool(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.QueryPoolOptions) (QueryPool, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return l.zeroQueryPool, common.VKErrorUnknown, err
	}

	var queryPoolHandle driver.VkQueryPool

	res, err := device.Driver().VkCreateQueryPool(device.Handle(), (*driver.VkQueryPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &queryPoolHandle)
	if err != nil {
		return l.zeroQueryPool, res, err
	}

	queryPool := iface.QueryPool(&VulkanQueryPool{driver: device.Driver(), device: device.Handle(), handle: queryPoolHandle})
	return queryPool.(QueryPool), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateRenderPass(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.RenderPassOptions) (RenderPass, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return l.zeroRenderPass, common.VKErrorUnknown, err
	}

	var renderPassHandle driver.VkRenderPass

	res, err := device.Driver().VkCreateRenderPass(device.Handle(), (*driver.VkRenderPassCreateInfo)(createInfo), allocationCallbacks.Handle(), &renderPassHandle)
	if err != nil {
		return l.zeroRenderPass, res, err
	}

	renderPass := iface.RenderPass(&VulkanRenderPass{driver: device.Driver(), device: device.Handle(), handle: renderPassHandle})
	return renderPass.(RenderPass), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) CreateSampler(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.SamplerOptions) (Sampler, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return l.zeroSampler, common.VKErrorUnknown, err
	}

	var samplerHandle driver.VkSampler

	res, err := device.Driver().VkCreateSampler(device.Handle(), (*driver.VkSamplerCreateInfo)(createInfo), allocationCallbacks.Handle(), &samplerHandle)
	if err != nil {
		return l.zeroSampler, res, err
	}

	sampler := iface.Sampler(&VulkanSampler{handle: samplerHandle, driver: device.Driver(), device: device.Handle()})
	return sampler.(Sampler), res, nil
}

// Free a slice of command buffers which should all have the same device/driver/pool
// guaranteed to have at least one element
func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) freeCommandBufferSlice(buffers []iface.CommandBuffer) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	bufferCount := len(buffers)
	bufferDriver := buffers[0].Driver()
	bufferDevice := buffers[0].DeviceHandle()
	bufferPool := buffers[0].CommandPoolHandle()

	size := bufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{}))
	bufferArrayPtr := (*driver.VkCommandBuffer)(allocator.Malloc(size))
	bufferArraySlice := ([]driver.VkCommandBuffer)(unsafe.Slice(bufferArrayPtr, bufferCount))

	for i := 0; i < bufferCount; i++ {
		bufferArraySlice[i] = buffers[i].Handle()
	}

	bufferDriver.VkFreeCommandBuffers(bufferDevice, bufferPool, driver.Uint32(bufferCount), bufferArrayPtr)
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) FreeCommandBuffers(buffers []iface.CommandBuffer) {
	bufferCount := len(buffers)
	if bufferCount == 0 {
		return
	}

	multimap := make(map[driver.VkCommandPool][]iface.CommandBuffer)
	for _, buffer := range buffers {
		poolHandle := buffer.CommandPoolHandle()
		existingSet := multimap[poolHandle]
		multimap[poolHandle] = append(existingSet, buffer)
	}

	for _, setBuffers := range multimap {
		l.freeCommandBufferSlice(setBuffers)
	}
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) AllocateCommandBuffers(o *options.CommandBufferOptions) ([]CommandBuffer, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	if o.CommandPool == nil {
		return nil, common.VKErrorUnknown, errors.New("no command pool provided to allocate from")
	}

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	device := o.CommandPool.Device()

	commandBufferPtr := (*driver.VkCommandBuffer)(arena.Malloc(o.BufferCount * int(unsafe.Sizeof([1]driver.VkCommandBuffer{}))))

	res, err := o.CommandPool.Driver().VkAllocateCommandBuffers(device, (*driver.VkCommandBufferAllocateInfo)(createInfo), commandBufferPtr)
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	commandBufferArray := ([]driver.VkCommandBuffer)(unsafe.Slice(commandBufferPtr, o.BufferCount))
	var result []CommandBuffer
	for i := 0; i < o.BufferCount; i++ {
		commandBuffer := iface.CommandBuffer(&VulkanCommandBuffer{driver: o.CommandPool.Driver(), pool: o.CommandPool.Handle(), device: device, handle: commandBufferArray[i]})
		result = append(result, commandBuffer.(CommandBuffer))
	}

	return result, res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) AllocateDescriptorSets(o *options.DescriptorSetOptions) ([]DescriptorSet, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	if o.DescriptorPool == nil {
		return nil, common.VKErrorUnknown, errors.New("no descriptor pool provided to allocate from")
	}

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	device := o.DescriptorPool.DeviceHandle()
	poolDriver := o.DescriptorPool.Driver()

	setCount := len(o.AllocationLayouts)
	descriptorSets := (*driver.VkDescriptorSet)(arena.Malloc(setCount * int(unsafe.Sizeof([1]C.VkDescriptorSet{}))))

	res, err := poolDriver.VkAllocateDescriptorSets(device, (*driver.VkDescriptorSetAllocateInfo)(createInfo), descriptorSets)
	if err != nil {
		return nil, res, err
	}

	var sets []DescriptorSet
	descriptorSetSlice := ([]driver.VkDescriptorSet)(unsafe.Slice(descriptorSets, setCount))
	for i := 0; i < setCount; i++ {
		descriptorSet := iface.DescriptorSet(&VulkanDescriptorSet{handle: descriptorSetSlice[i], driver: poolDriver, device: device, pool: o.DescriptorPool.Handle()})
		sets = append(sets, descriptorSet.(DescriptorSet))
	}

	return sets, res, nil
}

// Free a slice of descriptor sets which should all have the same device/driver/pool
// guaranteed to have at least one element
func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) freeDescriptorSetSlice(sets []iface.DescriptorSet) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	setSize := len(sets)
	arraySize := setSize * int(unsafe.Sizeof([1]C.VkDescriptorSet{}))

	arrayPtr := (*driver.VkDescriptorSet)(arena.Malloc(arraySize))
	arraySlice := ([]driver.VkDescriptorSet)(unsafe.Slice(arrayPtr, setSize))

	for i := 0; i < setSize; i++ {
		arraySlice[i] = sets[i].Handle()
	}

	setDriver := sets[0].Driver()
	pool := sets[0].PoolHandle()
	device := sets[0].DeviceHandle()

	return setDriver.VkFreeDescriptorSets(device, pool, driver.Uint32(setSize), arrayPtr)
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) FreeDescriptorSets(sets []iface.DescriptorSet) (common.VkResult, error) {
	poolMultimap := make(map[driver.VkDescriptorPool][]iface.DescriptorSet)

	for _, set := range sets {
		poolHandle := set.PoolHandle()
		existingSet := poolMultimap[poolHandle]
		poolMultimap[poolHandle] = append(existingSet, set)
	}

	var res common.VkResult
	var err error
	for _, set := range poolMultimap {
		res, err = l.freeDescriptorSetSlice(set)
		if err != nil {
			return res, err
		}
	}

	return res, err
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) AllocateMemory(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.DeviceMemoryOptions) (DeviceMemory, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return l.zeroDeviceMemory, common.VKErrorUnknown, err
	}

	var deviceMemoryHandle driver.VkDeviceMemory

	deviceDriver := device.Driver()
	deviceHandle := device.Handle()

	res, err := deviceDriver.VkAllocateMemory(deviceHandle, (*driver.VkMemoryAllocateInfo)(createInfo), allocationCallbacks.Handle(), &deviceMemoryHandle)
	if err != nil {
		return l.zeroDeviceMemory, res, err
	}

	deviceMemory := iface.DeviceMemory(&VulkanDeviceMemory{
		driver: deviceDriver,
		device: deviceHandle,
		handle: deviceMemoryHandle,
		size:   o.AllocationSize,
	})
	return deviceMemory.(DeviceMemory), res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) FreeDeviceMemory(deviceMemory iface.DeviceMemory, allocationCallbacks *driver.AllocationCallbacks) {
	// This is really only here for a kind of API symmetry
	freeDeviceMemory(deviceMemory, allocationCallbacks)
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) PhysicalDevices(instance iface.Instance) ([]PhysicalDevice, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	count := (*driver.Uint32)(allocator.Malloc(int(unsafe.Sizeof(driver.Uint32(0)))))

	res, err := instance.Driver().VkEnumeratePhysicalDevices(instance.Handle(), count, nil)
	if err != nil {
		return nil, res, err
	}

	if *count == 0 {
		return nil, res, nil
	}

	allocatedHandles := allocator.Malloc(int(uintptr(*count) * unsafe.Sizeof([1]driver.VkPhysicalDevice{})))

	deviceHandles := ([]driver.VkPhysicalDevice)(unsafe.Slice((*driver.VkPhysicalDevice)(allocatedHandles), int(*count)))
	res, err = instance.Driver().VkEnumeratePhysicalDevices(instance.Handle(), count, (*driver.VkPhysicalDevice)(allocatedHandles))
	if err != nil {
		return nil, res, err
	}

	goCount := uint32(*count)
	var devices []PhysicalDevice
	for ind := uint32(0); ind < goCount; ind++ {
		physicalDevice := iface.PhysicalDevice(&VulkanPhysicalDevice{driver: instance.Driver(), handle: deviceHandles[ind]})
		devices = append(devices, physicalDevice.(PhysicalDevice))
	}

	return devices, res, nil
}

func (l *VulkanLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence,
	Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler,
	Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout,
	Pipeline, PhysicalDevice, Queue]) DeviceQueue(device iface.Device, queueFamilyIndex int, queueIndex int) Queue {

	var queueHandle driver.VkQueue

	device.Driver().VkGetDeviceQueue(device.Handle(), driver.Uint32(queueFamilyIndex), driver.Uint32(queueIndex), &queueHandle)

	queue := iface.Queue(&VulkanQueue{driver: device.Driver(), handle: queueHandle})
	return queue.(Queue)
}
