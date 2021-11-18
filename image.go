package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"strings"
	"unsafe"
)

type vulkanImage struct {
	driver Driver
	handle VkImage
	device VkDevice
}

func CreateImageFromHandles(handle VkImage, device VkDevice, driver Driver) Image {
	return &vulkanImage{handle: handle, device: device, driver: driver}
}

func (i *vulkanImage) Handle() VkImage {
	return i.handle
}

func (i *vulkanImage) Destroy() error {
	return i.driver.VkDestroyImage(i.device, i.handle, nil)
}

func (i *vulkanImage) MemoryRequirements() (*common.MemoryRequirements, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	memRequirementsUnsafe := arena.Malloc(C.sizeof_struct_VkMemoryRequirements)

	err := i.driver.VkGetImageMemoryRequirements(i.device, i.handle, (*VkMemoryRequirements)(memRequirementsUnsafe))
	if err != nil {
		return nil, err
	}

	memRequirements := (*C.VkMemoryRequirements)(memRequirementsUnsafe)

	return &common.MemoryRequirements{
		Size:       int(memRequirements.size),
		Alignment:  int(memRequirements.alignment),
		MemoryType: uint32(memRequirements.memoryTypeBits),
	}, nil
}

func (i *vulkanImage) BindImageMemory(memory DeviceMemory, offset int) (VkResult, error) {
	if memory == nil {
		return VKErrorUnknown, errors.New("received nil DeviceMemory")
	}
	if offset < 0 {
		return VKErrorUnknown, errors.New("received negative offset")
	}

	return i.driver.VkBindImageMemory(i.device, i.handle, memory.Handle(), VkDeviceSize(offset))
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
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder
	for i := 0; i < 32; i++ {
		checkBit := ImageFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := imageFlagsToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteRune('|')
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}

type ImageOptions struct {
	Flags  ImageFlags
	Type   common.ImageType
	Format common.DataFormat
	Extent common.Extent3D

	MipLevels   uint32
	ArrayLayers uint32

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
	createInfo.imageType = C.VkImageType(o.Type)
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
