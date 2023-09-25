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

func (s *DocumentService) GetBookmarkedDocuments(userId uuid.UUID, tags []string) ([]*domain.Document, error) {
	return s.documentRepository.GetBookmarkedDocuments(userId, tags)
}

func (s *DocumentService) GetDocument(userId uuid.UUID, documentId uuid.UUID) (*domain.Document, error) {
	document, err := s.documentRepository.GetDocument(userId, documentId)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (s *DocumentService) CreateDocument(userId uuid.UUID, title string, description string, tagIds []uuid.UUID, file *multipart.FileHeader) (*domain.Document, error) {
	return s.documentRepository.CreateDocument(userId, title, description, tagIds, file)
}

func (s *DocumentService) UpdateDocument(userId uuid.UUID, documentId uuid.UUID, title string, description string, tagIds []uuid.UUID, file *multipart.FileHeader) (*domain.Document, error) {
	return s.documentRepository.UpdateDocument(userId, documentId, title, description, tagIds, file)
}

func (s *DocumentService) DeleteDocument(userId uuid.UUID, documentId uuid.UUID) error {
	return s.documentRepository.DeleteDocument(userId, documentId)
}

func (s *DocumentService) BookMark(userId uuid.UUID, documentId uuid.UUID) error {
	return s.documentRepository.BookMark(userId, documentId)
}

func (s *DocumentService) UnBookMark(userId uuid.UUID, documentId uuid.UUID) error {
	return s.documentRepository.UnBookMark(userId, documentId)
}

func (s *DocumentService) Reference(userId uuid.UUID, documentId uuid.UUID) error {
	return s.documentRepository.Reference(userId, documentId)
}

func (s *DocumentService) UnReference(userId uuid.UUID, documentId uuid.UUID) error {
	return s.documentRepository.UnReference(userId, documentId)
}