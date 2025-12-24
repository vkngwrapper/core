package mocks1_0

import (
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

func InternalGlobalDriver(loader loader.Loader) core1_0.GlobalDriver {
	return impl1_0.NewGlobalDriver(loader)
}

func InternalCoreInstanceDriver(loader loader.Loader) core1_0.CoreInstanceDriver {
	return impl1_0.NewInstanceDriver(loader)
}

func InternalDeviceDriver(loader loader.Loader) core1_0.DeviceDriver {
	return impl1_0.NewDeviceDriver(loader)
}

func InternalCoreDriver(loader loader.Loader) core1_0.CoreDeviceDriver {
	return impl1_0.NewCoreDriver(loader)
}
