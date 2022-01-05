#include <stdlib.h>
#include "vulkan/vulkan.h"

typedef struct DriverFuncPtrs {
    //VK 1.0
    
    //Platform
    PFN_vkGetInstanceProcAddr vkGetInstanceProcAddr;

    //Pre-instance
    PFN_vkEnumerateInstanceExtensionProperties vkEnumerateInstanceExtensionProperties;
    PFN_vkEnumerateInstanceLayerProperties vkEnumerateInstanceLayerProperties;
    PFN_vkCreateInstance vkCreateInstance;

    //Instance
    PFN_vkCreateDevice vkCreateDevice;
    PFN_vkDestroyInstance vkDestroyInstance;
    PFN_vkEnumerateDeviceExtensionProperties vkEnumerateDeviceExtensionProperties;
    PFN_vkEnumerateDeviceLayerProperties vkEnumerateDeviceLayerProperties; //Todo*
    PFN_vkEnumeratePhysicalDevices vkEnumeratePhysicalDevices;
    PFN_vkGetPhysicalDeviceFeatures vkGetPhysicalDeviceFeatures;
    PFN_vkGetPhysicalDeviceFormatProperties vkGetPhysicalDeviceFormatProperties;
    PFN_vkGetPhysicalDeviceImageFormatProperties vkGetPhysicalDeviceImageFormatProperties; //Todo*
    PFN_vkGetPhysicalDeviceMemoryProperties vkGetPhysicalDeviceMemoryProperties;
    PFN_vkGetPhysicalDeviceProperties vkGetPhysicalDeviceProperties;
    PFN_vkGetPhysicalDeviceQueueFamilyProperties vkGetPhysicalDeviceQueueFamilyProperties;
    PFN_vkGetPhysicalDeviceSparseImageFormatProperties vkGetPhysicalDeviceSparseImageFormatProperties; //Todo*

    //Device-Platform
    PFN_vkGetDeviceProcAddr vkGetDeviceProcAddr;

    //Device
    PFN_vkAllocateCommandBuffers vkAllocateCommandBuffers;
    PFN_vkAllocateDescriptorSets vkAllocateDescriptorSets;
    PFN_vkAllocateMemory vkAllocateMemory;
    PFN_vkBeginCommandBuffer vkBeginCommandBuffer;
    PFN_vkBindBufferMemory vkBindBufferMemory;
    PFN_vkBindImageMemory vkBindImageMemory;
    PFN_vkCmdBeginQuery vkCmdBeginQuery;
    PFN_vkCmdBeginRenderPass vkCmdBeginRenderPass;
    PFN_vkCmdBindDescriptorSets vkCmdBindDescriptorSets;
    PFN_vkCmdBindIndexBuffer vkCmdBindIndexBuffer;
    PFN_vkCmdBindPipeline vkCmdBindPipeline;
    PFN_vkCmdBindVertexBuffers vkCmdBindVertexBuffers;
    PFN_vkCmdBlitImage vkCmdBlitImage;
    PFN_vkCmdClearAttachments vkCmdClearAttachments; //Todo*
    PFN_vkCmdClearColorImage vkCmdClearColorImage;
    PFN_vkCmdClearDepthStencilImage vkCmdClearDepthStencilImage; //Todo*
    PFN_vkCmdCopyBuffer vkCmdCopyBuffer;
    PFN_vkCmdCopyBufferToImage vkCmdCopyBufferToImage;
    PFN_vkCmdCopyImage vkCmdCopyImage;
    PFN_vkCmdCopyImageToBuffer vkCmdCopyImageToBuffer; //Todo*
    PFN_vkCmdCopyQueryPoolResults vkCmdCopyQueryPoolResults;
    PFN_vkCmdDispatch vkCmdDispatch; //Todo*
    PFN_vkCmdDispatchIndirect vkCmdDispatchIndirect; //Todo*
    PFN_vkCmdDraw vkCmdDraw;
    PFN_vkCmdDrawIndexed vkCmdDrawIndexed;
    PFN_vkCmdDrawIndexedIndirect vkCmdDrawIndexedIndirect; //Todo*
    PFN_vkCmdDrawIndirect vkCmdDrawIndirect; //Todo*
    PFN_vkCmdEndQuery vkCmdEndQuery;
    PFN_vkCmdEndRenderPass vkCmdEndRenderPass;
    PFN_vkCmdExecuteCommands vkCmdExecuteCommands;
    PFN_vkCmdFillBuffer vkCmdFillBuffer; //Todo*
    PFN_vkCmdNextSubpass vkCmdNextSubpass;
    PFN_vkCmdPipelineBarrier vkCmdPipelineBarrier;
    PFN_vkCmdPushConstants vkCmdPushConstants;
    PFN_vkCmdResetEvent vkCmdResetEvent; //Todo*
    PFN_vkCmdResetQueryPool vkCmdResetQueryPool;
    PFN_vkCmdResolveImage vkCmdResolveImage; //Todo*
    PFN_vkCmdSetBlendConstants vkCmdSetBlendConstants; //Todo*
    PFN_vkCmdSetDepthBias vkCmdSetDepthBias; //Todo*
    PFN_vkCmdSetDepthBounds vkCmdSetDepthBounds; //Todo*
    PFN_vkCmdSetEvent vkCmdSetEvent;
    PFN_vkCmdSetLineWidth vkCmdSetLineWidth; //Todo*
    PFN_vkCmdSetScissor vkCmdSetScissor;
    PFN_vkCmdSetStencilCompareMask vkCmdSetStencilCompareMask; //Todo*
    PFN_vkCmdSetStencilReference vkCmdSetStencilReference; //Todo*
    PFN_vkCmdSetStencilWriteMask vkCmdSetStencilWriteMask; //Todo*
    PFN_vkCmdSetViewport vkCmdSetViewport;
    PFN_vkCmdUpdateBuffer vkCmdUpdateBuffer; //Todo*
    PFN_vkCmdWaitEvents vkCmdWaitEvents;
    PFN_vkCmdWriteTimestamp vkCmdWriteTimestamp; //Todo*
    PFN_vkCreateBuffer vkCreateBuffer;
    PFN_vkCreateBufferView vkCreateBufferView;
    PFN_vkCreateCommandPool vkCreateCommandPool;
    PFN_vkCreateComputePipelines vkCreateComputePipelines; //Todo*
    PFN_vkCreateDescriptorPool vkCreateDescriptorPool;
    PFN_vkCreateDescriptorSetLayout vkCreateDescriptorSetLayout;
    PFN_vkCreateEvent vkCreateEvent;
    PFN_vkCreateFence vkCreateFence;
    PFN_vkCreateFramebuffer vkCreateFramebuffer;
    PFN_vkCreateGraphicsPipelines vkCreateGraphicsPipelines;
    PFN_vkCreateImage vkCreateImage;
    PFN_vkCreateImageView vkCreateImageView;
    PFN_vkCreatePipelineCache vkCreatePipelineCache;
    PFN_vkCreatePipelineLayout vkCreatePipelineLayout;
    PFN_vkCreateQueryPool vkCreateQueryPool;
    PFN_vkCreateRenderPass vkCreateRenderPass;
    PFN_vkCreateSampler vkCreateSampler;
    PFN_vkCreateSemaphore vkCreateSemaphore;
    PFN_vkCreateShaderModule vkCreateShaderModule;
    PFN_vkDestroyBuffer vkDestroyBuffer;
    PFN_vkDestroyBufferView vkDestroyBufferView;
    PFN_vkDestroyCommandPool vkDestroyCommandPool;
    PFN_vkDestroyDescriptorPool vkDestroyDescriptorPool;
    PFN_vkDestroyDescriptorSetLayout vkDestroyDescriptorSetLayout;
    PFN_vkDestroyDevice vkDestroyDevice;
    PFN_vkDestroyEvent vkDestroyEvent;
    PFN_vkDestroyFence vkDestroyFence;
    PFN_vkDestroyFramebuffer vkDestroyFramebuffer;
    PFN_vkDestroyImage vkDestroyImage;
    PFN_vkDestroyImageView vkDestroyImageView;
    PFN_vkDestroyPipeline vkDestroyPipeline;
    PFN_vkDestroyPipelineCache vkDestroyPipelineCache;
    PFN_vkDestroyPipelineLayout vkDestroyPipelineLayout;
    PFN_vkDestroyQueryPool vkDestroyQueryPool;
    PFN_vkDestroyRenderPass vkDestroyRenderPass;
    PFN_vkDestroySampler vkDestroySampler;
    PFN_vkDestroySemaphore vkDestroySemaphore;
    PFN_vkDestroyShaderModule vkDestroyShaderModule;
    PFN_vkDeviceWaitIdle vkDeviceWaitIdle;
    PFN_vkEndCommandBuffer vkEndCommandBuffer;
    PFN_vkFlushMappedMemoryRanges vkFlushMappedMemoryRanges;
    PFN_vkFreeCommandBuffers vkFreeCommandBuffers;
    PFN_vkFreeDescriptorSets vkFreeDescriptorSets;
    PFN_vkFreeMemory vkFreeMemory;
    PFN_vkGetBufferMemoryRequirements vkGetBufferMemoryRequirements;
    PFN_vkGetDeviceMemoryCommitment vkGetDeviceMemoryCommitment; //Todo*
    PFN_vkGetDeviceQueue vkGetDeviceQueue;
    PFN_vkGetEventStatus vkGetEventStatus;
    PFN_vkGetFenceStatus vkGetFenceStatus; //Todo*
    PFN_vkGetImageMemoryRequirements vkGetImageMemoryRequirements;
    PFN_vkGetImageSparseMemoryRequirements vkGetImageSparseMemoryRequirements; //Todo*
    PFN_vkGetImageSubresourceLayout vkGetImageSubresourceLayout;
    PFN_vkGetPipelineCacheData vkGetPipelineCacheData;
    PFN_vkGetQueryPoolResults vkGetQueryPoolResults;
    PFN_vkGetRenderAreaGranularity vkGetRenderAreaGranularity; //Todo*
    PFN_vkInvalidateMappedMemoryRanges vkInvalidateMappedMemoryRanges; //Todo*
    PFN_vkMapMemory vkMapMemory;
    PFN_vkMergePipelineCaches vkMergePipelineCaches; //Todo*
    PFN_vkQueueBindSparse vkQueueBindSparse; //Todo*
    PFN_vkQueueSubmit vkQueueSubmit;
    PFN_vkQueueWaitIdle vkQueueWaitIdle;
    PFN_vkResetCommandBuffer vkResetCommandBuffer;
    PFN_vkResetCommandPool vkResetCommandPool; //Todo*
    PFN_vkResetDescriptorPool vkResetDescriptorPool; //Todo*
    PFN_vkResetEvent vkResetEvent;
    PFN_vkResetFences vkResetFences;
    PFN_vkSetEvent vkSetEvent;
    PFN_vkUnmapMemory vkUnmapMemory;
    PFN_vkUpdateDescriptorSets vkUpdateDescriptorSets;
    PFN_vkWaitForFences vkWaitForFences;
    
    // VK 1.1

    //Platform
    PFN_vkEnumerateInstanceVersion vkEnumerateInstanceVersion; //Todo

    //Instance
    PFN_vkEnumeratePhysicalDeviceGroups vkEnumeratePhysicalDeviceGroups; //Todo
    PFN_vkGetPhysicalDeviceFeatures2 vkGetPhysicalDeviceFeatures2; //Todo
    PFN_vkGetPhysicalDeviceProperties2 vkGetPhysicalDeviceProperties2; //Todo
    PFN_vkGetPhysicalDeviceFormatProperties2 vkGetPhysicalDeviceFormatProperties2; //Todo
    PFN_vkGetPhysicalDeviceImageFormatProperties2 vkGetPhysicalDeviceImageFormatProperties2; //Todo
    PFN_vkGetPhysicalDeviceQueueFamilyProperties2 vkGetPhysicalDeviceQueueFamilyProperties2; //Todo
    PFN_vkGetPhysicalDeviceMemoryProperties2 vkGetPhysicalDeviceMemoryProperties2; //Todo
    PFN_vkGetPhysicalDeviceSparseImageFormatProperties2 vkGetPhysicalDeviceSparseImageFormatProperties2; //Todo
    PFN_vkGetPhysicalDeviceExternalBufferProperties vkGetPhysicalDeviceExternalBufferProperties; //Todo
    PFN_vkGetPhysicalDeviceExternalFenceProperties vkGetPhysicalDeviceExternalFenceProperties; //Todo
    PFN_vkGetPhysicalDeviceExternalSemaphoreProperties vkGetPhysicalDeviceExternalSemaphoreProperties; //Todo

    //Device
    PFN_vkBindBufferMemory2 vkBindBufferMemory2; //Todo
    PFN_vkBindImageMemory2 vkBindImageMemory2; //Todo
    PFN_vkGetDeviceGroupPeerMemoryFeatures vkGetDeviceGroupPeerMemoryFeatures; //Todo
    PFN_vkCmdSetDeviceMask vkCmdSetDeviceMask; //Todo
    PFN_vkCmdDispatchBase vkCmdDispatchBase; //Todo
    PFN_vkGetImageMemoryRequirements2 vkGetImageMemoryRequirements2; //Todo
    PFN_vkGetBufferMemoryRequirements2 vkGetBufferMemoryRequirements2; //Todo
    PFN_vkGetImageSparseMemoryRequirements2 vkGetImageSparseMemoryRequirements2; //Todo
    PFN_vkTrimCommandPool vkTrimCommandPool; //Todo
    PFN_vkGetDeviceQueue2 vkGetDeviceQueue2; //Todo
    PFN_vkCreateSamplerYcbcrConversion vkCreateSamplerYcbcrConversion; //Todo
    PFN_vkDestroySamplerYcbcrConversion vkDestroySamplerYcbcrConversion; //Todo
    PFN_vkCreateDescriptorUpdateTemplate vkCreateDescriptorUpdateTemplate; //Todo
    PFN_vkDestroyDescriptorUpdateTemplate vkDestroyDescriptorUpdateTemplate; //Todo
    PFN_vkUpdateDescriptorSetWithTemplate vkUpdateDescriptorSetWithTemplate; //Todo
    PFN_vkGetDescriptorSetLayoutSupport vkGetDescriptorSetLayoutSupport; //Todo
    
    // VK 1.2
    
    //Device
    PFN_vkCmdDrawIndirectCount vkCmdDrawIndirectCount; //Todo
    PFN_vkCmdDrawIndexedIndirectCount vkCmdDrawIndexedIndirectCount; //Todo
    PFN_vkCreateRenderPass2 vkCreateRenderPass2; //Todo
    PFN_vkCmdBeginRenderPass2 vkCmdBeginRenderPass2; //Todo
    PFN_vkCmdNextSubpass2 vkCmdNextSubpass2; //Todo
    PFN_vkCmdEndRenderPass2 vkCmdEndRenderPass2; //Todo
    PFN_vkResetQueryPool vkResetQueryPool; //Todo
    PFN_vkGetSemaphoreCounterValue vkGetSemaphoreCounterValue; //Todo
    PFN_vkWaitSemaphores vkWaitSemaphores; //Todo
    PFN_vkSignalSemaphore vkSignalSemaphore; //Todo
    PFN_vkGetBufferDeviceAddress vkGetBufferDeviceAddress; //Todo
    PFN_vkGetBufferOpaqueCaptureAddress vkGetBufferOpaqueCaptureAddress; //Todo
    PFN_vkGetDeviceMemoryOpaqueCaptureAddress vkGetDeviceMemoryOpaqueCaptureAddress; //Todo
} DriverFuncPtrs;
