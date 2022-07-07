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

const (
	AttachmentDescriptionMayAlias AttachmentDescriptionFlags = C.VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT

	AttachmentLoadOpLoad     AttachmentLoadOp = C.VK_ATTACHMENT_LOAD_OP_LOAD
	AttachmentLoadOpClear    AttachmentLoadOp = C.VK_ATTACHMENT_LOAD_OP_CLEAR
	AttachmentLoadOpDontCare AttachmentLoadOp = C.VK_ATTACHMENT_LOAD_OP_DONT_CARE

	AttachmentStoreOpStore    AttachmentStoreOp = C.VK_ATTACHMENT_STORE_OP_STORE
	AttachmentStoreOpDontCare AttachmentStoreOp = C.VK_ATTACHMENT_STORE_OP_DONT_CARE

	DependencyByRegion DependencyFlags = C.VK_DEPENDENCY_BY_REGION_BIT

	PipelineBindPointGraphics PipelineBindPoint = C.VK_PIPELINE_BIND_POINT_GRAPHICS
	PipelineBindPointCompute  PipelineBindPoint = C.VK_PIPELINE_BIND_POINT_COMPUTE

	SubpassExternal = int(C.VK_SUBPASS_EXTERNAL)
)

func init() {
	AttachmentDescriptionMayAlias.Register("May Alias")

	AttachmentLoadOpLoad.Register("Load")
	AttachmentLoadOpClear.Register("Clear")
	AttachmentLoadOpDontCare.Register("Don't Care")

	AttachmentStoreOpStore.Register("Store")
	AttachmentStoreOpDontCare.Register("Don't Care")

	DependencyByRegion.Register("By Region")

	PipelineBindPointGraphics.Register("Graphics")
	PipelineBindPointCompute.Register("Compute")
}

type AttachmentDescription struct {
	Flags   AttachmentDescriptionFlags
	Format  Format
	Samples SampleCountFlags

	LoadOp         AttachmentLoadOp
	StoreOp        AttachmentStoreOp
	StencilLoadOp  AttachmentLoadOp
	StencilStoreOp AttachmentStoreOp

	InitialLayout ImageLayout
	FinalLayout   ImageLayout
}

type SubpassDependency struct {
	DependencyFlags DependencyFlags

	SrcSubpass int
	DstSubpass int

	SrcStageMask PipelineStageFlags
	DstStageMask PipelineStageFlags

	SrcAccessMask AccessFlags
	DstAccessMask AccessFlags
}

type SubpassDescription struct {
	Flags             SubpassDescriptionFlags
	PipelineBindPoint PipelineBindPoint

	InputAttachments       []AttachmentReference
	ColorAttachments       []AttachmentReference
	ResolveAttachments     []AttachmentReference
	DepthStencilAttachment *AttachmentReference
	PreserveAttachments    []int
}

type RenderPassCreateInfo struct {
	Flags               RenderPassCreateFlags
	Attachments         []AttachmentDescription
	Subpasses           []SubpassDescription
	SubpassDependencies []SubpassDependency

	common.NextOptions
}

func (o RenderPassCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

	subPassCount := len(o.Subpasses)
	createInfo.subpassCount = C.uint32_t(subPassCount)

	if subPassCount == 0 {
		createInfo.pSubpasses = nil
	} else {
		subPassPtr := (*C.VkSubpassDescription)(allocator.Malloc(subPassCount * int(unsafe.Sizeof(C.VkSubpassDescription{}))))
		createInfo.pSubpasses = subPassPtr
		subPassSlice := ([]C.VkSubpassDescription)(unsafe.Slice(subPassPtr, subPassCount))

		for i := 0; i < subPassCount; i++ {
			resolveAttachmentCount := len(o.Subpasses[i].ResolveAttachments)
			colorAttachmentCount := len(o.Subpasses[i].ColorAttachments)

			if resolveAttachmentCount > 0 && resolveAttachmentCount != colorAttachmentCount {
				return nil, errors.Newf("in subpass %d, %d color attachments are defined, but %d resolve attachments are defined", i, colorAttachmentCount, resolveAttachmentCount)
			}

			subPassSlice[i].flags = C.VkSubpassDescriptionFlags(o.Subpasses[i].Flags)
			subPassSlice[i].pipelineBindPoint = C.VkPipelineBindPoint(o.Subpasses[i].PipelineBindPoint)
			subPassSlice[i].inputAttachmentCount = C.uint32_t(len(o.Subpasses[i].InputAttachments))
			subPassSlice[i].pInputAttachments = createAttachmentReferences(allocator, o.Subpasses[i].InputAttachments)
			subPassSlice[i].colorAttachmentCount = C.uint32_t(colorAttachmentCount)
			subPassSlice[i].pColorAttachments = createAttachmentReferences(allocator, o.Subpasses[i].ColorAttachments)
			subPassSlice[i].pResolveAttachments = createAttachmentReferences(allocator, o.Subpasses[i].ResolveAttachments)
			subPassSlice[i].pDepthStencilAttachment = nil

			if o.Subpasses[i].DepthStencilAttachment != nil {
				subPassSlice[i].pDepthStencilAttachment = createAttachmentReferences(allocator, []AttachmentReference{
					*o.Subpasses[i].DepthStencilAttachment,
				})
			}

			preserveAttachmentCount := len(o.Subpasses[i].PreserveAttachments)
			subPassSlice[i].preserveAttachmentCount = C.uint32_t(preserveAttachmentCount)
			if preserveAttachmentCount == 0 {
				subPassSlice[i].pPreserveAttachments = nil
			} else {
				preserveAttachmentPtr := (*C.uint32_t)(allocator.Malloc(preserveAttachmentCount * int(unsafe.Sizeof(C.uint32_t(0)))))
				subPassSlice[i].pPreserveAttachments = preserveAttachmentPtr
				preserveAttachmentSlice := ([]C.uint32_t)(unsafe.Slice(preserveAttachmentPtr, preserveAttachmentCount))

				for attInd := 0; attInd < preserveAttachmentCount; attInd++ {
					preserveAttachmentSlice[attInd] = C.uint32_t(o.Subpasses[i].PreserveAttachments[attInd])
				}
			}
		}
	}

	dependencyCount := len(o.SubpassDependencies)
	createInfo.dependencyCount = C.uint32_t(dependencyCount)

	if dependencyCount == 0 {
		createInfo.pDependencies = nil
	} else {
		dependencyPtr := (*C.VkSubpassDependency)(allocator.Malloc(dependencyCount * int(unsafe.Sizeof(C.VkSubpassDependency{}))))
		createInfo.pDependencies = dependencyPtr
		dependencySlice := ([]C.VkSubpassDependency)(unsafe.Slice(dependencyPtr, dependencyCount))

		for i := 0; i < dependencyCount; i++ {
			dependencySlice[i].srcSubpass = C.uint32_t(o.SubpassDependencies[i].SrcSubpass)
			dependencySlice[i].dstSubpass = C.uint32_t(o.SubpassDependencies[i].DstSubpass)
			dependencySlice[i].srcStageMask = C.VkPipelineStageFlags(o.SubpassDependencies[i].SrcStageMask)
			dependencySlice[i].dstStageMask = C.VkPipelineStageFlags(o.SubpassDependencies[i].DstStageMask)
			dependencySlice[i].srcAccessMask = C.VkAccessFlags(o.SubpassDependencies[i].SrcAccessMask)
			dependencySlice[i].dstAccessMask = C.VkAccessFlags(o.SubpassDependencies[i].DstAccessMask)
			dependencySlice[i].dependencyFlags = C.VkDependencyFlags(o.SubpassDependencies[i].DependencyFlags)
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
		if references[i].Attachment < 0 {
			inputAttachmentsSlice[i].attachment = C.VK_ATTACHMENT_UNUSED
		} else {
			inputAttachmentsSlice[i].attachment = C.uint32_t(references[i].Attachment)
		}

		inputAttachmentsSlice[i].layout = C.VkImageLayout(references[i].Layout)
	}

	return inputAttachmentsPtr
}
