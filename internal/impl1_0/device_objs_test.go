package impl1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreateBufferView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	buffer := mocks.NewDummyBuffer(device)
	expectedBufferView := mocks.NewFakeBufferViewHandle()

	mockLoader.EXPECT().VkCreateBufferView(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkBufferViewCreateInfo, pAllocator *loader.VkAllocationCallbacks, pBufferView *loader.VkBufferView) (common.VkResult, error) {
			v := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, v.FieldByName("sType").Uint(), uint64(13)) // VK_STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, v.FieldByName("flags").Uint(), uint64(0))

			actualBuffer := (loader.VkBuffer)(unsafe.Pointer(v.FieldByName("buffer").Elem().UnsafeAddr()))
			require.Equal(t, actualBuffer, buffer.Handle())

			require.Equal(t, v.FieldByName("format").Uint(), uint64(103)) // VK_FORMAT_R32G32_SFLOAT
			require.Equal(t, v.FieldByName("offset").Uint(), uint64(5))
			require.Equal(t, v.FieldByName("_range").Uint(), uint64(7))

			*pBufferView = expectedBufferView
			return core1_0.VKSuccess, nil
		})

	bufferView, res, err := driver.CreateBufferView(device, nil, core1_0.BufferViewCreateInfo{
		Buffer: buffer,
		Format: core1_0.FormatR32G32SignedFloat,
		Offset: 5,
		Range:  7,
	})

	require.Equal(t, res, core1_0.VKSuccess)
	require.NoError(t, err)
	require.Equal(t, expectedBufferView, bufferView.Handle())
}
