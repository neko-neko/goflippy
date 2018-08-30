package handler

import (
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/neko-neko/goflippy/admin/service"
	"github.com/neko-neko/goflippy/pkg/collection"
	"github.com/neko-neko/goflippy/pkg/util"
	"gopkg.in/mgo.v2/bson"
)

type ProjectFeatureHandler struct {
	service *service.FeatureService
}

func NewProjectFeatureHandler(service *service.FeatureService) *ProjectFeatureHandler {
	return &ProjectFeatureHandler{
		service: service,
	}
}

type getFeaturesRequest struct {
	ProjectID string `validate:"required,alphanum"`
}

type getFeaturesResponse struct {
	Features []collection.Feature `json:"features"`
}

// GetFeatures returns all features
//
// Responses:
//  200: getFeaturesResponse
//  400: errorResponse
func (p *ProjectFeatureHandler) GetFeatures(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	param := getFeaturesRequest{
		ProjectID: vars["id"],
	}

	if err := util.ValidateStruct(&param); err != nil {
		return http.StatusBadRequest, nil, err
	}

	features, err := p.service.FetchFeatures(param.ProjectID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, getFeaturesResponse{
		Features: features,
	}, nil
}

type getFeatureRequest struct {
	ProjectID string `validate:"required,alphanum"`
	Key       string `validate:"required,alphanum"`
}

type getFeatureResponse struct {
	Feature collection.Feature `json:"feature"`
}

func (p *ProjectFeatureHandler) GetFeature(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	param := getFeatureRequest{
		ProjectID: vars["id"],
		Key:       vars["key"],
	}

	if err := util.ValidateStruct(&param); err != nil {
		return http.StatusBadRequest, nil, err
	}

	feature, err := p.service.FetchFeature(param.Key, param.ProjectID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, getFeatureResponse{
		Feature: feature,
	}, nil
}

type postFearuresRequest struct {
	ProjectID string `validate:"required,alphanum"`
	Key       string `validate:"required,alphanum"`
	Name      string `validate:"required"`
}

type postFeaturesResponse struct {
	Feature collection.Feature `json:"feature"`
}

func (p *ProjectFeatureHandler) PostFeature(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	param := postFearuresRequest{
		ProjectID: vars["id"],
	}

	if err := util.BindJSONParam(r, &param); err != nil {
		return http.StatusBadRequest, nil, err
	}
	if err := util.ValidateStruct(&param); err != nil {
		return http.StatusBadRequest, nil, err
	}

	feature := collection.NewFeature()
	feature.ProjectID = bson.ObjectIdHex(param.ProjectID)
	feature.Key = param.Key
	feature.Name = param.Name

	fmt.Println(feature)

	if err := p.service.RegisterFeature(*feature); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, postFeaturesResponse{
		Feature: *feature,
	}, nil
}
