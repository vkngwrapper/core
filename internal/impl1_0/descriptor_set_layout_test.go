package impl1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestDescriptorSetLayout_Create_SingleBinding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	layoutHandle := mocks.NewFakeDescriptorSetLayout()

	mockDriver.EXPECT().VkCreateDescriptorSetLayout(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkDescriptorSetLayoutCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDescriptorSetLayout *driver.VkDescriptorSetLayout) (common.VkResult, error) {
			v := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(32), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), v.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), v.FieldByName("bindingCount").Uint())

			bindingsPtr := (*driver.VkDescriptorSetLayoutBinding)(unsafe.Pointer(v.FieldByName("pBindings").Elem().UnsafeAddr()))
			bindingsSlice := ([]driver.VkDescriptorSetLayoutBinding)(unsafe.Slice(bindingsPtr, 1))

			bindingV := reflect.ValueOf(bindingsSlice[0])
			require.Equal(t, uint64(3), bindingV.FieldByName("binding").Uint())
			require.Equal(t, uint64(7), bindingV.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_STORAGE_BUFFER
			require.Equal(t, uint64(1), bindingV.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(8), bindingV.FieldByName("stageFlags").Uint()) // VK_SHADER_STAGE_GEOMETRY_BIT
			require.True(t, bindingV.FieldByName("pImmutableSamplers").IsNil())

			*pDescriptorSetLayout = layoutHandle
			return core1_0.VKSuccess, nil
		})

	layout, _, err := device.CreateDescriptorSetLayout(nil, core1_0.DescriptorSetLayoutCreateInfo{
		Flags: 0,
		Bindings: []core1_0.DescriptorSetLayoutBinding{
			{
				Binding:         3,
				DescriptorType:  core1_0.DescriptorTypeStorageBuffer,
				DescriptorCount: 1,
				StageFlags:      core1_0.StageGeometry,
			},
		},
	})

	require.NoError(t, err)
	require.NotNil(t, layout)
	require.Equal(t, layoutHandle, layout.Handle())
}

func TestDescriptorSetLayout_Create_SingleBindingImmutableSamplers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	layoutHandle := mocks.NewFakeDescriptorSetLayout()

	sampler1 := mocks.EasyMockSampler(ctrl)
	sampler2 := mocks.EasyMockSampler(ctrl)
	sampler3 := mocks.EasyMockSampler(ctrl)
	sampler4 := mocks.EasyMockSampler(ctrl)

	mockDriver.EXPECT().VkCreateDescriptorSetLayout(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkDescriptorSetLayoutCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDescriptorSetLayout *driver.VkDescriptorSetLayout) (common.VkResult, error) {
			v := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(32), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), v.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), v.FieldByName("bindingCount").Uint())

			bindingsPtr := (*driver.VkDescriptorSetLayoutBinding)(unsafe.Pointer(v.FieldByName("pBindings").Elem().UnsafeAddr()))
			bindingsSlice := ([]driver.VkDescriptorSetLayoutBinding)(unsafe.Slice(bindingsPtr, 1))

			bindingV := reflect.ValueOf(bindingsSlice[0])
			require.Equal(t, uint64(3), bindingV.FieldByName("binding").Uint())
			require.Equal(t, uint64(1), bindingV.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER
			require.Equal(t, uint64(4), bindingV.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(8), bindingV.FieldByName("stageFlags").Uint()) // VK_SHADER_STAGE_GEOMETRY_BIT

			samplersPtr := (*driver.VkSampler)(unsafe.Pointer(bindingV.FieldByName("pImmutableSamplers").Elem().UnsafeAddr()))
			samplersSlice := ([]driver.VkSampler)(unsafe.Slice(samplersPtr, 4))

			require.Equal(t, sampler1.Handle(), samplersSlice[0])
			require.Equal(t, sampler2.Handle(), samplersSlice[1])
			require.Equal(t, sampler3.Handle(), samplersSlice[2])
			require.Equal(t, sampler4.Handle(), samplersSlice[3])

			*pDescriptorSetLayout = layoutHandle
			return core1_0.VKSuccess, nil
		})

	layout, _, err := device.CreateDescriptorSetLayout(nil, core1_0.DescriptorSetLayoutCreateInfo{
		Flags: 0,
		Bindings: []core1_0.DescriptorSetLayoutBinding{
			{
				Binding:         3,
				DescriptorType:  core1_0.DescriptorTypeCombinedImageSampler,
				DescriptorCount: 4,
				StageFlags:      core1_0.StageGeometry,
				ImmutableSamplers: []core1_0.Sampler{
					sampler1, sampler2, sampler3, sampler4,
				},
			},
		},
	})

	require.NoError(t, err)
	require.NotNil(t, layout)
	require.Equal(t, layoutHandle, layout.Handle())
}

func TestDescriptorSetLayout_Create_FailBindingSamplerMismatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})

	sampler1 := mocks.EasyMockSampler(ctrl)
	sampler2 := mocks.EasyMockSampler(ctrl)
	sampler3 := mocks.EasyMockSampler(ctrl)
	sampler4 := mocks.EasyMockSampler(ctrl)

	_, _, err := device.CreateDescriptorSetLayout(nil, core1_0.DescriptorSetLayoutCreateInfo{
		Flags: 0,
		Bindings: []core1_0.DescriptorSetLayoutBinding{
			{
				Binding:         3,
				DescriptorType:  core1_0.DescriptorTypeCombinedImageSampler,
				DescriptorCount: 3,
				StageFlags:      core1_0.StageGeometry,
				ImmutableSamplers: []core1_0.Sampler{
					sampler1, sampler2, sampler3, sampler4,
				},
			},
		},
	})

	require.EqualError(t, err, "allocate descriptor set layout bindings: binding 0 has 3 descriptors, but 4 immutable samplers. if immutable samplers are provided, they must match the descriptor count")
}

func TestDescriptorSetLayout_Create_MultiBinding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	layoutHandle := mocks.NewFakeDescriptorSetLayout()

	mockDriver.EXPECT().VkCreateDescriptorSetLayout(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkDescriptorSetLayoutCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDescriptorSetLayout *driver.VkDescriptorSetLayout) (common.VkResult, error) {
			v := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(32), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), v.FieldByName("flags").Uint())
			require.Equal(t, uint64(3), v.FieldByName("bindingCount").Uint())

			bindingsPtr := (*driver.VkDescriptorSetLayoutBinding)(unsafe.Pointer(v.FieldByName("pBindings").Elem().UnsafeAddr()))
			bindingsSlice := ([]driver.VkDescriptorSetLayoutBinding)(unsafe.Slice(bindingsPtr, 3))

			bindingV := reflect.ValueOf(bindingsSlice[0])
			require.Equal(t, uint64(3), bindingV.FieldByName("binding").Uint())
			require.Equal(t, uint64(7), bindingV.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_STORAGE_BUFFER
			require.Equal(t, uint64(1), bindingV.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(8), bindingV.FieldByName("stageFlags").Uint()) // VK_SHADER_STAGE_GEOMETRY_BIT
			require.True(t, bindingV.FieldByName("pImmutableSamplers").IsNil())

			bindingV = reflect.ValueOf(bindingsSlice[1])
			require.Equal(t, uint64(11), bindingV.FieldByName("binding").Uint())
			require.Equal(t, uint64(10), bindingV.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT
			require.Equal(t, uint64(9), bindingV.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(8), bindingV.FieldByName("stageFlags").Uint()) // VK_SHADER_STAGE_GEOMETRY_BIT
			require.True(t, bindingV.FieldByName("pImmutableSamplers").IsNil())

			bindingV = reflect.ValueOf(bindingsSlice[2])
			require.Equal(t, uint64(12), bindingV.FieldByName("binding").Uint())
			require.Equal(t, uint64(10), bindingV.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT
			require.Equal(t, uint64(18), bindingV.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(8), bindingV.FieldByName("stageFlags").Uint()) // VK_SHADER_STAGE_GEOMETRY_BIT
			require.True(t, bindingV.FieldByName("pImmutableSamplers").IsNil())

			*pDescriptorSetLayout = layoutHandle
			return core1_0.VKSuccess, nil
		})

	layout, _, err := device.CreateDescriptorSetLayout(nil, core1_0.DescriptorSetLayoutCreateInfo{
		Flags: 0,
		Bindings: []core1_0.DescriptorSetLayoutBinding{
			{
				Binding:         3,
				DescriptorType:  core1_0.DescriptorTypeStorageBuffer,
				DescriptorCount: 1,
				StageFlags:      core1_0.StageGeometry,
			},
			{
				Binding:         11,
				DescriptorType:  core1_0.DescriptorTypeInputAttachment,
				DescriptorCount: 9,
				StageFlags:      core1_0.StageGeometry,
			},
			{
				Binding:         12,
				DescriptorType:  core1_0.DescriptorTypeInputAttachment,
				DescriptorCount: 18,
				StageFlags:      core1_0.StageGeometry,
			},
		},
	})

	require.NoError(t, err)
	require.NotNil(t, layout)
	require.Equal(t, layoutHandle, layout.Handle())
}
