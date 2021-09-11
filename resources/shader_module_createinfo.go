package resources

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type ShaderModuleOptions struct {
	SpirVByteCode []uint32

	Next core.Options
}

func (o *ShaderModuleOptions) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	byteCodeLen := len(o.SpirVByteCode)
	if byteCodeLen == 0 {
		return nil, errors.New("attempted to create a shader module with no shader bytecode")
	}

	createInfo := (*C.VkShaderModuleCreateInfo)(allocator.Malloc(int(unsafe.Sizeof([1]C.VkShaderModuleCreateInfo{}))))
	bytecodePtr := (*C.uint32_t)(allocator.Malloc(byteCodeLen * int(unsafe.Sizeof(C.uint32_t(0)))))

	byteCodeArray := ([]C.uint32_t)(unsafe.Slice(bytecodePtr, byteCodeLen))

	createInfo.sType = C.VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO
	createInfo.flags = 0
	createInfo.codeSize = C.size_t(byteCodeLen * 4)
	createInfo.pCode = bytecodePtr

	for i := 0; i < byteCodeLen; i++ {
		byteCodeArray[i] = C.uint32_t(o.SpirVByteCode[i])
	}

	var err error
	var next unsafe.Pointer
	if o.Next != nil {
		next, err = o.Next.AllocForC(allocator)
	}

	if err != nil {
		return nil, err
	}
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}
