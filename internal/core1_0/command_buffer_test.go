package core1_0_test

import (
	"bytes"
	"encoding/binary"
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/internal/universal"
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

	mockDriver.EXPECT().VkBeginCommandBuffer(mocks.Exactly(buffer.Handle()), gomock.Any()).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, pBeginInfo *driver.VkCommandBufferBeginInfo) (common.VkResult, error) {
			v := reflect.ValueOf(*pBeginInfo)
			require.Equal(t, uint64(42), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(4), v.FieldByName("flags").Uint()) // VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT
			require.True(t, v.FieldByName("pInheritanceInfo").IsNil())

			return common.VKSuccess, nil
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

	mockDriver.EXPECT().VkBeginCommandBuffer(mocks.Exactly(buffer.Handle()), gomock.Any()).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, pBeginInfo *driver.VkCommandBufferBeginInfo) (common.VkResult, error) {
			v := reflect.ValueOf(*pBeginInfo)
			require.Equal(t, uint64(42), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(4), v.FieldByName("flags").Uint()) // VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT
			require.False(t, v.FieldByName("pInheritanceInfo").IsNil())

			inheritance := v.FieldByName("pInheritanceInfo").Elem()
			require.Equal(t, inheritance.FieldByName("sType").Uint(), uint64(41)) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO
			require.True(t, inheritance.FieldByName("pNext").IsNil())
			require.Same(t, renderPass.Handle(), (driver.VkRenderPass)(unsafe.Pointer(inheritance.FieldByName("renderPass").Elem().UnsafeAddr())))
			require.Same(t, framebuffer.Handle(), (driver.VkFramebuffer)(unsafe.Pointer(inheritance.FieldByName("framebuffer").Elem().UnsafeAddr())))
			require.Equal(t, uint64(3), inheritance.FieldByName("subpass").Uint())
			require.Equal(t, uint64(1), inheritance.FieldByName("occlusionQueryEnable").Uint())
			require.Equal(t, uint64(1), inheritance.FieldByName("queryFlags").Uint())          // VK_QUERY_CONTROL_PRECISE_BIT
			require.Equal(t, uint64(32), inheritance.FieldByName("pipelineStatistics").Uint()) // VK_QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT

			return common.VKSuccess, nil
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

	mockDriver.EXPECT().VkEndCommandBuffer(mocks.Exactly(buffer.Handle())).Return(common.VKSuccess, nil)

	_, err := buffer.End()
	require.NoError(t, err)
}

func TestCommandBuffer_BeginRenderPass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer, renderPass, framebuffer := setupWithRenderPass(t, ctrl)

	mockDriver.EXPECT().VkCmdBeginRenderPass(mocks.Exactly(buffer.Handle()), gomock.Any(), driver.VkSubpassContents(1) /*VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS*/).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, pRenderPassBegin *driver.VkRenderPassBeginInfo, contents driver.VkSubpassContents) {
			v := reflect.ValueOf(*pRenderPassBegin)
			require.Equal(t, uint64(43), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Same(t, renderPass.Handle(), (driver.VkRenderPass)(unsafe.Pointer(v.FieldByName("renderPass").Elem().UnsafeAddr())))
			require.Same(t, framebuffer.Handle(), (driver.VkFramebuffer)(unsafe.Pointer(v.FieldByName("framebuffer").Elem().UnsafeAddr())))
			require.Equal(t, int64(1), v.FieldByName("renderArea").FieldByName("offset").FieldByName("x").Int())
			require.Equal(t, int64(2), v.FieldByName("renderArea").FieldByName("offset").FieldByName("y").Int())
			require.Equal(t, uint64(30), v.FieldByName("renderArea").FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(50), v.FieldByName("renderArea").FieldByName("extent").FieldByName("height").Uint())
			require.Equal(t, uint64(1), v.FieldByName("clearValueCount").Uint())

			clearValue := (*float32)(unsafe.Pointer(v.FieldByName("pClearValues").Elem().UnsafeAddr()))
			clearValueSlice := ([]float32)(unsafe.Slice(clearValue, 4))

			require.ElementsMatch(t, []float32{5, 6, 7, 8}, clearValueSlice)
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

	mockDriver.EXPECT().VkCmdEndRenderPass(mocks.Exactly(buffer.Handle()))

	buffer.CmdEndRenderPass()
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

	mockDriver.EXPECT().VkCmdBindPipeline(mocks.Exactly(buffer.Handle()), driver.VkPipelineBindPoint(0), pipeline.Handle())

	buffer.CmdBindPipeline(common.BindGraphics, pipeline)
}

func TestCommandBuffer_CmdDraw(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkCmdDraw(mocks.Exactly(buffer.Handle()), driver.Uint32(6), driver.Uint32(1), driver.Uint32(2), driver.Uint32(3))

	buffer.CmdDraw(6, 1, 2, 3)
}

func TestCommandBuffer_CmdDrawIndexed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkCmdDrawIndexed(mocks.Exactly(buffer.Handle()), driver.Uint32(1), driver.Uint32(2), driver.Uint32(3), driver.Int32(4), driver.Uint32(5))

	buffer.CmdDrawIndexed(1, 2, 3, 4, 5)
}

func TestVulkanCommandBuffer_CmdBindVertexBuffers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	bufferHandle := mocks.NewFakeBufferHandle()
	vertexBuffer := mocks.NewMockBuffer(ctrl)
	vertexBuffer.EXPECT().Handle().Return(bufferHandle)

	mockDriver.EXPECT().VkCmdBindVertexBuffers(mocks.Exactly(buffer.Handle()), driver.Uint32(0), driver.Uint32(1), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, firstBinding driver.Uint32, bindingCount driver.Uint32, pBuffers *driver.VkBuffer, pOffsets *driver.VkDeviceSize) {
			singleBuffer := ([]driver.VkBuffer)(unsafe.Slice(pBuffers, 1))
			singleOffset := ([]driver.VkDeviceSize)(unsafe.Slice(pOffsets, 1))

			require.Same(t, bufferHandle, singleBuffer[0])
			require.ElementsMatch(t, []driver.VkDeviceSize{2}, singleOffset)
		})

	buffer.CmdBindVertexBuffers([]core.Buffer{vertexBuffer}, []int{2})
}

func TestVulkanCommandBuffer_CmdBindIndexBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	bufferHandle := mocks.NewFakeBufferHandle()
	indexBuffer := mocks.NewMockBuffer(ctrl)
	indexBuffer.EXPECT().Handle().Return(bufferHandle)

	mockDriver.EXPECT().VkCmdBindIndexBuffer(mocks.Exactly(buffer.Handle()), mocks.Exactly(bufferHandle), driver.VkDeviceSize(2), driver.VkIndexType(1) /* VK_INDEX_TYPE_UINT32*/)

	buffer.CmdBindIndexBuffer(indexBuffer, 2, common.IndexUInt32)
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
		mocks.Exactly(buffer.Handle()),
		driver.VkPipelineBindPoint(1), /* VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR */
		mocks.Exactly(pipelineLayoutHandle),
		driver.Uint32(0),
		driver.Uint32(1),
		gomock.Not(nil),
		driver.Uint32(3),
		gomock.Not(nil)).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, bind driver.VkPipelineBindPoint, pipelineLayout driver.VkPipelineLayout, firstSet, descriptorSetCount driver.Uint32, pDescriptorSets *driver.VkDescriptorSet, dynamicOffsetCount driver.Uint32, pDynamicOffsets *driver.Uint32) {
			descriptorSetSlice := ([]driver.VkDescriptorSet)(unsafe.Slice(pDescriptorSets, 1))
			dynamicOffsetSlice := ([]driver.Uint32)(unsafe.Slice(pDynamicOffsets, 3))

			require.Same(t, descriptorSetHandle, descriptorSetSlice[0])
			require.ElementsMatch(t, []driver.Uint32{4, 5, 6}, dynamicOffsetSlice)
		})

	buffer.CmdBindDescriptorSets(common.BindCompute, pipelineLayout, []core.DescriptorSet{
		descriptorSet,
	}, []int{4, 5, 6})
}

func TestVulkanCommandBuffer_CmdCopyBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	src := mocks.EasyMockBuffer(ctrl)
	dest := mocks.EasyMockBuffer(ctrl)

	mockDriver.EXPECT().VkCmdCopyBuffer(mocks.Exactly(buffer.Handle()), mocks.Exactly(src.Handle()), mocks.Exactly(dest.Handle()), driver.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(buffer driver.VkCommandBuffer, src driver.VkBuffer, dest driver.VkBuffer, regionCount driver.Uint32, pRegions *driver.VkBufferCopy) {
			regionSlice := ([]driver.VkBufferCopy)(unsafe.Slice(pRegions, 2))

			regionVal := reflect.ValueOf(regionSlice[0])
			require.Equal(t, uint64(3), regionVal.FieldByName("srcOffset").Uint())
			require.Equal(t, uint64(5), regionVal.FieldByName("dstOffset").Uint())
			require.Equal(t, uint64(7), regionVal.FieldByName("size").Uint())

			regionVal = reflect.ValueOf(regionSlice[1])
			require.Equal(t, uint64(11), regionVal.FieldByName("srcOffset").Uint())
			require.Equal(t, uint64(13), regionVal.FieldByName("dstOffset").Uint())
			require.Equal(t, uint64(17), regionVal.FieldByName("size").Uint())
		})

	err := buffer.CmdCopyBuffer(src, dest, []universal.BufferCopy{
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

	mockDriver.EXPECT().VkCmdPipelineBarrier(mocks.Exactly(buffer.Handle()),
		driver.VkPipelineStageFlags(0x00010000), // VK_PIPELINE_STAGE_ALL_COMMANDS_BIT
		driver.VkPipelineStageFlags(0x00020000), // VK_PIPELINE_STAGE_COMMAND_PREPROCESS_BIT_NV
		driver.VkDependencyFlags(2),             // VK_DEPENDENCY_VIEW_LOCAL_BIT
		driver.Uint32(2),
		gomock.Not(nil),
		driver.Uint32(1),
		gomock.Not(nil),
		driver.Uint32(1),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, srcStage, dstStage driver.VkPipelineStageFlags, dependencies driver.VkDependencyFlags, memoryBarrierCount driver.Uint32, pMemoryBarriers *driver.VkMemoryBarrier, bufferMemoryBarrierCount driver.Uint32, pBufferMemoryBarriers *driver.VkBufferMemoryBarrier, imageMemoryBarrierCount driver.Uint32, pImageMemoryBarriers *driver.VkImageMemoryBarrier) {
			memoryBarrierSlice := reflect.ValueOf(([]driver.VkMemoryBarrier)(unsafe.Slice(pMemoryBarriers, 2)))
			bufferMemoryBarrierSlice := reflect.ValueOf(([]driver.VkBufferMemoryBarrier)(unsafe.Slice(pBufferMemoryBarriers, 1)))
			imageMemoryBarrierSlice := reflect.ValueOf(([]driver.VkImageMemoryBarrier)(unsafe.Slice(pImageMemoryBarriers, 1)))

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

			actualBuffer := (driver.VkBuffer)(unsafe.Pointer(bufferMemoryBarrier.FieldByName("buffer").Elem().UnsafeAddr()))
			require.Same(t, mockBuffer.Handle(), actualBuffer)

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

			actualImage := (driver.VkImage)(unsafe.Pointer(imageMemoryBarrier.FieldByName("image").Elem().UnsafeAddr()))
			require.Same(t, mockImage.Handle(), actualImage)

			subresource := imageMemoryBarrier.FieldByName("subresourceRange")
			require.Equal(t, uint64(0x00000010), subresource.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_PLANE_0_BIT
			require.Equal(t, uint64(17), subresource.FieldByName("baseMipLevel").Uint())
			require.Equal(t, uint64(19), subresource.FieldByName("levelCount").Uint())
			require.Equal(t, uint64(23), subresource.FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(29), subresource.FieldByName("layerCount").Uint())
		})

	err := buffer.CmdPipelineBarrier(
		common.PipelineStageAllCommands,
		common.PipelineStageCommandPreprocessNV,
		common.DependencyViewLocal,
		[]*core.MemoryBarrierOptions{
			{
				SrcAccessMask: common.AccessColorAttachmentRead,
				DstAccessMask: common.AccessConditionalRenderingReadEXT,
			},
			{
				SrcAccessMask: common.AccessDepthStencilAttachmentWrite,
				DstAccessMask: common.AccessFragmentDensityMapReadEXT,
			},
		},
		[]*core.BufferMemoryBarrierOptions{
			{
				SrcAccessMask:       common.AccessHostWrite,
				DstAccessMask:       common.AccessShaderWrite,
				SrcQueueFamilyIndex: 1,
				DstQueueFamilyIndex: 3,
				Buffer:              mockBuffer,
				Offset:              5,
				Size:                7,
			},
		},
		[]*core.ImageMemoryBarrierOptions{
			{
				SrcAccessMask:       common.AccessIndexRead,
				DstAccessMask:       common.AccessPreProcessReadNV,
				OldLayout:           common.LayoutGeneral,
				NewLayout:           common.LayoutDepthStencilAttachmentOptimal,
				SrcQueueFamilyIndex: 11,
				DstQueueFamilyIndex: 13,
				Image:               mockImage,
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

	mockDriver.EXPECT().VkCmdCopyBufferToImage(mocks.Exactly(buffer.Handle()),
		mocks.Exactly(mockBuffer.Handle()),
		mocks.Exactly(mockImage.Handle()),
		driver.VkImageLayout(1000241000), // VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL
		driver.Uint32(2),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, srcBuffer driver.VkBuffer, dstImage driver.VkImage, dstImageLayout driver.VkImageLayout, regionCount driver.Uint32, pRegions *driver.VkBufferImageCopy) {
			regionSlice := reflect.ValueOf(([]driver.VkBufferImageCopy)(unsafe.Slice(pRegions, 2)))

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

	mockDriver.EXPECT().VkCmdBlitImage(mocks.Exactly(buffer.Handle()),
		mocks.Exactly(sourceImage.Handle()),
		driver.VkImageLayout(1000241000), // VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL
		mocks.Exactly(destImage.Handle()),
		driver.VkImageLayout(2), // VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
		driver.Uint32(1),
		gomock.Not(nil),
		driver.VkFilter(1), // VK_FILTER_LINEAR
	).DoAndReturn(func(commandBuffer driver.VkCommandBuffer,
		sourceImage driver.VkImage,
		sourceImageLayout driver.VkImageLayout,
		destImage driver.VkImage,
		destImageLayout driver.VkImageLayout,
		regionCount driver.Uint32,
		pRegions *driver.VkImageBlit,
		filter driver.VkFilter) {

		regionSlice := reflect.ValueOf(([]driver.VkImageBlit)(unsafe.Slice(pRegions, 1)))
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

func TestVulkanCommandBuffer_CmdPushConstants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	pipelineLayout := mocks.EasyMockPipelineLayout(ctrl)

	mockDriver.EXPECT().VkCmdPushConstants(mocks.Exactly(buffer.Handle()),
		mocks.Exactly(pipelineLayout.Handle()),
		driver.VkShaderStageFlags(8), // VK_SHADER_STAGE_GEOMETRY_BIT
		driver.Uint32(1),
		driver.Uint32(4),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer,
			pipelineLayout driver.VkPipelineLayout,
			shaderStages driver.VkShaderStageFlags,
			offset driver.Uint32,
			size driver.Uint32,
			valuePtr unsafe.Pointer) {

			bytesPtr := (*byte)(valuePtr)
			bytesSlice := ([]byte)(unsafe.Slice(bytesPtr, 4))

			var intVal uint32
			err := binary.Read(bytes.NewBuffer(bytesSlice), binary.LittleEndian, &intVal)
			require.NoError(t, err)

			require.Equal(t, uint32(5), intVal)
		})

	writer := bytes.NewBuffer(make([]byte, 0, 4))
	err := binary.Write(writer, common.ByteOrder, uint32(5))
	require.NoError(t, err)

	buffer.CmdPushConstants(pipelineLayout, common.StageGeometry, 1, writer.Bytes())
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdSetViewport(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkCmdSetViewport(mocks.Exactly(buffer.Handle()), driver.Uint32(0), driver.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(
			commandBuffer driver.VkCommandBuffer,
			firstViewport driver.Uint32,
			viewportCount driver.Uint32,
			pViewports *driver.VkViewport) {

			viewportSlice := ([]driver.VkViewport)(unsafe.Slice(pViewports, 2))
			val := reflect.ValueOf(viewportSlice)

			viewport := val.Index(0)
			require.InDelta(t, 3, viewport.FieldByName("x").Float(), 0.0001)
			require.InDelta(t, 5, viewport.FieldByName("y").Float(), 0.0001)
			require.InDelta(t, 7, viewport.FieldByName("width").Float(), 0.0001)
			require.InDelta(t, 11, viewport.FieldByName("height").Float(), 0.0001)
			require.InDelta(t, 13, viewport.FieldByName("minDepth").Float(), 0.0001)
			require.InDelta(t, 17, viewport.FieldByName("maxDepth").Float(), 0.0001)

			viewport = val.Index(1)
			require.InDelta(t, 19, viewport.FieldByName("x").Float(), 0.0001)
			require.InDelta(t, 23, viewport.FieldByName("y").Float(), 0.0001)
			require.InDelta(t, 29, viewport.FieldByName("width").Float(), 0.0001)
			require.InDelta(t, 31, viewport.FieldByName("height").Float(), 0.0001)
			require.InDelta(t, 37, viewport.FieldByName("minDepth").Float(), 0.0001)
			require.InDelta(t, 41, viewport.FieldByName("maxDepth").Float(), 0.0001)
		})

	buffer.CmdSetViewport([]common.Viewport{
		{
			X:        3,
			Y:        5,
			Width:    7,
			Height:   11,
			MinDepth: 13,
			MaxDepth: 17,
		},
		{
			X:        19,
			Y:        23,
			Width:    29,
			Height:   31,
			MinDepth: 37,
			MaxDepth: 41,
		},
	})
}

func TestVulkanCommandBuffer_CmdSetScissor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkCmdSetScissor(mocks.Exactly(buffer.Handle()), driver.Uint32(0), driver.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(
			commandBuffer driver.VkCommandBuffer,
			firstScissor driver.Uint32,
			scissorCount driver.Uint32,
			pScissors *driver.VkRect2D) {

			scissorSlice := ([]driver.VkRect2D)(unsafe.Slice(pScissors, 2))
			val := reflect.ValueOf(scissorSlice)

			scissor := val.Index(0)
			offset := scissor.FieldByName("offset")
			require.Equal(t, int64(3), offset.FieldByName("x").Int())
			require.Equal(t, int64(5), offset.FieldByName("y").Int())
			extent := scissor.FieldByName("extent")
			require.Equal(t, uint64(7), extent.FieldByName("width").Uint())
			require.Equal(t, uint64(11), extent.FieldByName("height").Uint())

			scissor = val.Index(1)
			offset = scissor.FieldByName("offset")
			require.Equal(t, int64(13), offset.FieldByName("x").Int())
			require.Equal(t, int64(17), offset.FieldByName("y").Int())
			extent = scissor.FieldByName("extent")
			require.Equal(t, uint64(19), extent.FieldByName("width").Uint())
			require.Equal(t, uint64(23), extent.FieldByName("height").Uint())
		})

	buffer.CmdSetScissor([]common.Rect2D{
		{
			Offset: common.Offset2D{3, 5},
			Extent: common.Extent2D{7, 11},
		},
		{
			Offset: common.Offset2D{13, 17},
			Extent: common.Extent2D{19, 23},
		},
	})
}

func TestVulkanCommandBuffer_CmdCopyImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	srcImage := mocks.EasyMockImage(ctrl)
	dstImage := mocks.EasyMockImage(ctrl)

	mockDriver.EXPECT().VkCmdCopyImage(mocks.Exactly(buffer.Handle()),
		mocks.Exactly(srcImage.Handle()),
		driver.VkImageLayout(1000001002), // VK_IMAGE_LAYOUT_PRESENT_SRC_KHR
		mocks.Exactly(dstImage.Handle()),
		driver.VkImageLayout(5), // VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL
		driver.Uint32(2),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, srcImage driver.VkImage, srcImageLayout driver.VkImageLayout, dstImage driver.VkImage, dstImageLayout driver.VkImageLayout, regionCount driver.Uint32, pRegions *driver.VkImageCopy) {
			regionSlice := ([]driver.VkImageCopy)(unsafe.Slice(pRegions, 2))
			val := reflect.ValueOf(regionSlice)

			region := val.Index(0)
			srcSubresource := region.FieldByName("srcSubresource")
			dstSubresource := region.FieldByName("dstSubresource")
			srcOffset := region.FieldByName("srcOffset")
			dstOffset := region.FieldByName("dstOffset")
			extent := region.FieldByName("extent")

			require.Equal(t, uint64(8), srcSubresource.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_METADATA_BIT
			require.Equal(t, uint64(1), srcSubresource.FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(3), srcSubresource.FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(5), srcSubresource.FieldByName("layerCount").Uint())

			require.Equal(t, uint64(0x00000200), dstSubresource.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_MEMORY_PLANE_2_BIT_EXT
			require.Equal(t, uint64(7), dstSubresource.FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(11), dstSubresource.FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(13), dstSubresource.FieldByName("layerCount").Uint())

			require.Equal(t, int64(17), srcOffset.FieldByName("x").Int())
			require.Equal(t, int64(19), srcOffset.FieldByName("y").Int())
			require.Equal(t, int64(23), srcOffset.FieldByName("z").Int())

			require.Equal(t, int64(29), dstOffset.FieldByName("x").Int())
			require.Equal(t, int64(31), dstOffset.FieldByName("y").Int())
			require.Equal(t, int64(37), dstOffset.FieldByName("z").Int())

			require.Equal(t, uint64(41), extent.FieldByName("width").Uint())
			require.Equal(t, uint64(43), extent.FieldByName("height").Uint())
			require.Equal(t, uint64(47), extent.FieldByName("depth").Uint())

			region = val.Index(1)
			srcSubresource = region.FieldByName("srcSubresource")
			dstSubresource = region.FieldByName("dstSubresource")
			srcOffset = region.FieldByName("srcOffset")
			dstOffset = region.FieldByName("dstOffset")
			extent = region.FieldByName("extent")

			require.Equal(t, uint64(1), srcSubresource.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_COLOR_BIT
			require.Equal(t, uint64(53), srcSubresource.FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(59), srcSubresource.FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(61), srcSubresource.FieldByName("layerCount").Uint())

			require.Equal(t, uint64(2), dstSubresource.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_DEPTH_BIT
			require.Equal(t, uint64(67), dstSubresource.FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(71), dstSubresource.FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(73), dstSubresource.FieldByName("layerCount").Uint())

			require.Equal(t, int64(79), srcOffset.FieldByName("x").Int())
			require.Equal(t, int64(83), srcOffset.FieldByName("y").Int())
			require.Equal(t, int64(89), srcOffset.FieldByName("z").Int())

			require.Equal(t, int64(97), dstOffset.FieldByName("x").Int())
			require.Equal(t, int64(101), dstOffset.FieldByName("y").Int())
			require.Equal(t, int64(103), dstOffset.FieldByName("z").Int())

			require.Equal(t, uint64(107), extent.FieldByName("width").Uint())
			require.Equal(t, uint64(109), extent.FieldByName("height").Uint())
			require.Equal(t, uint64(113), extent.FieldByName("depth").Uint())
		})

	err := buffer.CmdCopyImage(srcImage, common.LayoutPresentSrcKHR, dstImage, common.LayoutShaderReadOnlyOptimal, []universal.ImageCopy{
		{
			SrcSubresource: common.ImageSubresourceLayers{
				AspectMask:     common.AspectMetadata,
				MipLevel:       1,
				BaseArrayLayer: 3,
				LayerCount:     5,
			},
			DstSubresource: common.ImageSubresourceLayers{
				AspectMask:     common.AspectMemoryPlane2EXT,
				MipLevel:       7,
				BaseArrayLayer: 11,
				LayerCount:     13,
			},
			SrcOffset: common.Offset3D{17, 19, 23},
			DstOffset: common.Offset3D{29, 31, 37},
			Extent:    common.Extent3D{41, 43, 47},
		},
		{
			SrcSubresource: common.ImageSubresourceLayers{
				AspectMask:     common.AspectColor,
				MipLevel:       53,
				BaseArrayLayer: 59,
				LayerCount:     61,
			},
			DstSubresource: common.ImageSubresourceLayers{
				AspectMask:     common.AspectDepth,
				MipLevel:       67,
				BaseArrayLayer: 71,
				LayerCount:     73,
			},
			SrcOffset: common.Offset3D{79, 83, 89},
			DstOffset: common.Offset3D{97, 101, 103},
			Extent:    common.Extent3D{107, 109, 113},
		},
	})
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdNextSubpass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkCmdNextSubpass(mocks.Exactly(buffer.Handle()),
		driver.VkSubpassContents(1), /* VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS */
	)

	buffer.CmdNextSubpass(core.ContentsSecondaryCommandBuffers)
}

func TestVulkanCommandBuffer_CmdWaitEvents(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	mockBuffer := mocks.EasyMockBuffer(ctrl)
	mockImage := mocks.EasyMockImage(ctrl)
	event1 := mocks.EasyMockEvent(ctrl)
	event2 := mocks.EasyMockEvent(ctrl)

	mockDriver.EXPECT().VkCmdWaitEvents(mocks.Exactly(buffer.Handle()),
		driver.Uint32(2),
		gomock.Not(nil),
		driver.VkPipelineStageFlags(0x00010000), // VK_PIPELINE_STAGE_ALL_COMMANDS_BIT
		driver.VkPipelineStageFlags(0x00020000), // VK_PIPELINE_STAGE_COMMAND_PREPROCESS_BIT_NV
		driver.Uint32(2),
		gomock.Not(nil),
		driver.Uint32(1),
		gomock.Not(nil),
		driver.Uint32(1),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, eventCount driver.Uint32, pEvents *driver.VkEvent, srcStage, dstStage driver.VkPipelineStageFlags, memoryBarrierCount driver.Uint32, pMemoryBarriers *driver.VkMemoryBarrier, bufferMemoryBarrierCount driver.Uint32, pBufferMemoryBarriers *driver.VkBufferMemoryBarrier, imageMemoryBarrierCount driver.Uint32, pImageMemoryBarriers *driver.VkImageMemoryBarrier) {
			eventSlice := ([]driver.VkEvent)(unsafe.Slice(pEvents, 2))
			memoryBarrierSlice := reflect.ValueOf(([]driver.VkMemoryBarrier)(unsafe.Slice(pMemoryBarriers, 2)))
			bufferMemoryBarrierSlice := reflect.ValueOf(([]driver.VkBufferMemoryBarrier)(unsafe.Slice(pBufferMemoryBarriers, 1)))
			imageMemoryBarrierSlice := reflect.ValueOf(([]driver.VkImageMemoryBarrier)(unsafe.Slice(pImageMemoryBarriers, 1)))

			require.Same(t, event1.Handle(), eventSlice[0])
			require.Same(t, event2.Handle(), eventSlice[1])

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

			actualBuffer := (driver.VkBuffer)(unsafe.Pointer(bufferMemoryBarrier.FieldByName("buffer").Elem().UnsafeAddr()))
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

			actualImage := (driver.VkImage)(unsafe.Pointer(imageMemoryBarrier.FieldByName("image").Elem().UnsafeAddr()))
			require.Same(t, mockImage.Handle(), actualImage)

			subresource := imageMemoryBarrier.FieldByName("subresourceRange")
			require.Equal(t, uint64(0x00000010), subresource.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_PLANE_0_BIT
			require.Equal(t, uint64(17), subresource.FieldByName("baseMipLevel").Uint())
			require.Equal(t, uint64(19), subresource.FieldByName("levelCount").Uint())
			require.Equal(t, uint64(23), subresource.FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(29), subresource.FieldByName("layerCount").Uint())
		})

	err := buffer.CmdWaitEvents(
		[]core.Event{event1, event2},
		common.PipelineStageAllCommands,
		common.PipelineStageCommandPreprocessNV,
		[]*core.MemoryBarrierOptions{
			{
				SrcAccessMask: common.AccessColorAttachmentRead,
				DstAccessMask: common.AccessConditionalRenderingReadEXT,
			},
			{
				SrcAccessMask: common.AccessDepthStencilAttachmentWrite,
				DstAccessMask: common.AccessFragmentDensityMapReadEXT,
			},
		},
		[]*core.BufferMemoryBarrierOptions{
			{
				SrcAccessMask:       common.AccessHostWrite,
				DstAccessMask:       common.AccessShaderWrite,
				SrcQueueFamilyIndex: 1,
				DstQueueFamilyIndex: 3,
				Buffer:              mockBuffer,
				Offset:              5,
				Size:                7,
			},
		},
		[]*core.ImageMemoryBarrierOptions{
			{
				SrcAccessMask:       common.AccessIndexRead,
				DstAccessMask:       common.AccessPreProcessReadNV,
				OldLayout:           common.LayoutGeneral,
				NewLayout:           common.LayoutDepthStencilAttachmentOptimal,
				SrcQueueFamilyIndex: 11,
				DstQueueFamilyIndex: 13,
				Image:               mockImage,
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

func TestVulkanCommandBuffer_CmdSetEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	event := mocks.EasyMockEvent(ctrl)

	mockDriver.EXPECT().VkCmdSetEvent(mocks.Exactly(buffer.Handle()), mocks.Exactly(event.Handle()), driver.VkPipelineStageFlags(0x80) /*VK_PIPELINE_STAGE_FRAGMENT_SHADER_BIT*/)

	buffer.CmdSetEvent(event, common.PipelineStageFragmentShader)
}

func TestVulkanCommandBuffer_CmdClearColorImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	image := mocks.EasyMockImage(ctrl)

	mockDriver.EXPECT().VkCmdClearColorImage(mocks.Exactly(buffer.Handle()),
		mocks.Exactly(image.Handle()),
		driver.VkImageLayout(3), // VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL
		gomock.Not(nil),
		driver.Uint32(2),
		gomock.Not(nil),
	).DoAndReturn(
		func(buffer driver.VkCommandBuffer, image driver.VkImage, imageLayout driver.VkImageLayout, pColor *driver.VkClearColorValue, rangeCount driver.Uint32, pRanges *driver.VkImageSubresourceRange) {
			colorFloat := unsafe.Slice((*float32)(unsafe.Pointer(pColor)), 4)
			require.InDelta(t, 0.2, colorFloat[0], 0.0001)
			require.InDelta(t, 0.3, colorFloat[1], 0.0001)
			require.InDelta(t, 0.4, colorFloat[2], 0.0001)
			require.InDelta(t, 0.5, colorFloat[3], 0.0001)

			rangeSlice := reflect.ValueOf(([]driver.VkImageSubresourceRange)(unsafe.Slice(pRanges, 2)))
			r := rangeSlice.Index(0)

			require.Equal(t, uint64(8), r.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_METADATA_BIT
			require.Equal(t, uint64(1), r.FieldByName("baseMipLevel").Uint())
			require.Equal(t, uint64(3), r.FieldByName("levelCount").Uint())
			require.Equal(t, uint64(5), r.FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(7), r.FieldByName("layerCount").Uint())

			r = rangeSlice.Index(1)
			require.Equal(t, uint64(2), r.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_DEPTH_BIT
			require.Equal(t, uint64(11), r.FieldByName("baseMipLevel").Uint())
			require.Equal(t, uint64(13), r.FieldByName("levelCount").Uint())
			require.Equal(t, uint64(17), r.FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(19), r.FieldByName("layerCount").Uint())
		})

	buffer.CmdClearColorImage(image, common.LayoutDepthStencilAttachmentOptimal, &core.ClearValueFloat{0.2, 0.3, 0.4, 0.5}, []common.ImageSubresourceRange{
		{
			AspectMask:     common.AspectMetadata,
			BaseMipLevel:   1,
			LevelCount:     3,
			BaseArrayLayer: 5,
			LayerCount:     7,
		},
		{
			AspectMask:     common.AspectDepth,
			BaseMipLevel:   11,
			LevelCount:     13,
			BaseArrayLayer: 17,
			LayerCount:     19,
		},
	})
}

func TestVulkanCommandBuffer_Reset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkResetCommandBuffer(mocks.Exactly(buffer.Handle()), driver.VkCommandBufferResetFlags(1)).Return(common.VKSuccess, nil)

	_, err := buffer.Reset(ResetReleaseResources)
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdResetQueryPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	queryPool := mocks.EasyMockQueryPool(ctrl)

	mockDriver.EXPECT().VkCmdResetQueryPool(mocks.Exactly(buffer.Handle()), mocks.Exactly(queryPool.Handle()), driver.Uint32(1), driver.Uint32(3))

	buffer.CmdResetQueryPool(queryPool, 1, 3)
}

func TestVulkanCommandBuffer_CmdBeginQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	queryPool := mocks.EasyMockQueryPool(ctrl)

	mockDriver.EXPECT().VkCmdBeginQuery(
		mocks.Exactly(buffer.Handle()),
		mocks.Exactly(queryPool.Handle()),
		driver.Uint32(5),
		driver.VkQueryControlFlags(1), // VK_QUERY_CONTROL_PRECISE_BIT
	)

	buffer.CmdBeginQuery(queryPool, 5, common.QueryPrecise)
}

func TestVulkanCommandBuffer_CmdEndQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	queryPool := mocks.EasyMockQueryPool(ctrl)

	mockDriver.EXPECT().VkCmdEndQuery(
		mocks.Exactly(buffer.Handle()),
		mocks.Exactly(queryPool.Handle()),
		driver.Uint32(5),
	)

	buffer.CmdEndQuery(queryPool, 5)
}

func TestVulkanCommandBuffer_CmdCopyQueryPoolResults(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	queryPool := mocks.EasyMockQueryPool(ctrl)
	dstBuffer := mocks.EasyMockBuffer(ctrl)

	mockDriver.EXPECT().VkCmdCopyQueryPoolResults(
		mocks.Exactly(buffer.Handle()),
		mocks.Exactly(queryPool.Handle()),
		driver.Uint32(1),
		driver.Uint32(3),
		mocks.Exactly(dstBuffer.Handle()),
		driver.VkDeviceSize(5),
		driver.VkDeviceSize(7),
		driver.VkQueryResultFlags(8), // VK_QUERY_RESULT_PARTIAL_BIT
	)

	buffer.CmdCopyQueryPoolResults(queryPool, 1, 3, dstBuffer, 5, 7, common.QueryResultPartial)
}

func TestVulkanCommandBuffer_CmdExecuteCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	commandBuffers := []core.CommandBuffer{
		mocks.EasyMockCommandBuffer(ctrl),
		mocks.EasyMockCommandBuffer(ctrl),
		mocks.EasyMockCommandBuffer(ctrl),
	}

	mockDriver.EXPECT().VkCmdExecuteCommands(
		mocks.Exactly(buffer.Handle()),
		driver.Uint32(3),
		gomock.Not(nil)).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, secondaryCount driver.Uint32, pSecondaryBuffers *driver.VkCommandBuffer) (common.VkResult, error) {
			secondaryBufferSlice := ([]driver.VkCommandBuffer)(unsafe.Slice(pSecondaryBuffers, 3))
			require.Same(t, commandBuffers[0].Handle(), secondaryBufferSlice[0])
			require.Same(t, commandBuffers[1].Handle(), secondaryBufferSlice[1])
			require.Same(t, commandBuffers[2].Handle(), secondaryBufferSlice[2])

			return common.VKSuccess, nil
		})

	buffer.CmdExecuteCommands(commandBuffers)
}

func TestVulkanCommandBuffer_CmdClearAttachments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkCmdClearAttachments(
		mocks.Exactly(buffer.Handle()),
		driver.Uint32(1),
		gomock.Not(nil),
		driver.Uint32(2),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, attachmentCount driver.Uint32, pAttachments *driver.VkClearAttachment, rectCount driver.Uint32, pRects *driver.VkClearRect) {
			attachmentSlice := ([]driver.VkClearAttachment)(unsafe.Slice(pAttachments, 1))
			rectSlice := ([]driver.VkClearRect)(unsafe.Slice(pRects, 2))

			val := reflect.ValueOf(attachmentSlice).Index(0)
			require.Equal(t, uint64(1), val.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_COLOR_BIT
			require.Equal(t, uint64(3), val.FieldByName("colorAttachment").Uint())
			floatClear := (*float32)(unsafe.Pointer(val.FieldByName("clearValue").UnsafeAddr()))
			floatSlice := ([]float32)(unsafe.Slice(floatClear, 4))
			require.InDelta(t, 5, floatSlice[0], 0.001)
			require.InDelta(t, 7, floatSlice[1], 0.001)
			require.InDelta(t, 11, floatSlice[2], 0.001)
			require.InDelta(t, 13, floatSlice[3], 0.001)

			val = reflect.ValueOf(rectSlice[0])
			require.Equal(t, uint64(17), val.FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(19), val.FieldByName("layerCount").Uint())
			require.Equal(t, int64(23), val.FieldByName("rect").FieldByName("offset").FieldByName("x").Int())
			require.Equal(t, int64(29), val.FieldByName("rect").FieldByName("offset").FieldByName("y").Int())
			require.Equal(t, uint64(31), val.FieldByName("rect").FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(37), val.FieldByName("rect").FieldByName("extent").FieldByName("height").Uint())

			val = reflect.ValueOf(rectSlice[1])
			require.Equal(t, uint64(41), val.FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(43), val.FieldByName("layerCount").Uint())
			require.Equal(t, int64(47), val.FieldByName("rect").FieldByName("offset").FieldByName("x").Int())
			require.Equal(t, int64(53), val.FieldByName("rect").FieldByName("offset").FieldByName("y").Int())
			require.Equal(t, uint64(59), val.FieldByName("rect").FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(61), val.FieldByName("rect").FieldByName("extent").FieldByName("height").Uint())
		})

	buffer.CmdClearAttachments([]ClearAttachment{
		{
			AspectMask:      common.AspectColor,
			ColorAttachment: 3,
			ClearValue:      core.ClearValueFloat{5, 7, 11, 13},
		},
	}, []ClearRect{
		{
			BaseArrayLayer: 17,
			LayerCount:     19,
			Rect: common.Rect2D{
				Offset: common.Offset2D{23, 29},
				Extent: common.Extent2D{31, 37},
			},
		},
		{
			BaseArrayLayer: 41,
			LayerCount:     43,
			Rect: common.Rect2D{
				Offset: common.Offset2D{47, 53},
				Extent: common.Extent2D{59, 61},
			},
		},
	})
}

func TestVulkanCommandBuffer_CmdClearDepthStencilImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	image := mocks.EasyMockImage(ctrl)

	mockDriver.EXPECT().VkCmdClearDepthStencilImage(mocks.Exactly(buffer.Handle()), mocks.Exactly(image.Handle()),
		driver.VkImageLayout(5), // VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL
		gomock.Not(nil),
		driver.Uint32(2),
		gomock.Not(nil),
	).DoAndReturn(func(commandBuffer driver.VkCommandBuffer, image driver.VkImage, imageLayout driver.VkImageLayout, pDepthStencil *driver.VkClearDepthStencilValue, rangeCount driver.Uint32, pRanges *driver.VkImageSubresourceRange) {
		depthStencil := reflect.ValueOf(pDepthStencil).Elem()

		require.InDelta(t, 0.5, depthStencil.FieldByName("depth").Float(), 0.00001)
		require.Equal(t, uint64(3), depthStencil.FieldByName("stencil").Uint())

		rangeSlice := reflect.ValueOf(([]driver.VkImageSubresourceRange)(unsafe.Slice(pRanges, 2)))

		val := rangeSlice.Index(0)
		require.Equal(t, uint64(1), val.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_COLOR_BIT
		require.Equal(t, uint64(5), val.FieldByName("baseMipLevel").Uint())
		require.Equal(t, uint64(7), val.FieldByName("levelCount").Uint())
		require.Equal(t, uint64(11), val.FieldByName("baseArrayLayer").Uint())
		require.Equal(t, uint64(13), val.FieldByName("layerCount").Uint())

		val = rangeSlice.Index(1)
		require.Equal(t, uint64(2), val.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_DEPTH_BIT
		require.Equal(t, uint64(17), val.FieldByName("baseMipLevel").Uint())
		require.Equal(t, uint64(19), val.FieldByName("levelCount").Uint())
		require.Equal(t, uint64(23), val.FieldByName("baseArrayLayer").Uint())
		require.Equal(t, uint64(29), val.FieldByName("layerCount").Uint())
	})

	buffer.CmdClearDepthStencilImage(image, common.LayoutShaderReadOnlyOptimal, &core.ClearValueDepthStencil{
		Depth:   0.5,
		Stencil: 3,
	}, []common.ImageSubresourceRange{
		{
			AspectMask:     common.AspectColor,
			BaseMipLevel:   5,
			LevelCount:     7,
			BaseArrayLayer: 11,
			LayerCount:     13,
		},
		{
			AspectMask:     common.AspectDepth,
			BaseMipLevel:   17,
			LevelCount:     19,
			BaseArrayLayer: 23,
			LayerCount:     29,
		},
	})
}

func TestVulkanCommandBuffer_CmdCopyImageToBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	image := mocks.EasyMockImage(ctrl)
	dstBuffer := mocks.EasyMockBuffer(ctrl)

	mockDriver.EXPECT().VkCmdCopyImageToBuffer(mocks.Exactly(buffer.Handle()), mocks.Exactly(image.Handle()),
		driver.VkImageLayout(1000241000), // VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL
		mocks.Exactly(dstBuffer.Handle()),
		driver.Uint32(1),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, srcImage driver.VkImage, srcImageLayout driver.VkImageLayout, dstBuffer driver.VkBuffer, regionCount driver.Uint32, pRegions *driver.VkBufferImageCopy) {
			regionSlice := ([]driver.VkBufferImageCopy)(unsafe.Slice(pRegions, 1))
			val := reflect.ValueOf(regionSlice)
			val = val.Index(0)

			require.Equal(t, uint64(1), val.FieldByName("bufferOffset").Uint())
			require.Equal(t, uint64(3), val.FieldByName("bufferRowLength").Uint())
			require.Equal(t, uint64(5), val.FieldByName("bufferImageHeight").Uint())
			require.Equal(t, uint64(1), val.FieldByName("imageSubresource").FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_COLOR_BIT
			require.Equal(t, uint64(7), val.FieldByName("imageSubresource").FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(11), val.FieldByName("imageSubresource").FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(13), val.FieldByName("imageSubresource").FieldByName("layerCount").Uint())
			require.Equal(t, int64(17), val.FieldByName("imageOffset").FieldByName("x").Int())
			require.Equal(t, int64(19), val.FieldByName("imageOffset").FieldByName("y").Int())
			require.Equal(t, int64(23), val.FieldByName("imageOffset").FieldByName("z").Int())
			require.Equal(t, uint64(29), val.FieldByName("imageExtent").FieldByName("width").Uint())
			require.Equal(t, uint64(31), val.FieldByName("imageExtent").FieldByName("height").Uint())
			require.Equal(t, uint64(37), val.FieldByName("imageExtent").FieldByName("depth").Uint())
		})

	err := buffer.CmdCopyImageToBuffer(image, common.LayoutDepthAttachmentOptimal, dstBuffer, []core.BufferImageCopy{
		{
			BufferOffset:      1,
			BufferRowLength:   3,
			BufferImageHeight: 5,
			ImageSubresource: common.ImageSubresourceLayers{
				AspectMask:     common.AspectColor,
				MipLevel:       7,
				BaseArrayLayer: 11,
				LayerCount:     13,
			},
			ImageOffset: common.Offset3D{17, 19, 23},
			ImageExtent: common.Extent3D{29, 31, 37},
		},
	})
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdDispatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkCmdDispatch(mocks.Exactly(buffer.Handle()), driver.Uint32(1), driver.Uint32(3), driver.Uint32(5))

	buffer.CmdDispatch(1, 3, 5)
}

func TestVulkanCommandBuffer_CmdDispatchIndirect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	indirectBuffer := mocks.EasyMockBuffer(ctrl)

	mockDriver.EXPECT().VkCmdDispatchIndirect(mocks.Exactly(buffer.Handle()), mocks.Exactly(indirectBuffer.Handle()), driver.VkDeviceSize(3))

	buffer.CmdDispatchIndirect(indirectBuffer, 3)
}

func TestVulkanCommandBuffer_CmdDrawIndexedIndirect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	indirectBuffer := mocks.EasyMockBuffer(ctrl)

	mockDriver.EXPECT().VkCmdDrawIndexedIndirect(mocks.Exactly(buffer.Handle()), mocks.Exactly(indirectBuffer.Handle()), driver.VkDeviceSize(3), driver.Uint32(5), driver.Uint32(7))

	buffer.CmdDrawIndexedIndirect(indirectBuffer, 3, 5, 7)
}

func TestVulkanCommandBuffer_CmdDrawIndirect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	indirectBuffer := mocks.EasyMockBuffer(ctrl)

	mockDriver.EXPECT().VkCmdDrawIndirect(mocks.Exactly(buffer.Handle()), mocks.Exactly(indirectBuffer.Handle()), driver.VkDeviceSize(3), driver.Uint32(5), driver.Uint32(7))

	buffer.CmdDrawIndirect(indirectBuffer, 3, 5, 7)
}

func TestVulkanCommandBuffer_CmdFillBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	fillBuffer := mocks.EasyMockBuffer(ctrl)

	mockDriver.EXPECT().VkCmdFillBuffer(mocks.Exactly(buffer.Handle()), mocks.Exactly(fillBuffer.Handle()), driver.VkDeviceSize(1), driver.VkDeviceSize(3), driver.Uint32(5))

	buffer.CmdFillBuffer(fillBuffer, 1, 3, 5)
}

func TestVulkanCommandBuffer_CmdResetEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	event := mocks.EasyMockEvent(ctrl)

	mockDriver.EXPECT().VkCmdResetEvent(
		mocks.Exactly(buffer.Handle()), mocks.Exactly(event.Handle()),
		driver.VkPipelineStageFlags(0x40000), // VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT
	)

	buffer.CmdResetEvent(event, common.PipelineStageConditionalRenderingEXT)
}

func TestVulkanCommandBuffer_CmdResolveImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	srcImage := mocks.EasyMockImage(ctrl)
	dstImage := mocks.EasyMockImage(ctrl)

	mockDriver.EXPECT().VkCmdResolveImage(
		mocks.Exactly(buffer.Handle()),
		mocks.Exactly(srcImage.Handle()),
		driver.VkImageLayout(5), // VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL
		mocks.Exactly(dstImage.Handle()),
		driver.VkImageLayout(1000001002), // VK_IMAGE_LAYOUT_PRESENT_SRC_KHR
		driver.Uint32(2),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, srcImage driver.VkImage, srcImageLayout driver.VkImageLayout, dstImage driver.VkImage, dstImageLayout driver.VkImageLayout, resolveCount driver.Uint32, pResolves *driver.VkImageResolve) {
			resolveSlice := ([]driver.VkImageResolve)(unsafe.Slice(pResolves, 2))
			sliceVal := reflect.ValueOf(resolveSlice)

			val := sliceVal.Index(0)
			require.Equal(t, uint64(1), val.FieldByName("srcSubresource").FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_COLOR_BIT
			require.Equal(t, uint64(1), val.FieldByName("srcSubresource").FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(3), val.FieldByName("srcSubresource").FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(5), val.FieldByName("srcSubresource").FieldByName("layerCount").Uint())

			require.Equal(t, int64(7), val.FieldByName("srcOffset").FieldByName("x").Int())
			require.Equal(t, int64(11), val.FieldByName("srcOffset").FieldByName("y").Int())
			require.Equal(t, int64(13), val.FieldByName("srcOffset").FieldByName("z").Int())

			require.Equal(t, uint64(2), val.FieldByName("dstSubresource").FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_DEPTH_BIT
			require.Equal(t, uint64(17), val.FieldByName("dstSubresource").FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(19), val.FieldByName("dstSubresource").FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(23), val.FieldByName("dstSubresource").FieldByName("layerCount").Uint())

			require.Equal(t, int64(29), val.FieldByName("dstOffset").FieldByName("x").Int())
			require.Equal(t, int64(31), val.FieldByName("dstOffset").FieldByName("y").Int())
			require.Equal(t, int64(37), val.FieldByName("dstOffset").FieldByName("z").Int())

			require.Equal(t, uint64(41), val.FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(43), val.FieldByName("extent").FieldByName("height").Uint())
			require.Equal(t, uint64(47), val.FieldByName("extent").FieldByName("depth").Uint())

			val = sliceVal.Index(1)
			require.Equal(t, uint64(8), val.FieldByName("srcSubresource").FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_METADATA_BIT
			require.Equal(t, uint64(53), val.FieldByName("srcSubresource").FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(59), val.FieldByName("srcSubresource").FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(61), val.FieldByName("srcSubresource").FieldByName("layerCount").Uint())

			require.Equal(t, int64(67), val.FieldByName("srcOffset").FieldByName("x").Int())
			require.Equal(t, int64(71), val.FieldByName("srcOffset").FieldByName("y").Int())
			require.Equal(t, int64(73), val.FieldByName("srcOffset").FieldByName("z").Int())

			require.Equal(t, uint64(4), val.FieldByName("dstSubresource").FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_STENCIL_BIT
			require.Equal(t, uint64(79), val.FieldByName("dstSubresource").FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(83), val.FieldByName("dstSubresource").FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(89), val.FieldByName("dstSubresource").FieldByName("layerCount").Uint())

			require.Equal(t, int64(97), val.FieldByName("dstOffset").FieldByName("x").Int())
			require.Equal(t, int64(101), val.FieldByName("dstOffset").FieldByName("y").Int())
			require.Equal(t, int64(103), val.FieldByName("dstOffset").FieldByName("z").Int())

			require.Equal(t, uint64(107), val.FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(109), val.FieldByName("extent").FieldByName("height").Uint())
			require.Equal(t, uint64(113), val.FieldByName("extent").FieldByName("depth").Uint())
		})

	buffer.CmdResolveImage(srcImage,
		common.LayoutShaderReadOnlyOptimal,
		dstImage,
		common.LayoutPresentSrcKHR,
		[]ImageResolve{
			{
				SrcSubresource: common.ImageSubresourceLayers{
					AspectMask:     common.AspectColor,
					MipLevel:       1,
					BaseArrayLayer: 3,
					LayerCount:     5,
				},
				SrcOffset: common.Offset3D{7, 11, 13},
				DstSubresource: common.ImageSubresourceLayers{
					AspectMask:     common.AspectDepth,
					MipLevel:       17,
					BaseArrayLayer: 19,
					LayerCount:     23,
				},
				DstOffset: common.Offset3D{29, 31, 37},
				Extent:    common.Extent3D{41, 43, 47},
			},
			{
				SrcSubresource: common.ImageSubresourceLayers{
					AspectMask:     common.AspectMetadata,
					MipLevel:       53,
					BaseArrayLayer: 59,
					LayerCount:     61,
				},
				SrcOffset: common.Offset3D{67, 71, 73},
				DstSubresource: common.ImageSubresourceLayers{
					AspectMask:     common.AspectStencil,
					MipLevel:       79,
					BaseArrayLayer: 83,
					LayerCount:     89,
				},
				DstOffset: common.Offset3D{97, 101, 103},
				Extent:    common.Extent3D{107, 109, 113},
			},
		})
}

func TestVulkanCommandBuffer_CmdSetBlendConstants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkCmdSetBlendConstants(mocks.Exactly(buffer.Handle()), gomock.Not(nil)).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, blendConstants *driver.Float) {
			blendConsts := unsafe.Slice(blendConstants, 4)
			require.InDelta(t, 1, float32(blendConsts[0]), 0.0001)
			require.InDelta(t, 3, float32(blendConsts[1]), 0.0001)
			require.InDelta(t, 5, float32(blendConsts[2]), 0.0001)
			require.InDelta(t, 7, float32(blendConsts[3]), 0.0001)
		})

	buffer.CmdSetBlendConstants([4]float32{1, 3, 5, 7})
}

func TestVulkanCommandBuffer_CmdSetDepthBias(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkCmdSetDepthBias(mocks.Exactly(buffer.Handle()), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, depthBiasConstantFactor driver.Float, depthBiasClamp driver.Float, depthBiasSlopeFactor driver.Float) {
			require.InDelta(t, 1, float32(depthBiasConstantFactor), 0.0001)
			require.InDelta(t, 3, float32(depthBiasClamp), 0.0001)
			require.InDelta(t, 5, float32(depthBiasSlopeFactor), 0.0001)
		})

	buffer.CmdSetDepthBias(1, 3, 5)
}

func TestVulkanCommandBuffer_CmdSetDepthBounds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkCmdSetDepthBounds(mocks.Exactly(buffer.Handle()), gomock.Any(), gomock.Any()).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, minDepthBounds driver.Float, maxDepthBounds driver.Float) {
			require.InDelta(t, 1, float32(minDepthBounds), 0.0001)
			require.InDelta(t, 3, float32(maxDepthBounds), 0.0001)
		})

	buffer.CmdSetDepthBounds(1, 3)
}

func TestVulkanCommandBuffer_CmdSetLineWidth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkCmdSetLineWidth(mocks.Exactly(buffer.Handle()), gomock.Any()).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, lineWidth driver.Float) {
			require.InDelta(t, 3, float32(lineWidth), 0.0001)
		})

	buffer.CmdSetLineWidth(3)
}

func TestVulkanCommandBuffer_CmdSetStencilCompareMask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkCmdSetStencilCompareMask(mocks.Exactly(buffer.Handle()),
		driver.VkStencilFaceFlags(1), // VK_STENCIL_FACE_FRONT_BIT
		driver.Uint32(3),
	)

	buffer.CmdSetStencilCompareMask(common.StencilFaceFront, 3)
}

func TestVulkanCommandBuffer_CmdSetStencilReference(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkCmdSetStencilReference(mocks.Exactly(buffer.Handle()),
		driver.VkStencilFaceFlags(1), // VK_STENCIL_FACE_FRONT_BIT
		driver.Uint32(3),
	)

	buffer.CmdSetStencilReference(common.StencilFaceFront, 3)
}

func TestVulkanCommandBuffer_CmdSetStencilWriteMask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)

	mockDriver.EXPECT().VkCmdSetStencilWriteMask(mocks.Exactly(buffer.Handle()),
		driver.VkStencilFaceFlags(1), // VK_STENCIL_FACE_FRONT_BIT
		driver.Uint32(3),
	)

	buffer.CmdSetStencilWriteMask(common.StencilFaceFront, 3)
}

func TestVulkanCommandBuffer_CmdUpdateBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	dstBuffer := mocks.EasyMockBuffer(ctrl)

	mockDriver.EXPECT().VkCmdUpdateBuffer(mocks.Exactly(buffer.Handle()), mocks.Exactly(dstBuffer.Handle()), driver.VkDeviceSize(1), driver.VkDeviceSize(3), gomock.Not(nil)).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, dstBuffer driver.VkBuffer, dstOffset driver.VkDeviceSize, dataSize driver.VkDeviceSize, pData unsafe.Pointer) {
			dataPtr := (*byte)(pData)
			dataSlice := unsafe.Slice(dataPtr, 4)
			require.Equal(t, []byte{5, 7, 11, 13}, dataSlice)
		})

	buffer.CmdUpdateBuffer(dstBuffer, 1, 3, []byte{5, 7, 11, 13})
}

func TestVulkanCommandBuffer_CmdWriteTimestamp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver, buffer := setup(t, ctrl)
	queryPool := mocks.EasyMockQueryPool(ctrl)

	mockDriver.EXPECT().VkCmdWriteTimestamp(mocks.Exactly(buffer.Handle()),
		driver.VkPipelineStageFlags(0x800), // VK_PIPELINE_STAGE_COMPUTE_SHADER_BIT
		mocks.Exactly(queryPool.Handle()),
		driver.Uint32(3),
	)

	buffer.CmdWriteTimestamp(common.PipelineStageComputeShader, queryPool, 3)
}
