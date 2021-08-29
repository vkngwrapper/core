package core

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type PipelineBindPoint int32

const (
	BindGraphics   = C.VK_PIPELINE_BIND_POINT_GRAPHICS
	BindCompute    = C.VK_PIPELINE_BIND_POINT_COMPUTE
	BindRayTracing = C.VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR
)

var pipelineBindPointToString = map[PipelineBindPoint]string{
	BindGraphics:   "Graphics",
	BindCompute:    "Compute",
	BindRayTracing: "Ray Tracing",
}

func (p PipelineBindPoint) String() string {
	return pipelineBindPointToString[p]
}
