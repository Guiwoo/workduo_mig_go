package wd_rsa

import (
	"fmt"
	"testing"
)

func TestGetPrivateKey(t *testing.T) {
	pk := GetPrivateKey()
	fmt.Println(pk.X)
	fmt.Println(pk.Y)
	fmt.Println(pk.D)
}
