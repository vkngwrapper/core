package core1_0

type Viewport struct {
	X        float32
	Y        float32
	Width    float32
	Height   float32
	MinDepth float32
	MaxDepth float32
}

type Rect2D struct {
	Offset Offset2D
	Extent Extent2D
}
