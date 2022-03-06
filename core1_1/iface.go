package core1_1

import (
	"github.com/CannibalVox/VKng/core/core1_0"
)

//go:generate mockgen -source ./iface.go -destination ../mocks/core1_1_mocks.go -package mocks -mock_names Buffer=MockBuffer1_1,BufferView=MockBufferView1_1,CommandBuffer=MockCommandBuffer1_1,CommandPool=MockCommandPool1_1,DescriptorPool=MockDescriptorPool1_1,DescriptorSet=MockDescriptorSet1_1,DescriptorSetLayout=MockDescriptorSetLayout1_1,DeviceMemory=MockDeviceMemory1_1,Device=MockDevice1_1,Event=MockEvent1_1,Fence=MockFence1_1,Framebuffer=MockFramebuffer1_1,Image=MockImage1_1,ImageView=MockImageView1_1,Instance=MockInstance1_1,PhysicalDevice=MockPhysicalDevice1_1,Pipeline=MockPipeline1_1,PipelineCache=MockPipelineCache1_1,PipelineLayout=MockPipelineLayout1_1,QueryPool=MockQueryPool1_1,Queue=MockQueue1_1,RenderPass=MockRenderPass1_1,Semaphore=MockSemaphore1_1,ShaderModule=MockShaderModule1_1,Sampler=MockSampler1_1

type Buffer interface {
	core1_0.Buffer
}

type BufferView interface {
	core1_0.BufferView
}

type CommandBuffer interface {
	core1_0.CommandBuffer
}

type CommandPool interface {
	core1_0.CommandPool
}

type DescriptorPool interface {
	core1_0.DescriptorPool
}

type DescriptorSet interface {
	core1_0.DescriptorSet
}

type DescriptorSetLayout interface {
	core1_0.DescriptorSetLayout
}

type DeviceMemory interface {
	core1_0.DeviceMemory
}

type Device interface {
	core1_0.Device
}

type Event interface {
	core1_0.Event
}

type Fence interface {
	core1_0.Fence
}

type Framebuffer interface {
	core1_0.Framebuffer
}

type Image interface {
	core1_0.Image
}

type ImageView interface {
	core1_0.ImageView
}

type Instance interface {
	core1_0.Instance
}

type PhysicalDevice interface {
	core1_0.PhysicalDevice
}

type Pipeline interface {
	core1_0.Pipeline
}

type PipelineCache interface {
	core1_0.PipelineCache
}

type PipelineLayout interface {
	core1_0.PipelineLayout
}

type QueryPool interface {
	core1_0.QueryPool
}

type Queue interface {
	core1_0.Queue
}

type RenderPass interface {
	core1_0.RenderPass
}

type Semaphore interface {
	core1_0.Semaphore
}

type ShaderModule interface {
	core1_0.ShaderModule
}

type Sampler interface {
	core1_0.Sampler
}
