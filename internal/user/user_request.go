package user

import (
	"golang.org/x/crypto/bcrypt"
)

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRequestLogin struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (ur *UserRequest) EncryptPassword() string {

	hash, err := bcrypt.GenerateFromPassword([]byte(ur.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	hashString := string(hash)

	return hashString

}
