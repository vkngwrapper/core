package core1_2_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/internal/dummies"
	"github.com/CannibalVox/VKng/core/mocks"
	ext_descriptor_indexing_driver "github.com/CannibalVox/VKng/extensions/ext_descriptor_indexing/driver"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestPhysicalDeviceDriverOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_2.PromoteInstanceScopedPhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice,
			pProperties *driver.VkPhysicalDeviceProperties2) {

			val := reflect.ValueOf(pProperties).Elem()
			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2

			next := (*driver.VkPhysicalDeviceDriverProperties)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000196000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES
			require.True(t, val.FieldByName("pNext").IsNil())

			*(*uint32)(unsafe.Pointer(val.FieldByName("driverID").UnsafeAddr())) = uint32(10) // VK_DRIVER_ID_GOOGLE_SWIFTSHADER
			*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("major").UnsafeAddr())) = uint8(1)
			*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("minor").UnsafeAddr())) = uint8(3)
			*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("subminor").UnsafeAddr())) = uint8(5)
			*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("patch").UnsafeAddr())) = uint8(7)

			driverNamePtr := (*driver.Char)(unsafe.Pointer(val.FieldByName("driverName").UnsafeAddr()))
			driverNameSlice := ([]driver.Char)(unsafe.Slice(driverNamePtr, 256))
			driverName := "Some Driver"
			for i, r := range []byte(driverName) {
				driverNameSlice[i] = driver.Char(r)
			}
			driverNameSlice[len(driverName)] = 0

			driverInfoPtr := (*driver.Char)(unsafe.Pointer(val.FieldByName("driverInfo").UnsafeAddr()))
			driverInfoSlice := ([]driver.Char)(unsafe.Slice(driverInfoPtr, 256))
			driverInfo := "Whooo Info"
			for i, r := range []byte(driverInfo) {
				driverInfoSlice[i] = driver.Char(r)
			}
			driverInfoSlice[len(driverInfo)] = 0
		})

	var driverOutData core1_2.PhysicalDeviceDriverProperties
	err := physicalDevice.Properties2(
		&core1_1.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&driverOutData},
		})
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceDriverProperties{
		DriverID:           core1_2.DriverIDGoogleSwiftshader,
		DriverName:         "Some Driver",
		DriverInfo:         "Whooo Info",
		ConformanceVersion: core1_2.ConformanceVersion{Major: 1, Minor: 3, Subminor: 5, Patch: 7},
	}, driverOutData)
}

func TestPhysicalDeviceDepthStencilResolveOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_2.PromoteInstanceScopedPhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pProperties *driver.VkPhysicalDeviceProperties2) {

		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2

		next := (*driver.VkPhysicalDeviceDepthStencilResolveProperties)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000199000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES
		require.True(t, val.FieldByName("pNext").IsNil())
		depthResolveModePtr := (*driver.VkResolveModeFlags)(unsafe.Pointer(val.FieldByName("supportedDepthResolveModes").UnsafeAddr()))
		*depthResolveModePtr = driver.VkResolveModeFlags(2) // VK_RESOLVE_MODE_AVERAGE_BIT
		stencilResolveModePtr := (*driver.VkResolveModeFlags)(unsafe.Pointer(val.FieldByName("supportedStencilResolveModes").UnsafeAddr()))
		*stencilResolveModePtr = driver.VkResolveModeFlags(8)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("independentResolveNone").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("independentResolve").UnsafeAddr())) = driver.VkBool32(1)
	})

	var outData core1_2.PhysicalDeviceDepthStencilResolveProperties
	err := physicalDevice.Properties2(
		&core1_1.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceDepthStencilResolveProperties{
		SupportedDepthResolveModes:   core1_2.ResolveModeAverage,
		SupportedStencilResolveModes: core1_2.ResolveModeMax,
		IndependentResolve:           true,
		IndependentResolveNone:       false,
	}, outData)
}

func TestPhysicalDeviceDescriptorIndexingOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_2.PromoteInstanceScopedPhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pProperties *driver.VkPhysicalDeviceProperties2) {

		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
		next := (*ext_descriptor_indexing_driver.VkPhysicalDeviceDescriptorIndexingPropertiesEXT)(val.FieldByName("pNext").UnsafePointer())

		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000161002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxUpdateAfterBindDescriptorsInAllPools").UnsafeAddr())) = driver.Uint32(1)

		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformBufferArrayNonUniformIndexingNative").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSampledImageArrayNonUniformIndexingNative").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageBufferArrayNonUniformIndexingNative").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageArrayNonUniformIndexingNative").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInputAttachmentArrayNonUniformIndexingNative").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("robustBufferAccessUpdateAfterBind").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("quadDivergentImplicitLod").UnsafeAddr())) = driver.VkBool32(1)

		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindSamplers").UnsafeAddr())) = driver.Uint32(3)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindUniformBuffers").UnsafeAddr())) = driver.Uint32(5)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindStorageBuffers").UnsafeAddr())) = driver.Uint32(7)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindSampledImages").UnsafeAddr())) = driver.Uint32(11)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindStorageImages").UnsafeAddr())) = driver.Uint32(13)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindInputAttachments").UnsafeAddr())) = driver.Uint32(17)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageUpdateAfterBindResources").UnsafeAddr())) = driver.Uint32(19)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindSamplers").UnsafeAddr())) = driver.Uint32(23)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindUniformBuffers").UnsafeAddr())) = driver.Uint32(29)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindUniformBuffersDynamic").UnsafeAddr())) = driver.Uint32(31)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageBuffers").UnsafeAddr())) = driver.Uint32(37)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageBuffersDynamic").UnsafeAddr())) = driver.Uint32(41)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindSampledImages").UnsafeAddr())) = driver.Uint32(43)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageImages").UnsafeAddr())) = driver.Uint32(47)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindInputAttachments").UnsafeAddr())) = driver.Uint32(51)
	})

	var outData core1_2.PhysicalDeviceDescriptorIndexingProperties
	err := physicalDevice.Properties2(
		&core1_1.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t,
		core1_2.PhysicalDeviceDescriptorIndexingProperties{
			MaxUpdateAfterBindDescriptorsInAllPools: 1,

			ShaderUniformBufferArrayNonUniformIndexingNative:   true,
			ShaderSampledImageArrayNonUniformIndexingNative:    false,
			ShaderStorageBufferArrayNonUniformIndexingNative:   true,
			ShaderStorageImageArrayNonUniformIndexingNative:    false,
			ShaderInputAttachmentArrayNonUniformIndexingNative: true,
			RobustBufferAccessUpdateAfterBind:                  false,
			QuadDivergentImplicitLod:                           true,

			MaxPerStageDescriptorUpdateAfterBindSamplers:         3,
			MaxPerStageDescriptorUpdateAfterBindUniformBuffers:   5,
			MaxPerStageDescriptorUpdateAfterBindStorageBuffers:   7,
			MaxPerStageDescriptorUpdateAfterBindSampledImages:    11,
			MaxPerStageDescriptorUpdateAfterBindStorageImages:    13,
			MaxPerStageDescriptorUpdateAfterBindInputAttachments: 17,
			MaxPerStageUpdateAfterBindResources:                  19,
			MaxDescriptorSetUpdateAfterBindSamplers:              23,
			MaxDescriptorSetUpdateAfterBindUniformBuffers:        29,
			MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic: 31,
			MaxDescriptorSetUpdateAfterBindStorageBuffers:        37,
			MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic: 41,
			MaxDescriptorSetUpdateAfterBindSampledImages:         43,
			MaxDescriptorSetUpdateAfterBindStorageImages:         47,
			MaxDescriptorSetUpdateAfterBindInputAttachments:      51,
		},
		outData)
}

func TestPhysicalDeviceFloatControlsOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_2.PromoteInstanceScopedPhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pProperties *driver.VkPhysicalDeviceProperties2) {

		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2

		next := (*driver.VkPhysicalDeviceFloatControlsProperties)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000197000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES
		require.True(t, val.FieldByName("pNext").IsNil())
		denormBehavior := (*driver.VkShaderFloatControlsIndependence)(unsafe.Pointer(val.FieldByName("denormBehaviorIndependence").UnsafeAddr()))
		*denormBehavior = driver.VkShaderFloatControlsIndependence(0) // VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY
		roundingMode := (*driver.VkShaderFloatControlsIndependence)(unsafe.Pointer(val.FieldByName("roundingModeIndependence").UnsafeAddr()))
		*roundingMode = driver.VkShaderFloatControlsIndependence(1) // VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL

		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat16").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat32").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat64").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat16").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat32").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat64").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat16").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat32").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat64").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat16").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat32").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat64").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat16").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat32").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat64").UnsafeAddr())) = driver.VkBool32(1)
	})

	var outData core1_2.PhysicalDeviceFloatControlsProperties
	err := physicalDevice.Properties2(
		&core1_1.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceFloatControlsProperties{
		DenormBehaviorIndependence: core1_2.ShaderFloatControlsIndependence32BitOnly,
		RoundingMoundIndependence:  core1_2.ShaderFloatControlsIndependenceAll,

		ShaderSignedZeroInfNanPreserveFloat16: true,
		ShaderSignedZeroInfNanPreserveFloat32: false,
		ShaderSignedZeroInfNanPreserveFloat64: true,
		ShaderDenormPreserveFloat16:           false,
		ShaderDenormPreserveFloat32:           true,
		ShaderDenormPreserveFloat64:           false,
		ShaderDenormFlushToZeroFloat16:        true,
		ShaderDenormFlushToZeroFloat32:        false,
		ShaderDenormFlushToZeroFloat64:        true,
		ShaderRoundingModeRTEFloat16:          false,
		ShaderRoundingModeRTEFloat32:          true,
		ShaderRoundingModeRTEFloat64:          false,
		ShaderRoundingModeRTZFloat16:          true,
		ShaderRoundingModeRTZFloat32:          false,
		ShaderRoundingModeRTZFloat64:          true,
	}, outData)
}

func TestPhysicalDeviceSamplerFilterMinmaxOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_2.PromoteInstanceScopedPhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pProperties *driver.VkPhysicalDeviceProperties2) {
		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2

		next := (*driver.VkPhysicalDeviceSamplerFilterMinmaxProperties)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000130000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("filterMinmaxSingleComponentFormats").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("filterMinmaxImageComponentMapping").UnsafeAddr())) = driver.VkBool32(1)
	})

	var outData core1_2.PhysicalDeviceSamplerFilterMinmaxProperties
	err := physicalDevice.Properties2(
		&core1_1.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceSamplerFilterMinmaxProperties{
		FilterMinmaxImageComponentMapping:  true,
		FilterMinmaxSingleComponentFormats: true,
	}, outData)
}

func TestPhysicalDeviceTimelineSemaphoreOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_2.PromoteInstanceScopedPhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties2) {
		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2

		next := (*driver.VkPhysicalDeviceTimelineSemaphoreProperties)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000207001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_PROPERTIES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.Uint64)(unsafe.Pointer(val.FieldByName("maxTimelineSemaphoreValueDifference").UnsafeAddr())) = driver.Uint64(3)
	})

	var outData core1_2.PhysicalDeviceTimelineSemaphoreProperties
	err := physicalDevice.Properties2(
		&core1_1.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceTimelineSemaphoreProperties{
		MaxTimelineSemaphoreValueDifference: 3,
	}, outData)
}

func TestPhysicalDeviceVulkan11OutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_2.PromoteInstanceScopedPhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	deviceUUID, err := uuid.NewRandom()
	require.NoError(t, err)

	driverUUID, err := uuid.NewRandom()
	require.NoError(t, err)

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties2) {
		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2

		next := (*driver.VkPhysicalDeviceVulkan11Properties)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(50), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_1_PROPERTIES
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

		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("deviceNodeMask").UnsafeAddr())) = driver.Uint32(3)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("deviceLUIDValid").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("subgroupSize").UnsafeAddr())) = driver.Uint32(5)
		*(*driver.VkShaderStageFlags)(unsafe.Pointer(val.FieldByName("subgroupSupportedStages").UnsafeAddr())) = driver.VkShaderStageFlags(8)
		*(*driver.VkSubgroupFeatureFlags)(unsafe.Pointer(val.FieldByName("subgroupSupportedOperations").UnsafeAddr())) = driver.VkSubgroupFeatureFlags(4)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("subgroupQuadOperationsInAllStages").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkPointClippingBehavior)(unsafe.Pointer(val.FieldByName("pointClippingBehavior").UnsafeAddr())) = driver.VkPointClippingBehavior(1)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxMultiviewViewCount").UnsafeAddr())) = driver.Uint32(7)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxMultiviewInstanceIndex").UnsafeAddr())) = driver.Uint32(11)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("protectedNoFault").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerSetDescriptors").UnsafeAddr())) = driver.Uint32(13)
		*(*driver.VkDeviceSize)(unsafe.Pointer(val.FieldByName("maxMemoryAllocationSize").UnsafeAddr())) = driver.VkDeviceSize(17)
	})

	var outData core1_2.PhysicalDeviceVulkan11Properties
	err = physicalDevice.Properties2(
		&core1_1.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceVulkan11Properties{
		DeviceUUID:                        deviceUUID,
		DriverUUID:                        driverUUID,
		DeviceLUID:                        0xdeadbeefdeadbeef,
		DeviceNodeMask:                    3,
		DeviceLUIDValid:                   true,
		SubgroupSize:                      5,
		SubgroupSupportedStages:           core1_0.StageGeometry,
		SubgroupSupportedOperations:       core1_1.SubgroupFeatureArithmetic,
		SubgroupQuadOperationsInAllStages: false,
		PointClippingBehavior:             core1_1.PointClippingUserClipPlanesOnly,
		MaxMultiviewViewCount:             7,
		MaxMultiviewInstanceIndex:         11,
		ProtectedNoFault:                  true,
		MaxPerSetDescriptors:              13,
		MaxMemoryAllocationSize:           17,
	}, outData)
}

func TestPhysicalDeviceVulkan12OutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_2.PromoteInstanceScopedPhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties2) {
		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2

		next := (*driver.VkPhysicalDeviceVulkan12Properties)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(52), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_2_PROPERTIES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*uint32)(unsafe.Pointer(val.FieldByName("driverID").UnsafeAddr())) = uint32(10) // VK_DRIVER_ID_GOOGLE_SWIFTSHADER
		*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("major").UnsafeAddr())) = uint8(1)
		*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("minor").UnsafeAddr())) = uint8(3)
		*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("subminor").UnsafeAddr())) = uint8(5)
		*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("patch").UnsafeAddr())) = uint8(7)

		driverNamePtr := (*driver.Char)(unsafe.Pointer(val.FieldByName("driverName").UnsafeAddr()))
		driverNameSlice := ([]driver.Char)(unsafe.Slice(driverNamePtr, 256))
		driverName := "Some Driver"
		for i, r := range []byte(driverName) {
			driverNameSlice[i] = driver.Char(r)
		}
		driverNameSlice[len(driverName)] = 0

		driverInfoPtr := (*driver.Char)(unsafe.Pointer(val.FieldByName("driverInfo").UnsafeAddr()))
		driverInfoSlice := ([]driver.Char)(unsafe.Slice(driverInfoPtr, 256))
		driverInfo := "Whooo Info"
		for i, r := range []byte(driverInfo) {
			driverInfoSlice[i] = driver.Char(r)
		}
		driverInfoSlice[len(driverInfo)] = 0

		*(*driver.VkShaderFloatControlsIndependence)(unsafe.Pointer(val.FieldByName("denormBehaviorIndependence").UnsafeAddr())) = driver.VkShaderFloatControlsIndependence(1)
		*(*driver.VkShaderFloatControlsIndependence)(unsafe.Pointer(val.FieldByName("roundingModeIndependence").UnsafeAddr())) = driver.VkShaderFloatControlsIndependence(2)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat16").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat32").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat64").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat16").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat32").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat64").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat16").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat32").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat64").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat16").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat32").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat64").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat16").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat32").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat64").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxUpdateAfterBindDescriptorsInAllPools").UnsafeAddr())) = driver.Uint32(3)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformBufferArrayNonUniformIndexingNative").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSampledImageArrayNonUniformIndexingNative").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageBufferArrayNonUniformIndexingNative").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageArrayNonUniformIndexingNative").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInputAttachmentArrayNonUniformIndexingNative").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("robustBufferAccessUpdateAfterBind").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("quadDivergentImplicitLod").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindSamplers").UnsafeAddr())) = driver.Uint32(5)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindUniformBuffers").UnsafeAddr())) = driver.Uint32(7)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindStorageBuffers").UnsafeAddr())) = driver.Uint32(11)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindSampledImages").UnsafeAddr())) = driver.Uint32(13)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindStorageImages").UnsafeAddr())) = driver.Uint32(17)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindInputAttachments").UnsafeAddr())) = driver.Uint32(19)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageUpdateAfterBindResources").UnsafeAddr())) = driver.Uint32(23)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindSamplers").UnsafeAddr())) = driver.Uint32(29)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindUniformBuffers").UnsafeAddr())) = driver.Uint32(31)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindUniformBuffersDynamic").UnsafeAddr())) = driver.Uint32(37)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageBuffers").UnsafeAddr())) = driver.Uint32(41)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageBuffersDynamic").UnsafeAddr())) = driver.Uint32(43)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindSampledImages").UnsafeAddr())) = driver.Uint32(47)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageImages").UnsafeAddr())) = driver.Uint32(53)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindInputAttachments").UnsafeAddr())) = driver.Uint32(59)
		*(*driver.VkResolveModeFlags)(unsafe.Pointer(val.FieldByName("supportedDepthResolveModes").UnsafeAddr())) = driver.VkResolveModeFlags(4)
		*(*driver.VkResolveModeFlags)(unsafe.Pointer(val.FieldByName("supportedStencilResolveModes").UnsafeAddr())) = driver.VkResolveModeFlags(8)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("independentResolveNone").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("independentResolve").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("filterMinmaxSingleComponentFormats").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("filterMinmaxImageComponentMapping").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.Uint64)(unsafe.Pointer(val.FieldByName("maxTimelineSemaphoreValueDifference").UnsafeAddr())) = driver.Uint64(61)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("framebufferIntegerColorSampleCounts").UnsafeAddr())) = driver.Uint32(16)
	})

	var outData core1_2.PhysicalDeviceVulkan12Properties
	err := physicalDevice.Properties2(
		&core1_1.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceVulkan12Properties{
		DriverID:   core1_2.DriverIDGoogleSwiftshader,
		DriverName: "Some Driver",
		DriverInfo: "Whooo Info",
		ConformanceVersion: core1_2.ConformanceVersion{
			Major:    1,
			Minor:    3,
			Subminor: 5,
			Patch:    7,
		},

		DenormBehaviorIndependence:                           core1_2.ShaderFloatControlsIndependenceAll,
		RoundingModeIndependence:                             core1_2.ShaderFloatControlsIndependenceNone,
		ShaderSignedZeroInfNanPreserveFloat16:                true,
		ShaderSignedZeroInfNanPreserveFloat32:                false,
		ShaderSignedZeroInfNanPreserveFloat64:                true,
		ShaderDenormPreserveFloat16:                          false,
		ShaderDenormPreserveFloat32:                          true,
		ShaderDenormPreserveFloat64:                          false,
		ShaderDenormFlushToZeroFloat16:                       true,
		ShaderDenormFlushToZeroFloat32:                       false,
		ShaderDenormFlushToZeroFloat64:                       true,
		ShaderRoundingModeRTEFloat16:                         false,
		ShaderRoundingModeRTEFloat32:                         true,
		ShaderRoundingModeRTEFloat64:                         false,
		ShaderRoundingModeRTZFloat16:                         true,
		ShaderRoundingModeRTZFloat32:                         false,
		ShaderRoundingModeRTZFloat64:                         true,
		MaxUpdateAfterBindDescriptorsInAllPools:              3,
		ShaderUniformBufferArrayNonUniformIndexingNative:     false,
		ShaderSampledImageArrayNonUniformIndexingNative:      true,
		ShaderStorageBufferArrayNonUniformIndexingNative:     false,
		ShaderStorageImageArrayNonUniformIndexingNative:      true,
		ShaderInputAttachmentArrayNonUniformIndexingNative:   false,
		RobustBufferAccessUpdateAfterBind:                    true,
		QuadDivergentImplicitLod:                             false,
		MaxPerStageDescriptorUpdateAfterBindSamplers:         5,
		MaxPerStageDescriptorUpdateAfterBindUniformBuffers:   7,
		MaxPerStageDescriptorUpdateAfterBindStorageBuffers:   11,
		MaxPerStageDescriptorUpdateAfterBindSampledImages:    13,
		MaxPerStageDescriptorUpdateAfterBindStorageImages:    17,
		MaxPerStageDescriptorUpdateAfterBindInputAttachments: 19,
		MaxPerStageUpdateAfterBindResources:                  23,
		MaxDescriptorSetUpdateAfterBindSamplers:              29,
		MaxDescriptorSetUpdateAfterBindUniformBuffers:        31,
		MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic: 37,
		MaxDescriptorSetUpdateAfterBindStorageBuffers:        41,
		MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic: 43,
		MaxDescriptorSetUpdateAfterBindSampledImages:         47,
		MaxDescriptorSetUpdateAfterBindStorageImages:         53,
		MaxDescriptorSetUpdateAfterBindInputAttachments:      59,
		SupportedDepthResolveModes:                           core1_2.ResolveModeMin,
		SupportedStencilResolveModes:                         core1_2.ResolveModeMax,
		IndependentResolveNone:                               true,
		IndependentResolve:                                   false,
		FilterMinmaxSingleComponentFormats:                   true,
		FilterMinmaxImageComponentMapping:                    false,
		MaxTimelineSemaphoreValueDifference:                  61,
		FramebufferIntegerColorSampleCounts:                  core1_0.Samples16,
	}, outData)
}
