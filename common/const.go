package common

/*
#include <stdlib.h>
#include "vulkan.h"
*/
import "C"

const (
	MaxMemoryTypes int = C.VK_MAX_MEMORY_TYPES
	MaxMemoryHeaps int = C.VK_MAX_MEMORY_HEAPS
	WholeSize      int = -1
)
