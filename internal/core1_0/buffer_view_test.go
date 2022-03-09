package core1_0_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/internal/universal"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanLoader1_0_CreateBufferView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := universal.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)

	buffer := mocks.EasyMockBuffer(ctrl)
	expectedBufferView := mocks.NewFakeBufferViewHandle()

	mockDriver.EXPECT().VkCreateBufferView(mocks.Exactly(device.Handle()), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkBufferViewCreateInfo, pAllocator *driver.VkAllocationCallbacks, pBufferView *driver.VkBufferView) (common.VkResult, error) {
			v := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, v.FieldByName("sType").Uint(), uint64(13)) // VK_STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, v.FieldByName("flags").Uint(), uint64(0))

			actualBuffer := (driver.VkBuffer)(unsafe.Pointer(v.FieldByName("buffer").Elem().UnsafeAddr()))
			require.Equal(t, actualBuffer, buffer.Handle())

			require.Equal(t, v.FieldByName("format").Uint(), uint64(103)) // VK_FORMAT_R32G32_SFLOAT
			require.Equal(t, v.FieldByName("offset").Uint(), uint64(5))
			require.Equal(t, v.FieldByName("_range").Uint(), uint64(7))

			*pBufferView = expectedBufferView
			return common.VKSuccess, nil
		})

	bufferView, res, err := loader.CreateBufferView(device, nil, &core.BufferViewOptions{
		Buffer: buffer,
		Format: common.FormatR32G32SignedFloat,
		Offset: 5,
		Range:  7,
	})

	require.Equal(t, res, common.VKSuccess)
	require.NoError(t, err)
	require.Same(t, expectedBufferView, bufferView.Handle())
}