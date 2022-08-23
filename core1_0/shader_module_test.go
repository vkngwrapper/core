package core1_0_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	internal_mocks "github.com/vkngwrapper/core/v2/internal/dummies"
	"github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanLoader1_0_CreateShaderModule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := internal_mocks.EasyDummyDevice(mockDriver)
	handle := mocks.NewFakeShaderModule()

	mockDriver.EXPECT().VkCreateShaderModule(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
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

			return core1_0.VKSuccess, nil
		})

	shaderModule, _, err := device.CreateShaderModule(nil, core1_0.ShaderModuleCreateInfo{
		Code: []uint32{1, 1, 2, 3, 5, 8, 13, 21},
	})
	require.NoError(t, err)
	require.NotNil(t, shaderModule)
	require.Equal(t, handle, shaderModule.Handle())
}
