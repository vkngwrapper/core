package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanLoader1_0 struct {
	driver driver.Driver
}

func CreateStaticLinkedLoader() (*VulkanLoader1_0, error) {
	return CreateLoaderFromProcAddr(unsafe.Pointer(C.vkGetInstanceProcAddr))
}

func CreateLoaderFromProcAddr(addr unsafe.Pointer) (*VulkanLoader1_0, error) {
	driver, err := driver.CreateDriverFromProcAddr(addr)
	if err != nil {
		return nil, err
	}

	return CreateLoaderFromDriver(driver)
}

func CreateLoaderFromDriver(driver driver.Driver) (*VulkanLoader1_0, error) {
	return &VulkanLoader1_0{driver: driver}, nil
}

func (l *VulkanLoader1_0) Version() common.APIVersion {
	return l.driver.Version()
}

func (l *VulkanLoader1_0) Driver() driver.Driver {
	return l.driver
}

func (l *VulkanLoader1_0) attemptAvailableExtensions(layerName *driver.Char) (map[string]*common.ExtensionProperties, common.VkResult, error) {
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

func (l *VulkanLoader1_0) AvailableExtensions() (map[string]*common.ExtensionProperties, common.VkResult, error) {
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

func (l *VulkanLoader1_0) AvailableExtensionsForLayer(layerName string) (map[string]*common.ExtensionProperties, common.VkResult, error) {
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

func (l *VulkanLoader1_0) attemptAvailableLayers() (map[string]*common.LayerProperties, common.VkResult, error) {
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

func (l *VulkanLoader1_0) AvailableLayers() (map[string]*common.LayerProperties, common.VkResult, error) {
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

func (l *VulkanLoader1_0) CreateAllocationCallbacks(o *AllocationCallbackOptions) *AllocationCallbacks {
	callbacks := &AllocationCallbacks{
		allocation:         o.Allocation,
		reallocation:       o.Reallocation,
		free:               o.Free,
		internalAllocation: o.InternalAllocation,
		internalFree:       o.InternalFree,
		userData:           o.UserData,
	}
	callbacks.initHandle()

	return callbacks
}

func (l *VulkanLoader1_0) CreateBufferView(device Device, allocationCallbacks *AllocationCallbacks, options *BufferViewOptions) (BufferView, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var bufferViewHandle driver.VkBufferView

	res, err := device.Driver().VkCreateBufferView(device.Handle(), (*driver.VkBufferViewCreateInfo)(createInfo), allocationCallbacks.Handle(), &bufferViewHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanBufferView{
		driver: device.Driver(),
		device: device.Handle(),
		handle: bufferViewHandle,
	}, res, nil
}

func (l *VulkanLoader1_0) CreateInstance(allocationCallbacks *AllocationCallbacks, options *InstanceOptions) (Instance, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var instanceHandle driver.VkInstance

	res, err := l.driver.VkCreateInstance((*driver.VkInstanceCreateInfo)(createInfo), allocationCallbacks.Handle(), &instanceHandle)
	if err != nil {
		return nil, res, err
	}

	instanceDriver, err := l.driver.CreateInstanceDriver((driver.VkInstance)(instanceHandle))
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	return &vulkanInstance{
		driver: instanceDriver,
		handle: instanceHandle,
	}, res, nil
}

func (l *VulkanLoader1_0) CreateDevice(physicalDevice PhysicalDevice, allocationCallbacks *AllocationCallbacks, options *DeviceOptions) (Device, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var deviceHandle driver.VkDevice
	res, err := physicalDevice.Driver().VkCreateDevice(physicalDevice.Handle(), (*driver.VkDeviceCreateInfo)(createInfo), allocationCallbacks.Handle(), &deviceHandle)
	if err != nil {
		return nil, res, err
	}

	deviceDriver, err := physicalDevice.Driver().CreateDeviceDriver(deviceHandle)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	return &vulkanDevice{driver: deviceDriver, handle: deviceHandle}, res, nil
}

func (l *VulkanLoader1_0) CreateShaderModule(device Device, allocationCallbacks *AllocationCallbacks, o *ShaderModuleOptions) (ShaderModule, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var shaderModule driver.VkShaderModule
	res, err := device.Driver().VkCreateShaderModule(device.Handle(), (*driver.VkShaderModuleCreateInfo)(createInfo), allocationCallbacks.Handle(), &shaderModule)
	if err != nil {
		return nil, res, err
	}

	return &vulkanShaderModule{driver: device.Driver(), handle: shaderModule, device: device.Handle()}, res, nil
}

func (l *VulkanLoader1_0) CreateImageView(device Device, allocationCallbacks *AllocationCallbacks, o *ImageViewOptions) (ImageView, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var imageViewHandle driver.VkImageView

	res, err := device.Driver().VkCreateImageView(device.Handle(), (*driver.VkImageViewCreateInfo)(createInfo), allocationCallbacks.Handle(), &imageViewHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanImageView{driver: device.Driver(), handle: imageViewHandle, device: device.Handle()}, res, nil
}

func (l *VulkanLoader1_0) CreateSemaphore(device Device, allocationCallbacks *AllocationCallbacks, o *SemaphoreOptions) (Semaphore, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var semaphoreHandle driver.VkSemaphore

	res, err := device.Driver().VkCreateSemaphore(device.Handle(), (*driver.VkSemaphoreCreateInfo)(createInfo), allocationCallbacks.Handle(), &semaphoreHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanSemaphore{driver: device.Driver(), device: device.Handle(), handle: semaphoreHandle}, res, nil
}

func (l *VulkanLoader1_0) CreateFence(device Device, allocationCallbacks *AllocationCallbacks, o *FenceOptions) (Fence, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var fenceHandle driver.VkFence

	res, err := device.Driver().VkCreateFence(device.Handle(), (*driver.VkFenceCreateInfo)(createInfo), allocationCallbacks.Handle(), &fenceHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanFence{driver: device.Driver(), device: device.Handle(), handle: fenceHandle}, res, nil
}

func (l *VulkanLoader1_0) CreateBuffer(device Device, allocationCallbacks *AllocationCallbacks, o *BufferOptions) (Buffer, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var buffer driver.VkBuffer

	res, err := device.Driver().VkCreateBuffer(device.Handle(), (*driver.VkBufferCreateInfo)(createInfo), allocationCallbacks.Handle(), &buffer)
	if err != nil {
		return nil, res, err
	}

	return &vulkanBuffer{driver: device.Driver(), handle: buffer, device: device.Handle()}, res, nil
}

func (l *VulkanLoader1_0) CreateDescriptorSetLayout(device Device, allocationCallbacks *AllocationCallbacks, o *DescriptorSetLayoutOptions) (DescriptorSetLayout, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var descriptorSetLayout driver.VkDescriptorSetLayout

	res, err := device.Driver().VkCreateDescriptorSetLayout(device.Handle(), (*driver.VkDescriptorSetLayoutCreateInfo)(createInfo), allocationCallbacks.Handle(), &descriptorSetLayout)
	if err != nil {
		return nil, res, err
	}

	return &vulkanDescriptorSetLayout{
		driver: device.Driver(),
		device: device.Handle(),
		handle: descriptorSetLayout,
	}, res, nil
}

func (l *VulkanLoader1_0) CreateDescriptorPool(device Device, allocationCallbacks *AllocationCallbacks, o *DescriptorPoolOptions) (DescriptorPool, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var descriptorPool driver.VkDescriptorPool

	res, err := device.Driver().VkCreateDescriptorPool(device.Handle(), (*driver.VkDescriptorPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &descriptorPool)
	if err != nil {
		return nil, res, err
	}

	return &vulkanDescriptorPool{
		driver: device.Driver(),
		handle: descriptorPool,
		device: device.Handle(),
	}, res, nil
}

func (l *VulkanLoader1_0) CreateCommandPool(device Device, allocationCallbacks *AllocationCallbacks, o *CommandPoolOptions) (CommandPool, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var cmdPoolHandle driver.VkCommandPool
	res, err := device.Driver().VkCreateCommandPool(device.Handle(), (*driver.VkCommandPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &cmdPoolHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanCommandPool{driver: device.Driver(), handle: cmdPoolHandle, device: device.Handle()}, res, nil
}

func (l *VulkanLoader1_0) CreateEvent(device Device, allocationCallbacks *AllocationCallbacks, o *EventOptions) (Event, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var eventHandle driver.VkEvent
	res, err := device.Driver().VkCreateEvent(device.Handle(), (*driver.VkEventCreateInfo)(createInfo), allocationCallbacks.Handle(), &eventHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanEvent{driver: device.Driver(), handle: eventHandle, device: device.Handle()}, res, nil
}

func (l *VulkanLoader1_0) CreateFrameBuffer(device Device, allocationCallbacks *AllocationCallbacks, o *FramebufferOptions) (Framebuffer, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var framebuffer driver.VkFramebuffer

	res, err := device.Driver().VkCreateFramebuffer(device.Handle(), (*driver.VkFramebufferCreateInfo)(createInfo), allocationCallbacks.Handle(), &framebuffer)
	if err != nil {
		return nil, res, err
	}

	return &vulkanFramebuffer{driver: device.Driver(), device: device.Handle(), handle: framebuffer}, res, nil
}

func (l *VulkanLoader1_0) CreateGraphicsPipelines(device Device, pipelineCache PipelineCache, allocationCallbacks *AllocationCallbacks, o []*GraphicsPipelineOptions) ([]Pipeline, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	pipelineCount := len(o)

	pipelineCreateInfosPtrUnsafe := arena.Malloc(pipelineCount * C.sizeof_struct_VkGraphicsPipelineCreateInfo)
	pipelineCreateInfosSlice := ([]C.VkGraphicsPipelineCreateInfo)(unsafe.Slice((*C.VkGraphicsPipelineCreateInfo)(pipelineCreateInfosPtrUnsafe), pipelineCount))
	for i := 0; i < pipelineCount; i++ {
		next, err := common.AllocNext(arena, o[i])
		if err != nil {
			return nil, common.VKErrorUnknown, err
		}

		err = o[i].populate(arena, &pipelineCreateInfosSlice[i], next)
		if err != nil {
			return nil, common.VKErrorUnknown, err
		}
	}

	pipelinePtr := (*driver.VkPipeline)(arena.Malloc(pipelineCount * int(unsafe.Sizeof([1]driver.VkPipeline{}))))

	var pipelineCacheHandle driver.VkPipelineCache
	if pipelineCache != nil {
		pipelineCacheHandle = pipelineCache.Handle()
	}

	res, err := device.Driver().VkCreateGraphicsPipelines(device.Handle(), pipelineCacheHandle, driver.Uint32(pipelineCount), (*driver.VkGraphicsPipelineCreateInfo)(pipelineCreateInfosPtrUnsafe), allocationCallbacks.Handle(), pipelinePtr)
	if err != nil {
		return nil, res, err
	}

	var output []Pipeline
	pipelineSlice := ([]driver.VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))
	for i := 0; i < pipelineCount; i++ {
		output = append(output, &vulkanPipeline{driver: device.Driver(), device: device.Handle(), handle: pipelineSlice[i]})
	}

	return output, res, nil
}

func (l *VulkanLoader1_0) CreateComputePipelines(device Device, pipelineCache PipelineCache, allocationCallbacks *AllocationCallbacks, o []*ComputePipelineOptions) ([]Pipeline, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	pipelineCount := len(o)

	pipelineCreateInfosPtrUnsafe := arena.Malloc(pipelineCount * C.sizeof_struct_VkComputePipelineCreateInfo)
	pipelineCreateInfosSlice := ([]C.VkComputePipelineCreateInfo)(unsafe.Slice((*C.VkComputePipelineCreateInfo)(pipelineCreateInfosPtrUnsafe), pipelineCount))
	for i := 0; i < pipelineCount; i++ {
		next, err := common.AllocNext(arena, o[i])
		if err != nil {
			return nil, common.VKErrorUnknown, err
		}

		err = o[i].populate(arena, &pipelineCreateInfosSlice[i], next)
		if err != nil {
			return nil, common.VKErrorUnknown, err
		}
	}

	pipelinePtr := (*driver.VkPipeline)(arena.Malloc(pipelineCount * int(unsafe.Sizeof([1]driver.VkPipeline{}))))

	var pipelineCacheHandle driver.VkPipelineCache
	if pipelineCache != nil {
		pipelineCacheHandle = pipelineCache.Handle()
	}

	res, err := device.Driver().VkCreateComputePipelines(device.Handle(), pipelineCacheHandle, driver.Uint32(pipelineCount), (*driver.VkComputePipelineCreateInfo)(pipelineCreateInfosPtrUnsafe), allocationCallbacks.Handle(), pipelinePtr)
	if err != nil {
		return nil, res, err
	}

	var output []Pipeline
	pipelineSlice := ([]driver.VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))
	for i := 0; i < pipelineCount; i++ {
		output = append(output, &vulkanPipeline{driver: device.Driver(), device: device.Handle(), handle: pipelineSlice[i]})
	}

	return output, res, nil
}

func (l *VulkanLoader1_0) CreateImage(device Device, allocationCallbacks *AllocationCallbacks, o *ImageOptions) (Image, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var image driver.VkImage
	res, err := device.Driver().VkCreateImage(device.Handle(), (*driver.VkImageCreateInfo)(createInfo), allocationCallbacks.Handle(), &image)
	if err != nil {
		return nil, res, err
	}

	return &vulkanImage{device: device.Handle(), handle: image, driver: device.Driver()}, res, nil
}

func (l *VulkanLoader1_0) CreatePipelineCache(device Device, allocationCallbacks *AllocationCallbacks, o *PipelineCacheOptions) (PipelineCache, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var pipelineCache driver.VkPipelineCache
	res, err := device.Driver().VkCreatePipelineCache(device.Handle(), (*driver.VkPipelineCacheCreateInfo)(createInfo), allocationCallbacks.Handle(), &pipelineCache)
	if err != nil {
		return nil, res, err
	}

	return &vulkanPipelineCache{driver: device.Driver(), handle: pipelineCache, device: device.Handle()}, res, nil
}

func (l *VulkanLoader1_0) CreatePipelineLayout(device Device, allocationCallbacks *AllocationCallbacks, o *PipelineLayoutOptions) (PipelineLayout, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var pipelineLayout driver.VkPipelineLayout
	res, err := device.Driver().VkCreatePipelineLayout(device.Handle(), (*driver.VkPipelineLayoutCreateInfo)(createInfo), allocationCallbacks.Handle(), &pipelineLayout)
	if err != nil {
		return nil, res, err
	}

	return &vulkanPipelineLayout{driver: device.Driver(), handle: pipelineLayout, device: device.Handle()}, res, nil
}

func (l *VulkanLoader1_0) CreateQueryPool(device Device, allocationCallbacks *AllocationCallbacks, o *QueryPoolOptions) (QueryPool, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var queryPool driver.VkQueryPool

	res, err := device.Driver().VkCreateQueryPool(device.Handle(), (*driver.VkQueryPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &queryPool)
	if err != nil {
		return nil, res, err
	}

	return &vulkanQueryPool{driver: device.Driver(), device: device.Handle(), handle: queryPool}, res, nil

}

func (l *VulkanLoader1_0) CreateRenderPass(device Device, allocationCallbacks *AllocationCallbacks, o *RenderPassOptions) (RenderPass, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var renderPass driver.VkRenderPass

	res, err := device.Driver().VkCreateRenderPass(device.Handle(), (*driver.VkRenderPassCreateInfo)(createInfo), allocationCallbacks.Handle(), &renderPass)
	if err != nil {
		return nil, res, err
	}

	return &vulkanRenderPass{driver: device.Driver(), device: device.Handle(), handle: renderPass}, res, nil
}

func (l *VulkanLoader1_0) CreateSampler(device Device, allocationCallbacks *AllocationCallbacks, o *SamplerOptions) (Sampler, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var sampler driver.VkSampler

	res, err := device.Driver().VkCreateSampler(device.Handle(), (*driver.VkSamplerCreateInfo)(createInfo), allocationCallbacks.Handle(), &sampler)
	if err != nil {
		return nil, res, err
	}

	return &vulkanSampler{handle: sampler, driver: device.Driver(), device: device.Handle()}, common.VKSuccess, nil
}
