package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	"time"
)

//go:generate mockgen -source ./iface.go -destination ../mocks/core1_2_mocks.go -package mocks -mock_names CommandBuffer=CommandBuffer1_2,CommandPool=CommandPool1_2,Device=Device1_2,DescriptorUpdateTemplate=DescriptorUpdateTemplate1_2,Instance=Instance1_2,PhysicalDevice=PhysicalDevice1_2,Buffer=Buffer1_2,BufferView=BufferView1_2,DescriptorPool=DescriptorPool1_2,DescriptorSet=DescriptorSet1_2,DescriptorSetLayout=DescriptorSetLayout1_2,DeviceMemory=DeviceMemory1_2,Event=Event1_2,Fence=Fence1_2,Framebuffer=Framebuffer1_2,Image=Image1_2,ImageView=ImageView1_2,Pipeline=Pipeline1_2,PipelineCache=PipelineCache1_2,PipelineLayout=PipelineLayout1_2,QueryPool=QueryPool1_2,Queue=Queue1_2,RenderPass=RenderPass1_2,Sampler=Sampler1_2,Semaphore=Semaphore1_2,ShaderModule=ShaderModule1_2,InstanceScopedPhysicalDevice=InstanceScopedPhysicalDevice1_2,SamplerYcbcrConversion=SamplerYcbcrConversion1_2,DescriptorUpdateTemplate=DescriptorUpdateTemplate1_2

type Buffer interface {
	core1_1.Buffer
}

type BufferView interface {
	core1_1.BufferView
}

type CommandBuffer interface {
	core1_1.CommandBuffer

	CmdBeginRenderPass2(renderPassBegin core1_0.RenderPassBeginInfo, subpassBegin SubpassBeginInfo) error
	CmdEndRenderPass2(subpassEnd SubpassEndInfo) error
	CmdNextSubpass2(subpassBegin SubpassBeginInfo, subpassEnd SubpassEndInfo) error
	CmdDrawIndexedIndirectCount(buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int)
	CmdDrawIndirectCount(buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int)
}

type CommandPool interface {
	core1_1.CommandPool
}

type DescriptorPool interface {
	core1_1.DescriptorPool
}

type DescriptorSet interface {
	core1_1.DescriptorSet
}

type DescriptorSetLayout interface {
	core1_1.DescriptorSetLayout
}

type Device interface {
	core1_1.Device

	CreateRenderPass2(allocator *driver.AllocationCallbacks, options RenderPassCreateOptions) (core1_0.RenderPass, common.VkResult, error)
	GetBufferDeviceAddress(o BufferDeviceAddressInfo) (uint64, error)
	GetBufferOpaqueCaptureAddress(o BufferDeviceAddressInfo) (uint64, error)
	GetDeviceMemoryOpaqueCaptureAddress(o DeviceMemoryOpaqueCaptureAddressInfo) (uint64, error)

	SignalSemaphore(o SemaphoreSignalInfo) (common.VkResult, error)
	WaitSemaphores(timeout time.Duration, o SemaphoreWaitInfo) (common.VkResult, error)
}

type DeviceMemory interface {
	core1_1.DeviceMemory
}

type DescriptorUpdateTemplate interface {
	core1_1.DescriptorUpdateTemplate
}

type Event interface {
	core1_1.Event
}

type Fence interface {
	core1_1.Fence
}

type Framebuffer interface {
	core1_1.Framebuffer
}

type Image interface {
	core1_1.Image
}

type ImageView interface {
	core1_1.ImageView
}

type Instance interface {
	core1_1.Instance
}

type InstanceScopedPhysicalDevice interface {
	core1_1.InstanceScopedPhysicalDevice
}

type PhysicalDevice interface {
	core1_1.PhysicalDevice

	InstanceScopedPhysicalDevice1_2() InstanceScopedPhysicalDevice
}

type Pipeline interface {
	core1_1.Pipeline
}

type PipelineCache interface {
	core1_1.PipelineCache
}

type PipelineLayout interface {
	core1_1.PipelineLayout
}

type QueryPool interface {
	core1_1.QueryPool

	Reset(firstQuery, queryCount int)
}

type Queue interface {
	core1_1.Queue
}

type RenderPass interface {
	core1_1.RenderPass
}

type Sampler interface {
	core1_1.Sampler
}

type SamplerYcbcrConversion interface {
	core1_1.SamplerYcbcrConversion
}

type Semaphore interface {
	core1_1.Semaphore

	CounterValue() (uint64, common.VkResult, error)
}

type ShaderModule interface {
	core1_1.ShaderModule
}
