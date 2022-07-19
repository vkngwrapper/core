package core1_0

// Viewport specifies a viewport
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewport.html
type Viewport struct {
	// X is the x-coordinate of the viewport's upper-left corner
	X float32
	// Y is the y-coordinate of the viewport's upper-left corner
	Y float32
	// Width is the viewport's width
	Width float32
	// Height is the viewport's height
	Height float32
	// MinDepth is the near depth range for the viewport
	MinDepth float32
	// MaxDepth is the far depth range for the viewport
	MaxDepth float32
}

// Rect2D specifies a two-dimensional subregion
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html
type Rect2D struct {
	// Offset specifies the rectangle offset
	Offset Offset2D
	// Extent specifies the rectangle extent
	Extent Extent2D
}
