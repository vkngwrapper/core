package core1_0

// MemoryRequirements specifies memory requirements
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html
type MemoryRequirements struct {
	// Size is the size, in bytes, of the memory allocation required for the resource
	Size int
	// Alignment is the alignment, in bytes, of the offset within the allocation required
	// for the resource
	Alignment int
	// MemoryTypeBits is a bitmask and contains one bit set for every supported memory type
	// for the resource. Bit i is set if and only if the memory type i in PhysicalDeviceMemoryProperties
	// for the PhysicalDevice is supported for the resource
	MemoryTypeBits uint32
}
