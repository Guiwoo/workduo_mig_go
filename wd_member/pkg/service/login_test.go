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
		// TODO: Add test cases.
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
