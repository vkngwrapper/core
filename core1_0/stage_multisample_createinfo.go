package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

type PipelineMultisampleStateCreateFlags uint32

var pipelineMultisampleStateCreateFlagsMapping = common.NewFlagStringMapping[PipelineMultisampleStateCreateFlags]()

func (f PipelineMultisampleStateCreateFlags) Register(str string) {
	pipelineMultisampleStateCreateFlagsMapping.Register(f, str)
}

func (f PipelineMultisampleStateCreateFlags) String() string {
	return pipelineMultisampleStateCreateFlagsMapping.FlagsToString(f)
}

////

type PipelineMultisampleStateCreateInfo struct {
	Flags                PipelineMultisampleStateCreateFlags
	RasterizationSamples SampleCountFlags

	SampleShadingEnable bool
	MinSampleShading    float32
	SampleMask          []uint32

	AlphaToCoverageEnable bool
	AlphaToOneEnable      bool

	common.NextOptions
}

func (o PipelineMultisampleStateCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineMultisampleStateCreateInfo)
	}
	createInfo := (*C.VkPipelineMultisampleStateCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO
	createInfo.flags = C.VkPipelineMultisampleStateCreateFlags(o.Flags)
	createInfo.pNext = next
	createInfo.rasterizationSamples = C.VkSampleCountFlagBits(o.RasterizationSamples)
	createInfo.sampleShadingEnable = C.VK_FALSE
	createInfo.alphaToCoverageEnable = C.VK_FALSE
	createInfo.alphaToOneEnable = C.VK_FALSE

	if o.SampleShadingEnable {
		createInfo.sampleShadingEnable = C.VK_TRUE
	}

	if o.AlphaToCoverageEnable {
		createInfo.alphaToCoverageEnable = C.VK_TRUE
	}

	if o.AlphaToOneEnable {
		createInfo.alphaToOneEnable = C.VK_TRUE
	}

	createInfo.minSampleShading = C.float(o.MinSampleShading)
	createInfo.pSampleMask = nil

	if len(o.SampleMask) > 0 {
		sampleCount := o.RasterizationSamples.Count()
		maskSize := sampleCount / 32
		if sampleCount%32 != 0 {
			maskSize++
		}

		if len(o.SampleMask) != maskSize {
			return nil, errors.Newf("expected a sample mask size of %d, because %d rasterization samples were specified- however, received a sample mask size of %d", maskSize, sampleCount, len(o.SampleMask))
		}

		sampleMaskPtr := (*C.VkSampleMask)(allocator.Malloc(maskSize * int(unsafe.Sizeof(C.VkSampleMask(0)))))
		sampleMaskSlice := ([]C.VkSampleMask)(unsafe.Slice(sampleMaskPtr, maskSize))
		for i := 0; i < maskSize; i++ {
			sampleMaskSlice[i] = C.VkSampleMask(o.SampleMask[i])
		}

		createInfo.pSampleMask = sampleMaskPtr
	}

	return preallocatedPointer, nil
}
