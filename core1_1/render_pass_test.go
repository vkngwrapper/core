package core1_1_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
	"go.uber.org/mock/gomock"
)

func TestInputAttachmentAspectOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_1)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_1, []string{}).(core1_1.Device)

	expectedRenderPass := mocks1_1.EasyMockRenderPass(ctrl)

	coreDriver.EXPECT().VkCreateRenderPass(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkDevice, pCreateInfo *loader.VkRenderPassCreateInfo, pAllocator *loader.VkAllocationCallbacks, pRenderPass *loader.VkRenderPass) (common.VkResult, error) {
			*pRenderPass = expectedRenderPass.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(38), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO

			aspectOptions := (*loader.VkRenderPassInputAttachmentAspectCreateInfo)(val.FieldByName("pNext").UnsafePointer())
			aspectVal := reflect.ValueOf(aspectOptions).Elem()
			require.Equal(t, uint64(1000117001), aspectVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO
			require.True(t, aspectVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), aspectVal.FieldByName("aspectReferenceCount").Uint())

			refsPtr := (*loader.VkInputAttachmentAspectReference)(aspectVal.FieldByName("pAspectReferences").UnsafePointer())
			refsSlice := ([]loader.VkInputAttachmentAspectReference)(unsafe.Slice(refsPtr, 2))
			refsVal := reflect.ValueOf(refsSlice)
			ref := refsVal.Index(0)
			require.Equal(t, uint64(1), ref.FieldByName("subpass").Uint())
			require.Equal(t, uint64(3), ref.FieldByName("inputAttachmentIndex").Uint())
			require.Equal(t, uint64(0x00000001), ref.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_COLOR_BIT

			ref = refsVal.Index(1)
			require.Equal(t, uint64(5), ref.FieldByName("subpass").Uint())
			require.Equal(t, uint64(7), ref.FieldByName("inputAttachmentIndex").Uint())
			require.Equal(t, uint64(0x00000008), ref.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_METADATA_BIT

			return core1_0.VKSuccess, nil
		})

	aspectOptions := core1_1.RenderPassInputAttachmentAspectCreateInfo{
		AspectReferences: []core1_1.InputAttachmentAspectReference{
			{
				Subpass:              1,
				InputAttachmentIndex: 3,
				AspectMask:           core1_0.ImageAspectColor,
			},
			{
				Subpass:              5,
				InputAttachmentIndex: 7,
				AspectMask:           core1_0.ImageAspectMetadata,
			},
		},
	}
	renderPass, _, err := device.CreateRenderPass(nil, core1_0.RenderPassCreateInfo{
		NextOptions: common.NextOptions{Next: aspectOptions},
	})
	require.NoError(t, err)
	require.Equal(t, expectedRenderPass.Handle(), renderPass.Handle())
}

func TestDeviceGroupRenderPassBeginOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	device := mocks1_1.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks1_1.EasyMockCommandPool(ctrl, device)

	builder := &impl1_1.DeviceObjectBuilderImpl{}
	commandBuffer := builder.CreateCommandBufferObject(coreDriver, commandPool.Handle(), device.Handle(), mocks.NewFakeCommandBufferHandle(), common.Vulkan1_1)
	renderPass := mocks1_1.EasyMockRenderPass(ctrl)
	framebuffer := mocks1_1.EasyMockFramebuffer(ctrl)

	coreDriver.EXPECT().VkCmdBeginRenderPass(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
		loader.VkSubpassContents(1), // VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS
	).DoAndReturn(func(
		commandBuffer loader.VkCommandBuffer,
		pRenderPassBegin *loader.VkRenderPassBeginInfo,
		contents loader.VkSubpassContents,
	) {
		val := reflect.ValueOf(pRenderPassBegin).Elem()
		require.Equal(t, uint64(43), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO
		require.Equal(t, renderPass.Handle(), (loader.VkRenderPass)(val.FieldByName("renderPass").UnsafePointer()))
		require.Equal(t, framebuffer.Handle(), (loader.VkFramebuffer)(val.FieldByName("framebuffer").UnsafePointer()))

		next := (*loader.VkDeviceGroupRenderPassBeginInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(7), val.FieldByName("deviceMask").Uint())
		require.Equal(t, uint64(2), val.FieldByName("deviceRenderAreaCount").Uint())

		areas := (*loader.VkRect2D)(val.FieldByName("pDeviceRenderAreas").UnsafePointer())
		areaSlice := ([]loader.VkRect2D)(unsafe.Slice(areas, 2))
		val = reflect.ValueOf(areaSlice)

		oneArea := val.Index(0)
		require.Equal(t, int64(1), oneArea.FieldByName("offset").FieldByName("x").Int())
		require.Equal(t, int64(3), oneArea.FieldByName("offset").FieldByName("y").Int())
		require.Equal(t, uint64(5), oneArea.FieldByName("extent").FieldByName("width").Uint())
		require.Equal(t, uint64(7), oneArea.FieldByName("extent").FieldByName("height").Uint())

		oneArea = val.Index(1)
		require.Equal(t, int64(11), oneArea.FieldByName("offset").FieldByName("x").Int())
		require.Equal(t, int64(13), oneArea.FieldByName("offset").FieldByName("y").Int())
		require.Equal(t, uint64(17), oneArea.FieldByName("extent").FieldByName("width").Uint())
		require.Equal(t, uint64(19), oneArea.FieldByName("extent").FieldByName("height").Uint())
	})

	err := commandBuffer.CmdBeginRenderPass(
		core1_0.SubpassContentsSecondaryCommandBuffers,
		core1_0.RenderPassBeginInfo{
			RenderPass:  renderPass,
			Framebuffer: framebuffer,
			NextOptions: common.NextOptions{
				core1_1.DeviceGroupRenderPassBeginInfo{
					DeviceMask: 7,
					DeviceRenderAreas: []core1_0.Rect2D{
						{
							Offset: core1_0.Offset2D{X: 1, Y: 3},
							Extent: core1_0.Extent2D{Width: 5, Height: 7},
						},
						{
							Offset: core1_0.Offset2D{X: 11, Y: 13},
							Extent: core1_0.Extent2D{Width: 17, Height: 19},
						},
					},
				},
			},
		},
	)
	require.NoError(t, err)
}

func TestRenderPassMultiviewOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_1)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_1, []string{}).(core1_1.Device)

	mockRenderPass := mocks1_1.EasyMockRenderPass(ctrl)

	coreDriver.EXPECT().VkCreateRenderPass(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pCreateInfo *loader.VkRenderPassCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pRenderPass *loader.VkRenderPass) (common.VkResult, error) {

		*pRenderPass = mockRenderPass.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(38), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO

		next := (*loader.VkRenderPassMultiviewCreateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000053000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("subpassCount").Uint())
		require.Equal(t, uint64(2), val.FieldByName("dependencyCount").Uint())
		require.Equal(t, uint64(1), val.FieldByName("correlationMaskCount").Uint())

		masks := (*loader.Uint32)(val.FieldByName("pViewMasks").UnsafePointer())
		maskSlice := ([]loader.Uint32)(unsafe.Slice(masks, 3))
		require.Equal(t, []loader.Uint32{1, 2, 7}, maskSlice)

		offsets := (*loader.Int32)(val.FieldByName("pViewOffsets").UnsafePointer())
		offsetSlice := ([]loader.Int32)(unsafe.Slice(offsets, 2))
		require.Equal(t, []loader.Int32{11, 13}, offsetSlice)

		correlationMasks := (*loader.Uint32)(val.FieldByName("pCorrelationMasks").UnsafePointer())
		correlationSlice := ([]loader.Uint32)(unsafe.Slice(correlationMasks, 1))
		require.Equal(t, []loader.Uint32{17}, correlationSlice)

		return core1_0.VKSuccess, nil
	})

	renderPass, _, err := device.CreateRenderPass(nil, core1_0.RenderPassCreateInfo{
		NextOptions: common.NextOptions{
			core1_1.RenderPassMultiviewCreateInfo{
				ViewMasks:        []uint32{1, 2, 7},
				ViewOffsets:      []int{11, 13},
				CorrelationMasks: []uint32{17},
			},
		},
	})

	require.NoError(t, err)
	require.Equal(t, mockRenderPass.Handle(), renderPass.Handle())

}
