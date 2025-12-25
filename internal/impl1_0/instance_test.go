package impl1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreateInstance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalGlobalDriver(mockLoader)

	instanceHandle := mocks.NewFakeInstanceHandle()

	mockLoader.EXPECT().VkCreateInstance(gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(pCreateInfo *loader.VkInstanceCreateInfo, pAllocator *loader.VkAllocationCallbacks, pInstance *loader.VkInstance) (common.VkResult, error) {
			*pInstance = instanceHandle

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(1), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), val.FieldByName("enabledExtensionCount").Uint())
			require.Equal(t, uint64(2), val.FieldByName("enabledLayerCount").Uint())

			layerNames := ([]*loader.Char)(unsafe.Slice((**loader.Char)(unsafe.Pointer(val.FieldByName("ppEnabledLayerNames").Elem().UnsafeAddr())), 2))
			layerNameSlice := ([]loader.Char)(unsafe.Slice(layerNames[0], 256))
			layerNameBytes := []byte("layer a")
			for idx, b := range layerNameBytes {
				require.Equal(t, loader.Char(b), layerNameSlice[idx], "mismatch at idx %d of %s", idx, string(layerNameBytes))
			}
			require.Equal(t, loader.Char(0), layerNameSlice[len(layerNameBytes)])

			layerNameSlice = ([]loader.Char)(unsafe.Slice(layerNames[1], 256))
			layerNameBytes = []byte("layer 2")
			for idx, b := range layerNameBytes {
				require.Equal(t, loader.Char(b), layerNameSlice[idx])
			}
			require.Equal(t, loader.Char(0), layerNameSlice[len(layerNameBytes)])

			extensionNames := ([]*loader.Char)(unsafe.Slice((**loader.Char)(unsafe.Pointer(val.FieldByName("ppEnabledExtensionNames").Elem().UnsafeAddr())), 1))
			extensionNameSlice := ([]loader.Char)(unsafe.Slice(extensionNames[0], 256))
			extensionNameBytes := []byte("extension")
			for idx, b := range extensionNameBytes {
				require.Equal(t, loader.Char(b), extensionNameSlice[idx])
			}
			require.Equal(t, loader.Char(0), extensionNameSlice[len(extensionNameBytes)])

			require.False(t, val.FieldByName("pApplicationInfo").IsNil())
			appInfo := val.FieldByName("pApplicationInfo").Elem()
			require.Equal(t, uint64(0), appInfo.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_APPLICATION_INFO
			require.True(t, appInfo.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(common.CreateVersion(2, 3, 4)), appInfo.FieldByName("applicationVersion").Uint())
			require.Equal(t, uint64(common.CreateVersion(3, 4, 5)), appInfo.FieldByName("engineVersion").Uint())
			require.Equal(t, uint64(1<<22), appInfo.FieldByName("apiVersion").Uint()) // VK_API_VERSION_1_0

			applicationNamePtr := (*loader.Char)(unsafe.Pointer(appInfo.FieldByName("pApplicationName").Elem().UnsafeAddr()))
			engineNamePtr := (*loader.Char)(unsafe.Pointer(appInfo.FieldByName("pEngineName").Elem().UnsafeAddr()))

			applicationNameSlice := ([]loader.Char)(unsafe.Slice(applicationNamePtr, 256))
			applicationNameBytes := []byte("test app")
			for idx, b := range applicationNameBytes {
				require.Equal(t, loader.Char(b), applicationNameSlice[idx])
			}
			require.Equal(t, loader.Char(0), applicationNameSlice[len(applicationNameBytes)])

			engineNameSlice := ([]loader.Char)(unsafe.Slice(engineNamePtr, 256))
			engineNameBytes := []byte("test engine")
			for idx, b := range engineNameBytes {
				require.Equal(t, loader.Char(b), engineNameSlice[idx])
			}
			require.Equal(t, loader.Char(0), engineNameSlice[len(engineNameBytes)])

			return core1_0.VKSuccess, nil
		})

	instance, _, err := driver.CreateInstance(nil, core1_0.InstanceCreateInfo{
		ApplicationName:       "test app",
		ApplicationVersion:    common.CreateVersion(2, 3, 4),
		EngineName:            "test engine",
		EngineVersion:         common.CreateVersion(3, 4, 5),
		APIVersion:            common.Vulkan1_0,
		EnabledExtensionNames: []string{"extension"},
		EnabledLayerNames:     []string{"layer a", "layer 2"},
	})
	require.NoError(t, err)
	require.NotNil(t, instance)
	require.Equal(t, instanceHandle, instance.Handle())
}

func TestVulkanInstance_PhysicalDevices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_0.InternalCoreInstanceDriver(mockLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	device1 := mocks.NewFakePhysicalDeviceHandle()
	device2 := mocks.NewFakePhysicalDeviceHandle()

	mockLoader.EXPECT().VkEnumeratePhysicalDevices(instance.Handle(), gomock.Not(nil), nil).DoAndReturn(
		func(instance loader.VkInstance, pPhysicalDeviceCount *loader.Uint32, pPhysicalDevices *loader.VkPhysicalDevice) (common.VkResult, error) {
			*pPhysicalDeviceCount = 2

			return core1_0.VKSuccess, nil
		})
	mockLoader.EXPECT().VkEnumeratePhysicalDevices(instance.Handle(), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(instance loader.VkInstance, pPhysicalDeviceCount *loader.Uint32, pPhysicalDevices *loader.VkPhysicalDevice) (common.VkResult, error) {
			*pPhysicalDeviceCount = 2

			deviceSlice := ([]loader.VkPhysicalDevice)(unsafe.Slice(pPhysicalDevices, 2))
			deviceSlice[0] = device1
			deviceSlice[1] = device2

			return core1_0.VKSuccess, nil
		})
	mockLoader.EXPECT().VkGetPhysicalDeviceProperties(device1, gomock.Not(nil)).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			val := reflect.ValueOf(pProperties).Elem()

			*(*uint32)(unsafe.Pointer(val.FieldByName("apiVersion").UnsafeAddr())) = uint32(1 << 22)
		})

	mockLoader.EXPECT().VkGetPhysicalDeviceProperties(device2, gomock.Not(nil)).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			val := reflect.ValueOf(pProperties).Elem()

			*(*uint32)(unsafe.Pointer(val.FieldByName("apiVersion").UnsafeAddr())) = uint32(1<<22 | 2<<12)
		})

	devices, _, err := driver.EnumeratePhysicalDevices(instance)
	require.NoError(t, err)
	require.Len(t, devices, 2)
	require.Equal(t, device1, devices[0].Handle())
	require.Equal(t, common.Vulkan1_0, devices[0].DeviceAPIVersion())
	require.Equal(t, device2, devices[1].Handle())
	require.Equal(t, common.Vulkan1_2, devices[1].DeviceAPIVersion())
}
