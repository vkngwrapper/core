package core1_0

// ImageFormatProperties specifies an Image format properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties.html
type ImageFormatProperties struct {
	// MaxExtent are the maximum Image dimensions
	MaxExtent Extent3D
	// MaxMipLevels is the maximum number of mipmap levels
	MaxMipLevels int
	// MaxArrayLayers is the maximum number of array layers
	MaxArrayLayers int
	// SampleCounts specifies all the supported sample counts for this Image
	SampleCounts SampleCountFlags
	// MaxResourceSize is an upper bound on the total Image size in bytes
	MaxResourceSize int
}
