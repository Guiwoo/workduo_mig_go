package wd_token

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

const ExpiredTime = time.Minute * 10

var privateKey *ecdsa.PrivateKey

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

func getPrivateKey() *ecdsa.PrivateKey {
	if privateKey == nil {
		pk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			log.Fatalf("fail to generate private, public key pair %+v", err)
		}
		privateKey = pk
	}
	return privateKey
}

func Generate(memberID, email string) (string, error) {
	t := Token{
		MemberID:  memberID,
		Email:     email,
		ExpiresAt: time.Now().Add(ExpiredTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, &t)
	return token.SignedString(getPrivateKey())
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
		return privateKey.Public(), nil
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
