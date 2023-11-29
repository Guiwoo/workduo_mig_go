package mock

import (
	"errors"
	"github.com/labstack/echo/v4"
	"wd_user/pkg/model"
)

var (
	_ model.MemberRepository = (*SuccessMock)(nil)
	_ model.MemberRepository = (*FailMock)(nil)
)

type SuccessMock struct {
}

func (s *SuccessMock) Find(ctx echo.Context, email string) (*model.Member, error) {
	panic("implement me")
}

func (s *SuccessMock) Create(ctx echo.Context, member *model.Member) error {
	return nil
}

type FailMock struct {
}

func (f *FailMock) Find(ctx echo.Context, email string) (*model.Member, error) {
	panic("implement me")
}

func (f *FailMock) Create(ctx echo.Context, member *model.Member) error {
	return errors.New("fail mock create")
}

func NewMockDB(caseType bool) model.MemberRepository {
	if caseType {
		return &SuccessMock{}
	} else {
		return &FailMock{}
	}
}
