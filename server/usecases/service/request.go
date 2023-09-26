package service

import (
	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/digicon-trap1-2023/backend/interfaces/repository"
	"github.com/google/uuid"
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
	return s.requestRepository.GetRequests()
}

func (s *RequestService) CreateRequest(userId uuid.UUID, Tags []uuid.UUID, Title string, Description string) (*domain.Request, error) {
	request := domain.NewRequest(Title, Description, Tags, userId)
	return s.requestRepository.CreateRequest(request)
}
