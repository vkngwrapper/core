package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVulkanQueue_WaitForIdle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	queue := mocks.EasyDummyQueue(t, device)

	driver.EXPECT().VkQueueWaitIdle(queue.Handle()).Return(core.VKSuccess, nil)

	_, err = queue.WaitForIdle()
	require.NoError(t, err)
}
