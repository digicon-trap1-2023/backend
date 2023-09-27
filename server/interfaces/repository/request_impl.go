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

func (r *RequestRepository) CreateRequest(request *domain.Request) (*domain.Request, error) {
	requestModel, err := model.RequestToModel(request)
	if err != nil {
		return nil, err
	}

	if err := r.conn.Create(requestModel).Error; err != nil {
		return nil, err
	}

	for _, tag := range request.Tags {
		tagRequest := &model.TagRequest{
			RequestId: request.Id.String(),
			TagId:     tag.String(),
		}
		if err := r.conn.Create(tagRequest).Error; err != nil {
			return nil, err
		}
	}

	return request, nil
}

func (r *RequestRepository) DeleteRequest(userId uuid.UUID, requestId uuid.UUID) error {
	var request *model.Request
	if err := r.conn.First(&request, "id = ? AND user_id = ?", requestId, userId).Error; err != nil {
		return err
	}

	if err := r.conn.Delete(&model.Request{}, requestId).Error; err != nil {
		return err
	}

	return nil
}

func (r *RequestRepository) GetRequestsWithDocument(userId uuid.UUID) ([]*domain.Request, error) {
	var requests []*model.Request
	var requestDocuments []*model.RequestDocument
	var documents []*model.Document
	if err := r.conn.Where("user_id = ?").Find(&requests).Error; err != nil {
		return nil, err
	}

	reqIds := make([]string, len(requests))
	for i, request := range requests {
		reqIds[i] = request.Id
	}

	if err := r.conn.Where("request_id IN ?", reqIds).Find(&requestDocuments).Error; err != nil {
		return nil, err
	}

	docIds := make([]string, len(requestDocuments))
	for i, requestDocument := range requestDocuments {
		docIds[i] = requestDocument.DocumentId
	}

	if err := r.conn.Where("id IN ?", docIds).Find(&documents).Error; err != nil {
		return nil, err
	}

	return model.RequestsToDomain(requests, requestDocuments, documents)
}