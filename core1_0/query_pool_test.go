package core1_0_test

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	internal_mocks "github.com/vkngwrapper/core/v2/internal/dummies"
	"github.com/vkngwrapper/core/v2/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreateQueryPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := internal_mocks.EasyDummyDevice(mockDriver)
	poolHandle := mocks.NewFakeQueryPool()

	mockDriver.EXPECT().VkCreateQueryPool(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkQueryPoolCreateInfo, pAllocator *driver.VkAllocationCallbacks, pQueryPool *driver.VkQueryPool) (common.VkResult, error) {
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

	queryPool, _, err := device.CreateQueryPool(nil, core1_0.QueryPoolCreateInfo{
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

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, mockDriver)
	queryPool := internal_mocks.EasyDummyQueryPool(mockDriver, device)

	mockDriver.EXPECT().VkGetQueryPoolResults(
		device.Handle(),
		queryPool.Handle(),
		driver.Uint32(1),
		driver.Uint32(3),
		driver.Size(40),
		gomock.Not(nil),
		driver.VkDeviceSize(8),
		driver.VkQueryResultFlags(8), // VK_QUERY_RESULT_PARTIAL_BIT
	).DoAndReturn(
		func(device driver.VkDevice,
			queryPool driver.VkQueryPool,
			firstQuery, queryCount driver.Uint32,
			dataSize driver.Size,
			pData unsafe.Pointer,
			stride driver.VkDeviceSize,
			flags driver.VkQueryResultFlags) (common.VkResult, error) {

			data := ([]uint64)(unsafe.Slice((*uint64)(pData), 5))
			data[0] = 1
			data[1] = 3
			data[2] = 5
			data[3] = 8
			data[4] = 13

			return core1_0.VKSuccess, nil
		})

	results := make([]byte, 40)
	_, err := queryPool.PopulateResults(1, 3, results, 8, core1_0.QueryResultPartial)
	require.NoError(t, err)
	require.Len(t, results, 40)

	longs := []uint64{uint64(0), uint64(0), uint64(0), uint64(0), uint64(0)}
	reader := bytes.NewReader(results)
	err = binary.Read(reader, common.ByteOrder, longs)
	require.NoError(t, err)

	require.Len(t, longs, 5)
	require.Equal(t, []uint64{uint64(1), uint64(3), uint64(5), uint64(8), uint64(13)}, longs)
}
