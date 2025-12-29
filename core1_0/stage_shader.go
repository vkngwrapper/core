package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3/common"
)

const (
	// StageVertex specifies the vertex stage
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html
	StageVertex ShaderStageFlags = C.VK_SHADER_STAGE_VERTEX_BIT
	// StageTessellationControl specifies the tessellation control stage
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html
	StageTessellationControl ShaderStageFlags = C.VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT
	// StageTessellationEvaluation specifies the tessellation evaluation stage
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html
	StageTessellationEvaluation ShaderStageFlags = C.VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT
	// StageGeometry specifies the geometry stage
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html
	StageGeometry ShaderStageFlags = C.VK_SHADER_STAGE_GEOMETRY_BIT
	// StageFragment specifies the fragment stage
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html
	StageFragment ShaderStageFlags = C.VK_SHADER_STAGE_FRAGMENT_BIT
	// StageCompute specifies the compute stage
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html
	StageCompute ShaderStageFlags = C.VK_SHADER_STAGE_COMPUTE_BIT
	// StageAllGraphics is a combination of bits used as shorthand to specify all graphics
	// stages (excluding the compute stage)
	StageAllGraphics ShaderStageFlags = C.VK_SHADER_STAGE_ALL_GRAPHICS
	// StageAll is a combination of bits used as a shorthand to specify all shader stages
	// supported by the Device, including all additional stages which are introduced by
	// extensions
	StageAll ShaderStageFlags = C.VK_SHADER_STAGE_ALL
)

func init() {
	StageVertex.Register("Vertex")
	StageTessellationControl.Register("Tessellation Control")
	StageTessellationEvaluation.Register("Tessellation Evaluation")
	StageGeometry.Register("Geometry")
	StageFragment.Register("Fragment")
	StageCompute.Register("Compute")
}

// PipelineShaderStageCreateInfo specifies parameters of a newly-created Pipeline shader stage
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html
type PipelineShaderStageCreateInfo struct {
	// Flags specifies how the Pipeline shader stage will be generated
	Flags PipelineShaderStageCreateFlags
	// Name is a string specifying the entry point name of the shader for this stage
	Name string
	// Stage specifies a single Pipeline stage
	Stage ShaderStageFlags
	// Module contains the shader code for this stage
	Module ShaderModule
	// SpecializationInfo is a map specifying specialization contents
	SpecializationInfo map[uint32]any

	common.NextOptions
}

func (s PipelineShaderStageCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if !s.Module.Initialized() {
		return nil, errors.New("PipelineShaderStageCreateInfo.Module cannot be left unset")
	}

	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineShaderStageCreateInfo)
	}

	createInfo := (*C.VkPipelineShaderStageCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO
	createInfo.flags = C.VkPipelineShaderStageCreateFlags(s.Flags)
	createInfo.pNext = next
	createInfo.stage = C.VkShaderStageFlagBits(s.Stage)
	createInfo.module = C.VkShaderModule(unsafe.Pointer(s.Module.Handle()))
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
