package model

import (
	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/google/uuid"
)

type User struct {
	Id   string `gorm:"type:char(36);not null;primaryKey"`
	Name string `gorm:"type:varchar(32)"`
	Role string `gorm:"type:varchar(32)"`
}

func (user *User) ToDomain() (*domain.User, error) {
	id, err := uuid.Parse(user.Id)
	if err != nil {
		return nil, err
	}
	return &domain.User {
		Id: id,
		Name: user.Name,
		Role: user.Role,
	}, nil
}