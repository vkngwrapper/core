package pipeline

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
	"github.com/palantir/stacktrace"
	"unsafe"
)

type MultisampleOptions struct {
	RasterizationSamples core.SampleCounts

	SampleShading    bool
	MinSampleShading float32
	SampleMask       []uint32

	AlphaToCoverage bool
	AlphaToOne      bool

	Next core.Options
}

func (o *MultisampleOptions) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	createInfo := (*C.VkPipelineMultisampleStateCreateInfo)(allocator.Malloc(C.sizeof_struct_VkPipelineMultisampleStateCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO
	createInfo.flags = 0
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
			return nil, stacktrace.NewError("expected a sample mask size of %d, because %d rasterization samples were specified- however, received a sample mask size of %d", maskSize, sampleCount, len(o.SampleMask))
		}

		sampleMaskPtr := (*C.VkSampleMask)(allocator.Malloc(maskSize * int(unsafe.Sizeof(C.VkSampleMask(0)))))
		sampleMaskSlice := ([]C.VkSampleMask)(unsafe.Slice(sampleMaskPtr, maskSize))
		for i := 0; i < maskSize; i++ {
			sampleMaskSlice[i] = C.VkSampleMask(o.SampleMask[i])
		}

		createInfo.pSampleMask = sampleMaskPtr
	}

	var err error
	var next unsafe.Pointer
	if o.Next != nil {
		next, err = o.Next.AllocForC(allocator)
	}

	if err != nil {
		return nil, err
	}
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}
