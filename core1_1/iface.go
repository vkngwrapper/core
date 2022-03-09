package core1_1

//go:generate mockgen -source ./iface.go -destination ../mocks/core1_1_mocks.go -package mocks -mock_names Buffer=MockBuffer1_1,BufferView=MockBufferView1_1,CommandBuffer=MockCommandBuffer1_1,CommandPool=MockCommandPool1_1,DescriptorPool=MockDescriptorPool1_1,DescriptorSet=MockDescriptorSet1_1,DescriptorSetLayout=MockDescriptorSetLayout1_1,DeviceMemory=MockDeviceMemory1_1,Device=MockDevice1_1,Event=MockEvent1_1,Fence=MockFence1_1,Framebuffer=MockFramebuffer1_1,Image=MockImage1_1,ImageView=MockImageView1_1,Instance=MockInstance1_1,PhysicalDevice=MockPhysicalDevice1_1,Pipeline=MockPipeline1_1,PipelineCache=MockPipelineCache1_1,PipelineLayout=MockPipelineLayout1_1,QueryPool=MockQueryPool1_1,Queue=MockQueue1_1,RenderPass=MockRenderPass1_1,Semaphore=MockSemaphore1_1,ShaderModule=MockShaderModule1_1,Sampler=MockSampler1_1

type Buffer interface {
}

type BufferView interface {
}

type CommandBuffer interface {
}

type CommandPool interface {
}

type DescriptorPool interface {
}

type DescriptorSet interface {
}

type DescriptorSetLayout interface {
}

type DeviceMemory interface {
}

type Device interface {
}

type Event interface {
}

type Fence interface {
}

type Framebuffer interface {
}

type Image interface {
}

type ImageView interface {
}

type Instance interface {
}

type PhysicalDevice interface {
}

type Pipeline interface {
}

type PipelineCache interface {
}

type PipelineLayout interface {
}

type QueryPool interface {
}

type Queue interface {
}

type RenderPass interface {
}

type Semaphore interface {
}

type ShaderModule interface {
}

type Sampler interface {
}
