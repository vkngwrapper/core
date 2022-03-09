package common

import "github.com/cockroachdb/errors"

type AttachmentLoadOp int32

var attachmentLoadOpMapping = make(map[AttachmentLoadOp]string)

func (e AttachmentLoadOp) Register(str string) {
	attachmentLoadOpMapping[e] = str
}

func (e AttachmentLoadOp) String() string {
	return attachmentLoadOpMapping[e]
}

////

type AttachmentStoreOp int32

var attachmentStoreOpMapping = make(map[AttachmentStoreOp]string)

func (e AttachmentStoreOp) Register(str string) {
	attachmentStoreOpMapping[e] = str
}

func (e AttachmentStoreOp) String() string {
	return attachmentStoreOpMapping[e]
}

////

type BlendFactor int32

var blendFactorMapping = make(map[BlendFactor]string)

func (e BlendFactor) Register(str string) {
	blendFactorMapping[e] = str
}

func (e BlendFactor) String() string {
	return blendFactorMapping[e]
}

////

type BlendOp int32

var blendOpMapping = make(map[BlendOp]string)

func (e BlendOp) Register(str string) {
	blendOpMapping[e] = str
}

func (e BlendOp) String() string {
	return blendOpMapping[e]
}

////

type BorderColor int32

var borderColorMapping = make(map[BorderColor]string)

func (e BorderColor) Register(str string) {
	borderColorMapping[e] = str
}

func (e BorderColor) String() string {
	return borderColorMapping[e]
}

////

type ColorSpace int32

var colorSpaceMapping = make(map[ColorSpace]string)

func (e ColorSpace) Register(str string) {
	colorSpaceMapping[e] = str
}

func (e ColorSpace) String() string {
	return colorSpaceMapping[e]
}

////

type CommandBufferLevel int32

var commandBufferLevelMapping = make(map[CommandBufferLevel]string)

func (e CommandBufferLevel) Register(str string) {
	commandBufferLevelMapping[e] = str
}

func (e CommandBufferLevel) String() string {
	return commandBufferLevelMapping[e]
}

////

type CompareOp int32

var compareOpMapping = make(map[CompareOp]string)

func (e CompareOp) Register(str string) {
	compareOpMapping[e] = str
}

func (e CompareOp) String() string {
	return compareOpMapping[e]
}

////

type ComponentSwizzle int32

var componentSwizzleMapping = make(map[ComponentSwizzle]string)

func (e ComponentSwizzle) Register(str string) {
	componentSwizzleMapping[e] = str
}

func (e ComponentSwizzle) String() string {
	return componentSwizzleMapping[e]
}

////

type DataFormat int32

var dataFormatMapping = make(map[DataFormat]string)

func (e DataFormat) Register(str string) {
	dataFormatMapping[e] = str
}

func (e DataFormat) String() string {
	return dataFormatMapping[e]
}

////

type DescriptorType int32

var descriptorTypeMapping = make(map[DescriptorType]string)

func (e DescriptorType) Register(str string) {
	descriptorTypeMapping[e] = str
}

func (e DescriptorType) String() string {
	return descriptorTypeMapping[e]
}

////

type DynamicState int32

var dynamicStateMapping = make(map[DynamicState]string)

func (e DynamicState) Register(str string) {
	dynamicStateMapping[e] = str
}

func (e DynamicState) String() string {
	return dynamicStateMapping[e]
}

////

type Filter int32

var filterMapping = make(map[Filter]string)

func (e Filter) Register(str string) {
	filterMapping[e] = str
}

func (e Filter) String() string {
	return filterMapping[e]
}

////

type FrontFace int32

var frontFaceMapping = make(map[FrontFace]string)

func (e FrontFace) Register(str string) {
	frontFaceMapping[e] = str
}

func (e FrontFace) String() string {
	return frontFaceMapping[e]
}

////

type ImageLayout int32

var imageLayoutMapping = make(map[ImageLayout]string)

func (e ImageLayout) Register(str string) {
	imageLayoutMapping[e] = str
}

func (e ImageLayout) String() string {
	return imageLayoutMapping[e]
}

////

type ImageTiling int32

var imageTilingMapping = make(map[ImageTiling]string)

func (e ImageTiling) Register(str string) {
	imageTilingMapping[e] = str
}

func (e ImageTiling) String() string {
	return imageTilingMapping[e]
}

////

type ImageType int32

var imageTypeMapping = make(map[ImageType]string)

func (e ImageType) Register(str string) {
	imageTypeMapping[e] = str
}

func (e ImageType) String() string {
	return imageTypeMapping[e]
}

////

type ImageViewType int32

var imageViewTypeMapping = make(map[ImageViewType]string)

func (e ImageViewType) Register(str string) {
	imageViewTypeMapping[e] = str
}

func (e ImageViewType) String() string {
	return imageViewTypeMapping[e]
}

////

type IndexType int32

var indexTypeMapping = make(map[IndexType]string)

func (e IndexType) Register(str string) {
	indexTypeMapping[e] = str
}

func (e IndexType) String() string {
	return indexTypeMapping[e]
}

////

type InputRate int32

var inputRateMapping = make(map[InputRate]string)

func (e InputRate) Register(str string) {
	inputRateMapping[e] = str
}

func (e InputRate) String() string {
	return inputRateMapping[e]
}

////

type InternalAllocationType int32

var internalAllocationTypeMapping = make(map[InternalAllocationType]string)

func (e InternalAllocationType) Register(str string) {
	internalAllocationTypeMapping[e] = str
}

func (e InternalAllocationType) String() string {
	return internalAllocationTypeMapping[e]
}

////

type LogicOp int32

var logicOpMapping = make(map[LogicOp]string)

func (e LogicOp) Register(str string) {
	logicOpMapping[e] = str
}

func (e LogicOp) String() string {
	return logicOpMapping[e]
}

////

type MipmapMode int32

var mipmapModeMapping = make(map[MipmapMode]string)

func (e MipmapMode) Register(str string) {
	mipmapModeMapping[e] = str
}

func (e MipmapMode) String() string {
	return mipmapModeMapping[e]
}

////

type ObjectType int32

var objectTypeMapping = make(map[ObjectType]string)

func (e ObjectType) Register(str string) {
	objectTypeMapping[e] = str
}

func (e ObjectType) String() string {
	return objectTypeMapping[e]
}

////

type PhysicalDeviceType int32

var physicalDeviceTypeMapping = make(map[PhysicalDeviceType]string)

func (e PhysicalDeviceType) Register(str string) {
	physicalDeviceTypeMapping[e] = str
}

func (e PhysicalDeviceType) String() string {
	return physicalDeviceTypeMapping[e]
}

////

type PipelineBindPoint int32

var pipelineBindPointMapping = make(map[PipelineBindPoint]string)

func (e PipelineBindPoint) Register(str string) {
	pipelineBindPointMapping[e] = str
}

func (e PipelineBindPoint) String() string {
	return pipelineBindPointMapping[e]
}

////

type PipelineCacheHeaderVersion int32

var pipelineCacheHeaderVersionMapping = make(map[PipelineCacheHeaderVersion]string)

func (e PipelineCacheHeaderVersion) Register(str string) {
	pipelineCacheHeaderVersionMapping[e] = str
}

func (e PipelineCacheHeaderVersion) String() string {
	return pipelineCacheHeaderVersionMapping[e]
}

////

type PolygonMode int32

var polygonModeMapping = make(map[PolygonMode]string)

func (e PolygonMode) Register(str string) {
	polygonModeMapping[e] = str
}

func (e PolygonMode) String() string {
	return polygonModeMapping[e]
}

////

type PrimitiveTopology int32

var primitiveTopologyMapping = make(map[PrimitiveTopology]string)

func (e PrimitiveTopology) Register(str string) {
	primitiveTopologyMapping[e] = str
}

func (e PrimitiveTopology) String() string {
	return primitiveTopologyMapping[e]
}

////

type QueryType int32

var queryTypeMapping = make(map[QueryType]string)

func (e QueryType) Register(str string) {
	queryTypeMapping[e] = str
}

func (e QueryType) String() string {
	return queryTypeMapping[e]
}

////

type SamplerAddressMode int32

var samplerAddressModeMapping = make(map[SamplerAddressMode]string)

func (e SamplerAddressMode) Register(str string) {
	samplerAddressModeMapping[e] = str
}

func (e SamplerAddressMode) String() string {
	return samplerAddressModeMapping[e]
}

////

type SharingMode int32

var sharingModeMapping = make(map[SharingMode]string)

func (e SharingMode) Register(str string) {
	sharingModeMapping[e] = str
}

func (e SharingMode) String() string {
	return sharingModeMapping[e]
}

////

type StencilOp int32

var stencilOpMapping = make(map[StencilOp]string)

func (e StencilOp) Register(str string) {
	stencilOpMapping[e] = str
}

func (e StencilOp) String() string {
	return stencilOpMapping[e]
}

////

type SubpassContents int32

var subpassContentsMapping = make(map[SubpassContents]string)

func (e SubpassContents) Register(str string) {
	subpassContentsMapping[e] = str
}

func (e SubpassContents) String() string {
	return subpassContentsMapping[e]
}

////

type SystemAllocationScope int32

var systemAllocationScopeMapping = make(map[SystemAllocationScope]string)

func (e SystemAllocationScope) Register(str string) {
	systemAllocationScopeMapping[e] = str
}

func (e SystemAllocationScope) String() string {
	return systemAllocationScopeMapping[e]
}

////

type VkResult int

var vkResultMapping = make(map[VkResult]string)

func (e VkResult) Register(str string) {
	vkResultMapping[e] = str
}

func (e VkResult) String() string {
	return vkResultMapping[e]
}

func (e VkResult) ToError() error {
	if e >= 0 {
		// All VKError* are <0
		return nil
	}

	return errors.WithStack(&VkResultError{e})
}
