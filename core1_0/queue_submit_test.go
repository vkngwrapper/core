package core1_0_test

import (
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

func TestSubmitToQueue_SignalSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	queue := mocks.NewDummyQueue(device)

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	fence := mocks.NewDummyFence(device)
	pool := mocks.NewDummyCommandPool(device)
	buffer := mocks.NewDummyCommandBuffer(pool, device)

	waitSemaphore1 := mocks.NewDummySemaphore(device)
	waitSemaphore2 := mocks.NewDummySemaphore(device)

	signalSemaphore1 := mocks.NewDummySemaphore(device)
	signalSemaphore2 := mocks.NewDummySemaphore(device)
	signalSemaphore3 := mocks.NewDummySemaphore(device)

	mockLoader.EXPECT().VkQueueSubmit(queue.Handle(), loader.Uint32(1), gomock.Not(nil), fence.Handle()).DoAndReturn(
		func(queue loader.VkQueue, submitCount loader.Uint32, pSubmits *loader.VkSubmitInfo, fence loader.VkFence) (common.VkResult, error) {
			submitSlices := ([]loader.VkSubmitInfo)(unsafe.Slice(pSubmits, int(submitCount)))

			for _, submit := range submitSlices {
				v := reflect.ValueOf(submit)
				require.Equal(t, uint64(4), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBMIT_INFO
				require.True(t, v.FieldByName("pNext").IsNil())
				require.Equal(t, uint64(2), v.FieldByName("waitSemaphoreCount").Uint())
				require.Equal(t, uint64(1), v.FieldByName("commandBufferCount").Uint())
				require.Equal(t, uint64(3), v.FieldByName("signalSemaphoreCount").Uint())

				waitSemaphorePtr := unsafe.Pointer(v.FieldByName("pWaitSemaphores").Elem().UnsafeAddr())
				waitSemaphoreSlice := ([]loader.VkSemaphore)(unsafe.Slice((*loader.VkSemaphore)(waitSemaphorePtr), 2))

				require.Equal(t, waitSemaphore1.Handle(), waitSemaphoreSlice[0])
				require.Equal(t, waitSemaphore2.Handle(), waitSemaphoreSlice[1])

				waitDstStageMaskPtr := unsafe.Pointer(v.FieldByName("pWaitDstStageMask").Elem().UnsafeAddr())
				waitDstStageMaskSlice := ([]loader.VkPipelineStageFlags)(unsafe.Slice((*loader.VkPipelineStageFlags)(waitDstStageMaskPtr), 2))

				require.ElementsMatch(t, []loader.VkPipelineStageFlags{8, 128}, waitDstStageMaskSlice)

				commandBufferPtr := unsafe.Pointer(v.FieldByName("pCommandBuffers").Elem().UnsafeAddr())
				commandBufferSlice := ([]loader.VkCommandBuffer)(unsafe.Slice((*loader.VkCommandBuffer)(commandBufferPtr), 1))

				require.Equal(t, buffer.Handle(), commandBufferSlice[0])

				signalSemaphorePtr := unsafe.Pointer(v.FieldByName("pSignalSemaphores").Elem().UnsafeAddr())
				signalSemaphoreSlice := ([]loader.VkSemaphore)(unsafe.Slice((*loader.VkSemaphore)(signalSemaphorePtr), 3))

				require.Equal(t, signalSemaphore1.Handle(), signalSemaphoreSlice[0])
				require.Equal(t, signalSemaphore2.Handle(), signalSemaphoreSlice[1])
				require.Equal(t, signalSemaphore3.Handle(), signalSemaphoreSlice[2])
			}

			return core1_0.VKSuccess, nil
		})

	_, err := driver.QueueSubmit(queue, &fence,
		core1_0.SubmitInfo{
			CommandBuffers:   []core1_0.CommandBuffer{buffer},
			WaitSemaphores:   []core1_0.Semaphore{waitSemaphore1, waitSemaphore2},
			WaitDstStageMask: []core1_0.PipelineStageFlags{core1_0.PipelineStageVertexShader, core1_0.PipelineStageFragmentShader},
			SignalSemaphores: []core1_0.Semaphore{signalSemaphore1, signalSemaphore2, signalSemaphore3},
		},
	)
	require.NoError(t, err)
}

func TestSubmitToQueue_NoSignalSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	queue := mocks.NewDummyQueue(device)

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	pool := mocks.NewDummyCommandPool(device)
	buffer := mocks.NewDummyCommandBuffer(pool, device)

	mockLoader.EXPECT().VkQueueSubmit(queue.Handle(), loader.Uint32(1), gomock.Not(nil), loader.VkFence(loader.NullHandle)).DoAndReturn(
		func(queue loader.VkQueue, submitCount loader.Uint32, pSubmits *loader.VkSubmitInfo, fence loader.VkFence) (common.VkResult, error) {
			submitSlices := ([]loader.VkSubmitInfo)(unsafe.Slice(pSubmits, int(submitCount)))

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
				commandBufferSlice := ([]loader.VkCommandBuffer)(unsafe.Slice((*loader.VkCommandBuffer)(commandBufferPtr), 1))

				require.Equal(t, buffer.Handle(), commandBufferSlice[0])
			}

			return core1_0.VKSuccess, nil
		})

	_, err := driver.QueueSubmit(queue, nil,
		core1_0.SubmitInfo{
			CommandBuffers:   []core1_0.CommandBuffer{buffer},
			WaitSemaphores:   []core1_0.Semaphore{},
			WaitDstStageMask: []core1_0.PipelineStageFlags{},
			SignalSemaphores: []core1_0.Semaphore{},
		},
	)
	require.NoError(t, err)
}

func TestSubmitToQueue_MismatchWaitSemaphores(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	queue := mocks.NewDummyQueue(device)

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	pool := mocks.NewDummyCommandPool(device)
	buffer := mocks.NewDummyCommandBuffer(pool, device)

	waitSemaphore1 := mocks.NewDummySemaphore(device)
	waitSemaphore2 := mocks.NewDummySemaphore(device)

	_, err := driver.QueueSubmit(queue, nil,
		core1_0.SubmitInfo{
			CommandBuffers:   []core1_0.CommandBuffer{buffer},
			WaitSemaphores:   []core1_0.Semaphore{waitSemaphore1, waitSemaphore2},
			WaitDstStageMask: []core1_0.PipelineStageFlags{core1_0.PipelineStageFragmentShader},
			SignalSemaphores: []core1_0.Semaphore{},
		},
	)
	require.EqualError(t, err, "attempted to submit with 2 wait semaphores but 1 dst stages- these should match")
}
