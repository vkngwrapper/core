package core1_1_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/internal/dummies"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestDeviceGroupSubmitOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)

	device := mocks.EasyMockDevice(ctrl, coreDriver)
	fence := mocks.EasyMockFence(ctrl)
	commandBuffer := mocks.EasyMockCommandBuffer(ctrl)

	semaphore1 := mocks.EasyMockSemaphore(ctrl)
	semaphore2 := mocks.EasyMockSemaphore(ctrl)
	semaphore3 := mocks.EasyMockSemaphore(ctrl)

	queue := dummies.EasyDummyQueue(coreDriver, device)

	coreDriver.EXPECT().VkQueueSubmit(
		queue.Handle(),
		driver.Uint32(1),
		gomock.Not(gomock.Nil()),
		fence.Handle(),
	).DoAndReturn(func(queue driver.VkQueue, submitCount driver.Uint32, pSubmits *driver.VkSubmitInfo, fence driver.VkFence) (common.VkResult, error) {
		val := reflect.ValueOf(pSubmits).Elem()

		require.Equal(t, uint64(4), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBMIT_INFO
		require.Equal(t, uint64(1), val.FieldByName("waitSemaphoreCount").Uint())
		require.Equal(t, uint64(1), val.FieldByName("commandBufferCount").Uint())
		require.Equal(t, uint64(2), val.FieldByName("signalSemaphoreCount").Uint())

		require.Equal(t, semaphore1.Handle(), driver.VkSemaphore(val.FieldByName("pWaitSemaphores").Elem().UnsafePointer()))
		require.Equal(t, uint64(0x00002000), val.FieldByName("pWaitDstStageMask").Elem().Uint()) // VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
		require.Equal(t, commandBuffer.Handle(), driver.VkCommandBuffer(val.FieldByName("pCommandBuffers").Elem().UnsafePointer()))

		semaphores := (*driver.VkSemaphore)(val.FieldByName("pSignalSemaphores").UnsafePointer())
		semaphoreSlice := ([]driver.VkSemaphore)(unsafe.Slice(semaphores, 2))
		require.Equal(t, semaphore2.Handle(), semaphoreSlice[0])
		require.Equal(t, semaphore3.Handle(), semaphoreSlice[1])

		next := (*driver.VkDeviceGroupSubmitInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060005), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("waitSemaphoreCount").Uint())
		require.Equal(t, uint64(1), val.FieldByName("commandBufferCount").Uint())
		require.Equal(t, uint64(2), val.FieldByName("signalSemaphoreCount").Uint())

		require.Equal(t, uint64(1), val.FieldByName("pWaitSemaphoreDeviceIndices").Elem().Uint())
		require.Equal(t, uint64(2), val.FieldByName("pCommandBufferDeviceMasks").Elem().Uint())

		indices := (*driver.Uint32)(val.FieldByName("pSignalSemaphoreDeviceIndices").UnsafePointer())
		indexSlice := ([]driver.Uint32)(unsafe.Slice(indices, 2))
		require.Equal(t, []driver.Uint32{3, 5}, indexSlice)

		return core1_0.VKSuccess, nil
	})

	_, err := queue.SubmitToQueue(fence, []core1_0.SubmitOptions{
		{
			CommandBuffers:   []core1_0.CommandBuffer{commandBuffer},
			WaitSemaphores:   []core1_0.Semaphore{semaphore1},
			SignalSemaphores: []core1_0.Semaphore{semaphore2, semaphore3},
			WaitDstStages:    []core1_0.PipelineStages{core1_0.PipelineStageBottomOfPipe},

			HaveNext: common.HaveNext{
				core1_1.DeviceGroupSubmitOptions{
					WaitSemaphoreDeviceIndices:   []int{1},
					CommandBufferDeviceMasks:     []uint32{2},
					SignalSemaphoreDeviceIndices: []int{3, 5},
				},
			},
		},
	})
	require.NoError(t, err)
}

func TestProtectedMemorySubmitOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)

	device := mocks.EasyMockDevice(ctrl, coreDriver)
	fence := mocks.EasyMockFence(ctrl)
	commandBuffer := mocks.EasyMockCommandBuffer(ctrl)

	semaphore1 := mocks.EasyMockSemaphore(ctrl)
	semaphore2 := mocks.EasyMockSemaphore(ctrl)
	semaphore3 := mocks.EasyMockSemaphore(ctrl)

	queue := dummies.EasyDummyQueue(coreDriver, device)

	coreDriver.EXPECT().VkQueueSubmit(
		queue.Handle(),
		driver.Uint32(1),
		gomock.Not(gomock.Nil()),
		fence.Handle(),
	).DoAndReturn(func(queue driver.VkQueue, submitCount driver.Uint32, pSubmits *driver.VkSubmitInfo, fence driver.VkFence) (common.VkResult, error) {
		val := reflect.ValueOf(pSubmits).Elem()

		require.Equal(t, uint64(4), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBMIT_INFO
		require.Equal(t, uint64(1), val.FieldByName("waitSemaphoreCount").Uint())
		require.Equal(t, uint64(1), val.FieldByName("commandBufferCount").Uint())
		require.Equal(t, uint64(2), val.FieldByName("signalSemaphoreCount").Uint())

		require.Equal(t, semaphore1.Handle(), driver.VkSemaphore(val.FieldByName("pWaitSemaphores").Elem().UnsafePointer()))
		require.Equal(t, uint64(0x00002000), val.FieldByName("pWaitDstStageMask").Elem().Uint()) // VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
		require.Equal(t, commandBuffer.Handle(), driver.VkCommandBuffer(val.FieldByName("pCommandBuffers").Elem().UnsafePointer()))

		semaphores := (*driver.VkSemaphore)(val.FieldByName("pSignalSemaphores").UnsafePointer())
		semaphoreSlice := ([]driver.VkSemaphore)(unsafe.Slice(semaphores, 2))
		require.Equal(t, semaphore2.Handle(), semaphoreSlice[0])
		require.Equal(t, semaphore3.Handle(), semaphoreSlice[1])

		next := (*driver.VkProtectedSubmitInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000145000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PROTECTED_SUBMIT_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("protectedSubmit").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := queue.SubmitToQueue(fence, []core1_0.SubmitOptions{
		{
			CommandBuffers:   []core1_0.CommandBuffer{commandBuffer},
			WaitSemaphores:   []core1_0.Semaphore{semaphore1},
			SignalSemaphores: []core1_0.Semaphore{semaphore2, semaphore3},
			WaitDstStages:    []core1_0.PipelineStages{core1_0.PipelineStageBottomOfPipe},

			HaveNext: common.HaveNext{
				core1_1.ProtectedSubmitOptions{
					ProtectedSubmit: true,
				},
			},
		},
	})
	require.NoError(t, err)
}
