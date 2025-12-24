package impl1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreateShaderModule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	handle := mocks.NewFakeShaderModule()

	mockDriver.EXPECT().VkCreateShaderModule(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkShaderModuleCreateInfo, pAllocator *loader.VkAllocationCallbacks, pShaderModule *loader.VkShaderModule) (common.VkResult, error) {
			*pShaderModule = handle
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(16), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(32), val.FieldByName("codeSize").Uint())

			codePtr := (*loader.Uint32)(unsafe.Pointer(val.FieldByName("pCode").Elem().UnsafeAddr()))
			codeSlice := ([]loader.Uint32)(unsafe.Slice(codePtr, 8))

			require.Equal(t, []loader.Uint32{1, 1, 2, 3, 5, 8, 13, 21}, codeSlice)

			return core1_0.VKSuccess, nil
		})

	shaderModule, _, err := device.CreateShaderModule(nil, core1_0.ShaderModuleCreateInfo{
		Code: []uint32{1, 1, 2, 3, 5, 8, 13, 21},
	})
	require.NoError(t, err)
	require.NotNil(t, shaderModule)
	require.Equal(t, handle, shaderModule.Handle())
}
