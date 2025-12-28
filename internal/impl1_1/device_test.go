package impl1_1_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
	"go.uber.org/mock/gomock"
)

func TestVulkanDevice_DescriptorSetLayoutSupport(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkGetDescriptorSetLayoutSupport(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkDescriptorSetLayoutCreateInfo, pSupport *loader.VkDescriptorSetLayoutSupport) {
			optionVal := reflect.ValueOf(pCreateInfo).Elem()

			require.Equal(t, uint64(32), optionVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO
			require.True(t, optionVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), optionVal.FieldByName("bindingCount").Uint())

			bindingPtr := (*loader.VkDescriptorSetLayoutBinding)(optionVal.FieldByName("pBindings").UnsafePointer())
			binding := reflect.ValueOf(bindingPtr).Elem()
			require.Equal(t, uint64(1), binding.FieldByName("binding").Uint())
			require.Equal(t, uint64(3), binding.FieldByName("descriptorCount").Uint())

			outDataVal := reflect.ValueOf(pSupport).Elem()

			require.Equal(t, uint64(1000168001), outDataVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT
			require.True(t, outDataVal.FieldByName("pNext").IsNil())

			*(*loader.VkBool32)(unsafe.Pointer(outDataVal.FieldByName("supported").UnsafeAddr())) = loader.VkBool32(1)
		})

	outData := &core1_1.DescriptorSetLayoutSupport{}
	err := driver.GetDescriptorSetLayoutSupport(device, core1_0.DescriptorSetLayoutCreateInfo{
		Bindings: []core1_0.DescriptorSetLayoutBinding{
			{
				Binding:         1,
				DescriptorCount: 3,
			},
		},
	}, outData)
	require.NoError(t, err)
	require.True(t, outData.Supported)
}

func TestVulkanDevice_BindBufferMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(device, coreLoader)

	buffer1 := mocks.NewDummyBuffer(device)
	buffer2 := mocks.NewDummyBuffer(device)

	memory1 := mocks.NewDummyDeviceMemory(device, 1)
	memory2 := mocks.NewDummyDeviceMemory(device, 1)

	coreLoader.EXPECT().VkBindBufferMemory2(device.Handle(), loader.Uint32(2), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkDevice, bindInfoCount loader.Uint32, pBindInfos *loader.VkBindBufferMemoryInfo) (common.VkResult, error) {
			bindInfoSlice := unsafe.Slice(pBindInfos, 2)
			val := reflect.ValueOf(bindInfoSlice)

			bind := val.Index(0)
			require.Equal(t, uint64(1000157000), bind.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO
			require.True(t, bind.FieldByName("pNext").IsNil())
			require.Equal(t, buffer1.Handle(), (loader.VkBuffer)(bind.FieldByName("buffer").UnsafePointer()))
			require.Equal(t, memory1.Handle(), (loader.VkDeviceMemory)(bind.FieldByName("memory").UnsafePointer()))
			require.Equal(t, uint64(1), bind.FieldByName("memoryOffset").Uint())

			bind = val.Index(1)
			require.Equal(t, uint64(1000157000), bind.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO
			require.True(t, bind.FieldByName("pNext").IsNil())
			require.Equal(t, buffer2.Handle(), (loader.VkBuffer)(bind.FieldByName("buffer").UnsafePointer()))
			require.Equal(t, memory2.Handle(), (loader.VkDeviceMemory)(bind.FieldByName("memory").UnsafePointer()))
			require.Equal(t, uint64(3), bind.FieldByName("memoryOffset").Uint())

			return core1_0.VKSuccess, nil
		})

	_, err := driver.BindBufferMemory2(
		core1_1.BindBufferMemoryInfo{
			Buffer:       buffer1,
			Memory:       memory1,
			MemoryOffset: 1,
		},
		core1_1.BindBufferMemoryInfo{
			Buffer:       buffer2,
			Memory:       memory2,
			MemoryOffset: 3,
		},
	)
	require.NoError(t, err)
}

func TestVulkanDevice_BindImageMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(device, coreLoader)

	image1 := mocks.NewDummyImage(device)
	image2 := mocks.NewDummyImage(device)

	memory1 := mocks.NewDummyDeviceMemory(device, 1)
	memory2 := mocks.NewDummyDeviceMemory(device, 1)

	coreLoader.EXPECT().VkBindImageMemory2(device.Handle(), loader.Uint32(2), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkDevice, bindInfoCount loader.Uint32, pBindInfos *loader.VkBindImageMemoryInfo) (common.VkResult, error) {
			bindInfoSlice := unsafe.Slice(pBindInfos, 2)
			val := reflect.ValueOf(bindInfoSlice)

			bind := val.Index(0)
			require.Equal(t, uint64(1000157001), bind.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO
			require.True(t, bind.FieldByName("pNext").IsNil())
			require.Equal(t, image1.Handle(), (loader.VkImage)(bind.FieldByName("image").UnsafePointer()))
			require.Equal(t, memory1.Handle(), (loader.VkDeviceMemory)(bind.FieldByName("memory").UnsafePointer()))
			require.Equal(t, uint64(1), bind.FieldByName("memoryOffset").Uint())

			bind = val.Index(1)
			require.Equal(t, uint64(1000157001), bind.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO
			require.True(t, bind.FieldByName("pNext").IsNil())
			require.Equal(t, image2.Handle(), (loader.VkImage)(bind.FieldByName("image").UnsafePointer()))
			require.Equal(t, memory2.Handle(), (loader.VkDeviceMemory)(bind.FieldByName("memory").UnsafePointer()))
			require.Equal(t, uint64(3), bind.FieldByName("memoryOffset").Uint())

			return core1_0.VKSuccess, nil
		})

	_, err := driver.BindImageMemory2(
		core1_1.BindImageMemoryInfo{
			Image:        image1,
			Memory:       memory1,
			MemoryOffset: 1,
		},
		core1_1.BindImageMemoryInfo{
			Image:        image2,
			Memory:       memory2,
			MemoryOffset: 3,
		},
	)
	require.NoError(t, err)
}

func TestVulkanDevice_BufferMemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(device, coreLoader)

	buffer := mocks.NewDummyBuffer(device)

	coreLoader.EXPECT().VkGetBufferMemoryRequirements2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pInfo *loader.VkBufferMemoryRequirementsInfo2,
		pMemoryRequirements *loader.VkMemoryRequirements2,
	) {
		val := reflect.ValueOf(pInfo).Elem()
		require.Equal(t, uint64(1000146000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, buffer.Handle(), loader.VkBuffer(val.FieldByName("buffer").UnsafePointer()))

		val = reflect.ValueOf(pMemoryRequirements).Elem()
		require.Equal(t, uint64(1000146003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2
		require.True(t, val.FieldByName("pNext").IsNil())

		val = val.FieldByName("memoryRequirements")
		*(*loader.VkDeviceSize)(unsafe.Pointer(val.FieldByName("size").UnsafeAddr())) = loader.VkDeviceSize(1)
		*(*loader.VkDeviceSize)(unsafe.Pointer(val.FieldByName("alignment").UnsafeAddr())) = loader.VkDeviceSize(3)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("memoryTypeBits").UnsafeAddr())) = loader.Uint32(5)
	})

	var outData core1_1.MemoryRequirements2
	err := driver.GetBufferMemoryRequirements2(core1_1.BufferMemoryRequirementsInfo2{
		Buffer: buffer,
	}, &outData)
	require.NoError(t, err)

	require.Equal(t, 1, outData.MemoryRequirements.Size)
	require.Equal(t, 3, outData.MemoryRequirements.Alignment)
	require.Equal(t, uint32(5), outData.MemoryRequirements.MemoryTypeBits)
}

func TestVulkanDevice_ImageMemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(device, coreLoader)

	image := mocks.NewDummyImage(device)

	coreLoader.EXPECT().VkGetImageMemoryRequirements2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pInfo *loader.VkImageMemoryRequirementsInfo2,
		pMemoryRequirements *loader.VkMemoryRequirements2,
	) {
		val := reflect.ValueOf(pInfo).Elem()
		require.Equal(t, uint64(1000146001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, image.Handle(), loader.VkImage(val.FieldByName("image").UnsafePointer()))

		val = reflect.ValueOf(pMemoryRequirements).Elem()
		require.Equal(t, uint64(1000146003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2
		require.True(t, val.FieldByName("pNext").IsNil())

		val = val.FieldByName("memoryRequirements")
		*(*loader.VkDeviceSize)(unsafe.Pointer(val.FieldByName("size").UnsafeAddr())) = loader.VkDeviceSize(1)
		*(*loader.VkDeviceSize)(unsafe.Pointer(val.FieldByName("alignment").UnsafeAddr())) = loader.VkDeviceSize(3)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("memoryTypeBits").UnsafeAddr())) = loader.Uint32(5)
	})

	var outData core1_1.MemoryRequirements2
	err := driver.GetImageMemoryRequirements2(core1_1.ImageMemoryRequirementsInfo2{
		Image: image,
	}, &outData)
	require.NoError(t, err)

	require.Equal(t, 1, outData.MemoryRequirements.Size)
	require.Equal(t, 3, outData.MemoryRequirements.Alignment)
	require.Equal(t, uint32(5), outData.MemoryRequirements.MemoryTypeBits)
}

func TestVulkanExtension_SparseImageMemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(device, coreLoader)

	image := mocks.NewDummyImage(device)

	coreLoader.EXPECT().VkGetImageSparseMemoryRequirements2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device loader.VkDevice,
			pInfo *loader.VkImageSparseMemoryRequirementsInfo2,
			pSparseMemoryRequirementCount *loader.Uint32,
			pSparseMemoryRequirements *loader.VkSparseImageMemoryRequirements2) {

			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000146002), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2
			require.True(t, options.FieldByName("pNext").IsNil())
			require.Equal(t, image.Handle(), (loader.VkImage)(options.FieldByName("image").UnsafePointer()))

			*pSparseMemoryRequirementCount = loader.Uint32(2)
		})

	coreLoader.EXPECT().VkGetImageSparseMemoryRequirements2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkDevice,
			pInfo *loader.VkImageSparseMemoryRequirementsInfo2,
			pSparseMemoryRequirementCount *loader.Uint32,
			pSparseMemoryRequirements *loader.VkSparseImageMemoryRequirements2) {

			require.Equal(t, loader.Uint32(2), *pSparseMemoryRequirementCount)

			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000146002), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2
			require.True(t, options.FieldByName("pNext").IsNil())
			require.Equal(t, image.Handle(), (loader.VkImage)(options.FieldByName("image").UnsafePointer()))

			requirementSlice := ([]loader.VkSparseImageMemoryRequirements2)(unsafe.Slice(pSparseMemoryRequirements, 2))
			outData := reflect.ValueOf(requirementSlice)
			element := outData.Index(0)
			require.Equal(t, uint64(1000146004), element.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2
			require.True(t, element.FieldByName("pNext").IsNil())

			memReqs := element.FieldByName("memoryRequirements")
			imageAspectFlags := (*loader.VkImageAspectFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("aspectMask").UnsafeAddr()))
			*imageAspectFlags = loader.VkImageAspectFlags(0x00000008) // VK_IMAGE_ASPECT_METADATA_BIT
			width := (*loader.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("width").UnsafeAddr()))
			*width = loader.Uint32(1)
			height := (*loader.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("height").UnsafeAddr()))
			*height = loader.Uint32(3)
			depth := (*loader.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("depth").UnsafeAddr()))
			*depth = loader.Uint32(5)
			flags := (*loader.VkSparseImageFormatFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("flags").UnsafeAddr()))
			*flags = loader.VkSparseImageFormatFlags(0x00000004) // VK_SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT
			*(*loader.Uint32)(unsafe.Pointer(memReqs.FieldByName("imageMipTailFirstLod").UnsafeAddr())) = loader.Uint32(7)
			*(*loader.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailSize").UnsafeAddr())) = loader.VkDeviceSize(17)
			*(*loader.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailOffset").UnsafeAddr())) = loader.VkDeviceSize(11)
			*(*loader.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailStride").UnsafeAddr())) = loader.VkDeviceSize(13)

			element = outData.Index(1)
			require.Equal(t, uint64(1000146004), element.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2
			require.True(t, element.FieldByName("pNext").IsNil())

			memReqs = element.FieldByName("memoryRequirements")
			imageAspectFlags = (*loader.VkImageAspectFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("aspectMask").UnsafeAddr()))
			*imageAspectFlags = loader.VkImageAspectFlags(0x00000004) // VK_IMAGE_ASPECT_STENCIL_BIT
			width = (*loader.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("width").UnsafeAddr()))
			*width = loader.Uint32(19)
			height = (*loader.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("height").UnsafeAddr()))
			*height = loader.Uint32(23)
			depth = (*loader.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("depth").UnsafeAddr()))
			*depth = loader.Uint32(29)
			flags = (*loader.VkSparseImageFormatFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("flags").UnsafeAddr()))
			*flags = loader.VkSparseImageFormatFlags(0)
			*(*loader.Uint32)(unsafe.Pointer(memReqs.FieldByName("imageMipTailFirstLod").UnsafeAddr())) = loader.Uint32(43)
			*(*loader.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailSize").UnsafeAddr())) = loader.VkDeviceSize(31)
			*(*loader.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailOffset").UnsafeAddr())) = loader.VkDeviceSize(41)
			*(*loader.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailStride").UnsafeAddr())) = loader.VkDeviceSize(37)
		})

	outData, err := driver.GetImageSparseMemoryRequirements2(core1_1.ImageSparseMemoryRequirementsInfo2{
		Image: image,
	}, nil)
	require.NoError(t, err)
	require.Equal(t, []*core1_1.SparseImageMemoryRequirements2{
		{
			MemoryRequirements: core1_0.SparseImageMemoryRequirements{
				FormatProperties: core1_0.SparseImageFormatProperties{
					AspectMask: core1_0.ImageAspectMetadata,
					ImageGranularity: core1_0.Extent3D{
						Width:  1,
						Height: 3,
						Depth:  5,
					},
					Flags: core1_0.SparseImageFormatNonstandardBlockSize,
				},
				ImageMipTailFirstLod: 7,
				ImageMipTailOffset:   11,
				ImageMipTailStride:   13,
				ImageMipTailSize:     17,
			},
		},
		{
			MemoryRequirements: core1_0.SparseImageMemoryRequirements{
				FormatProperties: core1_0.SparseImageFormatProperties{
					AspectMask: core1_0.ImageAspectStencil,
					ImageGranularity: core1_0.Extent3D{
						Width:  19,
						Height: 23,
						Depth:  29,
					},
					Flags: 0,
				},
				ImageMipTailSize:     31,
				ImageMipTailStride:   37,
				ImageMipTailOffset:   41,
				ImageMipTailFirstLod: 43,
			},
		},
	}, outData)
}

func TestVulkanDevice_GetDeviceGroupPeerMemoryFeatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkGetDeviceGroupPeerMemoryFeatures(
		device.Handle(),
		loader.Uint32(1),
		loader.Uint32(3),
		loader.Uint32(5),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		heapIndex, localDeviceIndex, remoteDeviceIndex loader.Uint32,
		pPeerMemoryFeatures *loader.VkPeerMemoryFeatureFlags,
	) {
		*pPeerMemoryFeatures = loader.VkPeerMemoryFeatureFlags(1) // VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT
	})

	features := driver.GetDeviceGroupPeerMemoryFeatures(
		device,
		1, 3, 5,
	)
	require.Equal(t, core1_1.PeerMemoryFeatureCopySrc, features)
}

func TestVulkanLoader_CreateSamplerYcbcrConversion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(device, coreLoader)

	mockYcbcr := mocks.NewDummySamplerYcbcrConversion(device)

	coreLoader.EXPECT().VkCreateSamplerYcbcrConversion(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkDevice,
			pCreateInfo *loader.VkSamplerYcbcrConversionCreateInfo,
			pAllocator *loader.VkAllocationCallbacks,
			pYcbcrConversion *loader.VkSamplerYcbcrConversion,
		) (common.VkResult, error) {
			*pYcbcrConversion = mockYcbcr.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(1000156000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1000156021), val.FieldByName("format").Uint())             // VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16_KHR
			require.Equal(t, uint64(2), val.FieldByName("ycbcrModel").Uint())                  // VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709_KHR
			require.Equal(t, uint64(1), val.FieldByName("ycbcrRange").Uint())                  // VK_SAMPLER_YCBCR_RANGE_ITU_NARROW_KHR
			require.Equal(t, uint64(4), val.FieldByName("components").FieldByName("r").Uint()) // VK_COMPONENT_SWIZZLE_G
			require.Equal(t, uint64(6), val.FieldByName("components").FieldByName("g").Uint()) // VK_COMPONENT_SWIZZLE_A
			require.Equal(t, uint64(0), val.FieldByName("components").FieldByName("b").Uint()) // VK_COMPONENT_SWIZZLE_IDENTITY
			require.Equal(t, uint64(2), val.FieldByName("components").FieldByName("a").Uint()) // VK_COMPONENT_SWIZZLE_ONE
			require.Equal(t, uint64(0), val.FieldByName("yChromaOffset").Uint())               // VK_CHROMA_LOCATION_COSITED_EVEN_KHR
			require.Equal(t, uint64(1), val.FieldByName("xChromaOffset").Uint())               // VK_CHROMA_LOCATION_MIDPOINT_KHR
			require.Equal(t, uint64(1), val.FieldByName("forceExplicitReconstruction").Uint())

			return core1_0.VKSuccess, nil
		})

	ycbcr, _, err := driver.CreateSamplerYcbcrConversion(
		device,
		core1_1.SamplerYcbcrConversionCreateInfo{
			Format:     core1_1.FormatB12X4G12X4R12X4G12X4HorizontalChromaComponentPacked,
			YcbcrModel: core1_1.SamplerYcbcrModelConversionYcbcr709,
			YcbcrRange: core1_1.SamplerYcbcrRangeITUNarrow,
			Components: core1_0.ComponentMapping{
				R: core1_0.ComponentSwizzleGreen,
				G: core1_0.ComponentSwizzleAlpha,
				B: core1_0.ComponentSwizzleIdentity,
				A: core1_0.ComponentSwizzleOne,
			},
			YChromaOffset:               core1_1.ChromaLocationCositedEven,
			XChromaOffset:               core1_1.ChromaLocationMidpoint,
			ChromaFilter:                core1_0.FilterLinear,
			ForceExplicitReconstruction: true,
		},
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, mockYcbcr.Handle(), ycbcr.Handle())

	coreLoader.EXPECT().VkDestroySamplerYcbcrConversion(
		device.Handle(),
		ycbcr.Handle(),
		gomock.Nil(),
	)

	driver.DestroySamplerYcbcrConversion(ycbcr, nil)
}

func TestVulkanLoader1_1_CreateDescriptorUpdateTemplate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(device, coreLoader)

	mockQueue := mocks.NewDummyQueue(device)

	coreLoader.EXPECT().VkGetDeviceQueue2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pInfo *loader.VkDeviceQueueInfo2,
		pQueue *loader.VkQueue,
	) {
		*pQueue = mockQueue.Handle()

		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000145003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_QUEUE_INFO_2
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT
		require.Equal(t, uint64(3), val.FieldByName("queueFamilyIndex").Uint())
		require.Equal(t, uint64(5), val.FieldByName("queueIndex").Uint())
	})

	queue, err := driver.GetDeviceQueue2(
		device,
		core1_1.DeviceQueueInfo2{
			Flags:            core1_1.DeviceQueueCreateProtected,
			QueueFamilyIndex: 3,
			QueueIndex:       5,
		})

	require.NoError(t, err)
	require.Equal(t, mockQueue.Handle(), queue.Handle())
}
