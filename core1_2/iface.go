package core1_2

import (
	"time"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
)

//go:generate mockgen -source ./iface.go -destination ../mocks/mocks1_2/mocks.go -package mocks1_2

type CoreInstanceDriver interface {
	core1_1.CoreInstanceDriver
}

type DeviceDriver interface {
	core1_1.DeviceDriver

	// CmdBeginRenderPass2 begins a new RenderPass
	//
	// commandBuffer - the CommandBuffer to record to
	//
	// renderPassBegin - Specifies the RenderPass to begin an instance of, and the Framebuffer the instance
	// uses
	//
	// subpassBegin - Contains information about the subpass which is about to begin rendering
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass2.html
	CmdBeginRenderPass2(commandBuffer core1_0.CommandBuffer, renderPassBegin core1_0.RenderPassBeginInfo, subpassBegin SubpassBeginInfo) error
	// CmdEndRenderPass2 ends the current RenderPass
	//
	// commandBuffer - the CommandBuffer to record to
	//
	// subpassEnd - Contains information about how the previous subpass will be ended
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdEndRenderPass2.html
	CmdEndRenderPass2(commandBuffer core1_0.CommandBuffer, subpassEnd SubpassEndInfo) error
	// CmdNextSubpass2 transitions to the next subpass of a RenderPass
	//
	// commandBuffer - the CommandBuffer to record to
	//
	// subpassBegin - Contains information about the subpass which is about to begin rendering.
	//
	// subpassEnd - Contains information about how the previous subpass will be ended.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass2.html
	CmdNextSubpass2(commandBuffer core1_0.CommandBuffer, subpassBegin SubpassBeginInfo, subpassEnd SubpassEndInfo) error
	// CmdDrawIndexedIndirectCount draws with indirect parameters, indexed vertices, and draw count
	//
	// commandBuffer - the CommandBuffer to record to
	//
	// buffer - The Buffer containing draw parameters
	//
	// offset - The byte offset into buffer where parameters begin
	//
	// countBuffer - The Buffer containing the draw count
	//
	// countBufferOffset - The byte offset into countBuffer where the draw count begins
	//
	// maxDrawCount - Specifies the maximum number of draws that will be executed.
	//
	// stride - The byte stride between successive sets of draw parameters
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexedIndirectCount.html
	CmdDrawIndexedIndirectCount(commandBuffer core1_0.CommandBuffer, buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int)
	// CmdDrawIndirectCount draws primitives with indirect parameters and draw count
	//
	// commandBuffer - the CommandBuffer to record to
	//
	// buffer - The Buffer containing draw parameters
	//
	// offset - The byte offset into buffer where parameters begin
	//
	// countBuffer - The Buffer containing the draw count
	//
	// countBufferOffset - The byte offset into countBuffer where the draw count begins
	//
	// maxDrawCount - Specifies the maximum number of draws that will be executed.
	//
	// stride - The byte stride between successive sets of draw parameters
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndirectCount.html
	CmdDrawIndirectCount(commandBuffer core1_0.CommandBuffer, buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int)

	// CreateRenderPass2 creates a new RenderPass object
	//
	// device - The Device used to create the RenderPass
	//
	// allocator - Controls host memory allocation behavior
	//
	// options - Describes the parameters of the RenderPass
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass2.html
	CreateRenderPass2(allocator *loader.AllocationCallbacks, options RenderPassCreateInfo2) (core1_0.RenderPass, common.VkResult, error)

	// GetBufferDeviceAddress queries an address of a Buffer
	//
	// o - Specifies the Buffer to retrieve an address for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html
	GetBufferDeviceAddress(o BufferDeviceAddressInfo) (uint64, error)
	// GetBufferOpaqueCaptureAddress queries an opaque capture address of a Buffer
	//
	// o - Specifies the Buffer to retrieve an address for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetBufferOpaqueCaptureAddress.html
	GetBufferOpaqueCaptureAddress(o BufferDeviceAddressInfo) (uint64, error)
	// GetDeviceMemoryOpaqueCaptureAddress queries an opaque capture address of a DeviceMemory object
	//
	// o - Specifies the DeviceMemory object to retrieve an address for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMemoryOpaqueCaptureAddress.html
	GetDeviceMemoryOpaqueCaptureAddress(o DeviceMemoryOpaqueCaptureAddressInfo) (uint64, error)

	// SignalSemaphore signals a timeline Semaphore on the host
	//
	// o - Contains information about the signal operation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkSignalSemaphore.html
	SignalSemaphore(o SemaphoreSignalInfo) (common.VkResult, error)
	// WaitSemaphores waits for timeline Semaphore objects on the host
	//
	// timeout - How long to wait before returning VKTimeout. May be common.NoTimeout to wait indefinitely.
	// The timeout is adjusted to the closest value allowed by the implementation timeout accuracy,
	// which may be substantially longer than the requested timeout.
	//
	// o - Contains information about the wait condition
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkWaitSemaphores.html
	WaitSemaphores(timeout time.Duration, o SemaphoreWaitInfo) (common.VkResult, error)

	// ResetQueryPool resets queries in this QueryPool
	//
	// queryPool - The QueryPool to reset
	//
	// firstQuery - The initial query index to reset
	//
	// queryCount - The number of queries to reset
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetQueryPool.html
	ResetQueryPool(queryPool core1_0.QueryPool, firstQuery, queryCount int)

	// GetSemaphoreCounterValue queries the current state of a timeline Semaphore
	//
	// semaphore - The timeline Semaphore to query
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreCounterValue.html
	GetSemaphoreCounterValue(semaphore core1_0.Semaphore) (uint64, common.VkResult, error)
}

type CoreDeviceDriver interface {
	InstanceDriver() core1_0.CoreInstanceDriver
	DeviceDriver
}
