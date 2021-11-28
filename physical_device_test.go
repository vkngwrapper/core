package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanPhysicalDevice_AvailableExtensions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	physicalDevice := mocks.EasyDummyPhysicalDevice(t, loader)

	driver.EXPECT().VkEnumerateDeviceExtensionProperties(physicalDevice.Handle(), nil, gomock.Not(nil), nil).DoAndReturn(
		func(physDevice core.VkPhysicalDevice, pLayerName *core.Char, pPropertyCount *core.Uint32, pProperties *core.VkExtensionProperties) (core.VkResult, error) {
			*pPropertyCount = core.Uint32(2)

			return core.VKSuccess, nil
		})
	driver.EXPECT().VkEnumerateDeviceExtensionProperties(physicalDevice.Handle(), nil, gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(physDevice core.VkPhysicalDevice, pLayerName *core.Char, pPropertyCount *core.Uint32, pProperties *core.VkExtensionProperties) (core.VkResult, error) {
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
	extensions, _, err := physicalDevice.AvailableExtensions()
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

func TestVulkanPhysicalDevice_QueueFamilyProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	physicalDevice := mocks.EasyDummyPhysicalDevice(t, loader)

	driver.EXPECT().VkGetPhysicalDeviceQueueFamilyProperties(physicalDevice.Handle(), gomock.Not(nil), nil).DoAndReturn(
		func(device core.VkPhysicalDevice, pPropertyCount *core.Uint32, pProperties *core.VkQueueFamilyProperties) {
			*pPropertyCount = 1
		})
	driver.EXPECT().VkGetPhysicalDeviceQueueFamilyProperties(physicalDevice.Handle(), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(device core.VkPhysicalDevice, pPropertyCount *core.Uint32, pProperties *core.VkQueueFamilyProperties) {
			*pPropertyCount = 1
			propertySlice := reflect.ValueOf(([]core.VkQueueFamilyProperties)(unsafe.Slice(pProperties, 3)))

			property := propertySlice.Index(0)
			*(*core.VkQueueFlags)(unsafe.Pointer(property.FieldByName("queueFlags").UnsafeAddr())) = core.VkQueueFlags(8)
			*(*uint32)(unsafe.Pointer(property.FieldByName("queueCount").UnsafeAddr())) = uint32(3)
			*(*uint32)(unsafe.Pointer(property.FieldByName("timestampValidBits").UnsafeAddr())) = uint32(5)

			propertyExtent := property.FieldByName("minImageTransferGranularity")
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("width").UnsafeAddr())) = uint32(7)
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("height").UnsafeAddr())) = uint32(11)
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("depth").UnsafeAddr())) = uint32(13)
		})

	queueProperties := physicalDevice.QueueFamilyProperties()
	require.Equal(t, uint32(3), queueProperties[0].QueueCount)
	require.Equal(t, uint32(5), queueProperties[0].TimestampValidBits)
	require.Equal(t, 7, queueProperties[0].MinImageTransferGranularity.Width)
	require.Equal(t, 11, queueProperties[0].MinImageTransferGranularity.Height)
	require.Equal(t, 13, queueProperties[0].MinImageTransferGranularity.Depth)
	require.Equal(t, common.QueueSparseBinding, queueProperties[0].Flags)
}

func TestVulkanPhysicalDevice_Properties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	physicalDevice := mocks.EasyDummyPhysicalDevice(t, loader)

	deviceUUID, err := uuid.NewUUID()
	require.NoError(t, err)

	driver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice.Handle(), gomock.Not(nil)).DoAndReturn(
		func(device core.VkPhysicalDevice, pProperties *core.VkPhysicalDeviceProperties) {
			pPropertySlice := reflect.ValueOf(unsafe.Slice(pProperties, 1))
			val := pPropertySlice.Index(0)

			*(*uint32)(unsafe.Pointer(val.FieldByName("apiVersion").UnsafeAddr())) = uint32(common.CreateVersion(1, 2, 3))
			*(*uint32)(unsafe.Pointer(val.FieldByName("driverVersion").UnsafeAddr())) = uint32(common.CreateVersion(3, 2, 1))
			*(*uint32)(unsafe.Pointer(val.FieldByName("vendorID").UnsafeAddr())) = uint32(3)
			*(*uint32)(unsafe.Pointer(val.FieldByName("deviceID").UnsafeAddr())) = uint32(5)
			*(*uint32)(unsafe.Pointer(val.FieldByName("deviceType").UnsafeAddr())) = uint32(2) // VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU
			deviceNamePtr := (*core.Char)(unsafe.Pointer(val.FieldByName("deviceName").UnsafeAddr()))
			deviceNameSlice := ([]core.Char)(unsafe.Slice(deviceNamePtr, 256))
			deviceName := "Some Device"
			for i, r := range []byte(deviceName) {
				deviceNameSlice[i] = core.Char(r)
			}
			deviceNameSlice[len(deviceName)] = 0

			uuidPtr := (*core.Char)(unsafe.Pointer(val.FieldByName("pipelineCacheUUID").UnsafeAddr()))
			uuidSlice := ([]core.Char)(unsafe.Slice(uuidPtr, 16))
			uuid, err := deviceUUID.MarshalBinary()
			require.NoError(t, err)

			for i, b := range uuid {
				uuidSlice[i] = core.Char(b)
			}

			limits := val.FieldByName("limits")
			*(*uint32)(unsafe.Pointer(limits.FieldByName("maxUniformBufferRange").UnsafeAddr())) = uint32(7)
			*(*uint32)(unsafe.Pointer(limits.FieldByName("maxVertexInputBindingStride").UnsafeAddr())) = uint32(11)
			workGroupCount := limits.FieldByName("maxComputeWorkGroupCount")
			*(*uint32)(unsafe.Pointer(workGroupCount.Index(0).UnsafeAddr())) = uint32(13)
			*(*uint32)(unsafe.Pointer(workGroupCount.Index(1).UnsafeAddr())) = uint32(17)
			*(*uint32)(unsafe.Pointer(workGroupCount.Index(2).UnsafeAddr())) = uint32(19)
			*(*float32)(unsafe.Pointer(limits.FieldByName("maxInterpolationOffset").UnsafeAddr())) = float32(23)
			*(*core.VkBool32)(unsafe.Pointer(limits.FieldByName("strictLines").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkDeviceSize)(unsafe.Pointer(limits.FieldByName("optimalBufferCopyRowPitchAlignment").UnsafeAddr())) = core.VkDeviceSize(29)

			sparseProperties := val.FieldByName("sparseProperties")
			*(*core.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyStandard2DBlockShape").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyStandard2DMultisampleBlockShape").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyStandard3DBlockShape").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyAlignedMipSize").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyNonResidentStrict").UnsafeAddr())) = core.VkBool32(1)
		})

	properties := physicalDevice.Properties()
	require.NotNil(t, properties)
	require.Equal(t, common.CreateVersion(1, 2, 3), properties.APIVersion)
	require.Equal(t, common.CreateVersion(3, 2, 1), properties.DriverVersion)
	require.Equal(t, uint32(3), properties.VendorID)
	require.Equal(t, uint32(5), properties.DeviceID)
	require.Equal(t, common.DeviceDiscreteGPU, properties.Type)
	require.Equal(t, "Some Device", properties.Name)
	require.Equal(t, deviceUUID, properties.PipelineCacheUUID)

	require.Equal(t, 7, properties.Limits.MaxUniformBufferRange)
	require.Equal(t, 11, properties.Limits.MaxVertexInputBindingStride)
	require.Equal(t, [3]int{13, 17, 19}, properties.Limits.MaxComputeWorkGroupCount)
	require.Equal(t, float32(23), properties.Limits.MaxInterpolationOffset)
	require.True(t, properties.Limits.StrictLines)
	require.Equal(t, 29, properties.Limits.OptimalBufferCopyRowPitchAlignment)

	require.True(t, properties.SparseProperties.ResidencyStandard2DBlockShape)
	require.False(t, properties.SparseProperties.ResidencyStandard2DMultisampleBlockShape)
	require.True(t, properties.SparseProperties.ResidencyStandard3DBlockShape)
	require.False(t, properties.SparseProperties.ResidencyAlignedMipSize)
	require.True(t, properties.SparseProperties.ResidencyNonResidentStrict)
}

func TestVulkanPhysicalDevice_Features(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	physicalDevice := mocks.EasyDummyPhysicalDevice(t, loader)

	driver.EXPECT().VkGetPhysicalDeviceFeatures(physicalDevice.Handle(), gomock.Not(nil)).DoAndReturn(
		func(physicalDevice core.VkPhysicalDevice, pFeatures *core.VkPhysicalDeviceFeatures) {
			featureSlice := reflect.ValueOf(unsafe.Slice(pFeatures, 1))
			val := featureSlice.Index(0)

			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("robustBufferAccess").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("fullDrawIndexUint32").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("imageCubeArray").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("independentBlend").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("geometryShader").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("tessellationShader").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("sampleRateShading").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("dualSrcBlend").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("logicOp").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("multiDrawIndirect").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("drawIndirectFirstInstance").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("depthClamp").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("depthBiasClamp").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("fillModeNonSolid").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("depthBounds").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("wideLines").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("largePoints").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("alphaToOne").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("multiViewport").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("samplerAnisotropy").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("textureCompressionETC2").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("textureCompressionASTC_LDR").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("textureCompressionBC").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("occlusionQueryPrecise").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("pipelineStatisticsQuery").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("vertexPipelineStoresAndAtomics").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("fragmentStoresAndAtomics").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("shaderTessellationAndGeometryPointSize").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("shaderImageGatherExtended").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageExtendedFormats").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageMultisample").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageReadWithoutFormat").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageWriteWithoutFormat").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformBufferArrayDynamicIndexing").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSampledImageArrayDynamicIndexing").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageBufferArrayDynamicIndexing").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageArrayDynamicIndexing").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("shaderClipDistance").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("shaderCullDistance").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("shaderFloat64").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInt64").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInt16").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("shaderResourceResidency").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("shaderResourceMinLod").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("sparseBinding").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("sparseResidencyBuffer").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("sparseResidencyImage2D").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("sparseResidencyImage3D").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("sparseResidency2Samples").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("sparseResidency4Samples").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("sparseResidency8Samples").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("sparseResidency16Samples").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("sparseResidencyAliased").UnsafeAddr())) = core.VkBool32(1)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("variableMultisampleRate").UnsafeAddr())) = core.VkBool32(0)
			*(*core.VkBool32)(unsafe.Pointer(val.FieldByName("inheritedQueries").UnsafeAddr())) = core.VkBool32(1)
		})

	features := physicalDevice.Features()
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

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	physicalDevice := mocks.EasyDummyPhysicalDevice(t, loader)

	driver.EXPECT().VkGetPhysicalDeviceFormatProperties(physicalDevice.Handle(),
		core.VkFormat(57), // VK_FORMAT_A8B8G8R8_SRGB_PACK32
		gomock.Not(nil)).DoAndReturn(
		func(device core.VkPhysicalDevice, format core.VkFormat, pFormatProperties *core.VkFormatProperties) {
			val := reflect.ValueOf(pFormatProperties).Elem()

			*(*uint32)(unsafe.Pointer(val.FieldByName("optimalTilingFeatures").UnsafeAddr())) = uint32(0x00000100) // VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT
			*(*uint32)(unsafe.Pointer(val.FieldByName("linearTilingFeatures").UnsafeAddr())) = uint32(0x00400000)  // VK_FORMAT_FEATURE_DISJOINT_BIT
			*(*uint32)(unsafe.Pointer(val.FieldByName("bufferFeatures").UnsafeAddr())) = uint32(0x00040000)        // VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT
		})

	props := physicalDevice.FormatProperties(common.FormatA8B8G8R8SRGB)
	require.NotNil(t, props)
	require.Equal(t, common.FormatFeatureColorAttachmentBlend, props.OptimalTilingFeatures)
	require.Equal(t, common.FormatFeatureDisjoint, props.LinearTilingFeatures)
	require.Equal(t, common.FormatFeatureSampledImageYcbcrConversionLinearFilter, props.BufferFeatures)
}

func TestVulkanPhysicalDevice_MemoryProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	physicalDevice := mocks.EasyDummyPhysicalDevice(t, loader)

	driver.EXPECT().VkGetPhysicalDeviceMemoryProperties(physicalDevice.Handle(), gomock.Not(nil)).DoAndReturn(
		func(physicalDevice core.VkPhysicalDevice, pProperties *core.VkPhysicalDeviceMemoryProperties) {
			propertySlice := reflect.ValueOf(unsafe.Slice(pProperties, 1))
			val := propertySlice.Index(0)
			*(*uint32)(unsafe.Pointer(val.FieldByName("memoryTypeCount").UnsafeAddr())) = uint32(1)
			*(*uint32)(unsafe.Pointer(val.FieldByName("memoryHeapCount").UnsafeAddr())) = uint32(1)

			memoryType := val.FieldByName("memoryTypes").Index(0)
			*(*uint32)(unsafe.Pointer(memoryType.FieldByName("heapIndex").UnsafeAddr())) = uint32(3)
			*(*int32)(unsafe.Pointer(memoryType.FieldByName("propertyFlags").UnsafeAddr())) = int32(16) // VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT

			memoryHeap := val.FieldByName("memoryHeaps").Index(0)
			*(*uint64)(unsafe.Pointer(memoryHeap.FieldByName("size").UnsafeAddr())) = uint64(99)
			*(*int32)(unsafe.Pointer(memoryHeap.FieldByName("flags").UnsafeAddr())) = int32(2) // VK_MEMORY_HEAP_MULTI_INSTANCE_BIT
		})

	memoryProps := physicalDevice.MemoryProperties()
	require.NotNil(t, memoryProps)
	require.Len(t, memoryProps.MemoryTypes, 1)
	require.Len(t, memoryProps.MemoryHeaps, 1)

	require.Equal(t, 3, memoryProps.MemoryTypes[0].HeapIndex)
	require.Equal(t, core.MemoryLazilyAllocated, memoryProps.MemoryTypes[0].Properties)

	require.Equal(t, uint64(99), memoryProps.MemoryHeaps[0].Size)
	require.Equal(t, core.HeapMultiInstance, memoryProps.MemoryHeaps[0].Flags)
}
