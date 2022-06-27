package core1_1_test

import (
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

func TestTessellationDomainOriginOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := dummies.EasyDummyDevice(coreDriver)
	expectedPipeline := mocks.EasyMockPipeline(ctrl)

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

	domainOriginState := core1_1.PipelineTessellationDomainOriginStateOptions{
		DomainOrigin: core1_1.TessellationDomainOriginLowerLeft,
	}
	pipelines, _, err := device.CreateGraphicsPipelines(nil, nil, []core1_0.GraphicsPipelineCreateOptions{
		{
			Tessellation: &core1_0.TessellationStateOptions{
				PatchControlPoints: 1,
				HaveNext:           common.HaveNext{Next: domainOriginState},
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.Equal(t, expectedPipeline.Handle(), pipelines[0].Handle())
}
