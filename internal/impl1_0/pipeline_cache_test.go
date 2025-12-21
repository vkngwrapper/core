package impl1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreatePipelineCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})

	pipelineCacheHandle := mocks.NewFakePipelineCache()

	mockDriver.EXPECT().VkCreatePipelineCache(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice,
			pCreateInfo *driver.VkPipelineCacheCreateInfo,
			pAllocator *driver.VkAllocationCallbacks,
			pPipelineCache *driver.VkPipelineCache) (common.VkResult, error) {
			*pPipelineCache = pipelineCacheHandle

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(17), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_CACHE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(4), val.FieldByName("initialDataSize").Uint())

			dataPtr := (*byte)(unsafe.Pointer(val.FieldByName("pInitialData").Pointer()))
			dataSlice := ([]byte)(unsafe.Slice(dataPtr, 4))

			require.Equal(t, []byte{1, 3, 5, 7}, dataSlice)

			return core1_0.VKSuccess, nil
		})

	pipelineCache, _, err := device.CreatePipelineCache(nil, core1_0.PipelineCacheCreateInfo{
		Flags:       0,
		InitialData: []byte{1, 3, 5, 7},
	})
	require.NoError(t, err)
	require.NotNil(t, pipelineCache)
	require.Equal(t, pipelineCacheHandle, pipelineCache.Handle())
}

func TestVulkanPipelineCache_CacheData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks1_0.EasyMockDevice(ctrl, mockDriver)
	builder := impl1_0.DeviceObjectBuilderImpl{}
	pipelineCache := builder.CreatePipelineCacheObject(mockDriver, device.Handle(), mocks.NewFakePipelineCache(), common.Vulkan1_0)

	mockDriver.EXPECT().VkGetPipelineCacheData(device.Handle(), pipelineCache.Handle(), gomock.Not(nil), unsafe.Pointer(nil)).DoAndReturn(
		func(device driver.VkDevice, pipelineCache driver.VkPipelineCache, pSize *driver.Size, pCacheData unsafe.Pointer) (common.VkResult, error) {
			*pSize = 8
			return core1_0.VKSuccess, nil
		})
	mockDriver.EXPECT().VkGetPipelineCacheData(device.Handle(), pipelineCache.Handle(), gomock.Not(nil), gomock.Not(unsafe.Pointer(nil))).DoAndReturn(
		func(device driver.VkDevice, pipelineCache driver.VkPipelineCache, pSize *driver.Size, pCacheData unsafe.Pointer) (common.VkResult, error) {
			require.Equal(t, driver.Size(8), *pSize)
			bytes := ([]byte)(unsafe.Slice((*byte)(pCacheData), 8))
			copy(bytes, []byte{1, 1, 2, 3, 5, 8, 13, 21})
			return core1_0.VKSuccess, nil
		})

	data, _, err := pipelineCache.CacheData()
	require.NoError(t, err)
	require.Equal(t, []byte{1, 1, 2, 3, 5, 8, 13, 21}, data)
}

func TestVulkanPipelineCache_MergePipelineCaches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks1_0.EasyMockDevice(ctrl, mockDriver)
	builder := impl1_0.DeviceObjectBuilderImpl{}
	pipelineCache := builder.CreatePipelineCacheObject(mockDriver, device.Handle(), mocks.NewFakePipelineCache(), common.Vulkan1_0)

	srcPipeline1 := mocks1_0.EasyMockPipelineCache(ctrl)
	srcPipeline2 := mocks1_0.EasyMockPipelineCache(ctrl)
	srcPipeline3 := mocks1_0.EasyMockPipelineCache(ctrl)

	mockDriver.EXPECT().VkMergePipelineCaches(device.Handle(), pipelineCache.Handle(), driver.Uint32(3), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, dstCache driver.VkPipelineCache, srcCacheCount driver.Uint32, pSrcCaches *driver.VkPipelineCache) (common.VkResult, error) {
			cacheSlice := unsafe.Slice(pSrcCaches, 3)

			require.Equal(t, srcPipeline1.Handle(), cacheSlice[0])
			require.Equal(t, srcPipeline2.Handle(), cacheSlice[1])
			require.Equal(t, srcPipeline3.Handle(), cacheSlice[2])

			return core1_0.VKSuccess, nil
		})

	_, err := pipelineCache.MergePipelineCaches([]core1_0.PipelineCache{
		srcPipeline1, srcPipeline2, srcPipeline3,
	})
	require.NoError(t, err)
}
