package universal

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
	handle driver.VkDescriptorSet
	driver driver.Driver
	device driver.VkDevice
	pool   driver.VkDescriptorPool
}

func (s *VulkanDescriptorSet) Handle() driver.VkDescriptorSet {
	return s.handle
}

func (s *VulkanDescriptorSet) Free() (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	vkDescriptorSet := (*driver.VkDescriptorSet)(arena.Malloc(int(unsafe.Sizeof([1]C.VkDescriptorSet{}))))
	commandBufferSlice := ([]driver.VkDescriptorSet)(unsafe.Slice(vkDescriptorSet, 1))
	commandBufferSlice[0] = s.handle

	return s.driver.VkFreeDescriptorSets(s.device, s.pool, 1, vkDescriptorSet)
}

func (s *VulkanDescriptorSet) PoolHandle() driver.VkDescriptorPool {
	return s.pool
}

func (s *VulkanDescriptorSet) DeviceHandle() driver.VkDevice {
	return s.device
}

func (s *VulkanDescriptorSet) Driver() driver.Driver {
	return s.driver
}
