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

func TestVulkanCommandBuffer_CmdPipelineBarrier(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	mockBuffer := mocks.EasyMockBuffer(ctrl)
	mockImage := mocks.EasyMockImage(ctrl)

	mockDriver.EXPECT().VkCmdPipelineBarrier(buffer.Handle(),
		core.VkPipelineStageFlags(0x00010000), // VK_PIPELINE_STAGE_ALL_COMMANDS_BIT
		core.VkPipelineStageFlags(0x00020000), // VK_PIPELINE_STAGE_COMMAND_PREPROCESS_BIT_NV
		core.VkDependencyFlags(2),             // VK_DEPENDENCY_VIEW_LOCAL_BIT
		core.Uint32(2),
		gomock.Not(nil),
		core.Uint32(1),
		gomock.Not(nil),
		core.Uint32(1),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer core.VkCommandBuffer, srcStage, dstStage core.VkPipelineStageFlags, dependencies core.VkDependencyFlags, memoryBarrierCount core.Uint32, pMemoryBarriers *core.VkMemoryBarrier, bufferMemoryBarrierCount core.Uint32, pBufferMemoryBarriers *core.VkBufferMemoryBarrier, imageMemoryBarrierCount core.Uint32, pImageMemoryBarriers *core.VkImageMemoryBarrier) error {
			memoryBarrierSlice := reflect.ValueOf(([]core.VkMemoryBarrier)(unsafe.Slice(pMemoryBarriers, 2)))
			bufferMemoryBarrierSlice := reflect.ValueOf(([]core.VkBufferMemoryBarrier)(unsafe.Slice(pBufferMemoryBarriers, 1)))
			imageMemoryBarrierSlice := reflect.ValueOf(([]core.VkImageMemoryBarrier)(unsafe.Slice(pImageMemoryBarriers, 1)))

			memoryBarrier := memoryBarrierSlice.Index(0)
			require.Equal(t, uint64(46), memoryBarrier.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_BARRIER
			require.True(t, memoryBarrier.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000080), memoryBarrier.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_COLOR_ATTACHMENT_READ_BIT
			require.Equal(t, uint64(0x00100000), memoryBarrier.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_CONDITIONAL_RENDERING_READ_BIT_EXT

			memoryBarrier = memoryBarrierSlice.Index(1)
			require.Equal(t, uint64(46), memoryBarrier.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_BARRIER
			require.True(t, memoryBarrier.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000400), memoryBarrier.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT
			require.Equal(t, uint64(0x01000000), memoryBarrier.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_FRAGMENT_DENSITY_MAP_READ_BIT_EXT

			bufferMemoryBarrier := bufferMemoryBarrierSlice.Index(0)
			require.Equal(t, uint64(44), bufferMemoryBarrier.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER
			require.True(t, memoryBarrier.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00004000), bufferMemoryBarrier.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_HOST_WRITE_BIT
			require.Equal(t, uint64(0x00000040), bufferMemoryBarrier.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_SHADER_WRITE_BIT
			require.Equal(t, uint64(1), bufferMemoryBarrier.FieldByName("srcQueueFamilyIndex").Uint())
			require.Equal(t, uint64(3), bufferMemoryBarrier.FieldByName("dstQueueFamilyIndex").Uint())

			actualBuffer := (core.VkBuffer)(unsafe.Pointer(bufferMemoryBarrier.FieldByName("buffer").Elem().UnsafeAddr()))
			require.Equal(t, mockBuffer.Handle(), actualBuffer)

			require.Equal(t, uint64(5), bufferMemoryBarrier.FieldByName("offset").Uint())
			require.Equal(t, uint64(7), bufferMemoryBarrier.FieldByName("size").Uint())

			imageMemoryBarrier := imageMemoryBarrierSlice.Index(0)
			require.Equal(t, uint64(45), imageMemoryBarrier.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER
			require.True(t, imageMemoryBarrier.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000002), imageMemoryBarrier.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_INDEX_READ_BIT
			require.Equal(t, uint64(0x00020000), imageMemoryBarrier.FieldByName("dstAccessMask").Uint())
			require.Equal(t, uint64(1), imageMemoryBarrier.FieldByName("oldLayout").Uint())
			require.Equal(t, uint64(3), imageMemoryBarrier.FieldByName("newLayout").Uint())
			require.Equal(t, uint64(11), imageMemoryBarrier.FieldByName("srcQueueFamilyIndex").Uint())
			require.Equal(t, uint64(13), imageMemoryBarrier.FieldByName("dstQueueFamilyIndex").Uint())

			actualImage := (core.VkImage)(unsafe.Pointer(imageMemoryBarrier.FieldByName("image").Elem().UnsafeAddr()))
			require.Equal(t, mockImage.Handle(), actualImage)

			subresource := imageMemoryBarrier.FieldByName("subresourceRange")
			require.Equal(t, uint64(0x00000010), subresource.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_PLANE_0_BIT
			require.Equal(t, uint64(17), subresource.FieldByName("baseMipLevel").Uint())
			require.Equal(t, uint64(19), subresource.FieldByName("levelCount").Uint())
			require.Equal(t, uint64(23), subresource.FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(29), subresource.FieldByName("layerCount").Uint())

			return nil
		})

	err := buffer.CmdPipelineBarrier(
		common.PipelineStageAllCommands,
		common.PipelineStageCommandPreprocessNV,
		common.DependencyViewLocal,
		[]*core.MemoryBarrierOptions{
			{
				SrcAccessMask:  common.AccessColorAttachmentRead,
				DestAccessMask: common.AccessConditionalRenderingReadEXT,
			},
			{
				SrcAccessMask:  common.AccessDepthStencilAttachmentWrite,
				DestAccessMask: common.AccessFragmentDensityMapReadEXT,
			},
		},
		[]*core.BufferMemoryBarrierOptions{
			{
				SrcAccessMask:        common.AccessHostWrite,
				DestAccessMask:       common.AccessShaderWrite,
				SrcQueueFamilyIndex:  1,
				DestQueueFamilyIndex: 3,
				Buffer:               mockBuffer,
				Offset:               5,
				Size:                 7,
			},
		},
		[]*core.ImageMemoryBarrierOptions{
			{
				SrcAccessMask:        common.AccessIndexRead,
				DestAccessMask:       common.AccessPreProcessReadNV,
				OldLayout:            common.LayoutGeneral,
				NewLayout:            common.LayoutDepthStencilAttachmentOptimal,
				SrcQueueFamilyIndex:  11,
				DestQueueFamilyIndex: 13,
				Image:                mockImage,
				SubresourceRange: common.ImageSubresourceRange{
					AspectMask:     common.AspectPlane0,
					BaseMipLevel:   17,
					LevelCount:     19,
					BaseArrayLayer: 23,
					LayerCount:     29,
				},
			},
		})
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdCopyBufferToImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	mockBuffer := mocks.EasyMockBuffer(ctrl)
	mockImage := mocks.EasyMockImage(ctrl)

	mockDriver.EXPECT().VkCmdCopyBufferToImage(buffer.Handle(),
		mockBuffer.Handle(),
		mockImage.Handle(),
		core.VkImageLayout(1000241000), // VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL
		core.Uint32(2),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer core.VkCommandBuffer, srcBuffer core.VkBuffer, dstImage core.VkImage, dstImageLayout core.VkImageLayout, regionCount core.Uint32, pRegions *core.VkBufferImageCopy) error {
			regionSlice := reflect.ValueOf(([]core.VkBufferImageCopy)(unsafe.Slice(pRegions, 2)))

			region := regionSlice.Index(0)
			require.Equal(t, uint64(1), region.FieldByName("bufferOffset").Uint())
			require.Equal(t, uint64(3), region.FieldByName("bufferRowLength").Uint())
			require.Equal(t, uint64(5), region.FieldByName("bufferImageHeight").Uint())

			subresource := region.FieldByName("imageSubresource")
			require.Equal(t, uint64(0x00000020), subresource.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_PLANE_1_BIT
			require.Equal(t, uint64(7), subresource.FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(11), subresource.FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(13), subresource.FieldByName("layerCount").Uint())

			offset := region.FieldByName("imageOffset")
			require.Equal(t, int64(17), offset.FieldByName("x").Int())
			require.Equal(t, int64(19), offset.FieldByName("y").Int())
			require.Equal(t, int64(23), offset.FieldByName("z").Int())

			extent := region.FieldByName("imageExtent")
			require.Equal(t, uint64(29), extent.FieldByName("width").Uint())
			require.Equal(t, uint64(31), extent.FieldByName("height").Uint())
			require.Equal(t, uint64(37), extent.FieldByName("depth").Uint())

			region = regionSlice.Index(1)
			require.Equal(t, uint64(41), region.FieldByName("bufferOffset").Uint())
			require.Equal(t, uint64(43), region.FieldByName("bufferRowLength").Uint())
			require.Equal(t, uint64(47), region.FieldByName("bufferImageHeight").Uint())

			subresource = region.FieldByName("imageSubresource")
			require.Equal(t, uint64(0x00000001), subresource.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_COLOR_BIT
			require.Equal(t, uint64(53), subresource.FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(59), subresource.FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(61), subresource.FieldByName("layerCount").Uint())

			offset = region.FieldByName("imageOffset")
			require.Equal(t, int64(67), offset.FieldByName("x").Int())
			require.Equal(t, int64(71), offset.FieldByName("y").Int())
			require.Equal(t, int64(73), offset.FieldByName("z").Int())

			extent = region.FieldByName("imageExtent")
			require.Equal(t, uint64(79), extent.FieldByName("width").Uint())
			require.Equal(t, uint64(83), extent.FieldByName("height").Uint())
			require.Equal(t, uint64(89), extent.FieldByName("depth").Uint())

			return nil
		})

	err := buffer.CmdCopyBufferToImage(mockBuffer, mockImage, common.LayoutDepthAttachmentOptimal, []*core.BufferImageCopy{
		{
			BufferOffset:      1,
			BufferRowLength:   3,
			BufferImageHeight: 5,
			ImageSubresource: common.ImageSubresourceLayers{
				AspectMask:     common.AspectPlane1,
				MipLevel:       7,
				BaseArrayLayer: 11,
				LayerCount:     13,
			},
			ImageOffset: common.Offset3D{
				X: 17,
				Y: 19,
				Z: 23,
			},
			ImageExtent: common.Extent3D{
				Width:  29,
				Height: 31,
				Depth:  37,
			},
		},
		{
			BufferOffset:      41,
			BufferRowLength:   43,
			BufferImageHeight: 47,
			ImageSubresource: common.ImageSubresourceLayers{
				AspectMask:     common.AspectColor,
				MipLevel:       53,
				BaseArrayLayer: 59,
				LayerCount:     61,
			},
			ImageOffset: common.Offset3D{
				X: 67,
				Y: 71,
				Z: 73,
			},
			ImageExtent: common.Extent3D{
				Width:  79,
				Height: 83,
				Depth:  89,
			},
		},
	})
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdBlitImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	sourceImage := mocks.EasyMockImage(ctrl)
	destImage := mocks.EasyMockImage(ctrl)

	mockDriver.EXPECT().VkCmdBlitImage(buffer.Handle(),
		sourceImage.Handle(),
		core.VkImageLayout(1000241000), // VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL
		destImage.Handle(),
		core.VkImageLayout(2), // VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
		core.Uint32(1),
		gomock.Not(nil),
		core.VkFilter(1), // VK_FILTER_LINEAR
	).DoAndReturn(func(commandBuffer core.VkCommandBuffer,
		sourceImage core.VkImage,
		sourceImageLayout core.VkImageLayout,
		destImage core.VkImage,
		destImageLayout core.VkImageLayout,
		regionCount core.Uint32,
		pRegions *core.VkImageBlit,
		filter core.VkFilter) error {

		regionSlice := reflect.ValueOf(([]core.VkImageBlit)(unsafe.Slice(pRegions, 1)))
		region := regionSlice.Index(0)

		srcSubresource := region.FieldByName("srcSubresource")
		require.Equal(t, uint64(8), srcSubresource.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_METADATA_BIT
		require.Equal(t, uint64(1), srcSubresource.FieldByName("mipLevel").Uint())
		require.Equal(t, uint64(3), srcSubresource.FieldByName("baseArrayLayer").Uint())
		require.Equal(t, uint64(5), srcSubresource.FieldByName("layerCount").Uint())

		srcOffsets := region.FieldByName("srcOffsets")
		offset := srcOffsets.Index(0)
		require.Equal(t, int64(7), offset.FieldByName("x").Int())
		require.Equal(t, int64(11), offset.FieldByName("y").Int())
		require.Equal(t, int64(13), offset.FieldByName("z").Int())

		offset = srcOffsets.Index(1)
		require.Equal(t, int64(17), offset.FieldByName("x").Int())
		require.Equal(t, int64(19), offset.FieldByName("y").Int())
		require.Equal(t, int64(23), offset.FieldByName("z").Int())

		dstSubresource := region.FieldByName("dstSubresource")
		require.Equal(t, uint64(4), dstSubresource.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_STENCIL_BIT
		require.Equal(t, uint64(29), dstSubresource.FieldByName("mipLevel").Uint())
		require.Equal(t, uint64(31), dstSubresource.FieldByName("baseArrayLayer").Uint())
		require.Equal(t, uint64(37), dstSubresource.FieldByName("layerCount").Uint())

		dstOffsets := region.FieldByName("dstOffsets")
		offset = dstOffsets.Index(0)
		require.Equal(t, int64(41), offset.FieldByName("x").Int())
		require.Equal(t, int64(43), offset.FieldByName("y").Int())
		require.Equal(t, int64(47), offset.FieldByName("z").Int())

		offset = dstOffsets.Index(1)
		require.Equal(t, int64(53), offset.FieldByName("x").Int())
		require.Equal(t, int64(59), offset.FieldByName("y").Int())
		require.Equal(t, int64(61), offset.FieldByName("z").Int())

		return nil
	})

	err := buffer.CmdBlitImage(sourceImage,
		common.LayoutDepthAttachmentOptimal,
		destImage,
		common.LayoutColorAttachmentOptimal,
		[]*core.ImageBlit{
			{
				SourceSubresource: common.ImageSubresourceLayers{
					AspectMask:     common.AspectMetadata,
					MipLevel:       1,
					BaseArrayLayer: 3,
					LayerCount:     5,
				},
				SourceOffsets: [2]common.Offset3D{
					{
						X: 7,
						Y: 11,
						Z: 13,
					},
					{
						X: 17,
						Y: 19,
						Z: 23,
					},
				},
				DestinationSubresource: common.ImageSubresourceLayers{
					AspectMask:     common.AspectStencil,
					MipLevel:       29,
					BaseArrayLayer: 31,
					LayerCount:     37,
				},
				DestinationOffsets: [2]common.Offset3D{
					{
						X: 41,
						Y: 43,
						Z: 47,
					},
					{
						X: 53,
						Y: 59,
						Z: 61,
					},
				},
			},
		},
		common.FilterLinear,
	)
	require.NoError(t, err)
}
