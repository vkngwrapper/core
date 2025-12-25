package impl1_0_test

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"github.com/vkngwrapper/core/v3/types"
	"go.uber.org/mock/gomock"
)

func setup(t *testing.T, ctrl *gomock.Controller) (*mock_loader.MockLoader, core1_0.DeviceDriver, types.Device, types.CommandBuffer) {
	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	pool := mocks.NewDummyCommandPool(device)
	buffer := mocks.NewDummyCommandBuffer(pool, device)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)

	return mockLoader, driver, device, buffer
}

func setupWithRenderPass(t *testing.T, ctrl *gomock.Controller) (*mock_loader.MockLoader, core1_0.DeviceDriver, types.Device, types.CommandBuffer, types.RenderPass, types.Framebuffer) {
	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	pool := mocks.NewDummyCommandPool(device)
	buffer := mocks.NewDummyCommandBuffer(pool, device)
	renderPass := mocks.NewDummyRenderPass(device)
	framebuffer := mocks.NewDummyFramebuffer(device)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)

	return mockLoader, driver, device, buffer, renderPass, framebuffer
}

func TestCommandBuffer_Begin_NoInheritance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkBeginCommandBuffer(buffer.Handle(), gomock.Any()).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, pBeginInfo *loader.VkCommandBufferBeginInfo) (common.VkResult, error) {
			v := reflect.ValueOf(*pBeginInfo)
			require.Equal(t, uint64(42), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(4), v.FieldByName("flags").Uint()) // VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT
			require.True(t, v.FieldByName("pInheritanceInfo").IsNil())

			return core1_0.VKSuccess, nil
		})

	_, err := driver.BeginCommandBuffer(buffer, core1_0.CommandBufferBeginInfo{
		Flags: core1_0.CommandBufferUsageSimultaneousUse,
	})
	require.NoError(t, err)
}

func TestCommandBuffer_Begin_WithInheritance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer, renderPass, framebuffer := setupWithRenderPass(t, ctrl)

	mockLoader.EXPECT().VkBeginCommandBuffer(buffer.Handle(), gomock.Any()).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, pBeginInfo *loader.VkCommandBufferBeginInfo) (common.VkResult, error) {
			v := reflect.ValueOf(*pBeginInfo)
			require.Equal(t, uint64(42), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(4), v.FieldByName("flags").Uint()) // VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT
			require.False(t, v.FieldByName("pInheritanceInfo").IsNil())

			inheritance := v.FieldByName("pInheritanceInfo").Elem()
			require.Equal(t, inheritance.FieldByName("sType").Uint(), uint64(41)) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO
			require.True(t, inheritance.FieldByName("pNext").IsNil())
			require.Equal(t, renderPass.Handle(), (loader.VkRenderPass)(unsafe.Pointer(inheritance.FieldByName("renderPass").Elem().UnsafeAddr())))
			require.Equal(t, framebuffer.Handle(), (loader.VkFramebuffer)(unsafe.Pointer(inheritance.FieldByName("framebuffer").Elem().UnsafeAddr())))
			require.Equal(t, uint64(3), inheritance.FieldByName("subpass").Uint())
			require.Equal(t, uint64(1), inheritance.FieldByName("occlusionQueryEnable").Uint())
			require.Equal(t, uint64(1), inheritance.FieldByName("queryFlags").Uint())          // VK_QUERY_CONTROL_PRECISE_BIT
			require.Equal(t, uint64(32), inheritance.FieldByName("pipelineStatistics").Uint()) // VK_QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT

			return core1_0.VKSuccess, nil
		})

	_, err := driver.BeginCommandBuffer(buffer, core1_0.CommandBufferBeginInfo{
		Flags: core1_0.CommandBufferUsageSimultaneousUse,
		InheritanceInfo: &core1_0.CommandBufferInheritanceInfo{
			Framebuffer:          framebuffer,
			RenderPass:           renderPass,
			Subpass:              3,
			OcclusionQueryEnable: true,
			QueryFlags:           core1_0.QueryControlPrecise,
			PipelineStatistics:   core1_0.QueryPipelineStatisticClippingInvocations,
		},
	})
	require.NoError(t, err)
}

func TestCommandBuffer_End(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkEndCommandBuffer(buffer.Handle()).Return(core1_0.VKSuccess, nil)

	_, err := driver.EndCommandBuffer(buffer)
	require.NoError(t, err)
}

func TestCommandBuffer_BeginRenderPass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer, renderPass, framebuffer := setupWithRenderPass(t, ctrl)

	mockLoader.EXPECT().VkCmdBeginRenderPass(buffer.Handle(), gomock.Any(), loader.VkSubpassContents(1) /*VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS*/).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, pRenderPassBegin *loader.VkRenderPassBeginInfo, contents loader.VkSubpassContents) {
			v := reflect.ValueOf(*pRenderPassBegin)
			require.Equal(t, uint64(43), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, renderPass.Handle(), (loader.VkRenderPass)(unsafe.Pointer(v.FieldByName("renderPass").Elem().UnsafeAddr())))
			require.Equal(t, framebuffer.Handle(), (loader.VkFramebuffer)(unsafe.Pointer(v.FieldByName("framebuffer").Elem().UnsafeAddr())))
			require.Equal(t, int64(1), v.FieldByName("renderArea").FieldByName("offset").FieldByName("x").Int())
			require.Equal(t, int64(2), v.FieldByName("renderArea").FieldByName("offset").FieldByName("y").Int())
			require.Equal(t, uint64(30), v.FieldByName("renderArea").FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(50), v.FieldByName("renderArea").FieldByName("extent").FieldByName("height").Uint())
			require.Equal(t, uint64(1), v.FieldByName("clearValueCount").Uint())

			clearValue := (*float32)(unsafe.Pointer(v.FieldByName("pClearValues").Elem().UnsafeAddr()))
			clearValueSlice := ([]float32)(unsafe.Slice(clearValue, 4))

			require.ElementsMatch(t, []float32{5, 6, 7, 8}, clearValueSlice)
		})

	err := driver.CmdBeginRenderPass(buffer, core1_0.SubpassContentsSecondaryCommandBuffers, core1_0.RenderPassBeginInfo{
		RenderPass:  renderPass,
		Framebuffer: framebuffer,
		RenderArea: core1_0.Rect2D{
			Offset: core1_0.Offset2D{X: 1, Y: 2},
			Extent: core1_0.Extent2D{Width: 30, Height: 50},
		},
		ClearValues: []core1_0.ClearValue{core1_0.ClearValueFloat{5, 6, 7, 8}},
	})
	require.NoError(t, err)
}

func TestCommandBuffer_EndRenderPass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkCmdEndRenderPass(buffer.Handle())

	driver.CmdEndRenderPass(buffer)
}

func TestCommandBuffer_CmdBindGraphicsPipeline(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	pool := mocks.NewDummyCommandPool(device)
	buffer := mocks.NewDummyCommandBuffer(pool, device)
	pipeline := mocks.NewDummyPipeline(device)

	mockLoader.EXPECT().VkCmdBindPipeline(buffer.Handle(), loader.VkPipelineBindPoint(0), pipeline.Handle())

	driver.CmdBindPipeline(buffer, core1_0.PipelineBindPointGraphics, pipeline)
}

func TestCommandBuffer_CmdDraw(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkCmdDraw(buffer.Handle(), loader.Uint32(6), loader.Uint32(1), loader.Uint32(2), loader.Uint32(3))

	driver.CmdDraw(buffer, 6, 1, 2, 3)
}

func TestCommandBuffer_CmdDrawIndexed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkCmdDrawIndexed(buffer.Handle(), loader.Uint32(1), loader.Uint32(2), loader.Uint32(3), loader.Int32(4), loader.Uint32(5))

	driver.CmdDrawIndexed(buffer, 1, 2, 3, 4, 5)
}

func TestVulkanCommandBuffer_CmdBindVertexBuffers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	vertexBuffer := mocks.NewDummyBuffer(device)

	mockLoader.EXPECT().VkCmdBindVertexBuffers(buffer.Handle(), loader.Uint32(0), loader.Uint32(1), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, firstBinding loader.Uint32, bindingCount loader.Uint32, pBuffers *loader.VkBuffer, pOffsets *loader.VkDeviceSize) {
			singleBuffer := ([]loader.VkBuffer)(unsafe.Slice(pBuffers, 1))
			singleOffset := ([]loader.VkDeviceSize)(unsafe.Slice(pOffsets, 1))

			require.Equal(t, vertexBuffer.Handle(), singleBuffer[0])
			require.ElementsMatch(t, []loader.VkDeviceSize{2}, singleOffset)
		})

	driver.CmdBindVertexBuffers(buffer, 0, []types.Buffer{vertexBuffer}, []int{2})
}

func TestVulkanCommandBuffer_CmdBindIndexBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	indexBuffer := mocks.NewDummyBuffer(device)

	mockLoader.EXPECT().VkCmdBindIndexBuffer(buffer.Handle(), indexBuffer.Handle(), loader.VkDeviceSize(2), loader.VkIndexType(1) /* VK_INDEX_TYPE_UINT32*/)

	driver.CmdBindIndexBuffer(buffer, indexBuffer, 2, core1_0.IndexTypeUInt32)
}

func TestVulkanCommandBuffer_CmdBindDescriptorSets(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	pipelineLayout := mocks.NewDummyPipelineLayout(device)
	descriptorPool := mocks.NewDummyDescriptorPool(device)
	descriptorSet := mocks.NewDummyDescriptorSet(descriptorPool, device)

	mockLoader.EXPECT().VkCmdBindDescriptorSets(
		buffer.Handle(),
		loader.VkPipelineBindPoint(1), /* VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR */
		pipelineLayout.Handle(),
		loader.Uint32(3),
		loader.Uint32(1),
		gomock.Not(nil),
		loader.Uint32(3),
		gomock.Not(nil)).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, bind loader.VkPipelineBindPoint, pipelineLayout loader.VkPipelineLayout, firstSet, descriptorSetCount loader.Uint32, pDescriptorSets *loader.VkDescriptorSet, dynamicOffsetCount loader.Uint32, pDynamicOffsets *loader.Uint32) {
			descriptorSetSlice := ([]loader.VkDescriptorSet)(unsafe.Slice(pDescriptorSets, 1))
			dynamicOffsetSlice := ([]loader.Uint32)(unsafe.Slice(pDynamicOffsets, 3))

			require.Equal(t, descriptorSet.Handle(), descriptorSetSlice[0])
			require.ElementsMatch(t, []loader.Uint32{4, 5, 6}, dynamicOffsetSlice)
		})

	driver.CmdBindDescriptorSets(buffer, core1_0.PipelineBindPointCompute, pipelineLayout, 3, []types.DescriptorSet{
		descriptorSet,
	}, []int{4, 5, 6})
}

func TestVulkanCommandBuffer_CmdCopyBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	src := mocks.NewDummyBuffer(device)
	dest := mocks.NewDummyBuffer(device)

	mockLoader.EXPECT().VkCmdCopyBuffer(buffer.Handle(), src.Handle(), dest.Handle(), loader.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(buffer loader.VkCommandBuffer, src loader.VkBuffer, dest loader.VkBuffer, regionCount loader.Uint32, pRegions *loader.VkBufferCopy) {
			regionSlice := ([]loader.VkBufferCopy)(unsafe.Slice(pRegions, 2))

			regionVal := reflect.ValueOf(regionSlice[0])
			require.Equal(t, uint64(3), regionVal.FieldByName("srcOffset").Uint())
			require.Equal(t, uint64(5), regionVal.FieldByName("dstOffset").Uint())
			require.Equal(t, uint64(7), regionVal.FieldByName("size").Uint())

			regionVal = reflect.ValueOf(regionSlice[1])
			require.Equal(t, uint64(11), regionVal.FieldByName("srcOffset").Uint())
			require.Equal(t, uint64(13), regionVal.FieldByName("dstOffset").Uint())
			require.Equal(t, uint64(17), regionVal.FieldByName("size").Uint())
		})

	err := driver.CmdCopyBuffer(buffer, src, dest,
		core1_0.BufferCopy{
			SrcOffset: 3,
			DstOffset: 5,
			Size:      7,
		},
		core1_0.BufferCopy{
			SrcOffset: 11,
			DstOffset: 13,
			Size:      17,
		},
	)
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdPipelineBarrier(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	mockBuffer := mocks.NewDummyBuffer(device)
	mockImage := mocks.NewDummyImage(device)

	mockLoader.EXPECT().VkCmdPipelineBarrier(buffer.Handle(),
		loader.VkPipelineStageFlags(0x00010000), // VK_PIPELINE_STAGE_ALL_COMMANDS_BIT
		loader.VkPipelineStageFlags(0x00000100), // VK_PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT
		loader.VkDependencyFlags(1),             // VK_DEPENDENCY_BY_REGION_BIT
		loader.Uint32(2),
		gomock.Not(nil),
		loader.Uint32(1),
		gomock.Not(nil),
		loader.Uint32(1),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, srcStage, dstStage loader.VkPipelineStageFlags, dependencies loader.VkDependencyFlags, memoryBarrierCount loader.Uint32, pMemoryBarriers *loader.VkMemoryBarrier, bufferMemoryBarrierCount loader.Uint32, pBufferMemoryBarriers *loader.VkBufferMemoryBarrier, imageMemoryBarrierCount loader.Uint32, pImageMemoryBarriers *loader.VkImageMemoryBarrier) {
			memoryBarrierSlice := reflect.ValueOf(([]loader.VkMemoryBarrier)(unsafe.Slice(pMemoryBarriers, 2)))
			bufferMemoryBarrierSlice := reflect.ValueOf(([]loader.VkBufferMemoryBarrier)(unsafe.Slice(pBufferMemoryBarriers, 1)))
			imageMemoryBarrierSlice := reflect.ValueOf(([]loader.VkImageMemoryBarrier)(unsafe.Slice(pImageMemoryBarriers, 1)))

			memoryBarrier := memoryBarrierSlice.Index(0)
			require.Equal(t, uint64(46), memoryBarrier.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_BARRIER
			require.True(t, memoryBarrier.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000080), memoryBarrier.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_COLOR_ATTACHMENT_READ_BIT
			require.Equal(t, uint64(0x00000010), memoryBarrier.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_INPUT_ATTACHMENT_READ_BIT

			memoryBarrier = memoryBarrierSlice.Index(1)
			require.Equal(t, uint64(46), memoryBarrier.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_BARRIER
			require.True(t, memoryBarrier.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000400), memoryBarrier.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT
			require.Equal(t, uint64(0x00004000), memoryBarrier.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_HOST_WRITE_BIT

			bufferMemoryBarrier := bufferMemoryBarrierSlice.Index(0)
			require.Equal(t, uint64(44), bufferMemoryBarrier.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER
			require.True(t, memoryBarrier.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00004000), bufferMemoryBarrier.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_HOST_WRITE_BIT
			require.Equal(t, uint64(0x00000040), bufferMemoryBarrier.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_SHADER_WRITE_BIT
			require.Equal(t, uint64(1), bufferMemoryBarrier.FieldByName("srcQueueFamilyIndex").Uint())
			require.Equal(t, uint64(3), bufferMemoryBarrier.FieldByName("dstQueueFamilyIndex").Uint())

			actualBuffer := (loader.VkBuffer)(unsafe.Pointer(bufferMemoryBarrier.FieldByName("buffer").Elem().UnsafeAddr()))
			require.Equal(t, mockBuffer.Handle(), actualBuffer)

			require.Equal(t, uint64(5), bufferMemoryBarrier.FieldByName("offset").Uint())
			require.Equal(t, uint64(7), bufferMemoryBarrier.FieldByName("size").Uint())

			imageMemoryBarrier := imageMemoryBarrierSlice.Index(0)
			require.Equal(t, uint64(45), imageMemoryBarrier.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER
			require.True(t, imageMemoryBarrier.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000002), imageMemoryBarrier.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_INDEX_READ_BIT
			require.Equal(t, uint64(0x00000200), imageMemoryBarrier.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_READ_BIT
			require.Equal(t, uint64(1), imageMemoryBarrier.FieldByName("oldLayout").Uint())
			require.Equal(t, uint64(3), imageMemoryBarrier.FieldByName("newLayout").Uint())
			require.Equal(t, uint64(11), imageMemoryBarrier.FieldByName("srcQueueFamilyIndex").Uint())
			require.Equal(t, uint64(13), imageMemoryBarrier.FieldByName("dstQueueFamilyIndex").Uint())

			actualImage := (loader.VkImage)(unsafe.Pointer(imageMemoryBarrier.FieldByName("image").Elem().UnsafeAddr()))
			require.Equal(t, mockImage.Handle(), actualImage)

			subresource := imageMemoryBarrier.FieldByName("subresourceRange")
			require.Equal(t, uint64(0x00000008), subresource.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_METADATA_BIT
			require.Equal(t, uint64(17), subresource.FieldByName("baseMipLevel").Uint())
			require.Equal(t, uint64(19), subresource.FieldByName("levelCount").Uint())
			require.Equal(t, uint64(23), subresource.FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(29), subresource.FieldByName("layerCount").Uint())
		})

	err := driver.CmdPipelineBarrier(
		buffer,
		core1_0.PipelineStageAllCommands,
		core1_0.PipelineStageEarlyFragmentTests,
		core1_0.DependencyByRegion,
		[]core1_0.MemoryBarrier{
			{
				SrcAccessMask: core1_0.AccessColorAttachmentRead,
				DstAccessMask: core1_0.AccessInputAttachmentRead,
			},
			{
				SrcAccessMask: core1_0.AccessDepthStencilAttachmentWrite,
				DstAccessMask: core1_0.AccessHostWrite,
			},
		},
		[]core1_0.BufferMemoryBarrier{
			{
				SrcAccessMask:       core1_0.AccessHostWrite,
				DstAccessMask:       core1_0.AccessShaderWrite,
				SrcQueueFamilyIndex: 1,
				DstQueueFamilyIndex: 3,
				Buffer:              mockBuffer,
				Offset:              5,
				Size:                7,
			},
		},
		[]core1_0.ImageMemoryBarrier{
			{
				SrcAccessMask:       core1_0.AccessIndexRead,
				DstAccessMask:       core1_0.AccessDepthStencilAttachmentRead,
				OldLayout:           core1_0.ImageLayoutGeneral,
				NewLayout:           core1_0.ImageLayoutDepthStencilAttachmentOptimal,
				SrcQueueFamilyIndex: 11,
				DstQueueFamilyIndex: 13,
				Image:               mockImage,
				SubresourceRange: core1_0.ImageSubresourceRange{
					AspectMask:     core1_0.ImageAspectMetadata,
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

	mockLoader, driver, device, buffer := setup(t, ctrl)
	mockBuffer := mocks.NewDummyBuffer(device)
	mockImage := mocks.NewDummyImage(device)

	mockLoader.EXPECT().VkCmdCopyBufferToImage(buffer.Handle(),
		mockBuffer.Handle(),
		mockImage.Handle(),
		loader.VkImageLayout(8), // VK_IMAGE_LAYOUT_PREINITIALIZED
		loader.Uint32(2),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, srcBuffer loader.VkBuffer, dstImage loader.VkImage, dstImageLayout loader.VkImageLayout, regionCount loader.Uint32, pRegions *loader.VkBufferImageCopy) {
			regionSlice := reflect.ValueOf(([]loader.VkBufferImageCopy)(unsafe.Slice(pRegions, 2)))

			region := regionSlice.Index(0)
			require.Equal(t, uint64(1), region.FieldByName("bufferOffset").Uint())
			require.Equal(t, uint64(3), region.FieldByName("bufferRowLength").Uint())
			require.Equal(t, uint64(5), region.FieldByName("bufferImageHeight").Uint())

			subresource := region.FieldByName("imageSubresource")
			require.Equal(t, uint64(0x00000002), subresource.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_DEPTH_BIT
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

	err := driver.CmdCopyBufferToImage(buffer, mockBuffer, mockImage, core1_0.ImageLayoutPreInitialized,
		core1_0.BufferImageCopy{
			BufferOffset:      1,
			BufferRowLength:   3,
			BufferImageHeight: 5,
			ImageSubresource: core1_0.ImageSubresourceLayers{
				AspectMask:     core1_0.ImageAspectDepth,
				MipLevel:       7,
				BaseArrayLayer: 11,
				LayerCount:     13,
			},
			ImageOffset: core1_0.Offset3D{
				X: 17,
				Y: 19,
				Z: 23,
			},
			ImageExtent: core1_0.Extent3D{
				Width:  29,
				Height: 31,
				Depth:  37,
			},
		},
		core1_0.BufferImageCopy{
			BufferOffset:      41,
			BufferRowLength:   43,
			BufferImageHeight: 47,
			ImageSubresource: core1_0.ImageSubresourceLayers{
				AspectMask:     core1_0.ImageAspectColor,
				MipLevel:       53,
				BaseArrayLayer: 59,
				LayerCount:     61,
			},
			ImageOffset: core1_0.Offset3D{
				X: 67,
				Y: 71,
				Z: 73,
			},
			ImageExtent: core1_0.Extent3D{
				Width:  79,
				Height: 83,
				Depth:  89,
			},
		},
	)
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdBlitImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	sourceImage := mocks.NewDummyImage(device)
	destImage := mocks.NewDummyImage(device)

	mockLoader.EXPECT().VkCmdBlitImage(buffer.Handle(),
		sourceImage.Handle(),
		loader.VkImageLayout(6), // VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL
		destImage.Handle(),
		loader.VkImageLayout(2), // VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
		loader.Uint32(1),
		gomock.Not(nil),
		loader.VkFilter(1), // VK_FILTER_LINEAR
	).DoAndReturn(func(commandBuffer loader.VkCommandBuffer,
		sourceImage loader.VkImage,
		sourceImageLayout loader.VkImageLayout,
		destImage loader.VkImage,
		destImageLayout loader.VkImageLayout,
		regionCount loader.Uint32,
		pRegions *loader.VkImageBlit,
		filter loader.VkFilter) {

		regionSlice := reflect.ValueOf(([]loader.VkImageBlit)(unsafe.Slice(pRegions, 1)))
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

	err := driver.CmdBlitImage(buffer, sourceImage,
		core1_0.ImageLayoutTransferSrcOptimal,
		destImage,
		core1_0.ImageLayoutColorAttachmentOptimal,
		[]core1_0.ImageBlit{
			{
				SrcSubresource: core1_0.ImageSubresourceLayers{
					AspectMask:     core1_0.ImageAspectMetadata,
					MipLevel:       1,
					BaseArrayLayer: 3,
					LayerCount:     5,
				},
				SrcOffsets: [2]core1_0.Offset3D{
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
				DstSubresource: core1_0.ImageSubresourceLayers{
					AspectMask:     core1_0.ImageAspectStencil,
					MipLevel:       29,
					BaseArrayLayer: 31,
					LayerCount:     37,
				},
				DstOffsets: [2]core1_0.Offset3D{
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
		core1_0.FilterLinear,
	)
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdPushConstants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	pipelineLayout := mocks.NewDummyPipelineLayout(device)

	mockLoader.EXPECT().VkCmdPushConstants(buffer.Handle(),
		pipelineLayout.Handle(),
		loader.VkShaderStageFlags(8), // VK_SHADER_STAGE_GEOMETRY_BIT
		loader.Uint32(1),
		loader.Uint32(4),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer,
			pipelineLayout loader.VkPipelineLayout,
			shaderStages loader.VkShaderStageFlags,
			offset loader.Uint32,
			size loader.Uint32,
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

	driver.CmdPushConstants(buffer, pipelineLayout, core1_0.StageGeometry, 1, writer.Bytes())
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdSetViewport(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkCmdSetViewport(buffer.Handle(), loader.Uint32(0), loader.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(
			commandBuffer loader.VkCommandBuffer,
			firstViewport loader.Uint32,
			viewportCount loader.Uint32,
			pViewports *loader.VkViewport) {

			viewportSlice := ([]loader.VkViewport)(unsafe.Slice(pViewports, 2))
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

	driver.CmdSetViewport(buffer,
		core1_0.Viewport{
			X:        3,
			Y:        5,
			Width:    7,
			Height:   11,
			MinDepth: 13,
			MaxDepth: 17,
		},
		core1_0.Viewport{
			X:        19,
			Y:        23,
			Width:    29,
			Height:   31,
			MinDepth: 37,
			MaxDepth: 41,
		},
	)
}

func TestVulkanCommandBuffer_CmdSetScissor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkCmdSetScissor(buffer.Handle(), loader.Uint32(0), loader.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(
			commandBuffer loader.VkCommandBuffer,
			firstScissor loader.Uint32,
			scissorCount loader.Uint32,
			pScissors *loader.VkRect2D) {

			scissorSlice := ([]loader.VkRect2D)(unsafe.Slice(pScissors, 2))
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

	driver.CmdSetScissor(buffer,
		core1_0.Rect2D{
			Offset: core1_0.Offset2D{3, 5},
			Extent: core1_0.Extent2D{7, 11},
		},
		core1_0.Rect2D{
			Offset: core1_0.Offset2D{13, 17},
			Extent: core1_0.Extent2D{19, 23},
		},
	)
}

func TestVulkanCommandBuffer_CmdCopyImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	srcImage := mocks.NewDummyImage(device)
	dstImage := mocks.NewDummyImage(device)

	mockLoader.EXPECT().VkCmdCopyImage(buffer.Handle(),
		srcImage.Handle(),
		loader.VkImageLayout(7), // VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL
		dstImage.Handle(),
		loader.VkImageLayout(5), // VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL
		loader.Uint32(2),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, srcImage loader.VkImage, srcImageLayout loader.VkImageLayout, dstImage loader.VkImage, dstImageLayout loader.VkImageLayout, regionCount loader.Uint32, pRegions *loader.VkImageCopy) {
			regionSlice := ([]loader.VkImageCopy)(unsafe.Slice(pRegions, 2))
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

			require.Equal(t, uint64(0x00000004), dstSubresource.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_STENCIL_BIT
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

	err := driver.CmdCopyImage(buffer, srcImage, core1_0.ImageLayoutTransferDstOptimal, dstImage, core1_0.ImageLayoutShaderReadOnlyOptimal,
		core1_0.ImageCopy{
			SrcSubresource: core1_0.ImageSubresourceLayers{
				AspectMask:     core1_0.ImageAspectMetadata,
				MipLevel:       1,
				BaseArrayLayer: 3,
				LayerCount:     5,
			},
			DstSubresource: core1_0.ImageSubresourceLayers{
				AspectMask:     core1_0.ImageAspectStencil,
				MipLevel:       7,
				BaseArrayLayer: 11,
				LayerCount:     13,
			},
			SrcOffset: core1_0.Offset3D{17, 19, 23},
			DstOffset: core1_0.Offset3D{29, 31, 37},
			Extent:    core1_0.Extent3D{41, 43, 47},
		},
		core1_0.ImageCopy{
			SrcSubresource: core1_0.ImageSubresourceLayers{
				AspectMask:     core1_0.ImageAspectColor,
				MipLevel:       53,
				BaseArrayLayer: 59,
				LayerCount:     61,
			},
			DstSubresource: core1_0.ImageSubresourceLayers{
				AspectMask:     core1_0.ImageAspectDepth,
				MipLevel:       67,
				BaseArrayLayer: 71,
				LayerCount:     73,
			},
			SrcOffset: core1_0.Offset3D{79, 83, 89},
			DstOffset: core1_0.Offset3D{97, 101, 103},
			Extent:    core1_0.Extent3D{107, 109, 113},
		},
	)
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdNextSubpass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkCmdNextSubpass(buffer.Handle(),
		loader.VkSubpassContents(1), /* VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS */
	)

	driver.CmdNextSubpass(buffer, core1_0.SubpassContentsSecondaryCommandBuffers)
}

func TestVulkanCommandBuffer_CmdWaitEvents(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	mockBuffer := mocks.NewDummyBuffer(device)
	mockImage := mocks.NewDummyImage(device)
	event1 := mocks.NewDummyEvent(device)
	event2 := mocks.NewDummyEvent(device)

	mockLoader.EXPECT().VkCmdWaitEvents(buffer.Handle(),
		loader.Uint32(2),
		gomock.Not(nil),
		loader.VkPipelineStageFlags(0x00010000), // VK_PIPELINE_STAGE_ALL_COMMANDS_BIT
		loader.VkPipelineStageFlags(0x00000010), // VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT
		loader.Uint32(2),
		gomock.Not(nil),
		loader.Uint32(1),
		gomock.Not(nil),
		loader.Uint32(1),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, eventCount loader.Uint32, pEvents *loader.VkEvent, srcStage, dstStage loader.VkPipelineStageFlags, memoryBarrierCount loader.Uint32, pMemoryBarriers *loader.VkMemoryBarrier, bufferMemoryBarrierCount loader.Uint32, pBufferMemoryBarriers *loader.VkBufferMemoryBarrier, imageMemoryBarrierCount loader.Uint32, pImageMemoryBarriers *loader.VkImageMemoryBarrier) {
			eventSlice := ([]loader.VkEvent)(unsafe.Slice(pEvents, 2))
			memoryBarrierSlice := reflect.ValueOf(([]loader.VkMemoryBarrier)(unsafe.Slice(pMemoryBarriers, 2)))
			bufferMemoryBarrierSlice := reflect.ValueOf(([]loader.VkBufferMemoryBarrier)(unsafe.Slice(pBufferMemoryBarriers, 1)))
			imageMemoryBarrierSlice := reflect.ValueOf(([]loader.VkImageMemoryBarrier)(unsafe.Slice(pImageMemoryBarriers, 1)))

			require.Equal(t, event1.Handle(), eventSlice[0])
			require.Equal(t, event2.Handle(), eventSlice[1])

			memoryBarrier := memoryBarrierSlice.Index(0)
			require.Equal(t, uint64(46), memoryBarrier.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_BARRIER
			require.True(t, memoryBarrier.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000080), memoryBarrier.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_COLOR_ATTACHMENT_READ_BIT
			require.Equal(t, uint64(0x00000008), memoryBarrier.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_UNIFORM_READ_BIT

			memoryBarrier = memoryBarrierSlice.Index(1)
			require.Equal(t, uint64(46), memoryBarrier.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_BARRIER
			require.True(t, memoryBarrier.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000400), memoryBarrier.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT
			require.Equal(t, uint64(0x00000004), memoryBarrier.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_VERTEX_ATTRIBUTE_READ_BIT

			bufferMemoryBarrier := bufferMemoryBarrierSlice.Index(0)
			require.Equal(t, uint64(44), bufferMemoryBarrier.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER
			require.True(t, memoryBarrier.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00004000), bufferMemoryBarrier.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_HOST_WRITE_BIT
			require.Equal(t, uint64(0x00000040), bufferMemoryBarrier.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_SHADER_WRITE_BIT
			require.Equal(t, uint64(1), bufferMemoryBarrier.FieldByName("srcQueueFamilyIndex").Uint())
			require.Equal(t, uint64(3), bufferMemoryBarrier.FieldByName("dstQueueFamilyIndex").Uint())

			actualBuffer := (loader.VkBuffer)(unsafe.Pointer(bufferMemoryBarrier.FieldByName("buffer").Elem().UnsafeAddr()))
			require.Equal(t, mockBuffer.Handle(), actualBuffer)

			require.Equal(t, uint64(5), bufferMemoryBarrier.FieldByName("offset").Uint())
			require.Equal(t, uint64(7), bufferMemoryBarrier.FieldByName("size").Uint())

			imageMemoryBarrier := imageMemoryBarrierSlice.Index(0)
			require.Equal(t, uint64(45), imageMemoryBarrier.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER
			require.True(t, imageMemoryBarrier.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000002), imageMemoryBarrier.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_INDEX_READ_BIT
			require.Equal(t, uint64(0x00000001), imageMemoryBarrier.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_INDIRECT_COMMAND_READ_BIT
			require.Equal(t, uint64(1), imageMemoryBarrier.FieldByName("oldLayout").Uint())
			require.Equal(t, uint64(3), imageMemoryBarrier.FieldByName("newLayout").Uint())
			require.Equal(t, uint64(11), imageMemoryBarrier.FieldByName("srcQueueFamilyIndex").Uint())
			require.Equal(t, uint64(13), imageMemoryBarrier.FieldByName("dstQueueFamilyIndex").Uint())

			actualImage := (loader.VkImage)(unsafe.Pointer(imageMemoryBarrier.FieldByName("image").Elem().UnsafeAddr()))
			require.Equal(t, mockImage.Handle(), actualImage)

			subresource := imageMemoryBarrier.FieldByName("subresourceRange")
			require.Equal(t, uint64(0x00000002), subresource.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_DEPTH_BIT
			require.Equal(t, uint64(17), subresource.FieldByName("baseMipLevel").Uint())
			require.Equal(t, uint64(19), subresource.FieldByName("levelCount").Uint())
			require.Equal(t, uint64(23), subresource.FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(29), subresource.FieldByName("layerCount").Uint())
		})

	err := driver.CmdWaitEvents(
		buffer,
		[]types.Event{event1, event2},
		core1_0.PipelineStageAllCommands,
		core1_0.PipelineStageTessellationControlShader,
		[]core1_0.MemoryBarrier{
			{
				SrcAccessMask: core1_0.AccessColorAttachmentRead,
				DstAccessMask: core1_0.AccessUniformRead,
			},
			{
				SrcAccessMask: core1_0.AccessDepthStencilAttachmentWrite,
				DstAccessMask: core1_0.AccessVertexAttributeRead,
			},
		},
		[]core1_0.BufferMemoryBarrier{
			{
				SrcAccessMask:       core1_0.AccessHostWrite,
				DstAccessMask:       core1_0.AccessShaderWrite,
				SrcQueueFamilyIndex: 1,
				DstQueueFamilyIndex: 3,
				Buffer:              mockBuffer,
				Offset:              5,
				Size:                7,
			},
		},
		[]core1_0.ImageMemoryBarrier{
			{
				SrcAccessMask:       core1_0.AccessIndexRead,
				DstAccessMask:       core1_0.AccessIndirectCommandRead,
				OldLayout:           core1_0.ImageLayoutGeneral,
				NewLayout:           core1_0.ImageLayoutDepthStencilAttachmentOptimal,
				SrcQueueFamilyIndex: 11,
				DstQueueFamilyIndex: 13,
				Image:               mockImage,
				SubresourceRange: core1_0.ImageSubresourceRange{
					AspectMask:     core1_0.ImageAspectDepth,
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

	mockLoader, driver, device, buffer := setup(t, ctrl)
	event := mocks.NewDummyEvent(device)

	mockLoader.EXPECT().VkCmdSetEvent(buffer.Handle(), event.Handle(), loader.VkPipelineStageFlags(0x80) /*VK_PIPELINE_STAGE_FRAGMENT_SHADER_BIT*/)

	driver.CmdSetEvent(buffer, event, core1_0.PipelineStageFragmentShader)
}

func TestVulkanCommandBuffer_CmdClearColorImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	image := mocks.NewDummyImage(device)

	mockLoader.EXPECT().VkCmdClearColorImage(buffer.Handle(),
		image.Handle(),
		loader.VkImageLayout(3), // VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL
		gomock.Not(nil),
		loader.Uint32(2),
		gomock.Not(nil),
	).DoAndReturn(
		func(buffer loader.VkCommandBuffer, image loader.VkImage, imageLayout loader.VkImageLayout, pColor *loader.VkClearColorValue, rangeCount loader.Uint32, pRanges *loader.VkImageSubresourceRange) {
			colorFloat := unsafe.Slice((*float32)(unsafe.Pointer(pColor)), 4)
			require.InDelta(t, 0.2, colorFloat[0], 0.0001)
			require.InDelta(t, 0.3, colorFloat[1], 0.0001)
			require.InDelta(t, 0.4, colorFloat[2], 0.0001)
			require.InDelta(t, 0.5, colorFloat[3], 0.0001)

			rangeSlice := reflect.ValueOf(([]loader.VkImageSubresourceRange)(unsafe.Slice(pRanges, 2)))
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

	driver.CmdClearColorImage(buffer, image, core1_0.ImageLayoutDepthStencilAttachmentOptimal, &core1_0.ClearValueFloat{0.2, 0.3, 0.4, 0.5},
		core1_0.ImageSubresourceRange{
			AspectMask:     core1_0.ImageAspectMetadata,
			BaseMipLevel:   1,
			LevelCount:     3,
			BaseArrayLayer: 5,
			LayerCount:     7,
		},
		core1_0.ImageSubresourceRange{
			AspectMask:     core1_0.ImageAspectDepth,
			BaseMipLevel:   11,
			LevelCount:     13,
			BaseArrayLayer: 17,
			LayerCount:     19,
		},
	)
}

func TestVulkanCommandBuffer_Reset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkResetCommandBuffer(buffer.Handle(), loader.VkCommandBufferResetFlags(1)).Return(core1_0.VKSuccess, nil)

	_, err := driver.ResetCommandBuffer(buffer, core1_0.CommandBufferResetReleaseResources)
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdResetQueryPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	queryPool := mocks.NewDummyQueryPool(device)

	mockLoader.EXPECT().VkCmdResetQueryPool(buffer.Handle(), queryPool.Handle(), loader.Uint32(1), loader.Uint32(3))

	driver.CmdResetQueryPool(buffer, queryPool, 1, 3)
}

func TestVulkanCommandBuffer_CmdBeginQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	queryPool := mocks.NewDummyQueryPool(device)

	mockLoader.EXPECT().VkCmdBeginQuery(
		buffer.Handle(),
		queryPool.Handle(),
		loader.Uint32(5),
		loader.VkQueryControlFlags(1), // VK_QUERY_CONTROL_PRECISE_BIT
	)

	driver.CmdBeginQuery(buffer, queryPool, 5, core1_0.QueryControlPrecise)
}

func TestVulkanCommandBuffer_CmdEndQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	queryPool := mocks.NewDummyQueryPool(device)

	mockLoader.EXPECT().VkCmdEndQuery(
		buffer.Handle(),
		queryPool.Handle(),
		loader.Uint32(5),
	)

	driver.CmdEndQuery(buffer, queryPool, 5)
}

func TestVulkanCommandBuffer_CmdCopyQueryPoolResults(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	queryPool := mocks.NewDummyQueryPool(device)
	dstBuffer := mocks.NewDummyBuffer(device)

	mockLoader.EXPECT().VkCmdCopyQueryPoolResults(
		buffer.Handle(),
		queryPool.Handle(),
		loader.Uint32(1),
		loader.Uint32(3),
		dstBuffer.Handle(),
		loader.VkDeviceSize(5),
		loader.VkDeviceSize(7),
		loader.VkQueryResultFlags(8), // VK_QUERY_RESULT_PARTIAL_BIT
	)

	driver.CmdCopyQueryPoolResults(buffer, queryPool, 1, 3, dstBuffer, 5, 7, core1_0.QueryResultPartial)
}

func TestVulkanCommandBuffer_CmdExecuteCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	pool := mocks.NewDummyCommandPool(device)

	cmd1 := mocks.NewDummyCommandBuffer(pool, device)
	cmd2 := mocks.NewDummyCommandBuffer(pool, device)
	cmd3 := mocks.NewDummyCommandBuffer(pool, device)

	commandBuffers := []types.CommandBuffer{
		cmd1, cmd2, cmd3,
	}

	mockLoader.EXPECT().VkCmdExecuteCommands(
		buffer.Handle(),
		loader.Uint32(3),
		gomock.Not(nil)).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, secondaryCount loader.Uint32, pSecondaryBuffers *loader.VkCommandBuffer) (common.VkResult, error) {
			secondaryBufferSlice := ([]loader.VkCommandBuffer)(unsafe.Slice(pSecondaryBuffers, 3))
			require.Equal(t, commandBuffers[0].Handle(), secondaryBufferSlice[0])
			require.Equal(t, commandBuffers[1].Handle(), secondaryBufferSlice[1])
			require.Equal(t, commandBuffers[2].Handle(), secondaryBufferSlice[2])

			return core1_0.VKSuccess, nil
		})

	driver.CmdExecuteCommands(buffer, commandBuffers...)
}

func TestVulkanCommandBuffer_CmdClearAttachments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkCmdClearAttachments(
		buffer.Handle(),
		loader.Uint32(1),
		gomock.Not(nil),
		loader.Uint32(2),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, attachmentCount loader.Uint32, pAttachments *loader.VkClearAttachment, rectCount loader.Uint32, pRects *loader.VkClearRect) {
			attachmentSlice := ([]loader.VkClearAttachment)(unsafe.Slice(pAttachments, 1))
			rectSlice := ([]loader.VkClearRect)(unsafe.Slice(pRects, 2))

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

	err := driver.CmdClearAttachments(buffer, []core1_0.ClearAttachment{
		{
			AspectMask:      core1_0.ImageAspectColor,
			ColorAttachment: 3,
			ClearValue:      core1_0.ClearValueFloat{5, 7, 11, 13},
		},
	}, []core1_0.ClearRect{
		{
			BaseArrayLayer: 17,
			LayerCount:     19,
			Rect: core1_0.Rect2D{
				Offset: core1_0.Offset2D{23, 29},
				Extent: core1_0.Extent2D{31, 37},
			},
		},
		{
			BaseArrayLayer: 41,
			LayerCount:     43,
			Rect: core1_0.Rect2D{
				Offset: core1_0.Offset2D{47, 53},
				Extent: core1_0.Extent2D{59, 61},
			},
		},
	})
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdClearDepthStencilImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	image := mocks.NewDummyImage(device)

	mockLoader.EXPECT().VkCmdClearDepthStencilImage(buffer.Handle(), image.Handle(),
		loader.VkImageLayout(5), // VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL
		gomock.Not(nil),
		loader.Uint32(2),
		gomock.Not(nil),
	).DoAndReturn(func(commandBuffer loader.VkCommandBuffer, image loader.VkImage, imageLayout loader.VkImageLayout, pDepthStencil *loader.VkClearDepthStencilValue, rangeCount loader.Uint32, pRanges *loader.VkImageSubresourceRange) {
		depthStencil := reflect.ValueOf(pDepthStencil).Elem()

		require.InDelta(t, 0.5, depthStencil.FieldByName("depth").Float(), 0.00001)
		require.Equal(t, uint64(3), depthStencil.FieldByName("stencil").Uint())

		rangeSlice := reflect.ValueOf(([]loader.VkImageSubresourceRange)(unsafe.Slice(pRanges, 2)))

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

	driver.CmdClearDepthStencilImage(buffer, image, core1_0.ImageLayoutShaderReadOnlyOptimal,
		&core1_0.ClearValueDepthStencil{
			Depth:   0.5,
			Stencil: 3,
		},
		core1_0.ImageSubresourceRange{
			AspectMask:     core1_0.ImageAspectColor,
			BaseMipLevel:   5,
			LevelCount:     7,
			BaseArrayLayer: 11,
			LayerCount:     13,
		},
		core1_0.ImageSubresourceRange{
			AspectMask:     core1_0.ImageAspectDepth,
			BaseMipLevel:   17,
			LevelCount:     19,
			BaseArrayLayer: 23,
			LayerCount:     29,
		},
	)
}

func TestVulkanCommandBuffer_CmdCopyImageToBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	image := mocks.NewDummyImage(device)
	dstBuffer := mocks.NewDummyBuffer(device)

	mockLoader.EXPECT().VkCmdCopyImageToBuffer(buffer.Handle(), image.Handle(),
		loader.VkImageLayout(1), // VK_IMAGE_LAYOUT_GENERAL
		dstBuffer.Handle(),
		loader.Uint32(1),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, srcImage loader.VkImage, srcImageLayout loader.VkImageLayout, dstBuffer loader.VkBuffer, regionCount loader.Uint32, pRegions *loader.VkBufferImageCopy) {
			regionSlice := ([]loader.VkBufferImageCopy)(unsafe.Slice(pRegions, 1))
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

	err := driver.CmdCopyImageToBuffer(buffer, image, core1_0.ImageLayoutGeneral, dstBuffer,
		core1_0.BufferImageCopy{
			BufferOffset:      1,
			BufferRowLength:   3,
			BufferImageHeight: 5,
			ImageSubresource: core1_0.ImageSubresourceLayers{
				AspectMask:     core1_0.ImageAspectColor,
				MipLevel:       7,
				BaseArrayLayer: 11,
				LayerCount:     13,
			},
			ImageOffset: core1_0.Offset3D{17, 19, 23},
			ImageExtent: core1_0.Extent3D{29, 31, 37},
		},
	)
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdDispatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkCmdDispatch(buffer.Handle(), loader.Uint32(1), loader.Uint32(3), loader.Uint32(5))

	driver.CmdDispatch(buffer, 1, 3, 5)
}

func TestVulkanCommandBuffer_CmdDispatchIndirect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	indirectBuffer := mocks.NewDummyBuffer(device)

	mockLoader.EXPECT().VkCmdDispatchIndirect(buffer.Handle(), indirectBuffer.Handle(), loader.VkDeviceSize(3))

	driver.CmdDispatchIndirect(buffer, indirectBuffer, 3)
}

func TestVulkanCommandBuffer_CmdDrawIndexedIndirect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	indirectBuffer := mocks.NewDummyBuffer(device)

	mockLoader.EXPECT().VkCmdDrawIndexedIndirect(buffer.Handle(), indirectBuffer.Handle(), loader.VkDeviceSize(3), loader.Uint32(5), loader.Uint32(7))

	driver.CmdDrawIndexedIndirect(buffer, indirectBuffer, 3, 5, 7)
}

func TestVulkanCommandBuffer_CmdDrawIndirect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	indirectBuffer := mocks.NewDummyBuffer(device)

	mockLoader.EXPECT().VkCmdDrawIndirect(buffer.Handle(), indirectBuffer.Handle(), loader.VkDeviceSize(3), loader.Uint32(5), loader.Uint32(7))

	driver.CmdDrawIndirect(buffer, indirectBuffer, 3, 5, 7)
}

func TestVulkanCommandBuffer_CmdFillBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	fillBuffer := mocks.NewDummyBuffer(device)

	mockLoader.EXPECT().VkCmdFillBuffer(buffer.Handle(), fillBuffer.Handle(), loader.VkDeviceSize(1), loader.VkDeviceSize(3), loader.Uint32(5))

	driver.CmdFillBuffer(buffer, fillBuffer, 1, 3, 5)
}

func TestVulkanCommandBuffer_CmdResetEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	event := mocks.NewDummyEvent(device)

	mockLoader.EXPECT().VkCmdResetEvent(
		buffer.Handle(), event.Handle(),
		loader.VkPipelineStageFlags(0x00004000), // VK_PIPELINE_STAGE_HOST_BIT
	)

	driver.CmdResetEvent(buffer, event, core1_0.PipelineStageHost)
}

func TestVulkanCommandBuffer_CmdResolveImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	srcImage := mocks.NewDummyImage(device)
	dstImage := mocks.NewDummyImage(device)

	mockLoader.EXPECT().VkCmdResolveImage(
		buffer.Handle(),
		srcImage.Handle(),
		loader.VkImageLayout(5), // VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL
		dstImage.Handle(),
		loader.VkImageLayout(2), // VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
		loader.Uint32(2),
		gomock.Not(nil),
	).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, srcImage loader.VkImage, srcImageLayout loader.VkImageLayout, dstImage loader.VkImage, dstImageLayout loader.VkImageLayout, resolveCount loader.Uint32, pResolves *loader.VkImageResolve) {
			resolveSlice := ([]loader.VkImageResolve)(unsafe.Slice(pResolves, 2))
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

	err := driver.CmdResolveImage(buffer, srcImage,
		core1_0.ImageLayoutShaderReadOnlyOptimal,
		dstImage,
		core1_0.ImageLayoutColorAttachmentOptimal,
		core1_0.ImageResolve{
			SrcSubresource: core1_0.ImageSubresourceLayers{
				AspectMask:     core1_0.ImageAspectColor,
				MipLevel:       1,
				BaseArrayLayer: 3,
				LayerCount:     5,
			},
			SrcOffset: core1_0.Offset3D{7, 11, 13},
			DstSubresource: core1_0.ImageSubresourceLayers{
				AspectMask:     core1_0.ImageAspectDepth,
				MipLevel:       17,
				BaseArrayLayer: 19,
				LayerCount:     23,
			},
			DstOffset: core1_0.Offset3D{29, 31, 37},
			Extent:    core1_0.Extent3D{41, 43, 47},
		},
		core1_0.ImageResolve{
			SrcSubresource: core1_0.ImageSubresourceLayers{
				AspectMask:     core1_0.ImageAspectMetadata,
				MipLevel:       53,
				BaseArrayLayer: 59,
				LayerCount:     61,
			},
			SrcOffset: core1_0.Offset3D{67, 71, 73},
			DstSubresource: core1_0.ImageSubresourceLayers{
				AspectMask:     core1_0.ImageAspectStencil,
				MipLevel:       79,
				BaseArrayLayer: 83,
				LayerCount:     89,
			},
			DstOffset: core1_0.Offset3D{97, 101, 103},
			Extent:    core1_0.Extent3D{107, 109, 113},
		},
	)
	require.NoError(t, err)
}

func TestVulkanCommandBuffer_CmdSetBlendConstants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkCmdSetBlendConstants(buffer.Handle(), gomock.Not(nil)).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, blendConstants *loader.Float) {
			blendConsts := unsafe.Slice(blendConstants, 4)
			require.InDelta(t, 1, float32(blendConsts[0]), 0.0001)
			require.InDelta(t, 3, float32(blendConsts[1]), 0.0001)
			require.InDelta(t, 5, float32(blendConsts[2]), 0.0001)
			require.InDelta(t, 7, float32(blendConsts[3]), 0.0001)
		})

	driver.CmdSetBlendConstants(buffer, [4]float32{1, 3, 5, 7})
}

func TestVulkanCommandBuffer_CmdSetDepthBias(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkCmdSetDepthBias(buffer.Handle(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, depthBiasConstantFactor loader.Float, depthBiasClamp loader.Float, depthBiasSlopeFactor loader.Float) {
			require.InDelta(t, 1, float32(depthBiasConstantFactor), 0.0001)
			require.InDelta(t, 3, float32(depthBiasClamp), 0.0001)
			require.InDelta(t, 5, float32(depthBiasSlopeFactor), 0.0001)
		})

	driver.CmdSetDepthBias(buffer, 1, 3, 5)
}

func TestVulkanCommandBuffer_CmdSetDepthBounds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkCmdSetDepthBounds(buffer.Handle(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, minDepthBounds loader.Float, maxDepthBounds loader.Float) {
			require.InDelta(t, 1, float32(minDepthBounds), 0.0001)
			require.InDelta(t, 3, float32(maxDepthBounds), 0.0001)
		})

	driver.CmdSetDepthBounds(buffer, 1, 3)
}

func TestVulkanCommandBuffer_CmdSetLineWidth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkCmdSetLineWidth(buffer.Handle(), gomock.Any()).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, lineWidth loader.Float) {
			require.InDelta(t, 3, float32(lineWidth), 0.0001)
		})

	driver.CmdSetLineWidth(buffer, 3)
}

func TestVulkanCommandBuffer_CmdSetStencilCompareMask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkCmdSetStencilCompareMask(buffer.Handle(),
		loader.VkStencilFaceFlags(1), // VK_STENCIL_FACE_FRONT_BIT
		loader.Uint32(3),
	)

	driver.CmdSetStencilCompareMask(buffer, core1_0.StencilFaceFront, 3)
}

func TestVulkanCommandBuffer_CmdSetStencilReference(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkCmdSetStencilReference(buffer.Handle(),
		loader.VkStencilFaceFlags(1), // VK_STENCIL_FACE_FRONT_BIT
		loader.Uint32(3),
	)

	driver.CmdSetStencilReference(buffer, core1_0.StencilFaceFront, 3)
}

func TestVulkanCommandBuffer_CmdSetStencilWriteMask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, _, buffer := setup(t, ctrl)

	mockLoader.EXPECT().VkCmdSetStencilWriteMask(buffer.Handle(),
		loader.VkStencilFaceFlags(1), // VK_STENCIL_FACE_FRONT_BIT
		loader.Uint32(3),
	)

	driver.CmdSetStencilWriteMask(buffer, core1_0.StencilFaceFront, 3)
}

func TestVulkanCommandBuffer_CmdUpdateBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	dstBuffer := mocks.NewDummyBuffer(device)

	mockLoader.EXPECT().VkCmdUpdateBuffer(buffer.Handle(), dstBuffer.Handle(), loader.VkDeviceSize(1), loader.VkDeviceSize(3), gomock.Not(nil)).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, dstBuffer loader.VkBuffer, dstOffset loader.VkDeviceSize, dataSize loader.VkDeviceSize, pData unsafe.Pointer) {
			dataPtr := (*byte)(pData)
			dataSlice := unsafe.Slice(dataPtr, 4)
			require.Equal(t, []byte{5, 7, 11, 13}, dataSlice)
		})

	driver.CmdUpdateBuffer(buffer, dstBuffer, 1, 3, []byte{5, 7, 11, 13})
}

func TestVulkanCommandBuffer_CmdWriteTimestamp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader, driver, device, buffer := setup(t, ctrl)
	queryPool := mocks.NewDummyQueryPool(device)

	mockLoader.EXPECT().VkCmdWriteTimestamp(buffer.Handle(),
		loader.VkPipelineStageFlags(0x800), // VK_PIPELINE_STAGE_COMPUTE_SHADER_BIT
		queryPool.Handle(),
		loader.Uint32(3),
	)

	driver.CmdWriteTimestamp(buffer, core1_0.PipelineStageComputeShader, queryPool, 3)
}
