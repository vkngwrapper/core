package core1_0

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/google/uuid"
)

type PhysicalDeviceSparseProperties struct {
	ResidencyStandard2DBlockShape            bool
	ResidencyStandard2DMultisampleBlockShape bool
	ResidencyStandard3DBlockShape            bool
	ResidencyAlignedMipSize                  bool
	ResidencyNonResidentStrict               bool
}

type PhysicalDeviceProperties struct {
	Type PhysicalDeviceType
	Name string

	APIVersion    common.APIVersion
	DriverVersion common.Version
	VendorID      uint32
	DeviceID      uint32

	PipelineCacheUUID uuid.UUID
	Limits            *PhysicalDeviceLimits
	SparseProperties  *PhysicalDeviceSparseProperties
}
