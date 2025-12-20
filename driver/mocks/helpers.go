package mock_driver

import (
	"github.com/vkngwrapper/core/v3/common"
	gomock "go.uber.org/mock/gomock"
)

func DriverForVersion(ctrl *gomock.Controller, version common.APIVersion) *MockDriver {
	mockDriver := NewMockDriver(ctrl)
	mockDriver.EXPECT().Version().Return(version).AnyTimes()

	return mockDriver
}
