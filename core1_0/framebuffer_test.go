package core1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	internal_mocks "github.com/vkngwrapper/core/v2/internal/dummies"
	"github.com/vkngwrapper/core/v2/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreateFrameBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := internal_mocks.EasyDummyDevice(mockDriver)
	renderPass := mocks.EasyMockRenderPass(ctrl)
	imageView1 := mocks.EasyMockImageView(ctrl)
	imageView2 := mocks.EasyMockImageView(ctrl)
	framebufferHandle := mocks.NewFakeFramebufferHandle()

	mockDriver.EXPECT().VkCreateFramebuffer(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkFramebufferCreateInfo, pAllocator *driver.VkAllocationCallbacks, pFramebuffer *driver.VkFramebuffer) (common.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(37), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(3), val.FieldByName("width").Uint())
			require.Equal(t, uint64(5), val.FieldByName("height").Uint())
			require.Equal(t, uint64(7), val.FieldByName("layers").Uint())
			require.Equal(t, uint64(2), val.FieldByName("attachmentCount").Uint())

			require.Equal(t, renderPass.Handle(), (driver.VkRenderPass)(unsafe.Pointer(val.FieldByName("renderPass").Elem().UnsafeAddr())))

			attachmentPtr := (*driver.VkImageView)(unsafe.Pointer(val.FieldByName("pAttachments").Elem().UnsafeAddr()))
			attachmentSlice := ([]driver.VkImageView)(unsafe.Slice(attachmentPtr, 2))
			require.Equal(t, imageView1.Handle(), attachmentSlice[0])
			require.Equal(t, imageView2.Handle(), attachmentSlice[1])

			*pFramebuffer = framebufferHandle
			return core1_0.VKSuccess, nil
		})

	framebuffer, _, err := device.CreateFramebuffer(nil, core1_0.FramebufferCreateInfo{
		Flags:      0,
		RenderPass: renderPass,
		Width:      3,
		Height:     5,
		Layers:     7,
		Attachments: []core1_0.ImageView{
			imageView1, imageView2,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, framebuffer)
	require.Equal(t, framebufferHandle, framebuffer.Handle())
}
