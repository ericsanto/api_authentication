package user

import (
	"fmt"
	"strings"

	"github.com/ericsanto/api_authentication/internal/customerrors"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(userModel UserModel) (*UserResponse, error) {

	if err := ur.db.Create(&userModel).Error; err != nil {
		messageDuplicateKeyUniqueContraint := customerrors.DupliateKeyUniqueConstraint()
		if strings.Contains(err.Error(), messageDuplicateKeyUniqueContraint) {
			return nil, fmt.Errorf("usuário com email %s %w", userModel.Email, customerrors.ErrDupliateKey)
		}
		return nil, fmt.Errorf("erro ao cadastrar usuário")
	}

	userResponse := UserResponse{
		ID:    userModel.ID,
		Name:  userModel.Name,
		Email: userModel.Email,
	}

	return &userResponse, nil
}

func (ur *UserRepository) FindUserByEmail(email string) (*UserModel, error) {

	var userResponse UserModel

	query := `SELECT * FROM user_models WHERE email=?`

	if err := ur.db.Raw(query, email).Scan(&userResponse).Error; err != nil {

		return nil, fmt.Errorf("erro ao buscar usuario no banco de dados %w", err)
	}

	if userResponse.Email == "" {
		return nil, fmt.Errorf("não existe usuário cadastrado com esse email %w", customerrors.ErrNotFound)
	}

	return &userResponse, nil
}
