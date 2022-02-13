package universal_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/internal/universal"
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

	mockDriver := mocks.NewMockDriver(ctrl)
	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)
	loader, err := universal.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	expectedHandle := mocks.NewFakeDescriptorPool()

	mockDriver.EXPECT().VkCreateDescriptorPool(mocks.Exactly(mockDevice.Handle()), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkDescriptorPoolCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDescriptorPool *driver.VkDescriptorPool) (common.VkResult, error) {
			v := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(33), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), v.FieldByName("flags").Uint()) // VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT
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

			require.ElementsMatch(t, []uint64{1, 1000351000}, sizeTypes)
			require.ElementsMatch(t, []uint64{4, 5}, sizeCounts)

			*pDescriptorPool = expectedHandle
			return common.VKSuccess, nil
		})

	pool, _, err := loader.CreateDescriptorPool(mockDevice, nil, &core.DescriptorPoolOptions{
		Flags:   DescriptorPoolUpdateAfterBind,
		MaxSets: 3,
		PoolSizes: []PoolSize{
			{
				Type:            common.DescriptorCombinedImageSampler,
				DescriptorCount: 5,
			},
			{
				Type:            common.DescriptorMutableValve,
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

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)
	pool := mocks.EasyDummyDescriptorPool(t, loader, mockDevice)

	setHandle := mocks.NewFakeDescriptorSet()
	layout := mocks.EasyDummyDescriptorSetLayout(t, loader, mockDevice)

	mockDriver.EXPECT().VkAllocateDescriptorSets(mocks.Exactly(mockDevice.Handle()), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pAllocateInfo *driver.VkDescriptorSetAllocateInfo, pSets *driver.VkDescriptorSet) (common.VkResult, error) {
			v := reflect.ValueOf(*pAllocateInfo)

			require.Equal(t, uint64(34), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())

			actualPool := (driver.VkDescriptorPool)(unsafe.Pointer(v.FieldByName("descriptorPool").Elem().UnsafeAddr()))
			require.Same(t, pool.Handle(), actualPool)

			require.Equal(t, uint64(1), v.FieldByName("descriptorSetCount").Uint())

			setLayoutPtr := (*driver.VkDescriptorSetLayout)(unsafe.Pointer(v.FieldByName("pSetLayouts").Elem().UnsafeAddr()))
			setLayoutSlice := ([]driver.VkDescriptorSetLayout)(unsafe.Slice(setLayoutPtr, 1))

			require.Same(t, layout.Handle(), setLayoutSlice[0])

			setSlice := ([]driver.VkDescriptorSet)(unsafe.Slice(pSets, 1))
			setSlice[0] = setHandle

			return common.VKSuccess, nil
		})

	mockDriver.EXPECT().VkFreeDescriptorSets(mocks.Exactly(mockDevice.Handle()), mocks.Exactly(pool.Handle()), driver.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, descriptorPool driver.VkDescriptorPool, setCount driver.Uint32, pSets *driver.VkDescriptorSet) (common.VkResult, error) {
			setSlice := ([]driver.VkDescriptorSet)(unsafe.Slice(pSets, 1))
			require.Same(t, setHandle, setSlice[0])

			return common.VKSuccess, nil
		})

	sets, _, err := pool.AllocateDescriptorSets(&core.DescriptorSetOptions{
		AllocationLayouts: []core.DescriptorSetLayout{layout},
	})
	require.NoError(t, err)

	require.Len(t, sets, 1)
	require.Same(t, setHandle, sets[0].Handle())

	_, err = pool.FreeDescriptorSets(sets)
	require.NoError(t, err)
}

func TestDescriptorPool_AllocAndFree_Multi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)
	pool := mocks.EasyDummyDescriptorPool(t, loader, mockDevice)

	setHandle1 := mocks.NewFakeDescriptorSet()
	setHandle2 := mocks.NewFakeDescriptorSet()
	setHandle3 := mocks.NewFakeDescriptorSet()
	layout1 := mocks.EasyDummyDescriptorSetLayout(t, loader, mockDevice)
	layout2 := mocks.EasyDummyDescriptorSetLayout(t, loader, mockDevice)
	layout3 := mocks.EasyDummyDescriptorSetLayout(t, loader, mockDevice)

	mockDriver.EXPECT().VkAllocateDescriptorSets(mocks.Exactly(mockDevice.Handle()), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pAllocateInfo *driver.VkDescriptorSetAllocateInfo, pSets *driver.VkDescriptorSet) (common.VkResult, error) {
			v := reflect.ValueOf(*pAllocateInfo)

			require.Equal(t, uint64(34), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())

			actualPool := (driver.VkDescriptorPool)(unsafe.Pointer(v.FieldByName("descriptorPool").Elem().UnsafeAddr()))
			require.Same(t, pool.Handle(), actualPool)

			require.Equal(t, uint64(3), v.FieldByName("descriptorSetCount").Uint())

			setLayoutPtr := (*driver.VkDescriptorSetLayout)(unsafe.Pointer(v.FieldByName("pSetLayouts").Elem().UnsafeAddr()))
			setLayoutSlice := ([]driver.VkDescriptorSetLayout)(unsafe.Slice(setLayoutPtr, 3))

			require.Same(t, layout1.Handle(), setLayoutSlice[0])
			require.Same(t, layout2.Handle(), setLayoutSlice[1])
			require.Same(t, layout3.Handle(), setLayoutSlice[2])

			setSlice := ([]driver.VkDescriptorSet)(unsafe.Slice(pSets, 3))
			setSlice[0] = setHandle1
			setSlice[1] = setHandle2
			setSlice[2] = setHandle3

			return common.VKSuccess, nil
		})

	mockDriver.EXPECT().VkFreeDescriptorSets(mocks.Exactly(mockDevice.Handle()), mocks.Exactly(pool.Handle()), driver.Uint32(3), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, descriptorPool driver.VkDescriptorPool, setCount driver.Uint32, pSets *driver.VkDescriptorSet) (common.VkResult, error) {
			setSlice := ([]driver.VkDescriptorSet)(unsafe.Slice(pSets, 3))
			require.Same(t, setHandle1, setSlice[0])
			require.Same(t, setHandle2, setSlice[1])
			require.Same(t, setHandle3, setSlice[2])

			return common.VKSuccess, nil
		})

	sets, _, err := pool.AllocateDescriptorSets(&core.DescriptorSetOptions{
		AllocationLayouts: []core.DescriptorSetLayout{layout1, layout2, layout3},
	})
	require.NoError(t, err)

	require.Len(t, sets, 3)
	require.Same(t, setHandle1, sets[0].Handle())
	require.Same(t, setHandle2, sets[1].Handle())
	require.Same(t, setHandle3, sets[2].Handle())

	_, err = pool.FreeDescriptorSets(sets)
	require.NoError(t, err)
}

func TestVulkanDescriptorPool_Reset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)
	pool := mocks.EasyDummyDescriptorPool(t, loader, mockDevice)

	mockDriver.EXPECT().VkResetDescriptorPool(mocks.Exactly(mockDevice.Handle()), mocks.Exactly(pool.Handle()), driver.VkDescriptorPoolResetFlags(3)).Return(common.VKSuccess, nil)

	_, err = pool.Reset(3)
	require.NoError(t, err)
}
