package service

import (
	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/digicon-trap1-2023/backend/interfaces/repository"
)

type RequestService struct {
	requestRepository *repository.RequestRepository
}

func NewRequestService(requestRepository *repository.RequestRepository) *RequestService {
	return &RequestService{
		requestRepository,
	}
}

func (s *RequestService) GetRequests() ([]*domain.Request, error) {
	requests, err := s.requestRepository.GetRequests()
	if err != nil {
		return nil, err
	}

	return requests, nil
}
