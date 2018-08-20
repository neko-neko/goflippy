package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neko-neko/goflippy/api/ctx"
	"github.com/neko-neko/goflippy/api/service"
	"github.com/neko-neko/goflippy/pkg/util"
)

// FeatureHandler is features resource handler
type FeatureHandler struct {
	service *service.FeatureService
}

// NewFeatureHandler returns new instance
func NewFeatureHandler(f *service.FeatureService) *FeatureHandler {
	return &FeatureHandler{
		service: f,
	}
}

// getFeatureRequest is GetFeatures request parameter
// swagger:parameters listFeatures
type getFeatureRequest struct {
	// key is feature key
	Key string `validate:"required"`

	// uuids is UUID array
	UUIDs []string `validate:"-"`
}

// getFeatureResponse is GetFeature response
// swagger:parameters getFeatureResponse
type getFeatureResponse struct {
	Enabled bool   `json:"enabled"`
	Key     string `json:"key" bson:"key"`
}

// GetFeature returns the feature
//
// swagger:route GET /features/{key} feature featurer resource
//
// feature resource
//
// Responses:
//  200: getFeatureResponse
//  400: errorResponse
func (f *FeatureHandler) GetFeature(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	param := &getFeatureRequest{
		Key: vars["key"],
	}
	if err := util.BindParam(r, param); err != nil {
		return http.StatusBadRequest, nil, err
	}
	if err := util.ValidateStruct(param); err != nil {
		return http.StatusBadRequest, nil, err
	}

	feature, enabled, err := f.service.FeatureEnabled(param.Key, ctx.GetProjectID(r.Context()))
	if rerr, ok := err.(*service.ResourceNotFoundError); ok {
		return http.StatusNotFound, nil, rerr
	} else if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, getFeatureResponse{
		Enabled: enabled,
		Key:     feature.Key,
	}, nil
}
