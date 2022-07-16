package utils

import "golang.org/x/crypto/bcrypt"

func Hash(pwd string) string {
	bytePass := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(hash)
}

func ComparePasswords(plainPwd string, hashedPwd string) bool {
	bytePass := []byte(plainPwd)
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePass)
	return err == nil
}
