package service

import (
	"github.com/rs/zerolog"
	"testing"
	"wd_user/pkg/model"
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
