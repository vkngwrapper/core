package impl1_0

import (
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

type Vulkan struct {
	Driver driver.Driver
}

var _ core1_0.Loader = &Vulkan{}
