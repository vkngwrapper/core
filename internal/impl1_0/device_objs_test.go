package impl1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreateBufferView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	buffer := mocks.NewDummyBuffer(device)
	expectedBufferView := mocks.NewFakeBufferViewHandle()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateBufferView(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkBufferViewCreateInfo, pAllocator *loader.VkAllocationCallbacks, pBufferView *loader.VkBufferView) (common.VkResult, error) {
			v := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, v.FieldByName("sType").Uint(), uint64(13)) // VK_STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, v.FieldByName("flags").Uint(), uint64(0))

			actualBuffer := (loader.VkBuffer)(unsafe.Pointer(v.FieldByName("buffer").Elem().UnsafeAddr()))
			require.Equal(t, actualBuffer, buffer.Handle())

			require.Equal(t, v.FieldByName("format").Uint(), uint64(103)) // VK_FORMAT_R32G32_SFLOAT
			require.Equal(t, v.FieldByName("offset").Uint(), uint64(5))
			require.Equal(t, v.FieldByName("_range").Uint(), uint64(7))

			*pBufferView = expectedBufferView
			return core1_0.VKSuccess, nil
		})

	bufferView, res, err := driver.CreateBufferView(nil, core1_0.BufferViewCreateInfo{
		Buffer: buffer,
		Format: core1_0.FormatR32G32SignedFloat,
		Offset: 5,
		Range:  7,
	})

	require.Equal(t, res, core1_0.VKSuccess)
	require.NoError(t, err)
	require.Equal(t, expectedBufferView, bufferView.Handle())
}

func TestVulkanDescriptorSet_Free(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	pool := mocks.NewDummyDescriptorPool(device)
	set := mocks.NewDummyDescriptorSet(pool, device)

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkFreeDescriptorSets(device.Handle(), pool.Handle(), loader.Uint32(1), gomock.Not(unsafe.Pointer(nil))).DoAndReturn(
		func(device loader.VkDevice, descriptorPool loader.VkDescriptorPool, descriptorSetCount loader.Uint32, pDescriptorSets *loader.VkDescriptorSet) (common.VkResult, error) {
			descriptorSetSlice := unsafe.Slice(pDescriptorSets, 1)
			require.Equal(t, set.Handle(), descriptorSetSlice[0])

			return core1_0.VKSuccess, nil
		})

	_, err := driver.FreeDescriptorSets(set)
	require.NoError(t, err)
}

func TestVulkanLoader1_0_CreateFrameBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	renderPass := mocks.NewDummyRenderPass(device)
	imageView1 := mocks.NewDummyImageView(device)
	imageView2 := mocks.NewDummyImageView(device)

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	framebufferHandle := mocks.NewFakeFramebufferHandle()

	mockLoader.EXPECT().VkCreateFramebuffer(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkFramebufferCreateInfo, pAllocator *loader.VkAllocationCallbacks, pFramebuffer *loader.VkFramebuffer) (common.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(37), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(3), val.FieldByName("width").Uint())
			require.Equal(t, uint64(5), val.FieldByName("height").Uint())
			require.Equal(t, uint64(7), val.FieldByName("layers").Uint())
			require.Equal(t, uint64(2), val.FieldByName("attachmentCount").Uint())

			require.Equal(t, renderPass.Handle(), (loader.VkRenderPass)(unsafe.Pointer(val.FieldByName("renderPass").Elem().UnsafeAddr())))

			attachmentPtr := (*loader.VkImageView)(unsafe.Pointer(val.FieldByName("pAttachments").Elem().UnsafeAddr()))
			attachmentSlice := ([]loader.VkImageView)(unsafe.Slice(attachmentPtr, 2))
			require.Equal(t, imageView1.Handle(), attachmentSlice[0])
			require.Equal(t, imageView2.Handle(), attachmentSlice[1])

			*pFramebuffer = framebufferHandle
			return core1_0.VKSuccess, nil
		})

	framebuffer, _, err := driver.CreateFramebuffer(nil, core1_0.FramebufferCreateInfo{
		Flags:      0,
		RenderPass: renderPass,
		Width:      3,
		Height:     5,
		Layers:     7,
		Attachments: []core1_0.ImageView{
			imageView1, imageView2,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, framebuffer)
	require.Equal(t, framebufferHandle, framebuffer.Handle())
}

func TestVulkanLoader1_0_CreateImageView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	imageViewHandle := mocks.NewFakeImageViewHandle()
	image := mocks.NewDummyImage(device)

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateImageView(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkImageViewCreateInfo, pAllocator *loader.VkAllocationCallbacks, pImageView *loader.VkImageView) (common.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(15), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, image.Handle(), (loader.VkImage)(unsafe.Pointer(val.FieldByName("image").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), val.FieldByName("viewType").Uint()) // VK_IMAGE_VIEW_TYPE_2D
			require.Equal(t, uint64(67), val.FieldByName("format").Uint())  // VK_FORMAT_A2B10G10R10_SSCALED_PACK32

			components := val.FieldByName("components")
			require.Equal(t, uint64(3), components.FieldByName("r").Uint()) // VK_COMPONENT_SWIZZLE_R
			require.Equal(t, uint64(4), components.FieldByName("g").Uint()) // VK_COMPONENT_SWIZZLE_G
			require.Equal(t, uint64(5), components.FieldByName("b").Uint()) // VK_COMPONENT_SWIZZLE_B
			require.Equal(t, uint64(6), components.FieldByName("a").Uint()) // VK_COMPONENT_SWIZZLE_A

			subresource := val.FieldByName("subresourceRange")
			require.Equal(t, uint64(1), subresource.FieldByName("baseMipLevel").Uint())
			require.Equal(t, uint64(2), subresource.FieldByName("levelCount").Uint())
			require.Equal(t, uint64(3), subresource.FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(5), subresource.FieldByName("layerCount").Uint())
			require.Equal(t, uint64(3), subresource.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_COLOR_BIT | VK_IMAGE_ASPECT_DEPTH_BIT

			*pImageView = imageViewHandle
			return core1_0.VKSuccess, nil
		})

	imageView, _, err := driver.CreateImageView(nil, core1_0.ImageViewCreateInfo{
		Image:    image,
		ViewType: core1_0.ImageViewType2D,
		Format:   core1_0.FormatA2B10G10R10SignedScaledPacked,
		Flags:    0,
		Components: core1_0.ComponentMapping{
			A: core1_0.ComponentSwizzleAlpha,
			R: core1_0.ComponentSwizzleRed,
			G: core1_0.ComponentSwizzleGreen,
			B: core1_0.ComponentSwizzleBlue,
		},
		SubresourceRange: core1_0.ImageSubresourceRange{
			BaseMipLevel:   1,
			LevelCount:     2,
			BaseArrayLayer: 3,
			LayerCount:     5,
			AspectMask:     core1_0.ImageAspectColor | core1_0.ImageAspectDepth,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, imageView)
	require.Equal(t, imageViewHandle, imageView.Handle())
}

func TestVulkanLoader1_0_CreateSampler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	samplerHandle := mocks.NewFakeSamplerHandle()

	mockLoader.EXPECT().VkCreateSampler(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkSamplerCreateInfo, pAllocator *loader.VkAllocationCallbacks, pSampler *loader.VkSampler) (common.VkResult, error) {
			*pSampler = samplerHandle

			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(31), val.FieldByName("sType").Uint())
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(0), val.FieldByName("magFilter").Uint())    // VK_FILTER_NEAREST
			require.Equal(t, uint64(1), val.FieldByName("minFilter").Uint())    // VK_FILTER_LINEAR
			require.Equal(t, uint64(1), val.FieldByName("mipmapMode").Uint())   // VK_SAMPLER_MIPMAP_MODE_LINEAR
			require.Equal(t, uint64(0), val.FieldByName("addressModeU").Uint()) // VK_SAMPLER_ADDRESS_MODE_REPEAT
			require.Equal(t, uint64(3), val.FieldByName("addressModeV").Uint()) // VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER
			require.Equal(t, uint64(1), val.FieldByName("addressModeW").Uint()) // VK_SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT
			require.InDelta(t, 1.2, val.FieldByName("mipLodBias").Float(), 0.0001)
			require.Equal(t, uint64(1), val.FieldByName("anisotropyEnable").Uint()) // VK_TRUE
			require.InDelta(t, 4.5, val.FieldByName("maxAnisotropy").Float(), 0.0001)
			require.Equal(t, uint64(1), val.FieldByName("compareEnable").Uint()) // VK_TRUE
			require.Equal(t, uint64(4), val.FieldByName("compareOp").Uint())     // VK_COMPARE_OP_GREATER
			require.InDelta(t, 2.3, val.FieldByName("minLod").Float(), 0.0001)
			require.InDelta(t, 3.4, val.FieldByName("maxLod").Float(), 0.0001)
			require.Equal(t, uint64(2), val.FieldByName("borderColor").Uint())             // VK_BORDER_COLOR_FLOAT_OPAQUE_BLACK
			require.Equal(t, uint64(1), val.FieldByName("unnormalizedCoordinates").Uint()) // VK_TRUE

			return core1_0.VKSuccess, nil
		})

	sampler, _, err := driver.CreateSampler(nil, core1_0.SamplerCreateInfo{
		Flags:                   0,
		MagFilter:               core1_0.FilterNearest,
		MinFilter:               core1_0.FilterLinear,
		MipmapMode:              core1_0.SamplerMipmapModeLinear,
		AddressModeU:            core1_0.SamplerAddressModeRepeat,
		AddressModeV:            core1_0.SamplerAddressModeClampToBorder,
		AddressModeW:            core1_0.SamplerAddressModeMirroredRepeat,
		MipLodBias:              1.2,
		MinLod:                  2.3,
		MaxLod:                  3.4,
		AnisotropyEnable:        true,
		MaxAnisotropy:           4.5,
		CompareEnable:           true,
		CompareOp:               core1_0.CompareOpGreater,
		BorderColor:             core1_0.BorderColorFloatOpaqueBlack,
		UnnormalizedCoordinates: true,
	})
	require.NoError(t, err)
	require.NotNil(t, sampler)
	require.Equal(t, samplerHandle, sampler.Handle())
}

func TestVulkanLoader1_0_CreateSemaphore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	semaphoreHandle := mocks.NewFakeSemaphore()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateSemaphore(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkSemaphoreCreateInfo, pAllocator *loader.VkAllocationCallbacks, pSemaphore *loader.VkSemaphore) (common.VkResult, error) {
			*pSemaphore = semaphoreHandle
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(9), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())

			return core1_0.VKSuccess, nil
		})

	semaphore, _, err := driver.CreateSemaphore(nil, core1_0.SemaphoreCreateInfo{})
	require.NoError(t, err)
	require.NotNil(t, semaphore)
	require.Equal(t, semaphoreHandle, semaphore.Handle())
}

func TestVulkanLoader1_0_CreateShaderModule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	handle := mocks.NewFakeShaderModule()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateShaderModule(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkShaderModuleCreateInfo, pAllocator *loader.VkAllocationCallbacks, pShaderModule *loader.VkShaderModule) (common.VkResult, error) {
			*pShaderModule = handle
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(16), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(32), val.FieldByName("codeSize").Uint())

			codePtr := (*loader.Uint32)(unsafe.Pointer(val.FieldByName("pCode").Elem().UnsafeAddr()))
			codeSlice := ([]loader.Uint32)(unsafe.Slice(codePtr, 8))

			require.Equal(t, []loader.Uint32{1, 1, 2, 3, 5, 8, 13, 21}, codeSlice)

			return core1_0.VKSuccess, nil
		})

	shaderModule, _, err := driver.CreateShaderModule(nil, core1_0.ShaderModuleCreateInfo{
		Code: []uint32{1, 1, 2, 3, 5, 8, 13, 21},
	})
	require.NoError(t, err)
	require.NotNil(t, shaderModule)
	require.Equal(t, handle, shaderModule.Handle())
}
