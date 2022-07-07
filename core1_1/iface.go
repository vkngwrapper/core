package core1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
)

//go:generate mockgen -source ./iface.go -destination ../mocks/core1_1_mocks.go -package mocks -mock_names CommandBuffer=CommandBuffer1_1,CommandPool=CommandPool1_1,Device=Device1_1,DescriptorUpdateTemplate=DescriptorUpdateTemplate1_1,Instance=Instance1_1,PhysicalDevice=PhysicalDevice1_1,Buffer=Buffer1_1,BufferView=BufferView1_1,DescriptorPool=DescriptorPool1_1,DescriptorSet=DescriptorSet1_1,DescriptorSetLayout=DescriptorSetLayout1_1,DeviceMemory=DeviceMemory1_1,Event=Event1_1,Fence=Fence1_1,Framebuffer=Framebuffer1_1,Image=Image1_1,ImageView=ImageView1_1,Pipeline=Pipeline1_1,PipelineCache=PipelineCache1_1,PipelineLayout=PipelineLayout1_1,QueryPool=QueryPool1_1,Queue=Queue1_1,RenderPass=RenderPass1_1,Sampler=Sampler1_1,Semaphore=Semaphore1_1,ShaderModule=ShaderModule1_1

type Buffer interface {
	core1_0.Buffer
}

type BufferView interface {
	core1_0.BufferView
}

type CommandBuffer interface {
	core1_0.CommandBuffer

	CmdDispatchBase(baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ int)
	CmdSetDeviceMask(deviceMask uint32)
}

type CommandPool interface {
	core1_0.CommandPool

	TrimCommandPool(flags CommandPoolTrimFlags)
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

type Device interface {
	core1_0.Device

	BindBufferMemory2(o []BindBufferMemoryInfo) (common.VkResult, error)
	BindImageMemory2(o []BindImageMemoryInfo) (common.VkResult, error)

	BufferMemoryRequirements2(o BufferMemoryRequirementsInfo2, out *MemoryRequirements2) error
	ImageMemoryRequirements2(o ImageMemoryRequirementsInfo2, out *MemoryRequirements2) error
	ImageSparseMemoryRequirements2(o ImageSparseMemoryRequirementsInfo2, outDataFactory func() *SparseImageMemoryRequirements2) ([]*SparseImageMemoryRequirements2, error)

	DescriptorSetLayoutSupport(o core1_0.DescriptorSetLayoutCreateInfo, outData *DescriptorSetLayoutSupport) error

	DeviceGroupPeerMemoryFeatures(heapIndex, localDeviceIndex, remoteDeviceIndex int) PeerMemoryFeatureFlags

	CreateDescriptorUpdateTemplate(o DescriptorUpdateTemplateCreateInfo, allocator *driver.AllocationCallbacks) (DescriptorUpdateTemplate, common.VkResult, error)
	CreateSamplerYcbcrConversion(o SamplerYcbcrConversionCreateInfo, allocator *driver.AllocationCallbacks) (SamplerYcbcrConversion, common.VkResult, error)

	GetQueue2(o DeviceQueueInfo2) (core1_0.Queue, error)
}

type DeviceMemory interface {
	core1_0.DeviceMemory
}

type DescriptorUpdateTemplate interface {
	Handle() driver.VkDescriptorUpdateTemplate
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(allocator *driver.AllocationCallbacks)
	UpdateDescriptorSetFromBuffer(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorBufferInfo)
	UpdateDescriptorSetFromImage(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorImageInfo)
	UpdateDescriptorSetFromObjectHandle(descriptorSet core1_0.DescriptorSet, data driver.VulkanHandle)
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

	EnumeratePhysicalDeviceGroups(outDataFactory func() *PhysicalDeviceGroupProperties) ([]*PhysicalDeviceGroupProperties, common.VkResult, error)
}

type InstanceScopedPhysicalDevice interface {
	core1_0.PhysicalDevice

	ExternalFenceProperties(o PhysicalDeviceExternalFenceInfo, outData *ExternalFenceProperties) error
	ExternalBufferProperties(o PhysicalDeviceExternalBufferInfo, outData *ExternalBufferProperties) error
	ExternalSemaphoreProperties(o PhysicalDeviceExternalSemaphoreInfo, outData *ExternalSemaphoreProperties) error

	Features2(out *PhysicalDeviceFeatures2) error
	FormatProperties2(format core1_0.Format, out *FormatProperties2) error
	ImageFormatProperties2(o PhysicalDeviceImageFormatInfo2, out *ImageFormatProperties2) (common.VkResult, error)
	MemoryProperties2(out *PhysicalDeviceMemoryProperties2) error
	Properties2(out *PhysicalDeviceProperties2) error
	QueueFamilyProperties2(outDataFactory func() *QueueFamilyProperties2) ([]*QueueFamilyProperties2, error)
	SparseImageFormatProperties2(o PhysicalDeviceSparseImageFormatInfo2, outDataFactory func() *SparseImageFormatProperties2) ([]*SparseImageFormatProperties2, error)
}

type PhysicalDevice interface {
	core1_0.PhysicalDevice

	InstanceScopedPhysicalDevice1_1() InstanceScopedPhysicalDevice
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

type SamplerYcbcrConversion interface {
	Handle() driver.VkSamplerYcbcrConversion
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(allocator *driver.AllocationCallbacks)
}

type Semaphore interface {
	core1_0.Semaphore
}

type ShaderModule interface {
	core1_0.ShaderModule
}
