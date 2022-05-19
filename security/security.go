package security

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePwd(plainPwd string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(plainPwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(password), nil
}

func ComparePwd(hashedPwd string, plainPwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		return err
	}
	return nil
}
