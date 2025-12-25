package mock_loader

import (
	"github.com/vkngwrapper/core/v3/common"
	gomock "go.uber.org/mock/gomock"
)

func LoaderForVersion(ctrl *gomock.Controller, version common.APIVersion) *MockLoader {
	mockDriver := NewMockLoader(ctrl)
	mockDriver.EXPECT().Version().Return(version).AnyTimes()

	return mockDriver
}
