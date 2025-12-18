package driver

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

type vulkanDriver struct {
	instance VkInstance
	device   VkDevice
	funcPtrs *C.DriverFuncPtrs

	version common.APIVersion

	objStore *VulkanObjectStore
}

func createVulkanDriver(funcPtrs *C.DriverFuncPtrs, objStore *VulkanObjectStore, instance VkInstance, device VkDevice) (*vulkanDriver, error) {
	version := common.Vulkan1_0
	driver := &vulkanDriver{
		funcPtrs: funcPtrs,
		instance: instance,
		device:   device,

		objStore: objStore,
	}

	if funcPtrs.vkEnumerateInstanceVersion != nil {
		var versionBits Uint32
		_, err := driver.VkEnumerateInstanceVersion(&versionBits)
		if err != nil {
			return nil, err
		}

		version = common.APIVersion(versionBits)
	}

	driver.version = version
	return driver, nil
}

func CreateDriverFromProcAddr(procAddr unsafe.Pointer) (*vulkanDriver, error) {
	baseFuncPtr := (C.PFN_vkGetInstanceProcAddr)(procAddr)
	funcPtrs := (*C.DriverFuncPtrs)(C.malloc(C.sizeof_struct_DriverFuncPtrs))
	C.driverFuncPtrs_populate(baseFuncPtr, funcPtrs)

	return createVulkanDriver(funcPtrs, NewObjectStore(), VkInstance(NullHandle), VkDevice(NullHandle))
}

func (l *vulkanDriver) ObjectStore() *VulkanObjectStore {
	return l.objStore
}

func (l *vulkanDriver) Destroy() {
	C.free(unsafe.Pointer(l.funcPtrs))
}

func (l *vulkanDriver) CreateInstanceDriver(instance VkInstance) (Driver, error) {
	instanceFuncPtrs := (*C.DriverFuncPtrs)(C.malloc(C.sizeof_struct_DriverFuncPtrs))
	C.instanceFuncPtrs_populate((C.VkInstance)(unsafe.Pointer(instance)), l.funcPtrs, instanceFuncPtrs)

	return createVulkanDriver(instanceFuncPtrs, l.objStore, instance, VkDevice(NullHandle))
}

func (l *vulkanDriver) CreateDeviceDriver(device VkDevice) (Driver, error) {
	if l.instance == VkInstance(NullHandle) {
		return nil, errors.New("attempted to call instance driver function on a basic driver")
	}

	deviceFuncPtrs := (*C.DriverFuncPtrs)(C.malloc(C.sizeof_struct_DriverFuncPtrs))
	C.deviceFuncPtrs_populate((C.VkDevice)(unsafe.Pointer(device)), l.funcPtrs, deviceFuncPtrs)

	return createVulkanDriver(deviceFuncPtrs, l.objStore, l.instance, device)
}

func (l *vulkanDriver) LoadProcAddr(name *Char) unsafe.Pointer {
	if l.device != VkDevice(NullHandle) {
		return unsafe.Pointer(C.device_proc_addr(l.funcPtrs, C.VkDevice(unsafe.Pointer(l.device)), (*C.char)(name)))
	} else {
		return unsafe.Pointer(C.instance_proc_addr(l.funcPtrs, C.VkInstance(unsafe.Pointer(l.instance)), (*C.char)(name)))
	}
}

func (l *vulkanDriver) LoadInstanceProcAddr(name *Char) unsafe.Pointer {
	return unsafe.Pointer(C.instance_proc_addr(l.funcPtrs, C.VkInstance(unsafe.Pointer(l.instance)), (*C.char)(name)))
}

func (l *vulkanDriver) Version() common.APIVersion {
	return l.version
}
