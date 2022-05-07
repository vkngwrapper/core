package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type ShaderModuleCreateOptions struct {
	SpirVByteCode []uint32

	common.HaveNext
}

func (o ShaderModuleCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	byteCodeLen := len(o.SpirVByteCode)
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
	createInfo.flags = 0
	createInfo.pNext = next
	createInfo.codeSize = C.size_t(byteCodeLen * 4)
	createInfo.pCode = bytecodePtr

	for i := 0; i < byteCodeLen; i++ {
		byteCodeArray[i] = C.uint32_t(o.SpirVByteCode[i])
	}

	return preallocatedPointer, nil
}

func (o ShaderModuleCreateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkShaderModuleCreateInfo)(cDataPointer)
	return createInfo.pNext, nil
}
