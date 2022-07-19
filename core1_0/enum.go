package core1_0

// AttachmentLoadOp specifies how contents of an attachment are treated at the beginning
// of a subpass
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentLoadOp.html
type AttachmentLoadOp int32

var attachmentLoadOpMapping = make(map[AttachmentLoadOp]string)

func (e AttachmentLoadOp) Register(str string) {
	attachmentLoadOpMapping[e] = str
}

func (e AttachmentLoadOp) String() string {
	return attachmentLoadOpMapping[e]
}

////

// AttachmentStoreOp specifies how contents of an attachment are treated at the end of a subpass
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html
type AttachmentStoreOp int32

var attachmentStoreOpMapping = make(map[AttachmentStoreOp]string)

func (e AttachmentStoreOp) Register(str string) {
	attachmentStoreOpMapping[e] = str
}

func (e AttachmentStoreOp) String() string {
	return attachmentStoreOpMapping[e]
}

////

// BlendFactor specifies Framebuffer blending factors
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html
type BlendFactor int32

var blendFactorMapping = make(map[BlendFactor]string)

func (e BlendFactor) Register(str string) {
	blendFactorMapping[e] = str
}

func (e BlendFactor) String() string {
	return blendFactorMapping[e]
}

////

// BlendOp specifies Framebuffer blending operations
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOp.html
type BlendOp int32

var blendOpMapping = make(map[BlendOp]string)

func (e BlendOp) Register(str string) {
	blendOpMapping[e] = str
}

func (e BlendOp) String() string {
	return blendOpMapping[e]
}

////

// BorderColor specifies border color used for texture lookups
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBorderColor.html
type BorderColor int32

var borderColorMapping = make(map[BorderColor]string)

func (e BorderColor) Register(str string) {
	borderColorMapping[e] = str
}

func (e BorderColor) String() string {
	return borderColorMapping[e]
}

////

// CommandBufferLevel specifies a CommandBuffer level
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferLevel.html
type CommandBufferLevel int32

var commandBufferLevelMapping = make(map[CommandBufferLevel]string)

func (e CommandBufferLevel) Register(str string) {
	commandBufferLevelMapping[e] = str
}

func (e CommandBufferLevel) String() string {
	return commandBufferLevelMapping[e]
}

////

// CompareOp is comparison operators for depth, stencil, and Sampler operations
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html
type CompareOp int32

var compareOpMapping = make(map[CompareOp]string)

func (e CompareOp) Register(str string) {
	compareOpMapping[e] = str
}

func (e CompareOp) String() string {
	return compareOpMapping[e]
}

////

// ComponentSwizzle specifies how a component is swizzled
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html
type ComponentSwizzle int32

var componentSwizzleMapping = make(map[ComponentSwizzle]string)

func (e ComponentSwizzle) Register(str string) {
	componentSwizzleMapping[e] = str
}

func (e ComponentSwizzle) String() string {
	return componentSwizzleMapping[e]
}

////

// Format specifies available image formats
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html
type Format int32

var FormatMapping = make(map[Format]string)

func (e Format) Register(str string) {
	FormatMapping[e] = str
}

func (e Format) String() string {
	return FormatMapping[e]
}

////

// DescriptorType specifies the type of a descriptor in a DescriptorSet
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html
type DescriptorType int32

var descriptorTypeMapping = make(map[DescriptorType]string)

func (e DescriptorType) Register(str string) {
	descriptorTypeMapping[e] = str
}

func (e DescriptorType) String() string {
	return descriptorTypeMapping[e]
}

////

// DynamicState indicates which dynamic state is taken from dynamic state commands
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html
type DynamicState int32

var dynamicStateMapping = make(map[DynamicState]string)

func (e DynamicState) Register(str string) {
	dynamicStateMapping[e] = str
}

func (e DynamicState) String() string {
	return dynamicStateMapping[e]
}

////

// Filter specifies filters to use for texture lookups
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilter.html
type Filter int32

var filterMapping = make(map[Filter]string)

func (e Filter) Register(str string) {
	filterMapping[e] = str
}

func (e Filter) String() string {
	return filterMapping[e]
}

////

// FrontFace interprets polygon front-facing orientation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrontFace.html
type FrontFace int32

var frontFaceMapping = make(map[FrontFace]string)

func (e FrontFace) Register(str string) {
	frontFaceMapping[e] = str
}

func (e FrontFace) String() string {
	return frontFaceMapping[e]
}

////

// ImageLayout represents the layout of Image and image subresources
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
type ImageLayout int32

var imageLayoutMapping = make(map[ImageLayout]string)

func (e ImageLayout) Register(str string) {
	imageLayoutMapping[e] = str
}

func (e ImageLayout) String() string {
	return imageLayoutMapping[e]
}

////

// ImageTiling specifies the tiling arrangement of data in an Image
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html
type ImageTiling int32

var imageTilingMapping = make(map[ImageTiling]string)

func (e ImageTiling) Register(str string) {
	imageTilingMapping[e] = str
}

func (e ImageTiling) String() string {
	return imageTilingMapping[e]
}

////

// ImageType specifies the type of an Image object
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html
type ImageType int32

var imageTypeMapping = make(map[ImageType]string)

func (e ImageType) Register(str string) {
	imageTypeMapping[e] = str
}

func (e ImageType) String() string {
	return imageTypeMapping[e]
}

////

// ImageViewType represents the type of ImageView objects that can be created
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html
type ImageViewType int32

var imageViewTypeMapping = make(map[ImageViewType]string)

func (e ImageViewType) Register(str string) {
	imageViewTypeMapping[e] = str
}

func (e ImageViewType) String() string {
	return imageViewTypeMapping[e]
}

////

// IndexType represents the type of index buffer indices
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html
type IndexType int32

var indexTypeMapping = make(map[IndexType]string)

func (e IndexType) Register(str string) {
	indexTypeMapping[e] = str
}

func (e IndexType) String() string {
	return indexTypeMapping[e]
}

////

// VertexInputRate specifies the rate at which vertex attributes are pulled from buffers
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputRate.html
type VertexInputRate int32

var inputRateMapping = make(map[VertexInputRate]string)

func (e VertexInputRate) Register(str string) {
	inputRateMapping[e] = str
}

func (e VertexInputRate) String() string {
	return inputRateMapping[e]
}

////

// LogicOp represents Framebuffer logical operations
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html
type LogicOp int32

var logicOpMapping = make(map[LogicOp]string)

func (e LogicOp) Register(str string) {
	logicOpMapping[e] = str
}

func (e LogicOp) String() string {
	return logicOpMapping[e]
}

////

// SamplerMipmapMode specifies the mipmap mode used for texture lookups
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerMipmapMode.html
type SamplerMipmapMode int32

var mipmapModeMapping = make(map[SamplerMipmapMode]string)

func (e SamplerMipmapMode) Register(str string) {
	mipmapModeMapping[e] = str
}

func (e SamplerMipmapMode) String() string {
	return mipmapModeMapping[e]
}

////

// ObjectType specifies an enumeration to track object handle types
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
type ObjectType int32

var objectTypeMapping = make(map[ObjectType]string)

func (e ObjectType) Register(str string) {
	objectTypeMapping[e] = str
}

func (e ObjectType) String() string {
	return objectTypeMapping[e]
}

////

// PhysicalDeviceType represents supported PhysicalDevice types
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceType.html
type PhysicalDeviceType int32

var physicalDeviceTypeMapping = make(map[PhysicalDeviceType]string)

func (e PhysicalDeviceType) Register(str string) {
	physicalDeviceTypeMapping[e] = str
}

func (e PhysicalDeviceType) String() string {
	return physicalDeviceTypeMapping[e]
}

////

// PipelineBindPoint specifies the bind point of a Pipeline object to a CommandBuffer
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html
type PipelineBindPoint int32

var pipelineBindPointMapping = make(map[PipelineBindPoint]string)

func (e PipelineBindPoint) Register(str string) {
	pipelineBindPointMapping[e] = str
}

func (e PipelineBindPoint) String() string {
	return pipelineBindPointMapping[e]
}

////

// PipelineCacheHeaderVersion encodes the PipelineCache version
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheHeaderVersion.html
type PipelineCacheHeaderVersion int32

var pipelineCacheHeaderVersionMapping = make(map[PipelineCacheHeaderVersion]string)

func (e PipelineCacheHeaderVersion) Register(str string) {
	pipelineCacheHeaderVersionMapping[e] = str
}

func (e PipelineCacheHeaderVersion) String() string {
	return pipelineCacheHeaderVersionMapping[e]
}

////

// PolygonMode controls polygon rasterization mode
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPolygonMode.html
type PolygonMode int32

var polygonModeMapping = make(map[PolygonMode]string)

func (e PolygonMode) Register(str string) {
	polygonModeMapping[e] = str
}

func (e PolygonMode) String() string {
	return polygonModeMapping[e]
}

////

// PrimitiveTopology represents supported primitive topologies
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrimitiveTopology.html
type PrimitiveTopology int32

var primitiveTopologyMapping = make(map[PrimitiveTopology]string)

func (e PrimitiveTopology) Register(str string) {
	primitiveTopologyMapping[e] = str
}

func (e PrimitiveTopology) String() string {
	return primitiveTopologyMapping[e]
}

////

// QueryType specifies the type of queries managed by a QueryPool
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html
type QueryType int32

var queryTypeMapping = make(map[QueryType]string)

func (e QueryType) Register(str string) {
	queryTypeMapping[e] = str
}

func (e QueryType) String() string {
	return queryTypeMapping[e]
}

////

// SamplerAddressMode specifies behavior of sampling with texture coordinates outside an Image
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html
type SamplerAddressMode int32

var samplerAddressModeMapping = make(map[SamplerAddressMode]string)

func (e SamplerAddressMode) Register(str string) {
	samplerAddressModeMapping[e] = str
}

func (e SamplerAddressMode) String() string {
	return samplerAddressModeMapping[e]
}

////

// SharingMode represents Buffer and Image sharing modes
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSharingMode.html
type SharingMode int32

var sharingModeMapping = make(map[SharingMode]string)

func (e SharingMode) Register(str string) {
	sharingModeMapping[e] = str
}

func (e SharingMode) String() string {
	return sharingModeMapping[e]
}

////

// StencilOp represents the stencil comparison function
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html
type StencilOp int32

var stencilOpMapping = make(map[StencilOp]string)

func (e StencilOp) Register(str string) {
	stencilOpMapping[e] = str
}

func (e StencilOp) String() string {
	return stencilOpMapping[e]
}

////

// SubpassContents specifies how commands in the first subpass of a RenderPass are provided
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassContents.html
type SubpassContents int32

var subpassContentsMapping = make(map[SubpassContents]string)

func (e SubpassContents) Register(str string) {
	subpassContentsMapping[e] = str
}

func (e SubpassContents) String() string {
	return subpassContentsMapping[e]
}
