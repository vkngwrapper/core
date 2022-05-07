package core1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
)

//go:generate mockgen -source ./iface.go -destination ../mocks/core1_1_mocks.go -package mocks -mock_names CommandBuffer=CommandBuffer1_1,CommandPool=CommandPool1_1,Device=Device1_1,DescriptorUpdateTemplate=DescriptorUpdateTemplate1_1,Instance=Instance1_1,PhysicalDevice=PhysicalDevice1_1,SamplerYcbcrConversion=SamplerYcbcrConversion1_1

type CommandBuffer interface {
	CmdDispatchBase(baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ int)
	CmdSetDeviceMask(deviceMask uint32)
}

type CommandPool interface {
	TrimCommandPool(flags CommandPoolTrimFlags)
}

type Device interface {
	BindBufferMemory(o []BindBufferMemoryOptions) (common.VkResult, error)
	BindImageMemory(o []BindImageMemoryOptions) (common.VkResult, error)

	BufferMemoryRequirements(o BufferMemoryRequirementsOptions, out *MemoryRequirementsOutData) error
	ImageMemoryRequirements(o ImageMemoryRequirementsOptions, out *MemoryRequirementsOutData) error
	SparseImageMemoryRequirements(o SparseImageRequirementsOptions, outDataFactory func() *SparseImageMemoryRequirementsOutData) ([]*SparseImageMemoryRequirementsOutData, error)

	DescriptorSetLayoutSupport(o core1_0.DescriptorSetLayoutCreateOptions, outData *DescriptorSetLayoutSupportOutData) error
}

type DescriptorUpdateTemplate interface {
	Handle() driver.VkDescriptorUpdateTemplate
	Destroy(allocator *driver.AllocationCallbacks)

	UpdateDescriptorSetFromBuffer(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorBufferInfo)
	UpdateDescriptorSetFromImage(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorImageInfo)
	UpdateDescriptorSetFromObjectHandle(descriptorSet core1_0.DescriptorSet, data driver.VulkanHandle)
}

type Instance interface {
	EnumeratePhysicalDeviceGroups(outDataFactory func() *DeviceGroupOutData) ([]*DeviceGroupOutData, common.VkResult, error)
}

type PhysicalDevice interface {
	ExternalFenceProperties(o ExternalFenceOptions, outData *ExternalFenceOutData) error
	ExternalBufferProperties(o ExternalBufferOptions, outData *ExternalBufferOutData) error
	ExternalSemaphoreProperties(o ExternalSemaphoreOptions, outData *ExternalSemaphoreOutData) error

	Features(out *DeviceFeaturesOutData) error
	FormatProperties(format common.DataFormat, out *FormatPropertiesOutData) error
	ImageFormatProperties(o ImageFormatOptions, out *ImageFormatOutData) (common.VkResult, error)
	MemoryProperties(out *MemoryPropertiesOutData) error
	Properties(out *DevicePropertiesOutData) error
	QueueFamilyProperties(outDataFactory func() *QueueFamilyOutData) ([]*QueueFamilyOutData, error)
	SparseImageFormatProperties(o SparseImageFormatOptions, outDataFactory func() *SparseImageFormatPropertiesOutData) ([]*SparseImageFormatPropertiesOutData, error)
}

type SamplerYcbcrConversion interface {
	Handle() driver.VkSamplerYcbcrConversion
	Destroy(allocator *driver.AllocationCallbacks)
}
