package loader

import "C"

/*
#cgo noescape instance_proc_addr
#cgo noescape device_proc_addr
#cgo noescape cgoCreateInstance
#cgo noescape cgoDestroyInstance
#cgo noescape cgoEnumeratePhysicalDevices
#cgo noescape cgoGetPhysicalDeviceFeatures
#cgo noescape cgoGetPhysicalDeviceFormatProperties
#cgo noescape cgoGetPhysicalDeviceImageFormatProperties
#cgo noescape cgoGetPhysicalDeviceProperties
#cgo noescape cgoGetPhysicalDeviceQueueFamilyProperties
#cgo noescape cgoGetPhysicalDeviceMemoryProperties
#cgo noescape cgoCreateDevice
#cgo noescape cgoDestroyDevice
#cgo noescape cgoEnumerateInstanceExtensionProperties
#cgo noescape cgoEnumerateInstanceLayerProperties
#cgo noescape cgoEnumerateDeviceExtensionProperties
#cgo noescape cgoEnumerateDeviceLayerProperties
#cgo noescape cgoGetDeviceQueue
#cgo noescape cgoQueueSubmit
#cgo noescape cgoQueueWaitIdle
#cgo noescape cgoDeviceWaitIdle
#cgo noescape cgoAllocateMemory
#cgo noescape cgoFreeMemory
#cgo noescape cgoMapMemory
#cgo noescape cgoUnmapMemory
#cgo noescape cgoFlushMappedMemoryRanges
#cgo noescape cgoInvalidateMappedMemoryRanges
#cgo noescape cgoGetDeviceMemoryCommitment
#cgo noescape cgoBindBufferMemory
#cgo noescape cgoBindImageMemory
#cgo noescape cgoGetBufferMemoryRequirements
#cgo noescape cgoGetImageMemoryRequirements
#cgo noescape cgoGetImageSparseMemoryRequirements
#cgo noescape cgoGetPhysicalDeviceSparseImageFormatProperties
#cgo noescape cgoQueueBindSparse
#cgo noescape cgoCreateFence
#cgo noescape cgoDestroyFence
#cgo noescape cgoResetFences
#cgo noescape cgoGetFenceStatus
#cgo noescape cgoWaitForFences
#cgo noescape cgoCreateSemaphore
#cgo noescape cgoDestroySemaphore
#cgo noescape cgoCreateEvent
#cgo noescape cgoDestroyEvent
#cgo noescape cgoGetEventStatus
#cgo noescape cgoSetEvent
#cgo noescape cgoResetEvent
#cgo noescape cgoCreateQueryPool
#cgo noescape cgoDestroyQueryPool
#cgo noescape cgoGetQueryPoolResults
#cgo noescape cgoCreateBuffer
#cgo noescape cgoDestroyBuffer
#cgo noescape cgoCreateBufferView
#cgo noescape cgoDestroyBufferView
#cgo noescape cgoCreateImage
#cgo noescape cgoDestroyImage
#cgo noescape cgoGetImageSubresourceLayout
#cgo noescape cgoCreateImageView
#cgo noescape cgoDestroyImageView
#cgo noescape cgoCreateShaderModule
#cgo noescape cgoDestroyShaderModule
#cgo noescape cgoCreatePipelineCache
#cgo noescape cgoDestroyPipelineCache
#cgo noescape cgoGetPipelineCacheData
#cgo noescape cgoMergePipelineCaches
#cgo noescape cgoCreateGraphicsPipelines
#cgo noescape cgoCreateComputePipelines
#cgo noescape cgoDestroyPipeline
#cgo noescape cgoCreatePipelineLayout
#cgo noescape cgoDestroyPipelineLayout
#cgo noescape cgoCreateSampler
#cgo noescape cgoDestroySampler
#cgo noescape cgoCreateDescriptorSetLayout
#cgo noescape cgoDestroyDescriptorSetLayout
#cgo noescape cgoCreateDescriptorPool
#cgo noescape cgoDestroyDescriptorPool
#cgo noescape cgoResetDescriptorPool
#cgo noescape cgoAllocateDescriptorSets
#cgo noescape cgoFreeDescriptorSets
#cgo noescape cgoUpdateDescriptorSets
#cgo noescape cgoCreateFramebuffer
#cgo noescape cgoDestroyFramebuffer
#cgo noescape cgoCreateRenderPass
#cgo noescape cgoDestroyRenderPass
#cgo noescape cgoGetRenderAreaGranularity
#cgo noescape cgoCreateCommandPool
#cgo noescape cgoDestroyCommandPool
#cgo noescape cgoResetCommandPool
#cgo noescape cgoAllocateCommandBuffers
#cgo noescape cgoFreeCommandBuffers
#cgo noescape cgoBeginCommandBuffer
#cgo noescape cgoEndCommandBuffer
#cgo noescape cgoResetCommandBuffer
#cgo noescape cgoCmdBindPipeline
#cgo noescape cgoCmdSetViewport
#cgo noescape cgoCmdSetScissor
#cgo noescape cgoCmdSetLineWidth
#cgo noescape cgoCmdSetDepthBias
#cgo noescape cgoCmdSetBlendConstants
#cgo noescape cgoCmdSetDepthBounds
#cgo noescape cgoCmdSetStencilCompareMask
#cgo noescape cgoCmdSetStencilWriteMask
#cgo noescape cgoCmdSetStencilReference
#cgo noescape cgoCmdBindDescriptorSets
#cgo noescape cgoCmdBindIndexBuffer
#cgo noescape cgoCmdBindVertexBuffers
#cgo noescape cgoCmdDraw
#cgo noescape cgoCmdDrawIndexed
#cgo noescape cgoCmdDrawIndirect
#cgo noescape cgoCmdDrawIndexedIndirect
#cgo noescape cgoCmdDispatch
#cgo noescape cgoCmdDispatchIndirect
#cgo noescape cgoCmdCopyBuffer
#cgo noescape cgoCmdCopyImage
#cgo noescape cgoCmdBlitImage
#cgo noescape cgoCmdCopyBufferToImage
#cgo noescape cgoCmdCopyImageToBuffer
#cgo noescape cgoCmdUpdateBuffer
#cgo noescape cgoCmdFillBuffer
#cgo noescape cgoCmdClearColorImage
#cgo noescape cgoCmdClearDepthStencilImage
#cgo noescape cgoCmdClearAttachments
#cgo noescape cgoCmdResolveImage
#cgo noescape cgoCmdSetEvent
#cgo noescape cgoCmdResetEvent
#cgo noescape cgoCmdWaitEvents
#cgo noescape cgoCmdPipelineBarrier
#cgo noescape cgoCmdBeginQuery
#cgo noescape cgoCmdEndQuery
#cgo noescape cgoCmdResetQueryPool
#cgo noescape cgoCmdWriteTimestamp
#cgo noescape cgoCmdCopyQueryPoolResults
#cgo noescape cgoCmdPushConstants
#cgo noescape cgoCmdBeginRenderPass
#cgo noescape cgoCmdNextSubpass
#cgo noescape cgoCmdEndRenderPass
#cgo noescape cgoCmdExecuteCommands
#cgo noescape cgoEnumerateInstanceVersion
#cgo noescape cgoBindBufferMemory
#cgo noescape cgoBindImageMemory
#cgo noescape cgoGetDeviceGroupPeerMemoryFeatures
#cgo noescape cgoCmdSetDeviceMask
#cgo noescape cgoCmdDispatchBase
#cgo noescape cgoEnumeratePhysicalDeviceGroups
#cgo noescape cgoGetImageMemoryRequirements
#cgo noescape cgoGetBufferMemoryRequirements
#cgo noescape cgoGetImageSparseMemoryRequirements
#cgo noescape cgoGetPhysicalDeviceFeatures
#cgo noescape cgoGetPhysicalDeviceProperties
#cgo noescape cgoGetPhysicalDeviceFormatProperties
#cgo noescape cgoGetPhysicalDeviceImageFormatProperties
#cgo noescape cgoGetPhysicalDeviceQueueFamilyProperties
#cgo noescape cgoGetPhysicalDeviceMemoryProperties
#cgo noescape cgoGetPhysicalDeviceSparseImageFormatProperties
#cgo noescape cgoTrimCommandPool
#cgo noescape cgoGetDeviceQueue
#cgo noescape cgoCreateSamplerYcbcrConversion
#cgo noescape cgoDestroySamplerYcbcrConversion
#cgo noescape cgoCreateDescriptorUpdateTemplate
#cgo noescape cgoDestroyDescriptorUpdateTemplate
#cgo noescape cgoUpdateDescriptorSetWithTemplate
#cgo noescape cgoGetPhysicalDeviceExternalBufferProperties
#cgo noescape cgoGetPhysicalDeviceExternalFenceProperties
#cgo noescape cgoGetPhysicalDeviceExternalSemaphoreProperties
#cgo noescape cgoGetDescriptorSetLayoutSupport
#cgo noescape cgoCmdDrawIndirectCount
#cgo noescape cgoCmdDrawIndexedIndirectCount
#cgo noescape cgoCreateRenderPass
#cgo noescape cgoCmdBeginRenderPass
#cgo noescape cgoCmdNextSubpass
#cgo noescape cgoCmdEndRenderPass
#cgo noescape cgoResetQueryPool
#cgo noescape cgoGetSemaphoreCounterValue
#cgo noescape cgoWaitSemaphores
#cgo noescape cgoSignalSemaphore
#cgo noescape cgoGetBufferDeviceAddress
#cgo noescape cgoGetBufferOpaqueCaptureAddress
#cgo noescape cgoGetDeviceMemoryOpaqueCaptureAddress

#include "func_ptrs.h"
#include "func_ptrs_def.h"

PFN_vkVoidFunction instance_proc_addr(DriverFuncPtrs *funcPtrs, VkInstance instance, const char *procName) {
	PFN_vkGetInstanceProcAddr procAddr = funcPtrs->vkGetInstanceProcAddr;
	return procAddr(instance, procName);
}

PFN_vkVoidFunction device_proc_addr(DriverFuncPtrs *funcPtrs, VkDevice device, const char *procName) {
	PFN_vkGetDeviceProcAddr procAddr = funcPtrs->vkGetDeviceProcAddr;
	return procAddr(device, procName);
}
*/
import "C"
import (
	"unsafe"

	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3/common"
)

type vulkanLoader struct {
	instance VkInstance
	device   VkDevice
	funcPtrs *C.DriverFuncPtrs

	version common.APIVersion
}

func createVulkanLoader(funcPtrs *C.DriverFuncPtrs, instance VkInstance, device VkDevice) (*vulkanLoader, error) {
	version := common.Vulkan1_0
	loader := &vulkanLoader{
		funcPtrs: funcPtrs,
		instance: instance,
		device:   device,
	}

	if funcPtrs.vkEnumerateInstanceVersion != nil {
		var versionBits Uint32
		_, err := loader.VkEnumerateInstanceVersion(&versionBits)
		if err != nil {
			return nil, err
		}

		version = common.APIVersion(versionBits)
	}

	loader.version = version
	return loader, nil
}

func CreateLoaderFromProcAddr(procAddr unsafe.Pointer) (Loader, error) {
	baseFuncPtr := (C.PFN_vkGetInstanceProcAddr)(procAddr)
	funcPtrs := (*C.DriverFuncPtrs)(C.malloc(C.sizeof_struct_DriverFuncPtrs))
	C.driverFuncPtrs_populate(baseFuncPtr, funcPtrs)

	return createVulkanLoader(funcPtrs, VkInstance(NullHandle), VkDevice(NullHandle))
}

func (l *vulkanLoader) InstanceHandle() VkInstance {
	return l.instance
}

func (l *vulkanLoader) DeviceHandle() VkDevice {
	return l.device
}

func (l *vulkanLoader) Destroy() {
	C.free(unsafe.Pointer(l.funcPtrs))
}

func (l *vulkanLoader) CreateInstanceLoader(instance VkInstance) (Loader, error) {
	instanceFuncPtrs := (*C.DriverFuncPtrs)(C.malloc(C.sizeof_struct_DriverFuncPtrs))
	C.instanceFuncPtrs_populate((C.VkInstance)(unsafe.Pointer(instance)), l.funcPtrs, instanceFuncPtrs)

	return createVulkanLoader(instanceFuncPtrs, instance, VkDevice(NullHandle))
}

func (l *vulkanLoader) CreateDeviceLoader(device VkDevice) (Loader, error) {
	if l.instance == VkInstance(NullHandle) {
		return nil, errors.New("attempted to call instance loader function on a basic loader")
	}

	deviceFuncPtrs := (*C.DriverFuncPtrs)(C.malloc(C.sizeof_struct_DriverFuncPtrs))
	C.deviceFuncPtrs_populate((C.VkDevice)(unsafe.Pointer(device)), l.funcPtrs, deviceFuncPtrs)

	return createVulkanLoader(deviceFuncPtrs, l.instance, device)
}

func (l *vulkanLoader) LoadProcAddr(name *Char) unsafe.Pointer {
	if l.device != VkDevice(NullHandle) {
		return unsafe.Pointer(C.device_proc_addr(l.funcPtrs, C.VkDevice(unsafe.Pointer(l.device)), (*C.char)(name)))
	} else {
		return unsafe.Pointer(C.instance_proc_addr(l.funcPtrs, C.VkInstance(unsafe.Pointer(l.instance)), (*C.char)(name)))
	}
}

func (l *vulkanLoader) LoadInstanceProcAddr(name *Char) unsafe.Pointer {
	return unsafe.Pointer(C.instance_proc_addr(l.funcPtrs, C.VkInstance(unsafe.Pointer(l.instance)), (*C.char)(name)))
}

func (l *vulkanLoader) Version() common.APIVersion {
	return l.version
}
