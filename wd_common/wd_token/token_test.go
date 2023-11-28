package wd_token

import (
	"reflect"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.args.memberID, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Generate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		tokenStr string
	}
	tests := []struct {
		name    string
		args    args
		want    *Token
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.tokenStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToken_Valid(t1 *testing.T) {
	type fields struct {
		MemberID    string
		Email       string
		ExpiredDate time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Token{
				MemberID:    tt.fields.MemberID,
				Email:       tt.fields.Email,
				ExpiredDate: tt.fields.ExpiredDate,
			}
			if err := t.Valid(); (err != nil) != tt.wantErr {
				t1.Errorf("Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
