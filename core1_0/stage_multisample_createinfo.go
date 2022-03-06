package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type MultisampleOptions struct {
	RasterizationSamples common.SampleCounts

	SampleShading    bool
	MinSampleShading float32
	SampleMask       []uint32

	AlphaToCoverage bool
	AlphaToOne      bool

	core.HaveNext
}

func (o MultisampleOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineMultisampleStateCreateInfo)
	}
	createInfo := (*C.VkPipelineMultisampleStateCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next
	createInfo.rasterizationSamples = C.VkSampleCountFlagBits(o.RasterizationSamples)
	createInfo.sampleShadingEnable = C.VK_FALSE
	createInfo.alphaToCoverageEnable = C.VK_FALSE
	createInfo.alphaToOneEnable = C.VK_FALSE

	if o.SampleShading {
		createInfo.sampleShadingEnable = C.VK_TRUE
	}

	if o.AlphaToCoverage {
		createInfo.alphaToCoverageEnable = C.VK_TRUE
	}

	if o.AlphaToOne {
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
