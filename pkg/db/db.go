package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"purple/configs"
)

type Db struct {
	*gorm.DB
}

func NewDb(config *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(config.DB.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &Db{db}
}
