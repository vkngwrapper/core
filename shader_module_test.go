package core_test

import (
	"github.com/CannibalVox/VKng/core"
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

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	handle := mocks.NewFakeShaderModule()

	driver.EXPECT().VkCreateShaderModule(mocks.Exactly(device.Handle()), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkShaderModuleCreateInfo, pAllocator *core.VkAllocationCallbacks, pShaderModule *core.VkShaderModule) (core.VkResult, error) {
			*pShaderModule = handle
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(16), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(32), val.FieldByName("codeSize").Uint())

			codePtr := (*core.Uint32)(unsafe.Pointer(val.FieldByName("pCode").Elem().UnsafeAddr()))
			codeSlice := ([]core.Uint32)(unsafe.Slice(codePtr, 8))

			require.Equal(t, []core.Uint32{1, 1, 2, 3, 5, 8, 13, 21}, codeSlice)

			return core.VKSuccess, nil
		})

	shaderModule, _, err := loader.CreateShaderModule(device, nil, &core.ShaderModuleOptions{
		SpirVByteCode: []uint32{1, 1, 2, 3, 5, 8, 13, 21},
	})
	require.NoError(t, err)
	require.NotNil(t, shaderModule)
	require.Same(t, handle, shaderModule.Handle())
}
