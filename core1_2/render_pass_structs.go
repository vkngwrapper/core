package core1_2

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

// SubpassBeginInfo specifies subpass begin information
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassBeginInfoKHR.html
type SubpassBeginInfo struct {
	// Contents specifies how the commands in the next subpass will be provided
	Contents core1_0.SubpassContents

	common.NextOptions
}

func (o SubpassBeginInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassBeginInfo{})))
	}

	info := (*C.VkSubpassBeginInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO
	info.pNext = next
	info.contents = C.VkSubpassContents(o.Contents)

	return preallocatedPointer, nil
}

////

// SubpassEndInfo specifies subpass end information
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassEndInfo.html
type SubpassEndInfo struct {
	common.NextOptions
}

func (o SubpassEndInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassEndInfo{})))
	}

	info := (*C.VkSubpassEndInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_END_INFO
	info.pNext = next

	return preallocatedPointer, nil
}

////

// AttachmentDescription2 specifies an attachment description
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription2.html
type AttachmentDescription2 struct {
	// Flags specifies additional properties of the attachment
	Flags core1_0.AttachmentDescriptionFlags
	// Format specifies the format of the Image that will be used for the attachment
	Format core1_0.Format
	// Samples specifies the number of samples of the Image
	Samples core1_0.SampleCountFlags
	// LoadOp specifies how the contents of color and depth components of the attachment
	// are treated at the beginning of the subpass where it is first used
	LoadOp core1_0.AttachmentLoadOp
	// StoreOp specifies how the contents of color and depth components of the attachment
	// are treated at the end of the subpass where it is last used
	StoreOp core1_0.AttachmentStoreOp
	// StencilLoadOp specifies how the contents of stencil components of the attachment
	// are treated at the beginning of the subpass where it is first used
	StencilLoadOp core1_0.AttachmentLoadOp
	// StencilStoreOp specifies how the contents of the stencil components of the attachment
	// are treated at the end of the last subpass where it is used
	StencilStoreOp core1_0.AttachmentStoreOp
	// InitialLayout is the layout of the attachment Image subresource will be in when
	// a RenderPass instance begins
	InitialLayout core1_0.ImageLayout
	// FinalLayout is the layout the attachment Image subresource will be transitioned to
	// when a RenderPass instance ends
	FinalLayout core1_0.ImageLayout

	common.NextOptions
}

func (o AttachmentDescription2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAttachmentDescription2{})))
	}

	info := (*C.VkAttachmentDescription2)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2
	info.pNext = next
	info.flags = C.VkAttachmentDescriptionFlags(o.Flags)
	info.format = C.VkFormat(o.Format)
	info.samples = C.VkSampleCountFlagBits(o.Samples)
	info.loadOp = C.VkAttachmentLoadOp(o.LoadOp)
	info.storeOp = C.VkAttachmentStoreOp(o.StoreOp)
	info.stencilLoadOp = C.VkAttachmentLoadOp(o.StencilLoadOp)
	info.stencilStoreOp = C.VkAttachmentStoreOp(o.StencilStoreOp)
	info.initialLayout = C.VkImageLayout(o.InitialLayout)
	info.finalLayout = C.VkImageLayout(o.FinalLayout)

	return preallocatedPointer, nil
}

////

// AttachmentReference2 specifies an attachment reference
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html
type AttachmentReference2 struct {
	// Attachment identifies an attachment at the corresponding index in
	// RenderPassCreateInfo2.Attachments, or core1_0.AttachmentUnused
	Attachment int
	// Layout specifies the layout the attachment uses during the subpass
	Layout core1_0.ImageLayout
	// AspectMask is a mask of which aspect(s) can be accessed within the specified
	// subpass as an input attachment
	AspectMask core1_0.ImageAspectFlags

	common.NextOptions
}

func (o AttachmentReference2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAttachmentReference2{})))
	}

	info := (*C.VkAttachmentReference2)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2
	info.pNext = next
	info.attachment = C.uint32_t(o.Attachment)
	info.layout = C.VkImageLayout(o.Layout)
	info.aspectMask = C.VkImageAspectFlags(o.AspectMask)

	return preallocatedPointer, nil
}

////

// SubpassDescription2 specifies a subpass description
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html
type SubpassDescription2 struct {
	// Flags specifies usage of the subpass
	Flags core1_0.SubpassDescriptionFlags
	// PipelineBindPoint specifies the Pipeline type supported for this subpass
	PipelineBindPoint core1_0.PipelineBindPoint
	// ViewMask describes which views rendering is broadcast to in this subpass, when
	// multiview is enabled
	ViewMask uint32
	// InputAttachments is a slice of AttachmentReference2 structures defining the input
	// attachments for this subpass and their layouts
	InputAttachments []AttachmentReference2
	// ColorAttachments is a slice of AttachmentReference2 structures defining the color
	// attachments for this subpass and their layouts
	ColorAttachments []AttachmentReference2
	// ResolveAttachments is a slice of AttachmentReference2 structures defining the resolve
	// attachments for this subpass and their layouts
	ResolveAttachments []AttachmentReference2
	// DepthStencilAttachment specifies the depth/stencil attachment for this subpass and
	// its layout
	DepthStencilAttachment *AttachmentReference2
	// PreserveAttachments is a slice of RenderPass attachment indices identifying attachments
	// that are not used by this subpass, but whose contents must be preserved throughout the
	// subpass
	PreserveAttachments []int

	common.NextOptions
}

func (o SubpassDescription2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassDescription2{})))
	}

	info := (*C.VkSubpassDescription2)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2
	info.pNext = next
	info.flags = C.VkSubpassDescriptionFlags(o.Flags)
	info.pipelineBindPoint = C.VkPipelineBindPoint(o.PipelineBindPoint)
	info.viewMask = C.uint32_t(o.ViewMask)

	inputAttachmentCount := len(o.InputAttachments)
	colorAttachmentCount := len(o.ColorAttachments)
	resolveAttachmentCount := len(o.ResolveAttachments)
	preserveAttachmentCount := len(o.PreserveAttachments)

	if resolveAttachmentCount > 0 && resolveAttachmentCount != colorAttachmentCount {
		return nil, errors.Newf("in this subpass, %d color attachments are defined, but %d resolve attachments are defined- they should be equal", colorAttachmentCount, resolveAttachmentCount)
	}

	info.inputAttachmentCount = C.uint32_t(inputAttachmentCount)
	info.pInputAttachments = nil
	info.colorAttachmentCount = C.uint32_t(colorAttachmentCount)
	info.pColorAttachments = nil
	info.pResolveAttachments = nil
	info.pDepthStencilAttachment = nil
	info.preserveAttachmentCount = C.uint32_t(preserveAttachmentCount)
	info.pPreserveAttachments = nil

	var err error
	if inputAttachmentCount > 0 {
		info.pInputAttachments, err = common.AllocOptionSlice[C.VkAttachmentReference2, AttachmentReference2](allocator, o.InputAttachments)
		if err != nil {
			return nil, err
		}
	}

	if colorAttachmentCount > 0 {
		info.pColorAttachments, err = common.AllocOptionSlice[C.VkAttachmentReference2, AttachmentReference2](allocator, o.ColorAttachments)
		if err != nil {
			return nil, err
		}

		info.pResolveAttachments, err = common.AllocOptionSlice[C.VkAttachmentReference2, AttachmentReference2](allocator, o.ResolveAttachments)
		if err != nil {
			return nil, err
		}
	}

	if o.DepthStencilAttachment != nil {
		depthStencilPtr, err := common.AllocOptions(allocator, o.DepthStencilAttachment)
		if err != nil {
			return nil, err
		}

		info.pDepthStencilAttachment = (*C.VkAttachmentReference2)(depthStencilPtr)
	}

	if preserveAttachmentCount > 0 {
		attachmentsPtr := (*C.uint32_t)(allocator.Malloc(preserveAttachmentCount * int(unsafe.Sizeof(C.uint32_t(0)))))
		attachmentsSlice := unsafe.Slice(attachmentsPtr, preserveAttachmentCount)
		for i := 0; i < preserveAttachmentCount; i++ {
			attachmentsSlice[i] = C.uint32_t(o.PreserveAttachments[i])
		}
		info.pPreserveAttachments = attachmentsPtr
	}

	return preallocatedPointer, nil
}

////

// SubpassDependency2 specifies a subpass dependency
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDependency2.html
type SubpassDependency2 struct {
	// SrcSubpass is the subpass index of the first subpass in the dependency, or
	// core1_0.SubpassExternal
	SrcSubpass int
	// DstSubpass is the subpass index of the second subpass in the dependency, or
	// core1_0.SubpassExternal
	DstSubpass int
	// SrcStageMask specifies the source stage mask
	SrcStageMask core1_0.PipelineStageFlags
	// DstStageMask specifies the destination stage mask
	DstStageMask core1_0.PipelineStageFlags
	// SrcAccessMask specifies a source access mask
	SrcAccessMask core1_0.AccessFlags
	// DstAccessMask specifies a source access mask
	DstAccessMask core1_0.AccessFlags
	// DependencyFlags is a set of dependency flags
	DependencyFlags core1_0.DependencyFlags
	// ViewOffset controls which views in the source subpass the views in the destination
	// subpass depend on
	ViewOffset int

	common.NextOptions
}

func (o SubpassDependency2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassDependency2{})))
	}

	info := (*C.VkSubpassDependency2)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2
	info.pNext = next
	info.srcSubpass = C.uint32_t(o.SrcSubpass)
	info.dstSubpass = C.uint32_t(o.DstSubpass)
	info.srcStageMask = C.VkPipelineStageFlags(o.SrcStageMask)
	info.dstStageMask = C.VkPipelineStageFlags(o.DstStageMask)
	info.srcAccessMask = C.VkAccessFlags(o.SrcAccessMask)
	info.dstAccessMask = C.VkAccessFlags(o.DstAccessMask)
	info.dependencyFlags = C.VkDependencyFlags(o.DependencyFlags)
	info.viewOffset = C.int32_t(o.ViewOffset)

	return preallocatedPointer, nil
}

////

// RenderPassCreateInfo2 specifies parameters of a newly-created RenderPass
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2.html
type RenderPassCreateInfo2 struct {
	// Flags is reserved for future use
	Flags core1_0.RenderPassCreateFlags

	// Attachments is a slice of AttachmentDescription2 structures describing the attachments
	// used by the RenderPass
	Attachments []AttachmentDescription2
	// Subpasses is a slice of SubpassDescription2 structures describing each subpass
	Subpasses []SubpassDescription2
	// Dependencies is a slice of SubpassDependency2 structures describing dependencies
	// between pairs of subpasses
	Dependencies []SubpassDependency2

	// CorrelatedViewMasks is a slice of view masks indicating sets of views that may be
	// more efficient to render concurrently
	CorrelatedViewMasks []uint32

	common.NextOptions
}

func (o RenderPassCreateInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderPassCreateInfo2{})))
	}

	info := (*C.VkRenderPassCreateInfo2)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2
	info.pNext = next
	info.flags = C.VkRenderPassCreateFlags(o.Flags)

	attachmentCount := len(o.Attachments)
	subpassCount := len(o.Subpasses)
	dependencyCount := len(o.Dependencies)
	viewMaskCount := len(o.CorrelatedViewMasks)

	info.attachmentCount = C.uint32_t(attachmentCount)
	info.pAttachments = nil
	info.subpassCount = C.uint32_t(subpassCount)
	info.pSubpasses = nil
	info.dependencyCount = C.uint32_t(dependencyCount)
	info.pDependencies = nil
	info.correlatedViewMaskCount = C.uint32_t(viewMaskCount)
	info.pCorrelatedViewMasks = nil

	var err error
	if attachmentCount > 0 {
		info.pAttachments, err = common.AllocOptionSlice[C.VkAttachmentDescription2, AttachmentDescription2](allocator, o.Attachments)
		if err != nil {
			return nil, err
		}
	}

	if subpassCount > 0 {
		info.pSubpasses, err = common.AllocOptionSlice[C.VkSubpassDescription2, SubpassDescription2](allocator, o.Subpasses)
		if err != nil {
			return nil, err
		}
	}

	if dependencyCount > 0 {
		info.pDependencies, err = common.AllocOptionSlice[C.VkSubpassDependency2, SubpassDependency2](allocator, o.Dependencies)
		if err != nil {
			return nil, err
		}
	}

	if viewMaskCount > 0 {
		viewMaskPtr := (*C.uint32_t)(allocator.Malloc(viewMaskCount * int(unsafe.Sizeof(C.uint32_t(0)))))
		viewMaskSlice := unsafe.Slice(viewMaskPtr, viewMaskCount)
		for i := 0; i < viewMaskCount; i++ {
			viewMaskSlice[i] = C.uint32_t(o.CorrelatedViewMasks[i])
		}
		info.pCorrelatedViewMasks = viewMaskPtr
	}

	return preallocatedPointer, nil
}

////

// AttachmentDescriptionStencilLayout specifies an attachment description
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionStencilLayout.html
type AttachmentDescriptionStencilLayout struct {
	// StencilInitialLayout is the layout of the stencil aspect of the attachment Image
	// subresource will be in when a RenderPass instance begins
	StencilInitialLayout core1_0.ImageLayout
	// StencilFinalLayout is the layout the stencil aspect of the attachment Image subresource
	// will be transitioned to when a RenderPass instance ends
	StencilFinalLayout core1_0.ImageLayout

	common.NextOptions
}

func (o AttachmentDescriptionStencilLayout) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAttachmentDescriptionStencilLayout{})))
	}

	info := (*C.VkAttachmentDescriptionStencilLayout)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_STENCIL_LAYOUT
	info.pNext = next
	info.stencilInitialLayout = C.VkImageLayout(o.StencilInitialLayout)
	info.stencilFinalLayout = C.VkImageLayout(o.StencilFinalLayout)

	return preallocatedPointer, nil
}

////

// AttachmentReferenceStencilLayout specifies an attachment description
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReferenceStencilLayout.html
type AttachmentReferenceStencilLayout struct {
	// StencilLayout specifies the layout the stencil aspect of the attachment uses during hte subpass
	StencilLayout core1_0.ImageLayout

	common.NextOptions
}

func (o AttachmentReferenceStencilLayout) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAttachmentReferenceStencilLayout{})))
	}

	info := (*C.VkAttachmentReferenceStencilLayout)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_STENCIL_LAYOUT
	info.pNext = next
	info.stencilLayout = C.VkImageLayout(o.StencilLayout)

	return preallocatedPointer, nil
}

////

// RenderPassAttachmentBeginInfo specifies Image objects to be used as Framebuffer attachments
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html
type RenderPassAttachmentBeginInfo struct {
	// Attachments is a slice of ImageView objects, each of which will be used as the corresponding
	// attachment in the RenderPass instance
	Attachments []core1_0.ImageView

	common.NextOptions
}

func (o RenderPassAttachmentBeginInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderPassAttachmentBeginInfo{})))
	}

	info := (*C.VkRenderPassAttachmentBeginInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO
	info.pNext = next

	count := len(o.Attachments)
	info.attachmentCount = C.uint32_t(count)
	info.pAttachments = nil

	if count > 0 {
		info.pAttachments = (*C.VkImageView)(allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkImageView{}))))
		attachmentSlice := unsafe.Slice(info.pAttachments, count)
		for i := 0; i < count; i++ {
			attachmentSlice[i] = C.VkImageView(unsafe.Pointer(o.Attachments[i].Handle()))
		}
	}

	return preallocatedPointer, nil
}

////

// SubpassDescriptionDepthStencilResolve specifies depth/stencil resolve operations for
// a subpass
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionDepthStencilResolve.html
type SubpassDescriptionDepthStencilResolve struct {
	// DepthResolveMode describes the depth resolve mode
	DepthResolveMode ResolveModeFlags
	// StencilResolveMode describes the stencil resolve mode
	StencilResolveMode ResolveModeFlags
	// DepthStencilResolveAttachment defines the depth/stencil resolve attachment
	// for this subpass and its layout
	DepthStencilResolveAttachment *AttachmentReference2

	common.NextOptions
}

func (o SubpassDescriptionDepthStencilResolve) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassDescriptionDepthStencilResolve{})))
	}

	info := (*C.VkSubpassDescriptionDepthStencilResolve)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE
	info.pNext = next
	info.depthResolveMode = C.VkResolveModeFlagBits(o.DepthResolveMode)
	info.stencilResolveMode = C.VkResolveModeFlagBits(o.StencilResolveMode)
	info.pDepthStencilResolveAttachment = nil

	if o.DepthStencilResolveAttachment != nil {
		attachment, err := common.AllocOptions(allocator, o.DepthStencilResolveAttachment)
		if err != nil {
			return nil, err
		}

		info.pDepthStencilResolveAttachment = (*C.VkAttachmentReference2)(attachment)
	}

	return preallocatedPointer, nil
}
