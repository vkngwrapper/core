package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanLoader1_0_CreateFrameBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	imageView1 := mocks.EasyMockImageView(ctrl)
	imageView2 := mocks.EasyMockImageView(ctrl)
	framebufferHandle := mocks.NewFakeFramebufferHandle()

	driver.EXPECT().VkCreateFramebuffer(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkFramebufferCreateInfo, pAllocator *core.VkAllocationCallbacks, pFramebuffer *core.VkFramebuffer) (core.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(37), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(3), val.FieldByName("width").Uint())
			require.Equal(t, uint64(5), val.FieldByName("height").Uint())
			require.Equal(t, uint64(7), val.FieldByName("layers").Uint())
			require.Equal(t, uint64(2), val.FieldByName("attachmentCount").Uint())

			require.Equal(t, renderPass.Handle(), (core.VkRenderPass)(unsafe.Pointer(val.FieldByName("renderPass").Elem().UnsafeAddr())))

			attachmentPtr := (*core.VkImageView)(unsafe.Pointer(val.FieldByName("pAttachments").Elem().UnsafeAddr()))
			attachmentSlice := ([]core.VkImageView)(unsafe.Slice(attachmentPtr, 2))
			require.Equal(t, imageView1.Handle(), attachmentSlice[0])
			require.Equal(t, imageView2.Handle(), attachmentSlice[1])

			*pFramebuffer = framebufferHandle
			return core.VKSuccess, nil
		})

	framebuffer, _, err := loader.CreateFrameBuffer(device, &core.FramebufferOptions{
		Flags:      core.FramebufferCreateImageless,
		RenderPass: renderPass,
		Width:      3,
		Height:     5,
		Layers:     7,
		Attachments: []core.ImageView{
			imageView1, imageView2,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, framebuffer)
	require.Equal(t, framebufferHandle, framebuffer.Handle())
}
