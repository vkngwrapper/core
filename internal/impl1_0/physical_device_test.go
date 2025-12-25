package impl1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func strToCharSlice(text string, slice []loader.Char) {
	byteSlice := []byte(text)
	for idx, b := range byteSlice {
		slice[idx] = loader.Char(b)
	}
	slice[len(byteSlice)] = loader.Char(0)
}

func TestVulkanPhysicalDevice_AvailableExtensionsForLayer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewInstanceDriver(mockLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	mockLoader.EXPECT().VkEnumerateDeviceExtensionProperties(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), gomock.Nil()).DoAndReturn(
		func(physDevice loader.VkPhysicalDevice, pLayerName *loader.Char, pPropertyCount *loader.Uint32, pProperties *loader.VkExtensionProperties) (common.VkResult, error) {
			pLayerBytes := (*byte)(unsafe.Pointer(pLayerName))

			layerByteSlice := ([]byte)(unsafe.Slice(pLayerBytes, 9))
			layer := string(layerByteSlice)

			require.Equal(t, "someLayer", layer)

			*pPropertyCount = 2

			return core1_0.VKSuccess, nil
		})

	mockLoader.EXPECT().VkEnumerateDeviceExtensionProperties(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(physDevice loader.VkPhysicalDevice, pLayerName *loader.Char, pPropertyCount *loader.Uint32, pProperties *loader.VkExtensionProperties) (common.VkResult, error) {
			pLayerBytes := (*byte)(unsafe.Pointer(pLayerName))

			layerByteSlice := ([]byte)(unsafe.Slice(pLayerBytes, 9))
			layer := string(layerByteSlice)

			require.Equal(t, "someLayer", layer)

			*pPropertyCount = 2
			propertySlice := reflect.ValueOf(([]loader.VkExtensionProperties)(unsafe.Slice(pProperties, 2)))

			extension := propertySlice.Index(0)

			*(*uint32)(unsafe.Pointer(extension.FieldByName("specVersion").UnsafeAddr())) = uint32(123)
			extensionName := ([]loader.Char)(unsafe.Slice((*loader.Char)(unsafe.Pointer(extension.FieldByName("extensionName").UnsafeAddr())), 256))
			strToCharSlice("extension 1", extensionName)

			extension = propertySlice.Index(1)
			*(*uint32)(unsafe.Pointer(extension.FieldByName("specVersion").UnsafeAddr())) = uint32(321)
			extensionName = ([]loader.Char)(unsafe.Slice((*loader.Char)(unsafe.Pointer(extension.FieldByName("extensionName").UnsafeAddr())), 256))
			strToCharSlice("extension A", extensionName)

			return core1_0.VKSuccess, nil
		})

	extensions, _, err := driver.EnumerateDeviceExtensionPropertiesForLayer(physicalDevice, "someLayer")
	require.NoError(t, err)
	require.Len(t, extensions, 2)

	extension := extensions["extension 1"]
	require.NotNil(t, extension)
	require.Equal(t, "extension 1", extension.ExtensionName)
	require.Equal(t, uint(123), extension.SpecVersion)

	extension = extensions["extension A"]
	require.NotNil(t, extension)
	require.Equal(t, "extension A", extension.ExtensionName)
	require.Equal(t, uint(321), extension.SpecVersion)
}

func TestVulkanPhysicalDevice_AvailableExtensions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewInstanceDriver(mockLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	mockLoader.EXPECT().VkEnumerateDeviceExtensionProperties(physicalDevice.Handle(), gomock.Nil(), gomock.Not(gomock.Nil()), gomock.Nil()).DoAndReturn(
		func(physDevice loader.VkPhysicalDevice, pLayerName *loader.Char, pPropertyCount *loader.Uint32, pProperties *loader.VkExtensionProperties) (common.VkResult, error) {
			*pPropertyCount = loader.Uint32(2)

			return core1_0.VKSuccess, nil
		})
	mockLoader.EXPECT().VkEnumerateDeviceExtensionProperties(physicalDevice.Handle(), gomock.Nil(), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(physDevice loader.VkPhysicalDevice, pLayerName *loader.Char, pPropertyCount *loader.Uint32, pProperties *loader.VkExtensionProperties) (common.VkResult, error) {
			*pPropertyCount = 2
			propertySlice := reflect.ValueOf(([]loader.VkExtensionProperties)(unsafe.Slice(pProperties, 2)))

			extension := propertySlice.Index(0)

			*(*uint32)(unsafe.Pointer(extension.FieldByName("specVersion").UnsafeAddr())) = uint32(123)
			extensionName := ([]loader.Char)(unsafe.Slice((*loader.Char)(unsafe.Pointer(extension.FieldByName("extensionName").UnsafeAddr())), 256))
			strToCharSlice("extension 1", extensionName)

			extension = propertySlice.Index(1)
			*(*uint32)(unsafe.Pointer(extension.FieldByName("specVersion").UnsafeAddr())) = uint32(321)
			extensionName = ([]loader.Char)(unsafe.Slice((*loader.Char)(unsafe.Pointer(extension.FieldByName("extensionName").UnsafeAddr())), 256))
			strToCharSlice("extension A", extensionName)

			return core1_0.VKSuccess, nil
		})
	extensions, _, err := driver.EnumerateDeviceExtensionProperties(physicalDevice)
	require.NoError(t, err)
	require.Len(t, extensions, 2)

	extension := extensions["extension 1"]
	require.NotNil(t, extension)
	require.Equal(t, "extension 1", extension.ExtensionName)
	require.Equal(t, uint(123), extension.SpecVersion)

	extension = extensions["extension A"]
	require.NotNil(t, extension)
	require.Equal(t, "extension A", extension.ExtensionName)
	require.Equal(t, uint(321), extension.SpecVersion)
}

func TestVulkanPhysicalDevice_AvailableExtensions_Incomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewInstanceDriver(mockLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	mockLoader.EXPECT().VkEnumerateDeviceExtensionProperties(physicalDevice.Handle(), nil, gomock.Not(gomock.Nil()), gomock.Nil()).DoAndReturn(
		func(physDevice loader.VkPhysicalDevice, pLayerName *loader.Char, pPropertyCount *loader.Uint32, pProperties *loader.VkExtensionProperties) (common.VkResult, error) {
			*pPropertyCount = loader.Uint32(2)

			return core1_0.VKSuccess, nil
		})
	mockLoader.EXPECT().VkEnumerateDeviceExtensionProperties(physicalDevice.Handle(), nil, gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(physDevice loader.VkPhysicalDevice, pLayerName *loader.Char, pPropertyCount *loader.Uint32, pProperties *loader.VkExtensionProperties) (common.VkResult, error) {
			return core1_0.VKIncomplete, nil
		})
	mockLoader.EXPECT().VkEnumerateDeviceExtensionProperties(physicalDevice.Handle(), nil, gomock.Not(gomock.Nil()), gomock.Nil()).DoAndReturn(
		func(physDevice loader.VkPhysicalDevice, pLayerName *loader.Char, pPropertyCount *loader.Uint32, pProperties *loader.VkExtensionProperties) (common.VkResult, error) {
			*pPropertyCount = loader.Uint32(2)

			return core1_0.VKSuccess, nil
		})
	mockLoader.EXPECT().VkEnumerateDeviceExtensionProperties(physicalDevice.Handle(), nil, gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(physDevice loader.VkPhysicalDevice, pLayerName *loader.Char, pPropertyCount *loader.Uint32, pProperties *loader.VkExtensionProperties) (common.VkResult, error) {
			*pPropertyCount = 2
			propertySlice := reflect.ValueOf(([]loader.VkExtensionProperties)(unsafe.Slice(pProperties, 2)))

			extension := propertySlice.Index(0)

			*(*uint32)(unsafe.Pointer(extension.FieldByName("specVersion").UnsafeAddr())) = uint32(123)
			extensionName := ([]loader.Char)(unsafe.Slice((*loader.Char)(unsafe.Pointer(extension.FieldByName("extensionName").UnsafeAddr())), 256))
			strToCharSlice("extension 1", extensionName)

			extension = propertySlice.Index(1)
			*(*uint32)(unsafe.Pointer(extension.FieldByName("specVersion").UnsafeAddr())) = uint32(321)
			extensionName = ([]loader.Char)(unsafe.Slice((*loader.Char)(unsafe.Pointer(extension.FieldByName("extensionName").UnsafeAddr())), 256))
			strToCharSlice("extension A", extensionName)

			return core1_0.VKSuccess, nil
		})
	extensions, _, err := driver.EnumerateDeviceExtensionProperties(physicalDevice)
	require.NoError(t, err)
	require.Len(t, extensions, 2)

	extension := extensions["extension 1"]
	require.NotNil(t, extension)
	require.Equal(t, "extension 1", extension.ExtensionName)
	require.Equal(t, uint(123), extension.SpecVersion)

	extension = extensions["extension A"]
	require.NotNil(t, extension)
	require.Equal(t, "extension A", extension.ExtensionName)
	require.Equal(t, uint(321), extension.SpecVersion)
}

func TestVulkanPhysicalDevice_AvailableLayers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewInstanceDriver(mockLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	mockLoader.EXPECT().VkEnumerateDeviceLayerProperties(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Nil()).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice, pPropertyCount *loader.Uint32, pProperties *loader.VkLayerProperties) (common.VkResult, error) {
			*pPropertyCount = loader.Uint32(2)

			return core1_0.VKSuccess, nil
		})
	mockLoader.EXPECT().VkEnumerateDeviceLayerProperties(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice, pPropertyCount *loader.Uint32, pProperties *loader.VkLayerProperties) (common.VkResult, error) {
			*pPropertyCount = 2
			propertySlice := reflect.ValueOf(([]loader.VkLayerProperties)(unsafe.Slice(pProperties, 2)))

			layer := propertySlice.Index(0)

			*(*uint32)(unsafe.Pointer(layer.FieldByName("specVersion").UnsafeAddr())) = uint32(common.CreateVersion(1, 2, 3))
			*(*uint32)(unsafe.Pointer(layer.FieldByName("implementationVersion").UnsafeAddr())) = uint32(common.CreateVersion(2, 1, 3))
			layerName := ([]loader.Char)(unsafe.Slice((*loader.Char)(unsafe.Pointer(layer.FieldByName("layerName").UnsafeAddr())), 256))
			strToCharSlice("layer 1", layerName)
			layerDesc := ([]loader.Char)(unsafe.Slice((*loader.Char)(unsafe.Pointer(layer.FieldByName("description").UnsafeAddr())), 256))
			strToCharSlice("a cool layer", layerDesc)

			layer = propertySlice.Index(1)

			*(*uint32)(unsafe.Pointer(layer.FieldByName("specVersion").UnsafeAddr())) = uint32(common.CreateVersion(3, 2, 1))
			*(*uint32)(unsafe.Pointer(layer.FieldByName("implementationVersion").UnsafeAddr())) = uint32(common.CreateVersion(2, 3, 1))
			layerName = ([]loader.Char)(unsafe.Slice((*loader.Char)(unsafe.Pointer(layer.FieldByName("layerName").UnsafeAddr())), 256))
			strToCharSlice("layer A", layerName)
			layerDesc = ([]loader.Char)(unsafe.Slice((*loader.Char)(unsafe.Pointer(layer.FieldByName("description").UnsafeAddr())), 256))
			strToCharSlice("a bad layer", layerDesc)

			return core1_0.VKSuccess, nil
		})

	layers, _, err := driver.EnumerateDeviceLayerProperties(physicalDevice)
	require.NoError(t, err)
	require.Len(t, layers, 2)

	layer := layers["layer 1"]
	require.NotNil(t, layer)
	require.Equal(t, "layer 1", layer.LayerName)
	require.Equal(t, "a cool layer", layer.Description)
	require.Equal(t, common.CreateVersion(1, 2, 3), layer.SpecVersion)
	require.Equal(t, common.CreateVersion(2, 1, 3), layer.ImplementationVersion)

	layer = layers["layer A"]
	require.NotNil(t, layer)
	require.Equal(t, "layer A", layer.LayerName)
	require.Equal(t, "a bad layer", layer.Description)
	require.Equal(t, common.CreateVersion(3, 2, 1), layer.SpecVersion)
	require.Equal(t, common.CreateVersion(2, 3, 1), layer.ImplementationVersion)
}

func TestVulkanPhysicalDevice_AvailableLayers_Incomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewInstanceDriver(mockLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	mockLoader.EXPECT().VkEnumerateDeviceLayerProperties(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Nil()).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice, pPropertyCount *loader.Uint32, pProperties *loader.VkLayerProperties) (common.VkResult, error) {
			*pPropertyCount = loader.Uint32(2)

			return core1_0.VKSuccess, nil
		})
	mockLoader.EXPECT().VkEnumerateDeviceLayerProperties(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice, pPropertyCount *loader.Uint32, pProperties *loader.VkLayerProperties) (common.VkResult, error) {
			return core1_0.VKIncomplete, nil
		})
	mockLoader.EXPECT().VkEnumerateDeviceLayerProperties(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Nil()).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice, pPropertyCount *loader.Uint32, pProperties *loader.VkLayerProperties) (common.VkResult, error) {
			*pPropertyCount = loader.Uint32(2)

			return core1_0.VKSuccess, nil
		})
	mockLoader.EXPECT().VkEnumerateDeviceLayerProperties(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice, pPropertyCount *loader.Uint32, pProperties *loader.VkLayerProperties) (common.VkResult, error) {
			*pPropertyCount = 2
			propertySlice := reflect.ValueOf(([]loader.VkLayerProperties)(unsafe.Slice(pProperties, 2)))

			layer := propertySlice.Index(0)

			*(*uint32)(unsafe.Pointer(layer.FieldByName("specVersion").UnsafeAddr())) = uint32(common.CreateVersion(1, 2, 3))
			*(*uint32)(unsafe.Pointer(layer.FieldByName("implementationVersion").UnsafeAddr())) = uint32(common.CreateVersion(2, 1, 3))
			layerName := ([]loader.Char)(unsafe.Slice((*loader.Char)(unsafe.Pointer(layer.FieldByName("layerName").UnsafeAddr())), 256))
			strToCharSlice("layer 1", layerName)
			layerDesc := ([]loader.Char)(unsafe.Slice((*loader.Char)(unsafe.Pointer(layer.FieldByName("description").UnsafeAddr())), 256))
			strToCharSlice("a cool layer", layerDesc)

			layer = propertySlice.Index(1)

			*(*uint32)(unsafe.Pointer(layer.FieldByName("specVersion").UnsafeAddr())) = uint32(common.CreateVersion(3, 2, 1))
			*(*uint32)(unsafe.Pointer(layer.FieldByName("implementationVersion").UnsafeAddr())) = uint32(common.CreateVersion(2, 3, 1))
			layerName = ([]loader.Char)(unsafe.Slice((*loader.Char)(unsafe.Pointer(layer.FieldByName("layerName").UnsafeAddr())), 256))
			strToCharSlice("layer A", layerName)
			layerDesc = ([]loader.Char)(unsafe.Slice((*loader.Char)(unsafe.Pointer(layer.FieldByName("description").UnsafeAddr())), 256))
			strToCharSlice("a bad layer", layerDesc)

			return core1_0.VKSuccess, nil
		})

	layers, _, err := driver.EnumerateDeviceLayerProperties(physicalDevice)
	require.NoError(t, err)
	require.Len(t, layers, 2)

	layer := layers["layer 1"]
	require.NotNil(t, layer)
	require.Equal(t, "layer 1", layer.LayerName)
	require.Equal(t, "a cool layer", layer.Description)
	require.Equal(t, common.CreateVersion(1, 2, 3), layer.SpecVersion)
	require.Equal(t, common.CreateVersion(2, 1, 3), layer.ImplementationVersion)

	layer = layers["layer A"]
	require.NotNil(t, layer)
	require.Equal(t, "layer A", layer.LayerName)
	require.Equal(t, "a bad layer", layer.Description)
	require.Equal(t, common.CreateVersion(3, 2, 1), layer.SpecVersion)
	require.Equal(t, common.CreateVersion(2, 3, 1), layer.ImplementationVersion)
}

func TestVulkanPhysicalDevice_QueueFamilyProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewInstanceDriver(mockLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	mockLoader.EXPECT().VkGetPhysicalDeviceQueueFamilyProperties(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Nil()).DoAndReturn(
		func(device loader.VkPhysicalDevice, pPropertyCount *loader.Uint32, pProperties *loader.VkQueueFamilyProperties) {
			*pPropertyCount = 1
		})
	mockLoader.EXPECT().VkGetPhysicalDeviceQueueFamilyProperties(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(device loader.VkPhysicalDevice, pPropertyCount *loader.Uint32, pProperties *loader.VkQueueFamilyProperties) {
			*pPropertyCount = 1
			propertySlice := reflect.ValueOf(([]loader.VkQueueFamilyProperties)(unsafe.Slice(pProperties, 1)))

			property := propertySlice.Index(0)
			*(*loader.VkQueueFlags)(unsafe.Pointer(property.FieldByName("queueFlags").UnsafeAddr())) = loader.VkQueueFlags(8) // VK_QUEUE_SPARSE_BINDING_BIT
			*(*uint32)(unsafe.Pointer(property.FieldByName("queueCount").UnsafeAddr())) = uint32(3)
			*(*uint32)(unsafe.Pointer(property.FieldByName("timestampValidBits").UnsafeAddr())) = uint32(5)

			propertyExtent := property.FieldByName("minImageTransferGranularity")
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("width").UnsafeAddr())) = uint32(7)
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("height").UnsafeAddr())) = uint32(11)
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("depth").UnsafeAddr())) = uint32(13)
		})

	queueProperties := driver.GetPhysicalDeviceQueueFamilyProperties(physicalDevice)
	require.Equal(t, 3, queueProperties[0].QueueCount)
	require.Equal(t, uint32(5), queueProperties[0].TimestampValidBits)
	require.Equal(t, 7, queueProperties[0].MinImageTransferGranularity.Width)
	require.Equal(t, 11, queueProperties[0].MinImageTransferGranularity.Height)
	require.Equal(t, 13, queueProperties[0].MinImageTransferGranularity.Depth)
	require.Equal(t, core1_0.QueueSparseBinding, queueProperties[0].QueueFlags)
}

func TestVulkanPhysicalDevice_Properties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewInstanceDriver(mockLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	deviceUUID, err := uuid.NewUUID()
	require.NoError(t, err)

	mockLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice.Handle(), gomock.Not(gomock.Nil())).DoAndReturn(
		func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			pPropertySlice := reflect.ValueOf(unsafe.Slice(pProperties, 1))
			val := pPropertySlice.Index(0)

			*(*uint32)(unsafe.Pointer(val.FieldByName("apiVersion").UnsafeAddr())) = uint32(common.Vulkan1_1)
			*(*uint32)(unsafe.Pointer(val.FieldByName("driverVersion").UnsafeAddr())) = uint32(common.CreateVersion(3, 2, 1))
			*(*uint32)(unsafe.Pointer(val.FieldByName("vendorID").UnsafeAddr())) = uint32(3)
			*(*uint32)(unsafe.Pointer(val.FieldByName("deviceID").UnsafeAddr())) = uint32(5)
			*(*uint32)(unsafe.Pointer(val.FieldByName("deviceType").UnsafeAddr())) = uint32(2) // VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU
			deviceNamePtr := (*loader.Char)(unsafe.Pointer(val.FieldByName("deviceName").UnsafeAddr()))
			deviceNameSlice := ([]loader.Char)(unsafe.Slice(deviceNamePtr, 256))
			deviceName := "Some Device"
			for i, r := range []byte(deviceName) {
				deviceNameSlice[i] = loader.Char(r)
			}
			deviceNameSlice[len(deviceName)] = 0

			uuidPtr := (*loader.Char)(unsafe.Pointer(val.FieldByName("pipelineCacheUUID").UnsafeAddr()))
			uuidSlice := ([]loader.Char)(unsafe.Slice(uuidPtr, 16))
			uuid, err := deviceUUID.MarshalBinary()
			require.NoError(t, err)

			for i, b := range uuid {
				uuidSlice[i] = loader.Char(b)
			}
			limits := val.FieldByName("limits")
			*(*uint32)(unsafe.Pointer(limits.FieldByName("maxUniformBufferRange").UnsafeAddr())) = uint32(7)
			*(*uint32)(unsafe.Pointer(limits.FieldByName("maxVertexInputBindingStride").UnsafeAddr())) = uint32(11)
			workGroupCount := limits.FieldByName("maxComputeWorkGroupCount")
			*(*uint32)(unsafe.Pointer(workGroupCount.Index(0).UnsafeAddr())) = uint32(13)
			*(*uint32)(unsafe.Pointer(workGroupCount.Index(1).UnsafeAddr())) = uint32(17)
			*(*uint32)(unsafe.Pointer(workGroupCount.Index(2).UnsafeAddr())) = uint32(19)
			*(*float32)(unsafe.Pointer(limits.FieldByName("maxInterpolationOffset").UnsafeAddr())) = float32(23)
			*(*loader.VkBool32)(unsafe.Pointer(limits.FieldByName("strictLines").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkDeviceSize)(unsafe.Pointer(limits.FieldByName("optimalBufferCopyRowPitchAlignment").UnsafeAddr())) = loader.VkDeviceSize(29)

			sparseProperties := val.FieldByName("sparseProperties")
			*(*loader.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyStandard2DBlockShape").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyStandard2DMultisampleBlockShape").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyStandard3DBlockShape").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyAlignedMipSize").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyNonResidentStrict").UnsafeAddr())) = loader.VkBool32(1)
		})

	properties, err := driver.GetPhysicalDeviceProperties(physicalDevice)
	require.NotNil(t, properties)
	require.NoError(t, err)
	require.Equal(t, common.Vulkan1_1, properties.APIVersion)
	require.Equal(t, common.CreateVersion(3, 2, 1), properties.DriverVersion)
	require.Equal(t, uint32(3), properties.VendorID)
	require.Equal(t, uint32(5), properties.DeviceID)
	require.Equal(t, core1_0.PhysicalDeviceTypeDiscreteGPU, properties.DriverType)
	require.Equal(t, "Some Device", properties.DriverName)
	require.Equal(t, deviceUUID, properties.PipelineCacheUUID)

	require.Equal(t, 7, properties.Limits.MaxUniformBufferRange)
	require.Equal(t, 11, properties.Limits.MaxVertexInputBindingStride)
	require.Equal(t, [3]int{13, 17, 19}, properties.Limits.MaxComputeWorkGroupCount)
	require.Equal(t, float32(23), properties.Limits.MaxInterpolationOffset)
	require.True(t, properties.Limits.StrictLines)
	require.Equal(t, 29, properties.Limits.OptimalBufferCopyRowPitchAlignment)

	require.True(t, properties.SparseProperties.ResidencyStandard2DBlockShape)
	require.False(t, properties.SparseProperties.ResidencyStandard2DMultisampleBlockShape)
	require.True(t, properties.SparseProperties.ResidencyStandard3DBlockShape)
	require.False(t, properties.SparseProperties.ResidencyAlignedMipSize)
	require.True(t, properties.SparseProperties.ResidencyNonResidentStrict)
}

func TestVulkanPhysicalDevice_Features(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewInstanceDriver(mockLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	mockLoader.EXPECT().VkGetPhysicalDeviceFeatures(physicalDevice.Handle(), gomock.Not(gomock.Nil())).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice, pFeatures *loader.VkPhysicalDeviceFeatures) {
			featureSlice := reflect.ValueOf(unsafe.Slice(pFeatures, 1))
			val := featureSlice.Index(0)

			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("robustBufferAccess").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("fullDrawIndexUint32").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("imageCubeArray").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("independentBlend").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("geometryShader").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("tessellationShader").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("sampleRateShading").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("dualSrcBlend").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("logicOp").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("multiDrawIndirect").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("drawIndirectFirstInstance").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("depthClamp").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("depthBiasClamp").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("fillModeNonSolid").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("depthBounds").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("wideLines").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("largePoints").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("alphaToOne").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("multiViewport").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("samplerAnisotropy").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("textureCompressionETC2").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("textureCompressionASTC_LDR").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("textureCompressionBC").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("occlusionQueryPrecise").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("pipelineStatisticsQuery").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("vertexPipelineStoresAndAtomics").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("fragmentStoresAndAtomics").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderTessellationAndGeometryPointSize").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderImageGatherExtended").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageExtendedFormats").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageMultisample").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageReadWithoutFormat").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageWriteWithoutFormat").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformBufferArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSampledImageArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageBufferArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderClipDistance").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderCullDistance").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderFloat64").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInt64").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInt16").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderResourceResidency").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderResourceMinLod").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("sparseBinding").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("sparseResidencyBuffer").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("sparseResidencyImage2D").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("sparseResidencyImage3D").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("sparseResidency2Samples").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("sparseResidency4Samples").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("sparseResidency8Samples").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("sparseResidency16Samples").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("sparseResidencyAliased").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("variableMultisampleRate").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("inheritedQueries").UnsafeAddr())) = loader.VkBool32(1)
		})

	features := driver.GetPhysicalDeviceFeatures(physicalDevice)
	require.NotNil(t, features)
	require.True(t, features.RobustBufferAccess)
	require.False(t, features.FullDrawIndexUint32)
	require.True(t, features.ImageCubeArray)
	require.False(t, features.IndependentBlend)
	require.True(t, features.GeometryShader)
	require.False(t, features.TessellationShader)
	require.True(t, features.SampleRateShading)
	require.False(t, features.DualSrcBlend)
	require.True(t, features.LogicOp)
	require.False(t, features.MultiDrawIndirect)
	require.True(t, features.DrawIndirectFirstInstance)
	require.False(t, features.DepthClamp)
	require.True(t, features.DepthBiasClamp)
	require.False(t, features.FillModeNonSolid)
	require.True(t, features.DepthBounds)
	require.False(t, features.WideLines)
	require.True(t, features.LargePoints)
	require.False(t, features.AlphaToOne)
	require.True(t, features.MultiViewport)
	require.False(t, features.SamplerAnisotropy)
	require.True(t, features.TextureCompressionEtc2)
	require.False(t, features.TextureCompressionAstcLdc)
	require.True(t, features.TextureCompressionBc)
	require.False(t, features.OcclusionQueryPrecise)
	require.True(t, features.PipelineStatisticsQuery)
	require.False(t, features.VertexPipelineStoresAndAtomics)
	require.True(t, features.FragmentStoresAndAtomics)
	require.False(t, features.ShaderTessellationAndGeometryPointSize)
	require.True(t, features.ShaderImageGatherExtended)
	require.False(t, features.ShaderStorageImageExtendedFormats)
	require.True(t, features.ShaderStorageImageMultisample)
	require.False(t, features.ShaderStorageImageReadWithoutFormat)
	require.True(t, features.ShaderStorageImageWriteWithoutFormat)
	require.False(t, features.ShaderUniformBufferArrayDynamicIndexing)
	require.True(t, features.ShaderSampledImageArrayDynamicIndexing)
	require.False(t, features.ShaderStorageBufferArrayDynamicIndexing)
	require.True(t, features.ShaderStorageImageArrayDynamicIndexing)
	require.False(t, features.ShaderClipDistance)
	require.True(t, features.ShaderCullDistance)
	require.False(t, features.ShaderFloat64)
	require.True(t, features.ShaderInt64)
	require.False(t, features.ShaderInt16)
	require.True(t, features.ShaderResourceResidency)
	require.False(t, features.ShaderResourceMinLod)
	require.True(t, features.SparseBinding)
	require.False(t, features.SparseResidencyBuffer)
	require.True(t, features.SparseResidencyImage2D)
	require.False(t, features.SparseResidencyImage3D)
	require.True(t, features.SparseResidency2Samples)
	require.False(t, features.SparseResidency4Samples)
	require.True(t, features.SparseResidency8Samples)
	require.False(t, features.SparseResidency16Samples)
	require.True(t, features.SparseResidencyAliased)
	require.False(t, features.VariableMultisampleRate)
	require.True(t, features.InheritedQueries)
}

func TestVulkanPhysicalDevice_FormatProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewInstanceDriver(mockLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	mockLoader.EXPECT().VkGetPhysicalDeviceFormatProperties(physicalDevice.Handle(),
		loader.VkFormat(57), // VK_FORMAT_A8B8G8R8_SRGB_PACK32
		gomock.Not(nil)).DoAndReturn(
		func(device loader.VkPhysicalDevice, format loader.VkFormat, pFormatProperties *loader.VkFormatProperties) {
			val := reflect.ValueOf(pFormatProperties).Elem()

			*(*uint32)(unsafe.Pointer(val.FieldByName("optimalTilingFeatures").UnsafeAddr())) = uint32(0x00000100) // VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT
			*(*uint32)(unsafe.Pointer(val.FieldByName("linearTilingFeatures").UnsafeAddr())) = uint32(0x00000400)  // VK_FORMAT_FEATURE_BLIT_SRC_BIT
			*(*uint32)(unsafe.Pointer(val.FieldByName("bufferFeatures").UnsafeAddr())) = uint32(0x00000010)        // VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT
		})

	props := driver.GetPhysicalDeviceFormatProperties(physicalDevice, core1_0.FormatA8B8G8R8SRGBPacked)
	require.NotNil(t, props)
	require.Equal(t, core1_0.FormatFeatureColorAttachmentBlend, props.OptimalTilingFeatures)
	require.Equal(t, core1_0.FormatFeatureBlitSource, props.LinearTilingFeatures)
	require.Equal(t, core1_0.FormatFeatureStorageTexelBuffer, props.BufferFeatures)
}

func TestVulkanPhysicalDevice_ImageFormatProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewInstanceDriver(mockLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	mockLoader.EXPECT().VkGetPhysicalDeviceImageFormatProperties(physicalDevice.Handle(),
		loader.VkFormat(57),                   // VK_FORMAT_A8B8G8R8_SRGB_PACK32
		loader.VkImageType(1),                 // VK_IMAGE_TYPE_2D
		loader.VkImageTiling(1),               // VK_IMAGE_TILING_LINEAR
		loader.VkImageUsageFlags(8),           // VK_IMAGE_USAGE_STORAGE_BIT
		loader.VkImageCreateFlags(0x00000004), // VK_IMAGE_CREATE_SPARSE_ALIASED_BIT
		gomock.Not(nil),
	).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice,
			format loader.VkFormat,
			imageType loader.VkImageType,
			imageTiling loader.VkImageTiling,
			imageUsages loader.VkImageUsageFlags,
			flags loader.VkImageCreateFlags,
			pProperties *loader.VkImageFormatProperties) (common.VkResult, error) {

			val := reflect.ValueOf(pProperties).Elem()

			*(*uint32)(unsafe.Pointer(val.FieldByName("maxMipLevels").UnsafeAddr())) = uint32(1)
			*(*uint32)(unsafe.Pointer(val.FieldByName("maxArrayLayers").UnsafeAddr())) = uint32(3)
			*(*uint64)(unsafe.Pointer(val.FieldByName("maxResourceSize").UnsafeAddr())) = uint64(5)
			*(*uint32)(unsafe.Pointer(val.FieldByName("sampleCounts").UnsafeAddr())) = uint32(core1_0.Samples8)
			*(*uint32)(unsafe.Pointer(val.FieldByName("maxExtent").FieldByName("width").UnsafeAddr())) = uint32(11)
			*(*uint32)(unsafe.Pointer(val.FieldByName("maxExtent").FieldByName("height").UnsafeAddr())) = uint32(13)
			*(*uint32)(unsafe.Pointer(val.FieldByName("maxExtent").FieldByName("depth").UnsafeAddr())) = uint32(17)

			return core1_0.VKSuccess, nil
		})

	props, _, err := driver.GetPhysicalDeviceImageFormatProperties(physicalDevice, core1_0.FormatA8B8G8R8SRGBPacked, core1_0.ImageType2D, core1_0.ImageTilingLinear, core1_0.ImageUsageStorage, core1_0.ImageCreateSparseAliased)
	require.NoError(t, err)
	require.NotNil(t, props)
	require.Equal(t, 1, props.MaxMipLevels)
	require.Equal(t, 3, props.MaxArrayLayers)
	require.Equal(t, 5, props.MaxResourceSize)
	require.Equal(t, core1_0.Samples8, props.SampleCounts)
	require.Equal(t, 11, props.MaxExtent.Width)
	require.Equal(t, 13, props.MaxExtent.Height)
	require.Equal(t, 17, props.MaxExtent.Depth)
}

func TestVulkanPhysicalDevice_SparseImageFormatProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewInstanceDriver(mockLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	mockLoader.EXPECT().VkGetPhysicalDeviceSparseImageFormatProperties(physicalDevice.Handle(),
		loader.VkFormat(68),                  // VK_FORMAT_A2B10G10R10_UINT_PACK32
		loader.VkImageType(2),                // VK_IMAGE_TYPE_3D
		loader.VkSampleCountFlagBits(8),      // VK_SAMPLE_COUNT_8_BIT
		loader.VkImageUsageFlags(0x00000080), // VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT
		loader.VkImageTiling(1),              // VK_IMAGE_TILING_LINEAR
		gomock.Not(nil),
		nil).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice, format loader.VkFormat, t loader.VkImageType, samples loader.VkSampleCountFlagBits, usage loader.VkImageUsageFlags, tiling loader.VkImageTiling, pPropertyCount *loader.Uint32, pProperties *loader.VkSparseImageFormatProperties) {
			*pPropertyCount = loader.Uint32(2)
		})

	mockLoader.EXPECT().VkGetPhysicalDeviceSparseImageFormatProperties(physicalDevice.Handle(),
		loader.VkFormat(68),                  // VK_FORMAT_A2B10G10R10_UINT_PACK32
		loader.VkImageType(2),                // VK_IMAGE_TYPE_3D
		loader.VkSampleCountFlagBits(8),      // VK_SAMPLE_COUNT_8_BIT
		loader.VkImageUsageFlags(0x00000080), // VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT
		loader.VkImageTiling(1),              // VK_IMAGE_TILING_LINEAR
		gomock.Not(nil),
		gomock.Not(nil)).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice, format loader.VkFormat, imageType loader.VkImageType, samples loader.VkSampleCountFlagBits, usage loader.VkImageUsageFlags, tiling loader.VkImageTiling, pPropertyCount *loader.Uint32, pProperties *loader.VkSparseImageFormatProperties) {
			require.Equal(t, loader.Uint32(2), *pPropertyCount)

			properties := ([]loader.VkSparseImageFormatProperties)(unsafe.Slice(pProperties, 2))
			val := reflect.ValueOf(properties)

			prop := val.Index(0)
			*(*uint32)(unsafe.Pointer(prop.FieldByName("aspectMask").UnsafeAddr())) = uint32(0x00000004) // VK_IMAGE_ASPECT_STENCIL_BIT
			*(*uint32)(unsafe.Pointer(prop.FieldByName("flags").UnsafeAddr())) = uint32(0x00000005)      // VK_SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT | VK_SPARSE_IMAGE_FORMAT_SINGLE_MIPTAIL_BIT

			granularity := prop.FieldByName("imageGranularity")
			*(*uint32)(unsafe.Pointer(granularity.FieldByName("width").UnsafeAddr())) = uint32(1)
			*(*uint32)(unsafe.Pointer(granularity.FieldByName("height").UnsafeAddr())) = uint32(3)
			*(*uint32)(unsafe.Pointer(granularity.FieldByName("depth").UnsafeAddr())) = uint32(5)

			prop = val.Index(1)
			*(*uint32)(unsafe.Pointer(prop.FieldByName("aspectMask").UnsafeAddr())) = uint32(0x00000001) // VK_IMAGE_ASPECT_COLOR_BIT
			*(*uint32)(unsafe.Pointer(prop.FieldByName("flags").UnsafeAddr())) = uint32(0)

			granularity = prop.FieldByName("imageGranularity")
			*(*uint32)(unsafe.Pointer(granularity.FieldByName("width").UnsafeAddr())) = uint32(7)
			*(*uint32)(unsafe.Pointer(granularity.FieldByName("height").UnsafeAddr())) = uint32(11)
			*(*uint32)(unsafe.Pointer(granularity.FieldByName("depth").UnsafeAddr())) = uint32(13)
		})

	props := driver.GetPhysicalDeviceSparseImageFormatProperties(physicalDevice, core1_0.FormatA2B10G10R10UnsignedIntPacked, core1_0.ImageType3D, core1_0.Samples8, core1_0.ImageUsageInputAttachment, core1_0.ImageTilingLinear)
	require.Len(t, props, 2)
	require.Equal(t, core1_0.ImageAspectStencil, props[0].AspectMask)
	require.Equal(t, core1_0.SparseImageFormatNonstandardBlockSize|core1_0.SparseImageFormatSingleMipTail, props[0].Flags)
	require.Equal(t, 1, props[0].ImageGranularity.Width)
	require.Equal(t, 3, props[0].ImageGranularity.Height)
	require.Equal(t, 5, props[0].ImageGranularity.Depth)

	require.Equal(t, core1_0.ImageAspectColor, props[1].AspectMask)
	require.Equal(t, core1_0.SparseImageFormatFlags(0), props[1].Flags)
	require.Equal(t, 7, props[1].ImageGranularity.Width)
	require.Equal(t, 11, props[1].ImageGranularity.Height)
	require.Equal(t, 13, props[1].ImageGranularity.Depth)
}

func TestVulkanPhysicalDevice_MemoryProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewInstanceDriver(mockLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	mockLoader.EXPECT().VkGetPhysicalDeviceMemoryProperties(physicalDevice.Handle(), gomock.Not(nil)).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceMemoryProperties) {
			propertySlice := reflect.ValueOf(unsafe.Slice(pProperties, 1))
			val := propertySlice.Index(0)
			*(*uint32)(unsafe.Pointer(val.FieldByName("memoryTypeCount").UnsafeAddr())) = uint32(1)
			*(*uint32)(unsafe.Pointer(val.FieldByName("memoryHeapCount").UnsafeAddr())) = uint32(1)

			memoryType := val.FieldByName("memoryTypes").Index(0)
			*(*uint32)(unsafe.Pointer(memoryType.FieldByName("heapIndex").UnsafeAddr())) = uint32(3)
			*(*int32)(unsafe.Pointer(memoryType.FieldByName("propertyFlags").UnsafeAddr())) = int32(16) // VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT

			memoryHeap := val.FieldByName("memoryHeaps").Index(0)
			*(*uint64)(unsafe.Pointer(memoryHeap.FieldByName("size").UnsafeAddr())) = uint64(99)
			*(*int32)(unsafe.Pointer(memoryHeap.FieldByName("flags").UnsafeAddr())) = int32(1) // VK_MEMORY_HEAP_DEVICE_LOCAL_BIT
		})

	memoryProps := driver.GetPhysicalDeviceMemoryProperties(physicalDevice)
	require.NotNil(t, memoryProps)
	require.Len(t, memoryProps.MemoryTypes, 1)
	require.Len(t, memoryProps.MemoryHeaps, 1)

	require.Equal(t, 3, memoryProps.MemoryTypes[0].HeapIndex)
	require.Equal(t, core1_0.MemoryPropertyLazilyAllocated, memoryProps.MemoryTypes[0].PropertyFlags)

	require.Equal(t, 99, memoryProps.MemoryHeaps[0].Size)
	require.Equal(t, core1_0.MemoryHeapDeviceLocal, memoryProps.MemoryHeaps[0].Flags)
}
