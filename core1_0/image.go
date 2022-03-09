package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

const (
	ImageCreateSparseBinding   common.ImageCreateFlags = C.VK_IMAGE_CREATE_SPARSE_BINDING_BIT
	ImageCreateSparseResidency common.ImageCreateFlags = C.VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT
	ImageCreateSparseAliased   common.ImageCreateFlags = C.VK_IMAGE_CREATE_SPARSE_ALIASED_BIT
	ImageCreateMutableFormat   common.ImageCreateFlags = C.VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT
	ImageCreateCubeCompatible  common.ImageCreateFlags = C.VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT

	ImageLayoutUndefined                     common.ImageLayout = C.VK_IMAGE_LAYOUT_UNDEFINED
	ImageLayoutGeneral                       common.ImageLayout = C.VK_IMAGE_LAYOUT_GENERAL
	ImageLayoutColorAttachmentOptimal        common.ImageLayout = C.VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
	ImageLayoutDepthStencilAttachmentOptimal common.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL
	ImageLayoutDepthStencilReadOnlyOptimal   common.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL
	ImageLayoutShaderReadOnlyOptimal         common.ImageLayout = C.VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL
	ImageLayoutTransferSrcOptimal            common.ImageLayout = C.VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL
	ImageLayoutTransferDstOptimal            common.ImageLayout = C.VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL
	ImageLayoutPreInitialized                common.ImageLayout = C.VK_IMAGE_LAYOUT_PREINITIALIZED

	ImageTilingOptimal common.ImageTiling = C.VK_IMAGE_TILING_OPTIMAL
	ImageTilingLinear  common.ImageTiling = C.VK_IMAGE_TILING_LINEAR

	ImageType1D common.ImageType = C.VK_IMAGE_TYPE_1D
	ImageType2D common.ImageType = C.VK_IMAGE_TYPE_2D
	ImageType3D common.ImageType = C.VK_IMAGE_TYPE_3D

	ImageUsageTransferSrc            common.ImageUsages = C.VK_IMAGE_USAGE_TRANSFER_SRC_BIT
	ImageUsageTransferDst            common.ImageUsages = C.VK_IMAGE_USAGE_TRANSFER_DST_BIT
	ImageUsageSampled                common.ImageUsages = C.VK_IMAGE_USAGE_SAMPLED_BIT
	ImageUsageStorage                common.ImageUsages = C.VK_IMAGE_USAGE_STORAGE_BIT
	ImageUsageColorAttachment        common.ImageUsages = C.VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT
	ImageUsageDepthStencilAttachment common.ImageUsages = C.VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT
	ImageUsageTransientAttachment    common.ImageUsages = C.VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT
	ImageUsageInputAttachment        common.ImageUsages = C.VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT

	Samples1  common.SampleCounts = C.VK_SAMPLE_COUNT_1_BIT
	Samples2  common.SampleCounts = C.VK_SAMPLE_COUNT_2_BIT
	Samples4  common.SampleCounts = C.VK_SAMPLE_COUNT_4_BIT
	Samples8  common.SampleCounts = C.VK_SAMPLE_COUNT_8_BIT
	Samples16 common.SampleCounts = C.VK_SAMPLE_COUNT_16_BIT
	Samples32 common.SampleCounts = C.VK_SAMPLE_COUNT_32_BIT
	Samples64 common.SampleCounts = C.VK_SAMPLE_COUNT_64_BIT
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

type ImageOptions struct {
	Flags     common.ImageCreateFlags
	ImageType common.ImageType
	Format    common.DataFormat
	Extent    common.Extent3D

	MipLevels   int
	ArrayLayers int

	Samples     common.SampleCounts
	Tiling      common.ImageTiling
	Usage       common.ImageUsages
	SharingMode common.SharingMode

	QueueFamilies []uint32

	InitialLayout common.ImageLayout

	common.HaveNext
}

func (o *ImageOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
	indexCount := len(o.QueueFamilies)
	createInfo.queueFamilyIndexCount = C.uint32_t(indexCount)
	createInfo.pQueueFamilyIndices = nil

	createInfo.extent.width = C.uint32_t(o.Extent.Width)
	createInfo.extent.height = C.uint32_t(o.Extent.Height)
	createInfo.extent.depth = C.uint32_t(o.Extent.Depth)

	if indexCount > 0 {
		queueIndicesPtr := (*C.uint32_t)(allocator.Malloc(indexCount * int(unsafe.Sizeof(C.uint32_t(0)))))
		queueIndicesSlice := ([]C.uint32_t)(unsafe.Slice(queueIndicesPtr, indexCount))

		for i := 0; i < indexCount; i++ {
			queueIndicesSlice[i] = C.uint32_t(o.QueueFamilies[i])
		}

		createInfo.pQueueFamilyIndices = queueIndicesPtr
	}

	return unsafe.Pointer(createInfo), nil
}