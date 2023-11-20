package service

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"net/http"
	"net/http/httptest"
	"testing"
	"wd_user/pkg/model"
)

var (
	_ model.MemberRepository = (*SuccessMock)(nil)
	_ model.MemberRepository = (*FailMock)(nil)
)

type SuccessMock struct {
}

func (s *SuccessMock) Create(ctx echo.Context, member *model.Member) error {
	return nil
}

type FailMock struct {
}

func (f *FailMock) Create(ctx echo.Context, member *model.Member) error {
	return nil
}

func NewMockDB(caseType bool) model.MemberRepository {
	if caseType {
		return &SuccessMock{}
	} else {
		return &FailMock{}
	}
}

func TestSignUp_Handle_validateRequest(t *testing.T) {

	type fields struct {
		repo model.MemberRepository
		log  zerolog.Logger
		req  signUpRequest
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name:    "회원가입 검증 실패 - 이름없음 실패",
			fields:  fields{req: signUpRequest{}},
			wantErr: ErrRequiredName,
		},
		{
			name:    "회원가입 검증 실패 - 잘못된 이메일 형태",
			fields:  fields{req: signUpRequest{Name: "test", Email: "abc#.com"}},
			wantErr: ErrWrongFormatEmail,
		},
		{
			name:    "회원가입 검증 실패 - 잘못된 비밀번호 형태",
			fields:  fields{req: signUpRequest{Name: "test", Email: "abc@abc.com", Password: "ab123123"}},
			wantErr: ErrWrongFormatPassword,
		},
		{
			name:    "회원가입 검증 실패 - 지역 선택이 없는경우",
			fields:  fields{req: signUpRequest{Name: "test", Email: "abc@abc.com", Password: "ab123123123@"}},
			wantErr: ErrEmptyArea,
		},
		{
			name:    "회원가입 검증 실패 - 운동 선택이 없는경우",
			fields:  fields{req: signUpRequest{Name: "test", Email: "abc@abc.com", Password: "ab123123123@", Areas: make([]string, 3)}},
			wantErr: ErrEmptyExercise,
		},
		{
			name: "회원가입 성공",
			fields: fields{req: signUpRequest{
				Name:      "test",
				Email:     "abc@abc.com",
				Password:  "ab123123123@",
				Areas:     make([]string, 3),
				Exercises: make([]string, 2),
			}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &SignUp{
				repo: tt.fields.repo,
				log:  tt.fields.log,
				req:  tt.fields.req,
			}
			err := service.validateRequest()
			if !errors.Is(err, tt.wantErr) {
				t.Error(err)
			}
		})
	}
}

func TestSignUp_validatePassword(t *testing.T) {
	type fields struct {
		req signUpRequest
	}
	tests := []struct {
		name     string
		fields   fields
		wantBool bool
	}{
		{
			name:     "비밀번호 검증 실패 - 비어있는 경우",
			fields:   fields{req: signUpRequest{Password: ""}},
			wantBool: false,
		},
		{
			name:     "비밀번호 검증 실패 - 특수문자가 없는경우",
			fields:   fields{req: signUpRequest{Password: "abcd1234"}},
			wantBool: false,
		},
		{
			name:     "비밀번호 검증 성공",
			fields:   fields{req: signUpRequest{Password: "abcd123$"}},
			wantBool: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &SignUp{
				req: tt.fields.req,
			}
			if got := service.validatePassword(); got != tt.wantBool {
				t.Errorf("validatePassword() = %v, want %v", got, tt.wantBool)
			}
		})
	}
}

func TestSignUp_Handle(t *testing.T) {

	c := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user/api/v1.0", nil)
	res := httptest.NewRecorder()
	ctx := c.NewContext(req, res)

	log := zerolog.Logger{}

	type fields struct {
		repo model.MemberRepository
		log  zerolog.Logger
		req  signUpRequest
	}

	type args struct {
		ctx echo.Context
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr string
	}{
		{
			name:    "회원가입 실패 - 요청파람 데이터 검증 실패",
			fields:  fields{repo: NewMockDB(false), log: log, req: signUpRequest{}},
			args:    args{ctx},
			wantErr: FailSignUp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &SignUp{
				repo: tt.fields.repo,
				log:  tt.fields.log,
				req:  tt.fields.req,
			}
			service.Handle(tt.args.ctx)
		})
	}
}
