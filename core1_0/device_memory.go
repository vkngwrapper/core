package core1_0

// MemoryType specifies memory type
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryType.html
type MemoryType struct {
	// PropertyFlags specifies properties for this memory type
	PropertyFlags MemoryPropertyFlags
	// HeapIndex describes which memory heap this memory type corresponds to
	HeapIndex int
}

// MemoryHeap specifies a memory heap
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryHeap.html
type MemoryHeap struct {
	// Size is the total memory size in bytes in the heap
	Size int
	// Flags specifies attribute flags for the heap
	Flags MemoryHeapFlags
}
