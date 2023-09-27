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
	var tagReqsMap map[string][]uuid.UUID
	var tagRequests []*model.TagRequest
	var tags []*model.Tag

	if err := r.conn.Find(&requests).Error; err != nil {
		return nil, err
	}

	if err := r.conn.Find(&tagRequests).Error; err != nil {
		return nil, err
	}

	tagReqsMap = make(map[string][]uuid.UUID)
	tagIds := make([]string, len(tagRequests))
	for _, tagRequest := range tagRequests {
		tagId, err := uuid.Parse(tagRequest.TagId)
		if err != nil {
			return nil, err
		}
		tagReqsMap[tagRequest.RequestId] = append(tagReqsMap[tagRequest.RequestId], tagId)
		tagIds = append(tagIds, tagRequest.TagId)
	}

	if err := r.conn.Where("id IN ?", tagIds).Find(&tags).Error; err != nil {
		return nil, err
	}

	tagMap := make(map[string]string)
	for _, tag := range tags {
		tagMap[tag.Id] = tag.Name
	}

	result := make([]*domain.Request, 0)
	for _, request := range requests {
		req, err := request.ToDomain(tagReqsMap[request.Id], tagMap)
		if err != nil {
			return nil, err
		}
		result = append(result, req)
	}

	return result, nil
}

func (r *RequestRepository) CreateRequest(request *domain.Request) (*domain.Request, error) {
	var tags []*model.Tag
	
	requestModel, err := model.RequestToModel(request)
	if err != nil {
		return nil, err
	}

	if err := r.conn.Create(requestModel).Error; err != nil {
		return nil, err
	}

	tagIDs := make([]string, 0)
	for _, tag := range request.Tags {
		tagRequest := &model.TagRequest{
			RequestId: request.Id.String(),
			TagId:     tag.String(),
		}
		if err := r.conn.Create(tagRequest).Error; err != nil {
			return nil, err
		}
		tagIDs = append(tagIDs, tag.String())
	}

	if err := r.conn.Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
		return nil, err
	}

	tagNames := make([]string, len(tags))
	for i, tag := range tags {
		tagNames[i] = tag.Name
	}

	request.TagNames = tagNames

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
	var users []*model.User
	if err := r.conn.Where("user_id = ?", userId).Find(&requests).Error; err != nil {
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

	userIds := make([]string, len(documents))
	for i, document := range documents {
		userIds[i] = document.UserId
	}

	if err := r.conn.Where("id IN ?", userIds).Find(&users).Error; err != nil {
		return nil, err
	}

	return model.RequestsToDomain(requests, requestDocuments, documents, users)
}