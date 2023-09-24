package repository

import (
	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/digicon-trap1-2023/backend/interfaces/repository/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DocumentRepository struct {
	conn *gorm.DB
}

func NewDocumentRepository(conn *gorm.DB) *DocumentRepository {
	return &DocumentRepository{conn}
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
		res, err := document.ToDomain(bookmarks, references)
		if err != nil {
			return nil, err
		}

		result = append(result, res)
	}

	return result, nil
}
