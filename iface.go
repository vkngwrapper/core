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
	APIVersion() common.APIVersion

	Core1_1() Loader1_1

	AvailableExtensions() (map[string]*common.ExtensionProperties, common.VkResult, error)
	AvailableExtensionsForLayer(layerName string) (map[string]*common.ExtensionProperties, common.VkResult, error)
	AvailableLayers() (map[string]*common.LayerProperties, common.VkResult, error)

	CreateBuffer(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.BufferCreateOptions) (core1_0.Buffer, common.VkResult, error)
	CreateBufferView(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.BufferViewCreateOptions) (core1_0.BufferView, common.VkResult, error)
	CreateCommandPool(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.CommandPoolCreateOptions) (core1_0.CommandPool, common.VkResult, error)
	CreateDescriptorPool(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.DescriptorPoolCreateOptions) (core1_0.DescriptorPool, common.VkResult, error)
	CreateDescriptorSetLayout(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.DescriptorSetLayoutCreateOptions) (core1_0.DescriptorSetLayout, common.VkResult, error)
	CreateDevice(physicalDevice core1_0.PhysicalDevice, allocationCallbacks *driver.AllocationCallbacks, options core1_0.DeviceCreateOptions) (core1_0.Device, common.VkResult, error)
	CreateEvent(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, options core1_0.EventCreateOptions) (core1_0.Event, common.VkResult, error)
	CreateFence(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.FenceCreateOptions) (core1_0.Fence, common.VkResult, error)
	CreateFrameBuffer(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.FramebufferCreateOptions) (core1_0.Framebuffer, common.VkResult, error)
	CreateGraphicsPipelines(device core1_0.Device, pipelineCache core1_0.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []core1_0.GraphicsPipelineCreateOptions) ([]core1_0.Pipeline, common.VkResult, error)
	CreateComputePipelines(device core1_0.Device, pipelineCache core1_0.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []core1_0.ComputePipelineCreateOptions) ([]core1_0.Pipeline, common.VkResult, error)
	CreateInstance(allocationCallbacks *driver.AllocationCallbacks, options core1_0.InstanceCreateOptions) (core1_0.Instance, common.VkResult, error)
	CreateImage(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, options core1_0.ImageCreateOptions) (core1_0.Image, common.VkResult, error)
	CreateImageView(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.ImageViewCreateOptions) (core1_0.ImageView, common.VkResult, error)
	CreatePipelineCache(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.PipelineCacheCreateOptions) (core1_0.PipelineCache, common.VkResult, error)
	CreatePipelineLayout(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.PipelineLayoutCreateOptions) (core1_0.PipelineLayout, common.VkResult, error)
	CreateQueryPool(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.QueryPoolCreateOptions) (core1_0.QueryPool, common.VkResult, error)
	CreateRenderPass(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.RenderPassCreateOptions) (core1_0.RenderPass, common.VkResult, error)
	CreateSampler(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.SamplerCreateOptions) (core1_0.Sampler, common.VkResult, error)
	CreateSemaphore(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.SemaphoreCreateOptions) (core1_0.Semaphore, common.VkResult, error)
	CreateShaderModule(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.ShaderModuleCreateOptions) (core1_0.ShaderModule, common.VkResult, error)

	GetQueue(device core1_0.Device, queueFamilyIndex int, queueIndex int) core1_0.Queue
	AllocateMemory(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.MemoryAllocateOptions) (core1_0.DeviceMemory, common.VkResult, error)
	FreeMemory(deviceMemory core1_0.DeviceMemory, allocationCallbacks *driver.AllocationCallbacks)
	PhysicalDevices(instance core1_0.Instance) ([]core1_0.PhysicalDevice, common.VkResult, error)

	AllocateCommandBuffers(o core1_0.CommandBufferAllocateOptions) ([]core1_0.CommandBuffer, common.VkResult, error)
	FreeCommandBuffers(buffers []core1_0.CommandBuffer)
	AllocateDescriptorSets(o core1_0.DescriptorSetAllocateOptions) ([]core1_0.DescriptorSet, common.VkResult, error)
	FreeDescriptorSets(sets []core1_0.DescriptorSet) (common.VkResult, error)
}

type Loader1_1 interface {
	CreateDescriptorUpdateTemplate(device core1_0.Device, o core1_1.DescriptorUpdateTemplateCreateOptions, allocator *driver.AllocationCallbacks) (core1_1.DescriptorUpdateTemplate, common.VkResult, error)
	CreateSamplerYcbcrConversion(device core1_0.Device, o core1_1.SamplerYcbcrConversionCreateOptions, allocator *driver.AllocationCallbacks) (core1_1.SamplerYcbcrConversion, common.VkResult, error)

	GetQueue(device core1_0.Device, o core1_1.DeviceQueueOptions) (core1_0.Queue, error)
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
