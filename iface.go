package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	"unsafe"
)

//go:generate mockgen -source ./iface.go -destination mocks/loader_mocks.go -package mocks

type Loader interface {
	Driver() driver.Driver
	Version() common.APIVersion

	Core1_1() Loader1_1

	AvailableExtensions() (map[string]*common.ExtensionProperties, common.VkResult, error)
	AvailableExtensionsForLayer(layerName string) (map[string]*common.ExtensionProperties, common.VkResult, error)
	AvailableLayers() (map[string]*common.LayerProperties, common.VkResult, error)

	CreateBuffer(device Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.BufferCreateOptions) (Buffer, common.VkResult, error)
	CreateBufferView(device Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.BufferViewCreateOptions) (BufferView, common.VkResult, error)
	CreateCommandPool(device Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.CommandPoolCreateOptions) (CommandPool, common.VkResult, error)
	CreateDescriptorPool(device Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.DescriptorPoolCreateOptions) (DescriptorPool, common.VkResult, error)
	CreateDescriptorSetLayout(device Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.DescriptorSetLayoutCreateOptions) (DescriptorSetLayout, common.VkResult, error)
	CreateDevice(physicalDevice PhysicalDevice, allocationCallbacks *driver.AllocationCallbacks, options core1_0.DeviceCreateOptions) (Device, common.VkResult, error)
	CreateEvent(device Device, allocationCallbacks *driver.AllocationCallbacks, options core1_0.EventCreateOptions) (Event, common.VkResult, error)
	CreateFence(device Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.FenceCreateOptions) (Fence, common.VkResult, error)
	CreateFrameBuffer(device Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.FramebufferCreateOptions) (Framebuffer, common.VkResult, error)
	CreateGraphicsPipelines(device Device, pipelineCache core1_0.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []core1_0.GraphicsPipelineCreateOptions) ([]Pipeline, common.VkResult, error)
	CreateComputePipelines(device Device, pipelineCache core1_0.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []core1_0.ComputePipelineCreateOptions) ([]Pipeline, common.VkResult, error)
	CreateInstance(allocationCallbacks *driver.AllocationCallbacks, options core1_0.InstanceCreateOptions) (Instance, common.VkResult, error)
	CreateImage(device Device, allocationCallbacks *driver.AllocationCallbacks, options core1_0.ImageCreateOptions) (Image, common.VkResult, error)
	CreateImageView(device Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.ImageViewCreateOptions) (ImageView, common.VkResult, error)
	CreatePipelineCache(device Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.PipelineCacheCreateOptions) (PipelineCache, common.VkResult, error)
	CreatePipelineLayout(device Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.PipelineLayoutCreateOptions) (PipelineLayout, common.VkResult, error)
	CreateQueryPool(device Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.QueryPoolCreateOptions) (QueryPool, common.VkResult, error)
	CreateRenderPass(device Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.RenderPassCreateOptions) (RenderPass, common.VkResult, error)
	CreateSampler(device Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.SamplerCreateOptions) (Sampler, common.VkResult, error)
	CreateSemaphore(device Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.SemaphoreCreateOptions) (Semaphore, common.VkResult, error)
	CreateShaderModule(device Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.ShaderModuleCreateOptions) (ShaderModule, common.VkResult, error)

	GetQueue(device Device, queueFamilyIndex int, queueIndex int) Queue
	AllocateMemory(device Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.MemoryAllocateOptions) (DeviceMemory, common.VkResult, error)
	FreeMemory(deviceMemory DeviceMemory, allocationCallbacks *driver.AllocationCallbacks)
	PhysicalDevices(instance Instance) ([]PhysicalDevice, common.VkResult, error)

	AllocateCommandBuffers(o core1_0.CommandBufferOptions) ([]CommandBuffer, common.VkResult, error)
	FreeCommandBuffers(buffers []CommandBuffer)
	AllocateDescriptorSets(o core1_0.DescriptorSetOptions) ([]DescriptorSet, common.VkResult, error)
	FreeDescriptorSets(sets []DescriptorSet) (common.VkResult, error)
}

type Loader1_1 interface {
	CreateDescriptorUpdateTemplate(device Device, o core1_1.DescriptorUpdateTemplateCreateOptions, allocator *driver.AllocationCallbacks) (DescriptorUpdateTemplate, common.VkResult, error)
	CreateSamplerYcbcrConversion(device Device, o core1_1.SamplerYcbcrConversionCreateOptions, allocator *driver.AllocationCallbacks) (SamplerYcbcrConversion, common.VkResult, error)

	GetQueue(device Device, o core1_1.DeviceQueueOptions) (Queue, error)
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
	return NewLoader(driver), nil
}
