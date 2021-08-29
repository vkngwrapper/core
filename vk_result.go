package core

/*
#include <vulkan/vulkan.h>
*/
import "C"
import "github.com/palantir/stacktrace"

type Result int

const (
	VKSuccess                                  Result = C.VK_SUCCESS
	VKNotReady                                 Result = C.VK_NOT_READY
	VKTimeout                                  Result = C.VK_TIMEOUT
	VKEventSet                                 Result = C.VK_EVENT_SET
	VKEventReset                               Result = C.VK_EVENT_RESET
	VKIncomplete                               Result = C.VK_INCOMPLETE
	VKErrorOutOfHostMemory                     Result = C.VK_ERROR_OUT_OF_HOST_MEMORY
	VKErrorOutOfDeviceMemory                   Result = C.VK_ERROR_OUT_OF_DEVICE_MEMORY
	VKErrorInitializationFailed                Result = C.VK_ERROR_INITIALIZATION_FAILED
	VKErrorDeviceLost                          Result = C.VK_ERROR_DEVICE_LOST
	VKErrorMemoryMapFailed                     Result = C.VK_ERROR_MEMORY_MAP_FAILED
	VKErrorLayerNotPresent                     Result = C.VK_ERROR_LAYER_NOT_PRESENT
	VKErrorExtensionNotPresent                 Result = C.VK_ERROR_EXTENSION_NOT_PRESENT
	VKErrorFeatureNotPresent                   Result = C.VK_ERROR_FEATURE_NOT_PRESENT
	VKErrorIncompatibleDriver                  Result = C.VK_ERROR_INCOMPATIBLE_DRIVER
	VKErrorTooManyObjects                      Result = C.VK_ERROR_TOO_MANY_OBJECTS
	VKErrorFormatNotSupported                  Result = C.VK_ERROR_FORMAT_NOT_SUPPORTED
	VKErrorFragmentedPool                      Result = C.VK_ERROR_FRAGMENTED_POOL
	VKErrorUnknown                             Result = C.VK_ERROR_UNKNOWN
	VKErrorOutOfPoolMemory                     Result = C.VK_ERROR_OUT_OF_POOL_MEMORY_KHR
	VKErrorInvalidExternalHandle               Result = C.VK_ERROR_INVALID_EXTERNAL_HANDLE_KHR
	VKErrorFragmentation                       Result = C.VK_ERROR_FRAGMENTATION_EXT
	VKErrorInvalidDeviceAddress                Result = C.VK_ERROR_INVALID_DEVICE_ADDRESS_EXT
	VKErrorSurfaceLost                         Result = C.VK_ERROR_SURFACE_LOST_KHR
	VKErrorNativeWindowInUse                   Result = C.VK_ERROR_NATIVE_WINDOW_IN_USE_KHR
	VKSuboptimal                               Result = C.VK_SUBOPTIMAL_KHR
	VKErrorOutOfDate                           Result = C.VK_ERROR_OUT_OF_DATE_KHR
	VKErrorIncompatibleDisplay                 Result = C.VK_ERROR_INCOMPATIBLE_DISPLAY_KHR
	VKErrorValidationFailed                    Result = C.VK_ERROR_VALIDATION_FAILED_EXT
	VKErrorInvalidShader                       Result = C.VK_ERROR_INVALID_SHADER_NV
	VKErrorInvalidDRMFormatModifierPlaneLayout Result = C.VK_ERROR_INVALID_DRM_FORMAT_MODIFIER_PLANE_LAYOUT_EXT
	VKErrorNotPermitted                        Result = C.VK_ERROR_NOT_PERMITTED_EXT
	VKErrorFullScreenExclusiveModeLost         Result = C.VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT
	VKThreadIdle                               Result = C.VK_THREAD_IDLE_KHR
	VKThreadDone                               Result = C.VK_THREAD_DONE_KHR
	VKOperationDeferred                        Result = C.VK_OPERATION_DEFERRED_KHR
	VKOperationNotDeferred                     Result = C.VK_OPERATION_NOT_DEFERRED_KHR
	VKPipelineCompileRequired                  Result = C.VK_PIPELINE_COMPILE_REQUIRED_EXT
)

var vkResultToString = map[Result]string{
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

func (r Result) String() string {
	return vkResultToString[r]
}

func (r Result) ToError() error {
	if r >= 0 {
		// All VKError* are <0
		return nil
	}

	return stacktrace.NewErrorWithCode(stacktrace.ErrorCode(r), "vulkan error: %s", r.String())
}
