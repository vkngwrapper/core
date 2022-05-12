package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestVulkanLoader_CreateSamplerYcbcrConversion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	mockYcbcr := mocks.EasyMockSamplerYcbcrConversion(ctrl)

	coreDriver.EXPECT().VkCreateSamplerYcbcrConversion(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice,
			pCreateInfo *driver.VkSamplerYcbcrConversionCreateInfo,
			pAllocator *driver.VkAllocationCallbacks,
			pYcbcrConversion *driver.VkSamplerYcbcrConversion,
		) (common.VkResult, error) {
			*pYcbcrConversion = mockYcbcr.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(1000156000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1000156021), val.FieldByName("format").Uint())             // VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16_KHR
			require.Equal(t, uint64(2), val.FieldByName("ycbcrModel").Uint())                  // VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709_KHR
			require.Equal(t, uint64(1), val.FieldByName("ycbcrRange").Uint())                  // VK_SAMPLER_YCBCR_RANGE_ITU_NARROW_KHR
			require.Equal(t, uint64(4), val.FieldByName("components").FieldByName("r").Uint()) // VK_COMPONENT_SWIZZLE_G
			require.Equal(t, uint64(6), val.FieldByName("components").FieldByName("g").Uint()) // VK_COMPONENT_SWIZZLE_A
			require.Equal(t, uint64(0), val.FieldByName("components").FieldByName("b").Uint()) // VK_COMPONENT_SWIZZLE_IDENTITY
			require.Equal(t, uint64(2), val.FieldByName("components").FieldByName("a").Uint()) // VK_COMPONENT_SWIZZLE_ONE
			require.Equal(t, uint64(0), val.FieldByName("yChromaOffset").Uint())               // VK_CHROMA_LOCATION_COSITED_EVEN_KHR
			require.Equal(t, uint64(1), val.FieldByName("xChromaOffset").Uint())               // VK_CHROMA_LOCATION_MIDPOINT_KHR
			require.Equal(t, uint64(1), val.FieldByName("forceExplicitReconstruction").Uint())

			return core1_0.VKSuccess, nil
		})

	ycbcr, _, err := loader.Core1_1().CreateSamplerYcbcrConversion(device,
		core1_1.SamplerYcbcrConversionCreateOptions{
			Format:     core1_1.DataFormatB12X4G12X4R12X4G12X4HorizontalChromaComponentPacked,
			YcbcrModel: core1_1.SamplerYcbcrModelConversionYcbcr709,
			YcbcrRange: core1_1.SamplerYcbcrRangeITUNarrow,
			Components: core1_0.ComponentMapping{
				R: core1_0.SwizzleGreen,
				G: core1_0.SwizzleAlpha,
				B: core1_0.SwizzleIdentity,
				A: core1_0.SwizzleOne,
			},
			ChromaOffsetY:               core1_1.ChromaLocationCositedEven,
			ChromaOffsetX:               core1_1.ChromaLocationMidpoint,
			ChromaFilter:                core1_0.FilterLinear,
			ForceExplicitReconstruction: true,
		},
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, mockYcbcr.Handle(), ycbcr.Handle())

	coreDriver.EXPECT().VkDestroySamplerYcbcrConversion(
		device.Handle(),
		ycbcr.Handle(),
		gomock.Nil(),
	)

	ycbcr.Destroy(nil)
}

func TestVulkanLoader1_1_CreateDescriptorUpdateTemplate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	mockQueue := mocks.EasyMockQueue(ctrl)

	coreDriver.EXPECT().VkGetDeviceQueue2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pInfo *driver.VkDeviceQueueInfo2,
		pQueue *driver.VkQueue,
	) {
		*pQueue = mockQueue.Handle()

		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000145003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_QUEUE_INFO_2
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT
		require.Equal(t, uint64(3), val.FieldByName("queueFamilyIndex").Uint())
		require.Equal(t, uint64(5), val.FieldByName("queueIndex").Uint())
	})

	queue, err := loader.Core1_1().GetQueue(
		device,
		core1_1.DeviceQueueOptions{
			Flags:            core1_1.DeviceQueueCreateProtected,
			QueueFamilyIndex: 3,
			QueueIndex:       5,
		})

	require.NoError(t, err)
	require.Equal(t, mockQueue.Handle(), queue.Handle())
}
