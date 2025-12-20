package impl1_0_test

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanDescriptorSet_Free(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	pool := mocks.EasyMockDescriptorPool(ctrl, device)

	builder := &impl1_0.DeviceObjectBuilderImpl{}
	set := builder.CreateDescriptorSetObject(mockDriver, device.Handle(), pool.Handle(), mocks.NewFakeDescriptorSet(), common.Vulkan1_0)

	mockDriver.EXPECT().VkFreeDescriptorSets(device.Handle(), pool.Handle(), driver.Uint32(1), gomock.Not(unsafe.Pointer(nil))).DoAndReturn(
		func(device driver.VkDevice, descriptorPool driver.VkDescriptorPool, descriptorSetCount driver.Uint32, pDescriptorSets *driver.VkDescriptorSet) (common.VkResult, error) {
			descriptorSetSlice := unsafe.Slice(pDescriptorSets, 1)
			require.Equal(t, set.Handle(), descriptorSetSlice[0])

			return core1_0.VKSuccess, nil
		})

	_, err := set.Free()
	require.NoError(t, err)
}
