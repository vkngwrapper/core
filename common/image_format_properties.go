package common

type ImageFormatProperties struct {
	MaxExtent       Extent3D
	MaxMipLevels    int
	MaxArrayLayers  int
	SampleCounts    SampleCounts
	MaxResourceSize int
}
