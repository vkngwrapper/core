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

// PeerMemoryFeatureFlags specifies supported peer memory features
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlagBits.html
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

	// PeerMemoryFeatureCopyDst specifies that the memory can be accessed as the destination of
	// any CommandBuffer.CmdCopy... command
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlagBits.html
	PeerMemoryFeatureCopyDst PeerMemoryFeatureFlags = C.VK_PEER_MEMORY_FEATURE_COPY_DST_BIT
	// PeerMemoryFeatureCopySrc specifies that the memory can be accessed as the source of any
	// CommandBuffer.CmdCopy... command
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlagBits.html
	PeerMemoryFeatureCopySrc PeerMemoryFeatureFlags = C.VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT
	// PeerMemoryFeatureGenericDst specifies that the memory can be written as any memory access type
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlagBits.html
	PeerMemoryFeatureGenericDst PeerMemoryFeatureFlags = C.VK_PEER_MEMORY_FEATURE_GENERIC_DST_BIT
	// PeerMemoryFeatureGenericSrc specifies that the memory can be read as any memory access type
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlagBits.html
	PeerMemoryFeatureGenericSrc PeerMemoryFeatureFlags = C.VK_PEER_MEMORY_FEATURE_GENERIC_SRC_BIT

	// QueueFamilyExternal represents any Queue external to the resource's current Vulkan instance,
	// as long as the Queue uses the same underlying Device group or PhysicalDevice
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QUEUE_FAMILY_EXTERNAL_KHR.html
	QueueFamilyExternal int = C.VK_QUEUE_FAMILY_EXTERNAL

	// DependencyDeviceGroup specifies that dependencies are non-device-local
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlagBits.html
	DependencyDeviceGroup core1_0.DependencyFlags = C.VK_DEPENDENCY_DEVICE_GROUP_BIT
	// DependencyViewLocal specifies that a subpass has more than one view
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlagBits.html
	DependencyViewLocal core1_0.DependencyFlags = C.VK_DEPENDENCY_VIEW_LOCAL_BIT
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

// DescriptorSetLayoutSupport returns information about whether a DescriptorSetLayout can be
// supported
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutSupport.html
type DescriptorSetLayoutSupport struct {
	// Supported specifies whether the DescriptorSetLayout can be created
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
