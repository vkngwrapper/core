package core1_2_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/common/extensions"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestFramebufferAttachmentsCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_2)
	mockFramebuffer := mocks.EasyMockFramebuffer(ctrl)

	coreDriver.EXPECT().VkCreateFramebuffer(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkFramebufferCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pFramebuffer *driver.VkFramebuffer) (common.VkResult, error) {

		*pFramebuffer = mockFramebuffer.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(37), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO

		next := (*driver.VkFramebufferAttachmentsCreateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000108001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENTS_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("attachmentImageInfoCount").Uint())

		imageInfos := (*driver.VkFramebufferAttachmentImageInfo)(val.FieldByName("pAttachmentImageInfos").UnsafePointer())
		imageInfoSlice := unsafe.Slice(imageInfos, 2)
		val = reflect.ValueOf(imageInfoSlice)

		info := val.Index(0)
		require.Equal(t, uint64(1000108002), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0x10), info.FieldByName("flags").Uint()) // VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT
		require.Equal(t, uint64(4), info.FieldByName("usage").Uint())    // VK_IMAGE_USAGE_SAMPLED_BIT
		require.Equal(t, uint64(1), info.FieldByName("width").Uint())
		require.Equal(t, uint64(3), info.FieldByName("height").Uint())
		require.Equal(t, uint64(5), info.FieldByName("layerCount").Uint())
		require.Equal(t, uint64(2), info.FieldByName("viewFormatCount").Uint())

		viewFormats := (*driver.VkFormat)(info.FieldByName("pViewFormats").UnsafePointer())
		viewFormatSlice := unsafe.Slice(viewFormats, 2)

		require.Equal(t, []driver.VkFormat{68, 53}, viewFormatSlice)

		info = val.Index(1)
		require.Equal(t, uint64(1000108002), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), info.FieldByName("flags").Uint())    // VK_IMAGE_CREATE_SPARSE_BINDING_BIT
		require.Equal(t, uint64(0x10), info.FieldByName("usage").Uint()) // VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT
		require.Equal(t, uint64(7), info.FieldByName("width").Uint())
		require.Equal(t, uint64(11), info.FieldByName("height").Uint())
		require.Equal(t, uint64(13), info.FieldByName("layerCount").Uint())
		require.Equal(t, uint64(3), info.FieldByName("viewFormatCount").Uint())

		viewFormats = (*driver.VkFormat)(info.FieldByName("pViewFormats").UnsafePointer())
		viewFormatSlice = unsafe.Slice(viewFormats, 3)

		require.Equal(t, []driver.VkFormat{161, 164, 163}, viewFormatSlice)

		return core1_0.VKSuccess, nil
	})

	framebuffer, _, err := device.CreateFramebuffer(
		nil,
		core1_0.FramebufferCreateOptions{
			HaveNext: common.HaveNext{
				core1_2.FramebufferAttachmentsCreateOptions{
					AttachmentImageInfos: []core1_2.FramebufferAttachmentImageOptions{
						{
							Flags:      core1_0.ImageCreateCubeCompatible,
							Usage:      core1_0.ImageUsageSampled,
							Width:      1,
							Height:     3,
							LayerCount: 5,
							ViewFormats: []core1_0.DataFormat{
								core1_0.DataFormatA2B10G10R10UnsignedIntPacked,
								core1_0.DataFormatA8B8G8R8UnsignedScaledPacked,
							},
						},
						{
							Flags:      core1_0.ImageCreateSparseBinding,
							Usage:      core1_0.ImageUsageColorAttachment,
							Width:      7,
							Height:     11,
							LayerCount: 13,
							ViewFormats: []core1_0.DataFormat{
								core1_0.DataFormatASTC5x5_UnsignedNormalized,
								core1_0.DataFormatASTC6x5_sRGB,
								core1_0.DataFormatASTC6x5_UnsignedNormalized,
							},
						},
					},
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockFramebuffer.Handle(), framebuffer.Handle())
}
