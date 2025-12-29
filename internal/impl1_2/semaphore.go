package impl1_2

import (
	"fmt"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) GetSemaphoreCounterValue(semaphore core1_0.Semaphore) (uint64, common.VkResult, error) {
	if !semaphore.Initialized() {
		return 0, core1_0.VKErrorUnknown, fmt.Errorf("semaphore cannot be uninitialized")
	}
	var value loader.Uint64
	res, err := v.LoaderObj.VkGetSemaphoreCounterValue(
		semaphore.DeviceHandle(),
		semaphore.Handle(),
		&value,
	)
	if err != nil {
		return 0, res, err
	}

	return uint64(value), res, nil
}
