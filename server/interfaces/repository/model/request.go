package model

import (
	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/google/uuid"
)

type Request struct {
	Id          string `gorm:"type:char(36);not null;primaryKey"`
	Title       string `gorm:"type:varchar(36)"`
	Description string `gorm:"type:varchar(200)"`
	UserId      string `gorm:"type:char(36)"`
}

type TagRequest struct {
	RequestId string `gorm:"type:char(36);not null;primaryKey"`
	TagId     string `gorm:"type:char(36);not null;primaryKey"`
}

type RequestDocument struct {
	RequestId  string `gorm:"type:char(36);not null;primaryKey"`
	DocumentId string `gorm:"type:char(36);not null;primaryKey"`
}

func (request *Request) ToDomain(tags []uuid.UUID, tagsMap map[string]string, userName string) (*domain.Request, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}

	tagNames := make([]string, len(tags))
	for i, tag := range tags {
		tagNames[i] = tagsMap[tag.String()]
	}

	return &domain.Request{
		Id:              id,
		Title:           request.Title,
		Description:     request.Description,
		Tags:            tags,
		TagNames:        tagNames,
		CreatedBy:       userId,
		CreatedUserName: "",
	}, nil
}

func RequestToModel(request *domain.Request) (*Request, error) {
	id := request.Id.String()
	userId := request.CreatedBy.String()
	return &Request{
		Id:          id,
		Title:       request.Title,
		Description: request.Description,
		UserId:      userId,
	}, nil
}

func RequestsToDomain(requests []*Request, requestDocuments []*RequestDocument, documents []*Document, users []*User) ([]*domain.Request, error) {
	requestsDomain := make([]*domain.Request, len(requests))
	requestDocumentsMap := make(map[string][]string)
	for _, requestDocument := range requestDocuments {
		requestDocumentsMap[requestDocument.RequestId] = append(requestDocumentsMap[requestDocument.RequestId], requestDocument.DocumentId)
	}
	documentsMap := make(map[string]*Document)
	for _, document := range documents {
		documentsMap[document.Id] = document
	}
	userNamesMap := make(map[string]string)
	for _, user := range users {
		userNamesMap[user.Id] = user.Name
	}

	for i, request := range requests {
		requestDomain, err := requestToDomain(request, requestDocumentsMap[request.Id], documentsMap, userNamesMap)
		if err != nil {
			return nil, err
		}
		requestsDomain[i] = requestDomain
	}
	return requestsDomain, nil
}

func requestToDomain(request *Request, documentIds []string, documentsMap map[string]*Document, userNamesMap map[string]string) (*domain.Request, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}
	userId, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}
	documents := make([]*domain.Document, len(documentIds))
	for i, documentId := range documentIds {
		documentModel := documentsMap[documentId]
		document, err := documentModel.ToDomain(nil, nil, nil, userNamesMap[documentModel.UserId])
		if err != nil {
			return nil, err
		}
		documents[i] = document
	}
	return &domain.Request{
		Id:               id,
		Title:            request.Title,
		Description:      request.Description,
		Tags:             nil,
		CreatedBy:        userId,
		CreatedUserName:  userNamesMap[request.UserId],
		RelatedDocuments: documents,
	}, nil
}
