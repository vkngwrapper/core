package core1_1_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestMemoryDedicatedAllocateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := impl1_1.NewDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	buffer := mocks.NewDummyBuffer(device)
	expectedMemory := mocks.NewDummyDeviceMemory(device, 1)

	coreLoader.EXPECT().VkAllocateMemory(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkDevice, pAllocateInfo *loader.VkMemoryAllocateInfo, pAllocator *loader.VkAllocationCallbacks, pMemory *loader.VkDeviceMemory) (common.VkResult, error) {
			*pMemory = expectedMemory.Handle()

			options := reflect.ValueOf(pAllocateInfo).Elem()
			require.Equal(t, uint64(5), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
			require.Equal(t, uint64(1), options.FieldByName("allocationSize").Uint())
			require.Equal(t, uint64(3), options.FieldByName("memoryTypeIndex").Uint())

			dedicatedPtr := (*loader.VkMemoryDedicatedAllocateInfo)(options.FieldByName("pNext").UnsafePointer())
			dedicated := reflect.ValueOf(dedicatedPtr).Elem()

			require.Equal(t, uint64(1000127001), dedicated.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO
			require.True(t, dedicated.FieldByName("pNext").IsNil())
			require.Equal(t, buffer.Handle(), loader.VkBuffer(dedicated.FieldByName("buffer").UnsafePointer()))
			require.True(t, dedicated.FieldByName("image").IsNil())

			return core1_0.VKSuccess, nil
		})

	memory, _, err := driver.AllocateMemory(device, nil, core1_0.MemoryAllocateInfo{
		AllocationSize:  1,
		MemoryTypeIndex: 3,
		NextOptions: common.NextOptions{Next: core1_1.MemoryDedicatedAllocateInfo{
			Buffer: buffer,
		}},
	})
	require.NoError(t, err)
	require.Equal(t, expectedMemory.Handle(), memory.Handle())
}

func TestDedicatedMemoryRequirementsOutData_Buffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := impl1_1.NewDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})
	buffer := mocks.NewDummyBuffer(device)

	coreLoader.EXPECT().VkGetBufferMemoryRequirements2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device loader.VkDevice,
			pInfo *loader.VkBufferMemoryRequirementsInfo2,
			pMemoryRequirements *loader.VkMemoryRequirements2,
		) {
			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000146000), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2
			require.True(t, options.FieldByName("pNext").IsNil())
			require.Equal(t, buffer.Handle(), loader.VkBuffer(options.FieldByName("buffer").UnsafePointer()))

			outData := reflect.ValueOf(pMemoryRequirements).Elem()
			require.Equal(t, uint64(1000146003), outData.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2

			memoryRequirements := outData.FieldByName("memoryRequirements")
			*(*loader.VkDeviceSize)(unsafe.Pointer(memoryRequirements.FieldByName("size").UnsafeAddr())) = loader.VkDeviceSize(1)
			*(*loader.VkDeviceSize)(unsafe.Pointer(memoryRequirements.FieldByName("alignment").UnsafeAddr())) = loader.VkDeviceSize(3)
			*(*loader.Uint32)(unsafe.Pointer(memoryRequirements.FieldByName("memoryTypeBits").UnsafeAddr())) = loader.Uint32(5)

			dedicatedPtr := (*loader.VkMemoryDedicatedRequirements)(outData.FieldByName("pNext").UnsafePointer())
			dedicated := reflect.ValueOf(dedicatedPtr).Elem()
			require.Equal(t, uint64(1000127000), dedicated.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS
			require.True(t, dedicated.FieldByName("pNext").IsNil())
			*(*loader.VkBool32)(unsafe.Pointer(dedicated.FieldByName("prefersDedicatedAllocation").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(dedicated.FieldByName("requiresDedicatedAllocation").UnsafeAddr())) = loader.VkBool32(0)
		})

	var memReqs core1_1.MemoryDedicatedRequirements
	var outData = core1_1.MemoryRequirements2{
		NextOutData: common.NextOutData{Next: &memReqs},
	}
	err := driver.GetBufferMemoryRequirements2(
		core1_1.BufferMemoryRequirementsInfo2{
			Buffer: buffer,
		}, &outData)
	require.NoError(t, err)
	require.False(t, memReqs.RequiresDedicatedAllocation)
	require.True(t, memReqs.PrefersDedicatedAllocation)

	require.Equal(t, 1, outData.MemoryRequirements.Size)
	require.Equal(t, 3, outData.MemoryRequirements.Alignment)
	require.Equal(t, uint32(5), outData.MemoryRequirements.MemoryTypeBits)
}

func TestDedicatedMemoryRequirementsOutData_Image(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := impl1_1.NewDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})
	image := mocks.NewDummyImage(device)

	coreLoader.EXPECT().VkGetImageMemoryRequirements2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device loader.VkDevice,
			pInfo *loader.VkImageMemoryRequirementsInfo2,
			pMemoryRequirements *loader.VkMemoryRequirements2,
		) {
			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000146001), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2
			require.True(t, options.FieldByName("pNext").IsNil())
			require.Equal(t, image.Handle(), loader.VkImage(options.FieldByName("image").UnsafePointer()))

			outData := reflect.ValueOf(pMemoryRequirements).Elem()
			require.Equal(t, uint64(1000146003), outData.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2

			memoryRequirements := outData.FieldByName("memoryRequirements")
			*(*loader.VkDeviceSize)(unsafe.Pointer(memoryRequirements.FieldByName("size").UnsafeAddr())) = loader.VkDeviceSize(1)
			*(*loader.VkDeviceSize)(unsafe.Pointer(memoryRequirements.FieldByName("alignment").UnsafeAddr())) = loader.VkDeviceSize(3)
			*(*loader.Uint32)(unsafe.Pointer(memoryRequirements.FieldByName("memoryTypeBits").UnsafeAddr())) = loader.Uint32(5)

			dedicatedPtr := (*loader.VkMemoryDedicatedRequirements)(outData.FieldByName("pNext").UnsafePointer())
			dedicated := reflect.ValueOf(dedicatedPtr).Elem()
			require.Equal(t, uint64(1000127000), dedicated.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS
			require.True(t, dedicated.FieldByName("pNext").IsNil())
			*(*loader.VkBool32)(unsafe.Pointer(dedicated.FieldByName("prefersDedicatedAllocation").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(dedicated.FieldByName("requiresDedicatedAllocation").UnsafeAddr())) = loader.VkBool32(0)
		})

	var memReqs core1_1.MemoryDedicatedRequirements
	var outData = core1_1.MemoryRequirements2{
		NextOutData: common.NextOutData{Next: &memReqs},
	}
	err := driver.GetImageMemoryRequirements2(
		core1_1.ImageMemoryRequirementsInfo2{
			Image: image,
		}, &outData)
	require.NoError(t, err)
	require.False(t, memReqs.RequiresDedicatedAllocation)
	require.True(t, memReqs.PrefersDedicatedAllocation)

	require.Equal(t, 1, outData.MemoryRequirements.Size)
	require.Equal(t, 3, outData.MemoryRequirements.Alignment)
	require.Equal(t, uint32(5), outData.MemoryRequirements.MemoryTypeBits)
}

func TestExternalMemoryBufferOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := impl1_1.NewDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})
	mockBuffer := mocks.NewDummyBuffer(device)

	coreLoader.EXPECT().VkCreateBuffer(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkDevice,
			pCreateInfo *loader.VkBufferCreateInfo,
			pAllocator *loader.VkAllocationCallbacks,
			pImage *loader.VkBuffer,
		) (common.VkResult, error) {
			*pImage = mockBuffer.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(12), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
			require.Equal(t, uint64(1), val.FieldByName("size").Uint())
			require.Equal(t, uint64(8), val.FieldByName("usage").Uint())

			next := (*loader.VkExternalMemoryImageCreateInfo)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000072000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(8), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT

			return core1_0.VKSuccess, nil
		})

	buffer, _, err := driver.CreateBuffer(
		device,
		nil,
		core1_0.BufferCreateInfo{
			Size:  1,
			Usage: core1_0.BufferUsageStorageTexelBuffer,

			NextOptions: common.NextOptions{
				core1_1.ExternalMemoryBufferCreateInfo{
					HandleTypes: core1_1.ExternalMemoryHandleTypeD3D11Texture,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockBuffer.Handle(), buffer.Handle())
}

func TestExternalMemoryImageOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := impl1_1.NewDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})
	mockImage := mocks.NewDummyImage(device)

	coreLoader.EXPECT().VkCreateImage(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkDevice,
			pCreateInfo *loader.VkImageCreateInfo,
			pAllocator *loader.VkAllocationCallbacks,
			pImage *loader.VkImage,
		) (common.VkResult, error) {
			*pImage = mockImage.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(14), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO
			require.Equal(t, uint64(1), val.FieldByName("mipLevels").Uint())
			require.Equal(t, uint64(3), val.FieldByName("arrayLayers").Uint())

			next := (*loader.VkExternalMemoryImageCreateInfo)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000072001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x20), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT

			return core1_0.VKSuccess, nil
		})

	image, _, err := driver.CreateImage(
		device,
		nil,
		core1_0.ImageCreateInfo{
			MipLevels:   1,
			ArrayLayers: 3,

			NextOptions: common.NextOptions{
				core1_1.ExternalMemoryImageCreateInfo{
					HandleTypes: core1_1.ExternalMemoryHandleTypeD3D12Heap,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockImage.Handle(), image.Handle())
}

func TestExternalImageFormatOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := impl1_1.NewInstanceDriver(coreLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	coreLoader.EXPECT().VkGetPhysicalDeviceImageFormatProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		physicalDevice loader.VkPhysicalDevice,
		pImageFormatInfo *loader.VkPhysicalDeviceImageFormatInfo2,
		pImageFormatProperties *loader.VkImageFormatProperties2,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pImageFormatInfo).Elem()

		require.Equal(t, uint64(1000059004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2_KHR
		require.Equal(t, uint64(68), val.FieldByName("format").Uint())        // VK_FORMAT_A2B10G10R10_UINT_PACK32

		next := (*loader.VkPhysicalDeviceExternalImageFormatInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000071000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("handleType").Uint()) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT_KHR

		val = reflect.ValueOf(pImageFormatProperties).Elem()

		require.Equal(t, uint64(1000059003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2_KHR

		outDataNext := (*loader.VkExternalImageFormatProperties)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(outDataNext).Elem()

		require.Equal(t, uint64(1000071001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalMemoryProperties").FieldByName("externalMemoryFeatures").UnsafeAddr())) = uint32(4)        // VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_KHR
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalMemoryProperties").FieldByName("exportFromImportedHandleTypes").UnsafeAddr())) = uint32(8) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT_KHR
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalMemoryProperties").FieldByName("compatibleHandleTypes").UnsafeAddr())) = uint32(0x20)      // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT_KHR

		return core1_0.VKSuccess, nil
	})

	var outData core1_1.ExternalImageFormatProperties
	format := core1_1.ImageFormatProperties2{
		NextOutData: common.NextOutData{&outData},
	}
	_, err := driver.GetPhysicalDeviceImageFormatProperties2(
		physicalDevice,
		core1_1.PhysicalDeviceImageFormatInfo2{
			Format: core1_0.FormatA2B10G10R10UnsignedIntPacked,
			NextOptions: common.NextOptions{
				core1_1.PhysicalDeviceExternalImageFormatInfo{
					HandleType: core1_1.ExternalMemoryHandleTypeOpaqueFD,
				},
			},
		},
		&format,
	)
	require.NoError(t, err)
	require.Equal(t, core1_1.ExternalImageFormatProperties{
		ExternalMemoryProperties: core1_1.ExternalMemoryProperties{
			ExternalMemoryFeatures:        core1_1.ExternalMemoryFeatureImportable,
			ExportFromImportedHandleTypes: core1_1.ExternalMemoryHandleTypeD3D11Texture,
			CompatibleHandleTypes:         core1_1.ExternalMemoryHandleTypeD3D12Heap,
		},
	}, outData)
}

func TestExternalMemoryAllocateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_1.NewDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})
	mockMemory := mocks.NewDummyDeviceMemory(device, 1)

	coreLoader.EXPECT().VkAllocateMemory(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkDevice,
			pAllocateInfo *loader.VkMemoryAllocateInfo,
			pAllocator *loader.VkAllocationCallbacks,
			pMemory *loader.VkDeviceMemory,
		) (common.VkResult, error) {
			*pMemory = mockMemory.Handle()

			val := reflect.ValueOf(pAllocateInfo).Elem()
			require.Equal(t, uint64(5), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
			require.Equal(t, uint64(1), val.FieldByName("allocationSize").Uint())
			require.Equal(t, uint64(3), val.FieldByName("memoryTypeIndex").Uint())

			next := (*loader.VkExportMemoryAllocateInfo)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000072002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x10), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT_KHR

			return core1_0.VKSuccess, nil
		})

	memory, _, err := driver.AllocateMemory(device, nil, core1_0.MemoryAllocateInfo{
		AllocationSize:  1,
		MemoryTypeIndex: 3,
		NextOptions: common.NextOptions{
			core1_1.ExportMemoryAllocateInfo{
				HandleTypes: core1_1.ExternalMemoryHandleTypeD3D11TextureKMT,
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, mockMemory.Handle(), memory.Handle())
}
