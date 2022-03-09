package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

const (
	AccessIndirectCommandRead         common.AccessFlags = C.VK_ACCESS_INDIRECT_COMMAND_READ_BIT
	AccessIndexRead                   common.AccessFlags = C.VK_ACCESS_INDEX_READ_BIT
	AccessVertexAttributeRead         common.AccessFlags = C.VK_ACCESS_VERTEX_ATTRIBUTE_READ_BIT
	AccessUniformRead                 common.AccessFlags = C.VK_ACCESS_UNIFORM_READ_BIT
	AccessInputAttachmentRead         common.AccessFlags = C.VK_ACCESS_INPUT_ATTACHMENT_READ_BIT
	AccessShaderRead                  common.AccessFlags = C.VK_ACCESS_SHADER_READ_BIT
	AccessShaderWrite                 common.AccessFlags = C.VK_ACCESS_SHADER_WRITE_BIT
	AccessColorAttachmentRead         common.AccessFlags = C.VK_ACCESS_COLOR_ATTACHMENT_READ_BIT
	AccessColorAttachmentWrite        common.AccessFlags = C.VK_ACCESS_COLOR_ATTACHMENT_WRITE_BIT
	AccessDepthStencilAttachmentRead  common.AccessFlags = C.VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_READ_BIT
	AccessDepthStencilAttachmentWrite common.AccessFlags = C.VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT
	AccessTransferRead                common.AccessFlags = C.VK_ACCESS_TRANSFER_READ_BIT
	AccessTransferWrite               common.AccessFlags = C.VK_ACCESS_TRANSFER_WRITE_BIT
	AccessHostRead                    common.AccessFlags = C.VK_ACCESS_HOST_READ_BIT
	AccessHostWrite                   common.AccessFlags = C.VK_ACCESS_HOST_WRITE_BIT
	AccessMemoryRead                  common.AccessFlags = C.VK_ACCESS_MEMORY_READ_BIT
	AccessMemoryWrite                 common.AccessFlags = C.VK_ACCESS_MEMORY_WRITE_BIT
)

func init() {
	AccessIndirectCommandRead.Register("Indirect Command Read")
	AccessIndexRead.Register("Index Read")
	AccessVertexAttributeRead.Register("Vertex Attribute Read")
	AccessUniformRead.Register("Uniform Read")
	AccessInputAttachmentRead.Register("Input Attachment Read")
	AccessShaderRead.Register("Shader Read")
	AccessShaderWrite.Register("Shader Write")
	AccessColorAttachmentRead.Register("Color Attachment Read")
	AccessColorAttachmentWrite.Register("Color Attachment Write")
	AccessDepthStencilAttachmentRead.Register("Depth/Stencil Attachment Read")
	AccessDepthStencilAttachmentWrite.Register("Depth/Stencil Attachment Write")
	AccessTransferRead.Register("Transfer Read")
	AccessTransferWrite.Register("Transfer Write")
	AccessHostRead.Register("Host Read")
	AccessHostWrite.Register("Host Write")
	AccessMemoryRead.Register("Memory Read")
	AccessMemoryWrite.Register("Memory Write")
}
