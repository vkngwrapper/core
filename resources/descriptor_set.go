package resources

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type DescriptorSetOptions struct {
	DescriptorPool DescriptorPool

	AllocationLayouts []DescriptorSetLayout

	core.HaveNext
}

func (o *DescriptorSetOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkDescriptorSetAllocateInfo)(allocator.Malloc(C.sizeof_struct_VkDescriptorSetAllocateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO
	createInfo.pNext = next
	createInfo.descriptorPool = C.VkDescriptorPool(unsafe.Pointer(o.DescriptorPool.Handle()))

	setCount := len(o.AllocationLayouts)
	createInfo.descriptorSetCount = C.uint32_t(setCount)
	createInfo.pSetLayouts = nil

	if setCount > 0 {
		layoutsPtr := (*C.VkDescriptorSetLayout)(allocator.Malloc(setCount * int(unsafe.Sizeof([1]C.VkDescriptorSetLayout{}))))
		layoutsSlice := ([]C.VkDescriptorSetLayout)(unsafe.Slice(layoutsPtr, setCount))

		for i := 0; i < setCount; i++ {
			layoutsSlice[i] = C.VkDescriptorSetLayout(unsafe.Pointer(o.AllocationLayouts[i].Handle()))
		}

		createInfo.pSetLayouts = layoutsPtr
	}

	return unsafe.Pointer(createInfo), nil
}

type vulkanDescriptorSet struct {
	handle loader.VkDescriptorSet
}

func (s *vulkanDescriptorSet) Handle() loader.VkDescriptorSet {
	return s.handle
}

type DescriptorImageInfo struct {
	Sampler     Sampler
	ImageView   ImageView
	ImageLayout core.ImageLayout
}

type DescriptorBufferInfo struct {
	Buffer Buffer
	Offset uint64
	Range  uint64
}

type WriteDescriptorSetOptions struct {
	Destination             DescriptorSet
	DestinationBinding      int
	DestinationArrayElement int

	DescriptorType core.DescriptorType

	ImageInfo       []DescriptorImageInfo
	BufferInfo      []DescriptorBufferInfo
	TexelBufferView []BufferView

	core.HaveNext
}

type WriteDescriptorSetExtensionSource interface {
	WriteDescriptorSetCount() int
}

func (o *WriteDescriptorSetOptions) populate(allocator *cgoparam.Allocator, createInfo *C.VkWriteDescriptorSet, next unsafe.Pointer) (unsafe.Pointer, error) {
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

		nextObj = nextObj.NextInChain()
	}

	if imageInfoCount > 0 && texelBufferCount > 0 {
		return nil, errors.New("a WriteDescriptorSetOptions may have one or more ImageInfo sources OR one or more TexelBufferView sources, but not both")
	}

	if imageInfoCount > 0 && bufferInfoCount > 0 {
		return nil, errors.New("a WriteDescriptorSetOptions may have one or more ImageInfo sources OR one or more BufferInfo sources, but not both")
	}

	if bufferInfoCount > 0 && texelBufferCount > 0 {
		return nil, errors.New("a WriteDescriptorSetOptions may have one or more BufferInfo sources OR one or more TexelBufferView sources, but not both")
	}

	if imageInfoCount == 0 && bufferInfoCount == 0 && texelBufferCount == 0 && extSource == nil {
		return nil, errors.New("a WriteDescriptorSetOptions must have a source to write the descriptor from: ImageInfo, BufferInfo, TexelBufferView, or an extension source")
	}

	if extSource != nil && (bufferInfoCount > 0 || texelBufferCount > 0 || imageInfoCount > 0) {
		return nil, errors.New("an extension descriptor source for a WriteDescriptorSetOptions has been included, but so has a traditional descriptor source: ImageInfo, BufferInfo, or TexelBufferView")
	}

	createInfo.sType = C.VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
	createInfo.pNext = next

	createInfo.dstSet = C.VkDescriptorSet(unsafe.Pointer(o.Destination.Handle()))
	createInfo.dstBinding = C.uint32_t(o.DestinationBinding)
	createInfo.dstArrayElement = C.uint32_t(o.DestinationArrayElement)

	createInfo.descriptorType = C.VkDescriptorType(o.DescriptorType)

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
			imageInfoSlice[i].sampler = C.VkSampler(unsafe.Pointer(o.ImageInfo[i].Sampler.Handle()))
			imageInfoSlice[i].imageView = C.VkImageView(unsafe.Pointer(o.ImageInfo[i].ImageView.Handle()))
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

	return unsafe.Pointer(createInfo), nil
}

func (o *WriteDescriptorSetOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkWriteDescriptorSet)(allocator.Malloc(C.sizeof_struct_VkWriteDescriptorSet))
	return o.populate(allocator, createInfo, next)
}

type CopyDescriptorSetOptions struct {
	Source             DescriptorSet
	SourceBinding      int
	SourceArrayElement int

	Destination             DescriptorSet
	DestinationBinding      int
	DestinationArrayElement int

	Count int

	core.HaveNext
}

func (o *CopyDescriptorSetOptions) populate(allocator *cgoparam.Allocator, createInfo *C.VkCopyDescriptorSet, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo.sType = C.VK_STRUCTURE_TYPE_COPY_DESCRIPTOR_SET
	createInfo.pNext = next

	createInfo.srcSet = (C.VkDescriptorSet)(unsafe.Pointer(o.Source.Handle()))
	createInfo.srcBinding = C.uint32_t(o.SourceBinding)
	createInfo.srcArrayElement = C.uint32_t(o.SourceArrayElement)

	createInfo.dstSet = (C.VkDescriptorSet)(unsafe.Pointer(o.Destination.Handle()))
	createInfo.dstBinding = C.uint32_t(o.DestinationBinding)
	createInfo.dstArrayElement = C.uint32_t(o.DestinationArrayElement)

	createInfo.descriptorCount = C.uint32_t(o.Count)

	return unsafe.Pointer(createInfo), nil
}

func (o *CopyDescriptorSetOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkCopyDescriptorSet)(allocator.Malloc(C.sizeof_struct_VkCopyDescriptorSet))
	return o.populate(allocator, createInfo, next)
}
