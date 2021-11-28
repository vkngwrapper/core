package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanLoader1_0 struct {
	driver Driver
}

func CreateStaticLinkedLoader() (*VulkanLoader1_0, error) {
	return CreateLoaderFromProcAddr(unsafe.Pointer(C.vkGetInstanceProcAddr))
}

func CreateLoaderFromProcAddr(addr unsafe.Pointer) (*VulkanLoader1_0, error) {
	driver, err := createDriverFromProcAddr(addr)
	if err != nil {
		return nil, err
	}

	return CreateLoaderFromDriver(driver)
}

func CreateLoaderFromDriver(driver Driver) (*VulkanLoader1_0, error) {
	return &VulkanLoader1_0{driver: driver}, nil
}

func (l *VulkanLoader1_0) Version() common.APIVersion {
	return l.driver.Version()
}

func (l *VulkanLoader1_0) Driver() Driver {
	return l.driver
}

func (l *VulkanLoader1_0) attemptAvailableExtensions(layerName *Char) (map[string]*common.ExtensionProperties, VkResult, error) {
	alloc := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(alloc)

	extensionCount := (*Uint32)(alloc.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := l.driver.VkEnumerateInstanceExtensionProperties(layerName, extensionCount, nil)
	if err != nil || *extensionCount == 0 {
		return nil, res, err
	}

	extensionsUnsafe := alloc.Malloc(int(*extensionCount) * int(unsafe.Sizeof(C.VkExtensionProperties{})))

	res, err = l.driver.VkEnumerateInstanceExtensionProperties(layerName, extensionCount, (*VkExtensionProperties)(extensionsUnsafe))
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

func (l *VulkanLoader1_0) AvailableExtensions() (map[string]*common.ExtensionProperties, VkResult, error) {
	// There may be a race condition that adds new available extensions between getting the
	// extension count & pulling the extensions, in which case, attemptAvailableExtensions will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*common.ExtensionProperties
	var result VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == VKIncomplete) {
		layers, result, err = l.attemptAvailableExtensions(nil)
	}
	return layers, result, err
}

func (l *VulkanLoader1_0) AvailableExtensionsForLayer(layerName string) (map[string]*common.ExtensionProperties, VkResult, error) {
	// There may be a race condition that adds new available extensions between getting the
	// extension count & pulling the extensions, in which case, attemptAvailableExtensions will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*common.ExtensionProperties
	var result VkResult
	var err error
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	layerNamePtr := (*Char)(allocator.CString(layerName))
	for doWhile := true; doWhile; doWhile = (result == VKIncomplete) {
		layers, result, err = l.attemptAvailableExtensions(layerNamePtr)
	}
	return layers, result, err
}

func (l *VulkanLoader1_0) attemptAvailableLayers() (map[string]*common.LayerProperties, VkResult, error) {
	alloc := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(alloc)

	layerCount := (*Uint32)(alloc.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := l.driver.VkEnumerateInstanceLayerProperties(layerCount, nil)
	if err != nil || *layerCount == 0 {
		return nil, res, err
	}

	layersUnsafe := alloc.Malloc(int(*layerCount) * int(unsafe.Sizeof(C.VkLayerProperties{})))

	res, err = l.driver.VkEnumerateInstanceLayerProperties(layerCount, (*VkLayerProperties)(layersUnsafe))
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

func (l *VulkanLoader1_0) AvailableLayers() (map[string]*common.LayerProperties, VkResult, error) {
	// There may be a race condition that adds new available layers between getting the
	// layer count & pulling the layers, in which case, attemptAvailableLayers will return
	// VK_INCOMPLETE.  In this case, we should try again.
	var layers map[string]*common.LayerProperties
	var result VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == VKIncomplete) {
		layers, result, err = l.attemptAvailableLayers()
	}
	return layers, result, err
}

func (l *VulkanLoader1_0) CreateInstance(options *InstanceOptions) (Instance, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var instanceHandle VkInstance

	res, err := l.driver.VkCreateInstance((*VkInstanceCreateInfo)(createInfo), nil, &instanceHandle)
	if err != nil {
		return nil, res, err
	}

	instanceDriver, err := l.driver.CreateInstanceDriver((VkInstance)(instanceHandle))
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	return &vulkanInstance{
		driver: instanceDriver,
		handle: instanceHandle,
	}, res, nil
}

func (l *VulkanLoader1_0) CreateDevice(physicalDevice PhysicalDevice, options *DeviceOptions) (Device, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var deviceHandle VkDevice
	res, err := physicalDevice.Driver().VkCreateDevice(physicalDevice.Handle(), (*VkDeviceCreateInfo)(createInfo), nil, &deviceHandle)
	if err != nil {
		return nil, res, err
	}

	deviceDriver, err := physicalDevice.Driver().CreateDeviceDriver(deviceHandle)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	return &vulkanDevice{driver: deviceDriver, handle: deviceHandle}, res, nil
}

func (l *VulkanLoader1_0) CreateShaderModule(device Device, o *ShaderModuleOptions) (ShaderModule, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var shaderModule VkShaderModule
	res, err := device.Driver().VkCreateShaderModule(device.Handle(), (*VkShaderModuleCreateInfo)(createInfo), nil, &shaderModule)
	if err != nil {
		return nil, res, err
	}

	return &vulkanShaderModule{driver: device.Driver(), handle: shaderModule, device: device.Handle()}, res, nil
}

func (l *VulkanLoader1_0) CreateImageView(device Device, o *ImageViewOptions) (ImageView, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var imageViewHandle VkImageView

	res, err := device.Driver().VkCreateImageView(device.Handle(), (*VkImageViewCreateInfo)(createInfo), nil, &imageViewHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanImageView{driver: device.Driver(), handle: imageViewHandle, device: device.Handle()}, res, nil
}

func (l *VulkanLoader1_0) CreateSemaphore(device Device, o *SemaphoreOptions) (Semaphore, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var semaphoreHandle VkSemaphore

	res, err := device.Driver().VkCreateSemaphore(device.Handle(), (*VkSemaphoreCreateInfo)(createInfo), nil, &semaphoreHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanSemaphore{driver: device.Driver(), device: device.Handle(), handle: semaphoreHandle}, res, nil
}

func (l *VulkanLoader1_0) CreateFence(device Device, o *FenceOptions) (Fence, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var fenceHandle VkFence

	res, err := device.Driver().VkCreateFence(device.Handle(), (*VkFenceCreateInfo)(createInfo), nil, &fenceHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanFence{driver: device.Driver(), device: device.Handle(), handle: fenceHandle}, res, nil
}

func (l *VulkanLoader1_0) CreateBuffer(device Device, o *BufferOptions) (Buffer, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var buffer VkBuffer

	res, err := device.Driver().VkCreateBuffer(device.Handle(), (*VkBufferCreateInfo)(createInfo), nil, &buffer)
	if err != nil {
		return nil, res, err
	}

	return &vulkanBuffer{driver: device.Driver(), handle: buffer, device: device.Handle()}, res, nil
}

func (l *VulkanLoader1_0) CreateDescriptorSetLayout(device Device, o *DescriptorSetLayoutOptions) (DescriptorSetLayout, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var descriptorSetLayout VkDescriptorSetLayout

	res, err := device.Driver().VkCreateDescriptorSetLayout(device.Handle(), (*VkDescriptorSetLayoutCreateInfo)(createInfo), nil, &descriptorSetLayout)
	if err != nil {
		return nil, res, err
	}

	return &vulkanDescriptorSetLayout{
		driver: device.Driver(),
		device: device.Handle(),
		handle: descriptorSetLayout,
	}, res, nil
}

func (l *VulkanLoader1_0) CreateDescriptorPool(device Device, o *DescriptorPoolOptions) (DescriptorPool, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var descriptorPool VkDescriptorPool

	res, err := device.Driver().VkCreateDescriptorPool(device.Handle(), (*VkDescriptorPoolCreateInfo)(createInfo), nil, &descriptorPool)
	if err != nil {
		return nil, res, err
	}

	return &vulkanDescriptorPool{
		driver: device.Driver(),
		handle: descriptorPool,
		device: device.Handle(),
	}, res, nil
}

func (l *VulkanLoader1_0) CreateCommandPool(device Device, o *CommandPoolOptions) (CommandPool, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var cmdPoolHandle VkCommandPool
	res, err := device.Driver().VkCreateCommandPool(device.Handle(), (*VkCommandPoolCreateInfo)(createInfo), nil, &cmdPoolHandle)
	if err != nil {
		return nil, res, err
	}

	return &vulkanCommandPool{driver: device.Driver(), handle: cmdPoolHandle, device: device.Handle()}, res, nil
}

func (l *VulkanLoader1_0) CreateFrameBuffer(device Device, o *FramebufferOptions) (Framebuffer, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var framebuffer VkFramebuffer

	res, err := device.Driver().VkCreateFramebuffer(device.Handle(), (*VkFramebufferCreateInfo)(createInfo), nil, &framebuffer)
	if err != nil {
		return nil, res, err
	}

	return &vulkanFramebuffer{driver: device.Driver(), device: device.Handle(), handle: framebuffer}, res, nil
}

func (l *VulkanLoader1_0) CreateGraphicsPipelines(device Device, o []*GraphicsPipelineOptions) ([]Pipeline, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	pipelineCount := len(o)

	pipelineCreateInfosPtrUnsafe := arena.Malloc(pipelineCount * C.sizeof_struct_VkGraphicsPipelineCreateInfo)
	pipelineCreateInfosSlice := ([]C.VkGraphicsPipelineCreateInfo)(unsafe.Slice((*C.VkGraphicsPipelineCreateInfo)(pipelineCreateInfosPtrUnsafe), pipelineCount))
	for i := 0; i < pipelineCount; i++ {
		next, err := common.AllocNext(arena, o[i])
		if err != nil {
			return nil, VKErrorUnknown, err
		}

		err = o[i].populate(arena, &pipelineCreateInfosSlice[i], next)
		if err != nil {
			return nil, VKErrorUnknown, err
		}
	}

	pipelinePtr := (*VkPipeline)(arena.Malloc(pipelineCount * int(unsafe.Sizeof([1]VkPipeline{}))))

	res, err := device.Driver().VkCreateGraphicsPipelines(device.Handle(), nil, Uint32(pipelineCount), (*VkGraphicsPipelineCreateInfo)(pipelineCreateInfosPtrUnsafe), nil, pipelinePtr)
	if err != nil {
		return nil, res, err
	}

	var output []Pipeline
	pipelineSlice := ([]VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))
	for i := 0; i < pipelineCount; i++ {
		output = append(output, &vulkanPipeline{driver: device.Driver(), device: device.Handle(), handle: pipelineSlice[i]})
	}

	return output, res, nil
}

func (l *VulkanLoader1_0) CreateImage(device Device, o *ImageOptions) (Image, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var image VkImage
	res, err := device.Driver().VkCreateImage(device.Handle(), (*VkImageCreateInfo)(createInfo), nil, &image)
	if err != nil {
		return nil, res, err
	}

	return &vulkanImage{device: device.Handle(), handle: image, driver: device.Driver()}, res, nil
}

func (l *VulkanLoader1_0) CreatePipelineCache(device Device, o *PipelineCacheOptions) (PipelineCache, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var pipelineCache VkPipelineCache
	res, err := device.Driver().VkCreatePipelineCache(device.Handle(), (*VkPipelineCacheCreateInfo)(createInfo), nil, &pipelineCache)
	if err != nil {
		return nil, res, err
	}

	return &vulkanPipelineCache{driver: device.Driver(), handle: pipelineCache, device: device.Handle()}, res, nil
}

func (l *VulkanLoader1_0) CreatePipelineLayout(device Device, o *PipelineLayoutOptions) (PipelineLayout, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var pipelineLayout VkPipelineLayout
	res, err := device.Driver().VkCreatePipelineLayout(device.Handle(), (*VkPipelineLayoutCreateInfo)(createInfo), nil, &pipelineLayout)
	if err != nil {
		return nil, res, err
	}

	return &vulkanPipelineLayout{driver: device.Driver(), handle: pipelineLayout, device: device.Handle()}, res, nil
}

func (l *VulkanLoader1_0) CreateRenderPass(device Device, o *RenderPassOptions) (RenderPass, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var renderPass VkRenderPass

	res, err := device.Driver().VkCreateRenderPass(device.Handle(), (*VkRenderPassCreateInfo)(createInfo), nil, &renderPass)
	if err != nil {
		return nil, res, err
	}

	return &vulkanRenderPass{driver: device.Driver(), device: device.Handle(), handle: renderPass}, res, nil
}

func (l *VulkanLoader1_0) CreateSampler(device Device, o *SamplerOptions) (Sampler, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var sampler VkSampler

	res, err := device.Driver().VkCreateSampler(device.Handle(), (*VkSamplerCreateInfo)(createInfo), nil, &sampler)
	if err != nil {
		return nil, res, err
	}

	return &vulkanSampler{handle: sampler, driver: device.Driver(), device: device.Handle()}, VKSuccess, nil
}
