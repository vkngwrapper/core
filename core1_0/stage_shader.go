package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

const (
	StageVertex                 common.ShaderStages = C.VK_SHADER_STAGE_VERTEX_BIT
	StageTessellationControl    common.ShaderStages = C.VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT
	StageTessellationEvaluation common.ShaderStages = C.VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT
	StageGeometry               common.ShaderStages = C.VK_SHADER_STAGE_GEOMETRY_BIT
	StageFragment               common.ShaderStages = C.VK_SHADER_STAGE_FRAGMENT_BIT
	StageCompute                common.ShaderStages = C.VK_SHADER_STAGE_COMPUTE_BIT
	StageAllGraphics            common.ShaderStages = C.VK_SHADER_STAGE_ALL_GRAPHICS
	StageAll                    common.ShaderStages = C.VK_SHADER_STAGE_ALL
)

func init() {
	StageVertex.Register("Vertex")
	StageTessellationControl.Register("Tessellation Control")
	StageTessellationEvaluation.Register("Tessellation Evaluation")
	StageGeometry.Register("Geometry")
	StageFragment.Register("Fragment")
	StageCompute.Register("Compute")
}

type ShaderStageOptions struct {
	Flags              common.ShaderStageCreateFlags
	Name               string
	Stage              common.ShaderStages
	Shader             ShaderModule
	SpecializationInfo map[uint32]interface{}

	common.HaveNext
}

func (s ShaderStageOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineShaderStageCreateInfo)
	}

	createInfo := (*C.VkPipelineShaderStageCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO
	createInfo.flags = C.VkPipelineShaderStageCreateFlags(s.Flags)
	createInfo.pNext = next
	createInfo.stage = C.VkShaderStageFlagBits(s.Stage)
	createInfo.module = C.VkShaderModule(unsafe.Pointer(s.Shader.Handle()))
	createInfo.pName = (*C.char)(allocator.CString(s.Name))
	createInfo.pSpecializationInfo = nil

	if s.SpecializationInfo != nil && len(s.SpecializationInfo) > 0 {
		specInfo := (*C.VkSpecializationInfo)(allocator.Malloc(int(unsafe.Sizeof(C.VkSpecializationInfo{}))))
		mapLen := len(s.SpecializationInfo)
		specInfo.mapEntryCount = C.uint32_t(mapLen)

		mapEntryPtr := (*C.VkSpecializationMapEntry)(allocator.Malloc(mapLen * int(unsafe.Sizeof(C.VkSpecializationMapEntry{}))))
		mapEntrySlice := ([]C.VkSpecializationMapEntry)(unsafe.Slice(mapEntryPtr, mapLen))
		dataBytes := new(bytes.Buffer)
		mapIndex := 0

		for constantID, val := range s.SpecializationInfo {
			mapEntrySlice[mapIndex].constantID = C.uint32_t(constantID)
			mapEntrySlice[mapIndex].offset = C.uint32_t(dataBytes.Len())

			boolVal, isBool := val.(bool)
			if isBool {
				val = uint32(C.VK_FALSE)
				if boolVal {
					val = uint32(C.VK_TRUE)
				}
			}

			err := binary.Write(dataBytes, common.ByteOrder, val)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to populate shader stage with specialization values: %d -> %v", constantID, val)
			}
			mapEntrySlice[mapIndex].size = C.size_t(binary.Size(val))

			mapIndex++
		}
		specInfo.pMapEntries = mapEntryPtr
		specInfo.dataSize = C.size_t(dataBytes.Len())
		specInfo.pData = allocator.CBytes(dataBytes.Bytes())
		createInfo.pSpecializationInfo = specInfo
	}

	return preallocatedPointer, nil
}
