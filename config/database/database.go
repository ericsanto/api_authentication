package database

import (
	"fmt"

	"github.com/ericsanto/api_authentication/config"
	"github.com/ericsanto/api_authentication/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() error {

	dsn := "host=db user=go password=go dbname=go port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&user.UserModel{}); err != nil {
		return fmt.Errorf("erro ao migrar tabelas para o banco de dados %s", err.Error())
	}

	config.DB = db

	return nil

}
