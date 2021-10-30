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

func setup(t *testing.T, ctrl *gomock.Controller) (*mocks.MockDriver, core.CommandBuffer) {
	mockDriver := mocks.NewMockDriver(ctrl)
	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)

	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockPool := mocks.EasyDummyCommandPool(t, loader, mockDevice)

	buffer := mocks.EasyDummyCommandBuffer(t, mockDevice, mockPool)

	return mockDriver, buffer
}

func setupWithRenderPass(t *testing.T, ctrl *gomock.Controller) (*mocks.MockDriver, core.CommandBuffer, core.RenderPass, core.Framebuffer) {
	mockDriver := mocks.NewMockDriver(ctrl)
	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)

	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockPool := mocks.EasyDummyCommandPool(t, loader, mockDevice)

	buffer := mocks.EasyDummyCommandBuffer(t, mockDevice, mockPool)
	renderPass := mocks.EasyDummyRenderPass(t, loader, mockDevice)
	framebuffer := mocks.EasyDummyFramebuffer(t, loader, mockDevice)

	return mockDriver, buffer, renderPass, framebuffer
}

func TestCommandBuffer_Begin_NoInheritance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkBeginCommandBuffer(buffer.Handle(), gomock.Any()).DoAndReturn(
		func(commandBuffer core.VkCommandBuffer, pBeginInfo *core.VkCommandBufferBeginInfo) (core.VkResult, error) {
			v := reflect.ValueOf(*pBeginInfo)
			require.Equal(t, uint64(42), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(4), v.FieldByName("flags").Uint()) // VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT
			require.True(t, v.FieldByName("pInheritanceInfo").IsNil())

			return core.VKSuccess, nil
		})

	_, err := buffer.Begin(&core.BeginOptions{
		Flags: core.BeginInfoSimultaneousUse,
	})
	require.NoError(t, err)
}

func TestCommandBuffer_Begin_WithInheritance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer, renderPass, framebuffer := setupWithRenderPass(t, ctrl)

	mockDriver.EXPECT().VkBeginCommandBuffer(buffer.Handle(), gomock.Any()).DoAndReturn(
		func(commandBuffer core.VkCommandBuffer, pBeginInfo *core.VkCommandBufferBeginInfo) (core.VkResult, error) {
			v := reflect.ValueOf(*pBeginInfo)
			require.Equal(t, uint64(42), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(4), v.FieldByName("flags").Uint()) // VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT
			require.False(t, v.FieldByName("pInheritanceInfo").IsNil())

			inheritance := v.FieldByName("pInheritanceInfo").Elem()
			require.Equal(t, inheritance.FieldByName("sType").Uint(), uint64(41)) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO
			require.True(t, inheritance.FieldByName("pNext").IsNil())
			require.Equal(t, renderPass.Handle(), (core.VkRenderPass)(unsafe.Pointer(inheritance.FieldByName("renderPass").Elem().UnsafeAddr())))
			require.Equal(t, framebuffer.Handle(), (core.VkFramebuffer)(unsafe.Pointer(inheritance.FieldByName("framebuffer").Elem().UnsafeAddr())))
			require.Equal(t, uint64(3), inheritance.FieldByName("subpass").Uint())
			require.Equal(t, uint64(1), inheritance.FieldByName("occlusionQueryEnable").Uint())
			require.Equal(t, uint64(1), inheritance.FieldByName("queryFlags").Uint())          // VK_QUERY_CONTROL_PRECISE_BIT
			require.Equal(t, uint64(32), inheritance.FieldByName("pipelineStatistics").Uint()) // VK_QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT

			return core.VKSuccess, nil
		})

	_, err := buffer.Begin(&core.BeginOptions{
		Flags: core.BeginInfoSimultaneousUse,
		InheritanceInfo: &core.CommandBufferInheritanceOptions{
			Framebuffer:          framebuffer,
			RenderPass:           renderPass,
			SubPass:              3,
			OcclusionQueryEnable: true,
			QueryFlags:           common.QueryPrecise,
			PipelineStatistics:   common.StatisticClippingInvocations,
		},
	})
	require.NoError(t, err)
}

func TestCommandBuffer_End(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkEndCommandBuffer(buffer.Handle()).Return(core.VKSuccess, nil)

	_, err := buffer.End()
	require.NoError(t, err)
}

func TestCommandBuffer_BeginRenderPass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer, renderPass, framebuffer := setupWithRenderPass(t, ctrl)

	mockDriver.EXPECT().VkCmdBeginRenderPass(buffer.Handle(), gomock.Any(), core.VkSubpassContents(1) /*VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS*/).DoAndReturn(
		func(commandBuffer core.VkCommandBuffer, pRenderPassBegin *core.VkRenderPassBeginInfo, contents core.VkSubpassContents) error {
			v := reflect.ValueOf(*pRenderPassBegin)
			require.Equal(t, uint64(43), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, renderPass.Handle(), (core.VkRenderPass)(unsafe.Pointer(v.FieldByName("renderPass").Elem().UnsafeAddr())))
			require.Equal(t, framebuffer.Handle(), (core.VkFramebuffer)(unsafe.Pointer(v.FieldByName("framebuffer").Elem().UnsafeAddr())))
			require.Equal(t, int64(1), v.FieldByName("renderArea").FieldByName("offset").FieldByName("x").Int())
			require.Equal(t, int64(2), v.FieldByName("renderArea").FieldByName("offset").FieldByName("y").Int())
			require.Equal(t, uint64(30), v.FieldByName("renderArea").FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(50), v.FieldByName("renderArea").FieldByName("extent").FieldByName("height").Uint())
			require.Equal(t, uint64(1), v.FieldByName("clearValueCount").Uint())

			clearValue := (*float32)(unsafe.Pointer(v.FieldByName("pClearValues").Elem().UnsafeAddr()))
			clearValueSlice := ([]float32)(unsafe.Slice(clearValue, 4))

			require.ElementsMatch(t, []float32{5, 6, 7, 8}, clearValueSlice)

			return nil
		})

	err := buffer.CmdBeginRenderPass(core.ContentsSecondaryCommandBuffers, &core.RenderPassBeginOptions{
		RenderPass:  renderPass,
		Framebuffer: framebuffer,
		RenderArea: common.Rect2D{
			Offset: common.Offset2D{X: 1, Y: 2},
			Extent: common.Extent2D{Width: 30, Height: 50},
		},
		ClearValues: []core.ClearValue{core.ClearValueFloat{5, 6, 7, 8}},
	})
	require.NoError(t, err)
}

func TestCommandBuffer_EndRenderPass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkCmdEndRenderPass(buffer.Handle()).Return(nil)

	err := buffer.CmdEndRenderPass()
	require.NoError(t, err)
}

func TestCommandBuffer_CmdBindGraphicsPipeline(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)

	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockPool := mocks.EasyDummyCommandPool(t, loader, mockDevice)
	buffer := mocks.EasyDummyCommandBuffer(t, mockDevice, mockPool)
	pipeline := mocks.EasyDummyGraphicsPipeline(t, loader, mockDevice)

	mockDriver.EXPECT().VkCmdBindPipeline(buffer.Handle(), core.VkPipelineBindPoint(0), pipeline.Handle()).Return(nil)

	err = buffer.CmdBindPipeline(common.BindGraphics, pipeline)
	require.NoError(t, err)
}

func TestCommandBuffer_CmdDraw(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkCmdDraw(buffer.Handle(), core.Uint32(6), core.Uint32(1), core.Uint32(2), core.Uint32(3))

	err := buffer.CmdDraw(6, 1, 2, 3)
	require.NoError(t, err)
}

func TestCommandBuffer_CmdDrawIndexed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkCmdDrawIndexed(buffer.Handle(), core.Uint32(1), core.Uint32(2), core.Uint32(3), core.Int32(4), core.Uint32(5)).Return(nil)

	err := buffer.CmdDrawIndexed(1, 2, 3, 4, 5)
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdBindVertexBuffers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	bufferHandle := mocks.NewFakeBufferHandle()
	vertexBuffer := mocks.NewMockBuffer(ctrl)
	vertexBuffer.EXPECT().Handle().Return(bufferHandle)

	mockDriver.EXPECT().VkCmdBindVertexBuffers(buffer.Handle(), core.Uint32(3), core.Uint32(1), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(commandBuffer core.VkCommandBuffer, firstBinding core.Uint32, bindingCount core.Uint32, pBuffers *core.VkBuffer, pOffsets *core.VkDeviceSize) error {
			singleBuffer := ([]core.VkBuffer)(unsafe.Slice(pBuffers, 1))
			singleOffset := ([]core.VkDeviceSize)(unsafe.Slice(pOffsets, 1))

			require.ElementsMatch(t, []core.VkBuffer{bufferHandle}, singleBuffer)
			require.ElementsMatch(t, []core.VkDeviceSize{2}, singleOffset)
			return nil
		})

	err := buffer.CmdBindVertexBuffers(3, []core.Buffer{vertexBuffer}, []int{2})
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdBindIndexBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	bufferHandle := mocks.NewFakeBufferHandle()
	indexBuffer := mocks.NewMockBuffer(ctrl)
	indexBuffer.EXPECT().Handle().Return(bufferHandle)

	mockDriver.EXPECT().VkCmdBindIndexBuffer(buffer.Handle(), bufferHandle, core.VkDeviceSize(2), core.VkIndexType(1) /* VK_INDEX_TYPE_UINT32*/).Return(nil)

	err := buffer.CmdBindIndexBuffer(indexBuffer, 2, common.IndexUInt32)
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdBindDescriptorSets(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	pipelineLayoutHandle := mocks.NewFakePipelineLayout()
	pipelineLayout := mocks.NewMockPipelineLayout(ctrl)
	pipelineLayout.EXPECT().Handle().Return(pipelineLayoutHandle)

	descriptorSetHandle := mocks.NewFakeDescriptorSet()
	descriptorSet := mocks.NewMockDescriptorSet(ctrl)
	descriptorSet.EXPECT().Handle().Return(descriptorSetHandle)

	mockDriver.EXPECT().VkCmdBindDescriptorSets(
		buffer.Handle(),
		core.VkPipelineBindPoint(1), /* VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR */
		pipelineLayoutHandle,
		core.Uint32(1),
		core.Uint32(1),
		gomock.Not(nil),
		core.Uint32(3),
		gomock.Not(nil)).DoAndReturn(
		func(commandBuffer core.VkCommandBuffer, bind core.VkPipelineBindPoint, pipelineLayout core.VkPipelineLayout, firstSet, descriptorSetCount core.Uint32, pDescriptorSets *core.VkDescriptorSet, dynamicOffsetCount core.Uint32, pDynamicOffsets *core.Uint32) error {
			descriptorSetSlice := ([]core.VkDescriptorSet)(unsafe.Slice(pDescriptorSets, 1))
			dynamicOffsetSlice := ([]core.Uint32)(unsafe.Slice(pDynamicOffsets, 3))

			require.ElementsMatch(t, []core.VkDescriptorSet{descriptorSetHandle}, descriptorSetSlice)
			require.ElementsMatch(t, []core.Uint32{4, 5, 6}, dynamicOffsetSlice)

			return nil
		})

	err := buffer.CmdBindDescriptorSets(common.BindCompute, pipelineLayout, 1, []core.DescriptorSet{
		descriptorSet,
	}, []int{4, 5, 6})
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdCopyBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	src := mocks.EasyMockBuffer(ctrl)
	dest := mocks.EasyMockBuffer(ctrl)

	mockDriver.EXPECT().VkCmdCopyBuffer(buffer.Handle(), src.Handle(), dest.Handle(), core.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(buffer core.VkCommandBuffer, src core.VkBuffer, dest core.VkBuffer, regionCount core.Uint32, pRegions *core.VkBufferCopy) error {
			regionSlice := ([]core.VkBufferCopy)(unsafe.Slice(pRegions, 2))

			regionVal := reflect.ValueOf(regionSlice[0])
			require.Equal(t, uint64(3), regionVal.FieldByName("srcOffset").Uint())
			require.Equal(t, uint64(5), regionVal.FieldByName("dstOffset").Uint())
			require.Equal(t, uint64(7), regionVal.FieldByName("size").Uint())

			regionVal = reflect.ValueOf(regionSlice[1])
			require.Equal(t, uint64(11), regionVal.FieldByName("srcOffset").Uint())
			require.Equal(t, uint64(13), regionVal.FieldByName("dstOffset").Uint())
			require.Equal(t, uint64(17), regionVal.FieldByName("size").Uint())

			return nil
		})

	err := buffer.CmdCopyBuffer(src, dest, []core.BufferCopy{
		{
			SrcOffset: 3,
			DstOffset: 5,
			Size:      7,
		},
		{
			SrcOffset: 11,
			DstOffset: 13,
			Size:      17,
		},
	})
	require.NoError(t, err)
}
