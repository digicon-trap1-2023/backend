package repository

import (
	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/google/uuid"
	"github.com/digicon-trap1-2023/backend/interfaces/repository/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{
		conn,
	}
}

func (r *UserRepository) GetUser(userId uuid.UUID) (*domain.User, error) {
	var user *model.User
	if err := r.conn.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return user.ToDomain()
}
