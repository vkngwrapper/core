package impl1_0_test

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreateQueryPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	poolHandle := mocks.NewFakeQueryPool()

	mockLoader.EXPECT().VkCreateQueryPool(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkQueryPoolCreateInfo, pAllocator *loader.VkAllocationCallbacks, pQueryPool *loader.VkQueryPool) (common.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(11), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(0), val.FieldByName("queryType").Uint()) //VK_QUERY_TYPE_OCCLUSION
			require.Equal(t, uint64(5), val.FieldByName("queryCount").Uint())
			require.Equal(t, uint64(0x10), val.FieldByName("pipelineStatistics").Uint()) // VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_PRIMITIVES_BIT

			*pQueryPool = poolHandle
			return core1_0.VKSuccess, nil
		})

	queryPool, _, err := driver.CreateQueryPool(device, nil, core1_0.QueryPoolCreateInfo{
		QueryType:          core1_0.QueryTypeOcclusion,
		QueryCount:         5,
		PipelineStatistics: core1_0.QueryPipelineStatisticGeometryShaderPrimitives,
	})
	require.NoError(t, err)
	require.NotNil(t, queryPool)
	require.Equal(t, poolHandle, queryPool.Handle())
}

func TestVulkanQueryPool_PopulateResults(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	queryPool := mocks.NewDummyQueryPool(device)

	mockLoader.EXPECT().VkGetQueryPoolResults(
		device.Handle(),
		queryPool.Handle(),
		loader.Uint32(1),
		loader.Uint32(3),
		loader.Size(40),
		gomock.Not(nil),
		loader.VkDeviceSize(8),
		loader.VkQueryResultFlags(8), // VK_QUERY_RESULT_PARTIAL_BIT
	).DoAndReturn(
		func(device loader.VkDevice,
			queryPool loader.VkQueryPool,
			firstQuery, queryCount loader.Uint32,
			dataSize loader.Size,
			pData unsafe.Pointer,
			stride loader.VkDeviceSize,
			flags loader.VkQueryResultFlags) (common.VkResult, error) {

			data := ([]uint64)(unsafe.Slice((*uint64)(pData), 5))
			data[0] = 1
			data[1] = 3
			data[2] = 5
			data[3] = 8
			data[4] = 13

			return core1_0.VKSuccess, nil
		})

	results := make([]byte, 40)
	_, err := driver.GetQueryPoolResults(queryPool, 1, 3, results, 8, core1_0.QueryResultPartial)
	require.NoError(t, err)
	require.Len(t, results, 40)

	longs := []uint64{uint64(0), uint64(0), uint64(0), uint64(0), uint64(0)}
	reader := bytes.NewReader(results)
	err = binary.Read(reader, common.ByteOrder, longs)
	require.NoError(t, err)

	require.Len(t, longs, 5)
	require.Equal(t, []uint64{uint64(1), uint64(3), uint64(5), uint64(8), uint64(13)}, longs)
}
