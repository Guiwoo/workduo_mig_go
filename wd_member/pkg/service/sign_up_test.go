package service

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"net/http"
	"net/http/httptest"
	"testing"
	"wd_user/pkg/model"
	"wd_user/pkg/service/mock"
)

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
	createReq := func() *http.Request {
		return httptest.NewRequest(http.MethodPost, "/user/api/v1.0", nil)
	}
	createResp := func() *httptest.ResponseRecorder {
		return httptest.NewRecorder()
	}
	createContext := func(req *http.Request, res *httptest.ResponseRecorder) echo.Context {
		c := echo.New()
		return c.NewContext(req, res)
	}

	log := zerolog.Logger{}

	type fields struct {
		repo model.MemberRepository
		log  zerolog.Logger
		req  signUpRequest
	}

	type args struct {
		req *http.Request
		res *httptest.ResponseRecorder
	}

	type response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   response
	}{
		{
			name:   "회원가입 실패 - 요청파람 데이터 검증 실패",
			fields: fields{repo: mock.NewMockDB(false), log: log, req: signUpRequest{}},
			args:   args{createReq(), createResp()},
			want:   response{Code: http.StatusBadRequest, Msg: FailSignUp},
		},
		{
			name: "회원가입 실패 - 패스워드 해시 실패",
			fields: fields{repo: mock.NewMockDB(false), log: log, req: signUpRequest{
				Name:      "test",
				Email:     "abc@abc.com",
				Password:  "", // 해시 실패 를 위해서 ""을 넘겨야 하지만 Parameter 검증에서 걸려 fail
				Areas:     make([]string, 3),
				Exercises: make([]string, 2),
			}},
			args: args{createReq(), createResp()},
			want: response{Code: http.StatusBadRequest, Msg: FailSignUp},
		},
		{
			name: "회원가입 실패 - DB 오류",
			fields: fields{repo: mock.NewMockDB(false), log: log, req: signUpRequest{
				Name:      "test",
				Email:     "abc@abc.com",
				Password:  "12312312312@@#@",
				Areas:     make([]string, 3),
				Exercises: make([]string, 2),
			}},
			args: args{createReq(), createResp()},
			want: response{Code: http.StatusInternalServerError, Msg: FailSignUp},
		},
		{
			name: "회원가입 성공",
			fields: fields{repo: mock.NewMockDB(true), log: log, req: signUpRequest{
				Name:      "test",
				Email:     "abc@abc.com",
				Password:  "123123@abc.com",
				Areas:     make([]string, 3),
				Exercises: make([]string, 2),
			}},
			args: args{createReq(), createResp()},
			want: response{Code: http.StatusOK, Msg: "success"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &SignUp{
				repo: tt.fields.repo,
				log:  tt.fields.log,
				req:  tt.fields.req,
			}

			if err := service.Handle(createContext(tt.args.req, tt.args.res)); err != nil {
				t.Error(err)
			}

			var data response

			if err := json.Unmarshal(tt.args.res.Body.Bytes(), &data); err != nil {
				t.Error(err)
			}

			if data.Code != tt.want.Code && data.Msg != tt.want.Msg {
				t.Errorf("fail test service.Handle() => got %+v, want %+v", data, tt.want)
			}
		})
	}
}
