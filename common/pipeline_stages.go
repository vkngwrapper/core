package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type PipelineStages int32

const (
	PipelineStageNoneKHR                          PipelineStages = C.VK_PIPELINE_STAGE_NONE_KHR
	PipelineStageTopOfPipe                        PipelineStages = C.VK_PIPELINE_STAGE_TOP_OF_PIPE_BIT
	PipelineStageDrawIndirect                     PipelineStages = C.VK_PIPELINE_STAGE_DRAW_INDIRECT_BIT
	PipelineStageVertexInput                      PipelineStages = C.VK_PIPELINE_STAGE_VERTEX_INPUT_BIT
	PipelineStageVertexShader                     PipelineStages = C.VK_PIPELINE_STAGE_VERTEX_SHADER_BIT
	PipelineStageTessellationControlShader        PipelineStages = C.VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT
	PipelineStageTessellationEvaluationShader     PipelineStages = C.VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT
	PipelineStageGeometryShader                   PipelineStages = C.VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT
	PipelineStageFragmentShader                   PipelineStages = C.VK_PIPELINE_STAGE_FRAGMENT_SHADER_BIT
	PipelineStageEarlyFragmentTests               PipelineStages = C.VK_PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT
	PipelineStageLateFragmentTests                PipelineStages = C.VK_PIPELINE_STAGE_LATE_FRAGMENT_TESTS_BIT
	PipelineStageColorAttachmentOutput            PipelineStages = C.VK_PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT
	PipelineStageComputeShader                    PipelineStages = C.VK_PIPELINE_STAGE_COMPUTE_SHADER_BIT
	PipelineStageTransfer                         PipelineStages = C.VK_PIPELINE_STAGE_TRANSFER_BIT
	PipelineStageBottomOfPipe                     PipelineStages = C.VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
	PipelineStageHost                             PipelineStages = C.VK_PIPELINE_STAGE_HOST_BIT
	PipelineStageAllGraphics                      PipelineStages = C.VK_PIPELINE_STAGE_ALL_GRAPHICS_BIT
	PipelineStageAllCommands                      PipelineStages = C.VK_PIPELINE_STAGE_ALL_COMMANDS_BIT
	PipelineStageTransformFeedbackEXT             PipelineStages = C.VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT
	PipelineStageConditionalRenderingEXT          PipelineStages = C.VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT
	PipelineStageAccelerationStructureBuildKHR    PipelineStages = C.VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR
	PipelineStageRayTracingShaderKHR              PipelineStages = C.VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR
	PipelineStageTaskShaderNV                     PipelineStages = C.VK_PIPELINE_STAGE_TASK_SHADER_BIT_NV
	PipelineStageMeshShaderNV                     PipelineStages = C.VK_PIPELINE_STAGE_MESH_SHADER_BIT_NV
	PipelineStageFragmentDensityProcessEXT        PipelineStages = C.VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT
	PipelineStageFragmentShadingRateAttachmentKHR PipelineStages = C.VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR
	PipelineStageCommandPreprocessNV              PipelineStages = C.VK_PIPELINE_STAGE_COMMAND_PREPROCESS_BIT_NV
)

var pipelineStageToString = map[PipelineStages]string{
	PipelineStageTopOfPipe:                        "Top Of Pipe",
	PipelineStageDrawIndirect:                     "Draw Indirect",
	PipelineStageVertexInput:                      "Vertex Input",
	PipelineStageVertexShader:                     "Vertex Shader",
	PipelineStageTessellationControlShader:        "Tessellation Control Shader",
	PipelineStageTessellationEvaluationShader:     "Tessellation Evaluation Shader",
	PipelineStageGeometryShader:                   "Geometry Shader",
	PipelineStageFragmentShader:                   "Fragment Shader",
	PipelineStageEarlyFragmentTests:               "Early Fragment Tests",
	PipelineStageLateFragmentTests:                "Late Fragment Tests",
	PipelineStageColorAttachmentOutput:            "Color Attachment Output",
	PipelineStageComputeShader:                    "Compute Shader",
	PipelineStageTransfer:                         "Transfer",
	PipelineStageBottomOfPipe:                     "Bottom Of Pipe",
	PipelineStageHost:                             "Host",
	PipelineStageAllGraphics:                      "All Graphics",
	PipelineStageAllCommands:                      "All Commands",
	PipelineStageTransformFeedbackEXT:             "Transform Feedback (Extension)",
	PipelineStageConditionalRenderingEXT:          "Conditional Rendering (Extension)",
	PipelineStageAccelerationStructureBuildKHR:    "Acceleration Structure Build (Khronos Extension)",
	PipelineStageRayTracingShaderKHR:              "Ray Tracing Shader (Khronos Extension)",
	PipelineStageTaskShaderNV:                     "Task Shader (Nvidia Extension)",
	PipelineStageMeshShaderNV:                     "Mesh Shader (Nvidia Extension)",
	PipelineStageFragmentDensityProcessEXT:        "Fragment Density Process (Extension)",
	PipelineStageFragmentShadingRateAttachmentKHR: "Fragment Shading Rate Attachment (Khronos Extension)",
	PipelineStageCommandPreprocessNV:              "Command Preprocess (Nvidia Extension)",
}

func (s PipelineStages) String() string {
	return FlagsToString(s, pipelineStageToString)
}
