package usecase

import (
	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"
)

type Encrypter_pbkdf2 struct{}

func (Encrypter_pbkdf2) EncryptString(stringToEncrypt, salt string) string {
	bytes := pbkdf2.Key([]byte(stringToEncrypt), []byte(salt), 4096, 32, sha256.New)
	return string(bytes)
}
