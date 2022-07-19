package core1_0

import "github.com/vkngwrapper/core/common"

// LayerProperties specifies layer properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerProperties.html
type LayerProperties struct {
	// LayerName is a string which is the name of the layer
	LayerName string
	// SpecVersion is the Vulkan version the layer was written to
	SpecVersion common.Version
	// ImplementationVersion is the version of this layer
	ImplementationVersion common.Version
	// Description is a string which provides additional details that can be used by the
	// application to identify the layer
	Description string
}
