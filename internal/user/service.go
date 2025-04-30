package user

import (
	"fmt"

	"github.com/ericsanto/api_authentication/internal/customerrors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *UserRepository
}

func NewUserService(userRepository *UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (us *UserService) CreateUserService(userRequest *UserRequest) (*UserResponse, error) {

	hashString := userRequest.EncryptPassword()

	UserModel := UserModel{
		Name:     userRequest.Name,
		Password: hashString,
		Email:    userRequest.Email,
	}

	userResponse, err := us.userRepository.CreateUser(UserModel)
	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	return userResponse, nil

}

func (us *UserService) Login(userRequest UserRequestLogin) (bool, error) {

	userModel, err := us.userRepository.FindUserByEmail(userRequest.Email)
	if err != nil {
		return false, fmt.Errorf("erro: %w", err)
	}

	fmt.Println(userModel.Password)

	if err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(userRequest.Password)); err != nil {
		return false, fmt.Errorf("%w", customerrors.ErrPasswordIncorrect)
	}

	return true, nil

}
