package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestDescriptorSetLayout_Create_SingleBinding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)
	layoutHandle := mocks.NewFakeDescriptorSetLayout()

	mockDriver.EXPECT().VkCreateDescriptorSetLayout(mockDevice.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkDescriptorSetLayoutCreateInfo, pAllocator *core.VkAllocationCallbacks, pDescriptorSetLayout *core.VkDescriptorSetLayout) (core.VkResult, error) {
			v := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(32), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(4), v.FieldByName("flags").Uint()) // VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_VALVE
			require.Equal(t, uint64(1), v.FieldByName("bindingCount").Uint())

			bindingsPtr := (*core.VkDescriptorSetLayoutBinding)(unsafe.Pointer(v.FieldByName("pBindings").Elem().UnsafeAddr()))
			bindingsSlice := ([]core.VkDescriptorSetLayoutBinding)(unsafe.Slice(bindingsPtr, 1))

			bindingV := reflect.ValueOf(bindingsSlice[0])
			require.Equal(t, uint64(3), bindingV.FieldByName("binding").Uint())
			require.Equal(t, uint64(7), bindingV.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_STORAGE_BUFFER
			require.Equal(t, uint64(1), bindingV.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(8), bindingV.FieldByName("stageFlags").Uint()) // VK_SHADER_STAGE_GEOMETRY_BIT
			require.True(t, bindingV.FieldByName("pImmutableSamplers").IsNil())

			*pDescriptorSetLayout = layoutHandle
			return core.VKSuccess, nil
		})

	layout, _, err := loader.CreateDescriptorSetLayout(mockDevice, &core.DescriptorSetLayoutOptions{
		Flags: core.DescriptorSetLayoutHostOnlyPoolValve,
		Bindings: []*core.DescriptorLayoutBinding{
			{
				Binding:      3,
				Type:         common.DescriptorStorageBuffer,
				Count:        1,
				ShaderStages: common.StageGeometry,
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

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)
	layoutHandle := mocks.NewFakeDescriptorSetLayout()

	sampler1 := mocks.EasyMockSampler(ctrl)
	sampler2 := mocks.EasyMockSampler(ctrl)
	sampler3 := mocks.EasyMockSampler(ctrl)
	sampler4 := mocks.EasyMockSampler(ctrl)

	mockDriver.EXPECT().VkCreateDescriptorSetLayout(mockDevice.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkDescriptorSetLayoutCreateInfo, pAllocator *core.VkAllocationCallbacks, pDescriptorSetLayout *core.VkDescriptorSetLayout) (core.VkResult, error) {
			v := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(32), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(4), v.FieldByName("flags").Uint()) // VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_VALVE
			require.Equal(t, uint64(1), v.FieldByName("bindingCount").Uint())

			bindingsPtr := (*core.VkDescriptorSetLayoutBinding)(unsafe.Pointer(v.FieldByName("pBindings").Elem().UnsafeAddr()))
			bindingsSlice := ([]core.VkDescriptorSetLayoutBinding)(unsafe.Slice(bindingsPtr, 1))

			bindingV := reflect.ValueOf(bindingsSlice[0])
			require.Equal(t, uint64(3), bindingV.FieldByName("binding").Uint())
			require.Equal(t, uint64(1), bindingV.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER
			require.Equal(t, uint64(4), bindingV.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(8), bindingV.FieldByName("stageFlags").Uint()) // VK_SHADER_STAGE_GEOMETRY_BIT

			samplersPtr := (*core.VkSampler)(unsafe.Pointer(bindingV.FieldByName("pImmutableSamplers").Elem().UnsafeAddr()))
			samplersSlice := ([]core.VkSampler)(unsafe.Slice(samplersPtr, 4))

			var samplerHandles []core.VkSampler
			for _, sampler := range samplersSlice {
				samplerHandles = append(samplerHandles, sampler)
			}

			require.ElementsMatch(t, []core.VkSampler{sampler1.Handle(), sampler2.Handle(), sampler3.Handle(), sampler4.Handle()},
				samplerHandles)

			*pDescriptorSetLayout = layoutHandle
			return core.VKSuccess, nil
		})

	layout, _, err := loader.CreateDescriptorSetLayout(mockDevice, &core.DescriptorSetLayoutOptions{
		Flags: core.DescriptorSetLayoutHostOnlyPoolValve,
		Bindings: []*core.DescriptorLayoutBinding{
			{
				Binding:      3,
				Type:         common.DescriptorCombinedImageSampler,
				Count:        4,
				ShaderStages: common.StageGeometry,
				ImmutableSamplers: []core.Sampler{
					sampler1, sampler2, sampler3, sampler4,
				},
			},
		},
	})

	require.NoError(t, err)
	require.NotNil(t, layout)
	require.Equal(t, layoutHandle, layout.Handle())
}

func TestDescriptorSetLayout_Create_MultiBinding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)
	layoutHandle := mocks.NewFakeDescriptorSetLayout()

	mockDriver.EXPECT().VkCreateDescriptorSetLayout(mockDevice.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkDescriptorSetLayoutCreateInfo, pAllocator *core.VkAllocationCallbacks, pDescriptorSetLayout *core.VkDescriptorSetLayout) (core.VkResult, error) {
			v := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(32), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(4), v.FieldByName("flags").Uint()) // VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_VALVE
			require.Equal(t, uint64(3), v.FieldByName("bindingCount").Uint())

			bindingsPtr := (*core.VkDescriptorSetLayoutBinding)(unsafe.Pointer(v.FieldByName("pBindings").Elem().UnsafeAddr()))
			bindingsSlice := ([]core.VkDescriptorSetLayoutBinding)(unsafe.Slice(bindingsPtr, 3))

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
			return core.VKSuccess, nil
		})

	layout, _, err := loader.CreateDescriptorSetLayout(mockDevice, &core.DescriptorSetLayoutOptions{
		Flags: core.DescriptorSetLayoutHostOnlyPoolValve,
		Bindings: []*core.DescriptorLayoutBinding{
			{
				Binding:      3,
				Type:         common.DescriptorStorageBuffer,
				Count:        1,
				ShaderStages: common.StageGeometry,
			},
			{
				Binding:      11,
				Type:         common.DescriptorInputAttachment,
				Count:        9,
				ShaderStages: common.StageGeometry,
			},
			{
				Binding:      12,
				Type:         common.DescriptorInputAttachment,
				Count:        18,
				ShaderStages: common.StageGeometry,
			},
		},
	})

	require.NoError(t, err)
	require.NotNil(t, layout)
	require.Equal(t, layoutHandle, layout.Handle())
}