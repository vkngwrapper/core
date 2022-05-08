package internal1_0_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	internal_mocks "github.com/CannibalVox/VKng/core/internal/dummies"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanLoader1_0_CreateImageView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := internal_mocks.EasyDummyDevice(t, ctrl, loader)
	imageViewHandle := mocks.NewFakeImageViewHandle()
	image := mocks.EasyMockImage(ctrl)

	mockDriver.EXPECT().VkCreateImageView(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkImageViewCreateInfo, pAllocator *driver.VkAllocationCallbacks, pImageView *driver.VkImageView) (common.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(15), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, image.Handle(), (driver.VkImage)(unsafe.Pointer(val.FieldByName("image").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), val.FieldByName("viewType").Uint()) // VK_IMAGE_VIEW_TYPE_2D
			require.Equal(t, uint64(67), val.FieldByName("format").Uint())  // VK_FORMAT_A2B10G10R10_SSCALED_PACK32

			components := val.FieldByName("components")
			require.Equal(t, uint64(3), components.FieldByName("r").Uint()) // VK_COMPONENT_SWIZZLE_R
			require.Equal(t, uint64(4), components.FieldByName("g").Uint()) // VK_COMPONENT_SWIZZLE_G
			require.Equal(t, uint64(5), components.FieldByName("b").Uint()) // VK_COMPONENT_SWIZZLE_B
			require.Equal(t, uint64(6), components.FieldByName("a").Uint()) // VK_COMPONENT_SWIZZLE_A

			subresource := val.FieldByName("subresourceRange")
			require.Equal(t, uint64(1), subresource.FieldByName("baseMipLevel").Uint())
			require.Equal(t, uint64(2), subresource.FieldByName("levelCount").Uint())
			require.Equal(t, uint64(3), subresource.FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(5), subresource.FieldByName("layerCount").Uint())
			require.Equal(t, uint64(3), subresource.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_COLOR_BIT | VK_IMAGE_ASPECT_DEPTH_BIT

			*pImageView = imageViewHandle
			return core1_0.VKSuccess, nil
		})

	imageView, _, err := loader.CreateImageView(device, nil, core1_0.ImageViewCreateOptions{
		Image:    image,
		ViewType: core1_0.ViewType2D,
		Format:   core1_0.DataFormatA2B10G10R10SignedScaledPacked,
		Flags:    0,
		Components: core1_0.ComponentMapping{
			A: core1_0.SwizzleAlpha,
			R: core1_0.SwizzleRed,
			G: core1_0.SwizzleGreen,
			B: core1_0.SwizzleBlue,
		},
		SubresourceRange: common.ImageSubresourceRange{
			BaseMipLevel:   1,
			LevelCount:     2,
			BaseArrayLayer: 3,
			LayerCount:     5,
			AspectMask:     core1_0.AspectColor | core1_0.AspectDepth,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, imageView)
	require.Equal(t, imageViewHandle, imageView.Handle())
}
