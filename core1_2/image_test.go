package core1_2_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/internal/impl1_2"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_2"
	"go.uber.org/mock/gomock"
)

func TestImageStencilUsageCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_2.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	mockImage := mocks1_2.EasyMockImage(ctrl)

	coreDriver.EXPECT().VkCreateImage(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pCreateInfo *loader.VkImageCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pImage *loader.VkImage) (common.VkResult, error) {

		*pImage = mockImage.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(14), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO

		next := (*loader.VkImageStencilUsageCreateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000246000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0x10), val.FieldByName("stencilUsage").Uint()) // VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT

		return core1_0.VKSuccess, nil
	})

	image, _, err := device.CreateImage(
		nil,
		core1_0.ImageCreateInfo{
			NextOptions: common.NextOptions{core1_2.ImageStencilUsageCreateInfo{
				StencilUsage: core1_0.ImageUsageColorAttachment,
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockImage.Handle(), image.Handle())
}

func TestImageFormatListCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_2)
	builder := &impl1_2.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_2, []string{})

	mockImage := mocks1_2.EasyMockImage(ctrl)

	coreDriver.EXPECT().VkCreateImage(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pCreateInfo *loader.VkImageCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pImage *loader.VkImage) (common.VkResult, error) {

		*pImage = mockImage.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(14), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO

		next := (*loader.VkImageFormatListCreateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000147000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("viewFormatCount").Uint())

		formatPtr := (*loader.VkFormat)(val.FieldByName("pViewFormats").UnsafePointer())
		formatSlice := ([]loader.VkFormat)(unsafe.Slice(formatPtr, 3))
		require.Equal(t, []loader.VkFormat{64, 57, 52}, formatSlice)

		return core1_0.VKSuccess, nil
	})

	image, _, err := device.CreateImage(
		nil,
		core1_0.ImageCreateInfo{
			NextOptions: common.NextOptions{
				core1_2.ImageFormatListCreateInfo{
					ViewFormats: []core1_0.Format{
						core1_0.FormatA2B10G10R10UnsignedNormalizedPacked,
						core1_0.FormatA8B8G8R8SRGBPacked,
						core1_0.FormatA8B8G8R8SignedNormalizedPacked,
					},
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockImage.Handle(), image.Handle())
}
