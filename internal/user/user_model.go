package user

import "github.com/ericsanto/api_authentication/internal/farm"

type UserModel struct {
	ID       uint        `gorm:"primaryKey;autoIncrement"`
	Name     string      `gorm:"size:100"`
	Email    string      `gorm:"size:255;unique"`
	Password string      `gorm:"size:255"`
	Farm     []farm.Farm `gorm:"foreignKey:UserID"`
}
