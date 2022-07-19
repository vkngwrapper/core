package core1_0

// CommandCounter stores the number of commands, draws, dispatches, etc. executed since the last time
// the CommandBuffer was restarted. These are stored in a struct so that different CommandBuffer
// objects with the same CommandBuffer.Handle can share these counts
type CommandCounter struct {
	CommandCount  int
	DrawCallCount int
	DispatchCount int
}
