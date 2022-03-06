package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	internal_core1_0 "github.com/CannibalVox/VKng/core/internal/core1_0"
)

type VulkanBufferView struct {
	internal_core1_0.VulkanBufferView
}

func PromoteBufferView(view core1_0.BufferView) core1_1.BufferView {
	goodView, ok := view.(core1_1.BufferView)
	if ok {
		return goodView
	}

	oldVulkanView, ok := view.(*internal_core1_0.VulkanBufferView)
	if ok && oldVulkanView.MaximumAPIVersion.IsAtLeast(common.Vulkan1_1) {
		return &VulkanBufferView{*oldVulkanView}
	}

	return nil
}
