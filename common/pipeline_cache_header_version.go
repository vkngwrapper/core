package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

const (
	PipelineCacheHeaderVersion1 PipelineCacheHeaderVersion = C.VK_PIPELINE_CACHE_HEADER_VERSION_ONE
)

func init() {
	PipelineCacheHeaderVersion1.Register("Version 1")
}
