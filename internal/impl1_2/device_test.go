package impl1_2_test

import (
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/internal/impl1_2"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestDevice_CreateRenderPass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := impl1_2.NewDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	mockRenderPass := mocks.NewDummyRenderPass(device)

	coreLoader.EXPECT().VkCreateRenderPass2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pCreateInfo *loader.VkRenderPassCreateInfo2,
		pAllocator *loader.VkAllocationCallbacks,
		pRenderPass *loader.VkRenderPass) (common.VkResult, error) {

		*pRenderPass = mockRenderPass.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()

		require.Equal(t, uint64(1000109004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2
		require.True(t, val.FieldByName("pNext").IsNil())

		require.Equal(t, uint64(1), val.FieldByName("attachmentCount").Uint())
		require.Equal(t, uint64(1), val.FieldByName("subpassCount").Uint())
		require.Equal(t, uint64(2), val.FieldByName("dependencyCount").Uint())
		require.Equal(t, uint64(3), val.FieldByName("correlatedViewMaskCount").Uint())

		attachmentsPtr := (*loader.VkAttachmentDescription2)(val.FieldByName("pAttachments").UnsafePointer())
		attachmentsSlice := unsafe.Slice(attachmentsPtr, 1)
		attachment := reflect.ValueOf(attachmentsSlice).Index(0)

		require.Equal(t, uint64(1000109000), attachment.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2
		require.True(t, attachment.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), attachment.FieldByName("flags").Uint())          // VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT
		require.Equal(t, uint64(68), attachment.FieldByName("format").Uint())        // VK_FORMAT_A2B10G10R10_UINT_PACK32
		require.Equal(t, uint64(8), attachment.FieldByName("samples").Uint())        // VK_SAMPLE_COUNT_8_BIT
		require.Equal(t, uint64(1), attachment.FieldByName("loadOp").Uint())         // VK_ATTACHMENT_LOAD_OP_CLEAR
		require.Equal(t, uint64(1), attachment.FieldByName("storeOp").Uint())        // VK_ATTACHMENT_STORE_OP_DONT_CARE
		require.Equal(t, uint64(2), attachment.FieldByName("stencilLoadOp").Uint())  // VK_ATTACHMENT_LOAD_OP_DONT_CARE
		require.Equal(t, uint64(0), attachment.FieldByName("stencilStoreOp").Uint()) // VK_ATTACHMENT_STORE_OP_STORE
		require.Equal(t, uint64(4), attachment.FieldByName("initialLayout").Uint())  // VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL
		require.Equal(t, uint64(8), attachment.FieldByName("finalLayout").Uint())    // VK_IMAGE_LAYOUT_PREINITIALIZED

		viewMasks := (*uint32)(val.FieldByName("pCorrelatedViewMasks").UnsafePointer())
		viewMaskSlice := ([]uint32)(unsafe.Slice(viewMasks, 3))
		require.Equal(t, []uint32{29, 31, 37}, viewMaskSlice)

		subpassPtr := (*loader.VkSubpassDescription2)(val.FieldByName("pSubpasses").UnsafePointer())
		subpassSlice := unsafe.Slice(subpassPtr, 1)
		subpass := reflect.ValueOf(subpassSlice).Index(0)

		require.Equal(t, uint64(1000109002), subpass.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2
		require.True(t, subpass.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0), subpass.FieldByName("flags").Uint())
		require.Equal(t, uint64(1), subpass.FieldByName("pipelineBindPoint").Uint()) // VK_PIPELINE_BIND_POINT_COMPUTE
		require.Equal(t, uint64(1), subpass.FieldByName("viewMask").Uint())
		require.Equal(t, uint64(2), subpass.FieldByName("inputAttachmentCount").Uint())
		require.Equal(t, uint64(1), subpass.FieldByName("colorAttachmentCount").Uint())
		require.Equal(t, uint64(2), subpass.FieldByName("preserveAttachmentCount").Uint())

		preserveAttachments := (*uint32)(subpass.FieldByName("pPreserveAttachments").UnsafePointer())
		preserveAttachmentSlice := ([]uint32)(unsafe.Slice(preserveAttachments, 2))
		require.Equal(t, []uint32{59, 61}, preserveAttachmentSlice)

		inputAttachmentPtr := (*loader.VkAttachmentReference2)(subpass.FieldByName("pInputAttachments").UnsafePointer())
		inputAttachmentSlice := ([]loader.VkAttachmentReference2)(unsafe.Slice(inputAttachmentPtr, 2))
		inputAttachment := reflect.ValueOf(inputAttachmentSlice)
		require.Equal(t, uint64(1000109001), inputAttachment.Index(0).FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2
		require.True(t, inputAttachment.Index(0).FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), inputAttachment.Index(0).FieldByName("attachment").Uint())
		require.Equal(t, uint64(6), inputAttachment.Index(0).FieldByName("layout").Uint())     // VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL
		require.Equal(t, uint64(4), inputAttachment.Index(0).FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_STENCIL_BIT

		require.Equal(t, uint64(1000109001), inputAttachment.Index(1).FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2
		require.True(t, inputAttachment.Index(1).FieldByName("pNext").IsNil())
		require.Equal(t, uint64(5), inputAttachment.Index(1).FieldByName("attachment").Uint())
		require.Equal(t, uint64(6), inputAttachment.Index(1).FieldByName("layout").Uint())     // VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL
		require.Equal(t, uint64(8), inputAttachment.Index(1).FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_METADATA_BIT

		colorAttachment := subpass.FieldByName("pColorAttachments").Elem()
		require.Equal(t, uint64(1000109001), colorAttachment.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2
		require.True(t, colorAttachment.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(41), colorAttachment.FieldByName("attachment").Uint())
		require.Equal(t, uint64(8), colorAttachment.FieldByName("layout").Uint())     // VK_IMAGE_LAYOUT_PREINITIALIZED
		require.Equal(t, uint64(1), colorAttachment.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_COLOR_BIT

		resolveAttachment := subpass.FieldByName("pResolveAttachments").Elem()
		require.Equal(t, uint64(1000109001), resolveAttachment.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2
		require.True(t, resolveAttachment.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(43), resolveAttachment.FieldByName("attachment").Uint())
		require.Equal(t, uint64(1), resolveAttachment.FieldByName("layout").Uint())     // VK_IMAGE_LAYOUT_GENERAL
		require.Equal(t, uint64(2), resolveAttachment.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_DEPTH_BIT

		depthAttachment := subpass.FieldByName("pDepthStencilAttachment").Elem()
		require.Equal(t, uint64(1000109001), depthAttachment.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2
		require.True(t, depthAttachment.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(47), depthAttachment.FieldByName("attachment").Uint())
		require.Equal(t, uint64(7), depthAttachment.FieldByName("layout").Uint())     // VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL
		require.Equal(t, uint64(1), depthAttachment.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_COLOR_BIT

		dependenciesPtr := (*loader.VkSubpassDependency2)(val.FieldByName("pDependencies").UnsafePointer())
		dependenciesSlice := ([]loader.VkSubpassDependency2)(unsafe.Slice(dependenciesPtr, 2))
		val = reflect.ValueOf(dependenciesSlice)
		require.Equal(t, uint64(1000109003), val.Index(0).FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2
		require.True(t, val.Index(0).FieldByName("pNext").IsNil())
		require.Equal(t, uint64(7), val.Index(0).FieldByName("srcSubpass").Uint())
		require.Equal(t, uint64(11), val.Index(0).FieldByName("dstSubpass").Uint())
		require.Equal(t, uint64(0x800), val.Index(0).FieldByName("srcStageMask").Uint())  // VK_PIPELINE_STAGE_COMPUTE_SHADER_BIT
		require.Equal(t, uint64(2), val.Index(0).FieldByName("dstStageMask").Uint())      // VK_PIPELINE_STAGE_DRAW_INDIRECT_BIT
		require.Equal(t, uint64(2), val.Index(0).FieldByName("srcAccessMask").Uint())     // VK_ACCESS_INDEX_READ_BIT
		require.Equal(t, uint64(0x100), val.Index(0).FieldByName("dstAccessMask").Uint()) // VK_ACCESS_COLOR_ATTACHMENT_WRITE_BIT
		require.Equal(t, uint64(1), val.Index(0).FieldByName("dependencyFlags").Uint())   // VK_DEPENDENCY_BY_REGION_BIT
		require.Equal(t, int64(13), val.Index(0).FieldByName("viewOffset").Int())

		require.Equal(t, uint64(1000109003), val.Index(1).FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2
		require.True(t, val.Index(1).FieldByName("pNext").IsNil())
		require.Equal(t, uint64(17), val.Index(1).FieldByName("srcSubpass").Uint())
		require.Equal(t, uint64(19), val.Index(1).FieldByName("dstSubpass").Uint())
		require.Equal(t, uint64(0x40), val.Index(1).FieldByName("srcStageMask").Uint())   // VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT
		require.Equal(t, uint64(0x4000), val.Index(1).FieldByName("dstStageMask").Uint()) // VK_PIPELINE_STAGE_HOST_BIT
		require.Equal(t, uint64(0x80), val.Index(1).FieldByName("srcAccessMask").Uint())  // VK_ACCESS_COLOR_ATTACHMENT_READ_BIT
		require.Equal(t, uint64(0x200), val.Index(1).FieldByName("dstAccessMask").Uint()) // VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_READ_BIT
		require.Equal(t, uint64(0), val.Index(1).FieldByName("dependencyFlags").Uint())
		require.Equal(t, int64(23), val.Index(1).FieldByName("viewOffset").Int())

		return core1_0.VKSuccess, nil
	})

	renderPass, _, err := driver.CreateRenderPass2(
		device,
		nil,
		core1_2.RenderPassCreateInfo2{
			Flags: 0,
			Attachments: []core1_2.AttachmentDescription2{
				{
					Flags:          core1_0.AttachmentDescriptionMayAlias,
					Format:         core1_0.FormatA2B10G10R10UnsignedIntPacked,
					Samples:        core1_0.Samples8,
					LoadOp:         core1_0.AttachmentLoadOpClear,
					StoreOp:        core1_0.AttachmentStoreOpDontCare,
					StencilLoadOp:  core1_0.AttachmentLoadOpDontCare,
					StencilStoreOp: core1_0.AttachmentStoreOpStore,
					InitialLayout:  core1_0.ImageLayoutDepthStencilReadOnlyOptimal,
					FinalLayout:    core1_0.ImageLayoutPreInitialized,
				},
			},
			Subpasses: []core1_2.SubpassDescription2{
				{
					Flags:             0,
					PipelineBindPoint: core1_0.PipelineBindPointCompute,
					ViewMask:          1,
					InputAttachments: []core1_2.AttachmentReference2{
						{
							Attachment: 3,
							Layout:     core1_0.ImageLayoutTransferSrcOptimal,
							AspectMask: core1_0.ImageAspectStencil,
						},
						{
							Attachment: 5,
							Layout:     core1_0.ImageLayoutTransferSrcOptimal,
							AspectMask: core1_0.ImageAspectMetadata,
						},
					},
					ColorAttachments: []core1_2.AttachmentReference2{
						{
							Attachment: 41,
							Layout:     core1_0.ImageLayoutPreInitialized,
							AspectMask: core1_0.ImageAspectColor,
						},
					},
					ResolveAttachments: []core1_2.AttachmentReference2{
						{
							Attachment: 43,
							Layout:     core1_0.ImageLayoutGeneral,
							AspectMask: core1_0.ImageAspectDepth,
						},
					},
					DepthStencilAttachment: &core1_2.AttachmentReference2{
						Attachment: 47,
						Layout:     core1_0.ImageLayoutTransferDstOptimal,
						AspectMask: core1_0.ImageAspectColor,
					},
					PreserveAttachments: []int{59, 61},
				},
			},
			Dependencies: []core1_2.SubpassDependency2{
				{
					SrcSubpass:      7,
					DstSubpass:      11,
					SrcStageMask:    core1_0.PipelineStageComputeShader,
					DstStageMask:    core1_0.PipelineStageDrawIndirect,
					SrcAccessMask:   core1_0.AccessIndexRead,
					DstAccessMask:   core1_0.AccessColorAttachmentWrite,
					DependencyFlags: core1_0.DependencyByRegion,
					ViewOffset:      13,
				},
				{
					SrcSubpass:      17,
					DstSubpass:      19,
					SrcStageMask:    core1_0.PipelineStageGeometryShader,
					DstStageMask:    core1_0.PipelineStageHost,
					SrcAccessMask:   core1_0.AccessColorAttachmentRead,
					DstAccessMask:   core1_0.AccessDepthStencilAttachmentRead,
					DependencyFlags: 0,
					ViewOffset:      23,
				},
			},
			CorrelatedViewMasks: []uint32{29, 31, 37},
		})
	require.NoError(t, err)
	require.Equal(t, mockRenderPass.Handle(), renderPass.Handle())
}

func TestDevice_GetBufferDeviceAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := impl1_2.NewDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	buffer := mocks.NewDummyBuffer(device)

	coreLoader.EXPECT().VkGetBufferDeviceAddress(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice, pInfo *loader.VkBufferDeviceAddressInfo) loader.VkDeviceAddress {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000244001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, buffer.Handle(), loader.VkBuffer(val.FieldByName("buffer").UnsafePointer()))

		return 5
	})

	address, err := driver.GetBufferDeviceAddress(
		core1_2.BufferDeviceAddressInfo{
			Buffer: buffer,
		})
	require.NoError(t, err)
	require.Equal(t, uint64(5), address)
}

func TestDevice_GetBufferOpaqueCaptureAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := impl1_2.NewDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	buffer := mocks.NewDummyBuffer(device)

	coreLoader.EXPECT().VkGetBufferOpaqueCaptureAddress(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice, pInfo *loader.VkBufferDeviceAddressInfo) loader.Uint64 {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000244001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, buffer.Handle(), loader.VkBuffer(val.FieldByName("buffer").UnsafePointer()))

		return 7
	})

	address, err := driver.GetBufferOpaqueCaptureAddress(
		core1_2.BufferDeviceAddressInfo{
			Buffer: buffer,
		})
	require.NoError(t, err)
	require.Equal(t, uint64(7), address)
}

func TestDevice_GetDeviceMemoryOpaqueCaptureAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := impl1_2.NewDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	deviceMemory := mocks.NewDummyDeviceMemory(device, 1)

	coreLoader.EXPECT().VkGetDeviceMemoryOpaqueCaptureAddress(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice, pInfo *loader.VkDeviceMemoryOpaqueCaptureAddressInfo) loader.Uint64 {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000257004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_MEMORY_OPAQUE_CAPTURE_ADDRESS_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, deviceMemory.Handle(), loader.VkDeviceMemory(val.FieldByName("memory").UnsafePointer()))

		return 11
	})

	address, err := driver.GetDeviceMemoryOpaqueCaptureAddress(
		core1_2.DeviceMemoryOpaqueCaptureAddressInfo{
			Memory: deviceMemory,
		})
	require.NoError(t, err)
	require.Equal(t, uint64(11), address)
}

func TestVulkanDevice_SignalSemaphore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := impl1_2.NewDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	semaphore := mocks.NewDummySemaphore(device)

	coreLoader.EXPECT().VkSignalSemaphore(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pSignalInfo *loader.VkSemaphoreSignalInfo) (common.VkResult, error) {

		val := reflect.ValueOf(pSignalInfo).Elem()
		require.Equal(t, uint64(1000207005), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_SIGNAL_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, semaphore.Handle(), loader.VkSemaphore(val.FieldByName("semaphore").UnsafePointer()))
		require.Equal(t, uint64(13), val.FieldByName("value").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := driver.SignalSemaphore(
		core1_2.SemaphoreSignalInfo{
			Semaphore: semaphore,
			Value:     uint64(13),
		})
	require.NoError(t, err)
}

func TestVulkanDevice_WaitSemaphores(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := impl1_2.NewDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})

	semaphore1 := mocks.NewDummySemaphore(device)
	semaphore2 := mocks.NewDummySemaphore(device)

	coreLoader.EXPECT().VkWaitSemaphores(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		loader.Uint64(60000000000),
	).DoAndReturn(func(device loader.VkDevice,
		pWaitInfo *loader.VkSemaphoreWaitInfo,
		timeout loader.Uint64) (common.VkResult, error) {

		val := reflect.ValueOf(pWaitInfo).Elem()
		require.Equal(t, uint64(1000207004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_WAIT_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_SEMAPHORE_WAIT_ANY_BIT
		require.Equal(t, uint64(2), val.FieldByName("semaphoreCount").Uint())

		semaphorePtr := (*loader.VkSemaphore)(val.FieldByName("pSemaphores").UnsafePointer())
		semaphoreSlice := unsafe.Slice(semaphorePtr, 2)
		require.Equal(t, []loader.VkSemaphore{semaphore1.Handle(), semaphore2.Handle()}, semaphoreSlice)

		valuesPtr := (*loader.Uint64)(val.FieldByName("pValues").UnsafePointer())
		valuesSlice := unsafe.Slice(valuesPtr, 2)
		require.Equal(t, []loader.Uint64{13, 19}, valuesSlice)

		return core1_0.VKSuccess, nil
	})

	_, err := driver.WaitSemaphores(
		time.Minute,
		core1_2.SemaphoreWaitInfo{
			Flags: core1_2.SemaphoreWaitAny,
			Semaphores: []core.Semaphore{
				semaphore1,
				semaphore2,
			},
			Values: []uint64{
				13,
				19,
			},
		})
	require.NoError(t, err)
}
