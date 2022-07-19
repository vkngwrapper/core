package core1_0

// ExtensionProperties specifies an extension's properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtensionProperties.html
type ExtensionProperties struct {
	// ExtensionName is a string which is the name of the extension
	ExtensionName string
	// SpecVersion is the version of this extension
	SpecVersion uint
}
