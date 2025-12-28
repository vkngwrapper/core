package core1_2_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_2"
	"go.uber.org/mock/gomock"
)

func TestPhysicalDevice8BitStorageFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_2, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice,
			pCreateInfo *loader.VkDeviceCreateInfo,
			pAllocator *loader.VkAllocationCallbacks,
			pDevice *loader.VkDevice) (common.VkResult, error) {

			*pDevice = mockDevice.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

			next := (*loader.VkPhysicalDevice8BitStorageFeatures)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000177000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("storageBuffer8BitAccess").Uint())
			require.Equal(t, uint64(0), val.FieldByName("uniformAndStorageBuffer8BitAccess").Uint())
			require.Equal(t, uint64(1), val.FieldByName("storagePushConstant8").Uint())

			return core1_0.VKSuccess, nil
		})

	device, _, err := driver.CreateDevice(
		physicalDevice,
		nil,
		core1_0.DeviceCreateInfo{
			QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
				{
					QueuePriorities: []float32{0},
				},
			},
			NextOptions: common.NextOptions{core1_2.PhysicalDevice8BitStorageFeatures{
				StoragePushConstant8:              true,
				UniformAndStorageBuffer8BitAccess: false,
				StorageBuffer8BitAccess:           true,
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDevice8BitStorageFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *loader.VkPhysicalDeviceFeatures2) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2

		next := (*loader.VkPhysicalDevice8BitStorageFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000177000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("storageBuffer8BitAccess").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("uniformAndStorageBuffer8BitAccess").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("storagePushConstant8").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData core1_2.PhysicalDevice8BitStorageFeatures
	err := driver.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&core1_1.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDevice8BitStorageFeatures{
		StorageBuffer8BitAccess:           true,
		UniformAndStorageBuffer8BitAccess: false,
		StoragePushConstant8:              true,
	}, outData)
}

func TestPhysicalDeviceBufferAddressFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	mockDevice := mocks.NewDummyDevice(common.Vulkan1_2, []string{})

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pCreateInfo *loader.VkDeviceCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDevice *loader.VkDevice) (common.VkResult, error) {

		*pDevice = mockDevice.Handle()
		val := reflect.ValueOf(pCreateInfo).Elem()

		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		next := (*loader.VkPhysicalDeviceBufferDeviceAddressFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000257000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("bufferDeviceAddress").Uint())
		require.Equal(t, uint64(0), val.FieldByName("bufferDeviceAddressCaptureReplay").Uint())
		require.Equal(t, uint64(1), val.FieldByName("bufferDeviceAddressMultiDevice").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := driver.CreateDevice(
		physicalDevice,
		nil,
		core1_0.DeviceCreateInfo{
			QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
				{
					QueuePriorities: []float32{0},
				},
			},

			NextOptions: common.NextOptions{core1_2.PhysicalDeviceBufferDeviceAddressFeatures{
				BufferDeviceAddress:            true,
				BufferDeviceAddressMultiDevice: true,
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceBufferAddressFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *loader.VkPhysicalDeviceFeatures2) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2

		next := (*loader.VkPhysicalDeviceBufferDeviceAddressFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000257000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("bufferDeviceAddress").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("bufferDeviceAddressCaptureReplay").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("bufferDeviceAddressMultiDevice").UnsafeAddr())) = loader.VkBool32(0)
	})

	var outData core1_2.PhysicalDeviceBufferDeviceAddressFeatures
	err := driver.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&core1_1.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceBufferDeviceAddressFeatures{
		BufferDeviceAddressCaptureReplay: true,
	}, outData)
}

func TestPhysicalDeviceDescriptorIndexingFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_2, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice,
			pCreateInfo *loader.VkDeviceCreateInfo,
			pAllocator *loader.VkAllocationCallbacks,
			pDevice *loader.VkDevice) (common.VkResult, error) {

			*pDevice = mockDevice.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

			next := (*loader.VkPhysicalDeviceDescriptorIndexingFeatures)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000161001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("shaderInputAttachmentArrayDynamicIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderUniformTexelBufferArrayDynamicIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderStorageTexelBufferArrayDynamicIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderUniformBufferArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderSampledImageArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderStorageBufferArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderStorageImageArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderInputAttachmentArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderUniformTexelBufferArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderStorageTexelBufferArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingUniformBufferUpdateAfterBind").Uint())
			require.Equal(t, uint64(0), val.FieldByName("descriptorBindingSampledImageUpdateAfterBind").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingStorageImageUpdateAfterBind").Uint())
			require.Equal(t, uint64(0), val.FieldByName("descriptorBindingStorageBufferUpdateAfterBind").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingUniformTexelBufferUpdateAfterBind").Uint())
			require.Equal(t, uint64(0), val.FieldByName("descriptorBindingStorageTexelBufferUpdateAfterBind").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingUpdateUnusedWhilePending").Uint())
			require.Equal(t, uint64(0), val.FieldByName("descriptorBindingPartiallyBound").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingVariableDescriptorCount").Uint())
			require.Equal(t, uint64(0), val.FieldByName("runtimeDescriptorArray").Uint())

			return core1_0.VKSuccess, nil
		})

	device, _, err := driver.CreateDevice(
		physicalDevice,
		nil,
		core1_0.DeviceCreateInfo{
			QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
				{
					QueuePriorities: []float32{0},
				},
			},
			NextOptions: common.NextOptions{core1_2.PhysicalDeviceDescriptorIndexingFeatures{
				ShaderInputAttachmentArrayDynamicIndexing:          true,
				ShaderUniformTexelBufferArrayDynamicIndexing:       false,
				ShaderStorageTexelBufferArrayDynamicIndexing:       true,
				ShaderUniformBufferArrayNonUniformIndexing:         false,
				ShaderSampledImageArrayNonUniformIndexing:          true,
				ShaderStorageBufferArrayNonUniformIndexing:         false,
				ShaderStorageImageArrayNonUniformIndexing:          true,
				ShaderInputAttachmentArrayNonUniformIndexing:       false,
				ShaderUniformTexelBufferArrayNonUniformIndexing:    true,
				ShaderStorageTexelBufferArrayNonUniformIndexing:    false,
				DescriptorBindingUniformBufferUpdateAfterBind:      true,
				DescriptorBindingSampledImageUpdateAfterBind:       false,
				DescriptorBindingStorageImageUpdateAfterBind:       true,
				DescriptorBindingStorageBufferUpdateAfterBind:      false,
				DescriptorBindingUniformTexelBufferUpdateAfterBind: true,
				DescriptorBindingStorageTexelBufferUpdateAfterBind: false,
				DescriptorBindingUpdateUnusedWhilePending:          true,
				DescriptorBindingPartiallyBound:                    false,
				DescriptorBindingVariableDescriptorCount:           true,
				RuntimeDescriptorArray:                             false,
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceDescriptorIndexingFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *loader.VkPhysicalDeviceFeatures2) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2

		next := (*loader.VkPhysicalDeviceDescriptorIndexingFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000161001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInputAttachmentArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformTexelBufferArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageTexelBufferArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformBufferArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSampledImageArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageBufferArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInputAttachmentArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformTexelBufferArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageTexelBufferArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingUniformBufferUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingSampledImageUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingStorageImageUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingStorageBufferUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingUniformTexelBufferUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingStorageTexelBufferUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingUpdateUnusedWhilePending").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingPartiallyBound").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingVariableDescriptorCount").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("runtimeDescriptorArray").UnsafeAddr())) = loader.VkBool32(0)
	})

	var outData core1_2.PhysicalDeviceDescriptorIndexingFeatures
	err := driver.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&core1_1.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceDescriptorIndexingFeatures{
		ShaderInputAttachmentArrayDynamicIndexing:          true,
		ShaderUniformTexelBufferArrayDynamicIndexing:       false,
		ShaderStorageTexelBufferArrayDynamicIndexing:       true,
		ShaderUniformBufferArrayNonUniformIndexing:         false,
		ShaderSampledImageArrayNonUniformIndexing:          true,
		ShaderStorageBufferArrayNonUniformIndexing:         false,
		ShaderStorageImageArrayNonUniformIndexing:          true,
		ShaderInputAttachmentArrayNonUniformIndexing:       false,
		ShaderUniformTexelBufferArrayNonUniformIndexing:    true,
		ShaderStorageTexelBufferArrayNonUniformIndexing:    false,
		DescriptorBindingUniformBufferUpdateAfterBind:      true,
		DescriptorBindingSampledImageUpdateAfterBind:       false,
		DescriptorBindingStorageImageUpdateAfterBind:       true,
		DescriptorBindingStorageBufferUpdateAfterBind:      false,
		DescriptorBindingUniformTexelBufferUpdateAfterBind: true,
		DescriptorBindingStorageTexelBufferUpdateAfterBind: false,
		DescriptorBindingUpdateUnusedWhilePending:          true,
		DescriptorBindingPartiallyBound:                    false,
		DescriptorBindingVariableDescriptorCount:           true,
		RuntimeDescriptorArray:                             false,
	}, outData)
}

func TestPhysicalDeviceHostQueryResetFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_2, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pCreateInfo *loader.VkDeviceCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDevice *loader.VkDevice) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		next := (*loader.VkPhysicalDeviceHostQueryResetFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000261000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_QUERY_RESET_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("hostQueryReset").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := driver.CreateDevice(
		physicalDevice,
		nil,
		core1_0.DeviceCreateInfo{
			QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
				{
					QueuePriorities: []float32{0},
				},
			},
			NextOptions: common.NextOptions{
				core1_2.PhysicalDeviceHostQueryResetFeatures{
					HostQueryReset: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceHostQueryResetFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *loader.VkPhysicalDeviceFeatures2) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2

		next := (*loader.VkPhysicalDeviceHostQueryResetFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000261000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_QUERY_RESET_FEATURES_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("hostQueryReset").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData core1_2.PhysicalDeviceHostQueryResetFeatures
	err := driver.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&core1_1.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceHostQueryResetFeatures{
		HostQueryReset: true,
	}, outData)
}

func TestPhysicalDeviceImagelessFramebufferFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_2, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pCreateInfo *loader.VkDeviceCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDevice *loader.VkDevice) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		next := (*loader.VkPhysicalDeviceImagelessFramebufferFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000108000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("imagelessFramebuffer").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := driver.CreateDevice(
		physicalDevice,
		nil,
		core1_0.DeviceCreateInfo{
			QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
				{
					QueuePriorities: []float32{0},
				},
			},
			NextOptions: common.NextOptions{
				core1_2.PhysicalDeviceImagelessFramebufferFeatures{
					ImagelessFramebuffer: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceImagelessFramebufferFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *loader.VkPhysicalDeviceFeatures2) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2

		next := (*loader.VkPhysicalDeviceImagelessFramebufferFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000108000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("imagelessFramebuffer").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData core1_2.PhysicalDeviceImagelessFramebufferFeatures
	err := driver.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&core1_1.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceImagelessFramebufferFeatures{
		ImagelessFramebuffer: true,
	}, outData)
}

func TestPhysicalDeviceScalarBlockLayoutFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_2, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pCreateInfo *loader.VkDeviceCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDevice *loader.VkDevice) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		next := (*loader.VkPhysicalDeviceScalarBlockLayoutFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000221000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SCALAR_BLOCK_LAYOUT_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("scalarBlockLayout").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := driver.CreateDevice(
		physicalDevice,
		nil,
		core1_0.DeviceCreateInfo{
			QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
				{
					QueuePriorities: []float32{0},
				},
			},
			NextOptions: common.NextOptions{
				core1_2.PhysicalDeviceScalarBlockLayoutFeatures{
					ScalarBlockLayout: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceScalarBlockLayoutFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *loader.VkPhysicalDeviceFeatures2) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2

		next := (*loader.VkPhysicalDeviceScalarBlockLayoutFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000221000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SCALAR_BLOCK_LAYOUT_FEATURES_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("scalarBlockLayout").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData core1_2.PhysicalDeviceScalarBlockLayoutFeatures
	err := driver.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&core1_1.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceScalarBlockLayoutFeatures{
		ScalarBlockLayout: true,
	}, outData)
}

func TestPhysicalDeviceSeparateDepthStencilLayoutsFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_2, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pCreateInfo *loader.VkDeviceCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDevice *loader.VkDevice) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		next := (*loader.VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000241000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SEPARATE_DEPTH_STENCIL_LAYOUTS_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("separateDepthStencilLayouts").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := driver.CreateDevice(
		physicalDevice,
		nil,
		core1_0.DeviceCreateInfo{
			QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
				{
					QueuePriorities: []float32{0},
				},
			},
			NextOptions: common.NextOptions{
				core1_2.PhysicalDeviceSeparateDepthStencilLayoutsFeatures{
					SeparateDepthStencilLayouts: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceSeparateDepthStencilLayoutsFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *loader.VkPhysicalDeviceFeatures2) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2

		next := (*loader.VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000241000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SEPARATE_DEPTH_STENCIL_LAYOUTS_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("separateDepthStencilLayouts").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData core1_2.PhysicalDeviceSeparateDepthStencilLayoutsFeatures
	err := driver.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&core1_1.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceSeparateDepthStencilLayoutsFeatures{
		SeparateDepthStencilLayouts: true,
	}, outData)
}

func TestPhysicalDeviceShaderAtomicInt64FeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_2, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pCreateInfo *loader.VkDeviceCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDevice *loader.VkDevice) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		next := (*loader.VkPhysicalDeviceShaderAtomicInt64Features)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000180000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_INT64_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("shaderBufferInt64Atomics").Uint())
		require.Equal(t, uint64(1), val.FieldByName("shaderSharedInt64Atomics").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := driver.CreateDevice(
		physicalDevice,
		nil,
		core1_0.DeviceCreateInfo{
			QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
				{
					QueuePriorities: []float32{0},
				},
			},
			NextOptions: common.NextOptions{
				core1_2.PhysicalDeviceShaderAtomicInt64Features{
					ShaderBufferInt64Atomics: true,
					ShaderSharedInt64Atomics: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceShaderAtomicInt64FeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *loader.VkPhysicalDeviceFeatures2) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2

		next := (*loader.VkPhysicalDeviceShaderAtomicInt64Features)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000180000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_INT64_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderBufferInt64Atomics").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSharedInt64Atomics").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData core1_2.PhysicalDeviceShaderAtomicInt64Features
	err := driver.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&core1_1.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceShaderAtomicInt64Features{
		ShaderBufferInt64Atomics: true,
		ShaderSharedInt64Atomics: true,
	}, outData)
}

func TestPhysicalDeviceShaderFloat16Int8FeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_2, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pCreateInfo *loader.VkDeviceCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDevice *loader.VkDevice) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		next := (*loader.VkPhysicalDeviceShaderFloat16Int8Features)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000082000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("shaderInt8").Uint())
		require.Equal(t, uint64(1), val.FieldByName("shaderFloat16").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := driver.CreateDevice(
		physicalDevice,
		nil,
		core1_0.DeviceCreateInfo{
			QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
				{
					QueuePriorities: []float32{0},
				},
			},
			NextOptions: common.NextOptions{
				core1_2.PhysicalDeviceShaderFloat16Int8Features{
					ShaderInt8:    true,
					ShaderFloat16: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceShaderFloat16Int8FeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *loader.VkPhysicalDeviceFeatures2) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

		next := (*loader.VkPhysicalDeviceShaderFloat16Int8Features)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000082000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInt8").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderFloat16").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData core1_2.PhysicalDeviceShaderFloat16Int8Features
	err := driver.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&core1_1.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceShaderFloat16Int8Features{
		ShaderFloat16: true,
		ShaderInt8:    true,
	}, outData)
}

func TestPhysicalDeviceShaderSubgroupExtendedTypesFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_2, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pCreateInfo *loader.VkDeviceCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDevice *loader.VkDevice) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		next := (*loader.VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000175000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_EXTENDED_TYPES_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("shaderSubgroupExtendedTypes").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := driver.CreateDevice(
		physicalDevice,
		nil,
		core1_0.DeviceCreateInfo{
			QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
				{
					QueuePriorities: []float32{0},
				},
			},
			NextOptions: common.NextOptions{
				core1_2.PhysicalDeviceShaderSubgroupExtendedTypesFeatures{
					ShaderSubgroupExtendedTypes: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceShaderSubgroupExtendedTypesFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *loader.VkPhysicalDeviceFeatures2) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2

		next := (*loader.VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000175000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_EXTENDED_TYPES_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSubgroupExtendedTypes").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData core1_2.PhysicalDeviceShaderSubgroupExtendedTypesFeatures
	err := driver.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&core1_1.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceShaderSubgroupExtendedTypesFeatures{
		ShaderSubgroupExtendedTypes: true,
	}, outData)
}

func TestPhysicalDeviceTimelineSemaphoreFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_2, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pCreateInfo *loader.VkDeviceCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDevice *loader.VkDevice) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		next := (*loader.VkPhysicalDeviceTimelineSemaphoreFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000207000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("timelineSemaphore").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := driver.CreateDevice(
		physicalDevice,
		nil,
		core1_0.DeviceCreateInfo{
			QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
				{
					QueuePriorities: []float32{0},
				},
			},
			NextOptions: common.NextOptions{
				core1_2.PhysicalDeviceTimelineSemaphoreFeatures{
					TimelineSemaphore: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceTimelineSemaphoreFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *loader.VkPhysicalDeviceFeatures2) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2

		next := (*loader.VkPhysicalDeviceTimelineSemaphoreFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000207000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("timelineSemaphore").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData core1_2.PhysicalDeviceTimelineSemaphoreFeatures
	err := driver.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&core1_1.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceTimelineSemaphoreFeatures{
		TimelineSemaphore: true,
	}, outData)
}

func TestPhysicalDeviceUniformBufferStandardLayoutFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_2, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pCreateInfo *loader.VkDeviceCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDevice *loader.VkDevice) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		next := (*loader.VkPhysicalDeviceUniformBufferStandardLayoutFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000253000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("uniformBufferStandardLayout").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := driver.CreateDevice(
		physicalDevice,
		nil,
		core1_0.DeviceCreateInfo{
			QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
				{
					QueuePriorities: []float32{0},
				},
			},
			NextOptions: common.NextOptions{
				core1_2.PhysicalDeviceUniformBufferStandardLayoutFeatures{
					UniformBufferStandardLayout: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceUniformBufferStandardLayoutFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *loader.VkPhysicalDeviceFeatures2) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2

		next := (*loader.VkPhysicalDeviceUniformBufferStandardLayoutFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000253000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("uniformBufferStandardLayout").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData core1_2.PhysicalDeviceUniformBufferStandardLayoutFeatures
	err := driver.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&core1_1.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceUniformBufferStandardLayoutFeatures{
		UniformBufferStandardLayout: true,
	}, outData)
}

func TestPhysicalDeviceVulkanMemoryModelFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_2, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice,
			pCreateInfo *loader.VkDeviceCreateInfo,
			pAllocator *loader.VkAllocationCallbacks,
			pDevice *loader.VkDevice) (common.VkResult, error) {

			*pDevice = mockDevice.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

			featuresPtr := (*loader.VkPhysicalDeviceVulkanMemoryModelFeatures)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(featuresPtr).Elem()

			require.Equal(t, uint64(1000211000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("vulkanMemoryModel").Uint())
			require.Equal(t, uint64(0), val.FieldByName("vulkanMemoryModelDeviceScope").Uint())
			require.Equal(t, uint64(1), val.FieldByName("vulkanMemoryModelAvailabilityVisibilityChains").Uint())

			return core1_0.VKSuccess, nil
		})

	device, _, err := driver.CreateDevice(physicalDevice, nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueuePriorities: []float32{0},
			},
		},

		NextOptions: common.NextOptions{Next: core1_2.PhysicalDeviceVulkanMemoryModelFeatures{
			VulkanMemoryModel:                             true,
			VulkanMemoryModelDeviceScope:                  false,
			VulkanMemoryModelAvailabilityVisibilityChains: true,
		}},
	})
	require.NoError(t, err)
	require.NotNil(t, device)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceVulkanMemoryModelFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(
			physicalDevice loader.VkPhysicalDevice,
			pFeatures *loader.VkPhysicalDeviceFeatures2,
		) {
			val := reflect.ValueOf(pFeatures).Elem()

			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

			outData := (*loader.VkPhysicalDeviceVulkanMemoryModelFeatures)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(outData).Elem()

			require.Equal(t, uint64(1000211000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())

			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("vulkanMemoryModel").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("vulkanMemoryModelDeviceScope").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("vulkanMemoryModelAvailabilityVisibilityChains").UnsafeAddr())) = loader.VkBool32(1)
		})

	var outData core1_2.PhysicalDeviceVulkanMemoryModelFeatures
	err := driver.GetPhysicalDeviceFeatures2(physicalDevice, &core1_1.PhysicalDeviceFeatures2{
		NextOutData: common.NextOutData{Next: &outData},
	})
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceVulkanMemoryModelFeatures{
		VulkanMemoryModel:                             true,
		VulkanMemoryModelDeviceScope:                  false,
		VulkanMemoryModelAvailabilityVisibilityChains: true,
	}, outData)
}

func TestPhysicalDeviceVulkan11FeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_2, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice,
			pCreateInfo *loader.VkDeviceCreateInfo,
			pAllocator *loader.VkAllocationCallbacks,
			pDevice *loader.VkDevice) (common.VkResult, error) {

			*pDevice = mockDevice.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

			featuresPtr := (*loader.VkPhysicalDeviceVulkan11Features)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(featuresPtr).Elem()

			require.Equal(t, uint64(49), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_1_FEATURES
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("storageBuffer16BitAccess").Uint())
			require.Equal(t, uint64(0), val.FieldByName("uniformAndStorageBuffer16BitAccess").Uint())
			require.Equal(t, uint64(1), val.FieldByName("storagePushConstant16").Uint())
			require.Equal(t, uint64(0), val.FieldByName("storageInputOutput16").Uint())
			require.Equal(t, uint64(1), val.FieldByName("multiview").Uint())
			require.Equal(t, uint64(0), val.FieldByName("multiviewGeometryShader").Uint())
			require.Equal(t, uint64(1), val.FieldByName("multiviewTessellationShader").Uint())
			require.Equal(t, uint64(0), val.FieldByName("variablePointersStorageBuffer").Uint())
			require.Equal(t, uint64(1), val.FieldByName("variablePointers").Uint())
			require.Equal(t, uint64(0), val.FieldByName("protectedMemory").Uint())
			require.Equal(t, uint64(1), val.FieldByName("samplerYcbcrConversion").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderDrawParameters").Uint())

			return core1_0.VKSuccess, nil
		})

	device, _, err := driver.CreateDevice(physicalDevice, nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueuePriorities: []float32{0},
			},
		},

		NextOptions: common.NextOptions{Next: core1_2.PhysicalDeviceVulkan11Features{
			StorageBuffer16BitAccess:    true,
			StoragePushConstant16:       true,
			Multiview:                   true,
			MultiviewTessellationShader: true,
			VariablePointers:            true,
			SamplerYcbcrConversion:      true,
		}},
	})
	require.NoError(t, err)
	require.NotNil(t, device)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceVulkan11FeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(
			physicalDevice loader.VkPhysicalDevice,
			pFeatures *loader.VkPhysicalDeviceFeatures2,
		) {
			val := reflect.ValueOf(pFeatures).Elem()

			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

			outData := (*loader.VkPhysicalDeviceVulkan11Features)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(outData).Elem()

			require.Equal(t, uint64(49), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_1_FEATURES
			require.True(t, val.FieldByName("pNext").IsNil())

			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("storageBuffer16BitAccess").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("uniformAndStorageBuffer16BitAccess").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("storagePushConstant16").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("storageInputOutput16").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("multiview").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("multiviewGeometryShader").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("multiviewTessellationShader").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("variablePointersStorageBuffer").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("variablePointers").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("protectedMemory").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("samplerYcbcrConversion").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDrawParameters").UnsafeAddr())) = loader.VkBool32(0)
		})

	var outData core1_2.PhysicalDeviceVulkan11Features
	err := driver.GetPhysicalDeviceFeatures2(physicalDevice, &core1_1.PhysicalDeviceFeatures2{
		NextOutData: common.NextOutData{Next: &outData},
	})
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceVulkan11Features{
		StorageBuffer16BitAccess:           true,
		UniformAndStorageBuffer16BitAccess: false,
		StoragePushConstant16:              true,
		StorageInputOutput16:               false,
		Multiview:                          true,
		MultiviewGeometryShader:            false,
		MultiviewTessellationShader:        true,
		VariablePointersStorageBuffer:      false,
		VariablePointers:                   true,
		ProtectedMemory:                    false,
		SamplerYcbcrConversion:             true,
		ShaderDrawParameters:               false,
	}, outData)
}

func TestPhysicalDeviceVulkan12FeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_2, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice,
			pCreateInfo *loader.VkDeviceCreateInfo,
			pAllocator *loader.VkAllocationCallbacks,
			pDevice *loader.VkDevice) (common.VkResult, error) {

			*pDevice = mockDevice.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

			featuresPtr := (*loader.VkPhysicalDeviceVulkan12Features)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(featuresPtr).Elem()

			require.Equal(t, uint64(51), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_2_FEATURES
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("samplerMirrorClampToEdge").Uint())
			require.Equal(t, uint64(0), val.FieldByName("drawIndirectCount").Uint())
			require.Equal(t, uint64(1), val.FieldByName("storageBuffer8BitAccess").Uint())
			require.Equal(t, uint64(0), val.FieldByName("uniformAndStorageBuffer8BitAccess").Uint())
			require.Equal(t, uint64(1), val.FieldByName("storagePushConstant8").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderBufferInt64Atomics").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderSharedInt64Atomics").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderFloat16").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderInt8").Uint())
			require.Equal(t, uint64(0), val.FieldByName("descriptorIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderInputAttachmentArrayDynamicIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderUniformTexelBufferArrayDynamicIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderStorageTexelBufferArrayDynamicIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderUniformBufferArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderSampledImageArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderStorageBufferArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderStorageImageArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderInputAttachmentArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderUniformTexelBufferArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderStorageTexelBufferArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingUniformBufferUpdateAfterBind").Uint())
			require.Equal(t, uint64(0), val.FieldByName("descriptorBindingSampledImageUpdateAfterBind").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingStorageImageUpdateAfterBind").Uint())
			require.Equal(t, uint64(0), val.FieldByName("descriptorBindingStorageBufferUpdateAfterBind").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingUniformTexelBufferUpdateAfterBind").Uint())
			require.Equal(t, uint64(0), val.FieldByName("descriptorBindingStorageTexelBufferUpdateAfterBind").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingUpdateUnusedWhilePending").Uint())
			require.Equal(t, uint64(0), val.FieldByName("descriptorBindingPartiallyBound").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingVariableDescriptorCount").Uint())
			require.Equal(t, uint64(0), val.FieldByName("runtimeDescriptorArray").Uint())
			require.Equal(t, uint64(1), val.FieldByName("samplerFilterMinmax").Uint())
			require.Equal(t, uint64(0), val.FieldByName("scalarBlockLayout").Uint())
			require.Equal(t, uint64(1), val.FieldByName("imagelessFramebuffer").Uint())
			require.Equal(t, uint64(0), val.FieldByName("uniformBufferStandardLayout").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderSubgroupExtendedTypes").Uint())
			require.Equal(t, uint64(0), val.FieldByName("separateDepthStencilLayouts").Uint())
			require.Equal(t, uint64(1), val.FieldByName("hostQueryReset").Uint())
			require.Equal(t, uint64(0), val.FieldByName("timelineSemaphore").Uint())
			require.Equal(t, uint64(1), val.FieldByName("bufferDeviceAddress").Uint())
			require.Equal(t, uint64(0), val.FieldByName("bufferDeviceAddressCaptureReplay").Uint())
			require.Equal(t, uint64(1), val.FieldByName("bufferDeviceAddressMultiDevice").Uint())
			require.Equal(t, uint64(0), val.FieldByName("vulkanMemoryModel").Uint())
			require.Equal(t, uint64(1), val.FieldByName("vulkanMemoryModelDeviceScope").Uint())
			require.Equal(t, uint64(0), val.FieldByName("vulkanMemoryModelAvailabilityVisibilityChains").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderOutputViewportIndex").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderOutputLayer").Uint())
			require.Equal(t, uint64(1), val.FieldByName("subgroupBroadcastDynamicId").Uint())

			return core1_0.VKSuccess, nil
		})

	device, _, err := driver.CreateDevice(physicalDevice, nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueuePriorities: []float32{0},
			},
		},

		NextOptions: common.NextOptions{Next: core1_2.PhysicalDeviceVulkan12Features{
			SamplerMirrorClampToEdge:                           true,
			StorageBuffer8BitAccess:                            true,
			StoragePushConstant8:                               true,
			ShaderSharedInt64Atomics:                           true,
			ShaderInt8:                                         true,
			ShaderInputAttachmentArrayDynamicIndexing:          true,
			ShaderStorageTexelBufferArrayDynamicIndexing:       true,
			ShaderSampledImageArrayNonUniformIndexing:          true,
			ShaderStorageImageArrayNonUniformIndexing:          true,
			ShaderUniformTexelBufferArrayNonUniformIndexing:    true,
			DescriptorBindingUniformBufferUpdateAfterBind:      true,
			DescriptorBindingStorageImageUpdateAfterBind:       true,
			DescriptorBindingUniformTexelBufferUpdateAfterBind: true,
			DescriptorBindingUpdateUnusedWhilePending:          true,
			DescriptorBindingVariableDescriptorCount:           true,
			SamplerFilterMinmax:                                true,
			ImagelessFramebuffer:                               true,
			ShaderSubgroupExtendedTypes:                        true,
			HostQueryReset:                                     true,
			BufferDeviceAddress:                                true,
			BufferDeviceAddressMultiDevice:                     true,
			VulkanMemoryModelDeviceScope:                       true,
			ShaderOutputViewportIndex:                          true,
			SubgroupBroadcastDynamicID:                         true,
		}},
	})
	require.NoError(t, err)
	require.NotNil(t, device)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceVulkan12FeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(
			physicalDevice loader.VkPhysicalDevice,
			pFeatures *loader.VkPhysicalDeviceFeatures2,
		) {
			val := reflect.ValueOf(pFeatures).Elem()

			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

			outData := (*loader.VkPhysicalDeviceVulkan12Features)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(outData).Elem()

			require.Equal(t, uint64(51), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_2_FEATURES
			require.True(t, val.FieldByName("pNext").IsNil())

			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("samplerMirrorClampToEdge").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("drawIndirectCount").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("storageBuffer8BitAccess").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("uniformAndStorageBuffer8BitAccess").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("storagePushConstant8").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderBufferInt64Atomics").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSharedInt64Atomics").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderFloat16").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInt8").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorIndexing").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInputAttachmentArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformTexelBufferArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageTexelBufferArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformBufferArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSampledImageArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageBufferArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInputAttachmentArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformTexelBufferArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageTexelBufferArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingUniformBufferUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingSampledImageUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingStorageImageUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingStorageBufferUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingUniformTexelBufferUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingStorageTexelBufferUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingUpdateUnusedWhilePending").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingPartiallyBound").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingVariableDescriptorCount").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("runtimeDescriptorArray").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("samplerFilterMinmax").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("scalarBlockLayout").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("imagelessFramebuffer").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("uniformBufferStandardLayout").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSubgroupExtendedTypes").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("separateDepthStencilLayouts").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("hostQueryReset").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("timelineSemaphore").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("bufferDeviceAddress").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("bufferDeviceAddressCaptureReplay").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("bufferDeviceAddressMultiDevice").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("vulkanMemoryModel").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("vulkanMemoryModelDeviceScope").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("vulkanMemoryModelAvailabilityVisibilityChains").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderOutputViewportIndex").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderOutputLayer").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("subgroupBroadcastDynamicId").UnsafeAddr())) = loader.VkBool32(1)
		})

	var outData core1_2.PhysicalDeviceVulkan12Features
	err := driver.GetPhysicalDeviceFeatures2(physicalDevice, &core1_1.PhysicalDeviceFeatures2{
		NextOutData: common.NextOutData{Next: &outData},
	})
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceVulkan12Features{
		SamplerMirrorClampToEdge:                           true,
		DrawIndirectCount:                                  false,
		StorageBuffer8BitAccess:                            true,
		UniformAndStorageBuffer8BitAccess:                  false,
		StoragePushConstant8:                               true,
		ShaderBufferInt64Atomics:                           false,
		ShaderSharedInt64Atomics:                           true,
		ShaderFloat16:                                      false,
		ShaderInt8:                                         true,
		DescriptorIndexing:                                 false,
		ShaderInputAttachmentArrayDynamicIndexing:          true,
		ShaderUniformTexelBufferArrayDynamicIndexing:       false,
		ShaderStorageTexelBufferArrayDynamicIndexing:       true,
		ShaderUniformBufferArrayNonUniformIndexing:         false,
		ShaderSampledImageArrayNonUniformIndexing:          true,
		ShaderStorageBufferArrayNonUniformIndexing:         false,
		ShaderStorageImageArrayNonUniformIndexing:          true,
		ShaderInputAttachmentArrayNonUniformIndexing:       false,
		ShaderUniformTexelBufferArrayNonUniformIndexing:    true,
		ShaderStorageTexelBufferArrayNonUniformIndexing:    false,
		DescriptorBindingUniformBufferUpdateAfterBind:      true,
		DescriptorBindingSampledImageUpdateAfterBind:       false,
		DescriptorBindingStorageImageUpdateAfterBind:       true,
		DescriptorBindingStorageBufferUpdateAfterBind:      false,
		DescriptorBindingUniformTexelBufferUpdateAfterBind: true,
		DescriptorBindingStorageTexelBufferUpdateAfterBind: false,
		DescriptorBindingUpdateUnusedWhilePending:          true,
		DescriptorBindingPartiallyBound:                    false,
		DescriptorBindingVariableDescriptorCount:           true,
		RuntimeDescriptorArray:                             false,
		SamplerFilterMinmax:                                true,
		ScalarBlockLayout:                                  false,
		ImagelessFramebuffer:                               true,
		UniformBufferStandardLayout:                        false,
		ShaderSubgroupExtendedTypes:                        true,
		SeparateDepthStencilLayouts:                        false,
		HostQueryReset:                                     true,
		TimelineSemaphore:                                  false,
		BufferDeviceAddress:                                true,
		BufferDeviceAddressCaptureReplay:                   false,
		BufferDeviceAddressMultiDevice:                     true,
		VulkanMemoryModel:                                  false,
		VulkanMemoryModelDeviceScope:                       true,
		VulkanMemoryModelAvailabilityVisibilityChains:      false,
		ShaderOutputViewportIndex:                          true,
		ShaderOutputLayer:                                  false,
		SubgroupBroadcastDynamicID:                         true,
	}, outData)
}
