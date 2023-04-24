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
	"github.com/vkngwrapper/core/v2/common"
)

// DescriptorSetAllocateInfo specifies the allocation parameters for DescritporSet objects
type DescriptorSetAllocateInfo struct {
	// DescriptorPool is the pool which the sets will be allocated from
	DescriptorPool DescriptorPool

	// SetLayouts is a slice of DescriptorSetLayout objects, which each member specifying how the
	// corresponding DescriptorSet is allocated
	SetLayouts []DescriptorSetLayout

	common.NextOptions
}

func (o DescriptorSetAllocateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.DescriptorPool == nil {
		return nil, errors.New("core1_0.DescriptorSetAllocateInfo.DescriptorPool cannot be nil")
	}
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDescriptorSetAllocateInfo)
	}

	createInfo := (*C.VkDescriptorSetAllocateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO
	createInfo.pNext = next
	createInfo.descriptorPool = C.VkDescriptorPool(unsafe.Pointer(o.DescriptorPool.Handle()))

	setCount := len(o.SetLayouts)
	createInfo.descriptorSetCount = C.uint32_t(setCount)
	createInfo.pSetLayouts = nil

	if setCount > 0 {
		layoutsPtr := (*C.VkDescriptorSetLayout)(allocator.Malloc(setCount * int(unsafe.Sizeof([1]C.VkDescriptorSetLayout{}))))
		layoutsSlice := ([]C.VkDescriptorSetLayout)(unsafe.Slice(layoutsPtr, setCount))

		for i := 0; i < setCount; i++ {
			if o.SetLayouts[i] == nil {
				return nil, errors.New("core1_0.DescriptorSetAllocateInfo.SetLayouts cannot contain any nil elements")
			}

			layoutsSlice[i] = C.VkDescriptorSetLayout(unsafe.Pointer(o.SetLayouts[i].Handle()))
		}

		createInfo.pSetLayouts = layoutsPtr
	}

	return preallocatedPointer, nil
}

// DescriptorImageInfo specifies descriptor Image information
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorImageInfo.html
type DescriptorImageInfo struct {
	// Sampler is a Sampler object, and is used in descriptor updates for DescriptorTypeSampler and
	// DescriptorTypeCombinedImageSampler descriptors if the binding being update does not use
	// immutable sampler
	Sampler Sampler
	// ImageView is an ImageView object, and is used in descriptor updates for DescriptorTypeSampledImage,
	// DescriptorTypeStorageImage, DescriptorTypeCombinedImageSampler, and DescriptorTypeInputAttachment
	ImageView ImageView
	// ImageLayout is the layout that the Image subresources accessible form ImageView will be in
	// at the time this descriptor is accessed
	ImageLayout ImageLayout
}

// DescriptorBufferInfo specifies descriptor Buffer information
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBufferInfo.html
type DescriptorBufferInfo struct {
	// Buffer is the Buffer resource
	Buffer Buffer
	// Offset is the offset in bytes from the start of Buffer
	Offset int
	// Range is the size in bytes that is used for this descriptor update
	Range int
}

// WriteDescriptorSet specifies the parameters of a DescriptorSet write operation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html
type WriteDescriptorSet struct {
	// DstSet is the destination DescriptorSet to update
	DstSet DescriptorSet
	// DstBinding is the descriptor binding within that set
	DstBinding int
	// DstArrayElement is the starting element in that array
	DstArrayElement int

	// DescriptorType specifies the type of each descriptor in ImageInfo, BufferInfo, or
	// TexelBufferView
	DescriptorType DescriptorType

	// ImageInfo is a slice of DescriptorImageInfo structures or is ignored
	ImageInfo []DescriptorImageInfo
	// BufferInfo is a slice of DescriptorBufferInfo structures or is ignored
	BufferInfo []DescriptorBufferInfo
	// TexelBufferView is a slice of BufferView objects or is ignored
	TexelBufferView []BufferView

	common.NextOptions
}

type WriteDescriptorSetExtensionSource interface {
	WriteDescriptorSetCount() int
}

func (o WriteDescriptorSet) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkWriteDescriptorSet)
	}

	createInfo := (*C.VkWriteDescriptorSet)(preallocatedPointer)
	imageInfoCount := len(o.ImageInfo)
	bufferInfoCount := len(o.BufferInfo)
	texelBufferCount := len(o.TexelBufferView)
	var extSource WriteDescriptorSetExtensionSource

	nextObj := o.Next
	for nextObj != nil {
		var isExtSource bool
		extSource, isExtSource = o.Next.(WriteDescriptorSetExtensionSource)
		if isExtSource {
			break
		}

		nextObj = nextObj.NextOptionsInChain()
	}

	if imageInfoCount > 0 && texelBufferCount > 0 {
		return nil, errors.New("a WriteDescriptorSet may have one or more ImageInfo sources OR one or more TexelBufferView sources, but not both")
	}

	if imageInfoCount > 0 && bufferInfoCount > 0 {
		return nil, errors.New("a WriteDescriptorSet may have one or more ImageInfo sources OR one or more BufferInfo sources, but not both")
	}

	if bufferInfoCount > 0 && texelBufferCount > 0 {
		return nil, errors.New("a WriteDescriptorSet may have one or more BufferInfo sources OR one or more TexelBufferView sources, but not both")
	}

	if imageInfoCount == 0 && bufferInfoCount == 0 && texelBufferCount == 0 && extSource == nil {
		return nil, errors.New("a WriteDescriptorSet must have a source to write the descriptor from: ImageInfo, BufferInfo, TexelBufferView, or an extension source")
	}

	if extSource != nil && (bufferInfoCount > 0 || texelBufferCount > 0 || imageInfoCount > 0) {
		return nil, errors.New("an extension descriptor source for a WriteDescriptorSet has been included, but so has a traditional descriptor source: ImageInfo, BufferInfo, or TexelBufferView")
	}

	if o.DstSet == nil {
		return nil, errors.New("core1_0.WriteDescriptorSet.DstSet cannot be nil")
	}

	createInfo.sType = C.VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
	createInfo.pNext = next

	createInfo.dstSet = C.VkDescriptorSet(unsafe.Pointer(o.DstSet.Handle()))
	createInfo.dstBinding = C.uint32_t(o.DstBinding)
	createInfo.dstArrayElement = C.uint32_t(o.DstArrayElement)

	createInfo.descriptorType = C.VkDescriptorType(o.DescriptorType)

	createInfo.descriptorCount = 0
	createInfo.pImageInfo = nil
	createInfo.pBufferInfo = nil
	createInfo.pTexelBufferView = nil

	if extSource != nil {
		createInfo.descriptorCount = C.uint32_t(extSource.WriteDescriptorSetCount())
	} else if imageInfoCount > 0 {
		createInfo.descriptorCount = C.uint32_t(imageInfoCount)
		imageInfoPtr := (*C.VkDescriptorImageInfo)(allocator.Malloc(imageInfoCount * C.sizeof_struct_VkDescriptorImageInfo))
		imageInfoSlice := ([]C.VkDescriptorImageInfo)(unsafe.Slice(imageInfoPtr, imageInfoCount))
		for i := 0; i < imageInfoCount; i++ {
			imageInfoSlice[i].sampler = nil
			imageInfoSlice[i].imageView = nil

			if o.ImageInfo[i].Sampler != nil {
				imageInfoSlice[i].sampler = C.VkSampler(unsafe.Pointer(o.ImageInfo[i].Sampler.Handle()))
			}

			if o.ImageInfo[i].ImageView != nil {
				imageInfoSlice[i].imageView = C.VkImageView(unsafe.Pointer(o.ImageInfo[i].ImageView.Handle()))
			}

			imageInfoSlice[i].imageLayout = C.VkImageLayout(o.ImageInfo[i].ImageLayout)
		}

		createInfo.pImageInfo = imageInfoPtr
	} else if bufferInfoCount > 0 {
		createInfo.descriptorCount = C.uint32_t(bufferInfoCount)
		bufferInfoPtr := (*C.VkDescriptorBufferInfo)(allocator.Malloc(bufferInfoCount * C.sizeof_struct_VkDescriptorBufferInfo))
		bufferInfoSlice := ([]C.VkDescriptorBufferInfo)(unsafe.Slice(bufferInfoPtr, bufferInfoCount))
		for i := 0; i < bufferInfoCount; i++ {
			bufferInfoSlice[i].buffer = nil

			if o.BufferInfo[i].Buffer != nil {
				bufferInfoSlice[i].buffer = C.VkBuffer(unsafe.Pointer(o.BufferInfo[i].Buffer.Handle()))
			}

			bufferInfoSlice[i].offset = C.VkDeviceSize(o.BufferInfo[i].Offset)
			bufferInfoSlice[i]._range = C.VkDeviceSize(o.BufferInfo[i].Range)
		}

		createInfo.pBufferInfo = bufferInfoPtr
	} else if texelBufferCount > 0 {
		createInfo.descriptorCount = C.uint32_t(texelBufferCount)
		texelBufferPtr := (*C.VkBufferView)(allocator.Malloc(texelBufferCount * int(unsafe.Sizeof([1]C.VkBufferView{}))))
		texelBufferSlice := ([]C.VkBufferView)(unsafe.Slice(texelBufferPtr, texelBufferCount))
		for i := 0; i < texelBufferCount; i++ {
			texelBufferSlice[i] = nil

			if o.TexelBufferView[i] != nil {
				texelBufferSlice[i] = C.VkBufferView(unsafe.Pointer(o.TexelBufferView[i].Handle()))
			}
		}

		createInfo.pTexelBufferView = texelBufferPtr
	}

	return preallocatedPointer, nil
}

// CopyDescriptorSet specifies a copy descriptor set operation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyDescriptorSet.html
type CopyDescriptorSet struct {
	// SrcSet is the source descriptor set
	SrcSet DescriptorSet
	// SrcBinding is the source descriptor binding
	SrcBinding int
	// SrcArrayElement is the source descriptor array element
	SrcArrayElement int

	// DstSet is the destination descriptor set
	DstSet DescriptorSet
	// DstBinding is the destination descriptor binding
	DstBinding int
	// DstArrayElement is the destination descriptor array element
	DstArrayElement int

	// DescriptorCount is number of descriptors to copy from source to destination
	DescriptorCount int

	common.NextOptions
}

func (o CopyDescriptorSet) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkCopyDescriptorSet)
	}
	if o.SrcSet == nil {
		return nil, errors.New("core1_0.CopyDescriptorSet.SrcSet cannot be nil")
	}
	if o.DstSet == nil {
		return nil, errors.New("core1_0.CopyDescriptorSet.DstSet cannot be nil")
	}

	createInfo := (*C.VkCopyDescriptorSet)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_COPY_DESCRIPTOR_SET
	createInfo.pNext = next

	createInfo.srcSet = (C.VkDescriptorSet)(unsafe.Pointer(o.SrcSet.Handle()))
	createInfo.srcBinding = C.uint32_t(o.SrcBinding)
	createInfo.srcArrayElement = C.uint32_t(o.SrcArrayElement)

	createInfo.dstSet = (C.VkDescriptorSet)(unsafe.Pointer(o.DstSet.Handle()))
	createInfo.dstBinding = C.uint32_t(o.DstBinding)
	createInfo.dstArrayElement = C.uint32_t(o.DstArrayElement)

	createInfo.descriptorCount = C.uint32_t(o.DescriptorCount)

	return preallocatedPointer, nil
}
