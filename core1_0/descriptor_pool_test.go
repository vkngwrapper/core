package core1_0_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	internal_mocks "github.com/CannibalVox/VKng/core/internal/dummies"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestDescriptorPool_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := internal_mocks.EasyDummyDevice(mockDriver)

	expectedHandle := mocks.NewFakeDescriptorPool()

	mockDriver.EXPECT().VkCreateDescriptorPool(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkDescriptorPoolCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDescriptorPool *driver.VkDescriptorPool) (common.VkResult, error) {
			v := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(33), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), v.FieldByName("flags").Uint()) // VK_DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT
			require.Equal(t, uint64(3), v.FieldByName("maxSets").Uint())
			require.Equal(t, uint64(2), v.FieldByName("poolSizeCount").Uint())

			poolSizePtr := unsafe.Pointer(v.FieldByName("pPoolSizes").Elem().UnsafeAddr())
			poolSizeSlice := ([]driver.VkDescriptorPoolSize)(unsafe.Slice((*driver.VkDescriptorPoolSize)(poolSizePtr), 2))

			var sizeTypes []uint64
			var sizeCounts []uint64

			for _, poolSize := range poolSizeSlice {
				poolSizeVal := reflect.ValueOf(poolSize)
				sizeTypes = append(sizeTypes, poolSizeVal.FieldByName("_type").Uint())
				sizeCounts = append(sizeCounts, poolSizeVal.FieldByName("descriptorCount").Uint())
			}

			require.ElementsMatch(t, []uint64{1, 6}, sizeTypes)
			require.ElementsMatch(t, []uint64{4, 5}, sizeCounts)

			*pDescriptorPool = expectedHandle
			return core1_0.VKSuccess, nil
		})

	pool, _, err := device.CreateDescriptorPool(nil, core1_0.DescriptorPoolCreateInfo{
		Flags:   core1_0.DescriptorPoolCreateFreeDescriptorSet,
		MaxSets: 3,
		PoolSizes: []core1_0.DescriptorPoolSize{
			{
				Type:            core1_0.DescriptorTypeCombinedImageSampler,
				DescriptorCount: 5,
			},
			{
				Type:            core1_0.DescriptorTypeUniformBuffer,
				DescriptorCount: 4,
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, pool)
	require.Equal(t, expectedHandle, pool.Handle())
}

func TestDescriptorPool_AllocAndFree_Single(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := internal_mocks.EasyDummyDevice(mockDriver)
	pool := mocks.EasyMockDescriptorPool(ctrl, device)

	setHandle := mocks.NewFakeDescriptorSet()
	layout := mocks.EasyMockDescriptorSetLayout(ctrl)

	mockDriver.EXPECT().VkAllocateDescriptorSets(device.Handle(), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pAllocateInfo *driver.VkDescriptorSetAllocateInfo, pSets *driver.VkDescriptorSet) (common.VkResult, error) {
			v := reflect.ValueOf(*pAllocateInfo)

			require.Equal(t, uint64(34), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())

			actualPool := (driver.VkDescriptorPool)(unsafe.Pointer(v.FieldByName("descriptorPool").Elem().UnsafeAddr()))
			require.Equal(t, pool.Handle(), actualPool)

			require.Equal(t, uint64(1), v.FieldByName("descriptorSetCount").Uint())

			setLayoutPtr := (*driver.VkDescriptorSetLayout)(unsafe.Pointer(v.FieldByName("pSetLayouts").Elem().UnsafeAddr()))
			setLayoutSlice := ([]driver.VkDescriptorSetLayout)(unsafe.Slice(setLayoutPtr, 1))

			require.Equal(t, layout.Handle(), setLayoutSlice[0])

			setSlice := ([]driver.VkDescriptorSet)(unsafe.Slice(pSets, 1))
			setSlice[0] = setHandle

			return core1_0.VKSuccess, nil
		})

	mockDriver.EXPECT().VkFreeDescriptorSets(device.Handle(), pool.Handle(), driver.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, descriptorPool driver.VkDescriptorPool, setCount driver.Uint32, pSets *driver.VkDescriptorSet) (common.VkResult, error) {
			setSlice := ([]driver.VkDescriptorSet)(unsafe.Slice(pSets, 1))
			require.Equal(t, setHandle, setSlice[0])

			return core1_0.VKSuccess, nil
		})

	sets, _, err := device.AllocateDescriptorSets(core1_0.DescriptorSetAllocateInfo{
		DescriptorPool: pool,
		SetLayouts:     []core1_0.DescriptorSetLayout{layout},
	})
	require.NoError(t, err)

	require.Len(t, sets, 1)
	require.Equal(t, setHandle, sets[0].Handle())

	_, err = device.FreeDescriptorSets(sets)
	require.NoError(t, err)
}

func TestDescriptorPool_AllocAndFree_Multi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := internal_mocks.EasyDummyDevice(mockDriver)
	pool := mocks.EasyMockDescriptorPool(ctrl, device)

	setHandle1 := mocks.NewFakeDescriptorSet()
	setHandle2 := mocks.NewFakeDescriptorSet()
	setHandle3 := mocks.NewFakeDescriptorSet()
	layout1 := mocks.EasyMockDescriptorSetLayout(ctrl)
	layout2 := mocks.EasyMockDescriptorSetLayout(ctrl)
	layout3 := mocks.EasyMockDescriptorSetLayout(ctrl)

	mockDriver.EXPECT().VkAllocateDescriptorSets(device.Handle(), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pAllocateInfo *driver.VkDescriptorSetAllocateInfo, pSets *driver.VkDescriptorSet) (common.VkResult, error) {
			v := reflect.ValueOf(*pAllocateInfo)

			require.Equal(t, uint64(34), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())

			actualPool := (driver.VkDescriptorPool)(unsafe.Pointer(v.FieldByName("descriptorPool").Elem().UnsafeAddr()))
			require.Equal(t, pool.Handle(), actualPool)

			require.Equal(t, uint64(3), v.FieldByName("descriptorSetCount").Uint())

			setLayoutPtr := (*driver.VkDescriptorSetLayout)(unsafe.Pointer(v.FieldByName("pSetLayouts").Elem().UnsafeAddr()))
			setLayoutSlice := ([]driver.VkDescriptorSetLayout)(unsafe.Slice(setLayoutPtr, 3))

			require.Equal(t, layout1.Handle(), setLayoutSlice[0])
			require.Equal(t, layout2.Handle(), setLayoutSlice[1])
			require.Equal(t, layout3.Handle(), setLayoutSlice[2])

			setSlice := ([]driver.VkDescriptorSet)(unsafe.Slice(pSets, 3))
			setSlice[0] = setHandle1
			setSlice[1] = setHandle2
			setSlice[2] = setHandle3

			return core1_0.VKSuccess, nil
		})

	mockDriver.EXPECT().VkFreeDescriptorSets(device.Handle(), pool.Handle(), driver.Uint32(3), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, descriptorPool driver.VkDescriptorPool, setCount driver.Uint32, pSets *driver.VkDescriptorSet) (common.VkResult, error) {
			setSlice := ([]driver.VkDescriptorSet)(unsafe.Slice(pSets, 3))
			require.Equal(t, setHandle1, setSlice[0])
			require.Equal(t, setHandle2, setSlice[1])
			require.Equal(t, setHandle3, setSlice[2])

			return core1_0.VKSuccess, nil
		})

	sets, _, err := device.AllocateDescriptorSets(core1_0.DescriptorSetAllocateInfo{
		DescriptorPool: pool,
		SetLayouts:     []core1_0.DescriptorSetLayout{layout1, layout2, layout3},
	})
	require.NoError(t, err)

	require.Len(t, sets, 3)
	require.Equal(t, setHandle1, sets[0].Handle())
	require.Equal(t, setHandle2, sets[1].Handle())
	require.Equal(t, setHandle3, sets[2].Handle())

	_, err = device.FreeDescriptorSets(sets)
	require.NoError(t, err)
}

func TestVulkanDescriptorPool_Reset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)
	pool := internal_mocks.EasyDummyDescriptorPool(mockDriver, mockDevice)

	mockDriver.EXPECT().VkResetDescriptorPool(mockDevice.Handle(), pool.Handle(), driver.VkDescriptorPoolResetFlags(3)).Return(core1_0.VKSuccess, nil)

	_, err := pool.Reset(3)
	require.NoError(t, err)
}
