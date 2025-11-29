package core1_1_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	"github.com/vkngwrapper/core/v2/internal/dummies"
	"github.com/vkngwrapper/core/v2/mocks"
	"go.uber.org/mock/gomock"
)

func TestImagePlaneMemoryRequirementsOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := core1_1.PromoteDevice(dummies.EasyDummyDevice(coreDriver))

	image := mocks.EasyMockImage(ctrl)

	coreDriver.EXPECT().VkGetImageMemoryRequirements2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pInfo *driver.VkImageMemoryRequirementsInfo2,
		pMemoryRequirements *driver.VkMemoryRequirements2,
	) {
		val := reflect.ValueOf(pInfo).Elem()
		require.Equal(t, uint64(1000146001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2
		require.Equal(t, image.Handle(), driver.VkImage(val.FieldByName("image").UnsafePointer()))

		next := (*driver.VkImagePlaneMemoryRequirementsInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000156003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0x20), val.FieldByName("planeAspect").Uint()) // VK_IMAGE_ASPECT_PLANE_1_BIT

		val = reflect.ValueOf(pMemoryRequirements).Elem()
		require.Equal(t, uint64(1000146003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2
		require.True(t, val.FieldByName("pNext").IsNil())

		*(*uint64)(unsafe.Pointer(val.FieldByName("memoryRequirements").FieldByName("size").UnsafeAddr())) = uint64(17)
		*(*uint64)(unsafe.Pointer(val.FieldByName("memoryRequirements").FieldByName("alignment").UnsafeAddr())) = uint64(19)
		*(*uint32)(unsafe.Pointer(val.FieldByName("memoryRequirements").FieldByName("memoryTypeBits").UnsafeAddr())) = uint32(7)
	})

	var outData core1_1.MemoryRequirements2
	err := device.ImageMemoryRequirements2(
		core1_1.ImageMemoryRequirementsInfo2{
			Image: image,
			NextOptions: common.NextOptions{
				core1_1.ImagePlaneMemoryRequirementsInfo{
					PlaneAspect: core1_1.ImageAspectPlane1,
				},
			},
		},
		&outData)
	require.NoError(t, err)
	require.Equal(t, core1_1.MemoryRequirements2{
		MemoryRequirements: core1_0.MemoryRequirements{
			Size:           17,
			Alignment:      19,
			MemoryTypeBits: 7,
		},
	}, outData)
}

func TestSamplerYcbcrConversionOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := dummies.EasyDummyDevice(coreDriver)
	image := mocks.EasyMockImage(ctrl)
	ycbcr := mocks.EasyMockSamplerYcbcrConversion(ctrl)
	mockImageView := mocks.EasyMockImageView(ctrl)

	coreDriver.EXPECT().VkCreateImageView(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkImageViewCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pView *driver.VkImageView,
	) (common.VkResult, error) {
		*pView = mockImageView.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(15), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO
		require.Equal(t, image.Handle(), driver.VkImage(val.FieldByName("image").UnsafePointer()))
		require.Equal(t, uint64(1000156028), val.FieldByName("format").Uint()) // VK_FORMAT_B16G16R16G16_422_UNORM

		next := (*driver.VkSamplerYcbcrConversionInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000156001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, ycbcr.Handle(), driver.VkSamplerYcbcrConversion(val.FieldByName("conversion").UnsafePointer()))

		return core1_0.VKSuccess, nil
	})

	imageView, _, err := device.CreateImageView(
		nil,
		core1_0.ImageViewCreateInfo{
			Image:  image,
			Format: core1_1.FormatB16G16R16G16HorizontalChroma,

			NextOptions: common.NextOptions{
				core1_1.SamplerYcbcrConversionInfo{
					Conversion: ycbcr,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockImageView.Handle(), imageView.Handle())
}

func TestSamplerYcbcrImageFormatOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := dummies.EasyDummyInstance(coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceImageFormatProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice,
			pImageFormatInfo *driver.VkPhysicalDeviceImageFormatInfo2,
			pImageFormatProperties *driver.VkImageFormatProperties2,
		) (common.VkResult, error) {
			val := reflect.ValueOf(pImageFormatInfo).Elem()
			require.Equal(t, uint64(1000059004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2
			require.True(t, val.FieldByName("pNext").IsNil())

			val = reflect.ValueOf(pImageFormatProperties).Elem()
			require.Equal(t, uint64(1000059003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2

			next := (*driver.VkSamplerYcbcrConversionImageFormatProperties)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000156005), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES
			require.True(t, val.FieldByName("pNext").IsNil())
			*(*uint32)(unsafe.Pointer(val.FieldByName("combinedImageSamplerDescriptorCount").UnsafeAddr())) = uint32(7)

			return core1_0.VKSuccess, nil
		})

	var outData core1_1.SamplerYcbcrConversionImageFormatProperties
	_, err := physicalDevice.InstanceScopedPhysicalDevice1_1().ImageFormatProperties2(
		core1_1.PhysicalDeviceImageFormatInfo2{},
		&core1_1.ImageFormatProperties2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, core1_1.SamplerYcbcrConversionImageFormatProperties{
		CombinedImageSamplerDescriptorCount: 7,
	}, outData)
}
