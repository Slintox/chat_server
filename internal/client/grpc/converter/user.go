package converter

import (
	"github.com/Slintox/chat_server/internal/model"
	desc "github.com/Slintox/user-service/pkg/user_v1"
)

func ToGetUserRequest(username string) *desc.GetRequest {
	return &desc.GetRequest{
		Username: username,
	}
}

func FromGetUserResponse(r *desc.GetResponse) *model.User {
	return &model.User{
		Username:  r.User.Username,
		Email:     r.User.Email,
		Role:      model.UserRole(r.User.Role),
		CreatedAt: r.User.CreatedAt.AsTime(),
		UpdatedAt: r.User.UpdatedAt.AsTime(),
	}
}
