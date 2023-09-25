package request

import "github.com/digicon-trap1-2023/backend/domain"

type GetMeResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func UserToGetMeResponse(user *domain.User) *GetMeResponse {
	return &GetMeResponse{
		Id: user.Id.String(),
		Name: user.Name,
		Role: user.Role,
	}
}
