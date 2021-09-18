package loader

/*
#cgo windows LDFLAGS: -lvulkan
#cgo linux freebsd openbsd pkg-config: vulkan
#include "loader.h"
*/
import "C"
import (
	"github.com/cockroachdb/errors"
	"unsafe"
)

func (l *VulkanLoader) VkEnumerateInstanceExtensionProperties(pLayerName *Char, pPropertyCount *Uint32, pProperties *VkExtensionProperties) (VkResult, error) {
	res := VkResult(C.cgoEnumerateInstanceExtensionProperties(l.funcPtrs.vkEnumerateInstanceExtensionProperties,
		(*C.char)(pLayerName),
		(*C.uint32_t)(pPropertyCount),
		(*C.VkExtensionProperties)(pProperties)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkEnumerateInstanceLayerProperties(pPropertyCount *Uint32, pProperties *VkLayerProperties) (VkResult, error) {
	res := VkResult(C.cgoEnumerateInstanceLayerProperties(l.funcPtrs.vkEnumerateInstanceLayerProperties,
		(*C.uint32_t)(pPropertyCount),
		(*C.VkLayerProperties)(pProperties)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkCreateInstance(pCreateInfo *VkInstanceCreateInfo, pAllocator *VkAllocationCallbacks, pInstance *VkInstance) (VkResult, error) {
	res := VkResult(C.cgoCreateInstance(l.funcPtrs.vkCreateInstance,
		(*C.VkInstanceCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkInstance)(pInstance)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkEnumeratePhysicalDevices(instance VkInstance, pPhysicalDeviceCount *Uint32, pPhysicalDevices *VkPhysicalDevice) (VkResult, error) {
	if l.instance == nil {
		return VKErrorUnknown, errors.New("attempted to call instance loader function on a basic loader")
	}

	res := VkResult(C.cgoEnumeratePhysicalDevices(l.funcPtrs.vkEnumeratePhysicalDevices,
		(C.VkInstance)(instance),
		(*C.uint32_t)(pPhysicalDeviceCount),
		(*C.VkPhysicalDevice)(pPhysicalDevices)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroyInstance(instance VkInstance, pAllocator *VkAllocationCallbacks) error {
	if l.instance == nil {
		return errors.New("attempted to call instance loader function on a basic loader")
	}

	C.cgoDestroyInstance(l.funcPtrs.vkDestroyInstance,
		(C.VkInstance)(instance),
		(*C.VkAllocationCallbacks)(pAllocator))
	return nil
}

func (l *VulkanLoader) VkGetPhysicalDeviceFeatures(physicalDevice VkPhysicalDevice, pFeatures *VkPhysicalDeviceFeatures) error {
	if l.instance == nil {
		return errors.New("attempted to call instance loader function on a basic loader")
	}

	C.cgoGetPhysicalDeviceFeatures(l.funcPtrs.vkGetPhysicalDeviceFeatures,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.VkPhysicalDeviceFeatures)(pFeatures))
	return nil
}

func (l *VulkanLoader) VkGetPhysicalDeviceFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, pFormatProperties *VkFormatProperties) error {
	if l.instance == nil {
		return errors.New("attempted to call instance loader function on a basic loader")
	}

	C.cgoGetPhysicalDeviceFormatProperties(l.funcPtrs.vkGetPhysicalDeviceFormatProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(C.VkFormat)(format),
		(*C.VkFormatProperties)(pFormatProperties))
	return nil
}

func (l *VulkanLoader) VkGetPhysicalDeviceImageFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, t VkImageType, tiling VkImageTiling, usage VkImageUsageFlags, flags VkImageCreateFlags, pImageFormatProperties *VkImageFormatProperties) (VkResult, error) {
	if l.instance == nil {
		return VKErrorUnknown, errors.New("attempted to call instance loader function on a basic loader")
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

func (l *VulkanLoader) VkGetPhysicalDeviceProperties(physicalDevice VkPhysicalDevice, pProperties *VkPhysicalDeviceProperties) error {
	if l.instance == nil {
		return errors.New("attempted to call instance loader function on a basic loader")
	}

	C.cgoGetPhysicalDeviceProperties(l.funcPtrs.vkGetPhysicalDeviceProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.VkPhysicalDeviceProperties)(pProperties))
	return nil
}

func (l *VulkanLoader) VkGetPhysicalDeviceQueueFamilyProperties(physicalDevice VkPhysicalDevice, pQueueFamilyPropertyCount *Uint32, pQueueFamilyProperties *VkQueueFamilyProperties) error {
	if l.instance == nil {
		return errors.New("attempted to call instance loader function on a basic loader")
	}

	C.cgoGetPhysicalDeviceQueueFamilyProperties(l.funcPtrs.vkGetPhysicalDeviceQueueFamilyProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.uint32_t)(pQueueFamilyPropertyCount),
		(*C.VkQueueFamilyProperties)(pQueueFamilyProperties))
	return nil
}

func (l *VulkanLoader) VkGetPhysicalDeviceMemoryProperties(physicalDevice VkPhysicalDevice, pMemoryProperties *VkPhysicalDeviceMemoryProperties) error {
	if l.instance == nil {
		return errors.New("attempted to call instance loader function on a basic loader")
	}

	C.cgoGetPhysicalDeviceMemoryProperties(l.funcPtrs.vkGetPhysicalDeviceMemoryProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.VkPhysicalDeviceMemoryProperties)(pMemoryProperties))
	return nil
}

func (l *VulkanLoader) VkEnumerateDeviceExtensionProperties(physicalDevice VkPhysicalDevice, pLayerName *Char, pPropertyCount *Uint32, pProperties *VkExtensionProperties) (VkResult, error) {
	if l.instance == nil {
		return VKErrorUnknown, errors.New("attempted to call instance loader function on a basic loader")
	}

	res := VkResult(C.cgoEnumerateDeviceExtensionProperties(l.funcPtrs.vkEnumerateDeviceExtensionProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.char)(pLayerName),
		(*C.uint32_t)(pPropertyCount),
		(*C.VkExtensionProperties)(pProperties)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkEnumerateDeviceLayerProperties(physicalDevice VkPhysicalDevice, pPropertyCount *Uint32, pProperties *VkLayerProperties) (VkResult, error) {
	if l.instance == nil {
		return VKErrorUnknown, errors.New("attempted to call instance loader function on a basic loader")
	}

	res := VkResult(C.cgoEnumerateDeviceLayerProperties(l.funcPtrs.vkEnumerateDeviceLayerProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.uint32_t)(pPropertyCount),
		(*C.VkLayerProperties)(pProperties)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkGetPhysicalDeviceSparseImageFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, t VkImageType, samples VkSampleCountFlagBits, usage VkImageUsageFlags, tiling VkImageTiling, pPropertyCount *Uint32, pProperties *VkSparseImageFormatProperties) error {
	if l.instance == nil {
		return errors.New("attempted to call instance loader function on a basic loader")
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
	return nil
}

func (l *VulkanLoader) VkCreateDevice(physicalDevice VkPhysicalDevice, pCreateInfo *VkDeviceCreateInfo, pAllocator *VkAllocationCallbacks, pDevice *VkDevice) (VkResult, error) {
	if l.instance == nil {
		return VKErrorUnknown, errors.New("attempted to call instance loader function on a basic loader")
	}

	res := VkResult(C.cgoCreateDevice(l.funcPtrs.vkCreateDevice,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.VkDeviceCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkDevice)(pDevice)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroyDevice(device VkDevice, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyDevice(l.funcPtrs.vkDestroyDevice,
		(C.VkDevice)(device),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkGetDeviceQueue(device VkDevice, queueFamilyIndex Uint32, queueIndex Uint32, pQueue *VkQueue) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoGetDeviceQueue(l.funcPtrs.vkGetDeviceQueue,
		(C.VkDevice)(device),
		(C.uint32_t)(queueFamilyIndex),
		(C.uint32_t)(queueIndex),
		(*C.VkQueue)(pQueue))

	return nil
}

func (l *VulkanLoader) VkQueueSubmit(queue VkQueue, submitCount Uint32, pSubmits *VkSubmitInfo, fence VkFence) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoQueueSubmit(l.funcPtrs.vkQueueSubmit,
		(C.VkQueue)(queue),
		(C.uint32_t)(submitCount),
		(*C.VkSubmitInfo)(pSubmits),
		(C.VkFence)(fence)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkQueueWaitIdle(queue VkQueue) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoQueueWaitIdle(l.funcPtrs.vkQueueWaitIdle,
		(C.VkQueue)(queue)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDeviceWaitIdle(device VkDevice) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoDeviceWaitIdle(l.funcPtrs.vkDeviceWaitIdle,
		(C.VkDevice)(device)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkAllocateMemory(device VkDevice, pAllocateInfo *VkMemoryAllocateInfo, pAllocator *VkAllocationCallbacks, pMemory *VkDeviceMemory) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoAllocateMemory(l.funcPtrs.vkAllocateMemory,
		(C.VkDevice)(device),
		(*C.VkMemoryAllocateInfo)(pAllocateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkDeviceMemory)(pMemory)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkFreeMemory(device VkDevice, memory VkDeviceMemory, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoFreeMemory(l.funcPtrs.vkFreeMemory,
		(C.VkDevice)(device),
		(C.VkDeviceMemory)(memory),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkMapMemory(device VkDevice, memory VkDeviceMemory, offset VkDeviceSize, size VkDeviceSize, flags VkMemoryMapFlags, ppData *unsafe.Pointer) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
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

func (l *VulkanLoader) VkUnmapMemory(device VkDevice, memory VkDeviceMemory) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoUnmapMemory(l.funcPtrs.vkUnmapMemory,
		(C.VkDevice)(device),
		(C.VkDeviceMemory)(memory))

	return nil
}

func (l *VulkanLoader) VkFlushMappedMemoryRanges(device VkDevice, memoryRangeCount Uint32, pMemoryRanges *VkMappedMemoryRange) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoFlushMappedMemoryRanges(l.funcPtrs.vkFlushMappedMemoryRanges,
		(C.VkDevice)(device),
		(C.uint32_t)(memoryRangeCount),
		(*C.VkMappedMemoryRange)(pMemoryRanges)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkInvalidateMappedMemoryRanges(device VkDevice, memoryRangeCount Uint32, pMemoryRanges *VkMappedMemoryRange) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoInvalidateMappedMemoryRanges(l.funcPtrs.vkInvalidateMappedMemoryRanges,
		(C.VkDevice)(device),
		(C.uint32_t)(memoryRangeCount),
		(*C.VkMappedMemoryRange)(pMemoryRanges)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkGetDeviceMemoryCommitment(device VkDevice, memory VkDeviceMemory, pCommittedMemoryInBytes *VkDeviceSize) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoGetDeviceMemoryCommitment(l.funcPtrs.vkGetDeviceMemoryCommitment,
		(C.VkDevice)(device),
		(C.VkDeviceMemory)(memory),
		(*C.VkDeviceSize)(pCommittedMemoryInBytes))

	return nil
}

func (l *VulkanLoader) VkBindBufferMemory(device VkDevice, buffer VkBuffer, memory VkDeviceMemory, memoryOffset VkDeviceSize) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoBindBufferMemory(l.funcPtrs.vkBindBufferMemory,
		(C.VkDevice)(device),
		(C.VkBuffer)(buffer),
		(C.VkDeviceMemory)(memory),
		(C.VkDeviceSize)(memoryOffset)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkBindImageMemory(device VkDevice, image VkImage, memory VkDeviceMemory, memoryOffset VkDeviceSize) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoBindImageMemory(l.funcPtrs.vkBindImageMemory,
		(C.VkDevice)(device),
		(C.VkImage)(image),
		(C.VkDeviceMemory)(memory),
		(C.VkDeviceSize)(memoryOffset)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkGetBufferMemoryRequirements(device VkDevice, buffer VkBuffer, pMemoryRequirements *VkMemoryRequirements) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoGetBufferMemoryRequirements(l.funcPtrs.vkGetBufferMemoryRequirements,
		(C.VkDevice)(device),
		(C.VkBuffer)(buffer),
		(*C.VkMemoryRequirements)(pMemoryRequirements))

	return nil
}

func (l *VulkanLoader) VkGetImageMemoryRequirements(device VkDevice, image VkImage, pMemoryRequirements *VkMemoryRequirements) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoGetImageMemoryRequirements(l.funcPtrs.vkGetImageMemoryRequirements,
		(C.VkDevice)(device),
		(C.VkImage)(image),
		(*C.VkMemoryRequirements)(pMemoryRequirements))

	return nil
}

func (l *VulkanLoader) VkGetImageSparseMemoryRequirements(device VkDevice, image VkImage, pSparseMemoryRequirementCount *Uint32, pSparseMemoryRequirements *VkSparseImageMemoryRequirements) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoGetImageSparseMemoryRequirements(l.funcPtrs.vkGetImageSparseMemoryRequirements,
		(C.VkDevice)(device),
		(C.VkImage)(image),
		(*C.uint32_t)(pSparseMemoryRequirementCount),
		(*C.VkSparseImageMemoryRequirements)(pSparseMemoryRequirements))

	return nil
}

func (l *VulkanLoader) VkQueueBindSparse(queue VkQueue, bindInfoCount Uint32, pBindInfo *VkBindSparseInfo, fence VkFence) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoQueueBindSparse(l.funcPtrs.vkQueueBindSparse,
		(C.VkQueue)(queue),
		(C.uint32_t)(bindInfoCount),
		(*C.VkBindSparseInfo)(pBindInfo),
		(C.VkFence)(fence)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkCreateFence(device VkDevice, pCreateInfo *VkFenceCreateInfo, pAllocator *VkAllocationCallbacks, pFence *VkFence) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoCreateFence(l.funcPtrs.vkCreateFence,
		(C.VkDevice)(device),
		(*C.VkFenceCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkFence)(pFence)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroyFence(device VkDevice, fence VkFence, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyFence(l.funcPtrs.vkDestroyFence,
		(C.VkDevice)(device),
		(C.VkFence)(fence),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkResetFences(device VkDevice, fenceCount Uint32, pFences *VkFence) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoResetFences(l.funcPtrs.vkResetFences,
		(C.VkDevice)(device),
		(C.uint32_t)(fenceCount),
		(*C.VkFence)(pFences)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkGetFenceStatus(device VkDevice, fence VkFence) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoGetFenceStatus(l.funcPtrs.vkGetFenceStatus,
		(C.VkDevice)(device),
		(C.VkFence)(fence)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkWaitForFences(device VkDevice, fenceCount Uint32, pFences *VkFence, waitAll VkBool32, timeout Uint64) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoWaitForFences(l.funcPtrs.vkWaitForFences,
		(C.VkDevice)(device),
		(C.uint32_t)(fenceCount),
		(*C.VkFence)(pFences),
		(C.VkBool32)(waitAll),
		(C.uint64_t)(timeout)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkCreateSemaphore(device VkDevice, pCreateInfo *VkSemaphoreCreateInfo, pAllocator *VkAllocationCallbacks, pSemaphore *VkSemaphore) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoCreateSemaphore(l.funcPtrs.vkCreateSemaphore,
		(C.VkDevice)(device),
		(*C.VkSemaphoreCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkSemaphore)(pSemaphore)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroySemaphore(device VkDevice, semaphore VkSemaphore, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroySemaphore(l.funcPtrs.vkDestroySemaphore,
		(C.VkDevice)(device),
		(C.VkSemaphore)(semaphore),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkCreateEvent(device VkDevice, pCreateInfo *VkEventCreateInfo, pAllocator *VkAllocationCallbacks, pEvent *VkEvent) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoCreateEvent(l.funcPtrs.vkCreateEvent,
		(C.VkDevice)(device),
		(*C.VkEventCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkEvent)(pEvent)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroyEvent(device VkDevice, event VkEvent, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyEvent(l.funcPtrs.vkDestroyEvent,
		(C.VkDevice)(device),
		(C.VkEvent)(event),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkGetEventStatus(device VkDevice, event VkEvent) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoGetEventStatus(l.funcPtrs.vkGetEventStatus,
		(C.VkDevice)(device),
		(C.VkEvent)(event)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkSetEvent(device VkDevice, event VkEvent) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoSetEvent(l.funcPtrs.vkSetEvent,
		(C.VkDevice)(device),
		(C.VkEvent)(event)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkResetEvent(device VkDevice, event VkEvent) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoResetEvent(l.funcPtrs.vkResetEvent,
		(C.VkDevice)(device),
		(C.VkEvent)(event)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkCreateQueryPool(device VkDevice, pCreateInfo *VkQueryPoolCreateInfo, pAllocator *VkAllocationCallbacks, pQueryPool *VkQueryPool) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoCreateQueryPool(l.funcPtrs.vkCreateQueryPool,
		(C.VkDevice)(device),
		(*C.VkQueryPoolCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkQueryPool)(pQueryPool)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroyQueryPool(device VkDevice, queryPool VkQueryPool, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyQueryPool(l.funcPtrs.vkDestroyQueryPool,
		(C.VkDevice)(device),
		(C.VkQueryPool)(queryPool),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkGetQueryPoolResults(device VkDevice, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32, dataSize Size, pData unsafe.Pointer, stride VkDeviceSize, flags VkQueryResultFlags) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
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

func (l *VulkanLoader) VkCreateBuffer(device VkDevice, pCreateInfo *VkBufferCreateInfo, pAllocator *VkAllocationCallbacks, pBuffer *VkBuffer) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoCreateBuffer(l.funcPtrs.vkCreateBuffer,
		(C.VkDevice)(device),
		(*C.VkBufferCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkBuffer)(pBuffer)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroyBuffer(device VkDevice, buffer VkBuffer, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyBuffer(l.funcPtrs.vkDestroyBuffer,
		(C.VkDevice)(device),
		(C.VkBuffer)(buffer),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkCreateBufferView(device VkDevice, pCreateInfo *VkBufferViewCreateInfo, pAllocator *VkAllocationCallbacks, pView *VkBufferView) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoCreateBufferView(l.funcPtrs.vkCreateBufferView,
		(C.VkDevice)(device),
		(*C.VkBufferViewCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkBufferView)(pView)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroyBufferView(device VkDevice, bufferView VkBufferView, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyBufferView(l.funcPtrs.vkDestroyBufferView,
		(C.VkDevice)(device),
		(C.VkBufferView)(bufferView),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkCreateImage(device VkDevice, pCreateInfo *VkImageCreateInfo, pAllocator *VkAllocationCallbacks, pImage *VkImage) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoCreateImage(l.funcPtrs.vkCreateImage,
		(C.VkDevice)(device),
		(*C.VkImageCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkImage)(pImage)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroyImage(device VkDevice, image VkImage, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyImage(l.funcPtrs.vkDestroyImage,
		(C.VkDevice)(device),
		(C.VkImage)(image),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkGetImageSubresourceLayout(device VkDevice, image VkImage, pSubresource *VkImageSubresource, pLayout *VkSubresourceLayout) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoGetImageSubresourceLayout(l.funcPtrs.vkGetImageSubresourceLayout,
		(C.VkDevice)(device),
		(C.VkImage)(image),
		(*C.VkImageSubresource)(pSubresource),
		(*C.VkSubresourceLayout)(pLayout))

	return nil
}

func (l *VulkanLoader) VkCreateImageView(device VkDevice, pCreateInfo *VkImageViewCreateInfo, pAllocator *VkAllocationCallbacks, pView *VkImageView) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoCreateImageView(l.funcPtrs.vkCreateImageView,
		(C.VkDevice)(device),
		(*C.VkImageViewCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkImageView)(pView)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroyImageView(device VkDevice, imageView VkImageView, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyImageView(l.funcPtrs.vkDestroyImageView,
		(C.VkDevice)(device),
		(C.VkImageView)(imageView),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkCreateShaderModule(device VkDevice, pCreateInfo *VkShaderModuleCreateInfo, pAllocator *VkAllocationCallbacks, pShaderModule *VkShaderModule) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoCreateShaderModule(l.funcPtrs.vkCreateShaderModule,
		(C.VkDevice)(device),
		(*C.VkShaderModuleCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkShaderModule)(pShaderModule)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroyShaderModule(device VkDevice, shaderModule VkShaderModule, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyShaderModule(l.funcPtrs.vkDestroyShaderModule,
		(C.VkDevice)(device),
		(C.VkShaderModule)(shaderModule),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkCreatePipelineCache(device VkDevice, pCreateInfo *VkPipelineCacheCreateInfo, pAllocator *VkAllocationCallbacks, pPipelineCache *VkPipelineCache) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoCreatePipelineCache(l.funcPtrs.vkCreatePipelineCache,
		(C.VkDevice)(device),
		(*C.VkPipelineCacheCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkPipelineCache)(pPipelineCache)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroyPipelineCache(device VkDevice, pipelineCache VkPipelineCache, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyPipelineCache(l.funcPtrs.vkDestroyPipelineCache,
		(C.VkDevice)(device),
		(C.VkPipelineCache)(pipelineCache),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkGetPipelineCacheData(device VkDevice, pipelineCache VkPipelineCache, pDataSize *Size, pData unsafe.Pointer) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoGetPipelineCacheData(l.funcPtrs.vkGetPipelineCacheData,
		(C.VkDevice)(device),
		(C.VkPipelineCache)(pipelineCache),
		(*C.size_t)(pDataSize),
		pData))
	return res, res.ToError()
}

func (l *VulkanLoader) VkMergePipelineCaches(device VkDevice, dstCache VkPipelineCache, srcCacheCount Uint32, pSrcCaches *VkPipelineCache) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoMergePipelineCaches(l.funcPtrs.vkMergePipelineCaches,
		(C.VkDevice)(device),
		(C.VkPipelineCache)(dstCache),
		(C.uint32_t)(srcCacheCount),
		(*C.VkPipelineCache)(pSrcCaches)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkCreateGraphicsPipelines(device VkDevice, pipelineCache VkPipelineCache, createInfoCount Uint32, pCreateInfos *VkGraphicsPipelineCreateInfo, pAllocator *VkAllocationCallbacks, pPipelines *VkPipeline) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
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

func (l *VulkanLoader) VkCreateComputePipelines(device VkDevice, pipelineCache VkPipelineCache, createInfoCount Uint32, pCreateInfos *VkComputePipelineCreateInfo, pAllocator *VkAllocationCallbacks, pPipelines *VkPipeline) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
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

func (l *VulkanLoader) VkDestroyPipeline(device VkDevice, pipeline VkPipeline, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyPipeline(l.funcPtrs.vkDestroyPipeline,
		(C.VkDevice)(device),
		(C.VkPipeline)(pipeline),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkCreatePipelineLayout(device VkDevice, pCreateInfo *VkPipelineLayoutCreateInfo, pAllocator *VkAllocationCallbacks, pPipelineLayout *VkPipelineLayout) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoCreatePipelineLayout(l.funcPtrs.vkCreatePipelineLayout,
		(C.VkDevice)(device),
		(*C.VkPipelineLayoutCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkPipelineLayout)(pPipelineLayout)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroyPipelineLayout(device VkDevice, pipelineLayout VkPipelineLayout, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyPipelineLayout(l.funcPtrs.vkDestroyPipelineLayout,
		(C.VkDevice)(device),
		(C.VkPipelineLayout)(pipelineLayout),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkCreateSampler(device VkDevice, pCreateInfo *VkSamplerCreateInfo, pAllocator *VkAllocationCallbacks, pSampler *VkSampler) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoCreateSampler(l.funcPtrs.vkCreateSampler,
		(C.VkDevice)(device),
		(*C.VkSamplerCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkSampler)(pSampler)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroySampler(device VkDevice, sampler VkSampler, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroySampler(l.funcPtrs.vkDestroySampler,
		(C.VkDevice)(device),
		(C.VkSampler)(sampler),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkCreateDescriptorSetLayout(device VkDevice, pCreateInfo *VkDescriptorSetLayoutCreateInfo, pAllocator *VkAllocationCallbacks, pSetLayout *VkDescriptorSetLayout) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoCreateDescriptorSetLayout(l.funcPtrs.vkCreateDescriptorSetLayout,
		(C.VkDevice)(device),
		(*C.VkDescriptorSetLayoutCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkDescriptorSetLayout)(pSetLayout)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroyDescriptorSetLayout(device VkDevice, descriptorSetLayout VkDescriptorSetLayout, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyDescriptorSetLayout(l.funcPtrs.vkDestroyDescriptorSetLayout,
		(C.VkDevice)(device),
		(C.VkDescriptorSetLayout)(descriptorSetLayout),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkCreateDescriptorPool(device VkDevice, pCreateInfo *VkDescriptorPoolCreateInfo, pAllocator *VkAllocationCallbacks, pDescriptorPool *VkDescriptorPool) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoCreateDescriptorPool(l.funcPtrs.vkCreateDescriptorPool,
		(C.VkDevice)(device),
		(*C.VkDescriptorPoolCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkDescriptorPool)(pDescriptorPool)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroyDescriptorPool(device VkDevice, descriptorPool VkDescriptorPool, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyDescriptorPool(l.funcPtrs.vkDestroyDescriptorPool,
		(C.VkDevice)(device),
		(C.VkDescriptorPool)(descriptorPool),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkResetDescriptorPool(device VkDevice, descriptorPool VkDescriptorPool, flags VkDescriptorPoolResetFlags) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoResetDescriptorPool(l.funcPtrs.vkResetDescriptorPool,
		(C.VkDevice)(device),
		(C.VkDescriptorPool)(descriptorPool),
		(C.VkDescriptorPoolResetFlags)(flags)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkAllocateDescriptorSets(device VkDevice, pAllocateInfo *VkDescriptorSetAllocateInfo, pDescriptorSets *VkDescriptorSet) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoAllocateDescriptorSets(l.funcPtrs.vkAllocateDescriptorSets,
		(C.VkDevice)(device),
		(*C.VkDescriptorSetAllocateInfo)(pAllocateInfo),
		(*C.VkDescriptorSet)(pDescriptorSets)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkFreeDescriptorSets(device VkDevice, descriptorPool VkDescriptorPool, descriptorSetCount Uint32, pDescriptorSets *VkDescriptorSet) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoFreeDescriptorSets(l.funcPtrs.vkFreeDescriptorSets,
		(C.VkDevice)(device),
		(C.VkDescriptorPool)(descriptorPool),
		(C.uint32_t)(descriptorSetCount),
		(*C.VkDescriptorSet)(pDescriptorSets)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkUpdateDescriptorSets(device VkDevice, descriptorWriteCount Uint32, pDescriptorWrites *VkWriteDescriptorSet, descriptorCopyCount Uint32, pDescriptorCopies *VkCopyDescriptorSet) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoUpdateDescriptorSets(l.funcPtrs.vkUpdateDescriptorSets,
		(C.VkDevice)(device),
		(C.uint32_t)(descriptorWriteCount),
		(*C.VkWriteDescriptorSet)(pDescriptorWrites),
		(C.uint32_t)(descriptorCopyCount),
		(*C.VkCopyDescriptorSet)(pDescriptorCopies))

	return nil
}

func (l *VulkanLoader) VkCreateFramebuffer(device VkDevice, pCreateInfo *VkFramebufferCreateInfo, pAllocator *VkAllocationCallbacks, pFramebuffer *VkFramebuffer) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoCreateFramebuffer(l.funcPtrs.vkCreateFramebuffer,
		(C.VkDevice)(device),
		(*C.VkFramebufferCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkFramebuffer)(pFramebuffer)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroyFramebuffer(device VkDevice, framebuffer VkFramebuffer, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyFramebuffer(l.funcPtrs.vkDestroyFramebuffer,
		(C.VkDevice)(device),
		(C.VkFramebuffer)(framebuffer),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkCreateRenderPass(device VkDevice, pCreateInfo *VkRenderPassCreateInfo, pAllocator *VkAllocationCallbacks, pRenderPass *VkRenderPass) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoCreateRenderPass(l.funcPtrs.vkCreateRenderPass,
		(C.VkDevice)(device),
		(*C.VkRenderPassCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkRenderPass)(pRenderPass)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroyRenderPass(device VkDevice, renderPass VkRenderPass, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyRenderPass(l.funcPtrs.vkDestroyRenderPass,
		(C.VkDevice)(device),
		(C.VkRenderPass)(renderPass),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkGetRenderAreaGranularity(device VkDevice, renderPass VkRenderPass, pGranularity *VkExtent2D) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoGetRenderAreaGranularity(l.funcPtrs.vkGetRenderAreaGranularity,
		(C.VkDevice)(device),
		(C.VkRenderPass)(renderPass),
		(*C.VkExtent2D)(pGranularity))

	return nil
}

func (l *VulkanLoader) VkCreateCommandPool(device VkDevice, pCreateInfo *VkCommandPoolCreateInfo, pAllocator *VkAllocationCallbacks, pCommandPool *VkCommandPool) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoCreateCommandPool(l.funcPtrs.vkCreateCommandPool,
		(C.VkDevice)(device),
		(*C.VkCommandPoolCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkCommandPool)(pCommandPool)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroyCommandPool(device VkDevice, commandPool VkCommandPool, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyCommandPool(l.funcPtrs.vkDestroyCommandPool,
		(C.VkDevice)(device),
		(C.VkCommandPool)(commandPool),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkResetCommandPool(device VkDevice, commandPool VkCommandPool, flags VkCommandPoolResetFlags) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoResetCommandPool(l.funcPtrs.vkResetCommandPool,
		(C.VkDevice)(device),
		(C.VkCommandPool)(commandPool),
		(C.VkCommandPoolResetFlags)(flags)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkAllocateCommandBuffers(device VkDevice, pAllocateInfo *VkCommandBufferAllocateInfo, pCommandBuffers *VkCommandBuffer) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoAllocateCommandBuffers(l.funcPtrs.vkAllocateCommandBuffers,
		(C.VkDevice)(device),
		(*C.VkCommandBufferAllocateInfo)(pAllocateInfo),
		(*C.VkCommandBuffer)(pCommandBuffers)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkFreeCommandBuffers(device VkDevice, commandPool VkCommandPool, commandBufferCount Uint32, pCommandBuffers *VkCommandBuffer) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoFreeCommandBuffers(l.funcPtrs.vkFreeCommandBuffers,
		(C.VkDevice)(device),
		(C.VkCommandPool)(commandPool),
		(C.uint32_t)(commandBufferCount),
		(*C.VkCommandBuffer)(pCommandBuffers))

	return nil
}

func (l *VulkanLoader) VkBeginCommandBuffer(commandBuffer VkCommandBuffer, pBeginInfo *VkCommandBufferBeginInfo) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoBeginCommandBuffer(l.funcPtrs.vkBeginCommandBuffer,
		(C.VkCommandBuffer)(commandBuffer),
		(*C.VkCommandBufferBeginInfo)(pBeginInfo)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkEndCommandBuffer(commandBuffer VkCommandBuffer) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoEndCommandBuffer(l.funcPtrs.vkEndCommandBuffer,
		(C.VkCommandBuffer)(commandBuffer)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkResetCommandBuffer(commandBuffer VkCommandBuffer, flags VkCommandBufferResetFlags) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoResetCommandBuffer(l.funcPtrs.vkResetCommandBuffer,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkCommandBufferResetFlags)(flags)))
	return res, res.ToError()
}

func (l *VulkanLoader) VkCmdBindPipeline(commandBuffer VkCommandBuffer, pipelineBindPoint VkPipelineBindPoint, pipeline VkPipeline) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdBindPipeline(l.funcPtrs.vkCmdBindPipeline,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkPipelineBindPoint)(pipelineBindPoint),
		(C.VkPipeline)(pipeline))

	return nil
}

func (l *VulkanLoader) VkCmdSetViewport(commandBuffer VkCommandBuffer, firstViewport Uint32, viewportCount Uint32, pViewports *VkViewport) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdSetViewport(l.funcPtrs.vkCmdSetViewport,
		(C.VkCommandBuffer)(commandBuffer),
		(C.uint32_t)(firstViewport),
		(C.uint32_t)(viewportCount),
		(*C.VkViewport)(pViewports))

	return nil
}

func (l *VulkanLoader) VkCmdSetScissor(commandBuffer VkCommandBuffer, firstScissor Uint32, scissorCount Uint32, pScissors *VkRect2D) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdSetScissor(l.funcPtrs.vkCmdSetScissor,
		(C.VkCommandBuffer)(commandBuffer),
		C.uint32_t(firstScissor),
		C.uint32_t(scissorCount),
		(*C.VkRect2D)(pScissors))

	return nil
}

func (l *VulkanLoader) VkCmdSetLineWidth(commandBuffer VkCommandBuffer, lineWidth Float) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}
	C.cgoCmdSetLineWidth(l.funcPtrs.vkCmdSetLineWidth,
		(C.VkCommandBuffer)(commandBuffer),
		(C.float)(lineWidth))

	return nil
}

func (l *VulkanLoader) VkCmdSetDepthBias(commandBuffer VkCommandBuffer, depthBiasConstantFactor Float, depthBiasClamp Float, depthBiasSlopeFactor Float) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdSetDepthBias(l.funcPtrs.vkCmdSetDepthBias,
		(C.VkCommandBuffer)(commandBuffer),
		(C.float)(depthBiasConstantFactor),
		(C.float)(depthBiasClamp),
		(C.float)(depthBiasSlopeFactor))

	return nil
}

func (l *VulkanLoader) VkCmdSetBlendConstants(commandBuffer VkCommandBuffer, blendConstants *Float) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdSetBlendConstants(l.funcPtrs.vkCmdSetBlendConstants,
		(C.VkCommandBuffer)(commandBuffer),
		(*C.float)(blendConstants),
	)

	return nil
}

func (l *VulkanLoader) VkCmdSetDepthBounds(commandBuffer VkCommandBuffer, minDepthBounds Float, maxDepthBounds Float) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdSetDepthBounds(l.funcPtrs.vkCmdSetDepthBounds,
		(C.VkCommandBuffer)(commandBuffer),
		(C.float)(minDepthBounds),
		(C.float)(maxDepthBounds))

	return nil
}

func (l *VulkanLoader) VkCmdSetStencilCompareMask(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, compareMask Uint32) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdSetStencilCompareMask(l.funcPtrs.vkCmdSetStencilCompareMask,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkStencilFaceFlags)(faceMask),
		(C.uint32_t)(compareMask))

	return nil
}

func (l *VulkanLoader) VkCmdSetStencilWriteMask(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, writeMask Uint32) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdSetStencilWriteMask(l.funcPtrs.vkCmdSetStencilWriteMask,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkStencilFaceFlags)(faceMask),
		(C.uint32_t)(writeMask))

	return nil
}

func (l *VulkanLoader) VkCmdSetStencilReference(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, reference Uint32) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdSetStencilReference(l.funcPtrs.vkCmdSetStencilReference,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkStencilFaceFlags)(faceMask),
		(C.uint32_t)(reference))

	return nil
}

func (l *VulkanLoader) VkCmdBindDescriptorSets(commandBuffer VkCommandBuffer, pipelineBindPoint VkPipelineBindPoint, layout VkPipelineLayout, firstSet Uint32, descriptorSetCount Uint32, pDescriptorSets *VkDescriptorSet, dynamicOffsetCount Uint32, pDynamicOffsets *Uint32) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
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

	return nil
}

func (l *VulkanLoader) VkCmdBindIndexBuffer(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, indexType VkIndexType) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdBindIndexBuffer(l.funcPtrs.vkCmdBindIndexBuffer,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkBuffer)(buffer),
		(C.VkDeviceSize)(offset),
		(C.VkIndexType)(indexType))

	return nil
}

func (l *VulkanLoader) VkCmdBindVertexBuffers(commandBuffer VkCommandBuffer, firstBinding Uint32, bindingCount Uint32, pBuffers *VkBuffer, pOffsets *VkDeviceSize) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdBindVertexBuffers(l.funcPtrs.vkCmdBindVertexBuffers,
		(C.VkCommandBuffer)(commandBuffer),
		C.uint32_t(firstBinding),
		C.uint32_t(bindingCount),
		(*C.VkBuffer)(pBuffers),
		(*C.VkDeviceSize)(pOffsets))

	return nil
}

func (l *VulkanLoader) VkCmdDraw(commandBuffer VkCommandBuffer, vertexCount Uint32, instanceCount Uint32, firstVertex Uint32, firstInstance Uint32) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdDraw(l.funcPtrs.vkCmdDraw,
		(C.VkCommandBuffer)(commandBuffer),
		(C.uint32_t)(vertexCount),
		(C.uint32_t)(instanceCount),
		(C.uint32_t)(firstVertex),
		(C.uint32_t)(firstInstance))

	return nil
}

func (l *VulkanLoader) VkCmdDrawIndexed(commandBuffer VkCommandBuffer, indexCount Uint32, instanceCount Uint32, firstIndex Uint32, vertexOffset Int32, firstInstance Uint32) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdDrawIndexed(l.funcPtrs.vkCmdDrawIndexed,
		(C.VkCommandBuffer)(commandBuffer),
		(C.uint32_t)(indexCount),
		(C.uint32_t)(instanceCount),
		(C.uint32_t)(firstIndex),
		(C.int32_t)(vertexOffset),
		(C.uint32_t)(firstInstance))

	return nil
}

func (l *VulkanLoader) VkCmdDrawIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, drawCount Uint32, stride Uint32) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdDrawIndirect(l.funcPtrs.vkCmdDrawIndirect,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkBuffer)(buffer),
		(C.VkDeviceSize)(offset),
		(C.uint32_t)(drawCount),
		(C.uint32_t)(stride))

	return nil
}

func (l *VulkanLoader) VkCmdDrawIndexedIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, drawCount Uint32, stride Uint32) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdDrawIndexedIndirect(l.funcPtrs.vkCmdDrawIndexedIndirect,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkBuffer)(buffer),
		(C.VkDeviceSize)(offset),
		(C.uint32_t)(drawCount),
		(C.uint32_t)(stride))

	return nil
}

func (l *VulkanLoader) VkCmdDispatch(commandBuffer VkCommandBuffer, groupCountX Uint32, groupCountY Uint32, groupCountZ Uint32) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdDispatch(l.funcPtrs.vkCmdDispatch,
		(C.VkCommandBuffer)(commandBuffer),
		(C.uint32_t)(groupCountX),
		(C.uint32_t)(groupCountY),
		(C.uint32_t)(groupCountZ))

	return nil
}

func (l *VulkanLoader) VkCmdDispatchIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdDispatchIndirect(l.funcPtrs.vkCmdDispatchIndirect,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkBuffer)(buffer),
		(C.VkDeviceSize)(offset))

	return nil
}

func (l *VulkanLoader) VkCmdCopyBuffer(commandBuffer VkCommandBuffer, srcBuffer VkBuffer, dstBuffer VkBuffer, regionCount Uint32, pRegions *VkBufferCopy) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdCopyBuffer(l.funcPtrs.vkCmdCopyBuffer,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkBuffer)(srcBuffer),
		(C.VkBuffer)(dstBuffer),
		(C.uint32_t)(regionCount),
		(*C.VkBufferCopy)(pRegions))

	return nil
}

func (l *VulkanLoader) VkCmdCopyImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkImageCopy) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdCopyImage(l.funcPtrs.vkCmdCopyImage,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkImage)(srcImage),
		(C.VkImageLayout)(srcImageLayout),
		(C.VkImage)(dstImage),
		(C.VkImageLayout)(dstImageLayout),
		(C.uint32_t)(regionCount),
		(*C.VkImageCopy)(pRegions))

	return nil
}

func (l *VulkanLoader) VkCmdBlitImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkImageBlit, filter VkFilter) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
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

	return nil
}

func (l *VulkanLoader) VkCmdCopyBufferToImage(commandBuffer VkCommandBuffer, srcBuffer VkBuffer, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkBufferImageCopy) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdCopyBufferToImage(l.funcPtrs.vkCmdCopyBufferToImage,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkBuffer)(srcBuffer),
		(C.VkImage)(dstImage),
		(C.VkImageLayout)(dstImageLayout),
		(C.uint32_t)(regionCount),
		(*C.VkBufferImageCopy)(pRegions))

	return nil
}

func (l *VulkanLoader) VkCmdCopyImageToBuffer(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstBuffer VkBuffer, regionCount Uint32, pRegions *VkBufferImageCopy) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdCopyImageToBuffer(l.funcPtrs.vkCmdCopyImageToBuffer,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkImage)(srcImage),
		(C.VkImageLayout)(srcImageLayout),
		(C.VkBuffer)(dstBuffer),
		(C.uint32_t)(regionCount),
		(*C.VkBufferImageCopy)(pRegions))

	return nil
}

func (l *VulkanLoader) VkCmdUpdateBuffer(commandBuffer VkCommandBuffer, dstBuffer VkBuffer, dstOffset VkDeviceSize, dataSize VkDeviceSize, pData unsafe.Pointer) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdUpdateBuffer(l.funcPtrs.vkCmdUpdateBuffer,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkBuffer)(dstBuffer),
		(C.VkDeviceSize)(dstOffset),
		(C.VkDeviceSize)(dataSize),
		pData)

	return nil
}

func (l *VulkanLoader) VkCmdFillBuffer(commandBuffer VkCommandBuffer, dstBuffer VkBuffer, dstOffset VkDeviceSize, size VkDeviceSize, data Uint32) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdFillBuffer(l.funcPtrs.vkCmdFillBuffer,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkBuffer)(dstBuffer),
		(C.VkDeviceSize)(dstOffset),
		(C.VkDeviceSize)(size),
		(C.uint32_t)(data))

	return nil
}

func (l *VulkanLoader) VkCmdClearColorImage(commandBuffer VkCommandBuffer, image VkImage, imageLayout VkImageLayout, pColor *VkClearColorValue, rangeCount Uint32, pRanges *VkImageSubresourceRange) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdClearColorImage(l.funcPtrs.vkCmdClearColorImage,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkImage)(image),
		(C.VkImageLayout)(imageLayout),
		(*C.VkClearColorValue)(pColor),
		(C.uint32_t)(rangeCount),
		(*C.VkImageSubresourceRange)(pRanges))

	return nil
}

func (l *VulkanLoader) VkCmdClearDepthStencilImage(commandBuffer VkCommandBuffer, image VkImage, imageLayout VkImageLayout, pDepthStencil *VkClearDepthStencilValue, rangeCount Uint32, pRanges *VkImageSubresourceRange) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdClearDepthStencilImage(l.funcPtrs.vkCmdClearDepthStencilImage,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkImage)(image),
		(C.VkImageLayout)(imageLayout),
		(*C.VkClearDepthStencilValue)(pDepthStencil),
		(C.uint32_t)(rangeCount),
		(*C.VkImageSubresourceRange)(pRanges))

	return nil
}

func (l *VulkanLoader) VkCmdClearAttachments(commandBuffer VkCommandBuffer, attachmentCount Uint32, pAttachments *VkClearAttachment, rectCount Uint32, pRects *VkClearRect) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdClearAttachments(l.funcPtrs.vkCmdClearAttachments,
		(C.VkCommandBuffer)(commandBuffer),
		(C.uint32_t)(attachmentCount),
		(*C.VkClearAttachment)(pAttachments),
		(C.uint32_t)(rectCount),
		(*C.VkClearRect)(pRects))

	return nil
}

func (l *VulkanLoader) VkCmdResolveImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkImageResolve) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdResolveImage(l.funcPtrs.vkCmdResolveImage,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkImage)(srcImage),
		(C.VkImageLayout)(srcImageLayout),
		(C.VkImage)(dstImage),
		(C.VkImageLayout)(dstImageLayout),
		(C.uint32_t)(regionCount),
		(*C.VkImageResolve)(pRegions))

	return nil
}

func (l *VulkanLoader) VkCmdSetEvent(commandBuffer VkCommandBuffer, event VkEvent, stageMask VkPipelineStageFlags) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdSetEvent(l.funcPtrs.vkCmdSetEvent,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkEvent)(event),
		(C.VkPipelineStageFlags)(stageMask))

	return nil
}

func (l *VulkanLoader) VkCmdResetEvent(commandBuffer VkCommandBuffer, event VkEvent, stageMask VkPipelineStageFlags) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdResetEvent(l.funcPtrs.vkCmdResetEvent,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkEvent)(event),
		(C.VkPipelineStageFlags)(stageMask))

	return nil
}

func (l *VulkanLoader) VkCmdWaitEvents(commandBuffer VkCommandBuffer, eventCount Uint32, pEvents *VkEvent, srcStageMask VkPipelineStageFlags, dstStageMask VkPipelineStageFlags, memoryBarrierCount Uint32, pMemoryBarriers *VkMemoryBarrier, bufferMemoryBarrierCount Uint32, pBufferMemoryBarriers *VkBufferMemoryBarrier, imageMemoryBarrierCount Uint32, pImageMemoryBarriers *VkImageMemoryBarrier) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
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

	return nil
}

func (l *VulkanLoader) VkCmdPipelineBarrier(commandBuffer VkCommandBuffer, srcStageMask VkPipelineStageFlags, dstStageMask VkPipelineStageFlags, dependencyFlags VkDependencyFlags, memoryBarrierCount Uint32, pMemoryBarriers *VkMemoryBarrier, bufferMemoryBarrierCount Uint32, pBufferMemoryBarriers *VkBufferMemoryBarrier, imageMemoryBarrierCount Uint32, pImageMemoryBarriers *VkImageMemoryBarrier) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
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

	return nil
}

func (l *VulkanLoader) VkCmdBeginQuery(commandBuffer VkCommandBuffer, queryPool VkQueryPool, query Uint32, flags VkQueryControlFlags) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdBeginQuery(l.funcPtrs.vkCmdBeginQuery,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkQueryPool)(queryPool),
		(C.uint32_t)(query),
		(C.VkQueryControlFlags)(flags))

	return nil
}

func (l *VulkanLoader) VkCmdEndQuery(commandBuffer VkCommandBuffer, queryPool VkQueryPool, query Uint32) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdEndQuery(l.funcPtrs.vkCmdEndQuery,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkQueryPool)(queryPool),
		(C.uint32_t)(query))

	return nil
}

func (l *VulkanLoader) VkCmdResetQueryPool(commandBuffer VkCommandBuffer, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdResetQueryPool(l.funcPtrs.vkCmdResetQueryPool,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkQueryPool)(queryPool),
		(C.uint32_t)(firstQuery),
		(C.uint32_t)(queryCount))

	return nil
}

func (l *VulkanLoader) VkCmdWriteTimestamp(commandBuffer VkCommandBuffer, pipelineStage VkPipelineStageFlags, queryPool VkQueryPool, query Uint32) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdWriteTimestamp(l.funcPtrs.vkCmdWriteTimestamp,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkPipelineStageFlagBits)(pipelineStage),
		(C.VkQueryPool)(queryPool),
		(C.uint32_t)(query))

	return nil
}

func (l *VulkanLoader) VkCmdCopyQueryPoolResults(commandBuffer VkCommandBuffer, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32, dstBuffer VkBuffer, dstOffset VkDeviceSize, stride VkDeviceSize, flags VkQueryResultFlags) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
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

	return nil
}

func (l *VulkanLoader) VkCmdPushConstants(commandBuffer VkCommandBuffer, layout VkPipelineLayout, stageFlags VkShaderStageFlags, offset Uint32, size Uint32, pValues unsafe.Pointer) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdPushConstants(l.funcPtrs.vkCmdPushConstants,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkPipelineLayout)(layout),
		(C.VkShaderStageFlags)(stageFlags),
		(C.uint32_t)(offset),
		(C.uint32_t)(size),
		pValues)

	return nil
}

func (l *VulkanLoader) VkCmdBeginRenderPass(commandBuffer VkCommandBuffer, pRenderPassBegin *VkRenderPassBeginInfo, contents VkSubpassContents) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdBeginRenderPass(l.funcPtrs.vkCmdBeginRenderPass,
		(C.VkCommandBuffer)(commandBuffer),
		(*C.VkRenderPassBeginInfo)(pRenderPassBegin),
		(C.VkSubpassContents)(contents))

	return nil
}

func (l *VulkanLoader) VkCmdNextSubpass(commandBuffer VkCommandBuffer, contents VkSubpassContents) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdNextSubpass(l.funcPtrs.vkCmdNextSubpass,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkSubpassContents)(contents))

	return nil
}

func (l *VulkanLoader) VkCmdEndRenderPass(commandBuffer VkCommandBuffer) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdEndRenderPass(l.funcPtrs.vkCmdEndRenderPass,
		(C.VkCommandBuffer)(commandBuffer))

	return nil
}

func (l *VulkanLoader) VkCmdExecuteCommands(commandBuffer VkCommandBuffer, commandBufferCount Uint32, pCommandBuffers *VkCommandBuffer) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdExecuteCommands(l.funcPtrs.vkCmdExecuteCommands,
		(C.VkCommandBuffer)(commandBuffer),
		(C.uint32_t)(commandBufferCount),
		(*C.VkCommandBuffer)(pCommandBuffers))

	return nil
}

func (l *VulkanLoader) VkEnumerateInstanceVersion(pApiVersion *Uint32) (VkResult, error) {
	if l.funcPtrs.vkEnumerateInstanceVersion == nil {
		return VKErrorUnknown, errors.New("attempted to call method 'vkEnumerateInstanceVersion' which is not present on this loader")
	}

	res := VkResult(C.cgoEnumerateInstanceVersion(l.funcPtrs.vkEnumerateInstanceVersion,
		(*C.uint32_t)(pApiVersion)))

	return res, res.ToError()
}

func (l *VulkanLoader) VkEnumeratePhysicalDeviceGroups(instance VkInstance, pPhysicalDeviceGroupCount *Uint32, pPhysicalDeviceGroupProperties *VkPhysicalDeviceGroupProperties) (VkResult, error) {
	if l.funcPtrs.vkEnumeratePhysicalDeviceGroups == nil {
		return VKErrorUnknown, errors.New("attempted to call method 'vkEnumeratePhysicalDeviceGroups' which is not present on this loader")
	}

	res := VkResult(C.cgoEnumeratePhysicalDeviceGroups(l.funcPtrs.vkEnumeratePhysicalDeviceGroups,
		C.VkInstance(instance),
		(*C.uint32_t)(pPhysicalDeviceGroupCount),
		(*C.VkPhysicalDeviceGroupProperties)(pPhysicalDeviceGroupProperties)))

	return res, res.ToError()
}

func (l *VulkanLoader) VkGetPhysicalDeviceFeatures2(physicalDevice VkPhysicalDevice, pFeatures *VkPhysicalDeviceFeatures2) error {
	if l.funcPtrs.vkGetPhysicalDeviceFeatures2 == nil {
		return errors.New("attempted to call method 'vkGetPhysicalDeviceFeatures2' which is not present on this loader")
	}

	C.cgoGetPhysicalDeviceFeatures2(l.funcPtrs.vkGetPhysicalDeviceFeatures2,
		C.VkPhysicalDevice(physicalDevice),
		(*C.VkPhysicalDeviceFeatures2)(pFeatures))

	return nil
}

func (l *VulkanLoader) VkGetPhysicalDeviceProperties2(physicalDevice VkPhysicalDevice, pProperties *VkPhysicalDeviceProperties2) error {
	if l.funcPtrs.vkGetPhysicalDeviceProperties2 == nil {
		return errors.New("attempted to call method 'vkGetPhysicalDeviceProperties2' which is not present on this loader")
	}

	C.cgoGetPhysicalDeviceProperties2(l.funcPtrs.vkGetPhysicalDeviceProperties2,
		C.VkPhysicalDevice(physicalDevice),
		(*C.VkPhysicalDeviceProperties2)(pProperties))

	return nil
}

func (l *VulkanLoader) VkGetPhysicalDeviceFormatProperties2(physicalDevice VkPhysicalDevice, format VkFormat, pFormatProperties *VkFormatProperties2) error {
	if l.funcPtrs.vkGetPhysicalDeviceFormatProperties2 == nil {
		return errors.New("attempted to call method 'vkGetPhysicalDeviceFormatProperties2' which is not present on this loader")
	}

	C.cgoGetPhysicalDeviceFormatProperties2(l.funcPtrs.vkGetPhysicalDeviceFormatProperties2,
		C.VkPhysicalDevice(physicalDevice),
		C.VkFormat(format),
		(*C.VkFormatProperties2)(pFormatProperties))

	return nil
}

func (l *VulkanLoader) VkGetPhysicalDeviceImageFormatProperties2(physicalDevice VkPhysicalDevice, pImageFormatInfo *VkPhysicalDeviceImageFormatInfo2, pImageFormatProperties *VkImageFormatProperties2) (VkResult, error) {
	if l.funcPtrs.vkGetPhysicalDeviceImageFormatProperties2 == nil {
		return VKErrorUnknown, errors.New("attempted to call method 'vkGetPhysicalDeviceImageFormatProperties2' which is not present on this loader")
	}

	res := VkResult(C.cgoGetPhysicalDeviceImageFormatProperties2(l.funcPtrs.vkGetPhysicalDeviceImageFormatProperties2,
		C.VkPhysicalDevice(physicalDevice),
		(*C.VkPhysicalDeviceImageFormatInfo2)(pImageFormatInfo),
		(*C.VkImageFormatProperties2)(pImageFormatProperties)))

	return res, res.ToError()
}

func (l *VulkanLoader) VkGetPhysicalDeviceQueueFamilyProperties2(physicalDevice VkPhysicalDevice, pQueueFamilyPropertyCount *Uint32, pQueueFamilyProperties *VkQueueFamilyProperties2) error {
	if l.funcPtrs.vkGetPhysicalDeviceQueueFamilyProperties2 == nil {
		return errors.New("attempted to call method 'vkGetPhysicalDeviceQueueFamilyProperties2' which is not present on this loader")
	}

	C.cgoGetPhysicalDeviceQueueFamilyProperties2(l.funcPtrs.vkGetPhysicalDeviceQueueFamilyProperties2,
		C.VkPhysicalDevice(physicalDevice),
		(*C.uint32_t)(pQueueFamilyPropertyCount),
		(*C.VkQueueFamilyProperties2)(pQueueFamilyProperties))

	return nil
}

func (l *VulkanLoader) VkGetPhysicalDeviceMemoryProperties2(physicalDevice VkPhysicalDevice, pMemoryProperties *VkPhysicalDeviceMemoryProperties2) error {
	if l.funcPtrs.vkGetPhysicalDeviceMemoryProperties2 == nil {
		return errors.New("attempted to call method 'vkGetPhysicalDeviceMemoryProperties2' which is not present on this loader")
	}

	C.cgoGetPhysicalDeviceMemoryProperties2(l.funcPtrs.vkGetPhysicalDeviceMemoryProperties2,
		C.VkPhysicalDevice(physicalDevice),
		(*C.VkPhysicalDeviceMemoryProperties2)(pMemoryProperties))

	return nil
}

func (l *VulkanLoader) VkGetPhysicalDeviceSparseImageFormatProperties2(physicalDevice VkPhysicalDevice, pFormatInfo *VkPhysicalDeviceSparseImageFormatInfo2, pPropertyCount *Uint32, pProperties *VkSparseImageFormatProperties2) error {
	if l.funcPtrs.vkGetPhysicalDeviceSparseImageFormatProperties2 == nil {
		return errors.New("attempted to call method 'vkGetPhysicalDeviceSparseImageFormatProperties2' which is not present on this loader")
	}

	C.cgoGetPhysicalDeviceSparseImageFormatProperties2(l.funcPtrs.vkGetPhysicalDeviceSparseImageFormatProperties2,
		C.VkPhysicalDevice(physicalDevice),
		(*C.VkPhysicalDeviceSparseImageFormatInfo2)(pFormatInfo),
		(*C.uint32_t)(pPropertyCount),
		(*C.VkSparseImageFormatProperties2)(pProperties))

	return nil
}

func (l *VulkanLoader) VkGetPhysicalDeviceExternalBufferProperties(physicalDevice VkPhysicalDevice, pExternalBufferInfo *VkPhysicalDeviceExternalBufferInfo, pExternalBufferProperties *VkExternalBufferProperties) error {
	if l.funcPtrs.vkGetPhysicalDeviceExternalBufferProperties == nil {
		return errors.New("attempted to call method 'vkGetPhysicalDeviceExternalBufferProperties' which is not present on this loader")
	}

	C.cgoGetPhysicalDeviceExternalBufferProperties(l.funcPtrs.vkGetPhysicalDeviceExternalBufferProperties,
		C.VkPhysicalDevice(physicalDevice),
		(*C.VkPhysicalDeviceExternalBufferInfo)(pExternalBufferInfo),
		(*C.VkExternalBufferProperties)(pExternalBufferProperties))

	return nil
}

func (l *VulkanLoader) VkGetPhysicalDeviceExternalFenceProperties(physicalDevice VkPhysicalDevice, pExternalFenceInfo *VkPhysicalDeviceExternalFenceInfo, pExternalFenceProperties *VkExternalFenceProperties) error {
	if l.funcPtrs.vkGetPhysicalDeviceExternalFenceProperties == nil {
		return errors.New("attempted to call method 'vkGetPhysicalDeviceExternalFenceProperties' which is not present on this loader")
	}

	C.cgoGetPhysicalDeviceExternalFenceProperties(l.funcPtrs.vkGetPhysicalDeviceExternalFenceProperties,
		C.VkPhysicalDevice(physicalDevice),
		(*C.VkPhysicalDeviceExternalFenceInfo)(pExternalFenceInfo),
		(*C.VkExternalFenceProperties)(pExternalFenceProperties))

	return nil
}

func (l *VulkanLoader) VkGetPhysicalDeviceExternalSemaphoreProperties(physicalDevice VkPhysicalDevice, pExternalSemaphoreInfo *VkPhysicalDeviceExternalSemaphoreInfo, pExternalSemaphoreProperties *VkExternalSemaphoreProperties) error {
	if l.funcPtrs.vkGetPhysicalDeviceExternalSemaphoreProperties == nil {
		return errors.New("attempted to call method 'vkGetPhysicalDeviceExternalSemaphoreProperties' which is not present on this loader")
	}

	C.cgoGetPhysicalDeviceExternalSemaphoreProperties(l.funcPtrs.vkGetPhysicalDeviceExternalSemaphoreProperties,
		C.VkPhysicalDevice(physicalDevice),
		(*C.VkPhysicalDeviceExternalSemaphoreInfo)(pExternalSemaphoreInfo),
		(*C.VkExternalSemaphoreProperties)(pExternalSemaphoreProperties))

	return nil
}

func (l *VulkanLoader) VkBindBufferMemory2(device VkDevice, bindInfoCount Uint32, pBindInfos *VkBindBufferMemoryInfo) (VkResult, error) {
	if l.funcPtrs.vkBindBufferMemory2 == nil {
		return VKErrorUnknown, errors.New("attempted to call method 'vkBindBufferMemory2' which is not present on this loader")
	}

	res := VkResult(C.cgoBindBufferMemory2(l.funcPtrs.vkBindBufferMemory2,
		C.VkDevice(device),
		C.uint32_t(bindInfoCount),
		(*C.VkBindBufferMemoryInfo)(pBindInfos)))

	return res, res.ToError()
}

func (l *VulkanLoader) VkBindImageMemory2(device VkDevice, bindInfoCount Uint32, pBindInfos *VkBindImageMemoryInfo) (VkResult, error) {
	if l.funcPtrs.vkBindImageMemory2 == nil {
		return VKErrorUnknown, errors.New("attempted to call method 'vkBindImageMemory2' which is not present on this loader")
	}

	res := VkResult(C.cgoBindImageMemory2(l.funcPtrs.vkBindImageMemory2,
		C.VkDevice(device),
		C.uint32_t(bindInfoCount),
		(*C.VkBindImageMemoryInfo)(pBindInfos)))

	return res, res.ToError()
}

func (l *VulkanLoader) VkGetDeviceGroupPeerMemoryFeatures(device VkDevice, heapIndex Uint32, localDeviceIndex Uint32, remoteDeviceIndex Uint32, pPeerMemoryFeatures *VkPeerMemoryFeatureFlags) error {
	if l.funcPtrs.vkGetDeviceGroupPeerMemoryFeatures == nil {
		return errors.New("attempted to call method 'vkGetDeviceGroupPeerMemoryFeatures' which is not present on this loader")
	}

	C.cgoGetDeviceGroupPeerMemoryFeatures(l.funcPtrs.vkGetDeviceGroupPeerMemoryFeatures,
		C.VkDevice(device),
		C.uint32_t(heapIndex),
		C.uint32_t(localDeviceIndex),
		C.uint32_t(remoteDeviceIndex),
		(*C.VkPeerMemoryFeatureFlags)(pPeerMemoryFeatures))

	return nil
}

func (l *VulkanLoader) VkCmdSetDeviceMask(commandBuffer VkCommandBuffer, deviceMask Uint32) error {
	if l.funcPtrs.vkCmdSetDeviceMask == nil {
		return errors.New("attempted to call method 'vkCmdSetDeviceMask' which is not present on this loader")
	}

	C.cgoCmdSetDeviceMask(l.funcPtrs.vkCmdSetDeviceMask,
		C.VkCommandBuffer(commandBuffer),
		C.uint32_t(deviceMask))

	return nil
}

func (l *VulkanLoader) VkCmdDispatchBase(commandBuffer VkCommandBuffer, baseGroupX Uint32, baseGroupY Uint32, baseGroupZ Uint32, groupCountX Uint32, groupCountY Uint32, groupCountZ Uint32) error {
	if l.funcPtrs.vkCmdDispatchBase == nil {
		return errors.New("attempted to call method 'vkCmdDispatchBase' which is not present on this loader")
	}

	C.cgoCmdDispatchBase(l.funcPtrs.vkCmdDispatchBase,
		C.VkCommandBuffer(commandBuffer),
		C.uint32_t(baseGroupX),
		C.uint32_t(baseGroupY),
		C.uint32_t(baseGroupZ),
		C.uint32_t(groupCountX),
		C.uint32_t(groupCountY),
		C.uint32_t(groupCountZ))

	return nil
}

func (l *VulkanLoader) VkGetImageMemoryRequirements2(device VkDevice, pInfo *VkImageMemoryRequirementsInfo2, pMemoryRequirements *VkMemoryRequirements2) error {
	if l.funcPtrs.vkGetImageMemoryRequirements2 == nil {
		return errors.New("attempted to call method 'vkGetImageMemoryRequirements2' which is not present on this loader")
	}

	C.cgoGetImageMemoryRequirements2(l.funcPtrs.vkGetImageMemoryRequirements2,
		C.VkDevice(device),
		(*C.VkImageMemoryRequirementsInfo2)(pInfo),
		(*C.VkMemoryRequirements2)(pMemoryRequirements))

	return nil
}

func (l *VulkanLoader) VkGetBufferMemoryRequirements2(device VkDevice, pInfo *VkBufferMemoryRequirementsInfo2, pMemoryRequirements *VkMemoryRequirements2) error {
	if l.funcPtrs.vkGetBufferMemoryRequirements2 == nil {
		return errors.New("attempted to call method 'vkGetBufferMemoryRequirements2' which is not present on this loader")
	}

	C.cgoGetBufferMemoryRequirements2(l.funcPtrs.vkGetBufferMemoryRequirements2,
		C.VkDevice(device),
		(*C.VkBufferMemoryRequirementsInfo2)(pInfo),
		(*C.VkMemoryRequirements2)(pMemoryRequirements))

	return nil
}

func (l *VulkanLoader) VkGetImageSparseMemoryRequirements2(device VkDevice, pInfo *VkImageSparseMemoryRequirementsInfo2, pSparseMemoryRequirementCount *Uint32, pSparseMemoryRequirements *VkSparseImageMemoryRequirements2) error {
	if l.funcPtrs.vkGetImageSparseMemoryRequirements2 == nil {
		return errors.New("attempted to call method 'vkGetImageSparseMemoryRequirements2' which is not present on this loader")
	}

	C.cgoGetImageSparseMemoryRequirements2(l.funcPtrs.vkGetImageSparseMemoryRequirements2,
		C.VkDevice(device),
		(*C.VkImageSparseMemoryRequirementsInfo2)(pInfo),
		(*C.uint32_t)(pSparseMemoryRequirementCount),
		(*C.VkSparseImageMemoryRequirements2)(pSparseMemoryRequirements))

	return nil
}

func (l *VulkanLoader) VkTrimCommandPool(device VkDevice, commandPool VkCommandPool, flags VkCommandPoolTrimFlags) error {
	if l.funcPtrs.vkTrimCommandPool == nil {
		return errors.New("attempted to call method 'vkTrimCommandPool' which is not present on this loader")
	}

	C.cgoTrimCommandPool(l.funcPtrs.vkTrimCommandPool,
		C.VkDevice(device),
		C.VkCommandPool(commandPool),
		C.VkCommandPoolTrimFlags(flags))

	return nil
}

func (l *VulkanLoader) VkGetDeviceQueue2(device VkDevice, pQueueInfo *VkDeviceQueueInfo2, pQueue *VkQueue) error {
	if l.funcPtrs.vkGetDeviceQueue2 == nil {
		return errors.New("attempted to call method 'vkGetDeviceQueue2' which is not present on this loader")
	}

	C.cgoGetDeviceQueue2(l.funcPtrs.vkGetDeviceQueue2,
		C.VkDevice(device),
		(*C.VkDeviceQueueInfo2)(pQueueInfo),
		(*C.VkQueue)(pQueue))

	return nil
}

func (l *VulkanLoader) VkCreateSamplerYcbcrConversion(device VkDevice, pCreateInfo *VkSamplerYcbcrConversionCreateInfo, pAllocator *VkAllocationCallbacks, pYcbcrConversion *VkSamplerYcbcrConversion) (VkResult, error) {
	if l.funcPtrs.vkCreateSamplerYcbcrConversion == nil {
		return VKErrorUnknown, errors.New("attempted to call method 'vkCreateSamplerYcbcrConversion' which is not present on this loader")
	}

	res := VkResult(C.cgoCreateSamplerYcbcrConversion(l.funcPtrs.vkCreateSamplerYcbcrConversion,
		C.VkDevice(device),
		(*C.VkSamplerYcbcrConversionCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkSamplerYcbcrConversion)(pYcbcrConversion)))

	return res, res.ToError()
}

func (l *VulkanLoader) VkDestroySamplerYcbcrConversion(device VkDevice, ycbcrConversion VkSamplerYcbcrConversion, pAllocator *VkAllocationCallbacks) error {
	if l.funcPtrs.vkDestroySamplerYcbcrConversion == nil {
		return errors.New("attempted to call method 'vkDestroySamplerYcbcrConversion' which is not present on this loader")
	}

	C.cgoDestroySamplerYcbcrConversion(l.funcPtrs.vkDestroySamplerYcbcrConversion,
		C.VkDevice(device),
		C.VkSamplerYcbcrConversion(ycbcrConversion),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkCreateDescriptorUpdateTemplate(device VkDevice, pCreateInfo *VkDescriptorUpdateTemplateCreateInfo, pAllocator *VkAllocationCallbacks, pDescriptorUpdateTemplate *VkDescriptorUpdateTemplate) (VkResult, error) {
	if l.funcPtrs.vkCreateDescriptorUpdateTemplate == nil {
		return VKErrorUnknown, errors.New("attempted to call method 'vkCreateDescriptorUpdateTemplate' which is not present on this loader")
	}

	res := VkResult(C.cgoCreateDescriptorUpdateTemplate(l.funcPtrs.vkCreateDescriptorUpdateTemplate,
		C.VkDevice(device),
		(*C.VkDescriptorUpdateTemplateCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkDescriptorUpdateTemplate)(pDescriptorUpdateTemplate)))

	return res, res.ToError()
}
func (l *VulkanLoader) VkDestroyDescriptorUpdateTemplate(device VkDevice, descriptorUpdateTemplate VkDescriptorUpdateTemplate, pAllocator *VkAllocationCallbacks) error {
	if l.funcPtrs.vkDestroyDescriptorUpdateTemplate == nil {
		return errors.New("attempted to call method 'vkDestroyDescriptorUpdateTemplate' which is not present on this loader")
	}

	C.cgoDestroyDescriptorUpdateTemplate(l.funcPtrs.vkDestroyDescriptorUpdateTemplate,
		C.VkDevice(device),
		C.VkDescriptorUpdateTemplate(descriptorUpdateTemplate),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *VulkanLoader) VkUpdateDescriptorSetWithTemplate(device VkDevice, descriptorSet VkDescriptorSet, descriptorUpdateTemplate VkDescriptorUpdateTemplate, pData unsafe.Pointer) error {
	if l.funcPtrs.vkUpdateDescriptorSetWithTemplate == nil {
		return errors.New("attempted to call method 'vkUpdateDescriptorSetWithTemplate' which is not present on this loader")
	}

	C.cgoUpdateDescriptorSetWithTemplate(l.funcPtrs.vkUpdateDescriptorSetWithTemplate,
		C.VkDevice(device),
		C.VkDescriptorSet(descriptorSet),
		C.VkDescriptorUpdateTemplate(descriptorUpdateTemplate),
		pData)

	return nil
}

func (l *VulkanLoader) VkGetDescriptorSetLayoutSupport(device VkDevice, pCreateInfo *VkDescriptorSetLayoutCreateInfo, pSupport *VkDescriptorSetLayoutSupport) error {
	if l.funcPtrs.vkGetDescriptorSetLayoutSupport == nil {
		return errors.New("attempted to call method 'vkGetDescriptorSetLayoutSupport' which is not present on this loader")
	}

	C.cgoGetDescriptorSetLayoutSupport(l.funcPtrs.vkGetDescriptorSetLayoutSupport,
		C.VkDevice(device),
		(*C.VkDescriptorSetLayoutCreateInfo)(pCreateInfo),
		(*C.VkDescriptorSetLayoutSupport)(pSupport))

	return nil
}

func (l *VulkanLoader) VkCmdDrawIndirectCount(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, countBuffer VkBuffer, countBufferOffset VkDeviceSize, maxDrawCount Uint32, stride Uint32) error {
	if l.funcPtrs.vkCmdDrawIndirectCount == nil {
		return errors.New("attempted to call method 'vkCmdDrawIndirectCount' which is not present on this loader")
	}

	C.cgoCmdDrawIndirectCount(l.funcPtrs.vkCmdDrawIndirectCount,
		C.VkCommandBuffer(commandBuffer),
		C.VkBuffer(buffer),
		C.VkDeviceSize(offset),
		C.VkBuffer(countBuffer),
		C.VkDeviceSize(countBufferOffset),
		C.uint32_t(maxDrawCount),
		C.uint32_t(stride))

	return nil
}

func (l *VulkanLoader) VkCmdDrawIndexedIndirectCount(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, countBuffer VkBuffer, countBufferOffset VkDeviceSize, maxDrawCount Uint32, stride Uint32) error {
	if l.funcPtrs.vkCmdDrawIndexedIndirectCount == nil {
		return errors.New("attempted to call method 'vkCmdDrawIndexedIndirectCount' which is not present on this loader")
	}

	C.cgoCmdDrawIndexedIndirectCount(l.funcPtrs.vkCmdDrawIndexedIndirectCount,
		C.VkCommandBuffer(commandBuffer),
		C.VkBuffer(buffer),
		C.VkDeviceSize(offset),
		C.VkBuffer(countBuffer),
		C.VkDeviceSize(countBufferOffset),
		C.uint32_t(maxDrawCount),
		C.uint32_t(stride))

	return nil
}

func (l *VulkanLoader) VkCreateRenderPass2(device VkDevice, pCreateInfo *VkRenderPassCreateInfo2, pAllocator *VkAllocationCallbacks, pRenderPass *VkRenderPass) (VkResult, error) {
	if l.funcPtrs.vkCreateRenderPass2 == nil {
		return VKErrorUnknown, errors.New("attempted to call method 'vkCreateRenderPass2' which is not present on this loader")
	}

	res := VkResult(C.cgoCreateRenderPass2(l.funcPtrs.vkCreateRenderPass2,
		C.VkDevice(device),
		(*C.VkRenderPassCreateInfo2)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkRenderPass)(pRenderPass)))

	return res, res.ToError()
}

func (l *VulkanLoader) VkCmdBeginRenderPass2(commandBuffer VkCommandBuffer, pRenderPassBegin *VkRenderPassBeginInfo, pSubpassBeginInfo *VkSubpassBeginInfo) error {
	if l.funcPtrs.vkCmdBeginRenderPass2 == nil {
		return errors.New("attempted to call method 'vkCmdBeginRenderPass2' which is not present on this loader")
	}

	C.cgoCmdBeginRenderPass2(l.funcPtrs.vkCmdBeginRenderPass2,
		C.VkCommandBuffer(commandBuffer),
		(*C.VkRenderPassBeginInfo)(pRenderPassBegin),
		(*C.VkSubpassBeginInfo)(pSubpassBeginInfo))

	return nil
}

func (l *VulkanLoader) VkCmdNextSubpass2(commandBuffer VkCommandBuffer, pSubpassBeginInfo *VkSubpassBeginInfo, pSubpassEndInfo *VkSubpassEndInfo) error {
	if l.funcPtrs.vkCmdNextSubpass2 == nil {
		return errors.New("attempted to call method 'vkCmdNextSubpass2' which is not present on this loader")
	}

	C.cgoCmdNextSubpass2(l.funcPtrs.vkCmdNextSubpass2,
		C.VkCommandBuffer(commandBuffer),
		(*C.VkSubpassBeginInfo)(pSubpassBeginInfo),
		(*C.VkSubpassEndInfo)(pSubpassEndInfo))

	return nil
}

func (l *VulkanLoader) VkCmdEndRenderPass2(commandBuffer VkCommandBuffer, pSubpassEndInfo *VkSubpassEndInfo) error {
	if l.funcPtrs.vkCmdEndRenderPass2 == nil {
		return errors.New("attempted to call method 'vkCmdEndRenderPass2' which is not present on this loader")
	}

	C.cgoCmdEndRenderPass2(l.funcPtrs.vkCmdEndRenderPass2,
		C.VkCommandBuffer(commandBuffer),
		(*C.VkSubpassEndInfo)(pSubpassEndInfo))

	return nil
}

func (l *VulkanLoader) VkResetQueryPool(device VkDevice, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32) error {
	if l.funcPtrs.vkResetQueryPool == nil {
		return errors.New("attempted to call method 'vkResetQueryPool' which is not present on this loader")
	}

	C.cgoResetQueryPool(l.funcPtrs.vkResetQueryPool,
		C.VkDevice(device),
		C.VkQueryPool(queryPool),
		C.uint32_t(firstQuery),
		C.uint32_t(queryCount))

	return nil
}

func (l *VulkanLoader) VkGetSemaphoreCounterValue(device VkDevice, semaphore VkSemaphore, pValue *Uint64) (VkResult, error) {
	if l.funcPtrs.vkGetSemaphoreCounterValue == nil {
		return VKErrorUnknown, errors.New("attempted to call method 'vkGetSemaphoreCounterValue' which is not present on this loader")
	}

	res := VkResult(C.cgoGetSemaphoreCounterValue(l.funcPtrs.vkGetSemaphoreCounterValue,
		C.VkDevice(device),
		C.VkSemaphore(semaphore),
		(*C.uint64_t)(pValue)))

	return res, res.ToError()
}

func (l *VulkanLoader) VkWaitSemaphores(device VkDevice, pWaitInfo *VkSemaphoreWaitInfo, timeout Uint64) (VkResult, error) {
	if l.funcPtrs.vkWaitSemaphores == nil {
		return VKErrorUnknown, errors.New("attempted to call method 'vkWaitSemaphores' which is not present on this loader")
	}

	res := VkResult(C.cgoWaitSemaphores(l.funcPtrs.vkWaitSemaphores,
		C.VkDevice(device),
		(*C.VkSemaphoreWaitInfo)(pWaitInfo),
		C.uint64_t(timeout)))

	return res, res.ToError()
}

func (l *VulkanLoader) VkSignalSemaphore(device VkDevice, pSignalInfo *VkSemaphoreSignalInfo) (VkResult, error) {
	if l.funcPtrs.vkSignalSemaphore == nil {
		return VKErrorUnknown, errors.New("attempted to call method 'vkSignalSemaphore' which is not present on this loader")
	}

	res := VkResult(C.cgoSignalSemaphore(l.funcPtrs.vkSignalSemaphore,
		C.VkDevice(device),
		(*C.VkSemaphoreSignalInfo)(pSignalInfo)))

	return res, res.ToError()
}

func (l *VulkanLoader) VkGetBufferDeviceAddress(device VkDevice, pInfo *VkBufferDeviceAddressInfo) (VkDeviceAddress, error) {
	if l.funcPtrs.vkGetBufferDeviceAddress == nil {
		return VkDeviceAddress(0), errors.New("attempted to call method 'vkGetBufferDeviceAddress' which is not present on this loader")
	}

	address := VkDeviceAddress(C.cgoGetBufferDeviceAddress(l.funcPtrs.vkGetBufferDeviceAddress,
		C.VkDevice(device),
		(*C.VkBufferDeviceAddressInfo)(pInfo)))

	return address, nil
}

func (l *VulkanLoader) VkGetBufferOpaqueCaptureAddress(device VkDevice, pInfo *VkBufferDeviceAddressInfo) (Uint64, error) {
	if l.funcPtrs.vkGetBufferOpaqueCaptureAddress == nil {
		return 0, errors.New("attempted to call method 'vkGetBufferOpaqueCaptureAddress' which is not present on this loader")
	}

	address := Uint64(C.cgoGetBufferOpaqueCaptureAddress(l.funcPtrs.vkGetBufferOpaqueCaptureAddress,
		C.VkDevice(device),
		(*C.VkBufferDeviceAddressInfo)(pInfo)))

	return address, nil
}

func (l *VulkanLoader) VkGetDeviceMemoryOpaqueCaptureAddress(device VkDevice, pInfo *VkDeviceMemoryOpaqueCaptureAddressInfo) (Uint64, error) {
	if l.funcPtrs.vkGetDeviceMemoryOpaqueCaptureAddress == nil {
		return 0, errors.New("attempted to call method 'vkGetDeviceMemoryOpaqueCaptureAddress' which is not present on this loader")
	}

	address := Uint64(C.cgoGetDeviceMemoryOpaqueCaptureAddress(l.funcPtrs.vkGetDeviceMemoryOpaqueCaptureAddress,
		C.VkDevice(device),
		(*C.VkDeviceMemoryOpaqueCaptureAddressInfo)(pInfo)))

	return address, nil
}
