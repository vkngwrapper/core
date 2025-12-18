package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

const (
	// ImageCreateSparseBinding specifies that the Image will be backed using sparse memory
	// binding
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
	ImageCreateSparseBinding ImageCreateFlags = C.VK_IMAGE_CREATE_SPARSE_BINDING_BIT
	// ImageCreateSparseResidency specifies that the Image can be partially backed using sparse
	// memory binding
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
	ImageCreateSparseResidency ImageCreateFlags = C.VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT
	// ImageCreateSparseAliased specifies that the Image will be backed using sparse memory binding
	// with memory ranges that might also simultaneously be backing another Image or another portion
	// of the same Image
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
	ImageCreateSparseAliased ImageCreateFlags = C.VK_IMAGE_CREATE_SPARSE_ALIASED_BIT
	// ImageCreateMutableFormat specifies that the Image can be used to create an ImageView with
	// a different format from the Image
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
	ImageCreateMutableFormat ImageCreateFlags = C.VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT
	// ImageCreateCubeCompatible specifies that the Image can be used to create an ImageView of
	// type ImageViewTypeCube or ImageViewTypeCubeArray
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
	ImageCreateCubeCompatible ImageCreateFlags = C.VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT

	// ImageLayoutUndefined specifies that the layout is unknown
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutUndefined ImageLayout = C.VK_IMAGE_LAYOUT_UNDEFINED
	// ImageLayoutGeneral supports all types of Device access
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutGeneral ImageLayout = C.VK_IMAGE_LAYOUT_GENERAL
	// ImageLayoutColorAttachmentOptimal must only be used as a color or resolve attachment
	// in a Framebuffer
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutColorAttachmentOptimal ImageLayout = C.VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
	// ImageLayoutDepthStencilAttachmentOptimal specifies a layout for both the depth and stencil
	// aspects of a depth/stencil format Image allowing read and write access as a depth/stencil
	// attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutDepthStencilAttachmentOptimal ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL
	// ImageLayoutDepthStencilReadOnlyOptimal specifies a layout for both the depth and stencil
	// aspects of a depth/stencil format Image allowing read only access as a depth/stencil attachment
	// or in shaders as a sampled Image, combined Image/Sampler, or input attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutDepthStencilReadOnlyOptimal ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL
	// ImageLayoutShaderReadOnlyOptimal specifies a layout allowing read-only access in a shader
	// as a sampled Image, combined Image/Sampler, or input attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutShaderReadOnlyOptimal ImageLayout = C.VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL
	// ImageLayoutTransferSrcOptimal must only be used as a source Image of a transfer command
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutTransferSrcOptimal ImageLayout = C.VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL
	// ImageLayoutTransferDstOptimal must only be used as a destination Image of a transfer
	// command
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutTransferDstOptimal ImageLayout = C.VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL
	// ImageLayoutPreInitialized specifies that an Image object's memory is in a defined layout
	// and can be populated by data, but that it has not yet been initialized by the driver
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutPreInitialized ImageLayout = C.VK_IMAGE_LAYOUT_PREINITIALIZED

	// ImageTilingOptimal specifies optimal tiling
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html
	ImageTilingOptimal ImageTiling = C.VK_IMAGE_TILING_OPTIMAL
	// ImageTilingLinear specifies linear tiling
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html
	ImageTilingLinear ImageTiling = C.VK_IMAGE_TILING_LINEAR

	// ImageType1D specifies a one-dimensional Image
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html
	ImageType1D ImageType = C.VK_IMAGE_TYPE_1D
	// ImageType2D specifies a two-dimensional Image
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html
	ImageType2D ImageType = C.VK_IMAGE_TYPE_2D
	// ImageType3D specifies a three-dimensional Image
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html
	ImageType3D ImageType = C.VK_IMAGE_TYPE_3D

	// ImageUsageTransferSrc specifies that the Image can be used as the source of a transfer
	// command
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html
	ImageUsageTransferSrc ImageUsageFlags = C.VK_IMAGE_USAGE_TRANSFER_SRC_BIT
	// ImageUsageTransferDst specifies that the Image can be used as the destination of a
	// transfer command
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html
	ImageUsageTransferDst ImageUsageFlags = C.VK_IMAGE_USAGE_TRANSFER_DST_BIT
	// ImageUsageSampled specifies that the Image can be used to create an ImageView suitable
	// for occupying a DescriptorSet slot either of DescriptorTypeSampledImage or
	// DescriptorTypeCombinedImageSampler, and be sampled by a shader
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html
	ImageUsageSampled ImageUsageFlags = C.VK_IMAGE_USAGE_SAMPLED_BIT
	// ImageUsageStorage specifies that the Image can be used to create an ImageView suitable for
	// occupying a DescriptorSet slot of type DescriptorTypeStorageImage
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html
	ImageUsageStorage ImageUsageFlags = C.VK_IMAGE_USAGE_STORAGE_BIT
	// ImageUsageColorAttachment specifies that the image can be used to create an ImageView
	// suitable for use as a color or resolve attachment in a Framebuffer
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html
	ImageUsageColorAttachment ImageUsageFlags = C.VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT
	// ImageUsageDepthStencilAttachment specifies that the Image can be used to create an ImageView
	// suitable for use as a depth/stencil or depth/stencil resolve attachment in a Framebuffer
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html
	ImageUsageDepthStencilAttachment ImageUsageFlags = C.VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT
	// ImageUsageTransientAttachment specifies that implementations may support using memory
	// allocations with MemoryPropertyLazilyAllocated to back an image with this usage
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html
	ImageUsageTransientAttachment ImageUsageFlags = C.VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT
	// ImageUsageInputAttachment specifies that the image can be used to create an ImageView
	// suitable for occupying a DescriptorSet slot of type DescriptorTypeInputAttachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html
	ImageUsageInputAttachment ImageUsageFlags = C.VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT

	// Samples1 specifies an Image with one sample per pixel
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html
	Samples1 SampleCountFlags = C.VK_SAMPLE_COUNT_1_BIT
	// Samples2 specifies an Image with 2 samples per pixel
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html
	Samples2 SampleCountFlags = C.VK_SAMPLE_COUNT_2_BIT
	// Samples4 specifies an Image with 4 samples per pixel
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html
	Samples4 SampleCountFlags = C.VK_SAMPLE_COUNT_4_BIT
	// Samples8 specifies an Image with 8 samples per pixel
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html
	Samples8 SampleCountFlags = C.VK_SAMPLE_COUNT_8_BIT
	// Samples16 specifies an Image with 16 samples per pixel
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html
	Samples16 SampleCountFlags = C.VK_SAMPLE_COUNT_16_BIT
	// Samples32 specifies an Image with 32 samples per pixel
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html
	Samples32 SampleCountFlags = C.VK_SAMPLE_COUNT_32_BIT
	// Samples64 specifies an Image with 64 samples per pixel
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html
	Samples64 SampleCountFlags = C.VK_SAMPLE_COUNT_64_BIT
)

func init() {
	ImageCreateSparseBinding.Register("Sparse Binding")
	ImageCreateSparseResidency.Register("Sparse Residency")
	ImageCreateSparseAliased.Register("Sparse Aliased")
	ImageCreateMutableFormat.Register("Mutable Format")
	ImageCreateCubeCompatible.Register("Cube Compatible")

	ImageLayoutUndefined.Register("Undefined")
	ImageLayoutGeneral.Register("General")
	ImageLayoutColorAttachmentOptimal.Register("Color Attachment")
	ImageLayoutDepthStencilAttachmentOptimal.Register("Depth & Stencil Attachment")
	ImageLayoutDepthStencilReadOnlyOptimal.Register("Depth & Stencil Read-Only")
	ImageLayoutShaderReadOnlyOptimal.Register("Shader Read-Only")
	ImageLayoutTransferSrcOptimal.Register("Transfer Source")
	ImageLayoutTransferDstOptimal.Register("Transfer Destination")
	ImageLayoutPreInitialized.Register("Pre-Initialized")

	ImageTilingOptimal.Register("Optimal")
	ImageTilingLinear.Register("Linear")

	ImageType1D.Register("1D")
	ImageType2D.Register("2D")
	ImageType3D.Register("3D")

	ImageUsageTransferSrc.Register("Transfer Source")
	ImageUsageTransferDst.Register("Transfer Destination")
	ImageUsageSampled.Register("Sampled")
	ImageUsageStorage.Register("Storage")
	ImageUsageColorAttachment.Register("Color Attachment")
	ImageUsageDepthStencilAttachment.Register("Depth Stencil Attachment")
	ImageUsageTransientAttachment.Register("Transient Attachment")
	ImageUsageInputAttachment.Register("Input Attachment")

	Samples1.RegisterSamples("1 Samples", 1)
	Samples2.RegisterSamples("2 Samples", 2)
	Samples4.RegisterSamples("4 Samples", 4)
	Samples8.RegisterSamples("8 Samples", 8)
	Samples16.RegisterSamples("16 Samples", 16)
	Samples32.RegisterSamples("32 Samples", 32)
	Samples64.RegisterSamples("64 Samples", 64)
}

// ImageCreateInfo specifies the parameters of a newly-created Image object
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html
type ImageCreateInfo struct {
	// Flags describes additional parameters of the Image
	Flags ImageCreateFlags
	// ImageType specifies the basic dimensionality of the Image
	ImageType ImageType
	// Format describes the format and type of the texel blocks that will be contained in the Image
	Format Format
	// Extent Describes the number of data elements in each dimension of the base level
	Extent Extent3D

	// MipLevels describes the number of levels of detail available for minified sampling of the image
	MipLevels int
	// ArrayLayers is the number of layers in the IMage
	ArrayLayers int

	// Samples specifies the number of samples per texel
	Samples SampleCountFlags
	// Tiling specifies the tiling arrangement of the texel blocks in memory
	Tiling ImageTiling
	// Usage describes the intended usage of the Image
	Usage ImageUsageFlags
	// SharingMode specifies the sharing mode of the Image when it will be accessed by multiple
	// Queue families
	SharingMode SharingMode

	// QueueFamilyIndices is a slice of queue families that will access this Image
	QueueFamilyIndices []uint32

	// InitialLayout specifies the initial ImageLayout of all Image subresources of the Image
	InitialLayout ImageLayout

	common.NextOptions
}

func (o ImageCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkImageCreateInfo{})))
	}

	createInfo := (*C.VkImageCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO
	createInfo.pNext = next
	createInfo.flags = C.VkImageCreateFlags(o.Flags)
	createInfo.imageType = C.VkImageType(o.ImageType)
	createInfo.format = C.VkFormat(o.Format)
	createInfo.mipLevels = C.uint32_t(o.MipLevels)
	createInfo.arrayLayers = C.uint32_t(o.ArrayLayers)
	createInfo.samples = C.VkSampleCountFlagBits(o.Samples)
	createInfo.tiling = C.VkImageTiling(o.Tiling)
	createInfo.usage = C.VkImageUsageFlags(o.Usage)
	createInfo.sharingMode = C.VkSharingMode(o.SharingMode)
	createInfo.initialLayout = C.VkImageLayout(o.InitialLayout)
	indexCount := len(o.QueueFamilyIndices)
	createInfo.queueFamilyIndexCount = C.uint32_t(indexCount)
	createInfo.pQueueFamilyIndices = nil

	createInfo.extent.width = C.uint32_t(o.Extent.Width)
	createInfo.extent.height = C.uint32_t(o.Extent.Height)
	createInfo.extent.depth = C.uint32_t(o.Extent.Depth)

	if indexCount > 0 {
		queueIndicesPtr := (*C.uint32_t)(allocator.Malloc(indexCount * int(unsafe.Sizeof(C.uint32_t(0)))))
		queueIndicesSlice := ([]C.uint32_t)(unsafe.Slice(queueIndicesPtr, indexCount))

		for i := 0; i < indexCount; i++ {
			queueIndicesSlice[i] = C.uint32_t(o.QueueFamilyIndices[i])
		}

		createInfo.pQueueFamilyIndices = queueIndicesPtr
	}

	return unsafe.Pointer(createInfo), nil
}
