package core1_0

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_1"
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

	DescriptorSet1_1 core1_1.DescriptorSet
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

	return s.DeviceDriver.VkFreeDescriptorSets(s.Device, s.DescriptorPool, 1, vkDescriptorSet)
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

func (s *VulkanDescriptorSet) Core1_1() core1_1.DescriptorSet {
	return s.DescriptorSet1_1
}
