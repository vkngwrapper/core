package core1_1_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
	"go.uber.org/mock/gomock"
)

func TestDeviceGroupCommandBufferBeginOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_1.InternalDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	coreLoader.EXPECT().VkBeginCommandBuffer(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(commandBuffer loader.VkCommandBuffer, pBeginInfo *loader.VkCommandBufferBeginInfo) (common.VkResult, error) {
		val := reflect.ValueOf(pBeginInfo).Elem()

		require.Equal(t, uint64(42), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
		require.Equal(t, uint64(1), val.FieldByName("flags").Uint())  // VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT
		require.True(t, val.FieldByName("pInheritanceInfo").IsNil())

		next := (*loader.VkDeviceGroupCommandBufferBeginInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("deviceMask").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := driver.BeginCommandBuffer(commandBuffer, core1_0.CommandBufferBeginInfo{
		Flags: core1_0.CommandBufferUsageOneTimeSubmit,
		NextOptions: common.NextOptions{Next: core1_1.DeviceGroupCommandBufferBeginInfo{
			DeviceMask: 3,
		}},
	})
	require.NoError(t, err)
}
