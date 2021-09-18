package common

type PushConstantRange struct {
	Stages ShaderStages
	Offset uint32
	Size   uint32
}
