package core1_0_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	internal_mocks "github.com/CannibalVox/VKng/core/internal/dummies"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestVulkanLoader1_0_CreateSampler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := internal_mocks.EasyDummyDevice(mockDriver)
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

	sampler, _, err := device.CreateSampler(nil, core1_0.SamplerCreateOptions{
		Flags:                   0,
		MagFilter:               core1_0.FilterNearest,
		MinFilter:               core1_0.FilterLinear,
		MipmapMode:              core1_0.MipmapLinear,
		AddressModeU:            core1_0.SamplerAddressModeRepeat,
		AddressModeV:            core1_0.SamplerAddressModeClampToBorder,
		AddressModeW:            core1_0.SamplerAddressModeMirroredRepeat,
		MipLodBias:              1.2,
		MinLod:                  2.3,
		MaxLod:                  3.4,
		AnisotropyEnable:        true,
		MaxAnisotropy:           4.5,
		CompareEnable:           true,
		CompareOp:               core1_0.CompareGreater,
		BorderColor:             core1_0.BorderColorFloatOpaqueBlack,
		UnnormalizedCoordinates: true,
	})
	require.NoError(t, err)
	require.NotNil(t, sampler)
	require.Equal(t, samplerHandle, sampler.Handle())
}
