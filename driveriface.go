package core

import (
	"github.com/CannibalVox/VKng/core/common"
	"unsafe"
)

type Driver interface {
	Destroy()
	CreateInstanceDriver(instance VkInstance) (Driver, error)
	CreateDeviceDriver(device VkDevice) (Driver, error)
	LoadProcAddr(name *Char) unsafe.Pointer
	Version() common.APIVersion

	VkEnumerateInstanceVersion(pApiVersion *Uint32) (VkResult, error)

	//Instance
	VkEnumerateInstanceExtensionProperties(pLayerName *Char, pPropertyCount *Uint32, pProperties *VkExtensionProperties) (VkResult, error)
	VkEnumerateInstanceLayerProperties(pPropertyCount *Uint32, pProperties *VkLayerProperties) (VkResult, error)
	VkCreateInstance(pCreateInfo *VkInstanceCreateInfo, pAllocator *VkAllocationCallbacks, pInstance *VkInstance) (VkResult, error)
	VkEnumeratePhysicalDevices(instance VkInstance, pPhysicalDeviceCount *Uint32, pPhysicalDevices *VkPhysicalDevice) (VkResult, error)
	VkDestroyInstance(instance VkInstance, pAllocator *VkAllocationCallbacks) error
	VkGetPhysicalDeviceFeatures(physicalDevice VkPhysicalDevice, pFeatures *VkPhysicalDeviceFeatures) error
	VkGetPhysicalDeviceFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, pFormatProperties *VkFormatProperties) error
	VkGetPhysicalDeviceImageFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, t VkImageType, tiling VkImageTiling, usage VkImageUsageFlags, flags VkImageCreateFlags, pImageFormatProperties *VkImageFormatProperties) (VkResult, error)
	VkGetPhysicalDeviceProperties(physicalDevice VkPhysicalDevice, pProperties *VkPhysicalDeviceProperties) error
	VkGetPhysicalDeviceQueueFamilyProperties(physicalDevice VkPhysicalDevice, pQueueFamilyPropertyCount *Uint32, pQueueFamilyProperties *VkQueueFamilyProperties) error
	VkGetPhysicalDeviceMemoryProperties(physicalDevice VkPhysicalDevice, pMemoryProperties *VkPhysicalDeviceMemoryProperties) error
	VkEnumerateDeviceExtensionProperties(physicalDevice VkPhysicalDevice, pLayerName *Char, pPropertyCount *Uint32, pProperties *VkExtensionProperties) (VkResult, error)
	VkEnumerateDeviceLayerProperties(physicalDevice VkPhysicalDevice, pPropertyCount *Uint32, pProperties *VkLayerProperties) (VkResult, error)
	VkGetPhysicalDeviceSparseImageFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, t VkImageType, samples VkSampleCountFlagBits, usage VkImageUsageFlags, tiling VkImageTiling, pPropertyCount *Uint32, pProperties *VkSparseImageFormatProperties) error
	VkCreateDevice(physicalDevice VkPhysicalDevice, pCreateInfo *VkDeviceCreateInfo, pAllocator *VkAllocationCallbacks, pDevice *VkDevice) (VkResult, error)

	VkEnumeratePhysicalDeviceGroups(instance VkInstance, pPhysicalDeviceGroupCount *Uint32, pPhysicalDeviceGroupProperties *VkPhysicalDeviceGroupProperties) (VkResult, error)
	VkGetPhysicalDeviceFeatures2(physicalDevice VkPhysicalDevice, pFeatures *VkPhysicalDeviceFeatures2) error
	VkGetPhysicalDeviceProperties2(physicalDevice VkPhysicalDevice, pProperties *VkPhysicalDeviceProperties2) error
	VkGetPhysicalDeviceFormatProperties2(physicalDevice VkPhysicalDevice, format VkFormat, pFormatProperties *VkFormatProperties2) error
	VkGetPhysicalDeviceImageFormatProperties2(physicalDevice VkPhysicalDevice, pImageFormatInfo *VkPhysicalDeviceImageFormatInfo2, pImageFormatProperties *VkImageFormatProperties2) (VkResult, error)
	VkGetPhysicalDeviceQueueFamilyProperties2(physicalDevice VkPhysicalDevice, pQueueFamilyPropertyCount *Uint32, pQueueFamilyProperties *VkQueueFamilyProperties2) error
	VkGetPhysicalDeviceMemoryProperties2(physicalDevice VkPhysicalDevice, pMemoryProperties *VkPhysicalDeviceMemoryProperties2) error
	VkGetPhysicalDeviceSparseImageFormatProperties2(physicalDevice VkPhysicalDevice, pFormatInfo *VkPhysicalDeviceSparseImageFormatInfo2, pPropertyCount *Uint32, pProperties *VkSparseImageFormatProperties2) error
	VkGetPhysicalDeviceExternalBufferProperties(physicalDevice VkPhysicalDevice, pExternalBufferInfo *VkPhysicalDeviceExternalBufferInfo, pExternalBufferProperties *VkExternalBufferProperties) error
	VkGetPhysicalDeviceExternalFenceProperties(physicalDevice VkPhysicalDevice, pExternalFenceInfo *VkPhysicalDeviceExternalFenceInfo, pExternalFenceProperties *VkExternalFenceProperties) error
	VkGetPhysicalDeviceExternalSemaphoreProperties(physicalDevice VkPhysicalDevice, pExternalSemaphoreInfo *VkPhysicalDeviceExternalSemaphoreInfo, pExternalSemaphoreProperties *VkExternalSemaphoreProperties) error

	//Device
	VkDestroyDevice(device VkDevice, pAllocator *VkAllocationCallbacks) error
	VkGetDeviceQueue(device VkDevice, queueFamilyIndex Uint32, queueIndex Uint32, pQueue *VkQueue) error
	VkQueueSubmit(queue VkQueue, submitCount Uint32, pSubmits *VkSubmitInfo, fence VkFence) (VkResult, error)
	VkQueueWaitIdle(queue VkQueue) (VkResult, error)
	VkDeviceWaitIdle(device VkDevice) (VkResult, error)
	VkAllocateMemory(device VkDevice, pAllocateInfo *VkMemoryAllocateInfo, pAllocator *VkAllocationCallbacks, pMemory *VkDeviceMemory) (VkResult, error)
	VkFreeMemory(device VkDevice, memory VkDeviceMemory, pAllocator *VkAllocationCallbacks) error
	VkMapMemory(device VkDevice, memory VkDeviceMemory, offset VkDeviceSize, size VkDeviceSize, flags VkMemoryMapFlags, ppData *unsafe.Pointer) (VkResult, error)
	VkUnmapMemory(device VkDevice, memory VkDeviceMemory) error
	VkFlushMappedMemoryRanges(device VkDevice, memoryRangeCount Uint32, pMemoryRanges *VkMappedMemoryRange) (VkResult, error)
	VkInvalidateMappedMemoryRanges(device VkDevice, memoryRangeCount Uint32, pMemoryRanges *VkMappedMemoryRange) (VkResult, error)
	VkGetDeviceMemoryCommitment(device VkDevice, memory VkDeviceMemory, pCommittedMemoryInBytes *VkDeviceSize) error
	VkBindBufferMemory(device VkDevice, buffer VkBuffer, memory VkDeviceMemory, memoryOffset VkDeviceSize) (VkResult, error)
	VkBindImageMemory(device VkDevice, image VkImage, memory VkDeviceMemory, memoryOffset VkDeviceSize) (VkResult, error)
	VkGetBufferMemoryRequirements(device VkDevice, buffer VkBuffer, pMemoryRequirements *VkMemoryRequirements) error
	VkGetImageMemoryRequirements(device VkDevice, image VkImage, pMemoryRequirements *VkMemoryRequirements) error
	VkGetImageSparseMemoryRequirements(device VkDevice, image VkImage, pSparseMemoryRequirementCount *Uint32, pSparseMemoryRequirements *VkSparseImageMemoryRequirements) error
	VkQueueBindSparse(queue VkQueue, bindInfoCount Uint32, pBindInfo *VkBindSparseInfo, fence VkFence) (VkResult, error)
	VkCreateFence(device VkDevice, pCreateInfo *VkFenceCreateInfo, pAllocator *VkAllocationCallbacks, pFence *VkFence) (VkResult, error)
	VkDestroyFence(device VkDevice, fence VkFence, pAllocator *VkAllocationCallbacks) error
	VkResetFences(device VkDevice, fenceCount Uint32, pFences *VkFence) (VkResult, error)
	VkGetFenceStatus(device VkDevice, fence VkFence) (VkResult, error)
	VkWaitForFences(device VkDevice, fenceCount Uint32, pFences *VkFence, waitAll VkBool32, timeout Uint64) (VkResult, error)
	VkCreateSemaphore(device VkDevice, pCreateInfo *VkSemaphoreCreateInfo, pAllocator *VkAllocationCallbacks, pSemaphore *VkSemaphore) (VkResult, error)
	VkDestroySemaphore(device VkDevice, semaphore VkSemaphore, pAllocator *VkAllocationCallbacks) error
	VkCreateEvent(device VkDevice, pCreateInfo *VkEventCreateInfo, pAllocator *VkAllocationCallbacks, pEvent *VkEvent) (VkResult, error)
	VkDestroyEvent(device VkDevice, event VkEvent, pAllocator *VkAllocationCallbacks) error
	VkGetEventStatus(device VkDevice, event VkEvent) (VkResult, error)
	VkSetEvent(device VkDevice, event VkEvent) (VkResult, error)
	VkResetEvent(device VkDevice, event VkEvent) (VkResult, error)
	VkCreateQueryPool(device VkDevice, pCreateInfo *VkQueryPoolCreateInfo, pAllocator *VkAllocationCallbacks, pQueryPool *VkQueryPool) (VkResult, error)
	VkDestroyQueryPool(device VkDevice, queryPool VkQueryPool, pAllocator *VkAllocationCallbacks) error
	VkGetQueryPoolResults(device VkDevice, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32, dataSize Size, pData unsafe.Pointer, stride VkDeviceSize, flags VkQueryResultFlags) (VkResult, error)
	VkCreateBuffer(device VkDevice, pCreateInfo *VkBufferCreateInfo, pAllocator *VkAllocationCallbacks, pBuffer *VkBuffer) (VkResult, error)
	VkDestroyBuffer(device VkDevice, buffer VkBuffer, pAllocator *VkAllocationCallbacks) error
	VkCreateBufferView(device VkDevice, pCreateInfo *VkBufferViewCreateInfo, pAllocator *VkAllocationCallbacks, pView *VkBufferView) (VkResult, error)
	VkDestroyBufferView(device VkDevice, bufferView VkBufferView, pAllocator *VkAllocationCallbacks) error
	VkCreateImage(device VkDevice, pCreateInfo *VkImageCreateInfo, pAllocator *VkAllocationCallbacks, pImage *VkImage) (VkResult, error)
	VkDestroyImage(device VkDevice, image VkImage, pAllocator *VkAllocationCallbacks) error
	VkGetImageSubresourceLayout(device VkDevice, image VkImage, pSubresource *VkImageSubresource, pLayout *VkSubresourceLayout) error
	VkCreateImageView(device VkDevice, pCreateInfo *VkImageViewCreateInfo, pAllocator *VkAllocationCallbacks, pView *VkImageView) (VkResult, error)
	VkDestroyImageView(device VkDevice, imageView VkImageView, pAllocator *VkAllocationCallbacks) error
	VkCreateShaderModule(device VkDevice, pCreateInfo *VkShaderModuleCreateInfo, pAllocator *VkAllocationCallbacks, pShaderModule *VkShaderModule) (VkResult, error)
	VkDestroyShaderModule(device VkDevice, shaderModule VkShaderModule, pAllocator *VkAllocationCallbacks) error
	VkCreatePipelineCache(device VkDevice, pCreateInfo *VkPipelineCacheCreateInfo, pAllocator *VkAllocationCallbacks, pPipelineCache *VkPipelineCache) (VkResult, error)
	VkDestroyPipelineCache(device VkDevice, pipelineCache VkPipelineCache, pAllocator *VkAllocationCallbacks) error
	VkGetPipelineCacheData(device VkDevice, pipelineCache VkPipelineCache, pDataSize *Size, pData unsafe.Pointer) (VkResult, error)
	VkMergePipelineCaches(device VkDevice, dstCache VkPipelineCache, srcCacheCount Uint32, pSrcCaches *VkPipelineCache) (VkResult, error)
	VkCreateGraphicsPipelines(device VkDevice, pipelineCache VkPipelineCache, createInfoCount Uint32, pCreateInfos *VkGraphicsPipelineCreateInfo, pAllocator *VkAllocationCallbacks, pPipelines *VkPipeline) (VkResult, error)
	VkCreateComputePipelines(device VkDevice, pipelineCache VkPipelineCache, createInfoCount Uint32, pCreateInfos *VkComputePipelineCreateInfo, pAllocator *VkAllocationCallbacks, pPipelines *VkPipeline) (VkResult, error)
	VkDestroyPipeline(device VkDevice, pipeline VkPipeline, pAllocator *VkAllocationCallbacks) error
	VkCreatePipelineLayout(device VkDevice, pCreateInfo *VkPipelineLayoutCreateInfo, pAllocator *VkAllocationCallbacks, pPipelineLayout *VkPipelineLayout) (VkResult, error)
	VkDestroyPipelineLayout(device VkDevice, pipelineLayout VkPipelineLayout, pAllocator *VkAllocationCallbacks) error
	VkCreateSampler(device VkDevice, pCreateInfo *VkSamplerCreateInfo, pAllocator *VkAllocationCallbacks, pSampler *VkSampler) (VkResult, error)
	VkDestroySampler(device VkDevice, sampler VkSampler, pAllocator *VkAllocationCallbacks) error
	VkCreateDescriptorSetLayout(device VkDevice, pCreateInfo *VkDescriptorSetLayoutCreateInfo, pAllocator *VkAllocationCallbacks, pSetLayout *VkDescriptorSetLayout) (VkResult, error)
	VkDestroyDescriptorSetLayout(device VkDevice, descriptorSetLayout VkDescriptorSetLayout, pAllocator *VkAllocationCallbacks) error
	VkCreateDescriptorPool(device VkDevice, pCreateInfo *VkDescriptorPoolCreateInfo, pAllocator *VkAllocationCallbacks, pDescriptorPool *VkDescriptorPool) (VkResult, error)
	VkDestroyDescriptorPool(device VkDevice, descriptorPool VkDescriptorPool, pAllocator *VkAllocationCallbacks) error
	VkResetDescriptorPool(device VkDevice, descriptorPool VkDescriptorPool, flags VkDescriptorPoolResetFlags) (VkResult, error)
	VkAllocateDescriptorSets(device VkDevice, pAllocateInfo *VkDescriptorSetAllocateInfo, pDescriptorSets *VkDescriptorSet) (VkResult, error)
	VkFreeDescriptorSets(device VkDevice, descriptorPool VkDescriptorPool, descriptorSetCount Uint32, pDescriptorSets *VkDescriptorSet) (VkResult, error)
	VkUpdateDescriptorSets(device VkDevice, descriptorWriteCount Uint32, pDescriptorWrites *VkWriteDescriptorSet, descriptorCopyCount Uint32, pDescriptorCopies *VkCopyDescriptorSet) error
	VkCreateFramebuffer(device VkDevice, pCreateInfo *VkFramebufferCreateInfo, pAllocator *VkAllocationCallbacks, pFramebuffer *VkFramebuffer) (VkResult, error)
	VkDestroyFramebuffer(device VkDevice, framebuffer VkFramebuffer, pAllocator *VkAllocationCallbacks) error
	VkCreateRenderPass(device VkDevice, pCreateInfo *VkRenderPassCreateInfo, pAllocator *VkAllocationCallbacks, pRenderPass *VkRenderPass) (VkResult, error)
	VkDestroyRenderPass(device VkDevice, renderPass VkRenderPass, pAllocator *VkAllocationCallbacks) error
	VkGetRenderAreaGranularity(device VkDevice, renderPass VkRenderPass, pGranularity *VkExtent2D) error
	VkCreateCommandPool(device VkDevice, pCreateInfo *VkCommandPoolCreateInfo, pAllocator *VkAllocationCallbacks, pCommandPool *VkCommandPool) (VkResult, error)
	VkDestroyCommandPool(device VkDevice, commandPool VkCommandPool, pAllocator *VkAllocationCallbacks) error
	VkResetCommandPool(device VkDevice, commandPool VkCommandPool, flags VkCommandPoolResetFlags) (VkResult, error)
	VkAllocateCommandBuffers(device VkDevice, pAllocateInfo *VkCommandBufferAllocateInfo, pCommandBuffers *VkCommandBuffer) (VkResult, error)
	VkFreeCommandBuffers(device VkDevice, commandPool VkCommandPool, commandBufferCount Uint32, pCommandBuffers *VkCommandBuffer) error
	VkBeginCommandBuffer(commandBuffer VkCommandBuffer, pBeginInfo *VkCommandBufferBeginInfo) (VkResult, error)
	VkEndCommandBuffer(commandBuffer VkCommandBuffer) (VkResult, error)
	VkResetCommandBuffer(commandBuffer VkCommandBuffer, flags VkCommandBufferResetFlags) (VkResult, error)
	VkCmdBindPipeline(commandBuffer VkCommandBuffer, pipelineBindPoint VkPipelineBindPoint, pipeline VkPipeline) error
	VkCmdSetViewport(commandBuffer VkCommandBuffer, firstViewport Uint32, viewportCount Uint32, pViewports *VkViewport) error
	VkCmdSetScissor(commandBuffer VkCommandBuffer, firstScissor Uint32, scissorCount Uint32, pScissors *VkRect2D) error
	VkCmdSetLineWidth(commandBuffer VkCommandBuffer, lineWidth Float) error
	VkCmdSetDepthBias(commandBuffer VkCommandBuffer, depthBiasConstantFactor Float, depthBiasClamp Float, depthBiasSlopeFactor Float) error
	VkCmdSetBlendConstants(commandBuffer VkCommandBuffer, blendConstants *Float) error
	VkCmdSetDepthBounds(commandBuffer VkCommandBuffer, minDepthBounds Float, maxDepthBounds Float) error
	VkCmdSetStencilCompareMask(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, compareMask Uint32) error
	VkCmdSetStencilWriteMask(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, writeMask Uint32) error
	VkCmdSetStencilReference(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, reference Uint32) error
	VkCmdBindDescriptorSets(commandBuffer VkCommandBuffer, pipelineBindPoint VkPipelineBindPoint, layout VkPipelineLayout, firstSet Uint32, descriptorSetCount Uint32, pDescriptorSets *VkDescriptorSet, dynamicOffsetCount Uint32, pDynamicOffsets *Uint32) error
	VkCmdBindIndexBuffer(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, indexType VkIndexType) error
	VkCmdBindVertexBuffers(commandBuffer VkCommandBuffer, firstBinding Uint32, bindingCount Uint32, pBuffers *VkBuffer, pOffsets *VkDeviceSize) error
	VkCmdDraw(commandBuffer VkCommandBuffer, vertexCount Uint32, instanceCount Uint32, firstVertex Uint32, firstInstance Uint32) error
	VkCmdDrawIndexed(commandBuffer VkCommandBuffer, indexCount Uint32, instanceCount Uint32, firstIndex Uint32, vertexOffset Int32, firstInstance Uint32) error
	VkCmdDrawIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, drawCount Uint32, stride Uint32) error
	VkCmdDrawIndexedIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, drawCount Uint32, stride Uint32) error
	VkCmdDispatch(commandBuffer VkCommandBuffer, groupCountX Uint32, groupCountY Uint32, groupCountZ Uint32) error
	VkCmdDispatchIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize) error
	VkCmdCopyBuffer(commandBuffer VkCommandBuffer, srcBuffer VkBuffer, dstBuffer VkBuffer, regionCount Uint32, pRegions *VkBufferCopy) error
	VkCmdCopyImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkImageCopy) error
	VkCmdBlitImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkImageBlit, filter VkFilter) error
	VkCmdCopyBufferToImage(commandBuffer VkCommandBuffer, srcBuffer VkBuffer, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkBufferImageCopy) error
	VkCmdCopyImageToBuffer(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstBuffer VkBuffer, regionCount Uint32, pRegions *VkBufferImageCopy) error
	VkCmdUpdateBuffer(commandBuffer VkCommandBuffer, dstBuffer VkBuffer, dstOffset VkDeviceSize, dataSize VkDeviceSize, pData unsafe.Pointer) error
	VkCmdFillBuffer(commandBuffer VkCommandBuffer, dstBuffer VkBuffer, dstOffset VkDeviceSize, size VkDeviceSize, data Uint32) error
	VkCmdClearColorImage(commandBuffer VkCommandBuffer, image VkImage, imageLayout VkImageLayout, pColor *VkClearColorValue, rangeCount Uint32, pRanges *VkImageSubresourceRange) error
	VkCmdClearDepthStencilImage(commandBuffer VkCommandBuffer, image VkImage, imageLayout VkImageLayout, pDepthStencil *VkClearDepthStencilValue, rangeCount Uint32, pRanges *VkImageSubresourceRange) error
	VkCmdClearAttachments(commandBuffer VkCommandBuffer, attachmentCount Uint32, pAttachments *VkClearAttachment, rectCount Uint32, pRects *VkClearRect) error
	VkCmdResolveImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkImageResolve) error
	VkCmdSetEvent(commandBuffer VkCommandBuffer, event VkEvent, stageMask VkPipelineStageFlags) error
	VkCmdResetEvent(commandBuffer VkCommandBuffer, event VkEvent, stageMask VkPipelineStageFlags) error
	VkCmdWaitEvents(commandBuffer VkCommandBuffer, eventCount Uint32, pEvents *VkEvent, srcStageMask VkPipelineStageFlags, dstStageMask VkPipelineStageFlags, memoryBarrierCount Uint32, pMemoryBarriers *VkMemoryBarrier, bufferMemoryBarrierCount Uint32, pBufferMemoryBarriers *VkBufferMemoryBarrier, imageMemoryBarrierCount Uint32, pImageMemoryBarriers *VkImageMemoryBarrier) error
	VkCmdPipelineBarrier(commandBuffer VkCommandBuffer, srcStageMask VkPipelineStageFlags, dstStageMask VkPipelineStageFlags, dependencyFlags VkDependencyFlags, memoryBarrierCount Uint32, pMemoryBarriers *VkMemoryBarrier, bufferMemoryBarrierCount Uint32, pBufferMemoryBarriers *VkBufferMemoryBarrier, imageMemoryBarrierCount Uint32, pImageMemoryBarriers *VkImageMemoryBarrier) error
	VkCmdBeginQuery(commandBuffer VkCommandBuffer, queryPool VkQueryPool, query Uint32, flags VkQueryControlFlags) error
	VkCmdEndQuery(commandBuffer VkCommandBuffer, queryPool VkQueryPool, query Uint32) error
	VkCmdResetQueryPool(commandBuffer VkCommandBuffer, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32) error
	VkCmdWriteTimestamp(commandBuffer VkCommandBuffer, pipelineStage VkPipelineStageFlags, queryPool VkQueryPool, query Uint32) error
	VkCmdCopyQueryPoolResults(commandBuffer VkCommandBuffer, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32, dstBuffer VkBuffer, dstOffset VkDeviceSize, stride VkDeviceSize, flags VkQueryResultFlags) error
	VkCmdPushConstants(commandBuffer VkCommandBuffer, layout VkPipelineLayout, stageFlags VkShaderStageFlags, offset Uint32, size Uint32, pValues unsafe.Pointer) error
	VkCmdBeginRenderPass(commandBuffer VkCommandBuffer, pRenderPassBegin *VkRenderPassBeginInfo, contents VkSubpassContents) error
	VkCmdNextSubpass(commandBuffer VkCommandBuffer, contents VkSubpassContents) error
	VkCmdEndRenderPass(commandBuffer VkCommandBuffer) error
	VkCmdExecuteCommands(commandBuffer VkCommandBuffer, commandBufferCount Uint32, pCommandBuffers *VkCommandBuffer) error

	VkBindBufferMemory2(device VkDevice, bindInfoCount Uint32, pBindInfos *VkBindBufferMemoryInfo) (VkResult, error)
	VkBindImageMemory2(device VkDevice, bindInfoCount Uint32, pBindInfos *VkBindImageMemoryInfo) (VkResult, error)
	VkGetDeviceGroupPeerMemoryFeatures(device VkDevice, heapIndex Uint32, localDeviceIndex Uint32, remoteDeviceIndex Uint32, pPeerMemoryFeatures *VkPeerMemoryFeatureFlags) error
	VkCmdSetDeviceMask(commandBuffer VkCommandBuffer, deviceMask Uint32) error
	VkCmdDispatchBase(commandBuffer VkCommandBuffer, baseGroupX Uint32, baseGroupY Uint32, baseGroupZ Uint32, groupCountX Uint32, groupCountY Uint32, groupCountZ Uint32) error
	VkGetImageMemoryRequirements2(device VkDevice, pInfo *VkImageMemoryRequirementsInfo2, pMemoryRequirements *VkMemoryRequirements2) error
	VkGetBufferMemoryRequirements2(device VkDevice, pInfo *VkBufferMemoryRequirementsInfo2, pMemoryRequirements *VkMemoryRequirements2) error
	VkGetImageSparseMemoryRequirements2(device VkDevice, pInfo *VkImageSparseMemoryRequirementsInfo2, pSparseMemoryRequirementCount *Uint32, pSparseMemoryRequirements *VkSparseImageMemoryRequirements2) error
	VkTrimCommandPool(device VkDevice, commandPool VkCommandPool, flags VkCommandPoolTrimFlags) error
	VkGetDeviceQueue2(device VkDevice, pQueueInfo *VkDeviceQueueInfo2, pQueue *VkQueue) error
	VkCreateSamplerYcbcrConversion(device VkDevice, pCreateInfo *VkSamplerYcbcrConversionCreateInfo, pAllocator *VkAllocationCallbacks, pYcbcrConversion *VkSamplerYcbcrConversion) (VkResult, error)
	VkDestroySamplerYcbcrConversion(device VkDevice, ycbcrConversion VkSamplerYcbcrConversion, pAllocator *VkAllocationCallbacks) error
	VkCreateDescriptorUpdateTemplate(device VkDevice, pCreateInfo *VkDescriptorUpdateTemplateCreateInfo, pAllocator *VkAllocationCallbacks, pDescriptorUpdateTemplate *VkDescriptorUpdateTemplate) (VkResult, error)
	VkDestroyDescriptorUpdateTemplate(device VkDevice, descriptorUpdateTemplate VkDescriptorUpdateTemplate, pAllocator *VkAllocationCallbacks) error
	VkUpdateDescriptorSetWithTemplate(device VkDevice, descriptorSet VkDescriptorSet, descriptorUpdateTemplate VkDescriptorUpdateTemplate, pData unsafe.Pointer) error
	VkGetDescriptorSetLayoutSupport(device VkDevice, pCreateInfo *VkDescriptorSetLayoutCreateInfo, pSupport *VkDescriptorSetLayoutSupport) error

	VkCmdDrawIndirectCount(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, countBuffer VkBuffer, countBufferOffset VkDeviceSize, maxDrawCount Uint32, stride Uint32) error
	VkCmdDrawIndexedIndirectCount(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, countBuffer VkBuffer, countBufferOffset VkDeviceSize, maxDrawCount Uint32, stride Uint32) error
	VkCreateRenderPass2(device VkDevice, pCreateInfo *VkRenderPassCreateInfo2, pAllocator *VkAllocationCallbacks, pRenderPass *VkRenderPass) (VkResult, error)
	VkCmdBeginRenderPass2(commandBuffer VkCommandBuffer, pRenderPassBegin *VkRenderPassBeginInfo, pSubpassBeginInfo *VkSubpassBeginInfo) error
	VkCmdNextSubpass2(commandBuffer VkCommandBuffer, pSubpassBeginInfo *VkSubpassBeginInfo, pSubpassEndInfo *VkSubpassEndInfo) error
	VkCmdEndRenderPass2(commandBuffer VkCommandBuffer, pSubpassEndInfo *VkSubpassEndInfo) error
	VkResetQueryPool(device VkDevice, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32) error
	VkGetSemaphoreCounterValue(device VkDevice, semaphore VkSemaphore, pValue *Uint64) (VkResult, error)
	VkWaitSemaphores(device VkDevice, pWaitInfo *VkSemaphoreWaitInfo, timeout Uint64) (VkResult, error)
	VkSignalSemaphore(device VkDevice, pSignalInfo *VkSemaphoreSignalInfo) (VkResult, error)
	VkGetBufferDeviceAddress(device VkDevice, pInfo *VkBufferDeviceAddressInfo) (VkDeviceAddress, error)
	VkGetBufferOpaqueCaptureAddress(device VkDevice, pInfo *VkBufferDeviceAddressInfo) (Uint64, error)
	VkGetDeviceMemoryOpaqueCaptureAddress(device VkDevice, pInfo *VkDeviceMemoryOpaqueCaptureAddressInfo) (Uint64, error)
}
