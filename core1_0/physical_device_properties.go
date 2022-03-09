package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/google/uuid"
)

const (
	QueueGraphics      common.QueueFlags = C.VK_QUEUE_GRAPHICS_BIT
	QueueCompute       common.QueueFlags = C.VK_QUEUE_COMPUTE_BIT
	QueueTransfer      common.QueueFlags = C.VK_QUEUE_TRANSFER_BIT
	QueueSparseBinding common.QueueFlags = C.VK_QUEUE_SPARSE_BINDING_BIT

	MemoryDeviceLocal     common.MemoryProperties = C.VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT
	MemoryHostVisible     common.MemoryProperties = C.VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT
	MemoryHostCoherent    common.MemoryProperties = C.VK_MEMORY_PROPERTY_HOST_COHERENT_BIT
	MemoryLazilyAllocated common.MemoryProperties = C.VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT

	HeapDeviceLocal common.MemoryHeapFlags = C.VK_MEMORY_HEAP_DEVICE_LOCAL_BIT

	DeviceOther         common.PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_OTHER
	DeviceIntegratedGPU common.PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU
	DeviceDiscreteGPU   common.PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU
	DeviceVirtualGPU    common.PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU
	DeviceCPU           common.PhysicalDeviceType = C.VK_PHYSICAL_DEVICE_TYPE_CPU
)

func init() {
	QueueGraphics.Register("Graphics")
	QueueCompute.Register("Compute")
	QueueTransfer.Register("Transfer")
	QueueSparseBinding.Register("Sparse Binding")

	MemoryDeviceLocal.Register("Device Local")
	MemoryHostVisible.Register("Host Visible")
	MemoryHostCoherent.Register("Host Coherent")
	MemoryLazilyAllocated.Register("Lazily Allocated")

	HeapDeviceLocal.Register("Device Local")

	DeviceOther.Register("Other")
	DeviceIntegratedGPU.Register("Integrated GPU")
	DeviceDiscreteGPU.Register("Discrete GPU")
	DeviceVirtualGPU.Register("Virtual GPU")
	DeviceCPU.Register("CPU")
}

type PhysicalDeviceSparseProperties struct {
	ResidencyStandard2DBlockShape            bool
	ResidencyStandard2DMultisampleBlockShape bool
	ResidencyStandard3DBlockShape            bool
	ResidencyAlignedMipSize                  bool
	ResidencyNonResidentStrict               bool
}

type PhysicalDeviceProperties struct {
	Type common.PhysicalDeviceType
	Name string

	APIVersion    common.APIVersion
	DriverVersion common.Version
	VendorID      uint32
	DeviceID      uint32

	PipelineCacheUUID uuid.UUID
	Limits            *PhysicalDeviceLimits
	SparseProperties  *PhysicalDeviceSparseProperties
}

type QueueFamily struct {
	Flags                       common.QueueFlags
	QueueCount                  uint32
	TimestampValidBits          uint32
	MinImageTransferGranularity common.Extent3D
}

type PhysicalDeviceMemoryProperties struct {
	MemoryTypes []common.MemoryType
	MemoryHeaps []common.MemoryHeap
}
