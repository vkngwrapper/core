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

func TestVulkanLoader1_0_CreateImageView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	imageViewHandle := mocks.NewFakeImageViewHandle()
	image := mocks.EasyMockImage(ctrl)

	driver.EXPECT().VkCreateImageView(mocks.Exactly(device.Handle()), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkImageViewCreateInfo, pAllocator *core.VkAllocationCallbacks, pImageView *core.VkImageView) (core.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(15), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), val.FieldByName("flags").Uint()) // VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DEFERRED_BIT_EXT
			require.Same(t, image.Handle(), (core.VkImage)(unsafe.Pointer(val.FieldByName("image").Elem().UnsafeAddr())))
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
			return core.VKSuccess, nil
		})

	imageView, _, err := loader.CreateImageView(device, nil, &core.ImageViewOptions{
		Image:    image,
		ViewType: common.ViewType2D,
		Format:   common.FormatA2B10G10R10SignedScaled,
		Flags:    core.ImageViewCreateFragmentDensityMapDeferredEXT,
		Components: common.ComponentMapping{
			A: common.SwizzleAlpha,
			R: common.SwizzleRed,
			G: common.SwizzleGreen,
			B: common.SwizzleBlue,
		},
		SubresourceRange: common.ImageSubresourceRange{
			BaseMipLevel:   1,
			LevelCount:     2,
			BaseArrayLayer: 3,
			LayerCount:     5,
			AspectMask:     common.AspectColor | common.AspectDepth,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, imageView)
	require.Same(t, imageViewHandle, imageView.Handle())
}
