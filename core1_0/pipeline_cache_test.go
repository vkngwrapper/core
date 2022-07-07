package core1_0_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
	mock_driver "github.com/vkngwrapper/core/driver/mocks"
	internal_mocks "github.com/vkngwrapper/core/internal/dummies"
	"github.com/vkngwrapper/core/mocks"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanLoader1_0_CreatePipelineCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := internal_mocks.EasyDummyDevice(mockDriver)
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
	device := mocks.EasyMockDevice(ctrl, mockDriver)
	pipelineCache := internal_mocks.EasyDummyPipelineCache(mockDriver, device)

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
	device := mocks.EasyMockDevice(ctrl, mockDriver)
	pipelineCache := internal_mocks.EasyDummyPipelineCache(mockDriver, device)

	srcPipeline1 := mocks.EasyMockPipelineCache(ctrl)
	srcPipeline2 := mocks.EasyMockPipelineCache(ctrl)
	srcPipeline3 := mocks.EasyMockPipelineCache(ctrl)

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
