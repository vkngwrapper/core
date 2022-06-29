package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

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
	TessellationDomainOriginUpperLeft TessellationDomainOrigin = C.VK_TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT
	TessellationDomainOriginLowerLeft TessellationDomainOrigin = C.VK_TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT

	PipelineCreateDispatchBase             core1_0.PipelineCreateFlags = C.VK_PIPELINE_CREATE_DISPATCH_BASE
	PipelineCreateViewIndexFromDeviceIndex core1_0.PipelineCreateFlags = C.VK_PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT
)

func init() {
	TessellationDomainOriginUpperLeft.Register("Upper Left")
	TessellationDomainOriginLowerLeft.Register("Lower Left")

	PipelineCreateDispatchBase.Register("Dispatch Base")
	PipelineCreateViewIndexFromDeviceIndex.Register("View Index From Device Index")
}

////

type PipelineTessellationDomainOriginStateOptions struct {
	DomainOrigin TessellationDomainOrigin
	common.HaveNext
}

func (o PipelineTessellationDomainOriginStateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPipelineTessellationDomainOriginStateCreateInfo{})))
	}

	createInfo := (*C.VkPipelineTessellationDomainOriginStateCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO
	createInfo.pNext = next
	createInfo.domainOrigin = (C.VkTessellationDomainOriginKHR)(o.DomainOrigin)

	return preallocatedPointer, nil
}

func (o PipelineTessellationDomainOriginStateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkPipelineTessellationDomainOriginStateCreateInfo)(cDataPointer)
	return createInfo.pNext, nil
}
