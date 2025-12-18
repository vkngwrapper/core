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
	"github.com/vkngwrapper/core/v3/common"
)

const (
	// AttachmentDescriptionMayAlias specifies that the attachment aliases the same DeviceMemory
	// as other attachments
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionFlagBits.html
	AttachmentDescriptionMayAlias AttachmentDescriptionFlags = C.VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT

	// AttachmentLoadOpLoad specifies that the previous contents of the Image within the render
	// area will be preserved
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentLoadOp.html
	AttachmentLoadOpLoad AttachmentLoadOp = C.VK_ATTACHMENT_LOAD_OP_LOAD
	// AttachmentLoadOpClear specifies that the contents within the rendera area will be cleared
	// to a uniform value, which is specified when a RenderPass instance is begun
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentLoadOp.html
	AttachmentLoadOpClear AttachmentLoadOp = C.VK_ATTACHMENT_LOAD_OP_CLEAR
	// AttachmentLoadOpDontCare specifies that the previous contents within the area need not
	// be preserved
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentLoadOp.html
	AttachmentLoadOpDontCare AttachmentLoadOp = C.VK_ATTACHMENT_LOAD_OP_DONT_CARE

	// AttachmentStoreOpStore specifies the contents generated during the RenderPass and within
	// the render area are written to memory
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html
	AttachmentStoreOpStore AttachmentStoreOp = C.VK_ATTACHMENT_STORE_OP_STORE
	// AttachmentStoreOpDontCare specifies the contents within the render area are not
	// needed after rendering, and may be discarded
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html
	AttachmentStoreOpDontCare AttachmentStoreOp = C.VK_ATTACHMENT_STORE_OP_DONT_CARE

	// DependencyByRegion specifies that dependencies will be Framebuffer local
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlagBits.html
	DependencyByRegion DependencyFlags = C.VK_DEPENDENCY_BY_REGION_BIT

	// PipelineBindPointGraphics specifies binding as a graphics Pipeline
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html
	PipelineBindPointGraphics PipelineBindPoint = C.VK_PIPELINE_BIND_POINT_GRAPHICS
	// PipelineBindPointCompute specifies binding as a compute Pipeline
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html
	PipelineBindPointCompute PipelineBindPoint = C.VK_PIPELINE_BIND_POINT_COMPUTE

	// SubpassExternal is a subpass index sentinel expanding synchronization scope outside a
	// subpass
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_SUBPASS_EXTERNAL.html
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

// AttachmentDescription specifies an attachment description
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription.html
type AttachmentDescription struct {
	// Flags specifies additional properties of the attachment
	Flags AttachmentDescriptionFlags
	// Format specifies the format of the ImageView that will be used for the attachment
	Format Format
	// Samples specifies the number of samples of the Image
	Samples SampleCountFlags

	// LoadOp specifies how the contents of color and depth components of the attachment are
	// treated at the beginning of the subpass where it is first used
	LoadOp AttachmentLoadOp
	// StoreOp specifies how the contents of color and depth components of the attachment are
	// treated at the end of the subpass where it is last used
	StoreOp AttachmentStoreOp
	// StencilLoadOp specifies how the contents of stencil components of the attachment are treated
	// at the beginning of the subpass where it is first used
	StencilLoadOp AttachmentLoadOp
	// StencilStoreOp specifies how the contents of stencil components of the attachment are treated
	// at the end of the subpass where it is last used
	StencilStoreOp AttachmentStoreOp

	// InitialLayout is the layout the attachment Image subresource will be in when a RenderPass
	// instance begins
	InitialLayout ImageLayout
	// FinalLayout is the layout the attachment Image subresource will be transitioned to when
	// a RenderPass instance ends
	FinalLayout ImageLayout
}

// SubpassDependency specifies a subpass dependency
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDependency.html
type SubpassDependency struct {
	// DependencyFlags is a bitmask of DependencyFlags
	DependencyFlags DependencyFlags

	// SrcSubpass is the subpass index of the first subpass in the dependency, or SubpassExternal
	SrcSubpass int
	// DstSubpass is the subpass index of the second subpass in the dependency, or SubpassExternal
	DstSubpass int

	// SrcStageMask specifies the source stage mask
	SrcStageMask PipelineStageFlags
	// DstStageMask specifies the destination stage mask
	DstStageMask PipelineStageFlags

	// SrcAccessMask specifies a source access mask
	SrcAccessMask AccessFlags
	// DstAccessMask specifies a destination access mask
	DstAccessMask AccessFlags
}

// SubpassDescription specifies a subpass description
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription.html
type SubpassDescription struct {
	// Flags specifies usage of the subpass
	Flags SubpassDescriptionFlags
	// PipelineBindPoint specifies the Pipeline type supported by this subpass
	PipelineBindPoint PipelineBindPoint

	// InputAttachments is a slice of AttachmentReference structures defining the input attachments
	// for this subpass and their layouts
	InputAttachments []AttachmentReference
	// ColorAttachments is a slice of AttachmentReference structures defining the color attachments
	// for this subpass and their layouts
	ColorAttachments []AttachmentReference
	// ResolveAttachments is a slice of AttachmentReference structures defining the resolve
	// attachments for this subpass and their layouts
	ResolveAttachments []AttachmentReference
	// DepthStencilAttachment specifies the depth/stencil attachment for this subpass and its
	// layout
	DepthStencilAttachment *AttachmentReference
	// PreserveAttachments is a slice of indices identifying attachments that are not used by
	// this subpass, but whose contents must be preserved throughout the subpass
	PreserveAttachments []int
}

// RenderPassCreateInfo specifies parameters of a newly-created RenderPass
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html
type RenderPassCreateInfo struct {
	// Flags is a bitmask of RenderPassCreateFlags
	Flags RenderPassCreateFlags
	// Attachments is a slice of AttachmentDescription structures describing the attachments
	// used by the RenderPass
	Attachments []AttachmentDescription
	// Subpasses is a slice of SubpassDescription structures describing each subpass
	Subpasses []SubpassDescription
	// SubpassDependencies is a slice of SubpassDependency structures describing dependencies
	// between pairs of subpasses
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
				return nil, errors.Errorf("in subpass %d, %d color attachments are defined, but %d resolve attachments are defined", i, colorAttachmentCount, resolveAttachmentCount)
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
