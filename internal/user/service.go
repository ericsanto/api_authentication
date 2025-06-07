package user

import (
	"fmt"
	"os"
	"time"

	"github.com/ericsanto/api_authentication/internal/customerrors"
	"github.com/golang-jwt/jwt"
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

func (us *UserService) Login(userRequest UserRequestLogin) (*UserResponse, error) {

	userModel, err := us.userRepository.FindUserByEmail(userRequest.Email)
	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	userResponse := UserResponse{
		ID:    userModel.ID,
		Email: userModel.Email,
		Name:  userModel.Name,
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(userRequest.Password)); err != nil {
		return nil, fmt.Errorf("%w", customerrors.ErrPasswordIncorrect)
	}

	return &userResponse, nil

}

func (us *UserService) GenerateToken(id uint, email string) (string, error) {

	secret := os.Getenv("JWT_SECRET_KEY")

	encryptEmail, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("erro ao encriptar email")
	}

	encryptEmailString := string(encryptEmail)

	claims := jwt.MapClaims{
		"id":    id,
		"email": encryptEmailString,
		"exp":   time.Now().Add(time.Hour * 3).Unix(),
	}

	tokenGenerate := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := tokenGenerate.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("erro ao gerar token")
	}

	return tokenString, nil

}
