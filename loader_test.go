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

func strToCharSlice(text string, slice []core.Char) {
	byteSlice := []byte(text)
	for idx, b := range byteSlice {
		slice[idx] = core.Char(b)
	}
	slice[len(byteSlice)] = core.Char(0)
}

func TestVulkanLoader1_0_AvailableExtensions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	driver.EXPECT().VkEnumerateInstanceExtensionProperties(nil, gomock.Not(nil), nil).DoAndReturn(
		func(pLayerName *core.Char, pPropertyCount *core.Uint32, pProperties *core.VkExtensionProperties) (core.VkResult, error) {
			*pPropertyCount = 2

			return core.VKSuccess, nil
		})

	driver.EXPECT().VkEnumerateInstanceExtensionProperties(nil, gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(pLayerName *core.Char, pPropertyCount *core.Uint32, pProperties *core.VkExtensionProperties) (core.VkResult, error) {
			*pPropertyCount = 2
			propertySlice := reflect.ValueOf(([]core.VkExtensionProperties)(unsafe.Slice(pProperties, 2)))

			extension := propertySlice.Index(0)

			*(*uint32)(unsafe.Pointer(extension.FieldByName("specVersion").UnsafeAddr())) = uint32(common.CreateVersion(1, 2, 3))
			extensionName := ([]core.Char)(unsafe.Slice((*core.Char)(unsafe.Pointer(extension.FieldByName("extensionName").UnsafeAddr())), 256))
			strToCharSlice("extension 1", extensionName)

			extension = propertySlice.Index(1)
			*(*uint32)(unsafe.Pointer(extension.FieldByName("specVersion").UnsafeAddr())) = uint32(common.CreateVersion(3, 2, 1))
			extensionName = ([]core.Char)(unsafe.Slice((*core.Char)(unsafe.Pointer(extension.FieldByName("extensionName").UnsafeAddr())), 256))
			strToCharSlice("extension A", extensionName)

			return core.VKSuccess, nil
		})

	extensions, _, err := loader.AvailableExtensions()
	require.NoError(t, err)
	require.Len(t, extensions, 2)

	extension := extensions["extension 1"]
	require.NotNil(t, extension)
	require.Equal(t, "extension 1", extension.ExtensionName)
	require.Equal(t, common.CreateVersion(1, 2, 3), extension.SpecVersion)

	extension = extensions["extension A"]
	require.NotNil(t, extension)
	require.Equal(t, "extension A", extension.ExtensionName)
	require.Equal(t, common.CreateVersion(3, 2, 1), extension.SpecVersion)
}

func TestVulkanLoader1_0_AvailableLayers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	driver.EXPECT().VkEnumerateInstanceLayerProperties(gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(pPropertyCount *core.Uint32, pProperties *core.VkLayerProperties) (core.VkResult, error) {
			*pPropertyCount = 2

			return core.VKSuccess, nil
		})
	driver.EXPECT().VkEnumerateInstanceLayerProperties(gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(pPropertyCount *core.Uint32, pProperties *core.VkLayerProperties) (core.VkResult, error) {
			*pPropertyCount = 2
			propertySlice := reflect.ValueOf(([]core.VkLayerProperties)(unsafe.Slice(pProperties, 2)))

			layer := propertySlice.Index(0)

			*(*uint32)(unsafe.Pointer(layer.FieldByName("specVersion").UnsafeAddr())) = uint32(common.CreateVersion(1, 2, 3))
			*(*uint32)(unsafe.Pointer(layer.FieldByName("implementationVersion").UnsafeAddr())) = uint32(common.CreateVersion(2, 1, 3))
			layerName := ([]core.Char)(unsafe.Slice((*core.Char)(unsafe.Pointer(layer.FieldByName("layerName").UnsafeAddr())), 256))
			strToCharSlice("layer 1", layerName)
			layerDesc := ([]core.Char)(unsafe.Slice((*core.Char)(unsafe.Pointer(layer.FieldByName("description").UnsafeAddr())), 256))
			strToCharSlice("a cool layer", layerDesc)

			layer = propertySlice.Index(1)

			*(*uint32)(unsafe.Pointer(layer.FieldByName("specVersion").UnsafeAddr())) = uint32(common.CreateVersion(3, 2, 1))
			*(*uint32)(unsafe.Pointer(layer.FieldByName("implementationVersion").UnsafeAddr())) = uint32(common.CreateVersion(2, 3, 1))
			layerName = ([]core.Char)(unsafe.Slice((*core.Char)(unsafe.Pointer(layer.FieldByName("layerName").UnsafeAddr())), 256))
			strToCharSlice("layer A", layerName)
			layerDesc = ([]core.Char)(unsafe.Slice((*core.Char)(unsafe.Pointer(layer.FieldByName("description").UnsafeAddr())), 256))
			strToCharSlice("a bad layer", layerDesc)

			return core.VKSuccess, nil
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
