package core

type PushConstantRange struct {
	Stages ShaderStages
	Offset uint32
	Size   uint32
}
