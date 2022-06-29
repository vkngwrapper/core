package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

const (
	AccessIndirectCommandRead         AccessFlags = C.VK_ACCESS_INDIRECT_COMMAND_READ_BIT
	AccessIndexRead                   AccessFlags = C.VK_ACCESS_INDEX_READ_BIT
	AccessVertexAttributeRead         AccessFlags = C.VK_ACCESS_VERTEX_ATTRIBUTE_READ_BIT
	AccessUniformRead                 AccessFlags = C.VK_ACCESS_UNIFORM_READ_BIT
	AccessInputAttachmentRead         AccessFlags = C.VK_ACCESS_INPUT_ATTACHMENT_READ_BIT
	AccessShaderRead                  AccessFlags = C.VK_ACCESS_SHADER_READ_BIT
	AccessShaderWrite                 AccessFlags = C.VK_ACCESS_SHADER_WRITE_BIT
	AccessColorAttachmentRead         AccessFlags = C.VK_ACCESS_COLOR_ATTACHMENT_READ_BIT
	AccessColorAttachmentWrite        AccessFlags = C.VK_ACCESS_COLOR_ATTACHMENT_WRITE_BIT
	AccessDepthStencilAttachmentRead  AccessFlags = C.VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_READ_BIT
	AccessDepthStencilAttachmentWrite AccessFlags = C.VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT
	AccessTransferRead                AccessFlags = C.VK_ACCESS_TRANSFER_READ_BIT
	AccessTransferWrite               AccessFlags = C.VK_ACCESS_TRANSFER_WRITE_BIT
	AccessHostRead                    AccessFlags = C.VK_ACCESS_HOST_READ_BIT
	AccessHostWrite                   AccessFlags = C.VK_ACCESS_HOST_WRITE_BIT
	AccessMemoryRead                  AccessFlags = C.VK_ACCESS_MEMORY_READ_BIT
	AccessMemoryWrite                 AccessFlags = C.VK_ACCESS_MEMORY_WRITE_BIT
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
