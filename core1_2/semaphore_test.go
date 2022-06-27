package core1_2_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/common/extensions"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_2"
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

func TestVulkanSemaphore_SemaphoreCounterValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	semaphore := core1_2.PromoteSemaphore(dummies.EasyDummySemaphore(coreDriver, device))

	coreDriver.EXPECT().VkGetSemaphoreCounterValue(
		device.Handle(),
		semaphore.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		semaphore driver.VkSemaphore,
		pValue *driver.Uint64) (common.VkResult, error) {

		*pValue = driver.Uint64(37)
		return core1_0.VKSuccess, nil
	})

	value, _, err := semaphore.CounterValue()
	require.NoError(t, err)
	require.Equal(t, uint64(37), value)
}

func TestSemaphoreTypeCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_2)
	mockSemaphore := mocks.EasyMockSemaphore(ctrl)

	coreDriver.EXPECT().VkCreateSemaphore(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkSemaphoreCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pSemaphore *driver.VkSemaphore) (common.VkResult, error) {

		*pSemaphore = mockSemaphore.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(9), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO

		next := (*driver.VkSemaphoreTypeCreateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000207002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_TYPE_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("semaphoreType").Uint()) // VK_SEMAPHORE_TYPE_TIMELINE
		require.Equal(t, uint64(13), val.FieldByName("initialValue").Uint())

		return core1_0.VKSuccess, nil
	})

	semaphore, _, err := device.CreateSemaphore(
		nil,
		core1_0.SemaphoreCreateOptions{
			HaveNext: common.HaveNext{core1_2.SemaphoreTypeCreateOptions{
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

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	queue := extensions.CreateQueueObject(coreDriver, device.Handle(), mocks.NewFakeQueue(), common.Vulkan1_2)
	fence := mocks.EasyMockFence(ctrl)

	coreDriver.EXPECT().VkQueueSubmit(
		queue.Handle(),
		driver.Uint32(1),
		gomock.Not(gomock.Nil()),
		fence.Handle(),
	).DoAndReturn(func(queue driver.VkQueue,
		submitCount driver.Uint32,
		pSubmits *driver.VkSubmitInfo,
		fence driver.VkFence) (common.VkResult, error) {

		val := reflect.ValueOf(pSubmits).Elem()
		require.Equal(t, uint64(4), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBMIT_INFO

		next := (*driver.VkTimelineSemaphoreSubmitInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000207003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_TIMELINE_SEMAPHORE_SUBMIT_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("waitSemaphoreValueCount").Uint())
		require.Equal(t, uint64(3), val.FieldByName("signalSemaphoreValueCount").Uint())

		waitPtr := (*driver.Uint64)(val.FieldByName("pWaitSemaphoreValues").UnsafePointer())
		waitSlice := unsafe.Slice(waitPtr, 2)
		require.Equal(t, []driver.Uint64{3, 5}, waitSlice)

		signalPtr := (*driver.Uint64)(val.FieldByName("pSignalSemaphoreValues").UnsafePointer())
		signalSlice := unsafe.Slice(signalPtr, 3)
		require.Equal(t, []driver.Uint64{7, 11, 13}, signalSlice)

		return core1_0.VKSuccess, nil
	})

	_, err := queue.SubmitToQueue(
		fence,
		[]core1_0.SubmitOptions{
			{
				HaveNext: common.HaveNext{
					core1_2.TimelineSemaphoreSubmitOptions{
						WaitSemaphoreValues:   []uint64{3, 5},
						SignalSemaphoreValues: []uint64{7, 11, 13},
					},
				},
			},
		})
	require.NoError(t, err)
}
