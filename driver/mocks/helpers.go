package mock_driver

import (
	"github.com/golang/mock/gomock"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/driver"
)

func DriverForVersion(ctrl *gomock.Controller, version common.APIVersion) *MockDriver {
	mockDriver := NewMockDriver(ctrl)
	mockDriver.EXPECT().Version().Return(version).AnyTimes()
	mockDriver.EXPECT().ObjectStore().Return(driver.NewObjectStore()).AnyTimes()

	return mockDriver
}
