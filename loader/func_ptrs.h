#include <stdlib.h>
#include "../vulkan/vulkan.h"

typedef struct LoaderFuncPtrs {
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
    PFN_vkEnumerateDeviceLayerProperties vkEnumerateDeviceLayerProperties; //Todo
    PFN_vkEnumeratePhysicalDevices vkEnumeratePhysicalDevices;
    PFN_vkGetPhysicalDeviceFeatures vkGetPhysicalDeviceFeatures;
    PFN_vkGetPhysicalDeviceFormatProperties vkGetPhysicalDeviceFormatProperties; //Todo
    PFN_vkGetPhysicalDeviceImageFormatProperties vkGetPhysicalDeviceImageFormatProperties; //Todo
    PFN_vkGetPhysicalDeviceMemoryProperties vkGetPhysicalDeviceMemoryProperties;
    PFN_vkGetPhysicalDeviceProperties vkGetPhysicalDeviceProperties;
    PFN_vkGetPhysicalDeviceQueueFamilyProperties vkGetPhysicalDeviceQueueFamilyProperties;
    PFN_vkGetPhysicalDeviceSparseImageFormatProperties vkGetPhysicalDeviceSparseImageFormatProperties; //Todo

    //Device-Platform
    PFN_vkGetDeviceProcAddr vkGetDeviceProcAddr;

    //Device
    PFN_vkAllocateCommandBuffers vkAllocateCommandBuffers;
    PFN_vkAllocateDescriptorSets vkAllocateDescriptorSets; //Todo
    PFN_vkAllocateMemory vkAllocateMemory;
    PFN_vkBeginCommandBuffer vkBeginCommandBuffer;
    PFN_vkBindBufferMemory vkBindBufferMemory;
    PFN_vkBindImageMemory vkBindImageMemory;   //Todo
    PFN_vkCmdBeginQuery vkCmdBeginQuery; //Todo
    PFN_vkCmdBeginRenderPass vkCmdBeginRenderPass;
    PFN_vkCmdBindDescriptorSets vkCmdBindDescriptorSets; //Todo
    PFN_vkCmdBindIndexBuffer vkCmdBindIndexBuffer;
    PFN_vkCmdBindPipeline vkCmdBindPipeline;
    PFN_vkCmdBindVertexBuffers vkCmdBindVertexBuffers;
    PFN_vkCmdBlitImage vkCmdBlitImage; //Todo
    PFN_vkCmdClearAttachments vkCmdClearAttachments; //Todo
    PFN_vkCmdClearColorImage vkCmdClearColorImage; //Todo
    PFN_vkCmdClearDepthStencilImage vkCmdClearDepthStencilImage; //Todo
    PFN_vkCmdCopyBuffer vkCmdCopyBuffer;
    PFN_vkCmdCopyBufferToImage vkCmdCopyBufferToImage; //Todo
    PFN_vkCmdCopyImage vkCmdCopyImage; //Todo
    PFN_vkCmdCopyImageToBuffer vkCmdCopyImageToBuffer; //Todo
    PFN_vkCmdCopyQueryPoolResults vkCmdCopyQueryPoolResults; //Todo
    PFN_vkCmdDispatch vkCmdDispatch; //Todo
    PFN_vkCmdDispatchIndirect vkCmdDispatchIndirect; //Todo
    PFN_vkCmdDraw vkCmdDraw;
    PFN_vkCmdDrawIndexed vkCmdDrawIndexed;
    PFN_vkCmdDrawIndexedIndirect vkCmdDrawIndexedIndirect; //Todo
    PFN_vkCmdDrawIndirect vkCmdDrawIndirect; //Todo
    PFN_vkCmdEndQuery vkCmdEndQuery; //Todo
    PFN_vkCmdEndRenderPass vkCmdEndRenderPass;
    PFN_vkCmdExecuteCommands vkCmdExecuteCommands; //Todo
    PFN_vkCmdFillBuffer vkCmdFillBuffer; //Todo
    PFN_vkCmdNextSubpass vkCmdNextSubpass; //Todo
    PFN_vkCmdPipelineBarrier vkCmdPipelineBarrier; //Todo
    PFN_vkCmdPushConstants vkCmdPushConstants; //Todo
    PFN_vkCmdResetEvent vkCmdResetEvent; //Todo
    PFN_vkCmdResetQueryPool vkCmdResetQueryPool; //Todo
    PFN_vkCmdResolveImage vkCmdResolveImage; //Todo
    PFN_vkCmdSetBlendConstants vkCmdSetBlendConstants; //Todo
    PFN_vkCmdSetDepthBias vkCmdSetDepthBias; //Todo
    PFN_vkCmdSetDepthBounds vkCmdSetDepthBounds; //Todo
    PFN_vkCmdSetEvent vkCmdSetEvent; //Todo
    PFN_vkCmdSetLineWidth vkCmdSetLineWidth; //Todo
    PFN_vkCmdSetScissor vkCmdSetScissor;   //Todo
    PFN_vkCmdSetStencilCompareMask vkCmdSetStencilCompareMask; //Todo
    PFN_vkCmdSetStencilReference vkCmdSetStencilReference; //Todo
    PFN_vkCmdSetStencilWriteMask vkCmdSetStencilWriteMask; //Todo
    PFN_vkCmdSetViewport vkCmdSetViewport; //Todo
    PFN_vkCmdUpdateBuffer vkCmdUpdateBuffer; //Todo
    PFN_vkCmdWaitEvents vkCmdWaitEvents; //Todo
    PFN_vkCmdWriteTimestamp vkCmdWriteTimestamp; //Todo
    PFN_vkCreateBuffer vkCreateBuffer;
    PFN_vkCreateBufferView vkCreateBufferView; //Todo
    PFN_vkCreateCommandPool vkCreateCommandPool;
    PFN_vkCreateComputePipelines vkCreateComputePipelines; //Todo
    PFN_vkCreateDescriptorPool vkCreateDescriptorPool; //Todo
    PFN_vkCreateDescriptorSetLayout vkCreateDescriptorSetLayout;
    PFN_vkCreateEvent vkCreateEvent; //Todo
    PFN_vkCreateFence vkCreateFence;
    PFN_vkCreateFramebuffer vkCreateFramebuffer;
    PFN_vkCreateGraphicsPipelines vkCreateGraphicsPipelines;
    PFN_vkCreateImage vkCreateImage;
    PFN_vkCreateImageView vkCreateImageView;
    PFN_vkCreatePipelineCache vkCreatePipelineCache; //Todo
    PFN_vkCreatePipelineLayout vkCreatePipelineLayout;
    PFN_vkCreateQueryPool vkCreateQueryPool; //Todo
    PFN_vkCreateRenderPass vkCreateRenderPass;
    PFN_vkCreateSampler vkCreateSampler; //Todo
    PFN_vkCreateSemaphore vkCreateSemaphore;
    PFN_vkCreateShaderModule vkCreateShaderModule;
    PFN_vkDestroyBuffer vkDestroyBuffer;
    PFN_vkDestroyBufferView vkDestroyBufferView; //Todo
    PFN_vkDestroyCommandPool vkDestroyCommandPool;
    PFN_vkDestroyDescriptorPool vkDestroyDescriptorPool; //Todo
    PFN_vkDestroyDescriptorSetLayout vkDestroyDescriptorSetLayout; //Todo
    PFN_vkDestroyDevice vkDestroyDevice;
    PFN_vkDestroyEvent vkDestroyEvent; //Todo
    PFN_vkDestroyFence vkDestroyFence;
    PFN_vkDestroyFramebuffer vkDestroyFramebuffer;
    PFN_vkDestroyImage vkDestroyImage;
    PFN_vkDestroyImageView vkDestroyImageView;
    PFN_vkDestroyPipeline vkDestroyPipeline;
    PFN_vkDestroyPipelineCache vkDestroyPipelineCache; //Todo
    PFN_vkDestroyPipelineLayout vkDestroyPipelineLayout;
    PFN_vkDestroyQueryPool vkDestroyQueryPool; //Todo
    PFN_vkDestroyRenderPass vkDestroyRenderPass;
    PFN_vkDestroySampler vkDestroySampler; //Todo
    PFN_vkDestroySemaphore vkDestroySemaphore;
    PFN_vkDestroyShaderModule vkDestroyShaderModule;
    PFN_vkDeviceWaitIdle vkDeviceWaitIdle;
    PFN_vkEndCommandBuffer vkEndCommandBuffer;
    PFN_vkFlushMappedMemoryRanges vkFlushMappedMemoryRanges; //Todo
    PFN_vkFreeCommandBuffers vkFreeCommandBuffers;
    PFN_vkFreeDescriptorSets vkFreeDescriptorSets; //Todo
    PFN_vkFreeMemory vkFreeMemory;
    PFN_vkGetBufferMemoryRequirements vkGetBufferMemoryRequirements;
    PFN_vkGetDeviceMemoryCommitment vkGetDeviceMemoryCommitment; //Todo
    PFN_vkGetDeviceQueue vkGetDeviceQueue;
    PFN_vkGetEventStatus vkGetEventStatus; //Todo
    PFN_vkGetFenceStatus vkGetFenceStatus; //Todo
    PFN_vkGetImageMemoryRequirements vkGetImageMemoryRequirements; //Todo
    PFN_vkGetImageSparseMemoryRequirements vkGetImageSparseMemoryRequirements; //Todo
    PFN_vkGetImageSubresourceLayout vkGetImageSubresourceLayout; //Todo
    PFN_vkGetPipelineCacheData vkGetPipelineCacheData; //Todo
    PFN_vkGetQueryPoolResults vkGetQueryPoolResults; //Todo
    PFN_vkGetRenderAreaGranularity vkGetRenderAreaGranularity; //Todo
    PFN_vkInvalidateMappedMemoryRanges vkInvalidateMappedMemoryRanges; //Todo
    PFN_vkMapMemory vkMapMemory;
    PFN_vkMergePipelineCaches vkMergePipelineCaches; //Todo
    PFN_vkQueueBindSparse vkQueueBindSparse; //Todo
    PFN_vkQueueSubmit vkQueueSubmit;
    PFN_vkQueueWaitIdle vkQueueWaitIdle;
    PFN_vkResetCommandBuffer vkResetCommandBuffer; //Todo
    PFN_vkResetCommandPool vkResetCommandPool; //Todo
    PFN_vkResetDescriptorPool vkResetDescriptorPool; //Todo
    PFN_vkResetEvent vkResetEvent; //Todo
    PFN_vkResetFences vkResetFences;
    PFN_vkSetEvent vkSetEvent; //Todo
    PFN_vkUnmapMemory vkUnmapMemory;
    PFN_vkUpdateDescriptorSets vkUpdateDescriptorSets; //Todo
    PFN_vkWaitForFences vkWaitForFences;
} LoaderFuncPtrs;
