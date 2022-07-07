package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

type PeerMemoryFeatureFlags int32

var peerMemoryFeaturesMapping = common.NewFlagStringMapping[PeerMemoryFeatureFlags]()

func (f PeerMemoryFeatureFlags) Register(str string) {
	peerMemoryFeaturesMapping.Register(f, str)
}
func (f PeerMemoryFeatureFlags) String() string {
	return peerMemoryFeaturesMapping.FlagsToString(f)
}

////

const (
	PeerMemoryFeatureCopyDst    PeerMemoryFeatureFlags = C.VK_PEER_MEMORY_FEATURE_COPY_DST_BIT
	PeerMemoryFeatureCopySrc    PeerMemoryFeatureFlags = C.VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT
	PeerMemoryFeatureGenericDst PeerMemoryFeatureFlags = C.VK_PEER_MEMORY_FEATURE_GENERIC_DST_BIT
	PeerMemoryFeatureGenericSrc PeerMemoryFeatureFlags = C.VK_PEER_MEMORY_FEATURE_GENERIC_SRC_BIT

	QueueFamilyExternal int = C.VK_QUEUE_FAMILY_EXTERNAL

	DependencyDeviceGroup core1_0.DependencyFlags = C.VK_DEPENDENCY_DEVICE_GROUP_BIT
	DependencyViewLocal   core1_0.DependencyFlags = C.VK_DEPENDENCY_VIEW_LOCAL_BIT
)

func init() {
	PeerMemoryFeatureCopyDst.Register("Copy Dst")
	PeerMemoryFeatureCopySrc.Register("Copy Src")
	PeerMemoryFeatureGenericDst.Register("Generic Dst")
	PeerMemoryFeatureGenericSrc.Register("Generic Src")

	DependencyDeviceGroup.Register("Device Group")
	DependencyViewLocal.Register("View Local")
}

////

type DescriptorSetLayoutSupport struct {
	Supported bool

	common.NextOutData
}

func (o *DescriptorSetLayoutSupport) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDescriptorSetLayoutSupport{})))
	}

	outData := (*C.VkDescriptorSetLayoutSupport)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *DescriptorSetLayoutSupport) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkDescriptorSetLayoutSupport)(cDataPointer)
	o.Supported = outData.supported != C.VkBool32(0)

	return outData.pNext, nil
}

////
