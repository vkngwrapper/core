package mocks

import (
	"fmt"
	"github.com/golang/mock/gomock"
)

type exactMatcher struct {
	targetPtr interface{}
}

func (m exactMatcher) Matches(x interface{}) bool {
	return m.targetPtr == x
}

func (m exactMatcher) String() string {
	return fmt.Sprintf("pointer %p", m.targetPtr)
}

func Exactly(target interface{}) gomock.Matcher {
	return exactMatcher{targetPtr: target}
}
