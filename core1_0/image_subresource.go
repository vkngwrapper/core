package core1_0

type ImageSubresource struct {
	AspectMask ImageAspectFlags
	MipLevel   uint32
	ArrayLayer uint32
}

type ImageSubresourceLayers struct {
	AspectMask     ImageAspectFlags
	MipLevel       int
	BaseArrayLayer int
	LayerCount     int
}

type SubresourceLayout struct {
	Offset     int
	Size       int
	RowPitch   int
	ArrayPitch int
	DepthPitch int
}

type ImageSubresourceRange struct {
	AspectMask     ImageAspectFlags
	BaseMipLevel   int
	LevelCount     int
	BaseArrayLayer int
	LayerCount     int
}
