package core1_2

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

type SamplerReductionMode int32

var samplerReductionModeMapping = make(map[SamplerReductionMode]string)

func (e SamplerReductionMode) Register(str string) {
	samplerReductionModeMapping[e] = str
}

func (e SamplerReductionMode) String() string {
	return samplerReductionModeMapping[e]
}

////

const (
	SamplerAddressModeMirrorClampToEdge core1_0.SamplerAddressMode = C.VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE
)

func init() {
	SamplerAddressModeMirrorClampToEdge.Register("Mirror Clamp To Edge")
}

////

const (
	FormatFeatureSampledImageFilterMinmax core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT

	SamplerReductionModeMax             SamplerReductionMode = C.VK_SAMPLER_REDUCTION_MODE_MAX
	SamplerReductionModeMin             SamplerReductionMode = C.VK_SAMPLER_REDUCTION_MODE_MIN
	SamplerReductionModeWeightedAverage SamplerReductionMode = C.VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE
)

func init() {
	FormatFeatureSampledImageFilterMinmax.Register("Sampled Image Filter Min-Max")

	SamplerReductionModeMin.Register("Min")
	SamplerReductionModeMax.Register("Max")
	SamplerReductionModeWeightedAverage.Register("Weighted Average")
}

////

type SamplerReductionModeCreateInfo struct {
	ReductionMode SamplerReductionMode

	common.NextOptions
}

func (o SamplerReductionModeCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSamplerReductionModeCreateInfo{})))
	}

	info := (*C.VkSamplerReductionModeCreateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO
	info.pNext = next
	info.reductionMode = C.VkSamplerReductionModeEXT(o.ReductionMode)

	return preallocatedPointer, nil
}
