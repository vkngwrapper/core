package common

type MemoryType struct {
	Properties MemoryProperties
	HeapIndex  int
}

type MemoryHeap struct {
	Size  uint64
	Flags MemoryHeapFlags
}
