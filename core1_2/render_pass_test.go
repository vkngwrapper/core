package core1_2_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/internal/impl1_2"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestAttachmentDescriptionStencilLayoutOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	builder := &impl1_2.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_2, []string{}).(core1_2.Device)
	mockRenderPass := mocks.EasyMockRenderPass(ctrl)

	coreDriver.EXPECT().VkCreateRenderPass2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkRenderPassCreateInfo2,
		pAllocator *driver.VkAllocationCallbacks,
		pRenderPass *driver.VkRenderPass) (common.VkResult, error) {

		*pRenderPass = mockRenderPass.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(1000109004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2
		require.True(t, val.FieldByName("pNext").IsNil())

		require.Equal(t, uint64(1), val.FieldByName("subpassCount").Uint())
		require.Equal(t, uint64(1), val.FieldByName("attachmentCount").Uint())

		attachment := val.FieldByName("pAttachments").Elem()
		require.Equal(t, uint64(1000109000), attachment.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2
		attachmentNext := (*driver.VkAttachmentDescriptionStencilLayout)(attachment.FieldByName("pNext").UnsafePointer())

		attachmentLayout := reflect.ValueOf(attachmentNext).Elem()
		require.Equal(t, uint64(1000241002), attachmentLayout.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_STENCIL_LAYOUT
		require.True(t, attachmentLayout.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1000241000), attachmentLayout.FieldByName("stencilInitialLayout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL
		require.Equal(t, uint64(1000241003), attachmentLayout.FieldByName("stencilFinalLayout").Uint())   // VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL

		subpass := val.FieldByName("pSubpasses").Elem()
		require.Equal(t, uint64(1000109002), subpass.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2
		require.True(t, subpass.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), subpass.FieldByName("inputAttachmentCount").Uint())

		inputAttachment := subpass.FieldByName("pInputAttachments").Elem()
		require.Equal(t, uint64(1000109001), inputAttachment.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2

		inputAttachmentNext := (*driver.VkAttachmentReferenceStencilLayout)(inputAttachment.FieldByName("pNext").UnsafePointer())
		stencilRef := reflect.ValueOf(inputAttachmentNext).Elem()
		require.Equal(t, uint64(1000241001), stencilRef.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_STENCIL_LAYOUT
		require.True(t, stencilRef.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1000241000), stencilRef.FieldByName("stencilLayout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL

		return core1_0.VKSuccess, nil
	})

	renderPass, _, err := device.CreateRenderPass2(
		nil,
		core1_2.RenderPassCreateInfo2{
			Attachments: []core1_2.AttachmentDescription2{
				{
					NextOptions: common.NextOptions{core1_2.AttachmentDescriptionStencilLayout{
						StencilInitialLayout: core1_2.ImageLayoutDepthAttachmentOptimal,
						StencilFinalLayout:   core1_2.ImageLayoutStencilReadOnlyOptimal,
					}},
				},
			},
			Subpasses: []core1_2.SubpassDescription2{
				{
					InputAttachments: []core1_2.AttachmentReference2{
						{
							NextOptions: common.NextOptions{
								core1_2.AttachmentReferenceStencilLayout{
									StencilLayout: core1_2.ImageLayoutDepthAttachmentOptimal,
								},
							},
						},
					},
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockRenderPass.Handle(), renderPass.Handle())
}

func TestRenderPassAttachmentBeginInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks.EasyMockCommandPool(ctrl, device)
	builder := &impl1_2.DeviceObjectBuilderImpl{}
	commandBuffer := builder.CreateCommandBufferObject(coreDriver, commandPool.Handle(), device.Handle(), mocks.NewFakeCommandBufferHandle(), common.Vulkan1_0)

	imageView1 := mocks.EasyMockImageView(ctrl)
	imageView2 := mocks.EasyMockImageView(ctrl)

	coreDriver.EXPECT().VkCmdBeginRenderPass(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
		driver.VkSubpassContents(0), // VK_SUBPASS_CONTENTS_INLINE
	).DoAndReturn(func(commandBuffer driver.VkCommandBuffer,
		pRenderPassBegin *driver.VkRenderPassBeginInfo,
		contents driver.VkSubpassContents) {

		val := reflect.ValueOf(pRenderPassBegin).Elem()
		require.Equal(t, uint64(43), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO

		next := (*driver.VkRenderPassAttachmentBeginInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000108003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("attachmentCount").Uint())

		firstImageView := val.FieldByName("pAttachments").UnsafePointer()
		require.Equal(t, imageView1.Handle(), *(*driver.VkImageView)(firstImageView))

		secondImageView := unsafe.Add(firstImageView, unsafe.Sizeof(uintptr(0)))
		require.Equal(t, imageView2.Handle(), *(*driver.VkImageView)(secondImageView))
	})

	err := commandBuffer.CmdBeginRenderPass(core1_0.SubpassContentsInline, core1_0.RenderPassBeginInfo{
		NextOptions: common.NextOptions{core1_2.RenderPassAttachmentBeginInfo{
			Attachments: []core1_0.ImageView{imageView1, imageView2},
		}},
	})
	require.NoError(t, err)
}

func TestSubpassDescriptionDepthStencilResolveOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	builder := &impl1_2.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_2, []string{}).(core1_2.Device)
	mockRenderPass := mocks.EasyMockRenderPass(ctrl)

	coreDriver.EXPECT().VkCreateRenderPass2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkRenderPassCreateInfo2,
		pAllocator *driver.VkAllocationCallbacks,
		pRenderPass *driver.VkRenderPass) (common.VkResult, error) {

		*pRenderPass = mockRenderPass.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(1000109004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("subpassCount").Uint())

		subpass := val.FieldByName("pSubpasses").Elem()
		require.Equal(t, uint64(1000109002), subpass.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2

		next := (*driver.VkSubpassDescriptionDepthStencilResolve)(subpass.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000199001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(4), val.FieldByName("depthResolveMode").Uint())   // VK_RESOLVE_MODE_MIN_BIT
		require.Equal(t, uint64(1), val.FieldByName("stencilResolveMode").Uint()) // VK_RESOLVE_MODE_SAMPLE_ZERO_BIT

		val = val.FieldByName("pDepthStencilResolveAttachment").Elem()

		require.Equal(t, uint64(1000109001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("attachment").Uint())
		require.Equal(t, uint64(7), val.FieldByName("layout").Uint())     // VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL
		require.Equal(t, uint64(4), val.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_STENCIL_BIT

		return core1_0.VKSuccess, nil
	})

	renderPass, _, err := device.CreateRenderPass2(nil,
		core1_2.RenderPassCreateInfo2{
			Subpasses: []core1_2.SubpassDescription2{
				{
					NextOptions: common.NextOptions{
						core1_2.SubpassDescriptionDepthStencilResolve{
							DepthResolveMode:   core1_2.ResolveModeMin,
							StencilResolveMode: core1_2.ResolveModeSampleZero,
							DepthStencilResolveAttachment: &core1_2.AttachmentReference2{
								Attachment: 3,
								Layout:     core1_0.ImageLayoutTransferDstOptimal,
								AspectMask: core1_0.ImageAspectStencil,
							},
						},
					},
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockRenderPass.Handle(), renderPass.Handle())

}
