package iface

func CommonBuffers[BufferImpl Buffer](buffers []BufferImpl) []Buffer {
	commonBuffer := make([]Buffer, len(buffers))
	for i := 0; i < len(buffers); i++ {
		commonBuffer[i] = Buffer(buffers[i])
	}

	return commonBuffer
}

func CommonBufferViews[BufferViewImpl BufferView](bufferViews []BufferViewImpl) []BufferView {
	commonBufferView := make([]BufferView, len(bufferViews))
	for i := 0; i < len(bufferViews); i++ {
		commonBufferView[i] = BufferView(bufferViews[i])
	}

	return commonBufferView
}

func CommonCommandBuffers[CommandBufferImpl CommandBuffer](commandBuffers []CommandBufferImpl) []CommandBuffer {
	commonCommandBuffers := make([]CommandBuffer, len(commandBuffers))
	for i := 0; i < len(commandBuffers); i++ {
		commonCommandBuffers[i] = CommandBuffer(commandBuffers[i])
	}

	return commonCommandBuffers
}

func CommonCommandPools[CommandPoolImpl CommandPool](commandPools []CommandPoolImpl) []CommandPool {
	commonCommandPools := make([]CommandPool, len(commandPools))
	for i := 0; i < len(commandPools); i++ {
		commonCommandPools[i] = CommandPool(commandPools[i])
	}

	return commonCommandPools
}

func CommonDescriptorPools[DescriptorPoolImpl DescriptorPool](descriptorPools []DescriptorPoolImpl) []DescriptorPool {
	commonDescriptorPools := make([]DescriptorPool, len(descriptorPools))
	for i := 0; i < len(commonDescriptorPools); i++ {
		commonDescriptorPools[i] = DescriptorPool(descriptorPools[i])
	}

	return commonDescriptorPools
}

func CommonDescriptorSets[DescriptorSetImpl DescriptorSet](descriptorSets []DescriptorSetImpl) []DescriptorSet {
	commonDescriptorSets := make([]DescriptorSet, len(descriptorSets))
	for i := 0; i < len(commonDescriptorSets); i++ {
		commonDescriptorSets[i] = DescriptorSet(descriptorSets[i])
	}

	return commonDescriptorSets
}

func CommonDescriptorSetLayouts[DescriptorSetLayoutImpl DescriptorSetLayout](descriptorSetLayouts []DescriptorSetLayoutImpl) []DescriptorSetLayout {
	commonDescriptorSetLayouts := make([]DescriptorSetLayout, len(descriptorSetLayouts))
	for i := 0; i < len(commonDescriptorSetLayouts); i++ {
		commonDescriptorSetLayouts[i] = DescriptorSetLayout(descriptorSetLayouts[i])
	}

	return commonDescriptorSetLayouts
}

func CommonDeviceMemories[DeviceMemoryImpl DeviceMemory](deviceMemories []DeviceMemoryImpl) []DeviceMemory {
	commonDeviceMemories := make([]DeviceMemory, len(deviceMemories))
	for i := 0; i < len(commonDeviceMemories); i++ {
		commonDeviceMemories[i] = DeviceMemory(deviceMemories[i])
	}

	return commonDeviceMemories
}

func CommonDevices[DeviceImpl Device](devices []DeviceImpl) []Device {
	commonDevices := make([]Device, len(devices))
	for i := 0; i < len(commonDevices); i++ {
		commonDevices[i] = Device(devices[i])
	}

	return commonDevices
}

func CommonEvents[EventImpl Event](events []EventImpl) []Event {
	commonEvents := make([]Event, len(events))
	for i := 0; i < len(commonEvents); i++ {
		commonEvents[i] = Event(events[i])
	}

	return commonEvents
}

func CommonFences[FenceImpl Fence](fences []FenceImpl) []Fence {
	commonFences := make([]Fence, len(fences))
	for i := 0; i < len(commonFences); i++ {
		commonFences[i] = Fence(fences[i])
	}

	return commonFences
}

func CommonFramebuffer[FramebufferImpl Framebuffer](framebuffers []FramebufferImpl) []Framebuffer {
	commonFramebuffers := make([]Framebuffer, len(framebuffers))
	for i := 0; i < len(commonFramebuffers); i++ {
		commonFramebuffers[i] = Framebuffer(framebuffers[i])
	}

	return commonFramebuffers
}

func CommonImages[ImageImpl Image](images []ImageImpl) []Image {
	commonImages := make([]Image, len(images))
	for i := 0; i < len(commonImages); i++ {
		commonImages[i] = Image(images[i])
	}

	return commonImages
}

func CommonImageViews[ImageViewImpl ImageView](imageViews []ImageViewImpl) []ImageView {
	commonImageViews := make([]ImageView, len(imageViews))
	for i := 0; i < len(commonImageViews); i++ {
		commonImageViews[i] = ImageView(imageViews[i])
	}

	return commonImageViews
}

func CommonInstances[InstanceImpl Instance](instances []InstanceImpl) []Instance {
	commonInstances := make([]Instance, len(instances))
	for i := 0; i < len(commonInstances); i++ {
		commonInstances[i] = Instance(instances[i])
	}

	return commonInstances
}

func CommonPhysicalDevices[PhysicalDeviceImpl PhysicalDevice](physicalDevices []PhysicalDeviceImpl) []PhysicalDevice {
	commonPhysicalDevices := make([]PhysicalDevice, len(physicalDevices))
	for i := 0; i < len(commonPhysicalDevices); i++ {
		commonPhysicalDevices[i] = PhysicalDevice(physicalDevices[i])
	}

	return commonPhysicalDevices
}

func CommonPipelines[PipelineImpl Pipeline](pipelines []PipelineImpl) []Pipeline {
	commonPipelines := make([]Pipeline, len(pipelines))
	for i := 0; i < len(commonPipelines); i++ {
		commonPipelines[i] = Pipeline(pipelines[i])
	}

	return commonPipelines
}

func CommonPipelineCaches[PipelineCacheImpl PipelineCache](pipelineCaches []PipelineCacheImpl) []PipelineCache {
	commonPipelineCaches := make([]PipelineCache, len(pipelineCaches))
	for i := 0; i < len(commonPipelineCaches); i++ {
		commonPipelineCaches[i] = PipelineCache(pipelineCaches[i])
	}

	return commonPipelineCaches
}

//
//
//type PipelineCache interface {
//	Handle() driver.VkPipelineCache
//}
//
//type PipelineLayout interface {
//	Handle() driver.VkPipelineLayout
//}
//
//type QueryPool interface {
//	Handle() driver.VkQueryPool
//}
//
//type Queue interface {
//	Handle() driver.VkQueue
//	Driver() driver.Driver
//}
//
//type RenderPass interface {
//	Handle() driver.VkRenderPass
//}
//
//type Semaphore interface {
//	Handle() driver.VkSemaphore
//}
//
//type ShaderModule interface {
//	Handle() driver.VkShaderModule
//}
//
//type Sampler interface {
//	Handle() driver.VkSampler
//}
