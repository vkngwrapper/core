package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanLoader1_0_CreateShaderModule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	handle := mocks.NewFakeShaderModule()

	mockDriver.EXPECT().VkCreateShaderModule(mocks.Exactly(device.Handle()), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkShaderModuleCreateInfo, pAllocator *driver.VkAllocationCallbacks, pShaderModule *driver.VkShaderModule) (common.VkResult, error) {
			*pShaderModule = handle
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(16), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(32), val.FieldByName("codeSize").Uint())

			codePtr := (*driver.Uint32)(unsafe.Pointer(val.FieldByName("pCode").Elem().UnsafeAddr()))
			codeSlice := ([]driver.Uint32)(unsafe.Slice(codePtr, 8))

			require.Equal(t, []driver.Uint32{1, 1, 2, 3, 5, 8, 13, 21}, codeSlice)

			return common.VKSuccess, nil
		})

	shaderModule, _, err := loader.CreateShaderModule(device, nil, &core.ShaderModuleOptions{
		SpirVByteCode: []uint32{1, 1, 2, 3, 5, 8, 13, 21},
	})
	require.NoError(t, err)
	require.NotNil(t, shaderModule)
	require.Same(t, handle, shaderModule.Handle())
}
