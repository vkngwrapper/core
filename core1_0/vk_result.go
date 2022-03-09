package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

const (
	VKSuccess                   common.VkResult = C.VK_SUCCESS
	VKNotReady                  common.VkResult = C.VK_NOT_READY
	VKTimeout                   common.VkResult = C.VK_TIMEOUT
	VKEventSet                  common.VkResult = C.VK_EVENT_SET
	VKEventReset                common.VkResult = C.VK_EVENT_RESET
	VKIncomplete                common.VkResult = C.VK_INCOMPLETE
	VKErrorOutOfHostMemory      common.VkResult = C.VK_ERROR_OUT_OF_HOST_MEMORY
	VKErrorOutOfDeviceMemory    common.VkResult = C.VK_ERROR_OUT_OF_DEVICE_MEMORY
	VKErrorInitializationFailed common.VkResult = C.VK_ERROR_INITIALIZATION_FAILED
	VKErrorDeviceLost           common.VkResult = C.VK_ERROR_DEVICE_LOST
	VKErrorMemoryMapFailed      common.VkResult = C.VK_ERROR_MEMORY_MAP_FAILED
	VKErrorLayerNotPresent      common.VkResult = C.VK_ERROR_LAYER_NOT_PRESENT
	VKErrorExtensionNotPresent  common.VkResult = C.VK_ERROR_EXTENSION_NOT_PRESENT
	VKErrorFeatureNotPresent    common.VkResult = C.VK_ERROR_FEATURE_NOT_PRESENT
	VKErrorIncompatibleDriver   common.VkResult = C.VK_ERROR_INCOMPATIBLE_DRIVER
	VKErrorTooManyObjects       common.VkResult = C.VK_ERROR_TOO_MANY_OBJECTS
	VKErrorFormatNotSupported   common.VkResult = C.VK_ERROR_FORMAT_NOT_SUPPORTED
	VKErrorFragmentedPool       common.VkResult = C.VK_ERROR_FRAGMENTED_POOL
	VKErrorUnknown              common.VkResult = C.VK_ERROR_UNKNOWN
)

func init() {
	VKSuccess.Register("Success")
	VKNotReady.Register("Not Ready")
	VKTimeout.Register("Timeout")
	VKEventSet.Register("Event Set")
	VKEventReset.Register("Event Reset")
	VKIncomplete.Register("Incomplete")
	VKErrorOutOfHostMemory.Register("out of host memory")
	VKErrorOutOfDeviceMemory.Register("out of device memory")
	VKErrorInitializationFailed.Register("initialization failed")
	VKErrorDeviceLost.Register("device lost")
	VKErrorMemoryMapFailed.Register("memory map failed")
	VKErrorLayerNotPresent.Register("layer not present")
	VKErrorExtensionNotPresent.Register("extension not present")
	VKErrorFeatureNotPresent.Register("feature not present")
	VKErrorIncompatibleDriver.Register("incompatible driver")
	VKErrorTooManyObjects.Register("too many objects")
	VKErrorFormatNotSupported.Register("format not supported")
	VKErrorFragmentedPool.Register("fragmented pool")
	VKErrorUnknown.Register("unknown")
}
