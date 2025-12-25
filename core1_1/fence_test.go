package core1_1_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestExportFenceOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_1.NewDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	mockFence := mocks.NewDummyFence(device)

	coreLoader.EXPECT().VkCreateFence(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice, pCreateInfo *loader.VkFenceCreateInfo, pAllocator *loader.VkAllocationCallbacks, pFence *loader.VkFence) (common.VkResult, error) {
		*pFence = mockFence.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(8), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FENCE_CREATE_INFO
		require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_FENCE_CREATE_SIGNALED_BIT

		next := (*loader.VkExportFenceCreateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000113000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT

		return core1_0.VKSuccess, nil
	})

	fence, _, err := driver.CreateFence(
		device,
		nil,
		core1_0.FenceCreateInfo{
			Flags: core1_0.FenceCreateSignaled,

			NextOptions: common.NextOptions{
				core1_1.ExportFenceCreateInfo{
					HandleTypes: core1_1.ExternalFenceHandleTypeOpaqueWin32,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockFence.Handle(), fence.Handle())
}
