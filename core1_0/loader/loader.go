package loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/internal/universal/loaderiface"
	"unsafe"
)

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
	return loaderiface.CreateLoaderFromDriver[core1_0.Buffer, core1_0.BufferView, core1_0.CommandPool,
		core1_0.DescriptorPool, core1_0.DescriptorSet, core1_0.Device, core1_0.Event, core1_0.Fence, core1_0.Framebuffer,
		core1_0.Instance, core1_0.Image, core1_0.ImageView, core1_0.PipelineCache, core1_0.PipelineLayout, core1_0.QueryPool,
		core1_0.RenderPass, core1_0.Sampler, core1_0.Semaphore, core1_0.ShaderModule, core1_0.CommandBuffer,
		core1_0.DeviceMemory, core1_0.DescriptorSetLayout, core1_0.Pipeline, core1_0.PhysicalDevice,
		core1_0.Queue](driver), nil
}

type Loader loaderiface.Loader1_0[core1_0.Buffer, core1_0.BufferView, core1_0.CommandPool,
	core1_0.DescriptorPool, core1_0.DescriptorSet, core1_0.Device, core1_0.Event, core1_0.Fence, core1_0.Framebuffer,
	core1_0.Instance, core1_0.Image, core1_0.ImageView, core1_0.PipelineCache, core1_0.PipelineLayout, core1_0.QueryPool,
	core1_0.RenderPass, core1_0.Sampler, core1_0.Semaphore, core1_0.ShaderModule, core1_0.CommandBuffer,
	core1_0.DeviceMemory, core1_0.DescriptorSetLayout, core1_0.Pipeline, core1_0.PhysicalDevice,
	core1_0.Queue]
