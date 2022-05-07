package core

import (
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
)

//go:generate mockgen -source ./objects.go -destination mocks/core1_0_mocks.go -package mocks

type Buffer interface {
	core1_0.Buffer
}

type BufferView interface {
	core1_0.BufferView
}

type CommandBuffer interface {
	core1_0.CommandBuffer

	Core1_1() core1_1.CommandBuffer
}

type CommandPool interface {
	core1_0.CommandPool

	Core1_1() core1_1.CommandPool
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

	Core1_1() core1_1.Device
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

	Core1_1() core1_1.Instance
}

type PhysicalDevice interface {
	core1_0.PhysicalDevice

	Core1_1() core1_1.PhysicalDevice
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

type Sampler interface {
	core1_0.Sampler
}

type Semaphore interface {
	core1_0.Semaphore
}

type ShaderModule interface {
	core1_0.ShaderModule
}

type DescriptorUpdateTemplate interface {
	core1_1.DescriptorUpdateTemplate
}

type SamplerYcbcrConversion interface {
	core1_1.SamplerYcbcrConversion
}
