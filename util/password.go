package util

import (
	"golang.org/x/crypto/scrypt"
)

func GetPassword(psw string, salt []byte) string {
	dk, _ := scrypt.Key([]byte(psw), salt, 16384, 8, 1, 32)
	return string(dk)
}
