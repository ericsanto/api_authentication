package user

type UserModel struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"size:100"`
	Email    string `gorm:"size:255;unique"`
	Password string `gorm:"size:255"`
}
