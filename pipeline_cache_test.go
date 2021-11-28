package core_test

import (
	"github.com/CannibalVox/VKng/core"
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

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	pipelineCacheHandle := mocks.NewFakePipelineCache()

	driver.EXPECT().VkCreatePipelineCache(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice,
			pCreateInfo *core.VkPipelineCacheCreateInfo,
			pAllocator *core.VkAllocationCallbacks,
			pPipelineCache *core.VkPipelineCache) (core.VkResult, error) {
			*pPipelineCache = pipelineCacheHandle

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(17), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_CACHE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT_EXT
			require.Equal(t, uint64(4), val.FieldByName("initialDataSize").Uint())

			dataPtr := (*byte)(unsafe.Pointer(val.FieldByName("pInitialData").Pointer()))
			dataSlice := ([]byte)(unsafe.Slice(dataPtr, 4))

			require.Equal(t, []byte{1, 3, 5, 7}, dataSlice)

			return core.VKSuccess, nil
		})

	pipelineCache, _, err := loader.CreatePipelineCache(device, &core.PipelineCacheOptions{
		Flags:       core.PipelineCacheExternallySynchronized,
		InitialData: []byte{1, 3, 5, 7},
	})
	require.NoError(t, err)
	require.NotNil(t, pipelineCache)
	require.Equal(t, pipelineCacheHandle, pipelineCache.Handle())
}
