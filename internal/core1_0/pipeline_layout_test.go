package internal1_0_test

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
	"unsafe"
)

func TestVulkanLoader1_0_CreatePipelineLayout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := internal_mocks.EasyDummyDevice(mockDriver)
	descriptorSetLayout1 := mocks.EasyMockDescriptorSetLayout(ctrl)
	descriptorSetLayout2 := mocks.EasyMockDescriptorSetLayout(ctrl)
	layoutHandle := mocks.NewFakePipelineLayout()

	mockDriver.EXPECT().VkCreatePipelineLayout(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkPipelineLayoutCreateInfo, pAllocator *driver.VkAllocationCallbacks, pPipelineLayout *driver.VkPipelineLayout) (common.VkResult, error) {
			*pPipelineLayout = layoutHandle

			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(30), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(2), val.FieldByName("setLayoutCount").Uint())
			require.Equal(t, uint64(3), val.FieldByName("pushConstantRangeCount").Uint())

			setLayoutsPtr := (*driver.VkDescriptorSetLayout)(unsafe.Pointer(val.FieldByName("pSetLayouts").Elem().UnsafeAddr()))
			setLayoutsSlice := ([]driver.VkDescriptorSetLayout)(unsafe.Slice(setLayoutsPtr, 2))

			require.Equal(t, descriptorSetLayout1.Handle(), setLayoutsSlice[0])
			require.Equal(t, descriptorSetLayout2.Handle(), setLayoutsSlice[1])

			pushConstantsPtr := (*driver.VkPushConstantRange)(unsafe.Pointer(val.FieldByName("pPushConstantRanges").Elem().UnsafeAddr()))
			pushConstantsSlice := reflect.ValueOf(([]driver.VkPushConstantRange)(unsafe.Slice(pushConstantsPtr, 3)))

			pushConstant := pushConstantsSlice.Index(0)
			require.Equal(t, uint64(0x00000010), pushConstant.FieldByName("stageFlags").Uint()) // VK_SHADER_STAGE_FRAGMENT_BIT
			require.Equal(t, uint64(1), pushConstant.FieldByName("offset").Uint())
			require.Equal(t, uint64(3), pushConstant.FieldByName("size").Uint())

			pushConstant = pushConstantsSlice.Index(1)
			require.Equal(t, uint64(1), pushConstant.FieldByName("stageFlags").Uint()) // VK_SHADER_STAGE_VERTEX_BIT
			require.Equal(t, uint64(5), pushConstant.FieldByName("offset").Uint())
			require.Equal(t, uint64(7), pushConstant.FieldByName("size").Uint())

			pushConstant = pushConstantsSlice.Index(2)
			require.Equal(t, uint64(0x00000002), pushConstant.FieldByName("stageFlags").Uint()) // VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT
			require.Equal(t, uint64(11), pushConstant.FieldByName("offset").Uint())
			require.Equal(t, uint64(13), pushConstant.FieldByName("size").Uint())

			return core1_0.VKSuccess, nil
		})

	layout, _, err := device.CreatePipelineLayout(nil, core1_0.PipelineLayoutCreateOptions{
		SetLayouts: []core1_0.DescriptorSetLayout{descriptorSetLayout1, descriptorSetLayout2},
		PushConstantRanges: []core1_0.PushConstantRange{
			{
				Stages: core1_0.StageFragment,
				Offset: 1,
				Size:   3,
			},
			{
				Stages: core1_0.StageVertex,
				Offset: 5,
				Size:   7,
			},
			{
				Stages: core1_0.StageTessellationControl,
				Offset: 11,
				Size:   13,
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, layout)
	require.Equal(t, layoutHandle, layout.Handle())
}
