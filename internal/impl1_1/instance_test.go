package impl1_1_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
	"go.uber.org/mock/gomock"
)

func TestVulkanInstance_EnumeratePhysicalDeviceGroups(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalCoreInstanceDriver(coreLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})

	physicalDevice1 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)
	physicalDevice2 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)
	physicalDevice3 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)
	physicalDevice4 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)
	physicalDevice5 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)
	physicalDevice6 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	coreLoader.EXPECT().VkEnumeratePhysicalDeviceGroups(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(instance loader.VkInstance, pCount *loader.Uint32, pProperties *loader.VkPhysicalDeviceGroupProperties) (common.VkResult, error) {
		*pCount = loader.Uint32(3)

		return core1_0.VKSuccess, nil
	})

	coreLoader.EXPECT().VkEnumeratePhysicalDeviceGroups(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(instance loader.VkInstance, pCount *loader.Uint32, pProperties *loader.VkPhysicalDeviceGroupProperties) (common.VkResult, error) {
		require.Equal(t, loader.Uint32(3), *pCount)

		propertySlice := ([]loader.VkPhysicalDeviceGroupProperties)(unsafe.Slice(pProperties, 3))
		props := reflect.ValueOf(propertySlice)
		prop := props.Index(0)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = loader.Uint32(1)
		*(*loader.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = loader.VkBool32(1)

		propDevices := prop.FieldByName("physicalDevices")
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice1.Handle()

		prop = props.Index(1)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = loader.Uint32(2)
		*(*loader.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = loader.VkBool32(0)

		propDevices = prop.FieldByName("physicalDevices")
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice2.Handle()
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice3.Handle()

		prop = props.Index(2)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = loader.Uint32(3)
		*(*loader.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = loader.VkBool32(1)

		propDevices = prop.FieldByName("physicalDevices")
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice4.Handle()
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice5.Handle()
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(2).UnsafeAddr())) = physicalDevice6.Handle()

		return core1_0.VKSuccess, nil
	})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice1.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice2.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice3.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice4.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice5.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice6.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	groups, _, err := driver.EnumeratePhysicalDeviceGroups(instance, nil)
	require.NoError(t, err)
	require.Len(t, groups, 3)
	require.True(t, groups[0].SubsetAllocation)
	require.Len(t, groups[0].PhysicalDevices, 1)
	require.Equal(t, physicalDevice1.Handle(), groups[0].PhysicalDevices[0].Handle())

	require.False(t, groups[1].SubsetAllocation)
	require.Len(t, groups[1].PhysicalDevices, 2)
	require.Equal(t, physicalDevice2.Handle(), groups[1].PhysicalDevices[0].Handle())
	require.Equal(t, physicalDevice3.Handle(), groups[1].PhysicalDevices[1].Handle())

	require.True(t, groups[2].SubsetAllocation)
	require.Len(t, groups[2].PhysicalDevices, 3)
	require.Equal(t, physicalDevice4.Handle(), groups[2].PhysicalDevices[0].Handle())
	require.Equal(t, physicalDevice5.Handle(), groups[2].PhysicalDevices[1].Handle())
	require.Equal(t, physicalDevice6.Handle(), groups[2].PhysicalDevices[2].Handle())
}

func TestVulkanInstance_EnumeratePhysicalDeviceGroups_Incomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalCoreInstanceDriver(coreLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})

	physicalDevice1 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)
	physicalDevice2 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)
	physicalDevice3 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)
	physicalDevice4 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)
	physicalDevice5 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)
	physicalDevice6 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	coreLoader.EXPECT().VkEnumeratePhysicalDeviceGroups(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(instance loader.VkInstance, pCount *loader.Uint32, pProperties *loader.VkPhysicalDeviceGroupProperties) (common.VkResult, error) {
		*pCount = loader.Uint32(2)

		return core1_0.VKSuccess, nil
	})

	coreLoader.EXPECT().VkEnumeratePhysicalDeviceGroups(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(instance loader.VkInstance, pCount *loader.Uint32, pProperties *loader.VkPhysicalDeviceGroupProperties) (common.VkResult, error) {
		require.Equal(t, loader.Uint32(2), *pCount)

		propertySlice := ([]loader.VkPhysicalDeviceGroupProperties)(unsafe.Slice(pProperties, 3))
		props := reflect.ValueOf(propertySlice)
		prop := props.Index(0)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = loader.Uint32(1)
		*(*loader.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = loader.VkBool32(1)

		propDevices := prop.FieldByName("physicalDevices")
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice1.Handle()

		prop = props.Index(1)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = loader.Uint32(2)
		*(*loader.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = loader.VkBool32(0)

		propDevices = prop.FieldByName("physicalDevices")
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice2.Handle()
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice3.Handle()

		return core1_0.VKIncomplete, nil
	})

	coreLoader.EXPECT().VkEnumeratePhysicalDeviceGroups(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(instance loader.VkInstance, pCount *loader.Uint32, pProperties *loader.VkPhysicalDeviceGroupProperties) (common.VkResult, error) {
		*pCount = loader.Uint32(3)

		return core1_0.VKSuccess, nil
	})

	coreLoader.EXPECT().VkEnumeratePhysicalDeviceGroups(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(instance loader.VkInstance, pCount *loader.Uint32, pProperties *loader.VkPhysicalDeviceGroupProperties) (common.VkResult, error) {
		require.Equal(t, loader.Uint32(3), *pCount)

		propertySlice := ([]loader.VkPhysicalDeviceGroupProperties)(unsafe.Slice(pProperties, 3))
		props := reflect.ValueOf(propertySlice)
		prop := props.Index(0)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = loader.Uint32(1)
		*(*loader.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = loader.VkBool32(1)

		propDevices := prop.FieldByName("physicalDevices")
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice1.Handle()

		prop = props.Index(1)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = loader.Uint32(2)
		*(*loader.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = loader.VkBool32(0)

		propDevices = prop.FieldByName("physicalDevices")
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice2.Handle()
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice3.Handle()

		prop = props.Index(2)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = loader.Uint32(3)
		*(*loader.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = loader.VkBool32(1)

		propDevices = prop.FieldByName("physicalDevices")
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice4.Handle()
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice5.Handle()
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(2).UnsafeAddr())) = physicalDevice6.Handle()

		return core1_0.VKSuccess, nil
	})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice1.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		}).Times(2)

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice2.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		}).Times(2)

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice3.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		}).Times(2)

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice4.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice5.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice6.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	groups, _, err := driver.EnumeratePhysicalDeviceGroups(instance, nil)
	require.NoError(t, err)
	require.Len(t, groups, 3)
	require.True(t, groups[0].SubsetAllocation)
	require.Len(t, groups[0].PhysicalDevices, 1)
	require.Equal(t, physicalDevice1.Handle(), groups[0].PhysicalDevices[0].Handle())

	require.False(t, groups[1].SubsetAllocation)
	require.Len(t, groups[1].PhysicalDevices, 2)
	require.Equal(t, physicalDevice2.Handle(), groups[1].PhysicalDevices[0].Handle())
	require.Equal(t, physicalDevice3.Handle(), groups[1].PhysicalDevices[1].Handle())

	require.True(t, groups[2].SubsetAllocation)
	require.Len(t, groups[2].PhysicalDevices, 3)
	require.Equal(t, physicalDevice4.Handle(), groups[2].PhysicalDevices[0].Handle())
	require.Equal(t, physicalDevice5.Handle(), groups[2].PhysicalDevices[1].Handle())
	require.Equal(t, physicalDevice6.Handle(), groups[2].PhysicalDevices[2].Handle())
}
