package mock_driver

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/golang/mock/gomock"
)

func DriverForVersion(ctrl *gomock.Controller, version common.APIVersion) *MockDriver {
	mockDriver := NewMockDriver(ctrl)
	mockDriver.EXPECT().Version().Return(version).AnyTimes()
	mockDriver.EXPECT().ObjectStore().Return(driver.NewObjectStore()).AnyTimes()

	return mockDriver
}
