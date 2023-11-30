package wd_token

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"testing"
	"time"
)

func parseSymmetricKWay(tokenStr string) (*Token, error) {
	t := &Token{}
	token, err := jwt.ParseWithClaims(tokenStr, t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Header["alg"]; !ok {
			return nil, fmt.Errorf("missing algorithm in token header")
		}
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}
		return "This is Secret Key", nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Token); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("fail to type cast token")
	}
}

func generateSymmetricWay(memberID, email string) (string, error) {
	t := Token{
		MemberID:  memberID,
		Email:     email,
		ExpiresAt: time.Now().Add(ExpiredTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &t)
	return token.SignedString([]byte("This is Secret Key"))
}

func BenchmarkGenerateSymmetricWay(b *testing.B) {
	randomStr := func() string {
		uuid, _ := uuid.NewUUID()
		return uuid.String()
	}

	b.Run("Rsa 방법", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Generate(randomStr(), randomStr())
		}
	})

	b.Run("대칭키 방법", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generateSymmetricWay(randomStr(), randomStr())
		}
	})
}

func BenchmarkParsingToken(b *testing.B) {
	randomStr := func() string {
		uuid, _ := uuid.NewUUID()
		return uuid.String()
	}
	rsa, _ := Generate(randomStr(), randomStr())
	symmetric, _ := generateSymmetricWay(randomStr(), randomStr())

	b.Run("Rsa 방법", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Parse(rsa)
		}
	})

	b.Run("대칭키 방법", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			parseSymmetricKWay(symmetric)
		}
	})
}
