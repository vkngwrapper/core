package impl1_2_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_2"
	"go.uber.org/mock/gomock"
)

func TestCommandBuffer_CmdBeginRenderPass2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)
	renderPass := mocks.NewDummyRenderPass(device)
	framebuffer := mocks.NewDummyFramebuffer(device)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkCmdBeginRenderPass2(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(commandBuffer loader.VkCommandBuffer,
		pRenderPassBegin *loader.VkRenderPassBeginInfo,
		pSubpassBeginInfo *loader.VkSubpassBeginInfo) {

		val := reflect.ValueOf(pRenderPassBegin).Elem()
		require.Equal(t, uint64(43), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, renderPass.Handle(), loader.VkRenderPass(val.FieldByName("renderPass").UnsafePointer()))
		require.Equal(t, framebuffer.Handle(), loader.VkFramebuffer(val.FieldByName("framebuffer").UnsafePointer()))
		require.Equal(t, int64(1), val.FieldByName("renderArea").FieldByName("offset").FieldByName("x").Int())
		require.Equal(t, int64(3), val.FieldByName("renderArea").FieldByName("offset").FieldByName("y").Int())
		require.Equal(t, uint64(5), val.FieldByName("renderArea").FieldByName("extent").FieldByName("width").Uint())
		require.Equal(t, uint64(7), val.FieldByName("renderArea").FieldByName("extent").FieldByName("height").Uint())
		require.Equal(t, uint64(1), val.FieldByName("clearValueCount").Uint())

		values := (*loader.Float)(unsafe.Pointer(val.FieldByName("pClearValues").Elem().UnsafeAddr()))
		valueSlice := ([]loader.Float)(unsafe.Slice(values, 4))
		val = reflect.ValueOf(valueSlice)
		require.InDelta(t, 1.0, val.Index(0).Float(), 0.0001)
		require.InDelta(t, 3.0, val.Index(1).Float(), 0.0001)
		require.InDelta(t, 5.0, val.Index(2).Float(), 0.0001)
		require.InDelta(t, 7.0, val.Index(3).Float(), 0.0001)

		val = reflect.ValueOf(pSubpassBeginInfo).Elem()
		require.Equal(t, uint64(1000109005), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("contents").Uint()) // VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS
	})

	err := driver.CmdBeginRenderPass2(
		commandBuffer,
		core1_0.RenderPassBeginInfo{
			RenderPass:  renderPass,
			Framebuffer: framebuffer,
			RenderArea:  core1_0.Rect2D{Offset: core1_0.Offset2D{1, 3}, Extent: core1_0.Extent2D{5, 7}},
			ClearValues: []core1_0.ClearValue{core1_0.ClearValueFloat{1, 3, 5, 7}},
		},
		core1_2.SubpassBeginInfo{
			Contents: core1_0.SubpassContentsSecondaryCommandBuffers,
		})
	require.NoError(t, err)
}

func TestCommandBuffer_CmdEndRenderPass2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkCmdEndRenderPass2(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(commandBuffer loader.VkCommandBuffer,
		pSubpassEndInfo *loader.VkSubpassEndInfo) {

		val := reflect.ValueOf(pSubpassEndInfo).Elem()

		require.Equal(t, uint64(1000109006), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_END_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
	})

	err := driver.CmdEndRenderPass2(commandBuffer, core1_2.SubpassEndInfo{})
	require.NoError(t, err)
}

func TestVulkanExtension_CmdNextSubpass2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkCmdNextSubpass2(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(commandBuffer loader.VkCommandBuffer,
		pSubpassBeginInfo *loader.VkSubpassBeginInfo,
		pSubpassEndInfo *loader.VkSubpassEndInfo) {

		val := reflect.ValueOf(pSubpassBeginInfo).Elem()
		require.Equal(t, uint64(1000109005), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("contents").Uint()) // VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS

		val = reflect.ValueOf(pSubpassEndInfo).Elem()
		require.Equal(t, uint64(1000109006), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_END_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
	})

	err := driver.CmdNextSubpass2(
		commandBuffer,
		core1_2.SubpassBeginInfo{
			Contents: core1_0.SubpassContentsSecondaryCommandBuffers,
		},
		core1_2.SubpassEndInfo{})
	require.NoError(t, err)
}

func TestCommandBuffer_CmdDrawIndexedIndirectCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalDeviceDriver(device, coreLoader)

	buffer := mocks.NewDummyBuffer(device)
	countBuffer := mocks.NewDummyBuffer(device)

	coreLoader.EXPECT().VkCmdDrawIndexedIndirectCount(
		commandBuffer.Handle(),
		buffer.Handle(),
		loader.VkDeviceSize(1),
		countBuffer.Handle(),
		loader.VkDeviceSize(3),
		loader.Uint32(5),
		loader.Uint32(7),
	)

	driver.CmdDrawIndexedIndirectCount(
		commandBuffer,
		buffer,
		uint64(1),
		countBuffer,
		uint64(3),
		5,
		7,
	)
}

func TestCommandBuffer_CmdDrawIndirectCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalDeviceDriver(device, coreLoader)

	buffer := mocks.NewDummyBuffer(device)
	countBuffer := mocks.NewDummyBuffer(device)

	coreLoader.EXPECT().VkCmdDrawIndirectCount(
		commandBuffer.Handle(),
		buffer.Handle(),
		loader.VkDeviceSize(11),
		countBuffer.Handle(),
		loader.VkDeviceSize(13),
		loader.Uint32(17),
		loader.Uint32(19),
	)

	driver.CmdDrawIndirectCount(
		commandBuffer,
		buffer,
		uint64(11),
		countBuffer,
		uint64(13),
		17,
		19,
	)
}
