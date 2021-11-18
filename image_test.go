package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
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

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	imageHandle := mocks.NewFakeImageHandle()

	driver.EXPECT().VkCreateImage(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkImageCreateInfo, pAllocator *core.VkAllocationCallbacks, pImage *core.VkImage) (core.VkResult, error) {
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

			indicesPtr := (*core.Uint32)(unsafe.Pointer(val.FieldByName("pQueueFamilyIndices").Elem().UnsafeAddr()))
			indicesSlice := ([]core.Uint32)(unsafe.Slice(indicesPtr, 3))

			require.Equal(t, []core.Uint32{13, 17, 19}, indicesSlice)

			return core.VKSuccess, nil
		})

	image, _, err := loader.CreateImage(device, &core.ImageOptions{
		Flags:  core.ImageProtected,
		Type:   common.ImageType2D,
		Format: common.FormatA2R10G10B10SignedNormalized,
		Extent: common.Extent3D{
			Width:  1,
			Height: 3,
			Depth:  5,
		},
		MipLevels:     7,
		ArrayLayers:   11,
		Samples:       common.Samples16,
		Tiling:        common.ImageTilingLinear,
		Usage:         common.ImageTransferSrc,
		SharingMode:   common.SharingConcurrent,
		QueueFamilies: []uint32{13, 17, 19},
		InitialLayout: common.LayoutDepthAttachmentStencilReadOnlyOptimal,
	})
	require.NoError(t, err)
	require.NotNil(t, image)
	require.Equal(t, imageHandle, image.Handle())
}

func TestVulkanImage_MemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	image := mocks.EasyDummyImage(t, loader, device)

	driver.EXPECT().VkGetImageMemoryRequirements(device.Handle(), image.Handle(), gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, image core.VkImage, pRequirements *core.VkMemoryRequirements) error {
			val := reflect.ValueOf(pRequirements).Elem()

			*(*core.VkDeviceSize)(unsafe.Pointer(val.FieldByName("size").UnsafeAddr())) = 1
			*(*core.VkDeviceSize)(unsafe.Pointer(val.FieldByName("alignment").UnsafeAddr())) = 3
			*(*core.Uint32)(unsafe.Pointer(val.FieldByName("memoryTypeBits").UnsafeAddr())) = 5

			return nil
		})

	reqs, err := image.MemoryRequirements()
	require.NoError(t, err)
	require.NotNil(t, reqs)
	require.Equal(t, 1, reqs.Size)
	require.Equal(t, 3, reqs.Alignment)
	require.Equal(t, uint32(5), reqs.MemoryType)
}

func TestVulkanImage_BindImageMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	memory := mocks.EasyMockDeviceMemory(ctrl)
	image := mocks.EasyDummyImage(t, loader, device)

	driver.EXPECT().VkBindImageMemory(device.Handle(), image.Handle(), memory.Handle(), core.VkDeviceSize(3)).Return(core.VKSuccess, nil)

	_, err = image.BindImageMemory(memory, 3)
	require.NoError(t, err)
}
