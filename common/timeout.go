package common

import "C"
import "time"

// NoTimeout can be passed to any Vulkan command or field which accepts a timeout in order
// to indicate that the method should take as long as necessary.
const NoTimeout = time.Duration(^int64(0))

// TimeoutNanoseconds converts a time.Duration into a timeout value acceptable to
// any Vulkan command or field which accepts a timeout.
func TimeoutNanoseconds(timeout time.Duration) uint64 {
	if timeout == NoTimeout {
		return ^uint64(0)
	}

	return uint64(timeout.Nanoseconds())
}
