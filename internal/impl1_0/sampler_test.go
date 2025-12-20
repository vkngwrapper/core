package impl1_0_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreateSampler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})

	samplerHandle := mocks.NewFakeSamplerHandle()

	mockDriver.EXPECT().VkCreateSampler(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkSamplerCreateInfo, pAllocator *driver.VkAllocationCallbacks, pSampler *driver.VkSampler) (common.VkResult, error) {
			*pSampler = samplerHandle

			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(31), val.FieldByName("sType").Uint())
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(0), val.FieldByName("magFilter").Uint())    // VK_FILTER_NEAREST
			require.Equal(t, uint64(1), val.FieldByName("minFilter").Uint())    // VK_FILTER_LINEAR
			require.Equal(t, uint64(1), val.FieldByName("mipmapMode").Uint())   // VK_SAMPLER_MIPMAP_MODE_LINEAR
			require.Equal(t, uint64(0), val.FieldByName("addressModeU").Uint()) // VK_SAMPLER_ADDRESS_MODE_REPEAT
			require.Equal(t, uint64(3), val.FieldByName("addressModeV").Uint()) // VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER
			require.Equal(t, uint64(1), val.FieldByName("addressModeW").Uint()) // VK_SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT
			require.InDelta(t, 1.2, val.FieldByName("mipLodBias").Float(), 0.0001)
			require.Equal(t, uint64(1), val.FieldByName("anisotropyEnable").Uint()) // VK_TRUE
			require.InDelta(t, 4.5, val.FieldByName("maxAnisotropy").Float(), 0.0001)
			require.Equal(t, uint64(1), val.FieldByName("compareEnable").Uint()) // VK_TRUE
			require.Equal(t, uint64(4), val.FieldByName("compareOp").Uint())     // VK_COMPARE_OP_GREATER
			require.InDelta(t, 2.3, val.FieldByName("minLod").Float(), 0.0001)
			require.InDelta(t, 3.4, val.FieldByName("maxLod").Float(), 0.0001)
			require.Equal(t, uint64(2), val.FieldByName("borderColor").Uint())             // VK_BORDER_COLOR_FLOAT_OPAQUE_BLACK
			require.Equal(t, uint64(1), val.FieldByName("unnormalizedCoordinates").Uint()) // VK_TRUE

			return core1_0.VKSuccess, nil
		})

	sampler, _, err := device.CreateSampler(nil, core1_0.SamplerCreateInfo{
		Flags:                   0,
		MagFilter:               core1_0.FilterNearest,
		MinFilter:               core1_0.FilterLinear,
		MipmapMode:              core1_0.SamplerMipmapModeLinear,
		AddressModeU:            core1_0.SamplerAddressModeRepeat,
		AddressModeV:            core1_0.SamplerAddressModeClampToBorder,
		AddressModeW:            core1_0.SamplerAddressModeMirroredRepeat,
		MipLodBias:              1.2,
		MinLod:                  2.3,
		MaxLod:                  3.4,
		AnisotropyEnable:        true,
		MaxAnisotropy:           4.5,
		CompareEnable:           true,
		CompareOp:               core1_0.CompareOpGreater,
		BorderColor:             core1_0.BorderColorFloatOpaqueBlack,
		UnnormalizedCoordinates: true,
	})
	require.NoError(t, err)
	require.NotNil(t, sampler)
	require.Equal(t, samplerHandle, sampler.Handle())
}
