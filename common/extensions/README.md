# core/common/extensions

This package contains utility methods to be used by developers writing wrappers for
 Vulkan extensions. If you are writing an applications, these should not be used.

If an extension creates a new Vulkan object in a method, you can retrieve the handles
 from your call to a vulkan driver method and then pass them to these methods in order
 to get the vkngwrapper object wrapping the new Vulkan object.
