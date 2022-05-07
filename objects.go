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

// PhysicalDevice straddles the line between the "instance" and "device" section
// of the vulkan spec. They have "instance" level functionality which will be available
// if the instance supports core 1.x, even if the device does not, and they have "device"
// level functionality which is only available if both the instance and device support 1.x
//
// As a result, PhysicalDevice's Core1_X() methods indicate whether they are for instance
// or device.
//
// It is hypothetically possible for a core 1.x version to have a Core1_XInstance() and
// Core1_XDevice method here, and one of those will return nil while the other does not
// in certain circumstances. No core version has had that situation so far, though.
type PhysicalDevice interface {
	core1_0.PhysicalDevice

	// Core1_1Instance retrieves the 1.1 methods for physical devices
	// that were promoted from instance extensions
	Core1_1Instance() core1_1.InstancePhysicalDevice
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
