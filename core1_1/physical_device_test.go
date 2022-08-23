package core1_1_test

import (
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/internal/dummies"
	"github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanPhysicalDevice_PhysicalDeviceExternalFenceProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)

	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceExternalFenceProperties(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pExternalFenceInfo *driver.VkPhysicalDeviceExternalFenceInfo,
		pExternalFenceProperties *driver.VkExternalFenceProperties,
	) {
		val := reflect.ValueOf(pExternalFenceInfo).Elem()
		require.Equal(t, uint64(1000112000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(4), val.FieldByName("handleType").Uint()) // VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT

		val = reflect.ValueOf(pExternalFenceProperties).Elem()
		*(*uint32)(unsafe.Pointer(val.FieldByName("exportFromImportedHandleTypes").UnsafeAddr())) = uint32(8) // VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT
		*(*uint32)(unsafe.Pointer(val.FieldByName("compatibleHandleTypes").UnsafeAddr())) = uint32(4)         // VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalFenceFeatures").UnsafeAddr())) = uint32(1)         // VK_EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT
	})

	var outData core1_1.ExternalFenceProperties
	err := physicalDevice.InstanceScopedPhysicalDevice1_1().ExternalFenceProperties(
		core1_1.PhysicalDeviceExternalFenceInfo{
			HandleType: core1_1.ExternalFenceHandleTypeOpaqueWin32KMT,
		},
		&outData,
	)
	require.NoError(t, err)
	require.Equal(t, core1_1.ExternalFenceProperties{
		ExportFromImportedHandleTypes: core1_1.ExternalFenceHandleTypeSyncFD,
		CompatibleHandleTypes:         core1_1.ExternalFenceHandleTypeOpaqueWin32KMT,
		ExternalFenceFeatures:         core1_1.ExternalFenceFeatureExportable,
	}, outData)
}

func TestVulkanPhysicalDevice_ExternalBufferProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceExternalBufferProperties(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pExternalBufferInfo *driver.VkPhysicalDeviceExternalBufferInfo, pExternalBufferProperties *driver.VkExternalBufferProperties) {
		val := reflect.ValueOf(pExternalBufferInfo).Elem()

		require.Equal(t, uint64(1000071002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("flags").Uint())               // VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT
		require.Equal(t, uint64(8), val.FieldByName("usage").Uint())               // VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT
		require.Equal(t, uint64(0x00000010), val.FieldByName("handleType").Uint()) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT

		val = reflect.ValueOf(pExternalBufferProperties).Elem()
		require.Equal(t, uint64(1000071003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalMemoryProperties").FieldByName("externalMemoryFeatures").UnsafeAddr())) = uint32(2)
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalMemoryProperties").FieldByName("exportFromImportedHandleTypes").UnsafeAddr())) = uint32(0x40)
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalMemoryProperties").FieldByName("compatibleHandleTypes").UnsafeAddr())) = uint32(2)
	})

	var outData core1_1.ExternalBufferProperties
	err := physicalDevice.InstanceScopedPhysicalDevice1_1().ExternalBufferProperties(
		core1_1.PhysicalDeviceExternalBufferInfo{
			Flags:      core1_0.BufferCreateSparseResidency,
			Usage:      core1_0.BufferUsageStorageTexelBuffer,
			HandleType: core1_1.ExternalMemoryHandleTypeD3D11TextureKMT,
		},
		&outData,
	)
	require.NoError(t, err)
	require.Equal(t, core1_1.ExternalBufferProperties{
		ExternalMemoryProperties: core1_1.ExternalMemoryProperties{
			ExternalMemoryFeatures:        core1_1.ExternalMemoryFeatureExportable,
			ExportFromImportedHandleTypes: core1_1.ExternalMemoryHandleTypeD3D12Resource,
			CompatibleHandleTypes:         core1_1.ExternalMemoryHandleTypeOpaqueWin32,
		},
	}, outData)
}

func TestVulkanPhysicalDevice_ExternalSemaphoreProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceExternalSemaphoreProperties(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice,
			pExternalSemaphoreInfo *driver.VkPhysicalDeviceExternalSemaphoreInfo,
			pExternalSemaphoreProperties *driver.VkExternalSemaphoreProperties,
		) {
			val := reflect.ValueOf(pExternalSemaphoreInfo).Elem()

			require.Equal(t, uint64(1000076000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x10), val.FieldByName("handleType").Uint()) // VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT

			val = reflect.ValueOf(pExternalSemaphoreProperties).Elem()
			require.Equal(t, uint64(1000076001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES
			require.True(t, val.FieldByName("pNext").IsNil())

			*(*uint32)(unsafe.Pointer(val.FieldByName("exportFromImportedHandleTypes").UnsafeAddr())) = uint32(1) // VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT
			*(*uint32)(unsafe.Pointer(val.FieldByName("compatibleHandleTypes").UnsafeAddr())) = uint32(4)         // VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
			*(*uint32)(unsafe.Pointer(val.FieldByName("externalSemaphoreFeatures").UnsafeAddr())) = uint32(2)     // VK_EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT
		})

	var outData core1_1.ExternalSemaphoreProperties
	err := physicalDevice.InstanceScopedPhysicalDevice1_1().ExternalSemaphoreProperties(
		core1_1.PhysicalDeviceExternalSemaphoreInfo{
			HandleType: core1_1.ExternalSemaphoreHandleTypeSyncFD,
		},
		&outData)
	require.NoError(t, err)
	require.Equal(t, core1_1.ExternalSemaphoreProperties{
		ExportFromImportedHandleTypes: core1_1.ExternalSemaphoreHandleTypeOpaqueFD,
		CompatibleHandleTypes:         core1_1.ExternalSemaphoreHandleTypeOpaqueWin32KMT,
		ExternalSemaphoreFeatures:     core1_1.ExternalSemaphoreFeatureImportable,
	}, outData)
}

func TestVulkanPhysicalDevice_Features(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceFeatures2(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
			pFeatures *driver.VkPhysicalDeviceFeatures2) {
			val := reflect.ValueOf(pFeatures).Elem()

			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
			require.True(t, val.FieldByName("pNext").IsNil())

			featureVal := val.FieldByName("features")
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("robustBufferAccess").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fullDrawIndexUint32").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("imageCubeArray").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("independentBlend").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("geometryShader").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("tessellationShader").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sampleRateShading").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("dualSrcBlend").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("logicOp").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("multiDrawIndirect").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("drawIndirectFirstInstance").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("depthClamp").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("depthBiasClamp").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fillModeNonSolid").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("depthBounds").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("wideLines").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("largePoints").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("alphaToOne").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("multiViewport").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("samplerAnisotropy").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("textureCompressionETC2").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("textureCompressionASTC_LDR").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("textureCompressionBC").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("occlusionQueryPrecise").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("pipelineStatisticsQuery").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("vertexPipelineStoresAndAtomics").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fragmentStoresAndAtomics").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderTessellationAndGeometryPointSize").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderImageGatherExtended").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageExtendedFormats").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageMultisample").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageReadWithoutFormat").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageWriteWithoutFormat").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderUniformBufferArrayDynamicIndexing").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderSampledImageArrayDynamicIndexing").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageBufferArrayDynamicIndexing").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageArrayDynamicIndexing").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderClipDistance").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderCullDistance").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderFloat64").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderInt64").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderInt16").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderResourceResidency").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderResourceMinLod").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseBinding").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyBuffer").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyImage2D").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyImage3D").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency2Samples").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency4Samples").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency8Samples").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency16Samples").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyAliased").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("variableMultisampleRate").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("inheritedQueries").UnsafeAddr())) = driver.VkBool32(1)

		})

	outData := &core1_1.PhysicalDeviceFeatures2{}
	err := physicalDevice.InstanceScopedPhysicalDevice1_1().Features2(outData)
	require.NoError(t, err)

	features := outData.Features
	require.NotNil(t, features)
	require.True(t, features.RobustBufferAccess)
	require.False(t, features.FullDrawIndexUint32)
	require.True(t, features.ImageCubeArray)
	require.False(t, features.IndependentBlend)
	require.True(t, features.GeometryShader)
	require.False(t, features.TessellationShader)
	require.True(t, features.SampleRateShading)
	require.False(t, features.DualSrcBlend)
	require.True(t, features.LogicOp)
	require.False(t, features.MultiDrawIndirect)
	require.True(t, features.DrawIndirectFirstInstance)
	require.False(t, features.DepthClamp)
	require.True(t, features.DepthBiasClamp)
	require.False(t, features.FillModeNonSolid)
	require.True(t, features.DepthBounds)
	require.False(t, features.WideLines)
	require.True(t, features.LargePoints)
	require.False(t, features.AlphaToOne)
	require.True(t, features.MultiViewport)
	require.False(t, features.SamplerAnisotropy)
	require.True(t, features.TextureCompressionEtc2)
	require.False(t, features.TextureCompressionAstcLdc)
	require.True(t, features.TextureCompressionBc)
	require.False(t, features.OcclusionQueryPrecise)
	require.True(t, features.PipelineStatisticsQuery)
	require.False(t, features.VertexPipelineStoresAndAtomics)
	require.True(t, features.FragmentStoresAndAtomics)
	require.False(t, features.ShaderTessellationAndGeometryPointSize)
	require.True(t, features.ShaderImageGatherExtended)
	require.False(t, features.ShaderStorageImageExtendedFormats)
	require.True(t, features.ShaderStorageImageMultisample)
	require.False(t, features.ShaderStorageImageReadWithoutFormat)
	require.True(t, features.ShaderStorageImageWriteWithoutFormat)
	require.False(t, features.ShaderUniformBufferArrayDynamicIndexing)
	require.True(t, features.ShaderSampledImageArrayDynamicIndexing)
	require.False(t, features.ShaderStorageBufferArrayDynamicIndexing)
	require.True(t, features.ShaderStorageImageArrayDynamicIndexing)
	require.False(t, features.ShaderClipDistance)
	require.True(t, features.ShaderCullDistance)
	require.False(t, features.ShaderFloat64)
	require.True(t, features.ShaderInt64)
	require.False(t, features.ShaderInt16)
	require.True(t, features.ShaderResourceResidency)
	require.False(t, features.ShaderResourceMinLod)
	require.True(t, features.SparseBinding)
	require.False(t, features.SparseResidencyBuffer)
	require.True(t, features.SparseResidencyImage2D)
	require.False(t, features.SparseResidencyImage3D)
	require.True(t, features.SparseResidency2Samples)
	require.False(t, features.SparseResidency4Samples)
	require.True(t, features.SparseResidency8Samples)
	require.False(t, features.SparseResidency16Samples)
	require.True(t, features.SparseResidencyAliased)
	require.False(t, features.VariableMultisampleRate)
	require.True(t, features.InheritedQueries)
}

func TestVulkanPhysicalDevice_FormatProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceFormatProperties2(
		physicalDevice.Handle(),
		driver.VkFormat(64), // VK_FORMAT_A2B10G10R10_UNORM_PACK32
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		format driver.VkFormat,
		pFormatProperties *driver.VkFormatProperties2) {

		val := reflect.ValueOf(pFormatProperties).Elem()
		require.Equal(t, uint64(1000059002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2
		require.True(t, val.FieldByName("pNext").IsNil())

		properties := val.FieldByName("formatProperties")
		*(*uint32)(unsafe.Pointer(properties.FieldByName("optimalTilingFeatures").UnsafeAddr())) = uint32(0x00000100) // VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT
		*(*uint32)(unsafe.Pointer(properties.FieldByName("linearTilingFeatures").UnsafeAddr())) = uint32(0x00000400)  // VK_FORMAT_FEATURE_BLIT_SRC_BIT
		*(*uint32)(unsafe.Pointer(properties.FieldByName("bufferFeatures").UnsafeAddr())) = uint32(0x00000010)        // VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT
	})

	outData := core1_1.FormatProperties2{}
	err := physicalDevice.InstanceScopedPhysicalDevice1_1().FormatProperties2(
		core1_0.FormatA2B10G10R10UnsignedNormalizedPacked,
		&outData)
	require.NoError(t, err)

	require.Equal(t, core1_0.FormatFeatureColorAttachmentBlend, outData.FormatProperties.OptimalTilingFeatures)
	require.Equal(t, core1_0.FormatFeatureBlitSource, outData.FormatProperties.LinearTilingFeatures)
	require.Equal(t, core1_0.FormatFeatureStorageTexelBuffer, outData.FormatProperties.BufferFeatures)
}

func TestVulkanPhysicalDevice_ImageFormatProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceImageFormatProperties2(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
			pImageFormatInfo *driver.VkPhysicalDeviceImageFormatInfo2,
			pImageFormatProperties *driver.VkImageFormatProperties2,
		) (common.VkResult, error) {
			optionVal := reflect.ValueOf(*pImageFormatInfo)

			require.Equal(t, uint64(1000059004), optionVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2
			require.True(t, optionVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(68), optionVal.FieldByName("format").Uint())        // VK_FORMAT_A2B10G10R10_UINT_PACK32
			require.Equal(t, uint64(1), optionVal.FieldByName("_type").Uint())          // VK_IMAGE_TYPE_2D
			require.Equal(t, uint64(0), optionVal.FieldByName("tiling").Uint())         // VK_IMAGE_TILING_OPTIMAL
			require.Equal(t, uint64(8), optionVal.FieldByName("usage").Uint())          // VK_IMAGE_USAGE_STORAGE_BIT
			require.Equal(t, uint64(0x00000010), optionVal.FieldByName("flags").Uint()) // VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT

			outDataVal := reflect.ValueOf(pImageFormatProperties).Elem()
			require.Equal(t, uint64(1000059003), outDataVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2
			require.True(t, outDataVal.FieldByName("pNext").IsNil())

			formatPropertiesVal := outDataVal.FieldByName("imageFormatProperties")

			*(*uint32)(unsafe.Pointer(formatPropertiesVal.FieldByName("maxMipLevels").UnsafeAddr())) = uint32(1)
			*(*uint32)(unsafe.Pointer(formatPropertiesVal.FieldByName("maxArrayLayers").UnsafeAddr())) = uint32(3)
			*(*uint64)(unsafe.Pointer(formatPropertiesVal.FieldByName("maxResourceSize").UnsafeAddr())) = uint64(5)
			*(*uint32)(unsafe.Pointer(formatPropertiesVal.FieldByName("sampleCounts").UnsafeAddr())) = uint32(core1_0.Samples8)
			*(*uint32)(unsafe.Pointer(formatPropertiesVal.FieldByName("maxExtent").FieldByName("width").UnsafeAddr())) = uint32(11)
			*(*uint32)(unsafe.Pointer(formatPropertiesVal.FieldByName("maxExtent").FieldByName("height").UnsafeAddr())) = uint32(13)
			*(*uint32)(unsafe.Pointer(formatPropertiesVal.FieldByName("maxExtent").FieldByName("depth").UnsafeAddr())) = uint32(17)

			return core1_0.VKSuccess, nil
		})

	outData := core1_1.ImageFormatProperties2{}
	_, err := physicalDevice.InstanceScopedPhysicalDevice1_1().ImageFormatProperties2(core1_1.PhysicalDeviceImageFormatInfo2{
		Format: core1_0.FormatA2B10G10R10UnsignedIntPacked,
		Type:   core1_0.ImageType2D,
		Tiling: core1_0.ImageTilingOptimal,
		Usage:  core1_0.ImageUsageStorage,
		Flags:  core1_0.ImageCreateCubeCompatible,
	}, &outData)
	require.NoError(t, err)

	require.Equal(t, 1, outData.ImageFormatProperties.MaxMipLevels)
	require.Equal(t, 3, outData.ImageFormatProperties.MaxArrayLayers)
	require.Equal(t, 5, outData.ImageFormatProperties.MaxResourceSize)
	require.Equal(t, core1_0.Samples8, outData.ImageFormatProperties.SampleCounts)
	require.Equal(t, 11, outData.ImageFormatProperties.MaxExtent.Width)
	require.Equal(t, 13, outData.ImageFormatProperties.MaxExtent.Height)
	require.Equal(t, 17, outData.ImageFormatProperties.MaxExtent.Depth)
}

func TestVulkanPhysicalDevice_MemoryProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceMemoryProperties2(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pMemoryProperties *driver.VkPhysicalDeviceMemoryProperties2) {
			val := reflect.ValueOf(pMemoryProperties).Elem()

			require.Equal(t, uint64(1000059006), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2
			require.True(t, val.FieldByName("pNext").IsNil())

			memory := val.FieldByName("memoryProperties")
			*(*uint32)(unsafe.Pointer(memory.FieldByName("memoryTypeCount").UnsafeAddr())) = uint32(1)
			*(*uint32)(unsafe.Pointer(memory.FieldByName("memoryHeapCount").UnsafeAddr())) = uint32(1)

			memoryType := memory.FieldByName("memoryTypes").Index(0)
			*(*uint32)(unsafe.Pointer(memoryType.FieldByName("heapIndex").UnsafeAddr())) = uint32(3)
			*(*int32)(unsafe.Pointer(memoryType.FieldByName("propertyFlags").UnsafeAddr())) = int32(16) // VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT

			memoryHeap := memory.FieldByName("memoryHeaps").Index(0)
			*(*uint64)(unsafe.Pointer(memoryHeap.FieldByName("size").UnsafeAddr())) = uint64(99)
			*(*int32)(unsafe.Pointer(memoryHeap.FieldByName("flags").UnsafeAddr())) = int32(1) // VK_MEMORY_HEAP_DEVICE_LOCAL_BIT
		})

	outData := core1_1.PhysicalDeviceMemoryProperties2{}
	err := physicalDevice.InstanceScopedPhysicalDevice1_1().MemoryProperties2(&outData)
	require.NoError(t, err)
	require.Equal(t, []core1_0.MemoryType{
		{
			PropertyFlags: core1_0.MemoryPropertyLazilyAllocated,
			HeapIndex:     3,
		},
	}, outData.MemoryProperties.MemoryTypes)
	require.Equal(t, []core1_0.MemoryHeap{
		{
			Flags: core1_0.MemoryHeapDeviceLocal,
			Size:  99,
		},
	}, outData.MemoryProperties.MemoryHeaps)
}

func TestVulkanPhysicalDevice_Properties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	deviceUUID, err := uuid.NewUUID()
	require.NoError(t, err)

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties2(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties2) {
			val := reflect.ValueOf(pProperties).Elem()

			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
			require.True(t, val.FieldByName("pNext").IsNil())

			properties := val.FieldByName("properties")

			*(*uint32)(unsafe.Pointer(properties.FieldByName("apiVersion").UnsafeAddr())) = uint32(common.Vulkan1_1)
			*(*uint32)(unsafe.Pointer(properties.FieldByName("driverVersion").UnsafeAddr())) = uint32(common.CreateVersion(3, 2, 1))
			*(*uint32)(unsafe.Pointer(properties.FieldByName("vendorID").UnsafeAddr())) = uint32(3)
			*(*uint32)(unsafe.Pointer(properties.FieldByName("deviceID").UnsafeAddr())) = uint32(5)
			*(*uint32)(unsafe.Pointer(properties.FieldByName("deviceType").UnsafeAddr())) = uint32(2) // VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU
			deviceNamePtr := (*driver.Char)(unsafe.Pointer(properties.FieldByName("deviceName").UnsafeAddr()))
			deviceNameSlice := ([]driver.Char)(unsafe.Slice(deviceNamePtr, 256))
			deviceName := "Some Device"
			for i, r := range []byte(deviceName) {
				deviceNameSlice[i] = driver.Char(r)
			}
			deviceNameSlice[len(deviceName)] = 0

			uuidPtr := (*driver.Char)(unsafe.Pointer(properties.FieldByName("pipelineCacheUUID").UnsafeAddr()))
			uuidSlice := ([]driver.Char)(unsafe.Slice(uuidPtr, 16))
			uuid, err := deviceUUID.MarshalBinary()
			require.NoError(t, err)

			for i, b := range uuid {
				uuidSlice[i] = driver.Char(b)
			}

			limits := properties.FieldByName("limits")
			*(*uint32)(unsafe.Pointer(limits.FieldByName("maxUniformBufferRange").UnsafeAddr())) = uint32(7)
			*(*uint32)(unsafe.Pointer(limits.FieldByName("maxVertexInputBindingStride").UnsafeAddr())) = uint32(11)
			workGroupCount := limits.FieldByName("maxComputeWorkGroupCount")
			*(*uint32)(unsafe.Pointer(workGroupCount.Index(0).UnsafeAddr())) = uint32(13)
			*(*uint32)(unsafe.Pointer(workGroupCount.Index(1).UnsafeAddr())) = uint32(17)
			*(*uint32)(unsafe.Pointer(workGroupCount.Index(2).UnsafeAddr())) = uint32(19)
			*(*float32)(unsafe.Pointer(limits.FieldByName("maxInterpolationOffset").UnsafeAddr())) = float32(23)
			*(*driver.VkBool32)(unsafe.Pointer(limits.FieldByName("strictLines").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkDeviceSize)(unsafe.Pointer(limits.FieldByName("optimalBufferCopyRowPitchAlignment").UnsafeAddr())) = driver.VkDeviceSize(29)

			sparseProperties := properties.FieldByName("sparseProperties")
			*(*driver.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyStandard2DBlockShape").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyStandard2DMultisampleBlockShape").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyStandard3DBlockShape").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyAlignedMipSize").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyNonResidentStrict").UnsafeAddr())) = driver.VkBool32(1)
		})

	outData := core1_1.PhysicalDeviceProperties2{}
	err = physicalDevice.InstanceScopedPhysicalDevice1_1().Properties2(&outData)
	require.NoError(t, err)

	require.Equal(t, common.Vulkan1_1, outData.Properties.APIVersion)
	require.Equal(t, common.CreateVersion(3, 2, 1), outData.Properties.DriverVersion)
	require.Equal(t, uint32(3), outData.Properties.VendorID)
	require.Equal(t, uint32(5), outData.Properties.DeviceID)
	require.Equal(t, core1_0.PhysicalDeviceTypeDiscreteGPU, outData.Properties.DriverType)
	require.Equal(t, "Some Device", outData.Properties.DriverName)
	require.Equal(t, deviceUUID, outData.Properties.PipelineCacheUUID)

	require.Equal(t, 7, outData.Properties.Limits.MaxUniformBufferRange)
	require.Equal(t, 11, outData.Properties.Limits.MaxVertexInputBindingStride)
	require.Equal(t, [3]int{13, 17, 19}, outData.Properties.Limits.MaxComputeWorkGroupCount)
	require.Equal(t, float32(23), outData.Properties.Limits.MaxInterpolationOffset)
	require.True(t, outData.Properties.Limits.StrictLines)
	require.Equal(t, 29, outData.Properties.Limits.OptimalBufferCopyRowPitchAlignment)

	require.True(t, outData.Properties.SparseProperties.ResidencyStandard2DBlockShape)
	require.False(t, outData.Properties.SparseProperties.ResidencyStandard2DMultisampleBlockShape)
	require.True(t, outData.Properties.SparseProperties.ResidencyStandard3DBlockShape)
	require.False(t, outData.Properties.SparseProperties.ResidencyAlignedMipSize)
	require.True(t, outData.Properties.SparseProperties.ResidencyNonResidentStrict)
}

func TestVulkanPhysicalDevice_QueueFamilyProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceQueueFamilyProperties2(physicalDevice.Handle(), gomock.Not(gomock.Nil()), nil).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pQueueFamilyPropertyCount *driver.Uint32, pQueueFamilyProperties *driver.VkQueueFamilyProperties2) {
			*pQueueFamilyPropertyCount = 2
		})

	coreDriver.EXPECT().VkGetPhysicalDeviceQueueFamilyProperties2(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pQueueFamilyPropertyCount *driver.Uint32, pQueueFamilyProperties *driver.VkQueueFamilyProperties2) {
			require.Equal(t, driver.Uint32(2), *pQueueFamilyPropertyCount)

			propertySlice := ([]driver.VkQueueFamilyProperties2)(unsafe.Slice(pQueueFamilyProperties, 2))
			val := reflect.ValueOf(propertySlice)
			property := val.Index(0)

			require.Equal(t, uint64(1000059005), property.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2
			require.True(t, property.FieldByName("pNext").IsNil())

			queueFamily := property.FieldByName("queueFamilyProperties")
			*(*driver.VkQueueFlags)(unsafe.Pointer(queueFamily.FieldByName("queueFlags").UnsafeAddr())) = driver.VkQueueFlags(8) // VK_QUEUE_SPARSE_BINDING_BIT
			*(*uint32)(unsafe.Pointer(queueFamily.FieldByName("queueCount").UnsafeAddr())) = uint32(3)
			*(*uint32)(unsafe.Pointer(queueFamily.FieldByName("timestampValidBits").UnsafeAddr())) = uint32(5)

			propertyExtent := queueFamily.FieldByName("minImageTransferGranularity")
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("width").UnsafeAddr())) = uint32(7)
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("height").UnsafeAddr())) = uint32(11)
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("depth").UnsafeAddr())) = uint32(13)

			property = val.Index(1)
			require.Equal(t, uint64(1000059005), property.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2
			require.True(t, property.FieldByName("pNext").IsNil())

			queueFamily = property.FieldByName("queueFamilyProperties")
			*(*driver.VkQueueFlags)(unsafe.Pointer(queueFamily.FieldByName("queueFlags").UnsafeAddr())) = driver.VkQueueFlags(2) // VK_QUEUE_COMPUTE_BIT
			*(*uint32)(unsafe.Pointer(queueFamily.FieldByName("queueCount").UnsafeAddr())) = uint32(17)
			*(*uint32)(unsafe.Pointer(queueFamily.FieldByName("timestampValidBits").UnsafeAddr())) = uint32(19)

			propertyExtent = queueFamily.FieldByName("minImageTransferGranularity")
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("width").UnsafeAddr())) = uint32(23)
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("height").UnsafeAddr())) = uint32(29)
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("depth").UnsafeAddr())) = uint32(31)
		})

	outData, err := physicalDevice.InstanceScopedPhysicalDevice1_1().QueueFamilyProperties2(nil)
	require.NoError(t, err)

	require.Equal(t, []*core1_1.QueueFamilyProperties2{
		{
			QueueFamilyProperties: core1_0.QueueFamilyProperties{
				QueueFlags:         core1_0.QueueSparseBinding,
				QueueCount:         3,
				TimestampValidBits: 5,
				MinImageTransferGranularity: core1_0.Extent3D{
					Width:  7,
					Height: 11,
					Depth:  13,
				},
			},
		},
		{
			QueueFamilyProperties: core1_0.QueueFamilyProperties{
				QueueFlags:         core1_0.QueueCompute,
				QueueCount:         17,
				TimestampValidBits: 19,
				MinImageTransferGranularity: core1_0.Extent3D{
					Width:  23,
					Height: 29,
					Depth:  31,
				},
			},
		},
	}, outData)
}

func TestVulkanPhysicalDevice_SparseImageFormatProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceSparseImageFormatProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pFormatInfo *driver.VkPhysicalDeviceSparseImageFormatInfo2,
		pPropertyCount *driver.Uint32,
		pProperties *driver.VkSparseImageFormatProperties2) {

		val := reflect.ValueOf(pFormatInfo).Elem()
		require.Equal(t, uint64(1000059008), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(66), val.FieldByName("format").Uint())        // VK_FORMAT_A2B10G10R10_USCALED_PACK32
		require.Equal(t, uint64(2), val.FieldByName("_type").Uint())          // VK_IMAGE_TYPE_3D
		require.Equal(t, uint64(32), val.FieldByName("samples").Uint())       // VK_SAMPLE_COUNT_32_BIT
		require.Equal(t, uint64(0x00000008), val.FieldByName("usage").Uint()) // VK_IMAGE_USAGE_STORAGE_BIT
		require.Equal(t, uint64(1), val.FieldByName("tiling").Uint())         // VK_IMAGE_TILING_LINEAR

		*pPropertyCount = 1
	})

	coreDriver.EXPECT().VkGetPhysicalDeviceSparseImageFormatProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pFormatInfo *driver.VkPhysicalDeviceSparseImageFormatInfo2,
		pPropertyCount *driver.Uint32,
		pProperties *driver.VkSparseImageFormatProperties2) {

		val := reflect.ValueOf(pFormatInfo).Elem()
		require.Equal(t, uint64(1000059008), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(66), val.FieldByName("format").Uint())        // VK_FORMAT_A2B10G10R10_USCALED_PACK32
		require.Equal(t, uint64(2), val.FieldByName("_type").Uint())          // VK_IMAGE_TYPE_3D
		require.Equal(t, uint64(32), val.FieldByName("samples").Uint())       // VK_SAMPLE_COUNT_32_BIT
		require.Equal(t, uint64(0x00000008), val.FieldByName("usage").Uint()) // VK_IMAGE_USAGE_STORAGE_BIT
		require.Equal(t, uint64(1), val.FieldByName("tiling").Uint())         // VK_IMAGE_TILING_LINEAR

		require.Equal(t, driver.Uint32(1), *pPropertyCount)

		propertySlice := ([]driver.VkSparseImageFormatProperties2)(unsafe.Slice(pProperties, 1))
		outData := reflect.ValueOf(propertySlice)
		prop := outData.Index(0)
		require.Equal(t, uint64(1000059007), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2
		require.True(t, prop.FieldByName("pNext").IsNil())

		sparseProps := prop.FieldByName("properties")
		*(*uint32)(unsafe.Pointer(sparseProps.FieldByName("aspectMask").UnsafeAddr())) = uint32(1) // VK_IMAGE_ASPECT_COLOR_BIT
		*(*int32)(unsafe.Pointer(sparseProps.FieldByName("imageGranularity").FieldByName("width").UnsafeAddr())) = int32(1)
		*(*int32)(unsafe.Pointer(sparseProps.FieldByName("imageGranularity").FieldByName("height").UnsafeAddr())) = int32(3)
		*(*int32)(unsafe.Pointer(sparseProps.FieldByName("imageGranularity").FieldByName("depth").UnsafeAddr())) = int32(5)
		*(*uint32)(unsafe.Pointer(sparseProps.FieldByName("flags").UnsafeAddr())) = uint32(4) // VK_SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT
	})

	outData, err := physicalDevice.InstanceScopedPhysicalDevice1_1().SparseImageFormatProperties2(
		core1_1.PhysicalDeviceSparseImageFormatInfo2{
			Format:  core1_0.FormatA2B10G10R10UnsignedScaledPacked,
			Type:    core1_0.ImageType3D,
			Samples: core1_0.Samples32,
			Usage:   core1_0.ImageUsageStorage,
			Tiling:  core1_0.ImageTilingLinear,
		}, nil)
	require.NoError(t, err)
	require.Equal(t, []*core1_1.SparseImageFormatProperties2{
		{
			Properties: core1_0.SparseImageFormatProperties{
				AspectMask: core1_0.ImageAspectColor,
				ImageGranularity: core1_0.Extent3D{
					Width:  1,
					Height: 3,
					Depth:  5,
				},
				Flags: core1_0.SparseImageFormatNonstandardBlockSize,
			},
		},
	}, outData)
}

func TestPhysicalDeviceIDOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

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

	err = physicalDevice.InstanceScopedPhysicalDevice1_1().Properties2(
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
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

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
	err := physicalDevice.InstanceScopedPhysicalDevice1_1().Properties2(outData)
	require.NoError(t, err)

	require.Equal(t, uint32(3), outData.Properties.VendorID)
	require.Equal(t, 5, maintOutData.MaxPerSetDescriptors)
	require.Equal(t, 7, maintOutData.MaxMemoryAllocationSize)
}

func TestMultiviewPropertiesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

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

	err := physicalDevice.InstanceScopedPhysicalDevice1_1().Properties2(&properties)
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
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

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

	err := physicalDevice.InstanceScopedPhysicalDevice1_1().Properties2(properties)
	require.NoError(t, err)

	require.Equal(t, uint32(3), properties.Properties.VendorID)
	require.InDelta(t, 5.0, properties.Properties.Limits.LineWidthGranularity, 0.001)

	require.Equal(t, core1_1.PointClippingUserClipPlanesOnly, pointClipping.PointClippingBehavior)
}

func TestPhysicalDeviceProtectedMemoryOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

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

	err := physicalDevice.InstanceScopedPhysicalDevice1_1().Properties2(properties)
	require.NoError(t, err)

	require.Equal(t, uint32(3), properties.Properties.VendorID)
	require.InDelta(t, 5.0, properties.Properties.Limits.LineWidthGranularity, 0.001)

	require.True(t, protectedMemory.ProtectedNoFault)
}

func TestPhysicalDeviceSubgroupOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

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

	err := physicalDevice.InstanceScopedPhysicalDevice1_1().Properties2(properties)
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

func TestVulkanExtension_PhysicalDeviceFeatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

	coreDriver.EXPECT().VkGetPhysicalDeviceFeatures2(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pFeatures *driver.VkPhysicalDeviceFeatures2) {
			val := reflect.ValueOf(pFeatures).Elem()

			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
			require.True(t, val.FieldByName("pNext").IsNil())

			featureVal := val.FieldByName("features")
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("robustBufferAccess").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fullDrawIndexUint32").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("imageCubeArray").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("independentBlend").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("geometryShader").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("tessellationShader").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sampleRateShading").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("dualSrcBlend").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("logicOp").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("multiDrawIndirect").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("drawIndirectFirstInstance").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("depthClamp").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("depthBiasClamp").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fillModeNonSolid").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("depthBounds").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("wideLines").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("largePoints").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("alphaToOne").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("multiViewport").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("samplerAnisotropy").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("textureCompressionETC2").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("textureCompressionASTC_LDR").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("textureCompressionBC").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("occlusionQueryPrecise").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("pipelineStatisticsQuery").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("vertexPipelineStoresAndAtomics").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fragmentStoresAndAtomics").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderTessellationAndGeometryPointSize").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderImageGatherExtended").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageExtendedFormats").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageMultisample").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageReadWithoutFormat").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageWriteWithoutFormat").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderUniformBufferArrayDynamicIndexing").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderSampledImageArrayDynamicIndexing").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageBufferArrayDynamicIndexing").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageArrayDynamicIndexing").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderClipDistance").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderCullDistance").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderFloat64").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderInt64").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderInt16").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderResourceResidency").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderResourceMinLod").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseBinding").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyBuffer").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyImage2D").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyImage3D").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency2Samples").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency4Samples").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency8Samples").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency16Samples").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyAliased").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("variableMultisampleRate").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("inheritedQueries").UnsafeAddr())) = driver.VkBool32(1)

		})

	outData := &core1_1.PhysicalDeviceFeatures2{}
	err := physicalDevice.InstanceScopedPhysicalDevice1_1().Features2(outData)
	require.NoError(t, err)

	features := outData.Features
	require.NotNil(t, features)
	require.True(t, features.RobustBufferAccess)
	require.False(t, features.FullDrawIndexUint32)
	require.True(t, features.ImageCubeArray)
	require.False(t, features.IndependentBlend)
	require.True(t, features.GeometryShader)
	require.False(t, features.TessellationShader)
	require.True(t, features.SampleRateShading)
	require.False(t, features.DualSrcBlend)
	require.True(t, features.LogicOp)
	require.False(t, features.MultiDrawIndirect)
	require.True(t, features.DrawIndirectFirstInstance)
	require.False(t, features.DepthClamp)
	require.True(t, features.DepthBiasClamp)
	require.False(t, features.FillModeNonSolid)
	require.True(t, features.DepthBounds)
	require.False(t, features.WideLines)
	require.True(t, features.LargePoints)
	require.False(t, features.AlphaToOne)
	require.True(t, features.MultiViewport)
	require.False(t, features.SamplerAnisotropy)
	require.True(t, features.TextureCompressionEtc2)
	require.False(t, features.TextureCompressionAstcLdc)
	require.True(t, features.TextureCompressionBc)
	require.False(t, features.OcclusionQueryPrecise)
	require.True(t, features.PipelineStatisticsQuery)
	require.False(t, features.VertexPipelineStoresAndAtomics)
	require.True(t, features.FragmentStoresAndAtomics)
	require.False(t, features.ShaderTessellationAndGeometryPointSize)
	require.True(t, features.ShaderImageGatherExtended)
	require.False(t, features.ShaderStorageImageExtendedFormats)
	require.True(t, features.ShaderStorageImageMultisample)
	require.False(t, features.ShaderStorageImageReadWithoutFormat)
	require.True(t, features.ShaderStorageImageWriteWithoutFormat)
	require.False(t, features.ShaderUniformBufferArrayDynamicIndexing)
	require.True(t, features.ShaderSampledImageArrayDynamicIndexing)
	require.False(t, features.ShaderStorageBufferArrayDynamicIndexing)
	require.True(t, features.ShaderStorageImageArrayDynamicIndexing)
	require.False(t, features.ShaderClipDistance)
	require.True(t, features.ShaderCullDistance)
	require.False(t, features.ShaderFloat64)
	require.True(t, features.ShaderInt64)
	require.False(t, features.ShaderInt16)
	require.True(t, features.ShaderResourceResidency)
	require.False(t, features.ShaderResourceMinLod)
	require.True(t, features.SparseBinding)
	require.False(t, features.SparseResidencyBuffer)
	require.True(t, features.SparseResidencyImage2D)
	require.False(t, features.SparseResidencyImage3D)
	require.True(t, features.SparseResidency2Samples)
	require.False(t, features.SparseResidency4Samples)
	require.True(t, features.SparseResidency8Samples)
	require.False(t, features.SparseResidency16Samples)
	require.True(t, features.SparseResidencyAliased)
	require.False(t, features.VariableMultisampleRate)
	require.True(t, features.InheritedQueries)
}

func TestVulkanDevice_CreateDeviceWithFeatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil).AnyTimes()

	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))
	device := mocks.EasyMockDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkCreateDevice(physicalDevice.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice, pCreateInfo *driver.VkDeviceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDevice *driver.VkDevice) (common.VkResult, error) {
			*pDevice = device.Handle()

			v := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(3), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
			require.Equal(t, uint64(0), v.FieldByName("flags").Uint())
			require.Equal(t, uint64(2), v.FieldByName("queueCreateInfoCount").Uint())
			require.Equal(t, uint64(2), v.FieldByName("enabledExtensionCount").Uint())
			require.Equal(t, uint64(0), v.FieldByName("enabledLayerCount").Uint())
			require.True(t, v.FieldByName("pEnabledFeatures").IsNil())

			extensionNamePtr := (**driver.Char)(unsafe.Pointer(v.FieldByName("ppEnabledExtensionNames").Elem().UnsafeAddr()))
			extensionNameSlice := ([]*driver.Char)(unsafe.Slice(extensionNamePtr, 2))

			var extensionNames []string
			for _, extensionNameBytes := range extensionNameSlice {
				var extensionNameRunes []rune
				extensionNameByteSlice := ([]driver.Char)(unsafe.Slice(extensionNameBytes, 1<<30))
				for _, nameByte := range extensionNameByteSlice {
					if nameByte == 0 {
						break
					}

					extensionNameRunes = append(extensionNameRunes, rune(nameByte))
				}

				extensionNames = append(extensionNames, string(extensionNameRunes))
			}

			require.ElementsMatch(t, []string{"a", "b"}, extensionNames)

			require.True(t, v.FieldByName("ppEnabledLayerNames").IsNil())

			queueCreateInfoPtr := (*driver.VkDeviceQueueCreateInfo)(unsafe.Pointer(v.FieldByName("pQueueCreateInfos").Elem().UnsafeAddr()))
			queueCreateInfoSlice := ([]driver.VkDeviceQueueCreateInfo)(unsafe.Slice(queueCreateInfoPtr, 2))

			queueInfoV := reflect.ValueOf(queueCreateInfoSlice[0])
			require.Equal(t, uint64(2), queueInfoV.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
			require.True(t, queueInfoV.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), queueInfoV.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), queueInfoV.FieldByName("queueFamilyIndex").Uint())
			require.Equal(t, uint64(3), queueInfoV.FieldByName("queueCount").Uint())

			priorityPtr := (*float32)(unsafe.Pointer(queueInfoV.FieldByName("pQueuePriorities").Elem().UnsafeAddr()))
			prioritySlice := ([]float32)(unsafe.Slice(priorityPtr, 3))
			require.Equal(t, float32(3), prioritySlice[0])
			require.Equal(t, float32(5), prioritySlice[1])
			require.Equal(t, float32(7), prioritySlice[2])

			queueInfoV = reflect.ValueOf(queueCreateInfoSlice[1])
			require.Equal(t, uint64(2), queueInfoV.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
			require.True(t, queueInfoV.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), queueInfoV.FieldByName("flags").Uint())
			require.Equal(t, uint64(11), queueInfoV.FieldByName("queueFamilyIndex").Uint())
			require.Equal(t, uint64(1), queueInfoV.FieldByName("queueCount").Uint())

			priorityPtr = (*float32)(unsafe.Pointer(queueInfoV.FieldByName("pQueuePriorities").Elem().UnsafeAddr()))
			prioritySlice = ([]float32)(unsafe.Slice(priorityPtr, 1))
			require.Equal(t, float32(13), prioritySlice[0])

			nextPtr := (*driver.VkPhysicalDeviceFeatures2)(v.FieldByName("pNext").UnsafePointer())
			nextVal := reflect.ValueOf(nextPtr).Elem()
			require.Equal(t, uint64(1000059000), nextVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
			require.True(t, nextVal.FieldByName("pNext").IsNil())

			features := nextVal.FieldByName("features")
			require.Equal(t, uint64(1), features.FieldByName("textureCompressionETC2").Uint())
			require.Equal(t, uint64(1), features.FieldByName("depthBounds").Uint())
			require.Equal(t, uint64(0), features.FieldByName("samplerAnisotropy").Uint())

			return core1_0.VKSuccess, nil
		})

	options := core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueueFamilyIndex: 1,
				QueuePriorities:  []float32{3, 5, 7},
			},
			{
				QueueFamilyIndex: 11,
				QueuePriorities:  []float32{13},
			},
		},
		EnabledExtensionNames: []string{"a", "b"},
	}
	features := core1_1.PhysicalDeviceFeatures2{
		Features: core1_0.PhysicalDeviceFeatures{
			TextureCompressionEtc2: true,
			DepthBounds:            true,
		},
	}
	options.Next = features

	actualDevice, _, err := physicalDevice.CreateDevice(nil, options)
	require.NoError(t, err)
	require.NotNil(t, actualDevice)
	require.Equal(t, device.Handle(), actualDevice.Handle())
}

func TestDevice16BitStorageOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

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

	err := physicalDevice.InstanceScopedPhysicalDevice1_1().Features2(features)
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
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))
	mockDevice := mocks.EasyMockDevice(ctrl, coreDriver)

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
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

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

	err := physicalDevice.InstanceScopedPhysicalDevice1_1().Features2(&features)
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
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))
	mockDevice := mocks.EasyMockDevice(ctrl, coreDriver)

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
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

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

	err := physicalDevice.InstanceScopedPhysicalDevice1_1().Features2(&features)
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
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))
	mockDevice := mocks.EasyMockDevice(ctrl, coreDriver)

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
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

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

	err := physicalDevice.InstanceScopedPhysicalDevice1_1().Features2(
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
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))
	mockDevice := mocks.EasyMockDevice(ctrl, coreDriver)

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
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

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

	err := physicalDevice.InstanceScopedPhysicalDevice1_1().Features2(&features)
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
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))
	mockDevice := mocks.EasyMockDevice(ctrl, coreDriver)

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
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))

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

	err := physicalDevice.InstanceScopedPhysicalDevice1_1().Features2(&core1_1.PhysicalDeviceFeatures2{
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
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice1 := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))
	physicalDevice2 := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice3 := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

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
	device := dummies.EasyDummyDevice(coreDriver)

	mockMemory := mocks.EasyMockDeviceMemory(ctrl)

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

	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := core1_1.PromotePhysicalDevice(dummies.EasyDummyPhysicalDevice(coreDriver, instance))
	expectedDevice := mocks.EasyMockDevice(ctrl, coreDriver)

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
