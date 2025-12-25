package impl1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestDescriptorPool_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	expectedHandle := mocks.NewFakeDescriptorPool()

	mockLoader.EXPECT().VkCreateDescriptorPool(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkDescriptorPoolCreateInfo, pAllocator *loader.VkAllocationCallbacks, pDescriptorPool *loader.VkDescriptorPool) (common.VkResult, error) {
			v := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(33), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), v.FieldByName("flags").Uint()) // VK_DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT
			require.Equal(t, uint64(3), v.FieldByName("maxSets").Uint())
			require.Equal(t, uint64(2), v.FieldByName("poolSizeCount").Uint())

			poolSizePtr := unsafe.Pointer(v.FieldByName("pPoolSizes").Elem().UnsafeAddr())
			poolSizeSlice := ([]loader.VkDescriptorPoolSize)(unsafe.Slice((*loader.VkDescriptorPoolSize)(poolSizePtr), 2))

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

	pool, _, err := driver.CreateDescriptorPool(device, nil, core1_0.DescriptorPoolCreateInfo{
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

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	pool := mocks.NewDummyDescriptorPool(device)
	layout := mocks.NewDummyDescriptorSetLayout(device)
	set := mocks.NewDummyDescriptorSet(pool, device)

	mockLoader.EXPECT().VkAllocateDescriptorSets(device.Handle(), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pAllocateInfo *loader.VkDescriptorSetAllocateInfo, pSets *loader.VkDescriptorSet) (common.VkResult, error) {
			v := reflect.ValueOf(*pAllocateInfo)

			require.Equal(t, uint64(34), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())

			actualPool := (loader.VkDescriptorPool)(unsafe.Pointer(v.FieldByName("descriptorPool").Elem().UnsafeAddr()))
			require.Equal(t, pool.Handle(), actualPool)

			require.Equal(t, uint64(1), v.FieldByName("descriptorSetCount").Uint())

			setLayoutPtr := (*loader.VkDescriptorSetLayout)(unsafe.Pointer(v.FieldByName("pSetLayouts").Elem().UnsafeAddr()))
			setLayoutSlice := ([]loader.VkDescriptorSetLayout)(unsafe.Slice(setLayoutPtr, 1))

			require.Equal(t, layout.Handle(), setLayoutSlice[0])

			setSlice := ([]loader.VkDescriptorSet)(unsafe.Slice(pSets, 1))
			setSlice[0] = set.Handle()

			return core1_0.VKSuccess, nil
		})

	mockLoader.EXPECT().VkFreeDescriptorSets(device.Handle(), pool.Handle(), loader.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, descriptorPool loader.VkDescriptorPool, setCount loader.Uint32, pSets *loader.VkDescriptorSet) (common.VkResult, error) {
			setSlice := ([]loader.VkDescriptorSet)(unsafe.Slice(pSets, 1))
			require.Equal(t, set.Handle(), setSlice[0])

			return core1_0.VKSuccess, nil
		})

	sets, _, err := driver.AllocateDescriptorSets(core1_0.DescriptorSetAllocateInfo{
		DescriptorPool: pool,
		SetLayouts:     []core.DescriptorSetLayout{layout},
	})
	require.NoError(t, err)

	require.Len(t, sets, 1)
	require.Equal(t, set.Handle(), sets[0].Handle())

	_, err = driver.FreeDescriptorSets(sets...)
	require.NoError(t, err)
}

func TestDescriptorPool_AllocAndFree_Multi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	pool := mocks.NewDummyDescriptorPool(device)

	set1 := mocks.NewDummyDescriptorSet(pool, device)
	set2 := mocks.NewDummyDescriptorSet(pool, device)
	set3 := mocks.NewDummyDescriptorSet(pool, device)

	layout1 := mocks.NewDummyDescriptorSetLayout(device)
	layout2 := mocks.NewDummyDescriptorSetLayout(device)
	layout3 := mocks.NewDummyDescriptorSetLayout(device)

	mockLoader.EXPECT().VkAllocateDescriptorSets(device.Handle(), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pAllocateInfo *loader.VkDescriptorSetAllocateInfo, pSets *loader.VkDescriptorSet) (common.VkResult, error) {
			v := reflect.ValueOf(*pAllocateInfo)

			require.Equal(t, uint64(34), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())

			actualPool := (loader.VkDescriptorPool)(unsafe.Pointer(v.FieldByName("descriptorPool").Elem().UnsafeAddr()))
			require.Equal(t, pool.Handle(), actualPool)

			require.Equal(t, uint64(3), v.FieldByName("descriptorSetCount").Uint())

			setLayoutPtr := (*loader.VkDescriptorSetLayout)(unsafe.Pointer(v.FieldByName("pSetLayouts").Elem().UnsafeAddr()))
			setLayoutSlice := ([]loader.VkDescriptorSetLayout)(unsafe.Slice(setLayoutPtr, 3))

			require.Equal(t, layout1.Handle(), setLayoutSlice[0])
			require.Equal(t, layout2.Handle(), setLayoutSlice[1])
			require.Equal(t, layout3.Handle(), setLayoutSlice[2])

			setSlice := ([]loader.VkDescriptorSet)(unsafe.Slice(pSets, 3))
			setSlice[0] = set1.Handle()
			setSlice[1] = set2.Handle()
			setSlice[2] = set3.Handle()

			return core1_0.VKSuccess, nil
		})

	mockLoader.EXPECT().VkFreeDescriptorSets(device.Handle(), pool.Handle(), loader.Uint32(3), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, descriptorPool loader.VkDescriptorPool, setCount loader.Uint32, pSets *loader.VkDescriptorSet) (common.VkResult, error) {
			setSlice := ([]loader.VkDescriptorSet)(unsafe.Slice(pSets, 3))
			require.Equal(t, set1.Handle(), setSlice[0])
			require.Equal(t, set2.Handle(), setSlice[1])
			require.Equal(t, set3.Handle(), setSlice[2])

			return core1_0.VKSuccess, nil
		})

	sets, _, err := driver.AllocateDescriptorSets(core1_0.DescriptorSetAllocateInfo{
		DescriptorPool: pool,
		SetLayouts:     []core.DescriptorSetLayout{layout1, layout2, layout3},
	})
	require.NoError(t, err)

	require.Len(t, sets, 3)
	require.Equal(t, set1.Handle(), sets[0].Handle())
	require.Equal(t, set2.Handle(), sets[1].Handle())
	require.Equal(t, set3.Handle(), sets[2].Handle())

	_, err = driver.FreeDescriptorSets(sets...)
	require.NoError(t, err)
}

func TestVulkanDescriptorPool_Reset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	pool := mocks.NewDummyDescriptorPool(device)

	mockLoader.EXPECT().VkResetDescriptorPool(device.Handle(), pool.Handle(), loader.VkDescriptorPoolResetFlags(3)).Return(core1_0.VKSuccess, nil)

	_, err := driver.ResetDescriptorPool(pool, 3)
	require.NoError(t, err)
}
