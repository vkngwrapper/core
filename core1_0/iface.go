package core1_0

import (
	"time"
	"unsafe"

	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
)

//go:generate mockgen -source ./iface.go -destination ../mocks/mocks1_0/mocks.go -package mocks1_0

type GlobalDriver interface {
	Loader() loader.Loader

	BuildInstanceDriver(instance core.Instance) (CoreInstanceDriver, error)

	// AvailableExtensions returns all of the instance extensions available on this Loader,
	// in the form of a map of extension name to ExtensionProperties
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateInstanceExtensionProperties.html
	AvailableExtensions() (map[string]*ExtensionProperties, common.VkResult, error)
	// AvailableExtensionsForLayer returns all of the layer extensions available on this Loader
	// for the requested layer, in the form of a map of extension name to ExtensionProperties
	//
	// layerName - a string naming the layer to retrieve extensions from
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateInstanceExtensionProperties.html
	AvailableExtensionsForLayer(layerName string) (map[string]*ExtensionProperties, common.VkResult, error)
	// AvailableLayers returns all of the layers available on this Loader, in the form of a
	// map of layer name to LayerProperties
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateInstanceLayerProperties.html
	AvailableLayers() (map[string]*LayerProperties, common.VkResult, error)
	// CreateInstance creates a new Vulkan Instance
	//
	// allocationCallbacks - controls host memory allocation
	//
	// options - Controls creation of the Instance
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateInstance.html
	CreateInstance(allocationCallbacks *loader.AllocationCallbacks, options InstanceCreateInfo) (core.Instance, common.VkResult, error)
}

type CoreInstanceDriver interface {
	GlobalDriver
	Instance() core.Instance
	BuildDeviceDriver(device core.Device) (CoreDeviceDriver, error)

	// DestroyInstance destroys the Instance object underlying this driver and the underlying
	// structures. **Warning** after destruction, the object will continue to exist, but the
	// Vulkan object handle that backs it will be invalid. Do not call further methods on this driver
	// or using the Instance.
	//
	// instance - The Vulkan instance to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyInstance.html
	DestroyInstance(callbacks *loader.AllocationCallbacks)

	// CreateDevice creates a new logical device as a connection to a PhysicalDevice
	//
	// physicalDevice - The PhysicalDevice to connect to
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// options - Parameters affecting the creation of the Device
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateDevice.html
	CreateDevice(physicalDevice core.PhysicalDevice, allocationCallbacks *loader.AllocationCallbacks, options DeviceCreateInfo) (core.Device, common.VkResult, error)

	// EnumeratePhysicalDevices enumerates the physical devices accessible to an Instance
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDevices.html
	EnumeratePhysicalDevices() ([]core.PhysicalDevice, common.VkResult, error)

	// GetPhysicalDeviceQueueFamilyProperties reports properties of the queues of a PhysicalDevice
	//
	// physicalDevice - the PhysicalDevice to retrieve queues for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties.html
	GetPhysicalDeviceQueueFamilyProperties(physicalDevice core.PhysicalDevice) []*QueueFamilyProperties

	// GetPhysicalDeviceProperties returns properties of a PhysicalDevice
	//
	// physicalDevice - the PhysicalDevice to retrieve properties for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties.html
	GetPhysicalDeviceProperties(physicalDevice core.PhysicalDevice) (*PhysicalDeviceProperties, error)
	// GetPhysicalDeviceFeatures reports capabilities of a PhysicalDevice
	//
	// physicalDevice - The PhysicalDevice to retrieve features for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures.html
	GetPhysicalDeviceFeatures(physicalDevice core.PhysicalDevice) *PhysicalDeviceFeatures
	// EnumerateDeviceExtensionProperties returns properties of available PhysicalDevice extensions
	//
	// physicalDevice - The PhysicalDevice to retrieve device extensions for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkEnumerateDeviceExtensionProperties.html
	EnumerateDeviceExtensionProperties(physicalDevice core.PhysicalDevice) (map[string]*ExtensionProperties, common.VkResult, error)
	// EnumerateDeviceExtensionPropertiesForLayer returns properties of available PhysicalDevice extensions
	// for the specifies layer
	//
	// physicalDevice - The PhysicalDevice to retrieve extension properties for
	//
	// layerName - Name of the layer to retrieve extensions from
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkEnumerateDeviceExtensionProperties.html
	EnumerateDeviceExtensionPropertiesForLayer(physicalDevice core.PhysicalDevice, layerName string) (map[string]*ExtensionProperties, common.VkResult, error)
	// EnumerateDeviceLayerProperties returns properties of available PhysicalDevice layers
	//
	// physicalDevice - The PhysicalDevice to retrieve layer properties for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkEnumerateDeviceLayerProperties.html
	EnumerateDeviceLayerProperties(physicalDevice core.PhysicalDevice) (map[string]*LayerProperties, common.VkResult, error)
	// GetPhysicalDeviceMemoryProperties reports memory information for a PhysicalDevice
	//
	// physicalDevice - The PhysicalDevice to retrieve memory properties for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMemoryProperties.html
	GetPhysicalDeviceMemoryProperties(physicalDevice core.PhysicalDevice) *PhysicalDeviceMemoryProperties
	// GetPhysicalDeviceFormatProperties lists a PhysicalDevice object's format capabilities
	//
	// physicalDevice - The PhysicalDevice to retrieve format properties for
	//
	// format - The format whose properties are queried
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html
	GetPhysicalDeviceFormatProperties(physicalDevice core.PhysicalDevice, format Format) *FormatProperties
	// GetPhysicalDeviceImageFormatProperties lists a PhysicalDevice object's image format capabilities
	//
	// physicalDevice - The PhysicalDevice to retrieve image format properties for
	//
	// format - Specifies the Image format
	//
	// imageType - Specifies the Image type
	//
	// tiling - Specifies the Image tiling
	//
	// usages - Specifies the intended usage of the Image
	//
	// flags - Specifies additional parmeters of the Image
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html
	GetPhysicalDeviceImageFormatProperties(physicalDevice core.PhysicalDevice, format Format, imageType ImageType, tiling ImageTiling, usages ImageUsageFlags, flags ImageCreateFlags) (*ImageFormatProperties, common.VkResult, error)
	// GetPhysicalDeviceSparseImageFormatProperties retrieves properties of an image format applied to sparse images
	//
	// physicalDevice - The PhysicalDevice to retireve sparse image format properties for
	//
	// format - The Image format
	//
	// imageType - The dimensionality of the Image
	//
	// samples - Specifies the number of samples per texel
	//
	// usages - Describes the intended usage of the Image
	//
	// tiling - The tiling arrangement of the texel blocks in memory
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties.html
	GetPhysicalDeviceSparseImageFormatProperties(physicalDevice core.PhysicalDevice, format Format, imageType ImageType, samples SampleCountFlags, usages ImageUsageFlags, tiling ImageTiling) []SparseImageFormatProperties
}

type DeviceDriver interface {
	Loader() loader.Loader
	Device() core.Device

	// DestroyBuffer deletes a buffer and underlying structures from the device. **Warning**
	// after destruction, the object will still exist, but the Vulkan object handle
	// that backs it will be invalid. Do not call further methods with this Buffer.
	//
	// buffer - The Buffer object to destroy
	//
	// callbacks - An set of allocation callbacks to control the memory free behavior of this command
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyBuffer.html
	DestroyBuffer(buffer core.Buffer, callbacks *loader.AllocationCallbacks)
	// GetBufferMemoryRequirements returns the memory requirements for a Buffer.
	//
	// buffer - The Buffer object to retrieve memory requirements for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetBufferMemoryRequirements.html
	GetBufferMemoryRequirements(buffer core.Buffer) *MemoryRequirements
	// BindBufferMemory binds DeviceMemory to a Buffer
	//
	// buffer - The Buffer object to bind to memory
	//
	// memory - A DeviceMemory object describing the device memory to attach
	//
	// offset - The start offset of the region of memory which is to be bound to the buffer.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkBindBufferMemory.html
	BindBufferMemory(buffer core.Buffer, memory core.DeviceMemory, offset int) (common.VkResult, error)
	// DestroyBufferView deletes a buffer view and the underlying structures from the device. **Warning**
	// after destruction, the object will continue to exist, but the Vulkan object handle
	// that backs it will be invalid. Do not call further methods with the BufferView object.
	//
	// bufferView - The BufferView object to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyBufferView.html
	DestroyBufferView(bufferView core.BufferView, callbacks *loader.AllocationCallbacks)
	// FreeCommandBuffers frees the provided CommandBuffer objects and usually returns the underlying
	// memory to the CommandPool.
	//
	// commandBuffers - one or more CommandBuffer objects to free
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkFreeCommandBuffers.html
	FreeCommandBuffers(commandBuffers ...core.CommandBuffer)
	// BeginCommandBuffer starts recording on a CommandBuffer
	//
	// commandBuffer - The CommandBuffer object to begin recording on
	//
	// o - Defines additional information about how the CommandBuffer begins recording
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkBeginCommandBuffer.html
	BeginCommandBuffer(commandBuffer core.CommandBuffer, o CommandBufferBeginInfo) (common.VkResult, error)
	// EndCommandBuffer finishes recording on a CommandBuffer
	//
	// commandBuffer - The CommandBuffer object to end recording on
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkEndCommandBuffer.html
	EndCommandBuffer(commandBuffer core.CommandBuffer) (common.VkResult, error)
	// ResetCommandBuffer returns a CommandBuffer to its initial state
	//
	// commandBuffer - The CommandBuffer object to reset
	//
	// flags - Options controlling the reset operation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetCommandBuffer.html
	ResetCommandBuffer(commandBuffer core.CommandBuffer, flags CommandBufferResetFlags) (common.VkResult, error)

	// CmdBeginRenderPass begins a new RenderPass
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// contents - Specifies how the commands in the first subpass will be provided
	//
	// o - Specifies the RenderPass to begin an instance of, and the Framebuffer the instance uses
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass.html
	CmdBeginRenderPass(commandBuffer core.CommandBuffer, contents SubpassContents, o RenderPassBeginInfo) error
	// CmdEndRenderPass ends the current RenderPass
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdEndRenderPass.html
	CmdEndRenderPass(commandBuffer core.CommandBuffer)
	// CmdBindPipeline binds a pipeline object to this CommandBuffer
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// bindPoint - Specifies to which bind point the Pipeline is bound
	//
	// pipeline - The Pipeline to be bound
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBindPipeline.html
	CmdBindPipeline(commandBuffer core.CommandBuffer, bindPoint PipelineBindPoint, pipeline core.Pipeline)
	// CmdDraw draws primitives without indexing the vertices
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// vertexCount - The number of vertices to draw
	//
	// instanceCount - The number of instances to draw
	//
	// firstVertex - The index of the first vertex to draw
	//
	// firstInstance - The instance ID of the first instance to draw
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDraw.html
	CmdDraw(commandBuffer core.CommandBuffer, vertexCount, instanceCount int, firstVertex, firstInstance uint32)
	// CmdDrawIndexed draws primitives with indexed vertices
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// indexCount - The number of vertices to draw
	//
	// instanceCount - The number of instances to draw
	//
	// firstIndex - The base index within the index Buffer
	//
	// vertexOffset - The value added to the vertex index before indexing into the vertex Buffer
	//
	// firstInstance - The instance ID of the first instance to draw
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexed.html
	CmdDrawIndexed(commandBuffer core.CommandBuffer, indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32)
	// CmdBindVertexBuffers binds vertex Buffers to this CommandBuffer
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// firstBinding - The index of the first input binding whose state is updated by the command
	//
	// buffers - A slice of Buffer objects
	//
	// bufferOffsets - A slice of Buffer offsets
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBindVertexBuffers.html
	CmdBindVertexBuffers(commandBuffer core.CommandBuffer, firstBinding int, buffers []core.Buffer, bufferOffsets []int)
	// CmdBindIndexBuffer binds an index Buffer to this CommandBuffer
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// buffer - The Buffer being bound
	//
	// offset - The starting offset in bytes within Buffer, used in index Buffer address calculations
	//
	// indexType - Specifies the size of the indices
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer.html
	CmdBindIndexBuffer(commandBuffer core.CommandBuffer, buffer core.Buffer, offset int, indexType IndexType)
	// CmdCopyBuffer copies data between Buffer regions
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// srcBuffer - The source Buffer
	//
	// dstBuffer - The destination Buffer
	//
	// copyRegions - A slice of structures specifying the regions to copy
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBuffer.html
	CmdCopyBuffer(commandBuffer core.CommandBuffer, srcBuffer core.Buffer, dstBuffer core.Buffer, copyRegions ...BufferCopy) error
	// CmdBindDescriptorSets binds DescriptorSets to this CommandBuffer
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// bindPoint - Indicates the type of the pipeline that will use the descriptors
	//
	// layout - A PipelineLayout object used to program the bindings
	//
	// firstSet - The set number of the first DescriptorSet to be bound
	//
	// sets - A slice of DescriptorSet objects describing the DescriptorSets to bind
	//
	// dynamicOffsets - A slice of values specifying dynamic offsets
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorSets.html
	CmdBindDescriptorSets(commandBuffer core.CommandBuffer, bindPoint PipelineBindPoint, layout core.PipelineLayout, firstSet int, sets []core.DescriptorSet, dynamicOffsets []int)
	// CmdPipelineBarrier inserts a memory dependency into the recorded commands
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// srcStageMask - Specifies the source stages
	//
	// dstStageMask - Specifies the destination stages
	//
	// dependencies - Specifies how execution and memory dependencies are formed
	//
	// memoryBarriers - A slice of MemoryBarrier structures
	//
	// bufferMemoryBarriers - A slice of BufferMemoryBarrier structures
	//
	// imageMemoryBarriers - A slice of ImageMemoryBarrier structures
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdPipelineBarrier.html
	CmdPipelineBarrier(commandBuffer core.CommandBuffer, srcStageMask, dstStageMask PipelineStageFlags, dependencies DependencyFlags, memoryBarriers []MemoryBarrier, bufferMemoryBarriers []BufferMemoryBarrier, imageMemoryBarriers []ImageMemoryBarrier) error
	// CmdCopyBufferToImage copies data from a Buffer to an Image
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// buffer - The source buffer
	//
	// image - The destination Image
	//
	// layout - The layout of the destination Image subresources for the copy
	//
	// regions - A slice of BufferImageCopy structures specifying the regions to copy
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBufferToImage.html
	CmdCopyBufferToImage(commandBuffer core.CommandBuffer, buffer core.Buffer, image core.Image, layout ImageLayout, regions ...BufferImageCopy) error
	// CmdBlitImage copies regions of an Image, potentially performing format conversion
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// sourceImage - The source Image
	//
	// sourceImageLayout - The layout of the source Image subresources for the blit
	//
	// destinationImage - The destination Image
	//
	// destinationImageLayout - The layout of the destination Image subresources for the blit
	//
	// regions - A slice of ImageBlit structures specifying the regions to blit
	//
	// filter - Specifies the filter to apply if the blits require scaling
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBlitImage.html
	CmdBlitImage(commandBuffer core.CommandBuffer, sourceImage core.Image, sourceImageLayout ImageLayout, destinationImage core.Image, destinationImageLayout ImageLayout, regions []ImageBlit, filter Filter) error
	// CmdPushConstants updates the values of push constants
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// layout - The pipeline layout used to program the push constant updates
	//
	// stageFlags - Specifies the shader stages that will use the push constants in the updated range
	//
	// offset - The start offset of the push constant range to update, in units of bytes
	//
	// valueBytes - A slice of bytes containing the new push constant values
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdPushConstants.html
	CmdPushConstants(commandBuffer core.CommandBuffer, layout core.PipelineLayout, stageFlags ShaderStageFlags, offset int, valueBytes []byte)
	// CmdSetViewport sets the viewport dynamically for a CommandBuffer
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// viewports - A slice of Viewport structures specifying viewport parameters
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewport.html
	CmdSetViewport(commandBuffer core.CommandBuffer, viewports ...Viewport)
	// CmdSetScissor sets scissor rectangles dynamically for a CommandBuffer
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// scissors - A slice of Rect2D structures specifying scissor rectangles
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetScissor.html
	CmdSetScissor(commandBuffer core.CommandBuffer, scissors ...Rect2D)
	// CmdCopyImage copies data between Images
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// srcImage - The source Image
	//
	// srcImageLayout - The current layout of the source Image subresource
	//
	// dstImage - The destination Image
	//
	// dstImageLayout - The current layout of the destination Image subresource
	//
	// regions - A slice of ImageCopy structures specifying the regions to copy
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImage.html
	CmdCopyImage(commandBuffer core.CommandBuffer, srcImage core.Image, srcImageLayout ImageLayout, dstImage core.Image, dstImageLayout ImageLayout, regions ...ImageCopy) error
	// CmdNextSubpass transitions to the next subpass of a RenderPass
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// contents - Specifies how the commands in the next subpass will be provided
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass.html
	CmdNextSubpass(commandBuffer core.CommandBuffer, contents SubpassContents)
	// CmdWaitEvents waits for one or more events and inserts a set of memory
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// events - A slice of Event objects to wait on
	//
	// srcStageMask - Specifies the source stage mask
	//
	// dstStageMask - Specifies the destination stage mask
	//
	// memoryBarriers - A slice of MemoryBarrier structures
	//
	// bufferMemoryBarriers - A slice of BufferMemoryBarrier structures
	//
	// imageMemoryBarriers - A slice of ImageMemoryBarrier structures
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents.html
	CmdWaitEvents(commandBuffer core.CommandBuffer, events []core.Event, srcStageMask PipelineStageFlags, dstStageMask PipelineStageFlags, memoryBarriers []MemoryBarrier, bufferMemoryBarriers []BufferMemoryBarrier, imageMemoryBarriers []ImageMemoryBarrier) error
	// CmdSetEvent sets an Event object to the signaled state
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// event - The Event that will be signaled
	//
	// stageMask - Specifies teh source stage mask used to determine the first synchronization scope
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent.html
	CmdSetEvent(commandBuffer core.CommandBuffer, event core.Event, stageMask PipelineStageFlags)
	// CmdClearColorImage clears regions of a color Image
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// image - The Image to be cleared
	//
	// imageLayout - Specifies the current layout of the Image subresource ranges to be cleared
	//
	// color - A ClearColorValue containing the values that the Image subresource ranges will be cleared to
	//
	// ranges - A slice of ImageSubresourceRange structures describing a range of mipmap levels, array layers,
	// and aspects to be cleared.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdClearColorImage.html
	CmdClearColorImage(commandBuffer core.CommandBuffer, image core.Image, imageLayout ImageLayout, color ClearColorValue, ranges ...ImageSubresourceRange)
	// CmdResetQueryPool resets queries in a QueryPool
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// queryPool - The QueryPool managing the queries being reset
	//
	// startQuery - The initial query index to reset
	//
	// queryCount - The number of queries to reset
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdResetQueryPool.html
	CmdResetQueryPool(commandBuffer core.CommandBuffer, queryPool core.QueryPool, startQuery, queryCount int)
	// CmdBeginQuery begins a query
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// queryPool - The QueryPool that will manage the results of the query
	//
	// query - The query index within the QueryPool that will contain the results
	//
	// flags - Specifies constraints on the types of queries that can be performed
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBeginQuery.html
	CmdBeginQuery(commandBuffer core.CommandBuffer, queryPool core.QueryPool, query int, flags QueryControlFlags)
	// CmdEndQuery ends a query
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// queryPool - The QueryPool that is managing the results of the query
	//
	// query - The query index within the QueryPool where the result is stored
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdEndQuery.html
	CmdEndQuery(commandBuffer core.CommandBuffer, queryPool core.QueryPool, query int)
	// CmdCopyQueryPoolResults copies the results of queries in a QueryPool to a Buffer object
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// queryPool - The QueryPool managing the queries containing the desired results
	//
	// firstQuery - The initial query index
	//
	// queryCount - The number of queries
	//
	// dstBuffer - A Buffer object that will receive the results of the copy command
	//
	// dstOffset - An offset into the destination Buffer
	//
	// stride - The stride in bytes between the results for individual queries within the destination Buffer
	//
	// flags - Specifies how and when results are returned
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdCopyQueryPoolResults.html
	CmdCopyQueryPoolResults(commandBuffer core.CommandBuffer, queryPool core.QueryPool, firstQuery, queryCount int, dstBuffer core.Buffer, dstOffset, stride int, flags QueryResultFlags)
	// CmdExecuteCommands executes a secondary CommandBuffer from a primary CommandBuffer
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// commandBuffers - A slice of CommandBuffer objects, which are recorded to execute in the primary CommandBuffer
	// in the order they are listed in the slice
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdExecuteCommands.html
	CmdExecuteCommands(commandBuffer core.CommandBuffer, commandBuffers ...core.CommandBuffer)
	// CmdClearAttachments clears regions within bound Framebuffer attachments
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// attachments - A slice of ClearAttachment structures defining the attachments to clear and the clear values to use.
	//
	// rects - A slice of ClearRect structures defining regions within each selected attachment to clear
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdClearAttachments.html
	CmdClearAttachments(commandBuffer core.CommandBuffer, attachments []ClearAttachment, rects []ClearRect) error
	// CmdClearDepthStencilImage fills regions of a combined depth/stencil image
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// image - The Image to be cleared
	//
	// imageLayout - Specifies the current layout of the Image subresource ranges to be cleared
	//
	// depthStencil - Contains the values that the depth and stencil images will be cleared to
	//
	// ranges - A slice of ImageSubrsourceRange structures describing a range of mipmap levels, array layers,
	// and aspects to be cleared
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdClearDepthStencilImage.html
	CmdClearDepthStencilImage(commandBuffer core.CommandBuffer, image core.Image, imageLayout ImageLayout, depthStencil *ClearValueDepthStencil, ranges ...ImageSubresourceRange)
	// CmdCopyImageToBuffer copies image data into a buffer
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// srcImage - The source Image
	//
	// srcImageLayout - The layout of the source Image subresources for the copy
	//
	// dstBuffer - The desination Buffer
	//
	// regions - A slice of BufferImageCopy structures specifying the regions to copy
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImageToBuffer.html
	CmdCopyImageToBuffer(commandBuffer core.CommandBuffer, srcImage core.Image, srcImageLayout ImageLayout, dstBuffer core.Buffer, regions ...BufferImageCopy) error
	// CmdDispatch dispatches compute work items
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// groupCountX - the number of local workgroups to dispatch in the X dimension
	//
	// groupCountY - the number of local workgroups to dispatch in the Y dimension
	//
	// groupCountZ - the number of local workgroups to dispatch in the Z dimension
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDispatch.html
	CmdDispatch(commandBuffer core.CommandBuffer, groupCountX, groupCountY, groupCountZ int)
	// CmdDispatchIndirect dispatches compute work items with indirect parameters
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// buffer - The Buffer containing dispatch parameters
	//
	// offset - The byte offset into the Buffer where parameters begin
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchIndirect.html
	CmdDispatchIndirect(commandBuffer core.CommandBuffer, buffer core.Buffer, offset int)
	// CmdDrawIndexedIndirect draws primitives with indirect parameters and indexed vertices
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// buffer - The Buffer containing draw parameters
	//
	// offset - The byte offset into the Buffer where parameters begin
	//
	// drawCount - The number of draws to execute, which can be zero
	//
	// stride - The byte stride between successive sets of draw parameters
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexedIndirect.html
	CmdDrawIndexedIndirect(commandBuffer core.CommandBuffer, buffer core.Buffer, offset int, drawCount, stride int)
	// CmdDrawIndirect draws primitives with indirect parameters
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// buffer - The buffer containing draw parameters
	//
	// offset - The byte offset into the Buffer where parameters begin
	//
	// drawCount - The number of draws to execute, which can be zero
	//
	// stride - The byte stride between successive sets of draw parameters
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndirect.html
	CmdDrawIndirect(commandBuffer core.CommandBuffer, buffer core.Buffer, offset int, drawCount, stride int)
	// CmdFillBuffer fills a region of a buffer with a fixed value
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// dstBuffer - The Buffer to be filled
	//
	// dstOffset - The byte offset into the Buffer at which to start filling, must be a multiple of 4
	//
	// size - The number of bytes to fill
	//
	// data - The 4-byte word written repeatedly to the Buffer to fill size bytes of data.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdFillBuffer.html
	CmdFillBuffer(commandBuffer core.CommandBuffer, dstBuffer core.Buffer, dstOffset int, size int, data uint32)
	// CmdResetEvent resets an Event object to non-signaled state
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// event - The Event that will be unsignaled
	//
	// stageMask - Specifies the source stage mask used to determine when the Event is unsignaled
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdResetEvent.html
	CmdResetEvent(commandBuffer core.CommandBuffer, event core.Event, stageMask PipelineStageFlags)
	// CmdResolveImage resolves regions of an Image
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// srcImage - The source Image
	//
	// srcImageLayout - The layout of the source Image subresources for the resolve
	//
	// dstImage - The destination Image
	//
	// dstImageLayout - The layout of the destination Image subresources for the resolve
	//
	// regions - A slice of ImageResolve structure specifying the regions to resolve
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdResolveImage.html
	CmdResolveImage(commandBuffer core.CommandBuffer, srcImage core.Image, srcImageLayout ImageLayout, dstImage core.Image, dstImageLayout ImageLayout, regions ...ImageResolve) error
	// CmdSetBlendConstants sets the values of the blend constants
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// blendConstants - An array of four values specifying the R, G, B, and A components of the blend
	// color used in blending, depending on the blend factor
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetBlendConstants.html
	CmdSetBlendConstants(commandBuffer core.CommandBuffer, blendConstants [4]float32)
	// CmdSetDepthBias sets depth bias factors and clamp dynamically for the CommandBuffer
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// depthBiasConstantFactor - The scalar factor controlling the constant depth value added to each fragment
	//
	// depthBiasClamp - The maximum (or minimum) depth bias of a fragment
	//
	// depthBiasSlopeFactor - The scalar factor applied to a fragment's slope in depth bias calculations
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBias.html
	CmdSetDepthBias(commandBuffer core.CommandBuffer, depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor float32)
	// CmdSetDepthBounds sets depth bounds range dynamically for the CommandBuffer
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// min - The minimum depth bound
	//
	// max - The maximum depth bound
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBounds.html
	CmdSetDepthBounds(commandBuffer core.CommandBuffer, min, max float32)
	// CmdSetLineWidth sets line width dynamically for the CommandBuffer
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// lineWidth - The width of rasterized line segments
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineWidth.html
	CmdSetLineWidth(commandBuffer core.CommandBuffer, lineWidth float32)
	// CmdSetStencilCompareMask sets the stencil compare mask dynamically for the CommandBuffer
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// faceMask - Specifies the set of stencil state for which to update the compare mask
	//
	// compareMask - The new value to use as the stencil compare mask
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilCompareMask.html
	CmdSetStencilCompareMask(commandBuffer core.CommandBuffer, faceMask StencilFaceFlags, compareMask uint32)
	// CmdSetStencilReference sets stencil reference value dynamically for the CommandBuffer
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// faceMask - Specifies the set of stencil state for which to update the reference value
	//
	// reference - The new value to use as the stencil reference value
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilReference.html
	CmdSetStencilReference(commandBuffer core.CommandBuffer, faceMask StencilFaceFlags, reference uint32)
	// CmdSetStencilWriteMask sets the stencil write mask dynamically for the CommandBuffer
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// faceMask - Specifies the set of stencil state for which to update the write mask
	//
	// reference - The new value to use as the stencil write mask
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilWriteMask.html
	CmdSetStencilWriteMask(commandBuffer core.CommandBuffer, faceMask StencilFaceFlags, writeMask uint32)
	// CmdUpdateBuffer updates a buffer's contents from host memory
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// dstBuffer - The Buffer to be updated
	//
	// dstOffset - The byte offset into the Buffer to start updating, must be a multiple of 4
	//
	// dataSize - The number of bytes to update, must be a multiple of 4
	//
	// data - The source data for the buffer update, must be at least dataSize bytes in size
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdUpdateBuffer.html
	CmdUpdateBuffer(commandBuffer core.CommandBuffer, dstBuffer core.Buffer, dstOffset int, dataSize int, data []byte)
	// CmdWriteTimestamp writes a device timestamp into a query object
	//
	// commandBuffer - The CommandBuffer object to record to
	//
	// pipelineStage - Specifies a stage of the pipeline
	//
	// queryPool - The QueryPool that will manage the timestamp
	//
	// query - The query within the QueryPool that will contain the timestamp
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdWriteTimestamp.html
	CmdWriteTimestamp(commandBuffer core.CommandBuffer, pipelineStage PipelineStageFlags, queryPool core.QueryPool, query int)
	// DestroyCommandPool destroys a CommandPool object and the underlying structures. **Warning** after
	// destruction, the object will continue to exist, but the Vulkan object handle that backs
	// it will be invalid. Do not call further methods on the CommandPool.
	//
	// commandPool - The CommandPool object to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyCommandPool.html
	DestroyCommandPool(commandPool core.CommandPool, callbacks *loader.AllocationCallbacks)
	// ResetCommandPool resets a CommandPool, recycling all the resources from all the CommandBuffer objects
	// allocated from the CommandPool back to the CommandPool.  All CommandBuffer objects that
	// have been allocated from the CommandPool are put in the initial state.
	//
	// commandPool - The CommandPool object to reset
	//
	// flags - Controls the reset operation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetCommandPool.html
	ResetCommandPool(commandPool core.CommandPool, flags CommandPoolResetFlags) (common.VkResult, error)
	// DestroyDescriptorPool destroys a DescriptorPool object and the underlying structures. **Warning** after
	// destruction, the object will continue to exist, but the Vulkan object handle that backs it will
	// be invalid. Do not call further methods with the DescriptorPool
	//
	// descriptorPool - The DescriptorPool object to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorPool.html
	DestroyDescriptorPool(descriptorPool core.DescriptorPool, callbacks *loader.AllocationCallbacks)
	// ResetDescriptorPool resets the DescriptorPool and recycles all of the resources from all of the
	// DescriptorSet objects allocated from the DescriptorPool back to the DescriptorPool, and the
	// DescriptorSet objects are implicitly freed.
	//
	// descriptorPool - The DescriptorPool object to reset
	//
	// flags - Reserved (always 0)
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetDescriptorPool.html
	ResetDescriptorPool(descriptorPool core.DescriptorPool, flags DescriptorPoolResetFlags) (common.VkResult, error)
	// FreeDescriptorSets frees the provided DescriptorSet objects
	//
	// sets - One or more DescriptorSet objects to free
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkFreeDescriptorSets.html
	FreeDescriptorSets(sets ...core.DescriptorSet) (common.VkResult, error)
	// DestroyDescriptorSetLayout destroys a DescriptorSetLayout object and the underlying structures.
	// **Warning** after destruction, the object will continue to exist, but the Vulkan object handle
	// that backs it will be invalid. Do not call further methods with the DescriptorSetLayout
	//
	// descriptorSetLayout - DescriptorSetLayout object to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorSetLayout.html
	DestroyDescriptorSetLayout(descriptorSetLayout core.DescriptorSetLayout, callbacks *loader.AllocationCallbacks)
	// MapMemory maps a memory object into application address space
	//
	// memory - DeviceMemory object to map
	//
	// offset - A zero-based byte offset from the beginning of the memory object
	//
	// size - The size of the memory range to map, or -1 to map from offset to the end of the
	// allocation
	//
	// flags - Reserved (always 0)
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkMapMemory.html
	MapMemory(memory core.DeviceMemory, offset int, size int, flags MemoryMapFlags) (unsafe.Pointer, common.VkResult, error)
	// UnmapMemory unmaps a previously-mapped memory object
	//
	// memory - The DeviceMemory object to unmap
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkUnmapMemory.html
	UnmapMemory(memory core.DeviceMemory)
	// FreeMemory frees the DeviceMemory. **Warning** after freeing, the object will continue to exist,
	// but the Vulkan object handle that backs it will be invalid. Do not call further methods with
	// the DeviceMemory object.
	//
	// memory - The DeviceMemory object to free
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkFreeMemory.html
	FreeMemory(memory core.DeviceMemory, callbacks *loader.AllocationCallbacks)
	// GetDeviceMemoryCommitment returns the current number of bytes currently committed to a DeviceMemory
	//
	// memory - The DeviceMemory object to get memory commitment for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMemoryCommitment.html
	GetDeviceMemoryCommitment(memory core.DeviceMemory) int
	// FlushMappedMemoryRanges flushes the provided mapped memory ranges to a device
	//
	// ranges - The memory ranges to flush to the device
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkFlushMappedMemoryRanges.html
	FlushMappedMemoryRanges(ranges ...MappedMemoryRange) (common.VkResult, error)
	// InvalidateMappedMemoryRanges invalidates the provided mapped memory ranges in the host cache
	//
	// ranges - The memory ranges to invalidate in the cache
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkInvalidateMappedMemoryRanges.html
	InvalidateMappedMemoryRanges(ranges ...MappedMemoryRange) (common.VkResult, error)
	// CreateBuffer creates a new Buffer object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the Buffer
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateBuffer.html
	CreateBuffer(allocationCallbacks *loader.AllocationCallbacks, o BufferCreateInfo) (core.Buffer, common.VkResult, error)
	// CreateBufferView creates a new BufferView object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the BufferView
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateBufferView.html
	CreateBufferView(allocationCallbacks *loader.AllocationCallbacks, o BufferViewCreateInfo) (core.BufferView, common.VkResult, error)
	// CreateCommandPool creates a new CommandPool object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the CommandPool
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateCommandPool.html
	CreateCommandPool(allocationCallbacks *loader.AllocationCallbacks, o CommandPoolCreateInfo) (core.CommandPool, common.VkResult, error)
	// CreateDescriptorPool creates a new DescriptorPool object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the DescriptorPool
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorPool.html
	CreateDescriptorPool(allocationCallbacks *loader.AllocationCallbacks, o DescriptorPoolCreateInfo) (core.DescriptorPool, common.VkResult, error)
	// CreateDescriptorSetLayout creates a new DescriptorSetLayout object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the DescriptorSetLayout
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorSetLayout.html
	CreateDescriptorSetLayout(allocationCallbacks *loader.AllocationCallbacks, o DescriptorSetLayoutCreateInfo) (core.DescriptorSetLayout, common.VkResult, error)
	// CreateEvent creates a new Event object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the Event
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateEvent.html
	CreateEvent(allocationCallbacks *loader.AllocationCallbacks, options EventCreateInfo) (core.Event, common.VkResult, error)
	// CreateFence creates a new Fence object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the Fence
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateFence.html
	CreateFence(allocationCallbacks *loader.AllocationCallbacks, o FenceCreateInfo) (core.Fence, common.VkResult, error)
	// CreateFramebuffer creates a new Framebuffer object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the Framebuffer
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateFramebuffer.html
	CreateFramebuffer(allocationCallbacks *loader.AllocationCallbacks, o FramebufferCreateInfo) (core.Framebuffer, common.VkResult, error)
	// CreateGraphicsPipelines creates a slice of new Pipeline objects which can be used for drawing graphics
	//
	// pipelineCache - A PipelineCache object which can be used to accelerate pipeline creation
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - A slice of GraphicsPipelineCreateInfo structures containing parameters affecting the creation of the Pipeline objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateGraphicsPipelines.html
	CreateGraphicsPipelines(pipelineCache *core.PipelineCache, allocationCallbacks *loader.AllocationCallbacks, o ...GraphicsPipelineCreateInfo) ([]core.Pipeline, common.VkResult, error)
	// CreateComputePipelines creates a slice of new Pipeline objects which can be used for dispatching compute workloads
	//
	// pipelineCache - A PipelineCache object which can be used to accelerate pipeline creation
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - A slice of ComputePipelineCreateInfo structures containing parameters affecting the creation of the Pipeline objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateComputePipelines.html
	CreateComputePipelines(pipelineCache *core.PipelineCache, allocationCallbacks *loader.AllocationCallbacks, o ...ComputePipelineCreateInfo) ([]core.Pipeline, common.VkResult, error)
	// CreateImage creates a new Image object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the Image
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html
	CreateImage(allocationCallbacks *loader.AllocationCallbacks, options ImageCreateInfo) (core.Image, common.VkResult, error)
	// CreateImageView creates a new ImageView object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the ImageView
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateImageView.html
	CreateImageView(allocationCallbacks *loader.AllocationCallbacks, o ImageViewCreateInfo) (core.ImageView, common.VkResult, error)
	// CreatePipelineCache creates a new PipelineCache object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the PipelineCache
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreatePipelineCache.html
	CreatePipelineCache(allocationCallbacks *loader.AllocationCallbacks, o PipelineCacheCreateInfo) (core.PipelineCache, common.VkResult, error)
	// CreatePipelineLayout creates a new PipelineLayout object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the PipelineLayout
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreatePipelineLayout.html
	CreatePipelineLayout(allocationCallbacks *loader.AllocationCallbacks, o PipelineLayoutCreateInfo) (core.PipelineLayout, common.VkResult, error)
	// CreateQueryPool creates a new QueryPool object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the QueryPool
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateQueryPool.html
	CreateQueryPool(allocationCallbacks *loader.AllocationCallbacks, o QueryPoolCreateInfo) (core.QueryPool, common.VkResult, error)
	// CreateRenderPass creates a new RenderPass object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the RenderPass
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass.html
	CreateRenderPass(allocationCallbacks *loader.AllocationCallbacks, o RenderPassCreateInfo) (core.RenderPass, common.VkResult, error)
	// CreateSampler creates a new Sampler object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the Sampler
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateSampler.html
	CreateSampler(allocationCallbacks *loader.AllocationCallbacks, o SamplerCreateInfo) (core.Sampler, common.VkResult, error)
	// CreateSemaphore creates a new Semaphore object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the Semaphore
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateSemaphore.html
	CreateSemaphore(allocationCallbacks *loader.AllocationCallbacks, o SemaphoreCreateInfo) (core.Semaphore, common.VkResult, error)
	// CreateShaderModule creates a new ShaderModule object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the ShaderModule
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateShaderModule.html
	CreateShaderModule(allocationCallbacks *loader.AllocationCallbacks, o ShaderModuleCreateInfo) (core.ShaderModule, common.VkResult, error)

	// GetQueue gets a Queue object from the Device
	//
	// queueFamilyIndex - The index of the queue family to which the Queue belongs
	//
	// queueIndex - The index within this queue family of the Queue to retrieve
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetDeviceQueue.html
	GetQueue(queueFamilyIndex int, queueIndex int) core.Queue
	// AllocateMemory allocates DeviceMemory
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Describes the parameters of the allocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkAllocateMemory.html
	AllocateMemory(allocationCallbacks *loader.AllocationCallbacks, o MemoryAllocateInfo) (core.DeviceMemory, common.VkResult, error)

	// AllocateCommandBuffers allocates CommandBuffer objects from an existing CommandPool
	//
	// o - Describes parameters of the allocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkAllocateCommandBuffers.html
	AllocateCommandBuffers(o CommandBufferAllocateInfo) ([]core.CommandBuffer, common.VkResult, error)
	// AllocateDescriptorSets allocates one or more DescriptorSet objects from a DescriptorPool
	//
	// device - The Device to allocate the DescriptorSet objects on
	//
	// o - Describes the parameters of the allocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkAllocateDescriptorSets.html
	AllocateDescriptorSets(o DescriptorSetAllocateInfo) ([]core.DescriptorSet, common.VkResult, error)

	// DestroyDevice destroys the logical Device object backing this driver.  **Warning** after
	// destruction, the object will continue to exist, but the Vulkan object handle that backs
	// it will be invalid. Do not call further methods on this driver or with the underlying
	// Device.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyDevice.html
	DestroyDevice(callbacks *loader.AllocationCallbacks)
	// DeviceWaitIdle waits for the Device to become idle
	//
	// device - The Device to wait on
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDeviceWaitIdle.html
	DeviceWaitIdle() (common.VkResult, error)
	// WaitForFences waits for one or more Fence objects to become signaled
	//
	// waitForAll - If true, then the call will wait until all fences in `fences` are signaled. If
	// false, the call will wait until any fence in `fences` is signaled.
	//
	// timeout - How long to wait before returning VKTimeout. May be common.NoTimeout to wait indefinitely.
	// The timeout is adjusted to the closest value allowed by the implementation timeout accuracy,
	// which may be substantially longer than the requested timeout.
	//
	// fences - A slice of Fence objects to wait for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkWaitForFences.html
	WaitForFences(waitForAll bool, timeout time.Duration, fences ...core.Fence) (common.VkResult, error)
	// ResetFences resets one or more objects to the unsignaled state
	//
	// fences - A slice of Fence objects to reset
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetFences.html
	ResetFences(fences ...core.Fence) (common.VkResult, error)
	// UpdateDescriptorSets updates the contents of one or more DescriptorSet objects
	//
	// writes - A slice of WriteDescriptorSet structures describing the DescriptorSet objects to
	// write to
	//
	// copies - A slice of CopyDescriptorSet structures describing the DescriptorSet objects to
	// copy between
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSets.html
	UpdateDescriptorSets(writes []WriteDescriptorSet, copies []CopyDescriptorSet) error
	// DestroyEvent destroys an Event and the underlying structures. **Warning** after destruction,
	// the object will continue to exist, but the Vulkan object handle that backs it will be invalid.
	// Do not call further methods with the Event
	//
	// event - The Event to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	DestroyEvent(event core.Event, callbacks *loader.AllocationCallbacks)
	// SetEvent sets the Event to the signaled state
	//
	// event - The Event object to signal
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkSetEvent.html
	SetEvent(event core.Event) (common.VkResult, error)
	// ResetEvent sets this Event to the unsignaled state
	//
	// event - The Event object to unsignal
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetEvent.html
	ResetEvent(event core.Event) (common.VkResult, error)
	// GetEventStatus retrieves the status of this Event
	//
	// event - The Event object to retrieve the status from
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetEventStatus.html
	GetEventStatus(event core.Event) (common.VkResult, error)
	// DestroyFence destroys the Fence object and the underlying structures.  **Warning** after destruction,
	// the object will continue to exist, but the Vulkan object handle that backs it will be invalid.
	// Do not call further methods with the Fence.
	//
	// fence - The Fence object to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyFence.html
	DestroyFence(fence core.Fence, callbacks *loader.AllocationCallbacks)
	// GetFenceStatus returns the status of a Fence
	//
	// fence - The Fence object to get the status of
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetFenceStatus.html
	GetFenceStatus(fence core.Fence) (common.VkResult, error)
	// DestroyFramebuffer destroys a Framebuffer object and the underlying structures. **Warning** after destruction,
	// the object will continue to exist, but the Vulkan object handle that backs it will be invalid. Do
	// not call further methods with the Framebuffer.
	//
	// fence - The Framebuffer object to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyFramebuffer.html
	DestroyFramebuffer(framebuffer core.Framebuffer, callbacks *loader.AllocationCallbacks)
	// DestroyImage destroys an Image object and the underlying structures. **Warning** after destruction,
	// the object will continue to exist, but the Vulkan object handle that backs it will be invalid.
	// Do not call further methods with the Image
	//
	// image - The Image object to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyImage.html
	DestroyImage(image core.Image, callbacks *loader.AllocationCallbacks)
	// GetImageMemoryRequirements returns the memory requirements for an Image
	//
	// image - The Image to get the memory requirements for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements.html
	GetImageMemoryRequirements(image core.Image) *MemoryRequirements
	// BindImageMemory binds a DeviceMemory object to this Image object
	//
	// image - The Image to bind to memory
	//
	// memory - Describes the DeviceMemory to attach
	//
	// offset - The start offset of the region of memory which is to be bound to the image.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkBindImageMemory.html
	BindImageMemory(image core.Image, memory core.DeviceMemory, offset int) (common.VkResult, error)
	// GetImageSubresourceLayout retrieves information about an Image subresource
	//
	// image - The Image to retrieve subresource data for
	//
	// subresource - Selects a specific subresource from the Image
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout.html
	GetImageSubresourceLayout(image core.Image, subresource *ImageSubresource) *SubresourceLayout
	// GetImageSparseMemoryRequirements queries the memory requirements for a sparse image
	//
	// image - The Image to retrieve sparse memory requirements for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetImageSparseMemoryRequirements.html
	GetImageSparseMemoryRequirements(image core.Image) []SparseImageMemoryRequirements
	// DestroyImageView destroys an ImageView object and the underlying structures. **Warning** after destruction,
	// the object will continue to exist, but the Vulkan object handle that backs it will be invalid. Do not
	// call further methods with the image
	DestroyImageView(image core.ImageView, callbacks *loader.AllocationCallbacks)

	// DestroyPipeline destroys a Pipeline object and the underlying structures. **Warning** after
	// destruction, the object will continue to exist, but the Vulkan object handle that backs it will be
	// invalid. Do not call further methods with the Pipeline.
	//
	// pipeline - The Pipeline to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyPipeline.html
	DestroyPipeline(pipeline core.Pipeline, callbacks *loader.AllocationCallbacks)

	// DestroyPipelineCache destroys a PipelineCache object and the underlying structures. **Warning** after
	// destruction, the object will continue to exist, but the Vulkan object handle that backs it will
	// be invalid. Do not call further methods with the PipelineCache
	//
	// pipelineCache - The PipelineCache object to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyPipelineCache.html
	DestroyPipelineCache(cache core.PipelineCache, callbacks *loader.AllocationCallbacks)
	// GetPipelineCacheData gets the data store from this PipelineCache
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPipelineCacheData.html
	GetPipelineCacheData(cache core.PipelineCache) ([]byte, common.VkResult, error)
	// MergePipelineCaches combines the data stores of multiple PipelineCache object into another one
	//
	// dstCache - The PipelineCache that the source caches will be merged into
	//
	// srcCaches - A slice of PipelineCache objects which will be merged into the destination PipelineCache
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkMergePipelineCaches.html
	MergePipelineCaches(dstCache core.PipelineCache, srcCaches ...core.PipelineCache) (common.VkResult, error)

	// DestroyPipelineLayout destroys a PipelineLayout object and the underlying structures. **Warning** after
	// destruction, the object will continue to exist, but the Vulkan object handle that backs it will be
	// invalid. Do not call further methods with the PipelineLayout
	//
	// layout - The PipelineLayout object to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyPipelineLayout.html
	DestroyPipelineLayout(layout core.PipelineLayout, callbacks *loader.AllocationCallbacks)

	// DestroyQueryPool destroys a QueryPool object and the underlying structures. **Warning** after destruction,
	// the object will continue to exist, but the Vulkan object handle that backs it will be invalid.
	// Do not call further methods with the QueryPool object
	//
	// queryPool - The QueryPool object to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyQueryPool.html
	DestroyQueryPool(queryPool core.QueryPool, callbacks *loader.AllocationCallbacks)
	// GetQueryPoolResults retrieves the status and results for a set of queries, and populates those results
	// into a preallocated byte array
	//
	// queryPool - The QueryPool object to get results for
	//
	// firstQuery - The initial query index
	//
	// queryCount - The number of queries to read
	//
	// results - A user-allocated slice of bytes where the results will be written
	//
	// resultStride - The stride in bytes between results for individual queries
	//
	// flags - Specifies how and when results are returned
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetQueryPoolResults.html
	GetQueryPoolResults(queryPool core.QueryPool, firstQuery, queryCount int, results []byte, resultStride int, flags QueryResultFlags) (common.VkResult, error)

	// QueueWaitIdle waits for a Queue to become idle
	//
	// queue - The Queue to wait on
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkQueueWaitIdle.html
	QueueWaitIdle(queue core.Queue) (common.VkResult, error)
	// QueueSubmit submits a sequence of Semaphore or CommandBuffer objects to a queue
	//
	// queue - The Queue to submit to
	//
	// fence - An optional Fence object to be signaled once all submitted CommandBuffer objects have
	// completed execution.
	//
	// o - A slice of SubmitInfo structures, each specifying a CommandBuffer submission batch
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkQueueSubmit.html
	QueueSubmit(queue core.Queue, fence *core.Fence, o ...SubmitInfo) (common.VkResult, error)
	// QueueBindSparse binds DeviceMemory to a sparse resource object
	//
	// queue -The queue to submit the sparse binding operation to
	//
	// fence - An optional Fence object to be signaled.
	//
	// bindInfos - A slice of BindSparseInfo structures, each speicfying a sparse binding submission batch
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkQueueBindSparse.html
	QueueBindSparse(queue core.Queue, fence *core.Fence, bindInfos ...BindSparseInfo) (common.VkResult, error)

	// DestroyRenderPass destroys a RenderPass object and the underlying structures.  **Warning** after
	// destruction, the object will continue to exist, but the Vulkan object handle that backs it will
	// be invalid. Do not call further methods with the RenderPass
	//
	// renderPass - The RenderPass object to destroy
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyRenderPass.html
	DestroyRenderPass(renderPass core.RenderPass, callbacks *loader.AllocationCallbacks)
	// GetRenderAreaGranularity returns the granularity for optimal render area
	//
	// renderPass - The RenderPass object to get granularity information for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetRenderAreaGranularity.html
	GetRenderAreaGranularity(renderPass core.RenderPass) Extent2D

	// DestroySampler destroys a Sampler object and the underlying structures. **Warning** after destruction,
	// the object will continue to exist, but the Vulkan objec thandle that backs it will be invalid.
	// Do not call further methods with the Sampler
	//
	// sampler - The Sampler object to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroySampler.html
	DestroySampler(sampler core.Sampler, callbacks *loader.AllocationCallbacks)

	// DestroySemaphore destroys a Semaphore object and the underlying structures. **Warning** after destruction,
	// the object will continue to exist, but the Vulkan object handle that backs it will be invalid.
	// Do not call further methods with the ShaderModule
	//
	// semaphore - The Semaphore object to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroySemaphore.html
	DestroySemaphore(semaphore core.Semaphore, callbacks *loader.AllocationCallbacks)

	// DestroyShaderModule destroys a ShaderModule object and the underlying structures. **Warning** after
	// destruction, the object will continue to exist, but the Vulkan object handle that backs it will
	// be invalid. Do not call further methods with the ShaderModule
	//
	// shaderModule - The SHaderModule object to destroy
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyShaderModule.html
	DestroyShaderModule(shaderModule core.ShaderModule, callbacks *loader.AllocationCallbacks)
}

type CoreDeviceDriver interface {
	InstanceDriver() CoreInstanceDriver
	DeviceDriver
}
