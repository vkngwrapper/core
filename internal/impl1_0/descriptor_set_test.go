package impl1_0_test

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestVulkanDescriptorSet_Free(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)

	device := mocks1_0.EasyMockDevice(ctrl, mockDriver)
	pool := mocks1_0.EasyMockDescriptorPool(ctrl, device)

	builder := &impl1_0.DeviceObjectBuilderImpl{}
	set := builder.CreateDescriptorSetObject(mockDriver, device.Handle(), pool.Handle(), mocks.NewFakeDescriptorSet(), common.Vulkan1_0)

	mockDriver.EXPECT().VkFreeDescriptorSets(device.Handle(), pool.Handle(), loader.Uint32(1), gomock.Not(unsafe.Pointer(nil))).DoAndReturn(
		func(device loader.VkDevice, descriptorPool loader.VkDescriptorPool, descriptorSetCount loader.Uint32, pDescriptorSets *loader.VkDescriptorSet) (common.VkResult, error) {
			descriptorSetSlice := unsafe.Slice(pDescriptorSets, 1)
			require.Equal(t, set.Handle(), descriptorSetSlice[0])

			return core1_0.VKSuccess, nil
		})

	_, err := set.Free()
	require.NoError(t, err)
}
