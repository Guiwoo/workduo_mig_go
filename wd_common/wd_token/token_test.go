package wd_token

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	type args struct {
		memberID string
		email    string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "토큰생성 성공",
			args:    args{"MB_abc", "email@email.com"},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Generate(tt.args.memberID, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestParse(t *testing.T) {
	current := time.Now()
	tokenStr, err := Generate("MB_abc", "email@email.com")
	fmt.Println(tokenStr, err)
	type args struct {
		tokenStr string
	}
	tests := []struct {
		name    string
		args    args
		want    *Token
		wantErr bool
	}{
		{
			name: "토큰파싱 성공",
			args: args{
				tokenStr: tokenStr,
			},
			want: &Token{
				MemberID: "MB_abc",
				Email:    "email@email.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.tokenStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.MemberID != "MB_abc" && got.Email != "email@email.com" {
				t.Errorf("got %+v", got)
				return
			}
			if current.Add(ExpiredTime - 1).After(got.ExpiresAt) {
				t.Errorf("fail time should be +10mins %+v", got.ExpiresAt)
				return
			}
		})
	}
}

func TestToken_Valid(t1 *testing.T) {
	type fields struct {
		ExpiredDate time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "토큰검증 실패 - 만료기한 지난 토큰",
			fields: fields{
				ExpiredDate: time.Now().Add(-(10 * time.Minute)),
			},
			wantErr: true,
		},
		{
			name: "토큰검증 성공",
			fields: fields{
				ExpiredDate: time.Now().Add(-(time.Second * 599)),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Token{}
			if err := t.Valid(); (err != nil) != tt.wantErr {
				t1.Errorf("Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
