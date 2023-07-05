package mocktest_test

import (
	"github.com/golang/mock/gomock"
	"go-examples/pkg/mocks"
	"testing"
)

func TestMockDoer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)

	mockDoer.EXPECT().DoSomething(1, "abc").Return(nil).Times(1)

	mockDoer.DoSomething(1, "abc")
}
