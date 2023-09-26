package repository

import (
	"gorm.io/gorm"
)

type RequestRepository struct {
	conn *gorm.DB
}

func NewRequestRepository(conn *gorm.DB) *RequestRepository {
	return &RequestRepository{conn}
}
