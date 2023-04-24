package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v2/common"
)

// ShaderModuleCreateFlags is reserved for future use
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateFlags.html
type ShaderModuleCreateFlags int32

var shaderModuleCreateFlagsMapping = common.NewFlagStringMapping[ShaderModuleCreateFlags]()

func (f ShaderModuleCreateFlags) Register(str string) {
	shaderModuleCreateFlagsMapping.Register(f, str)
}

func (f ShaderModuleCreateFlags) String() string {
	return shaderModuleCreateFlagsMapping.FlagsToString(f)
}

////

// ShaderModuleCreateInfo specifies parameters of a newly-created ShaderModule
type ShaderModuleCreateInfo struct {
	// Code is the code that is used to create the ShaderModule. The type and format of the code
	// is determined from the content of the data
	Code []uint32
	// Flags is reserved for future use
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
