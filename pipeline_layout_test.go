package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
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

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	descriptorSetLayout1 := mocks.EasyMockDescriptorSetLayout(ctrl)
	descriptorSetLayout2 := mocks.EasyMockDescriptorSetLayout(ctrl)
	layoutHandle := mocks.NewFakePipelineLayout()

	driver.EXPECT().VkCreatePipelineLayout(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkPipelineLayoutCreateInfo, pAllocator *core.VkAllocationCallbacks, pPipelineLayout *core.VkPipelineLayout) (core.VkResult, error) {
			*pPipelineLayout = layoutHandle

			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(30), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(2), val.FieldByName("setLayoutCount").Uint())
			require.Equal(t, uint64(3), val.FieldByName("pushConstantRangeCount").Uint())

			setLayoutsPtr := (*core.VkDescriptorSetLayout)(unsafe.Pointer(val.FieldByName("pSetLayouts").Elem().UnsafeAddr()))
			setLayoutsSlice := ([]core.VkDescriptorSetLayout)(unsafe.Slice(setLayoutsPtr, 2))

			require.Equal(t, descriptorSetLayout1.Handle(), setLayoutsSlice[0])
			require.Equal(t, descriptorSetLayout2.Handle(), setLayoutsSlice[1])

			pushConstantsPtr := (*core.VkPushConstantRange)(unsafe.Pointer(val.FieldByName("pPushConstantRanges").Elem().UnsafeAddr()))
			pushConstantsSlice := reflect.ValueOf(([]core.VkPushConstantRange)(unsafe.Slice(pushConstantsPtr, 3)))

			pushConstant := pushConstantsSlice.Index(0)
			require.Equal(t, uint64(0x00000010), pushConstant.FieldByName("stageFlags").Uint()) // VK_SHADER_STAGE_FRAGMENT_BIT
			require.Equal(t, uint64(1), pushConstant.FieldByName("offset").Uint())
			require.Equal(t, uint64(3), pushConstant.FieldByName("size").Uint())

			pushConstant = pushConstantsSlice.Index(1)
			require.Equal(t, uint64(1), pushConstant.FieldByName("stageFlags").Uint()) // VK_SHADER_STAGE_VERTEX_BIT
			require.Equal(t, uint64(5), pushConstant.FieldByName("offset").Uint())
			require.Equal(t, uint64(7), pushConstant.FieldByName("size").Uint())

			pushConstant = pushConstantsSlice.Index(2)
			require.Equal(t, uint64(0x00000200), pushConstant.FieldByName("stageFlags").Uint()) // VK_SHADER_STAGE_ANY_HIT_BIT_KHR
			require.Equal(t, uint64(11), pushConstant.FieldByName("offset").Uint())
			require.Equal(t, uint64(13), pushConstant.FieldByName("size").Uint())

			return core.VKSuccess, nil
		})

	layout, _, err := loader.CreatePipelineLayout(device, &core.PipelineLayoutOptions{
		SetLayouts: []core.DescriptorSetLayout{descriptorSetLayout1, descriptorSetLayout2},
		PushConstantRanges: []common.PushConstantRange{
			{
				Stages: common.StageFragment,
				Offset: 1,
				Size:   3,
			},
			{
				Stages: common.StageVertex,
				Offset: 5,
				Size:   7,
			},
			{
				Stages: common.StageAnyHitKHR,
				Offset: 11,
				Size:   13,
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, layout)
	require.Equal(t, layoutHandle, layout.Handle())
}
