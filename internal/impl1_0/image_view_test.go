package impl1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreateImageView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	imageViewHandle := mocks.NewFakeImageViewHandle()
	image := mocks1_0.EasyMockImage(ctrl)

	mockDriver.EXPECT().VkCreateImageView(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkImageViewCreateInfo, pAllocator *loader.VkAllocationCallbacks, pImageView *loader.VkImageView) (common.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(15), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, image.Handle(), (loader.VkImage)(unsafe.Pointer(val.FieldByName("image").Elem().UnsafeAddr())))
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

	imageView, _, err := device.CreateImageView(nil, core1_0.ImageViewCreateInfo{
		Image:    image,
		ViewType: core1_0.ImageViewType2D,
		Format:   core1_0.FormatA2B10G10R10SignedScaledPacked,
		Flags:    0,
		Components: core1_0.ComponentMapping{
			A: core1_0.ComponentSwizzleAlpha,
			R: core1_0.ComponentSwizzleRed,
			G: core1_0.ComponentSwizzleGreen,
			B: core1_0.ComponentSwizzleBlue,
		},
		SubresourceRange: core1_0.ImageSubresourceRange{
			BaseMipLevel:   1,
			LevelCount:     2,
			BaseArrayLayer: 3,
			LayerCount:     5,
			AspectMask:     core1_0.ImageAspectColor | core1_0.ImageAspectDepth,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, imageView)
	require.Equal(t, imageViewHandle, imageView.Handle())
}
