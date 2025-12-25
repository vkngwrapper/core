package core1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreatePipelineLayout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	descriptorSetLayout1 := mocks.NewDummyDescriptorSetLayout(device)
	descriptorSetLayout2 := mocks.NewDummyDescriptorSetLayout(device)
	layoutHandle := mocks.NewFakePipelineLayout()

	mockLoader.EXPECT().VkCreatePipelineLayout(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkPipelineLayoutCreateInfo, pAllocator *loader.VkAllocationCallbacks, pPipelineLayout *loader.VkPipelineLayout) (common.VkResult, error) {
			*pPipelineLayout = layoutHandle

			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(30), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(2), val.FieldByName("setLayoutCount").Uint())
			require.Equal(t, uint64(3), val.FieldByName("pushConstantRangeCount").Uint())

			setLayoutsPtr := (*loader.VkDescriptorSetLayout)(unsafe.Pointer(val.FieldByName("pSetLayouts").Elem().UnsafeAddr()))
			setLayoutsSlice := ([]loader.VkDescriptorSetLayout)(unsafe.Slice(setLayoutsPtr, 2))

			require.Equal(t, descriptorSetLayout1.Handle(), setLayoutsSlice[0])
			require.Equal(t, descriptorSetLayout2.Handle(), setLayoutsSlice[1])

			pushConstantsPtr := (*loader.VkPushConstantRange)(unsafe.Pointer(val.FieldByName("pPushConstantRanges").Elem().UnsafeAddr()))
			pushConstantsSlice := reflect.ValueOf(([]loader.VkPushConstantRange)(unsafe.Slice(pushConstantsPtr, 3)))

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

	layout, _, err := driver.CreatePipelineLayout(device, nil, core1_0.PipelineLayoutCreateInfo{
		SetLayouts: []core.DescriptorSetLayout{descriptorSetLayout1, descriptorSetLayout2},
		PushConstantRanges: []core1_0.PushConstantRange{
			{
				StageFlags: core1_0.StageFragment,
				Offset:     1,
				Size:       3,
			},
			{
				StageFlags: core1_0.StageVertex,
				Offset:     5,
				Size:       7,
			},
			{
				StageFlags: core1_0.StageTessellationControl,
				Offset:     11,
				Size:       13,
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, layout)
	require.Equal(t, layoutHandle, layout.Handle())
}
