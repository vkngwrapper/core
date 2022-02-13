package core1_0

import "github.com/CannibalVox/VKng/core/common"

type ImageFormatProperties struct {
	MaxExtent       common.Extent3D
	MaxMipLevels    int
	MaxArrayLayers  int
	SampleCounts    common.SampleCounts
	MaxResourceSize int
}
