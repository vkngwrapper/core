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

func TestVulkanLoader1_0_CreateRenderPass_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	renderPassHandle := mocks.NewFakeRenderPassHandle()

	driver.EXPECT().VkCreateRenderPass(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(deviceHandle core.VkDevice, pCreateInfo *core.VkRenderPassCreateInfo, pAllocator *core.VkAllocationCallbacks, pRenderPass *core.VkRenderPass) (core.VkResult, error) {
			*pRenderPass = renderPassHandle

			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(38), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), val.FieldByName("flags").Uint()) // VK_RENDER_PASS_CREATE_TRANSFORM_BIT_QCOM

			require.Equal(t, uint64(2), val.FieldByName("attachmentCount").Uint())
			require.Equal(t, uint64(1), val.FieldByName("subpassCount").Uint())
			require.Equal(t, uint64(3), val.FieldByName("dependencyCount").Uint())

			attachmentsPtr := (*core.VkAttachmentDescription)(unsafe.Pointer(val.FieldByName("pAttachments").Elem().UnsafeAddr()))
			attachmentsSlice := reflect.ValueOf(([]core.VkAttachmentDescription)(unsafe.Slice(attachmentsPtr, 2)))

			attachment := attachmentsSlice.Index(0)
			require.Equal(t, uint64(1), attachment.FieldByName("flags").Uint())                // VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT
			require.Equal(t, uint64(69), attachment.FieldByName("format").Uint())              // VK_FORMAT_A2B10G10R10_SINT_PACK32
			require.Equal(t, uint64(4), attachment.FieldByName("samples").Uint())              // VK_SAMPLE_COUNT_4_BIT
			require.Equal(t, uint64(1), attachment.FieldByName("loadOp").Uint())               // VK_ATTACHMENT_LOAD_OP_CLEAR
			require.Equal(t, uint64(0), attachment.FieldByName("storeOp").Uint())              // VK_ATTACHMENT_STORE_OP_STORE
			require.Equal(t, uint64(2), attachment.FieldByName("stencilLoadOp").Uint())        // VK_ATTACHMENT_LOAD_OP_DONT_CARE
			require.Equal(t, uint64(1), attachment.FieldByName("stencilStoreOp").Uint())       // VK_ATTACHMENT_STORE_OP_DONT_CARE
			require.Equal(t, uint64(2), attachment.FieldByName("initialLayout").Uint())        // VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
			require.Equal(t, uint64(1000241001), attachment.FieldByName("finalLayout").Uint()) //VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL = )

			attachment = attachmentsSlice.Index(1)
			require.Equal(t, uint64(0), attachment.FieldByName("flags").Uint())
			require.Equal(t, uint64(63), attachment.FieldByName("format").Uint())          // VK_FORMAT_A2R10G10B10_SINT_PACK32
			require.Equal(t, uint64(64), attachment.FieldByName("samples").Uint())         // VK_SAMPLE_COUNT_64_BIT
			require.Equal(t, uint64(0), attachment.FieldByName("loadOp").Uint())           // VK_ATTACHMENT_LOAD_OP_LOAD
			require.Equal(t, uint64(1000301000), attachment.FieldByName("storeOp").Uint()) // VK_ATTACHMENT_STORE_OP_NONE_EXT
			require.Equal(t, uint64(1), attachment.FieldByName("stencilLoadOp").Uint())    // VK_ATTACHMENT_LOAD_OP_CLEAR
			require.Equal(t, uint64(0), attachment.FieldByName("stencilStoreOp").Uint())   // VK_ATTACHMENT_STORE_OP_STORE
			require.Equal(t, uint64(1), attachment.FieldByName("initialLayout").Uint())    // VK_IMAGE_LAYOUT_GENERAL
			require.Equal(t, uint64(2), attachment.FieldByName("finalLayout").Uint())      // VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL

			subpassesPtr := (*core.VkSubpassDescription)(unsafe.Pointer(val.FieldByName("pSubpasses").Elem().UnsafeAddr()))
			subpassesSlice := reflect.ValueOf(([]core.VkSubpassDescription)(unsafe.Slice(subpassesPtr, 1)))

			subpass := subpassesSlice.Index(0)
			require.Equal(t, uint64(8), subpass.FieldByName("flags").Uint())             // VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM
			require.Equal(t, uint64(1), subpass.FieldByName("pipelineBindPoint").Uint()) // VK_PIPELINE_BIND_POINT_COMPUTE
			require.Equal(t, uint64(1), subpass.FieldByName("inputAttachmentCount").Uint())
			require.Equal(t, uint64(2), subpass.FieldByName("colorAttachmentCount").Uint())
			require.Equal(t, uint64(1), subpass.FieldByName("preserveAttachmentCount").Uint())

			inputAttachmentPtr := (*core.VkAttachmentReference)(unsafe.Pointer(subpass.FieldByName("pInputAttachments").Elem().UnsafeAddr()))
			inputAttachmentSlice := reflect.ValueOf(([]core.VkAttachmentReference)(unsafe.Slice(inputAttachmentPtr, 1)))

			attach := inputAttachmentSlice.Index(0)
			require.Equal(t, uint64(0), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(1), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_GENERAL

			colorAttachmentPtr := (*core.VkAttachmentReference)(unsafe.Pointer(subpass.FieldByName("pColorAttachments").Elem().UnsafeAddr()))
			colorAttachmentSlice := reflect.ValueOf(([]core.VkAttachmentReference)(unsafe.Slice(colorAttachmentPtr, 2)))

			attach = colorAttachmentSlice.Index(0)
			require.Equal(t, uint64(1), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(1000241000), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL

			attach = colorAttachmentSlice.Index(1)
			require.Equal(t, uint64(2), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(3), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL

			resolveAttachmentPtr := (*core.VkAttachmentReference)(unsafe.Pointer(subpass.FieldByName("pResolveAttachments").Elem().UnsafeAddr()))
			resolveAttachmentSlice := reflect.ValueOf(([]core.VkAttachmentReference)(unsafe.Slice(resolveAttachmentPtr, 2)))

			attach = resolveAttachmentSlice.Index(0)
			require.Equal(t, uint64(3), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(4), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL

			attach = resolveAttachmentSlice.Index(1)
			require.Equal(t, uint64(5), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(1000117001), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL

			attach = reflect.ValueOf((*core.VkAttachmentReference)(unsafe.Pointer(subpass.FieldByName("pDepthStencilAttachment").Elem().UnsafeAddr()))).Elem()
			require.Equal(t, uint64(11), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(6), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL

			preservePtr := (*core.Uint32)(unsafe.Pointer(subpass.FieldByName("pPreserveAttachments").Elem().UnsafeAddr()))
			preserveSlice := reflect.ValueOf(([]core.Uint32)(unsafe.Slice(preservePtr, 1)))
			require.Equal(t, uint64(17), preserveSlice.Index(0).Uint())

			dependencyPtr := (*core.VkSubpassDependency)(unsafe.Pointer(val.FieldByName("pDependencies").Elem().UnsafeAddr()))
			dependencySlice := reflect.ValueOf(([]core.VkSubpassDependency)(unsafe.Slice(dependencyPtr, 3)))

			dependency := dependencySlice.Index(0)
			require.Equal(t, uint64(4), dependency.FieldByName("dependencyFlags").Uint()) // VK_DEPENDENCY_DEVICE_GROUP_BIT
			require.Equal(t, uint64(17), dependency.FieldByName("srcSubpass").Uint())
			require.Equal(t, uint64(19), dependency.FieldByName("dstSubpass").Uint())
			require.Equal(t, uint64(8), dependency.FieldByName("srcStageMask").Uint())           // VK_PIPELINE_STAGE_VERTEX_SHADER_BIT
			require.Equal(t, uint64(0x02000000), dependency.FieldByName("dstStageMask").Uint())  // VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR
			require.Equal(t, uint64(0x00080000), dependency.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT
			require.Equal(t, uint64(0x00100000), dependency.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_CONDITIONAL_RENDERING_READ_BIT_EXT

			dependency = dependencySlice.Index(1)
			require.Equal(t, uint64(2), dependency.FieldByName("dependencyFlags").Uint()) // VK_DEPENDENCY_VIEW_LOCAL_BIT
			require.Equal(t, uint64(23), dependency.FieldByName("srcSubpass").Uint())
			require.Equal(t, uint64(29), dependency.FieldByName("dstSubpass").Uint())
			require.Equal(t, uint64(0x00040000), dependency.FieldByName("srcStageMask").Uint())  // VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT
			require.Equal(t, uint64(0x00002000), dependency.FieldByName("dstStageMask").Uint())  // VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
			require.Equal(t, uint64(0x00000080), dependency.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_COLOR_ATTACHMENT_READ_BIT
			require.Equal(t, uint64(0x01000000), dependency.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_FRAGMENT_DENSITY_MAP_READ_BIT_EXT

			dependency = dependencySlice.Index(2)
			require.Equal(t, uint64(0), dependency.FieldByName("dependencyFlags").Uint())
			require.Equal(t, uint64(31), dependency.FieldByName("srcSubpass").Uint())
			require.Equal(t, uint64(37), dependency.FieldByName("dstSubpass").Uint())
			require.Equal(t, uint64(0x00002000), dependency.FieldByName("srcStageMask").Uint())  // VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
			require.Equal(t, uint64(0x00008000), dependency.FieldByName("dstStageMask").Uint())  // VK_PIPELINE_STAGE_ALL_GRAPHICS_BIT
			require.Equal(t, uint64(0x00000400), dependency.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT
			require.Equal(t, uint64(2), dependency.FieldByName("dstAccessMask").Uint())          // VK_ACCESS_INDEX_READ_BIT

			return core.VKSuccess, nil
		})

	renderPass, _, err := loader.CreateRenderPass(device, &core.RenderPassOptions{
		Flags: core.RenderPassCreateTransformBitQCOM,
		Attachments: []core.AttachmentDescription{
			{
				Flags:          common.AttachmentMayAlias,
				Format:         common.FormatA2B10G10R10SignedInt,
				Samples:        common.Samples4,
				LoadOp:         common.LoadOpClear,
				StoreOp:        common.StoreOpStore,
				StencilLoadOp:  common.LoadOpDontCare,
				StencilStoreOp: common.StoreOpDontCare,
				InitialLayout:  common.LayoutColorAttachmentOptimal,
				FinalLayout:    common.LayoutDepthReadOnlyOptimal,
			},
			{
				Flags:          0,
				Format:         common.FormatA2R10G10B10SignedInt,
				Samples:        common.Samples64,
				LoadOp:         common.LoadOpLoad,
				StoreOp:        common.StoreOpNoneEXT,
				StencilLoadOp:  common.LoadOpClear,
				StencilStoreOp: common.StoreOpStore,
				InitialLayout:  common.LayoutGeneral,
				FinalLayout:    common.LayoutColorAttachmentOptimal,
			},
		},
		SubPasses: []core.SubPass{
			{
				Flags:     core.SubPassShaderResolveQCOM,
				BindPoint: common.BindCompute,
				InputAttachments: []common.AttachmentReference{
					{
						AttachmentIndex: 0,
						Layout:          common.LayoutGeneral,
					},
				},
				ColorAttachments: []common.AttachmentReference{
					{
						AttachmentIndex: 1,
						Layout:          common.LayoutDepthAttachmentOptimal,
					},
					{
						AttachmentIndex: 2,
						Layout:          common.LayoutDepthStencilAttachmentOptimal,
					},
				},
				ResolveAttachments: []common.AttachmentReference{
					{
						AttachmentIndex: 3,
						Layout:          common.LayoutDepthStencilReadOnlyOptimal,
					},
					{
						AttachmentIndex: 5,
						Layout:          common.LayoutDepthAttachmentStencilReadOnlyOptimal,
					},
				},
				DepthStencilAttachment: &common.AttachmentReference{
					AttachmentIndex: 11,
					Layout:          common.LayoutTransferSrcOptimal,
				},
				PreservedAttachmentIndices: []int{17},
			},
		},
		SubPassDependencies: []core.SubPassDependency{
			{
				Flags:           common.DependencyDeviceGroup,
				SrcSubPassIndex: 17,
				DstSubPassIndex: 19,
				SrcStageMask:    common.PipelineStageVertexShader,
				DstStageMask:    common.PipelineStageAccelerationStructureBuildKHR,
				SrcAccessMask:   common.AccessColorAttachmentReadNonCoherentEXT,
				DstAccessMask:   common.AccessConditionalRenderingReadEXT,
			},
			{
				Flags:           common.DependencyViewLocal,
				SrcSubPassIndex: 23,
				DstSubPassIndex: 29,
				SrcStageMask:    common.PipelineStageConditionalRenderingEXT,
				DstStageMask:    common.PipelineStageBottomOfPipe,
				SrcAccessMask:   common.AccessColorAttachmentRead,
				DstAccessMask:   common.AccessFragmentDensityMapReadEXT,
			},
			{
				Flags:           0,
				SrcSubPassIndex: 31,
				DstSubPassIndex: 37,
				SrcStageMask:    common.PipelineStageBottomOfPipe,
				DstStageMask:    common.PipelineStageAllGraphics,
				SrcAccessMask:   common.AccessDepthStencilAttachmentWrite,
				DstAccessMask:   common.AccessIndexRead,
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

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	renderPassHandle := mocks.NewFakeRenderPassHandle()

	driver.EXPECT().VkCreateRenderPass(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(deviceHandle core.VkDevice, pCreateInfo *core.VkRenderPassCreateInfo, pAllocator *core.VkAllocationCallbacks, pRenderPass *core.VkRenderPass) (core.VkResult, error) {
			*pRenderPass = renderPassHandle

			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(38), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), val.FieldByName("flags").Uint()) // VK_RENDER_PASS_CREATE_TRANSFORM_BIT_QCOM

			require.Equal(t, uint64(2), val.FieldByName("attachmentCount").Uint())
			require.Equal(t, uint64(1), val.FieldByName("subpassCount").Uint())
			require.Equal(t, uint64(3), val.FieldByName("dependencyCount").Uint())

			attachmentsPtr := (*core.VkAttachmentDescription)(unsafe.Pointer(val.FieldByName("pAttachments").Elem().UnsafeAddr()))
			attachmentsSlice := reflect.ValueOf(([]core.VkAttachmentDescription)(unsafe.Slice(attachmentsPtr, 2)))

			attachment := attachmentsSlice.Index(0)
			require.Equal(t, uint64(1), attachment.FieldByName("flags").Uint())                // VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT
			require.Equal(t, uint64(69), attachment.FieldByName("format").Uint())              // VK_FORMAT_A2B10G10R10_SINT_PACK32
			require.Equal(t, uint64(4), attachment.FieldByName("samples").Uint())              // VK_SAMPLE_COUNT_4_BIT
			require.Equal(t, uint64(1), attachment.FieldByName("loadOp").Uint())               // VK_ATTACHMENT_LOAD_OP_CLEAR
			require.Equal(t, uint64(0), attachment.FieldByName("storeOp").Uint())              // VK_ATTACHMENT_STORE_OP_STORE
			require.Equal(t, uint64(2), attachment.FieldByName("stencilLoadOp").Uint())        // VK_ATTACHMENT_LOAD_OP_DONT_CARE
			require.Equal(t, uint64(1), attachment.FieldByName("stencilStoreOp").Uint())       // VK_ATTACHMENT_STORE_OP_DONT_CARE
			require.Equal(t, uint64(2), attachment.FieldByName("initialLayout").Uint())        // VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
			require.Equal(t, uint64(1000241001), attachment.FieldByName("finalLayout").Uint()) //VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL = )

			attachment = attachmentsSlice.Index(1)
			require.Equal(t, uint64(0), attachment.FieldByName("flags").Uint())
			require.Equal(t, uint64(63), attachment.FieldByName("format").Uint())          // VK_FORMAT_A2R10G10B10_SINT_PACK32
			require.Equal(t, uint64(64), attachment.FieldByName("samples").Uint())         // VK_SAMPLE_COUNT_64_BIT
			require.Equal(t, uint64(0), attachment.FieldByName("loadOp").Uint())           // VK_ATTACHMENT_LOAD_OP_LOAD
			require.Equal(t, uint64(1000301000), attachment.FieldByName("storeOp").Uint()) // VK_ATTACHMENT_STORE_OP_NONE_EXT
			require.Equal(t, uint64(1), attachment.FieldByName("stencilLoadOp").Uint())    // VK_ATTACHMENT_LOAD_OP_CLEAR
			require.Equal(t, uint64(0), attachment.FieldByName("stencilStoreOp").Uint())   // VK_ATTACHMENT_STORE_OP_STORE
			require.Equal(t, uint64(1), attachment.FieldByName("initialLayout").Uint())    // VK_IMAGE_LAYOUT_GENERAL
			require.Equal(t, uint64(2), attachment.FieldByName("finalLayout").Uint())      // VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL

			subpassesPtr := (*core.VkSubpassDescription)(unsafe.Pointer(val.FieldByName("pSubpasses").Elem().UnsafeAddr()))
			subpassesSlice := reflect.ValueOf(([]core.VkSubpassDescription)(unsafe.Slice(subpassesPtr, 1)))

			subpass := subpassesSlice.Index(0)
			require.Equal(t, uint64(8), subpass.FieldByName("flags").Uint())             // VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM
			require.Equal(t, uint64(1), subpass.FieldByName("pipelineBindPoint").Uint()) // VK_PIPELINE_BIND_POINT_COMPUTE
			require.Equal(t, uint64(1), subpass.FieldByName("inputAttachmentCount").Uint())
			require.Equal(t, uint64(2), subpass.FieldByName("colorAttachmentCount").Uint())
			require.Equal(t, uint64(1), subpass.FieldByName("preserveAttachmentCount").Uint())

			inputAttachmentPtr := (*core.VkAttachmentReference)(unsafe.Pointer(subpass.FieldByName("pInputAttachments").Elem().UnsafeAddr()))
			inputAttachmentSlice := reflect.ValueOf(([]core.VkAttachmentReference)(unsafe.Slice(inputAttachmentPtr, 1)))

			attach := inputAttachmentSlice.Index(0)
			require.Equal(t, uint64(0), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(1), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_GENERAL

			colorAttachmentPtr := (*core.VkAttachmentReference)(unsafe.Pointer(subpass.FieldByName("pColorAttachments").Elem().UnsafeAddr()))
			colorAttachmentSlice := reflect.ValueOf(([]core.VkAttachmentReference)(unsafe.Slice(colorAttachmentPtr, 2)))

			attach = colorAttachmentSlice.Index(0)
			require.Equal(t, uint64(1), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(1000241000), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL

			attach = colorAttachmentSlice.Index(1)
			require.Equal(t, uint64(2), attach.FieldByName("attachment").Uint())
			require.Equal(t, uint64(3), attach.FieldByName("layout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL

			require.True(t, subpass.FieldByName("pResolveAttachments").IsNil())
			require.True(t, subpass.FieldByName("pDepthStencilAttachment").IsNil())

			preservePtr := (*core.Uint32)(unsafe.Pointer(subpass.FieldByName("pPreserveAttachments").Elem().UnsafeAddr()))
			preserveSlice := reflect.ValueOf(([]core.Uint32)(unsafe.Slice(preservePtr, 1)))
			require.Equal(t, uint64(17), preserveSlice.Index(0).Uint())

			dependencyPtr := (*core.VkSubpassDependency)(unsafe.Pointer(val.FieldByName("pDependencies").Elem().UnsafeAddr()))
			dependencySlice := reflect.ValueOf(([]core.VkSubpassDependency)(unsafe.Slice(dependencyPtr, 3)))

			dependency := dependencySlice.Index(0)
			require.Equal(t, uint64(4), dependency.FieldByName("dependencyFlags").Uint()) // VK_DEPENDENCY_DEVICE_GROUP_BIT
			require.Equal(t, uint64(17), dependency.FieldByName("srcSubpass").Uint())
			require.Equal(t, uint64(19), dependency.FieldByName("dstSubpass").Uint())
			require.Equal(t, uint64(8), dependency.FieldByName("srcStageMask").Uint())           // VK_PIPELINE_STAGE_VERTEX_SHADER_BIT
			require.Equal(t, uint64(0x02000000), dependency.FieldByName("dstStageMask").Uint())  // VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR
			require.Equal(t, uint64(0x00080000), dependency.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT
			require.Equal(t, uint64(0x00100000), dependency.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_CONDITIONAL_RENDERING_READ_BIT_EXT

			dependency = dependencySlice.Index(1)
			require.Equal(t, uint64(2), dependency.FieldByName("dependencyFlags").Uint()) // VK_DEPENDENCY_VIEW_LOCAL_BIT
			require.Equal(t, uint64(23), dependency.FieldByName("srcSubpass").Uint())
			require.Equal(t, uint64(29), dependency.FieldByName("dstSubpass").Uint())
			require.Equal(t, uint64(0x00040000), dependency.FieldByName("srcStageMask").Uint())  // VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT
			require.Equal(t, uint64(0x00002000), dependency.FieldByName("dstStageMask").Uint())  // VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
			require.Equal(t, uint64(0x00000080), dependency.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_COLOR_ATTACHMENT_READ_BIT
			require.Equal(t, uint64(0x01000000), dependency.FieldByName("dstAccessMask").Uint()) // VK_ACCESS_FRAGMENT_DENSITY_MAP_READ_BIT_EXT

			dependency = dependencySlice.Index(2)
			require.Equal(t, uint64(0), dependency.FieldByName("dependencyFlags").Uint())
			require.Equal(t, uint64(31), dependency.FieldByName("srcSubpass").Uint())
			require.Equal(t, uint64(37), dependency.FieldByName("dstSubpass").Uint())
			require.Equal(t, uint64(0x00002000), dependency.FieldByName("srcStageMask").Uint())  // VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
			require.Equal(t, uint64(0x00008000), dependency.FieldByName("dstStageMask").Uint())  // VK_PIPELINE_STAGE_ALL_GRAPHICS_BIT
			require.Equal(t, uint64(0x00000400), dependency.FieldByName("srcAccessMask").Uint()) // VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT
			require.Equal(t, uint64(2), dependency.FieldByName("dstAccessMask").Uint())          // VK_ACCESS_INDEX_READ_BIT

			return core.VKSuccess, nil
		})

	renderPass, _, err := loader.CreateRenderPass(device, &core.RenderPassOptions{
		Flags: core.RenderPassCreateTransformBitQCOM,
		Attachments: []core.AttachmentDescription{
			{
				Flags:          common.AttachmentMayAlias,
				Format:         common.FormatA2B10G10R10SignedInt,
				Samples:        common.Samples4,
				LoadOp:         common.LoadOpClear,
				StoreOp:        common.StoreOpStore,
				StencilLoadOp:  common.LoadOpDontCare,
				StencilStoreOp: common.StoreOpDontCare,
				InitialLayout:  common.LayoutColorAttachmentOptimal,
				FinalLayout:    common.LayoutDepthReadOnlyOptimal,
			},
			{
				Flags:          0,
				Format:         common.FormatA2R10G10B10SignedInt,
				Samples:        common.Samples64,
				LoadOp:         common.LoadOpLoad,
				StoreOp:        common.StoreOpNoneEXT,
				StencilLoadOp:  common.LoadOpClear,
				StencilStoreOp: common.StoreOpStore,
				InitialLayout:  common.LayoutGeneral,
				FinalLayout:    common.LayoutColorAttachmentOptimal,
			},
		},
		SubPasses: []core.SubPass{
			{
				Flags:     core.SubPassShaderResolveQCOM,
				BindPoint: common.BindCompute,
				InputAttachments: []common.AttachmentReference{
					{
						AttachmentIndex: 0,
						Layout:          common.LayoutGeneral,
					},
				},
				ColorAttachments: []common.AttachmentReference{
					{
						AttachmentIndex: 1,
						Layout:          common.LayoutDepthAttachmentOptimal,
					},
					{
						AttachmentIndex: 2,
						Layout:          common.LayoutDepthStencilAttachmentOptimal,
					},
				},
				ResolveAttachments:         []common.AttachmentReference{},
				PreservedAttachmentIndices: []int{17},
			},
		},
		SubPassDependencies: []core.SubPassDependency{
			{
				Flags:           common.DependencyDeviceGroup,
				SrcSubPassIndex: 17,
				DstSubPassIndex: 19,
				SrcStageMask:    common.PipelineStageVertexShader,
				DstStageMask:    common.PipelineStageAccelerationStructureBuildKHR,
				SrcAccessMask:   common.AccessColorAttachmentReadNonCoherentEXT,
				DstAccessMask:   common.AccessConditionalRenderingReadEXT,
			},
			{
				Flags:           common.DependencyViewLocal,
				SrcSubPassIndex: 23,
				DstSubPassIndex: 29,
				SrcStageMask:    common.PipelineStageConditionalRenderingEXT,
				DstStageMask:    common.PipelineStageBottomOfPipe,
				SrcAccessMask:   common.AccessColorAttachmentRead,
				DstAccessMask:   common.AccessFragmentDensityMapReadEXT,
			},
			{
				Flags:           0,
				SrcSubPassIndex: 31,
				DstSubPassIndex: 37,
				SrcStageMask:    common.PipelineStageBottomOfPipe,
				DstStageMask:    common.PipelineStageAllGraphics,
				SrcAccessMask:   common.AccessDepthStencilAttachmentWrite,
				DstAccessMask:   common.AccessIndexRead,
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

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)

	_, _, err = loader.CreateRenderPass(device, &core.RenderPassOptions{
		Flags: core.RenderPassCreateTransformBitQCOM,
		Attachments: []core.AttachmentDescription{
			{
				Flags:          common.AttachmentMayAlias,
				Format:         common.FormatA2B10G10R10SignedInt,
				Samples:        common.Samples4,
				LoadOp:         common.LoadOpClear,
				StoreOp:        common.StoreOpStore,
				StencilLoadOp:  common.LoadOpDontCare,
				StencilStoreOp: common.StoreOpDontCare,
				InitialLayout:  common.LayoutColorAttachmentOptimal,
				FinalLayout:    common.LayoutDepthReadOnlyOptimal,
			},
			{
				Flags:          0,
				Format:         common.FormatA2R10G10B10SignedInt,
				Samples:        common.Samples64,
				LoadOp:         common.LoadOpLoad,
				StoreOp:        common.StoreOpNoneEXT,
				StencilLoadOp:  common.LoadOpClear,
				StencilStoreOp: common.StoreOpStore,
				InitialLayout:  common.LayoutGeneral,
				FinalLayout:    common.LayoutColorAttachmentOptimal,
			},
		},
		SubPasses: []core.SubPass{
			{
				Flags:     core.SubPassShaderResolveQCOM,
				BindPoint: common.BindCompute,
				InputAttachments: []common.AttachmentReference{
					{
						AttachmentIndex: 0,
						Layout:          common.LayoutGeneral,
					},
				},
				ColorAttachments: []common.AttachmentReference{
					{
						AttachmentIndex: 1,
						Layout:          common.LayoutDepthAttachmentOptimal,
					},
					{
						AttachmentIndex: 2,
						Layout:          common.LayoutDepthStencilAttachmentOptimal,
					},
				},
				ResolveAttachments: []common.AttachmentReference{
					{
						AttachmentIndex: 3,
						Layout:          common.LayoutDepthStencilReadOnlyOptimal,
					},
					{
						AttachmentIndex: 5,
						Layout:          common.LayoutDepthAttachmentStencilReadOnlyOptimal,
					},
					{
						AttachmentIndex: 0,
						Layout:          common.LayoutStencilReadOnlyOptimal,
					},
				},
				DepthStencilAttachment: &common.AttachmentReference{
					AttachmentIndex: 11,
					Layout:          common.LayoutTransferSrcOptimal,
				},
				PreservedAttachmentIndices: []int{17},
			},
		},
		SubPassDependencies: []core.SubPassDependency{
			{
				Flags:           common.DependencyDeviceGroup,
				SrcSubPassIndex: 17,
				DstSubPassIndex: 19,
				SrcStageMask:    common.PipelineStageVertexShader,
				DstStageMask:    common.PipelineStageAccelerationStructureBuildKHR,
				SrcAccessMask:   common.AccessColorAttachmentReadNonCoherentEXT,
				DstAccessMask:   common.AccessConditionalRenderingReadEXT,
			},
			{
				Flags:           common.DependencyViewLocal,
				SrcSubPassIndex: 23,
				DstSubPassIndex: 29,
				SrcStageMask:    common.PipelineStageConditionalRenderingEXT,
				DstStageMask:    common.PipelineStageBottomOfPipe,
				SrcAccessMask:   common.AccessColorAttachmentRead,
				DstAccessMask:   common.AccessFragmentDensityMapReadEXT,
			},
			{
				Flags:           0,
				SrcSubPassIndex: 31,
				DstSubPassIndex: 37,
				SrcStageMask:    common.PipelineStageBottomOfPipe,
				DstStageMask:    common.PipelineStageAllGraphics,
				SrcAccessMask:   common.AccessDepthStencilAttachmentWrite,
				DstAccessMask:   common.AccessIndexRead,
			},
		},
	})
	require.EqualError(t, err, "in subpass 0, 2 color attachments are defined, but 3 resolve attachments are defined")
}
