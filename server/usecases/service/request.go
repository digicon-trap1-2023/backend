package service

import "github.com/digicon-trap1-2023/backend/interfaces/repository"

type RequestService struct {
	requestRepository *repository.RequestRepository
}

func NewRequestService(requestRepository *repository.RequestRepository) *RequestService {
	return &RequestService{
		requestRepository,
	}
}
