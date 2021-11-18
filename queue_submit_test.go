package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestSubmitToQueue_SignalSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockDevice := mocks.EasyDummyDevice(t, ctrl, loader)
	queue := mocks.EasyDummyQueue(t, mockDevice)

	fence := mocks.EasyDummyFence(t, loader, mockDevice)

	pool := mocks.EasyDummyCommandPool(t, loader, mockDevice)
	buffer := mocks.EasyDummyCommandBuffer(t, mockDevice, pool)

	waitSemaphore1 := mocks.EasyDummySemaphore(t, loader, mockDevice)
	waitSemaphore2 := mocks.EasyDummySemaphore(t, loader, mockDevice)

	signalSemaphore1 := mocks.EasyDummySemaphore(t, loader, mockDevice)
	signalSemaphore2 := mocks.EasyDummySemaphore(t, loader, mockDevice)
	signalSemaphore3 := mocks.EasyDummySemaphore(t, loader, mockDevice)

	mockDriver.EXPECT().VkQueueSubmit(queue.Handle(), core.Uint32(1), gomock.Not(nil), fence.Handle()).DoAndReturn(
		func(queue core.VkQueue, submitCount core.Uint32, pSubmits *core.VkSubmitInfo, fence core.VkFence) (core.VkResult, error) {
			submitSlices := ([]core.VkSubmitInfo)(unsafe.Slice(pSubmits, int(submitCount)))

			for _, submit := range submitSlices {
				v := reflect.ValueOf(submit)
				require.Equal(t, uint64(4), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBMIT_INFO
				require.True(t, v.FieldByName("pNext").IsNil())
				require.Equal(t, uint64(2), v.FieldByName("waitSemaphoreCount").Uint())
				require.Equal(t, uint64(1), v.FieldByName("commandBufferCount").Uint())
				require.Equal(t, uint64(3), v.FieldByName("signalSemaphoreCount").Uint())

				waitSemaphorePtr := unsafe.Pointer(v.FieldByName("pWaitSemaphores").Elem().UnsafeAddr())
				waitSemaphoreSlice := ([]core.VkSemaphore)(unsafe.Slice((*core.VkSemaphore)(waitSemaphorePtr), 2))

				require.ElementsMatch(t, []core.VkSemaphore{waitSemaphore1.Handle(), waitSemaphore2.Handle()}, waitSemaphoreSlice)

				waitDstStageMaskPtr := unsafe.Pointer(v.FieldByName("pWaitDstStageMask").Elem().UnsafeAddr())
				waitDstStageMaskSlice := ([]core.VkPipelineStageFlags)(unsafe.Slice((*core.VkPipelineStageFlags)(waitDstStageMaskPtr), 2))

				require.ElementsMatch(t, []core.VkPipelineStageFlags{8, 128}, waitDstStageMaskSlice)

				commandBufferPtr := unsafe.Pointer(v.FieldByName("pCommandBuffers").Elem().UnsafeAddr())
				commandBufferSlice := ([]core.VkCommandBuffer)(unsafe.Slice((*core.VkCommandBuffer)(commandBufferPtr), 1))

				require.ElementsMatch(t, []core.VkCommandBuffer{buffer.Handle()}, commandBufferSlice)

				signalSemaphorePtr := unsafe.Pointer(v.FieldByName("pSignalSemaphores").Elem().UnsafeAddr())
				signalSemaphoreSlice := ([]core.VkSemaphore)(unsafe.Slice((*core.VkSemaphore)(signalSemaphorePtr), 3))

				require.ElementsMatch(t, []core.VkSemaphore{signalSemaphore1.Handle(), signalSemaphore2.Handle(), signalSemaphore3.Handle()}, signalSemaphoreSlice)
			}

			return core.VKSuccess, nil
		})

	_, err = queue.SubmitToQueue(fence, []*core.SubmitOptions{
		{
			CommandBuffers:   []core.CommandBuffer{buffer},
			WaitSemaphores:   []core.Semaphore{waitSemaphore1, waitSemaphore2},
			WaitDstStages:    []common.PipelineStages{common.PipelineStageVertexShader, common.PipelineStageFragmentShader},
			SignalSemaphores: []core.Semaphore{signalSemaphore1, signalSemaphore2, signalSemaphore3},
		},
	})
	require.NoError(t, err)
}

func TestSubmitToQueue_NoSignalSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockDevice := mocks.EasyDummyDevice(t, ctrl, loader)
	queue := mocks.EasyDummyQueue(t, mockDevice)

	pool := mocks.EasyDummyCommandPool(t, loader, mockDevice)
	buffer := mocks.EasyDummyCommandBuffer(t, mockDevice, pool)

	mockDriver.EXPECT().VkQueueSubmit(queue.Handle(), core.Uint32(1), gomock.Not(nil), nil).DoAndReturn(
		func(queue core.VkQueue, submitCount core.Uint32, pSubmits *core.VkSubmitInfo, fence core.VkFence) (core.VkResult, error) {
			submitSlices := ([]core.VkSubmitInfo)(unsafe.Slice(pSubmits, int(submitCount)))

			for _, submit := range submitSlices {
				v := reflect.ValueOf(submit)
				require.Equal(t, uint64(4), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBMIT_INFO
				require.True(t, v.FieldByName("pNext").IsNil())
				require.Equal(t, uint64(0), v.FieldByName("waitSemaphoreCount").Uint())
				require.Equal(t, uint64(1), v.FieldByName("commandBufferCount").Uint())
				require.Equal(t, uint64(0), v.FieldByName("signalSemaphoreCount").Uint())

				require.True(t, v.FieldByName("pWaitSemaphores").IsNil())
				require.True(t, v.FieldByName("pWaitDstStageMask").IsNil())
				require.True(t, v.FieldByName("pSignalSemaphores").IsNil())

				commandBufferPtr := unsafe.Pointer(v.FieldByName("pCommandBuffers").Elem().UnsafeAddr())
				commandBufferSlice := ([]core.VkCommandBuffer)(unsafe.Slice((*core.VkCommandBuffer)(commandBufferPtr), 1))

				require.ElementsMatch(t, []core.VkCommandBuffer{buffer.Handle()}, commandBufferSlice)
			}

			return core.VKSuccess, nil
		})

	_, err = queue.SubmitToQueue(nil, []*core.SubmitOptions{
		{
			CommandBuffers:   []core.CommandBuffer{buffer},
			WaitSemaphores:   []core.Semaphore{},
			WaitDstStages:    []common.PipelineStages{},
			SignalSemaphores: []core.Semaphore{},
		},
	})
	require.NoError(t, err)
}

func TestSubmitToQueue_MismatchWaitSemaphores(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockDevice := mocks.EasyDummyDevice(t, ctrl, loader)
	queue := mocks.EasyDummyQueue(t, mockDevice)

	pool := mocks.EasyDummyCommandPool(t, loader, mockDevice)
	buffer := mocks.EasyDummyCommandBuffer(t, mockDevice, pool)

	waitSemaphore1 := mocks.EasyDummySemaphore(t, loader, mockDevice)
	waitSemaphore2 := mocks.EasyDummySemaphore(t, loader, mockDevice)

	_, err = queue.SubmitToQueue(nil, []*core.SubmitOptions{
		{
			CommandBuffers:   []core.CommandBuffer{buffer},
			WaitSemaphores:   []core.Semaphore{waitSemaphore1, waitSemaphore2},
			WaitDstStages:    []common.PipelineStages{common.PipelineStageFragmentShader},
			SignalSemaphores: []core.Semaphore{},
		},
	})
	require.EqualError(t, err, "attempted to submit with 2 wait semaphores but 1 dst stages- these should match")
}
