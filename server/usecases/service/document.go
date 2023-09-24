package service

import (
	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/digicon-trap1-2023/backend/interfaces/repository"
	"github.com/google/uuid"
)

type DocumentService struct {
	documentRepository *repository.DocumentRepository
}

func NewDocumentService(documentRepository *repository.DocumentRepository) *DocumentService {
	return &DocumentService{
		documentRepository: documentRepository,
	}
}

func (s *DocumentService) GetDocuments(userId uuid.UUID, tags []string) ([]*domain.Document, error) {
	return s.documentRepository.GetDocuments(userId, tags)
}
