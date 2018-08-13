package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neko-neko/goflippy/api/ctx"
	"github.com/neko-neko/goflippy/api/service"
	"github.com/neko-neko/goflippy/pkg/collection"
	"github.com/neko-neko/goflippy/pkg/util"
)

// UserFeatureHandler is users resource handler
type UserFeatureHandler struct {
	service *service.FeatureService
}

// NewUserFeatureHandler returns new instance
func NewUserFeatureHandler(f *service.FeatureService) *UserFeatureHandler {
	return &UserFeatureHandler{
		service: f,
	}
}

// getUsersFeaturesRequest is GetFeatures request parameter
// swagger:parameters
type getUsersFeaturesRequest struct {
	UUID string `validate:"required"`
	Key  string `validate:"required"`
}

// postUsersResponse is PostUsers response
// swagger:parameters postUsersResponse
type getUsersFeaturesResponse struct {
	Enabled bool               `json:"enabled"`
	Feature collection.Feature `json:"feature"`
}

// GetFeatures get features associated the user
//
// swagger:route GET /users/{uuid}/features/{key} get features
//
// Get features associated the user
//
// Responses:
//  200: userResponse
//  400: errorResponse
func (u *UserFeatureHandler) GetFeatures(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	param := getUsersFeaturesRequest{
		UUID: vars["uuid"],
		Key:  vars["key"],
	}
	if err := util.ValidateStruct(&param); err != nil {
		return http.StatusBadRequest, nil, err
	}

	feature, enabled, err := u.service.FeatureEnabled(param.UUID, param.Key, ctx.GetProjectID(r.Context()))
	if nerr, ok := err.(*service.ResourceNotFoundError); ok {
		return http.StatusNotFound, nil, nerr
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, getUsersFeaturesResponse{
		Enabled: enabled,
		Feature: feature,
	}, nil
}
