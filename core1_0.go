package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	internal1_0 "github.com/CannibalVox/VKng/core/internal/core1_0"
	"github.com/CannibalVox/VKng/core/internal/objects"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type VulkanLoader1_0 struct {
	driver driver.Driver

	loader1_1 Loader1_1
}

func NewLoader(
	driver driver.Driver,
) *VulkanLoader1_0 {

	var loader1_1 Loader1_1

	if driver.Version().IsAtLeast(common.Vulkan1_1) {
		loader1_1 = &VulkanLoader1_1{
			driver: driver,
		}
	}

	return &VulkanLoader1_0{
		driver:    driver,
		loader1_1: loader1_1,
	}
}

func (l *VulkanLoader1_0) Version() common.APIVersion {
	return l.driver.Version()
}

func (l *VulkanLoader1_0) Driver() driver.Driver {
	return l.driver
}

func (l *VulkanLoader1_0) Core1_1() Loader1_1 {
	return l.loader1_1
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
	if err != nil || res == core1_0.VKIncomplete {
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
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
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
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
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
	if err != nil || res == core1_0.VKIncomplete {
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
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		layers, result, err = l.attemptAvailableLayers()
	}
	return layers, result, err
}

func (l *VulkanLoader1_0) CreateBufferView(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, options *core1_0.BufferViewOptions) (core1_0.BufferView, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var bufferViewHandle driver.VkBufferView

	res, err := device.Driver().VkCreateBufferView(device.Handle(), (*driver.VkBufferViewCreateInfo)(createInfo), allocationCallbacks.Handle(), &bufferViewHandle)
	if err != nil {
		return nil, res, err
	}

	bufferView := objects.CreateBufferView(device, bufferViewHandle)

	return bufferView, res, nil
}

func (l *VulkanLoader1_0) CreateInstance(allocationCallbacks *driver.AllocationCallbacks, options *core1_0.InstanceOptions) (core1_0.Instance, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var instanceHandle driver.VkInstance

	res, err := l.driver.VkCreateInstance((*driver.VkInstanceCreateInfo)(createInfo), allocationCallbacks.Handle(), &instanceHandle)
	if err != nil {
		return nil, res, err
	}

	instanceDriver, err := l.driver.CreateInstanceDriver((driver.VkInstance)(instanceHandle))
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	version := l.Version().Min(options.VulkanVersion)
	instance := objects.CreateInstance(instanceDriver, instanceHandle, version)

	return instance, res, nil
}

func (l *VulkanLoader1_0) CreateDevice(physicalDevice core1_0.PhysicalDevice, allocationCallbacks *driver.AllocationCallbacks, options *core1_0.DeviceOptions) (core1_0.Device, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var deviceHandle driver.VkDevice
	res, err := physicalDevice.Driver().VkCreateDevice(physicalDevice.Handle(), (*driver.VkDeviceCreateInfo)(createInfo), allocationCallbacks.Handle(), &deviceHandle)
	if err != nil {
		return nil, res, err
	}

	deviceDriver, err := physicalDevice.Driver().CreateDeviceDriver(deviceHandle)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	device := objects.CreateDevice(deviceDriver, deviceHandle, physicalDevice.APIVersion())

	return device, res, nil
}

func (l *VulkanLoader1_0) CreateShaderModule(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.ShaderModuleOptions) (core1_0.ShaderModule, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var shaderModuleHandle driver.VkShaderModule
	res, err := device.Driver().VkCreateShaderModule(device.Handle(), (*driver.VkShaderModuleCreateInfo)(createInfo), allocationCallbacks.Handle(), &shaderModuleHandle)
	if err != nil {
		return nil, res, err
	}

	shaderModule := objects.CreateShaderModule(device, shaderModuleHandle)

	return shaderModule, res, nil
}

func (l *VulkanLoader1_0) CreateImageView(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.ImageViewOptions) (core1_0.ImageView, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var imageViewHandle driver.VkImageView

	res, err := device.Driver().VkCreateImageView(device.Handle(), (*driver.VkImageViewCreateInfo)(createInfo), allocationCallbacks.Handle(), &imageViewHandle)
	if err != nil {
		return nil, res, err
	}

	imageView := objects.CreateImageView(device, imageViewHandle)

	return imageView, res, nil
}

func (l *VulkanLoader1_0) CreateSemaphore(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.SemaphoreOptions) (core1_0.Semaphore, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var semaphoreHandle driver.VkSemaphore

	res, err := device.Driver().VkCreateSemaphore(device.Handle(), (*driver.VkSemaphoreCreateInfo)(createInfo), allocationCallbacks.Handle(), &semaphoreHandle)
	if err != nil {
		return nil, res, err
	}

	semaphore := objects.CreateSemaphore(device, semaphoreHandle)

	return semaphore, res, nil
}

func (l *VulkanLoader1_0) CreateFence(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.FenceOptions) (core1_0.Fence, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var fenceHandle driver.VkFence

	res, err := device.Driver().VkCreateFence(device.Handle(), (*driver.VkFenceCreateInfo)(createInfo), allocationCallbacks.Handle(), &fenceHandle)
	if err != nil {
		return nil, res, err
	}

	fence := objects.CreateFence(device, fenceHandle)

	return fence, res, nil
}

func (l *VulkanLoader1_0) CreateBuffer(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.BufferOptions) (core1_0.Buffer, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var bufferHandle driver.VkBuffer

	res, err := device.Driver().VkCreateBuffer(device.Handle(), (*driver.VkBufferCreateInfo)(createInfo), allocationCallbacks.Handle(), &bufferHandle)
	if err != nil {
		return nil, res, err
	}

	buffer := objects.CreateBuffer(device, bufferHandle)

	return buffer, res, nil
}

func (l *VulkanLoader1_0) CreateDescriptorSetLayout(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.DescriptorSetLayoutOptions) (core1_0.DescriptorSetLayout, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var descriptorSetLayoutHandle driver.VkDescriptorSetLayout

	res, err := device.Driver().VkCreateDescriptorSetLayout(device.Handle(), (*driver.VkDescriptorSetLayoutCreateInfo)(createInfo), allocationCallbacks.Handle(), &descriptorSetLayoutHandle)
	if err != nil {
		return nil, res, err
	}

	descriptorSetLayout := objects.CreateDescriptorSetLayout(device, descriptorSetLayoutHandle)

	return descriptorSetLayout, res, nil
}

func (l *VulkanLoader1_0) CreateDescriptorPool(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.DescriptorPoolOptions) (core1_0.DescriptorPool, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var descriptorPoolHandle driver.VkDescriptorPool

	res, err := device.Driver().VkCreateDescriptorPool(device.Handle(), (*driver.VkDescriptorPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &descriptorPoolHandle)
	if err != nil {
		return nil, res, err
	}

	descriptorPool := objects.CreateDescriptorPool(device, descriptorPoolHandle)

	return descriptorPool, res, nil
}

func (l *VulkanLoader1_0) CreateCommandPool(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.CommandPoolOptions) (core1_0.CommandPool, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var cmdPoolHandle driver.VkCommandPool
	res, err := device.Driver().VkCreateCommandPool(device.Handle(), (*driver.VkCommandPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &cmdPoolHandle)
	if err != nil {
		return nil, res, err
	}

	commandPool := objects.CreateCommandPool(device, cmdPoolHandle)

	return commandPool, res, nil
}

func (l *VulkanLoader1_0) CreateEvent(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.EventOptions) (core1_0.Event, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var eventHandle driver.VkEvent
	res, err := device.Driver().VkCreateEvent(device.Handle(), (*driver.VkEventCreateInfo)(createInfo), allocationCallbacks.Handle(), &eventHandle)
	if err != nil {
		return nil, res, err
	}

	event := objects.CreateEvent(device, eventHandle)

	return event, res, nil
}

func (l *VulkanLoader1_0) CreateFrameBuffer(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.FramebufferOptions) (core1_0.Framebuffer, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var framebufferHandle driver.VkFramebuffer

	res, err := device.Driver().VkCreateFramebuffer(device.Handle(), (*driver.VkFramebufferCreateInfo)(createInfo), allocationCallbacks.Handle(), &framebufferHandle)
	if err != nil {
		return nil, res, err
	}

	framebuffer := objects.CreateFramebuffer(device, framebufferHandle)

	return framebuffer, res, nil
}

func (l *VulkanLoader1_0) CreateGraphicsPipelines(device core1_0.Device, pipelineCache core1_0.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []core1_0.GraphicsPipelineOptions) ([]core1_0.Pipeline, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	pipelineCount := len(o)

	pipelineCreateInfosPtr, err := common.AllocOptionSlice[C.VkGraphicsPipelineCreateInfo, core1_0.GraphicsPipelineOptions](arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
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

	var output []core1_0.Pipeline
	pipelineSlice := ([]driver.VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))

	for i := 0; i < pipelineCount; i++ {
		pipeline := objects.CreatePipeline(device, pipelineSlice[i])
		output = append(output, pipeline)
	}

	return output, res, nil
}

func (l *VulkanLoader1_0) CreateComputePipelines(device core1_0.Device, pipelineCache core1_0.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []core1_0.ComputePipelineOptions) ([]core1_0.Pipeline, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	pipelineCount := len(o)

	pipelineCreateInfosPtr, err := common.AllocOptionSlice[C.VkComputePipelineCreateInfo, core1_0.ComputePipelineOptions](arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
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

	var output []core1_0.Pipeline
	pipelineSlice := ([]driver.VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))

	for i := 0; i < pipelineCount; i++ {
		pipeline := objects.CreatePipeline(device, pipelineSlice[i])

		output = append(output, pipeline)
	}

	return output, res, nil
}

func (l *VulkanLoader1_0) CreateImage(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.ImageOptions) (core1_0.Image, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var imageHandle driver.VkImage
	res, err := device.Driver().VkCreateImage(device.Handle(), (*driver.VkImageCreateInfo)(createInfo), allocationCallbacks.Handle(), &imageHandle)
	if err != nil {
		return nil, res, err
	}

	image := objects.CreateImage(device, imageHandle)

	return image, res, nil
}

func (l *VulkanLoader1_0) CreatePipelineCache(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.PipelineCacheOptions) (core1_0.PipelineCache, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var pipelineCacheHandle driver.VkPipelineCache
	res, err := device.Driver().VkCreatePipelineCache(device.Handle(), (*driver.VkPipelineCacheCreateInfo)(createInfo), allocationCallbacks.Handle(), &pipelineCacheHandle)
	if err != nil {
		return nil, res, err
	}

	pipelineCache := objects.CreatePipelineCache(device, pipelineCacheHandle)

	return pipelineCache, res, nil
}

func (l *VulkanLoader1_0) CreatePipelineLayout(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.PipelineLayoutOptions) (core1_0.PipelineLayout, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var pipelineLayoutHandle driver.VkPipelineLayout
	res, err := device.Driver().VkCreatePipelineLayout(device.Handle(), (*driver.VkPipelineLayoutCreateInfo)(createInfo), allocationCallbacks.Handle(), &pipelineLayoutHandle)
	if err != nil {
		return nil, res, err
	}

	pipelineLayout := objects.CreatePipelineLayout(device, pipelineLayoutHandle)

	return pipelineLayout, res, nil
}

func (l *VulkanLoader1_0) CreateQueryPool(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.QueryPoolOptions) (core1_0.QueryPool, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var queryPoolHandle driver.VkQueryPool

	res, err := device.Driver().VkCreateQueryPool(device.Handle(), (*driver.VkQueryPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &queryPoolHandle)
	if err != nil {
		return nil, res, err
	}

	queryPool := objects.CreateQueryPool(device, queryPoolHandle)
	return queryPool, res, nil
}

func (l *VulkanLoader1_0) CreateRenderPass(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.RenderPassOptions) (core1_0.RenderPass, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var renderPassHandle driver.VkRenderPass

	res, err := device.Driver().VkCreateRenderPass(device.Handle(), (*driver.VkRenderPassCreateInfo)(createInfo), allocationCallbacks.Handle(), &renderPassHandle)
	if err != nil {
		return nil, res, err
	}

	renderPass := objects.CreateRenderPass(device, renderPassHandle)

	return renderPass, res, nil
}

func (l *VulkanLoader1_0) CreateSampler(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.SamplerOptions) (core1_0.Sampler, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var samplerHandle driver.VkSampler

	res, err := device.Driver().VkCreateSampler(device.Handle(), (*driver.VkSamplerCreateInfo)(createInfo), allocationCallbacks.Handle(), &samplerHandle)
	if err != nil {
		return nil, res, err
	}

	sampler := objects.CreateSampler(device, samplerHandle)

	return sampler, res, nil
}

// Free a slice of command buffers which should all have the same device/driver/pool
// guaranteed to have at least one element
func (l *VulkanLoader1_0) freeCommandBufferSlice(buffers []core1_0.CommandBuffer) {
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

	objStore := l.driver.ObjectStore()
	for i := 0; i < bufferCount; i++ {
		objStore.Delete(driver.VulkanHandle(buffers[i].Handle()), buffers[i])
	}
}

func (l *VulkanLoader1_0) FreeCommandBuffers(buffers []core1_0.CommandBuffer) {
	bufferCount := len(buffers)
	if bufferCount == 0 {
		return
	}

	multimap := make(map[driver.VkCommandPool][]core1_0.CommandBuffer)
	for _, buffer := range buffers {
		poolHandle := buffer.CommandPoolHandle()
		existingSet := multimap[poolHandle]
		multimap[poolHandle] = append(existingSet, buffer)
	}

	for _, setBuffers := range multimap {
		l.freeCommandBufferSlice(setBuffers)
	}
}

func (l *VulkanLoader1_0) AllocateCommandBuffers(o *core1_0.CommandBufferOptions) ([]core1_0.CommandBuffer, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	if o.CommandPool == nil {
		return nil, core1_0.VKErrorUnknown, errors.New("no command pool provided to allocate from")
	}

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	device := o.CommandPool.Device()

	commandBufferPtr := (*driver.VkCommandBuffer)(arena.Malloc(o.BufferCount * int(unsafe.Sizeof([1]driver.VkCommandBuffer{}))))

	res, err := o.CommandPool.Driver().VkAllocateCommandBuffers(device, (*driver.VkCommandBufferAllocateInfo)(createInfo), commandBufferPtr)
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	commandBufferArray := ([]driver.VkCommandBuffer)(unsafe.Slice(commandBufferPtr, o.BufferCount))
	var result []core1_0.CommandBuffer

	for i := 0; i < o.BufferCount; i++ {
		commandBuffer := objects.CreateCommandBuffer(o.CommandPool, commandBufferArray[i])

		result = append(result, commandBuffer)
	}

	return result, res, nil
}

func (l *VulkanLoader1_0) AllocateDescriptorSets(o *core1_0.DescriptorSetOptions) ([]core1_0.DescriptorSet, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	if o.DescriptorPool == nil {
		return nil, core1_0.VKErrorUnknown, errors.New("no descriptor pool provided to allocate from")
	}

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	device := o.DescriptorPool.DeviceHandle()
	poolDriver := o.DescriptorPool.Driver()

	setCount := len(o.AllocationLayouts)
	descriptorSets := (*driver.VkDescriptorSet)(arena.Malloc(setCount * int(unsafe.Sizeof([1]C.VkDescriptorSet{}))))

	res, err := poolDriver.VkAllocateDescriptorSets(device, (*driver.VkDescriptorSetAllocateInfo)(createInfo), descriptorSets)
	if err != nil {
		return nil, res, err
	}

	var sets []core1_0.DescriptorSet
	descriptorSetSlice := ([]driver.VkDescriptorSet)(unsafe.Slice(descriptorSets, setCount))

	for i := 0; i < setCount; i++ {
		descriptorSet := objects.CreateDescriptorSet(o.DescriptorPool, descriptorSetSlice[i])

		sets = append(sets, descriptorSet)
	}

	return sets, res, nil
}

// Free a slice of descriptor sets which should all have the same device/driver/pool
// guaranteed to have at least one element
func (l *VulkanLoader1_0) freeDescriptorSetSlice(sets []core1_0.DescriptorSet) (common.VkResult, error) {
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

	res, err := setDriver.VkFreeDescriptorSets(device, pool, driver.Uint32(setSize), arrayPtr)
	if err != nil {
		return res, err
	}

	objStore := setDriver.ObjectStore()
	for i := 0; i < setSize; i++ {
		objStore.Delete(driver.VulkanHandle(sets[i].Handle()), sets[i])
	}

	return res, nil
}

func (l *VulkanLoader1_0) FreeDescriptorSets(sets []core1_0.DescriptorSet) (common.VkResult, error) {
	poolMultimap := make(map[driver.VkDescriptorPool][]core1_0.DescriptorSet)

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

func (l *VulkanLoader1_0) GetQueue(device core1_0.Device, queueFamilyIndex int, queueIndex int) core1_0.Queue {

	var queueHandle driver.VkQueue

	device.Driver().VkGetDeviceQueue(device.Handle(), driver.Uint32(queueFamilyIndex), driver.Uint32(queueIndex), &queueHandle)

	queue := objects.CreateQueue(device, queueHandle)

	return queue
}

func (l *VulkanLoader1_0) AllocateMemory(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.DeviceMemoryOptions) (core1_0.DeviceMemory, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var deviceMemoryHandle driver.VkDeviceMemory

	deviceDriver := device.Driver()
	deviceHandle := device.Handle()

	res, err := deviceDriver.VkAllocateMemory(deviceHandle, (*driver.VkMemoryAllocateInfo)(createInfo), allocationCallbacks.Handle(), &deviceMemoryHandle)
	if err != nil {
		return nil, res, err
	}

	deviceMemory := objects.CreateDeviceMemory(device, deviceMemoryHandle, o.AllocationSize)

	return deviceMemory, res, nil
}

func (l *VulkanLoader1_0) FreeMemory(deviceMemory core1_0.DeviceMemory, allocationCallbacks *driver.AllocationCallbacks) {
	// This is really only here for a kind of API symmetry
	deviceMemory.Free(allocationCallbacks)
}

func (l *VulkanLoader1_0) PhysicalDevices(instance core1_0.Instance) ([]core1_0.PhysicalDevice, common.VkResult, error) {
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
	var devices []core1_0.PhysicalDevice
	for ind := uint32(0); ind < goCount; ind++ {
		propertiesUnsafe := allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceProperties{})))

		instance.Driver().VkGetPhysicalDeviceProperties(deviceHandles[ind], (*driver.VkPhysicalDeviceProperties)(propertiesUnsafe))

		properties := internal1_0.CreatePhysicalDeviceProperties(propertiesUnsafe)

		version := instance.APIVersion().Min(properties.APIVersion)
		physicalDevice := objects.CreatePhysicalDevice(instance, deviceHandles[ind], version)

		devices = append(devices, physicalDevice)
	}

	return devices, res, nil
}
