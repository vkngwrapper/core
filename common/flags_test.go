package common_test

import (
	"fmt"
	"github.com/vkngwrapper/core/common"
)

type CoolFlags int32

var coolFlagsMapping = common.NewFlagStringMapping[CoolFlags]()

func (f CoolFlags) Register(str string) {
	coolFlagsMapping.Register(f, str)
}
func (f CoolFlags) String() string {
	return coolFlagsMapping.FlagsToString(f)
}

const (
	CoolFlagsRed  CoolFlags = 1
	CoolFlagsBlue CoolFlags = 2
)

func init() {
	CoolFlagsRed.Register("Red")
	CoolFlagsBlue.Register("Blue")
}
func ExampleFlagStringMapping() {
	fmt.Println(CoolFlagsBlue | CoolFlagsRed)

	// Output: Red|Blue
}
