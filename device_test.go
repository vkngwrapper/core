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

func TestDevice_Create_Success(t *testing.T) {
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

func TestDevice_Create_FailNoQueueFamilies(t *testing.T) {
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

func TestDevice_Create_FailFamilyWithoutPriorities(t *testing.T) {
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
