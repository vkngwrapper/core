package internal1_1_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/internal/dummies"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanPhysicalDevice_PhysicalDeviceExternalFenceProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)

	physicalDevice := dummies.EasyDummyPhysicalDevice(t, loader)

	coreDriver.EXPECT().VkGetPhysicalDeviceExternalFenceProperties(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pExternalFenceInfo *driver.VkPhysicalDeviceExternalFenceInfo,
		pExternalFenceProperties *driver.VkExternalFenceProperties,
	) {
		val := reflect.ValueOf(pExternalFenceInfo).Elem()
		require.Equal(t, uint64(1000112000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(4), val.FieldByName("handleType").Uint()) // VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT

		val = reflect.ValueOf(pExternalFenceProperties).Elem()
		*(*uint32)(unsafe.Pointer(val.FieldByName("exportFromImportedHandleTypes").UnsafeAddr())) = uint32(8) // VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT
		*(*uint32)(unsafe.Pointer(val.FieldByName("compatibleHandleTypes").UnsafeAddr())) = uint32(4)         // VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalFenceFeatures").UnsafeAddr())) = uint32(1)         // VK_EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT
	})

	var outData core1_1.ExternalFenceOutData
	err = physicalDevice.Core1_1Instance().ExternalFenceProperties(
		core1_1.ExternalFenceOptions{
			HandleType: core1_1.ExternalFenceHandleTypeOpaqueWin32KMT,
		},
		&outData,
	)
	require.NoError(t, err)
	require.Equal(t, core1_1.ExternalFenceOutData{
		ExportFromImportedHandleTypes: core1_1.ExternalFenceHandleTypeSyncFD,
		CompatibleHandleTypes:         core1_1.ExternalFenceHandleTypeOpaqueWin32KMT,
		ExternalFenceFeatures:         core1_1.ExternalFenceFeatureExportable,
	}, outData)
}

func TestExportFenceOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	mockFence := mocks.EasyMockFence(ctrl)

	coreDriver.EXPECT().VkCreateFence(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice, pCreateInfo *driver.VkFenceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pFence *driver.VkFence) (common.VkResult, error) {
		*pFence = mockFence.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(8), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FENCE_CREATE_INFO
		require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_FENCE_CREATE_SIGNALED_BIT

		next := (*driver.VkExportFenceCreateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000113000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT

		return core1_0.VKSuccess, nil
	})

	fence, _, err := loader.CreateFence(
		device,
		nil,
		core1_0.FenceCreateOptions{
			Flags: core1_0.FenceCreateSignaled,

			HaveNext: common.HaveNext{
				core1_1.ExportFenceOptions{
					HandleTypes: core1_1.ExternalFenceHandleTypeOpaqueWin32,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockFence.Handle(), fence.Handle())
}
