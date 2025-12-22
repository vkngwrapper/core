package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroyDevice(device types.Device, callbacks *driver.AllocationCallbacks) {
	if device.Handle() == 0 {
		panic("device was uninitialized")
	}

	v.Driver.VkDestroyDevice(device.Handle(), callbacks.Handle())
}

func (v *Vulkan) DeviceWaitIdle(device types.Device) (common.VkResult, error) {
	if device.Handle() == 0 {
		return core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	return v.Driver.VkDeviceWaitIdle(device.Handle())
}

func (v *Vulkan) WaitForFences(device types.Device, waitForAll bool, timeout time.Duration, fences []types.Fence) (common.VkResult, error) {
	if device.Handle() == 0 {
		return core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver.VkFence)(fenceUnsafePtr)

	fenceSlice := ([]driver.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		if fences[i].Handle() == 0 {
			panic(fmt.Sprintf("element %d of slice fences is uninitialized", i))
		}
		fenceSlice[i] = fences[i].Handle()
	}

	waitAllConst := C.VK_FALSE
	if waitForAll {
		waitAllConst = C.VK_TRUE
	}

	return v.Driver.VkWaitForFences(device.Handle(), driver.Uint32(fenceCount), fencePtr, driver.VkBool32(waitAllConst), driver.Uint64(common.TimeoutNanoseconds(timeout)))
}

func (v *Vulkan) ResetFences(device types.Device, fences []types.Fence) (common.VkResult, error) {
	if device.Handle() == 0 {
		return core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver.VkFence)(fenceUnsafePtr)
	fenceSlice := ([]driver.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		if fences[i].Handle() == 0 {
			panic(fmt.Sprintf("element %d of slice fences is uninitialized", i))
		}
		fenceSlice[i] = fences[i].Handle()
	}

	return v.Driver.VkResetFences(device.Handle(), driver.Uint32(fenceCount), fencePtr)
}

func (v *Vulkan) UpdateDescriptorSets(device types.Device, writes []core1_0.WriteDescriptorSet, copies []core1_0.CopyDescriptorSet) error {
	if device.Handle() == 0 {
		return errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	writeCount := len(writes)
	copyCount := len(copies)

	var err error
	var writePtr *C.VkWriteDescriptorSet
	var copyPtr *C.VkCopyDescriptorSet

	if writeCount > 0 {
		writePtr, err = common.AllocOptionSlice[C.VkWriteDescriptorSet, core1_0.WriteDescriptorSet](arena, writes)
		if err != nil {
			return err
		}
	}

	if copyCount > 0 {
		copyPtr, err = common.AllocOptionSlice[C.VkCopyDescriptorSet, core1_0.CopyDescriptorSet](arena, copies)
		if err != nil {
			return err
		}
	}

	v.Driver.VkUpdateDescriptorSets(device.Handle(), driver.Uint32(writeCount), (*driver.VkWriteDescriptorSet)(unsafe.Pointer(writePtr)), driver.Uint32(copyCount), (*driver.VkCopyDescriptorSet)(unsafe.Pointer(copyPtr)))
	return nil
}

func (v *Vulkan) FlushMappedMemoryRanges(device types.Device, ranges []core1_0.MappedMemoryRange) (common.VkResult, error) {
	if device.Handle() == 0 {
		return core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	rangeCount := len(ranges)
	createInfos, err := common.AllocOptionSlice[C.VkMappedMemoryRange, core1_0.MappedMemoryRange](arena, ranges)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return v.Driver.VkFlushMappedMemoryRanges(device.Handle(), driver.Uint32(rangeCount), (*driver.VkMappedMemoryRange)(unsafe.Pointer(createInfos)))
}

func (v *Vulkan) InvalidateMappedMemoryRanges(device types.Device, ranges []core1_0.MappedMemoryRange) (common.VkResult, error) {
	if device.Handle() == 0 {
		return core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	rangeCount := len(ranges)
	createInfos, err := common.AllocOptionSlice[C.VkMappedMemoryRange, core1_0.MappedMemoryRange](arena, ranges)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return v.Driver.VkInvalidateMappedMemoryRanges(device.Handle(), driver.Uint32(rangeCount), (*driver.VkMappedMemoryRange)(unsafe.Pointer(createInfos)))
}

func (v *Vulkan) CreateBufferView(device types.Device, allocationCallbacks *driver.AllocationCallbacks, options core1_0.BufferViewCreateInfo) (types.BufferView, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.BufferView{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return types.BufferView{}, core1_0.VKErrorUnknown, err
	}

	var bufferViewHandle driver.VkBufferView

	res, err := v.Driver.VkCreateBufferView(device.Handle(), (*driver.VkBufferViewCreateInfo)(createInfo), allocationCallbacks.Handle(), &bufferViewHandle)
	if err != nil {
		return types.BufferView{}, res, err
	}

	bufferView := types.InternalBufferView(device.Handle(), bufferViewHandle, device.APIVersion())

	return bufferView, res, nil
}

func (v *Vulkan) CreateShaderModule(device types.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.ShaderModuleCreateInfo) (types.ShaderModule, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.ShaderModule{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.ShaderModule{}, core1_0.VKErrorUnknown, err
	}

	var shaderModuleHandle driver.VkShaderModule
	res, err := v.Driver.VkCreateShaderModule(device.Handle(), (*driver.VkShaderModuleCreateInfo)(createInfo), allocationCallbacks.Handle(), &shaderModuleHandle)
	if err != nil {
		return types.ShaderModule{}, res, err
	}

	shaderModule := types.InternalShaderModule(device.Handle(), shaderModuleHandle, device.APIVersion())

	return shaderModule, res, nil
}

func (v *Vulkan) CreateImageView(device types.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.ImageViewCreateInfo) (types.ImageView, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.ImageView{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.ImageView{}, core1_0.VKErrorUnknown, err
	}

	var imageViewHandle driver.VkImageView

	res, err := v.Driver.VkCreateImageView(device.Handle(), (*driver.VkImageViewCreateInfo)(createInfo), allocationCallbacks.Handle(), &imageViewHandle)
	if err != nil {
		return types.ImageView{}, res, err
	}

	imageView := types.InternalImageView(device.Handle(), imageViewHandle, device.APIVersion())

	return imageView, res, nil
}

func (v *Vulkan) CreateSemaphore(device types.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.SemaphoreCreateInfo) (types.Semaphore, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.Semaphore{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.Semaphore{}, core1_0.VKErrorUnknown, err
	}

	var semaphoreHandle driver.VkSemaphore

	res, err := v.Driver.VkCreateSemaphore(device.Handle(), (*driver.VkSemaphoreCreateInfo)(createInfo), allocationCallbacks.Handle(), &semaphoreHandle)
	if err != nil {
		return types.Semaphore{}, res, err
	}

	semaphore := types.InternalSemaphore(device.Handle(), semaphoreHandle, device.APIVersion())

	return semaphore, res, nil
}

func (v *Vulkan) CreateFence(device types.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.FenceCreateInfo) (types.Fence, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.Fence{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.Fence{}, core1_0.VKErrorUnknown, err
	}

	var fenceHandle driver.VkFence

	res, err := v.Driver.VkCreateFence(device.Handle(), (*driver.VkFenceCreateInfo)(createInfo), allocationCallbacks.Handle(), &fenceHandle)
	if err != nil {
		return types.Fence{}, res, err
	}

	fence := types.InternalFence(device.Handle(), fenceHandle, device.APIVersion())

	return fence, res, nil
}

func (v *Vulkan) CreateBuffer(device types.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.BufferCreateInfo) (types.Buffer, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.Buffer{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.Buffer{}, core1_0.VKErrorUnknown, err
	}

	var bufferHandle driver.VkBuffer

	res, err := v.Driver.VkCreateBuffer(device.Handle(), (*driver.VkBufferCreateInfo)(createInfo), allocationCallbacks.Handle(), &bufferHandle)
	if err != nil {
		return types.Buffer{}, res, err
	}

	buffer := types.InternalBuffer(device.Handle(), bufferHandle, device.APIVersion())

	return buffer, res, nil
}

func (v *Vulkan) CreateDescriptorSetLayout(device types.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.DescriptorSetLayoutCreateInfo) (types.DescriptorSetLayout, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.DescriptorSetLayout{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.DescriptorSetLayout{}, core1_0.VKErrorUnknown, err
	}

	var descriptorSetLayoutHandle driver.VkDescriptorSetLayout

	res, err := v.Driver.VkCreateDescriptorSetLayout(device.Handle(), (*driver.VkDescriptorSetLayoutCreateInfo)(createInfo), allocationCallbacks.Handle(), &descriptorSetLayoutHandle)
	if err != nil {
		return types.DescriptorSetLayout{}, res, err
	}

	descriptorSetLayout := types.InternalDescriptorSetLayout(device.Handle(), descriptorSetLayoutHandle, device.APIVersion())

	return descriptorSetLayout, res, nil
}

func (v *Vulkan) CreateDescriptorPool(device types.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.DescriptorPoolCreateInfo) (types.DescriptorPool, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.DescriptorPool{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.DescriptorPool{}, core1_0.VKErrorUnknown, err
	}

	var descriptorPoolHandle driver.VkDescriptorPool

	res, err := v.Driver.VkCreateDescriptorPool(device.Handle(), (*driver.VkDescriptorPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &descriptorPoolHandle)
	if err != nil {
		return types.DescriptorPool{}, res, err
	}

	descriptorPool := types.InternalDescriptorPool(device.Handle(), descriptorPoolHandle, device.APIVersion())

	return descriptorPool, res, nil
}

func (v *Vulkan) CreateCommandPool(device types.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.CommandPoolCreateInfo) (types.CommandPool, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.CommandPool{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.CommandPool{}, core1_0.VKErrorUnknown, err
	}

	var cmdPoolHandle driver.VkCommandPool
	res, err := v.Driver.VkCreateCommandPool(device.Handle(), (*driver.VkCommandPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &cmdPoolHandle)
	if err != nil {
		return types.CommandPool{}, res, err
	}

	commandPool := types.InternalCommandPool(device.Handle(), cmdPoolHandle, device.APIVersion())

	return commandPool, res, nil
}

func (v *Vulkan) CreateEvent(device types.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.EventCreateInfo) (types.Event, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.Event{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.Event{}, core1_0.VKErrorUnknown, err
	}

	var eventHandle driver.VkEvent
	res, err := v.Driver.VkCreateEvent(device.Handle(), (*driver.VkEventCreateInfo)(createInfo), allocationCallbacks.Handle(), &eventHandle)
	if err != nil {
		return types.Event{}, res, err
	}

	event := types.InternalEvent(device.Handle(), eventHandle, device.APIVersion())

	return event, res, nil
}

func (v *Vulkan) CreateFramebuffer(device types.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.FramebufferCreateInfo) (types.Framebuffer, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.Framebuffer{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.Framebuffer{}, core1_0.VKErrorUnknown, err
	}

	var framebufferHandle driver.VkFramebuffer

	res, err := v.Driver.VkCreateFramebuffer(device.Handle(), (*driver.VkFramebufferCreateInfo)(createInfo), allocationCallbacks.Handle(), &framebufferHandle)
	if err != nil {
		return types.Framebuffer{}, res, err
	}

	framebuffer := types.InternalFramebuffer(device.Handle(), framebufferHandle, device.APIVersion())

	return framebuffer, res, nil
}

func (v *Vulkan) CreateGraphicsPipelines(device types.Device, pipelineCache types.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o ...core1_0.GraphicsPipelineCreateInfo) ([]types.Pipeline, common.VkResult, error) {
	if device.Handle() == 0 {
		return nil, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	pipelineCount := len(o)

	pipelineCreateInfosPtr, err := common.AllocOptionSlice[C.VkGraphicsPipelineCreateInfo, core1_0.GraphicsPipelineCreateInfo](arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	pipelinePtr := (*driver.VkPipeline)(arena.Malloc(pipelineCount * int(unsafe.Sizeof([1]driver.VkPipeline{}))))

	var pipelineCacheHandle driver.VkPipelineCache
	if pipelineCache.Handle() != 0 {
		pipelineCacheHandle = pipelineCache.Handle()
	}

	res, err := v.Driver.VkCreateGraphicsPipelines(device.Handle(), pipelineCacheHandle, driver.Uint32(pipelineCount), (*driver.VkGraphicsPipelineCreateInfo)(unsafe.Pointer(pipelineCreateInfosPtr)), allocationCallbacks.Handle(), pipelinePtr)
	if err != nil {
		return nil, res, err
	}

	var output []types.Pipeline
	pipelineSlice := ([]driver.VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))

	for i := 0; i < pipelineCount; i++ {
		pipeline := types.InternalPipeline(device.Handle(), pipelineSlice[i], device.APIVersion())
		output = append(output, pipeline)
	}

	return output, res, nil
}

func (v *Vulkan) CreateComputePipelines(device types.Device, pipelineCache types.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o ...core1_0.ComputePipelineCreateInfo) ([]types.Pipeline, common.VkResult, error) {
	if device.Handle() == 0 {
		return nil, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	pipelineCount := len(o)

	pipelineCreateInfosPtr, err := common.AllocOptionSlice[C.VkComputePipelineCreateInfo, core1_0.ComputePipelineCreateInfo](arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	pipelinePtr := (*driver.VkPipeline)(arena.Malloc(pipelineCount * int(unsafe.Sizeof([1]driver.VkPipeline{}))))

	var pipelineCacheHandle driver.VkPipelineCache
	if pipelineCache.Handle() != 0 {
		pipelineCacheHandle = pipelineCache.Handle()
	}

	res, err := v.Driver.VkCreateComputePipelines(device.Handle(), pipelineCacheHandle, driver.Uint32(pipelineCount), (*driver.VkComputePipelineCreateInfo)(unsafe.Pointer(pipelineCreateInfosPtr)), allocationCallbacks.Handle(), pipelinePtr)
	if err != nil {
		return nil, res, err
	}

	var output []types.Pipeline
	pipelineSlice := ([]driver.VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))

	for i := 0; i < pipelineCount; i++ {
		pipeline := types.InternalPipeline(device.Handle(), pipelineSlice[i], device.APIVersion())

		output = append(output, pipeline)
	}

	return output, res, nil
}

func (v *Vulkan) CreateImage(device types.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.ImageCreateInfo) (types.Image, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.Image{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.Image{}, core1_0.VKErrorUnknown, err
	}

	var imageHandle driver.VkImage
	res, err := v.Driver.VkCreateImage(device.Handle(), (*driver.VkImageCreateInfo)(createInfo), allocationCallbacks.Handle(), &imageHandle)
	if err != nil {
		return types.Image{}, res, err
	}

	image := types.InternalImage(device.Handle(), imageHandle, device.APIVersion())

	return image, res, nil
}

func (v *Vulkan) CreatePipelineCache(device types.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.PipelineCacheCreateInfo) (types.PipelineCache, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.PipelineCache{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.PipelineCache{}, core1_0.VKErrorUnknown, err
	}

	var pipelineCacheHandle driver.VkPipelineCache
	res, err := v.Driver.VkCreatePipelineCache(device.Handle(), (*driver.VkPipelineCacheCreateInfo)(createInfo), allocationCallbacks.Handle(), &pipelineCacheHandle)
	if err != nil {
		return types.PipelineCache{}, res, err
	}

	pipelineCache := types.InternalPipelineCache(device.Handle(), pipelineCacheHandle, device.APIVersion())

	return pipelineCache, res, nil
}

func (v *Vulkan) CreatePipelineLayout(device types.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.PipelineLayoutCreateInfo) (types.PipelineLayout, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.PipelineLayout{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.PipelineLayout{}, core1_0.VKErrorUnknown, err
	}

	var pipelineLayoutHandle driver.VkPipelineLayout
	res, err := v.Driver.VkCreatePipelineLayout(device.Handle(), (*driver.VkPipelineLayoutCreateInfo)(createInfo), allocationCallbacks.Handle(), &pipelineLayoutHandle)
	if err != nil {
		return types.PipelineLayout{}, res, err
	}

	pipelineLayout := types.InternalPipelineLayout(device.Handle(), pipelineLayoutHandle, device.APIVersion())

	return pipelineLayout, res, nil
}

func (v *Vulkan) CreateQueryPool(device types.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.QueryPoolCreateInfo) (types.QueryPool, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.QueryPool{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.QueryPool{}, core1_0.VKErrorUnknown, err
	}

	var queryPoolHandle driver.VkQueryPool

	res, err := v.Driver.VkCreateQueryPool(device.Handle(), (*driver.VkQueryPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &queryPoolHandle)
	if err != nil {
		return types.QueryPool{}, res, err
	}

	queryPool := types.InternalQueryPool(device.Handle(), queryPoolHandle, device.APIVersion())
	return queryPool, res, nil
}

func (v *Vulkan) CreateRenderPass(device types.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.RenderPassCreateInfo) (types.RenderPass, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.RenderPass{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.RenderPass{}, core1_0.VKErrorUnknown, err
	}

	var renderPassHandle driver.VkRenderPass

	res, err := v.Driver.VkCreateRenderPass(device.Handle(), (*driver.VkRenderPassCreateInfo)(createInfo), allocationCallbacks.Handle(), &renderPassHandle)
	if err != nil {
		return types.RenderPass{}, res, err
	}

	renderPass := types.InternalRenderPass(device.Handle(), renderPassHandle, device.APIVersion())

	return renderPass, res, nil
}

func (v *Vulkan) CreateSampler(device types.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.SamplerCreateInfo) (types.Sampler, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.Sampler{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.Sampler{}, core1_0.VKErrorUnknown, err
	}

	var samplerHandle driver.VkSampler

	res, err := v.Driver.VkCreateSampler(device.Handle(), (*driver.VkSamplerCreateInfo)(createInfo), allocationCallbacks.Handle(), &samplerHandle)
	if err != nil {
		return types.Sampler{}, res, err
	}

	sampler := types.InternalSampler(device.Handle(), samplerHandle, device.APIVersion())

	return sampler, res, nil
}

func (v *Vulkan) GetQueue(device types.Device, queueFamilyIndex int, queueIndex int) types.Queue {
	if device.Handle() == 0 {
		panic("device was uninitialized")
	}

	var queueHandle driver.VkQueue

	v.Driver.VkGetDeviceQueue(device.Handle(), driver.Uint32(queueFamilyIndex), driver.Uint32(queueIndex), &queueHandle)

	queue := types.InternalQueue(device.Handle(), queueHandle, device.APIVersion())

	return queue
}

func (v *Vulkan) AllocateMemory(device types.Device, allocationCallbacks *driver.AllocationCallbacks, o core1_0.MemoryAllocateInfo) (types.DeviceMemory, common.VkResult, error) {
	if device.Handle() == 0 {
		return types.DeviceMemory{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return types.DeviceMemory{}, core1_0.VKErrorUnknown, err
	}

	var deviceMemoryHandle driver.VkDeviceMemory

	deviceDriver := v.Driver
	deviceHandle := device.Handle()

	res, err := deviceDriver.VkAllocateMemory(deviceHandle, (*driver.VkMemoryAllocateInfo)(createInfo), allocationCallbacks.Handle(), &deviceMemoryHandle)
	if err != nil {
		return types.DeviceMemory{}, res, err
	}

	deviceMemory := types.InternalDeviceMemory(device.Handle(), deviceMemoryHandle, device.APIVersion(), o.AllocationSize)

	return deviceMemory, res, nil
}

// Free a slice of command buffers which should all have the same device/driver/pool
// guaranteed to have at least one element
func (v *Vulkan) freeCommandBufferSlice(buffers []types.CommandBuffer) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	bufferCount := len(buffers)
	bufferDevice := buffers[0].DeviceHandle()
	bufferPool := buffers[0].CommandPoolHandle()

	size := bufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{}))
	bufferArrayPtr := (*driver.VkCommandBuffer)(allocator.Malloc(size))
	bufferArraySlice := ([]driver.VkCommandBuffer)(unsafe.Slice(bufferArrayPtr, bufferCount))

	for i := 0; i < bufferCount; i++ {
		bufferArraySlice[i] = driver.VkCommandBuffer(0)

		if buffers[i].Handle() != 0 {
			bufferArraySlice[i] = buffers[i].Handle()
		}
	}

	v.Driver.VkFreeCommandBuffers(bufferDevice, bufferPool, driver.Uint32(bufferCount), bufferArrayPtr)
}

func (v *Vulkan) FreeCommandBuffers(buffers []types.CommandBuffer) {
	bufferCount := len(buffers)
	if bufferCount == 0 {
		return
	}

	multimap := make(map[driver.VkCommandPool][]types.CommandBuffer)
	for _, buffer := range buffers {
		poolHandle := buffer.CommandPoolHandle()
		existingSet := multimap[poolHandle]
		multimap[poolHandle] = append(existingSet, buffer)
	}

	for _, setBuffers := range multimap {
		v.freeCommandBufferSlice(setBuffers)
	}
}

func (v *Vulkan) AllocateCommandBuffers(o core1_0.CommandBufferAllocateInfo) ([]types.CommandBuffer, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	if o.CommandPool.Handle() == 0 {
		return nil, core1_0.VKErrorUnknown, errors.New("no command pool provided to allocate from")
	}

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	device := o.CommandPool.DeviceHandle()
	version := o.CommandPool.APIVersion()

	commandBufferPtr := (*driver.VkCommandBuffer)(arena.Malloc(o.CommandBufferCount * int(unsafe.Sizeof([1]driver.VkCommandBuffer{}))))

	res, err := v.Driver.VkAllocateCommandBuffers(device, (*driver.VkCommandBufferAllocateInfo)(createInfo), commandBufferPtr)
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	commandBufferArray := ([]driver.VkCommandBuffer)(unsafe.Slice(commandBufferPtr, o.CommandBufferCount))
	var result []types.CommandBuffer

	for i := 0; i < o.CommandBufferCount; i++ {
		commandBuffer := types.InternalCommandBuffer(device, o.CommandPool.Handle(), commandBufferArray[i], version)

		result = append(result, commandBuffer)
	}

	return result, res, nil
}

func (v *Vulkan) AllocateDescriptorSets(o core1_0.DescriptorSetAllocateInfo) ([]types.DescriptorSet, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	if o.DescriptorPool.Handle() == 0 {
		return nil, core1_0.VKErrorUnknown, errors.New("no descriptor pool provided to allocate from")
	}

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	device := o.DescriptorPool.DeviceHandle()
	version := o.DescriptorPool.APIVersion()

	setCount := len(o.SetLayouts)
	descriptorSets := (*driver.VkDescriptorSet)(arena.Malloc(setCount * int(unsafe.Sizeof([1]C.VkDescriptorSet{}))))

	res, err := v.Driver.VkAllocateDescriptorSets(device, (*driver.VkDescriptorSetAllocateInfo)(createInfo), descriptorSets)
	if err != nil {
		return nil, res, err
	}

	var sets []types.DescriptorSet
	descriptorSetSlice := ([]driver.VkDescriptorSet)(unsafe.Slice(descriptorSets, setCount))

	for i := 0; i < setCount; i++ {
		descriptorSet := types.InternalDescriptorSet(device, o.DescriptorPool.Handle(), descriptorSetSlice[i], version)

		sets = append(sets, descriptorSet)
	}

	return sets, res, nil
}

// Free a slice of descriptor sets which should all have the same device/driver/pool
// guaranteed to have at least one element
func (v *Vulkan) freeDescriptorSetSlice(sets []types.DescriptorSet) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	setSize := len(sets)
	arraySize := setSize * int(unsafe.Sizeof([1]C.VkDescriptorSet{}))

	arrayPtr := (*driver.VkDescriptorSet)(arena.Malloc(arraySize))
	arraySlice := ([]driver.VkDescriptorSet)(unsafe.Slice(arrayPtr, setSize))

	for i := 0; i < setSize; i++ {
		arraySlice[i] = driver.VkDescriptorSet(0)
		if sets[i].Handle() != 0 {
			arraySlice[i] = sets[i].Handle()
		}
	}

	pool := sets[0].DescriptorPoolHandle()
	device := sets[0].DeviceHandle()

	res, err := v.Driver.VkFreeDescriptorSets(device, pool, driver.Uint32(setSize), arrayPtr)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (v *Vulkan) FreeDescriptorSets(device types.Device, sets []types.DescriptorSet) (common.VkResult, error) {
	if device.Handle() == 0 {
		return core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	poolMultimap := make(map[driver.VkDescriptorPool][]types.DescriptorSet)

	for _, set := range sets {
		poolHandle := set.DescriptorPoolHandle()
		existingSet := poolMultimap[poolHandle]
		poolMultimap[poolHandle] = append(existingSet, set)
	}

	var res common.VkResult
	var err error
	for _, set := range poolMultimap {
		res, err = v.freeDescriptorSetSlice(set)
		if err != nil {
			return res, err
		}
	}

	return res, err
}
