package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/driver"
	"unsafe"
)

// VulkanDescriptorSet is an implementation of the DescriptorSet interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDescriptorSet struct {
	descriptorSetHandle driver.VkDescriptorSet
	deviceDriver        driver.Driver
	device              driver.VkDevice
	descriptorPool      driver.VkDescriptorPool

	maximumAPIVersion common.APIVersion
}

func (s *VulkanDescriptorSet) Handle() driver.VkDescriptorSet {
	return s.descriptorSetHandle
}

func (s *VulkanDescriptorSet) APIVersion() common.APIVersion {
	return s.maximumAPIVersion
}

func (s *VulkanDescriptorSet) Free() (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	vkDescriptorSet := (*driver.VkDescriptorSet)(arena.Malloc(int(unsafe.Sizeof([1]C.VkDescriptorSet{}))))
	commandBufferSlice := ([]driver.VkDescriptorSet)(unsafe.Slice(vkDescriptorSet, 1))
	commandBufferSlice[0] = s.descriptorSetHandle

	res, err := s.deviceDriver.VkFreeDescriptorSets(s.device, s.descriptorPool, 1, vkDescriptorSet)
	if err != nil {
		return res, err
	}

	s.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(s.descriptorSetHandle))
	return res, nil
}

func (s *VulkanDescriptorSet) DescriptorPoolHandle() driver.VkDescriptorPool {
	return s.descriptorPool
}

func (s *VulkanDescriptorSet) DeviceHandle() driver.VkDevice {
	return s.device
}

func (s *VulkanDescriptorSet) Driver() driver.Driver {
	return s.deviceDriver
}
