package core1_1_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
	"go.uber.org/mock/gomock"
)

func TestDeviceGroupSubmitOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)

	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})
	driver := mocks1_1.InternalDeviceDriver(coreLoader)
	fence := mocks.NewDummyFence(device)
	pool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(pool, device)

	semaphore1 := mocks.NewDummySemaphore(device)
	semaphore2 := mocks.NewDummySemaphore(device)
	semaphore3 := mocks.NewDummySemaphore(device)

	queue := mocks.NewDummyQueue(device)

	coreLoader.EXPECT().VkQueueSubmit(
		queue.Handle(),
		loader.Uint32(1),
		gomock.Not(gomock.Nil()),
		fence.Handle(),
	).DoAndReturn(func(queue loader.VkQueue, submitCount loader.Uint32, pSubmits *loader.VkSubmitInfo, fence loader.VkFence) (common.VkResult, error) {
		val := reflect.ValueOf(pSubmits).Elem()

		require.Equal(t, uint64(4), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBMIT_INFO
		require.Equal(t, uint64(1), val.FieldByName("waitSemaphoreCount").Uint())
		require.Equal(t, uint64(1), val.FieldByName("commandBufferCount").Uint())
		require.Equal(t, uint64(2), val.FieldByName("signalSemaphoreCount").Uint())

		require.Equal(t, semaphore1.Handle(), loader.VkSemaphore(val.FieldByName("pWaitSemaphores").Elem().UnsafePointer()))
		require.Equal(t, uint64(0x00002000), val.FieldByName("pWaitDstStageMask").Elem().Uint()) // VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
		require.Equal(t, commandBuffer.Handle(), loader.VkCommandBuffer(val.FieldByName("pCommandBuffers").Elem().UnsafePointer()))

		semaphores := (*loader.VkSemaphore)(val.FieldByName("pSignalSemaphores").UnsafePointer())
		semaphoreSlice := ([]loader.VkSemaphore)(unsafe.Slice(semaphores, 2))
		require.Equal(t, semaphore2.Handle(), semaphoreSlice[0])
		require.Equal(t, semaphore3.Handle(), semaphoreSlice[1])

		next := (*loader.VkDeviceGroupSubmitInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060005), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("waitSemaphoreCount").Uint())
		require.Equal(t, uint64(1), val.FieldByName("commandBufferCount").Uint())
		require.Equal(t, uint64(2), val.FieldByName("signalSemaphoreCount").Uint())

		require.Equal(t, uint64(1), val.FieldByName("pWaitSemaphoreDeviceIndices").Elem().Uint())
		require.Equal(t, uint64(2), val.FieldByName("pCommandBufferDeviceMasks").Elem().Uint())

		indices := (*loader.Uint32)(val.FieldByName("pSignalSemaphoreDeviceIndices").UnsafePointer())
		indexSlice := ([]loader.Uint32)(unsafe.Slice(indices, 2))
		require.Equal(t, []loader.Uint32{3, 5}, indexSlice)

		return core1_0.VKSuccess, nil
	})

	_, err := driver.QueueSubmit(queue, &fence,
		core1_0.SubmitInfo{
			CommandBuffers:   []core.CommandBuffer{commandBuffer},
			WaitSemaphores:   []core.Semaphore{semaphore1},
			SignalSemaphores: []core.Semaphore{semaphore2, semaphore3},
			WaitDstStageMask: []core1_0.PipelineStageFlags{core1_0.PipelineStageBottomOfPipe},

			NextOptions: common.NextOptions{
				core1_1.DeviceGroupSubmitInfo{
					WaitSemaphoreDeviceIndices:   []int{1},
					CommandBufferDeviceMasks:     []uint32{2},
					SignalSemaphoreDeviceIndices: []int{3, 5},
				},
			},
		},
	)
	require.NoError(t, err)
}

func TestProtectedMemorySubmitOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})
	fence := mocks.NewDummyFence(device)
	pool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(pool, device)

	semaphore1 := mocks.NewDummySemaphore(device)
	semaphore2 := mocks.NewDummySemaphore(device)
	semaphore3 := mocks.NewDummySemaphore(device)

	queue := mocks.NewDummyQueue(device)

	coreLoader.EXPECT().VkQueueSubmit(
		queue.Handle(),
		loader.Uint32(1),
		gomock.Not(gomock.Nil()),
		fence.Handle(),
	).DoAndReturn(func(queue loader.VkQueue, submitCount loader.Uint32, pSubmits *loader.VkSubmitInfo, fence loader.VkFence) (common.VkResult, error) {
		val := reflect.ValueOf(pSubmits).Elem()

		require.Equal(t, uint64(4), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBMIT_INFO
		require.Equal(t, uint64(1), val.FieldByName("waitSemaphoreCount").Uint())
		require.Equal(t, uint64(1), val.FieldByName("commandBufferCount").Uint())
		require.Equal(t, uint64(2), val.FieldByName("signalSemaphoreCount").Uint())

		require.Equal(t, semaphore1.Handle(), loader.VkSemaphore(val.FieldByName("pWaitSemaphores").Elem().UnsafePointer()))
		require.Equal(t, uint64(0x00002000), val.FieldByName("pWaitDstStageMask").Elem().Uint()) // VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
		require.Equal(t, commandBuffer.Handle(), loader.VkCommandBuffer(val.FieldByName("pCommandBuffers").Elem().UnsafePointer()))

		semaphores := (*loader.VkSemaphore)(val.FieldByName("pSignalSemaphores").UnsafePointer())
		semaphoreSlice := ([]loader.VkSemaphore)(unsafe.Slice(semaphores, 2))
		require.Equal(t, semaphore2.Handle(), semaphoreSlice[0])
		require.Equal(t, semaphore3.Handle(), semaphoreSlice[1])

		next := (*loader.VkProtectedSubmitInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000145000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PROTECTED_SUBMIT_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("protectedSubmit").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := driver.QueueSubmit(queue, &fence,
		core1_0.SubmitInfo{
			CommandBuffers:   []core.CommandBuffer{commandBuffer},
			WaitSemaphores:   []core.Semaphore{semaphore1},
			SignalSemaphores: []core.Semaphore{semaphore2, semaphore3},
			WaitDstStageMask: []core1_0.PipelineStageFlags{core1_0.PipelineStageBottomOfPipe},

			NextOptions: common.NextOptions{
				core1_1.ProtectedSubmitInfo{
					ProtectedSubmit: true,
				},
			},
		},
	)
	require.NoError(t, err)
}
