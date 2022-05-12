package internal1_1_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/internal/dummies"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestMemoryDedicatedAllocateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil).AnyTimes()
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)

	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	deviceHandle := mocks.EasyMockDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkCreateDevice(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pCreateInfo *driver.VkDeviceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDevice *driver.VkDevice) (common.VkResult, error) {
			*pDevice = deviceHandle.Handle()

			return core1_0.VKSuccess, nil
		})

	device, _, err := loader.CreateDevice(physicalDevice, nil, core1_0.DeviceCreateOptions{
		QueueFamilies: []core1_0.DeviceQueueCreateOptions{
			{
				CreatedQueuePriorities: []float32{0},
			},
		},
	})
	require.NoError(t, err)

	buffer := mocks.EasyMockBuffer(ctrl)
	expectedMemory := mocks.EasyMockDeviceMemory(ctrl)

	coreDriver.EXPECT().VkAllocateMemory(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pAllocateInfo *driver.VkMemoryAllocateInfo, pAllocator *driver.VkAllocationCallbacks, pMemory *driver.VkDeviceMemory) (common.VkResult, error) {
			*pMemory = expectedMemory.Handle()

			options := reflect.ValueOf(pAllocateInfo).Elem()
			require.Equal(t, uint64(5), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
			require.Equal(t, uint64(1), options.FieldByName("allocationSize").Uint())
			require.Equal(t, uint64(3), options.FieldByName("memoryTypeIndex").Uint())

			dedicatedPtr := (*driver.VkMemoryDedicatedAllocateInfo)(options.FieldByName("pNext").UnsafePointer())
			dedicated := reflect.ValueOf(dedicatedPtr).Elem()

			require.Equal(t, uint64(1000127001), dedicated.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO
			require.True(t, dedicated.FieldByName("pNext").IsNil())
			require.Equal(t, buffer.Handle(), driver.VkBuffer(dedicated.FieldByName("buffer").UnsafePointer()))
			require.True(t, dedicated.FieldByName("image").IsNil())

			return core1_0.VKSuccess, nil
		})

	memory, _, err := loader.AllocateMemory(device, nil, core1_0.MemoryAllocateOptions{
		AllocationSize:  1,
		MemoryTypeIndex: 3,
		HaveNext: common.HaveNext{Next: core1_1.MemoryDedicatedAllocationOptions{
			Buffer: buffer,
		}},
	})
	require.NoError(t, err)
	require.Equal(t, expectedMemory.Handle(), memory.Handle())
}

func TestDedicatedMemoryRequirementsOutData_Buffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)

	device := dummies.EasyDummyDevice(t, ctrl, loader)
	buffer := mocks.EasyMockBuffer(ctrl)

	coreDriver.EXPECT().VkGetBufferMemoryRequirements2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device driver.VkDevice,
			pInfo *driver.VkBufferMemoryRequirementsInfo2,
			pMemoryRequirements *driver.VkMemoryRequirements2,
		) {
			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000146000), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2
			require.True(t, options.FieldByName("pNext").IsNil())
			require.Equal(t, buffer.Handle(), driver.VkBuffer(options.FieldByName("buffer").UnsafePointer()))

			outData := reflect.ValueOf(pMemoryRequirements).Elem()
			require.Equal(t, uint64(1000146003), outData.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2

			memoryRequirements := outData.FieldByName("memoryRequirements")
			*(*driver.VkDeviceSize)(unsafe.Pointer(memoryRequirements.FieldByName("size").UnsafeAddr())) = driver.VkDeviceSize(1)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memoryRequirements.FieldByName("alignment").UnsafeAddr())) = driver.VkDeviceSize(3)
			*(*driver.Uint32)(unsafe.Pointer(memoryRequirements.FieldByName("memoryTypeBits").UnsafeAddr())) = driver.Uint32(5)

			dedicatedPtr := (*driver.VkMemoryDedicatedRequirements)(outData.FieldByName("pNext").UnsafePointer())
			dedicated := reflect.ValueOf(dedicatedPtr).Elem()
			require.Equal(t, uint64(1000127000), dedicated.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS
			require.True(t, dedicated.FieldByName("pNext").IsNil())
			*(*driver.VkBool32)(unsafe.Pointer(dedicated.FieldByName("prefersDedicatedAllocation").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(dedicated.FieldByName("requiresDedicatedAllocation").UnsafeAddr())) = driver.VkBool32(0)
		})

	var memReqs core1_1.MemoryDedicatedAllocationOutData
	var outData = core1_1.MemoryRequirementsOutData{
		HaveNext: common.HaveNext{Next: &memReqs},
	}
	err = device.Core1_1().BufferMemoryRequirements(
		core1_1.BufferMemoryRequirementsOptions{
			Buffer: buffer,
		}, &outData)
	require.NoError(t, err)
	require.False(t, memReqs.DedicatedRequired)
	require.True(t, memReqs.DedicatedPreferred)

	require.Equal(t, 1, outData.MemoryRequirements.Size)
	require.Equal(t, 3, outData.MemoryRequirements.Alignment)
	require.Equal(t, uint32(5), outData.MemoryRequirements.MemoryType)
}

func TestDedicatedMemoryRequirementsOutData_Image(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)
	device := dummies.EasyDummyDevice(t, ctrl, loader)
	image := mocks.EasyMockImage(ctrl)

	coreDriver.EXPECT().VkGetImageMemoryRequirements2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device driver.VkDevice,
			pInfo *driver.VkImageMemoryRequirementsInfo2,
			pMemoryRequirements *driver.VkMemoryRequirements2,
		) {
			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000146001), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2
			require.True(t, options.FieldByName("pNext").IsNil())
			require.Equal(t, image.Handle(), driver.VkImage(options.FieldByName("image").UnsafePointer()))

			outData := reflect.ValueOf(pMemoryRequirements).Elem()
			require.Equal(t, uint64(1000146003), outData.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2

			memoryRequirements := outData.FieldByName("memoryRequirements")
			*(*driver.VkDeviceSize)(unsafe.Pointer(memoryRequirements.FieldByName("size").UnsafeAddr())) = driver.VkDeviceSize(1)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memoryRequirements.FieldByName("alignment").UnsafeAddr())) = driver.VkDeviceSize(3)
			*(*driver.Uint32)(unsafe.Pointer(memoryRequirements.FieldByName("memoryTypeBits").UnsafeAddr())) = driver.Uint32(5)

			dedicatedPtr := (*driver.VkMemoryDedicatedRequirements)(outData.FieldByName("pNext").UnsafePointer())
			dedicated := reflect.ValueOf(dedicatedPtr).Elem()
			require.Equal(t, uint64(1000127000), dedicated.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS
			require.True(t, dedicated.FieldByName("pNext").IsNil())
			*(*driver.VkBool32)(unsafe.Pointer(dedicated.FieldByName("prefersDedicatedAllocation").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(dedicated.FieldByName("requiresDedicatedAllocation").UnsafeAddr())) = driver.VkBool32(0)
		})

	var memReqs core1_1.MemoryDedicatedAllocationOutData
	var outData = core1_1.MemoryRequirementsOutData{
		HaveNext: common.HaveNext{Next: &memReqs},
	}
	err = device.Core1_1().ImageMemoryRequirements(
		core1_1.ImageMemoryRequirementsOptions{
			Image: image,
		}, &outData)
	require.NoError(t, err)
	require.False(t, memReqs.DedicatedRequired)
	require.True(t, memReqs.DedicatedPreferred)

	require.Equal(t, 1, outData.MemoryRequirements.Size)
	require.Equal(t, 3, outData.MemoryRequirements.Alignment)
	require.Equal(t, uint32(5), outData.MemoryRequirements.MemoryType)
}

func TestExternalMemoryBufferOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	mockBuffer := mocks.EasyMockBuffer(ctrl)

	coreDriver.EXPECT().VkCreateBuffer(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice,
			pCreateInfo *driver.VkBufferCreateInfo,
			pAllocator *driver.VkAllocationCallbacks,
			pImage *driver.VkBuffer,
		) (common.VkResult, error) {
			*pImage = mockBuffer.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(12), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
			require.Equal(t, uint64(1), val.FieldByName("size").Uint())
			require.Equal(t, uint64(8), val.FieldByName("usage").Uint())

			next := (*driver.VkExternalMemoryImageCreateInfo)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000072000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(8), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT

			return core1_0.VKSuccess, nil
		})

	buffer, _, err := loader.CreateBuffer(
		device,
		nil,
		core1_0.BufferCreateOptions{
			BufferSize: 1,
			Usage:      core1_0.BufferUsageStorageTexelBuffer,

			HaveNext: common.HaveNext{
				core1_1.ExternalMemoryBufferOptions{
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

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	mockImage := mocks.EasyMockImage(ctrl)

	coreDriver.EXPECT().VkCreateImage(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice,
			pCreateInfo *driver.VkImageCreateInfo,
			pAllocator *driver.VkAllocationCallbacks,
			pImage *driver.VkImage,
		) (common.VkResult, error) {
			*pImage = mockImage.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(14), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO
			require.Equal(t, uint64(1), val.FieldByName("mipLevels").Uint())
			require.Equal(t, uint64(3), val.FieldByName("arrayLayers").Uint())

			next := (*driver.VkExternalMemoryImageCreateInfo)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000072001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x20), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT

			return core1_0.VKSuccess, nil
		})

	image, _, err := loader.CreateImage(
		device,
		nil,
		core1_0.ImageCreateOptions{
			MipLevels:   1,
			ArrayLayers: 3,

			HaveNext: common.HaveNext{
				core1_1.ExternalMemoryImageOptions{
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

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)
	physicalDevice := dummies.EasyDummyPhysicalDevice(t, loader)

	coreDriver.EXPECT().VkGetPhysicalDeviceImageFormatProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		physicalDevice driver.VkPhysicalDevice,
		pImageFormatInfo *driver.VkPhysicalDeviceImageFormatInfo2,
		pImageFormatProperties *driver.VkImageFormatProperties2,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pImageFormatInfo).Elem()

		require.Equal(t, uint64(1000059004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2_KHR
		require.Equal(t, uint64(68), val.FieldByName("format").Uint())        // VK_FORMAT_A2B10G10R10_UINT_PACK32

		next := (*driver.VkPhysicalDeviceExternalImageFormatInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000071000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("handleType").Uint()) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT_KHR

		val = reflect.ValueOf(pImageFormatProperties).Elem()

		require.Equal(t, uint64(1000059003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2_KHR

		outDataNext := (*driver.VkExternalImageFormatProperties)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(outDataNext).Elem()

		require.Equal(t, uint64(1000071001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalMemoryProperties").FieldByName("externalMemoryFeatures").UnsafeAddr())) = uint32(4)        // VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_KHR
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalMemoryProperties").FieldByName("exportFromImportedHandleTypes").UnsafeAddr())) = uint32(8) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT_KHR
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalMemoryProperties").FieldByName("compatibleHandleTypes").UnsafeAddr())) = uint32(0x20)      // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT_KHR

		return core1_0.VKSuccess, nil
	})

	var outData core1_1.ExternalImageFormatOutData
	format := core1_1.ImageFormatPropertiesOutData{
		HaveNext: common.HaveNext{&outData},
	}
	_, err = physicalDevice.Core1_1Instance().ImageFormatProperties(
		core1_1.ImageFormatOptions{
			Format: core1_0.DataFormatA2B10G10R10UnsignedIntPacked,
			HaveNext: common.HaveNext{
				core1_1.PhysicalDeviceExternalImageFormatOptions{
					HandleType: core1_1.ExternalMemoryHandleTypeOpaqueFD,
				},
			},
		},
		&format,
	)
	require.NoError(t, err)
	require.Equal(t, core1_1.ExternalImageFormatOutData{
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

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	mockMemory := mocks.EasyMockDeviceMemory(ctrl)

	coreDriver.EXPECT().VkAllocateMemory(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice,
			pAllocateInfo *driver.VkMemoryAllocateInfo,
			pAllocator *driver.VkAllocationCallbacks,
			pMemory *driver.VkDeviceMemory,
		) (common.VkResult, error) {
			*pMemory = mockMemory.Handle()

			val := reflect.ValueOf(pAllocateInfo).Elem()
			require.Equal(t, uint64(5), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
			require.Equal(t, uint64(1), val.FieldByName("allocationSize").Uint())
			require.Equal(t, uint64(3), val.FieldByName("memoryTypeIndex").Uint())

			next := (*driver.VkExportMemoryAllocateInfo)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000072002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x10), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT_KHR

			return core1_0.VKSuccess, nil
		})

	memory, _, err := loader.AllocateMemory(device, nil, core1_0.MemoryAllocateOptions{
		AllocationSize:  1,
		MemoryTypeIndex: 3,
		HaveNext: common.HaveNext{
			core1_1.ExportMemoryAllocateOptions{
				HandleTypes: core1_1.ExternalMemoryHandleTypeD3D11TextureKMT,
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, mockMemory.Handle(), memory.Handle())
}
