package impl1_1_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
	"go.uber.org/mock/gomock"
)

func TestVulkanPhysicalDevice_PhysicalDeviceExternalFenceProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceExternalFenceProperties(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pExternalFenceInfo *loader.VkPhysicalDeviceExternalFenceInfo,
		pExternalFenceProperties *loader.VkExternalFenceProperties,
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
	err := driver.GetPhysicalDeviceExternalFenceProperties(
		physicalDevice,
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

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceExternalBufferProperties(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice, pExternalBufferInfo *loader.VkPhysicalDeviceExternalBufferInfo, pExternalBufferProperties *loader.VkExternalBufferProperties) {
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
	err := driver.GetPhysicalDeviceExternalBufferProperties(
		physicalDevice,
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

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceExternalSemaphoreProperties(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice,
			pExternalSemaphoreInfo *loader.VkPhysicalDeviceExternalSemaphoreInfo,
			pExternalSemaphoreProperties *loader.VkExternalSemaphoreProperties,
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
	err := driver.GetPhysicalDeviceExternalSemaphoreProperties(
		physicalDevice,
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

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFeatures2(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
			pFeatures *loader.VkPhysicalDeviceFeatures2) {
			val := reflect.ValueOf(pFeatures).Elem()

			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
			require.True(t, val.FieldByName("pNext").IsNil())

			featureVal := val.FieldByName("features")
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("robustBufferAccess").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fullDrawIndexUint32").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("imageCubeArray").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("independentBlend").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("geometryShader").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("tessellationShader").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sampleRateShading").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("dualSrcBlend").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("logicOp").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("multiDrawIndirect").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("drawIndirectFirstInstance").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("depthClamp").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("depthBiasClamp").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fillModeNonSolid").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("depthBounds").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("wideLines").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("largePoints").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("alphaToOne").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("multiViewport").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("samplerAnisotropy").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("textureCompressionETC2").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("textureCompressionASTC_LDR").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("textureCompressionBC").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("occlusionQueryPrecise").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("pipelineStatisticsQuery").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("vertexPipelineStoresAndAtomics").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fragmentStoresAndAtomics").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderTessellationAndGeometryPointSize").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderImageGatherExtended").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageExtendedFormats").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageMultisample").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageReadWithoutFormat").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageWriteWithoutFormat").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderUniformBufferArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderSampledImageArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageBufferArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderClipDistance").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderCullDistance").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderFloat64").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderInt64").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderInt16").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderResourceResidency").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderResourceMinLod").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseBinding").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyBuffer").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyImage2D").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyImage3D").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency2Samples").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency4Samples").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency8Samples").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency16Samples").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyAliased").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("variableMultisampleRate").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("inheritedQueries").UnsafeAddr())) = loader.VkBool32(1)

		})

	outData := &core1_1.PhysicalDeviceFeatures2{}
	err := driver.GetPhysicalDeviceFeatures2(physicalDevice, outData)
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

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFormatProperties2(
		physicalDevice.Handle(),
		loader.VkFormat(64), // VK_FORMAT_A2B10G10R10_UNORM_PACK32
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		format loader.VkFormat,
		pFormatProperties *loader.VkFormatProperties2) {

		val := reflect.ValueOf(pFormatProperties).Elem()
		require.Equal(t, uint64(1000059002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2
		require.True(t, val.FieldByName("pNext").IsNil())

		properties := val.FieldByName("formatProperties")
		*(*uint32)(unsafe.Pointer(properties.FieldByName("optimalTilingFeatures").UnsafeAddr())) = uint32(0x00000100) // VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT
		*(*uint32)(unsafe.Pointer(properties.FieldByName("linearTilingFeatures").UnsafeAddr())) = uint32(0x00000400)  // VK_FORMAT_FEATURE_BLIT_SRC_BIT
		*(*uint32)(unsafe.Pointer(properties.FieldByName("bufferFeatures").UnsafeAddr())) = uint32(0x00000010)        // VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT
	})

	outData := core1_1.FormatProperties2{}
	err := driver.GetPhysicalDeviceFormatProperties2(
		physicalDevice,
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

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceImageFormatProperties2(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
			pImageFormatInfo *loader.VkPhysicalDeviceImageFormatInfo2,
			pImageFormatProperties *loader.VkImageFormatProperties2,
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
	_, err := driver.GetPhysicalDeviceImageFormatProperties2(physicalDevice, core1_1.PhysicalDeviceImageFormatInfo2{
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

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceMemoryProperties2(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice loader.VkPhysicalDevice, pMemoryProperties *loader.VkPhysicalDeviceMemoryProperties2) {
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
	err := driver.GetPhysicalDeviceMemoryProperties2(physicalDevice, &outData)
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

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, coreLoader)

	deviceUUID, err := uuid.NewUUID()
	require.NoError(t, err)

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties2(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties2) {
			val := reflect.ValueOf(pProperties).Elem()

			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
			require.True(t, val.FieldByName("pNext").IsNil())

			properties := val.FieldByName("properties")

			*(*uint32)(unsafe.Pointer(properties.FieldByName("apiVersion").UnsafeAddr())) = uint32(common.Vulkan1_1)
			*(*uint32)(unsafe.Pointer(properties.FieldByName("driverVersion").UnsafeAddr())) = uint32(common.CreateVersion(3, 2, 1))
			*(*uint32)(unsafe.Pointer(properties.FieldByName("vendorID").UnsafeAddr())) = uint32(3)
			*(*uint32)(unsafe.Pointer(properties.FieldByName("deviceID").UnsafeAddr())) = uint32(5)
			*(*uint32)(unsafe.Pointer(properties.FieldByName("deviceType").UnsafeAddr())) = uint32(2) // VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU
			deviceNamePtr := (*loader.Char)(unsafe.Pointer(properties.FieldByName("deviceName").UnsafeAddr()))
			deviceNameSlice := ([]loader.Char)(unsafe.Slice(deviceNamePtr, 256))
			deviceName := "Some Device"
			for i, r := range []byte(deviceName) {
				deviceNameSlice[i] = loader.Char(r)
			}
			deviceNameSlice[len(deviceName)] = 0

			uuidPtr := (*loader.Char)(unsafe.Pointer(properties.FieldByName("pipelineCacheUUID").UnsafeAddr()))
			uuidSlice := ([]loader.Char)(unsafe.Slice(uuidPtr, 16))
			uuid, err := deviceUUID.MarshalBinary()
			require.NoError(t, err)

			for i, b := range uuid {
				uuidSlice[i] = loader.Char(b)
			}

			limits := properties.FieldByName("limits")
			*(*uint32)(unsafe.Pointer(limits.FieldByName("maxUniformBufferRange").UnsafeAddr())) = uint32(7)
			*(*uint32)(unsafe.Pointer(limits.FieldByName("maxVertexInputBindingStride").UnsafeAddr())) = uint32(11)
			workGroupCount := limits.FieldByName("maxComputeWorkGroupCount")
			*(*uint32)(unsafe.Pointer(workGroupCount.Index(0).UnsafeAddr())) = uint32(13)
			*(*uint32)(unsafe.Pointer(workGroupCount.Index(1).UnsafeAddr())) = uint32(17)
			*(*uint32)(unsafe.Pointer(workGroupCount.Index(2).UnsafeAddr())) = uint32(19)
			*(*float32)(unsafe.Pointer(limits.FieldByName("maxInterpolationOffset").UnsafeAddr())) = float32(23)
			*(*loader.VkBool32)(unsafe.Pointer(limits.FieldByName("strictLines").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkDeviceSize)(unsafe.Pointer(limits.FieldByName("optimalBufferCopyRowPitchAlignment").UnsafeAddr())) = loader.VkDeviceSize(29)

			sparseProperties := properties.FieldByName("sparseProperties")
			*(*loader.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyStandard2DBlockShape").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyStandard2DMultisampleBlockShape").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyStandard3DBlockShape").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyAlignedMipSize").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyNonResidentStrict").UnsafeAddr())) = loader.VkBool32(1)
		})

	outData := core1_1.PhysicalDeviceProperties2{}
	err = driver.GetPhysicalDeviceProperties2(physicalDevice, &outData)
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

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceQueueFamilyProperties2(physicalDevice.Handle(), gomock.Not(gomock.Nil()), nil).
		DoAndReturn(func(physicalDevice loader.VkPhysicalDevice, pQueueFamilyPropertyCount *loader.Uint32, pQueueFamilyProperties *loader.VkQueueFamilyProperties2) {
			*pQueueFamilyPropertyCount = 2
		})

	coreLoader.EXPECT().VkGetPhysicalDeviceQueueFamilyProperties2(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice loader.VkPhysicalDevice, pQueueFamilyPropertyCount *loader.Uint32, pQueueFamilyProperties *loader.VkQueueFamilyProperties2) {
			require.Equal(t, loader.Uint32(2), *pQueueFamilyPropertyCount)

			propertySlice := ([]loader.VkQueueFamilyProperties2)(unsafe.Slice(pQueueFamilyProperties, 2))
			val := reflect.ValueOf(propertySlice)
			property := val.Index(0)

			require.Equal(t, uint64(1000059005), property.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2
			require.True(t, property.FieldByName("pNext").IsNil())

			queueFamily := property.FieldByName("queueFamilyProperties")
			*(*loader.VkQueueFlags)(unsafe.Pointer(queueFamily.FieldByName("queueFlags").UnsafeAddr())) = loader.VkQueueFlags(8) // VK_QUEUE_SPARSE_BINDING_BIT
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
			*(*loader.VkQueueFlags)(unsafe.Pointer(queueFamily.FieldByName("queueFlags").UnsafeAddr())) = loader.VkQueueFlags(2) // VK_QUEUE_COMPUTE_BIT
			*(*uint32)(unsafe.Pointer(queueFamily.FieldByName("queueCount").UnsafeAddr())) = uint32(17)
			*(*uint32)(unsafe.Pointer(queueFamily.FieldByName("timestampValidBits").UnsafeAddr())) = uint32(19)

			propertyExtent = queueFamily.FieldByName("minImageTransferGranularity")
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("width").UnsafeAddr())) = uint32(23)
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("height").UnsafeAddr())) = uint32(29)
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("depth").UnsafeAddr())) = uint32(31)
		})

	outData, err := driver.GetPhysicalDeviceQueueFamilyProperties2(physicalDevice, nil)
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

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceSparseImageFormatProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFormatInfo *loader.VkPhysicalDeviceSparseImageFormatInfo2,
		pPropertyCount *loader.Uint32,
		pProperties *loader.VkSparseImageFormatProperties2) {

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

	coreLoader.EXPECT().VkGetPhysicalDeviceSparseImageFormatProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFormatInfo *loader.VkPhysicalDeviceSparseImageFormatInfo2,
		pPropertyCount *loader.Uint32,
		pProperties *loader.VkSparseImageFormatProperties2) {

		val := reflect.ValueOf(pFormatInfo).Elem()
		require.Equal(t, uint64(1000059008), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(66), val.FieldByName("format").Uint())        // VK_FORMAT_A2B10G10R10_USCALED_PACK32
		require.Equal(t, uint64(2), val.FieldByName("_type").Uint())          // VK_IMAGE_TYPE_3D
		require.Equal(t, uint64(32), val.FieldByName("samples").Uint())       // VK_SAMPLE_COUNT_32_BIT
		require.Equal(t, uint64(0x00000008), val.FieldByName("usage").Uint()) // VK_IMAGE_USAGE_STORAGE_BIT
		require.Equal(t, uint64(1), val.FieldByName("tiling").Uint())         // VK_IMAGE_TILING_LINEAR

		require.Equal(t, loader.Uint32(1), *pPropertyCount)

		propertySlice := ([]loader.VkSparseImageFormatProperties2)(unsafe.Slice(pProperties, 1))
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

	outData, err := driver.GetPhysicalDeviceSparseImageFormatProperties2(
		physicalDevice,
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

func TestVulkanExtension_PhysicalDeviceFeatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkGetPhysicalDeviceFeatures2(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice loader.VkPhysicalDevice, pFeatures *loader.VkPhysicalDeviceFeatures2) {
			val := reflect.ValueOf(pFeatures).Elem()

			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
			require.True(t, val.FieldByName("pNext").IsNil())

			featureVal := val.FieldByName("features")
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("robustBufferAccess").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fullDrawIndexUint32").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("imageCubeArray").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("independentBlend").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("geometryShader").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("tessellationShader").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sampleRateShading").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("dualSrcBlend").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("logicOp").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("multiDrawIndirect").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("drawIndirectFirstInstance").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("depthClamp").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("depthBiasClamp").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fillModeNonSolid").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("depthBounds").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("wideLines").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("largePoints").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("alphaToOne").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("multiViewport").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("samplerAnisotropy").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("textureCompressionETC2").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("textureCompressionASTC_LDR").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("textureCompressionBC").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("occlusionQueryPrecise").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("pipelineStatisticsQuery").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("vertexPipelineStoresAndAtomics").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fragmentStoresAndAtomics").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderTessellationAndGeometryPointSize").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderImageGatherExtended").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageExtendedFormats").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageMultisample").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageReadWithoutFormat").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageWriteWithoutFormat").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderUniformBufferArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderSampledImageArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageBufferArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderClipDistance").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderCullDistance").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderFloat64").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderInt64").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderInt16").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderResourceResidency").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderResourceMinLod").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseBinding").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyBuffer").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyImage2D").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyImage3D").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency2Samples").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency4Samples").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency8Samples").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency16Samples").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyAliased").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("variableMultisampleRate").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("inheritedQueries").UnsafeAddr())) = loader.VkBool32(1)

		})

	outData := &core1_1.PhysicalDeviceFeatures2{}
	err := driver.GetPhysicalDeviceFeatures2(physicalDevice, outData)
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

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(physicalDevice.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice, pCreateInfo *loader.VkDeviceCreateInfo, pAllocator *loader.VkAllocationCallbacks, pDevice *loader.VkDevice) (common.VkResult, error) {
			*pDevice = device.Handle()

			v := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(3), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
			require.Equal(t, uint64(0), v.FieldByName("flags").Uint())
			require.Equal(t, uint64(2), v.FieldByName("queueCreateInfoCount").Uint())
			require.Equal(t, uint64(2), v.FieldByName("enabledExtensionCount").Uint())
			require.Equal(t, uint64(0), v.FieldByName("enabledLayerCount").Uint())
			require.True(t, v.FieldByName("pEnabledFeatures").IsNil())

			extensionNamePtr := (**loader.Char)(unsafe.Pointer(v.FieldByName("ppEnabledExtensionNames").Elem().UnsafeAddr()))
			extensionNameSlice := ([]*loader.Char)(unsafe.Slice(extensionNamePtr, 2))

			var extensionNames []string
			for _, extensionNameBytes := range extensionNameSlice {
				var extensionNameRunes []rune
				extensionNameByteSlice := ([]loader.Char)(unsafe.Slice(extensionNameBytes, 1<<30))
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

			queueCreateInfoPtr := (*loader.VkDeviceQueueCreateInfo)(unsafe.Pointer(v.FieldByName("pQueueCreateInfos").Elem().UnsafeAddr()))
			queueCreateInfoSlice := ([]loader.VkDeviceQueueCreateInfo)(unsafe.Slice(queueCreateInfoPtr, 2))

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

			nextPtr := (*loader.VkPhysicalDeviceFeatures2)(v.FieldByName("pNext").UnsafePointer())
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

	actualDevice, _, err := driver.CreateDevice(physicalDevice, nil, options)
	require.NoError(t, err)
	require.NotNil(t, actualDevice)
	require.Equal(t, device.Handle(), actualDevice.Device().Handle())
}
