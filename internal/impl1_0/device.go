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
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) DestroyDevice(device core.Device, callbacks *loader.AllocationCallbacks) {
	if device.Handle() == 0 {
		panic("device was uninitialized")
	}

	v.LoaderObj.VkDestroyDevice(device.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) DeviceWaitIdle(device core.Device) (common.VkResult, error) {
	if device.Handle() == 0 {
		return core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	return v.LoaderObj.VkDeviceWaitIdle(device.Handle())
}

func (v *DeviceVulkanDriver) WaitForFences(waitForAll bool, timeout time.Duration, fences ...core.Fence) (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*loader.VkFence)(fenceUnsafePtr)

	fenceSlice := ([]loader.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		if fences[i].Handle() == 0 {
			panic(fmt.Sprintf("element %d of slice fences is uninitialized", i))
		}
		if fences[i].DeviceHandle() != v.LoaderObj.DeviceHandle() {
			panic(fmt.Sprintf("element %d of slice fences was not created by this driver's device", i))
		}
		fenceSlice[i] = fences[i].Handle()
	}

	waitAllConst := C.VK_FALSE
	if waitForAll {
		waitAllConst = C.VK_TRUE
	}

	return v.LoaderObj.VkWaitForFences(v.LoaderObj.DeviceHandle(), loader.Uint32(fenceCount), fencePtr, loader.VkBool32(waitAllConst), loader.Uint64(common.TimeoutNanoseconds(timeout)))
}

func (v *DeviceVulkanDriver) ResetFences(fences ...core.Fence) (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*loader.VkFence)(fenceUnsafePtr)
	fenceSlice := ([]loader.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		if fences[i].Handle() == 0 {
			panic(fmt.Sprintf("element %d of slice fences is uninitialized", i))
		}
		if fences[i].Handle() == 0 {
			panic(fmt.Sprintf("element %d of slice fences was not created by this driver's device", i))
		}
		fenceSlice[i] = fences[i].Handle()
	}

	return v.LoaderObj.VkResetFences(v.LoaderObj.DeviceHandle(), loader.Uint32(fenceCount), fencePtr)
}

func (v *DeviceVulkanDriver) UpdateDescriptorSets(device core.Device, writes []core1_0.WriteDescriptorSet, copies []core1_0.CopyDescriptorSet) error {
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

	v.LoaderObj.VkUpdateDescriptorSets(device.Handle(), loader.Uint32(writeCount), (*loader.VkWriteDescriptorSet)(unsafe.Pointer(writePtr)), loader.Uint32(copyCount), (*loader.VkCopyDescriptorSet)(unsafe.Pointer(copyPtr)))
	return nil
}

func (v *DeviceVulkanDriver) FlushMappedMemoryRanges(ranges ...core1_0.MappedMemoryRange) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	for i, r := range ranges {
		if r.Memory.Handle() == 0 {
			return core1_0.VKErrorUnknown, fmt.Errorf("received uninitialized DeviceMemory at element %d", i)
		}
		if v.LoaderObj.DeviceHandle() != r.Memory.DeviceHandle() {
			return core1_0.VKErrorUnknown, fmt.Errorf("received memory that was not allocated by this driver's device at element %d", i)
		}
	}

	rangeCount := len(ranges)
	createInfos, err := common.AllocOptionSlice[C.VkMappedMemoryRange, core1_0.MappedMemoryRange](arena, ranges)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return v.LoaderObj.VkFlushMappedMemoryRanges(v.LoaderObj.DeviceHandle(), loader.Uint32(rangeCount), (*loader.VkMappedMemoryRange)(unsafe.Pointer(createInfos)))
}

func (v *DeviceVulkanDriver) InvalidateMappedMemoryRanges(ranges ...core1_0.MappedMemoryRange) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	for i, r := range ranges {
		if r.Memory.Handle() == 0 {
			return core1_0.VKErrorUnknown, fmt.Errorf("received uninitialized DeviceMemory at element %d", i)
		}
		if v.LoaderObj.DeviceHandle() != r.Memory.DeviceHandle() {
			return core1_0.VKErrorUnknown, fmt.Errorf("received memory that was not allocated by this driver's device at element %d", i)
		}
	}

	rangeCount := len(ranges)
	createInfos, err := common.AllocOptionSlice[C.VkMappedMemoryRange, core1_0.MappedMemoryRange](arena, ranges)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return v.LoaderObj.VkInvalidateMappedMemoryRanges(v.LoaderObj.DeviceHandle(), loader.Uint32(rangeCount), (*loader.VkMappedMemoryRange)(unsafe.Pointer(createInfos)))
}

func (v *DeviceVulkanDriver) CreateBufferView(device core.Device, allocationCallbacks *loader.AllocationCallbacks, options core1_0.BufferViewCreateInfo) (core.BufferView, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.BufferView{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return core.BufferView{}, core1_0.VKErrorUnknown, err
	}

	var bufferViewHandle loader.VkBufferView

	res, err := v.LoaderObj.VkCreateBufferView(device.Handle(), (*loader.VkBufferViewCreateInfo)(createInfo), allocationCallbacks.Handle(), &bufferViewHandle)
	if err != nil {
		return core.BufferView{}, res, err
	}

	bufferView := core.InternalBufferView(device.Handle(), bufferViewHandle, device.APIVersion())

	return bufferView, res, nil
}

func (v *DeviceVulkanDriver) CreateShaderModule(device core.Device, allocationCallbacks *loader.AllocationCallbacks, o core1_0.ShaderModuleCreateInfo) (core.ShaderModule, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.ShaderModule{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.ShaderModule{}, core1_0.VKErrorUnknown, err
	}

	var shaderModuleHandle loader.VkShaderModule
	res, err := v.LoaderObj.VkCreateShaderModule(device.Handle(), (*loader.VkShaderModuleCreateInfo)(createInfo), allocationCallbacks.Handle(), &shaderModuleHandle)
	if err != nil {
		return core.ShaderModule{}, res, err
	}

	shaderModule := core.InternalShaderModule(device.Handle(), shaderModuleHandle, device.APIVersion())

	return shaderModule, res, nil
}

func (v *DeviceVulkanDriver) CreateImageView(device core.Device, allocationCallbacks *loader.AllocationCallbacks, o core1_0.ImageViewCreateInfo) (core.ImageView, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.ImageView{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.ImageView{}, core1_0.VKErrorUnknown, err
	}

	var imageViewHandle loader.VkImageView

	res, err := v.LoaderObj.VkCreateImageView(device.Handle(), (*loader.VkImageViewCreateInfo)(createInfo), allocationCallbacks.Handle(), &imageViewHandle)
	if err != nil {
		return core.ImageView{}, res, err
	}

	imageView := core.InternalImageView(device.Handle(), imageViewHandle, device.APIVersion())

	return imageView, res, nil
}

func (v *DeviceVulkanDriver) CreateSemaphore(device core.Device, allocationCallbacks *loader.AllocationCallbacks, o core1_0.SemaphoreCreateInfo) (core.Semaphore, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.Semaphore{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.Semaphore{}, core1_0.VKErrorUnknown, err
	}

	var semaphoreHandle loader.VkSemaphore

	res, err := v.LoaderObj.VkCreateSemaphore(device.Handle(), (*loader.VkSemaphoreCreateInfo)(createInfo), allocationCallbacks.Handle(), &semaphoreHandle)
	if err != nil {
		return core.Semaphore{}, res, err
	}

	semaphore := core.InternalSemaphore(device.Handle(), semaphoreHandle, device.APIVersion())

	return semaphore, res, nil
}

func (v *DeviceVulkanDriver) CreateFence(device core.Device, allocationCallbacks *loader.AllocationCallbacks, o core1_0.FenceCreateInfo) (core.Fence, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.Fence{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.Fence{}, core1_0.VKErrorUnknown, err
	}

	var fenceHandle loader.VkFence

	res, err := v.LoaderObj.VkCreateFence(device.Handle(), (*loader.VkFenceCreateInfo)(createInfo), allocationCallbacks.Handle(), &fenceHandle)
	if err != nil {
		return core.Fence{}, res, err
	}

	fence := core.InternalFence(device.Handle(), fenceHandle, device.APIVersion())

	return fence, res, nil
}

func (v *DeviceVulkanDriver) CreateBuffer(device core.Device, allocationCallbacks *loader.AllocationCallbacks, o core1_0.BufferCreateInfo) (core.Buffer, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.Buffer{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.Buffer{}, core1_0.VKErrorUnknown, err
	}

	var bufferHandle loader.VkBuffer

	res, err := v.LoaderObj.VkCreateBuffer(device.Handle(), (*loader.VkBufferCreateInfo)(createInfo), allocationCallbacks.Handle(), &bufferHandle)
	if err != nil {
		return core.Buffer{}, res, err
	}

	buffer := core.InternalBuffer(device.Handle(), bufferHandle, device.APIVersion())

	return buffer, res, nil
}

func (v *DeviceVulkanDriver) CreateDescriptorSetLayout(device core.Device, allocationCallbacks *loader.AllocationCallbacks, o core1_0.DescriptorSetLayoutCreateInfo) (core.DescriptorSetLayout, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.DescriptorSetLayout{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.DescriptorSetLayout{}, core1_0.VKErrorUnknown, err
	}

	var descriptorSetLayoutHandle loader.VkDescriptorSetLayout

	res, err := v.LoaderObj.VkCreateDescriptorSetLayout(device.Handle(), (*loader.VkDescriptorSetLayoutCreateInfo)(createInfo), allocationCallbacks.Handle(), &descriptorSetLayoutHandle)
	if err != nil {
		return core.DescriptorSetLayout{}, res, err
	}

	descriptorSetLayout := core.InternalDescriptorSetLayout(device.Handle(), descriptorSetLayoutHandle, device.APIVersion())

	return descriptorSetLayout, res, nil
}

func (v *DeviceVulkanDriver) CreateDescriptorPool(device core.Device, allocationCallbacks *loader.AllocationCallbacks, o core1_0.DescriptorPoolCreateInfo) (core.DescriptorPool, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.DescriptorPool{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.DescriptorPool{}, core1_0.VKErrorUnknown, err
	}

	var descriptorPoolHandle loader.VkDescriptorPool

	res, err := v.LoaderObj.VkCreateDescriptorPool(device.Handle(), (*loader.VkDescriptorPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &descriptorPoolHandle)
	if err != nil {
		return core.DescriptorPool{}, res, err
	}

	descriptorPool := core.InternalDescriptorPool(device.Handle(), descriptorPoolHandle, device.APIVersion())

	return descriptorPool, res, nil
}

func (v *DeviceVulkanDriver) CreateCommandPool(device core.Device, allocationCallbacks *loader.AllocationCallbacks, o core1_0.CommandPoolCreateInfo) (core.CommandPool, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.CommandPool{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.CommandPool{}, core1_0.VKErrorUnknown, err
	}

	var cmdPoolHandle loader.VkCommandPool
	res, err := v.LoaderObj.VkCreateCommandPool(device.Handle(), (*loader.VkCommandPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &cmdPoolHandle)
	if err != nil {
		return core.CommandPool{}, res, err
	}

	commandPool := core.InternalCommandPool(device.Handle(), cmdPoolHandle, device.APIVersion())

	return commandPool, res, nil
}

func (v *DeviceVulkanDriver) CreateEvent(device core.Device, allocationCallbacks *loader.AllocationCallbacks, o core1_0.EventCreateInfo) (core.Event, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.Event{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.Event{}, core1_0.VKErrorUnknown, err
	}

	var eventHandle loader.VkEvent
	res, err := v.LoaderObj.VkCreateEvent(device.Handle(), (*loader.VkEventCreateInfo)(createInfo), allocationCallbacks.Handle(), &eventHandle)
	if err != nil {
		return core.Event{}, res, err
	}

	event := core.InternalEvent(device.Handle(), eventHandle, device.APIVersion())

	return event, res, nil
}

func (v *DeviceVulkanDriver) CreateFramebuffer(device core.Device, allocationCallbacks *loader.AllocationCallbacks, o core1_0.FramebufferCreateInfo) (core.Framebuffer, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.Framebuffer{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.Framebuffer{}, core1_0.VKErrorUnknown, err
	}

	var framebufferHandle loader.VkFramebuffer

	res, err := v.LoaderObj.VkCreateFramebuffer(device.Handle(), (*loader.VkFramebufferCreateInfo)(createInfo), allocationCallbacks.Handle(), &framebufferHandle)
	if err != nil {
		return core.Framebuffer{}, res, err
	}

	framebuffer := core.InternalFramebuffer(device.Handle(), framebufferHandle, device.APIVersion())

	return framebuffer, res, nil
}

func (v *DeviceVulkanDriver) CreateGraphicsPipelines(device core.Device, pipelineCache *core.PipelineCache, allocationCallbacks *loader.AllocationCallbacks, o ...core1_0.GraphicsPipelineCreateInfo) ([]core.Pipeline, common.VkResult, error) {
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

	pipelinePtr := (*loader.VkPipeline)(arena.Malloc(pipelineCount * int(unsafe.Sizeof([1]loader.VkPipeline{}))))

	var pipelineCacheHandle loader.VkPipelineCache
	if pipelineCache != nil {
		pipelineCacheHandle = pipelineCache.Handle()
	}

	res, err := v.LoaderObj.VkCreateGraphicsPipelines(device.Handle(), pipelineCacheHandle, loader.Uint32(pipelineCount), (*loader.VkGraphicsPipelineCreateInfo)(unsafe.Pointer(pipelineCreateInfosPtr)), allocationCallbacks.Handle(), pipelinePtr)
	if err != nil {
		return nil, res, err
	}

	var output []core.Pipeline
	pipelineSlice := ([]loader.VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))

	for i := 0; i < pipelineCount; i++ {
		pipeline := core.InternalPipeline(device.Handle(), pipelineSlice[i], device.APIVersion())
		output = append(output, pipeline)
	}

	return output, res, nil
}

func (v *DeviceVulkanDriver) CreateComputePipelines(device core.Device, pipelineCache *core.PipelineCache, allocationCallbacks *loader.AllocationCallbacks, o ...core1_0.ComputePipelineCreateInfo) ([]core.Pipeline, common.VkResult, error) {
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

	pipelinePtr := (*loader.VkPipeline)(arena.Malloc(pipelineCount * int(unsafe.Sizeof([1]loader.VkPipeline{}))))

	var pipelineCacheHandle loader.VkPipelineCache
	if pipelineCache != nil {
		pipelineCacheHandle = pipelineCache.Handle()
	}

	res, err := v.LoaderObj.VkCreateComputePipelines(device.Handle(), pipelineCacheHandle, loader.Uint32(pipelineCount), (*loader.VkComputePipelineCreateInfo)(unsafe.Pointer(pipelineCreateInfosPtr)), allocationCallbacks.Handle(), pipelinePtr)
	if err != nil {
		return nil, res, err
	}

	var output []core.Pipeline
	pipelineSlice := ([]loader.VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))

	for i := 0; i < pipelineCount; i++ {
		pipeline := core.InternalPipeline(device.Handle(), pipelineSlice[i], device.APIVersion())

		output = append(output, pipeline)
	}

	return output, res, nil
}

func (v *DeviceVulkanDriver) CreateImage(device core.Device, allocationCallbacks *loader.AllocationCallbacks, o core1_0.ImageCreateInfo) (core.Image, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.Image{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.Image{}, core1_0.VKErrorUnknown, err
	}

	var imageHandle loader.VkImage
	res, err := v.LoaderObj.VkCreateImage(device.Handle(), (*loader.VkImageCreateInfo)(createInfo), allocationCallbacks.Handle(), &imageHandle)
	if err != nil {
		return core.Image{}, res, err
	}

	image := core.InternalImage(device.Handle(), imageHandle, device.APIVersion())

	return image, res, nil
}

func (v *DeviceVulkanDriver) CreatePipelineCache(device core.Device, allocationCallbacks *loader.AllocationCallbacks, o core1_0.PipelineCacheCreateInfo) (core.PipelineCache, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.PipelineCache{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.PipelineCache{}, core1_0.VKErrorUnknown, err
	}

	var pipelineCacheHandle loader.VkPipelineCache
	res, err := v.LoaderObj.VkCreatePipelineCache(device.Handle(), (*loader.VkPipelineCacheCreateInfo)(createInfo), allocationCallbacks.Handle(), &pipelineCacheHandle)
	if err != nil {
		return core.PipelineCache{}, res, err
	}

	pipelineCache := core.InternalPipelineCache(device.Handle(), pipelineCacheHandle, device.APIVersion())

	return pipelineCache, res, nil
}

func (v *DeviceVulkanDriver) CreatePipelineLayout(device core.Device, allocationCallbacks *loader.AllocationCallbacks, o core1_0.PipelineLayoutCreateInfo) (core.PipelineLayout, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.PipelineLayout{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.PipelineLayout{}, core1_0.VKErrorUnknown, err
	}

	var pipelineLayoutHandle loader.VkPipelineLayout
	res, err := v.LoaderObj.VkCreatePipelineLayout(device.Handle(), (*loader.VkPipelineLayoutCreateInfo)(createInfo), allocationCallbacks.Handle(), &pipelineLayoutHandle)
	if err != nil {
		return core.PipelineLayout{}, res, err
	}

	pipelineLayout := core.InternalPipelineLayout(device.Handle(), pipelineLayoutHandle, device.APIVersion())

	return pipelineLayout, res, nil
}

func (v *DeviceVulkanDriver) CreateQueryPool(device core.Device, allocationCallbacks *loader.AllocationCallbacks, o core1_0.QueryPoolCreateInfo) (core.QueryPool, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.QueryPool{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.QueryPool{}, core1_0.VKErrorUnknown, err
	}

	var queryPoolHandle loader.VkQueryPool

	res, err := v.LoaderObj.VkCreateQueryPool(device.Handle(), (*loader.VkQueryPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &queryPoolHandle)
	if err != nil {
		return core.QueryPool{}, res, err
	}

	queryPool := core.InternalQueryPool(device.Handle(), queryPoolHandle, device.APIVersion())
	return queryPool, res, nil
}

func (v *DeviceVulkanDriver) CreateRenderPass(device core.Device, allocationCallbacks *loader.AllocationCallbacks, o core1_0.RenderPassCreateInfo) (core.RenderPass, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.RenderPass{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.RenderPass{}, core1_0.VKErrorUnknown, err
	}

	var renderPassHandle loader.VkRenderPass

	res, err := v.LoaderObj.VkCreateRenderPass(device.Handle(), (*loader.VkRenderPassCreateInfo)(createInfo), allocationCallbacks.Handle(), &renderPassHandle)
	if err != nil {
		return core.RenderPass{}, res, err
	}

	renderPass := core.InternalRenderPass(device.Handle(), renderPassHandle, device.APIVersion())

	return renderPass, res, nil
}

func (v *DeviceVulkanDriver) CreateSampler(device core.Device, allocationCallbacks *loader.AllocationCallbacks, o core1_0.SamplerCreateInfo) (core.Sampler, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.Sampler{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.Sampler{}, core1_0.VKErrorUnknown, err
	}

	var samplerHandle loader.VkSampler

	res, err := v.LoaderObj.VkCreateSampler(device.Handle(), (*loader.VkSamplerCreateInfo)(createInfo), allocationCallbacks.Handle(), &samplerHandle)
	if err != nil {
		return core.Sampler{}, res, err
	}

	sampler := core.InternalSampler(device.Handle(), samplerHandle, device.APIVersion())

	return sampler, res, nil
}

func (v *DeviceVulkanDriver) GetQueue(device core.Device, queueFamilyIndex int, queueIndex int) core.Queue {
	if device.Handle() == 0 {
		panic("device was uninitialized")
	}

	var queueHandle loader.VkQueue

	v.LoaderObj.VkGetDeviceQueue(device.Handle(), loader.Uint32(queueFamilyIndex), loader.Uint32(queueIndex), &queueHandle)

	queue := core.InternalQueue(device.Handle(), queueHandle, device.APIVersion())

	return queue
}

func (v *DeviceVulkanDriver) AllocateMemory(device core.Device, allocationCallbacks *loader.AllocationCallbacks, o core1_0.MemoryAllocateInfo) (core.DeviceMemory, common.VkResult, error) {
	if device.Handle() == 0 {
		return core.DeviceMemory{}, core1_0.VKErrorUnknown, errors.New("device was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.DeviceMemory{}, core1_0.VKErrorUnknown, err
	}

	var deviceMemoryHandle loader.VkDeviceMemory

	deviceDriver := v.LoaderObj
	deviceHandle := device.Handle()

	res, err := deviceDriver.VkAllocateMemory(deviceHandle, (*loader.VkMemoryAllocateInfo)(createInfo), allocationCallbacks.Handle(), &deviceMemoryHandle)
	if err != nil {
		return core.DeviceMemory{}, res, err
	}

	deviceMemory := core.InternalDeviceMemory(device.Handle(), deviceMemoryHandle, device.APIVersion(), o.AllocationSize)

	return deviceMemory, res, nil
}

// Free a slice of command buffers which should all have the same device/loader/pool
// guaranteed to have at least one element
func (v *DeviceVulkanDriver) freeCommandBufferSlice(buffers []core.CommandBuffer) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	bufferCount := len(buffers)
	bufferDevice := buffers[0].DeviceHandle()
	bufferPool := buffers[0].CommandPoolHandle()

	size := bufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{}))
	bufferArrayPtr := (*loader.VkCommandBuffer)(allocator.Malloc(size))
	bufferArraySlice := ([]loader.VkCommandBuffer)(unsafe.Slice(bufferArrayPtr, bufferCount))

	for i := 0; i < bufferCount; i++ {
		bufferArraySlice[i] = loader.VkCommandBuffer(0)

		if buffers[i].Handle() != 0 {
			bufferArraySlice[i] = buffers[i].Handle()
		}
	}

	v.LoaderObj.VkFreeCommandBuffers(bufferDevice, bufferPool, loader.Uint32(bufferCount), bufferArrayPtr)
}

func (v *DeviceVulkanDriver) FreeCommandBuffers(buffers ...core.CommandBuffer) {
	bufferCount := len(buffers)
	if bufferCount == 0 {
		return
	}

	multimap := make(map[loader.VkCommandPool][]core.CommandBuffer)
	for _, buffer := range buffers {
		poolHandle := buffer.CommandPoolHandle()
		existingSet := multimap[poolHandle]
		multimap[poolHandle] = append(existingSet, buffer)
	}

	for _, setBuffers := range multimap {
		v.freeCommandBufferSlice(setBuffers)
	}
}

func (v *DeviceVulkanDriver) AllocateCommandBuffers(o core1_0.CommandBufferAllocateInfo) ([]core.CommandBuffer, common.VkResult, error) {
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

	commandBufferPtr := (*loader.VkCommandBuffer)(arena.Malloc(o.CommandBufferCount * int(unsafe.Sizeof([1]loader.VkCommandBuffer{}))))

	res, err := v.LoaderObj.VkAllocateCommandBuffers(device, (*loader.VkCommandBufferAllocateInfo)(createInfo), commandBufferPtr)
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	commandBufferArray := ([]loader.VkCommandBuffer)(unsafe.Slice(commandBufferPtr, o.CommandBufferCount))
	var result []core.CommandBuffer

	for i := 0; i < o.CommandBufferCount; i++ {
		commandBuffer := core.InternalCommandBuffer(device, o.CommandPool.Handle(), commandBufferArray[i], version)

		result = append(result, commandBuffer)
	}

	return result, res, nil
}

func (v *DeviceVulkanDriver) AllocateDescriptorSets(o core1_0.DescriptorSetAllocateInfo) ([]core.DescriptorSet, common.VkResult, error) {
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
	descriptorSets := (*loader.VkDescriptorSet)(arena.Malloc(setCount * int(unsafe.Sizeof([1]C.VkDescriptorSet{}))))

	res, err := v.LoaderObj.VkAllocateDescriptorSets(device, (*loader.VkDescriptorSetAllocateInfo)(createInfo), descriptorSets)
	if err != nil {
		return nil, res, err
	}

	var sets []core.DescriptorSet
	descriptorSetSlice := ([]loader.VkDescriptorSet)(unsafe.Slice(descriptorSets, setCount))

	for i := 0; i < setCount; i++ {
		descriptorSet := core.InternalDescriptorSet(device, o.DescriptorPool.Handle(), descriptorSetSlice[i], version)

		sets = append(sets, descriptorSet)
	}

	return sets, res, nil
}

// Free a slice of descriptor sets which should all have the same device/loader/pool
// guaranteed to have at least one element
func (v *DeviceVulkanDriver) freeDescriptorSetSlice(sets []core.DescriptorSet) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	setSize := len(sets)
	arraySize := setSize * int(unsafe.Sizeof([1]C.VkDescriptorSet{}))

	arrayPtr := (*loader.VkDescriptorSet)(arena.Malloc(arraySize))
	arraySlice := ([]loader.VkDescriptorSet)(unsafe.Slice(arrayPtr, setSize))

	for i := 0; i < setSize; i++ {
		arraySlice[i] = loader.VkDescriptorSet(0)
		if sets[i].Handle() != 0 {
			arraySlice[i] = sets[i].Handle()
		}
	}

	pool := sets[0].DescriptorPoolHandle()
	device := sets[0].DeviceHandle()

	res, err := v.LoaderObj.VkFreeDescriptorSets(device, pool, loader.Uint32(setSize), arrayPtr)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (v *DeviceVulkanDriver) FreeDescriptorSets(sets ...core.DescriptorSet) (common.VkResult, error) {
	poolMultimap := make(map[loader.VkDescriptorPool][]core.DescriptorSet)

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
