package wd_token

import (
	"crypto/aes"
	"encoding/hex"
	"encoding/json"
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

func generateAESWay(memberID, email string) (string, error) {
	t := Token{
		MemberID:  memberID,
		Email:     email,
		ExpiresAt: time.Now().Add(ExpiredTime),
	}

	token, err := json.Marshal(&t)
	if err != nil {
		return "", err
	}

	cip, err := aes.NewCipher([]byte("This is must!@#$"))
	if err != nil {
		return "", err
	}
	length := (len(token) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, token)
	pad := byte(len(plain) - len(token))
	for i := len(token); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	for bs, be := 0, cip.BlockSize(); bs <= len(token); bs, be = bs+cip.BlockSize(), be+cip.BlockSize() {
		cip.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return hex.EncodeToString(encrypted), nil
}

func parseAESWay(token string) (*Token, error) {
	encrypted, err := hex.DecodeString("This is must!@#$")
	if err != nil {
		return nil, err
	}
	cip, err := aes.NewCipher([]byte("This is must!@#$"))
	if err != nil {
		return nil, err
	}

	decrypted := make([]byte, len(encrypted))
	for bs, be := 0, cip.BlockSize(); bs < len(encrypted); bs, be = bs+cip.BlockSize(), be+cip.BlockSize() {
		cip.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	data := string(decrypted[:trim])
	var parsed Token
	if err := json.Unmarshal([]byte(data), &token); err != nil {
		return nil, err
	}
	return &parsed, nil
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

	b.Run("해쉬 방법", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generateSymmetricWay(randomStr(), randomStr())
		}
	})

	b.Run("대칭키 방법", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generateAESWay(randomStr(), randomStr())
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
	aes, _ := generateAESWay(randomStr(), randomStr())

	b.Run("Rsa 방법", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Parse(rsa)
		}
	})

	b.Run("해쉬 방법", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			parseSymmetricKWay(symmetric)
		}
	})

	b.Run("대칭키 방법", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			parseAESWay(aes)
		}
	})
}
