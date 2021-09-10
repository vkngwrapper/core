package loader

/*
#cgo windows LDFLAGS: -lvulkan
#cgo linux freebsd darwin openbsd pkg-config: vulkan
#include "loader.h"
*/
import "C"
import (
	"github.com/cockroachdb/errors"
	"unsafe"
)

func (l *Loader) VkEnumerateInstanceExtensionProperties(pLayerName *Char, pPropertyCount *Uint32, pProperties *VkExtensionProperties) (VkResult, error) {
	res := VkResult(C.cgoEnumerateInstanceExtensionProperties(l.funcPtrs.vkEnumerateInstanceExtensionProperties,
		(*C.char)(pLayerName),
		(*C.uint32_t)(pPropertyCount),
		(*C.VkExtensionProperties)(pProperties)))
	return res, res.ToError()
}

func (l *Loader) VkEnumerateInstanceLayerProperties(pPropertyCount *Uint32, pProperties *VkLayerProperties) (VkResult, error) {
	res := VkResult(C.cgoEnumerateInstanceLayerProperties(l.funcPtrs.vkEnumerateInstanceLayerProperties,
		(*C.uint32_t)(pPropertyCount),
		(*C.VkLayerProperties)(pProperties)))
	return res, res.ToError()
}

func (l *Loader) VkCreateInstance(pCreateInfo *VkInstanceCreateInfo, pAllocator *VkAllocationCallbacks, pInstance *VkInstance) (VkResult, error) {
	res := VkResult(C.cgoCreateInstance(l.funcPtrs.vkCreateInstance,
		(*C.VkInstanceCreateInfo)(pCreateInfo),
		(*C.VkAllocationCallbacks)(pAllocator),
		(*C.VkInstance)(pInstance)))
	return res, res.ToError()
}

func (l *Loader) VkEnumeratePhysicalDevices(instance VkInstance, pPhysicalDeviceCount *Uint32, pPhysicalDevices *VkPhysicalDevice) (VkResult, error) {
	if l.instance == nil {
		return VKErrorUnknown, errors.New("attempted to call instance loader function on a basic loader")
	}

	res := VkResult(C.cgoEnumeratePhysicalDevices(l.funcPtrs.vkEnumeratePhysicalDevices,
		(C.VkInstance)(instance),
		(*C.uint32_t)(pPhysicalDeviceCount),
		(*C.VkPhysicalDevice)(pPhysicalDevices)))
	return res, res.ToError()
}

func (l *Loader) VkDestroyInstance(instance VkInstance, pAllocator *VkAllocationCallbacks) error {
	if l.instance == nil {
		return errors.New("attempted to call instance loader function on a basic loader")
	}

	C.cgoDestroyInstance(l.funcPtrs.vkDestroyInstance,
		(C.VkInstance)(instance),
		(*C.VkAllocationCallbacks)(pAllocator))
	return nil
}

func (l *Loader) VkGetPhysicalDeviceFeatures(physicalDevice VkPhysicalDevice, pFeatures *VkPhysicalDeviceFeatures) error {
	if l.instance == nil {
		return errors.New("attempted to call instance loader function on a basic loader")
	}

	C.cgoGetPhysicalDeviceFeatures(l.funcPtrs.vkGetPhysicalDeviceFeatures,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.VkPhysicalDeviceFeatures)(pFeatures))
	return nil
}

func (l *Loader) VkGetPhysicalDeviceFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, pFormatProperties *VkFormatProperties) error {
	if l.instance == nil {
		return errors.New("attempted to call instance loader function on a basic loader")
	}

	C.cgoGetPhysicalDeviceFormatProperties(l.funcPtrs.vkGetPhysicalDeviceFormatProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(C.VkFormat)(format),
		(*C.VkFormatProperties)(pFormatProperties))
	return nil
}

func (l *Loader) VkGetPhysicalDeviceImageFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, t VkImageType, tiling VkImageTiling, usage VkImageUsageFlags, flags VkImageCreateFlags, pImageFormatProperties *VkImageFormatProperties) (VkResult, error) {
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

func (l *Loader) VkGetPhysicalDeviceProperties(physicalDevice VkPhysicalDevice, pProperties *VkPhysicalDeviceProperties) error {
	if l.instance == nil {
		return errors.New("attempted to call instance loader function on a basic loader")
	}

	C.cgoGetPhysicalDeviceProperties(l.funcPtrs.vkGetPhysicalDeviceProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.VkPhysicalDeviceProperties)(pProperties))
	return nil
}

func (l *Loader) VkGetPhysicalDeviceQueueFamilyProperties(physicalDevice VkPhysicalDevice, pQueueFamilyPropertyCount *Uint32, pQueueFamilyProperties *VkQueueFamilyProperties) error {
	if l.instance == nil {
		return errors.New("attempted to call instance loader function on a basic loader")
	}

	C.cgoGetPhysicalDeviceQueueFamilyProperties(l.funcPtrs.vkGetPhysicalDeviceQueueFamilyProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.uint32_t)(pQueueFamilyPropertyCount),
		(*C.VkQueueFamilyProperties)(pQueueFamilyProperties))
	return nil
}

func (l *Loader) VkGetPhysicalDeviceMemoryProperties(physicalDevice VkPhysicalDevice, pMemoryProperties *VkPhysicalDeviceMemoryProperties) error {
	if l.instance == nil {
		return errors.New("attempted to call instance loader function on a basic loader")
	}

	C.cgoGetPhysicalDeviceMemoryProperties(l.funcPtrs.vkGetPhysicalDeviceMemoryProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.VkPhysicalDeviceMemoryProperties)(pMemoryProperties))
	return nil
}

func (l *Loader) VkEnumerateDeviceExtensionProperties(physicalDevice VkPhysicalDevice, pLayerName *Char, pPropertyCount *Uint32, pProperties *VkExtensionProperties) (VkResult, error) {
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

func (l *Loader) VkEnumerateDeviceLayerProperties(physicalDevice VkPhysicalDevice, pPropertyCount *Uint32, pProperties *VkLayerProperties) (VkResult, error) {
	if l.instance == nil {
		return VKErrorUnknown, errors.New("attempted to call instance loader function on a basic loader")
	}

	res := VkResult(C.cgoEnumerateDeviceLayerProperties(l.funcPtrs.vkEnumerateDeviceLayerProperties,
		(C.VkPhysicalDevice)(physicalDevice),
		(*C.uint32_t)(pPropertyCount),
		(*C.VkLayerProperties)(pProperties)))
	return res, res.ToError()
}

func (l *Loader) VkGetPhysicalDeviceSparseImageFormatProperties(physicalDevice VkPhysicalDevice, format VkFormat, t VkImageType, samples VkSampleCountFlagBits, usage VkImageUsageFlags, tiling VkImageTiling, pPropertyCount *Uint32, pProperties *VkSparseImageFormatProperties) error {
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

func (l *Loader) VkCreateDevice(physicalDevice VkPhysicalDevice, pCreateInfo *VkDeviceCreateInfo, pAllocator *VkAllocationCallbacks, pDevice *VkDevice) (VkResult, error) {
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

func (l *Loader) VkDestroyDevice(device VkDevice, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyDevice(l.funcPtrs.vkDestroyDevice,
		(C.VkDevice)(device),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkGetDeviceQueue(device VkDevice, queueFamilyIndex Uint32, queueIndex Uint32, pQueue *VkQueue) error {
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

func (l *Loader) VkQueueSubmit(queue VkQueue, submitCount Uint32, pSubmits *VkSubmitInfo, fence VkFence) (VkResult, error) {
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

func (l *Loader) VkQueueWaitIdle(queue VkQueue) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoQueueWaitIdle(l.funcPtrs.vkQueueWaitIdle,
		(C.VkQueue)(queue)))
	return res, res.ToError()
}

func (l *Loader) VkDeviceWaitIdle(device VkDevice) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoDeviceWaitIdle(l.funcPtrs.vkDeviceWaitIdle,
		(C.VkDevice)(device)))
	return res, res.ToError()
}

func (l *Loader) VkAllocateMemory(device VkDevice, pAllocateInfo *VkMemoryAllocateInfo, pAllocator *VkAllocationCallbacks, pMemory *VkDeviceMemory) (VkResult, error) {
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

func (l *Loader) VkFreeMemory(device VkDevice, memory VkDeviceMemory, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoFreeMemory(l.funcPtrs.vkFreeMemory,
		(C.VkDevice)(device),
		(C.VkDeviceMemory)(memory),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkMapMemory(device VkDevice, memory VkDeviceMemory, offset VkDeviceSize, size VkDeviceSize, flags VkMemoryMapFlags, ppData *unsafe.Pointer) (VkResult, error) {
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

func (l *Loader) VkUnmapMemory(device VkDevice, memory VkDeviceMemory) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoUnmapMemory(l.funcPtrs.vkUnmapMemory,
		(C.VkDevice)(device),
		(C.VkDeviceMemory)(memory))

	return nil
}

func (l *Loader) VkFlushMappedMemoryRanges(device VkDevice, memoryRangeCount Uint32, pMemoryRanges *VkMappedMemoryRange) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoFlushMappedMemoryRanges(l.funcPtrs.vkFlushMappedMemoryRanges,
		(C.VkDevice)(device),
		(C.uint32_t)(memoryRangeCount),
		(*C.VkMappedMemoryRange)(pMemoryRanges)))
	return res, res.ToError()
}

func (l *Loader) VkInvalidateMappedMemoryRanges(device VkDevice, memoryRangeCount Uint32, pMemoryRanges *VkMappedMemoryRange) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoInvalidateMappedMemoryRanges(l.funcPtrs.vkInvalidateMappedMemoryRanges,
		(C.VkDevice)(device),
		(C.uint32_t)(memoryRangeCount),
		(*C.VkMappedMemoryRange)(pMemoryRanges)))
	return res, res.ToError()
}

func (l *Loader) VkGetDeviceMemoryCommitment(device VkDevice, memory VkDeviceMemory, pCommittedMemoryInBytes *VkDeviceSize) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoGetDeviceMemoryCommitment(l.funcPtrs.vkGetDeviceMemoryCommitment,
		(C.VkDevice)(device),
		(C.VkDeviceMemory)(memory),
		(*C.VkDeviceSize)(pCommittedMemoryInBytes))

	return nil
}

func (l *Loader) VkBindBufferMemory(device VkDevice, buffer VkBuffer, memory VkDeviceMemory, memoryOffset VkDeviceSize) (VkResult, error) {
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

func (l *Loader) VkBindImageMemory(device VkDevice, image VkImage, memory VkDeviceMemory, memoryOffset VkDeviceSize) (VkResult, error) {
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

func (l *Loader) VkGetBufferMemoryRequirements(device VkDevice, buffer VkBuffer, pMemoryRequirements *VkMemoryRequirements) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoGetBufferMemoryRequirements(l.funcPtrs.vkGetBufferMemoryRequirements,
		(C.VkDevice)(device),
		(C.VkBuffer)(buffer),
		(*C.VkMemoryRequirements)(pMemoryRequirements))

	return nil
}

func (l *Loader) VkGetImageMemoryRequirements(device VkDevice, image VkImage, pMemoryRequirements *VkMemoryRequirements) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoGetImageMemoryRequirements(l.funcPtrs.vkGetImageMemoryRequirements,
		(C.VkDevice)(device),
		(C.VkImage)(image),
		(*C.VkMemoryRequirements)(pMemoryRequirements))

	return nil
}

func (l *Loader) VkGetImageSparseMemoryRequirements(device VkDevice, image VkImage, pSparseMemoryRequirementCount *Uint32, pSparseMemoryRequirements *VkSparseImageMemoryRequirements) error {
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

func (l *Loader) VkQueueBindSparse(queue VkQueue, bindInfoCount Uint32, pBindInfo *VkBindSparseInfo, fence VkFence) (VkResult, error) {
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

func (l *Loader) VkCreateFence(device VkDevice, pCreateInfo *VkFenceCreateInfo, pAllocator *VkAllocationCallbacks, pFence *VkFence) (VkResult, error) {
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

func (l *Loader) VkDestroyFence(device VkDevice, fence VkFence, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyFence(l.funcPtrs.vkDestroyFence,
		(C.VkDevice)(device),
		(C.VkFence)(fence),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkResetFences(device VkDevice, fenceCount Uint32, pFences *VkFence) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoResetFences(l.funcPtrs.vkResetFences,
		(C.VkDevice)(device),
		(C.uint32_t)(fenceCount),
		(*C.VkFence)(pFences)))
	return res, res.ToError()
}

func (l *Loader) VkGetFenceStatus(device VkDevice, fence VkFence) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoGetFenceStatus(l.funcPtrs.vkGetFenceStatus,
		(C.VkDevice)(device),
		(C.VkFence)(fence)))
	return res, res.ToError()
}

func (l *Loader) VkWaitForFences(device VkDevice, fenceCount Uint32, pFences *VkFence, waitAll VkBool32, timeout Uint64) (VkResult, error) {
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

func (l *Loader) VkCreateSemaphore(device VkDevice, pCreateInfo *VkSemaphoreCreateInfo, pAllocator *VkAllocationCallbacks, pSemaphore *VkSemaphore) (VkResult, error) {
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

func (l *Loader) VkDestroySemaphore(device VkDevice, semaphore VkSemaphore, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroySemaphore(l.funcPtrs.vkDestroySemaphore,
		(C.VkDevice)(device),
		(C.VkSemaphore)(semaphore),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkCreateEvent(device VkDevice, pCreateInfo *VkEventCreateInfo, pAllocator *VkAllocationCallbacks, pEvent *VkEvent) (VkResult, error) {
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

func (l *Loader) VkDestroyEvent(device VkDevice, event VkEvent, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyEvent(l.funcPtrs.vkDestroyEvent,
		(C.VkDevice)(device),
		(C.VkEvent)(event),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkGetEventStatus(device VkDevice, event VkEvent) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoGetEventStatus(l.funcPtrs.vkGetEventStatus,
		(C.VkDevice)(device),
		(C.VkEvent)(event)))
	return res, res.ToError()
}

func (l *Loader) VkSetEvent(device VkDevice, event VkEvent) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoSetEvent(l.funcPtrs.vkSetEvent,
		(C.VkDevice)(device),
		(C.VkEvent)(event)))
	return res, res.ToError()
}

func (l *Loader) VkResetEvent(device VkDevice, event VkEvent) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoResetEvent(l.funcPtrs.vkResetEvent,
		(C.VkDevice)(device),
		(C.VkEvent)(event)))
	return res, res.ToError()
}

func (l *Loader) VkCreateQueryPool(device VkDevice, pCreateInfo *VkQueryPoolCreateInfo, pAllocator *VkAllocationCallbacks, pQueryPool *VkQueryPool) (VkResult, error) {
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

func (l *Loader) VkDestroyQueryPool(device VkDevice, queryPool VkQueryPool, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyQueryPool(l.funcPtrs.vkDestroyQueryPool,
		(C.VkDevice)(device),
		(C.VkQueryPool)(queryPool),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkGetQueryPoolResults(device VkDevice, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32, dataSize C.size_t, pData unsafe.Pointer, stride VkDeviceSize, flags VkQueryResultFlags) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoGetQueryPoolResults(l.funcPtrs.vkGetQueryPoolResults,
		(C.VkDevice)(device),
		(C.VkQueryPool)(queryPool),
		(C.uint32_t)(firstQuery),
		(C.uint32_t)(queryCount),
		dataSize,
		pData,
		(C.VkDeviceSize)(stride),
		(C.VkQueryResultFlags)(flags)))
	return res, res.ToError()
}

func (l *Loader) VkCreateBuffer(device VkDevice, pCreateInfo *VkBufferCreateInfo, pAllocator *VkAllocationCallbacks, pBuffer *VkBuffer) (VkResult, error) {
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

func (l *Loader) VkDestroyBuffer(device VkDevice, buffer VkBuffer, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyBuffer(l.funcPtrs.vkDestroyBuffer,
		(C.VkDevice)(device),
		(C.VkBuffer)(buffer),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkCreateBufferView(device VkDevice, pCreateInfo *VkBufferViewCreateInfo, pAllocator *VkAllocationCallbacks, pView *VkBufferView) (VkResult, error) {
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

func (l *Loader) VkDestroyBufferView(device VkDevice, bufferView VkBufferView, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyBufferView(l.funcPtrs.vkDestroyBufferView,
		(C.VkDevice)(device),
		(C.VkBufferView)(bufferView),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkCreateImage(device VkDevice, pCreateInfo *VkImageCreateInfo, pAllocator *VkAllocationCallbacks, pImage *VkImage) (VkResult, error) {
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

func (l *Loader) VkDestroyImage(device VkDevice, image VkImage, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyImage(l.funcPtrs.vkDestroyImage,
		(C.VkDevice)(device),
		(C.VkImage)(image),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkGetImageSubresourceLayout(device VkDevice, image VkImage, pSubresource *VkImageSubresource, pLayout *VkSubresourceLayout) error {
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

func (l *Loader) VkCreateImageView(device VkDevice, pCreateInfo *VkImageViewCreateInfo, pAllocator *VkAllocationCallbacks, pView *VkImageView) (VkResult, error) {
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

func (l *Loader) VkDestroyImageView(device VkDevice, imageView VkImageView, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyImageView(l.funcPtrs.vkDestroyImageView,
		(C.VkDevice)(device),
		(C.VkImageView)(imageView),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkCreateShaderModule(device VkDevice, pCreateInfo *VkShaderModuleCreateInfo, pAllocator *VkAllocationCallbacks, pShaderModule *VkShaderModule) (VkResult, error) {
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

func (l *Loader) VkDestroyShaderModule(device VkDevice, shaderModule VkShaderModule, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyShaderModule(l.funcPtrs.vkDestroyShaderModule,
		(C.VkDevice)(device),
		(C.VkShaderModule)(shaderModule),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkCreatePipelineCache(device VkDevice, pCreateInfo *VkPipelineCacheCreateInfo, pAllocator *VkAllocationCallbacks, pPipelineCache *VkPipelineCache) (VkResult, error) {
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

func (l *Loader) VkDestroyPipelineCache(device VkDevice, pipelineCache VkPipelineCache, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyPipelineCache(l.funcPtrs.vkDestroyPipelineCache,
		(C.VkDevice)(device),
		(C.VkPipelineCache)(pipelineCache),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkGetPipelineCacheData(device VkDevice, pipelineCache VkPipelineCache, pDataSize *C.size_t, pData unsafe.Pointer) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoGetPipelineCacheData(l.funcPtrs.vkGetPipelineCacheData,
		(C.VkDevice)(device),
		(C.VkPipelineCache)(pipelineCache),
		pDataSize,
		pData))
	return res, res.ToError()
}

func (l *Loader) VkMergePipelineCaches(device VkDevice, dstCache VkPipelineCache, srcCacheCount Uint32, pSrcCaches *VkPipelineCache) (VkResult, error) {
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

func (l *Loader) VkCreateGraphicsPipelines(device VkDevice, pipelineCache VkPipelineCache, createInfoCount Uint32, pCreateInfos *VkGraphicsPipelineCreateInfo, pAllocator *VkAllocationCallbacks, pPipelines *VkPipeline) (VkResult, error) {
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

func (l *Loader) VkCreateComputePipelines(device VkDevice, pipelineCache VkPipelineCache, createInfoCount Uint32, pCreateInfos *VkComputePipelineCreateInfo, pAllocator *VkAllocationCallbacks, pPipelines *VkPipeline) (VkResult, error) {
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

func (l *Loader) VkDestroyPipeline(device VkDevice, pipeline VkPipeline, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyPipeline(l.funcPtrs.vkDestroyPipeline,
		(C.VkDevice)(device),
		(C.VkPipeline)(pipeline),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkCreatePipelineLayout(device VkDevice, pCreateInfo *VkPipelineLayoutCreateInfo, pAllocator *VkAllocationCallbacks, pPipelineLayout *VkPipelineLayout) (VkResult, error) {
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

func (l *Loader) VkDestroyPipelineLayout(device VkDevice, pipelineLayout VkPipelineLayout, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyPipelineLayout(l.funcPtrs.vkDestroyPipelineLayout,
		(C.VkDevice)(device),
		(C.VkPipelineLayout)(pipelineLayout),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkCreateSampler(device VkDevice, pCreateInfo *VkSamplerCreateInfo, pAllocator *VkAllocationCallbacks, pSampler *VkSampler) (VkResult, error) {
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

func (l *Loader) VkDestroySampler(device VkDevice, sampler VkSampler, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroySampler(l.funcPtrs.vkDestroySampler,
		(C.VkDevice)(device),
		(C.VkSampler)(sampler),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkCreateDescriptorSetLayout(device VkDevice, pCreateInfo *VkDescriptorSetLayoutCreateInfo, pAllocator *VkAllocationCallbacks, pSetLayout *VkDescriptorSetLayout) (VkResult, error) {
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

func (l *Loader) VkDestroyDescriptorSetLayout(device VkDevice, descriptorSetLayout VkDescriptorSetLayout, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyDescriptorSetLayout(l.funcPtrs.vkDestroyDescriptorSetLayout,
		(C.VkDevice)(device),
		(C.VkDescriptorSetLayout)(descriptorSetLayout),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkCreateDescriptorPool(device VkDevice, pCreateInfo *VkDescriptorPoolCreateInfo, pAllocator *VkAllocationCallbacks, pDescriptorPool *VkDescriptorPool) (VkResult, error) {
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

func (l *Loader) VkDestroyDescriptorPool(device VkDevice, descriptorPool VkDescriptorPool, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyDescriptorPool(l.funcPtrs.vkDestroyDescriptorPool,
		(C.VkDevice)(device),
		(C.VkDescriptorPool)(descriptorPool),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkResetDescriptorPool(device VkDevice, descriptorPool VkDescriptorPool, flags VkDescriptorPoolResetFlags) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoResetDescriptorPool(l.funcPtrs.vkResetDescriptorPool,
		(C.VkDevice)(device),
		(C.VkDescriptorPool)(descriptorPool),
		(C.VkDescriptorPoolResetFlags)(flags)))
	return res, res.ToError()
}

func (l *Loader) VkAllocateDescriptorSets(device VkDevice, pAllocateInfo *VkDescriptorSetAllocateInfo, pDescriptorSets *VkDescriptorSet) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoAllocateDescriptorSets(l.funcPtrs.vkAllocateDescriptorSets,
		(C.VkDevice)(device),
		(*C.VkDescriptorSetAllocateInfo)(pAllocateInfo),
		(*C.VkDescriptorSet)(pDescriptorSets)))
	return res, res.ToError()
}

func (l *Loader) VkFreeDescriptorSets(device VkDevice, descriptorPool VkDescriptorPool, descriptorSetCount Uint32, pDescriptorSets *VkDescriptorSet) (VkResult, error) {
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

func (l *Loader) VkUpdateDescriptorSets(device VkDevice, descriptorWriteCount Uint32, pDescriptorWrites *VkWriteDescriptorSet, descriptorCopyCount Uint32, pDescriptorCopies *VkCopyDescriptorSet) error {
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

func (l *Loader) VkCreateFramebuffer(device VkDevice, pCreateInfo *VkFramebufferCreateInfo, pAllocator *VkAllocationCallbacks, pFramebuffer *VkFramebuffer) (VkResult, error) {
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

func (l *Loader) VkDestroyFramebuffer(device VkDevice, framebuffer VkFramebuffer, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyFramebuffer(l.funcPtrs.vkDestroyFramebuffer,
		(C.VkDevice)(device),
		(C.VkFramebuffer)(framebuffer),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkCreateRenderPass(device VkDevice, pCreateInfo *VkRenderPassCreateInfo, pAllocator *VkAllocationCallbacks, pRenderPass *VkRenderPass) (VkResult, error) {
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

func (l *Loader) VkDestroyRenderPass(device VkDevice, renderPass VkRenderPass, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyRenderPass(l.funcPtrs.vkDestroyRenderPass,
		(C.VkDevice)(device),
		(C.VkRenderPass)(renderPass),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkGetRenderAreaGranularity(device VkDevice, renderPass VkRenderPass, pGranularity *VkExtent2D) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoGetRenderAreaGranularity(l.funcPtrs.vkGetRenderAreaGranularity,
		(C.VkDevice)(device),
		(C.VkRenderPass)(renderPass),
		(*C.VkExtent2D)(pGranularity))

	return nil
}

func (l *Loader) VkCreateCommandPool(device VkDevice, pCreateInfo *VkCommandPoolCreateInfo, pAllocator *VkAllocationCallbacks, pCommandPool *VkCommandPool) (VkResult, error) {
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

func (l *Loader) VkDestroyCommandPool(device VkDevice, commandPool VkCommandPool, pAllocator *VkAllocationCallbacks) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoDestroyCommandPool(l.funcPtrs.vkDestroyCommandPool,
		(C.VkDevice)(device),
		(C.VkCommandPool)(commandPool),
		(*C.VkAllocationCallbacks)(pAllocator))

	return nil
}

func (l *Loader) VkResetCommandPool(device VkDevice, commandPool VkCommandPool, flags VkCommandPoolResetFlags) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoResetCommandPool(l.funcPtrs.vkResetCommandPool,
		(C.VkDevice)(device),
		(C.VkCommandPool)(commandPool),
		(C.VkCommandPoolResetFlags)(flags)))
	return res, res.ToError()
}

func (l *Loader) VkAllocateCommandBuffers(device VkDevice, pAllocateInfo *VkCommandBufferAllocateInfo, pCommandBuffers *VkCommandBuffer) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoAllocateCommandBuffers(l.funcPtrs.vkAllocateCommandBuffers,
		(C.VkDevice)(device),
		(*C.VkCommandBufferAllocateInfo)(pAllocateInfo),
		(*C.VkCommandBuffer)(pCommandBuffers)))
	return res, res.ToError()
}

func (l *Loader) VkFreeCommandBuffers(device VkDevice, commandPool VkCommandPool, commandBufferCount Uint32, pCommandBuffers *VkCommandBuffer) error {
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

func (l *Loader) VkBeginCommandBuffer(commandBuffer VkCommandBuffer, pBeginInfo *VkCommandBufferBeginInfo) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoBeginCommandBuffer(l.funcPtrs.vkBeginCommandBuffer,
		(C.VkCommandBuffer)(commandBuffer),
		(*C.VkCommandBufferBeginInfo)(pBeginInfo)))
	return res, res.ToError()
}

func (l *Loader) VkEndCommandBuffer(commandBuffer VkCommandBuffer) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoEndCommandBuffer(l.funcPtrs.vkEndCommandBuffer,
		(C.VkCommandBuffer)(commandBuffer)))
	return res, res.ToError()
}

func (l *Loader) VkResetCommandBuffer(commandBuffer VkCommandBuffer, flags VkCommandBufferResetFlags) (VkResult, error) {
	if l.device == nil {
		return VKErrorUnknown, errors.New("attempted device loader function on a non-device loader")
	}

	res := VkResult(C.cgoResetCommandBuffer(l.funcPtrs.vkResetCommandBuffer,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkCommandBufferResetFlags)(flags)))
	return res, res.ToError()
}

func (l *Loader) VkCmdBindPipeline(commandBuffer VkCommandBuffer, pipelineBindPoint VkPipelineBindPoint, pipeline VkPipeline) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdBindPipeline(l.funcPtrs.vkCmdBindPipeline,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkPipelineBindPoint)(pipelineBindPoint),
		(C.VkPipeline)(pipeline))

	return nil
}

func (l *Loader) VkCmdSetViewport(commandBuffer VkCommandBuffer, firstViewport Uint32, viewportCount Uint32, pViewports *VkViewport) error {
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

func (l *Loader) VkCmdSetScissor(commandBuffer VkCommandBuffer, firstScissor Uint32, scissorCount Uint32, pScissors *VkRect2D) error {
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

func (l *Loader) VkCmdSetLineWidth(commandBuffer VkCommandBuffer, lineWidth Float) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}
	C.cgoCmdSetLineWidth(l.funcPtrs.vkCmdSetLineWidth,
		(C.VkCommandBuffer)(commandBuffer),
		(C.float)(lineWidth))

	return nil
}

func (l *Loader) VkCmdSetDepthBias(commandBuffer VkCommandBuffer, depthBiasConstantFactor Float, depthBiasClamp Float, depthBiasSlopeFactor Float) error {
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

func (l *Loader) VkCmdSetBlendConstants(commandBuffer VkCommandBuffer, blendConstants *Float) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdSetBlendConstants(l.funcPtrs.vkCmdSetBlendConstants,
		(C.VkCommandBuffer)(commandBuffer),
		(*C.float)(blendConstants),
	)

	return nil
}

func (l *Loader) VkCmdSetDepthBounds(commandBuffer VkCommandBuffer, minDepthBounds Float, maxDepthBounds Float) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdSetDepthBounds(l.funcPtrs.vkCmdSetDepthBounds,
		(C.VkCommandBuffer)(commandBuffer),
		(C.float)(minDepthBounds),
		(C.float)(maxDepthBounds))

	return nil
}

func (l *Loader) VkCmdSetStencilCompareMask(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, compareMask Uint32) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdSetStencilCompareMask(l.funcPtrs.vkCmdSetStencilCompareMask,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkStencilFaceFlags)(faceMask),
		(C.uint32_t)(compareMask))

	return nil
}

func (l *Loader) VkCmdSetStencilWriteMask(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, writeMask Uint32) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdSetStencilWriteMask(l.funcPtrs.vkCmdSetStencilWriteMask,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkStencilFaceFlags)(faceMask),
		(C.uint32_t)(writeMask))

	return nil
}

func (l *Loader) VkCmdSetStencilReference(commandBuffer VkCommandBuffer, faceMask VkStencilFaceFlags, reference Uint32) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdSetStencilReference(l.funcPtrs.vkCmdSetStencilReference,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkStencilFaceFlags)(faceMask),
		(C.uint32_t)(reference))

	return nil
}

func (l *Loader) VkCmdBindDescriptorSets(commandBuffer VkCommandBuffer, pipelineBindPoint VkPipelineBindPoint, layout VkPipelineLayout, firstSet Uint32, descriptorSetCount Uint32, pDescriptorSets *VkDescriptorSet, dynamicOffsetCount Uint32, pDynamicOffsets *Uint32) error {
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

func (l *Loader) VkCmdBindIndexBuffer(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, indexType VkIndexType) error {
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

func (l *Loader) VkCmdBindVertexBuffers(commandBuffer VkCommandBuffer, firstBinding Uint32, bindingCount Uint32, pBuffers *VkBuffer, pOffsets *VkDeviceSize) error {
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

func (l *Loader) VkCmdDraw(commandBuffer VkCommandBuffer, vertexCount Uint32, instanceCount Uint32, firstVertex Uint32, firstInstance Uint32) error {
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

func (l *Loader) VkCmdDrawIndexed(commandBuffer VkCommandBuffer, indexCount Uint32, instanceCount Uint32, firstIndex Uint32, vertexOffset Int32, firstInstance Uint32) error {
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

func (l *Loader) VkCmdDrawIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, drawCount Uint32, stride Uint32) error {
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

func (l *Loader) VkCmdDrawIndexedIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize, drawCount Uint32, stride Uint32) error {
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

func (l *Loader) VkCmdDispatch(commandBuffer VkCommandBuffer, groupCountX Uint32, groupCountY Uint32, groupCountZ Uint32) error {
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

func (l *Loader) VkCmdDispatchIndirect(commandBuffer VkCommandBuffer, buffer VkBuffer, offset VkDeviceSize) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdDispatchIndirect(l.funcPtrs.vkCmdDispatchIndirect,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkBuffer)(buffer),
		(C.VkDeviceSize)(offset))

	return nil
}

func (l *Loader) VkCmdCopyBuffer(commandBuffer VkCommandBuffer, srcBuffer VkBuffer, dstBuffer VkBuffer, regionCount Uint32, pRegions *VkBufferCopy) error {
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

func (l *Loader) VkCmdCopyImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkImageCopy) error {
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

func (l *Loader) VkCmdBlitImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkImageBlit, filter VkFilter) error {
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

func (l *Loader) VkCmdCopyBufferToImage(commandBuffer VkCommandBuffer, srcBuffer VkBuffer, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkBufferImageCopy) error {
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

func (l *Loader) VkCmdCopyImageToBuffer(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstBuffer VkBuffer, regionCount Uint32, pRegions *VkBufferImageCopy) error {
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

func (l *Loader) VkCmdUpdateBuffer(commandBuffer VkCommandBuffer, dstBuffer VkBuffer, dstOffset VkDeviceSize, dataSize VkDeviceSize, pData unsafe.Pointer) error {
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

func (l *Loader) VkCmdFillBuffer(commandBuffer VkCommandBuffer, dstBuffer VkBuffer, dstOffset VkDeviceSize, size VkDeviceSize, data Uint32) error {
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

func (l *Loader) VkCmdClearColorImage(commandBuffer VkCommandBuffer, image VkImage, imageLayout VkImageLayout, pColor *VkClearColorValue, rangeCount Uint32, pRanges *VkImageSubresourceRange) error {
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

func (l *Loader) VkCmdClearDepthStencilImage(commandBuffer VkCommandBuffer, image VkImage, imageLayout VkImageLayout, pDepthStencil *VkClearDepthStencilValue, rangeCount Uint32, pRanges *VkImageSubresourceRange) error {
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

func (l *Loader) VkCmdClearAttachments(commandBuffer VkCommandBuffer, attachmentCount Uint32, pAttachments *VkClearAttachment, rectCount Uint32, pRects *VkClearRect) error {
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

func (l *Loader) VkCmdResolveImage(commandBuffer VkCommandBuffer, srcImage VkImage, srcImageLayout VkImageLayout, dstImage VkImage, dstImageLayout VkImageLayout, regionCount Uint32, pRegions *VkImageResolve) error {
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

func (l *Loader) VkCmdSetEvent(commandBuffer VkCommandBuffer, event VkEvent, stageMask VkPipelineStageFlags) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdSetEvent(l.funcPtrs.vkCmdSetEvent,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkEvent)(event),
		(C.VkPipelineStageFlags)(stageMask))

	return nil
}

func (l *Loader) VkCmdResetEvent(commandBuffer VkCommandBuffer, event VkEvent, stageMask VkPipelineStageFlags) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdResetEvent(l.funcPtrs.vkCmdResetEvent,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkEvent)(event),
		(C.VkPipelineStageFlags)(stageMask))

	return nil
}

func (l *Loader) VkCmdWaitEvents(commandBuffer VkCommandBuffer, eventCount Uint32, pEvents *VkEvent, srcStageMask VkPipelineStageFlags, dstStageMask VkPipelineStageFlags, memoryBarrierCount Uint32, pMemoryBarriers *VkMemoryBarrier, bufferMemoryBarrierCount Uint32, pBufferMemoryBarriers *VkBufferMemoryBarrier, imageMemoryBarrierCount Uint32, pImageMemoryBarriers *VkImageMemoryBarrier) error {
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

func (l *Loader) VkCmdPipelineBarrier(commandBuffer VkCommandBuffer, srcStageMask VkPipelineStageFlags, dstStageMask VkPipelineStageFlags, dependencyFlags VkDependencyFlags, memoryBarrierCount Uint32, pMemoryBarriers *VkMemoryBarrier, bufferMemoryBarrierCount Uint32, pBufferMemoryBarriers *VkBufferMemoryBarrier, imageMemoryBarrierCount Uint32, pImageMemoryBarriers *VkImageMemoryBarrier) error {
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

func (l *Loader) VkCmdBeginQuery(commandBuffer VkCommandBuffer, queryPool VkQueryPool, query Uint32, flags VkQueryControlFlags) error {
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

func (l *Loader) VkCmdEndQuery(commandBuffer VkCommandBuffer, queryPool VkQueryPool, query Uint32) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdEndQuery(l.funcPtrs.vkCmdEndQuery,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkQueryPool)(queryPool),
		(C.uint32_t)(query))

	return nil
}

func (l *Loader) VkCmdResetQueryPool(commandBuffer VkCommandBuffer, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32) error {
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

func (l *Loader) VkCmdWriteTimestamp(commandBuffer VkCommandBuffer, pipelineStage VkPipelineStageFlags, queryPool VkQueryPool, query Uint32) error {
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

func (l *Loader) VkCmdCopyQueryPoolResults(commandBuffer VkCommandBuffer, queryPool VkQueryPool, firstQuery Uint32, queryCount Uint32, dstBuffer VkBuffer, dstOffset VkDeviceSize, stride VkDeviceSize, flags VkQueryResultFlags) error {
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

func (l *Loader) VkCmdPushConstants(commandBuffer VkCommandBuffer, layout VkPipelineLayout, stageFlags VkShaderStageFlags, offset Uint32, size Uint32, pValues unsafe.Pointer) error {
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

func (l *Loader) VkCmdBeginRenderPass(commandBuffer VkCommandBuffer, pRenderPassBegin *VkRenderPassBeginInfo, contents VkSubpassContents) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdBeginRenderPass(l.funcPtrs.vkCmdBeginRenderPass,
		(C.VkCommandBuffer)(commandBuffer),
		(*C.VkRenderPassBeginInfo)(pRenderPassBegin),
		(C.VkSubpassContents)(contents))

	return nil
}

func (l *Loader) VkCmdNextSubpass(commandBuffer VkCommandBuffer, contents VkSubpassContents) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdNextSubpass(l.funcPtrs.vkCmdNextSubpass,
		(C.VkCommandBuffer)(commandBuffer),
		(C.VkSubpassContents)(contents))

	return nil
}

func (l *Loader) VkCmdEndRenderPass(commandBuffer VkCommandBuffer) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdEndRenderPass(l.funcPtrs.vkCmdEndRenderPass,
		(C.VkCommandBuffer)(commandBuffer))

	return nil
}

func (l *Loader) VkCmdExecuteCommands(commandBuffer VkCommandBuffer, commandBufferCount Uint32, pCommandBuffers *VkCommandBuffer) error {
	if l.device == nil {
		return errors.New("attempted device loader function on a non-device loader")
	}

	C.cgoCmdExecuteCommands(l.funcPtrs.vkCmdExecuteCommands,
		(C.VkCommandBuffer)(commandBuffer),
		(C.uint32_t)(commandBufferCount),
		(*C.VkCommandBuffer)(pCommandBuffers))

	return nil
}
