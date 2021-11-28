package common

type ImageSubresource struct {
	AspectMask ImageAspectFlags
	MipLevel   uint32
	ArrayLayer uint32
}
