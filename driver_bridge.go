package core

/*
#cgo windows LDFLAGS: -lvulkan
#cgo linux freebsd openbsd pkg-config: vulkan
#include "driver.h"
*/
import "C"
import (
	"unsafe"
)

func (l *vulkanDriver) VkEnumerateInstanceExtensionProperties(pLayerName *Char, pPropertyCount *Uint32, pProperties *VkExtensionProperties) (VkResult, error) {
	res := VkResult(C.cgoEnumerateInstanceExtensionProperties(l.funcPtrs.vkEnumerateInstanceExtensionProperties,
		(*C.char)(pLayerName),
		(*C.uint32_t)(pPropertyCount),
		(*C.VkExtensionProperties)(pProperties)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkEnumerateInstanceLayerProperties(pPropertyCount *Uint32, pProperties *VkLayerProperties) (VkResult, error) {
	res := VkResult(C.cgoEnumerateInstanceLayerProperties(l.funcPtrs.vkEnumerateInstanceLayerProperties,
		(*C.uint32_t)(pPropertyCount),
		(*C.VkLayerProperties)(pProperties)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkCreateInstance(pCreateInfo *VkInstanceCreateInfo, pAllocator *VkAllocationCallbacks, pInstance *VkInstance) (VkResult, error) {
	res := VkResult(C.cgoCreateInstance(l.funcPtrs.vkCreateInstance,
		(*C.VkInstanceCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkInstance)(pInstance)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkEnumeratePhysicalDevices(instance VkInstance, pPhysicalDeviceCount *Uint32, pPhysicalDevices *VkPhysicalDevice) (VkResult, error) {
	if l.instance == nil {
		panic("attempted to call instance driver function on a basic driver")
	}

	res := VkResult(C.cgoEnumeratePhysicalDevices(l.funcPtrs.vkEnumeratePhysicalDevices,
		(C.VkInstance)(instance),
		(*C.uint32_t)(pPhysicalDeviceCount),
		(*C.VkPhysicalDevice)(pPhysicalDevices)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyInstance(instance VkInstance, pAllocator *VkAllocationCallbacks) {
	if l.instance == nil {
		panic("attempted to call instance driver function on a basic driver")
	}

	C.cgoDestroyInstance(l.funcPtrs.vkDestroyInstance,
		(C.VkInstance)(instance),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkGetPhysicalDeviceFeatures(physicalDevice VkPhysicalDevice, pFeatures *VkPhysicalDeviceFeatures) {
	if l.instance == nil {
		panic("attempted to call instance driver function on a basic driver")
	}

	C.cgoGetPhysicalDeviceFeatures(l.funcPtrs.vkGetPhysicalDeviceFeatures,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.VkPhysicalDeviceFeatures)(pFeatures))
}

func (l *vulkanDriver) VkGetPhysicalDeviceFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, pFormatProperties *VkFormatProperties) {
	if l.instance == nil {
		panic("attempted to call instance driver function on a basic driver")
	}

	C.cgoGetPhysicalDeviceFormatProperties(l.funcPtrs.vkGetPhysicalDeviceFormatProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(C.VkFormat)(format),
		(*C.VkFormatProperties)(pFormatProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceImageFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, t VkImageType, tiling VkImageTiling, usage VkImageUsageFlags, flags VkImageCreateFlags, pImageFormatProperties *VkImageFormatProperties) (VkResult, error) {
	if l.instance == nil {
		panic("attempted to call instance driver function on a basic driver")
	}

	res := VkResult(C.cgoGetPhysicalDeviceImageFormatProperties(l.funcPtrs.vkGetPhysicalDeviceImageFormatProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(C.VkFormat)(format),
		(C.VkImageType)(t),
		(C.VkImageTiling)(tiling),
		(C.VkImageUsageFlags)(usage),
		(C.VkImageCreateFlags)(flags),
		(*C.VkImageFormatProperties)(pImageFormatProperties)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkGetPhysicalDeviceProperties(physicalDevice VkPhysicalDevice, pProperties *VkPhysicalDeviceProperties) {
	if l.instance == nil {
		panic("attempted to call instance driver function on a basic driver")
	}

	C.cgoGetPhysicalDeviceProperties(l.funcPtrs.vkGetPhysicalDeviceProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.VkPhysicalDeviceProperties)(pProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceQueueFamilyProperties(physicalDevice VkPhysicalDevice, pQueueFamilyPropertyCount *Uint32, pQueueFamilyProperties *VkQueueFamilyProperties) {
	if l.instance == nil {
		panic("attempted to call instance driver function on a basic driver")
	}

	C.cgoGetPhysicalDeviceQueueFamilyProperties(l.funcPtrs.vkGetPhysicalDeviceQueueFamilyProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.uint32_t)(pQueueFamilyPropertyCount),
		(*C.VkQueueFamilyProperties)(pQueueFamilyProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceMemoryProperties(physicalDevice VkPhysicalDevice, pMemoryProperties *VkPhysicalDeviceMemoryProperties) {
	if l.instance == nil {
		panic("attempted to call instance driver function on a basic driver")
	}

	C.cgoGetPhysicalDeviceMemoryProperties(l.funcPtrs.vkGetPhysicalDeviceMemoryProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.VkPhysicalDeviceMemoryProperties)(pMemoryProperties))
}

func (l *vulkanDriver) VkEnumerateDeviceExtensionProperties(physicalDevice VkPhysicalDevice, pLayerName *Char, pPropertyCount *Uint32, pProperties *VkExtensionProperties) (VkResult, error) {
	if l.instance == nil {
		panic("attempted to call instance driver function on a basic driver")
	}

	res := VkResult(C.cgoEnumerateDeviceExtensionProperties(l.funcPtrs.vkEnumerateDeviceExtensionProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.char)(pLayerName),
		(*C.uint32_t)(pPropertyCount),
		(*C.VkExtensionProperties)(pProperties)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkEnumerateDeviceLayerProperties(physicalDevice VkPhysicalDevice, pPropertyCount *Uint32, pProperties *VkLayerProperties) (VkResult, error) {
	if l.instance == nil {
		panic("attempted to call instance driver function on a basic driver")
	}

	res := VkResult(C.cgoEnumerateDeviceLayerProperties(l.funcPtrs.vkEnumerateDeviceLayerProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.uint32_t)(pPropertyCount),
		(*C.VkLayerProperties)(pProperties)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkGetPhysicalDeviceSparseImageFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, t VkImageType, samples VkSampleCountFlagBits, usage VkImageUsageFlags, tiling VkImageTiling, pPropertyCount *Uint32, pProperties *VkSparseImageFormatProperties) {
	if l.instance == nil {
		panic("attempted to call instance driver function on a basic driver")
	}

	C.cgoGetPhysicalDeviceSparseImageFormatProperties(l.funcPtrs.vkGetPhysicalDeviceSparseImageFormatProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(C.VkFormat)(format),
		(C.VkImageType)(t),
		(C.VkSampleCountFlagBits)(samples),
		(C.VkImageUsageFlags)(usage),
		(C.VkImageTiling)(tiling),
		(*C.uint32_t)(pPropertyCount),
		(*C.VkSparseImageFormatProperties)(pProperties))
}

func (l *vulkanDriver) VkCreateDevice(physicalDevice VkPhysicalDevice, pCreateInfo *VkDeviceCreateInfo, pAllocator *VkAllocationCallbacks, pDevice *VkDevice) (VkResult, error) {
	if l.instance == nil {
		panic("attempted to call instance driver function on a basic driver")
	}

	res := VkResult(C.cgoCreateDevice(l.funcPtrs.vkCreateDevice,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.VkDeviceCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkDevice)(pDevice)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyDevice(device VkDevice, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyDevice(l.funcPtrs.vkDestroyDevice,
		(C.VkDevice)(device),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkGetDeviceQueue(device VkDevice, queueFamilyIndex Uint32, queueIndex Uint32, pQueue *VkQueue) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoGetDeviceQueue(l.funcPtrs.vkGetDeviceQueue,
		(C.VkDevice)(device),
		(C.uint32_t)(queueFamilyIndex),
		(C.uint32_t)(queueIndex),
		(*C.VkQueue)(pQueue))
}

func (l *vulkanDriver) VkQueueSubmit(queue VkQueue, submitCount Uint32, pSubmits *VkSubmitInfo, fence VkFence) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoQueueSubmit(l.funcPtrs.vkQueueSubmit,
		(C.VkQueue)(queue),
		(C.uint32_t)(submitCount),
		(*C.VkSubmitInfo)(pSubmits),
		(C.VkFence)(fence)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkQueueWaitIdle(queue VkQueue) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoQueueWaitIdle(l.funcPtrs.vkQueueWaitIdle,
		(C.VkQueue)(queue)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDeviceWaitIdle(device VkDevice) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoDeviceWaitIdle(l.funcPtrs.vkDeviceWaitIdle,
		(C.VkDevice)(device)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkAllocateMemory(device VkDevice, pAllocateInfo *VkMemoryAllocateInfo, pAllocator *VkAllocationCallbacks, pMemory *VkDeviceMemory) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoAllocateMemory(l.funcPtrs.vkAllocateMemory,
		(C.VkDevice)(device),
		(*C.VkMemoryAllocateInfo)(pAllocateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkDeviceMemory)(pMemory)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkFreeMemory(device VkDevice, memory VkDeviceMemory, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoFreeMemory(l.funcPtrs.vkFreeMemory,
		(C.VkDevice)(device),
		(C.VkDeviceMemory)(memory),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkMapMemory(device VkDevice, memory VkDeviceMemory, offset VkDeviceSize, size VkDeviceSize, flags VkMemoryMapFlags, ppData *unsafe.Pointer) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoMapMemory(l.funcPtrs.vkMapMemory,
		(C.VkDevice)(device),
		(C.VkDeviceMemory)(memory),
		(C.VkDeviceSize)(offset),
		(C.VkDeviceSize)(size),
		(C.VkMemoryMapFlags)(flags),
		ppData))
	return res, res.ToError()
}

func (l *vulkanDriver) VkUnmapMemory(device VkDevice, memory VkDeviceMemory) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoUnmapMemory(l.funcPtrs.vkUnmapMemory,
		(C.VkDevice)(device),
		(C.VkDeviceMemory)(memory))
}

func (l *vulkanDriver) VkFlushMappedMemoryRanges(device VkDevice, memoryRangeCount Uint32, pMemoryRanges *VkMappedMemoryRange) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoFlushMappedMemoryRanges(l.funcPtrs.vkFlushMappedMemoryRanges,
		(C.VkDevice)(device),
		(C.uint32_t)(memoryRangeCount),
		(*C.VkMappedMemoryRange)(pMemoryRanges)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkInvalidateMappedMemoryRanges(device VkDevice, memoryRangeCount Uint32, pMemoryRanges *VkMappedMemoryRange) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoInvalidateMappedMemoryRanges(l.funcPtrs.vkInvalidateMappedMemoryRanges,
		(C.VkDevice)(device),
		(C.uint32_t)(memoryRangeCount),
		(*C.VkMappedMemoryRange)(pMemoryRanges)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkGetDeviceMemoryCommitment(device VkDevice, memory VkDeviceMemory, pCommittedMemoryInBytes *VkDeviceSize) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoGetDeviceMemoryCommitment(l.funcPtrs.vkGetDeviceMemoryCommitment,
		(C.VkDevice)(device),
		(C.VkDeviceMemory)(memory),
		(*C.VkDeviceSize)(pCommittedMemoryInBytes))
}

func (l *vulkanDriver) VkBindBufferMemory(device VkDevice, buffer VkBuffer, memory VkDeviceMemory, memoryOffset VkDeviceSize) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoBindBufferMemory(l.funcPtrs.vkBindBufferMemory,
		(C.VkDevice)(device),
		(C.VkBuffer)(buffer),
		(C.VkDeviceMemory)(memory),
		(C.VkDeviceSize)(memoryOffset)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkBindImageMemory(device VkDevice, image VkImage, memory VkDeviceMemory, memoryOffset VkDeviceSize) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoBindImageMemory(l.funcPtrs.vkBindImageMemory,
		(C.VkDevice)(device),
		(C.VkImage)(image),
		(C.VkDeviceMemory)(memory),
		(C.VkDeviceSize)(memoryOffset)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkGetBufferMemoryRequirements(device VkDevice, buffer VkBuffer, pMemoryRequirements *VkMemoryRequirements) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoGetBufferMemoryRequirements(l.funcPtrs.vkGetBufferMemoryRequirements,
		(C.VkDevice)(device),
		(C.VkBuffer)(buffer),
		(*C.VkMemoryRequirements)(pMemoryRequirements))
}

func (l *vulkanDriver) VkGetImageMemoryRequirements(device VkDevice, image VkImage, pMemoryRequirements *VkMemoryRequirements) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoGetImageMemoryRequirements(l.funcPtrs.vkGetImageMemoryRequirements,
		(C.VkDevice)(device),
		(C.VkImage)(image),
		(*C.VkMemoryRequirements)(pMemoryRequirements))
}

func (l *vulkanDriver) VkGetImageSparseMemoryRequirements(device VkDevice, image VkImage, pSparseMemoryRequirementCount *Uint32, pSparseMemoryRequirements *VkSparseImageMemoryRequirements) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoGetImageSparseMemoryRequirements(l.funcPtrs.vkGetImageSparseMemoryRequirements,
		(C.VkDevice)(device),
		(C.VkImage)(image),
		(*C.uint32_t)(pSparseMemoryRequirementCount),
		(*C.VkSparseImageMemoryRequirements)(pSparseMemoryRequirements))
}

func (l *vulkanDriver) VkQueueBindSparse(queue VkQueue, bindInfoCount Uint32, pBindInfo *VkBindSparseInfo, fence VkFence) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoQueueBindSparse(l.funcPtrs.vkQueueBindSparse,
		(C.VkQueue)(queue),
		(C.uint32_t)(bindInfoCount),
		(*C.VkBindSparseInfo)(pBindInfo),
		(C.VkFence)(fence)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkCreateFence(device VkDevice, pCreateInfo *VkFenceCreateInfo, pAllocator *VkAllocationCallbacks, pFence *VkFence) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreateFence(l.funcPtrs.vkCreateFence,
		(C.VkDevice)(device),
		(*C.VkFenceCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkFence)(pFence)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyFence(device VkDevice, fence VkFence, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyFence(l.funcPtrs.vkDestroyFence,
		(C.VkDevice)(device),
		(C.VkFence)(fence),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkResetFences(device VkDevice, fenceCount Uint32, pFences *VkFence) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoResetFences(l.funcPtrs.vkResetFences,
		(C.VkDevice)(device),
		(C.uint32_t)(fenceCount),
		(*C.VkFence)(pFences)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkGetFenceStatus(device VkDevice, fence VkFence) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoGetFenceStatus(l.funcPtrs.vkGetFenceStatus,
		(C.VkDevice)(device),
		(C.VkFence)(fence)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkWaitForFences(device VkDevice, fenceCount Uint32, pFences *VkFence, waitAll VkBool32, timeout Uint64) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoWaitForFences(l.funcPtrs.vkWaitForFences,
		(C.VkDevice)(device),
		(C.uint32_t)(fenceCount),
		(*C.VkFence)(pFences),
		(C.VkBool32)(waitAll),
		(C.uint64_t)(timeout)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkCreateSemaphore(device VkDevice, pCreateInfo *VkSemaphoreCreateInfo, pAllocator *VkAllocationCallbacks, pSemaphore *VkSemaphore) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreateSemaphore(l.funcPtrs.vkCreateSemaphore,
		(C.VkDevice)(device),
		(*C.VkSemaphoreCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkSemaphore)(pSemaphore)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroySemaphore(device VkDevice, semaphore VkSemaphore, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroySemaphore(l.funcPtrs.vkDestroySemaphore,
		(C.VkDevice)(device),
		(C.VkSemaphore)(semaphore),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateEvent(device VkDevice, pCreateInfo *VkEventCreateInfo, pAllocator *VkAllocationCallbacks, pEvent *VkEvent) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreateEvent(l.funcPtrs.vkCreateEvent,
		(C.VkDevice)(device),
		(*C.VkEventCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkEvent)(pEvent)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyEvent(device VkDevice, event VkEvent, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyEvent(l.funcPtrs.vkDestroyEvent,
		(C.VkDevice)(device),
		(C.VkEvent)(event),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkGetEventStatus(device VkDevice, event VkEvent) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoGetEventStatus(l.funcPtrs.vkGetEventStatus,
		(C.VkDevice)(device),
		(C.VkEvent)(event)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkSetEvent(device VkDevice, event VkEvent) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoSetEvent(l.funcPtrs.vkSetEvent,
		(C.VkDevice)(device),
		(C.VkEvent)(event)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkResetEvent(device VkDevice, event VkEvent) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoResetEvent(l.funcPtrs.vkResetEvent,
		(C.VkDevice)(device),
		(C.VkEvent)(event)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkCreateQueryPool(device VkDevice, pCreateInfo *VkQueryPoolCreateInfo, pAllocator *VkAllocationCallbacks, pQueryPool *VkQueryPool) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreateQueryPool(l.funcPtrs.vkCreateQueryPool,
		(C.VkDevice)(device),
		(*C.VkQueryPoolCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkQueryPool)(pQueryPool)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyQueryPool(device VkDevice, queryPool VkQueryPool, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyQueryPool(l.funcPtrs.vkDestroyQueryPool,
		(C.VkDevice)(device),
		(C.VkQueryPool)(queryPool),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkGetQueryPoolResults(device VkDevice, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32, dataSize Size, pData unsafe.Pointer, stride VkDeviceSize, flags VkQueryResultFlags) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoGetQueryPoolResults(l.funcPtrs.vkGetQueryPoolResults,
		(C.VkDevice)(device),
		(C.VkQueryPool)(queryPool),
		(C.uint32_t)(firstQuery),
		(C.uint32_t)(queryCount),
		(C.size_t)(dataSize),
		pData,
		(C.VkDeviceSize)(stride),
		(C.VkQueryResultFlags)(flags)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkCreateBuffer(device VkDevice, pCreateInfo *VkBufferCreateInfo, pAllocator *VkAllocationCallbacks, pBuffer *VkBuffer) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreateBuffer(l.funcPtrs.vkCreateBuffer,
		(C.VkDevice)(device),
		(*C.VkBufferCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkBuffer)(pBuffer)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyBuffer(device VkDevice, buffer VkBuffer, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyBuffer(l.funcPtrs.vkDestroyBuffer,
		(C.VkDevice)(device),
		(C.VkBuffer)(buffer),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateBufferView(device VkDevice, pCreateInfo *VkBufferViewCreateInfo, pAllocator *VkAllocationCallbacks, pView *VkBufferView) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreateBufferView(l.funcPtrs.vkCreateBufferView,
		(C.VkDevice)(device),
		(*C.VkBufferViewCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkBufferView)(pView)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyBufferView(device VkDevice, bufferView VkBufferView, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyBufferView(l.funcPtrs.vkDestroyBufferView,
		(C.VkDevice)(device),
		(C.VkBufferView)(bufferView),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateImage(device VkDevice, pCreateInfo *VkImageCreateInfo, pAllocator *VkAllocationCallbacks, pImage *VkImage) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreateImage(l.funcPtrs.vkCreateImage,
		(C.VkDevice)(device),
		(*C.VkImageCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkImage)(pImage)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyImage(device VkDevice, image VkImage, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyImage(l.funcPtrs.vkDestroyImage,
		(C.VkDevice)(device),
		(C.VkImage)(image),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkGetImageSubresourceLayout(device VkDevice, image VkImage, pSubresource *VkImageSubresource, pLayout *VkSubresourceLayout) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoGetImageSubresourceLayout(l.funcPtrs.vkGetImageSubresourceLayout,
		(C.VkDevice)(device),
		(C.VkImage)(image),
		(*C.VkImageSubresource)(pSubresource),
		(*C.VkSubresourceLayout)(pLayout))
}

func (l *vulkanDriver) VkCreateImageView(device VkDevice, pCreateInfo *VkImageViewCreateInfo, pAllocator *VkAllocationCallbacks, pView *VkImageView) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreateImageView(l.funcPtrs.vkCreateImageView,
		(C.VkDevice)(device),
		(*C.VkImageViewCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkImageView)(pView)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyImageView(device VkDevice, imageView VkImageView, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyImageView(l.funcPtrs.vkDestroyImageView,
		(C.VkDevice)(device),
		(C.VkImageView)(imageView),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateShaderModule(device VkDevice, pCreateInfo *VkShaderModuleCreateInfo, pAllocator *VkAllocationCallbacks, pShaderModule *VkShaderModule) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreateShaderModule(l.funcPtrs.vkCreateShaderModule,
		(C.VkDevice)(device),
		(*C.VkShaderModuleCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkShaderModule)(pShaderModule)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyShaderModule(device VkDevice, shaderModule VkShaderModule, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyShaderModule(l.funcPtrs.vkDestroyShaderModule,
		(C.VkDevice)(device),
		(C.VkShaderModule)(shaderModule),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreatePipelineCache(device VkDevice, pCreateInfo *VkPipelineCacheCreateInfo, pAllocator *VkAllocationCallbacks, pPipelineCache *VkPipelineCache) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreatePipelineCache(l.funcPtrs.vkCreatePipelineCache,
		(C.VkDevice)(device),
		(*C.VkPipelineCacheCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkPipelineCache)(pPipelineCache)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyPipelineCache(device VkDevice, pipelineCache VkPipelineCache, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyPipelineCache(l.funcPtrs.vkDestroyPipelineCache,
		(C.VkDevice)(device),
		(C.VkPipelineCache)(pipelineCache),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkGetPipelineCacheData(device VkDevice, pipelineCache VkPipelineCache, pDataSize *Size, pData unsafe.Pointer) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoGetPipelineCacheData(l.funcPtrs.vkGetPipelineCacheData,
		(C.VkDevice)(device),
		(C.VkPipelineCache)(pipelineCache),
		(*C.size_t)(pDataSize),
		pData))
	return res, res.ToError()
}

func (l *vulkanDriver) VkMergePipelineCaches(device VkDevice, dstCache VkPipelineCache, srcCacheCount Uint32, pSrcCaches *VkPipelineCache) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoMergePipelineCaches(l.funcPtrs.vkMergePipelineCaches,
		(C.VkDevice)(device),
		(C.VkPipelineCache)(dstCache),
		(C.uint32_t)(srcCacheCount),
		(*C.VkPipelineCache)(pSrcCaches)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkCreateGraphicsPipelines(device VkDevice, pipelineCache VkPipelineCache, createInfoCount Uint32, pCreateInfos *VkGraphicsPipelineCreateInfo, pAllocator *VkAllocationCallbacks, pPipelines *VkPipeline) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreateGraphicsPipelines(l.funcPtrs.vkCreateGraphicsPipelines,
		(C.VkDevice)(device),
		(C.VkPipelineCache)(pipelineCache),
		(C.uint32_t)(createInfoCount),
		(*C.VkGraphicsPipelineCreateInfo)(pCreateInfos),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkPipeline)(pPipelines)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkCreateComputePipelines(device VkDevice, pipelineCache VkPipelineCache, createInfoCount Uint32, pCreateInfos *VkComputePipelineCreateInfo, pAllocator *VkAllocationCallbacks, pPipelines *VkPipeline) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreateComputePipelines(l.funcPtrs.vkCreateComputePipelines,
		(C.VkDevice)(device),
		(C.VkPipelineCache)(pipelineCache),
		(C.uint32_t)(createInfoCount),
		(*C.VkComputePipelineCreateInfo)(pCreateInfos),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkPipeline)(pPipelines)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyPipeline(device VkDevice, pipeline VkPipeline, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyPipeline(l.funcPtrs.vkDestroyPipeline,
		(C.VkDevice)(device),
		(C.VkPipeline)(pipeline),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreatePipelineLayout(device VkDevice, pCreateInfo *VkPipelineLayoutCreateInfo, pAllocator *VkAllocationCallbacks, pPipelineLayout *VkPipelineLayout) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreatePipelineLayout(l.funcPtrs.vkCreatePipelineLayout,
		(C.VkDevice)(device),
		(*C.VkPipelineLayoutCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkPipelineLayout)(pPipelineLayout)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyPipelineLayout(device VkDevice, pipelineLayout VkPipelineLayout, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyPipelineLayout(l.funcPtrs.vkDestroyPipelineLayout,
		(C.VkDevice)(device),
		(C.VkPipelineLayout)(pipelineLayout),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateSampler(device VkDevice, pCreateInfo *VkSamplerCreateInfo, pAllocator *VkAllocationCallbacks, pSampler *VkSampler) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreateSampler(l.funcPtrs.vkCreateSampler,
		(C.VkDevice)(device),
		(*C.VkSamplerCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkSampler)(pSampler)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroySampler(device VkDevice, sampler VkSampler, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroySampler(l.funcPtrs.vkDestroySampler,
		(C.VkDevice)(device),
		(C.VkSampler)(sampler),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateDescriptorSetLayout(device VkDevice, pCreateInfo *VkDescriptorSetLayoutCreateInfo, pAllocator *VkAllocationCallbacks, pSetLayout *VkDescriptorSetLayout) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreateDescriptorSetLayout(l.funcPtrs.vkCreateDescriptorSetLayout,
		(C.VkDevice)(device),
		(*C.VkDescriptorSetLayoutCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkDescriptorSetLayout)(pSetLayout)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyDescriptorSetLayout(device VkDevice, descriptorSetLayout VkDescriptorSetLayout, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyDescriptorSetLayout(l.funcPtrs.vkDestroyDescriptorSetLayout,
		(C.VkDevice)(device),
		(C.VkDescriptorSetLayout)(descriptorSetLayout),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateDescriptorPool(device VkDevice, pCreateInfo *VkDescriptorPoolCreateInfo, pAllocator *VkAllocationCallbacks, pDescriptorPool *VkDescriptorPool) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreateDescriptorPool(l.funcPtrs.vkCreateDescriptorPool,
		(C.VkDevice)(device),
		(*C.VkDescriptorPoolCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkDescriptorPool)(pDescriptorPool)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyDescriptorPool(device VkDevice, descriptorPool VkDescriptorPool, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyDescriptorPool(l.funcPtrs.vkDestroyDescriptorPool,
		(C.VkDevice)(device),
		(C.VkDescriptorPool)(descriptorPool),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkResetDescriptorPool(device VkDevice, descriptorPool VkDescriptorPool, flags VkDescriptorPoolResetFlags) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoResetDescriptorPool(l.funcPtrs.vkResetDescriptorPool,
		(C.VkDevice)(device),
		(C.VkDescriptorPool)(descriptorPool),
		(C.VkDescriptorPoolResetFlags)(flags)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkAllocateDescriptorSets(device VkDevice, pAllocateInfo *VkDescriptorSetAllocateInfo, pDescriptorSets *VkDescriptorSet) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoAllocateDescriptorSets(l.funcPtrs.vkAllocateDescriptorSets,
		(C.VkDevice)(device),
		(*C.VkDescriptorSetAllocateInfo)(pAllocateInfo),
		(*C.VkDescriptorSet)(pDescriptorSets)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkFreeDescriptorSets(device VkDevice, descriptorPool VkDescriptorPool, descriptorSetCount Uint32, pDescriptorSets *VkDescriptorSet) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoFreeDescriptorSets(l.funcPtrs.vkFreeDescriptorSets,
		(C.VkDevice)(device),
		(C.VkDescriptorPool)(descriptorPool),
		(C.uint32_t)(descriptorSetCount),
		(*C.VkDescriptorSet)(pDescriptorSets)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkUpdateDescriptorSets(device VkDevice, descriptorWriteCount Uint32, pDescriptorWrites *VkWriteDescriptorSet, descriptorCopyCount Uint32, pDescriptorCopies *VkCopyDescriptorSet) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoUpdateDescriptorSets(l.funcPtrs.vkUpdateDescriptorSets,
		(C.VkDevice)(device),
		(C.uint32_t)(descriptorWriteCount),
		(*C.VkWriteDescriptorSet)(pDescriptorWrites),
		(C.uint32_t)(descriptorCopyCount),
		(*C.VkCopyDescriptorSet)(pDescriptorCopies))
}

func (l *vulkanDriver) VkCreateFramebuffer(device VkDevice, pCreateInfo *VkFramebufferCreateInfo, pAllocator *VkAllocationCallbacks, pFramebuffer *VkFramebuffer) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreateFramebuffer(l.funcPtrs.vkCreateFramebuffer,
		(C.VkDevice)(device),
		(*C.VkFramebufferCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkFramebuffer)(pFramebuffer)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyFramebuffer(device VkDevice, framebuffer VkFramebuffer, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyFramebuffer(l.funcPtrs.vkDestroyFramebuffer,
		(C.VkDevice)(device),
		(C.VkFramebuffer)(framebuffer),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateRenderPass(device VkDevice, pCreateInfo *VkRenderPassCreateInfo, pAllocator *VkAllocationCallbacks, pRenderPass *VkRenderPass) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreateRenderPass(l.funcPtrs.vkCreateRenderPass,
		(C.VkDevice)(device),
		(*C.VkRenderPassCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkRenderPass)(pRenderPass)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyRenderPass(device VkDevice, renderPass VkRenderPass, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyRenderPass(l.funcPtrs.vkDestroyRenderPass,
		(C.VkDevice)(device),
		(C.VkRenderPass)(renderPass),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkGetRenderAreaGranularity(device VkDevice, renderPass VkRenderPass, pGranularity *VkExtent2D) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoGetRenderAreaGranularity(l.funcPtrs.vkGetRenderAreaGranularity,
		(C.VkDevice)(device),
		(C.VkRenderPass)(renderPass),
		(*C.VkExtent2D)(pGranularity))
}

func (l *vulkanDriver) VkCreateCommandPool(device VkDevice, pCreateInfo *VkCommandPoolCreateInfo, pAllocator *VkAllocationCallbacks, pCommandPool *VkCommandPool) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoCreateCommandPool(l.funcPtrs.vkCreateCommandPool,
		(C.VkDevice)(device),
		(*C.VkCommandPoolCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkCommandPool)(pCommandPool)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyCommandPool(device VkDevice, commandPool VkCommandPool, pAllocator *VkAllocationCallbacks) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyCommandPool(l.funcPtrs.vkDestroyCommandPool,
		(C.VkDevice)(device),
		(C.VkCommandPool)(commandPool),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkResetCommandPool(device VkDevice, commandPool VkCommandPool, flags VkCommandPoolResetFlags) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoResetCommandPool(l.funcPtrs.vkResetCommandPool,
		(C.VkDevice)(device),
		(C.VkCommandPool)(commandPool),
		(C.VkCommandPoolResetFlags)(flags)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkAllocateCommandBuffers(device VkDevice, pAllocateInfo *VkCommandBufferAllocateInfo, pCommandBuffers *VkCommandBuffer) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoAllocateCommandBuffers(l.funcPtrs.vkAllocateCommandBuffers,
		(C.VkDevice)(device),
		(*C.VkCommandBufferAllocateInfo)(pAllocateInfo),
		(*C.VkCommandBuffer)(pCommandBuffers)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkFreeCommandBuffers(device VkDevice, commandPool VkCommandPool, commandBufferCount Uint32, pCommandBuffers *VkCommandBuffer) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoFreeCommandBuffers(l.funcPtrs.vkFreeCommandBuffers,
		(C.VkDevice)(device),
		(C.VkCommandPool)(commandPool),
		(C.uint32_t)(commandBufferCount),
		(*C.VkCommandBuffer)(pCommandBuffers))
}

func (l *vulkanDriver) VkBeginCommandBuffer(commandBuffer VkCommandBuffer, pBeginInfo *VkCommandBufferBeginInfo) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoBeginCommandBuffer(l.funcPtrs.vkBeginCommandBuffer,
		(C.VkCommandBuffer)(commandBuffer),
		(*C.VkCommandBufferBeginInfo)(pBeginInfo)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkEndCommandBuffer(commandBuffer VkCommandBuffer) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoEndCommandBuffer(l.funcPtrs.vkEndCommandBuffer,
		(C.VkCommandBuffer)(commandBuffer)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkResetCommandBuffer(commandBuffer VkCommandBuffer, flags VkCommandBufferResetFlags) (VkResult, error) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	res := VkResult(C.cgoResetCommandBuffer(l.funcPtrs.vkResetCommandBuffer,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkCommandBufferResetFlags)(flags)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkCmdBindPipeline(commandBuffer VkCommandBuffer, pipelineBindPoint VkPipelineBindPoint, pipeline VkPipeline) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdBindPipeline(l.funcPtrs.vkCmdBindPipeline,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkPipelineBindPoint)(pipelineBindPoint),
		(C.VkPipeline)(pipeline))
}

func (l *vulkanDriver) VkCmdSetViewport(commandBuffer VkCommandBuffer, firstViewport Uint32, viewportCount Uint32, pViewports *VkViewport) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetViewport(l.funcPtrs.vkCmdSetViewport,
		(C.VkCommandBuffer)(commandBuffer),
		(C.uint32_t)(firstViewport),
		(C.uint32_t)(viewportCount),
		(*C.VkViewport)(pViewports))
}

func (l *vulkanDriver) VkCmdSetScissor(commandBuffer VkCommandBuffer, firstScissor Uint32, scissorCount Uint32, pScissors *VkRect2D) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetScissor(l.funcPtrs.vkCmdSetScissor,
		(C.VkCommandBuffer)(commandBuffer),
		C.uint32_t(firstScissor),
		C.uint32_t(scissorCount),
		(*C.VkRect2D)(pScissors))
}

func (l *vulkanDriver) VkCmdSetLineWidth(commandBuffer VkCommandBuffer, lineWidth Float) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}
	C.cgoCmdSetLineWidth(l.funcPtrs.vkCmdSetLineWidth,
		(C.VkCommandBuffer)(commandBuffer),
		(C.float)(lineWidth))
}

func (l *vulkanDriver) VkCmdSetDepthBias(commandBuffer VkCommandBuffer, depthBiasConstantFactor Float, depthBiasClamp Float, depthBiasSlopeFactor Float) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetDepthBias(l.funcPtrs.vkCmdSetDepthBias,
		(C.VkCommandBuffer)(commandBuffer),
		(C.float)(depthBiasConstantFactor),
		(C.float)(depthBiasClamp),
		(C.float)(depthBiasSlopeFactor))
}

func (l *vulkanDriver) VkCmdSetBlendConstants(commandBuffer VkCommandBuffer, blendConstants *Float) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetBlendConstants(l.funcPtrs.vkCmdSetBlendConstants,
		(C.VkCommandBuffer)(commandBuffer),
		(*C.float)(blendConstants),
	)
}

func (l *vulkanDriver) VkCmdSetDepthBounds(commandBuffer VkCommandBuffer, minDepthBounds Float, maxDepthBounds Float) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetDepthBounds(l.funcPtrs.vkCmdSetDepthBounds,
		(C.VkCommandBuffer)(commandBuffer),
		(C.float)(minDepthBounds),
		(C.float)(maxDepthBounds))
}

func (l *vulkanDriver) VkCmdSetStencilCompareMask(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, compareMask Uint32) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetStencilCompareMask(l.funcPtrs.vkCmdSetStencilCompareMask,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkStencilFaceFlags)(faceMask),
		(C.uint32_t)(compareMask))
}

func (l *vulkanDriver) VkCmdSetStencilWriteMask(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, writeMask Uint32) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetStencilWriteMask(l.funcPtrs.vkCmdSetStencilWriteMask,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkStencilFaceFlags)(faceMask),
		(C.uint32_t)(writeMask))
}

func (l *vulkanDriver) VkCmdSetStencilReference(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, reference Uint32) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetStencilReference(l.funcPtrs.vkCmdSetStencilReference,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkStencilFaceFlags)(faceMask),
		(C.uint32_t)(reference))
}

func (l *vulkanDriver) VkCmdBindDescriptorSets(commandBuffer VkCommandBuffer, pipelineBindPoint VkPipelineBindPoint, layout VkPipelineLayout, firstSet Uint32, descriptorSetCount Uint32, pDescriptorSets *VkDescriptorSet, dynamicOffsetCount Uint32, pDynamicOffsets *Uint32) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdBindDescriptorSets(l.funcPtrs.vkCmdBindDescriptorSets,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkPipelineBindPoint)(pipelineBindPoint),
		(C.VkPipelineLayout)(layout),
		(C.uint32_t)(firstSet),
		(C.uint32_t)(descriptorSetCount),
		(*C.VkDescriptorSet)(pDescriptorSets),
		(C.uint32_t)(dynamicOffsetCount),
		(*C.uint32_t)(pDynamicOffsets))
}

func (l *vulkanDriver) VkCmdBindIndexBuffer(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, indexType VkIndexType) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdBindIndexBuffer(l.funcPtrs.vkCmdBindIndexBuffer,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkBuffer)(buffer),
		(C.VkDeviceSize)(offset),
		(C.VkIndexType)(indexType))
}

func (l *vulkanDriver) VkCmdBindVertexBuffers(commandBuffer VkCommandBuffer, firstBinding Uint32, bindingCount Uint32, pBuffers *VkBuffer, pOffsets *VkDeviceSize) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdBindVertexBuffers(l.funcPtrs.vkCmdBindVertexBuffers,
		(C.VkCommandBuffer)(commandBuffer),
		C.uint32_t(firstBinding),
		C.uint32_t(bindingCount),
		(*C.VkBuffer)(pBuffers),
		(*C.VkDeviceSize)(pOffsets))
}

func (l *vulkanDriver) VkCmdDraw(commandBuffer VkCommandBuffer, vertexCount Uint32, instanceCount Uint32, firstVertex Uint32, firstInstance Uint32) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdDraw(l.funcPtrs.vkCmdDraw,
		(C.VkCommandBuffer)(commandBuffer),
		(C.uint32_t)(vertexCount),
		(C.uint32_t)(instanceCount),
		(C.uint32_t)(firstVertex),
		(C.uint32_t)(firstInstance))
}

func (l *vulkanDriver) VkCmdDrawIndexed(commandBuffer VkCommandBuffer, indexCount Uint32, instanceCount Uint32, firstIndex Uint32, vertexOffset Int32, firstInstance Uint32) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdDrawIndexed(l.funcPtrs.vkCmdDrawIndexed,
		(C.VkCommandBuffer)(commandBuffer),
		(C.uint32_t)(indexCount),
		(C.uint32_t)(instanceCount),
		(C.uint32_t)(firstIndex),
		(C.int32_t)(vertexOffset),
		(C.uint32_t)(firstInstance))
}

func (l *vulkanDriver) VkCmdDrawIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, drawCount Uint32, stride Uint32) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdDrawIndirect(l.funcPtrs.vkCmdDrawIndirect,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkBuffer)(buffer),
		(C.VkDeviceSize)(offset),
		(C.uint32_t)(drawCount),
		(C.uint32_t)(stride))
}

func (l *vulkanDriver) VkCmdDrawIndexedIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, drawCount Uint32, stride Uint32) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdDrawIndexedIndirect(l.funcPtrs.vkCmdDrawIndexedIndirect,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkBuffer)(buffer),
		(C.VkDeviceSize)(offset),
		(C.uint32_t)(drawCount),
		(C.uint32_t)(stride))
}

func (l *vulkanDriver) VkCmdDispatch(commandBuffer VkCommandBuffer, groupCountX Uint32, groupCountY Uint32, groupCountZ Uint32) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdDispatch(l.funcPtrs.vkCmdDispatch,
		(C.VkCommandBuffer)(commandBuffer),
		(C.uint32_t)(groupCountX),
		(C.uint32_t)(groupCountY),
		(C.uint32_t)(groupCountZ))
}

func (l *vulkanDriver) VkCmdDispatchIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdDispatchIndirect(l.funcPtrs.vkCmdDispatchIndirect,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkBuffer)(buffer),
		(C.VkDeviceSize)(offset))
}

func (l *vulkanDriver) VkCmdCopyBuffer(commandBuffer VkCommandBuffer, srcBuffer VkBuffer, dstBuffer VkBuffer, regionCount Uint32, pRegions *VkBufferCopy) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdCopyBuffer(l.funcPtrs.vkCmdCopyBuffer,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkBuffer)(srcBuffer),
		(C.VkBuffer)(dstBuffer),
		(C.uint32_t)(regionCount),
		(*C.VkBufferCopy)(pRegions))
}

func (l *vulkanDriver) VkCmdCopyImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkImageCopy) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdCopyImage(l.funcPtrs.vkCmdCopyImage,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkImage)(srcImage),
		(C.VkImageLayout)(srcImageLayout),
		(C.VkImage)(dstImage),
		(C.VkImageLayout)(dstImageLayout),
		(C.uint32_t)(regionCount),
		(*C.VkImageCopy)(pRegions))
}

func (l *vulkanDriver) VkCmdBlitImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkImageBlit, filter VkFilter) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdBlitImage(l.funcPtrs.vkCmdBlitImage,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkImage)(srcImage),
		(C.VkImageLayout)(srcImageLayout),
		(C.VkImage)(dstImage),
		(C.VkImageLayout)(dstImageLayout),
		(C.uint32_t)(regionCount),
		(*C.VkImageBlit)(pRegions),
		(C.VkFilter)(filter),
	)
}

func (l *vulkanDriver) VkCmdCopyBufferToImage(commandBuffer VkCommandBuffer, srcBuffer VkBuffer, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkBufferImageCopy) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdCopyBufferToImage(l.funcPtrs.vkCmdCopyBufferToImage,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkBuffer)(srcBuffer),
		(C.VkImage)(dstImage),
		(C.VkImageLayout)(dstImageLayout),
		(C.uint32_t)(regionCount),
		(*C.VkBufferImageCopy)(pRegions))
}

func (l *vulkanDriver) VkCmdCopyImageToBuffer(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstBuffer VkBuffer, regionCount Uint32, pRegions *VkBufferImageCopy) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdCopyImageToBuffer(l.funcPtrs.vkCmdCopyImageToBuffer,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkImage)(srcImage),
		(C.VkImageLayout)(srcImageLayout),
		(C.VkBuffer)(dstBuffer),
		(C.uint32_t)(regionCount),
		(*C.VkBufferImageCopy)(pRegions))
}

func (l *vulkanDriver) VkCmdUpdateBuffer(commandBuffer VkCommandBuffer, dstBuffer VkBuffer, dstOffset VkDeviceSize, dataSize VkDeviceSize, pData unsafe.Pointer) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdUpdateBuffer(l.funcPtrs.vkCmdUpdateBuffer,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkBuffer)(dstBuffer),
		(C.VkDeviceSize)(dstOffset),
		(C.VkDeviceSize)(dataSize),
		pData)
}

func (l *vulkanDriver) VkCmdFillBuffer(commandBuffer VkCommandBuffer, dstBuffer VkBuffer, dstOffset VkDeviceSize, size VkDeviceSize, data Uint32) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdFillBuffer(l.funcPtrs.vkCmdFillBuffer,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkBuffer)(dstBuffer),
		(C.VkDeviceSize)(dstOffset),
		(C.VkDeviceSize)(size),
		(C.uint32_t)(data))
}

func (l *vulkanDriver) VkCmdClearColorImage(commandBuffer VkCommandBuffer, image VkImage, imageLayout VkImageLayout, pColor *VkClearColorValue, rangeCount Uint32, pRanges *VkImageSubresourceRange) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdClearColorImage(l.funcPtrs.vkCmdClearColorImage,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkImage)(image),
		(C.VkImageLayout)(imageLayout),
		(*C.VkClearColorValue)(pColor),
		(C.uint32_t)(rangeCount),
		(*C.VkImageSubresourceRange)(pRanges))
}

func (l *vulkanDriver) VkCmdClearDepthStencilImage(commandBuffer VkCommandBuffer, image VkImage, imageLayout VkImageLayout, pDepthStencil *VkClearDepthStencilValue, rangeCount Uint32, pRanges *VkImageSubresourceRange) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdClearDepthStencilImage(l.funcPtrs.vkCmdClearDepthStencilImage,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkImage)(image),
		(C.VkImageLayout)(imageLayout),
		(*C.VkClearDepthStencilValue)(pDepthStencil),
		(C.uint32_t)(rangeCount),
		(*C.VkImageSubresourceRange)(pRanges))
}

func (l *vulkanDriver) VkCmdClearAttachments(commandBuffer VkCommandBuffer, attachmentCount Uint32, pAttachments *VkClearAttachment, rectCount Uint32, pRects *VkClearRect) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdClearAttachments(l.funcPtrs.vkCmdClearAttachments,
		(C.VkCommandBuffer)(commandBuffer),
		(C.uint32_t)(attachmentCount),
		(*C.VkClearAttachment)(pAttachments),
		(C.uint32_t)(rectCount),
		(*C.VkClearRect)(pRects))
}

func (l *vulkanDriver) VkCmdResolveImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkImageResolve) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdResolveImage(l.funcPtrs.vkCmdResolveImage,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkImage)(srcImage),
		(C.VkImageLayout)(srcImageLayout),
		(C.VkImage)(dstImage),
		(C.VkImageLayout)(dstImageLayout),
		(C.uint32_t)(regionCount),
		(*C.VkImageResolve)(pRegions))
}

func (l *vulkanDriver) VkCmdSetEvent(commandBuffer VkCommandBuffer, event VkEvent, stageMask VkPipelineStageFlags) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetEvent(l.funcPtrs.vkCmdSetEvent,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkEvent)(event),
		(C.VkPipelineStageFlags)(stageMask))
}

func (l *vulkanDriver) VkCmdResetEvent(commandBuffer VkCommandBuffer, event VkEvent, stageMask VkPipelineStageFlags) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdResetEvent(l.funcPtrs.vkCmdResetEvent,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkEvent)(event),
		(C.VkPipelineStageFlags)(stageMask))
}

func (l *vulkanDriver) VkCmdWaitEvents(commandBuffer VkCommandBuffer, eventCount Uint32, pEvents *VkEvent, srcStageMask VkPipelineStageFlags, dstStageMask VkPipelineStageFlags, memoryBarrierCount Uint32, pMemoryBarriers *VkMemoryBarrier, bufferMemoryBarrierCount Uint32, pBufferMemoryBarriers *VkBufferMemoryBarrier, imageMemoryBarrierCount Uint32, pImageMemoryBarriers *VkImageMemoryBarrier) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdWaitEvents(l.funcPtrs.vkCmdWaitEvents,
		(C.VkCommandBuffer)(commandBuffer),
		(C.uint32_t)(eventCount),
		(*C.VkEvent)(pEvents),
		(C.VkPipelineStageFlags)(srcStageMask),
		(C.VkPipelineStageFlags)(dstStageMask),
		(C.uint32_t)(memoryBarrierCount),
		(*C.VkMemoryBarrier)(pMemoryBarriers),
		(C.uint32_t)(bufferMemoryBarrierCount),
		(*C.VkBufferMemoryBarrier)(pBufferMemoryBarriers),
		(C.uint32_t)(imageMemoryBarrierCount),
		(*C.VkImageMemoryBarrier)(pImageMemoryBarriers))
}

func (l *vulkanDriver) VkCmdPipelineBarrier(commandBuffer VkCommandBuffer, srcStageMask VkPipelineStageFlags, dstStageMask VkPipelineStageFlags, dependencyFlags VkDependencyFlags, memoryBarrierCount Uint32, pMemoryBarriers *VkMemoryBarrier, bufferMemoryBarrierCount Uint32, pBufferMemoryBarriers *VkBufferMemoryBarrier, imageMemoryBarrierCount Uint32, pImageMemoryBarriers *VkImageMemoryBarrier) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdPipelineBarrier(l.funcPtrs.vkCmdPipelineBarrier,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkPipelineStageFlags)(srcStageMask),
		(C.VkPipelineStageFlags)(dstStageMask),
		(C.VkDependencyFlags)(dependencyFlags),
		(C.uint32_t)(memoryBarrierCount),
		(*C.VkMemoryBarrier)(pMemoryBarriers),
		(C.uint32_t)(bufferMemoryBarrierCount),
		(*C.VkBufferMemoryBarrier)(pBufferMemoryBarriers),
		(C.uint32_t)(imageMemoryBarrierCount),
		(*C.VkImageMemoryBarrier)(pImageMemoryBarriers))
}

func (l *vulkanDriver) VkCmdBeginQuery(commandBuffer VkCommandBuffer, queryPool VkQueryPool, query Uint32, flags VkQueryControlFlags) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdBeginQuery(l.funcPtrs.vkCmdBeginQuery,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkQueryPool)(queryPool),
		(C.uint32_t)(query),
		(C.VkQueryControlFlags)(flags))
}

func (l *vulkanDriver) VkCmdEndQuery(commandBuffer VkCommandBuffer, queryPool VkQueryPool, query Uint32) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdEndQuery(l.funcPtrs.vkCmdEndQuery,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkQueryPool)(queryPool),
		(C.uint32_t)(query))
}

func (l *vulkanDriver) VkCmdResetQueryPool(commandBuffer VkCommandBuffer, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdResetQueryPool(l.funcPtrs.vkCmdResetQueryPool,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkQueryPool)(queryPool),
		(C.uint32_t)(firstQuery),
		(C.uint32_t)(queryCount))
}

func (l *vulkanDriver) VkCmdWriteTimestamp(commandBuffer VkCommandBuffer, pipelineStage VkPipelineStageFlags, queryPool VkQueryPool, query Uint32) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdWriteTimestamp(l.funcPtrs.vkCmdWriteTimestamp,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkPipelineStageFlagBits)(pipelineStage),
		(C.VkQueryPool)(queryPool),
		(C.uint32_t)(query))
}

func (l *vulkanDriver) VkCmdCopyQueryPoolResults(commandBuffer VkCommandBuffer, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32, dstBuffer VkBuffer, dstOffset VkDeviceSize, stride VkDeviceSize, flags VkQueryResultFlags) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdCopyQueryPoolResults(l.funcPtrs.vkCmdCopyQueryPoolResults,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkQueryPool)(queryPool),
		(C.uint32_t)(firstQuery),
		(C.uint32_t)(queryCount),
		(C.VkBuffer)(dstBuffer),
		(C.VkDeviceSize)(dstOffset),
		(C.VkDeviceSize)(stride),
		(C.VkQueryResultFlags)(flags))
}

func (l *vulkanDriver) VkCmdPushConstants(commandBuffer VkCommandBuffer, layout VkPipelineLayout, stageFlags VkShaderStageFlags, offset Uint32, size Uint32, pValues unsafe.Pointer) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdPushConstants(l.funcPtrs.vkCmdPushConstants,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkPipelineLayout)(layout),
		(C.VkShaderStageFlags)(stageFlags),
		(C.uint32_t)(offset),
		(C.uint32_t)(size),
		pValues)
}

func (l *vulkanDriver) VkCmdBeginRenderPass(commandBuffer VkCommandBuffer, pRenderPassBegin *VkRenderPassBeginInfo, contents VkSubpassContents) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdBeginRenderPass(l.funcPtrs.vkCmdBeginRenderPass,
		(C.VkCommandBuffer)(commandBuffer),
		(*C.VkRenderPassBeginInfo)(pRenderPassBegin),
		(C.VkSubpassContents)(contents))
}

func (l *vulkanDriver) VkCmdNextSubpass(commandBuffer VkCommandBuffer, contents VkSubpassContents) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdNextSubpass(l.funcPtrs.vkCmdNextSubpass,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkSubpassContents)(contents))
}

func (l *vulkanDriver) VkCmdEndRenderPass(commandBuffer VkCommandBuffer) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdEndRenderPass(l.funcPtrs.vkCmdEndRenderPass,
		(C.VkCommandBuffer)(commandBuffer))
}

func (l *vulkanDriver) VkCmdExecuteCommands(commandBuffer VkCommandBuffer, commandBufferCount Uint32, pCommandBuffers *VkCommandBuffer) {
	if l.device == nil {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdExecuteCommands(l.funcPtrs.vkCmdExecuteCommands,
		(C.VkCommandBuffer)(commandBuffer),
		(C.uint32_t)(commandBufferCount),
		(*C.VkCommandBuffer)(pCommandBuffers))
}

func (l *vulkanDriver) VkEnumerateInstanceVersion(pApiVersion *Uint32) (VkResult, error) {
	if l.funcPtrs.vkEnumerateInstanceVersion == nil {
		panic("attempted to call method 'vkEnumerateInstanceVersion' which is not present on this driver")
	}

	res := VkResult(C.cgoEnumerateInstanceVersion(l.funcPtrs.vkEnumerateInstanceVersion,
		(*C.uint32_t)(pApiVersion)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkEnumeratePhysicalDeviceGroups(instance VkInstance, pPhysicalDeviceGroupCount *Uint32, pPhysicalDeviceGroupProperties *VkPhysicalDeviceGroupProperties) (VkResult, error) {
	if l.funcPtrs.vkEnumeratePhysicalDeviceGroups == nil {
		panic("attempted to call method 'vkEnumeratePhysicalDeviceGroups' which is not present on this driver")
	}

	res := VkResult(C.cgoEnumeratePhysicalDeviceGroups(l.funcPtrs.vkEnumeratePhysicalDeviceGroups,
		C.VkInstance(instance),
		(*C.uint32_t)(pPhysicalDeviceGroupCount),
		(*C.VkPhysicalDeviceGroupProperties)(pPhysicalDeviceGroupProperties)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkGetPhysicalDeviceFeatures2(physicalDevice VkPhysicalDevice, pFeatures *VkPhysicalDeviceFeatures2) {
	if l.funcPtrs.vkGetPhysicalDeviceFeatures2 == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceFeatures2' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceFeatures2(l.funcPtrs.vkGetPhysicalDeviceFeatures2,
		C.VkPhysicalDevice(physicalDevice),
		(*C.VkPhysicalDeviceFeatures2)(pFeatures))
}

func (l *vulkanDriver) VkGetPhysicalDeviceProperties2(physicalDevice VkPhysicalDevice, pProperties *VkPhysicalDeviceProperties2) {
	if l.funcPtrs.vkGetPhysicalDeviceProperties2 == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceProperties2' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceProperties2(l.funcPtrs.vkGetPhysicalDeviceProperties2,
		C.VkPhysicalDevice(physicalDevice),
		(*C.VkPhysicalDeviceProperties2)(pProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceFormatProperties2(physicalDevice VkPhysicalDevice, format VkFormat, pFormatProperties *VkFormatProperties2) {
	if l.funcPtrs.vkGetPhysicalDeviceFormatProperties2 == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceFormatProperties2' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceFormatProperties2(l.funcPtrs.vkGetPhysicalDeviceFormatProperties2,
		C.VkPhysicalDevice(physicalDevice),
		C.VkFormat(format),
		(*C.VkFormatProperties2)(pFormatProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceImageFormatProperties2(physicalDevice VkPhysicalDevice, pImageFormatInfo *VkPhysicalDeviceImageFormatInfo2, pImageFormatProperties *VkImageFormatProperties2) (VkResult, error) {
	if l.funcPtrs.vkGetPhysicalDeviceImageFormatProperties2 == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceImageFormatProperties2' which is not present on this driver")
	}

	res := VkResult(C.cgoGetPhysicalDeviceImageFormatProperties2(l.funcPtrs.vkGetPhysicalDeviceImageFormatProperties2,
		C.VkPhysicalDevice(physicalDevice),
		(*C.VkPhysicalDeviceImageFormatInfo2)(pImageFormatInfo),
		(*C.VkImageFormatProperties2)(pImageFormatProperties)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkGetPhysicalDeviceQueueFamilyProperties2(physicalDevice VkPhysicalDevice, pQueueFamilyPropertyCount *Uint32, pQueueFamilyProperties *VkQueueFamilyProperties2) {
	if l.funcPtrs.vkGetPhysicalDeviceQueueFamilyProperties2 == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceQueueFamilyProperties2' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceQueueFamilyProperties2(l.funcPtrs.vkGetPhysicalDeviceQueueFamilyProperties2,
		C.VkPhysicalDevice(physicalDevice),
		(*C.uint32_t)(pQueueFamilyPropertyCount),
		(*C.VkQueueFamilyProperties2)(pQueueFamilyProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceMemoryProperties2(physicalDevice VkPhysicalDevice, pMemoryProperties *VkPhysicalDeviceMemoryProperties2) {
	if l.funcPtrs.vkGetPhysicalDeviceMemoryProperties2 == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceMemoryProperties2' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceMemoryProperties2(l.funcPtrs.vkGetPhysicalDeviceMemoryProperties2,
		C.VkPhysicalDevice(physicalDevice),
		(*C.VkPhysicalDeviceMemoryProperties2)(pMemoryProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceSparseImageFormatProperties2(physicalDevice VkPhysicalDevice, pFormatInfo *VkPhysicalDeviceSparseImageFormatInfo2, pPropertyCount *Uint32, pProperties *VkSparseImageFormatProperties2) {
	if l.funcPtrs.vkGetPhysicalDeviceSparseImageFormatProperties2 == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceSparseImageFormatProperties2' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceSparseImageFormatProperties2(l.funcPtrs.vkGetPhysicalDeviceSparseImageFormatProperties2,
		C.VkPhysicalDevice(physicalDevice),
		(*C.VkPhysicalDeviceSparseImageFormatInfo2)(pFormatInfo),
		(*C.uint32_t)(pPropertyCount),
		(*C.VkSparseImageFormatProperties2)(pProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceExternalBufferProperties(physicalDevice VkPhysicalDevice, pExternalBufferInfo *VkPhysicalDeviceExternalBufferInfo, pExternalBufferProperties *VkExternalBufferProperties) {
	if l.funcPtrs.vkGetPhysicalDeviceExternalBufferProperties == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceExternalBufferProperties' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceExternalBufferProperties(l.funcPtrs.vkGetPhysicalDeviceExternalBufferProperties,
		C.VkPhysicalDevice(physicalDevice),
		(*C.VkPhysicalDeviceExternalBufferInfo)(pExternalBufferInfo),
		(*C.VkExternalBufferProperties)(pExternalBufferProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceExternalFenceProperties(physicalDevice VkPhysicalDevice, pExternalFenceInfo *VkPhysicalDeviceExternalFenceInfo, pExternalFenceProperties *VkExternalFenceProperties) {
	if l.funcPtrs.vkGetPhysicalDeviceExternalFenceProperties == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceExternalFenceProperties' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceExternalFenceProperties(l.funcPtrs.vkGetPhysicalDeviceExternalFenceProperties,
		C.VkPhysicalDevice(physicalDevice),
		(*C.VkPhysicalDeviceExternalFenceInfo)(pExternalFenceInfo),
		(*C.VkExternalFenceProperties)(pExternalFenceProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceExternalSemaphoreProperties(physicalDevice VkPhysicalDevice, pExternalSemaphoreInfo *VkPhysicalDeviceExternalSemaphoreInfo, pExternalSemaphoreProperties *VkExternalSemaphoreProperties) {
	if l.funcPtrs.vkGetPhysicalDeviceExternalSemaphoreProperties == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceExternalSemaphoreProperties' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceExternalSemaphoreProperties(l.funcPtrs.vkGetPhysicalDeviceExternalSemaphoreProperties,
		C.VkPhysicalDevice(physicalDevice),
		(*C.VkPhysicalDeviceExternalSemaphoreInfo)(pExternalSemaphoreInfo),
		(*C.VkExternalSemaphoreProperties)(pExternalSemaphoreProperties))
}

func (l *vulkanDriver) VkBindBufferMemory2(device VkDevice, bindInfoCount Uint32, pBindInfos *VkBindBufferMemoryInfo) (VkResult, error) {
	if l.funcPtrs.vkBindBufferMemory2 == nil {
		panic("attempted to call method 'vkBindBufferMemory2' which is not present on this driver")
	}

	res := VkResult(C.cgoBindBufferMemory2(l.funcPtrs.vkBindBufferMemory2,
		C.VkDevice(device),
		C.uint32_t(bindInfoCount),
		(*C.VkBindBufferMemoryInfo)(pBindInfos)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkBindImageMemory2(device VkDevice, bindInfoCount Uint32, pBindInfos *VkBindImageMemoryInfo) (VkResult, error) {
	if l.funcPtrs.vkBindImageMemory2 == nil {
		panic("attempted to call method 'vkBindImageMemory2' which is not present on this driver")
	}

	res := VkResult(C.cgoBindImageMemory2(l.funcPtrs.vkBindImageMemory2,
		C.VkDevice(device),
		C.uint32_t(bindInfoCount),
		(*C.VkBindImageMemoryInfo)(pBindInfos)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkGetDeviceGroupPeerMemoryFeatures(device VkDevice, heapIndex Uint32, localDeviceIndex Uint32, remoteDeviceIndex Uint32, pPeerMemoryFeatures *VkPeerMemoryFeatureFlags) {
	if l.funcPtrs.vkGetDeviceGroupPeerMemoryFeatures == nil {
		panic("attempted to call method 'vkGetDeviceGroupPeerMemoryFeatures' which is not present on this driver")
	}

	C.cgoGetDeviceGroupPeerMemoryFeatures(l.funcPtrs.vkGetDeviceGroupPeerMemoryFeatures,
		C.VkDevice(device),
		C.uint32_t(heapIndex),
		C.uint32_t(localDeviceIndex),
		C.uint32_t(remoteDeviceIndex),
		(*C.VkPeerMemoryFeatureFlags)(pPeerMemoryFeatures))
}

func (l *vulkanDriver) VkCmdSetDeviceMask(commandBuffer VkCommandBuffer, deviceMask Uint32) {
	if l.funcPtrs.vkCmdSetDeviceMask == nil {
		panic("attempted to call method 'vkCmdSetDeviceMask' which is not present on this driver")
	}

	C.cgoCmdSetDeviceMask(l.funcPtrs.vkCmdSetDeviceMask,
		C.VkCommandBuffer(commandBuffer),
		C.uint32_t(deviceMask))
}

func (l *vulkanDriver) VkCmdDispatchBase(commandBuffer VkCommandBuffer, baseGroupX Uint32, baseGroupY Uint32, baseGroupZ Uint32, groupCountX Uint32, groupCountY Uint32, groupCountZ Uint32) {
	if l.funcPtrs.vkCmdDispatchBase == nil {
		panic("attempted to call method 'vkCmdDispatchBase' which is not present on this driver")
	}

	C.cgoCmdDispatchBase(l.funcPtrs.vkCmdDispatchBase,
		C.VkCommandBuffer(commandBuffer),
		C.uint32_t(baseGroupX),
		C.uint32_t(baseGroupY),
		C.uint32_t(baseGroupZ),
		C.uint32_t(groupCountX),
		C.uint32_t(groupCountY),
		C.uint32_t(groupCountZ))
}

func (l *vulkanDriver) VkGetImageMemoryRequirements2(device VkDevice, pInfo *VkImageMemoryRequirementsInfo2, pMemoryRequirements *VkMemoryRequirements2) {
	if l.funcPtrs.vkGetImageMemoryRequirements2 == nil {
		panic("attempted to call method 'vkGetImageMemoryRequirements2' which is not present on this driver")
	}

	C.cgoGetImageMemoryRequirements2(l.funcPtrs.vkGetImageMemoryRequirements2,
		C.VkDevice(device),
		(*C.VkImageMemoryRequirementsInfo2)(pInfo),
		(*C.VkMemoryRequirements2)(pMemoryRequirements))
}

func (l *vulkanDriver) VkGetBufferMemoryRequirements2(device VkDevice, pInfo *VkBufferMemoryRequirementsInfo2, pMemoryRequirements *VkMemoryRequirements2) {
	if l.funcPtrs.vkGetBufferMemoryRequirements2 == nil {
		panic("attempted to call method 'vkGetBufferMemoryRequirements2' which is not present on this driver")
	}

	C.cgoGetBufferMemoryRequirements2(l.funcPtrs.vkGetBufferMemoryRequirements2,
		C.VkDevice(device),
		(*C.VkBufferMemoryRequirementsInfo2)(pInfo),
		(*C.VkMemoryRequirements2)(pMemoryRequirements))
}

func (l *vulkanDriver) VkGetImageSparseMemoryRequirements2(device VkDevice, pInfo *VkImageSparseMemoryRequirementsInfo2, pSparseMemoryRequirementCount *Uint32, pSparseMemoryRequirements *VkSparseImageMemoryRequirements2) {
	if l.funcPtrs.vkGetImageSparseMemoryRequirements2 == nil {
		panic("attempted to call method 'vkGetImageSparseMemoryRequirements2' which is not present on this driver")
	}

	C.cgoGetImageSparseMemoryRequirements2(l.funcPtrs.vkGetImageSparseMemoryRequirements2,
		C.VkDevice(device),
		(*C.VkImageSparseMemoryRequirementsInfo2)(pInfo),
		(*C.uint32_t)(pSparseMemoryRequirementCount),
		(*C.VkSparseImageMemoryRequirements2)(pSparseMemoryRequirements))
}

func (l *vulkanDriver) VkTrimCommandPool(device VkDevice, commandPool VkCommandPool, flags VkCommandPoolTrimFlags) {
	if l.funcPtrs.vkTrimCommandPool == nil {
		panic("attempted to call method 'vkTrimCommandPool' which is not present on this driver")
	}

	C.cgoTrimCommandPool(l.funcPtrs.vkTrimCommandPool,
		C.VkDevice(device),
		C.VkCommandPool(commandPool),
		C.VkCommandPoolTrimFlags(flags))
}

func (l *vulkanDriver) VkGetDeviceQueue2(device VkDevice, pQueueInfo *VkDeviceQueueInfo2, pQueue *VkQueue) {
	if l.funcPtrs.vkGetDeviceQueue2 == nil {
		panic("attempted to call method 'vkGetDeviceQueue2' which is not present on this driver")
	}

	C.cgoGetDeviceQueue2(l.funcPtrs.vkGetDeviceQueue2,
		C.VkDevice(device),
		(*C.VkDeviceQueueInfo2)(pQueueInfo),
		(*C.VkQueue)(pQueue))
}

func (l *vulkanDriver) VkCreateSamplerYcbcrConversion(device VkDevice, pCreateInfo *VkSamplerYcbcrConversionCreateInfo, pAllocator *VkAllocationCallbacks, pYcbcrConversion *VkSamplerYcbcrConversion) (VkResult, error) {
	if l.funcPtrs.vkCreateSamplerYcbcrConversion == nil {
		panic("attempted to call method 'vkCreateSamplerYcbcrConversion' which is not present on this driver")
	}

	res := VkResult(C.cgoCreateSamplerYcbcrConversion(l.funcPtrs.vkCreateSamplerYcbcrConversion,
		C.VkDevice(device),
		(*C.VkSamplerYcbcrConversionCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkSamplerYcbcrConversion)(pYcbcrConversion)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroySamplerYcbcrConversion(device VkDevice, ycbcrConversion VkSamplerYcbcrConversion, pAllocator *VkAllocationCallbacks) {
	if l.funcPtrs.vkDestroySamplerYcbcrConversion == nil {
		panic("attempted to call method 'vkDestroySamplerYcbcrConversion' which is not present on this driver")
	}

	C.cgoDestroySamplerYcbcrConversion(l.funcPtrs.vkDestroySamplerYcbcrConversion,
		C.VkDevice(device),
		C.VkSamplerYcbcrConversion(ycbcrConversion),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateDescriptorUpdateTemplate(device VkDevice, pCreateInfo *VkDescriptorUpdateTemplateCreateInfo, pAllocator *VkAllocationCallbacks, pDescriptorUpdateTemplate *VkDescriptorUpdateTemplate) (VkResult, error) {
	if l.funcPtrs.vkCreateDescriptorUpdateTemplate == nil {
		panic("attempted to call method 'vkCreateDescriptorUpdateTemplate' which is not present on this driver")
	}

	res := VkResult(C.cgoCreateDescriptorUpdateTemplate(l.funcPtrs.vkCreateDescriptorUpdateTemplate,
		C.VkDevice(device),
		(*C.VkDescriptorUpdateTemplateCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkDescriptorUpdateTemplate)(pDescriptorUpdateTemplate)))

	return res, res.ToError()
}
func (l *vulkanDriver) VkDestroyDescriptorUpdateTemplate(device VkDevice, descriptorUpdateTemplate VkDescriptorUpdateTemplate, pAllocator *VkAllocationCallbacks) {
	if l.funcPtrs.vkDestroyDescriptorUpdateTemplate == nil {
		panic("attempted to call method 'vkDestroyDescriptorUpdateTemplate' which is not present on this driver")
	}

	C.cgoDestroyDescriptorUpdateTemplate(l.funcPtrs.vkDestroyDescriptorUpdateTemplate,
		C.VkDevice(device),
		C.VkDescriptorUpdateTemplate(descriptorUpdateTemplate),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkUpdateDescriptorSetWithTemplate(device VkDevice, descriptorSet VkDescriptorSet, descriptorUpdateTemplate VkDescriptorUpdateTemplate, pData unsafe.Pointer) {
	if l.funcPtrs.vkUpdateDescriptorSetWithTemplate == nil {
		panic("attempted to call method 'vkUpdateDescriptorSetWithTemplate' which is not present on this driver")
	}

	C.cgoUpdateDescriptorSetWithTemplate(l.funcPtrs.vkUpdateDescriptorSetWithTemplate,
		C.VkDevice(device),
		C.VkDescriptorSet(descriptorSet),
		C.VkDescriptorUpdateTemplate(descriptorUpdateTemplate),
		pData)
}

func (l *vulkanDriver) VkGetDescriptorSetLayoutSupport(device VkDevice, pCreateInfo *VkDescriptorSetLayoutCreateInfo, pSupport *VkDescriptorSetLayoutSupport) {
	if l.funcPtrs.vkGetDescriptorSetLayoutSupport == nil {
		panic("attempted to call method 'vkGetDescriptorSetLayoutSupport' which is not present on this driver")
	}

	C.cgoGetDescriptorSetLayoutSupport(l.funcPtrs.vkGetDescriptorSetLayoutSupport,
		C.VkDevice(device),
		(*C.VkDescriptorSetLayoutCreateInfo)(pCreateInfo),
		(*C.VkDescriptorSetLayoutSupport)(pSupport))
}

func (l *vulkanDriver) VkCmdDrawIndirectCount(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, countBuffer VkBuffer, countBufferOffset VkDeviceSize, maxDrawCount Uint32, stride Uint32) {
	if l.funcPtrs.vkCmdDrawIndirectCount == nil {
		panic("attempted to call method 'vkCmdDrawIndirectCount' which is not present on this driver")
	}

	C.cgoCmdDrawIndirectCount(l.funcPtrs.vkCmdDrawIndirectCount,
		C.VkCommandBuffer(commandBuffer),
		C.VkBuffer(buffer),
		C.VkDeviceSize(offset),
		C.VkBuffer(countBuffer),
		C.VkDeviceSize(countBufferOffset),
		C.uint32_t(maxDrawCount),
		C.uint32_t(stride))
}

func (l *vulkanDriver) VkCmdDrawIndexedIndirectCount(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, countBuffer VkBuffer, countBufferOffset VkDeviceSize, maxDrawCount Uint32, stride Uint32) {
	if l.funcPtrs.vkCmdDrawIndexedIndirectCount == nil {
		panic("attempted to call method 'vkCmdDrawIndexedIndirectCount' which is not present on this driver")
	}

	C.cgoCmdDrawIndexedIndirectCount(l.funcPtrs.vkCmdDrawIndexedIndirectCount,
		C.VkCommandBuffer(commandBuffer),
		C.VkBuffer(buffer),
		C.VkDeviceSize(offset),
		C.VkBuffer(countBuffer),
		C.VkDeviceSize(countBufferOffset),
		C.uint32_t(maxDrawCount),
		C.uint32_t(stride))
}

func (l *vulkanDriver) VkCreateRenderPass2(device VkDevice, pCreateInfo *VkRenderPassCreateInfo2, pAllocator *VkAllocationCallbacks, pRenderPass *VkRenderPass) (VkResult, error) {
	if l.funcPtrs.vkCreateRenderPass2 == nil {
		panic("attempted to call method 'vkCreateRenderPass2' which is not present on this driver")
	}

	res := VkResult(C.cgoCreateRenderPass2(l.funcPtrs.vkCreateRenderPass2,
		C.VkDevice(device),
		(*C.VkRenderPassCreateInfo2)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkRenderPass)(pRenderPass)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkCmdBeginRenderPass2(commandBuffer VkCommandBuffer, pRenderPassBegin *VkRenderPassBeginInfo, pSubpassBeginInfo *VkSubpassBeginInfo) {
	if l.funcPtrs.vkCmdBeginRenderPass2 == nil {
		panic("attempted to call method 'vkCmdBeginRenderPass2' which is not present on this driver")
	}

	C.cgoCmdBeginRenderPass2(l.funcPtrs.vkCmdBeginRenderPass2,
		C.VkCommandBuffer(commandBuffer),
		(*C.VkRenderPassBeginInfo)(pRenderPassBegin),
		(*C.VkSubpassBeginInfo)(pSubpassBeginInfo))
}

func (l *vulkanDriver) VkCmdNextSubpass2(commandBuffer VkCommandBuffer, pSubpassBeginInfo *VkSubpassBeginInfo, pSubpassEndInfo *VkSubpassEndInfo) {
	if l.funcPtrs.vkCmdNextSubpass2 == nil {
		panic("attempted to call method 'vkCmdNextSubpass2' which is not present on this driver")
	}

	C.cgoCmdNextSubpass2(l.funcPtrs.vkCmdNextSubpass2,
		C.VkCommandBuffer(commandBuffer),
		(*C.VkSubpassBeginInfo)(pSubpassBeginInfo),
		(*C.VkSubpassEndInfo)(pSubpassEndInfo))
}

func (l *vulkanDriver) VkCmdEndRenderPass2(commandBuffer VkCommandBuffer, pSubpassEndInfo *VkSubpassEndInfo) {
	if l.funcPtrs.vkCmdEndRenderPass2 == nil {
		panic("attempted to call method 'vkCmdEndRenderPass2' which is not present on this driver")
	}

	C.cgoCmdEndRenderPass2(l.funcPtrs.vkCmdEndRenderPass2,
		C.VkCommandBuffer(commandBuffer),
		(*C.VkSubpassEndInfo)(pSubpassEndInfo))
}

func (l *vulkanDriver) VkResetQueryPool(device VkDevice, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32) {
	if l.funcPtrs.vkResetQueryPool == nil {
		panic("attempted to call method 'vkResetQueryPool' which is not present on this driver")
	}

	C.cgoResetQueryPool(l.funcPtrs.vkResetQueryPool,
		C.VkDevice(device),
		C.VkQueryPool(queryPool),
		C.uint32_t(firstQuery),
		C.uint32_t(queryCount))
}

func (l *vulkanDriver) VkGetSemaphoreCounterValue(device VkDevice, semaphore VkSemaphore, pValue *Uint64) (VkResult, error) {
	if l.funcPtrs.vkGetSemaphoreCounterValue == nil {
		panic("attempted to call method 'vkGetSemaphoreCounterValue' which is not present on this driver")
	}

	res := VkResult(C.cgoGetSemaphoreCounterValue(l.funcPtrs.vkGetSemaphoreCounterValue,
		C.VkDevice(device),
		C.VkSemaphore(semaphore),
		(*C.uint64_t)(pValue)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkWaitSemaphores(device VkDevice, pWaitInfo *VkSemaphoreWaitInfo, timeout Uint64) (VkResult, error) {
	if l.funcPtrs.vkWaitSemaphores == nil {
		panic("attempted to call method 'vkWaitSemaphores' which is not present on this driver")
	}

	res := VkResult(C.cgoWaitSemaphores(l.funcPtrs.vkWaitSemaphores,
		C.VkDevice(device),
		(*C.VkSemaphoreWaitInfo)(pWaitInfo),
		C.uint64_t(timeout)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkSignalSemaphore(device VkDevice, pSignalInfo *VkSemaphoreSignalInfo) (VkResult, error) {
	if l.funcPtrs.vkSignalSemaphore == nil {
		panic("attempted to call method 'vkSignalSemaphore' which is not present on this driver")
	}

	res := VkResult(C.cgoSignalSemaphore(l.funcPtrs.vkSignalSemaphore,
		C.VkDevice(device),
		(*C.VkSemaphoreSignalInfo)(pSignalInfo)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkGetBufferDeviceAddress(device VkDevice, pInfo *VkBufferDeviceAddressInfo) VkDeviceAddress {
	if l.funcPtrs.vkGetBufferDeviceAddress == nil {
		panic("attempted to call method 'vkGetBufferDeviceAddress' which is not present on this driver")
	}

	address := VkDeviceAddress(C.cgoGetBufferDeviceAddress(l.funcPtrs.vkGetBufferDeviceAddress,
		C.VkDevice(device),
		(*C.VkBufferDeviceAddressInfo)(pInfo)))

	return address
}

func (l *vulkanDriver) VkGetBufferOpaqueCaptureAddress(device VkDevice, pInfo *VkBufferDeviceAddressInfo) Uint64 {
	if l.funcPtrs.vkGetBufferOpaqueCaptureAddress == nil {
		panic("attempted to call method 'vkGetBufferOpaqueCaptureAddress' which is not present on this driver")
	}

	address := Uint64(C.cgoGetBufferOpaqueCaptureAddress(l.funcPtrs.vkGetBufferOpaqueCaptureAddress,
		C.VkDevice(device),
		(*C.VkBufferDeviceAddressInfo)(pInfo)))

	return address
}

func (l *vulkanDriver) VkGetDeviceMemoryOpaqueCaptureAddress(device VkDevice, pInfo *VkDeviceMemoryOpaqueCaptureAddressInfo) Uint64 {
	if l.funcPtrs.vkGetDeviceMemoryOpaqueCaptureAddress == nil {
		panic("attempted to call method 'vkGetDeviceMemoryOpaqueCaptureAddress' which is not present on this driver")
	}

	address := Uint64(C.cgoGetDeviceMemoryOpaqueCaptureAddress(l.funcPtrs.vkGetDeviceMemoryOpaqueCaptureAddress,
		C.VkDevice(device),
		(*C.VkDeviceMemoryOpaqueCaptureAddressInfo)(pInfo)))

	return address
}
