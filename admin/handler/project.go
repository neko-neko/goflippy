package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neko-neko/goflippy/admin/service"
	"github.com/neko-neko/goflippy/pkg/collection"
	"github.com/neko-neko/goflippy/pkg/util"
)

// ProjectHandler is projects resource handler
type ProjectHandler struct {
	service *service.ProjectService
}

// NewProjectHandler returns new instance
func NewProjectHandler(service *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		service: service,
	}
}

// getProjectsResponse is GetProjects response
type getProjectsResponse struct {
	Projects []collection.Project `json:"projects"`
}

// GetProjects returns all projects
//
// Responses:
//  200: getProjectsResponse
//  400: errorResponse
func (p *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	projects, err := p.service.FetchAllProjects()
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, getProjectsResponse{
		Projects: projects,
	}, nil
}

type getProjectRequest struct {
	ID string `json:"id" validate:"required,alphanum"`
}

type getProjectResponse struct {
	Project collection.Project `json:"project"`
}

func (p *ProjectHandler) GetProject(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	param := &getProjectRequest{
		ID: vars["id"],
	}

	if err := util.ValidateStruct(param); err != nil {
		return http.StatusBadRequest, nil, err
	}

	project, err := p.service.FetchProject(param.ID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, getProjectResponse{
		Project: project,
	}, nil
}

type postProjectsRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type postProjectsResponse struct {
	Project collection.Project `json:"project"`
}

func (p *ProjectHandler) PostProjects(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	param := postProjectsRequest{}

	if err := util.BindJSONParam(r, &param); err != nil {
		return http.StatusBadRequest, nil, err
	}
	if err := util.ValidateStruct(&param); err != nil {
		return http.StatusBadRequest, nil, err
	}

	project := collection.NewProject()
	project.Name = param.Name
	project.Description = param.Description

	if err := p.service.RegisterProject(project); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, postProjectsResponse{
		Project: *project,
	}, nil
}

type patchProjectsRequest struct {
	ID          string `json:"project" validate:"required,alphanum"`
	Name        string `json:"name" validate:"omitempty,required"`
	Description string `json:"description" validate:"omitempty,required"`
}

type patchProjectResponse struct {
	Project collection.Project `json:"project"`
}

func (p *ProjectHandler) PatchProjects(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	param := &patchProjectsRequest{
		ID: vars["id"],
	}

	if err := util.BindJSONParam(r, &param); err != nil {
		return http.StatusBadRequest, nil, err
	}
	if err := util.ValidateStruct(param); err != nil {
		return http.StatusBadRequest, nil, err
	}

	project, err := p.service.UpdateProjectInfo(param.ID, param.Name, param.Description)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, patchProjectResponse{
		Project: project,
	}, nil
}
