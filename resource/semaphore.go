package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/loader"

type VulkanSemaphore struct {
	loader *loader.Loader
	device loader.VkDevice
	handle loader.VkSemaphore
}

func (s *VulkanSemaphore) Handle() loader.VkSemaphore {
	return s.handle
}

func (s *VulkanSemaphore) Destroy() error {
	return s.loader.VkDestroySemaphore(s.device, s.handle, nil)
}
