package core1_0

// Extent2D specifies a two-dimensional extent
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html
type Extent2D struct {
	// Width is the width of the extent
	Width int
	// Height is the height of the extent
	Height int
}

// Extent3D specifies a three-dimensional extent
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent3D.html
type Extent3D struct {
	// Width is the width of the extent
	Width int
	// Height is the height of the extent
	Height int
	// Depth is the depth of the extent
	Depth int
}
