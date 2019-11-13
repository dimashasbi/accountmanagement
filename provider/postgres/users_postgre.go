package postgres

import (
	"AccountManagement/engine"
	"AccountManagement/model"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type (
	usersRepository struct {
		session *gorm.DB
	}
)

func newUsersRepository(db *gorm.DB) engine.UsersRepository {
	return &usersRepository{db}
}

func (s usersRepository) Insert(k *model.Users) error {
	result := s.session.Create(&k)
	if result.Error != nil {
		return errors.Errorf("Error Insert Users : %v", result.Error)
	}
	return nil
}

func (s usersRepository) Select(k *model.Users) (*model.Users, error) {
	result := s.session.Where("user_name = ?", k.UserName).First(&k)
	if result.Error != nil {
		return k, errors.Errorf("Error Select Users : %v", result.Error)
	}
	return k, nil
}
