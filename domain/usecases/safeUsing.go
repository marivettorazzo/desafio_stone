package usecases

import (
	"crypto/sha256"
	"fmt"

	"exemple.com/DesafioStone/desafio_stone/domain"
)

func AsSha256(o interface{}) string { //hash generator function
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}
func PasswordGenerator(s string) string { //class method that changes password to hash
	var k domain.KeySecret
	k.Key = AsSha256(s)
	return k.Key
}
