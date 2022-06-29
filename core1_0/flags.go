package core1_0

import "github.com/CannibalVox/VKng/core/common"

type AccessFlags int32

var accessFlagsMapping = common.NewFlagStringMapping[AccessFlags]()

func (f AccessFlags) Register(str string) {
	accessFlagsMapping.Register(f, str)
}
func (f AccessFlags) String() string {
	return accessFlagsMapping.FlagsToString(f)
}

////

type AttachmentDescriptionFlags int32

var attachmentDescriptionFlagsMapping = common.NewFlagStringMapping[AttachmentDescriptionFlags]()

func (f AttachmentDescriptionFlags) Register(str string) {
	attachmentDescriptionFlagsMapping.Register(f, str)
}

func (f AttachmentDescriptionFlags) String() string {
	return attachmentDescriptionFlagsMapping.FlagsToString(f)
}

////

type BeginInfoFlags int32

var beginInfoFlagsMapping = common.NewFlagStringMapping[BeginInfoFlags]()

func (f BeginInfoFlags) Register(str string) {
	beginInfoFlagsMapping.Register(f, str)
}

func (f BeginInfoFlags) String() string {
	return beginInfoFlagsMapping.FlagsToString(f)
}

////

type BufferCreateFlags int32

var bufferCreateFlagsMapping = common.NewFlagStringMapping[BufferCreateFlags]()

func (f BufferCreateFlags) Register(str string) {
	bufferCreateFlagsMapping.Register(f, str)
}

func (f BufferCreateFlags) String() string {
	return bufferCreateFlagsMapping.FlagsToString(f)
}

////

type BufferUsages int32

var bufferUsagesMapping = common.NewFlagStringMapping[BufferUsages]()

func (f BufferUsages) Register(str string) {
	bufferUsagesMapping.Register(f, str)
}

func (f BufferUsages) String() string {
	return bufferUsagesMapping.FlagsToString(f)
}

////

type CommandBufferResetFlags int32

var commandBufferResetFlagsMapping = common.NewFlagStringMapping[CommandBufferResetFlags]()

func (f CommandBufferResetFlags) Register(str string) {
	commandBufferResetFlagsMapping.Register(f, str)
}

func (f CommandBufferResetFlags) String() string {
	return commandBufferResetFlagsMapping.FlagsToString(f)
}

////

type CommandPoolCreateFlags int32

var commandPoolCreateFlagsMapping = common.NewFlagStringMapping[CommandPoolCreateFlags]()

func (f CommandPoolCreateFlags) Register(str string) {
	commandPoolCreateFlagsMapping.Register(f, str)
}

func (f CommandPoolCreateFlags) String() string {
	return commandPoolCreateFlagsMapping.FlagsToString(f)
}

////

type CommandPoolResetFlags int32

var commandPoolResetFlagsMapping = common.NewFlagStringMapping[CommandPoolResetFlags]()

func (f CommandPoolResetFlags) Register(str string) {
	commandPoolResetFlagsMapping.Register(f, str)
}

func (f CommandPoolResetFlags) String() string {
	return commandPoolResetFlagsMapping.FlagsToString(f)
}

////

type CompositeAlphaModes int32

var compositeAlphaModesMapping = common.NewFlagStringMapping[CompositeAlphaModes]()

func (f CompositeAlphaModes) Register(str string) {
	compositeAlphaModesMapping.Register(f, str)
}

func (f CompositeAlphaModes) String() string {
	return compositeAlphaModesMapping.FlagsToString(f)
}

////

type CullModes int32

var cullModesMapping = common.NewFlagStringMapping[CullModes]()

func (f CullModes) Register(str string) {
	cullModesMapping.Register(f, str)
}

func (f CullModes) String() string {
	return cullModesMapping.FlagsToString(f)
}

////

type DependencyFlags int32

var dependencyFlagsMapping = common.NewFlagStringMapping[DependencyFlags]()

func (f DependencyFlags) Register(str string) {
	dependencyFlagsMapping.Register(f, str)
}

func (f DependencyFlags) String() string {
	return dependencyFlagsMapping.FlagsToString(f)
}

////

type DescriptorPoolCreateFlags int32

var descriptorPoolCreateFlagsMapping = common.NewFlagStringMapping[DescriptorPoolCreateFlags]()

func (f DescriptorPoolCreateFlags) Register(str string) {
	descriptorPoolCreateFlagsMapping.Register(f, str)
}

func (f DescriptorPoolCreateFlags) String() string {
	return descriptorPoolCreateFlagsMapping.FlagsToString(f)
}

////

type DescriptorPoolResetFlags int32

var descriptorPoolResetFlagsMapping = common.NewFlagStringMapping[DescriptorPoolResetFlags]()

func (f DescriptorPoolResetFlags) Register(str string) {
	descriptorPoolResetFlagsMapping.Register(f, str)
}

func (f DescriptorPoolResetFlags) String() string {
	return descriptorPoolResetFlagsMapping.FlagsToString(f)
}

////

type DescriptorSetLayoutCreateFlags int32

var descriptorSetLayoutCreateFlagsMapping = common.NewFlagStringMapping[DescriptorSetLayoutCreateFlags]()

func (f DescriptorSetLayoutCreateFlags) Register(str string) {
	descriptorSetLayoutCreateFlagsMapping.Register(f, str)
}

func (f DescriptorSetLayoutCreateFlags) String() string {
	return descriptorSetLayoutCreateFlagsMapping.FlagsToString(f)
}

////

type DeviceQueueCreateFlags int32

var deviceQueueCreateFlagsMapping = common.NewFlagStringMapping[DeviceQueueCreateFlags]()

func (f DeviceQueueCreateFlags) Register(str string) {
	deviceQueueCreateFlagsMapping.Register(f, str)
}

func (f DeviceQueueCreateFlags) String() string {
	return deviceQueueCreateFlagsMapping.FlagsToString(f)
}

////

type EventCreateFlags int32

var eventCreateFlagsMapping = common.NewFlagStringMapping[EventCreateFlags]()

func (f EventCreateFlags) Register(str string) {
	eventCreateFlagsMapping.Register(f, str)
}

func (f EventCreateFlags) String() string {
	return eventCreateFlagsMapping.FlagsToString(f)
}

////

type FenceCreateFlags int32

var fenceCreateFlagsMapping = common.NewFlagStringMapping[FenceCreateFlags]()

func (f FenceCreateFlags) Register(str string) {
	fenceCreateFlagsMapping.Register(f, str)
}

func (f FenceCreateFlags) String() string {
	return fenceCreateFlagsMapping.FlagsToString(f)
}

////

type FormatFeatures int32

var formatFeaturesMapping = common.NewFlagStringMapping[FormatFeatures]()

func (f FormatFeatures) Register(str string) {
	formatFeaturesMapping.Register(f, str)
}

func (f FormatFeatures) String() string {
	return formatFeaturesMapping.FlagsToString(f)
}

////

type FramebufferCreateFlags int32

var framebufferCreateFlagsMapping = common.NewFlagStringMapping[FramebufferCreateFlags]()

func (f FramebufferCreateFlags) Register(str string) {
	framebufferCreateFlagsMapping.Register(f, str)
}

func (f FramebufferCreateFlags) String() string {
	return framebufferCreateFlagsMapping.FlagsToString(f)
}

////

type ImageAspectFlags int32

var imageAspectFlagsMapping = common.NewFlagStringMapping[ImageAspectFlags]()

func (f ImageAspectFlags) Register(str string) {
	imageAspectFlagsMapping.Register(f, str)
}

func (f ImageAspectFlags) String() string {
	return imageAspectFlagsMapping.FlagsToString(f)
}

////

type ImageCreateFlags int32

var imageCreateFlagsMapping = common.NewFlagStringMapping[ImageCreateFlags]()

func (f ImageCreateFlags) Register(str string) {
	imageCreateFlagsMapping.Register(f, str)
}

func (f ImageCreateFlags) String() string {
	return imageCreateFlagsMapping.FlagsToString(f)
}

////

type ImageViewCreateFlags int32

var imageViewCreateFlagsMapping = common.NewFlagStringMapping[ImageViewCreateFlags]()

func (f ImageViewCreateFlags) Register(str string) {
	imageViewCreateFlagsMapping.Register(f, str)
}

func (f ImageViewCreateFlags) String() string {
	return imageViewCreateFlagsMapping.FlagsToString(f)
}

////

type ImageUsages int32

var imageUsagesMapping = common.NewFlagStringMapping[ImageUsages]()

func (f ImageUsages) Register(str string) {
	imageUsagesMapping.Register(f, str)
}

func (f ImageUsages) String() string {
	return imageUsagesMapping.FlagsToString(f)
}

////

type MemoryHeapFlags int32

var memoryHeapFlagsMapping = common.NewFlagStringMapping[MemoryHeapFlags]()

func (f MemoryHeapFlags) Register(str string) {
	memoryHeapFlagsMapping.Register(f, str)
}

func (f MemoryHeapFlags) String() string {
	return memoryHeapFlagsMapping.FlagsToString(f)
}

////

type MemoryProperties int32

var memoryPropertiesMapping = common.NewFlagStringMapping[MemoryProperties]()

func (f MemoryProperties) Register(str string) {
	memoryPropertiesMapping.Register(f, str)
}

func (f MemoryProperties) String() string {
	return memoryPropertiesMapping.FlagsToString(f)
}

////

type PipelineCacheCreateFlags int32

var pipelineCacheCreateFlagsMapping = common.NewFlagStringMapping[PipelineCacheCreateFlags]()

func (f PipelineCacheCreateFlags) Register(str string) {
	pipelineCacheCreateFlagsMapping.Register(f, str)
}

func (f PipelineCacheCreateFlags) String() string {
	return pipelineCacheCreateFlagsMapping.FlagsToString(f)
}

////

type PipelineCreateFlags int32

var pipelineCreateFlagsMapping = common.NewFlagStringMapping[PipelineCreateFlags]()

func (f PipelineCreateFlags) Register(str string) {
	pipelineCreateFlagsMapping.Register(f, str)
}

func (f PipelineCreateFlags) String() string {
	return pipelineCreateFlagsMapping.FlagsToString(f)
}

////

type PipelineStages int32

var pipelineStagesMapping = common.NewFlagStringMapping[PipelineStages]()

func (f PipelineStages) Register(str string) {
	pipelineStagesMapping.Register(f, str)
}

func (f PipelineStages) String() string {
	return pipelineStagesMapping.FlagsToString(f)
}

////

type PipelineStatistics int32

var pipelineStatisticsMapping = common.NewFlagStringMapping[PipelineStatistics]()

func (f PipelineStatistics) Register(str string) {
	pipelineStatisticsMapping.Register(f, str)
}

func (f PipelineStatistics) String() string {
	return pipelineStatisticsMapping.FlagsToString(f)
}

////

type QueryControlFlags int32

var queryControlFlagsMapping = common.NewFlagStringMapping[QueryControlFlags]()

func (f QueryControlFlags) Register(str string) {
	queryControlFlagsMapping.Register(f, str)
}

func (f QueryControlFlags) String() string {
	return queryControlFlagsMapping.FlagsToString(f)
}

////

type QueryPipelineStatisticFlags int32

var queryPipelineStatisticFlagsMapping = common.NewFlagStringMapping[QueryPipelineStatisticFlags]()

func (f QueryPipelineStatisticFlags) Register(str string) {
	queryPipelineStatisticFlagsMapping.Register(f, str)
}

func (f QueryPipelineStatisticFlags) String() string {
	return queryPipelineStatisticFlagsMapping.FlagsToString(f)
}

////

type QueryResultFlags int32

var queryResultFlagsMapping = common.NewFlagStringMapping[QueryResultFlags]()

func (f QueryResultFlags) Register(str string) {
	queryResultFlagsMapping.Register(f, str)
}

func (f QueryResultFlags) String() string {
	return queryResultFlagsMapping.FlagsToString(f)
}

////

type QueueFlags int32

var queueFlagsMapping = common.NewFlagStringMapping[QueueFlags]()

func (f QueueFlags) Register(str string) {
	queueFlagsMapping.Register(f, str)
}

func (f QueueFlags) String() string {
	return queueFlagsMapping.FlagsToString(f)
}

////

type RenderPassCreateFlags int32

var renderPassCreateFlagsMapping = common.NewFlagStringMapping[RenderPassCreateFlags]()

func (f RenderPassCreateFlags) Register(str string) {
	renderPassCreateFlagsMapping.Register(f, str)
}

func (f RenderPassCreateFlags) String() string {
	return renderPassCreateFlagsMapping.FlagsToString(f)
}

////

type SampleCounts int32

var sampleCountsMapping = common.NewFlagStringMapping[SampleCounts]()
var sampleCountsToCount = make(map[SampleCounts]int)

func (f SampleCounts) RegisterSamples(str string, sampleCount int) {
	sampleCountsMapping.Register(f, str)
	sampleCountsToCount[f] = sampleCount
}

func (f SampleCounts) String() string {
	return sampleCountsMapping.FlagsToString(f)
}

func (f SampleCounts) Count() int {
	var outCount int
	for i := 0; i < 32; i++ {
		checkBit := SampleCounts(1 << i)
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

type SamplerCreateFlags int32

var samplerCreateFlagsMapping = common.NewFlagStringMapping[SamplerCreateFlags]()

func (f SamplerCreateFlags) Register(str string) {
	samplerCreateFlagsMapping.Register(f, str)
}

func (f SamplerCreateFlags) String() string {
	return samplerCreateFlagsMapping.FlagsToString(f)
}

////

type ShaderStageCreateFlags int32

var shaderStageCreateMapping = common.NewFlagStringMapping[ShaderStageCreateFlags]()

func (f ShaderStageCreateFlags) Register(str string) {
	shaderStageCreateMapping.Register(f, str)
}

func (f ShaderStageCreateFlags) String() string {
	return shaderStageCreateMapping.FlagsToString(f)
}

////

type ShaderStages int32

var shaderStagesMapping = common.NewFlagStringMapping[ShaderStages]()

func (f ShaderStages) Register(str string) {
	shaderStagesMapping.Register(f, str)
}

func (f ShaderStages) String() string {
	return shaderStagesMapping.FlagsToString(f)
}

////

type SparseImageFormatFlags int32

var sparseImageFormatFlagsMapping = common.NewFlagStringMapping[SparseImageFormatFlags]()

func (f SparseImageFormatFlags) Register(str string) {
	sparseImageFormatFlagsMapping.Register(f, str)
}

func (f SparseImageFormatFlags) String() string {
	return sparseImageFormatFlagsMapping.FlagsToString(f)
}

////

type SparseMemoryBindFlags int32

var sparseMemoryBindFlagsMapping = common.NewFlagStringMapping[SparseMemoryBindFlags]()

func (f SparseMemoryBindFlags) Register(str string) {
	sparseMemoryBindFlagsMapping.Register(f, str)
}

func (f SparseMemoryBindFlags) String() string {
	return sparseMemoryBindFlagsMapping.FlagsToString(f)
}

////

type StencilFaces int32

var stencilFacesMapping = common.NewFlagStringMapping[StencilFaces]()

func (f StencilFaces) Register(str string) {
	stencilFacesMapping.Register(f, str)
}

func (f StencilFaces) String() string {
	return stencilFacesMapping.FlagsToString(f)
}

////

type SubPassDescriptionFlags int32

var subPassFlagsMapping = common.NewFlagStringMapping[SubPassDescriptionFlags]()

func (f SubPassDescriptionFlags) Register(str string) {
	subPassFlagsMapping.Register(f, str)
}

func (f SubPassDescriptionFlags) String() string {
	return subPassFlagsMapping.FlagsToString(f)
}
