package core1_0_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/internal/universal"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanLoader1_0_CreatePipelineCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := universal.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	pipelineCacheHandle := mocks.NewFakePipelineCache()

	mockDriver.EXPECT().VkCreatePipelineCache(mocks.Exactly(device.Handle()), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice,
			pCreateInfo *driver.VkPipelineCacheCreateInfo,
			pAllocator *driver.VkAllocationCallbacks,
			pPipelineCache *driver.VkPipelineCache) (common.VkResult, error) {
			*pPipelineCache = pipelineCacheHandle

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(17), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_CACHE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT_EXT
			require.Equal(t, uint64(4), val.FieldByName("initialDataSize").Uint())

			dataPtr := (*byte)(unsafe.Pointer(val.FieldByName("pInitialData").Pointer()))
			dataSlice := ([]byte)(unsafe.Slice(dataPtr, 4))

			require.Equal(t, []byte{1, 3, 5, 7}, dataSlice)

			return common.VKSuccess, nil
		})

	pipelineCache, _, err := loader.CreatePipelineCache(device, nil, &core.PipelineCacheOptions{
		Flags:       core.PipelineCacheExternallySynchronized,
		InitialData: []byte{1, 3, 5, 7},
	})
	require.NoError(t, err)
	require.NotNil(t, pipelineCache)
	require.Same(t, pipelineCacheHandle, pipelineCache.Handle())
}

func TestVulkanPipelineCache_CacheData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := universal.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	pipelineCache := mocks.EasyDummyPipelineCache(t, device, loader)

	mockDriver.EXPECT().VkGetPipelineCacheData(mocks.Exactly(device.Handle()), mocks.Exactly(pipelineCache.Handle()), gomock.Not(nil), unsafe.Pointer(nil)).DoAndReturn(
		func(device driver.VkDevice, pipelineCache driver.VkPipelineCache, pSize *driver.Size, pCacheData unsafe.Pointer) (common.VkResult, error) {
			*pSize = 8
			return common.VKSuccess, nil
		})
	mockDriver.EXPECT().VkGetPipelineCacheData(mocks.Exactly(device.Handle()), mocks.Exactly(pipelineCache.Handle()), gomock.Not(nil), gomock.Not(unsafe.Pointer(nil))).DoAndReturn(
		func(device driver.VkDevice, pipelineCache driver.VkPipelineCache, pSize *driver.Size, pCacheData unsafe.Pointer) (common.VkResult, error) {
			require.Equal(t, driver.Size(8), *pSize)
			bytes := ([]byte)(unsafe.Slice((*byte)(pCacheData), 8))
			copy(bytes, []byte{1, 1, 2, 3, 5, 8, 13, 21})
			return common.VKSuccess, nil
		})

	data, _, err := pipelineCache.CacheData()
	require.NoError(t, err)
	require.Equal(t, []byte{1, 1, 2, 3, 5, 8, 13, 21}, data)
}

func TestVulkanPipelineCache_MergePipelineCaches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	pipelineCache := mocks.EasyDummyPipelineCache(t, device, loader)

	srcPipeline1 := mocks.EasyMockPipelineCache(ctrl)
	srcPipeline2 := mocks.EasyMockPipelineCache(ctrl)
	srcPipeline3 := mocks.EasyMockPipelineCache(ctrl)

	mockDriver.EXPECT().VkMergePipelineCaches(mocks.Exactly(device.Handle()), mocks.Exactly(pipelineCache.Handle()), driver.Uint32(3), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, dstCache driver.VkPipelineCache, srcCacheCount driver.Uint32, pSrcCaches *driver.VkPipelineCache) (common.VkResult, error) {
			cacheSlice := unsafe.Slice(pSrcCaches, 3)

			require.Same(t, srcPipeline1.Handle(), cacheSlice[0])
			require.Same(t, srcPipeline2.Handle(), cacheSlice[1])
			require.Same(t, srcPipeline3.Handle(), cacheSlice[2])

			return common.VKSuccess, nil
		})

	_, err = pipelineCache.MergePipelineCaches([]core.PipelineCache{
		srcPipeline1, srcPipeline2, srcPipeline3,
	})
	require.NoError(t, err)
}