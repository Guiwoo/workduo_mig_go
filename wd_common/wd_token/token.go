package wd_token

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const ExpiredTime = time.Minute * 10

var PrivateKey *ecdsa.PrivateKey

type Token struct {
	MemberID  string
	Email     string
	ExpiresAt time.Time
}

func (t *Token) Valid() error {
	if t.ExpiresAt.Before(time.Now()) {
		return errors.New("token is expired")
	}
	return nil
}

func Generate(memberID, email string) (string, error) {
	t := Token{
		MemberID:  memberID,
		Email:     email,
		ExpiresAt: time.Now().Add(ExpiredTime),
	}

	PrivateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, &t)
	return token.SignedString(PrivateKey)
}

func Parse(tokenStr string) (*Token, error) {
	t := &Token{}
	token, err := jwt.ParseWithClaims(tokenStr, t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Header["alg"]; !ok {
			return nil, fmt.Errorf("missing algorithm in token header")
		}
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}
		return PrivateKey.Public(), nil
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
