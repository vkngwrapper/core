package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type vulkanImage struct {
	driver driver.Driver
	handle driver.VkImage
	device driver.VkDevice
}

func CreateImageFromHandles(handle driver.VkImage, device driver.VkDevice, driver driver.Driver) Image {
	return &vulkanImage{handle: handle, device: device, driver: driver}
}

func (i *vulkanImage) Handle() driver.VkImage {
	return i.handle
}

func (i *vulkanImage) Destroy(callbacks *AllocationCallbacks) {
	i.driver.VkDestroyImage(i.device, i.handle, callbacks.Handle())
}

func (i *vulkanImage) MemoryRequirements() *common.MemoryRequirements {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	memRequirementsUnsafe := arena.Malloc(C.sizeof_struct_VkMemoryRequirements)

	i.driver.VkGetImageMemoryRequirements(i.device, i.handle, (*driver.VkMemoryRequirements)(memRequirementsUnsafe))

	memRequirements := (*C.VkMemoryRequirements)(memRequirementsUnsafe)

	return &common.MemoryRequirements{
		Size:       int(memRequirements.size),
		Alignment:  int(memRequirements.alignment),
		MemoryType: uint32(memRequirements.memoryTypeBits),
	}
}

func (i *vulkanImage) BindImageMemory(memory DeviceMemory, offset int) (common.VkResult, error) {
	if memory == nil {
		return common.VKErrorUnknown, errors.New("received nil DeviceMemory")
	}
	if offset < 0 {
		return common.VKErrorUnknown, errors.New("received negative offset")
	}

	return i.driver.VkBindImageMemory(i.device, i.handle, memory.Handle(), driver.VkDeviceSize(offset))
}

func (i *vulkanImage) SubresourceLayout(subresource *common.ImageSubresource) *common.SubresourceLayout {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	subresourcePtr := (*C.VkImageSubresource)(arena.Malloc(C.sizeof_struct_VkImageSubresource))
	subresourceLayoutUnsafe := arena.Malloc(C.sizeof_struct_VkSubresourceLayout)

	subresourcePtr.aspectMask = C.VkImageAspectFlags(subresource.AspectMask)
	subresourcePtr.mipLevel = C.uint32_t(subresource.MipLevel)
	subresourcePtr.arrayLayer = C.uint32_t(subresource.ArrayLayer)

	i.driver.VkGetImageSubresourceLayout(i.device, i.handle, (*driver.VkImageSubresource)(unsafe.Pointer(subresourcePtr)), (*driver.VkSubresourceLayout)(subresourceLayoutUnsafe))

	subresourceLayout := (*C.VkSubresourceLayout)(subresourceLayoutUnsafe)
	return &common.SubresourceLayout{
		Offset:     int(subresourceLayout.offset),
		Size:       int(subresourceLayout.size),
		RowPitch:   int(subresourceLayout.rowPitch),
		ArrayPitch: int(subresourceLayout.arrayPitch),
		DepthPitch: int(subresourceLayout.depthPitch),
	}
}

type ImageFlags int32

const (
	ImageSparseBinding                     ImageFlags = C.VK_IMAGE_CREATE_SPARSE_BINDING_BIT
	ImageSparseAliased                     ImageFlags = C.VK_IMAGE_CREATE_SPARSE_ALIASED_BIT
	ImageMutableFormat                     ImageFlags = C.VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT
	ImageCubeCompatible                    ImageFlags = C.VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT
	ImageAlias                             ImageFlags = C.VK_IMAGE_CREATE_ALIAS_BIT
	ImageSplitInstanceBindRegions          ImageFlags = C.VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT
	Image2DArrayCompatible                 ImageFlags = C.VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT
	ImageBlockTexelViewCompatible          ImageFlags = C.VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT
	ImageExtendedUsage                     ImageFlags = C.VK_IMAGE_CREATE_EXTENDED_USAGE_BIT
	ImageProtected                         ImageFlags = C.VK_IMAGE_CREATE_PROTECTED_BIT
	ImageDisjoint                          ImageFlags = C.VK_IMAGE_CREATE_DISJOINT_BIT
	ImageCornerSampledNV                   ImageFlags = C.VK_IMAGE_CREATE_CORNER_SAMPLED_BIT_NV
	ImageSampleLocationsCompatibleDepthEXT ImageFlags = C.VK_IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT
	ImageSubsampledEXT                     ImageFlags = C.VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT
)

var imageFlagsToString = map[ImageFlags]string{
	ImageSparseBinding:                     "Sparse Binding",
	ImageSparseAliased:                     "Sparse Aliased",
	ImageMutableFormat:                     "Mutable Format",
	ImageCubeCompatible:                    "Cube Compatible",
	ImageAlias:                             "Alias",
	ImageSplitInstanceBindRegions:          "Split Instance Bind Regions",
	Image2DArrayCompatible:                 "2D Array Compatible",
	ImageBlockTexelViewCompatible:          "Block Texel View Compatible",
	ImageExtendedUsage:                     "Extended Usage",
	ImageProtected:                         "Protected",
	ImageDisjoint:                          "Disjoint",
	ImageCornerSampledNV:                   "Corner Sampled (Nvidia Extension)",
	ImageSampleLocationsCompatibleDepthEXT: "Sample Locations Compatible Depth (Extension)",
	ImageSubsampledEXT:                     "Subsampled (Extension)",
}

func (f ImageFlags) String() string {
	return common.FlagsToString(f, imageFlagsToString)
}

type ImageOptions struct {
	Flags     ImageFlags
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

func (o *ImageOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkImageCreateInfo)(allocator.Malloc(int(unsafe.Sizeof([1]C.VkImageCreateInfo{}))))

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
