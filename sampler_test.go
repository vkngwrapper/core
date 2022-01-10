package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestVulkanLoader1_0_CreateSampler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	samplerHandle := mocks.NewFakeSamplerHandle()

	driver.EXPECT().VkCreateSampler(mocks.Exactly(device.Handle()), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkSamplerCreateInfo, pAllocator *core.VkAllocationCallbacks, pSampler *core.VkSampler) (core.VkResult, error) {
			*pSampler = samplerHandle

			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(31), val.FieldByName("sType").Uint())
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("flags").Uint())              // VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT
			require.Equal(t, uint64(1000015000), val.FieldByName("magFilter").Uint()) // VK_FILTER_CUBIC_IMG
			require.Equal(t, uint64(1), val.FieldByName("minFilter").Uint())          // VK_FILTER_LINEAR
			require.Equal(t, uint64(1), val.FieldByName("mipmapMode").Uint())         // VK_SAMPLER_MIPMAP_MODE_LINEAR
			require.Equal(t, uint64(4), val.FieldByName("addressModeU").Uint())       // VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE
			require.Equal(t, uint64(3), val.FieldByName("addressModeV").Uint())       // VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER
			require.Equal(t, uint64(1), val.FieldByName("addressModeW").Uint())       // VK_SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT
			require.InDelta(t, 1.2, val.FieldByName("mipLodBias").Float(), 0.0001)
			require.Equal(t, uint64(1), val.FieldByName("anisotropyEnable").Uint()) // VK_TRUE
			require.InDelta(t, 4.5, val.FieldByName("maxAnisotropy").Float(), 0.0001)
			require.Equal(t, uint64(1), val.FieldByName("compareEnable").Uint()) // VK_TRUE
			require.Equal(t, uint64(4), val.FieldByName("compareOp").Uint())     // VK_COMPARE_OP_GREATER
			require.InDelta(t, 2.3, val.FieldByName("minLod").Float(), 0.0001)
			require.InDelta(t, 3.4, val.FieldByName("maxLod").Float(), 0.0001)
			require.Equal(t, uint64(2), val.FieldByName("borderColor").Uint())             // VK_BORDER_COLOR_FLOAT_OPAQUE_BLACK
			require.Equal(t, uint64(1), val.FieldByName("unnormalizedCoordinates").Uint()) // VK_TRUE

			return core.VKSuccess, nil
		})

	sampler, _, err := loader.CreateSampler(device, nil, &core.SamplerOptions{
		Flags:                   core.SamplerSubsampledEXT,
		MagFilter:               common.FilterCubic,
		MinFilter:               common.FilterLinear,
		MipmapMode:              common.MipmapLinear,
		AddressModeU:            common.AddressModeMirrorClampToEdge,
		AddressModeV:            common.AddressModeClampToBorder,
		AddressModeW:            common.AddressModeMirroredRepeat,
		MipLodBias:              1.2,
		MinLod:                  2.3,
		MaxLod:                  3.4,
		AnisotropyEnable:        true,
		MaxAnisotropy:           4.5,
		CompareEnable:           true,
		CompareOp:               common.CompareGreater,
		BorderColor:             common.BorderColorFloatOpaqueBlack,
		UnnormalizedCoordinates: true,
	})
	require.NoError(t, err)
	require.NotNil(t, sampler)
	require.Same(t, samplerHandle, sampler.Handle())
}
