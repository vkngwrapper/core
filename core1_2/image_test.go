package core1_2_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/common/extensions"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestImageStencilUsageCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)
	mockImage := mocks.EasyMockImage(ctrl)

	coreDriver.EXPECT().VkCreateImage(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkImageCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pImage *driver.VkImage) (common.VkResult, error) {

		*pImage = mockImage.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(14), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO

		next := (*driver.VkImageStencilUsageCreateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000246000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0x10), val.FieldByName("stencilUsage").Uint()) // VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT

		return core1_0.VKSuccess, nil
	})

	image, _, err := device.CreateImage(
		nil,
		core1_0.ImageCreateOptions{
			NextOptions: common.NextOptions{core1_2.ImageStencilUsageCreateOptions{
				StencilUsage: core1_0.ImageUsageColorAttachment,
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockImage.Handle(), image.Handle())
}

func TestImageFormatListCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_2)
	mockImage := mocks.EasyMockImage(ctrl)

	coreDriver.EXPECT().VkCreateImage(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkImageCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pImage *driver.VkImage) (common.VkResult, error) {

		*pImage = mockImage.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(14), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO

		next := (*driver.VkImageFormatListCreateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000147000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("viewFormatCount").Uint())

		formatPtr := (*driver.VkFormat)(val.FieldByName("pViewFormats").UnsafePointer())
		formatSlice := ([]driver.VkFormat)(unsafe.Slice(formatPtr, 3))
		require.Equal(t, []driver.VkFormat{64, 57, 52}, formatSlice)

		return core1_0.VKSuccess, nil
	})

	image, _, err := device.CreateImage(
		nil,
		core1_0.ImageCreateOptions{
			NextOptions: common.NextOptions{
				core1_2.ImageFormatListCreateOptions{
					ViewFormats: []core1_0.DataFormat{
						core1_0.DataFormatA2B10G10R10UnsignedNormalizedPacked,
						core1_0.DataFormatA8B8G8R8SRGBPacked,
						core1_0.DataFormatA8B8G8R8SignedNormalizedPacked,
					},
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockImage.Handle(), image.Handle())
}
