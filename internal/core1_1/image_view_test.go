package internal1_1_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestImageViewUsageOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, coreDriver)
	image := mocks.EasyMockImage(ctrl)
	expectedImageView := mocks.EasyMockImageView(ctrl)

	coreDriver.EXPECT().VkCreateImageView(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pCreateInfo *driver.VkImageViewCreateInfo, pAllocator *driver.VkAllocationCallbacks, pImageView *driver.VkImageView) (common.VkResult, error) {
			*pImageView = expectedImageView.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(15), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO
			require.Equal(t, image.Handle(), (driver.VkImage)(val.FieldByName("image").UnsafePointer()))

			viewUsagePtr := (*driver.VkImageViewUsageCreateInfo)(val.FieldByName("pNext").UnsafePointer())
			viewUsage := reflect.ValueOf(viewUsagePtr).Elem()
			require.Equal(t, uint64(1000117002), viewUsage.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO
			require.True(t, viewUsage.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000080), viewUsage.FieldByName("usage").Uint()) // VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT

			return core1_0.VKSuccess, nil
		})

	imageView, _, err := loader.CreateImageView(device, nil, core1_0.ImageViewCreateOptions{
		Image: image,
		HaveNext: common.HaveNext{Next: core1_1.ImageViewUsageOptions{
			Usage: core1_0.ImageUsageInputAttachment,
		}},
	})

	require.NoError(t, err)
	require.Equal(t, expectedImageView.Handle(), imageView.Handle())
}
