package core1_1_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
	"go.uber.org/mock/gomock"
)

func TestTessellationDomainOriginOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_1, []string{}).(core1_1.Device)
	expectedPipeline := mocks1_1.EasyMockPipeline(ctrl)

	coreDriver.EXPECT().VkCreateGraphicsPipelines(device.Handle(), driver.VkPipelineCache(0), driver.Uint32(1), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pipelineCache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pPipelines *driver.VkPipeline) (common.VkResult, error) {
			pipelineSlice := ([]driver.VkPipeline)(unsafe.Slice(pPipelines, 1))
			pipelineSlice[0] = expectedPipeline.Handle()

			createInfoSlice := ([]driver.VkGraphicsPipelineCreateInfo)(unsafe.Slice(pCreateInfos, 1))
			val := reflect.ValueOf(createInfoSlice[0])

			require.Equal(t, uint64(28), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())

			tessellation := (*driver.VkPipelineTessellationStateCreateInfo)(val.FieldByName("pTessellationState").UnsafePointer())
			tessVal := reflect.ValueOf(tessellation).Elem()

			require.Equal(t, uint64(21), tessVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO
			require.Equal(t, uint64(1), tessVal.FieldByName("patchControlPoints").Uint())

			domain := (*driver.VkPipelineTessellationDomainOriginStateCreateInfo)(tessVal.FieldByName("pNext").UnsafePointer())
			domainVal := reflect.ValueOf(domain).Elem()

			require.Equal(t, uint64(1000117003), domainVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO
			require.True(t, domainVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), domainVal.FieldByName("domainOrigin").Uint())

			return core1_0.VKSuccess, nil
		})

	domainOriginState := core1_1.PipelineTessellationDomainOriginStateCreateInfo{
		DomainOrigin: core1_1.TessellationDomainOriginLowerLeft,
	}
	pipelines, _, err := device.CreateGraphicsPipelines(nil, nil, []core1_0.GraphicsPipelineCreateInfo{
		{
			TessellationState: &core1_0.PipelineTessellationStateCreateInfo{
				PatchControlPoints: 1,
				NextOptions:        common.NextOptions{Next: domainOriginState},
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.Equal(t, expectedPipeline.Handle(), pipelines[0].Handle())
}
