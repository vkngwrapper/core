package core1_0

// ImageSubresource specifies an Image subresource
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource.html
type ImageSubresource struct {
	// AspectMask selects the Image aspect
	AspectMask ImageAspectFlags
	// MipLevel selects the mipmap level
	MipLevel uint32
	// ArrayLayer selects the array layer
	ArrayLayer uint32
}

// ImageSubresourceLayers specifies an Image subresource layers
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html
type ImageSubresourceLayers struct {
	// AspectMask selects the color, depth, and/or stencil aspects to be copied
	AspectMask ImageAspectFlags
	// MipLevel is the mipmap level to copy
	MipLevel int
	// BaseArrayLayer is the starting layer to copy
	BaseArrayLayer int
	// LayerCount is the number of layers to copy
	LayerCount int
}

// SubresourceLayout specifies the subresource layer
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout.html
type SubresourceLayout struct {
	// Offset is the byte offset from the start of the Image or the plane where the Image
	// subresource begins
	Offset int
	// Size is the size in bytes of the image subresource
	Size int
	// RowPitch describes the number of bytes between each row of texels in an Image
	RowPitch int
	// ArrayPitch describes the number of bytes between each layer of an Image
	ArrayPitch int
	// DepthPitch describes the number of bytes between each slice of a 3D image
	DepthPitch int
}

// ImageSubresourceRange specifies an Image subresource range
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html
type ImageSubresourceRange struct {
	// AspectMask specifies which aspect(s) of the Image are included in the view
	AspectMask ImageAspectFlags
	// BaseMipLevel is the first mipmap level accessible to the view
	BaseMipLevel int
	// LevelCount is the number of mipmap levels accessible to the view
	LevelCount int
	// BaseArrayLayer is the first array layer accessible to the view
	BaseArrayLayer int
	// LayerCount is the number of array layers accessible to the view
	LayerCount int
}
