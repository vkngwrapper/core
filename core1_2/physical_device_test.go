package core1_2_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/google/uuid"
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

func TestPhysicalDeviceDriverOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice,
			pProperties *loader.VkPhysicalDeviceProperties2) {

			val := reflect.ValueOf(pProperties).Elem()
			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2

			next := (*loader.VkPhysicalDeviceDriverProperties)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000196000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES
			require.True(t, val.FieldByName("pNext").IsNil())

			*(*uint32)(unsafe.Pointer(val.FieldByName("driverID").UnsafeAddr())) = uint32(10) // VK_DRIVER_ID_GOOGLE_SWIFTSHADER
			*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("major").UnsafeAddr())) = uint8(1)
			*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("minor").UnsafeAddr())) = uint8(3)
			*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("subminor").UnsafeAddr())) = uint8(5)
			*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("patch").UnsafeAddr())) = uint8(7)

			driverNamePtr := (*loader.Char)(unsafe.Pointer(val.FieldByName("driverName").UnsafeAddr()))
			driverNameSlice := ([]loader.Char)(unsafe.Slice(driverNamePtr, 256))
			driverName := "Some loader"
			for i, r := range []byte(driverName) {
				driverNameSlice[i] = loader.Char(r)
			}
			driverNameSlice[len(driverName)] = 0

			driverInfoPtr := (*loader.Char)(unsafe.Pointer(val.FieldByName("driverInfo").UnsafeAddr()))
			driverInfoSlice := ([]loader.Char)(unsafe.Slice(driverInfoPtr, 256))
			driverInfo := "Whooo Info"
			for i, r := range []byte(driverInfo) {
				driverInfoSlice[i] = loader.Char(r)
			}
			driverInfoSlice[len(driverInfo)] = 0
		})

	var driverOutData core1_2.PhysicalDeviceDriverProperties
	err := driver.GetPhysicalDeviceProperties2(
		physicalDevice,
		&core1_1.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&driverOutData},
		})
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceDriverProperties{
		DriverID:           core1_2.DriverIDGoogleSwiftshader,
		DriverName:         "Some loader",
		DriverInfo:         "Whooo Info",
		ConformanceVersion: core1_2.ConformanceVersion{Major: 1, Minor: 3, Subminor: 5, Patch: 7},
	}, driverOutData)
}

func TestPhysicalDeviceDepthStencilResolveOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pProperties *loader.VkPhysicalDeviceProperties2) {

		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2

		next := (*loader.VkPhysicalDeviceDepthStencilResolveProperties)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000199000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES
		require.True(t, val.FieldByName("pNext").IsNil())
		depthResolveModePtr := (*loader.VkResolveModeFlags)(unsafe.Pointer(val.FieldByName("supportedDepthResolveModes").UnsafeAddr()))
		*depthResolveModePtr = loader.VkResolveModeFlags(2) // VK_RESOLVE_MODE_AVERAGE_BIT
		stencilResolveModePtr := (*loader.VkResolveModeFlags)(unsafe.Pointer(val.FieldByName("supportedStencilResolveModes").UnsafeAddr()))
		*stencilResolveModePtr = loader.VkResolveModeFlags(8)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("independentResolveNone").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("independentResolve").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData core1_2.PhysicalDeviceDepthStencilResolveProperties
	err := driver.GetPhysicalDeviceProperties2(
		physicalDevice,
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

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pProperties *loader.VkPhysicalDeviceProperties2) {

		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
		next := (*loader.VkPhysicalDeviceDescriptorIndexingProperties)(val.FieldByName("pNext").UnsafePointer())

		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000161002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxUpdateAfterBindDescriptorsInAllPools").UnsafeAddr())) = loader.Uint32(1)

		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformBufferArrayNonUniformIndexingNative").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSampledImageArrayNonUniformIndexingNative").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageBufferArrayNonUniformIndexingNative").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageArrayNonUniformIndexingNative").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInputAttachmentArrayNonUniformIndexingNative").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("robustBufferAccessUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("quadDivergentImplicitLod").UnsafeAddr())) = loader.VkBool32(1)

		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindSamplers").UnsafeAddr())) = loader.Uint32(3)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindUniformBuffers").UnsafeAddr())) = loader.Uint32(5)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindStorageBuffers").UnsafeAddr())) = loader.Uint32(7)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindSampledImages").UnsafeAddr())) = loader.Uint32(11)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindStorageImages").UnsafeAddr())) = loader.Uint32(13)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindInputAttachments").UnsafeAddr())) = loader.Uint32(17)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageUpdateAfterBindResources").UnsafeAddr())) = loader.Uint32(19)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindSamplers").UnsafeAddr())) = loader.Uint32(23)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindUniformBuffers").UnsafeAddr())) = loader.Uint32(29)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindUniformBuffersDynamic").UnsafeAddr())) = loader.Uint32(31)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageBuffers").UnsafeAddr())) = loader.Uint32(37)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageBuffersDynamic").UnsafeAddr())) = loader.Uint32(41)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindSampledImages").UnsafeAddr())) = loader.Uint32(43)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageImages").UnsafeAddr())) = loader.Uint32(47)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindInputAttachments").UnsafeAddr())) = loader.Uint32(51)
	})

	var outData core1_2.PhysicalDeviceDescriptorIndexingProperties
	err := driver.GetPhysicalDeviceProperties2(
		physicalDevice,
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

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pProperties *loader.VkPhysicalDeviceProperties2) {

		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2

		next := (*loader.VkPhysicalDeviceFloatControlsProperties)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000197000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES
		require.True(t, val.FieldByName("pNext").IsNil())
		denormBehavior := (*loader.VkShaderFloatControlsIndependence)(unsafe.Pointer(val.FieldByName("denormBehaviorIndependence").UnsafeAddr()))
		*denormBehavior = loader.VkShaderFloatControlsIndependence(0) // VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY
		roundingMode := (*loader.VkShaderFloatControlsIndependence)(unsafe.Pointer(val.FieldByName("roundingModeIndependence").UnsafeAddr()))
		*roundingMode = loader.VkShaderFloatControlsIndependence(1) // VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL

		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat16").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat32").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat64").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat16").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat32").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat64").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat16").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat32").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat64").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat16").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat32").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat64").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat16").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat32").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat64").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData core1_2.PhysicalDeviceFloatControlsProperties
	err := driver.GetPhysicalDeviceProperties2(
		physicalDevice,
		&core1_1.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceFloatControlsProperties{
		DenormBehaviorIndependence: core1_2.ShaderFloatControlsIndependence32BitOnly,
		RoundingModeIndependence:   core1_2.ShaderFloatControlsIndependenceAll,

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

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pProperties *loader.VkPhysicalDeviceProperties2) {
		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2

		next := (*loader.VkPhysicalDeviceSamplerFilterMinmaxProperties)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000130000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("filterMinmaxSingleComponentFormats").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("filterMinmaxImageComponentMapping").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData core1_2.PhysicalDeviceSamplerFilterMinmaxProperties
	err := driver.GetPhysicalDeviceProperties2(
		physicalDevice,
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

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties2) {
		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2

		next := (*loader.VkPhysicalDeviceTimelineSemaphoreProperties)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000207001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_PROPERTIES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.Uint64)(unsafe.Pointer(val.FieldByName("maxTimelineSemaphoreValueDifference").UnsafeAddr())) = loader.Uint64(3)
	})

	var outData core1_2.PhysicalDeviceTimelineSemaphoreProperties
	err := driver.GetPhysicalDeviceProperties2(
		physicalDevice,
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

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	deviceUUID, err := uuid.NewRandom()
	require.NoError(t, err)

	driverUUID, err := uuid.NewRandom()
	require.NoError(t, err)

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties2) {
		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2

		next := (*loader.VkPhysicalDeviceVulkan11Properties)(val.FieldByName("pNext").UnsafePointer())
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

		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("deviceNodeMask").UnsafeAddr())) = loader.Uint32(3)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("deviceLUIDValid").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("subgroupSize").UnsafeAddr())) = loader.Uint32(5)
		*(*loader.VkShaderStageFlags)(unsafe.Pointer(val.FieldByName("subgroupSupportedStages").UnsafeAddr())) = loader.VkShaderStageFlags(8)
		*(*loader.VkSubgroupFeatureFlags)(unsafe.Pointer(val.FieldByName("subgroupSupportedOperations").UnsafeAddr())) = loader.VkSubgroupFeatureFlags(4)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("subgroupQuadOperationsInAllStages").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkPointClippingBehavior)(unsafe.Pointer(val.FieldByName("pointClippingBehavior").UnsafeAddr())) = loader.VkPointClippingBehavior(1)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxMultiviewViewCount").UnsafeAddr())) = loader.Uint32(7)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxMultiviewInstanceIndex").UnsafeAddr())) = loader.Uint32(11)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("protectedNoFault").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerSetDescriptors").UnsafeAddr())) = loader.Uint32(13)
		*(*loader.VkDeviceSize)(unsafe.Pointer(val.FieldByName("maxMemoryAllocationSize").UnsafeAddr())) = loader.VkDeviceSize(17)
	})

	var outData core1_2.PhysicalDeviceVulkan11Properties
	err = driver.GetPhysicalDeviceProperties2(
		physicalDevice,
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

	instance := mocks.NewDummyInstance(common.Vulkan1_2, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_2)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties2) {
		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2

		next := (*loader.VkPhysicalDeviceVulkan12Properties)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(52), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_2_PROPERTIES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*uint32)(unsafe.Pointer(val.FieldByName("driverID").UnsafeAddr())) = uint32(10) // VK_DRIVER_ID_GOOGLE_SWIFTSHADER
		*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("major").UnsafeAddr())) = uint8(1)
		*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("minor").UnsafeAddr())) = uint8(3)
		*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("subminor").UnsafeAddr())) = uint8(5)
		*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("patch").UnsafeAddr())) = uint8(7)

		driverNamePtr := (*loader.Char)(unsafe.Pointer(val.FieldByName("driverName").UnsafeAddr()))
		driverNameSlice := ([]loader.Char)(unsafe.Slice(driverNamePtr, 256))
		driverName := "Some loader"
		for i, r := range []byte(driverName) {
			driverNameSlice[i] = loader.Char(r)
		}
		driverNameSlice[len(driverName)] = 0

		driverInfoPtr := (*loader.Char)(unsafe.Pointer(val.FieldByName("driverInfo").UnsafeAddr()))
		driverInfoSlice := ([]loader.Char)(unsafe.Slice(driverInfoPtr, 256))
		driverInfo := "Whooo Info"
		for i, r := range []byte(driverInfo) {
			driverInfoSlice[i] = loader.Char(r)
		}
		driverInfoSlice[len(driverInfo)] = 0

		*(*loader.VkShaderFloatControlsIndependence)(unsafe.Pointer(val.FieldByName("denormBehaviorIndependence").UnsafeAddr())) = loader.VkShaderFloatControlsIndependence(1)
		*(*loader.VkShaderFloatControlsIndependence)(unsafe.Pointer(val.FieldByName("roundingModeIndependence").UnsafeAddr())) = loader.VkShaderFloatControlsIndependence(2)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat16").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat32").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat64").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat16").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat32").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat64").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat16").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat32").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat64").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat16").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat32").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat64").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat16").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat32").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat64").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxUpdateAfterBindDescriptorsInAllPools").UnsafeAddr())) = loader.Uint32(3)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformBufferArrayNonUniformIndexingNative").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSampledImageArrayNonUniformIndexingNative").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageBufferArrayNonUniformIndexingNative").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageArrayNonUniformIndexingNative").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInputAttachmentArrayNonUniformIndexingNative").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("robustBufferAccessUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("quadDivergentImplicitLod").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindSamplers").UnsafeAddr())) = loader.Uint32(5)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindUniformBuffers").UnsafeAddr())) = loader.Uint32(7)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindStorageBuffers").UnsafeAddr())) = loader.Uint32(11)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindSampledImages").UnsafeAddr())) = loader.Uint32(13)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindStorageImages").UnsafeAddr())) = loader.Uint32(17)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindInputAttachments").UnsafeAddr())) = loader.Uint32(19)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageUpdateAfterBindResources").UnsafeAddr())) = loader.Uint32(23)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindSamplers").UnsafeAddr())) = loader.Uint32(29)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindUniformBuffers").UnsafeAddr())) = loader.Uint32(31)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindUniformBuffersDynamic").UnsafeAddr())) = loader.Uint32(37)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageBuffers").UnsafeAddr())) = loader.Uint32(41)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageBuffersDynamic").UnsafeAddr())) = loader.Uint32(43)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindSampledImages").UnsafeAddr())) = loader.Uint32(47)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageImages").UnsafeAddr())) = loader.Uint32(53)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindInputAttachments").UnsafeAddr())) = loader.Uint32(59)
		*(*loader.VkResolveModeFlags)(unsafe.Pointer(val.FieldByName("supportedDepthResolveModes").UnsafeAddr())) = loader.VkResolveModeFlags(4)
		*(*loader.VkResolveModeFlags)(unsafe.Pointer(val.FieldByName("supportedStencilResolveModes").UnsafeAddr())) = loader.VkResolveModeFlags(8)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("independentResolveNone").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("independentResolve").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("filterMinmaxSingleComponentFormats").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("filterMinmaxImageComponentMapping").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.Uint64)(unsafe.Pointer(val.FieldByName("maxTimelineSemaphoreValueDifference").UnsafeAddr())) = loader.Uint64(61)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("framebufferIntegerColorSampleCounts").UnsafeAddr())) = loader.Uint32(16)
	})

	var outData core1_2.PhysicalDeviceVulkan12Properties
	err := driver.GetPhysicalDeviceProperties2(
		physicalDevice,
		&core1_1.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, core1_2.PhysicalDeviceVulkan12Properties{
		DriverID:   core1_2.DriverIDGoogleSwiftshader,
		DriverName: "Some loader",
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
