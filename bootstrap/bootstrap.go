package bootstrap

import (
	"unsafe"

	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/internal/impl1_2"
	"github.com/vkngwrapper/core/v3/loader"
)

func CreateDriverFromProcAddr(procAddr unsafe.Pointer) (core1_0.GlobalDriver, error) {
	loaderObj, err := loader.CreateLoaderFromProcAddr(procAddr)
	if err != nil {
		return nil, err
	}

	return &impl1_0.GlobalVulkanDriver{
		LoaderObj:             loaderObj,
		InstanceDriverFactory: buildInstanceDriver,
		DeviceDriverFactory:   buildDeviceDriver,
	}, nil
}

func buildInstanceDriver(driver *impl1_0.GlobalVulkanDriver, instance core.Instance) (core1_0.CoreInstanceDriver, error) {
	loaderObj, err := driver.LoaderObj.CreateInstanceLoader(instance.Handle())
	if err != nil {
		return nil, err
	}

	switch instance.APIVersion() {
	case common.Vulkan1_2:
		return &impl1_2.InstanceVulkanDriver{
			InstanceVulkanDriver: impl1_1.InstanceVulkanDriver{
				InstanceVulkanDriver: impl1_0.InstanceVulkanDriver{
					GlobalVulkanDriver: impl1_0.GlobalVulkanDriver{
						LoaderObj:             loaderObj,
						InstanceDriverFactory: buildInstanceDriver,
						DeviceDriverFactory:   buildDeviceDriver,
					},
				},
			},
		}, nil
	case common.Vulkan1_1:
		return &impl1_1.InstanceVulkanDriver{
			InstanceVulkanDriver: impl1_0.InstanceVulkanDriver{
				GlobalVulkanDriver: impl1_0.GlobalVulkanDriver{
					LoaderObj:             loaderObj,
					InstanceDriverFactory: buildInstanceDriver,
					DeviceDriverFactory:   buildDeviceDriver,
				},
			},
		}, nil
	default:
		return &impl1_0.InstanceVulkanDriver{
			GlobalVulkanDriver: impl1_0.GlobalVulkanDriver{
				LoaderObj:             loaderObj,
				InstanceDriverFactory: buildInstanceDriver,
				DeviceDriverFactory:   buildDeviceDriver,
			},
		}, nil
	}
}

func buildDeviceDriver(driver *impl1_0.GlobalVulkanDriver, device core.Device) (core1_0.CoreDeviceDriver, error) {
	loaderObj, err := driver.LoaderObj.CreateDeviceLoader(device.Handle())
	if err != nil {
		return nil, err
	}

	switch device.APIVersion() {
	case common.Vulkan1_2:
		return &impl1_2.CoreVulkanDriver{
			InstanceVulkanDriver: impl1_2.InstanceVulkanDriver{
				InstanceVulkanDriver: impl1_1.InstanceVulkanDriver{
					InstanceVulkanDriver: impl1_0.InstanceVulkanDriver{
						GlobalVulkanDriver: impl1_0.GlobalVulkanDriver{
							LoaderObj:             loaderObj,
							InstanceDriverFactory: buildInstanceDriver,
							DeviceDriverFactory:   buildDeviceDriver,
						},
					},
				},
			},
			DeviceVulkanDriver: impl1_2.DeviceVulkanDriver{
				DeviceVulkanDriver: impl1_1.DeviceVulkanDriver{
					DeviceVulkanDriver: impl1_0.DeviceVulkanDriver{
						LoaderObj: loaderObj,
					},
				},
			},
		}, nil
	case common.Vulkan1_1:
		return &impl1_1.CoreVulkanDriver{
			InstanceVulkanDriver: impl1_1.InstanceVulkanDriver{
				InstanceVulkanDriver: impl1_0.InstanceVulkanDriver{
					GlobalVulkanDriver: impl1_0.GlobalVulkanDriver{
						LoaderObj:             loaderObj,
						InstanceDriverFactory: buildInstanceDriver,
						DeviceDriverFactory:   buildDeviceDriver,
					},
				},
			},
			DeviceVulkanDriver: impl1_1.DeviceVulkanDriver{
				DeviceVulkanDriver: impl1_0.DeviceVulkanDriver{
					LoaderObj: loaderObj,
				},
			},
		}, nil
	default:
		return &impl1_0.CoreVulkanDriver{
			InstanceVulkanDriver: impl1_0.InstanceVulkanDriver{
				GlobalVulkanDriver: impl1_0.GlobalVulkanDriver{
					LoaderObj:             loaderObj,
					InstanceDriverFactory: buildInstanceDriver,
					DeviceDriverFactory:   buildDeviceDriver,
				},
			},
			DeviceVulkanDriver: impl1_0.DeviceVulkanDriver{
				LoaderObj: loaderObj,
			},
		}, nil
	}
}
