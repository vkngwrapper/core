package internal1_0

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanDescriptorSet struct {
	DescriptorSetHandle driver.VkDescriptorSet
	DeviceDriver        driver.Driver
	Device              driver.VkDevice
	DescriptorPool      driver.VkDescriptorPool

	MaximumAPIVersion common.APIVersion
}

func (s *VulkanDescriptorSet) Handle() driver.VkDescriptorSet {
	return s.DescriptorSetHandle
}

func (s *VulkanDescriptorSet) Free() (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	vkDescriptorSet := (*driver.VkDescriptorSet)(arena.Malloc(int(unsafe.Sizeof([1]C.VkDescriptorSet{}))))
	commandBufferSlice := ([]driver.VkDescriptorSet)(unsafe.Slice(vkDescriptorSet, 1))
	commandBufferSlice[0] = s.DescriptorSetHandle

	res, err := s.DeviceDriver.VkFreeDescriptorSets(s.Device, s.DescriptorPool, 1, vkDescriptorSet)
	if err != nil {
		return res, err
	}

	s.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(s.DescriptorSetHandle), s)
	return res, nil
}

func (s *VulkanDescriptorSet) PoolHandle() driver.VkDescriptorPool {
	return s.DescriptorPool
}

func (s *VulkanDescriptorSet) DeviceHandle() driver.VkDevice {
	return s.Device
}

func (s *VulkanDescriptorSet) Driver() driver.Driver {
	return s.DeviceDriver
}
