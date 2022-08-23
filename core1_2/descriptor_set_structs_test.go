package core1_2_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/common/extensions"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/core1_2"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	"github.com/vkngwrapper/core/v2/internal/dummies"
	"github.com/vkngwrapper/core/v2/mocks"
	"reflect"
	"testing"
	"unsafe"
)

func TestDescriptorSetVariableDescriptorCountAllocateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_2)
	descriptorPool := mocks.EasyMockDescriptorPool(ctrl, device)
	descriptorLayout1 := mocks.EasyMockDescriptorSetLayout(ctrl)
	descriptorLayout2 := mocks.EasyMockDescriptorSetLayout(ctrl)
	descriptorLayout3 := mocks.EasyMockDescriptorSetLayout(ctrl)
	descriptorLayout4 := mocks.EasyMockDescriptorSetLayout(ctrl)

	mockDescriptorSet1 := mocks.EasyMockDescriptorSet(ctrl)
	mockDescriptorSet2 := mocks.EasyMockDescriptorSet(ctrl)
	mockDescriptorSet3 := mocks.EasyMockDescriptorSet(ctrl)
	mockDescriptorSet4 := mocks.EasyMockDescriptorSet(ctrl)

	coreDriver.EXPECT().VkAllocateDescriptorSets(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pAllocateInfo *driver.VkDescriptorSetAllocateInfo,
		pDescriptorSets *driver.VkDescriptorSet) (common.VkResult, error) {

		sets := unsafe.Slice(pDescriptorSets, 4)
		sets[0] = mockDescriptorSet1.Handle()
		sets[1] = mockDescriptorSet2.Handle()
		sets[2] = mockDescriptorSet3.Handle()
		sets[3] = mockDescriptorSet4.Handle()

		val := reflect.ValueOf(pAllocateInfo).Elem()
		require.Equal(t, uint64(34), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO

		next := (*driver.VkDescriptorSetVariableDescriptorCountAllocateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000161003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(4), val.FieldByName("descriptorSetCount").Uint())

		countsPtr := (*driver.Uint32)(val.FieldByName("pDescriptorCounts").UnsafePointer())
		countSlice := unsafe.Slice(countsPtr, 4)

		require.Equal(t, []driver.Uint32{1, 3, 5, 7}, countSlice)

		return core1_0.VKSuccess, nil
	})

	sets, _, err := device.AllocateDescriptorSets(core1_0.DescriptorSetAllocateInfo{
		DescriptorPool: descriptorPool,
		SetLayouts: []core1_0.DescriptorSetLayout{
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
	require.Equal(t, []driver.VkDescriptorSet{
		mockDescriptorSet1.Handle(),
		mockDescriptorSet2.Handle(),
		mockDescriptorSet3.Handle(),
		mockDescriptorSet4.Handle(),
	}, []driver.VkDescriptorSet{
		sets[0].Handle(),
		sets[1].Handle(),
		sets[2].Handle(),
		sets[3].Handle(),
	})
}

func TestDescriptorSetLayoutBindingFlagsCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)
	mockDescriptorSetLayout := mocks.EasyMockDescriptorSetLayout(ctrl)

	coreDriver.EXPECT().VkCreateDescriptorSetLayout(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice, pCreateInfo *driver.VkDescriptorSetLayoutCreateInfo, pAllocator *driver.VkAllocationCallbacks, pSetLayout *driver.VkDescriptorSetLayout) (common.VkResult, error) {
		*pSetLayout = mockDescriptorSetLayout.Handle()
		val := reflect.ValueOf(pCreateInfo).Elem()

		require.Equal(t, uint64(32), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO

		next := (*driver.VkDescriptorSetLayoutBindingFlagsCreateInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000161000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("bindingCount").Uint())
		flagsPtr := (*driver.VkDescriptorBindingFlags)(val.FieldByName("pBindingFlags").UnsafePointer())
		flagSlice := unsafe.Slice(flagsPtr, 2)

		require.Equal(t, []driver.VkDescriptorBindingFlags{8, 1}, flagSlice)

		return core1_0.VKSuccess, nil
	})

	descriptorSetLayout, _, err := device.CreateDescriptorSetLayout(
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

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	device := core1_2.PromoteDevice(dummies.EasyDummyDevice(coreDriver))

	coreDriver.EXPECT().VkGetDescriptorSetLayoutSupport(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkDescriptorSetLayoutCreateInfo,
		pSupport *driver.VkDescriptorSetLayoutSupport) {
		val := reflect.ValueOf(pSupport).Elem()

		require.Equal(t, uint64(1000168001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT
		next := (*driver.VkDescriptorSetVariableDescriptorCountLayoutSupport)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000161004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxVariableDescriptorCount").UnsafeAddr())) = driver.Uint32(7)
	})

	var outData core1_2.DescriptorSetVariableDescriptorCountLayoutSupport
	err := device.DescriptorSetLayoutSupport(
		core1_0.DescriptorSetLayoutCreateInfo{},
		&core1_1.DescriptorSetLayoutSupport{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, core1_2.DescriptorSetVariableDescriptorCountLayoutSupport{
		MaxVariableDescriptorCount: 7,
	}, outData)
}
