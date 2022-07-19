package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"

const (
	// PipelineCacheHeaderVersionOne specifies version 1 of the PipelineCache
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheHeaderVersion.html
	PipelineCacheHeaderVersionOne PipelineCacheHeaderVersion = C.VK_PIPELINE_CACHE_HEADER_VERSION_ONE
)

func init() {
	PipelineCacheHeaderVersionOne.Register("APIVersion 1")
}
