package core1_2_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_2"
	"github.com/vkngwrapper/core/driver"
	mock_driver "github.com/vkngwrapper/core/driver/mocks"
	"github.com/vkngwrapper/core/internal/dummies"
	"github.com/vkngwrapper/core/mocks"
	"reflect"
	"testing"
	"unsafe"
)

func TestCommandBuffer_CmdBeginRenderPass2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks.EasyMockCommandPool(ctrl, device)
	commandBuffer := core1_2.PromoteCommandBuffer(dummies.EasyDummyCommandBuffer(coreDriver, device, commandPool))
	renderPass := mocks.EasyMockRenderPass(ctrl)
	framebuffer := mocks.EasyMockFramebuffer(ctrl)

	coreDriver.EXPECT().VkCmdBeginRenderPass2(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(commandBuffer driver.VkCommandBuffer,
		pRenderPassBegin *driver.VkRenderPassBeginInfo,
		pSubpassBeginInfo *driver.VkSubpassBeginInfo) {

		val := reflect.ValueOf(pRenderPassBegin).Elem()
		require.Equal(t, uint64(43), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, renderPass.Handle(), driver.VkRenderPass(val.FieldByName("renderPass").UnsafePointer()))
		require.Equal(t, framebuffer.Handle(), driver.VkFramebuffer(val.FieldByName("framebuffer").UnsafePointer()))
		require.Equal(t, int64(1), val.FieldByName("renderArea").FieldByName("offset").FieldByName("x").Int())
		require.Equal(t, int64(3), val.FieldByName("renderArea").FieldByName("offset").FieldByName("y").Int())
		require.Equal(t, uint64(5), val.FieldByName("renderArea").FieldByName("extent").FieldByName("width").Uint())
		require.Equal(t, uint64(7), val.FieldByName("renderArea").FieldByName("extent").FieldByName("height").Uint())
		require.Equal(t, uint64(1), val.FieldByName("clearValueCount").Uint())

		values := (*driver.Float)(unsafe.Pointer(val.FieldByName("pClearValues").Elem().UnsafeAddr()))
		valueSlice := ([]driver.Float)(unsafe.Slice(values, 4))
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

	err := commandBuffer.CmdBeginRenderPass2(
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

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks.EasyMockCommandPool(ctrl, device)
	commandBuffer := core1_2.PromoteCommandBuffer(dummies.EasyDummyCommandBuffer(coreDriver, device, commandPool))

	coreDriver.EXPECT().VkCmdEndRenderPass2(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(commandBuffer driver.VkCommandBuffer,
		pSubpassEndInfo *driver.VkSubpassEndInfo) {

		val := reflect.ValueOf(pSubpassEndInfo).Elem()

		require.Equal(t, uint64(1000109006), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_END_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
	})

	err := commandBuffer.CmdEndRenderPass2(core1_2.SubpassEndInfo{})
	require.NoError(t, err)
}

func TestVulkanExtension_CmdNextSubpass2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks.EasyMockCommandPool(ctrl, device)
	commandBuffer := core1_2.PromoteCommandBuffer(dummies.EasyDummyCommandBuffer(coreDriver, device, commandPool))

	coreDriver.EXPECT().VkCmdNextSubpass2(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(commandBuffer driver.VkCommandBuffer,
		pSubpassBeginInfo *driver.VkSubpassBeginInfo,
		pSubpassEndInfo *driver.VkSubpassEndInfo) {

		val := reflect.ValueOf(pSubpassBeginInfo).Elem()
		require.Equal(t, uint64(1000109005), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("contents").Uint()) // VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS

		val = reflect.ValueOf(pSubpassEndInfo).Elem()
		require.Equal(t, uint64(1000109006), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_END_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
	})

	err := commandBuffer.CmdNextSubpass2(
		core1_2.SubpassBeginInfo{
			Contents: core1_0.SubpassContentsSecondaryCommandBuffers,
		},
		core1_2.SubpassEndInfo{})
	require.NoError(t, err)
}

func TestCommandBuffer_CmdDrawIndexedIndirectCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks.EasyMockCommandPool(ctrl, device)
	commandBuffer := core1_2.PromoteCommandBuffer(dummies.EasyDummyCommandBuffer(coreDriver, device, commandPool))
	buffer := mocks.EasyMockBuffer(ctrl)
	countBuffer := mocks.EasyMockBuffer(ctrl)

	coreDriver.EXPECT().VkCmdDrawIndexedIndirectCount(
		commandBuffer.Handle(),
		buffer.Handle(),
		driver.VkDeviceSize(1),
		countBuffer.Handle(),
		driver.VkDeviceSize(3),
		driver.Uint32(5),
		driver.Uint32(7),
	)

	commandBuffer.CmdDrawIndexedIndirectCount(
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

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks.EasyMockCommandPool(ctrl, device)
	commandBuffer := core1_2.PromoteCommandBuffer(dummies.EasyDummyCommandBuffer(coreDriver, device, commandPool))
	buffer := mocks.EasyMockBuffer(ctrl)
	countBuffer := mocks.EasyMockBuffer(ctrl)

	coreDriver.EXPECT().VkCmdDrawIndirectCount(
		commandBuffer.Handle(),
		buffer.Handle(),
		driver.VkDeviceSize(11),
		countBuffer.Handle(),
		driver.VkDeviceSize(13),
		driver.Uint32(17),
		driver.Uint32(19),
	)

	commandBuffer.CmdDrawIndirectCount(
		buffer,
		uint64(11),
		countBuffer,
		uint64(13),
		17,
		19,
	)
}
