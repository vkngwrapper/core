package core_test

import (
	"bytes"
	"encoding/binary"
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanLoader1_0_CreateQueryPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	poolHandle := mocks.NewFakeQueryPool()

	driver.EXPECT().VkCreateQueryPool(mocks.Exactly(device.Handle()), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkQueryPoolCreateInfo, pAllocator *core.VkAllocationCallbacks, pQueryPool *core.VkQueryPool) (core.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(11), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(1000150000), val.FieldByName("queryType").Uint()) //VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR
			require.Equal(t, uint64(5), val.FieldByName("queryCount").Uint())
			require.Equal(t, uint64(0x10), val.FieldByName("pipelineStatistics").Uint()) // VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_PRIMITIVES_BIT

			*pQueryPool = poolHandle
			return core.VKSuccess, nil
		})

	queryPool, _, err := loader.CreateQueryPool(device, nil, &core.QueryPoolOptions{
		QueryType:          common.QueryTypeAccelerationStructureCompactedSizeKHR,
		QueryCount:         5,
		PipelineStatistics: common.PipelineStatisticGeometryShaderPrimitives,
	})
	require.NoError(t, err)
	require.NotNil(t, queryPool)
	require.Same(t, poolHandle, queryPool.Handle())
}

func TestVulkanQueryPool_PopulateResults(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	queryPool := mocks.EasyDummyQueryPool(t, loader, device)

	driver.EXPECT().VkGetQueryPoolResults(
		mocks.Exactly(device.Handle()),
		mocks.Exactly(queryPool.Handle()),
		core.Uint32(1),
		core.Uint32(3),
		core.Size(40),
		gomock.Not(nil),
		core.VkDeviceSize(8),
		core.VkQueryResultFlags(8), // VK_QUERY_RESULT_PARTIAL_BIT
	).DoAndReturn(
		func(device core.VkDevice,
			queryPool core.VkQueryPool,
			firstQuery, queryCount core.Uint32,
			dataSize core.Size,
			pData unsafe.Pointer,
			stride core.VkDeviceSize,
			flags core.VkQueryResultFlags) (core.VkResult, error) {

			data := ([]uint64)(unsafe.Slice((*uint64)(pData), 5))
			data[0] = 1
			data[1] = 3
			data[2] = 5
			data[3] = 8
			data[4] = 13

			return core.VKSuccess, nil
		})

	results, _, err := queryPool.PopulateResults(1, 3, 40, 8, common.QueryResultPartial)
	require.NoError(t, err)
	require.Len(t, results, 40)

	longs := []uint64{uint64(0), uint64(0), uint64(0), uint64(0), uint64(0)}
	reader := bytes.NewReader(results)
	err = binary.Read(reader, common.ByteOrder, longs)
	require.NoError(t, err)

	require.Len(t, longs, 5)
	require.Equal(t, []uint64{uint64(1), uint64(3), uint64(5), uint64(8), uint64(13)}, longs)
}
