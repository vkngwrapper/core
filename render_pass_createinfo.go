package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"strings"
	"unsafe"
)

type RenderPassFlags int32

const (
	RenderPassCreateTransformBitQCOM RenderPassFlags = C.VK_RENDER_PASS_CREATE_TRANSFORM_BIT_QCOM
)

var renderPassFlagsToString = map[RenderPassFlags]string{
	RenderPassCreateTransformBitQCOM: "Create Transform Bit (Qualcomm)",
}

func (f RenderPassFlags) String() string {
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder
	for i := 0; i < 32; i++ {
		checkBit := RenderPassFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := renderPassFlagsToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteString("|")
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}

type SubPassFlags int32

const (
	SubPassPerViewAttributesNVX    = C.VK_SUBPASS_DESCRIPTION_PER_VIEW_ATTRIBUTES_BIT_NVX
	SubPassPerViewPositionXOnlyNVX = C.VK_SUBPASS_DESCRIPTION_PER_VIEW_POSITION_X_ONLY_BIT_NVX
	SubPassFragmentRegionQCOM      = C.VK_SUBPASS_DESCRIPTION_FRAGMENT_REGION_BIT_QCOM
	SubPassShaderResolveQCOM       = C.VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM
)

var subPassFlagsToString = map[SubPassFlags]string{
	SubPassPerViewAttributesNVX:    "Per-View Attributes (NVidia Experimental)",
	SubPassPerViewPositionXOnlyNVX: "Per-View Position X Only (NVidia Experimental)",
	SubPassFragmentRegionQCOM:      "Fragment Region (Qualcomm)",
	SubPassShaderResolveQCOM:       "Shader Resolve (Qualcomm)",
}

func (f SubPassFlags) String() string {
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder
	for i := 0; i < 32; i++ {
		checkBit := SubPassFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := subPassFlagsToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteString("|")
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}

type AttachmentDescription struct {
	Flags   common.AttachmentDescriptionFlags
	Format  common.DataFormat
	Samples common.SampleCounts

	LoadOp         common.AttachmentLoadOp
	StoreOp        common.AttachmentStoreOp
	StencilLoadOp  common.AttachmentLoadOp
	StencilStoreOp common.AttachmentStoreOp

	InitialLayout common.ImageLayout
	FinalLayout   common.ImageLayout
}

const SubpassExternal = int(C.VK_SUBPASS_EXTERNAL)

type SubPassDependency struct {
	Flags common.DependencyFlags

	SrcSubPassIndex int
	DstSubPassIndex int

	SrcStageMask common.PipelineStages
	DstStageMask common.PipelineStages

	SrcAccess common.AccessFlags
	DstAccess common.AccessFlags
}

type SubPass struct {
	Flags     SubPassFlags
	BindPoint common.PipelineBindPoint

	InputAttachments           []common.AttachmentReference
	ColorAttachments           []common.AttachmentReference
	ResolveAttachments         []common.AttachmentReference
	DepthStencilAttachment     *common.AttachmentReference
	PreservedAttachmentIndices []int
}

type RenderPassOptions struct {
	Flags               RenderPassFlags
	Attachments         []AttachmentDescription
	SubPasses           []SubPass
	SubPassDependencies []SubPassDependency

	common.HaveNext
}

func (o *RenderPassOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkRenderPassCreateInfo)(allocator.Malloc(int(unsafe.Sizeof(C.VkRenderPassCreateInfo{}))))
	createInfo.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO
	createInfo.flags = C.VkRenderPassCreateFlags(o.Flags)
	createInfo.pNext = next

	attachmentCount := len(o.Attachments)
	createInfo.attachmentCount = C.uint32_t(attachmentCount)

	if attachmentCount == 0 {
		createInfo.pAttachments = nil
	} else {
		attachmentPtr := (*C.VkAttachmentDescription)(allocator.Malloc(attachmentCount * int(unsafe.Sizeof(C.VkAttachmentDescription{}))))
		createInfo.pAttachments = attachmentPtr
		attachmentSlice := ([]C.VkAttachmentDescription)(unsafe.Slice(attachmentPtr, attachmentCount))

		for i := 0; i < attachmentCount; i++ {
			attachmentSlice[i].flags = C.VkAttachmentDescriptionFlags(o.Attachments[i].Flags)
			attachmentSlice[i].format = C.VkFormat(o.Attachments[i].Format)
			attachmentSlice[i].samples = C.VkSampleCountFlagBits(o.Attachments[i].Samples)
			attachmentSlice[i].loadOp = C.VkAttachmentLoadOp(o.Attachments[i].LoadOp)
			attachmentSlice[i].storeOp = C.VkAttachmentStoreOp(o.Attachments[i].StoreOp)
			attachmentSlice[i].stencilLoadOp = C.VkAttachmentLoadOp(o.Attachments[i].StencilLoadOp)
			attachmentSlice[i].stencilStoreOp = C.VkAttachmentStoreOp(o.Attachments[i].StencilStoreOp)
			attachmentSlice[i].initialLayout = C.VkImageLayout(o.Attachments[i].InitialLayout)
			attachmentSlice[i].finalLayout = C.VkImageLayout(o.Attachments[i].FinalLayout)
		}
	}

	subPassCount := len(o.SubPasses)
	createInfo.subpassCount = C.uint32_t(subPassCount)

	if subPassCount == 0 {
		createInfo.pSubpasses = nil
	} else {
		subPassPtr := (*C.VkSubpassDescription)(allocator.Malloc(subPassCount * int(unsafe.Sizeof(C.VkSubpassDescription{}))))
		createInfo.pSubpasses = subPassPtr
		subPassSlice := ([]C.VkSubpassDescription)(unsafe.Slice(subPassPtr, subPassCount))

		for i := 0; i < subPassCount; i++ {
			resolveAttachmentCount := len(o.SubPasses[i].ResolveAttachments)
			colorAttachmentCount := len(o.SubPasses[i].ColorAttachments)

			if resolveAttachmentCount > 0 && resolveAttachmentCount != colorAttachmentCount {
				return nil, errors.Newf("in subpass %d, %d color attachments are defined, but %d resolve attachments are defined", i, colorAttachmentCount, resolveAttachmentCount)
			}

			subPassSlice[i].flags = C.VkSubpassDescriptionFlags(o.SubPasses[i].Flags)
			subPassSlice[i].pipelineBindPoint = C.VkPipelineBindPoint(o.SubPasses[i].BindPoint)
			subPassSlice[i].inputAttachmentCount = C.uint32_t(len(o.SubPasses[i].InputAttachments))
			subPassSlice[i].pInputAttachments = createAttachmentReferences(allocator, o.SubPasses[i].InputAttachments)
			subPassSlice[i].colorAttachmentCount = C.uint32_t(colorAttachmentCount)
			subPassSlice[i].pColorAttachments = createAttachmentReferences(allocator, o.SubPasses[i].ColorAttachments)
			subPassSlice[i].pResolveAttachments = createAttachmentReferences(allocator, o.SubPasses[i].ResolveAttachments)
			subPassSlice[i].pDepthStencilAttachment = nil

			if o.SubPasses[i].DepthStencilAttachment != nil {
				subPassSlice[i].pDepthStencilAttachment = createAttachmentReferences(allocator, []common.AttachmentReference{
					*o.SubPasses[i].DepthStencilAttachment,
				})
			}

			preserveAttachmentCount := len(o.SubPasses[i].PreservedAttachmentIndices)
			subPassSlice[i].preserveAttachmentCount = C.uint32_t(preserveAttachmentCount)
			if preserveAttachmentCount == 0 {
				subPassSlice[i].pPreserveAttachments = nil
			} else {
				preserveAttachmentPtr := (*C.uint32_t)(allocator.Malloc(preserveAttachmentCount * int(unsafe.Sizeof(C.uint32_t(0)))))
				subPassSlice[i].pPreserveAttachments = preserveAttachmentPtr
				preserveAttachmentSlice := ([]C.uint32_t)(unsafe.Slice(preserveAttachmentPtr, preserveAttachmentCount))

				for attInd := 0; attInd < preserveAttachmentCount; attInd++ {
					preserveAttachmentSlice[attInd] = C.uint32_t(o.SubPasses[i].PreservedAttachmentIndices[attInd])
				}
			}
		}
	}

	dependencyCount := len(o.SubPassDependencies)
	createInfo.dependencyCount = C.uint32_t(dependencyCount)

	if dependencyCount == 0 {
		createInfo.pDependencies = nil
	} else {
		dependencyPtr := (*C.VkSubpassDependency)(allocator.Malloc(dependencyCount * int(unsafe.Sizeof(C.VkSubpassDependency{}))))
		createInfo.pDependencies = dependencyPtr
		dependencySlice := ([]C.VkSubpassDependency)(unsafe.Slice(dependencyPtr, dependencyCount))

		for i := 0; i < dependencyCount; i++ {
			dependencySlice[i].srcSubpass = C.uint32_t(o.SubPassDependencies[i].SrcSubPassIndex)
			dependencySlice[i].dstSubpass = C.uint32_t(o.SubPassDependencies[i].DstSubPassIndex)
			dependencySlice[i].srcStageMask = C.VkPipelineStageFlags(o.SubPassDependencies[i].SrcStageMask)
			dependencySlice[i].dstStageMask = C.VkPipelineStageFlags(o.SubPassDependencies[i].DstStageMask)
			dependencySlice[i].srcAccessMask = C.VkAccessFlags(o.SubPassDependencies[i].SrcAccess)
			dependencySlice[i].dstAccessMask = C.VkAccessFlags(o.SubPassDependencies[i].DstAccess)
			dependencySlice[i].dependencyFlags = C.VkDependencyFlags(o.SubPassDependencies[i].Flags)
		}
	}

	return unsafe.Pointer(createInfo), nil
}

func createAttachmentReferences(allocator *cgoparam.Allocator, references []common.AttachmentReference) *C.VkAttachmentReference {
	count := len(references)
	if count == 0 {
		return nil
	}

	inputAttachmentsPtr := (*C.VkAttachmentReference)(allocator.Malloc(count * int(unsafe.Sizeof(C.VkAttachmentReference{}))))
	inputAttachmentsSlice := ([]C.VkAttachmentReference)(unsafe.Slice(inputAttachmentsPtr, count))

	for i := 0; i < count; i++ {
		if references[i].AttachmentIndex < 0 {
			inputAttachmentsSlice[i].attachment = C.VK_ATTACHMENT_UNUSED
		} else {
			inputAttachmentsSlice[i].attachment = C.uint32_t(references[i].AttachmentIndex)
		}

		inputAttachmentsSlice[i].layout = C.VkImageLayout(references[i].Layout)
	}

	return inputAttachmentsPtr
}
