package common

type ImageSubresourceRange struct {
	AspectMask     ImageAspectFlags
	BaseMipLevel   int
	LevelCount     int
	BaseArrayLayer int
	LayerCount     int
}
