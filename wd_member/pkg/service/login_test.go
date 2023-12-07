package service

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"net/http"
	"net/http/httptest"
	"testing"
	"wd_user/pkg/model"
	"wd_user/pkg/service/mock"
)

func TestLogin_validateEmail(t *testing.T) {
	type fields struct {
		repo model.MemberRepository
		log  zerolog.Logger
		req  loginRequest
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "로그인 실패 - 이메일 검증 실패 @ 없는경우",
			fields: fields{
				repo: nil,
				log:  zerolog.Logger{},
				req: loginRequest{
					Email: "abc#abc.com",
				},
			},
			want: false,
		},
		{
			name: "로그인 실패 - 이메일 검증 실패 특수문자 없는경우",
			fields: fields{
				repo: nil,
				log:  zerolog.Logger{},
				req: loginRequest{
					Email: "abcabc2abc.com",
				},
			},
			want: false,
		},
		{
			name: "로그인 실패 - 이메일 검증 성공",
			fields: fields{
				repo: nil,
				log:  zerolog.Logger{},
				req: loginRequest{
					Email: "abcabc@abc.com",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Login{
				repo: tt.fields.repo,
				log:  tt.fields.log,
				req:  tt.fields.req,
			}
			if got := service.validateEmail(); got != tt.want {
				t.Errorf("validateEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogin_getHashedPassword(t *testing.T) {
	type fields struct {
		repo model.MemberRepository
		log  zerolog.Logger
		req  loginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "해쉬 패스워드 획득 - 성공",
			fields: fields{
				repo: nil,
				log:  zerolog.Logger{},
				req: loginRequest{
					Password: "123",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Login{
				repo: tt.fields.repo,
				log:  tt.fields.log,
				req:  tt.fields.req,
			}
			_, err := service.getHashedPassword()
			if (err != nil) != tt.wantErr {
				t.Errorf("getHashedPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestLogin_Handle(t *testing.T) {
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
		req  loginRequest
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
			name:   "로그인 실패 - 요청파람 데이터 검증 실패",
			fields: fields{repo: mock.NewMockDB(false), log: log, req: loginRequest{}},
			args:   args{createReq(), createResp()},
			want:   response{Code: http.StatusBadRequest, Msg: FailLogin},
		},
		{
			name:   "로그인 실패 - 멤버 조회 실패",
			fields: fields{repo: mock.NewMockDB(false), log: log, req: loginRequest{}},
			args:   args{createReq(), createResp()},
			want:   response{Code: http.StatusInternalServerError, Msg: FailLogin},
		},
		{
			name: "로그인 성공",
			fields: fields{repo: mock.NewMockDB(true), log: log, req: loginRequest{
				Email:    "abc@abc.com",
				Password: "test12344@", // 해시 실패 를 위해서 ""을 넘겨야 하지만 Parameter 검증에서 걸려 fail 해쉬 테스트로 대체
			}},
			args: args{createReq(), createResp()},
			want: response{Code: http.StatusBadRequest, Msg: FailLogin},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Login{
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
