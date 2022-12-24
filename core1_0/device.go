package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"fmt"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/driver"
	"time"
	"unsafe"
)

// VulkanDevice is an implementation of the Device interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDevice struct {
	deviceDriver driver.Driver
	deviceHandle driver.VkDevice

	maximumAPIVersion      common.APIVersion
	activeDeviceExtensions map[string]struct{}
}

func (d *VulkanDevice) Driver() driver.Driver {
	return d.deviceDriver
}

func (d *VulkanDevice) Handle() driver.VkDevice {
	return d.deviceHandle
}

func (d *VulkanDevice) APIVersion() common.APIVersion {
	return d.maximumAPIVersion
}

func (d *VulkanDevice) IsDeviceExtensionActive(extensionName string) bool {
	_, active := d.activeDeviceExtensions[extensionName]
	return active
}

func (d *VulkanDevice) Destroy(callbacks *driver.AllocationCallbacks) {
	d.deviceDriver.VkDestroyDevice(d.deviceHandle, callbacks.Handle())
	d.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(d.deviceHandle))
}

func (d *VulkanDevice) WaitIdle() (common.VkResult, error) {
	return d.deviceDriver.VkDeviceWaitIdle(d.deviceHandle)
}

func (d *VulkanDevice) WaitForFences(waitForAll bool, timeout time.Duration, fences []Fence) (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver.VkFence)(fenceUnsafePtr)

	fenceSlice := ([]driver.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		if fences[i] == nil {
			panic(fmt.Sprintf("element %d of slice fences is nil", i))
		}
		fenceSlice[i] = fences[i].Handle()
	}

	waitAllConst := C.VK_FALSE
	if waitForAll {
		waitAllConst = C.VK_TRUE
	}

	return d.deviceDriver.VkWaitForFences(d.deviceHandle, driver.Uint32(fenceCount), fencePtr, driver.VkBool32(waitAllConst), driver.Uint64(common.TimeoutNanoseconds(timeout)))
}

func (d *VulkanDevice) ResetFences(fences []Fence) (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver.VkFence)(fenceUnsafePtr)
	fenceSlice := ([]driver.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		if fences[i] == nil {
			panic(fmt.Sprintf("element %d of slice fences is nil", i))
		}
		fenceSlice[i] = fences[i].Handle()
	}

	return d.deviceDriver.VkResetFences(d.deviceHandle, driver.Uint32(fenceCount), fencePtr)
}

func (d *VulkanDevice) UpdateDescriptorSets(writes []WriteDescriptorSet, copies []CopyDescriptorSet) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	writeCount := len(writes)
	copyCount := len(copies)

	var err error
	var writePtr *C.VkWriteDescriptorSet
	var copyPtr *C.VkCopyDescriptorSet

	if writeCount > 0 {
		writePtr, err = common.AllocOptionSlice[C.VkWriteDescriptorSet, WriteDescriptorSet](arena, writes)
		if err != nil {
			return err
		}
	}

	if copyCount > 0 {
		copyPtr, err = common.AllocOptionSlice[C.VkCopyDescriptorSet, CopyDescriptorSet](arena, copies)
		if err != nil {
			return err
		}
	}

	d.deviceDriver.VkUpdateDescriptorSets(d.deviceHandle, driver.Uint32(writeCount), (*driver.VkWriteDescriptorSet)(unsafe.Pointer(writePtr)), driver.Uint32(copyCount), (*driver.VkCopyDescriptorSet)(unsafe.Pointer(copyPtr)))
	return nil
}

func (d *VulkanDevice) FlushMappedMemoryRanges(ranges []MappedMemoryRange) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	rangeCount := len(ranges)
	createInfos, err := common.AllocOptionSlice[C.VkMappedMemoryRange, MappedMemoryRange](arena, ranges)
	if err != nil {
		return VKErrorUnknown, err
	}

	return d.deviceDriver.VkFlushMappedMemoryRanges(d.deviceHandle, driver.Uint32(rangeCount), (*driver.VkMappedMemoryRange)(unsafe.Pointer(createInfos)))
}

func (d *VulkanDevice) InvalidateMappedMemoryRanges(ranges []MappedMemoryRange) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	rangeCount := len(ranges)
	createInfos, err := common.AllocOptionSlice[C.VkMappedMemoryRange, MappedMemoryRange](arena, ranges)
	if err != nil {
		return VKErrorUnknown, err
	}

	return d.deviceDriver.VkInvalidateMappedMemoryRanges(d.deviceHandle, driver.Uint32(rangeCount), (*driver.VkMappedMemoryRange)(unsafe.Pointer(createInfos)))
}

func (d *VulkanDevice) CreateBufferView(allocationCallbacks *driver.AllocationCallbacks, options BufferViewCreateInfo) (BufferView, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var bufferViewHandle driver.VkBufferView

	res, err := d.deviceDriver.VkCreateBufferView(d.deviceHandle, (*driver.VkBufferViewCreateInfo)(createInfo), allocationCallbacks.Handle(), &bufferViewHandle)
	if err != nil {
		return nil, res, err
	}

	bufferView := createBufferViewObject(d.deviceDriver, d.deviceHandle, bufferViewHandle, d.maximumAPIVersion)

	return bufferView, res, nil
}

func (d *VulkanDevice) CreateShaderModule(allocationCallbacks *driver.AllocationCallbacks, o ShaderModuleCreateInfo) (ShaderModule, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var shaderModuleHandle driver.VkShaderModule
	res, err := d.deviceDriver.VkCreateShaderModule(d.deviceHandle, (*driver.VkShaderModuleCreateInfo)(createInfo), allocationCallbacks.Handle(), &shaderModuleHandle)
	if err != nil {
		return nil, res, err
	}

	shaderModule := createShaderModuleObject(d.deviceDriver, d.deviceHandle, shaderModuleHandle, d.maximumAPIVersion)

	return shaderModule, res, nil
}

func (d *VulkanDevice) CreateImageView(allocationCallbacks *driver.AllocationCallbacks, o ImageViewCreateInfo) (ImageView, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var imageViewHandle driver.VkImageView

	res, err := d.deviceDriver.VkCreateImageView(d.deviceHandle, (*driver.VkImageViewCreateInfo)(createInfo), allocationCallbacks.Handle(), &imageViewHandle)
	if err != nil {
		return nil, res, err
	}

	imageView := createImageViewObject(d.deviceDriver, d.deviceHandle, imageViewHandle, d.maximumAPIVersion)

	return imageView, res, nil
}

func (d *VulkanDevice) CreateSemaphore(allocationCallbacks *driver.AllocationCallbacks, o SemaphoreCreateInfo) (Semaphore, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var semaphoreHandle driver.VkSemaphore

	res, err := d.deviceDriver.VkCreateSemaphore(d.deviceHandle, (*driver.VkSemaphoreCreateInfo)(createInfo), allocationCallbacks.Handle(), &semaphoreHandle)
	if err != nil {
		return nil, res, err
	}

	semaphore := createSemaphoreObject(d.deviceDriver, d.deviceHandle, semaphoreHandle, d.maximumAPIVersion)

	return semaphore, res, nil
}

func (d *VulkanDevice) CreateFence(allocationCallbacks *driver.AllocationCallbacks, o FenceCreateInfo) (Fence, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var fenceHandle driver.VkFence

	res, err := d.deviceDriver.VkCreateFence(d.deviceHandle, (*driver.VkFenceCreateInfo)(createInfo), allocationCallbacks.Handle(), &fenceHandle)
	if err != nil {
		return nil, res, err
	}

	fence := createFenceObject(d.deviceDriver, d.deviceHandle, fenceHandle, d.maximumAPIVersion)

	return fence, res, nil
}

func (d *VulkanDevice) CreateBuffer(allocationCallbacks *driver.AllocationCallbacks, o BufferCreateInfo) (Buffer, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var bufferHandle driver.VkBuffer

	res, err := d.deviceDriver.VkCreateBuffer(d.deviceHandle, (*driver.VkBufferCreateInfo)(createInfo), allocationCallbacks.Handle(), &bufferHandle)
	if err != nil {
		return nil, res, err
	}

	buffer := createBufferObject(d.deviceDriver, d.deviceHandle, bufferHandle, d.maximumAPIVersion)

	return buffer, res, nil
}

func (d *VulkanDevice) CreateDescriptorSetLayout(allocationCallbacks *driver.AllocationCallbacks, o DescriptorSetLayoutCreateInfo) (DescriptorSetLayout, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var descriptorSetLayoutHandle driver.VkDescriptorSetLayout

	res, err := d.deviceDriver.VkCreateDescriptorSetLayout(d.deviceHandle, (*driver.VkDescriptorSetLayoutCreateInfo)(createInfo), allocationCallbacks.Handle(), &descriptorSetLayoutHandle)
	if err != nil {
		return nil, res, err
	}

	descriptorSetLayout := createDescriptorSetLayoutObject(d.deviceDriver, d.deviceHandle, descriptorSetLayoutHandle, d.maximumAPIVersion)

	return descriptorSetLayout, res, nil
}

func (d *VulkanDevice) CreateDescriptorPool(allocationCallbacks *driver.AllocationCallbacks, o DescriptorPoolCreateInfo) (DescriptorPool, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var descriptorPoolHandle driver.VkDescriptorPool

	res, err := d.deviceDriver.VkCreateDescriptorPool(d.deviceHandle, (*driver.VkDescriptorPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &descriptorPoolHandle)
	if err != nil {
		return nil, res, err
	}

	descriptorPool := createDescriptorPoolObject(d.deviceDriver, d.deviceHandle, descriptorPoolHandle, d.maximumAPIVersion)

	return descriptorPool, res, nil
}

func (d *VulkanDevice) CreateCommandPool(allocationCallbacks *driver.AllocationCallbacks, o CommandPoolCreateInfo) (CommandPool, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var cmdPoolHandle driver.VkCommandPool
	res, err := d.deviceDriver.VkCreateCommandPool(d.deviceHandle, (*driver.VkCommandPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &cmdPoolHandle)
	if err != nil {
		return nil, res, err
	}

	commandPool := createCommandPoolObject(d.deviceDriver, d.deviceHandle, cmdPoolHandle, d.maximumAPIVersion)

	return commandPool, res, nil
}

func (d *VulkanDevice) CreateEvent(allocationCallbacks *driver.AllocationCallbacks, o EventCreateInfo) (Event, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var eventHandle driver.VkEvent
	res, err := d.deviceDriver.VkCreateEvent(d.deviceHandle, (*driver.VkEventCreateInfo)(createInfo), allocationCallbacks.Handle(), &eventHandle)
	if err != nil {
		return nil, res, err
	}

	event := createEventObject(d.deviceDriver, d.deviceHandle, eventHandle, d.maximumAPIVersion)

	return event, res, nil
}

func (d *VulkanDevice) CreateFramebuffer(allocationCallbacks *driver.AllocationCallbacks, o FramebufferCreateInfo) (Framebuffer, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var framebufferHandle driver.VkFramebuffer

	res, err := d.deviceDriver.VkCreateFramebuffer(d.deviceHandle, (*driver.VkFramebufferCreateInfo)(createInfo), allocationCallbacks.Handle(), &framebufferHandle)
	if err != nil {
		return nil, res, err
	}

	framebuffer := createFramebufferObject(d.deviceDriver, d.deviceHandle, framebufferHandle, d.maximumAPIVersion)

	return framebuffer, res, nil
}

func (d *VulkanDevice) CreateGraphicsPipelines(pipelineCache PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []GraphicsPipelineCreateInfo) ([]Pipeline, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	pipelineCount := len(o)

	pipelineCreateInfosPtr, err := common.AllocOptionSlice[C.VkGraphicsPipelineCreateInfo, GraphicsPipelineCreateInfo](arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	pipelinePtr := (*driver.VkPipeline)(arena.Malloc(pipelineCount * int(unsafe.Sizeof([1]driver.VkPipeline{}))))

	var pipelineCacheHandle driver.VkPipelineCache
	if pipelineCache != nil {
		pipelineCacheHandle = pipelineCache.Handle()
	}

	res, err := d.deviceDriver.VkCreateGraphicsPipelines(d.deviceHandle, pipelineCacheHandle, driver.Uint32(pipelineCount), (*driver.VkGraphicsPipelineCreateInfo)(unsafe.Pointer(pipelineCreateInfosPtr)), allocationCallbacks.Handle(), pipelinePtr)
	if err != nil {
		return nil, res, err
	}

	var output []Pipeline
	pipelineSlice := ([]driver.VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))

	for i := 0; i < pipelineCount; i++ {
		pipeline := createPipelineObject(d.deviceDriver, d.deviceHandle, pipelineSlice[i], d.maximumAPIVersion)
		output = append(output, pipeline)
	}

	return output, res, nil
}

func (d *VulkanDevice) CreateComputePipelines(pipelineCache PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []ComputePipelineCreateInfo) ([]Pipeline, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	pipelineCount := len(o)

	pipelineCreateInfosPtr, err := common.AllocOptionSlice[C.VkComputePipelineCreateInfo, ComputePipelineCreateInfo](arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	pipelinePtr := (*driver.VkPipeline)(arena.Malloc(pipelineCount * int(unsafe.Sizeof([1]driver.VkPipeline{}))))

	var pipelineCacheHandle driver.VkPipelineCache
	if pipelineCache != nil {
		pipelineCacheHandle = pipelineCache.Handle()
	}

	res, err := d.deviceDriver.VkCreateComputePipelines(d.deviceHandle, pipelineCacheHandle, driver.Uint32(pipelineCount), (*driver.VkComputePipelineCreateInfo)(unsafe.Pointer(pipelineCreateInfosPtr)), allocationCallbacks.Handle(), pipelinePtr)
	if err != nil {
		return nil, res, err
	}

	var output []Pipeline
	pipelineSlice := ([]driver.VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))

	for i := 0; i < pipelineCount; i++ {
		pipeline := createPipelineObject(d.deviceDriver, d.deviceHandle, pipelineSlice[i], d.maximumAPIVersion)

		output = append(output, pipeline)
	}

	return output, res, nil
}

func (d *VulkanDevice) CreateImage(allocationCallbacks *driver.AllocationCallbacks, o ImageCreateInfo) (Image, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var imageHandle driver.VkImage
	res, err := d.deviceDriver.VkCreateImage(d.deviceHandle, (*driver.VkImageCreateInfo)(createInfo), allocationCallbacks.Handle(), &imageHandle)
	if err != nil {
		return nil, res, err
	}

	image := createImageObject(d.deviceDriver, d.deviceHandle, imageHandle, d.maximumAPIVersion)

	return image, res, nil
}

func (d *VulkanDevice) CreatePipelineCache(allocationCallbacks *driver.AllocationCallbacks, o PipelineCacheCreateInfo) (PipelineCache, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var pipelineCacheHandle driver.VkPipelineCache
	res, err := d.deviceDriver.VkCreatePipelineCache(d.deviceHandle, (*driver.VkPipelineCacheCreateInfo)(createInfo), allocationCallbacks.Handle(), &pipelineCacheHandle)
	if err != nil {
		return nil, res, err
	}

	pipelineCache := createPipelineCacheObject(d.deviceDriver, d.deviceHandle, pipelineCacheHandle, d.maximumAPIVersion)

	return pipelineCache, res, nil
}

func (d *VulkanDevice) CreatePipelineLayout(allocationCallbacks *driver.AllocationCallbacks, o PipelineLayoutCreateInfo) (PipelineLayout, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var pipelineLayoutHandle driver.VkPipelineLayout
	res, err := d.deviceDriver.VkCreatePipelineLayout(d.deviceHandle, (*driver.VkPipelineLayoutCreateInfo)(createInfo), allocationCallbacks.Handle(), &pipelineLayoutHandle)
	if err != nil {
		return nil, res, err
	}

	pipelineLayout := createPipelineLayoutObject(d.deviceDriver, d.deviceHandle, pipelineLayoutHandle, d.maximumAPIVersion)

	return pipelineLayout, res, nil
}

func (d *VulkanDevice) CreateQueryPool(allocationCallbacks *driver.AllocationCallbacks, o QueryPoolCreateInfo) (QueryPool, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var queryPoolHandle driver.VkQueryPool

	res, err := d.deviceDriver.VkCreateQueryPool(d.deviceHandle, (*driver.VkQueryPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &queryPoolHandle)
	if err != nil {
		return nil, res, err
	}

	queryPool := createQueryPoolObject(d.deviceDriver, d.deviceHandle, queryPoolHandle, d.maximumAPIVersion)
	return queryPool, res, nil
}

func (d *VulkanDevice) CreateRenderPass(allocationCallbacks *driver.AllocationCallbacks, o RenderPassCreateInfo) (RenderPass, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var renderPassHandle driver.VkRenderPass

	res, err := d.deviceDriver.VkCreateRenderPass(d.deviceHandle, (*driver.VkRenderPassCreateInfo)(createInfo), allocationCallbacks.Handle(), &renderPassHandle)
	if err != nil {
		return nil, res, err
	}

	renderPass := createRenderPassObject(d.deviceDriver, d.deviceHandle, renderPassHandle, d.maximumAPIVersion)

	return renderPass, res, nil
}

func (d *VulkanDevice) CreateSampler(allocationCallbacks *driver.AllocationCallbacks, o SamplerCreateInfo) (Sampler, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var samplerHandle driver.VkSampler

	res, err := d.deviceDriver.VkCreateSampler(d.deviceHandle, (*driver.VkSamplerCreateInfo)(createInfo), allocationCallbacks.Handle(), &samplerHandle)
	if err != nil {
		return nil, res, err
	}

	sampler := createSamplerObject(d.deviceDriver, d.deviceHandle, samplerHandle, d.maximumAPIVersion)

	return sampler, res, nil
}

func (d *VulkanDevice) GetQueue(queueFamilyIndex int, queueIndex int) Queue {

	var queueHandle driver.VkQueue

	d.deviceDriver.VkGetDeviceQueue(d.deviceHandle, driver.Uint32(queueFamilyIndex), driver.Uint32(queueIndex), &queueHandle)

	queue := createQueueObject(d.deviceDriver, d.deviceHandle, queueHandle, d.maximumAPIVersion)

	return queue
}

func (d *VulkanDevice) AllocateMemory(allocationCallbacks *driver.AllocationCallbacks, o MemoryAllocateInfo) (DeviceMemory, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var deviceMemoryHandle driver.VkDeviceMemory

	deviceDriver := d.deviceDriver
	deviceHandle := d.deviceHandle

	res, err := deviceDriver.VkAllocateMemory(deviceHandle, (*driver.VkMemoryAllocateInfo)(createInfo), allocationCallbacks.Handle(), &deviceMemoryHandle)
	if err != nil {
		return nil, res, err
	}

	deviceMemory := createDeviceMemoryObject(deviceDriver, deviceHandle, deviceMemoryHandle, d.maximumAPIVersion, o.AllocationSize)

	return deviceMemory, res, nil
}

func (d *VulkanDevice) FreeMemory(deviceMemory DeviceMemory, allocationCallbacks *driver.AllocationCallbacks) {
	// This is really only here for a kind of API symmetry
	deviceMemory.Free(allocationCallbacks)
}

// Free a slice of command buffers which should all have the same device/driver/pool
// guaranteed to have at least one element
func (d *VulkanDevice) freeCommandBufferSlice(buffers []CommandBuffer) {
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
		bufferArraySlice[i] = driver.VkCommandBuffer(0)

		if buffers[i] != nil {
			bufferArraySlice[i] = buffers[i].Handle()
		}
	}

	bufferDriver.VkFreeCommandBuffers(bufferDevice, bufferPool, driver.Uint32(bufferCount), bufferArrayPtr)

	objStore := d.deviceDriver.ObjectStore()
	for i := 0; i < bufferCount; i++ {
		if buffers[i] != nil {
			objStore.Delete(driver.VulkanHandle(buffers[i].Handle()))
		}
	}
}

func (d *VulkanDevice) FreeCommandBuffers(buffers []CommandBuffer) {
	bufferCount := len(buffers)
	if bufferCount == 0 {
		return
	}

	multimap := make(map[driver.VkCommandPool][]CommandBuffer)
	for _, buffer := range buffers {
		poolHandle := buffer.CommandPoolHandle()
		existingSet := multimap[poolHandle]
		multimap[poolHandle] = append(existingSet, buffer)
	}

	for _, setBuffers := range multimap {
		d.freeCommandBufferSlice(setBuffers)
	}
}

func (d *VulkanDevice) AllocateCommandBuffers(o CommandBufferAllocateInfo) ([]CommandBuffer, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	if o.CommandPool == nil {
		return nil, VKErrorUnknown, errors.New("no command pool provided to allocate from")
	}

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	device := o.CommandPool.DeviceHandle()

	commandBufferPtr := (*driver.VkCommandBuffer)(arena.Malloc(o.CommandBufferCount * int(unsafe.Sizeof([1]driver.VkCommandBuffer{}))))

	res, err := o.CommandPool.Driver().VkAllocateCommandBuffers(device, (*driver.VkCommandBufferAllocateInfo)(createInfo), commandBufferPtr)
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	commandBufferArray := ([]driver.VkCommandBuffer)(unsafe.Slice(commandBufferPtr, o.CommandBufferCount))
	var result []CommandBuffer

	for i := 0; i < o.CommandBufferCount; i++ {
		commandBuffer := createCommandBufferObject(o.CommandPool.Driver(), o.CommandPool.Handle(), device, commandBufferArray[i], o.CommandPool.APIVersion())

		result = append(result, commandBuffer)
	}

	return result, res, nil
}

func (d *VulkanDevice) AllocateDescriptorSets(o DescriptorSetAllocateInfo) ([]DescriptorSet, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	if o.DescriptorPool == nil {
		return nil, VKErrorUnknown, errors.New("no descriptor pool provided to allocate from")
	}

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	device := o.DescriptorPool.DeviceHandle()
	poolDriver := o.DescriptorPool.Driver()

	setCount := len(o.SetLayouts)
	descriptorSets := (*driver.VkDescriptorSet)(arena.Malloc(setCount * int(unsafe.Sizeof([1]C.VkDescriptorSet{}))))

	res, err := poolDriver.VkAllocateDescriptorSets(device, (*driver.VkDescriptorSetAllocateInfo)(createInfo), descriptorSets)
	if err != nil {
		return nil, res, err
	}

	var sets []DescriptorSet
	descriptorSetSlice := ([]driver.VkDescriptorSet)(unsafe.Slice(descriptorSets, setCount))

	for i := 0; i < setCount; i++ {
		descriptorSet := createDescriptorSetObject(poolDriver, device, o.DescriptorPool.Handle(), descriptorSetSlice[i], o.DescriptorPool.APIVersion())

		sets = append(sets, descriptorSet)
	}

	return sets, res, nil
}

// Free a slice of descriptor sets which should all have the same device/driver/pool
// guaranteed to have at least one element
func (d *VulkanDevice) freeDescriptorSetSlice(sets []DescriptorSet) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	setSize := len(sets)
	arraySize := setSize * int(unsafe.Sizeof([1]C.VkDescriptorSet{}))

	arrayPtr := (*driver.VkDescriptorSet)(arena.Malloc(arraySize))
	arraySlice := ([]driver.VkDescriptorSet)(unsafe.Slice(arrayPtr, setSize))

	for i := 0; i < setSize; i++ {
		arraySlice[i] = driver.VkDescriptorSet(0)
		if sets[i] != nil {
			arraySlice[i] = sets[i].Handle()
		}
	}

	setDriver := sets[0].Driver()
	pool := sets[0].DescriptorPoolHandle()
	device := sets[0].DeviceHandle()

	res, err := setDriver.VkFreeDescriptorSets(device, pool, driver.Uint32(setSize), arrayPtr)
	if err != nil {
		return res, err
	}

	objStore := setDriver.ObjectStore()
	for i := 0; i < setSize; i++ {
		if sets[i] != nil {
			objStore.Delete(driver.VulkanHandle(sets[i].Handle()))
		}
	}

	return res, nil
}

func (d *VulkanDevice) FreeDescriptorSets(sets []DescriptorSet) (common.VkResult, error) {
	poolMultimap := make(map[driver.VkDescriptorPool][]DescriptorSet)

	for _, set := range sets {
		poolHandle := set.DescriptorPoolHandle()
		existingSet := poolMultimap[poolHandle]
		poolMultimap[poolHandle] = append(existingSet, set)
	}

	var res common.VkResult
	var err error
	for _, set := range poolMultimap {
		res, err = d.freeDescriptorSetSlice(set)
		if err != nil {
			return res, err
		}
	}

	return res, err
}
