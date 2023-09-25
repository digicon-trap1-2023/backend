package repository

import (
	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/digicon-trap1-2023/backend/infrastructure"
	"github.com/digicon-trap1-2023/backend/interfaces/repository/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DocumentRepository struct {
	conn *gorm.DB
	s3   *infrastructure.S3Client
}

func NewDocumentRepository(conn *gorm.DB, client *infrastructure.S3Client) *DocumentRepository {
	return &DocumentRepository{
		conn: conn,
		s3:   client,
	}
}

func (r *DocumentRepository) GetDocuments(userId uuid.UUID, tags []string) ([]*domain.Document, error) {
	var tagDocuments []model.TagDocument
	var docIds []string
	var documents []*model.Document
	var bookmarks []*model.BookMark
	var references []*model.Reference

	if err := r.conn.Where("tag_id IN ?", tags).Find(&tagDocuments).Error; err != nil {
		return nil, err
	}

	for _, tagDocument := range tagDocuments {
		docIds = append(docIds, tagDocument.DocumentId)
	}

	if err := r.conn.Where("id IN ?", docIds).Find(&documents).Error; err != nil {
		return nil, err
	}

	if err := r.conn.Where("document_id IN ?", docIds).Where("user_id = ?", userId).Find(&bookmarks).Error; err != nil {
		return nil, err
	}

	if err := r.conn.Where("document_id IN ?", docIds).Find(&references).Error; err != nil {
		return nil, err
	}

	var result []*domain.Document
	for _, document := range documents {
		res, err := document.ToDomain(bookmarks, references, nil)
		if err != nil {
			return nil, err
		}

		result = append(result, res)
	}

	return result, nil
}

func (r *DocumentRepository) GetDocument(userId uuid.UUID, documentId uuid.UUID) (*domain.Document, error) {
	var document model.Document
	var bookmarks []*model.BookMark
	var references []*model.Reference
	var tagDocuments []*model.TagDocument
	var tags []*model.Tag

	if err := r.conn.Where("id = ?", documentId).First(&document).Error; err != nil {
		return nil, err
	}

	if err := r.conn.Where("document_id = ?", documentId).Where("user_id = ?", userId).Find(&bookmarks).Error; err != nil {
		return nil, err
	}

	if err := r.conn.Where("document_id = ?", documentId).Find(&references).Error; err != nil {
		return nil, err
	}

	if err := r.conn.Where("document_id = ?", documentId).Find(&tagDocuments).Error; err != nil {
		return nil, err
	}

	tagIds := make([]string, len(tags))
	for i, tag := range tagDocuments {
		tagIds[i] = tag.TagId
	}

	if err := r.conn.Where("id IN ?", tagIds).Find(&tags).Error; err != nil {
		return nil, err
	}

	return document.ToDomain(bookmarks, references, tags)
}
