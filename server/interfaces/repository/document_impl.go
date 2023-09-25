package repository

import (
	"fmt"
	"mime/multipart"

	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/digicon-trap1-2023/backend/infrastructure"
	"github.com/digicon-trap1-2023/backend/interfaces/repository/model"
	"github.com/digicon-trap1-2023/backend/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	if len(tags) == 0 {
		if err := r.conn.Find(&documents).Error; err != nil {
			return nil, err
		}
		for _, document := range documents {
			docIds = append(docIds, document.Id)
		}
	} else {
		if err := r.conn.Where("tag_id IN ?", tags).Find(&tagDocuments).Error; err != nil {
			return nil, err
		}
		for _, tagDocument := range tagDocuments {
			docIds = append(docIds, tagDocument.DocumentId)
		}
		if err := r.conn.Where("id IN ?", docIds).Find(&documents).Error; err != nil {
			return nil, err
		}
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

	fmt.Println("result", result)
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
	for _, tag := range tagDocuments {
		tagIds = append(tagIds, tag.TagId)
	}

	if err := r.conn.Where("id IN ?", tagIds).Find(&tags).Error; err != nil {
		return nil, err
	}

	return document.ToDomain(bookmarks, references, tags)
}

func (r *DocumentRepository) CreateDocument(userId uuid.UUID, title string, description string, tagIds []uuid.UUID, file *multipart.FileHeader) (*domain.Document, error) {
	var tagModels []*model.Tag

	fileId := util.NewID().String()

	fileUrl, err := r.s3.PutObject(fileId, file); 
	if err != nil {
		return nil, err
	}

	document := &model.Document{
		Id:          util.NewID().String(),
		UserId:      userId.String(),
		Title:       title,
		File:        fileUrl,
		Description: description,
	}
	var tagDocuments []*model.TagDocument
	for _, tagID := range tagIds {
		tagDocument := &model.TagDocument{
			TagId:      tagID.String(),
			DocumentId: document.Id,
		}
		tagDocuments = append(tagDocuments, tagDocument)
	}

	if err := r.conn.Create(&tagDocuments).Error; err != nil {
		return nil, err
	}

	if err := r.conn.Create(document).Error; err != nil {
		return nil, err
	}

	if err := r.conn.Where("id IN ?", tagIds).Find(&tagModels).Error; err != nil {
		return nil, err
	}

	return document.ToDomain(nil, nil, tagModels)
}

func (r *DocumentRepository) UpdateDocument(userId uuid.UUID, documentId uuid.UUID, title string, description string, tagIds []uuid.UUID, file *multipart.FileHeader) (*domain.Document, error) {
	var tagModels []*model.Tag
	var document model.Document
	if err := r.conn.Where("id = ?", documentId).First(&document).Error; err != nil {
		return nil, err
	}

	if document.UserId != userId.String() {
		return nil, fmt.Errorf("this document is not yours")
	}

	if file != nil {
		fileUrl, err := r.s3.PutObject(document.File, file)
		if err != nil {
			return nil, err
		}

		document.File = fileUrl
	}

	if title != "" {
		document.Title = title
	}

	if description != "" {
		document.Description = description
	}

	var tagDocuments []*model.TagDocument
	for _, tagID := range tagIds {
		tagDocument := &model.TagDocument{
			TagId:      tagID.String(),
			DocumentId: document.Id,
		}
		tagDocuments = append(tagDocuments, tagDocument)
	}

	if err := r.conn.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&tagDocuments).Error; err != nil {
		return nil, err
	}

	if err := r.conn.Model(&document).Updates(model.Document{
		Title:       document.Title,
		Description: document.Description,
		File:        document.File,
	}).Error; err != nil {
		return nil, err
	}

	if err := r.conn.Where("id IN ?", tagIds).Find(&tagModels).Error; err != nil {
		return nil, err
	}

	return document.ToDomain(nil, nil, tagModels)
}

func (r *DocumentRepository) DeleteDocument(userId uuid.UUID, documentId uuid.UUID) error {
	var document model.Document
	if err := r.conn.Where("id = ?", documentId).First(&document).Error; err != nil {
		return err
	}

	if document.UserId != userId.String() {
		return fmt.Errorf("this document is not yours")
	}

	if err := r.conn.Where("id = ?", documentId).Delete(&model.Document{}).Error; err != nil {
		return err
	}

	if err := r.conn.Where("document_id = ?", documentId).Delete(&model.TagDocument{}).Error; err != nil {
		return err
	}

	if err := r.conn.Where("document_id = ?", documentId).Delete(&model.BookMark{}).Error; err != nil {
		return err
	}

	if err := r.conn.Where("document_id = ?", documentId).Delete(&model.Reference{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *DocumentRepository) BookMark(userId uuid.UUID, documentId uuid.UUID) error {
	bookmark := &model.BookMark{
		UserId:     userId.String(),
		DocumentId: documentId.String(),
	}

	if err := r.conn.Create(bookmark).Error; err != nil {
		return err
	}

	return nil
}

func (r *DocumentRepository) UnBookMark(userId uuid.UUID, documentId uuid.UUID) error {
	if err := r.conn.Where("user_id = ?", userId).Where("document_id = ?", documentId).Delete(&model.BookMark{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *DocumentRepository) Reference(userId uuid.UUID, documentId uuid.UUID) error {
	reference := &model.Reference{
		UserId:     userId.String(),
		DocumentId: documentId.String(),
	}

	if err := r.conn.Create(reference).Error; err != nil {
		return err
	}

	return nil
}

func (r *DocumentRepository) UnReference(userId uuid.UUID, documentId uuid.UUID) error {
	if err := r.conn.Where("user_id = ?", userId).Where("document_id = ?", documentId).Delete(&model.Reference{}).Error; err != nil {
		return err
	}

	return nil
}
