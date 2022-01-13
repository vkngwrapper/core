package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type ShaderStages int32

const (
	StageVertex                 ShaderStages = C.VK_SHADER_STAGE_VERTEX_BIT
	StageTessellationControl    ShaderStages = C.VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT
	StageTessellationEvaluation ShaderStages = C.VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT
	StageGeometry               ShaderStages = C.VK_SHADER_STAGE_GEOMETRY_BIT
	StageFragment               ShaderStages = C.VK_SHADER_STAGE_FRAGMENT_BIT
	StageCompute                ShaderStages = C.VK_SHADER_STAGE_COMPUTE_BIT
	StageAllGraphics            ShaderStages = C.VK_SHADER_STAGE_ALL_GRAPHICS
	StageAll                    ShaderStages = C.VK_SHADER_STAGE_ALL
	StageRayGenKHR              ShaderStages = C.VK_SHADER_STAGE_RAYGEN_BIT_KHR
	StageAnyHitKHR              ShaderStages = C.VK_SHADER_STAGE_ANY_HIT_BIT_KHR
	StageClosestHitKHR          ShaderStages = C.VK_SHADER_STAGE_CLOSEST_HIT_BIT_KHR
	StageMissKHR                ShaderStages = C.VK_SHADER_STAGE_MISS_BIT_KHR
	StageIntersectionKHR        ShaderStages = C.VK_SHADER_STAGE_INTERSECTION_BIT_KHR
	StageCallableKHR            ShaderStages = C.VK_SHADER_STAGE_CALLABLE_BIT_KHR
	StageTaskNV                 ShaderStages = C.VK_SHADER_STAGE_TASK_BIT_NV
	StageMeshNV                 ShaderStages = C.VK_SHADER_STAGE_MESH_BIT_NV
)

var shaderStageToString = map[ShaderStages]string{
	StageVertex:                 "Vertex",
	StageTessellationControl:    "Tessellation Control",
	StageTessellationEvaluation: "Tessellation Evaluation",
	StageGeometry:               "Geometry",
	StageFragment:               "Fragment",
	StageCompute:                "Compute",
	StageRayGenKHR:              "Ray Gen (Khronos Extension)",
	StageAnyHitKHR:              "Any Hit (Khronos Extension)",
	StageClosestHitKHR:          "Closest Hit (Khronos Extension)",
	StageMissKHR:                "Miss (Khronos Extension)",
	StageIntersectionKHR:        "Intersection (Khronos Extension)",
	StageCallableKHR:            "Callable (Khronos Extension)",
	StageTaskNV:                 "Task (Nvidia Extension)",
	StageMeshNV:                 "Mesh (Nvidia Extension)",
}

func (s ShaderStages) String() string {
	return FlagsToString(s, shaderStageToString)
}
