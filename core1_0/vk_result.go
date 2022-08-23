package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/v2/common"

const (
	// VKSuccess indicates the command was successfully completed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKSuccess common.VkResult = C.VK_SUCCESS
	// VKNotReady indicates a Fence or query has not yet completed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKNotReady common.VkResult = C.VK_NOT_READY
	// VKTimeout indicates a wait operation has not completed in the specified time
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKTimeout common.VkResult = C.VK_TIMEOUT
	// VKEventSet indicates an Event is signaled
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKEventSet common.VkResult = C.VK_EVENT_SET
	// VKEventReset indicates an Event is unsignaled
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKEventReset common.VkResult = C.VK_EVENT_RESET
	// VKIncomplete indicates a return array was too small for the result
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKIncomplete common.VkResult = C.VK_INCOMPLETE
	// VKErrorOutOfHostMemory indicates a host memory allocation has failed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKErrorOutOfHostMemory common.VkResult = C.VK_ERROR_OUT_OF_HOST_MEMORY
	// VKErrorOutOfDeviceMemory indicates a DeviceMemory allocation has failed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKErrorOutOfDeviceMemory common.VkResult = C.VK_ERROR_OUT_OF_DEVICE_MEMORY
	// VKErrorInitializationFailed indicates initialization of an object could not be completed
	// for implementation-specific reasons
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKErrorInitializationFailed common.VkResult = C.VK_ERROR_INITIALIZATION_FAILED
	// VKErrorDeviceLost indicates the logical or physical device has been lost
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKErrorDeviceLost common.VkResult = C.VK_ERROR_DEVICE_LOST
	// VKErrorMemoryMapFailed indicates mapping of a memory object has failed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKErrorMemoryMapFailed common.VkResult = C.VK_ERROR_MEMORY_MAP_FAILED
	// VKErrorLayerNotPresent indicates a requested layer is not present or could not be loaded
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKErrorLayerNotPresent common.VkResult = C.VK_ERROR_LAYER_NOT_PRESENT
	// VKErrorExtensionNotPresent indicates a requested extension is not supported
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKErrorExtensionNotPresent common.VkResult = C.VK_ERROR_EXTENSION_NOT_PRESENT
	// VKErrorFeatureNotPresent indicates a requested feature is not supported
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKErrorFeatureNotPresent common.VkResult = C.VK_ERROR_FEATURE_NOT_PRESENT
	// VKErrorIncompatibleDriver indicates the requested version of Vulkan is not supported
	// by the driver or is otherwise incompatible for implementation-specific reasons
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKErrorIncompatibleDriver common.VkResult = C.VK_ERROR_INCOMPATIBLE_DRIVER
	// VKErrorTooManyObjects indicates too many objects of the type have already been created
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKErrorTooManyObjects common.VkResult = C.VK_ERROR_TOO_MANY_OBJECTS
	// VKErrorFormatNotSupported indicates a requested format is not supported on this Device
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKErrorFormatNotSupported common.VkResult = C.VK_ERROR_FORMAT_NOT_SUPPORTED
	// VKErrorFragmentedPool indicates a pool allocation has failed due to fragmentation of the
	// pool's memory
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKErrorFragmentedPool common.VkResult = C.VK_ERROR_FRAGMENTED_POOL
	// VKErrorUnknown indicates an unknown error has occurred, either the application has
	// provided invalid input, or an implementation failure has occurred
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKErrorUnknown common.VkResult = C.VK_ERROR_UNKNOWN
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
