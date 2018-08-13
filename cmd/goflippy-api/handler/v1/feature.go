package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neko-neko/goflippy/cmd/goflippy-api/ctx"
	"github.com/neko-neko/goflippy/cmd/goflippy-api/service"
	"github.com/neko-neko/goflippy/collection"
	"github.com/neko-neko/goflippy/util"
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

// getFeaturesResponse is GetFeatures response
// swagger:parameters getFeaturesResponse
type getFeaturesResponse struct {
	Features []collection.Feature `json:"features"`
}

// GetFeatures returns all features
//
// swagger:route GET /features features listFeatures
//
// Lists features filtered by some parameters
//
// Responses:
//  200: getFeaturesResponse
//  400: errorResponse
func (f *FeatureHandler) GetFeatures(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	features, err := f.service.FetchFeatures(ctx.GetProjectID(r.Context()))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, getFeaturesResponse{
		Features: features,
	}, nil
}

// getFeatureRequest is GetFeatures request parameter
// swagger:parameters listFeatures
type getFeatureRequest struct {
	Key string `validate:"required"`
}

// getFeatureResponse is GetFeature response
// swagger:parameters getFeatureResponse
type getFeatureResponse struct {
	Feature collection.Feature `json:"feature"`
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

	feature, err := f.service.FetchFeature(param.Key, ctx.GetProjectID(r.Context()))
	if rerr, ok := err.(*service.ResourceNotFoundError); ok {
		return http.StatusNotFound, nil, rerr
	} else if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, getFeatureResponse{
		Feature: feature,
	}, nil
}
