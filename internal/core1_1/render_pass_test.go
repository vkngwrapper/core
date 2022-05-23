package internal1_1_test

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

func TestInputAttachmentAspectOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := dummies.EasyDummyDevice(coreDriver)
	expectedRenderPass := mocks.EasyMockRenderPass(ctrl)

	coreDriver.EXPECT().VkCreateRenderPass(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pCreateInfo *driver.VkRenderPassCreateInfo, pAllocator *driver.VkAllocationCallbacks, pRenderPass *driver.VkRenderPass) (common.VkResult, error) {
			*pRenderPass = expectedRenderPass.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(38), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO

			aspectOptions := (*driver.VkRenderPassInputAttachmentAspectCreateInfo)(val.FieldByName("pNext").UnsafePointer())
			aspectVal := reflect.ValueOf(aspectOptions).Elem()
			require.Equal(t, uint64(1000117001), aspectVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO
			require.True(t, aspectVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), aspectVal.FieldByName("aspectReferenceCount").Uint())

			refsPtr := (*driver.VkInputAttachmentAspectReference)(aspectVal.FieldByName("pAspectReferences").UnsafePointer())
			refsSlice := ([]driver.VkInputAttachmentAspectReference)(unsafe.Slice(refsPtr, 2))
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

	aspectOptions := core1_1.RenderPassInputAttachmentAspectOptions{
		AspectReferences: []core1_1.InputAttachmentAspectReference{
			{
				Subpass:              1,
				InputAttachmentIndex: 3,
				AspectMask:           core1_0.AspectColor,
			},
			{
				Subpass:              5,
				InputAttachmentIndex: 7,
				AspectMask:           core1_0.AspectMetadata,
			},
		},
	}
	renderPass, _, err := device.CreateRenderPass(nil, core1_0.RenderPassCreateOptions{
		HaveNext: common.HaveNext{Next: aspectOptions},
	})
	require.NoError(t, err)
	require.Equal(t, expectedRenderPass.Handle(), renderPass.Handle())
}

func TestDeviceGroupRenderPassBeginOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks.EasyMockCommandPool(ctrl, device)

	commandBuffer := dummies.EasyDummyCommandBuffer(coreDriver, device, commandPool)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	framebuffer := mocks.EasyMockFramebuffer(ctrl)

	coreDriver.EXPECT().VkCmdBeginRenderPass(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
		driver.VkSubpassContents(1), // VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS
	).DoAndReturn(func(
		commandBuffer driver.VkCommandBuffer,
		pRenderPassBegin *driver.VkRenderPassBeginInfo,
		contents driver.VkSubpassContents,
	) {
		val := reflect.ValueOf(pRenderPassBegin).Elem()
		require.Equal(t, uint64(43), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO
		require.Equal(t, renderPass.Handle(), (driver.VkRenderPass)(val.FieldByName("renderPass").UnsafePointer()))
		require.Equal(t, framebuffer.Handle(), (driver.VkFramebuffer)(val.FieldByName("framebuffer").UnsafePointer()))

		next := (*driver.VkDeviceGroupRenderPassBeginInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(7), val.FieldByName("deviceMask").Uint())
		require.Equal(t, uint64(2), val.FieldByName("deviceRenderAreaCount").Uint())

		areas := (*driver.VkRect2D)(val.FieldByName("pDeviceRenderAreas").UnsafePointer())
		areaSlice := ([]driver.VkRect2D)(unsafe.Slice(areas, 2))
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
		core1_0.RenderPassBeginOptions{
			RenderPass:  renderPass,
			Framebuffer: framebuffer,
			HaveNext: common.HaveNext{
				core1_1.DeviceGroupRenderPassBeginOptions{
					DeviceMask: 7,
					DeviceRenderAreas: []common.Rect2D{
						{
							Offset: common.Offset2D{X: 1, Y: 3},
							Extent: common.Extent2D{Width: 5, Height: 7},
						},
						{
							Offset: common.Offset2D{X: 11, Y: 13},
							Extent: common.Extent2D{Width: 17, Height: 19},
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

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := dummies.EasyDummyDevice(coreDriver)
	mockRenderPass := mocks.EasyMockRenderPass(ctrl)

	coreDriver.EXPECT().VkCreateRenderPass(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkRenderPassCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pRenderPass *driver.VkRenderPass) (common.VkResult, error) {

		*pRenderPass = mockRenderPass.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(38), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO

		next := (*driver.VkRenderPassMultiviewCreateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000053000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("subpassCount").Uint())
		require.Equal(t, uint64(2), val.FieldByName("dependencyCount").Uint())
		require.Equal(t, uint64(1), val.FieldByName("correlationMaskCount").Uint())

		masks := (*driver.Uint32)(val.FieldByName("pViewMasks").UnsafePointer())
		maskSlice := ([]driver.Uint32)(unsafe.Slice(masks, 3))
		require.Equal(t, []driver.Uint32{1, 2, 7}, maskSlice)

		offsets := (*driver.Int32)(val.FieldByName("pViewOffsets").UnsafePointer())
		offsetSlice := ([]driver.Int32)(unsafe.Slice(offsets, 2))
		require.Equal(t, []driver.Int32{11, 13}, offsetSlice)

		correlationMasks := (*driver.Uint32)(val.FieldByName("pCorrelationMasks").UnsafePointer())
		correlationSlice := ([]driver.Uint32)(unsafe.Slice(correlationMasks, 1))
		require.Equal(t, []driver.Uint32{17}, correlationSlice)

		return core1_0.VKSuccess, nil
	})

	renderPass, _, err := device.CreateRenderPass(nil, core1_0.RenderPassCreateOptions{
		HaveNext: common.HaveNext{
			core1_1.RenderPassMultiviewOptions{
				SubpassViewMasks:      []uint32{1, 2, 7},
				DependencyViewOffsets: []int{11, 13},
				CorrelationMasks:      []uint32{17},
			},
		},
	})

	require.NoError(t, err)
	require.Equal(t, mockRenderPass.Handle(), renderPass.Handle())

}
