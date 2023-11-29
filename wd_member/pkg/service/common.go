package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

const (
	WrongBodyForm = "올바른 바디의 데이터 타입이 아닙니다."
)

const (
	ErrCodeEmail       = 1101
	ErrCodePassword    = 1102
	ErrCodePauseMember = 1103
	ErrCodeQuitMember  = 1104
)

var (
	ErrRequiredName        = errors.New("name is required")
	ErrWrongFormatEmail    = errors.New("email format is wrong")
	ErrEmptyArea           = errors.New("area needs at least once")
	ErrEmptyExercise       = errors.New("exercise needs at least once")
	ErrWrongFormatPassword = errors.New("password should have special character and longer than 8 characters")
)

func IDGenerator(target string) string {
	uuid, _ := uuid.NewUUID()
	return fmt.Sprintf("%s_%s", target, uuid.String())
}
