package core1_1

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

// TessellationDomainOrigin describes tessellation domain origin
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTessellationDomainOrigin.html
type TessellationDomainOrigin int32

var tessellationDomainOriginMapping = make(map[TessellationDomainOrigin]string)

func (e TessellationDomainOrigin) Register(str string) {
	tessellationDomainOriginMapping[e] = str
}

func (e TessellationDomainOrigin) String() string {
	return tessellationDomainOriginMapping[e]
}

////

const (
	// TessellationDomainOriginUpperLeft specifies that the origin of the domain space
	// is in the upper left corner, as shown in figure
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTessellationDomainOrigin.html
	TessellationDomainOriginUpperLeft TessellationDomainOrigin = C.VK_TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT
	// TessellationDomainOriginLowerLeft specifies that the origin of the domain space
	// is in the lower left corner, as shown in figure
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTessellationDomainOrigin.html
	TessellationDomainOriginLowerLeft TessellationDomainOrigin = C.VK_TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT

	// PipelineCreateDispatchBase specifies that a compute pipeline can be used with
	// CommandBuffer.CmdDispatchBase with a non-zero base workgroup
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html
	PipelineCreateDispatchBase core1_0.PipelineCreateFlags = C.VK_PIPELINE_CREATE_DISPATCH_BASE
	// PipelineCreateViewIndexFromDeviceIndex specifies that any shader input variables
	// decorated as ViewIndex will be assigned values as if they were decorated as DeviceIndex
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html
	PipelineCreateViewIndexFromDeviceIndex core1_0.PipelineCreateFlags = C.VK_PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT
)

func init() {
	TessellationDomainOriginUpperLeft.Register("Upper Left")
	TessellationDomainOriginLowerLeft.Register("Lower Left")

	PipelineCreateDispatchBase.Register("Dispatch Base")
	PipelineCreateViewIndexFromDeviceIndex.Register("View Index From Device Index")
}

////

// PipelineTessellationDomainOriginStateCreateInfo specifies the origin of the tessellation domain
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationDomainOriginStateCreateInfo.html
type PipelineTessellationDomainOriginStateCreateInfo struct {
	// DomainOrigin controls the origin of the tessellation domain space
	DomainOrigin TessellationDomainOrigin

	common.NextOptions
}

func (o PipelineTessellationDomainOriginStateCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPipelineTessellationDomainOriginStateCreateInfo{})))
	}

	createInfo := (*C.VkPipelineTessellationDomainOriginStateCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO
	createInfo.pNext = next
	createInfo.domainOrigin = (C.VkTessellationDomainOriginKHR)(o.DomainOrigin)

	return preallocatedPointer, nil
}
