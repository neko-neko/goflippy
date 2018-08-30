package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neko-neko/goflippy/admin/service"
	"github.com/neko-neko/goflippy/pkg/collection"
	"github.com/neko-neko/goflippy/pkg/util"
)

type ProjectUserHandler struct {
	service *service.UserService
}

func NewProjectUserHandler(service *service.UserService) *ProjectUserHandler {
	return &ProjectUserHandler{
		service: service,
	}
}

type getUsersRequest struct {
	ProjectID string `validate:"required"`
}

type getUsersResponse struct {
	Users []collection.User `json:"users"`
}

// GetUsers returns all users
//
// Responses:
//  200: getUsersResponse
//  400: errorResponse
func (p *ProjectUserHandler) GetUsers(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	param := getUsersRequest{
		ProjectID: vars["id"],
	}

	if err := util.ValidateStruct(&param); err != nil {
		return http.StatusBadRequest, nil, err
	}

	users, err := p.service.FetchUsers(param.ProjectID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, getUsersResponse{
		Users: users,
	}, nil
}
