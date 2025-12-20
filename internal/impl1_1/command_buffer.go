package impl1_1

import (
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
)

// VulkanCommandBuffer is an implementation of the CommandBuffer interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanCommandBuffer struct {
	impl1_0.VulkanCommandBuffer
}

func (c *VulkanCommandBuffer) CmdDispatchBase(baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ int) {
	c.Driver().VkCmdDispatchBase(c.Handle(),
		driver.Uint32(baseGroupX),
		driver.Uint32(baseGroupY),
		driver.Uint32(baseGroupZ),
		driver.Uint32(groupCountX),
		driver.Uint32(groupCountY),
		driver.Uint32(groupCountZ))

	counter := c.CommandCounter()
	counter.CommandCount++
	counter.DispatchCount++
}

func (c *VulkanCommandBuffer) CmdSetDeviceMask(deviceMask uint32) {
	c.Driver().VkCmdSetDeviceMask(c.Handle(), driver.Uint32(deviceMask))
	c.CommandCounter().CommandCount++
}
