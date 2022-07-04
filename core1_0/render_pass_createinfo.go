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

const (
	AttachmentDescriptionMayAlias AttachmentDescriptionFlags = C.VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT

	LoadOpLoad     AttachmentLoadOp = C.VK_ATTACHMENT_LOAD_OP_LOAD
	LoadOpClear    AttachmentLoadOp = C.VK_ATTACHMENT_LOAD_OP_CLEAR
	LoadOpDontCare AttachmentLoadOp = C.VK_ATTACHMENT_LOAD_OP_DONT_CARE

	StoreOpStore    AttachmentStoreOp = C.VK_ATTACHMENT_STORE_OP_STORE
	StoreOpDontCare AttachmentStoreOp = C.VK_ATTACHMENT_STORE_OP_DONT_CARE

	DependencyByRegion DependencyFlags = C.VK_DEPENDENCY_BY_REGION_BIT

	BindGraphics PipelineBindPoint = C.VK_PIPELINE_BIND_POINT_GRAPHICS
	BindCompute  PipelineBindPoint = C.VK_PIPELINE_BIND_POINT_COMPUTE

	SubpassExternal = int(C.VK_SUBPASS_EXTERNAL)
)

func init() {
	AttachmentDescriptionMayAlias.Register("May Alias")

	LoadOpLoad.Register("Load")
	LoadOpClear.Register("Clear")
	LoadOpDontCare.Register("Don't Care")

	StoreOpStore.Register("Store")
	StoreOpDontCare.Register("Don't Care")

	DependencyByRegion.Register("By Region")

	BindGraphics.Register("Graphics")
	BindCompute.Register("Compute")
}

type AttachmentDescription struct {
	Flags   AttachmentDescriptionFlags
	Format  DataFormat
	Samples SampleCounts

	LoadOp         AttachmentLoadOp
	StoreOp        AttachmentStoreOp
	StencilLoadOp  AttachmentLoadOp
	StencilStoreOp AttachmentStoreOp

	InitialLayout ImageLayout
	FinalLayout   ImageLayout
}

type SubPassDependency struct {
	Flags DependencyFlags

	SrcSubPassIndex int
	DstSubPassIndex int

	SrcStageMask PipelineStages
	DstStageMask PipelineStages

	SrcAccessMask AccessFlags
	DstAccessMask AccessFlags
}

type SubPassDescription struct {
	Flags     SubPassDescriptionFlags
	BindPoint PipelineBindPoint

	InputAttachments           []AttachmentReference
	ColorAttachments           []AttachmentReference
	ResolveAttachments         []AttachmentReference
	DepthStencilAttachment     *AttachmentReference
	PreservedAttachmentIndices []int
}

type RenderPassCreateOptions struct {
	Flags               RenderPassCreateFlags
	Attachments         []AttachmentDescription
	SubPassDescriptions []SubPassDescription
	SubPassDependencies []SubPassDependency

	common.NextOptions
}

func (o RenderPassCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderPassCreateInfo{})))
	}
	createInfo := (*C.VkRenderPassCreateInfo)(preallocatedPointer)
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

	subPassCount := len(o.SubPassDescriptions)
	createInfo.subpassCount = C.uint32_t(subPassCount)

	if subPassCount == 0 {
		createInfo.pSubpasses = nil
	} else {
		subPassPtr := (*C.VkSubpassDescription)(allocator.Malloc(subPassCount * int(unsafe.Sizeof(C.VkSubpassDescription{}))))
		createInfo.pSubpasses = subPassPtr
		subPassSlice := ([]C.VkSubpassDescription)(unsafe.Slice(subPassPtr, subPassCount))

		for i := 0; i < subPassCount; i++ {
			resolveAttachmentCount := len(o.SubPassDescriptions[i].ResolveAttachments)
			colorAttachmentCount := len(o.SubPassDescriptions[i].ColorAttachments)

			if resolveAttachmentCount > 0 && resolveAttachmentCount != colorAttachmentCount {
				return nil, errors.Newf("in subpass %d, %d color attachments are defined, but %d resolve attachments are defined", i, colorAttachmentCount, resolveAttachmentCount)
			}

			subPassSlice[i].flags = C.VkSubpassDescriptionFlags(o.SubPassDescriptions[i].Flags)
			subPassSlice[i].pipelineBindPoint = C.VkPipelineBindPoint(o.SubPassDescriptions[i].BindPoint)
			subPassSlice[i].inputAttachmentCount = C.uint32_t(len(o.SubPassDescriptions[i].InputAttachments))
			subPassSlice[i].pInputAttachments = createAttachmentReferences(allocator, o.SubPassDescriptions[i].InputAttachments)
			subPassSlice[i].colorAttachmentCount = C.uint32_t(colorAttachmentCount)
			subPassSlice[i].pColorAttachments = createAttachmentReferences(allocator, o.SubPassDescriptions[i].ColorAttachments)
			subPassSlice[i].pResolveAttachments = createAttachmentReferences(allocator, o.SubPassDescriptions[i].ResolveAttachments)
			subPassSlice[i].pDepthStencilAttachment = nil

			if o.SubPassDescriptions[i].DepthStencilAttachment != nil {
				subPassSlice[i].pDepthStencilAttachment = createAttachmentReferences(allocator, []AttachmentReference{
					*o.SubPassDescriptions[i].DepthStencilAttachment,
				})
			}

			preserveAttachmentCount := len(o.SubPassDescriptions[i].PreservedAttachmentIndices)
			subPassSlice[i].preserveAttachmentCount = C.uint32_t(preserveAttachmentCount)
			if preserveAttachmentCount == 0 {
				subPassSlice[i].pPreserveAttachments = nil
			} else {
				preserveAttachmentPtr := (*C.uint32_t)(allocator.Malloc(preserveAttachmentCount * int(unsafe.Sizeof(C.uint32_t(0)))))
				subPassSlice[i].pPreserveAttachments = preserveAttachmentPtr
				preserveAttachmentSlice := ([]C.uint32_t)(unsafe.Slice(preserveAttachmentPtr, preserveAttachmentCount))

				for attInd := 0; attInd < preserveAttachmentCount; attInd++ {
					preserveAttachmentSlice[attInd] = C.uint32_t(o.SubPassDescriptions[i].PreservedAttachmentIndices[attInd])
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
			dependencySlice[i].srcAccessMask = C.VkAccessFlags(o.SubPassDependencies[i].SrcAccessMask)
			dependencySlice[i].dstAccessMask = C.VkAccessFlags(o.SubPassDependencies[i].DstAccessMask)
			dependencySlice[i].dependencyFlags = C.VkDependencyFlags(o.SubPassDependencies[i].Flags)
		}
	}

	return unsafe.Pointer(createInfo), nil
}

func createAttachmentReferences(allocator *cgoparam.Allocator, references []AttachmentReference) *C.VkAttachmentReference {
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
