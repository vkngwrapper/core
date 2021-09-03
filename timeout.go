package core

import "C"
import "time"

const NoTimeout = time.Duration(^int64(0))

func TimeoutNanoseconds(timeout time.Duration) uint64 {
	if timeout == NoTimeout {
		return ^uint64(0)
	}

	return uint64(timeout.Nanoseconds())
}
