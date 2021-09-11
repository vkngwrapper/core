package render_pass

import "github.com/CannibalVox/VKng/core/loader"

//go:generate mockgen -source iface.go -destination ../mocks/render_pass.go -package=mocks

type Framebuffer interface {
	Handle() loader.VkFramebuffer
	Destroy() error
}

type RenderPass interface {
	Handle() loader.VkRenderPass
	Destroy() error
}
