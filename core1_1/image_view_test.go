package core1_1_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	"github.com/vkngwrapper/core/v2/internal/dummies"
	"github.com/vkngwrapper/core/v2/mocks"
	"go.uber.org/mock/gomock"
)

func TestImageViewUsageOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)

	device := dummies.EasyDummyDevice(coreDriver)
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

	imageView, _, err := device.CreateImageView(nil, core1_0.ImageViewCreateInfo{
		Image: image,
		NextOptions: common.NextOptions{Next: core1_1.ImageViewUsageCreateInfo{
			Usage: core1_0.ImageUsageInputAttachment,
		}},
	})

	require.NoError(t, err)
	require.Equal(t, expectedImageView.Handle(), imageView.Handle())
}
