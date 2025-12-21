package core1_1_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
	"go.uber.org/mock/gomock"
)

func TestPhysicalDeviceIDOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)

	deviceUUID, err := uuid.NewRandom()
	require.NoError(t, err)

	driverUUID, err := uuid.NewRandom()
	require.NoError(t, err)

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(
			physicalDevice driver.VkPhysicalDevice,
			pProperties *driver.VkPhysicalDeviceProperties2,
		) {
			val := reflect.ValueOf(pProperties).Elem()
			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2

			next := (*driver.VkPhysicalDeviceIDProperties)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000071004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES
			require.True(t, val.FieldByName("pNext").IsNil())

			for i := 0; i < len(deviceUUID); i++ {
				*(*byte)(unsafe.Pointer(val.FieldByName("deviceUUID").Index(i).UnsafeAddr())) = deviceUUID[i]
				*(*byte)(unsafe.Pointer(val.FieldByName("driverUUID").Index(i).UnsafeAddr())) = driverUUID[i]
			}

			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(0).UnsafeAddr())) = byte(0xef)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(1).UnsafeAddr())) = byte(0xbe)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(2).UnsafeAddr())) = byte(0xad)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(3).UnsafeAddr())) = byte(0xde)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(4).UnsafeAddr())) = byte(0xef)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(5).UnsafeAddr())) = byte(0xbe)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(6).UnsafeAddr())) = byte(0xad)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(7).UnsafeAddr())) = byte(0xde)

			*(*uint32)(unsafe.Pointer(val.FieldByName("deviceNodeMask").UnsafeAddr())) = uint32(7)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("deviceLUIDValid").UnsafeAddr())) = driver.VkBool32(1)
		})

	var properties core1_1.PhysicalDeviceProperties2
	var outData core1_1.PhysicalDeviceIDProperties
	properties.NextOutData = common.NextOutData{&outData}

	err = physicalDevice.Properties2(
		&properties,
	)
	require.NoError(t, err)
	require.Equal(t, core1_1.PhysicalDeviceIDProperties{
		DeviceUUID:      deviceUUID,
		DriverUUID:      driverUUID,
		DeviceLUID:      0xdeadbeefdeadbeef,
		DeviceNodeMask:  7,
		DeviceLUIDValid: true,
	}, outData)
}

func TestMaintenance3OutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties2(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties2) {
			val := reflect.ValueOf(pProperties).Elem()

			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2

			props := val.FieldByName("properties")
			*(*driver.Uint32)(unsafe.Pointer(props.FieldByName("vendorID").UnsafeAddr())) = driver.Uint32(3)

			maintPtr := (*driver.VkPhysicalDeviceMaintenance3Properties)(val.FieldByName("pNext").UnsafePointer())
			maint := reflect.ValueOf(maintPtr).Elem()

			require.Equal(t, uint64(1000168000), maint.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES
			require.True(t, maint.FieldByName("pNext").IsNil())

			*(*driver.Uint32)(unsafe.Pointer(maint.FieldByName("maxPerSetDescriptors").UnsafeAddr())) = driver.Uint32(5)
			*(*driver.Uint64)(unsafe.Pointer(maint.FieldByName("maxMemoryAllocationSize").UnsafeAddr())) = driver.Uint64(7)
		})

	maintOutData := &core1_1.PhysicalDeviceMaintenance3Properties{}
	outData := &core1_1.PhysicalDeviceProperties2{
		NextOutData: common.NextOutData{Next: maintOutData},
	}
	err := physicalDevice.Properties2(outData)
	require.NoError(t, err)

	require.Equal(t, uint32(3), outData.Properties.VendorID)
	require.Equal(t, 5, maintOutData.MaxPerSetDescriptors)
	require.Equal(t, 7, maintOutData.MaxMemoryAllocationSize)
}

func TestMultiviewPropertiesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pProperties *driver.VkPhysicalDeviceProperties2,
	) {
		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2

		next := (*driver.VkPhysicalDeviceMultiviewProperties)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000053002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*uint32)(unsafe.Pointer(val.FieldByName("maxMultiviewViewCount").UnsafeAddr())) = uint32(5)
		*(*uint32)(unsafe.Pointer(val.FieldByName("maxMultiviewInstanceIndex").UnsafeAddr())) = uint32(3)
	})

	var outData core1_1.PhysicalDeviceMultiviewProperties
	properties := core1_1.PhysicalDeviceProperties2{
		NextOutData: common.NextOutData{&outData},
	}

	err := physicalDevice.Properties2(&properties)
	require.NoError(t, err)
	require.Equal(t, core1_1.PhysicalDeviceMultiviewProperties{
		MaxMultiviewInstanceIndex: 3,
		MaxMultiviewViewCount:     5,
	}, outData)
}

func TestPointClippingOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties2(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties2) {
			val := reflect.ValueOf(pProperties).Elem()

			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
			properties := val.FieldByName("properties")
			*(*uint32)(unsafe.Pointer(properties.FieldByName("vendorID").UnsafeAddr())) = uint32(3)

			limits := properties.FieldByName("limits")
			*(*float32)(unsafe.Pointer(limits.FieldByName("lineWidthGranularity").UnsafeAddr())) = float32(5)

			pointClippingPtr := (*driver.VkPhysicalDevicePointClippingProperties)(val.FieldByName("pNext").UnsafePointer())
			pointClipping := reflect.ValueOf(pointClippingPtr).Elem()

			require.Equal(t, uint64(1000117000), pointClipping.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES
			require.True(t, pointClipping.FieldByName("pNext").IsNil())

			behavior := (*driver.VkPointClippingBehavior)(unsafe.Pointer(pointClipping.FieldByName("pointClippingBehavior").UnsafeAddr()))
			*behavior = driver.VkPointClippingBehavior(1) // VK_POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY
		})

	pointClipping := &core1_1.PhysicalDevicePointClippingProperties{}
	properties := &core1_1.PhysicalDeviceProperties2{
		NextOutData: common.NextOutData{Next: pointClipping},
	}

	err := physicalDevice.Properties2(properties)
	require.NoError(t, err)

	require.Equal(t, uint32(3), properties.Properties.VendorID)
	require.InDelta(t, 5.0, properties.Properties.Limits.LineWidthGranularity, 0.001)

	require.Equal(t, core1_1.PointClippingUserClipPlanesOnly, pointClipping.PointClippingBehavior)
}

func TestPhysicalDeviceProtectedMemoryOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties2(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties2) {
			val := reflect.ValueOf(pProperties).Elem()

			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
			properties := val.FieldByName("properties")
			*(*uint32)(unsafe.Pointer(properties.FieldByName("vendorID").UnsafeAddr())) = uint32(3)

			limits := properties.FieldByName("limits")
			*(*float32)(unsafe.Pointer(limits.FieldByName("lineWidthGranularity").UnsafeAddr())) = float32(5)

			protectedMemoryPtr := (*driver.VkPhysicalDeviceProtectedMemoryProperties)(val.FieldByName("pNext").UnsafePointer())
			protectedMemory := reflect.ValueOf(protectedMemoryPtr).Elem()

			require.Equal(t, uint64(1000145002), protectedMemory.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_PROPERTIES
			require.True(t, protectedMemory.FieldByName("pNext").IsNil())

			noFault := (*driver.VkBool32)(unsafe.Pointer(protectedMemory.FieldByName("protectedNoFault").UnsafeAddr()))
			*noFault = driver.VkBool32(1)
		})

	protectedMemory := &core1_1.PhysicalDeviceProtectedMemoryProperties{}
	properties := &core1_1.PhysicalDeviceProperties2{
		NextOutData: common.NextOutData{Next: protectedMemory},
	}

	err := physicalDevice.Properties2(properties)
	require.NoError(t, err)

	require.Equal(t, uint32(3), properties.Properties.VendorID)
	require.InDelta(t, 5.0, properties.Properties.Limits.LineWidthGranularity, 0.001)

	require.True(t, protectedMemory.ProtectedNoFault)
}

func TestPhysicalDeviceSubgroupOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties2(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties2) {
			val := reflect.ValueOf(pProperties).Elem()

			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
			properties := val.FieldByName("properties")
			*(*uint32)(unsafe.Pointer(properties.FieldByName("vendorID").UnsafeAddr())) = uint32(3)

			limits := properties.FieldByName("limits")
			*(*float32)(unsafe.Pointer(limits.FieldByName("lineWidthGranularity").UnsafeAddr())) = float32(5)

			subgroupPtr := (*driver.VkPhysicalDeviceSubgroupProperties)(val.FieldByName("pNext").UnsafePointer())
			subgroup := reflect.ValueOf(subgroupPtr).Elem()

			require.Equal(t, uint64(1000094000), subgroup.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_PROPERTIES
			require.True(t, subgroup.FieldByName("pNext").IsNil())

			*(*uint32)(unsafe.Pointer(subgroup.FieldByName("subgroupSize").UnsafeAddr())) = uint32(1)
			stages := (*driver.VkShaderStageFlags)(unsafe.Pointer(subgroup.FieldByName("supportedStages").UnsafeAddr()))
			*stages = driver.VkShaderStageFlags(0x10) // VK_SHADER_STAGE_FRAGMENT_BIT

			operations := (*driver.VkSubgroupFeatureFlags)(unsafe.Pointer(subgroup.FieldByName("supportedOperations").UnsafeAddr()))
			*operations = driver.VkSubgroupFeatureFlags(8) // VK_SUBGROUP_FEATURE_BALLOT_BIT

			*(*driver.VkBool32)(unsafe.Pointer(subgroup.FieldByName("quadOperationsInAllStages").UnsafeAddr())) = driver.VkBool32(1)
		})

	subgroups := &core1_1.PhysicalDeviceSubgroupProperties{}
	properties := &core1_1.PhysicalDeviceProperties2{
		NextOutData: common.NextOutData{Next: subgroups},
	}

	err := physicalDevice.Properties2(properties)
	require.NoError(t, err)

	require.Equal(t, uint32(3), properties.Properties.VendorID)
	require.InDelta(t, 5.0, properties.Properties.Limits.LineWidthGranularity, 0.001)

	require.Equal(t, subgroups, &core1_1.PhysicalDeviceSubgroupProperties{
		SubgroupSize:              1,
		SupportedStages:           core1_0.StageFragment,
		SupportedOperations:       core1_1.SubgroupFeatureBallot,
		QuadOperationsInAllStages: true,
	})
}

func TestDevice16BitStorageOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)

	coreDriver.EXPECT().VkGetPhysicalDeviceFeatures2(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pFeatures *driver.VkPhysicalDeviceFeatures2) {
			val := reflect.ValueOf(pFeatures).Elem()

			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2

			featureVal := val.FieldByName("features")
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fillModeNonSolid").UnsafeAddr())) = driver.VkBool32(1)

			outDataPtr := (*driver.VkPhysicalDevice16BitStorageFeatures)(val.FieldByName("pNext").UnsafePointer())
			outDataVal := reflect.ValueOf(outDataPtr).Elem()
			*(*driver.VkBool32)(unsafe.Pointer(outDataVal.FieldByName("storageBuffer16BitAccess").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(outDataVal.FieldByName("uniformAndStorageBuffer16BitAccess").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(outDataVal.FieldByName("storagePushConstant16").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(outDataVal.FieldByName("storageInputOutput16").UnsafeAddr())) = driver.VkBool32(1)
		})

	outData := &core1_1.PhysicalDevice16BitStorageFeatures{}
	features := &core1_1.PhysicalDeviceFeatures2{
		NextOutData: common.NextOutData{Next: outData},
	}

	err := physicalDevice.Features2(features)
	require.NoError(t, err)

	require.True(t, outData.StoragePushConstant16)
	require.False(t, outData.UniformAndStorageBuffer16BitAccess)
	require.True(t, outData.StorageInputOutput16)
	require.False(t, outData.StorageBuffer16BitAccess)

	require.True(t, features.Features.FillModeNonSolid)
}

func TestMultiviewFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)
	mockDevice := mocks1_1.EasyMockDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pCreateInfo *driver.VkDeviceCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pDevice *driver.VkDevice,
	) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
		require.Equal(t, uint64(1), val.FieldByName("queueCreateInfoCount").Uint())

		queueCreate := (*driver.VkDeviceQueueCreateInfo)(val.FieldByName("pQueueCreateInfos").UnsafePointer())

		queueFamilyVal := reflect.ValueOf(queueCreate).Elem()
		require.Equal(t, uint64(2), queueFamilyVal.FieldByName("sType").Uint()) //VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
		require.True(t, queueFamilyVal.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), queueFamilyVal.FieldByName("queueCount").Uint())

		next := (*driver.VkPhysicalDeviceMultiviewFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000053001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("multiview").Uint())
		require.Equal(t, uint64(0), val.FieldByName("multiviewGeometryShader").Uint())
		require.Equal(t, uint64(1), val.FieldByName("multiviewTessellationShader").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := physicalDevice.CreateDevice(nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueuePriorities: []float32{3, 2, 1},
			},
		},
		NextOptions: common.NextOptions{
			core1_1.PhysicalDeviceMultiviewFeatures{
				Multiview:                   true,
				MultiviewTessellationShader: true,
				MultiviewGeometryShader:     false,
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestMultiviewFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)

	coreDriver.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pFeatures *driver.VkPhysicalDeviceFeatures2,
	) {
		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2

		next := (*driver.VkPhysicalDeviceMultiviewFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000053001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("multiview").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("multiviewGeometryShader").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("multiviewTessellationShader").UnsafeAddr())) = driver.VkBool32(0)
	})

	var outData core1_1.PhysicalDeviceMultiviewFeatures
	features := core1_1.PhysicalDeviceFeatures2{
		NextOutData: common.NextOutData{&outData},
	}

	err := physicalDevice.Features2(&features)
	require.NoError(t, err)
	require.Equal(t, core1_1.PhysicalDeviceMultiviewFeatures{
		Multiview:                   true,
		MultiviewTessellationShader: false,
		MultiviewGeometryShader:     true,
	}, outData)
}

func TestPhysicalDeviceProtectedMemoryFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)

	mockDevice := mocks1_1.EasyMockDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pCreateInfo *driver.VkDeviceCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pDevice *driver.VkDevice,
	) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
		require.Equal(t, uint64(1), val.FieldByName("queueCreateInfoCount").Uint())

		queueCreate := (*driver.VkDeviceQueueCreateInfo)(val.FieldByName("pQueueCreateInfos").UnsafePointer())

		queueFamilyVal := reflect.ValueOf(queueCreate).Elem()
		require.Equal(t, uint64(2), queueFamilyVal.FieldByName("sType").Uint()) //VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
		require.True(t, queueFamilyVal.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), queueFamilyVal.FieldByName("queueCount").Uint())

		next := (*driver.VkPhysicalDeviceProtectedMemoryFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000145001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("protectedMemory").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := physicalDevice.CreateDevice(nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueuePriorities: []float32{3, 2, 1},
			},
		},
		NextOptions: common.NextOptions{
			core1_1.PhysicalDeviceProtectedMemoryFeatures{
				ProtectedMemory: true,
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceProtectedMemoryFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)

	coreDriver.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pFeatures *driver.VkPhysicalDeviceFeatures2,
	) {
		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2

		next := (*driver.VkPhysicalDeviceProtectedMemoryFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000145001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("protectedMemory").UnsafeAddr())) = driver.VkBool32(1)
	})

	var outData core1_1.PhysicalDeviceProtectedMemoryFeatures
	features := core1_1.PhysicalDeviceFeatures2{
		NextOutData: common.NextOutData{&outData},
	}

	err := physicalDevice.Features2(&features)
	require.NoError(t, err)
	require.Equal(t, core1_1.PhysicalDeviceProtectedMemoryFeatures{
		ProtectedMemory: true,
	}, outData)
}

func TestSamplerYcbcrFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)
	mockDevice := mocks1_1.EasyMockDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice,
			pCreateInfo *driver.VkDeviceCreateInfo,
			pAllocator *driver.VkAllocationCallbacks,
			pDevice *driver.VkDevice,
		) (common.VkResult, error) {
			*pDevice = mockDevice.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

			next := (*driver.VkPhysicalDeviceSamplerYcbcrConversionFeatures)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000156004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("samplerYcbcrConversion").Uint())

			return core1_0.VKSuccess, nil
		})

	device, _, err := physicalDevice.CreateDevice(
		nil,
		core1_0.DeviceCreateInfo{
			QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
				{
					QueuePriorities: []float32{0},
				},
			},

			NextOptions: common.NextOptions{
				core1_1.PhysicalDeviceSamplerYcbcrConversionFeatures{
					SamplerYcbcrConversion: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestSamplerYcbcrFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)

	coreDriver.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice,
			pFeatures *driver.VkPhysicalDeviceFeatures2,
		) {
			val := reflect.ValueOf(pFeatures).Elem()
			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2

			next := (*driver.VkPhysicalDeviceSamplerYcbcrConversionFeatures)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000156004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES
			require.True(t, val.FieldByName("pNext").IsNil())
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("samplerYcbcrConversion").UnsafeAddr())) = driver.VkBool32(1)
		})

	var outData core1_1.PhysicalDeviceSamplerYcbcrConversionFeatures

	err := physicalDevice.Features2(
		&core1_1.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{
				&outData,
			},
		})
	require.NoError(t, err)
	require.Equal(t, core1_1.PhysicalDeviceSamplerYcbcrConversionFeatures{
		SamplerYcbcrConversion: true,
	}, outData)
}

func TestPhysicalDeviceShaderDrawParametersFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)
	mockDevice := mocks1_1.EasyMockDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pCreateInfo *driver.VkDeviceCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pDevice *driver.VkDevice,
	) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
		require.Equal(t, uint64(1), val.FieldByName("queueCreateInfoCount").Uint())

		queueCreate := (*driver.VkDeviceQueueCreateInfo)(val.FieldByName("pQueueCreateInfos").UnsafePointer())

		queueFamilyVal := reflect.ValueOf(queueCreate).Elem()
		require.Equal(t, uint64(2), queueFamilyVal.FieldByName("sType").Uint()) //VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
		require.True(t, queueFamilyVal.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), queueFamilyVal.FieldByName("queueCount").Uint())

		next := (*driver.VkPhysicalDeviceShaderDrawParametersFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000063000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETERS_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("shaderDrawParameters").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := physicalDevice.CreateDevice(nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueuePriorities: []float32{3, 2, 1},
			},
		},
		NextOptions: common.NextOptions{
			core1_1.PhysicalDeviceShaderDrawParametersFeatures{
				ShaderDrawParameters: true,
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceShaderDrawParametersFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)

	coreDriver.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pFeatures *driver.VkPhysicalDeviceFeatures2,
	) {
		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2

		next := (*driver.VkPhysicalDeviceShaderDrawParametersFeatures)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000063000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETERS_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDrawParameters").UnsafeAddr())) = driver.VkBool32(1)
	})

	var outData core1_1.PhysicalDeviceShaderDrawParametersFeatures
	features := core1_1.PhysicalDeviceFeatures2{
		NextOutData: common.NextOutData{&outData},
	}

	err := physicalDevice.Features2(&features)
	require.NoError(t, err)
	require.Equal(t, core1_1.PhysicalDeviceShaderDrawParametersFeatures{
		ShaderDrawParameters: true,
	}, outData)
}

func TestVariablePointersFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)
	mockDevice := mocks1_1.EasyMockDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice,
			pCreateInfo *driver.VkDeviceCreateInfo,
			pAllocator *driver.VkAllocationCallbacks,
			pDevice *driver.VkDevice) (common.VkResult, error) {

			*pDevice = mockDevice.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

			featuresPtr := (*driver.VkPhysicalDeviceVariablePointersFeatures)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(featuresPtr).Elem()

			require.Equal(t, uint64(1000120000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("variablePointers").Uint())
			require.Equal(t, uint64(0), val.FieldByName("variablePointersStorageBuffer").Uint())

			return core1_0.VKSuccess, nil
		})

	device, _, err := physicalDevice.CreateDevice(nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueuePriorities: []float32{0},
			},
		},

		NextOptions: common.NextOptions{Next: core1_1.PhysicalDeviceVariablePointersFeatures{
			VariablePointers:              true,
			VariablePointersStorageBuffer: false,
		}},
	})
	require.NoError(t, err)
	require.NotNil(t, device)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestVariablePointersFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)

	var pointersOutData core1_1.PhysicalDeviceVariablePointersFeatures

	coreDriver.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(
			physicalDevice driver.VkPhysicalDevice,
			pFeatures *driver.VkPhysicalDeviceFeatures2,
		) {
			val := reflect.ValueOf(pFeatures).Elem()

			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2

			outData := (*driver.VkPhysicalDeviceVariablePointersFeatures)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(outData).Elem()

			require.Equal(t, uint64(1000120000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES
			require.True(t, val.FieldByName("pNext").IsNil())

			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("variablePointers").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("variablePointersStorageBuffer").UnsafeAddr())) = driver.VkBool32(1)
		})

	err := physicalDevice.Features2(&core1_1.PhysicalDeviceFeatures2{
		NextOutData: common.NextOutData{Next: &pointersOutData},
	})
	require.NoError(t, err)
	require.True(t, pointersOutData.VariablePointersStorageBuffer)
	require.False(t, pointersOutData.VariablePointers)
}

func TestDeviceGroupOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil)
	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice1 := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)
	physicalDevice2 := mocks1_1.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice3 := mocks1_1.EasyMockPhysicalDevice(ctrl, coreDriver)

	handle := mocks.NewFakeDeviceHandle()

	coreDriver.EXPECT().VkCreateDevice(
		physicalDevice1.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pCreateInfo *driver.VkDeviceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDevice *driver.VkDevice) (common.VkResult, error) {
		*pDevice = handle

		val := reflect.ValueOf(pCreateInfo).Elem()

		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		optionsPtr := (*driver.VkDeviceGroupDeviceCreateInfo)(val.FieldByName("pNext").UnsafePointer())
		options := reflect.ValueOf(optionsPtr).Elem()

		require.Equal(t, uint64(1000070001), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO
		require.True(t, options.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), options.FieldByName("physicalDeviceCount").Uint())

		devicePtr := (*driver.VkPhysicalDevice)(options.FieldByName("pPhysicalDevices").UnsafePointer())
		deviceSlice := ([]driver.VkPhysicalDevice)(unsafe.Slice(devicePtr, 3))
		require.Equal(t, physicalDevice1.Handle(), deviceSlice[0])
		require.Equal(t, physicalDevice2.Handle(), deviceSlice[1])
		require.Equal(t, physicalDevice3.Handle(), deviceSlice[2])

		return core1_0.VKSuccess, nil
	})

	device, _, err := physicalDevice1.CreateDevice(nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueuePriorities: []float32{0},
			},
		},
		NextOptions: common.NextOptions{Next: core1_1.DeviceGroupDeviceCreateInfo{
			PhysicalDevices: []core1_0.PhysicalDevice{physicalDevice1, physicalDevice2, physicalDevice3},
		}},
	})
	require.NoError(t, err)
	require.Equal(t, handle, device.Handle())
}

func TestMemoryAllocateFlagsOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_1, []string{}).(core1_1.Device)

	mockMemory := mocks1_1.EasyMockDeviceMemory(ctrl)

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

			next := (*driver.VkMemoryAllocateFlagsInfo)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000060000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) //VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT
			require.Equal(t, uint64(5), val.FieldByName("deviceMask").Uint())

			return core1_0.VKSuccess, nil
		})

	memory, _, err := device.AllocateMemory(nil,
		core1_0.MemoryAllocateInfo{
			AllocationSize:  1,
			MemoryTypeIndex: 3,
			NextOptions: common.NextOptions{Next: core1_1.MemoryAllocateFlagsInfo{
				Flags:      core1_1.MemoryAllocateDeviceMask,
				DeviceMask: 5,
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockMemory.Handle(), memory.Handle())
}

func TestDevice16BitStorageOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil)

	instance := mocks1_1.EasyMockInstance(ctrl, coreDriver)
	builder := &impl1_1.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_1, common.Vulkan1_1).(core1_1.PhysicalDevice)
	expectedDevice := mocks1_1.EasyMockDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkCreateDevice(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pCreateInfo *driver.VkDeviceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDevice *driver.VkDevice) (common.VkResult, error) {
			*pDevice = expectedDevice.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())

			storageFeatures := (*driver.VkPhysicalDevice16BitStorageFeatures)(val.FieldByName("pNext").UnsafePointer())
			storageVal := reflect.ValueOf(storageFeatures).Elem()

			require.Equal(t, uint64(1000083000), storageVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES
			require.True(t, storageVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), storageVal.FieldByName("storageBuffer16BitAccess").Uint())
			require.Equal(t, uint64(1), storageVal.FieldByName("uniformAndStorageBuffer16BitAccess").Uint())
			require.Equal(t, uint64(0), storageVal.FieldByName("storagePushConstant16").Uint())
			require.Equal(t, uint64(1), storageVal.FieldByName("storageInputOutput16").Uint())

			return core1_0.VKSuccess, nil
		})

	storage := core1_1.PhysicalDevice16BitStorageFeatures{
		StorageInputOutput16:               true,
		UniformAndStorageBuffer16BitAccess: true,
		StoragePushConstant16:              false,
		StorageBuffer16BitAccess:           false,
	}
	device, _, err := physicalDevice.CreateDevice(nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueuePriorities: []float32{0},
			},
		},

		NextOptions: common.NextOptions{Next: storage},
	})

	require.NoError(t, err)
	require.NotNil(t, device)
	require.Equal(t, expectedDevice.Handle(), device.Handle())
}
