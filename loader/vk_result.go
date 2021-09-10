package loader

/*
#include <vulkan/vulkan.h>
*/
import "C"
import (
	"fmt"
	"github.com/cockroachdb/errors"
)

type VkResult int

const (
	VKSuccess                                  VkResult = C.VK_SUCCESS
	VKNotReady                                 VkResult = C.VK_NOT_READY
	VKTimeout                                  VkResult = C.VK_TIMEOUT
	VKEventSet                                 VkResult = C.VK_EVENT_SET
	VKEventReset                               VkResult = C.VK_EVENT_RESET
	VKIncomplete                               VkResult = C.VK_INCOMPLETE
	VKErrorOutOfHostMemory                     VkResult = C.VK_ERROR_OUT_OF_HOST_MEMORY
	VKErrorOutOfDeviceMemory                   VkResult = C.VK_ERROR_OUT_OF_DEVICE_MEMORY
	VKErrorInitializationFailed                VkResult = C.VK_ERROR_INITIALIZATION_FAILED
	VKErrorDeviceLost                          VkResult = C.VK_ERROR_DEVICE_LOST
	VKErrorMemoryMapFailed                     VkResult = C.VK_ERROR_MEMORY_MAP_FAILED
	VKErrorLayerNotPresent                     VkResult = C.VK_ERROR_LAYER_NOT_PRESENT
	VKErrorExtensionNotPresent                 VkResult = C.VK_ERROR_EXTENSION_NOT_PRESENT
	VKErrorFeatureNotPresent                   VkResult = C.VK_ERROR_FEATURE_NOT_PRESENT
	VKErrorIncompatibleDriver                  VkResult = C.VK_ERROR_INCOMPATIBLE_DRIVER
	VKErrorTooManyObjects                      VkResult = C.VK_ERROR_TOO_MANY_OBJECTS
	VKErrorFormatNotSupported                  VkResult = C.VK_ERROR_FORMAT_NOT_SUPPORTED
	VKErrorFragmentedPool                      VkResult = C.VK_ERROR_FRAGMENTED_POOL
	VKErrorUnknown                             VkResult = C.VK_ERROR_UNKNOWN
	VKErrorOutOfPoolMemory                     VkResult = C.VK_ERROR_OUT_OF_POOL_MEMORY_KHR
	VKErrorInvalidExternalHandle               VkResult = C.VK_ERROR_INVALID_EXTERNAL_HANDLE_KHR
	VKErrorFragmentation                       VkResult = C.VK_ERROR_FRAGMENTATION_EXT
	VKErrorInvalidDeviceAddress                VkResult = C.VK_ERROR_INVALID_DEVICE_ADDRESS_EXT
	VKErrorSurfaceLost                         VkResult = C.VK_ERROR_SURFACE_LOST_KHR
	VKErrorNativeWindowInUse                   VkResult = C.VK_ERROR_NATIVE_WINDOW_IN_USE_KHR
	VKSuboptimal                               VkResult = C.VK_SUBOPTIMAL_KHR
	VKErrorOutOfDate                           VkResult = C.VK_ERROR_OUT_OF_DATE_KHR
	VKErrorIncompatibleDisplay                 VkResult = C.VK_ERROR_INCOMPATIBLE_DISPLAY_KHR
	VKErrorValidationFailed                    VkResult = C.VK_ERROR_VALIDATION_FAILED_EXT
	VKErrorInvalidShader                       VkResult = C.VK_ERROR_INVALID_SHADER_NV
	VKErrorInvalidDRMFormatModifierPlaneLayout VkResult = C.VK_ERROR_INVALID_DRM_FORMAT_MODIFIER_PLANE_LAYOUT_EXT
	VKErrorNotPermitted                        VkResult = C.VK_ERROR_NOT_PERMITTED_EXT
	VKErrorFullScreenExclusiveModeLost         VkResult = C.VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT
	VKThreadIdle                               VkResult = C.VK_THREAD_IDLE_KHR
	VKThreadDone                               VkResult = C.VK_THREAD_DONE_KHR
	VKOperationDeferred                        VkResult = C.VK_OPERATION_DEFERRED_KHR
	VKOperationNotDeferred                     VkResult = C.VK_OPERATION_NOT_DEFERRED_KHR
	VKPipelineCompileRequired                  VkResult = C.VK_PIPELINE_COMPILE_REQUIRED_EXT
)

var vkResultToString = map[VkResult]string{
	VKSuccess:                                  "Success",
	VKNotReady:                                 "Not Ready",
	VKTimeout:                                  "Timeout",
	VKEventSet:                                 "Event Set",
	VKEventReset:                               "Event Reset",
	VKIncomplete:                               "Incomplete",
	VKErrorOutOfHostMemory:                     "out of host memory",
	VKErrorOutOfDeviceMemory:                   "out of device memory",
	VKErrorInitializationFailed:                "initialization failed",
	VKErrorDeviceLost:                          "device lost",
	VKErrorMemoryMapFailed:                     "memory map failed",
	VKErrorLayerNotPresent:                     "layer not present",
	VKErrorExtensionNotPresent:                 "extension not present",
	VKErrorFeatureNotPresent:                   "feature not present",
	VKErrorIncompatibleDriver:                  "incompatible driver",
	VKErrorTooManyObjects:                      "too many objects",
	VKErrorFormatNotSupported:                  "format not supported",
	VKErrorFragmentedPool:                      "fragmented pool",
	VKErrorUnknown:                             "unknown",
	VKErrorOutOfPoolMemory:                     "out of pool memory",
	VKErrorInvalidExternalHandle:               "invalid external handle",
	VKErrorFragmentation:                       "fragmentation",
	VKErrorInvalidDeviceAddress:                "invalid device address",
	VKErrorSurfaceLost:                         "surface lost",
	VKErrorNativeWindowInUse:                   "native window in use",
	VKSuboptimal:                               "Suboptimal",
	VKErrorOutOfDate:                           "out of date",
	VKErrorIncompatibleDisplay:                 "incompatible display",
	VKErrorValidationFailed:                    "validation failed",
	VKErrorInvalidShader:                       "invalid shader",
	VKErrorInvalidDRMFormatModifierPlaneLayout: "invalid drm format modifier plane layout",
	VKErrorNotPermitted:                        "not permitted",
	VKErrorFullScreenExclusiveModeLost:         "full-screen exclusive mode lost",
	VKThreadIdle:                               "Thread Idle",
	VKThreadDone:                               "Thread Done",
	VKOperationDeferred:                        "Operation Deferred",
	VKOperationNotDeferred:                     "Operation Not Deferred",
	VKPipelineCompileRequired:                  "Pipeline Compile Required",
}

func (r VkResult) String() string {
	return vkResultToString[r]
}

func (r VkResult) ToError() error {
	if r >= 0 {
		// All VKError* are <0
		return nil
	}

	return errors.WithStack(&VkResultError{r})
}

type VkResultError struct {
	code VkResult
}

func (err *VkResultError) Error() string {
	return fmt.Sprintf("vulkan error: %s", err.code.String())
}

func ResultFromError(err error) VkResult {
	if err == nil {
		return VKSuccess
	}

	var target VkResultError
	if errors.As(err, &target) {
		return target.code
	}

	return VKErrorUnknown
}
