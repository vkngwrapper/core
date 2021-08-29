package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "strings"

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
	StageRayGen                 ShaderStages = C.VK_SHADER_STAGE_RAYGEN_BIT_KHR
	StageAnyHit                 ShaderStages = C.VK_SHADER_STAGE_ANY_HIT_BIT_KHR
	StageClosestHit             ShaderStages = C.VK_SHADER_STAGE_CLOSEST_HIT_BIT_KHR
	StageMiss                   ShaderStages = C.VK_SHADER_STAGE_MISS_BIT_KHR
	StageIntersection           ShaderStages = C.VK_SHADER_STAGE_INTERSECTION_BIT_KHR
	StageCallable               ShaderStages = C.VK_SHADER_STAGE_CALLABLE_BIT_KHR
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
	StageRayGen:                 "Ray Gen",
	StageAnyHit:                 "Any Hit",
	StageClosestHit:             "Closest Hit",
	StageMiss:                   "Miss",
	StageIntersection:           "Intersection",
	StageCallable:               "Callable",
	StageTaskNV:                 "Task (Nvidia)",
	StageMeshNV:                 "Mesh (Nvidia)",
}

func (s ShaderStages) String() string {
	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := ShaderStages(1 << i)
		if (s & checkBit) != 0 {
			if hasOne {
				sb.WriteString("|")
			}
			sb.WriteString(shaderStageToString[s])
			hasOne = true
		}
	}

	return sb.String()
}
