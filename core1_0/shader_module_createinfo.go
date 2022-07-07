package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

type ShaderModuleCreateFlags int32

var shaderModuleCreateFlagsMapping = common.NewFlagStringMapping[ShaderModuleCreateFlags]()

func (f ShaderModuleCreateFlags) Register(str string) {
	shaderModuleCreateFlagsMapping.Register(f, str)
}

func (f ShaderModuleCreateFlags) String() string {
	return shaderModuleCreateFlagsMapping.FlagsToString(f)
}

////

type ShaderModuleCreateInfo struct {
	Code  []uint32
	Flags ShaderModuleCreateFlags

	common.NextOptions
}

func (o ShaderModuleCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	byteCodeLen := len(o.Code)
	if byteCodeLen == 0 {
		return nil, errors.New("attempted to create a shader module with no shader bytecode")
	}

	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkShaderModuleCreateInfo{})))
	}

	createInfo := (*C.VkShaderModuleCreateInfo)(preallocatedPointer)
	bytecodePtr := (*C.uint32_t)(allocator.Malloc(byteCodeLen * int(unsafe.Sizeof(C.uint32_t(0)))))

	byteCodeArray := ([]C.uint32_t)(unsafe.Slice(bytecodePtr, byteCodeLen))

	createInfo.sType = C.VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO
	createInfo.flags = C.VkShaderModuleCreateFlags(o.Flags)
	createInfo.pNext = next
	createInfo.codeSize = C.size_t(byteCodeLen * 4)
	createInfo.pCode = bytecodePtr

	for i := 0; i < byteCodeLen; i++ {
		byteCodeArray[i] = C.uint32_t(o.Code[i])
	}

	return preallocatedPointer, nil
}
