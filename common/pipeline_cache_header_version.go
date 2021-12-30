package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type PipelineCacheHeaderVersion int32

const (
	PipelineCacheHeaderVersion1 PipelineCacheHeaderVersion = C.VK_PIPELINE_CACHE_HEADER_VERSION_ONE
)

var pipelineCacheHeaderVersionToString = map[PipelineCacheHeaderVersion]string{
	PipelineCacheHeaderVersion1: "Version 1",
}

func (v PipelineCacheHeaderVersion) String() string {
	return pipelineCacheHeaderVersionToString[v]
}