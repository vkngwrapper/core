package core1_0_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanLoader1_0_CreateImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	imageHandle := mocks.NewFakeImageHandle()

	mockDriver.EXPECT().VkCreateImage(mocks.Exactly(device.Handle()), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkImageCreateInfo, pAllocator *driver.VkAllocationCallbacks, pImage *driver.VkImage) (common.VkResult, error) {
			*pImage = imageHandle

			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(14), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000800), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), val.FieldByName("imageType").Uint()) // VK_IMAGE_TYPE_2D
			require.Equal(t, uint64(59), val.FieldByName("format").Uint())   // VK_FORMAT_A2R10G10B10_SNORM_PACK32
			require.Equal(t, uint64(7), val.FieldByName("mipLevels").Uint())
			require.Equal(t, uint64(11), val.FieldByName("arrayLayers").Uint())
			require.Equal(t, uint64(16), val.FieldByName("samples").Uint())    // VK_SAMPLE_COUNT_16_BIT
			require.Equal(t, uint64(1), val.FieldByName("tiling").Uint())      // VK_IMAGE_TILING_LINEAR
			require.Equal(t, uint64(1), val.FieldByName("usage").Uint())       // VK_IMAGE_USAGE_TRANSFER_SRC_BIT
			require.Equal(t, uint64(1), val.FieldByName("sharingMode").Uint()) // VK_SHARING_MODE_CONCURRENT
			require.Equal(t, uint64(3), val.FieldByName("queueFamilyIndexCount").Uint())
			require.Equal(t, uint64(1000117001), val.FieldByName("initialLayout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL

			extent := val.FieldByName("extent")
			require.Equal(t, uint64(1), extent.FieldByName("width").Uint())
			require.Equal(t, uint64(3), extent.FieldByName("height").Uint())
			require.Equal(t, uint64(5), extent.FieldByName("depth").Uint())

			indicesPtr := (*driver.Uint32)(unsafe.Pointer(val.FieldByName("pQueueFamilyIndices").Elem().UnsafeAddr()))
			indicesSlice := ([]driver.Uint32)(unsafe.Slice(indicesPtr, 3))

			require.Equal(t, []driver.Uint32{13, 17, 19}, indicesSlice)

			return common.VKSuccess, nil
		})

	image, _, err := loader.CreateImage(device, nil, &core.ImageOptions{
		Flags:     core.ImageProtected,
		ImageType: common.ImageType2D,
		Format:    common.FormatA2R10G10B10SignedNormalized,
		Extent: common.Extent3D{
			Width:  1,
			Height: 3,
			Depth:  5,
		},
		MipLevels:     7,
		ArrayLayers:   11,
		Samples:       common.Samples16,
		Tiling:        common.ImageTilingLinear,
		Usage:         common.ImageUsageTransferSrc,
		SharingMode:   common.SharingConcurrent,
		QueueFamilies: []uint32{13, 17, 19},
		InitialLayout: common.LayoutDepthAttachmentStencilReadOnlyOptimal,
	})
	require.NoError(t, err)
	require.NotNil(t, image)
	require.Same(t, imageHandle, image.Handle())
}

func TestVulkanImage_MemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	image := mocks.EasyDummyImage(t, loader, device)

	mockDriver.EXPECT().VkGetImageMemoryRequirements(mocks.Exactly(device.Handle()), mocks.Exactly(image.Handle()), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, image driver.VkImage, pRequirements *driver.VkMemoryRequirements) {
			val := reflect.ValueOf(pRequirements).Elem()

			*(*driver.VkDeviceSize)(unsafe.Pointer(val.FieldByName("size").UnsafeAddr())) = 1
			*(*driver.VkDeviceSize)(unsafe.Pointer(val.FieldByName("alignment").UnsafeAddr())) = 3
			*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("memoryTypeBits").UnsafeAddr())) = 5
		})

	reqs := image.MemoryRequirements()
	require.NotNil(t, reqs)
	require.Equal(t, 1, reqs.Size)
	require.Equal(t, 3, reqs.Alignment)
	require.Equal(t, uint32(5), reqs.MemoryType)
}

func TestVulkanImage_BindImageMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	memory := mocks.EasyMockDeviceMemory(ctrl)
	image := mocks.EasyDummyImage(t, loader, device)

	mockDriver.EXPECT().VkBindImageMemory(mocks.Exactly(device.Handle()), mocks.Exactly(image.Handle()), memory.Handle(), driver.VkDeviceSize(3)).Return(common.VKSuccess, nil)

	_, err = image.BindImageMemory(memory, 3)
	require.NoError(t, err)
}

func TestVulkanImage_SubresourceLayout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	image := mocks.EasyDummyImage(t, loader, device)

	mockDriver.EXPECT().VkGetImageSubresourceLayout(mocks.Exactly(device.Handle()), mocks.Exactly(image.Handle()), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, image driver.VkImage, pSubresource *driver.VkImageSubresource, pLayout *driver.VkSubresourceLayout) {
			val := reflect.ValueOf(pSubresource).Elem()

			require.Equal(t, uint64(0x00000200), val.FieldByName("aspectMask").Uint())
			require.Equal(t, uint64(1), val.FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(3), val.FieldByName("arrayLayer").Uint())

			val = reflect.ValueOf(pLayout).Elem()
			*(*uint64)(unsafe.Pointer(val.FieldByName("offset").UnsafeAddr())) = 5
			*(*uint64)(unsafe.Pointer(val.FieldByName("size").UnsafeAddr())) = 7
			*(*uint64)(unsafe.Pointer(val.FieldByName("rowPitch").UnsafeAddr())) = 11
			*(*uint64)(unsafe.Pointer(val.FieldByName("depthPitch").UnsafeAddr())) = 13
			*(*uint64)(unsafe.Pointer(val.FieldByName("arrayPitch").UnsafeAddr())) = 17
		})

	layout := image.SubresourceLayout(&common.ImageSubresource{
		AspectMask: common.AspectMemoryPlane2EXT,
		MipLevel:   1,
		ArrayLayer: 3,
	})
	require.NotNil(t, layout)
	require.Equal(t, 5, layout.Offset)
	require.Equal(t, 7, layout.Size)
	require.Equal(t, 11, layout.RowPitch)
	require.Equal(t, 13, layout.DepthPitch)
	require.Equal(t, 17, layout.ArrayPitch)
}

func TestVulkanImage_SparseMemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	image := mocks.EasyDummyImage(t, loader, device)

	mockDriver.EXPECT().VkGetImageSparseMemoryRequirements(mocks.Exactly(device.Handle()), mocks.Exactly(image.Handle()), gomock.Not(nil), nil).DoAndReturn(
		func(device driver.VkDevice, image driver.VkImage, pReqCount *driver.Uint32, pRequirements *driver.VkSparseImageMemoryRequirements) {
			*pReqCount = 2
		})

	mockDriver.EXPECT().VkGetImageSparseMemoryRequirements(mocks.Exactly(device.Handle()), mocks.Exactly(image.Handle()), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, image driver.VkImage, pReqCount *driver.Uint32, pRequirements *driver.VkSparseImageMemoryRequirements) {
			require.Equal(t, driver.Uint32(2), *pReqCount)

			requirementSlice := unsafe.Slice(pRequirements, 2)
			reqVal := reflect.ValueOf(requirementSlice)

			req := reqVal.Index(0)

			*(*uint32)(unsafe.Pointer(req.FieldByName("formatProperties").FieldByName("aspectMask").UnsafeAddr())) = uint32(1) // VK_IMAGE_ASPECT_COLOR_BIT
			*(*int32)(unsafe.Pointer(req.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("width").UnsafeAddr())) = int32(1)
			*(*int32)(unsafe.Pointer(req.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("height").UnsafeAddr())) = int32(3)
			*(*int32)(unsafe.Pointer(req.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("depth").UnsafeAddr())) = int32(5)
			*(*uint32)(unsafe.Pointer(req.FieldByName("formatProperties").FieldByName("flags").UnsafeAddr())) = uint32(4) // VK_SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT
			*(*uint32)(unsafe.Pointer(req.FieldByName("imageMipTailFirstLod").UnsafeAddr())) = uint32(7)
			*(*uint64)(unsafe.Pointer(req.FieldByName("imageMipTailOffset").UnsafeAddr())) = uint64(11)
			*(*uint64)(unsafe.Pointer(req.FieldByName("imageMipTailSize").UnsafeAddr())) = uint64(13)
			*(*uint64)(unsafe.Pointer(req.FieldByName("imageMipTailStride").UnsafeAddr())) = uint64(17)

			req = reqVal.Index(1)

			*(*uint32)(unsafe.Pointer(req.FieldByName("formatProperties").FieldByName("aspectMask").UnsafeAddr())) = uint32(2) // VK_IMAGE_ASPECT_DEPTH_BIT
			*(*int32)(unsafe.Pointer(req.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("width").UnsafeAddr())) = int32(19)
			*(*int32)(unsafe.Pointer(req.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("height").UnsafeAddr())) = int32(23)
			*(*int32)(unsafe.Pointer(req.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("depth").UnsafeAddr())) = int32(29)
			*(*uint32)(unsafe.Pointer(req.FieldByName("formatProperties").FieldByName("flags").UnsafeAddr())) = uint32(2) // VK_SPARSE_IMAGE_FORMAT_ALIGNED_MIP_SIZE_BIT
			*(*uint32)(unsafe.Pointer(req.FieldByName("imageMipTailFirstLod").UnsafeAddr())) = uint32(31)
			*(*uint64)(unsafe.Pointer(req.FieldByName("imageMipTailOffset").UnsafeAddr())) = uint64(37)
			*(*uint64)(unsafe.Pointer(req.FieldByName("imageMipTailSize").UnsafeAddr())) = uint64(41)
			*(*uint64)(unsafe.Pointer(req.FieldByName("imageMipTailStride").UnsafeAddr())) = uint64(43)
		})

	reqs := image.SparseMemoryRequirements()
	require.Equal(t, []core.SparseImageMemoryRequirements{
		{
			FormatProperties: core.SparseImageFormatProperties{
				AspectMask:       common.AspectColor,
				ImageGranularity: common.Extent3D{1, 3, 5},
				Flags:            core.SparseImageFormatNonstandardBlockSize,
			},
			ImageMipTailFirstLod: 7,
			ImageMipTailOffset:   11,
			ImageMipTailSize:     13,
			ImageMipTailStride:   17,
		},
		{
			FormatProperties: core.SparseImageFormatProperties{
				AspectMask:       common.AspectDepth,
				ImageGranularity: common.Extent3D{19, 23, 29},
				Flags:            core.SparseImageFormatAlignedMipSize,
			},
			ImageMipTailFirstLod: 31,
			ImageMipTailOffset:   37,
			ImageMipTailSize:     41,
			ImageMipTailStride:   43,
		},
	}, reqs)
}
