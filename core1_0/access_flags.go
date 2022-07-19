package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

const (
	// AccessIndirectCommandRead specifies read access to indirect command data read as part
	// of an indirect build, trace, drawing or dispatching command. Such access occurs in the
	// PipelineStageDrawIndirect pipeline stage.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
	AccessIndirectCommandRead AccessFlags = C.VK_ACCESS_INDIRECT_COMMAND_READ_BIT
	// AccessIndexRead specifies read access to an index buffer as part of an indexed drawing
	// command, bound by vkCmdBindIndexBuffer. Such access occurs in the
	// PipelineStageVertexInput pipeline stage.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
	AccessIndexRead AccessFlags = C.VK_ACCESS_INDEX_READ_BIT
	// AccessVertexAttributeRead specifies read access to a vertex buffer as part of a drawing
	// command, bound by vkCmdBindVertexBuffers. Such access occurs in the
	// PipelineStageVertexInput pipeline stage.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
	AccessVertexAttributeRead AccessFlags = C.VK_ACCESS_VERTEX_ATTRIBUTE_READ_BIT
	// AccessUniformRead specifies read access to a uniform buffer in any shader pipeline stage.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
	AccessUniformRead AccessFlags = C.VK_ACCESS_UNIFORM_READ_BIT
	// AccessInputAttachmentRead specifies read access to an input attachment within a render
	// pass during subpass shading or fragment shading. Such access occurs in the
	// PipelineStage2SubpassShadingHuawei or PipelineStageFragmentShader
	// pipeline stage.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
	AccessInputAttachmentRead AccessFlags = C.VK_ACCESS_INPUT_ATTACHMENT_READ_BIT
	// AccessShaderRead specifies read access to a uniform buffer, uniform texel buffer,
	// sampled image, storage buffer, physical storage buffer, shader binding table, storage
	// texel buffer, or storage image in any shader pipeline stage.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
	AccessShaderRead AccessFlags = C.VK_ACCESS_SHADER_READ_BIT
	// AccessShaderWrite specifies write access to a storage buffer, physical storage buffer,
	// storage texel buffer, or storage image in any shader pipeline stage.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
	AccessShaderWrite AccessFlags = C.VK_ACCESS_SHADER_WRITE_BIT
	// AccessColorAttachmentRead specifies read access to a color attachment, such as via
	// blending, logic operations, or via certain subpass load operations. It does not include
	// advanced blend operations. Such access occurs in the
	// PipelineStageColorAttachmentOutput pipeline stage.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
	AccessColorAttachmentRead AccessFlags = C.VK_ACCESS_COLOR_ATTACHMENT_READ_BIT
	// AccessColorAttachmentWrite specifies write access to a color, resolve, or depth/stencil
	// resolve attachment during a render pass or via certain subpass load and store operations.
	// Such access occurs in the PipelineStageColorAttachmentOutput pipeline stage.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
	AccessColorAttachmentWrite AccessFlags = C.VK_ACCESS_COLOR_ATTACHMENT_WRITE_BIT
	// AccessDepthStencilAttachmentRead specifies read access to a depth/stencil attachment, via
	// depth or stencil operations or via certain subpass load operations. Such access occurs in
	// the PipelineStageEarlyFragmentTests or PipelineStageLateFramentTests
	// pipeline stages.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
	AccessDepthStencilAttachmentRead AccessFlags = C.VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_READ_BIT
	// AccessDepthStencilAttachmentWrite specifies write access to a depth/stencil attachment,
	// via depth or stencil operations or via certain subpass load and store operations. Such
	// access occurs in the PipelineStageEarlyFragmentTests or
	// PipelineStageLateFragmentTests pipeline stages.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
	AccessDepthStencilAttachmentWrite AccessFlags = C.VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT
	// AccessTransferRead specifies read access to an image or buffer in a copy operation. Such
	// access occurs in the PipelineStage2AllTransfer pipeline stage.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
	AccessTransferRead AccessFlags = C.VK_ACCESS_TRANSFER_READ_BIT
	// AccessTransferWrite specifies write access to an image or buffer in a clear or copy
	// operation. Such access occurs in the PipelineStage2AllTransfer pipeline stage.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
	AccessTransferWrite AccessFlags = C.VK_ACCESS_TRANSFER_WRITE_BIT
	// AccessHostRead specifies read access by a host operation. Accesses of this type are not
	// performed through a resource, but directly on memory. Such access occurs in the
	// PipelineStageHost pipeline stage.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
	AccessHostRead AccessFlags = C.VK_ACCESS_HOST_READ_BIT
	// AccessHostWrite specifies write access by a host operation. Accesses of this type are not
	// performed through a resource, but directly on memory. Such access occurs in the
	// PipelineStageHost pipeline stage.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
	AccessHostWrite AccessFlags = C.VK_ACCESS_HOST_WRITE_BIT
	// AccessMemoryRead specifies all read accesses. It is always valid in any access mask,
	// and is treated as equivalent to setting all READ access flags that are valid where it is
	// used.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
	AccessMemoryRead AccessFlags = C.VK_ACCESS_MEMORY_READ_BIT
	// AccessMemoryWrite specifies all write accesses. It is always valid in any access mask,
	// and is treated as equivalent to setting all WRITE access flags that are valid where it
	// is used.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html
	AccessMemoryWrite AccessFlags = C.VK_ACCESS_MEMORY_WRITE_BIT
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
