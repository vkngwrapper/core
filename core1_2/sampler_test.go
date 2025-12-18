package core1_2_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/common/extensions"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestSamplerReductionModeCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_2)
	mockSampler := mocks.EasyMockSampler(ctrl)

	coreDriver.EXPECT().VkCreateSampler(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkSamplerCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pSampler *driver.VkSampler) (common.VkResult, error) {
		*pSampler = mockSampler.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(31), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SAMPLER_CREATE_INFO

		next := (*driver.VkSamplerReductionModeCreateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000130001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("reductionMode").Uint()) // VK_SAMPLER_REDUCTION_MODE_MAX

		return core1_0.VKSuccess, nil
	})

	sampler, _, err := device.CreateSampler(
		nil,
		core1_0.SamplerCreateInfo{
			NextOptions: common.NextOptions{core1_2.SamplerReductionModeCreateInfo{
				ReductionMode: core1_2.SamplerReductionModeMax,
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockSampler.Handle(), sampler.Handle())
}
