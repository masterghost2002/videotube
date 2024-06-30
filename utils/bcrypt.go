package utils

import (
	config "github.com/masterghost2002/videotube/configs"
	"golang.org/x/crypto/bcrypt"
)

var cryptSalt = config.ENVS.CryptSalt

func HashString(str string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(str+cryptSalt), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
func ChechString(hash, str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str+cryptSalt))
	return err == nil
}
