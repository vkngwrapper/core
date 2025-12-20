package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanDescriptorSet is an implementation of the DescriptorSet interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
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

func (s *VulkanDescriptorSet) APIVersion() common.APIVersion {
	return s.MaximumAPIVersion
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

	return res, nil
}

func (s *VulkanDescriptorSet) DescriptorPoolHandle() driver.VkDescriptorPool {
	return s.DescriptorPool
}

func (s *VulkanDescriptorSet) DeviceHandle() driver.VkDevice {
	return s.Device
}

func (s *VulkanDescriptorSet) Driver() driver.Driver {
	return s.DeviceDriver
}
