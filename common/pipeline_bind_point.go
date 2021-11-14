package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type PipelineBindPoint int32

const (
	BindGraphics      = C.VK_PIPELINE_BIND_POINT_GRAPHICS
	BindCompute       = C.VK_PIPELINE_BIND_POINT_COMPUTE
	BindRayTracingKHR = C.VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR
)

var pipelineBindPointToString = map[PipelineBindPoint]string{
	BindGraphics:      "Graphics",
	BindCompute:       "Compute",
	BindRayTracingKHR: "Ray Tracing (Khronos Extension)",
}

func (p PipelineBindPoint) String() string {
	return pipelineBindPointToString[p]
}
