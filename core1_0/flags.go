package core1_0

import "github.com/vkngwrapper/core/common"

// AccessFlags specifies memory access types that will participate in a memory dependency
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
type AccessFlags int32

var accessFlagsMapping = common.NewFlagStringMapping[AccessFlags]()

func (f AccessFlags) Register(str string) {
	accessFlagsMapping.Register(f, str)
}
func (f AccessFlags) String() string {
	return accessFlagsMapping.FlagsToString(f)
}

////

// AttachmentDescriptionFlags specifies additional properties of an attachment
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionFlagBits.html
type AttachmentDescriptionFlags int32

var attachmentDescriptionFlagsMapping = common.NewFlagStringMapping[AttachmentDescriptionFlags]()

func (f AttachmentDescriptionFlags) Register(str string) {
	attachmentDescriptionFlagsMapping.Register(f, str)
}

func (f AttachmentDescriptionFlags) String() string {
	return attachmentDescriptionFlagsMapping.FlagsToString(f)
}

////

// CommandBufferUsageFlags specifies usage behavior for a CommandBuffer
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferUsageFlagBits.html
type CommandBufferUsageFlags int32

var beginInfoFlagsMapping = common.NewFlagStringMapping[CommandBufferUsageFlags]()

func (f CommandBufferUsageFlags) Register(str string) {
	beginInfoFlagsMapping.Register(f, str)
}

func (f CommandBufferUsageFlags) String() string {
	return beginInfoFlagsMapping.FlagsToString(f)
}

////

// BufferCreateFlags specifies additional parameters of a Buffer
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateFlagBits.html
type BufferCreateFlags int32

var bufferCreateFlagsMapping = common.NewFlagStringMapping[BufferCreateFlags]()

func (f BufferCreateFlags) Register(str string) {
	bufferCreateFlagsMapping.Register(f, str)
}

func (f BufferCreateFlags) String() string {
	return bufferCreateFlagsMapping.FlagsToString(f)
}

////

// BufferUsageFlags specifies allowed usage of a Buffer
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html
type BufferUsageFlags int32

var bufferUsagesMapping = common.NewFlagStringMapping[BufferUsageFlags]()

func (f BufferUsageFlags) Register(str string) {
	bufferUsagesMapping.Register(f, str)
}

func (f BufferUsageFlags) String() string {
	return bufferUsagesMapping.FlagsToString(f)
}

////

// CommandBufferResetFlags controls behavior of a CommandBuffer reset
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferResetFlagBits.html
type CommandBufferResetFlags int32

var commandBufferResetFlagsMapping = common.NewFlagStringMapping[CommandBufferResetFlags]()

func (f CommandBufferResetFlags) Register(str string) {
	commandBufferResetFlagsMapping.Register(f, str)
}

func (f CommandBufferResetFlags) String() string {
	return commandBufferResetFlagsMapping.FlagsToString(f)
}

////

// CommandPoolCreateFlags specifies usage behavior for a CommandPool
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateFlagBits.html
type CommandPoolCreateFlags int32

var commandPoolCreateFlagsMapping = common.NewFlagStringMapping[CommandPoolCreateFlags]()

func (f CommandPoolCreateFlags) Register(str string) {
	commandPoolCreateFlagsMapping.Register(f, str)
}

func (f CommandPoolCreateFlags) String() string {
	return commandPoolCreateFlagsMapping.FlagsToString(f)
}

////

// CommandPoolResetFlags controls behavior of a CommandPool reset
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolResetFlagBits.html
type CommandPoolResetFlags int32

var commandPoolResetFlagsMapping = common.NewFlagStringMapping[CommandPoolResetFlags]()

func (f CommandPoolResetFlags) Register(str string) {
	commandPoolResetFlagsMapping.Register(f, str)
}

func (f CommandPoolResetFlags) String() string {
	return commandPoolResetFlagsMapping.FlagsToString(f)
}

////

// CullModeFlags controls triangle culling
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCullModeFlagBits.html
type CullModeFlags int32

var cullModesMapping = common.NewFlagStringMapping[CullModeFlags]()

func (f CullModeFlags) Register(str string) {
	cullModesMapping.Register(f, str)
}

func (f CullModeFlags) String() string {
	return cullModesMapping.FlagsToString(f)
}

////

// DependencyFlags specifies how execution and memory dependencies are formed
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlagBits.html
type DependencyFlags int32

var dependencyFlagsMapping = common.NewFlagStringMapping[DependencyFlags]()

func (f DependencyFlags) Register(str string) {
	dependencyFlagsMapping.Register(f, str)
}

func (f DependencyFlags) String() string {
	return dependencyFlagsMapping.FlagsToString(f)
}

////

// DescriptorPoolCreateFlags specifies parameters of a newly-created DescriptorPool
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateInfo.html
type DescriptorPoolCreateFlags int32

var descriptorPoolCreateFlagsMapping = common.NewFlagStringMapping[DescriptorPoolCreateFlags]()

func (f DescriptorPoolCreateFlags) Register(str string) {
	descriptorPoolCreateFlagsMapping.Register(f, str)
}

func (f DescriptorPoolCreateFlags) String() string {
	return descriptorPoolCreateFlagsMapping.FlagsToString(f)
}

////

// DescriptorPoolResetFlags is reserved for future use
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolResetFlags.html
type DescriptorPoolResetFlags int32

var descriptorPoolResetFlagsMapping = common.NewFlagStringMapping[DescriptorPoolResetFlags]()

func (f DescriptorPoolResetFlags) Register(str string) {
	descriptorPoolResetFlagsMapping.Register(f, str)
}

func (f DescriptorPoolResetFlags) String() string {
	return descriptorPoolResetFlagsMapping.FlagsToString(f)
}

////

// DescriptorSetLayoutCreateFlags specifies DescriptorSetLayout properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateFlagBits.html
type DescriptorSetLayoutCreateFlags int32

var descriptorSetLayoutCreateFlagsMapping = common.NewFlagStringMapping[DescriptorSetLayoutCreateFlags]()

func (f DescriptorSetLayoutCreateFlags) Register(str string) {
	descriptorSetLayoutCreateFlagsMapping.Register(f, str)
}

func (f DescriptorSetLayoutCreateFlags) String() string {
	return descriptorSetLayoutCreateFlagsMapping.FlagsToString(f)
}

////

// DeviceQueueCreateFlags specifies behavior of a Queue
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateFlagBits.html
type DeviceQueueCreateFlags int32

var deviceQueueCreateFlagsMapping = common.NewFlagStringMapping[DeviceQueueCreateFlags]()

func (f DeviceQueueCreateFlags) Register(str string) {
	deviceQueueCreateFlagsMapping.Register(f, str)
}

func (f DeviceQueueCreateFlags) String() string {
	return deviceQueueCreateFlagsMapping.FlagsToString(f)
}

////

// EventCreateFlags represents event creation flag bits
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEventCreateFlagBits.html
type EventCreateFlags int32

var eventCreateFlagsMapping = common.NewFlagStringMapping[EventCreateFlags]()

func (f EventCreateFlags) Register(str string) {
	eventCreateFlagsMapping.Register(f, str)
}

func (f EventCreateFlags) String() string {
	return eventCreateFlagsMapping.FlagsToString(f)
}

////

// FenceCreateFlags specifies initial state and behavior of a Fence
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceCreateFlagBits.html
type FenceCreateFlags int32

var fenceCreateFlagsMapping = common.NewFlagStringMapping[FenceCreateFlags]()

func (f FenceCreateFlags) Register(str string) {
	fenceCreateFlagsMapping.Register(f, str)
}

func (f FenceCreateFlags) String() string {
	return fenceCreateFlagsMapping.FlagsToString(f)
}

////

// FormatFeatureFlags specifies features supported by a Buffer
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
type FormatFeatureFlags int32

var formatFeaturesMapping = common.NewFlagStringMapping[FormatFeatureFlags]()

func (f FormatFeatureFlags) Register(str string) {
	formatFeaturesMapping.Register(f, str)
}

func (f FormatFeatureFlags) String() string {
	return formatFeaturesMapping.FlagsToString(f)
}

////

// FramebufferCreateFlags specifies Framebuffer properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateFlagBits.html
type FramebufferCreateFlags int32

var framebufferCreateFlagsMapping = common.NewFlagStringMapping[FramebufferCreateFlags]()

func (f FramebufferCreateFlags) Register(str string) {
	framebufferCreateFlagsMapping.Register(f, str)
}

func (f FramebufferCreateFlags) String() string {
	return framebufferCreateFlagsMapping.FlagsToString(f)
}

////

// ImageAspectFlags specifies which aspects of an Image are included in a view
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html
type ImageAspectFlags int32

var imageAspectFlagsMapping = common.NewFlagStringMapping[ImageAspectFlags]()

func (f ImageAspectFlags) Register(str string) {
	imageAspectFlagsMapping.Register(f, str)
}

func (f ImageAspectFlags) String() string {
	return imageAspectFlagsMapping.FlagsToString(f)
}

////

// ImageCreateFlags specifies additional parameters of an Image
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
type ImageCreateFlags int32

var imageCreateFlagsMapping = common.NewFlagStringMapping[ImageCreateFlags]()

func (f ImageCreateFlags) Register(str string) {
	imageCreateFlagsMapping.Register(f, str)
}

func (f ImageCreateFlags) String() string {
	return imageCreateFlagsMapping.FlagsToString(f)
}

////

// ImageViewCreateFlags specifies additional parameters of an ImageView
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateFlagBits.html
type ImageViewCreateFlags int32

var imageViewCreateFlagsMapping = common.NewFlagStringMapping[ImageViewCreateFlags]()

func (f ImageViewCreateFlags) Register(str string) {
	imageViewCreateFlagsMapping.Register(f, str)
}

func (f ImageViewCreateFlags) String() string {
	return imageViewCreateFlagsMapping.FlagsToString(f)
}

////

// ImageUsageFlags specifies intended usage of an Image
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html
type ImageUsageFlags int32

var imageUsagesMapping = common.NewFlagStringMapping[ImageUsageFlags]()

func (f ImageUsageFlags) Register(str string) {
	imageUsagesMapping.Register(f, str)
}

func (f ImageUsageFlags) String() string {
	return imageUsagesMapping.FlagsToString(f)
}

////

// MemoryHeapFlags specifies attribute flags for a heap
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryHeapFlagBits.html
type MemoryHeapFlags int32

var memoryHeapFlagsMapping = common.NewFlagStringMapping[MemoryHeapFlags]()

func (f MemoryHeapFlags) Register(str string) {
	memoryHeapFlagsMapping.Register(f, str)
}

func (f MemoryHeapFlags) String() string {
	return memoryHeapFlagsMapping.FlagsToString(f)
}

////

// MemoryPropertyFlags specifies properties for a memory type
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryPropertyFlagBits.html
type MemoryPropertyFlags int32

var memoryPropertiesMapping = common.NewFlagStringMapping[MemoryPropertyFlags]()

func (f MemoryPropertyFlags) Register(str string) {
	memoryPropertiesMapping.Register(f, str)
}

func (f MemoryPropertyFlags) String() string {
	return memoryPropertiesMapping.FlagsToString(f)
}

////

// PipelineCacheCreateFlags specifies the behavior of the PipelineCache
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheCreateFlagBits.html
type PipelineCacheCreateFlags int32

var pipelineCacheCreateFlagsMapping = common.NewFlagStringMapping[PipelineCacheCreateFlags]()

func (f PipelineCacheCreateFlags) Register(str string) {
	pipelineCacheCreateFlagsMapping.Register(f, str)
}

func (f PipelineCacheCreateFlags) String() string {
	return pipelineCacheCreateFlagsMapping.FlagsToString(f)
}

////

// PipelineCreateFlags controls how a Pipeline is created
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html
type PipelineCreateFlags int32

var pipelineCreateFlagsMapping = common.NewFlagStringMapping[PipelineCreateFlags]()

func (f PipelineCreateFlags) Register(str string) {
	pipelineCreateFlagsMapping.Register(f, str)
}

func (f PipelineCreateFlags) String() string {
	return pipelineCreateFlagsMapping.FlagsToString(f)
}

////

// PipelineStageFlags specifies Pipeline stages
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html
type PipelineStageFlags int32

var pipelineStagesMapping = common.NewFlagStringMapping[PipelineStageFlags]()

func (f PipelineStageFlags) Register(str string) {
	pipelineStagesMapping.Register(f, str)
}

func (f PipelineStageFlags) String() string {
	return pipelineStagesMapping.FlagsToString(f)
}

////

// QueryControlFlags specifies constraints on a query
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryControlFlagBits.html
type QueryControlFlags int32

var queryControlFlagsMapping = common.NewFlagStringMapping[QueryControlFlags]()

func (f QueryControlFlags) Register(str string) {
	queryControlFlagsMapping.Register(f, str)
}

func (f QueryControlFlags) String() string {
	return queryControlFlagsMapping.FlagsToString(f)
}

////

// QueryPipelineStatisticFlags specifies queried Pipeline statistics
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html
type QueryPipelineStatisticFlags int32

var queryPipelineStatisticFlagsMapping = common.NewFlagStringMapping[QueryPipelineStatisticFlags]()

func (f QueryPipelineStatisticFlags) Register(str string) {
	queryPipelineStatisticFlagsMapping.Register(f, str)
}

func (f QueryPipelineStatisticFlags) String() string {
	return queryPipelineStatisticFlagsMapping.FlagsToString(f)
}

////

// QueryResultFlags specifies how and when query results are returned
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultFlagBits.html
type QueryResultFlags int32

var queryResultFlagsMapping = common.NewFlagStringMapping[QueryResultFlags]()

func (f QueryResultFlags) Register(str string) {
	queryResultFlagsMapping.Register(f, str)
}

func (f QueryResultFlags) String() string {
	return queryResultFlagsMapping.FlagsToString(f)
}

////

// QueueFlags specifies capabilities of Queue objects in a Queue family
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFlagBits.html
type QueueFlags int32

var queueFlagsMapping = common.NewFlagStringMapping[QueueFlags]()

func (f QueueFlags) Register(str string) {
	queueFlagsMapping.Register(f, str)
}

func (f QueueFlags) String() string {
	return queueFlagsMapping.FlagsToString(f)
}

////

// RenderPassCreateFlags specifies additional properties of a RenderPass
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateFlagBits.html
type RenderPassCreateFlags int32

var renderPassCreateFlagsMapping = common.NewFlagStringMapping[RenderPassCreateFlags]()

func (f RenderPassCreateFlags) Register(str string) {
	renderPassCreateFlagsMapping.Register(f, str)
}

func (f RenderPassCreateFlags) String() string {
	return renderPassCreateFlagsMapping.FlagsToString(f)
}

////

// SampleCountFlags specifies sample counts supported for an Image used for storage operations
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html
type SampleCountFlags int32

var sampleCountsMapping = common.NewFlagStringMapping[SampleCountFlags]()
var sampleCountsToCount = make(map[SampleCountFlags]int)

func (f SampleCountFlags) RegisterSamples(str string, sampleCount int) {
	sampleCountsMapping.Register(f, str)
	sampleCountsToCount[f] = sampleCount
}

func (f SampleCountFlags) String() string {
	return sampleCountsMapping.FlagsToString(f)
}

func (f SampleCountFlags) Count() int {
	var outCount int
	for i := 0; i < 32; i++ {
		checkBit := SampleCountFlags(1 << i)
		if (f & checkBit) != 0 {
			count, hasCount := sampleCountsToCount[checkBit]
			if hasCount && count > outCount {
				outCount = count
			}
		}
	}

	return outCount
}

////

// SamplerCreateFlags specifies additional parameters of a Sampler
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateFlagBits.html
type SamplerCreateFlags int32

var samplerCreateFlagsMapping = common.NewFlagStringMapping[SamplerCreateFlags]()

func (f SamplerCreateFlags) Register(str string) {
	samplerCreateFlagsMapping.Register(f, str)
}

func (f SamplerCreateFlags) String() string {
	return samplerCreateFlagsMapping.FlagsToString(f)
}

////

// PipelineShaderStageCreateFlags controls how a Pipeline shader stage is created
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateFlagBits.html
type PipelineShaderStageCreateFlags int32

var shaderStageCreateMapping = common.NewFlagStringMapping[PipelineShaderStageCreateFlags]()

func (f PipelineShaderStageCreateFlags) Register(str string) {
	shaderStageCreateMapping.Register(f, str)
}

func (f PipelineShaderStageCreateFlags) String() string {
	return shaderStageCreateMapping.FlagsToString(f)
}

////

// ShaderStageFlags specifies a Pipeline stage
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html
type ShaderStageFlags int32

var shaderStagesMapping = common.NewFlagStringMapping[ShaderStageFlags]()

func (f ShaderStageFlags) Register(str string) {
	shaderStagesMapping.Register(f, str)
}

func (f ShaderStageFlags) String() string {
	return shaderStagesMapping.FlagsToString(f)
}

////

// SparseImageFormatFlags specifies additional information about a sparse Image resource
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatFlagBits.html
type SparseImageFormatFlags int32

var sparseImageFormatFlagsMapping = common.NewFlagStringMapping[SparseImageFormatFlags]()

func (f SparseImageFormatFlags) Register(str string) {
	sparseImageFormatFlagsMapping.Register(f, str)
}

func (f SparseImageFormatFlags) String() string {
	return sparseImageFormatFlagsMapping.FlagsToString(f)
}

////

// SparseMemoryBindFlags specifies usage of a sparse memory binding operation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseMemoryBindFlagBits.html
type SparseMemoryBindFlags int32

var sparseMemoryBindFlagsMapping = common.NewFlagStringMapping[SparseMemoryBindFlags]()

func (f SparseMemoryBindFlags) Register(str string) {
	sparseMemoryBindFlagsMapping.Register(f, str)
}

func (f SparseMemoryBindFlags) String() string {
	return sparseMemoryBindFlagsMapping.FlagsToString(f)
}

////

// StencilFaceFlags specifies sets of stencil state for which to update the compare mask
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilFaceFlagBits.html
type StencilFaceFlags int32

var stencilFacesMapping = common.NewFlagStringMapping[StencilFaceFlags]()

func (f StencilFaceFlags) Register(str string) {
	stencilFacesMapping.Register(f, str)
}

func (f StencilFaceFlags) String() string {
	return stencilFacesMapping.FlagsToString(f)
}

////

// SubpassDescriptionFlags specifies usage of a subpass
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionFlagBits.html
type SubpassDescriptionFlags int32

var subPassFlagsMapping = common.NewFlagStringMapping[SubpassDescriptionFlags]()

func (f SubpassDescriptionFlags) Register(str string) {
	subPassFlagsMapping.Register(f, str)
}

func (f SubpassDescriptionFlags) String() string {
	return subPassFlagsMapping.FlagsToString(f)
}
