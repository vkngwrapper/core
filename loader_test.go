package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/cockroachdb/errors"
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

func TestVulkanLoader1_0_CreateInstance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	instanceHandle := mocks.NewFakeInstanceHandle()

	driver.EXPECT().CreateInstanceDriver(gomock.Any()).Return(driver, nil)
	driver.EXPECT().VkCreateInstance(gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(pCreateInfo *core.VkInstanceCreateInfo, pAllocator *core.VkAllocationCallbacks, pInstance *core.VkInstance) (core.VkResult, error) {
			*pInstance = instanceHandle

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(1), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), val.FieldByName("enabledExtensionCount").Uint())
			require.Equal(t, uint64(2), val.FieldByName("enabledLayerCount").Uint())

			layerNames := ([]*core.Char)(unsafe.Slice((**core.Char)(unsafe.Pointer(val.FieldByName("ppEnabledLayerNames").Elem().UnsafeAddr())), 2))
			layerNameSlice := ([]core.Char)(unsafe.Slice(layerNames[0], 256))
			layerNameBytes := []byte("layer a")
			for idx, b := range layerNameBytes {
				require.Equal(t, core.Char(b), layerNameSlice[idx], "mismatch at idx %d of %s", idx, string(layerNameBytes))
			}
			require.Equal(t, core.Char(0), layerNameSlice[len(layerNameBytes)])

			layerNameSlice = ([]core.Char)(unsafe.Slice(layerNames[1], 256))
			layerNameBytes = []byte("layer 2")
			for idx, b := range layerNameBytes {
				require.Equal(t, core.Char(b), layerNameSlice[idx])
			}
			require.Equal(t, core.Char(0), layerNameSlice[len(layerNameBytes)])

			extensionNames := ([]*core.Char)(unsafe.Slice((**core.Char)(unsafe.Pointer(val.FieldByName("ppEnabledExtensionNames").Elem().UnsafeAddr())), 1))
			extensionNameSlice := ([]core.Char)(unsafe.Slice(extensionNames[0], 256))
			extensionNameBytes := []byte("extension")
			for idx, b := range extensionNameBytes {
				require.Equal(t, core.Char(b), extensionNameSlice[idx])
			}
			require.Equal(t, core.Char(0), extensionNameSlice[len(extensionNameBytes)])

			require.False(t, val.FieldByName("pApplicationInfo").IsNil())
			appInfo := val.FieldByName("pApplicationInfo").Elem()
			require.Equal(t, uint64(0), appInfo.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_APPLICATION_INFO
			require.True(t, appInfo.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(common.CreateVersion(2, 3, 4)), appInfo.FieldByName("applicationVersion").Uint())
			require.Equal(t, uint64(common.CreateVersion(3, 4, 5)), appInfo.FieldByName("engineVersion").Uint())
			require.Equal(t, uint64(1<<22), appInfo.FieldByName("apiVersion").Uint()) // VK_API_VERSION_1_0

			applicationNamePtr := (*core.Char)(unsafe.Pointer(appInfo.FieldByName("pApplicationName").Elem().UnsafeAddr()))
			engineNamePtr := (*core.Char)(unsafe.Pointer(appInfo.FieldByName("pEngineName").Elem().UnsafeAddr()))

			applicationNameSlice := ([]core.Char)(unsafe.Slice(applicationNamePtr, 256))
			applicationNameBytes := []byte("test app")
			for idx, b := range applicationNameBytes {
				require.Equal(t, core.Char(b), applicationNameSlice[idx])
			}
			require.Equal(t, core.Char(0), applicationNameSlice[len(applicationNameBytes)])

			engineNameSlice := ([]core.Char)(unsafe.Slice(engineNamePtr, 256))
			engineNameBytes := []byte("test engine")
			for idx, b := range engineNameBytes {
				require.Equal(t, core.Char(b), engineNameSlice[idx])
			}
			require.Equal(t, core.Char(0), engineNameSlice[len(engineNameBytes)])

			return core.VKSuccess, nil
		})

	instance, _, err := loader.CreateInstance(&core.InstanceOptions{
		ApplicationName:    "test app",
		ApplicationVersion: common.CreateVersion(2, 3, 4),
		EngineName:         "test engine",
		EngineVersion:      common.CreateVersion(3, 4, 5),
		VulkanVersion:      common.Vulkan1_0,
		ExtensionNames:     []string{"extension"},
		LayerNames:         []string{"layer a", "layer 2"},
	})
	require.NoError(t, err)
	require.NotNil(t, instance)
	require.Equal(t, instanceHandle, instance.Handle())
}

func TestVulkanLoader1_0_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockPhysicalDevice := mocks.EasyMockPhysicalDevice(ctrl, mockDriver)
	deviceHandle := mocks.NewFakeDeviceHandle()

	mockDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(mockDriver, nil)
	mockDriver.EXPECT().VkCreateDevice(mockPhysicalDevice.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(physicalDevice core.VkPhysicalDevice, pCreateInfo *core.VkDeviceCreateInfo, pAllocator *core.VkAllocationCallbacks, pDevice *core.VkDevice) (core.VkResult, error) {
			v := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(3), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), v.FieldByName("flags").Uint())
			require.Equal(t, uint64(2), v.FieldByName("queueCreateInfoCount").Uint())
			require.Equal(t, uint64(3), v.FieldByName("enabledExtensionCount").Uint())
			require.Equal(t, uint64(2), v.FieldByName("enabledLayerCount").Uint())

			featuresV := v.FieldByName("pEnabledFeatures").Elem()

			require.Equal(t, uint64(1), featuresV.FieldByName("occlusionQueryPrecise").Uint())
			require.Equal(t, uint64(1), featuresV.FieldByName("tessellationShader").Uint())
			require.Equal(t, uint64(0), featuresV.FieldByName("samplerAnisotropy").Uint())

			extensionNamePtr := (**core.Char)(unsafe.Pointer(v.FieldByName("ppEnabledExtensionNames").Elem().UnsafeAddr()))
			extensionNameSlice := ([]*core.Char)(unsafe.Slice(extensionNamePtr, 3))

			var extensionNames []string
			for _, extensionNameBytes := range extensionNameSlice {
				var extensionNameRunes []rune
				extensionNameByteSlice := ([]core.Char)(unsafe.Slice(extensionNameBytes, 1<<30))
				for _, nameByte := range extensionNameByteSlice {
					if nameByte == 0 {
						break
					}

					extensionNameRunes = append(extensionNameRunes, rune(nameByte))
				}

				extensionNames = append(extensionNames, string(extensionNameRunes))
			}

			require.ElementsMatch(t, []string{"A", "B", "C"}, extensionNames)

			layerNamePtr := (**core.Char)(unsafe.Pointer(v.FieldByName("ppEnabledLayerNames").Elem().UnsafeAddr()))
			layerNameSlice := ([]*core.Char)(unsafe.Slice(layerNamePtr, 2))

			var layerNames []string
			for _, layerNameBytes := range layerNameSlice {
				var layerNameRunes []rune
				layerNameByteSlice := ([]core.Char)(unsafe.Slice(layerNameBytes, 1<<30))
				for _, nameByte := range layerNameByteSlice {
					if nameByte == 0 {
						break
					}

					layerNameRunes = append(layerNameRunes, rune(nameByte))
				}

				layerNames = append(layerNames, string(layerNameRunes))
			}

			require.ElementsMatch(t, []string{"D", "E"}, layerNames)

			queueCreateInfoPtr := (*core.VkDeviceQueueCreateInfo)(unsafe.Pointer(v.FieldByName("pQueueCreateInfos").Elem().UnsafeAddr()))
			queueCreateInfoSlice := ([]core.VkDeviceQueueCreateInfo)(unsafe.Slice(queueCreateInfoPtr, 2))

			queueInfoV := reflect.ValueOf(queueCreateInfoSlice[0])
			require.Equal(t, uint64(2), queueInfoV.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
			require.True(t, queueInfoV.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), queueInfoV.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), queueInfoV.FieldByName("queueFamilyIndex").Uint())
			require.Equal(t, uint64(3), queueInfoV.FieldByName("queueCount").Uint())

			priorityPtr := (*float32)(unsafe.Pointer(queueInfoV.FieldByName("pQueuePriorities").Elem().UnsafeAddr()))
			prioritySlice := ([]float32)(unsafe.Slice(priorityPtr, 3))
			require.Equal(t, float32(1.0), prioritySlice[0])
			require.Equal(t, float32(0.0), prioritySlice[1])
			require.Equal(t, float32(0.5), prioritySlice[2])

			queueInfoV = reflect.ValueOf(queueCreateInfoSlice[1])
			require.Equal(t, uint64(2), queueInfoV.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
			require.True(t, queueInfoV.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), queueInfoV.FieldByName("flags").Uint())
			require.Equal(t, uint64(3), queueInfoV.FieldByName("queueFamilyIndex").Uint())
			require.Equal(t, uint64(1), queueInfoV.FieldByName("queueCount").Uint())

			priorityPtr = (*float32)(unsafe.Pointer(queueInfoV.FieldByName("pQueuePriorities").Elem().UnsafeAddr()))
			prioritySlice = ([]float32)(unsafe.Slice(priorityPtr, 1))
			require.Equal(t, float32(0.5), prioritySlice[0])

			*pDevice = deviceHandle
			return core.VKSuccess, nil
		})

	device, _, err := loader.CreateDevice(mockPhysicalDevice, &core.DeviceOptions{
		QueueFamilies: []*core.QueueFamilyOptions{
			{
				QueueFamilyIndex: 1,
				QueuePriorities:  []float32{1, 0, 0.5},
			},
			{
				QueueFamilyIndex: 3,
				QueuePriorities:  []float32{0.5},
			},
		},
		ExtensionNames: []string{"A", "B", "C"},
		LayerNames:     []string{"D", "E"},
		EnabledFeatures: &common.PhysicalDeviceFeatures{
			OcclusionQueryPrecise: true,
			TessellationShader:    true,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, device)
	require.Equal(t, deviceHandle, device.Handle())
}

func TestVulkanLoader1_0_Create_FailNoQueueFamilies(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	mockPhysicalDevice := mocks.EasyMockPhysicalDevice(ctrl, mockDriver)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	_, _, err = loader.CreateDevice(mockPhysicalDevice, &core.DeviceOptions{
		QueueFamilies:  []*core.QueueFamilyOptions{},
		ExtensionNames: []string{"A", "B", "C"},
		LayerNames:     []string{"D", "E"},
		EnabledFeatures: &common.PhysicalDeviceFeatures{
			OcclusionQueryPrecise: true,
			TessellationShader:    true,
		},
	})
	require.Error(t, err, errors.New("alloc DeviceOptions: no queue families added"))
}

func TestVulkanLoader1_0_Create_FailFamilyWithoutPriorities(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockPhysicalDevice := mocks.EasyMockPhysicalDevice(ctrl, mockDriver)

	_, _, err = loader.CreateDevice(mockPhysicalDevice, &core.DeviceOptions{
		QueueFamilies: []*core.QueueFamilyOptions{
			{
				QueueFamilyIndex: 1,
				QueuePriorities:  []float32{1, 0, 0.5},
			},
			{
				QueueFamilyIndex: 3,
				QueuePriorities:  []float32{},
			},
		},
		ExtensionNames: []string{"A", "B", "C"},
		LayerNames:     []string{"D", "E"},
		EnabledFeatures: &common.PhysicalDeviceFeatures{
			OcclusionQueryPrecise: true,
			TessellationShader:    true,
		},
	})
	require.Error(t, errors.New("alloc DeviceOptions: queue family 1 had no queue priorities"))
}
