package service

import (
	"mime/multipart"

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

func (s *DocumentService) GetDocument(userId uuid.UUID, documentId uuid.UUID) (*domain.Document, error) {
	document, err := s.documentRepository.GetDocument(userId, documentId)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (s *DocumentService) CreateDocument(userId uuid.UUID, title string, description string, tags []string, file *multipart.FileHeader) (*domain.Document, error) {
	return s.documentRepository.CreateDocument(userId, title, description, tags, file)
}

func (s *DocumentService) UpdateDocument(userId uuid.UUID, documentId uuid.UUID, title string, description string, tags []string, file *multipart.FileHeader) (*domain.Document, error) {
	return s.documentRepository.UpdateDocument(userId, documentId, title, description, tags, file)
}
