package core1_2

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"unsafe"
)

// SamplerReductionMode specifies reduction mode for texture filtering
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionMode.html
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
	// SamplerAddressModeMirrorClampToEdge specifies that the mirror clamp to edge wrap mode will
	// be used
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html
	SamplerAddressModeMirrorClampToEdge core1_0.SamplerAddressMode = C.VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE
)

func init() {
	SamplerAddressModeMirrorClampToEdge.Register("Mirror Clamp To Edge")
}

////

const (
	// FormatFeatureSampledImageFilterMinmax specifies the Image can be used as a sampled Image
	// with a min or max SamplerReductionMode
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureSampledImageFilterMinmax core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT

	// SamplerReductionModeMax specifies that texel values are combined by taking
	// the component-wise maximum of values in the footprint with non-zero weights
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionMode.html
	SamplerReductionModeMax SamplerReductionMode = C.VK_SAMPLER_REDUCTION_MODE_MAX
	// SamplerReductionModeMin specifies that texel values are combined by taking the
	// component-wise minimum of values in the footprint with non-zero weights
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionMode.html
	SamplerReductionModeMin SamplerReductionMode = C.VK_SAMPLER_REDUCTION_MODE_MIN
	// SamplerReductionModeWeightedAverage specifies that texel values are combined by
	// computing a weighted average of values in the footprint
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionMode.html
	SamplerReductionModeWeightedAverage SamplerReductionMode = C.VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE
)

func init() {
	FormatFeatureSampledImageFilterMinmax.Register("Sampled Image Filter Min-Max")

	SamplerReductionModeMin.Register("Min")
	SamplerReductionModeMax.Register("Max")
	SamplerReductionModeWeightedAverage.Register("Weighted Average")
}

////

// SamplerReductionModeCreateInfo specifies a Sampler reduction mode
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionModeCreateInfoEXT.html
type SamplerReductionModeCreateInfo struct {
	// ReductionMode controls how texture filtering combines texel values
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
