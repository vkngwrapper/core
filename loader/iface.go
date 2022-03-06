package loader

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"unsafe"
)

//go:generate mockgen -source ./iface.go -destination ../mocks/loader_mocks.go -package mocks

type Loader interface {
	Driver() driver.Driver
	Version() common.APIVersion

	AvailableExtensions() (map[string]*common.ExtensionProperties, common.VkResult, error)
	AvailableExtensionsForLayer(layerName string) (map[string]*common.ExtensionProperties, common.VkResult, error)
	AvailableLayers() (map[string]*common.LayerProperties, common.VkResult, error)

	CreateBuffer(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.BufferOptions) (core1_0.Buffer, common.VkResult, error)
	CreateBufferView(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.BufferViewOptions) (core1_0.BufferView, common.VkResult, error)
	CreateCommandPool(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.CommandPoolOptions) (core1_0.CommandPool, common.VkResult, error)
	CreateDescriptorPool(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.DescriptorPoolOptions) (core1_0.DescriptorPool, common.VkResult, error)
	CreateDescriptorSetLayout(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.DescriptorSetLayoutOptions) (core1_0.DescriptorSetLayout, common.VkResult, error)
	CreateDevice(physicalDevice core1_0.PhysicalDevice, allocationCallbacks *driver.AllocationCallbacks, options *core1_0.DeviceOptions) (core1_0.Device, common.VkResult, error)
	CreateEvent(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, options *core1_0.EventOptions) (core1_0.Event, common.VkResult, error)
	CreateFence(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.FenceOptions) (core1_0.Fence, common.VkResult, error)
	CreateFrameBuffer(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.FramebufferOptions) (core1_0.Framebuffer, common.VkResult, error)
	CreateGraphicsPipelines(device core1_0.Device, pipelineCache core1_0.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []core1_0.GraphicsPipelineOptions) ([]core1_0.Pipeline, common.VkResult, error)
	CreateComputePipelines(device core1_0.Device, pipelineCache core1_0.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []core1_0.ComputePipelineOptions) ([]core1_0.Pipeline, common.VkResult, error)
	CreateInstance(allocationCallbacks *driver.AllocationCallbacks, options *core1_0.InstanceOptions) (core1_0.Instance, common.VkResult, error)
	CreateImage(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, options *core1_0.ImageOptions) (core1_0.Image, common.VkResult, error)
	CreateImageView(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.ImageViewOptions) (core1_0.ImageView, common.VkResult, error)
	CreatePipelineCache(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.PipelineCacheOptions) (core1_0.PipelineCache, common.VkResult, error)
	CreatePipelineLayout(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.PipelineLayoutOptions) (core1_0.PipelineLayout, common.VkResult, error)
	CreateQueryPool(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.QueryPoolOptions) (core1_0.QueryPool, common.VkResult, error)
	CreateRenderPass(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.RenderPassOptions) (core1_0.RenderPass, common.VkResult, error)
	CreateSampler(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.SamplerOptions) (core1_0.Sampler, common.VkResult, error)
	CreateSemaphore(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.SemaphoreOptions) (core1_0.Semaphore, common.VkResult, error)
	CreateShaderModule(device core1_0.Device, allocationCallbacks *driver.AllocationCallbacks, o *core1_0.ShaderModuleOptions) (core1_0.ShaderModule, common.VkResult, error)

	AllocateCommandBuffers(o *core1_0.CommandBufferOptions) ([]core1_0.CommandBuffer, common.VkResult, error)
	FreeCommandBuffers(buffers []core1_0.CommandBuffer)
	AllocateDescriptorSets(o *core1_0.DescriptorSetOptions) ([]core1_0.DescriptorSet, common.VkResult, error)
	FreeDescriptorSets(sets []core1_0.DescriptorSet) (common.VkResult, error)
}

type Loader1_1 interface {
	Loader

	SomeOneOneMethod()
}

func CreateStaticLinkedLoader() (Loader, error) {
	return CreateLoaderFromProcAddr(unsafe.Pointer(C.vkGetInstanceProcAddr))
}

func CreateLoaderFromProcAddr(addr unsafe.Pointer) (Loader, error) {
	driver, err := driver.CreateDriverFromProcAddr(addr)
	if err != nil {
		return nil, err
	}

	return CreateLoaderFromDriver(driver)
}

func CreateLoaderFromDriver(driver driver.Driver) (Loader, error) {
	return NewLoader(driver), nil
}
