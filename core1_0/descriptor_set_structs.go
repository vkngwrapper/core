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

type DescriptorSetAllocateInfo struct {
	DescriptorPool DescriptorPool

	SetLayouts []DescriptorSetLayout

	common.NextOptions
}

func (o DescriptorSetAllocateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
			layoutsSlice[i] = C.VkDescriptorSetLayout(unsafe.Pointer(o.SetLayouts[i].Handle()))
		}

		createInfo.pSetLayouts = layoutsPtr
	}

	return preallocatedPointer, nil
}

type DescriptorImageInfo struct {
	Sampler     Sampler
	ImageView   ImageView
	ImageLayout ImageLayout
}

type DescriptorBufferInfo struct {
	Buffer Buffer
	Offset int
	Range  int
}

type WriteDescriptorSet struct {
	DstSet          DescriptorSet
	DstBinding      int
	DstArrayElement int

	DescriptorType DescriptorType

	ImageInfo       []DescriptorImageInfo
	BufferInfo      []DescriptorBufferInfo
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
			bufferInfoSlice[i].buffer = C.VkBuffer(unsafe.Pointer(o.BufferInfo[i].Buffer.Handle()))
			bufferInfoSlice[i].offset = C.VkDeviceSize(o.BufferInfo[i].Offset)
			bufferInfoSlice[i]._range = C.VkDeviceSize(o.BufferInfo[i].Range)
		}

		createInfo.pBufferInfo = bufferInfoPtr
	} else if texelBufferCount > 0 {
		createInfo.descriptorCount = C.uint32_t(texelBufferCount)
		texelBufferPtr := (*C.VkBufferView)(allocator.Malloc(texelBufferCount * int(unsafe.Sizeof([1]C.VkBufferView{}))))
		texelBufferSlice := ([]C.VkBufferView)(unsafe.Slice(texelBufferPtr, texelBufferCount))
		for i := 0; i < texelBufferCount; i++ {
			texelBufferSlice[i] = C.VkBufferView(unsafe.Pointer(o.TexelBufferView[i].Handle()))
		}

		createInfo.pTexelBufferView = texelBufferPtr
	}

	return preallocatedPointer, nil
}

type CopyDescriptorSet struct {
	SrcSet          DescriptorSet
	SrcBinding      int
	SrcArrayElement int

	DstSet          DescriptorSet
	DstBinding      int
	DstArrayElement int

	DescriptorCount int

	common.NextOptions
}

func (o CopyDescriptorSet) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkCopyDescriptorSet)
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
