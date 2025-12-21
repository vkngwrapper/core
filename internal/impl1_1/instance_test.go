package impl1_1_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
	"go.uber.org/mock/gomock"
)

func TestVulkanInstance_EnumeratePhysicalDeviceGroups(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := impl1_1.CreateInstanceObject(coreDriver, mocks.NewFakeInstanceHandle(), common.Vulkan1_1, []string{}).(core1_1.Instance)

	physicalDevice1 := mocks1_1.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice2 := mocks1_1.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice3 := mocks1_1.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice4 := mocks1_1.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice5 := mocks1_1.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice6 := mocks1_1.EasyMockPhysicalDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkEnumeratePhysicalDeviceGroups(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(instance driver.VkInstance, pCount *driver.Uint32, pProperties *driver.VkPhysicalDeviceGroupProperties) (common.VkResult, error) {
		*pCount = driver.Uint32(3)

		return core1_0.VKSuccess, nil
	})

	coreDriver.EXPECT().VkEnumeratePhysicalDeviceGroups(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(instance driver.VkInstance, pCount *driver.Uint32, pProperties *driver.VkPhysicalDeviceGroupProperties) (common.VkResult, error) {
		require.Equal(t, driver.Uint32(3), *pCount)

		propertySlice := ([]driver.VkPhysicalDeviceGroupProperties)(unsafe.Slice(pProperties, 3))
		props := reflect.ValueOf(propertySlice)
		prop := props.Index(0)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = driver.Uint32(1)
		*(*driver.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = driver.VkBool32(1)

		propDevices := prop.FieldByName("physicalDevices")
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice1.Handle()

		prop = props.Index(1)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = driver.Uint32(2)
		*(*driver.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = driver.VkBool32(0)

		propDevices = prop.FieldByName("physicalDevices")
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice2.Handle()
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice3.Handle()

		prop = props.Index(2)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = driver.Uint32(3)
		*(*driver.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = driver.VkBool32(1)

		propDevices = prop.FieldByName("physicalDevices")
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice4.Handle()
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice5.Handle()
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(2).UnsafeAddr())) = physicalDevice6.Handle()

		return core1_0.VKSuccess, nil
	})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice1.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice2.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice3.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice4.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice5.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice6.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	groups, _, err := instance.EnumeratePhysicalDeviceGroups(nil)
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

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	instance := impl1_1.CreateInstanceObject(coreDriver, mocks.NewFakeInstanceHandle(), common.Vulkan1_1, []string{}).(core1_1.Instance)

	physicalDevice1 := mocks1_1.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice2 := mocks1_1.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice3 := mocks1_1.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice4 := mocks1_1.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice5 := mocks1_1.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice6 := mocks1_1.EasyMockPhysicalDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkEnumeratePhysicalDeviceGroups(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(instance driver.VkInstance, pCount *driver.Uint32, pProperties *driver.VkPhysicalDeviceGroupProperties) (common.VkResult, error) {
		*pCount = driver.Uint32(2)

		return core1_0.VKSuccess, nil
	})

	coreDriver.EXPECT().VkEnumeratePhysicalDeviceGroups(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(instance driver.VkInstance, pCount *driver.Uint32, pProperties *driver.VkPhysicalDeviceGroupProperties) (common.VkResult, error) {
		require.Equal(t, driver.Uint32(2), *pCount)

		propertySlice := ([]driver.VkPhysicalDeviceGroupProperties)(unsafe.Slice(pProperties, 3))
		props := reflect.ValueOf(propertySlice)
		prop := props.Index(0)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = driver.Uint32(1)
		*(*driver.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = driver.VkBool32(1)

		propDevices := prop.FieldByName("physicalDevices")
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice1.Handle()

		prop = props.Index(1)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = driver.Uint32(2)
		*(*driver.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = driver.VkBool32(0)

		propDevices = prop.FieldByName("physicalDevices")
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice2.Handle()
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice3.Handle()

		return core1_0.VKIncomplete, nil
	})

	coreDriver.EXPECT().VkEnumeratePhysicalDeviceGroups(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(instance driver.VkInstance, pCount *driver.Uint32, pProperties *driver.VkPhysicalDeviceGroupProperties) (common.VkResult, error) {
		*pCount = driver.Uint32(3)

		return core1_0.VKSuccess, nil
	})

	coreDriver.EXPECT().VkEnumeratePhysicalDeviceGroups(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(instance driver.VkInstance, pCount *driver.Uint32, pProperties *driver.VkPhysicalDeviceGroupProperties) (common.VkResult, error) {
		require.Equal(t, driver.Uint32(3), *pCount)

		propertySlice := ([]driver.VkPhysicalDeviceGroupProperties)(unsafe.Slice(pProperties, 3))
		props := reflect.ValueOf(propertySlice)
		prop := props.Index(0)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = driver.Uint32(1)
		*(*driver.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = driver.VkBool32(1)

		propDevices := prop.FieldByName("physicalDevices")
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice1.Handle()

		prop = props.Index(1)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = driver.Uint32(2)
		*(*driver.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = driver.VkBool32(0)

		propDevices = prop.FieldByName("physicalDevices")
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice2.Handle()
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice3.Handle()

		prop = props.Index(2)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = driver.Uint32(3)
		*(*driver.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = driver.VkBool32(1)

		propDevices = prop.FieldByName("physicalDevices")
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice4.Handle()
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice5.Handle()
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(2).UnsafeAddr())) = physicalDevice6.Handle()

		return core1_0.VKSuccess, nil
	})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice1.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		}).Times(2)

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice2.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		}).Times(2)

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice3.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		}).Times(2)

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice4.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice5.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice6.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	groups, _, err := instance.EnumeratePhysicalDeviceGroups(nil)
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
