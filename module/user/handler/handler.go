package handler

import (
	"context"

	"Oasyce-backend/api/kitex/kitex_gen/user"
)

type Handler struct {
}

func (h *Handler) GetUser(ctx context.Context, req *user.GetUserRequest) (r *user.GetUserResponse, err error) {
	return &user.GetUserResponse{
		User: &user.User{
			Id:               req.GetId(),
			Username:         req.GetUsername(),
			DisplayName:      nil,
			Email:            "",
			TenantId:         "",
			CreatedAt:        "",
			Emails:           nil,
			Type:             "",
			Status:           "",
			AvatarURL:        nil,
			CanPinRepository: false,
			Features:         nil,
			Bio:              nil,
			Location:         nil,
			ThemeMode:        nil,
			Language:         nil,
			Properties:       nil,
		},
	}, nil
}
