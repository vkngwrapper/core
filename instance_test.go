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

func TestVulkanInstance_PhysicalDevices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	instance := mocks.EasyDummyInstance(t, loader)
	device1 := mocks.NewFakePhysicalDeviceHandle()
	device2 := mocks.NewFakePhysicalDeviceHandle()

	driver.EXPECT().VkEnumeratePhysicalDevices(instance.Handle(), gomock.Not(nil), nil).DoAndReturn(
		func(instance core.VkInstance, pPhysicalDeviceCount *core.Uint32, pPhysicalDevices *core.VkPhysicalDevice) (core.VkResult, error) {
			*pPhysicalDeviceCount = 2

			return core.VKSuccess, nil
		})
	driver.EXPECT().VkEnumeratePhysicalDevices(instance.Handle(), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(instance core.VkInstance, pPhysicalDeviceCount *core.Uint32, pPhysicalDevices *core.VkPhysicalDevice) (core.VkResult, error) {
			*pPhysicalDeviceCount = 2

			deviceSlice := ([]core.VkPhysicalDevice)(unsafe.Slice(pPhysicalDevices, 2))
			deviceSlice[0] = device1
			deviceSlice[1] = device2

			return core.VKSuccess, nil
		})

	devices, _, err := instance.PhysicalDevices()
	require.NoError(t, err)
	require.Len(t, devices, 2)
	require.Equal(t, device1, devices[0].Handle())
	require.Equal(t, device2, devices[1].Handle())
}
