package render_pass

import "github.com/CannibalVox/VKng/core/loader"

type Framebuffer interface {
	Handle() loader.VkFramebuffer
	Destroy() error
}

type RenderPass interface {
	Handle() loader.VkRenderPass
	Destroy() error
}
