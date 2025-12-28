package core1_2_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_2"
	"go.uber.org/mock/gomock"
)

func TestDescriptorSetVariableDescriptorCountAllocateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	descriptorPool := mocks.NewDummyDescriptorPool(device)
	descriptorLayout1 := mocks.NewDummyDescriptorSetLayout(device)
	descriptorLayout2 := mocks.NewDummyDescriptorSetLayout(device)
	descriptorLayout3 := mocks.NewDummyDescriptorSetLayout(device)
	descriptorLayout4 := mocks.NewDummyDescriptorSetLayout(device)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalDeviceDriver(device, coreLoader)

	mockDescriptorSet1 := mocks.NewDummyDescriptorSet(descriptorPool, device)
	mockDescriptorSet2 := mocks.NewDummyDescriptorSet(descriptorPool, device)
	mockDescriptorSet3 := mocks.NewDummyDescriptorSet(descriptorPool, device)
	mockDescriptorSet4 := mocks.NewDummyDescriptorSet(descriptorPool, device)

	coreLoader.EXPECT().VkAllocateDescriptorSets(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pAllocateInfo *loader.VkDescriptorSetAllocateInfo,
		pDescriptorSets *loader.VkDescriptorSet) (common.VkResult, error) {

		sets := unsafe.Slice(pDescriptorSets, 4)
		sets[0] = mockDescriptorSet1.Handle()
		sets[1] = mockDescriptorSet2.Handle()
		sets[2] = mockDescriptorSet3.Handle()
		sets[3] = mockDescriptorSet4.Handle()

		val := reflect.ValueOf(pAllocateInfo).Elem()
		require.Equal(t, uint64(34), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO

		next := (*loader.VkDescriptorSetVariableDescriptorCountAllocateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000161003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(4), val.FieldByName("descriptorSetCount").Uint())

		countsPtr := (*loader.Uint32)(val.FieldByName("pDescriptorCounts").UnsafePointer())
		countSlice := unsafe.Slice(countsPtr, 4)

		require.Equal(t, []loader.Uint32{1, 3, 5, 7}, countSlice)

		return core1_0.VKSuccess, nil
	})

	sets, _, err := driver.AllocateDescriptorSets(core1_0.DescriptorSetAllocateInfo{
		DescriptorPool: descriptorPool,
		SetLayouts: []core.DescriptorSetLayout{
			descriptorLayout1,
			descriptorLayout2,
			descriptorLayout3,
			descriptorLayout4,
		},
		NextOptions: common.NextOptions{
			core1_2.DescriptorSetVariableDescriptorCountAllocateInfo{
				DescriptorCounts: []int{1, 3, 5, 7},
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, sets, 4)
	require.Equal(t, []loader.VkDescriptorSet{
		mockDescriptorSet1.Handle(),
		mockDescriptorSet2.Handle(),
		mockDescriptorSet3.Handle(),
		mockDescriptorSet4.Handle(),
	}, []loader.VkDescriptorSet{
		sets[0].Handle(),
		sets[1].Handle(),
		sets[2].Handle(),
		sets[3].Handle(),
	})
}

func TestDescriptorSetLayoutBindingFlagsCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	mockDescriptorSetLayout := mocks.NewDummyDescriptorSetLayout(device)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkCreateDescriptorSetLayout(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice, pCreateInfo *loader.VkDescriptorSetLayoutCreateInfo, pAllocator *loader.VkAllocationCallbacks, pSetLayout *loader.VkDescriptorSetLayout) (common.VkResult, error) {
		*pSetLayout = mockDescriptorSetLayout.Handle()
		val := reflect.ValueOf(pCreateInfo).Elem()

		require.Equal(t, uint64(32), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO

		next := (*loader.VkDescriptorSetLayoutBindingFlagsCreateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000161000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("bindingCount").Uint())
		flagsPtr := (*loader.VkDescriptorBindingFlags)(val.FieldByName("pBindingFlags").UnsafePointer())
		flagSlice := unsafe.Slice(flagsPtr, 2)

		require.Equal(t, []loader.VkDescriptorBindingFlags{8, 1}, flagSlice)

		return core1_0.VKSuccess, nil
	})

	descriptorSetLayout, _, err := driver.CreateDescriptorSetLayout(
		nil,
		core1_0.DescriptorSetLayoutCreateInfo{
			NextOptions: common.NextOptions{
				core1_2.DescriptorSetLayoutBindingFlagsCreateInfo{
					BindingFlags: []core1_2.DescriptorBindingFlags{
						core1_2.DescriptorBindingVariableDescriptorCount,
						core1_2.DescriptorBindingUpdateAfterBind,
					},
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDescriptorSetLayout.Handle(), descriptorSetLayout.Handle())
}

func TestDescriptorSetVariableDescriptorCountLayoutSupportOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkGetDescriptorSetLayoutSupport(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pCreateInfo *loader.VkDescriptorSetLayoutCreateInfo,
		pSupport *loader.VkDescriptorSetLayoutSupport) {
		val := reflect.ValueOf(pSupport).Elem()

		require.Equal(t, uint64(1000168001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT
		next := (*loader.VkDescriptorSetVariableDescriptorCountLayoutSupport)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000161004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxVariableDescriptorCount").UnsafeAddr())) = loader.Uint32(7)
	})

	var outData core1_2.DescriptorSetVariableDescriptorCountLayoutSupport
	err := driver.GetDescriptorSetLayoutSupport(
		core1_0.DescriptorSetLayoutCreateInfo{},
		&core1_1.DescriptorSetLayoutSupport{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, core1_2.DescriptorSetVariableDescriptorCountLayoutSupport{
		MaxVariableDescriptorCount: 7,
	}, outData)
}
