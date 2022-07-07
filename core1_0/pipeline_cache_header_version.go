package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"

const (
	PipelineCacheHeaderVersionOne PipelineCacheHeaderVersion = C.VK_PIPELINE_CACHE_HEADER_VERSION_ONE
)

func init() {
	PipelineCacheHeaderVersionOne.Register("APIVersion 1")
}
