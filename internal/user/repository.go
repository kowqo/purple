package user

import (
	"purple/pkg/db"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	res := repo.Database.DB.Create(user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (repo *UserRepository) FindById(email string) (*User, error) {
	var user User
	res := repo.Database.DB.First(&user, email)

	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
