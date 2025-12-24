#include "func_ptrs.h"

VkResult cgoCreateInstance(PFN_vkCreateInstance fn, VkInstanceCreateInfo *pCreateInfo, VkAllocationCallbacks *pAllocator, VkInstance *pInstance) {
    return fn(pCreateInfo, pAllocator, pInstance);
}

void cgoDestroyInstance(PFN_vkDestroyInstance fn, VkInstance instance, VkAllocationCallbacks* pAllocator) {
    fn(instance, pAllocator);
}

VkResult cgoEnumeratePhysicalDevices(PFN_vkEnumeratePhysicalDevices fn, VkInstance instance, uint32_t* pPhysicalDeviceCount, VkPhysicalDevice* pPhysicalDevices) {
    return fn(instance, pPhysicalDeviceCount, pPhysicalDevices);
}

void cgoGetPhysicalDeviceFeatures(PFN_vkGetPhysicalDeviceFeatures fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceFeatures* pFeatures) {
    fn(physicalDevice, pFeatures);
}

void cgoGetPhysicalDeviceFormatProperties(PFN_vkGetPhysicalDeviceFormatProperties fn, VkPhysicalDevice physicalDevice, VkFormat format, VkFormatProperties* pFormatProperties) {
    fn(physicalDevice, format, pFormatProperties);
}

VkResult cgoGetPhysicalDeviceImageFormatProperties(PFN_vkGetPhysicalDeviceImageFormatProperties fn, VkPhysicalDevice physicalDevice, VkFormat format, VkImageType type, VkImageTiling tiling, VkImageUsageFlags usage, VkImageCreateFlags flags, VkImageFormatProperties* pImageFormatProperties) {
    return fn(physicalDevice, format, type, tiling, usage, flags, pImageFormatProperties);
}

void cgoGetPhysicalDeviceProperties(PFN_vkGetPhysicalDeviceProperties fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties* pProperties) {
    fn(physicalDevice, pProperties);
}

void cgoGetPhysicalDeviceQueueFamilyProperties(PFN_vkGetPhysicalDeviceQueueFamilyProperties fn, VkPhysicalDevice physicalDevice, uint32_t* pQueueFamilyPropertyCount, VkQueueFamilyProperties* pQueueFamilyProperties) {
    fn(physicalDevice, pQueueFamilyPropertyCount, pQueueFamilyProperties);
}

void cgoGetPhysicalDeviceMemoryProperties(PFN_vkGetPhysicalDeviceMemoryProperties fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceMemoryProperties* pMemoryProperties) {
    fn(physicalDevice, pMemoryProperties);
}

VkResult cgoCreateDevice(PFN_vkCreateDevice fn, VkPhysicalDevice physicalDevice, VkDeviceCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkDevice* pDevice) {
    return fn(physicalDevice, pCreateInfo, pAllocator, pDevice);
}

void cgoDestroyDevice(PFN_vkDestroyDevice fn, VkDevice device, VkAllocationCallbacks* pAllocator) {
    fn(device, pAllocator);
}

VkResult cgoEnumerateInstanceExtensionProperties(PFN_vkEnumerateInstanceExtensionProperties fn, char* pLayerName, uint32_t* pPropertyCount, VkExtensionProperties* pProperties) {
    return fn(pLayerName, pPropertyCount, pProperties);
}

VkResult cgoEnumerateInstanceLayerProperties(PFN_vkEnumerateInstanceLayerProperties fn, uint32_t* pPropertyCount, VkLayerProperties* pProperties) {
    return fn(pPropertyCount, pProperties);
}

VkResult cgoEnumerateDeviceExtensionProperties(PFN_vkEnumerateDeviceExtensionProperties fn, VkPhysicalDevice physicalDevice, char* pLayerName, uint32_t* pPropertyCount, VkExtensionProperties* pProperties) {
    return fn(physicalDevice, pLayerName, pPropertyCount, pProperties);
}

VkResult cgoEnumerateDeviceLayerProperties(PFN_vkEnumerateDeviceLayerProperties fn, VkPhysicalDevice physicalDevice, uint32_t* pPropertyCount, VkLayerProperties* pProperties) {
    return fn(physicalDevice, pPropertyCount, pProperties);
}

void cgoGetDeviceQueue(PFN_vkGetDeviceQueue fn, VkDevice device, uint32_t queueFamilyIndex, uint32_t queueIndex, VkQueue* pQueue) {
    fn(device, queueFamilyIndex, queueIndex, pQueue);
}

VkResult cgoQueueSubmit(PFN_vkQueueSubmit fn, VkQueue queue, uint32_t submitCount, VkSubmitInfo* pSubmits, VkFence fence) {
    return fn(queue, submitCount, pSubmits, fence);
}

VkResult cgoQueueWaitIdle(PFN_vkQueueWaitIdle fn, VkQueue queue) {
    return fn(queue);
}

VkResult cgoDeviceWaitIdle(PFN_vkDeviceWaitIdle fn, VkDevice device) {
    return fn(device);
}

VkResult cgoAllocateMemory(PFN_vkAllocateMemory fn, VkDevice device, VkMemoryAllocateInfo* pAllocateInfo, VkAllocationCallbacks* pAllocator, VkDeviceMemory* pMemory) {
    return fn(device, pAllocateInfo, pAllocator, pMemory);
}

void cgoFreeMemory(PFN_vkFreeMemory fn, VkDevice device, VkDeviceMemory memory, VkAllocationCallbacks* pAllocator) {
    fn(device, memory, pAllocator);
}

VkResult cgoMapMemory(PFN_vkMapMemory fn, VkDevice device, VkDeviceMemory memory, VkDeviceSize offset, VkDeviceSize size, VkMemoryMapFlags flags, void** ppData) {
    return fn(device, memory, offset, size, flags, ppData);
}

void cgoUnmapMemory(PFN_vkUnmapMemory fn, VkDevice device, VkDeviceMemory memory) {
    fn(device, memory);
}

VkResult cgoFlushMappedMemoryRanges(PFN_vkFlushMappedMemoryRanges fn, VkDevice device, uint32_t memoryRangeCount, VkMappedMemoryRange* pMemoryRanges) {
    return fn(device, memoryRangeCount, pMemoryRanges);
}

VkResult cgoInvalidateMappedMemoryRanges(PFN_vkInvalidateMappedMemoryRanges fn, VkDevice device, uint32_t memoryRangeCount, VkMappedMemoryRange* pMemoryRanges) {
    return fn(device, memoryRangeCount, pMemoryRanges);
}

void cgoGetDeviceMemoryCommitment(PFN_vkGetDeviceMemoryCommitment fn, VkDevice device, VkDeviceMemory memory, VkDeviceSize* pCommittedMemoryInBytes) {
    fn(device, memory, pCommittedMemoryInBytes);
}

VkResult cgoBindBufferMemory(PFN_vkBindBufferMemory fn, VkDevice device, VkBuffer buffer, VkDeviceMemory memory, VkDeviceSize memoryOffset) {
    return fn(device, buffer, memory, memoryOffset);
}

VkResult cgoBindImageMemory(PFN_vkBindImageMemory fn, VkDevice device, VkImage image, VkDeviceMemory memory, VkDeviceSize memoryOffset) {
    return fn(device, image, memory, memoryOffset);
}

void cgoGetBufferMemoryRequirements(PFN_vkGetBufferMemoryRequirements fn, VkDevice device, VkBuffer buffer, VkMemoryRequirements* pMemoryRequirements) {
    fn(device, buffer, pMemoryRequirements);
}

void cgoGetImageMemoryRequirements(PFN_vkGetImageMemoryRequirements fn, VkDevice device, VkImage image, VkMemoryRequirements* pMemoryRequirements) {
    fn(device, image, pMemoryRequirements);
}

void cgoGetImageSparseMemoryRequirements(PFN_vkGetImageSparseMemoryRequirements fn, VkDevice device, VkImage image, uint32_t* pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements* pSparseMemoryRequirements) {
    fn(device, image, pSparseMemoryRequirementCount, pSparseMemoryRequirements);
}

void cgoGetPhysicalDeviceSparseImageFormatProperties(PFN_vkGetPhysicalDeviceSparseImageFormatProperties fn, VkPhysicalDevice physicalDevice, VkFormat format, VkImageType type, VkSampleCountFlagBits samples, VkImageUsageFlags usage, VkImageTiling tiling, uint32_t* pPropertyCount, VkSparseImageFormatProperties* pProperties) {
    fn(physicalDevice, format, type, samples, usage, tiling, pPropertyCount, pProperties);
}

VkResult cgoQueueBindSparse(PFN_vkQueueBindSparse fn, VkQueue queue, uint32_t bindInfoCount, VkBindSparseInfo* pBindInfo, VkFence fence) {
    return fn(queue, bindInfoCount, pBindInfo, fence);
}

VkResult cgoCreateFence(PFN_vkCreateFence fn, VkDevice device, VkFenceCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkFence* pFence) {
    return fn(device, pCreateInfo, pAllocator, pFence);
}

void cgoDestroyFence(PFN_vkDestroyFence fn, VkDevice device, VkFence fence, VkAllocationCallbacks* pAllocator) {
    fn(device, fence, pAllocator);
}

VkResult cgoResetFences(PFN_vkResetFences fn, VkDevice device, uint32_t fenceCount, VkFence* pFences) {
    return fn(device, fenceCount, pFences);
}

VkResult cgoGetFenceStatus(PFN_vkGetFenceStatus fn, VkDevice device, VkFence fence) {
    return fn(device, fence);
}

VkResult cgoWaitForFences(PFN_vkWaitForFences fn, VkDevice device, uint32_t fenceCount, VkFence* pFences, VkBool32 waitAll, uint64_t timeout) {
    return fn(device, fenceCount, pFences, waitAll, timeout);
}

VkResult cgoCreateSemaphore(PFN_vkCreateSemaphore fn, VkDevice device, VkSemaphoreCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkSemaphore* pSemaphore) {
    return fn(device, pCreateInfo, pAllocator, pSemaphore);
}

void cgoDestroySemaphore(PFN_vkDestroySemaphore fn, VkDevice device, VkSemaphore semaphore, VkAllocationCallbacks* pAllocator) {
    fn(device, semaphore, pAllocator);
}

VkResult cgoCreateEvent(PFN_vkCreateEvent fn, VkDevice device, VkEventCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkEvent* pEvent) {
    return fn(device, pCreateInfo, pAllocator, pEvent);
}

void cgoDestroyEvent(PFN_vkDestroyEvent fn, VkDevice device, VkEvent event, VkAllocationCallbacks* pAllocator) {
    fn(device, event, pAllocator);
}

VkResult cgoGetEventStatus(PFN_vkGetEventStatus fn, VkDevice device, VkEvent event) {
    return fn(device, event);
}

VkResult cgoSetEvent(PFN_vkSetEvent fn, VkDevice device, VkEvent event) {
    return fn(device, event);
}

VkResult cgoResetEvent(PFN_vkResetEvent fn, VkDevice device, VkEvent event) {
    return fn(device, event);
}

VkResult cgoCreateQueryPool(PFN_vkCreateQueryPool fn, VkDevice device, VkQueryPoolCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkQueryPool* pQueryPool) {
    return fn(device, pCreateInfo, pAllocator, pQueryPool);
}

void cgoDestroyQueryPool(PFN_vkDestroyQueryPool fn, VkDevice device, VkQueryPool queryPool, VkAllocationCallbacks* pAllocator) {
    fn(device, queryPool, pAllocator);
}

VkResult cgoGetQueryPoolResults(PFN_vkGetQueryPoolResults fn, VkDevice device, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount, size_t dataSize, void* pData, VkDeviceSize stride, VkQueryResultFlags flags) {
    return fn(device, queryPool, firstQuery, queryCount, dataSize, pData, stride, flags);
}

VkResult cgoCreateBuffer(PFN_vkCreateBuffer fn, VkDevice device, VkBufferCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkBuffer* pBuffer) {
    return fn(device, pCreateInfo, pAllocator, pBuffer);
}

void cgoDestroyBuffer(PFN_vkDestroyBuffer fn, VkDevice device, VkBuffer buffer, VkAllocationCallbacks* pAllocator) {
    fn(device, buffer, pAllocator);
}

VkResult cgoCreateBufferView(PFN_vkCreateBufferView fn, VkDevice device, VkBufferViewCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkBufferView* pView) {
    return fn(device, pCreateInfo, pAllocator, pView);
}

void cgoDestroyBufferView(PFN_vkDestroyBufferView fn, VkDevice device, VkBufferView bufferView, VkAllocationCallbacks* pAllocator) {
    fn(device, bufferView, pAllocator);
}

VkResult cgoCreateImage(PFN_vkCreateImage fn, VkDevice device, VkImageCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkImage* pImage) {
    return fn(device, pCreateInfo, pAllocator, pImage);
}

void cgoDestroyImage(PFN_vkDestroyImage fn, VkDevice device, VkImage image, VkAllocationCallbacks* pAllocator) {
    fn(device, image, pAllocator);
}

void cgoGetImageSubresourceLayout(PFN_vkGetImageSubresourceLayout fn, VkDevice device, VkImage image, VkImageSubresource* pSubresource, VkSubresourceLayout* pLayout) {
    fn(device, image, pSubresource, pLayout);
}

VkResult cgoCreateImageView(PFN_vkCreateImageView fn, VkDevice device, VkImageViewCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkImageView* pView) {
    return fn(device, pCreateInfo, pAllocator, pView);
}

void cgoDestroyImageView(PFN_vkDestroyImageView fn, VkDevice device, VkImageView imageView, VkAllocationCallbacks* pAllocator) {
    fn(device, imageView, pAllocator);
}

VkResult cgoCreateShaderModule(PFN_vkCreateShaderModule fn, VkDevice device, VkShaderModuleCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkShaderModule* pShaderModule) {
    return fn(device, pCreateInfo, pAllocator, pShaderModule);
}

void cgoDestroyShaderModule(PFN_vkDestroyShaderModule fn, VkDevice device, VkShaderModule shaderModule, VkAllocationCallbacks* pAllocator) {
    fn(device, shaderModule, pAllocator);
}

VkResult cgoCreatePipelineCache(PFN_vkCreatePipelineCache fn, VkDevice device, VkPipelineCacheCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkPipelineCache* pPipelineCache) {
    return fn(device, pCreateInfo, pAllocator, pPipelineCache);
}

void cgoDestroyPipelineCache(PFN_vkDestroyPipelineCache fn, VkDevice device, VkPipelineCache pipelineCache, VkAllocationCallbacks* pAllocator) {
    fn(device, pipelineCache, pAllocator);
}

VkResult cgoGetPipelineCacheData(PFN_vkGetPipelineCacheData fn, VkDevice device, VkPipelineCache pipelineCache, size_t* pDataSize, void* pData) {
    return fn(device, pipelineCache, pDataSize, pData);
}

VkResult cgoMergePipelineCaches(PFN_vkMergePipelineCaches fn, VkDevice device, VkPipelineCache dstCache, uint32_t srcCacheCount, VkPipelineCache* pSrcCaches) {
    return fn(device, dstCache, srcCacheCount, pSrcCaches);
}

VkResult cgoCreateGraphicsPipelines(PFN_vkCreateGraphicsPipelines fn, VkDevice device, VkPipelineCache pipelineCache, uint32_t createInfoCount, VkGraphicsPipelineCreateInfo* pCreateInfos, VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines) {
    return fn(device, pipelineCache, createInfoCount, pCreateInfos, pAllocator, pPipelines);
}

VkResult cgoCreateComputePipelines(PFN_vkCreateComputePipelines fn, VkDevice device, VkPipelineCache pipelineCache, uint32_t createInfoCount, VkComputePipelineCreateInfo* pCreateInfos, VkAllocationCallbacks* pAllocator, VkPipeline* pPipelines) {
    return fn(device, pipelineCache, createInfoCount, pCreateInfos, pAllocator, pPipelines);
}

void cgoDestroyPipeline(PFN_vkDestroyPipeline fn, VkDevice device, VkPipeline pipeline, VkAllocationCallbacks* pAllocator) {
    fn(device, pipeline, pAllocator);
}

VkResult cgoCreatePipelineLayout(PFN_vkCreatePipelineLayout fn, VkDevice device, VkPipelineLayoutCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkPipelineLayout* pPipelineLayout) {
    return fn(device, pCreateInfo, pAllocator, pPipelineLayout);
}

void cgoDestroyPipelineLayout(PFN_vkDestroyPipelineLayout fn, VkDevice device, VkPipelineLayout pipelineLayout, VkAllocationCallbacks* pAllocator) {
    fn(device, pipelineLayout, pAllocator);
}

VkResult cgoCreateSampler(PFN_vkCreateSampler fn, VkDevice device, VkSamplerCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkSampler* pSampler) {
    return fn(device, pCreateInfo, pAllocator, pSampler);
}

void cgoDestroySampler(PFN_vkDestroySampler fn, VkDevice device, VkSampler sampler, VkAllocationCallbacks* pAllocator) {
    fn(device, sampler, pAllocator);
}

VkResult cgoCreateDescriptorSetLayout(PFN_vkCreateDescriptorSetLayout fn, VkDevice device, VkDescriptorSetLayoutCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkDescriptorSetLayout* pSetLayout) {
    return fn(device, pCreateInfo, pAllocator, pSetLayout);
}

void cgoDestroyDescriptorSetLayout(PFN_vkDestroyDescriptorSetLayout fn, VkDevice device, VkDescriptorSetLayout descriptorSetLayout, VkAllocationCallbacks* pAllocator) {
    fn(device, descriptorSetLayout, pAllocator);
}

VkResult cgoCreateDescriptorPool(PFN_vkCreateDescriptorPool fn, VkDevice device, VkDescriptorPoolCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkDescriptorPool* pDescriptorPool) {
    return fn(device, pCreateInfo, pAllocator, pDescriptorPool);
}

void cgoDestroyDescriptorPool(PFN_vkDestroyDescriptorPool fn, VkDevice device, VkDescriptorPool descriptorPool, VkAllocationCallbacks* pAllocator) {
    fn(device, descriptorPool, pAllocator);
}

VkResult cgoResetDescriptorPool(PFN_vkResetDescriptorPool fn, VkDevice device, VkDescriptorPool descriptorPool, VkDescriptorPoolResetFlags flags) {
    return fn(device, descriptorPool, flags);
}

VkResult cgoAllocateDescriptorSets(PFN_vkAllocateDescriptorSets fn, VkDevice device, VkDescriptorSetAllocateInfo* pAllocateInfo, VkDescriptorSet* pDescriptorSets) {
    return fn(device, pAllocateInfo, pDescriptorSets);
}

VkResult cgoFreeDescriptorSets(PFN_vkFreeDescriptorSets fn, VkDevice device, VkDescriptorPool descriptorPool, uint32_t descriptorSetCount, VkDescriptorSet* pDescriptorSets) {
    return fn(device, descriptorPool, descriptorSetCount, pDescriptorSets);
}

void cgoUpdateDescriptorSets(PFN_vkUpdateDescriptorSets fn, VkDevice device, uint32_t descriptorWriteCount, VkWriteDescriptorSet* pDescriptorWrites, uint32_t descriptorCopyCount, VkCopyDescriptorSet* pDescriptorCopies) {
    fn(device, descriptorWriteCount, pDescriptorWrites, descriptorCopyCount, pDescriptorCopies);
}

VkResult cgoCreateFramebuffer(PFN_vkCreateFramebuffer fn, VkDevice device, VkFramebufferCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkFramebuffer* pFramebuffer) {
    return fn(device, pCreateInfo, pAllocator, pFramebuffer);
}

void cgoDestroyFramebuffer(PFN_vkDestroyFramebuffer fn, VkDevice device, VkFramebuffer framebuffer, VkAllocationCallbacks* pAllocator) {
    fn(device, framebuffer, pAllocator);
}

VkResult cgoCreateRenderPass(PFN_vkCreateRenderPass fn, VkDevice device, VkRenderPassCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkRenderPass* pRenderPass) {
    return fn(device, pCreateInfo, pAllocator, pRenderPass);
}

void cgoDestroyRenderPass(PFN_vkDestroyRenderPass fn, VkDevice device, VkRenderPass renderPass, VkAllocationCallbacks* pAllocator) {
    fn(device, renderPass, pAllocator);
}

void cgoGetRenderAreaGranularity(PFN_vkGetRenderAreaGranularity fn, VkDevice device, VkRenderPass renderPass, VkExtent2D* pGranularity) {
    fn(device, renderPass, pGranularity);
}

VkResult cgoCreateCommandPool(PFN_vkCreateCommandPool fn, VkDevice device, VkCommandPoolCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkCommandPool* pCommandPool) {
    return fn(device, pCreateInfo, pAllocator, pCommandPool);
}

void cgoDestroyCommandPool(PFN_vkDestroyCommandPool fn, VkDevice device, VkCommandPool commandPool, VkAllocationCallbacks* pAllocator) {
    fn(device, commandPool, pAllocator);
}

VkResult cgoResetCommandPool(PFN_vkResetCommandPool fn, VkDevice device, VkCommandPool commandPool, VkCommandPoolResetFlags flags) {
    return fn(device, commandPool, flags);
}

VkResult cgoAllocateCommandBuffers(PFN_vkAllocateCommandBuffers fn, VkDevice device, VkCommandBufferAllocateInfo* pAllocateInfo, VkCommandBuffer* pCommandBuffers) {
    return fn(device, pAllocateInfo, pCommandBuffers);
}

void cgoFreeCommandBuffers(PFN_vkFreeCommandBuffers fn, VkDevice device, VkCommandPool commandPool, uint32_t commandBufferCount, VkCommandBuffer* pCommandBuffers) {
    fn(device, commandPool, commandBufferCount, pCommandBuffers);
}

VkResult cgoBeginCommandBuffer(PFN_vkBeginCommandBuffer fn, VkCommandBuffer commandBuffer, VkCommandBufferBeginInfo* pBeginInfo) {
    return fn(commandBuffer, pBeginInfo);
}

VkResult cgoEndCommandBuffer(PFN_vkEndCommandBuffer fn, VkCommandBuffer commandBuffer) {
    return fn(commandBuffer);
}

VkResult cgoResetCommandBuffer(PFN_vkResetCommandBuffer fn, VkCommandBuffer commandBuffer, VkCommandBufferResetFlags flags) {
    return fn(commandBuffer, flags);
}

void cgoCmdBindPipeline(PFN_vkCmdBindPipeline fn, VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipeline pipeline) {
    fn(commandBuffer, pipelineBindPoint, pipeline);
}

void cgoCmdSetViewport(PFN_vkCmdSetViewport fn, VkCommandBuffer commandBuffer, uint32_t firstViewport, uint32_t viewportCount, VkViewport* pViewports) {
    fn(commandBuffer, firstViewport, viewportCount, pViewports);
}

void cgoCmdSetScissor(PFN_vkCmdSetScissor fn, VkCommandBuffer commandBuffer, uint32_t firstScissor, uint32_t scissorCount, VkRect2D* pScissors) {
    fn(commandBuffer, firstScissor, scissorCount, pScissors);
}

void cgoCmdSetLineWidth(PFN_vkCmdSetLineWidth fn, VkCommandBuffer commandBuffer, float lineWidth) {
    fn(commandBuffer, lineWidth);
}

void cgoCmdSetDepthBias(PFN_vkCmdSetDepthBias fn, VkCommandBuffer commandBuffer, float depthBiasConstantFactor, float depthBiasClamp, float depthBiasSlopeFactor) {
    fn(commandBuffer, depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor);
}

void cgoCmdSetBlendConstants(PFN_vkCmdSetBlendConstants fn, VkCommandBuffer commandBuffer, float blendConstants[4]) {
    fn(commandBuffer, blendConstants);
}

void cgoCmdSetDepthBounds(PFN_vkCmdSetDepthBounds fn, VkCommandBuffer commandBuffer, float minDepthBounds, float maxDepthBounds) {
    fn(commandBuffer, minDepthBounds, maxDepthBounds);
}

void cgoCmdSetStencilCompareMask(PFN_vkCmdSetStencilCompareMask fn, VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t compareMask) {
    fn(commandBuffer, faceMask, compareMask);
}

void cgoCmdSetStencilWriteMask(PFN_vkCmdSetStencilWriteMask fn, VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t writeMask) {
    fn(commandBuffer, faceMask, writeMask);
}

void cgoCmdSetStencilReference(PFN_vkCmdSetStencilReference fn, VkCommandBuffer commandBuffer, VkStencilFaceFlags faceMask, uint32_t reference) {
    fn(commandBuffer, faceMask, reference);
}

void cgoCmdBindDescriptorSets(PFN_vkCmdBindDescriptorSets fn, VkCommandBuffer commandBuffer, VkPipelineBindPoint pipelineBindPoint, VkPipelineLayout layout, uint32_t firstSet, uint32_t descriptorSetCount, VkDescriptorSet* pDescriptorSets, uint32_t dynamicOffsetCount, uint32_t* pDynamicOffsets) {
    fn(commandBuffer, pipelineBindPoint, layout, firstSet, descriptorSetCount, pDescriptorSets, dynamicOffsetCount, pDynamicOffsets);
}

void cgoCmdBindIndexBuffer(PFN_vkCmdBindIndexBuffer fn, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkIndexType indexType) {
    fn(commandBuffer, buffer, offset, indexType);
}

void cgoCmdBindVertexBuffers(PFN_vkCmdBindVertexBuffers fn, VkCommandBuffer commandBuffer, uint32_t firstBinding, uint32_t bindingCount, VkBuffer* pBuffers, VkDeviceSize* pOffsets) {
    fn(commandBuffer, firstBinding, bindingCount, pBuffers, pOffsets);
}

void cgoCmdDraw(PFN_vkCmdDraw fn, VkCommandBuffer commandBuffer, uint32_t vertexCount, uint32_t instanceCount, uint32_t firstVertex, uint32_t firstInstance) {
    fn(commandBuffer, vertexCount, instanceCount, firstVertex, firstInstance);
}

void cgoCmdDrawIndexed(PFN_vkCmdDrawIndexed fn, VkCommandBuffer commandBuffer, uint32_t indexCount, uint32_t instanceCount, uint32_t firstIndex, int32_t vertexOffset, uint32_t firstInstance) {
    fn(commandBuffer, indexCount, instanceCount, firstIndex, vertexOffset, firstInstance);
}

void cgoCmdDrawIndirect(PFN_vkCmdDrawIndirect fn, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, uint32_t drawCount, uint32_t stride) {
    fn(commandBuffer, buffer, offset, drawCount, stride);
}

void cgoCmdDrawIndexedIndirect(PFN_vkCmdDrawIndexedIndirect fn, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, uint32_t drawCount, uint32_t stride) {
    fn(commandBuffer, buffer, offset, drawCount, stride);
}

void cgoCmdDispatch(PFN_vkCmdDispatch fn, VkCommandBuffer commandBuffer, uint32_t groupCountX, uint32_t groupCountY, uint32_t groupCountZ) {
    fn(commandBuffer, groupCountX, groupCountY, groupCountZ);
}

void cgoCmdDispatchIndirect(PFN_vkCmdDispatchIndirect fn, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset) {
    fn(commandBuffer, buffer, offset);
}

void cgoCmdCopyBuffer(PFN_vkCmdCopyBuffer fn, VkCommandBuffer commandBuffer, VkBuffer srcBuffer, VkBuffer dstBuffer, uint32_t regionCount, VkBufferCopy* pRegions) {
    fn(commandBuffer, srcBuffer, dstBuffer, regionCount, pRegions);
}

void cgoCmdCopyImage(PFN_vkCmdCopyImage fn, VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, VkImageCopy* pRegions) {
    fn(commandBuffer, srcImage, srcImageLayout, dstImage, dstImageLayout, regionCount, pRegions);
}

void cgoCmdBlitImage(PFN_vkCmdBlitImage fn, VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, VkImageBlit* pRegions, VkFilter filter) {
    fn(commandBuffer, srcImage, srcImageLayout, dstImage, dstImageLayout, regionCount, pRegions, filter);
}

void cgoCmdCopyBufferToImage(PFN_vkCmdCopyBufferToImage fn, VkCommandBuffer commandBuffer, VkBuffer srcBuffer, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, VkBufferImageCopy* pRegions) {
    fn(commandBuffer, srcBuffer, dstImage, dstImageLayout, regionCount, pRegions);
}

void cgoCmdCopyImageToBuffer(PFN_vkCmdCopyImageToBuffer fn, VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkBuffer dstBuffer, uint32_t regionCount, VkBufferImageCopy* pRegions) {
    fn(commandBuffer, srcImage, srcImageLayout, dstBuffer, regionCount, pRegions);
}

void cgoCmdUpdateBuffer(PFN_vkCmdUpdateBuffer fn, VkCommandBuffer commandBuffer, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize dataSize, void* pData) {
    fn(commandBuffer, dstBuffer, dstOffset, dataSize, pData);
}

void cgoCmdFillBuffer(PFN_vkCmdFillBuffer fn, VkCommandBuffer commandBuffer, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize size, uint32_t data) {
    fn(commandBuffer, dstBuffer, dstOffset, size, data);
}

void cgoCmdClearColorImage(PFN_vkCmdClearColorImage fn, VkCommandBuffer commandBuffer, VkImage image, VkImageLayout imageLayout, VkClearColorValue* pColor, uint32_t rangeCount, VkImageSubresourceRange* pRanges) {
    fn(commandBuffer, image, imageLayout, pColor, rangeCount, pRanges);
}

void cgoCmdClearDepthStencilImage(PFN_vkCmdClearDepthStencilImage fn, VkCommandBuffer commandBuffer, VkImage image, VkImageLayout imageLayout, VkClearDepthStencilValue* pDepthStencil, uint32_t rangeCount, VkImageSubresourceRange* pRanges) {
    fn(commandBuffer, image, imageLayout, pDepthStencil, rangeCount, pRanges);
}

void cgoCmdClearAttachments(PFN_vkCmdClearAttachments fn, VkCommandBuffer commandBuffer, uint32_t attachmentCount, VkClearAttachment* pAttachments, uint32_t rectCount, VkClearRect* pRects) {
    fn(commandBuffer, attachmentCount, pAttachments, rectCount, pRects);
}

void cgoCmdResolveImage(PFN_vkCmdResolveImage fn, VkCommandBuffer commandBuffer, VkImage srcImage, VkImageLayout srcImageLayout, VkImage dstImage, VkImageLayout dstImageLayout, uint32_t regionCount, VkImageResolve* pRegions) {
    fn(commandBuffer, srcImage, srcImageLayout, dstImage, dstImageLayout, regionCount, pRegions);
}

void cgoCmdSetEvent(PFN_vkCmdSetEvent fn, VkCommandBuffer commandBuffer, VkEvent event, VkPipelineStageFlags stageMask) {
    fn(commandBuffer, event, stageMask);
}

void cgoCmdResetEvent(PFN_vkCmdResetEvent fn, VkCommandBuffer commandBuffer, VkEvent event, VkPipelineStageFlags stageMask) {
    fn(commandBuffer, event, stageMask);
}

void cgoCmdWaitEvents(PFN_vkCmdWaitEvents fn, VkCommandBuffer commandBuffer, uint32_t eventCount, VkEvent* pEvents, VkPipelineStageFlags srcStageMask, VkPipelineStageFlags dstStageMask, uint32_t memoryBarrierCount, VkMemoryBarrier* pMemoryBarriers, uint32_t bufferMemoryBarrierCount, VkBufferMemoryBarrier* pBufferMemoryBarriers, uint32_t imageMemoryBarrierCount, VkImageMemoryBarrier* pImageMemoryBarriers) {
    fn(commandBuffer, eventCount, pEvents, srcStageMask, dstStageMask, memoryBarrierCount, pMemoryBarriers, bufferMemoryBarrierCount, pBufferMemoryBarriers, imageMemoryBarrierCount, pImageMemoryBarriers);
}

void cgoCmdPipelineBarrier(PFN_vkCmdPipelineBarrier fn, VkCommandBuffer commandBuffer, VkPipelineStageFlags srcStageMask, VkPipelineStageFlags dstStageMask, VkDependencyFlags dependencyFlags, uint32_t memoryBarrierCount, VkMemoryBarrier* pMemoryBarriers, uint32_t bufferMemoryBarrierCount, VkBufferMemoryBarrier* pBufferMemoryBarriers, uint32_t imageMemoryBarrierCount, VkImageMemoryBarrier* pImageMemoryBarriers) {
    fn(commandBuffer, srcStageMask, dstStageMask, dependencyFlags, memoryBarrierCount, pMemoryBarriers, bufferMemoryBarrierCount, pBufferMemoryBarriers, imageMemoryBarrierCount, pImageMemoryBarriers);
}

void cgoCmdBeginQuery(PFN_vkCmdBeginQuery fn, VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t query, VkQueryControlFlags flags) {
    fn(commandBuffer, queryPool, query, flags);
}

void cgoCmdEndQuery(PFN_vkCmdEndQuery fn, VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t query) {
    fn(commandBuffer, queryPool, query);
}

void cgoCmdResetQueryPool(PFN_vkCmdResetQueryPool fn, VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount) {
    fn(commandBuffer, queryPool, firstQuery, queryCount);
}

void cgoCmdWriteTimestamp(PFN_vkCmdWriteTimestamp fn, VkCommandBuffer commandBuffer, VkPipelineStageFlagBits pipelineStage, VkQueryPool queryPool, uint32_t query) {
    fn(commandBuffer, pipelineStage, queryPool, query);
}

void cgoCmdCopyQueryPoolResults(PFN_vkCmdCopyQueryPoolResults fn, VkCommandBuffer commandBuffer, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount, VkBuffer dstBuffer, VkDeviceSize dstOffset, VkDeviceSize stride, VkQueryResultFlags flags) {
    fn(commandBuffer, queryPool, firstQuery, queryCount, dstBuffer, dstOffset, stride, flags);
}

void cgoCmdPushConstants(PFN_vkCmdPushConstants fn, VkCommandBuffer commandBuffer, VkPipelineLayout layout, VkShaderStageFlags stageFlags, uint32_t offset, uint32_t size, void* pValues) {
    fn(commandBuffer, layout, stageFlags, offset, size, pValues);
}

void cgoCmdBeginRenderPass(PFN_vkCmdBeginRenderPass fn, VkCommandBuffer commandBuffer, VkRenderPassBeginInfo* pRenderPassBegin, VkSubpassContents contents) {
    fn(commandBuffer, pRenderPassBegin, contents);
}

void cgoCmdNextSubpass(PFN_vkCmdNextSubpass fn, VkCommandBuffer commandBuffer, VkSubpassContents contents) {
    fn(commandBuffer, contents);
}

void cgoCmdEndRenderPass(PFN_vkCmdEndRenderPass fn, VkCommandBuffer commandBuffer) {
    fn(commandBuffer);
}

void cgoCmdExecuteCommands(PFN_vkCmdExecuteCommands fn, VkCommandBuffer commandBuffer, uint32_t commandBufferCount, VkCommandBuffer* pCommandBuffers) {
    fn(commandBuffer, commandBufferCount, pCommandBuffers);
}

VkResult cgoEnumerateInstanceVersion(PFN_vkEnumerateInstanceVersion fn, uint32_t* pApiVersion) {
    return fn(pApiVersion);
}

VkResult cgoBindBufferMemory2(PFN_vkBindBufferMemory2 fn, VkDevice device, uint32_t bindInfoCount, VkBindBufferMemoryInfo* pBindInfos) {
    return fn(device, bindInfoCount, pBindInfos);
}

VkResult cgoBindImageMemory2(PFN_vkBindImageMemory2 fn, VkDevice device, uint32_t bindInfoCount, VkBindImageMemoryInfo* pBindInfos) {
    return fn(device, bindInfoCount, pBindInfos);
}

void cgoGetDeviceGroupPeerMemoryFeatures(PFN_vkGetDeviceGroupPeerMemoryFeatures fn, VkDevice device, uint32_t heapIndex, uint32_t localDeviceIndex, uint32_t remoteDeviceIndex, VkPeerMemoryFeatureFlags* pPeerMemoryFeatures) {
    fn(device, heapIndex, localDeviceIndex, remoteDeviceIndex, pPeerMemoryFeatures);
}

void cgoCmdSetDeviceMask(PFN_vkCmdSetDeviceMask fn, VkCommandBuffer commandBuffer, uint32_t deviceMask) {
    fn(commandBuffer, deviceMask);
}

void cgoCmdDispatchBase(PFN_vkCmdDispatchBase fn, VkCommandBuffer commandBuffer, uint32_t baseGroupX, uint32_t baseGroupY, uint32_t baseGroupZ, uint32_t groupCountX, uint32_t groupCountY, uint32_t groupCountZ) {
    fn(commandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ);
}

VkResult cgoEnumeratePhysicalDeviceGroups(PFN_vkEnumeratePhysicalDeviceGroups fn, VkInstance instance, uint32_t* pPhysicalDeviceGroupCount, VkPhysicalDeviceGroupProperties* pPhysicalDeviceGroupProperties) {
    return fn(instance, pPhysicalDeviceGroupCount, pPhysicalDeviceGroupProperties);
}

void cgoGetImageMemoryRequirements2(PFN_vkGetImageMemoryRequirements2 fn, VkDevice device, VkImageMemoryRequirementsInfo2* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
    fn(device, pInfo, pMemoryRequirements);
}

void cgoGetBufferMemoryRequirements2(PFN_vkGetBufferMemoryRequirements2 fn, VkDevice device, VkBufferMemoryRequirementsInfo2* pInfo, VkMemoryRequirements2* pMemoryRequirements) {
    fn(device, pInfo, pMemoryRequirements);
}

void cgoGetImageSparseMemoryRequirements2(PFN_vkGetImageSparseMemoryRequirements2 fn, VkDevice device, VkImageSparseMemoryRequirementsInfo2* pInfo, uint32_t* pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements2* pSparseMemoryRequirements) {
    fn(device, pInfo, pSparseMemoryRequirementCount, pSparseMemoryRequirements);
}

void cgoGetPhysicalDeviceFeatures2(PFN_vkGetPhysicalDeviceFeatures2 fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceFeatures2* pFeatures) {
    fn(physicalDevice, pFeatures);
}

void cgoGetPhysicalDeviceProperties2(PFN_vkGetPhysicalDeviceProperties2 fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties2* pProperties) {
    fn(physicalDevice, pProperties);
}

void cgoGetPhysicalDeviceFormatProperties2(PFN_vkGetPhysicalDeviceFormatProperties2 fn, VkPhysicalDevice physicalDevice, VkFormat format, VkFormatProperties2* pFormatProperties) {
    fn(physicalDevice, format, pFormatProperties);
}

VkResult cgoGetPhysicalDeviceImageFormatProperties2(PFN_vkGetPhysicalDeviceImageFormatProperties2 fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceImageFormatInfo2* pImageFormatInfo, VkImageFormatProperties2* pImageFormatProperties) {
    return fn(physicalDevice, pImageFormatInfo, pImageFormatProperties);
}

void cgoGetPhysicalDeviceQueueFamilyProperties2(PFN_vkGetPhysicalDeviceQueueFamilyProperties2 fn, VkPhysicalDevice physicalDevice, uint32_t* pQueueFamilyPropertyCount, VkQueueFamilyProperties2* pQueueFamilyProperties) {
    fn(physicalDevice, pQueueFamilyPropertyCount, pQueueFamilyProperties);
}

void cgoGetPhysicalDeviceMemoryProperties2(PFN_vkGetPhysicalDeviceMemoryProperties2 fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceMemoryProperties2* pMemoryProperties) {
    fn(physicalDevice, pMemoryProperties);
}

void cgoGetPhysicalDeviceSparseImageFormatProperties2(PFN_vkGetPhysicalDeviceSparseImageFormatProperties2 fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceSparseImageFormatInfo2* pFormatInfo, uint32_t* pPropertyCount, VkSparseImageFormatProperties2* pProperties) {
    fn(physicalDevice, pFormatInfo, pPropertyCount, pProperties);
}

void cgoTrimCommandPool(PFN_vkTrimCommandPool fn, VkDevice device, VkCommandPool commandPool, VkCommandPoolTrimFlags flags) {
    fn(device, commandPool, flags);
}

void cgoGetDeviceQueue2(PFN_vkGetDeviceQueue2 fn, VkDevice device, VkDeviceQueueInfo2* pQueueInfo, VkQueue* pQueue) {
    fn(device, pQueueInfo, pQueue);
}

VkResult cgoCreateSamplerYcbcrConversion(PFN_vkCreateSamplerYcbcrConversion fn, VkDevice device, VkSamplerYcbcrConversionCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkSamplerYcbcrConversion* pYcbcrConversion) {
    return fn(device, pCreateInfo, pAllocator, pYcbcrConversion);
}

void cgoDestroySamplerYcbcrConversion(PFN_vkDestroySamplerYcbcrConversion fn, VkDevice device, VkSamplerYcbcrConversion ycbcrConversion, VkAllocationCallbacks* pAllocator) {
    fn(device, ycbcrConversion, pAllocator);
}

VkResult cgoCreateDescriptorUpdateTemplate(PFN_vkCreateDescriptorUpdateTemplate fn, VkDevice device, VkDescriptorUpdateTemplateCreateInfo* pCreateInfo, VkAllocationCallbacks* pAllocator, VkDescriptorUpdateTemplate* pDescriptorUpdateTemplate) {
    return fn(device, pCreateInfo, pAllocator, pDescriptorUpdateTemplate);
}

void cgoDestroyDescriptorUpdateTemplate(PFN_vkDestroyDescriptorUpdateTemplate fn, VkDevice device, VkDescriptorUpdateTemplate descriptorUpdateTemplate, VkAllocationCallbacks* pAllocator) {
    fn(device, descriptorUpdateTemplate, pAllocator);
}

void cgoUpdateDescriptorSetWithTemplate(PFN_vkUpdateDescriptorSetWithTemplate fn, VkDevice device, VkDescriptorSet descriptorSet, VkDescriptorUpdateTemplate descriptorUpdateTemplate, void* pData) {
    fn(device, descriptorSet, descriptorUpdateTemplate, pData);
}

void cgoGetPhysicalDeviceExternalBufferProperties(PFN_vkGetPhysicalDeviceExternalBufferProperties fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceExternalBufferInfo* pExternalBufferInfo, VkExternalBufferProperties* pExternalBufferProperties) {
    fn(physicalDevice, pExternalBufferInfo, pExternalBufferProperties);
}

void cgoGetPhysicalDeviceExternalFenceProperties(PFN_vkGetPhysicalDeviceExternalFenceProperties fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceExternalFenceInfo* pExternalFenceInfo, VkExternalFenceProperties* pExternalFenceProperties) {
    fn(physicalDevice, pExternalFenceInfo, pExternalFenceProperties);
}

void cgoGetPhysicalDeviceExternalSemaphoreProperties(PFN_vkGetPhysicalDeviceExternalSemaphoreProperties fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceExternalSemaphoreInfo* pExternalSemaphoreInfo, VkExternalSemaphoreProperties* pExternalSemaphoreProperties) {
    fn(physicalDevice, pExternalSemaphoreInfo, pExternalSemaphoreProperties);
}

void cgoGetDescriptorSetLayoutSupport(PFN_vkGetDescriptorSetLayoutSupport fn, VkDevice device, VkDescriptorSetLayoutCreateInfo* pCreateInfo, VkDescriptorSetLayoutSupport* pSupport) {
    fn(device, pCreateInfo, pSupport);
}

void cgoCmdDrawIndirectCount(PFN_vkCmdDrawIndirectCount fn, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
    fn(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}

void cgoCmdDrawIndexedIndirectCount(PFN_vkCmdDrawIndexedIndirectCount fn, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
    fn(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}

VkResult cgoCreateRenderPass2(PFN_vkCreateRenderPass2 fn, VkDevice device, VkRenderPassCreateInfo2* pCreateInfo, VkAllocationCallbacks* pAllocator, VkRenderPass* pRenderPass) {
    return fn(device, pCreateInfo, pAllocator, pRenderPass);
}

void cgoCmdBeginRenderPass2(PFN_vkCmdBeginRenderPass2 fn, VkCommandBuffer commandBuffer, VkRenderPassBeginInfo*      pRenderPassBegin, VkSubpassBeginInfo*      pSubpassBeginInfo) {
    fn(commandBuffer, pRenderPassBegin, pSubpassBeginInfo);
}

void cgoCmdNextSubpass2(PFN_vkCmdNextSubpass2 fn, VkCommandBuffer commandBuffer, VkSubpassBeginInfo*      pSubpassBeginInfo, VkSubpassEndInfo*        pSubpassEndInfo) {
    fn(commandBuffer, pSubpassBeginInfo, pSubpassEndInfo);
}

void cgoCmdEndRenderPass2(PFN_vkCmdEndRenderPass2 fn, VkCommandBuffer commandBuffer, VkSubpassEndInfo*        pSubpassEndInfo) {
    fn(commandBuffer, pSubpassEndInfo);
}

void cgoResetQueryPool(PFN_vkResetQueryPool fn, VkDevice device, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount) {
    fn(device, queryPool, firstQuery, queryCount);
}

VkResult cgoGetSemaphoreCounterValue(PFN_vkGetSemaphoreCounterValue fn, VkDevice device, VkSemaphore semaphore, uint64_t* pValue) {
    return fn(device, semaphore, pValue);
}

VkResult cgoWaitSemaphores(PFN_vkWaitSemaphores fn, VkDevice device, VkSemaphoreWaitInfo* pWaitInfo, uint64_t timeout) {
    return fn(device, pWaitInfo, timeout);
}

VkResult cgoSignalSemaphore(PFN_vkSignalSemaphore fn, VkDevice device, VkSemaphoreSignalInfo* pSignalInfo) {
    return fn(device, pSignalInfo);
}

VkDeviceAddress cgoGetBufferDeviceAddress(PFN_vkGetBufferDeviceAddress fn, VkDevice device, VkBufferDeviceAddressInfo* pInfo) {
    return fn(device, pInfo);
}

uint64_t cgoGetBufferOpaqueCaptureAddress(PFN_vkGetBufferOpaqueCaptureAddress fn, VkDevice device, VkBufferDeviceAddressInfo* pInfo) {
    return fn(device, pInfo);
}

uint64_t cgoGetDeviceMemoryOpaqueCaptureAddress(PFN_vkGetDeviceMemoryOpaqueCaptureAddress fn, VkDevice device, VkDeviceMemoryOpaqueCaptureAddressInfo* pInfo) {
    return fn(device, pInfo);
}


