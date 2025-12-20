package impl1_1_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestExportSemaphoreOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_1, []string{}).(core1_1.Device)

	mockSemaphore := mocks.EasyMockSemaphore(ctrl)

	coreDriver.EXPECT().VkCreateSemaphore(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkSemaphoreCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pSemaphore *driver.VkSemaphore,
	) (common.VkResult, error) {
		*pSemaphore = mockSemaphore.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(9), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO

		next := (*driver.VkExportSemaphoreCreateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000077000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(4), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT

		return core1_0.VKSuccess, nil
	})

	semaphore, _, err := device.CreateSemaphore(nil, core1_0.SemaphoreCreateInfo{
		NextOptions: common.NextOptions{
			core1_1.ExportSemaphoreCreateInfo{
				HandleTypes: core1_1.ExternalSemaphoreHandleTypeOpaqueWin32KMT,
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, mockSemaphore.Handle(), semaphore.Handle())
}
