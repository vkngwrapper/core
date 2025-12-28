package core1_2_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_2"
	"go.uber.org/mock/gomock"
)

func TestSemaphoreTypeCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	mockSemaphore := mocks.NewDummySemaphore(device)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkCreateSemaphore(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pCreateInfo *loader.VkSemaphoreCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pSemaphore *loader.VkSemaphore) (common.VkResult, error) {

		*pSemaphore = mockSemaphore.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(9), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO

		next := (*loader.VkSemaphoreTypeCreateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000207002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_TYPE_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("semaphoreType").Uint()) // VK_SEMAPHORE_TYPE_TIMELINE
		require.Equal(t, uint64(13), val.FieldByName("initialValue").Uint())

		return core1_0.VKSuccess, nil
	})

	semaphore, _, err := driver.CreateSemaphore(
		nil,
		core1_0.SemaphoreCreateInfo{
			NextOptions: common.NextOptions{core1_2.SemaphoreTypeCreateInfo{
				SemaphoreType: core1_2.SemaphoreTypeTimeline,
				InitialValue:  uint64(13),
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockSemaphore.Handle(), semaphore.Handle())
}

func TestTimelineSemaphoreSubmitOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	queue := mocks.NewDummyQueue(device)
	fence := mocks.NewDummyFence(device)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkQueueSubmit(
		queue.Handle(),
		loader.Uint32(1),
		gomock.Not(gomock.Nil()),
		fence.Handle(),
	).DoAndReturn(func(queue loader.VkQueue,
		submitCount loader.Uint32,
		pSubmits *loader.VkSubmitInfo,
		fence loader.VkFence) (common.VkResult, error) {

		val := reflect.ValueOf(pSubmits).Elem()
		require.Equal(t, uint64(4), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBMIT_INFO

		next := (*loader.VkTimelineSemaphoreSubmitInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000207003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_TIMELINE_SEMAPHORE_SUBMIT_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("waitSemaphoreValueCount").Uint())
		require.Equal(t, uint64(3), val.FieldByName("signalSemaphoreValueCount").Uint())

		waitPtr := (*loader.Uint64)(val.FieldByName("pWaitSemaphoreValues").UnsafePointer())
		waitSlice := unsafe.Slice(waitPtr, 2)
		require.Equal(t, []loader.Uint64{3, 5}, waitSlice)

		signalPtr := (*loader.Uint64)(val.FieldByName("pSignalSemaphoreValues").UnsafePointer())
		signalSlice := unsafe.Slice(signalPtr, 3)
		require.Equal(t, []loader.Uint64{7, 11, 13}, signalSlice)

		return core1_0.VKSuccess, nil
	})

	_, err := driver.QueueSubmit(
		queue,
		&fence,

		core1_0.SubmitInfo{
			NextOptions: common.NextOptions{
				core1_2.TimelineSemaphoreSubmitInfo{
					WaitSemaphoreValues:   []uint64{3, 5},
					SignalSemaphoreValues: []uint64{7, 11, 13},
				},
			},
		},
	)
	require.NoError(t, err)
}
