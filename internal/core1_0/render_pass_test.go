package core1_0_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	internal_mocks "github.com/CannibalVox/VKng/core/internal/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanLoader1_0_CreateRenderPass_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	renderPassHandle := mocks.NewFakeRenderPassHandle()

	mockDriver.EXPECT().VkCreateRenderPass(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(deviceHandle driver.VkDevice, pCreateInfo *driver.VkRenderPassCreateInfo, pAllocator *driver.VkAllocationCallbacks, pRenderPass *driver.VkRenderPass) (common.VkResult, error) {
			*pRenderPass = renderPassHandle

			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(38), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())

			require.Equal(t, uint64(2), val.FieldByName("attachmentCount").Uint())
			require.Equal(t, uint64(1), val.FieldByName("subpassCount").Uint())
			require.Equal(t, uint64(3), val.FieldByName("dependencyCount").Uint())

			attachmentsPtr := (*driver.VkAttachmentDescription)(unsafe.Pointer(val.FieldByName("pAttachments").Elem().UnsafeAddr()))
			attachmentsSlice := reflect.ValueOf(([]driver.VkAttachmentDescription)(unsafe.Slice(attachmentsPtr, 2)))

			attachment := attachmentsSlice.Index(0)
			require.Equal(t, uint64(1), attachment.FieldByName("flags").Uint())          // VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT
			require.Equal(t, uint64(69), attachment.FieldByName("format").Uint())        // VK_FORMAT_A2B10G10R10_SINT_PACK32
			require.Equal(t, uint64(4), attachment.FieldByName("samples").Uint())        // VK_SAMPLE_COUNT_4_BIT
			require.Equal(t, uint64(1), attachment.FieldByName("loadOp").Uint())         // VK_ATTACHMENT_LOAD_OP_CLEAR
			require.Equal(t, uint64(0), attachment.FieldByName("storeOp").Uint())        // VK_ATTACHMENT_STORE_OP_STORE
			require.Equal(t, uint64(2), attachment.FieldByName("stencilLoadOp").Uint())  // VK_ATTACHMENT_LOAD_OP_DONT_CARE
			require.Equal(t, uint64(1), attachment.FieldByName("stencilStoreOp").Uint()) // VK_ATTACHMENT_STORE_OP_DONT_CARE
			require.Equal(t, uint64(2), attachment.FieldByName("initialLayout").Uint())  // VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
			require.Equal(t, uint64(4), attachment.FieldByName("finalLayout").Uint())    // VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL

			attachment = attachmentsSlice.Index(1)
			require.Equal(t, uint64(0), attachment.FieldByName("flags").Uint())
			require.Equal(t, uint64(63), attachment.FieldByName("format").Uint())        // VK_FORMAT_A2R10G10B10_SINT_PACK32
			require.Equal(t, uint64(64), attachment.FieldByName("samples").Uint())       // VK_SAMPLE_COUNT_64_BIT
			require.Equal(t, uint64(0), attachment.FieldByName("loadOp").Uint())         // VK_ATTACHMENT_LOAD_OP_LOAD
			require.Equal(t, uint64(1), attachment.FieldByName("storeOp").Uint())        // VK_ATTACHMENT_STORE_OP_DONT_CARE
			require.Equal(t, uint64(1), attachment.FieldByName("stencilLoadOp").Uint())  // VK_ATTACHMENT_LOAD_OP_CLEAR
			require.Equal(t, uint64(0), attachment.FieldByName("stencilStoreOp").Uint()) // VK_ATTACHMENT_STORE_OP_STORE
			require.Equal(t, uint64(1), attachment.FieldByName("initialLayout").Uint())  // VK_IMAGE_LAYOUT_GENERAL
			require.Equal(t, uint64(2), attachment.FieldByName("finalLayout").Uint())    // VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL

			subpassesPtr := (*driver.VkSubpassDescription)(unsafe.Pointer(val.FieldByName("pSubpasses").Elem().UnsafeAddr()))
			subpassesSlice := reflect.ValueOf(([]driver.VkSubpassDescription)(unsafe.Slice(subpassesPtr, 1)))

			subpass := subpassesSlice.Index(0)
			require.Equal(t, uint64(0), subpass.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), subpass.FieldByName("pipelineBindPoint").Uint()) // VK_PIPELINE_BIND_POINT_COMPUTE
			require.Equal(t, uint64(1), subpass.FieldByName("inputAttachmentCount").Uint())
			require.Equal(t, uint64(2), subpass.FieldByName("colorAttachmentCount").Uint())
			require.Equal(t, uint64(1), subpass.FieldByName("preserveAttachmentCount").Uint())

			inputAttachmentPtr := (*driver.VkAttachmentReference)(unsafe.Pointer(subpass.FieldByName("pInputAttachments").Elem().UnsafeAddr()))
			inputAttachmentSlice := reflect.ValueOf(([]driver.VkAttachmentReference)(unsafe.Slice(inputAttachmentPtr, 1)))

			attach := inputAttachmentSlice.Index(0)
			require.Equal(t, uint64(0), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(1), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_GENERAL

			colorAttachmentPtr := (*driver.VkAttachmentReference)(unsafe.Pointer(subpass.FieldByName("pColorAttachments").Elem().UnsafeAddr()))
			colorAttachmentSlice := reflect.ValueOf(([]driver.VkAttachmentReference)(unsafe.Slice(colorAttachmentPtr, 2)))

			attach = colorAttachmentSlice.Index(0)
			require.Equal(t, uint64(1), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(8), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_PREINITIALIZED

			attach = colorAttachmentSlice.Index(1)
			require.Equal(t, uint64(2), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(3), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL

			resolveAttachmentPtr := (*driver.VkAttachmentReference)(unsafe.Pointer(subpass.FieldByName("pResolveAttachments").Elem().UnsafeAddr()))
			resolveAttachmentSlice := reflect.ValueOf(([]driver.VkAttachmentReference)(unsafe.Slice(resolveAttachmentPtr, 2)))

			attach = resolveAttachmentSlice.Index(0)
			require.Equal(t, uint64(3), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(4), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL

			attach = resolveAttachmentSlice.Index(1)
			require.Equal(t, uint64(5), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(5), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL

			attach = reflect.ValueOf((*driver.VkAttachmentReference)(unsafe.Pointer(subpass.FieldByName("pDepthStencilAttachment").Elem().UnsafeAddr()))).Elem()
			require.Equal(t, uint64(11), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(6), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL

			preservePtr := (*driver.Uint32)(unsafe.Pointer(subpass.FieldByName("pPreserveAttachments").Elem().UnsafeAddr()))
			preserveSlice := reflect.ValueOf(([]driver.Uint32)(unsafe.Slice(preservePtr, 1)))
			require.Equal(t, uint64(17), preserveSlice.Index(0).Uint())

			dependencyPtr := (*driver.VkSubpassDependency)(unsafe.Pointer(val.FieldByName("pDependencies").Elem().UnsafeAddr()))
			dependencySlice := reflect.ValueOf(([]driver.VkSubpassDependency)(unsafe.Slice(dependencyPtr, 3)))

			dependency := dependencySlice.Index(0)
			require.Equal(t, uint64(0x00000001), dependency.FieldByName("dependencyFlags").Uint()) // VK_DEPENDENCY_BY_REGION_BIT
			require.Equal(t, uint64(17), dependency.FieldByName("srcSubpass").Uint())
			require.Equal(t, uint64(19), dependency.FieldByName("dstSubpass").Uint())
			require.Equal(t, uint64(8), dependency.FieldByName("srcStageMask").Uint())           // VK_PIPELINE_STAGE_VERTEX_SHADER_BIT
			require.Equal(t, uint64(0x00002000), dependency.FieldByName("dstStageMask").Uint())  // VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
			require.Equal(t, uint64(0x00000001), dependency.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_INDIRECT_COMMAND_READ_BIT
			require.Equal(t, uint64(0x00000040), dependency.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_SHADER_WRITE_BIT

			dependency = dependencySlice.Index(1)
			require.Equal(t, uint64(0), dependency.FieldByName("dependencyFlags").Uint())
			require.Equal(t, uint64(23), dependency.FieldByName("srcSubpass").Uint())
			require.Equal(t, uint64(29), dependency.FieldByName("dstSubpass").Uint())
			require.Equal(t, uint64(0x00000100), dependency.FieldByName("srcStageMask").Uint())  // VK_PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT
			require.Equal(t, uint64(0x00002000), dependency.FieldByName("dstStageMask").Uint())  // VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
			require.Equal(t, uint64(0x00000080), dependency.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_COLOR_ATTACHMENT_READ_BIT
			require.Equal(t, uint64(0x00000008), dependency.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_UNIFORM_READ_BIT

			dependency = dependencySlice.Index(2)
			require.Equal(t, uint64(0), dependency.FieldByName("dependencyFlags").Uint())
			require.Equal(t, uint64(31), dependency.FieldByName("srcSubpass").Uint())
			require.Equal(t, uint64(37), dependency.FieldByName("dstSubpass").Uint())
			require.Equal(t, uint64(0x00002000), dependency.FieldByName("srcStageMask").Uint())  // VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
			require.Equal(t, uint64(0x00008000), dependency.FieldByName("dstStageMask").Uint())  // VK_PIPELINE_STAGE_ALL_GRAPHICS_BIT
			require.Equal(t, uint64(0x00000400), dependency.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT
			require.Equal(t, uint64(2), dependency.FieldByName("dstAccessMask").Uint())          // VK_ACCESS_INDEX_READ_BIT

			return core1_0.VKSuccess, nil
		})

	renderPass, _, err := loader.CreateRenderPass(device, nil, core1_0.RenderPassOptions{
		Flags: 0,
		Attachments: []core1_0.AttachmentDescription{
			{
				Flags:          core1_0.AttachmentMayAlias,
				Format:         core1_0.DataFormatA2B10G10R10SignedIntPacked,
				Samples:        core1_0.Samples4,
				LoadOp:         core1_0.LoadOpClear,
				StoreOp:        core1_0.StoreOpStore,
				StencilLoadOp:  core1_0.LoadOpDontCare,
				StencilStoreOp: core1_0.StoreOpDontCare,
				InitialLayout:  core1_0.ImageLayoutColorAttachmentOptimal,
				FinalLayout:    core1_0.ImageLayoutDepthStencilReadOnlyOptimal,
			},
			{
				Flags:          0,
				Format:         core1_0.DataFormatA2R10G10B10SignedIntPacked,
				Samples:        core1_0.Samples64,
				LoadOp:         core1_0.LoadOpLoad,
				StoreOp:        core1_0.StoreOpDontCare,
				StencilLoadOp:  core1_0.LoadOpClear,
				StencilStoreOp: core1_0.StoreOpStore,
				InitialLayout:  core1_0.ImageLayoutGeneral,
				FinalLayout:    core1_0.ImageLayoutColorAttachmentOptimal,
			},
		},
		SubPasses: []core1_0.SubPass{
			{
				Flags:     0,
				BindPoint: core1_0.BindCompute,
				InputAttachments: []common.AttachmentReference{
					{
						AttachmentIndex: 0,
						Layout:          core1_0.ImageLayoutGeneral,
					},
				},
				ColorAttachments: []common.AttachmentReference{
					{
						AttachmentIndex: 1,
						Layout:          core1_0.ImageLayoutPreInitialized,
					},
					{
						AttachmentIndex: 2,
						Layout:          core1_0.ImageLayoutDepthStencilAttachmentOptimal,
					},
				},
				ResolveAttachments: []common.AttachmentReference{
					{
						AttachmentIndex: 3,
						Layout:          core1_0.ImageLayoutDepthStencilReadOnlyOptimal,
					},
					{
						AttachmentIndex: 5,
						Layout:          core1_0.ImageLayoutShaderReadOnlyOptimal,
					},
				},
				DepthStencilAttachment: &common.AttachmentReference{
					AttachmentIndex: 11,
					Layout:          core1_0.ImageLayoutTransferSrcOptimal,
				},
				PreservedAttachmentIndices: []int{17},
			},
		},
		SubPassDependencies: []core1_0.SubPassDependency{
			{
				Flags:           core1_0.DependencyByRegion,
				SrcSubPassIndex: 17,
				DstSubPassIndex: 19,
				SrcStageMask:    core1_0.PipelineStageVertexShader,
				DstStageMask:    core1_0.PipelineStageBottomOfPipe,
				SrcAccessMask:   core1_0.AccessIndirectCommandRead,
				DstAccessMask:   core1_0.AccessShaderWrite,
			},
			{
				Flags:           0,
				SrcSubPassIndex: 23,
				DstSubPassIndex: 29,
				SrcStageMask:    core1_0.PipelineStageEarlyFragmentTests,
				DstStageMask:    core1_0.PipelineStageBottomOfPipe,
				SrcAccessMask:   core1_0.AccessColorAttachmentRead,
				DstAccessMask:   core1_0.AccessUniformRead,
			},
			{
				Flags:           0,
				SrcSubPassIndex: 31,
				DstSubPassIndex: 37,
				SrcStageMask:    core1_0.PipelineStageBottomOfPipe,
				DstStageMask:    core1_0.PipelineStageAllGraphics,
				SrcAccessMask:   core1_0.AccessDepthStencilAttachmentWrite,
				DstAccessMask:   core1_0.AccessIndexRead,
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, renderPass)
	require.Equal(t, renderPassHandle, renderPass.Handle())
}

func TestVulkanLoader1_0_CreateRenderPass_SuccessNoNonColorAttachments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	renderPassHandle := mocks.NewFakeRenderPassHandle()

	mockDriver.EXPECT().VkCreateRenderPass(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(deviceHandle driver.VkDevice, pCreateInfo *driver.VkRenderPassCreateInfo, pAllocator *driver.VkAllocationCallbacks, pRenderPass *driver.VkRenderPass) (common.VkResult, error) {
			*pRenderPass = renderPassHandle

			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(38), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())

			require.Equal(t, uint64(2), val.FieldByName("attachmentCount").Uint())
			require.Equal(t, uint64(1), val.FieldByName("subpassCount").Uint())
			require.Equal(t, uint64(3), val.FieldByName("dependencyCount").Uint())

			attachmentsPtr := (*driver.VkAttachmentDescription)(unsafe.Pointer(val.FieldByName("pAttachments").Elem().UnsafeAddr()))
			attachmentsSlice := reflect.ValueOf(([]driver.VkAttachmentDescription)(unsafe.Slice(attachmentsPtr, 2)))

			attachment := attachmentsSlice.Index(0)
			require.Equal(t, uint64(1), attachment.FieldByName("flags").Uint())          // VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT
			require.Equal(t, uint64(69), attachment.FieldByName("format").Uint())        // VK_FORMAT_A2B10G10R10_SINT_PACK32
			require.Equal(t, uint64(4), attachment.FieldByName("samples").Uint())        // VK_SAMPLE_COUNT_4_BIT
			require.Equal(t, uint64(1), attachment.FieldByName("loadOp").Uint())         // VK_ATTACHMENT_LOAD_OP_CLEAR
			require.Equal(t, uint64(0), attachment.FieldByName("storeOp").Uint())        // VK_ATTACHMENT_STORE_OP_STORE
			require.Equal(t, uint64(2), attachment.FieldByName("stencilLoadOp").Uint())  // VK_ATTACHMENT_LOAD_OP_DONT_CARE
			require.Equal(t, uint64(1), attachment.FieldByName("stencilStoreOp").Uint()) // VK_ATTACHMENT_STORE_OP_DONT_CARE
			require.Equal(t, uint64(2), attachment.FieldByName("initialLayout").Uint())  // VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
			require.Equal(t, uint64(8), attachment.FieldByName("finalLayout").Uint())    //VK_IMAGE_LAYOUT_PREINITIALIZED

			attachment = attachmentsSlice.Index(1)
			require.Equal(t, uint64(0), attachment.FieldByName("flags").Uint())
			require.Equal(t, uint64(63), attachment.FieldByName("format").Uint())        // VK_FORMAT_A2R10G10B10_SINT_PACK32
			require.Equal(t, uint64(64), attachment.FieldByName("samples").Uint())       // VK_SAMPLE_COUNT_64_BIT
			require.Equal(t, uint64(0), attachment.FieldByName("loadOp").Uint())         // VK_ATTACHMENT_LOAD_OP_LOAD
			require.Equal(t, uint64(1), attachment.FieldByName("storeOp").Uint())        // VK_ATTACHMENT_STORE_OP_DONT_CARE
			require.Equal(t, uint64(1), attachment.FieldByName("stencilLoadOp").Uint())  // VK_ATTACHMENT_LOAD_OP_CLEAR
			require.Equal(t, uint64(0), attachment.FieldByName("stencilStoreOp").Uint()) // VK_ATTACHMENT_STORE_OP_STORE
			require.Equal(t, uint64(1), attachment.FieldByName("initialLayout").Uint())  // VK_IMAGE_LAYOUT_GENERAL
			require.Equal(t, uint64(2), attachment.FieldByName("finalLayout").Uint())    // VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL

			subpassesPtr := (*driver.VkSubpassDescription)(unsafe.Pointer(val.FieldByName("pSubpasses").Elem().UnsafeAddr()))
			subpassesSlice := reflect.ValueOf(([]driver.VkSubpassDescription)(unsafe.Slice(subpassesPtr, 1)))

			subpass := subpassesSlice.Index(0)
			require.Equal(t, uint64(0), subpass.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), subpass.FieldByName("pipelineBindPoint").Uint()) // VK_PIPELINE_BIND_POINT_COMPUTE
			require.Equal(t, uint64(1), subpass.FieldByName("inputAttachmentCount").Uint())
			require.Equal(t, uint64(2), subpass.FieldByName("colorAttachmentCount").Uint())
			require.Equal(t, uint64(1), subpass.FieldByName("preserveAttachmentCount").Uint())

			inputAttachmentPtr := (*driver.VkAttachmentReference)(unsafe.Pointer(subpass.FieldByName("pInputAttachments").Elem().UnsafeAddr()))
			inputAttachmentSlice := reflect.ValueOf(([]driver.VkAttachmentReference)(unsafe.Slice(inputAttachmentPtr, 1)))

			attach := inputAttachmentSlice.Index(0)
			require.Equal(t, uint64(0), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(1), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_GENERAL

			colorAttachmentPtr := (*driver.VkAttachmentReference)(unsafe.Pointer(subpass.FieldByName("pColorAttachments").Elem().UnsafeAddr()))
			colorAttachmentSlice := reflect.ValueOf(([]driver.VkAttachmentReference)(unsafe.Slice(colorAttachmentPtr, 2)))

			attach = colorAttachmentSlice.Index(0)
			require.Equal(t, uint64(1), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(7), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL

			attach = colorAttachmentSlice.Index(1)
			require.Equal(t, uint64(2), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(3), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL

			require.True(t, subpass.FieldByName("pResolveAttachments").IsNil())
			require.True(t, subpass.FieldByName("pDepthStencilAttachment").IsNil())

			preservePtr := (*driver.Uint32)(unsafe.Pointer(subpass.FieldByName("pPreserveAttachments").Elem().UnsafeAddr()))
			preserveSlice := reflect.ValueOf(([]driver.Uint32)(unsafe.Slice(preservePtr, 1)))
			require.Equal(t, uint64(17), preserveSlice.Index(0).Uint())

			dependencyPtr := (*driver.VkSubpassDependency)(unsafe.Pointer(val.FieldByName("pDependencies").Elem().UnsafeAddr()))
			dependencySlice := reflect.ValueOf(([]driver.VkSubpassDependency)(unsafe.Slice(dependencyPtr, 3)))

			dependency := dependencySlice.Index(0)
			require.Equal(t, uint64(1), dependency.FieldByName("dependencyFlags").Uint()) // VK_DEPENDENCY_BY_REGION_BIT
			require.Equal(t, uint64(17), dependency.FieldByName("srcSubpass").Uint())
			require.Equal(t, uint64(19), dependency.FieldByName("dstSubpass").Uint())
			require.Equal(t, uint64(8), dependency.FieldByName("srcStageMask").Uint())           // VK_PIPELINE_STAGE_VERTEX_SHADER_BIT
			require.Equal(t, uint64(0x00000800), dependency.FieldByName("dstStageMask").Uint())  // VK_PIPELINE_STAGE_COMPUTE_SHADER_BIT
			require.Equal(t, uint64(0x00000004), dependency.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_VERTEX_ATTRIBUTE_READ_BIT
			require.Equal(t, uint64(0x00000100), dependency.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_COLOR_ATTACHMENT_WRITE_BIT

			dependency = dependencySlice.Index(1)
			require.Equal(t, uint64(0), dependency.FieldByName("dependencyFlags").Uint())
			require.Equal(t, uint64(23), dependency.FieldByName("srcSubpass").Uint())
			require.Equal(t, uint64(29), dependency.FieldByName("dstSubpass").Uint())
			require.Equal(t, uint64(0x00000002), dependency.FieldByName("srcStageMask").Uint())  // VK_PIPELINE_STAGE_DRAW_INDIRECT_BIT
			require.Equal(t, uint64(0x00002000), dependency.FieldByName("dstStageMask").Uint())  // VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
			require.Equal(t, uint64(0x00000080), dependency.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_COLOR_ATTACHMENT_READ_BIT
			require.Equal(t, uint64(0x00002000), dependency.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_HOST_READ_BIT

			dependency = dependencySlice.Index(2)
			require.Equal(t, uint64(0), dependency.FieldByName("dependencyFlags").Uint())
			require.Equal(t, uint64(31), dependency.FieldByName("srcSubpass").Uint())
			require.Equal(t, uint64(37), dependency.FieldByName("dstSubpass").Uint())
			require.Equal(t, uint64(0x00002000), dependency.FieldByName("srcStageMask").Uint())  // VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
			require.Equal(t, uint64(0x00008000), dependency.FieldByName("dstStageMask").Uint())  // VK_PIPELINE_STAGE_ALL_GRAPHICS_BIT
			require.Equal(t, uint64(0x00000400), dependency.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT
			require.Equal(t, uint64(2), dependency.FieldByName("dstAccessMask").Uint())          // VK_ACCESS_INDEX_READ_BIT

			return core1_0.VKSuccess, nil
		})

	renderPass, _, err := loader.CreateRenderPass(device, nil, core1_0.RenderPassOptions{
		Flags: 0,
		Attachments: []core1_0.AttachmentDescription{
			{
				Flags:          core1_0.AttachmentMayAlias,
				Format:         core1_0.DataFormatA2B10G10R10SignedIntPacked,
				Samples:        core1_0.Samples4,
				LoadOp:         core1_0.LoadOpClear,
				StoreOp:        core1_0.StoreOpStore,
				StencilLoadOp:  core1_0.LoadOpDontCare,
				StencilStoreOp: core1_0.StoreOpDontCare,
				InitialLayout:  core1_0.ImageLayoutColorAttachmentOptimal,
				FinalLayout:    core1_0.ImageLayoutPreInitialized,
			},
			{
				Flags:          0,
				Format:         core1_0.DataFormatA2R10G10B10SignedIntPacked,
				Samples:        core1_0.Samples64,
				LoadOp:         core1_0.LoadOpLoad,
				StoreOp:        core1_0.StoreOpDontCare,
				StencilLoadOp:  core1_0.LoadOpClear,
				StencilStoreOp: core1_0.StoreOpStore,
				InitialLayout:  core1_0.ImageLayoutGeneral,
				FinalLayout:    core1_0.ImageLayoutColorAttachmentOptimal,
			},
		},
		SubPasses: []core1_0.SubPass{
			{
				Flags:     0,
				BindPoint: core1_0.BindCompute,
				InputAttachments: []common.AttachmentReference{
					{
						AttachmentIndex: 0,
						Layout:          core1_0.ImageLayoutGeneral,
					},
				},
				ColorAttachments: []common.AttachmentReference{
					{
						AttachmentIndex: 1,
						Layout:          core1_0.ImageLayoutTransferDstOptimal,
					},
					{
						AttachmentIndex: 2,
						Layout:          core1_0.ImageLayoutDepthStencilAttachmentOptimal,
					},
				},
				ResolveAttachments:         []common.AttachmentReference{},
				PreservedAttachmentIndices: []int{17},
			},
		},
		SubPassDependencies: []core1_0.SubPassDependency{
			{
				Flags:           core1_0.DependencyByRegion,
				SrcSubPassIndex: 17,
				DstSubPassIndex: 19,
				SrcStageMask:    core1_0.PipelineStageVertexShader,
				DstStageMask:    core1_0.PipelineStageComputeShader,
				SrcAccessMask:   core1_0.AccessVertexAttributeRead,
				DstAccessMask:   core1_0.AccessColorAttachmentWrite,
			},
			{
				Flags:           0,
				SrcSubPassIndex: 23,
				DstSubPassIndex: 29,
				SrcStageMask:    core1_0.PipelineStageDrawIndirect,
				DstStageMask:    core1_0.PipelineStageBottomOfPipe,
				SrcAccessMask:   core1_0.AccessColorAttachmentRead,
				DstAccessMask:   core1_0.AccessHostRead,
			},
			{
				Flags:           0,
				SrcSubPassIndex: 31,
				DstSubPassIndex: 37,
				SrcStageMask:    core1_0.PipelineStageBottomOfPipe,
				DstStageMask:    core1_0.PipelineStageAllGraphics,
				SrcAccessMask:   core1_0.AccessDepthStencilAttachmentWrite,
				DstAccessMask:   core1_0.AccessIndexRead,
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, renderPass)
	require.Equal(t, renderPassHandle, renderPass.Handle())
}

func TestVulkanLoader1_0_CreateRenderPass_MismatchResolve(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)

	_, _, err = loader.CreateRenderPass(device, nil, core1_0.RenderPassOptions{
		Flags: 0,
		Attachments: []core1_0.AttachmentDescription{
			{
				Flags:          core1_0.AttachmentMayAlias,
				Format:         core1_0.DataFormatA2B10G10R10SignedIntPacked,
				Samples:        core1_0.Samples4,
				LoadOp:         core1_0.LoadOpClear,
				StoreOp:        core1_0.StoreOpStore,
				StencilLoadOp:  core1_0.LoadOpDontCare,
				StencilStoreOp: core1_0.StoreOpDontCare,
				InitialLayout:  core1_0.ImageLayoutColorAttachmentOptimal,
				FinalLayout:    core1_0.ImageLayoutPreInitialized,
			},
			{
				Flags:          0,
				Format:         core1_0.DataFormatA2R10G10B10SignedIntPacked,
				Samples:        core1_0.Samples64,
				LoadOp:         core1_0.LoadOpLoad,
				StoreOp:        core1_0.StoreOpDontCare,
				StencilLoadOp:  core1_0.LoadOpClear,
				StencilStoreOp: core1_0.StoreOpStore,
				InitialLayout:  core1_0.ImageLayoutGeneral,
				FinalLayout:    core1_0.ImageLayoutColorAttachmentOptimal,
			},
		},
		SubPasses: []core1_0.SubPass{
			{
				Flags:     0,
				BindPoint: core1_0.BindCompute,
				InputAttachments: []common.AttachmentReference{
					{
						AttachmentIndex: 0,
						Layout:          core1_0.ImageLayoutGeneral,
					},
				},
				ColorAttachments: []common.AttachmentReference{
					{
						AttachmentIndex: 1,
						Layout:          core1_0.ImageLayoutDepthStencilReadOnlyOptimal,
					},
					{
						AttachmentIndex: 2,
						Layout:          core1_0.ImageLayoutDepthStencilAttachmentOptimal,
					},
				},
				ResolveAttachments: []common.AttachmentReference{
					{
						AttachmentIndex: 3,
						Layout:          core1_0.ImageLayoutDepthStencilReadOnlyOptimal,
					},
					{
						AttachmentIndex: 5,
						Layout:          core1_0.ImageLayoutColorAttachmentOptimal,
					},
					{
						AttachmentIndex: 0,
						Layout:          core1_0.ImageLayoutUndefined,
					},
				},
				DepthStencilAttachment: &common.AttachmentReference{
					AttachmentIndex: 11,
					Layout:          core1_0.ImageLayoutTransferSrcOptimal,
				},
				PreservedAttachmentIndices: []int{17},
			},
		},
		SubPassDependencies: []core1_0.SubPassDependency{
			{
				Flags:           0,
				SrcSubPassIndex: 17,
				DstSubPassIndex: 19,
				SrcStageMask:    core1_0.PipelineStageVertexShader,
				DstStageMask:    core1_0.PipelineStageTessellationEvaluationShader,
				SrcAccessMask:   core1_0.AccessDepthStencilAttachmentRead,
				DstAccessMask:   core1_0.AccessHostWrite,
			},
			{
				Flags:           core1_0.DependencyByRegion,
				SrcSubPassIndex: 23,
				DstSubPassIndex: 29,
				SrcStageMask:    core1_0.PipelineStageLateFragmentTests,
				DstStageMask:    core1_0.PipelineStageBottomOfPipe,
				SrcAccessMask:   core1_0.AccessColorAttachmentRead,
				DstAccessMask:   core1_0.AccessInputAttachmentRead,
			},
			{
				Flags:           0,
				SrcSubPassIndex: 31,
				DstSubPassIndex: 37,
				SrcStageMask:    core1_0.PipelineStageBottomOfPipe,
				DstStageMask:    core1_0.PipelineStageAllGraphics,
				SrcAccessMask:   core1_0.AccessDepthStencilAttachmentWrite,
				DstAccessMask:   core1_0.AccessIndexRead,
			},
		},
	})
	require.EqualError(t, err, "in subpass 0, 2 color attachments are defined, but 3 resolve attachments are defined")
}

func TestVulkanRenderPass_RenderAreaGranularity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	renderPass := internal_mocks.EasyDummyRenderPass(t, loader, device)

	mockDriver.EXPECT().VkGetRenderAreaGranularity(device.Handle(), renderPass.Handle(), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, renderPass driver.VkRenderPass, pGranularity *driver.VkExtent2D) {
			val := reflect.ValueOf(pGranularity).Elem()

			*(*uint32)(unsafe.Pointer(val.FieldByName("width").UnsafeAddr())) = uint32(1)
			*(*uint32)(unsafe.Pointer(val.FieldByName("height").UnsafeAddr())) = uint32(3)
		})

	gran := renderPass.RenderAreaGranularity()
	require.Equal(t, 1, gran.Width)
	require.Equal(t, 3, gran.Height)
}
