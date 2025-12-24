package impl1_1

import (
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *DeviceVulkanDriver) CmdDispatchBase(commandBuffer types.CommandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ int) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}
	v.LoaderObj.VkCmdDispatchBase(commandBuffer.Handle(),
		loader.Uint32(baseGroupX),
		loader.Uint32(baseGroupY),
		loader.Uint32(baseGroupZ),
		loader.Uint32(groupCountX),
		loader.Uint32(groupCountY),
		loader.Uint32(groupCountZ))
}

func (v *DeviceVulkanDriver) CmdSetDeviceMask(commandBuffer types.CommandBuffer, deviceMask uint32) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}
	v.LoaderObj.VkCmdSetDeviceMask(commandBuffer.Handle(), loader.Uint32(deviceMask))
}
