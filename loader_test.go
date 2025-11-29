package core_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	"go.uber.org/mock/gomock"
)

func strToCharSlice(text string, slice []driver.Char) {
	byteSlice := []byte(text)
	for idx, b := range byteSlice {
		slice[idx] = driver.Char(b)
	}
	slice[len(byteSlice)] = driver.Char(0)
}

func TestVulkanLoader1_0_AvailableExtensionsForLayer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockDriver.EXPECT().VkEnumerateInstanceExtensionProperties(gomock.Not(nil), gomock.Not(nil), nil).DoAndReturn(
		func(pLayerName *driver.Char, pPropertyCount *driver.Uint32, pProperties *driver.VkExtensionProperties) (common.VkResult, error) {
			pLayerBytes := (*byte)(unsafe.Pointer(pLayerName))

			layerByteSlice := ([]byte)(unsafe.Slice(pLayerBytes, 9))
			layer := string(layerByteSlice)

			require.Equal(t, "someLayer", layer)

			*pPropertyCount = 2

			return core1_0.VKSuccess, nil
		})

	mockDriver.EXPECT().VkEnumerateInstanceExtensionProperties(gomock.Not(nil), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(pLayerName *driver.Char, pPropertyCount *driver.Uint32, pProperties *driver.VkExtensionProperties) (common.VkResult, error) {
			pLayerBytes := (*byte)(unsafe.Pointer(pLayerName))

			layerByteSlice := ([]byte)(unsafe.Slice(pLayerBytes, 9))
			layer := string(layerByteSlice)

			require.Equal(t, "someLayer", layer)

			*pPropertyCount = 2
			propertySlice := reflect.ValueOf(([]driver.VkExtensionProperties)(unsafe.Slice(pProperties, 2)))

			extension := propertySlice.Index(0)

			*(*uint32)(unsafe.Pointer(extension.FieldByName("specVersion").UnsafeAddr())) = uint32(123)
			extensionName := ([]driver.Char)(unsafe.Slice((*driver.Char)(unsafe.Pointer(extension.FieldByName("extensionName").UnsafeAddr())), 256))
			strToCharSlice("extension 1", extensionName)

			extension = propertySlice.Index(1)
			*(*uint32)(unsafe.Pointer(extension.FieldByName("specVersion").UnsafeAddr())) = uint32(321)
			extensionName = ([]driver.Char)(unsafe.Slice((*driver.Char)(unsafe.Pointer(extension.FieldByName("extensionName").UnsafeAddr())), 256))
			strToCharSlice("extension A", extensionName)

			return core1_0.VKSuccess, nil
		})

	extensions, _, err := loader.AvailableExtensionsForLayer("someLayer")
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

func TestVulkanLoader1_0_AvailableExtensions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockDriver.EXPECT().VkEnumerateInstanceExtensionProperties(nil, gomock.Not(nil), nil).DoAndReturn(
		func(pLayerName *driver.Char, pPropertyCount *driver.Uint32, pProperties *driver.VkExtensionProperties) (common.VkResult, error) {
			*pPropertyCount = 2

			return core1_0.VKSuccess, nil
		})

	mockDriver.EXPECT().VkEnumerateInstanceExtensionProperties(nil, gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(pLayerName *driver.Char, pPropertyCount *driver.Uint32, pProperties *driver.VkExtensionProperties) (common.VkResult, error) {
			*pPropertyCount = 2
			propertySlice := reflect.ValueOf(([]driver.VkExtensionProperties)(unsafe.Slice(pProperties, 2)))

			extension := propertySlice.Index(0)

			*(*uint32)(unsafe.Pointer(extension.FieldByName("specVersion").UnsafeAddr())) = uint32(123)
			extensionName := ([]driver.Char)(unsafe.Slice((*driver.Char)(unsafe.Pointer(extension.FieldByName("extensionName").UnsafeAddr())), 256))
			strToCharSlice("extension 1", extensionName)

			extension = propertySlice.Index(1)
			*(*uint32)(unsafe.Pointer(extension.FieldByName("specVersion").UnsafeAddr())) = uint32(321)
			extensionName = ([]driver.Char)(unsafe.Slice((*driver.Char)(unsafe.Pointer(extension.FieldByName("extensionName").UnsafeAddr())), 256))
			strToCharSlice("extension A", extensionName)

			return core1_0.VKSuccess, nil
		})

	extensions, _, err := loader.AvailableExtensions()
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

func TestVulkanLoader1_0_AvailableExtensions_Incomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockDriver.EXPECT().VkEnumerateInstanceExtensionProperties(nil, gomock.Not(nil), nil).DoAndReturn(
		func(pLayerName *driver.Char, pPropertyCount *driver.Uint32, pProperties *driver.VkExtensionProperties) (common.VkResult, error) {
			*pPropertyCount = 2

			return core1_0.VKSuccess, nil
		})
	mockDriver.EXPECT().VkEnumerateInstanceExtensionProperties(nil, gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(pLayerName *driver.Char, pPropertyCount *driver.Uint32, pProperties *driver.VkExtensionProperties) (common.VkResult, error) {
			return core1_0.VKIncomplete, nil
		})
	mockDriver.EXPECT().VkEnumerateInstanceExtensionProperties(nil, gomock.Not(nil), nil).DoAndReturn(
		func(pLayerName *driver.Char, pPropertyCount *driver.Uint32, pProperties *driver.VkExtensionProperties) (common.VkResult, error) {
			*pPropertyCount = 2

			return core1_0.VKSuccess, nil
		})
	mockDriver.EXPECT().VkEnumerateInstanceExtensionProperties(nil, gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(pLayerName *driver.Char, pPropertyCount *driver.Uint32, pProperties *driver.VkExtensionProperties) (common.VkResult, error) {
			*pPropertyCount = 2
			propertySlice := reflect.ValueOf(([]driver.VkExtensionProperties)(unsafe.Slice(pProperties, 2)))

			extension := propertySlice.Index(0)

			*(*uint32)(unsafe.Pointer(extension.FieldByName("specVersion").UnsafeAddr())) = uint32(123)
			extensionName := ([]driver.Char)(unsafe.Slice((*driver.Char)(unsafe.Pointer(extension.FieldByName("extensionName").UnsafeAddr())), 256))
			strToCharSlice("extension 1", extensionName)

			extension = propertySlice.Index(1)
			*(*uint32)(unsafe.Pointer(extension.FieldByName("specVersion").UnsafeAddr())) = uint32(321)
			extensionName = ([]driver.Char)(unsafe.Slice((*driver.Char)(unsafe.Pointer(extension.FieldByName("extensionName").UnsafeAddr())), 256))
			strToCharSlice("extension A", extensionName)

			return core1_0.VKSuccess, nil
		})

	extensions, _, err := loader.AvailableExtensions()
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

func TestVulkanLoader1_0_AvailableLayers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockDriver.EXPECT().VkEnumerateInstanceLayerProperties(gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(pPropertyCount *driver.Uint32, pProperties *driver.VkLayerProperties) (common.VkResult, error) {
			*pPropertyCount = 2

			return core1_0.VKSuccess, nil
		})
	mockDriver.EXPECT().VkEnumerateInstanceLayerProperties(gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(pPropertyCount *driver.Uint32, pProperties *driver.VkLayerProperties) (common.VkResult, error) {
			*pPropertyCount = 2
			propertySlice := reflect.ValueOf(([]driver.VkLayerProperties)(unsafe.Slice(pProperties, 2)))

			layer := propertySlice.Index(0)

			*(*uint32)(unsafe.Pointer(layer.FieldByName("specVersion").UnsafeAddr())) = uint32(common.CreateVersion(1, 2, 3))
			*(*uint32)(unsafe.Pointer(layer.FieldByName("implementationVersion").UnsafeAddr())) = uint32(common.CreateVersion(2, 1, 3))
			layerName := ([]driver.Char)(unsafe.Slice((*driver.Char)(unsafe.Pointer(layer.FieldByName("layerName").UnsafeAddr())), 256))
			strToCharSlice("layer 1", layerName)
			layerDesc := ([]driver.Char)(unsafe.Slice((*driver.Char)(unsafe.Pointer(layer.FieldByName("description").UnsafeAddr())), 256))
			strToCharSlice("a cool layer", layerDesc)

			layer = propertySlice.Index(1)

			*(*uint32)(unsafe.Pointer(layer.FieldByName("specVersion").UnsafeAddr())) = uint32(common.CreateVersion(3, 2, 1))
			*(*uint32)(unsafe.Pointer(layer.FieldByName("implementationVersion").UnsafeAddr())) = uint32(common.CreateVersion(2, 3, 1))
			layerName = ([]driver.Char)(unsafe.Slice((*driver.Char)(unsafe.Pointer(layer.FieldByName("layerName").UnsafeAddr())), 256))
			strToCharSlice("layer A", layerName)
			layerDesc = ([]driver.Char)(unsafe.Slice((*driver.Char)(unsafe.Pointer(layer.FieldByName("description").UnsafeAddr())), 256))
			strToCharSlice("a bad layer", layerDesc)

			return core1_0.VKSuccess, nil
		})

	layers, _, err := loader.AvailableLayers()
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

func TestVulkanLoader1_0_AvailableLayers_Incomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockDriver.EXPECT().VkEnumerateInstanceLayerProperties(gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(pPropertyCount *driver.Uint32, pProperties *driver.VkLayerProperties) (common.VkResult, error) {
			*pPropertyCount = 2

			return core1_0.VKSuccess, nil
		})
	mockDriver.EXPECT().VkEnumerateInstanceLayerProperties(gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(pPropertyCount *driver.Uint32, pProperties *driver.VkLayerProperties) (common.VkResult, error) {
			return core1_0.VKIncomplete, nil
		})
	mockDriver.EXPECT().VkEnumerateInstanceLayerProperties(gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(pPropertyCount *driver.Uint32, pProperties *driver.VkLayerProperties) (common.VkResult, error) {
			*pPropertyCount = 2

			return core1_0.VKSuccess, nil
		})
	mockDriver.EXPECT().VkEnumerateInstanceLayerProperties(gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(pPropertyCount *driver.Uint32, pProperties *driver.VkLayerProperties) (common.VkResult, error) {
			*pPropertyCount = 2
			propertySlice := reflect.ValueOf(([]driver.VkLayerProperties)(unsafe.Slice(pProperties, 2)))

			layer := propertySlice.Index(0)

			*(*uint32)(unsafe.Pointer(layer.FieldByName("specVersion").UnsafeAddr())) = uint32(common.CreateVersion(1, 2, 3))
			*(*uint32)(unsafe.Pointer(layer.FieldByName("implementationVersion").UnsafeAddr())) = uint32(common.CreateVersion(2, 1, 3))
			layerName := ([]driver.Char)(unsafe.Slice((*driver.Char)(unsafe.Pointer(layer.FieldByName("layerName").UnsafeAddr())), 256))
			strToCharSlice("layer 1", layerName)
			layerDesc := ([]driver.Char)(unsafe.Slice((*driver.Char)(unsafe.Pointer(layer.FieldByName("description").UnsafeAddr())), 256))
			strToCharSlice("a cool layer", layerDesc)

			layer = propertySlice.Index(1)

			*(*uint32)(unsafe.Pointer(layer.FieldByName("specVersion").UnsafeAddr())) = uint32(common.CreateVersion(3, 2, 1))
			*(*uint32)(unsafe.Pointer(layer.FieldByName("implementationVersion").UnsafeAddr())) = uint32(common.CreateVersion(2, 3, 1))
			layerName = ([]driver.Char)(unsafe.Slice((*driver.Char)(unsafe.Pointer(layer.FieldByName("layerName").UnsafeAddr())), 256))
			strToCharSlice("layer A", layerName)
			layerDesc = ([]driver.Char)(unsafe.Slice((*driver.Char)(unsafe.Pointer(layer.FieldByName("description").UnsafeAddr())), 256))
			strToCharSlice("a bad layer", layerDesc)

			return core1_0.VKSuccess, nil
		})

	layers, _, err := loader.AvailableLayers()
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
