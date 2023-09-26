package repository

import (
	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/digicon-trap1-2023/backend/interfaces/repository/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RequestRepository struct {
	conn *gorm.DB
}

func NewRequestRepository(conn *gorm.DB) *RequestRepository {
	return &RequestRepository{conn}
}

func (r *RequestRepository) GetRequests() ([]*domain.Request, error) {
	var requests []*model.Request
	var tags map[string][]uuid.UUID
	var tagRequests []*model.TagRequest
	if err := r.conn.Find(&requests).Error; err != nil {
		return nil, err
	}

	if err := r.conn.Find(&tagRequests).Error; err != nil {
		return nil, err
	}

	tags = make(map[string][]uuid.UUID)
	for _, tagRequest := range tagRequests {
		tagId, err := uuid.Parse(tagRequest.TagId)
		if err != nil {
			return nil, err
		}
		tags[tagRequest.RequestId] = append(tags[tagRequest.RequestId], tagId)
	}

	result := make([]*domain.Request, len(requests))
	for _, request := range requests {
		req, err := request.ToDomain(tags[request.Id])
		if err != nil {
			return nil, err
		}
		result = append(result, req)
	}

	return result, nil
}
