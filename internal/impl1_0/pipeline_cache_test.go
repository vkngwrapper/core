package impl1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreatePipelineCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	pipelineCacheHandle := mocks.NewFakePipelineCache()

	mockLoader.EXPECT().VkCreatePipelineCache(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice,
			pCreateInfo *loader.VkPipelineCacheCreateInfo,
			pAllocator *loader.VkAllocationCallbacks,
			pPipelineCache *loader.VkPipelineCache) (common.VkResult, error) {
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

	pipelineCache, _, err := driver.CreatePipelineCache(device, nil, core1_0.PipelineCacheCreateInfo{
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

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	pipelineCache := mocks.NewDummyPipelineCache(device)

	mockLoader.EXPECT().VkGetPipelineCacheData(device.Handle(), pipelineCache.Handle(), gomock.Not(nil), unsafe.Pointer(nil)).DoAndReturn(
		func(device loader.VkDevice, pipelineCache loader.VkPipelineCache, pSize *loader.Size, pCacheData unsafe.Pointer) (common.VkResult, error) {
			*pSize = 8
			return core1_0.VKSuccess, nil
		})
	mockLoader.EXPECT().VkGetPipelineCacheData(device.Handle(), pipelineCache.Handle(), gomock.Not(nil), gomock.Not(unsafe.Pointer(nil))).DoAndReturn(
		func(device loader.VkDevice, pipelineCache loader.VkPipelineCache, pSize *loader.Size, pCacheData unsafe.Pointer) (common.VkResult, error) {
			require.Equal(t, loader.Size(8), *pSize)
			bytes := ([]byte)(unsafe.Slice((*byte)(pCacheData), 8))
			copy(bytes, []byte{1, 1, 2, 3, 5, 8, 13, 21})
			return core1_0.VKSuccess, nil
		})

	data, _, err := driver.GetPipelineCacheData(pipelineCache)
	require.NoError(t, err)
	require.Equal(t, []byte{1, 1, 2, 3, 5, 8, 13, 21}, data)
}

func TestVulkanPipelineCache_MergePipelineCaches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	driver := impl1_0.NewDeviceDriver(mockLoader)
	pipelineCache := mocks.NewDummyPipelineCache(device)

	srcPipeline1 := mocks.NewDummyPipelineCache(device)
	srcPipeline2 := mocks.NewDummyPipelineCache(device)
	srcPipeline3 := mocks.NewDummyPipelineCache(device)

	mockLoader.EXPECT().VkMergePipelineCaches(device.Handle(), pipelineCache.Handle(), loader.Uint32(3), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, dstCache loader.VkPipelineCache, srcCacheCount loader.Uint32, pSrcCaches *loader.VkPipelineCache) (common.VkResult, error) {
			cacheSlice := unsafe.Slice(pSrcCaches, 3)

			require.Equal(t, srcPipeline1.Handle(), cacheSlice[0])
			require.Equal(t, srcPipeline2.Handle(), cacheSlice[1])
			require.Equal(t, srcPipeline3.Handle(), cacheSlice[2])

			return core1_0.VKSuccess, nil
		})

	_, err := driver.MergePipelineCaches(pipelineCache,
		srcPipeline1, srcPipeline2, srcPipeline3,
	)
	require.NoError(t, err)
}
