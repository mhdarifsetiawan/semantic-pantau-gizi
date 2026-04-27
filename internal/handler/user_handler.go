package handler

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/semantic-digital-nusantara/semantic-pantau-gizi/internal/domain"
	"github.com/semantic-digital-nusantara/semantic-pantau-gizi/internal/service"
	apperrors "github.com/semantic-digital-nusantara/semantic-pantau-gizi/pkg/errors"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// — Input & Output structs —

type GetAllUsersInput struct{}

type GetAllUsersOutput struct {
	Body struct {
		Data  []*domain.User `json:"data"`
		Total int            `json:"total"`
	}
}

type GetUserByIDInput struct {
	ID string `path:"id" doc:"User ID" example:"1"`
}

type GetUserByIDOutput struct {
	Body struct {
		Data *domain.User `json:"data"`
	}
}

// — Handlers —

func (h *UserHandler) GetAllUsers(ctx context.Context, _ *GetAllUsersInput) (*GetAllUsersOutput, error) {
	users, err := h.svc.GetAllUsers(ctx)
	if err != nil {
		return nil, toHumaError(err)
	}

	out := &GetAllUsersOutput{}
	out.Body.Data = users
	out.Body.Total = len(users)
	return out, nil
}

func (h *UserHandler) GetUserByID(ctx context.Context, input *GetUserByIDInput) (*GetUserByIDOutput, error) {
	user, err := h.svc.GetUserByID(ctx, input.ID)
	if err != nil {
		return nil, toHumaError(err)
	}

	out := &GetUserByIDOutput{}
	out.Body.Data = user
	return out, nil
}

// — Register routes —

func RegisterUserRoutes(api huma.API, h *UserHandler) {
	huma.Register(api, huma.Operation{
		OperationID: "get-all-users",
		Method:      http.MethodGet,
		Path:        "/api/v1/users",
		Summary:     "Get all users",
		Tags:        []string{"Users"},
	}, h.GetAllUsers)

	huma.Register(api, huma.Operation{
		OperationID: "get-user-by-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/users/{id}",
		Summary:     "Get user by ID",
		Tags:        []string{"Users"},
	}, h.GetUserByID)
}

func toHumaError(err error) error {
	var appErr *apperrors.AppError
	if e, ok := err.(*apperrors.AppError); ok {
		appErr = e
	} else {
		appErr = apperrors.Internal(err)
	}
	return huma.NewError(appErr.HTTPStatus(), appErr.Message)
}
