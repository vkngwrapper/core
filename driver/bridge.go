package driver

/*
#include "driver.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

func (l *vulkanDriver) VkEnumerateInstanceExtensionProperties(pLayerName *Char, pPropertyCount *Uint32, pProperties *VkExtensionProperties) (common.VkResult, error) {
	res := common.VkResult(C.cgoEnumerateInstanceExtensionProperties(l.funcPtrs.vkEnumerateInstanceExtensionProperties,
		(*C.char)(pLayerName),
		(*C.uint32_t)(pPropertyCount),
		(*C.VkExtensionProperties)(pProperties)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkEnumerateInstanceLayerProperties(pPropertyCount *Uint32, pProperties *VkLayerProperties) (common.VkResult, error) {
	res := common.VkResult(C.cgoEnumerateInstanceLayerProperties(l.funcPtrs.vkEnumerateInstanceLayerProperties,
		(*C.uint32_t)(pPropertyCount),
		(*C.VkLayerProperties)(pProperties)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkCreateInstance(pCreateInfo *VkInstanceCreateInfo, pAllocator *VkAllocationCallbacks, pInstance *VkInstance) (common.VkResult, error) {
	res := common.VkResult(C.cgoCreateInstance(l.funcPtrs.vkCreateInstance,
		(*C.VkInstanceCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkInstance)(unsafe.Pointer(pInstance))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkEnumeratePhysicalDevices(instance VkInstance, pPhysicalDeviceCount *Uint32, pPhysicalDevices *VkPhysicalDevice) (common.VkResult, error) {
	if VulkanHandle(l.instance) == NullHandle {
		panic("attempted to call instance driver function on a basic driver")
	}

	res := common.VkResult(C.cgoEnumeratePhysicalDevices(l.funcPtrs.vkEnumeratePhysicalDevices,
		(C.VkInstance)(unsafe.Pointer(instance)),
		(*C.uint32_t)(pPhysicalDeviceCount),
		(*C.VkPhysicalDevice)(unsafe.Pointer(pPhysicalDevices))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyInstance(instance VkInstance, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.instance) == NullHandle {
		panic("attempted to call instance driver function on a basic driver")
	}

	C.cgoDestroyInstance(l.funcPtrs.vkDestroyInstance,
		(C.VkInstance)(unsafe.Pointer(instance)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkGetPhysicalDeviceFeatures(physicalDevice VkPhysicalDevice, pFeatures *VkPhysicalDeviceFeatures) {
	if VulkanHandle(l.instance) == NullHandle {
		panic("attempted to call instance driver function on a basic driver")
	}

	C.cgoGetPhysicalDeviceFeatures(l.funcPtrs.vkGetPhysicalDeviceFeatures,
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceFeatures)(pFeatures))
}

func (l *vulkanDriver) VkGetPhysicalDeviceFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, pFormatProperties *VkFormatProperties) {
	if VulkanHandle(l.instance) == NullHandle {
		panic("attempted to call instance driver function on a basic driver")
	}

	C.cgoGetPhysicalDeviceFormatProperties(l.funcPtrs.vkGetPhysicalDeviceFormatProperties,
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(C.VkFormat)(format),
		(*C.VkFormatProperties)(pFormatProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceImageFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, t VkImageType, tiling VkImageTiling, usage VkImageUsageFlags, flags VkImageCreateFlags, pImageFormatProperties *VkImageFormatProperties) (common.VkResult, error) {
	if VulkanHandle(l.instance) == NullHandle {
		panic("attempted to call instance driver function on a basic driver")
	}

	res := common.VkResult(C.cgoGetPhysicalDeviceImageFormatProperties(l.funcPtrs.vkGetPhysicalDeviceImageFormatProperties,
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(C.VkFormat)(format),
		(C.VkImageType)(t),
		(C.VkImageTiling)(tiling),
		(C.VkImageUsageFlags)(usage),
		(C.VkImageCreateFlags)(flags),
		(*C.VkImageFormatProperties)(pImageFormatProperties)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkGetPhysicalDeviceProperties(physicalDevice VkPhysicalDevice, pProperties *VkPhysicalDeviceProperties) {
	if VulkanHandle(l.instance) == NullHandle {
		panic("attempted to call instance driver function on a basic driver")
	}

	C.cgoGetPhysicalDeviceProperties(l.funcPtrs.vkGetPhysicalDeviceProperties,
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceProperties)(pProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceQueueFamilyProperties(physicalDevice VkPhysicalDevice, pQueueFamilyPropertyCount *Uint32, pQueueFamilyProperties *VkQueueFamilyProperties) {
	if VulkanHandle(l.instance) == NullHandle {
		panic("attempted to call instance driver function on a basic driver")
	}

	C.cgoGetPhysicalDeviceQueueFamilyProperties(l.funcPtrs.vkGetPhysicalDeviceQueueFamilyProperties,
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(*C.uint32_t)(pQueueFamilyPropertyCount),
		(*C.VkQueueFamilyProperties)(pQueueFamilyProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceMemoryProperties(physicalDevice VkPhysicalDevice, pMemoryProperties *VkPhysicalDeviceMemoryProperties) {
	if VulkanHandle(l.instance) == NullHandle {
		panic("attempted to call instance driver function on a basic driver")
	}

	C.cgoGetPhysicalDeviceMemoryProperties(l.funcPtrs.vkGetPhysicalDeviceMemoryProperties,
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceMemoryProperties)(pMemoryProperties))
}

func (l *vulkanDriver) VkEnumerateDeviceExtensionProperties(physicalDevice VkPhysicalDevice, pLayerName *Char, pPropertyCount *Uint32, pProperties *VkExtensionProperties) (common.VkResult, error) {
	if VulkanHandle(l.instance) == NullHandle {
		panic("attempted to call instance driver function on a basic driver")
	}

	res := common.VkResult(C.cgoEnumerateDeviceExtensionProperties(l.funcPtrs.vkEnumerateDeviceExtensionProperties,
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(*C.char)(pLayerName),
		(*C.uint32_t)(pPropertyCount),
		(*C.VkExtensionProperties)(pProperties)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkEnumerateDeviceLayerProperties(physicalDevice VkPhysicalDevice, pPropertyCount *Uint32, pProperties *VkLayerProperties) (common.VkResult, error) {
	if VulkanHandle(l.instance) == NullHandle {
		panic("attempted to call instance driver function on a basic driver")
	}

	res := common.VkResult(C.cgoEnumerateDeviceLayerProperties(l.funcPtrs.vkEnumerateDeviceLayerProperties,
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(*C.uint32_t)(pPropertyCount),
		(*C.VkLayerProperties)(pProperties)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkGetPhysicalDeviceSparseImageFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, t VkImageType, samples VkSampleCountFlagBits, usage VkImageUsageFlags, tiling VkImageTiling, pPropertyCount *Uint32, pProperties *VkSparseImageFormatProperties) {
	if VulkanHandle(l.instance) == NullHandle {
		panic("attempted to call instance driver function on a basic driver")
	}

	C.cgoGetPhysicalDeviceSparseImageFormatProperties(l.funcPtrs.vkGetPhysicalDeviceSparseImageFormatProperties,
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(C.VkFormat)(format),
		(C.VkImageType)(t),
		(C.VkSampleCountFlagBits)(samples),
		(C.VkImageUsageFlags)(usage),
		(C.VkImageTiling)(tiling),
		(*C.uint32_t)(pPropertyCount),
		(*C.VkSparseImageFormatProperties)(pProperties))
}

func (l *vulkanDriver) VkCreateDevice(physicalDevice VkPhysicalDevice, pCreateInfo *VkDeviceCreateInfo, pAllocator *VkAllocationCallbacks, pDevice *VkDevice) (common.VkResult, error) {
	if VulkanHandle(l.instance) == NullHandle {
		panic("attempted to call instance driver function on a basic driver")
	}

	res := common.VkResult(C.cgoCreateDevice(l.funcPtrs.vkCreateDevice,
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(*C.VkDeviceCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkDevice)(unsafe.Pointer(pDevice))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyDevice(device VkDevice, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyDevice(l.funcPtrs.vkDestroyDevice,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkGetDeviceQueue(device VkDevice, queueFamilyIndex Uint32, queueIndex Uint32, pQueue *VkQueue) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoGetDeviceQueue(l.funcPtrs.vkGetDeviceQueue,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.uint32_t)(queueFamilyIndex),
		(C.uint32_t)(queueIndex),
		(*C.VkQueue)(unsafe.Pointer(pQueue)))
}

func (l *vulkanDriver) VkQueueSubmit(queue VkQueue, submitCount Uint32, pSubmits *VkSubmitInfo, fence VkFence) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoQueueSubmit(l.funcPtrs.vkQueueSubmit,
		(C.VkQueue)(unsafe.Pointer(queue)),
		(C.uint32_t)(submitCount),
		(*C.VkSubmitInfo)(pSubmits),
		(C.VkFence)(unsafe.Pointer(fence))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkQueueWaitIdle(queue VkQueue) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoQueueWaitIdle(l.funcPtrs.vkQueueWaitIdle,
		(C.VkQueue)(unsafe.Pointer(queue))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDeviceWaitIdle(device VkDevice) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoDeviceWaitIdle(l.funcPtrs.vkDeviceWaitIdle,
		(C.VkDevice)(unsafe.Pointer(device))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkAllocateMemory(device VkDevice, pAllocateInfo *VkMemoryAllocateInfo, pAllocator *VkAllocationCallbacks, pMemory *VkDeviceMemory) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoAllocateMemory(l.funcPtrs.vkAllocateMemory,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkMemoryAllocateInfo)(pAllocateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkDeviceMemory)(unsafe.Pointer(pMemory))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkFreeMemory(device VkDevice, memory VkDeviceMemory, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoFreeMemory(l.funcPtrs.vkFreeMemory,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkDeviceMemory)(unsafe.Pointer(memory)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkMapMemory(device VkDevice, memory VkDeviceMemory, offset VkDeviceSize, size VkDeviceSize, flags VkMemoryMapFlags, ppData *unsafe.Pointer) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoMapMemory(l.funcPtrs.vkMapMemory,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkDeviceMemory)(unsafe.Pointer(memory)),
		(C.VkDeviceSize)(offset),
		(C.VkDeviceSize)(size),
		(C.VkMemoryMapFlags)(flags),
		ppData))
	return res, res.ToError()
}

func (l *vulkanDriver) VkUnmapMemory(device VkDevice, memory VkDeviceMemory) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoUnmapMemory(l.funcPtrs.vkUnmapMemory,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkDeviceMemory)(unsafe.Pointer(memory)))
}

func (l *vulkanDriver) VkFlushMappedMemoryRanges(device VkDevice, memoryRangeCount Uint32, pMemoryRanges *VkMappedMemoryRange) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoFlushMappedMemoryRanges(l.funcPtrs.vkFlushMappedMemoryRanges,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.uint32_t)(memoryRangeCount),
		(*C.VkMappedMemoryRange)(pMemoryRanges)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkInvalidateMappedMemoryRanges(device VkDevice, memoryRangeCount Uint32, pMemoryRanges *VkMappedMemoryRange) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoInvalidateMappedMemoryRanges(l.funcPtrs.vkInvalidateMappedMemoryRanges,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.uint32_t)(memoryRangeCount),
		(*C.VkMappedMemoryRange)(pMemoryRanges)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkGetDeviceMemoryCommitment(device VkDevice, memory VkDeviceMemory, pCommittedMemoryInBytes *VkDeviceSize) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoGetDeviceMemoryCommitment(l.funcPtrs.vkGetDeviceMemoryCommitment,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkDeviceMemory)(unsafe.Pointer(memory)),
		(*C.VkDeviceSize)(pCommittedMemoryInBytes))
}

func (l *vulkanDriver) VkBindBufferMemory(device VkDevice, buffer VkBuffer, memory VkDeviceMemory, memoryOffset VkDeviceSize) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoBindBufferMemory(l.funcPtrs.vkBindBufferMemory,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkBuffer)(unsafe.Pointer(buffer)),
		(C.VkDeviceMemory)(unsafe.Pointer(memory)),
		(C.VkDeviceSize)(memoryOffset)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkBindImageMemory(device VkDevice, image VkImage, memory VkDeviceMemory, memoryOffset VkDeviceSize) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoBindImageMemory(l.funcPtrs.vkBindImageMemory,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkImage)(unsafe.Pointer(image)),
		(C.VkDeviceMemory)(unsafe.Pointer(memory)),
		(C.VkDeviceSize)(memoryOffset)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkGetBufferMemoryRequirements(device VkDevice, buffer VkBuffer, pMemoryRequirements *VkMemoryRequirements) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoGetBufferMemoryRequirements(l.funcPtrs.vkGetBufferMemoryRequirements,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkBuffer)(unsafe.Pointer(buffer)),
		(*C.VkMemoryRequirements)(pMemoryRequirements))
}

func (l *vulkanDriver) VkGetImageMemoryRequirements(device VkDevice, image VkImage, pMemoryRequirements *VkMemoryRequirements) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoGetImageMemoryRequirements(l.funcPtrs.vkGetImageMemoryRequirements,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkImage)(unsafe.Pointer(image)),
		(*C.VkMemoryRequirements)(pMemoryRequirements))
}

func (l *vulkanDriver) VkGetImageSparseMemoryRequirements(device VkDevice, image VkImage, pSparseMemoryRequirementCount *Uint32, pSparseMemoryRequirements *VkSparseImageMemoryRequirements) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoGetImageSparseMemoryRequirements(l.funcPtrs.vkGetImageSparseMemoryRequirements,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkImage)(unsafe.Pointer(image)),
		(*C.uint32_t)(pSparseMemoryRequirementCount),
		(*C.VkSparseImageMemoryRequirements)(pSparseMemoryRequirements))
}

func (l *vulkanDriver) VkQueueBindSparse(queue VkQueue, bindInfoCount Uint32, pBindInfo *VkBindSparseInfo, fence VkFence) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoQueueBindSparse(l.funcPtrs.vkQueueBindSparse,
		(C.VkQueue)(unsafe.Pointer(queue)),
		(C.uint32_t)(bindInfoCount),
		(*C.VkBindSparseInfo)(pBindInfo),
		(C.VkFence)(unsafe.Pointer(fence))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkCreateFence(device VkDevice, pCreateInfo *VkFenceCreateInfo, pAllocator *VkAllocationCallbacks, pFence *VkFence) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreateFence(l.funcPtrs.vkCreateFence,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkFenceCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkFence)(unsafe.Pointer(pFence))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyFence(device VkDevice, fence VkFence, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyFence(l.funcPtrs.vkDestroyFence,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkFence)(unsafe.Pointer(fence)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkResetFences(device VkDevice, fenceCount Uint32, pFences *VkFence) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoResetFences(l.funcPtrs.vkResetFences,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.uint32_t)(fenceCount),
		(*C.VkFence)(unsafe.Pointer(pFences))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkGetFenceStatus(device VkDevice, fence VkFence) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoGetFenceStatus(l.funcPtrs.vkGetFenceStatus,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkFence)(unsafe.Pointer(fence))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkWaitForFences(device VkDevice, fenceCount Uint32, pFences *VkFence, waitAll VkBool32, timeout Uint64) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoWaitForFences(l.funcPtrs.vkWaitForFences,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.uint32_t)(fenceCount),
		(*C.VkFence)(unsafe.Pointer(pFences)),
		(C.VkBool32)(waitAll),
		(C.uint64_t)(timeout)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkCreateSemaphore(device VkDevice, pCreateInfo *VkSemaphoreCreateInfo, pAllocator *VkAllocationCallbacks, pSemaphore *VkSemaphore) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreateSemaphore(l.funcPtrs.vkCreateSemaphore,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkSemaphoreCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkSemaphore)(unsafe.Pointer(pSemaphore))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroySemaphore(device VkDevice, semaphore VkSemaphore, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroySemaphore(l.funcPtrs.vkDestroySemaphore,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkSemaphore)(unsafe.Pointer(semaphore)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateEvent(device VkDevice, pCreateInfo *VkEventCreateInfo, pAllocator *VkAllocationCallbacks, pEvent *VkEvent) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreateEvent(l.funcPtrs.vkCreateEvent,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkEventCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkEvent)(unsafe.Pointer(pEvent))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyEvent(device VkDevice, event VkEvent, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyEvent(l.funcPtrs.vkDestroyEvent,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkEvent)(unsafe.Pointer(event)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkGetEventStatus(device VkDevice, event VkEvent) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoGetEventStatus(l.funcPtrs.vkGetEventStatus,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkEvent)(unsafe.Pointer(event))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkSetEvent(device VkDevice, event VkEvent) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoSetEvent(l.funcPtrs.vkSetEvent,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkEvent)(unsafe.Pointer(event))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkResetEvent(device VkDevice, event VkEvent) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoResetEvent(l.funcPtrs.vkResetEvent,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkEvent)(unsafe.Pointer(event))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkCreateQueryPool(device VkDevice, pCreateInfo *VkQueryPoolCreateInfo, pAllocator *VkAllocationCallbacks, pQueryPool *VkQueryPool) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreateQueryPool(l.funcPtrs.vkCreateQueryPool,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkQueryPoolCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkQueryPool)(unsafe.Pointer(pQueryPool))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyQueryPool(device VkDevice, queryPool VkQueryPool, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyQueryPool(l.funcPtrs.vkDestroyQueryPool,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkQueryPool)(unsafe.Pointer(queryPool)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkGetQueryPoolResults(device VkDevice, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32, dataSize Size, pData unsafe.Pointer, stride VkDeviceSize, flags VkQueryResultFlags) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoGetQueryPoolResults(l.funcPtrs.vkGetQueryPoolResults,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkQueryPool)(unsafe.Pointer(queryPool)),
		(C.uint32_t)(firstQuery),
		(C.uint32_t)(queryCount),
		(C.size_t)(dataSize),
		pData,
		(C.VkDeviceSize)(stride),
		(C.VkQueryResultFlags)(flags)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkCreateBuffer(device VkDevice, pCreateInfo *VkBufferCreateInfo, pAllocator *VkAllocationCallbacks, pBuffer *VkBuffer) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreateBuffer(l.funcPtrs.vkCreateBuffer,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkBufferCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkBuffer)(unsafe.Pointer(pBuffer))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyBuffer(device VkDevice, buffer VkBuffer, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyBuffer(l.funcPtrs.vkDestroyBuffer,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkBuffer)(unsafe.Pointer(buffer)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateBufferView(device VkDevice, pCreateInfo *VkBufferViewCreateInfo, pAllocator *VkAllocationCallbacks, pView *VkBufferView) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreateBufferView(l.funcPtrs.vkCreateBufferView,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkBufferViewCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkBufferView)(unsafe.Pointer(pView))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyBufferView(device VkDevice, bufferView VkBufferView, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyBufferView(l.funcPtrs.vkDestroyBufferView,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkBufferView)(unsafe.Pointer(bufferView)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateImage(device VkDevice, pCreateInfo *VkImageCreateInfo, pAllocator *VkAllocationCallbacks, pImage *VkImage) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreateImage(l.funcPtrs.vkCreateImage,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkImageCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkImage)(unsafe.Pointer(pImage))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyImage(device VkDevice, image VkImage, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyImage(l.funcPtrs.vkDestroyImage,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkImage)(unsafe.Pointer(image)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkGetImageSubresourceLayout(device VkDevice, image VkImage, pSubresource *VkImageSubresource, pLayout *VkSubresourceLayout) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoGetImageSubresourceLayout(l.funcPtrs.vkGetImageSubresourceLayout,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkImage)(unsafe.Pointer(image)),
		(*C.VkImageSubresource)(pSubresource),
		(*C.VkSubresourceLayout)(pLayout))
}

func (l *vulkanDriver) VkCreateImageView(device VkDevice, pCreateInfo *VkImageViewCreateInfo, pAllocator *VkAllocationCallbacks, pView *VkImageView) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreateImageView(l.funcPtrs.vkCreateImageView,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkImageViewCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkImageView)(unsafe.Pointer(pView))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyImageView(device VkDevice, imageView VkImageView, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyImageView(l.funcPtrs.vkDestroyImageView,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkImageView)(unsafe.Pointer(imageView)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateShaderModule(device VkDevice, pCreateInfo *VkShaderModuleCreateInfo, pAllocator *VkAllocationCallbacks, pShaderModule *VkShaderModule) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreateShaderModule(l.funcPtrs.vkCreateShaderModule,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkShaderModuleCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkShaderModule)(unsafe.Pointer(pShaderModule))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyShaderModule(device VkDevice, shaderModule VkShaderModule, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyShaderModule(l.funcPtrs.vkDestroyShaderModule,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkShaderModule)(unsafe.Pointer(shaderModule)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreatePipelineCache(device VkDevice, pCreateInfo *VkPipelineCacheCreateInfo, pAllocator *VkAllocationCallbacks, pPipelineCache *VkPipelineCache) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreatePipelineCache(l.funcPtrs.vkCreatePipelineCache,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkPipelineCacheCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkPipelineCache)(unsafe.Pointer(pPipelineCache))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyPipelineCache(device VkDevice, pipelineCache VkPipelineCache, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyPipelineCache(l.funcPtrs.vkDestroyPipelineCache,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkPipelineCache)(unsafe.Pointer(pipelineCache)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkGetPipelineCacheData(device VkDevice, pipelineCache VkPipelineCache, pDataSize *Size, pData unsafe.Pointer) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoGetPipelineCacheData(l.funcPtrs.vkGetPipelineCacheData,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkPipelineCache)(unsafe.Pointer(pipelineCache)),
		(*C.size_t)(pDataSize),
		pData))
	return res, res.ToError()
}

func (l *vulkanDriver) VkMergePipelineCaches(device VkDevice, dstCache VkPipelineCache, srcCacheCount Uint32, pSrcCaches *VkPipelineCache) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoMergePipelineCaches(l.funcPtrs.vkMergePipelineCaches,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkPipelineCache)(unsafe.Pointer(dstCache)),
		(C.uint32_t)(srcCacheCount),
		(*C.VkPipelineCache)(unsafe.Pointer(pSrcCaches))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkCreateGraphicsPipelines(device VkDevice, pipelineCache VkPipelineCache, createInfoCount Uint32, pCreateInfos *VkGraphicsPipelineCreateInfo, pAllocator *VkAllocationCallbacks, pPipelines *VkPipeline) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreateGraphicsPipelines(l.funcPtrs.vkCreateGraphicsPipelines,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkPipelineCache)(unsafe.Pointer(pipelineCache)),
		(C.uint32_t)(createInfoCount),
		(*C.VkGraphicsPipelineCreateInfo)(pCreateInfos),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkPipeline)(unsafe.Pointer(pPipelines))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkCreateComputePipelines(device VkDevice, pipelineCache VkPipelineCache, createInfoCount Uint32, pCreateInfos *VkComputePipelineCreateInfo, pAllocator *VkAllocationCallbacks, pPipelines *VkPipeline) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreateComputePipelines(l.funcPtrs.vkCreateComputePipelines,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkPipelineCache)(unsafe.Pointer(pipelineCache)),
		(C.uint32_t)(createInfoCount),
		(*C.VkComputePipelineCreateInfo)(pCreateInfos),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkPipeline)(unsafe.Pointer(pPipelines))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyPipeline(device VkDevice, pipeline VkPipeline, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyPipeline(l.funcPtrs.vkDestroyPipeline,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkPipeline)(unsafe.Pointer(pipeline)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreatePipelineLayout(device VkDevice, pCreateInfo *VkPipelineLayoutCreateInfo, pAllocator *VkAllocationCallbacks, pPipelineLayout *VkPipelineLayout) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreatePipelineLayout(l.funcPtrs.vkCreatePipelineLayout,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkPipelineLayoutCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkPipelineLayout)(unsafe.Pointer(pPipelineLayout))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyPipelineLayout(device VkDevice, pipelineLayout VkPipelineLayout, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyPipelineLayout(l.funcPtrs.vkDestroyPipelineLayout,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkPipelineLayout)(unsafe.Pointer(pipelineLayout)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateSampler(device VkDevice, pCreateInfo *VkSamplerCreateInfo, pAllocator *VkAllocationCallbacks, pSampler *VkSampler) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreateSampler(l.funcPtrs.vkCreateSampler,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkSamplerCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkSampler)(unsafe.Pointer(pSampler))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroySampler(device VkDevice, sampler VkSampler, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroySampler(l.funcPtrs.vkDestroySampler,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkSampler)(unsafe.Pointer(sampler)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateDescriptorSetLayout(device VkDevice, pCreateInfo *VkDescriptorSetLayoutCreateInfo, pAllocator *VkAllocationCallbacks, pSetLayout *VkDescriptorSetLayout) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreateDescriptorSetLayout(l.funcPtrs.vkCreateDescriptorSetLayout,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkDescriptorSetLayoutCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkDescriptorSetLayout)(unsafe.Pointer(pSetLayout))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyDescriptorSetLayout(device VkDevice, descriptorSetLayout VkDescriptorSetLayout, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyDescriptorSetLayout(l.funcPtrs.vkDestroyDescriptorSetLayout,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkDescriptorSetLayout)(unsafe.Pointer(descriptorSetLayout)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateDescriptorPool(device VkDevice, pCreateInfo *VkDescriptorPoolCreateInfo, pAllocator *VkAllocationCallbacks, pDescriptorPool *VkDescriptorPool) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreateDescriptorPool(l.funcPtrs.vkCreateDescriptorPool,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkDescriptorPoolCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkDescriptorPool)(unsafe.Pointer(pDescriptorPool))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyDescriptorPool(device VkDevice, descriptorPool VkDescriptorPool, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyDescriptorPool(l.funcPtrs.vkDestroyDescriptorPool,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkDescriptorPool)(unsafe.Pointer(descriptorPool)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkResetDescriptorPool(device VkDevice, descriptorPool VkDescriptorPool, flags VkDescriptorPoolResetFlags) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoResetDescriptorPool(l.funcPtrs.vkResetDescriptorPool,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkDescriptorPool)(unsafe.Pointer(descriptorPool)),
		(C.VkDescriptorPoolResetFlags)(flags)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkAllocateDescriptorSets(device VkDevice, pAllocateInfo *VkDescriptorSetAllocateInfo, pDescriptorSets *VkDescriptorSet) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoAllocateDescriptorSets(l.funcPtrs.vkAllocateDescriptorSets,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkDescriptorSetAllocateInfo)(pAllocateInfo),
		(*C.VkDescriptorSet)(unsafe.Pointer(pDescriptorSets))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkFreeDescriptorSets(device VkDevice, descriptorPool VkDescriptorPool, descriptorSetCount Uint32, pDescriptorSets *VkDescriptorSet) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoFreeDescriptorSets(l.funcPtrs.vkFreeDescriptorSets,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkDescriptorPool)(unsafe.Pointer(descriptorPool)),
		(C.uint32_t)(descriptorSetCount),
		(*C.VkDescriptorSet)(unsafe.Pointer(pDescriptorSets))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkUpdateDescriptorSets(device VkDevice, descriptorWriteCount Uint32, pDescriptorWrites *VkWriteDescriptorSet, descriptorCopyCount Uint32, pDescriptorCopies *VkCopyDescriptorSet) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoUpdateDescriptorSets(l.funcPtrs.vkUpdateDescriptorSets,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.uint32_t)(descriptorWriteCount),
		(*C.VkWriteDescriptorSet)(pDescriptorWrites),
		(C.uint32_t)(descriptorCopyCount),
		(*C.VkCopyDescriptorSet)(pDescriptorCopies))
}

func (l *vulkanDriver) VkCreateFramebuffer(device VkDevice, pCreateInfo *VkFramebufferCreateInfo, pAllocator *VkAllocationCallbacks, pFramebuffer *VkFramebuffer) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreateFramebuffer(l.funcPtrs.vkCreateFramebuffer,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkFramebufferCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkFramebuffer)(unsafe.Pointer(pFramebuffer))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyFramebuffer(device VkDevice, framebuffer VkFramebuffer, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyFramebuffer(l.funcPtrs.vkDestroyFramebuffer,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkFramebuffer)(unsafe.Pointer(framebuffer)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateRenderPass(device VkDevice, pCreateInfo *VkRenderPassCreateInfo, pAllocator *VkAllocationCallbacks, pRenderPass *VkRenderPass) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreateRenderPass(l.funcPtrs.vkCreateRenderPass,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkRenderPassCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkRenderPass)(unsafe.Pointer(pRenderPass))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyRenderPass(device VkDevice, renderPass VkRenderPass, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyRenderPass(l.funcPtrs.vkDestroyRenderPass,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkRenderPass)(unsafe.Pointer(renderPass)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkGetRenderAreaGranularity(device VkDevice, renderPass VkRenderPass, pGranularity *VkExtent2D) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoGetRenderAreaGranularity(l.funcPtrs.vkGetRenderAreaGranularity,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkRenderPass)(unsafe.Pointer(renderPass)),
		(*C.VkExtent2D)(pGranularity))
}

func (l *vulkanDriver) VkCreateCommandPool(device VkDevice, pCreateInfo *VkCommandPoolCreateInfo, pAllocator *VkAllocationCallbacks, pCommandPool *VkCommandPool) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoCreateCommandPool(l.funcPtrs.vkCreateCommandPool,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkCommandPoolCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkCommandPool)(unsafe.Pointer(pCommandPool))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroyCommandPool(device VkDevice, commandPool VkCommandPool, pAllocator *VkAllocationCallbacks) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoDestroyCommandPool(l.funcPtrs.vkDestroyCommandPool,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkCommandPool)(unsafe.Pointer(commandPool)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkResetCommandPool(device VkDevice, commandPool VkCommandPool, flags VkCommandPoolResetFlags) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoResetCommandPool(l.funcPtrs.vkResetCommandPool,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkCommandPool)(unsafe.Pointer(commandPool)),
		(C.VkCommandPoolResetFlags)(flags)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkAllocateCommandBuffers(device VkDevice, pAllocateInfo *VkCommandBufferAllocateInfo, pCommandBuffers *VkCommandBuffer) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoAllocateCommandBuffers(l.funcPtrs.vkAllocateCommandBuffers,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkCommandBufferAllocateInfo)(pAllocateInfo),
		(*C.VkCommandBuffer)(unsafe.Pointer(pCommandBuffers))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkFreeCommandBuffers(device VkDevice, commandPool VkCommandPool, commandBufferCount Uint32, pCommandBuffers *VkCommandBuffer) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoFreeCommandBuffers(l.funcPtrs.vkFreeCommandBuffers,
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkCommandPool)(unsafe.Pointer(commandPool)),
		(C.uint32_t)(commandBufferCount),
		(*C.VkCommandBuffer)(unsafe.Pointer(pCommandBuffers)))
}

func (l *vulkanDriver) VkBeginCommandBuffer(commandBuffer VkCommandBuffer, pBeginInfo *VkCommandBufferBeginInfo) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoBeginCommandBuffer(l.funcPtrs.vkBeginCommandBuffer,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(*C.VkCommandBufferBeginInfo)(pBeginInfo)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkEndCommandBuffer(commandBuffer VkCommandBuffer) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoEndCommandBuffer(l.funcPtrs.vkEndCommandBuffer,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer))))
	return res, res.ToError()
}

func (l *vulkanDriver) VkResetCommandBuffer(commandBuffer VkCommandBuffer, flags VkCommandBufferResetFlags) (common.VkResult, error) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	res := common.VkResult(C.cgoResetCommandBuffer(l.funcPtrs.vkResetCommandBuffer,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkCommandBufferResetFlags)(flags)))
	return res, res.ToError()
}

func (l *vulkanDriver) VkCmdBindPipeline(commandBuffer VkCommandBuffer, pipelineBindPoint VkPipelineBindPoint, pipeline VkPipeline) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdBindPipeline(l.funcPtrs.vkCmdBindPipeline,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkPipelineBindPoint)(pipelineBindPoint),
		(C.VkPipeline)(unsafe.Pointer(pipeline)))
}

func (l *vulkanDriver) VkCmdSetViewport(commandBuffer VkCommandBuffer, firstViewport Uint32, viewportCount Uint32, pViewports *VkViewport) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetViewport(l.funcPtrs.vkCmdSetViewport,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.uint32_t)(firstViewport),
		(C.uint32_t)(viewportCount),
		(*C.VkViewport)(pViewports))
}

func (l *vulkanDriver) VkCmdSetScissor(commandBuffer VkCommandBuffer, firstScissor Uint32, scissorCount Uint32, pScissors *VkRect2D) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetScissor(l.funcPtrs.vkCmdSetScissor,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		C.uint32_t(firstScissor),
		C.uint32_t(scissorCount),
		(*C.VkRect2D)(pScissors))
}

func (l *vulkanDriver) VkCmdSetLineWidth(commandBuffer VkCommandBuffer, lineWidth Float) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}
	C.cgoCmdSetLineWidth(l.funcPtrs.vkCmdSetLineWidth,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.float)(lineWidth))
}

func (l *vulkanDriver) VkCmdSetDepthBias(commandBuffer VkCommandBuffer, depthBiasConstantFactor Float, depthBiasClamp Float, depthBiasSlopeFactor Float) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetDepthBias(l.funcPtrs.vkCmdSetDepthBias,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.float)(depthBiasConstantFactor),
		(C.float)(depthBiasClamp),
		(C.float)(depthBiasSlopeFactor))
}

func (l *vulkanDriver) VkCmdSetBlendConstants(commandBuffer VkCommandBuffer, blendConstants *Float) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetBlendConstants(l.funcPtrs.vkCmdSetBlendConstants,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(*C.float)(blendConstants),
	)
}

func (l *vulkanDriver) VkCmdSetDepthBounds(commandBuffer VkCommandBuffer, minDepthBounds Float, maxDepthBounds Float) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetDepthBounds(l.funcPtrs.vkCmdSetDepthBounds,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.float)(minDepthBounds),
		(C.float)(maxDepthBounds))
}

func (l *vulkanDriver) VkCmdSetStencilCompareMask(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, compareMask Uint32) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetStencilCompareMask(l.funcPtrs.vkCmdSetStencilCompareMask,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkStencilFaceFlags)(faceMask),
		(C.uint32_t)(compareMask))
}

func (l *vulkanDriver) VkCmdSetStencilWriteMask(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, writeMask Uint32) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetStencilWriteMask(l.funcPtrs.vkCmdSetStencilWriteMask,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkStencilFaceFlags)(faceMask),
		(C.uint32_t)(writeMask))
}

func (l *vulkanDriver) VkCmdSetStencilReference(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, reference Uint32) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetStencilReference(l.funcPtrs.vkCmdSetStencilReference,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkStencilFaceFlags)(faceMask),
		(C.uint32_t)(reference))
}

func (l *vulkanDriver) VkCmdBindDescriptorSets(commandBuffer VkCommandBuffer, pipelineBindPoint VkPipelineBindPoint, layout VkPipelineLayout, firstSet Uint32, descriptorSetCount Uint32, pDescriptorSets *VkDescriptorSet, dynamicOffsetCount Uint32, pDynamicOffsets *Uint32) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdBindDescriptorSets(l.funcPtrs.vkCmdBindDescriptorSets,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkPipelineBindPoint)(pipelineBindPoint),
		(C.VkPipelineLayout)(unsafe.Pointer(layout)),
		(C.uint32_t)(firstSet),
		(C.uint32_t)(descriptorSetCount),
		(*C.VkDescriptorSet)(unsafe.Pointer(pDescriptorSets)),
		(C.uint32_t)(dynamicOffsetCount),
		(*C.uint32_t)(pDynamicOffsets))
}

func (l *vulkanDriver) VkCmdBindIndexBuffer(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, indexType VkIndexType) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdBindIndexBuffer(l.funcPtrs.vkCmdBindIndexBuffer,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkBuffer)(unsafe.Pointer(buffer)),
		(C.VkDeviceSize)(offset),
		(C.VkIndexType)(indexType))
}

func (l *vulkanDriver) VkCmdBindVertexBuffers(commandBuffer VkCommandBuffer, firstBinding Uint32, bindingCount Uint32, pBuffers *VkBuffer, pOffsets *VkDeviceSize) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdBindVertexBuffers(l.funcPtrs.vkCmdBindVertexBuffers,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		C.uint32_t(firstBinding),
		C.uint32_t(bindingCount),
		(*C.VkBuffer)(unsafe.Pointer(pBuffers)),
		(*C.VkDeviceSize)(pOffsets))
}

func (l *vulkanDriver) VkCmdDraw(commandBuffer VkCommandBuffer, vertexCount Uint32, instanceCount Uint32, firstVertex Uint32, firstInstance Uint32) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdDraw(l.funcPtrs.vkCmdDraw,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.uint32_t)(vertexCount),
		(C.uint32_t)(instanceCount),
		(C.uint32_t)(firstVertex),
		(C.uint32_t)(firstInstance))
}

func (l *vulkanDriver) VkCmdDrawIndexed(commandBuffer VkCommandBuffer, indexCount Uint32, instanceCount Uint32, firstIndex Uint32, vertexOffset Int32, firstInstance Uint32) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdDrawIndexed(l.funcPtrs.vkCmdDrawIndexed,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.uint32_t)(indexCount),
		(C.uint32_t)(instanceCount),
		(C.uint32_t)(firstIndex),
		(C.int32_t)(vertexOffset),
		(C.uint32_t)(firstInstance))
}

func (l *vulkanDriver) VkCmdDrawIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, drawCount Uint32, stride Uint32) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdDrawIndirect(l.funcPtrs.vkCmdDrawIndirect,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkBuffer)(unsafe.Pointer(buffer)),
		(C.VkDeviceSize)(offset),
		(C.uint32_t)(drawCount),
		(C.uint32_t)(stride))
}

func (l *vulkanDriver) VkCmdDrawIndexedIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, drawCount Uint32, stride Uint32) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdDrawIndexedIndirect(l.funcPtrs.vkCmdDrawIndexedIndirect,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkBuffer)(unsafe.Pointer(buffer)),
		(C.VkDeviceSize)(offset),
		(C.uint32_t)(drawCount),
		(C.uint32_t)(stride))
}

func (l *vulkanDriver) VkCmdDispatch(commandBuffer VkCommandBuffer, groupCountX Uint32, groupCountY Uint32, groupCountZ Uint32) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdDispatch(l.funcPtrs.vkCmdDispatch,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.uint32_t)(groupCountX),
		(C.uint32_t)(groupCountY),
		(C.uint32_t)(groupCountZ))
}

func (l *vulkanDriver) VkCmdDispatchIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdDispatchIndirect(l.funcPtrs.vkCmdDispatchIndirect,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkBuffer)(unsafe.Pointer(buffer)),
		(C.VkDeviceSize)(offset))
}

func (l *vulkanDriver) VkCmdCopyBuffer(commandBuffer VkCommandBuffer, srcBuffer VkBuffer, dstBuffer VkBuffer, regionCount Uint32, pRegions *VkBufferCopy) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdCopyBuffer(l.funcPtrs.vkCmdCopyBuffer,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkBuffer)(unsafe.Pointer(srcBuffer)),
		(C.VkBuffer)(unsafe.Pointer(dstBuffer)),
		(C.uint32_t)(regionCount),
		(*C.VkBufferCopy)(pRegions))
}

func (l *vulkanDriver) VkCmdCopyImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkImageCopy) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdCopyImage(l.funcPtrs.vkCmdCopyImage,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkImage)(unsafe.Pointer(srcImage)),
		(C.VkImageLayout)(srcImageLayout),
		(C.VkImage)(unsafe.Pointer(dstImage)),
		(C.VkImageLayout)(dstImageLayout),
		(C.uint32_t)(regionCount),
		(*C.VkImageCopy)(pRegions))
}

func (l *vulkanDriver) VkCmdBlitImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkImageBlit, filter VkFilter) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdBlitImage(l.funcPtrs.vkCmdBlitImage,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkImage)(unsafe.Pointer(srcImage)),
		(C.VkImageLayout)(srcImageLayout),
		(C.VkImage)(unsafe.Pointer(dstImage)),
		(C.VkImageLayout)(dstImageLayout),
		(C.uint32_t)(regionCount),
		(*C.VkImageBlit)(pRegions),
		(C.VkFilter)(filter),
	)
}

func (l *vulkanDriver) VkCmdCopyBufferToImage(commandBuffer VkCommandBuffer, srcBuffer VkBuffer, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkBufferImageCopy) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdCopyBufferToImage(l.funcPtrs.vkCmdCopyBufferToImage,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkBuffer)(unsafe.Pointer(srcBuffer)),
		(C.VkImage)(unsafe.Pointer(dstImage)),
		(C.VkImageLayout)(dstImageLayout),
		(C.uint32_t)(regionCount),
		(*C.VkBufferImageCopy)(pRegions))
}

func (l *vulkanDriver) VkCmdCopyImageToBuffer(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstBuffer VkBuffer, regionCount Uint32, pRegions *VkBufferImageCopy) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdCopyImageToBuffer(l.funcPtrs.vkCmdCopyImageToBuffer,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkImage)(unsafe.Pointer(srcImage)),
		(C.VkImageLayout)(srcImageLayout),
		(C.VkBuffer)(unsafe.Pointer(dstBuffer)),
		(C.uint32_t)(regionCount),
		(*C.VkBufferImageCopy)(pRegions))
}

func (l *vulkanDriver) VkCmdUpdateBuffer(commandBuffer VkCommandBuffer, dstBuffer VkBuffer, dstOffset VkDeviceSize, dataSize VkDeviceSize, pData unsafe.Pointer) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdUpdateBuffer(l.funcPtrs.vkCmdUpdateBuffer,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkBuffer)(unsafe.Pointer(dstBuffer)),
		(C.VkDeviceSize)(dstOffset),
		(C.VkDeviceSize)(dataSize),
		pData)
}

func (l *vulkanDriver) VkCmdFillBuffer(commandBuffer VkCommandBuffer, dstBuffer VkBuffer, dstOffset VkDeviceSize, size VkDeviceSize, data Uint32) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdFillBuffer(l.funcPtrs.vkCmdFillBuffer,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkBuffer)(unsafe.Pointer(dstBuffer)),
		(C.VkDeviceSize)(dstOffset),
		(C.VkDeviceSize)(size),
		(C.uint32_t)(data))
}

func (l *vulkanDriver) VkCmdClearColorImage(commandBuffer VkCommandBuffer, image VkImage, imageLayout VkImageLayout, pColor *VkClearColorValue, rangeCount Uint32, pRanges *VkImageSubresourceRange) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdClearColorImage(l.funcPtrs.vkCmdClearColorImage,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkImage)(unsafe.Pointer(image)),
		(C.VkImageLayout)(imageLayout),
		(*C.VkClearColorValue)(pColor),
		(C.uint32_t)(rangeCount),
		(*C.VkImageSubresourceRange)(pRanges))
}

func (l *vulkanDriver) VkCmdClearDepthStencilImage(commandBuffer VkCommandBuffer, image VkImage, imageLayout VkImageLayout, pDepthStencil *VkClearDepthStencilValue, rangeCount Uint32, pRanges *VkImageSubresourceRange) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdClearDepthStencilImage(l.funcPtrs.vkCmdClearDepthStencilImage,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkImage)(unsafe.Pointer(image)),
		(C.VkImageLayout)(imageLayout),
		(*C.VkClearDepthStencilValue)(pDepthStencil),
		(C.uint32_t)(rangeCount),
		(*C.VkImageSubresourceRange)(pRanges))
}

func (l *vulkanDriver) VkCmdClearAttachments(commandBuffer VkCommandBuffer, attachmentCount Uint32, pAttachments *VkClearAttachment, rectCount Uint32, pRects *VkClearRect) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdClearAttachments(l.funcPtrs.vkCmdClearAttachments,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.uint32_t)(attachmentCount),
		(*C.VkClearAttachment)(pAttachments),
		(C.uint32_t)(rectCount),
		(*C.VkClearRect)(pRects))
}

func (l *vulkanDriver) VkCmdResolveImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkImageResolve) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdResolveImage(l.funcPtrs.vkCmdResolveImage,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkImage)(unsafe.Pointer(srcImage)),
		(C.VkImageLayout)(srcImageLayout),
		(C.VkImage)(unsafe.Pointer(dstImage)),
		(C.VkImageLayout)(dstImageLayout),
		(C.uint32_t)(regionCount),
		(*C.VkImageResolve)(pRegions))
}

func (l *vulkanDriver) VkCmdSetEvent(commandBuffer VkCommandBuffer, event VkEvent, stageMask VkPipelineStageFlags) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdSetEvent(l.funcPtrs.vkCmdSetEvent,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkEvent)(unsafe.Pointer(event)),
		(C.VkPipelineStageFlags)(stageMask))
}

func (l *vulkanDriver) VkCmdResetEvent(commandBuffer VkCommandBuffer, event VkEvent, stageMask VkPipelineStageFlags) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdResetEvent(l.funcPtrs.vkCmdResetEvent,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkEvent)(unsafe.Pointer(event)),
		(C.VkPipelineStageFlags)(stageMask))
}

func (l *vulkanDriver) VkCmdWaitEvents(commandBuffer VkCommandBuffer, eventCount Uint32, pEvents *VkEvent, srcStageMask VkPipelineStageFlags, dstStageMask VkPipelineStageFlags, memoryBarrierCount Uint32, pMemoryBarriers *VkMemoryBarrier, bufferMemoryBarrierCount Uint32, pBufferMemoryBarriers *VkBufferMemoryBarrier, imageMemoryBarrierCount Uint32, pImageMemoryBarriers *VkImageMemoryBarrier) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdWaitEvents(l.funcPtrs.vkCmdWaitEvents,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.uint32_t)(eventCount),
		(*C.VkEvent)(unsafe.Pointer(pEvents)),
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
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdPipelineBarrier(l.funcPtrs.vkCmdPipelineBarrier,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
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
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdBeginQuery(l.funcPtrs.vkCmdBeginQuery,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkQueryPool)(unsafe.Pointer(queryPool)),
		(C.uint32_t)(query),
		(C.VkQueryControlFlags)(flags))
}

func (l *vulkanDriver) VkCmdEndQuery(commandBuffer VkCommandBuffer, queryPool VkQueryPool, query Uint32) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdEndQuery(l.funcPtrs.vkCmdEndQuery,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkQueryPool)(unsafe.Pointer(queryPool)),
		(C.uint32_t)(query))
}

func (l *vulkanDriver) VkCmdResetQueryPool(commandBuffer VkCommandBuffer, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdResetQueryPool(l.funcPtrs.vkCmdResetQueryPool,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkQueryPool)(unsafe.Pointer(queryPool)),
		(C.uint32_t)(firstQuery),
		(C.uint32_t)(queryCount))
}

func (l *vulkanDriver) VkCmdWriteTimestamp(commandBuffer VkCommandBuffer, pipelineStage VkPipelineStageFlags, queryPool VkQueryPool, query Uint32) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdWriteTimestamp(l.funcPtrs.vkCmdWriteTimestamp,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkPipelineStageFlagBits)(pipelineStage),
		(C.VkQueryPool)(unsafe.Pointer(queryPool)),
		(C.uint32_t)(query))
}

func (l *vulkanDriver) VkCmdCopyQueryPoolResults(commandBuffer VkCommandBuffer, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32, dstBuffer VkBuffer, dstOffset VkDeviceSize, stride VkDeviceSize, flags VkQueryResultFlags) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdCopyQueryPoolResults(l.funcPtrs.vkCmdCopyQueryPoolResults,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkQueryPool)(unsafe.Pointer(queryPool)),
		(C.uint32_t)(firstQuery),
		(C.uint32_t)(queryCount),
		(C.VkBuffer)(unsafe.Pointer(dstBuffer)),
		(C.VkDeviceSize)(dstOffset),
		(C.VkDeviceSize)(stride),
		(C.VkQueryResultFlags)(flags))
}

func (l *vulkanDriver) VkCmdPushConstants(commandBuffer VkCommandBuffer, layout VkPipelineLayout, stageFlags VkShaderStageFlags, offset Uint32, size Uint32, pValues unsafe.Pointer) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdPushConstants(l.funcPtrs.vkCmdPushConstants,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkPipelineLayout)(unsafe.Pointer(layout)),
		(C.VkShaderStageFlags)(stageFlags),
		(C.uint32_t)(offset),
		(C.uint32_t)(size),
		pValues)
}

func (l *vulkanDriver) VkCmdBeginRenderPass(commandBuffer VkCommandBuffer, pRenderPassBegin *VkRenderPassBeginInfo, contents VkSubpassContents) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdBeginRenderPass(l.funcPtrs.vkCmdBeginRenderPass,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(*C.VkRenderPassBeginInfo)(pRenderPassBegin),
		(C.VkSubpassContents)(contents))
}

func (l *vulkanDriver) VkCmdNextSubpass(commandBuffer VkCommandBuffer, contents VkSubpassContents) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdNextSubpass(l.funcPtrs.vkCmdNextSubpass,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkSubpassContents)(contents))
}

func (l *vulkanDriver) VkCmdEndRenderPass(commandBuffer VkCommandBuffer) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdEndRenderPass(l.funcPtrs.vkCmdEndRenderPass,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)))
}

func (l *vulkanDriver) VkCmdExecuteCommands(commandBuffer VkCommandBuffer, commandBufferCount Uint32, pCommandBuffers *VkCommandBuffer) {
	if VulkanHandle(l.device) == NullHandle {
		panic("attempted device driver function on a non-device driver")
	}

	C.cgoCmdExecuteCommands(l.funcPtrs.vkCmdExecuteCommands,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.uint32_t)(commandBufferCount),
		(*C.VkCommandBuffer)(unsafe.Pointer(pCommandBuffers)))
}

func (l *vulkanDriver) VkEnumerateInstanceVersion(pApiVersion *Uint32) (common.VkResult, error) {
	if l.funcPtrs.vkEnumerateInstanceVersion == nil {
		panic("attempted to call method 'vkEnumerateInstanceVersion' which is not present on this driver")
	}

	res := common.VkResult(C.cgoEnumerateInstanceVersion(l.funcPtrs.vkEnumerateInstanceVersion,
		(*C.uint32_t)(pApiVersion)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkEnumeratePhysicalDeviceGroups(instance VkInstance, pPhysicalDeviceGroupCount *Uint32, pPhysicalDeviceGroupProperties *VkPhysicalDeviceGroupProperties) (common.VkResult, error) {
	if l.funcPtrs.vkEnumeratePhysicalDeviceGroups == nil {
		panic("attempted to call method 'vkEnumeratePhysicalDeviceGroups' which is not present on this driver")
	}

	res := common.VkResult(C.cgoEnumeratePhysicalDeviceGroups(l.funcPtrs.vkEnumeratePhysicalDeviceGroups,
		C.VkInstance(unsafe.Pointer(instance)),
		(*C.uint32_t)(pPhysicalDeviceGroupCount),
		(*C.VkPhysicalDeviceGroupProperties)(pPhysicalDeviceGroupProperties)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkGetPhysicalDeviceFeatures2(physicalDevice VkPhysicalDevice, pFeatures *VkPhysicalDeviceFeatures2) {
	if l.funcPtrs.vkGetPhysicalDeviceFeatures2 == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceFeatures2' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceFeatures2(l.funcPtrs.vkGetPhysicalDeviceFeatures2,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceFeatures2)(pFeatures))
}

func (l *vulkanDriver) VkGetPhysicalDeviceProperties2(physicalDevice VkPhysicalDevice, pProperties *VkPhysicalDeviceProperties2) {
	if l.funcPtrs.vkGetPhysicalDeviceProperties2 == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceProperties2' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceProperties2(l.funcPtrs.vkGetPhysicalDeviceProperties2,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceProperties2)(pProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceFormatProperties2(physicalDevice VkPhysicalDevice, format VkFormat, pFormatProperties *VkFormatProperties2) {
	if l.funcPtrs.vkGetPhysicalDeviceFormatProperties2 == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceFormatProperties2' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceFormatProperties2(l.funcPtrs.vkGetPhysicalDeviceFormatProperties2,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		C.VkFormat(format),
		(*C.VkFormatProperties2)(pFormatProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceImageFormatProperties2(physicalDevice VkPhysicalDevice, pImageFormatInfo *VkPhysicalDeviceImageFormatInfo2, pImageFormatProperties *VkImageFormatProperties2) (common.VkResult, error) {
	if l.funcPtrs.vkGetPhysicalDeviceImageFormatProperties2 == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceImageFormatProperties2' which is not present on this driver")
	}

	res := common.VkResult(C.cgoGetPhysicalDeviceImageFormatProperties2(l.funcPtrs.vkGetPhysicalDeviceImageFormatProperties2,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceImageFormatInfo2)(pImageFormatInfo),
		(*C.VkImageFormatProperties2)(pImageFormatProperties)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkGetPhysicalDeviceQueueFamilyProperties2(physicalDevice VkPhysicalDevice, pQueueFamilyPropertyCount *Uint32, pQueueFamilyProperties *VkQueueFamilyProperties2) {
	if l.funcPtrs.vkGetPhysicalDeviceQueueFamilyProperties2 == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceQueueFamilyProperties2' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceQueueFamilyProperties2(l.funcPtrs.vkGetPhysicalDeviceQueueFamilyProperties2,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.uint32_t)(pQueueFamilyPropertyCount),
		(*C.VkQueueFamilyProperties2)(pQueueFamilyProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceMemoryProperties2(physicalDevice VkPhysicalDevice, pMemoryProperties *VkPhysicalDeviceMemoryProperties2) {
	if l.funcPtrs.vkGetPhysicalDeviceMemoryProperties2 == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceMemoryProperties2' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceMemoryProperties2(l.funcPtrs.vkGetPhysicalDeviceMemoryProperties2,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceMemoryProperties2)(pMemoryProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceSparseImageFormatProperties2(physicalDevice VkPhysicalDevice, pFormatInfo *VkPhysicalDeviceSparseImageFormatInfo2, pPropertyCount *Uint32, pProperties *VkSparseImageFormatProperties2) {
	if l.funcPtrs.vkGetPhysicalDeviceSparseImageFormatProperties2 == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceSparseImageFormatProperties2' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceSparseImageFormatProperties2(l.funcPtrs.vkGetPhysicalDeviceSparseImageFormatProperties2,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceSparseImageFormatInfo2)(pFormatInfo),
		(*C.uint32_t)(pPropertyCount),
		(*C.VkSparseImageFormatProperties2)(pProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceExternalBufferProperties(physicalDevice VkPhysicalDevice, pExternalBufferInfo *VkPhysicalDeviceExternalBufferInfo, pExternalBufferProperties *VkExternalBufferProperties) {
	if l.funcPtrs.vkGetPhysicalDeviceExternalBufferProperties == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceExternalBufferProperties' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceExternalBufferProperties(l.funcPtrs.vkGetPhysicalDeviceExternalBufferProperties,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceExternalBufferInfo)(pExternalBufferInfo),
		(*C.VkExternalBufferProperties)(pExternalBufferProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceExternalFenceProperties(physicalDevice VkPhysicalDevice, pExternalFenceInfo *VkPhysicalDeviceExternalFenceInfo, pExternalFenceProperties *VkExternalFenceProperties) {
	if l.funcPtrs.vkGetPhysicalDeviceExternalFenceProperties == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceExternalFenceProperties' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceExternalFenceProperties(l.funcPtrs.vkGetPhysicalDeviceExternalFenceProperties,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceExternalFenceInfo)(pExternalFenceInfo),
		(*C.VkExternalFenceProperties)(pExternalFenceProperties))
}

func (l *vulkanDriver) VkGetPhysicalDeviceExternalSemaphoreProperties(physicalDevice VkPhysicalDevice, pExternalSemaphoreInfo *VkPhysicalDeviceExternalSemaphoreInfo, pExternalSemaphoreProperties *VkExternalSemaphoreProperties) {
	if l.funcPtrs.vkGetPhysicalDeviceExternalSemaphoreProperties == nil {
		panic("attempted to call method 'vkGetPhysicalDeviceExternalSemaphoreProperties' which is not present on this driver")
	}

	C.cgoGetPhysicalDeviceExternalSemaphoreProperties(l.funcPtrs.vkGetPhysicalDeviceExternalSemaphoreProperties,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceExternalSemaphoreInfo)(pExternalSemaphoreInfo),
		(*C.VkExternalSemaphoreProperties)(pExternalSemaphoreProperties))
}

func (l *vulkanDriver) VkBindBufferMemory2(device VkDevice, bindInfoCount Uint32, pBindInfos *VkBindBufferMemoryInfo) (common.VkResult, error) {
	if l.funcPtrs.vkBindBufferMemory2 == nil {
		panic("attempted to call method 'vkBindBufferMemory2' which is not present on this driver")
	}

	res := common.VkResult(C.cgoBindBufferMemory2(l.funcPtrs.vkBindBufferMemory2,
		C.VkDevice(unsafe.Pointer(device)),
		C.uint32_t(bindInfoCount),
		(*C.VkBindBufferMemoryInfo)(pBindInfos)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkBindImageMemory2(device VkDevice, bindInfoCount Uint32, pBindInfos *VkBindImageMemoryInfo) (common.VkResult, error) {
	if l.funcPtrs.vkBindImageMemory2 == nil {
		panic("attempted to call method 'vkBindImageMemory2' which is not present on this driver")
	}

	res := common.VkResult(C.cgoBindImageMemory2(l.funcPtrs.vkBindImageMemory2,
		C.VkDevice(unsafe.Pointer(device)),
		C.uint32_t(bindInfoCount),
		(*C.VkBindImageMemoryInfo)(pBindInfos)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkGetDeviceGroupPeerMemoryFeatures(device VkDevice, heapIndex Uint32, localDeviceIndex Uint32, remoteDeviceIndex Uint32, pPeerMemoryFeatures *VkPeerMemoryFeatureFlags) {
	if l.funcPtrs.vkGetDeviceGroupPeerMemoryFeatures == nil {
		panic("attempted to call method 'vkGetDeviceGroupPeerMemoryFeatures' which is not present on this driver")
	}

	C.cgoGetDeviceGroupPeerMemoryFeatures(l.funcPtrs.vkGetDeviceGroupPeerMemoryFeatures,
		C.VkDevice(unsafe.Pointer(device)),
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
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		C.uint32_t(deviceMask))
}

func (l *vulkanDriver) VkCmdDispatchBase(commandBuffer VkCommandBuffer, baseGroupX Uint32, baseGroupY Uint32, baseGroupZ Uint32, groupCountX Uint32, groupCountY Uint32, groupCountZ Uint32) {
	if l.funcPtrs.vkCmdDispatchBase == nil {
		panic("attempted to call method 'vkCmdDispatchBase' which is not present on this driver")
	}

	C.cgoCmdDispatchBase(l.funcPtrs.vkCmdDispatchBase,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
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
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkImageMemoryRequirementsInfo2)(pInfo),
		(*C.VkMemoryRequirements2)(pMemoryRequirements))
}

func (l *vulkanDriver) VkGetBufferMemoryRequirements2(device VkDevice, pInfo *VkBufferMemoryRequirementsInfo2, pMemoryRequirements *VkMemoryRequirements2) {
	if l.funcPtrs.vkGetBufferMemoryRequirements2 == nil {
		panic("attempted to call method 'vkGetBufferMemoryRequirements2' which is not present on this driver")
	}

	C.cgoGetBufferMemoryRequirements2(l.funcPtrs.vkGetBufferMemoryRequirements2,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkBufferMemoryRequirementsInfo2)(pInfo),
		(*C.VkMemoryRequirements2)(pMemoryRequirements))
}

func (l *vulkanDriver) VkGetImageSparseMemoryRequirements2(device VkDevice, pInfo *VkImageSparseMemoryRequirementsInfo2, pSparseMemoryRequirementCount *Uint32, pSparseMemoryRequirements *VkSparseImageMemoryRequirements2) {
	if l.funcPtrs.vkGetImageSparseMemoryRequirements2 == nil {
		panic("attempted to call method 'vkGetImageSparseMemoryRequirements2' which is not present on this driver")
	}

	C.cgoGetImageSparseMemoryRequirements2(l.funcPtrs.vkGetImageSparseMemoryRequirements2,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkImageSparseMemoryRequirementsInfo2)(pInfo),
		(*C.uint32_t)(pSparseMemoryRequirementCount),
		(*C.VkSparseImageMemoryRequirements2)(pSparseMemoryRequirements))
}

func (l *vulkanDriver) VkTrimCommandPool(device VkDevice, commandPool VkCommandPool, flags VkCommandPoolTrimFlags) {
	if l.funcPtrs.vkTrimCommandPool == nil {
		panic("attempted to call method 'vkTrimCommandPool' which is not present on this driver")
	}

	C.cgoTrimCommandPool(l.funcPtrs.vkTrimCommandPool,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkCommandPool(unsafe.Pointer(commandPool)),
		C.VkCommandPoolTrimFlags(flags))
}

func (l *vulkanDriver) VkGetDeviceQueue2(device VkDevice, pQueueInfo *VkDeviceQueueInfo2, pQueue *VkQueue) {
	if l.funcPtrs.vkGetDeviceQueue2 == nil {
		panic("attempted to call method 'vkGetDeviceQueue2' which is not present on this driver")
	}

	C.cgoGetDeviceQueue2(l.funcPtrs.vkGetDeviceQueue2,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkDeviceQueueInfo2)(pQueueInfo),
		(*C.VkQueue)(unsafe.Pointer(pQueue)))
}

func (l *vulkanDriver) VkCreateSamplerYcbcrConversion(device VkDevice, pCreateInfo *VkSamplerYcbcrConversionCreateInfo, pAllocator *VkAllocationCallbacks, pYcbcrConversion *VkSamplerYcbcrConversion) (common.VkResult, error) {
	if l.funcPtrs.vkCreateSamplerYcbcrConversion == nil {
		panic("attempted to call method 'vkCreateSamplerYcbcrConversion' which is not present on this driver")
	}

	res := common.VkResult(C.cgoCreateSamplerYcbcrConversion(l.funcPtrs.vkCreateSamplerYcbcrConversion,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkSamplerYcbcrConversionCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkSamplerYcbcrConversion)(unsafe.Pointer(pYcbcrConversion))))

	return res, res.ToError()
}

func (l *vulkanDriver) VkDestroySamplerYcbcrConversion(device VkDevice, ycbcrConversion VkSamplerYcbcrConversion, pAllocator *VkAllocationCallbacks) {
	if l.funcPtrs.vkDestroySamplerYcbcrConversion == nil {
		panic("attempted to call method 'vkDestroySamplerYcbcrConversion' which is not present on this driver")
	}

	C.cgoDestroySamplerYcbcrConversion(l.funcPtrs.vkDestroySamplerYcbcrConversion,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSamplerYcbcrConversion(unsafe.Pointer(ycbcrConversion)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkCreateDescriptorUpdateTemplate(device VkDevice, pCreateInfo *VkDescriptorUpdateTemplateCreateInfo, pAllocator *VkAllocationCallbacks, pDescriptorUpdateTemplate *VkDescriptorUpdateTemplate) (common.VkResult, error) {
	if l.funcPtrs.vkCreateDescriptorUpdateTemplate == nil {
		panic("attempted to call method 'vkCreateDescriptorUpdateTemplate' which is not present on this driver")
	}

	res := common.VkResult(C.cgoCreateDescriptorUpdateTemplate(l.funcPtrs.vkCreateDescriptorUpdateTemplate,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkDescriptorUpdateTemplateCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkDescriptorUpdateTemplate)(unsafe.Pointer(pDescriptorUpdateTemplate))))

	return res, res.ToError()
}
func (l *vulkanDriver) VkDestroyDescriptorUpdateTemplate(device VkDevice, descriptorUpdateTemplate VkDescriptorUpdateTemplate, pAllocator *VkAllocationCallbacks) {
	if l.funcPtrs.vkDestroyDescriptorUpdateTemplate == nil {
		panic("attempted to call method 'vkDestroyDescriptorUpdateTemplate' which is not present on this driver")
	}

	C.cgoDestroyDescriptorUpdateTemplate(l.funcPtrs.vkDestroyDescriptorUpdateTemplate,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkDescriptorUpdateTemplate(unsafe.Pointer(descriptorUpdateTemplate)),
		(*C.VkAllocationCallbacks)(pAllocator))
}

func (l *vulkanDriver) VkUpdateDescriptorSetWithTemplate(device VkDevice, descriptorSet VkDescriptorSet, descriptorUpdateTemplate VkDescriptorUpdateTemplate, pData unsafe.Pointer) {
	if l.funcPtrs.vkUpdateDescriptorSetWithTemplate == nil {
		panic("attempted to call method 'vkUpdateDescriptorSetWithTemplate' which is not present on this driver")
	}

	C.cgoUpdateDescriptorSetWithTemplate(l.funcPtrs.vkUpdateDescriptorSetWithTemplate,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkDescriptorSet(unsafe.Pointer(descriptorSet)),
		C.VkDescriptorUpdateTemplate(unsafe.Pointer(descriptorUpdateTemplate)),
		pData)
}

func (l *vulkanDriver) VkGetDescriptorSetLayoutSupport(device VkDevice, pCreateInfo *VkDescriptorSetLayoutCreateInfo, pSupport *VkDescriptorSetLayoutSupport) {
	if l.funcPtrs.vkGetDescriptorSetLayoutSupport == nil {
		panic("attempted to call method 'vkGetDescriptorSetLayoutSupport' which is not present on this driver")
	}

	C.cgoGetDescriptorSetLayoutSupport(l.funcPtrs.vkGetDescriptorSetLayoutSupport,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkDescriptorSetLayoutCreateInfo)(pCreateInfo),
		(*C.VkDescriptorSetLayoutSupport)(pSupport))
}

func (l *vulkanDriver) VkCmdDrawIndirectCount(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, countBuffer VkBuffer, countBufferOffset VkDeviceSize, maxDrawCount Uint32, stride Uint32) {
	if l.funcPtrs.vkCmdDrawIndirectCount == nil {
		panic("attempted to call method 'vkCmdDrawIndirectCount' which is not present on this driver")
	}

	C.cgoCmdDrawIndirectCount(l.funcPtrs.vkCmdDrawIndirectCount,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		C.VkBuffer(unsafe.Pointer(buffer)),
		C.VkDeviceSize(offset),
		C.VkBuffer(unsafe.Pointer(countBuffer)),
		C.VkDeviceSize(countBufferOffset),
		C.uint32_t(maxDrawCount),
		C.uint32_t(stride))
}

func (l *vulkanDriver) VkCmdDrawIndexedIndirectCount(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, countBuffer VkBuffer, countBufferOffset VkDeviceSize, maxDrawCount Uint32, stride Uint32) {
	if l.funcPtrs.vkCmdDrawIndexedIndirectCount == nil {
		panic("attempted to call method 'vkCmdDrawIndexedIndirectCount' which is not present on this driver")
	}

	C.cgoCmdDrawIndexedIndirectCount(l.funcPtrs.vkCmdDrawIndexedIndirectCount,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		C.VkBuffer(unsafe.Pointer(buffer)),
		C.VkDeviceSize(offset),
		C.VkBuffer(unsafe.Pointer(countBuffer)),
		C.VkDeviceSize(countBufferOffset),
		C.uint32_t(maxDrawCount),
		C.uint32_t(stride))
}

func (l *vulkanDriver) VkCreateRenderPass2(device VkDevice, pCreateInfo *VkRenderPassCreateInfo2, pAllocator *VkAllocationCallbacks, pRenderPass *VkRenderPass) (common.VkResult, error) {
	if l.funcPtrs.vkCreateRenderPass2 == nil {
		panic("attempted to call method 'vkCreateRenderPass2' which is not present on this driver")
	}

	res := common.VkResult(C.cgoCreateRenderPass2(l.funcPtrs.vkCreateRenderPass2,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkRenderPassCreateInfo2)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkRenderPass)(unsafe.Pointer(pRenderPass))))

	return res, res.ToError()
}

func (l *vulkanDriver) VkCmdBeginRenderPass2(commandBuffer VkCommandBuffer, pRenderPassBegin *VkRenderPassBeginInfo, pSubpassBeginInfo *VkSubpassBeginInfo) {
	if l.funcPtrs.vkCmdBeginRenderPass2 == nil {
		panic("attempted to call method 'vkCmdBeginRenderPass2' which is not present on this driver")
	}

	C.cgoCmdBeginRenderPass2(l.funcPtrs.vkCmdBeginRenderPass2,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkRenderPassBeginInfo)(pRenderPassBegin),
		(*C.VkSubpassBeginInfo)(pSubpassBeginInfo))
}

func (l *vulkanDriver) VkCmdNextSubpass2(commandBuffer VkCommandBuffer, pSubpassBeginInfo *VkSubpassBeginInfo, pSubpassEndInfo *VkSubpassEndInfo) {
	if l.funcPtrs.vkCmdNextSubpass2 == nil {
		panic("attempted to call method 'vkCmdNextSubpass2' which is not present on this driver")
	}

	C.cgoCmdNextSubpass2(l.funcPtrs.vkCmdNextSubpass2,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkSubpassBeginInfo)(pSubpassBeginInfo),
		(*C.VkSubpassEndInfo)(pSubpassEndInfo))
}

func (l *vulkanDriver) VkCmdEndRenderPass2(commandBuffer VkCommandBuffer, pSubpassEndInfo *VkSubpassEndInfo) {
	if l.funcPtrs.vkCmdEndRenderPass2 == nil {
		panic("attempted to call method 'vkCmdEndRenderPass2' which is not present on this driver")
	}

	C.cgoCmdEndRenderPass2(l.funcPtrs.vkCmdEndRenderPass2,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkSubpassEndInfo)(pSubpassEndInfo))
}

func (l *vulkanDriver) VkResetQueryPool(device VkDevice, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32) {
	if l.funcPtrs.vkResetQueryPool == nil {
		panic("attempted to call method 'vkResetQueryPool' which is not present on this driver")
	}

	C.cgoResetQueryPool(l.funcPtrs.vkResetQueryPool,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkQueryPool(unsafe.Pointer(queryPool)),
		C.uint32_t(firstQuery),
		C.uint32_t(queryCount))
}

func (l *vulkanDriver) VkGetSemaphoreCounterValue(device VkDevice, semaphore VkSemaphore, pValue *Uint64) (common.VkResult, error) {
	if l.funcPtrs.vkGetSemaphoreCounterValue == nil {
		panic("attempted to call method 'vkGetSemaphoreCounterValue' which is not present on this driver")
	}

	res := common.VkResult(C.cgoGetSemaphoreCounterValue(l.funcPtrs.vkGetSemaphoreCounterValue,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSemaphore(unsafe.Pointer(semaphore)),
		(*C.uint64_t)(pValue)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkWaitSemaphores(device VkDevice, pWaitInfo *VkSemaphoreWaitInfo, timeout Uint64) (common.VkResult, error) {
	if l.funcPtrs.vkWaitSemaphores == nil {
		panic("attempted to call method 'vkWaitSemaphores' which is not present on this driver")
	}

	res := common.VkResult(C.cgoWaitSemaphores(l.funcPtrs.vkWaitSemaphores,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkSemaphoreWaitInfo)(pWaitInfo),
		C.uint64_t(timeout)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkSignalSemaphore(device VkDevice, pSignalInfo *VkSemaphoreSignalInfo) (common.VkResult, error) {
	if l.funcPtrs.vkSignalSemaphore == nil {
		panic("attempted to call method 'vkSignalSemaphore' which is not present on this driver")
	}

	res := common.VkResult(C.cgoSignalSemaphore(l.funcPtrs.vkSignalSemaphore,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkSemaphoreSignalInfo)(pSignalInfo)))

	return res, res.ToError()
}

func (l *vulkanDriver) VkGetBufferDeviceAddress(device VkDevice, pInfo *VkBufferDeviceAddressInfo) VkDeviceAddress {
	if l.funcPtrs.vkGetBufferDeviceAddress == nil {
		panic("attempted to call method 'vkGetBufferDeviceAddress' which is not present on this driver")
	}

	address := VkDeviceAddress(C.cgoGetBufferDeviceAddress(l.funcPtrs.vkGetBufferDeviceAddress,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkBufferDeviceAddressInfo)(pInfo)))

	return address
}

func (l *vulkanDriver) VkGetBufferOpaqueCaptureAddress(device VkDevice, pInfo *VkBufferDeviceAddressInfo) Uint64 {
	if l.funcPtrs.vkGetBufferOpaqueCaptureAddress == nil {
		panic("attempted to call method 'vkGetBufferOpaqueCaptureAddress' which is not present on this driver")
	}

	address := Uint64(C.cgoGetBufferOpaqueCaptureAddress(l.funcPtrs.vkGetBufferOpaqueCaptureAddress,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkBufferDeviceAddressInfo)(pInfo)))

	return address
}

func (l *vulkanDriver) VkGetDeviceMemoryOpaqueCaptureAddress(device VkDevice, pInfo *VkDeviceMemoryOpaqueCaptureAddressInfo) Uint64 {
	if l.funcPtrs.vkGetDeviceMemoryOpaqueCaptureAddress == nil {
		panic("attempted to call method 'vkGetDeviceMemoryOpaqueCaptureAddress' which is not present on this driver")
	}

	address := Uint64(C.cgoGetDeviceMemoryOpaqueCaptureAddress(l.funcPtrs.vkGetDeviceMemoryOpaqueCaptureAddress,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkDeviceMemoryOpaqueCaptureAddressInfo)(pInfo)))

	return address
}
